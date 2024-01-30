package algorithm

import (
	"algoliteos/database"
	"algoliteos/global"
	"algoliteos/logger"
	"algoliteos/mvc"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ReceiveAlarmApi struct{}

var Rand *rand.Rand

func init() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	Rand = rand.New(source)
}

// 调试用
func (b *ReceiveAlarmApi) AlarmRev1(c *gin.Context) {
	// reqBody, _ := io.ReadAll(c.Request.Body)
	// logger.Info("接受告警:%s", string(reqBody))

	var alarmDate mvc.AlarmDate
	if err := c.ShouldBindJSON(&alarmDate); err != nil {
		logger.Error("告警接收失败，参数错误:%v", err)
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	logger.Info("接受告警，%s", alarmDate.TaskID)

	dir := global.SystemConf.Picture.Dir + "/" + alarmDate.TaskID
	if !mvc.FileIsExisted(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}

	now := time.Now()
	bigPicName := fmt.Sprintf("%c%c%d", Rand.Intn(26)+'A', Rand.Intn(26)+'A', now.Unix()) + ".jpg"
	if err := jpegSave(&alarmDate.SceneImageBase64, dir+"/"+bigPicName); err != nil {
		logger.Error("图片保存错误:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "picture save error"))
		return
	}

	for _, event := range alarmDate.AnalyzeEvents {
		eventDir := global.SystemConf.Picture.Dir + "/" + alarmDate.TaskID + "/" + strconv.Itoa(event.Type)
		if !mvc.FileIsExisted(eventDir) {
			os.MkdirAll(eventDir, os.ModePerm)
		}
		smallPicName := fmt.Sprintf("%c%c%d", Rand.Intn(26)+'A', Rand.Intn(26)+'A', now.Unix()) + ".jpg"
		if err := jpegSave(&event.ImageBase64, eventDir+"/"+smallPicName); err != nil {
			logger.Error("图片保存错误:%v", err)
			continue
		}

		jsonData, err := json.Marshal(event.Extend)
		if err != nil {
			logger.Error("解析Extend错误")
		}

		record := mvc.Record{
			TaskId:             alarmDate.TaskID,
			SrcID:              alarmDate.SrcID,
			FrameIndex:         alarmDate.FrameIndex,
			BigPictureFilename: bigPicName,

			Type:                 event.Type,
			Date:                 now.Unix(),
			SamllPictureFilename: smallPicName,
			LeftTopY:             event.Box.LeftTopY,
			RightBtmY:            event.Box.RightBtmY,
			LeftTopX:             event.Box.LeftTopX,
			RightBtmX:            event.Box.RightBtmX,
			Extend:               string(jsonData),
		}
		if err := SaveRecord(record); err != nil {
			logger.Error("写数据库错误:%v", err)
			continue
		}

	}

	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *ReceiveAlarmApi) AlarmRev(c *gin.Context) {
	// reqBody, _ := io.ReadAll(c.Request.Body)
	// logger.Info("接受告警:%s", string(reqBody))

	var alarmDates []mvc.AlarmDate
	if err := c.ShouldBindJSON(&alarmDates); err != nil {
		logger.Error("告警接收失败，参数错误:%v", err)
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	logger.Info("接受告警，%s", alarmDates[0].TaskID)

	for _, alarmDate := range alarmDates {
		dir := global.SystemConf.Picture.Dir + "/" + alarmDate.TaskID
		if !mvc.FileIsExisted(dir) {
			os.MkdirAll(dir, os.ModePerm)
		}

		now := time.Now()
		bigPicName := fmt.Sprintf("%c%c%d", Rand.Intn(26)+'A', Rand.Intn(26)+'A', now.Unix()) + ".jpg"
		if err := jpegSave(&alarmDate.SceneImageBase64, dir+"/"+bigPicName); err != nil {
			logger.Error("图片保存错误:%v", err)
			c.JSON(http.StatusOK, mvc.Fail(-1, "picture save error"))
			return
		}

		for _, event := range alarmDate.AnalyzeEvents {
			eventDir := global.SystemConf.Picture.Dir + "/" + alarmDate.TaskID + "/" + strconv.Itoa(event.Type)
			if !mvc.FileIsExisted(eventDir) {
				os.MkdirAll(eventDir, os.ModePerm)
			}
			smallPicName := fmt.Sprintf("%c%c%d", Rand.Intn(26)+'A', Rand.Intn(26)+'A', now.Unix()) + ".jpg"
			if err := jpegSave(&event.ImageBase64, eventDir+"/"+smallPicName); err != nil {
				logger.Error("图片保存错误:%v", err)
				continue
			}

			jsonData, err := json.Marshal(event.Extend)
			if err != nil {
				logger.Error("解析Extend错误")
			}

			record := mvc.Record{
				TaskId:             alarmDate.TaskID,
				SrcID:              alarmDate.SrcID,
				FrameIndex:         alarmDate.FrameIndex,
				BigPictureFilename: bigPicName,

				Type:                 event.Type,
				Date:                 now.Unix(),
				SamllPictureFilename: smallPicName,
				LeftTopY:             event.Box.LeftTopY,
				RightBtmY:            event.Box.RightBtmY,
				LeftTopX:             event.Box.LeftTopX,
				RightBtmX:            event.Box.RightBtmX,
				Extend:               string(jsonData),
			}
			if err := SaveRecord(record); err != nil {
				logger.Error("写数据库错误:%v", err)
				continue
			}
		}

	}

	c.JSON(http.StatusOK, mvc.Ok())
}

func jpegSave(base64ImageData *string, name string) error {
	// 将Base64数据解码成字节数组
	imageData, err := base64.StdEncoding.DecodeString(*base64ImageData)
	if err != nil {
		logger.Error("解码Base64数据失败:%v", err)
		return err
	}

	img, _, err := image.Decode(strings.NewReader(string(imageData)))
	if err != nil {
		logger.Error("解码图片失败:%v", err)
		return err
	}

	// 设置 JPEG 编码器的选项，包括图像质量（1-100，100表示最高质量）,数值越高，图片越清晰，磁盘占用也越高
	options := jpeg.Options{Quality: int(global.SystemConf.Picture.Quality)}

	// 图片保存
	outputFile, err := os.Create(name)
	if err != nil {
		logger.Error("创建输出文件失败:", err)
		return err
	}
	defer outputFile.Close()

	// 保存图片为JPEG格式
	err = jpeg.Encode(outputFile, img, &options)
	if err != nil {
		logger.Error("保存图片失败:", err)
		return err
	}
	return nil
}

// 保存告警
func SaveRecord(record mvc.Record) error {
	db := database.DB.Create(&record)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}
