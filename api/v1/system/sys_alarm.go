package system

import (
	"io"
	"net/http"

	"sophliteos/global"
	"sophliteos/logger"

	"sophliteos/pkg/buserr"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/service"

	"github.com/gin-gonic/gin"
)

type AlarmApi struct{}

func (b *AlarmApi) AlarmQuery(c *gin.Context) {
	ctrl, _, err := service.GetCtrlBasic()
	handle.HandleError(err)
	c.JSON(http.StatusOK, handle.Success(ctrl.Configure.AlarmThreshold))
}

func (b *AlarmApi) AlarmSet(c *gin.Context) {

	data, _ := io.ReadAll(c.Request.Body)

	var err error
	switch global.DeviceType {
	case "SE5", "SE7", "SE9", "PCIE":
		_, err = service.SetAlarmSe5(data)
	case "SE6", "SE8":
		_, err = service.SetAlarm(data)
	}

	if err != nil {
		logger.Error("设置告警失败：%v", err)
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设置告警失败"))
		return
	}

	ssmResult, err := service.SubscribeAlarm()
	if err != nil {
		logger.Error("设置告警失败：%v", err)
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设置告警失败"))
		return
	}
	service.SaveOptLog(c.Request, "设置告警", "设置告警")
	c.JSON(http.StatusOK, handle.Handle(ssmResult, buserr.SetAlarmErr))
}
