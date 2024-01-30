package algorithm

import (
	"algoliteos/config"
	"algoliteos/database"
	"algoliteos/global"
	"algoliteos/logger"
	"algoliteos/mvc"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type QueryApi struct{}

func init() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("0 0 0,12 * * ?", func() {
		clearDisk()
	})
	if err != nil {
		fmt.Println("err:", err)
	}

	c.Start()
}

func (b *QueryApi) GetRecord(c *gin.Context) {
	var algoQuery mvc.AlgoQuery
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &algoQuery)

	records, total := queryRecord(algoQuery)
	var items []mvc.AlarmReqItem

	for _, record := range records {

		smallImagePath := fmt.Sprintf("/algorithm/image?taskId=%s&type=%d&fileName=%s", record.TaskId, record.Type, record.SamllPictureFilename)
		bigImagePath := fmt.Sprintf("/algorithm/image?taskId=%s&fileName=%s", record.TaskId, record.BigPictureFilename)
		item := mvc.AlarmReqItem{
			Id:         int(record.ID),
			TaskId:     record.TaskId,
			SrcID:      record.SrcID,
			FrameIndex: record.FrameIndex,
			Type:       record.Type,
			BigImage:   bigImagePath,
			SmallImage: smallImagePath,
			Time:       record.Date,
			Box: mvc.Box{
				LeftTopY:  record.LeftTopY,
				LeftTopX:  record.LeftTopX,
				RightBtmY: record.RightBtmY,
				RightBtmX: record.RightBtmX,
			},
			Extend: record.Extend,
		}
		items = append(items, item)
	}

	alarmReq := mvc.AlarmReq{
		Total:     total,
		PageSize:  algoQuery.PageSize,
		PageNo:    algoQuery.PageNo,
		PageCount: total/algoQuery.PageSize + 1,
		Items:     items,
	}
	usize, msize := getSize()
	alarmReq.UsedSize = formatSize(usize)
	alarmReq.MaxSize = strconv.Itoa(int(msize) / 1024 / 1024)
	c.JSON(http.StatusOK, mvc.Success(alarmReq))

}

func (b *QueryApi) GetQueryInfo(c *gin.Context) {
	var taskIds, srcIds []string
	var types []int
	database.DB.Model(&mvc.Record{}).Pluck("distinct(task_id)", &taskIds)
	database.DB.Model(&mvc.Record{}).Pluck("distinct(src_id)", &srcIds)
	database.DB.Model(&mvc.Record{}).Pluck("distinct(type)", &types)

	queryList := mvc.QueryInfo{
		Types:   types,
		TaskIds: taskIds,
		SrcIds:  srcIds,
	}

	c.JSON(http.StatusOK, mvc.Success(queryList))

}

func (b *QueryApi) ModSize(c *gin.Context) {
	var maxSize struct {
		MaxSize int64 `json:"maxSize"`
	}
	if err := c.ShouldBindJSON(&maxSize); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	global.SystemConf.Picture.MaxSize = maxSize.MaxSize

	err := config.SaveConfig()
	if err != nil {
		c.JSON(http.StatusOK, mvc.Fail(-1, "set error"))
		return
	}
	go clearDisk()
	c.JSON(http.StatusOK, mvc.Ok())
}

func queryRecord(algoQuery mvc.AlgoQuery) ([]mvc.Record, int) {
	db := database.DB.Model(&mvc.Record{})
	var records []mvc.Record
	// 构建基本的查询语句
	query := ""
	// 构建条件语句
	var vars []interface{}

	// 添加 BeginTime 和 EndTime 条件
	if algoQuery.BeginTime != 0 {
		query = query + " and date >= ?  "
		vars = append(vars, &algoQuery.BeginTime)
	}
	if algoQuery.EndTime != 0 {
		query = query + " and date <= ?  "
		vars = append(vars, &algoQuery.EndTime)
	}

	// 添加 task_id 条件
	if algoQuery.TaskID != "" {
		query = query + " and task_id = ? "
		vars = append(vars, algoQuery.TaskID)
	}

	// 添加 task_id 条件
	if algoQuery.SrcID != "" {
		query = query + " and src_id = ? "
		vars = append(vars, algoQuery.SrcID)
	}

	// 添加 Alarms 条件
	if len(algoQuery.Types) > 0 {
		alarmConditions := []string{}
		for _, alarm := range algoQuery.Types {
			alarmConditions = append(alarmConditions, "type = ?")
			vars = append(vars, alarm)
		}
		query = query + " and (" + strings.Join(alarmConditions, " or ") + ")"
	}

	var total int
	if len(query) > 0 {
		db.Where(query[4:], vars...).Count(&total)
		db.Where(query[4:], vars...).Offset((algoQuery.PageNo - 1) * algoQuery.PageSize).Limit(algoQuery.PageSize).Order("date desc").Find(&records)
	} else {
		db.Count(&total)
		db.Where(query, vars...).Offset((algoQuery.PageNo - 1) * algoQuery.PageSize).Limit(algoQuery.PageSize).Order("date desc").Find(&records)
	}
	return records, total
}

func getSize() (int64, int64) {
	size, err := calculateSize(global.SystemConf.Picture.Dir)
	if err != nil {
		logger.Error("Unable to parse du output")
		return 0, 0
	}

	return size, global.SystemConf.Picture.MaxSize * 1024 * 1024
}

func formatSize(size int64) string {
	const (
		B  = 1
		KB = 1024 * B
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
	)

	switch {
	case size >= TB:
		return fmt.Sprintf("%.2f TB", float64(size)/float64(TB))
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func calculateSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	return size, err
}

func clearDisk() {
	num := 0
	for i := 0; i < 800; i++ {
		usedSize, maxSize := getSize()
		if usedSize > maxSize {
			removeAlarmRecode(1)
			num++
		} else {
			break
		}
	}
	logger.Info("删除告警记录，数量：%d\n", num)
}

func removeAlarmRecode(num int64) {
	db := database.DB.Model(&mvc.Record{})
	var records []mvc.Record

	// 查询最旧的num条记录
	db.Order("date").Limit(num).Find(&records)

	// 删除记录
	for _, record := range records {
		filePath := global.SystemConf.Picture.Dir + "/" + record.TaskId + "/" + strconv.Itoa(record.Type) + "/" + record.SamllPictureFilename
		err := os.Remove(filePath)
		if err != nil {
			logger.Error("无法删除文件：%v\n", err)
		}
		// 删除记录
		if err := db.Delete(&record).Error; err != nil {
			logger.Error("无法删除记录：%v\n", err)
		}
		if isNotInRecodes(record.BigPictureFilename) {
			filePath := global.SystemConf.Picture.Dir + "/" + record.TaskId + "/" + record.BigPictureFilename
			err := os.Remove(filePath)
			if err != nil {
				logger.Error("无法删除文件：%v\n", err)
			}
		}

	}
}

func isNotInRecodes(name string) bool {
	var record mvc.Record
	notFound := database.DB.Where("big_picture_filename = ?", name).First(&record).RecordNotFound()
	return notFound
}
