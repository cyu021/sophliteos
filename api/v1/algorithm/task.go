package algorithm

import (
	"algoliteos/database"
	"algoliteos/global"
	"algoliteos/logger"
	"algoliteos/mvc"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 算法应用api
const (
	taskFind   = "/task/list"
	taskCancle = "/task/delete"
	taskSetup  = "/task/create"
	taskQuery  = "/task/query"
)

var HEADER = map[string]string{"content-type": "application/json;charset=UTF-8"}

type TaskApi struct{}

// 给前端返回算法能力列表，示例如下
//
//	"1": "吸烟检测",
//	"2": "机动车违停",
//	"3": "未戴口罩",
//	"4": "非机动车乱停",
//	"5": "火焰监测",
//	"6": "突发性事件",
//	"7": "占道经营",
//	"8": "电动车检测",
//	"9": "离岗检测",
//	"10": "施工占道"
func (b *TaskApi) AbilitiesList(c *gin.Context) {
	var res []mvc.AbilityList
	for key, value := range global.SystemConf.Abilities {
		res = append(res, mvc.AbilityList{
			Type: key,
			Name: value,
		})
	}

	c.JSON(http.StatusOK, mvc.Success(res))
}

// 添加算法任务
func (b *TaskApi) AddTask(c *gin.Context) {
	algoTask := mvc.AlgoTask{}
	if err := c.ShouldBindJSON(&algoTask); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}
	// 参数判断
	if algoTask.DeviceId == "" || algoTask.DeviceName == "" || algoTask.Url == "" || algoTask.TaskId == "" {
		logger.Error("添加任务参数错误")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误，不能为空"))
		return
	}

	if len(algoTask.Abilities) == 0 {
		logger.Error("算法能力为空")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "算法能力为空"))
		return
	}

	if mvc.ContainsSpecialCharacters(algoTask.TaskId) {
		logger.Error("任务名称不合法")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务名称不能包含符号，只能为中英文和数字"))
		return
	}

	// 通过 TaskId 查找数据
	var task mvc.AlgoTaskSql
	result := database.DB.Where("task_id = ?", algoTask.TaskId).First(&task)

	if result.Error != gorm.ErrRecordNotFound {
		logger.Error("任务已存在")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务已存在"))
		return
	}

	// 获取视频通道信息，检测区域默认设置为全屏
	deviceMap := getMediaDevices()
	if deviceMap == nil {
		logger.Error("流媒体服务异常")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "流媒体服务异常"))
		return
	}
	deviceInfo := deviceMap[algoTask.DeviceId]
	if deviceInfo.Resolution == "unknown" {
		logger.Error("视频未巡检")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "视频未巡检或不在线"))
		return
	}

	logger.Info("%s", mvc.StructPrint(deviceInfo))
	xMax, yMax, err := mvc.ExtractResolution(deviceInfo.Resolution)
	if err != nil {
		logger.Error("deviceInfo.Resolution:%s", deviceInfo.Resolution)
		logger.Error("视频分辨率获取失败")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "视频分辨率获取失败"))
		return
	}
	detectInfos := []mvc.DetectInfo{
		{
			HotArea: []mvc.Point2D{
				{
					X: 0,
					Y: 0,
				},
				{
					X: xMax,
					Y: 0,
				},
				{
					X: xMax,
					Y: yMax,
				},
				{
					X: 0,
					Y: yMax,
				},
			},
		},
	}

	var algorithms []mvc.Algorithm
	for _, algoType := range algoTask.Abilities {
		algorithm := mvc.Algorithm{
			Type:           algoType,
			TrackInterval:  1,
			DetectInterval: 200,
			TargetSize: mvc.TargetSize{
				MinDetect: 30,
				MaxDetect: 250,
			},
			DetectInfos: detectInfos,
		}
		algorithms = append(algorithms, algorithm)
	}

	str, _ := mvc.StructToString(algorithms)

	algoTaskSql := mvc.AlgoTaskSql{
		TaskId:      algoTask.TaskId,
		Status:      0,
		DeviceName:  algoTask.DeviceName,
		DeviceId:    algoTask.DeviceId,
		Url:         algoTask.Url,
		VideoWidth:  xMax,
		VideoHeight: yMax,
		Abilities:   str,
	}

	// 保存到数据库
	_ = database.DB.Create(&algoTaskSql)

	logger.Info("任务创建成功\n%s", mvc.StructPrint(algoTaskSql))

	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *TaskApi) ModTask(c *gin.Context) {
	algoTask := mvc.AlgoTaskMod{}
	if err := c.ShouldBindJSON(&algoTask); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	// 通过 TaskId 查找数据
	var task mvc.AlgoTaskSql
	db := database.DB.Where("task_id = ?", algoTask.TaskId).First(&task)

	if db.Error == gorm.ErrRecordNotFound {
		logger.Error("任务不存在")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务不存在"))
		return
	}

	if task.DeviceId != algoTask.DeviceId && algoTask.DeviceId != "" {
		deviceMpa := getMediaDevices()
		device := deviceMpa[algoTask.DeviceId]
		task.DeviceId = device.DeviceId
		task.DeviceName = device.DeviceName
		task.Url = device.Url

		var err error
		task.VideoWidth, task.VideoHeight, err = mvc.ExtractResolution(device.Resolution)
		if err != nil {
			logger.Error("deviceInfo.Resolution:%s", device.Resolution)
			logger.Error("视频分辨率获取失败")
			c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "视频分辨率获取失败"))
			return
		}
	}

	var abilities []mvc.Algorithm
	if err := json.Unmarshal([]byte(task.Abilities), &abilities); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}

	// 获取视频通道信息，检测区域默认设置为全屏
	deviceMap := getMediaDevices()
	if deviceMap == nil {
		logger.Error("流媒体服务异常")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "流媒体服务异常"))
		return
	}
	deviceInfo := deviceMap[algoTask.DeviceId]
	if deviceInfo.Resolution == "unknown" {
		logger.Error("视频未巡检")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "视频未巡检或不在线"))
		return
	}

	detectInfos := []mvc.DetectInfo{
		{
			HotArea: []mvc.Point2D{
				{
					X: 0,
					Y: 0,
				},
				{
					X: task.VideoWidth,
					Y: 0,
				},
				{
					X: task.VideoWidth,
					Y: task.VideoHeight,
				},
				{
					X: 0,
					Y: task.VideoHeight,
				},
			},
		},
	}

	var algorithms []mvc.Algorithm
	for _, algoType := range algoTask.Abilities {
		if exist, item := hasAlgorithmWithType(abilities, algoType); exist == true {
			algorithms = append(algorithms, item)
			continue
		}

		algorithm := mvc.Algorithm{
			Type:           algoType,
			TrackInterval:  1,
			DetectInterval: 5,
			TargetSize: mvc.TargetSize{
				MinDetect: 30,
				MaxDetect: 250,
			},
			DetectInfos: detectInfos,
		}
		algorithms = append(algorithms, algorithm)
	}

	str, _ := mvc.StructToString(algorithms)

	task.Abilities = str
	task.Status = 0

	// 保存到数据库
	_ = database.DB.Save(&task)
	logger.Info("任务修改成功%s", mvc.StructPrint(task))

	stopAlgoTask(task.DeviceName)

	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *TaskApi) DeleteTask(c *gin.Context) {
	var req mvc.TaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	// 通过 TaskName 查找数据
	var task mvc.AlgoTaskSql
	db := database.DB.Where("task_id = ?", req.TaskId).First(&task)

	if task.Status == 1 {
		stopAlgoTask(task.DeviceName)
	}

	// 删除找到的数据
	db = database.DB.Delete(&task)
	if db.Error != nil {
		logger.Error("删除数据时出错:%v", db.Error)
		return
	}
	logger.Info("任务删除成功%v", task)

	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *TaskApi) StartTask(c *gin.Context) {
	var req mvc.TaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	// 通过 TaskId 查找任务
	var task mvc.AlgoTaskSql
	db := database.DB.Where("task_id = ?", req.TaskId).First(&task)
	if db.Error == gorm.ErrRecordNotFound {
		logger.Error("任务不存在")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务不存在"))
		return
	}

	if task.Status != 0 {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务正在运行"))
		return
	}

	// if mediaPullMap[task.DeviceId].MediaPull == 0 {
	// 	logger.Error("视频源未巡检")
	// 	c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "视频源未巡检或拉流失败"))
	// 	return
	// }

	var algorithmAbilites []mvc.Algorithm
	var err error
	if algorithmAbilites, err = mvc.StrToAbilities(task.Abilities, task.VideoHeight, task.VideoWidth); err != nil {
		logger.Error("算法能力配置错误")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "算法能力配置错误"))
		return
	}
	taskInfo := mvc.Task{
		TaskId: task.TaskId,
		InputSrc: mvc.InputSource{
			SrcID: task.DeviceName,
			StreamSrc: mvc.StreamSource{
				Address: task.Url,
			},
		},
		Algorithm: algorithmAbilites,
		Reporting: mvc.Reporting{
			ReportUrlList: []string{"http://" + global.SystemConf.UploadHost + "/algorithm/upload"},
		},
	}

	requestBody, _ := mvc.StructToString(taskInfo)
	logger.Info("启动任务:%s", requestBody)

	rec := mvc.AlgoReq{
		Code: -1,
	}
	data := mvc.NewRequestWithHeaders(global.SystemConf.AlgorithmHost+taskSetup, "POST", HEADER, []byte(requestBody))
	json.Unmarshal(data, &rec)

	if rec.Code != 0 {
		if rec.Msg == "" {
			c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务启动失败"))
			return
		}
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, rec.Msg))
		return
	}
	task.Status = 1
	_ = database.DB.Save(&task)
	c.JSON(http.StatusOK, mvc.Ok())

}

