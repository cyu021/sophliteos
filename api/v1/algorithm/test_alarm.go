package algorithm

import (
	"algoliteos/global"
	"algoliteos/logger"
	"algoliteos/mvc"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type TestApi struct{}

// 简易的告警上报测试接口，用于开发调试，不涉及业务
func (b *TestApi) TestAlarm(c *gin.Context) {
	var req TestAlarmInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	var events []mvc.AnalyzeEvent
	for _, i := range req.AnalyzeEvents {
		event := mvc.AnalyzeEvent{
			Type: i,
			Box: mvc.Box{
				LeftTopX:  20,
				LeftTopY:  20,
				RightBtmY: 200,
				RightBtmX: 200,
			},
			ImageBase64: ImageToBase64(req.Path),
		}
		events = append(events, event)
	}

	alarmDates := []mvc.AlarmDate{
		{
			TaskID:           req.TaskId,
			SrcID:            req.SrcID,
			FrameIndex:       req.FrameIndex,
			SceneImageBase64: ImageToBase64(req.Path),
			AnalyzeEvents:    events,
		},
	}

	body, _ := mvc.StructToString(alarmDates)

	// logger.Info(body)

	var msg Msg
	res := mvc.NewRequestWithHeaders(global.SystemConf.UploadHost+"/algorithm/upload", "POST", nil, []byte(body))
	err := json.Unmarshal(res, &msg)
	if err != nil {
		logger.Error("请求失败:%v", err)
		return
	}

	logger.Info("%v", msg)

	c.JSON(http.StatusOK, mvc.Ok())
}

type TestAlarmInfo struct {
	TaskId        string `json:"taskId"`
	Path          string `json:"path"`
	SrcID         string `json:"srcID"`
	FrameIndex    int    `json:"frameIndex"`
	AnalyzeEvents []int  `json:"analyzeEvents"`
}

type Msg struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func ImageToBase64(filePath string) string {
	// 读取文件内容
	imageBytes, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}

	// 将字节内容编码为 Base64
	base64Encoding := base64.StdEncoding.EncodeToString(imageBytes)

	return base64Encoding
}
