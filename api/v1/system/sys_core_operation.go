package system

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"sophliteos/pkg/buserr"
	"sophliteos/pkg/dto"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/service"
	"sophliteos/pkg/utils/i18n"

	"github.com/gin-gonic/gin"
)

type CoreOperationApi struct{}

func init() {
	i18n.SetString(i18n.Zh, getCode(dto.Reboot), "重启")
	i18n.SetString(i18n.En, getCode(dto.Reboot), "Reboot")
	i18n.SetString(i18n.Zh, getCode(dto.Shutdown), "关机")
	i18n.SetString(i18n.En, getCode(dto.Shutdown), "Shutdown")
}

func getCode(code int) string {
	return fmt.Sprintf("CoreOperation:%v", code)
}

func (b *CoreOperationApi) PostOperation(c *gin.Context) {

	ope := dto.CoreOperation{}
	data, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(data, &ope)
	result, err := service.CoreOperate(ope.Number, ope.Type)
	operation := i18n.GetString(service.GetLang(c.Request), getCode(ope.Type))
	handle.HandleResult(result, err, buserr.DeviceOperationErr)
	service.SaveOptLog(c.Request, "核心板操作：%s", operation)
	c.JSON(http.StatusOK, handle.Success(nil))
}