func (b *TaskApi) StopTask(c *gin.Context) {
	var req mvc.TaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	var task mvc.AlgoTaskSql
	db := database.DB.Where("task_id = ?", req.TaskId).First(&task)
	if db.Error == gorm.ErrRecordNotFound {
		logger.Error("任务不存在")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务不存在"))
		return
	}

	if !stopAlgoTask(task.TaskId) {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "停止任务失败"))
		return
	}

	task.Status = 0
	_ = database.DB.Save(&task)

	c.JSON(http.StatusOK, mvc.Ok())
}

func stopAlgoTask(taskId string) bool {

	body := mvc.TaskRes{
		TaskId: taskId,
	}
	data, _ := json.Marshal(body)

	logger.Info("停止任务：%s", string(data))

	req := mvc.AlgoReq{
		Code: -1,
	}
	res := mvc.NewRequestWithHeaders(global.SystemConf.AlgorithmHost+taskCancle, "POST", HEADER, data)
	json.Unmarshal(res, &req)

	if req.Code != 0 {
		return false
	}
	return true
}

func (b *TaskApi) List(c *gin.Context) {
	var page struct {
		PageNo   int `json:"pageNo"`
		PageSize int `json:"pageSize"`
	}
	reqBody, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(reqBody, &page)

	var algoTasks []mvc.AlgoTaskSql
	database.DB.Find(&algoTasks)

	var getTaskLists mvc.TaskList

	res := mvc.NewRequestWithHeaders(global.SystemConf.AlgorithmHost+taskFind, "POST", HEADER, []byte("{}"))
	err := json.Unmarshal(res, &getTaskLists)
	if err != nil {
		logger.Error("向算法服务请求  %s  失败", global.SystemConf.AlgorithmHost+taskFind)
	}

	var taskList mvc.TaskListRes
	var items []mvc.TaskItem
	taskList.Total = len(algoTasks)
	taskList.PageCount = taskList.Total / page.PageSize
	taskList.PageSize = page.PageSize

	// logger.Info("%s", mvc.StructPrint(getTaskLists))
	for i := range algoTasks {

		status := containsTask(getTaskLists.Result, algoTasks[i].TaskId)
		// logger.Info("%s %d", algoTasks[i].TaskId, status)

		if algoTasks[i].Status != status {
			algoTasks[i].Status = status
			database.DB.Save(&algoTasks[i])
		}

		var algorithms []mvc.Algorithm
		mvc.ParseJSONString(algoTasks[i].Abilities, &algorithms)
		var abilities []string
		var types []int

		for _, algorithm := range algorithms {
			types = append(types, algorithm.Type)
			abilities = append(abilities, global.SystemConf.Abilities[algorithm.Type])
		}
		item := mvc.TaskItem{
			TaskId:      algoTasks[i].TaskId,
			DeviceName:  algoTasks[i].DeviceName,
			DeviceId:    algoTasks[i].DeviceId,
			Status:      algoTasks[i].Status,
			ErrorReason: "",
			Abilities:   abilities,
			Types:       types,
		}
		items = append(items, item)

	}
	taskList.Items = items
	c.JSON(http.StatusOK, mvc.Success(taskList))
}

