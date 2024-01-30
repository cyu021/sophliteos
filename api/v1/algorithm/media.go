package algorithm

import (
	"algoliteos/config"
	"algoliteos/global"
	"algoliteos/logger"
	"algoliteos/mvc"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type MediaApi struct{}

var mediaPullMap = make(map[string]mvc.DetectDeviceInfo)
var mu sync.RWMutex

func init() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("0 0 0 * * ?", func() {
		autoDetect()
	})
	if err != nil {
		fmt.Println("err:", err)
	}

	c.Start()
}

// 配置流媒体服务器地址
// 写入全局变量，再写入文件
func (b *MediaApi) ConfMedia(c *gin.Context) {
	var req mvc.MediaHost
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	global.SystemConf.MediaHost = req.Ip + ":" + strconv.Itoa(req.Port)

	err := config.SaveConfig()
	if err != nil {
		c.JSON(http.StatusOK, mvc.Fail(-1, "配置失败"))
		return
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

// 获取流媒体服务器地址
func (b *MediaApi) GetMedia(c *gin.Context) {
	host, port, err := mvc.ExtractIP(global.SystemConf.MediaHost)
	portInt, _ := strconv.Atoi(port)
	if err != nil {
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体地址错误"))
	}
	mediaHost := mvc.MediaHost{
		Ip:   host,
		Port: portInt,
	}
	c.JSON(http.StatusOK, mvc.Success(mediaHost))
}

// 添加视频设备到流媒体服务
func (b *MediaApi) AddDev(c *gin.Context) {
	var deviceInfo mvc.DeviceInfo
	if err := c.ShouldBindJSON(&deviceInfo); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	if mvc.ContainsSpecialCharacters(deviceInfo.Name) {
		logger.Error("设备名称不合法")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "设备名称不能包含符号，只能为中英文和数字"))
		return
	}

	deviceMap := getMediaDevices()
	for _, value := range deviceMap {
		if deviceInfo.Name == value.DeviceName {
			c.JSON(http.StatusOK, mvc.Fail(-1, "设备已存在"))
			return
		}
	}

	body, _ := json.Marshal(deviceInfo)

	var devices mvc.MediaRes
	res := mvc.NewRequestWithHeaders(global.SystemConf.MediaHost+"/addDevice", "POST", nil, body)
	err := json.Unmarshal(res, &devices)
	if err != nil {
		logger.Error("流媒体设备添加失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备添加失败"))
		return
	}
	if devices.Code != 0 {
		logger.Error("请求流媒体设备信息失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备添加失败"+devices.Msg))
		return
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

// 修改流媒体服务中设备
func (b *MediaApi) ModDev(c *gin.Context) {
	var deviceMod mvc.DeviceMod
	if err := c.ShouldBindJSON(&deviceMod); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	deviceMap := getMediaDevices()
	if _, exists := deviceMap[deviceMod.DeviceId]; !exists {
		c.JSON(http.StatusOK, mvc.Fail(-1, "设备不存在"))
		return
	}

	body, _ := json.Marshal(deviceMod)

	var devices mvc.MediaRes
	res := mvc.NewRequestWithHeaders(global.SystemConf.MediaHost+"/modDevice", "POST", nil, body)
	err := json.Unmarshal(res, &devices)
	if err != nil {
		logger.Error("流媒体设备修改失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备修改失败"))
		return
	}
	if devices.Code != 0 {
		logger.Error("流媒体设备修改失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备修改失败"+devices.Msg))
		return
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

// 删除设备
func (b *MediaApi) DelDev(c *gin.Context) {
	var deviceDel struct {
		Devices []string `json:"device"`
	}
	body, _ := io.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &deviceDel)
	if err != nil {
		logger.Error("JSON解析失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "JSON解析失败"))
		return
	}

	var devices mvc.MediaRes
	res := mvc.NewRequestWithHeaders(global.SystemConf.MediaHost+"/delDevice", "POST", nil, body)
	err = json.Unmarshal(res, &devices)
	if err != nil {
		logger.Error("流媒体设备删除失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备删除失败"))
		return
	}
	if devices.Code != 0 {
		logger.Error("流媒体设备删除失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备删除失败"+devices.Msg))
		return
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

// 获取流媒体设备的实时播放地址
func (b *MediaApi) GetLiveUrl(c *gin.Context) {
	deviceId := c.Query("deviceId")

	body := fmt.Sprintf("{\"deviceId\":\"%s\"}", deviceId)
	res := mvc.NewRequestWithHeaders(global.SystemConf.MediaHost+"/getLiveUrl", "POST", nil, []byte(body))

	var liveUrl mvc.LiveUrl
	err := json.Unmarshal(res, &liveUrl)
	if err != nil {
		logger.Error("解析JSON失败:%v", err)
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "media server error"))
		return
	}

	c.JSON(http.StatusOK, mvc.Success(modMediaHost(liveUrl.TsUrl)))
}

