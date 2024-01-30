package system

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"sophliteos/logger"

	"sophliteos/pkg/buserr"
	"sophliteos/pkg/dto"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/service"

	"github.com/gin-gonic/gin"
)

type information struct {
	Model string `json:"model"`
}

type BasicApi struct{}

func (b *BasicApi) BasicMod(c *gin.Context) {
	var req dto.BasicSettings
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	result, err := service.SetBasic(req)
	handle.HandleError(err, buserr.SetDeviceInfoErr)
	service.SaveOptLog(c.Request, "基础信息设置", "设备名称修改为："+req.Name)
	c.JSON(http.StatusOK, handle.Success(handle.Handle(result, buserr.SetDeviceInfoErr)))
}

func (b *BasicApi) GetBasic(c *gin.Context) {

	ctrlBasic, _, err := service.GetCtrlBasic()
	if err != nil {
		logger.Error("error is :%v", err)
		logger.Error("ctrlBasic is :%v", ctrlBasic)
	}

	c.JSON(http.StatusOK, handle.Success(ctrlBasic))
}

// 云平台信息设置
func (b *BasicApi) ModBasic(c *gin.Context) {

	body, _ := io.ReadAll(c.Request.Body)
	_, err := service.ModBasic(body)
	if err != nil {
		logger.Error("error is :%v", err)
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设置失败"))
		return
	}

	service.SaveOptLog(c.Request, "基础信息设置", "云平台信息设置")
	c.JSON(http.StatusOK, handle.Success(handle.Ok()))
}

func GetDeviceType() string {
	data, err := os.ReadFile("/sys/bus/i2c/devices/1-0017/information")
	if err != nil {
		logger.Error("读取文件错误:%v", err)
		return ""
	}

	// 解析JSON数据
	var info information
	if err := json.Unmarshal(data, &info); err != nil {
		logger.Error("解析JSON错误:", err)
		return ""
	}

	if strings.Contains(info.Model, "SE5") {
		return "SE5"
	} else if strings.Contains(info.Model, "SE6") {
		return "SE6"
	} else if strings.Contains(info.Model, "BM1684X EVB") {
		return "SE7"
	}
	return ""
}
