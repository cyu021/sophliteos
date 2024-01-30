package system

import (
	"net/http"

	"sophliteos/global"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/repo"
	"sophliteos/pkg/service"

	"github.com/gin-gonic/gin"
)

type LogApi struct{}

var layout string = "2006-01-02+15:04:05"

func (b *LogApi) AlarmLogQuery(c *gin.Context) {

	r := repo.WrapperQuery(c.Request)
	componentType := r.Get("componentType")
	startTime := r.GetDate("startTime")
	endTime := r.GetDate("endTime")
	pageNo := r.Required().GetInt("pageNo")
	pageSize := r.Required().GetInt("pageSize")

	var alarm repo.Alarm
	alarm.ComponentType = componentType
	page := repo.QueryAlarms(&alarm, *pageNo, *pageSize, startTime, endTime)

	if page.Total > 0 {
		cr, _ := service.GetCtrlResource()
		dMap := make(map[string]string)
		dMap[cr.DeviceSn] = cr.CentralProcessingUnit.NetCard[0].IP
		for _, board := range cr.CoreComputingUnit.Board {
			if len(board.CoreSys.NetCards) > 0 {
				dMap[board.BoardSn] = board.CoreSys.NetCards[0].IP
			}
		}
		items := page.Items.([]repo.Alarm)
		for i := 0; i < len(items); i++ {
			items[i].DeviceIp = dMap[items[i].CoreUnitBoardSn]
		}
		page.Items = items
	}
	c.JSON(http.StatusOK, handle.Success(page))
}

func (b *LogApi) OptLogQuery(c *gin.Context) {
	r := repo.WrapperQuery(c.Request)
	startTime := r.GetDate("startTime")
	endTime := r.GetDate("endTime")
	pageNo := r.Required().GetInt("pageNo")
	pageSize := r.Required().GetInt("pageSize")

	var optLog repo.OptLog
	optLog.UserName = c.Query("username")
	optLog.OperationFunc = c.Query("operationFunc")
	optLog.OperationIP = c.Query("operationIp")
	optLog.OperationContent = c.Query("operationContent")

	page := repo.QueryOptLogs(&optLog, *pageNo, *pageSize, startTime, endTime)
	c.JSON(http.StatusOK, handle.Success(page))

}

func (b *LogApi) GetQueryInfo(c *gin.Context) {
	var operationFunc, username, ip []string
	global.DB.Model(&repo.OptLog{}).Pluck("distinct(operation_func)", &operationFunc)
	global.DB.Model(&repo.OptLog{}).Pluck("distinct(user_name)", &username)
	global.DB.Model(&repo.OptLog{}).Pluck("distinct(operation_ip)", &ip)

	res := repo.OptLogInfo{
		UserName:      username,
		OperationIP:   ip,
		OperationFunc: operationFunc,
	}

	c.JSON(http.StatusOK, handle.Success(res))
}