// 流媒体视频播放地址host必须是本机ip，替换流媒体地址,优先使用流媒体配置地址,其次
func modMediaHost(orginUrl string) string {
	// 解析原始 URL
	parsedURL, err := url.Parse(orginUrl)
	if err != nil {
		logger.Error("Error parsing URL:%v", err)
		return ""
	}

	// 优先使用流媒体返回地址，如果流媒体配置是127.0.0.1，则更换视频播放地址
	if !strings.Contains(orginUrl, "127.0.0.1") {
		return orginUrl
	}

	var deviceIp string
	// 使用流媒体配置地址
	if strings.Contains(global.SystemConf.MediaHost, "127.0.0.1") {
		// 如果配置地址也是127.1，则获取设备地址
		deviceIp = getLoaclHost()
		if deviceIp == "" {
			deviceIp = global.SystemConf.LocalHost
		}
		parsedURL.Host = fmt.Sprintf("%s:%s", deviceIp, parsedURL.Port())
	} else {
		deviceIp = global.SystemConf.MediaHost
		parsedURL.Host = deviceIp
	}

	// 重新构建 URL 字符串
	return parsedURL.String()
}

// 通过sophliteos获取本机地址
func getLoaclHost() string {
	res := mvc.NewRequestWithHeaders("127.0.0.1:8080/api/localhost", "GET", nil, nil)

	var ip mvc.DeviceIP
	err := json.Unmarshal(res, &ip)
	if err != nil {
		logger.Error("解析JSON失败:%v", err)
		return ""
	}
	return ip.Result
}

// 查询流媒体的所有设备
func (b *MediaApi) GetDevices(c *gin.Context) {
	var page struct {
		PageNo   int `json:"pageNo"`
		PageSize int `json:"pageSize"`
	}
	reqBody, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(reqBody, &page)

	deviceMap := getMediaDevices()

	if deviceMap == nil {
		logger.Error("请求流媒体设备信息失败")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "请求流媒体设备信息失败"))
		return
	}

	devices := mvc.DevicesList{
		TotalSize: len(deviceMap),
		PageCount: 0, //int(math.Ceil(float64(len(deviceMap)) / float64(page.PageSize)))
		PageNo:    page.PageNo,
		PageSize:  page.PageSize,
	}

	keys := make([]string, 0, len(deviceMap))
	for key := range deviceMap {
		keys = append(keys, key)
	}
	// 对键进行排序
	sort.Strings(keys)

	for _, key := range keys {
		value := deviceMap[key]
		value.MediaPull = mediaPullMap[key].MediaPull
		devices.Devices = append(devices.Devices, value)
	}

	c.JSON(http.StatusOK, mvc.Success(devices))
}

// 接受巡检上报
func (b *MediaApi) DeviceDetectRev(c *gin.Context) {
	var detectInfo mvc.Detect
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &detectInfo)

	logger.Info("revice detect result:%v", detectInfo)
	mu.Lock()
	defer mu.Unlock()

	for _, item := range detectInfo.List {
		mediaPullMap[item.DeviceID] = item
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

// 主动巡检设备
func (b *MediaApi) DetectDev(c *gin.Context) {
	var devs struct {
		Devs []string `json:"deviceIds"`
	}
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &devs)
	logger.Info("巡检设备:%v", devs.Devs)
	detectDevice(devs.Devs)

	c.JSON(http.StatusOK, mvc.Ok())
}

// 自动巡检设备
func detectDevice(deviceId []string) bool {
	var deviceIds []mvc.DeviceId

	for _, dev := range deviceId {
		devID := mvc.DeviceId{
			DeviceID: dev,
		}
		deviceIds = append(deviceIds, devID)
	}

	logger.Info("上传地址:%s", global.SystemConf.UploadHost)

	body := mvc.DeviceDetectBody{
		DetectReportURL: "http://" + global.SystemConf.UploadHost + "/algorithm/media/detect",
		DevList:         deviceIds,
	}
	bodyByte, _ := json.Marshal(body)

	var req mvc.MediaReq
	res := mvc.NewRequestWithHeaders(global.SystemConf.MediaHost+"/manualDeviceDetect", "POST", nil, bodyByte)
	json.Unmarshal(res, &req)

	return req.Status == 0
}

// 获取流媒体设备信息
func getMediaDevices() map[string]mvc.Device {
	var devices mvc.DevicesList
	res := mvc.NewRequestWithHeaders(global.SystemConf.MediaHost+"/getDevListEx", "POST", nil, []byte("{}"))
	err := json.Unmarshal(res, &devices)
	if err != nil {
		logger.Error("请求流媒体设备信息失败:%v", err)
		return nil
	}
	if devices.Code != 0 {
		logger.Error("请求流媒体设备信息失败:%v", err)
		return nil
	}

	deviceMpa := make(map[string]mvc.Device)

	for i := range devices.Devices {
		devices.Devices[i].MediaServer = global.SystemConf.MediaHost

		device := devices.Devices[i]
		if device.Type != "camera" {
			continue
		}

		deviceMpa[device.DeviceId] = device
	}

	return deviceMpa
}

// 自动巡检
func autoDetect() {
	deviceMpa := getMediaDevices()

	var result []string
	for key := range deviceMpa {
		result = append(result, key)
	}
	logger.Info("巡检设备:%v", result)
	detectDevice(result)
}