func containsTask(taskInfos []mvc.TaskInfo, taskId string) int {
	for _, item := range taskInfos {
		if item.TaskId == taskId {
			return item.Status
		}
	}
	return 0
}

func (b *TaskApi) GetTaskConfig(c *gin.Context) {
	var req mvc.TaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	var task mvc.AlgoTaskSql
	db := database.DB.Where("task_id = ?", req.TaskId).First(&task)

	if db.Error == gorm.ErrRecordNotFound {
		logger.Error("任务不存在")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务不存在"))
		return
	}

	var algorithms []mvc.Algorithm
	mvc.ParseJSONString(task.Abilities, &algorithms)

	deviceMap := getMediaDevices()

	res := mvc.AlgorithmConf{
		Device:     deviceMap[task.DeviceId],
		Algorithms: algorithms,
	}

	c.JSON(http.StatusOK, mvc.Success(res))
}

// 修改算法任务参数配置
func (b *TaskApi) ModTaskConfig(c *gin.Context) {
	var req mvc.TaskConf
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	var task mvc.AlgoTaskSql
	db := database.DB.Where("task_id = ?", req.TaskId).First(&task)

	if db.Error == gorm.ErrRecordNotFound {
		logger.Error("任务不存在")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务不存在"))
		return
	}

	var algorithms []mvc.Algorithm
	mvc.ParseJSONString(task.Abilities, &algorithms)

	for i, algorithm := range algorithms {
		if algorithm.Type == req.Algorithm.Type {
			// algorithms[i] = req.Algorithm
			CopyAlgorithmValues(&algorithms[i], &req.Algorithm)
		}
	}

	str, _ := mvc.StructToString(algorithms)
	task.Abilities = str

	// 保存到数据库
	_ = database.DB.Save(&task)
	logger.Info("任务参数修改成功%s", mvc.StructPrint(task))

	c.JSON(http.StatusOK, mvc.Ok())
}

func CopyAlgorithmValues(dst, src *mvc.Algorithm) {

	dst.TrackInterval = src.TrackInterval
	dst.TargetSize = src.TargetSize
	dst.DetectInterval = src.DetectInterval

	if len(src.DetectInfos) > 0 {
		dst.DetectInfos = src.DetectInfos
	}
	if src.Extend != nil {
		dst.Extend = src.Extend
	}
}

// 判断是否存在指定 Type 的 Algorithm
func hasAlgorithmWithType(algorithms []mvc.Algorithm, targetType int) (bool, mvc.Algorithm) {
	for _, algo := range algorithms {
		if algo.Type == targetType {
			return true, algo
		}
	}
	return false, mvc.Algorithm{}
}
