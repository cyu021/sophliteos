package system

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"sophliteos/global"
	"sophliteos/logger"

	"sophliteos/pkg/buserr"
	"sophliteos/pkg/dto"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/repo"
	"sophliteos/pkg/service"
	"sophliteos/pkg/utils/httpclient"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

// Login
func (b *BaseApi) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	userName := req.UserName
	password := req.Password
	if userName == "" || password == "" {
		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的用户名或密码"))
		return
	}

	user, _ := repo.QueryUserWithName(userName)
	var token string
	if user == nil || user.Password != password { // 验证密码
		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的用户名或密码"))
		return
	} else {
		now := time.Now()
		if now.After(user.ExpireTime) {
			token = strings.ReplaceAll(uuid.New().String(), "-", "")
			user.Token = token
			user.LoginTime = now
			user.ExpireTime = now.Add(time.Hour * 2)
			repo.UpdateUser(user)
		} else {
			token = user.Token
		}
	}
	service.SetUser(token, user)
	service.SaveOptLog(c.Request, "登录", "登录")

	c.JSON(http.StatusOK, handle.Success(dto.LoginResponse{
		Token: token,
	}))
}

func (b *BaseApi) Logout(c *gin.Context) {
	var req dto.LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if req.Token != "" {
		user, err := repo.QueryUserWithToken(req.Token)
		if err == nil && user != nil {
			user.Token = ""
			user.ExpireTime = time.Now()
			repo.UpdateUser(user)
			service.SaveOptLog(c.Request, "注销登录", "注销登录")
			service.RemoveUser(req.Token)
			c.JSON(http.StatusOK, handle.Success(nil))

		} else {
			c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的token"))
			return
		}
	} else {
		c.JSON(http.StatusOK, handle.Fail(buserr.InvalidUsernameOrPassword, "无效的token"))
	}
}

func (b *BaseApi) AlarmListen(c *gin.Context) {
	var alarmRec repo.AlarmRec
	if err := c.ShouldBindJSON(&alarmRec); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if alarmRec.DiskName == "/dev/mmcblk0p4" || alarmRec.DiskName == "/dev/mmcblk0p2" || alarmRec.DiskName == "/dev/mmcblk0p5" {
		c.JSON(http.StatusOK, handle.Ok())
		return
	}

	alarm := repo.Alarm{
		DeviceSn:      alarmRec.DeviceSn,
		DeviceIp:      "",
		CreatedAt:     time.Now(),
		ComponentType: getType(alarmRec.Code),
		Code:          alarmRec.Code,
		Msg:           alarmRec.Msg,
	}

	if alarm.ComponentType == "disk" && alarmRec.Code < 0 {
		alarm.Msg = "磁盘" + alarmRec.DiskName + ": " + alarmRec.Msg
	}

	alarm.CoreUnitBoardSn = alarmRec.BoardSn
	if alarm.CoreUnitBoardSn == "" {
		alarm.CoreUnitBoardSn = alarmRec.DeviceSn
	}

	err := repo.SaveAlarm(alarm)
	handle.HandleError(err)
	c.JSON(http.StatusOK, handle.Ok())
}

func (b *BaseApi) AlgoRegister(c *gin.Context) {
	global.AlgoFlag = true
	c.JSON(http.StatusOK, handle.Ok())
}

func (b *BaseApi) AlgoExist(c *gin.Context) {
	c.JSON(http.StatusOK, handle.Success(global.AlgoFlag))
}

func (b *BaseApi) GetBase(c *gin.Context) {

	resource := GetArmResource(c)
	base := Base{
		DeviceSn:  resource.DeviceSn,
		DevicType: resource.DeviceType,
		IP:        resource.WanIP,
	}
	c.JSON(http.StatusOK, handle.Success(base))
}

func (b *BaseApi) GetDeviceIp(c *gin.Context) {

	resource := GetArmResource(c)

	c.JSON(http.StatusOK, handle.Success(resource.DeviceIP))
}

func Register() {
	var req struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}
	logger.Info("尝试注册algoliteos服务")

	data, _ := httpclient.NewRequest("127.0.0.1:8081/algorithm/register", "GET", nil, nil)
	json.Unmarshal(data, &req)

	if req.Msg != "ok" {
		logger.Info("algoliteos未运行服务")
		return
	}
	logger.Info("注册到algoliteos服务成功")
	global.AlgoFlag = true
}

func getType(code int) string {
	if code < 0 {
		code = -code
	}
	code = code / 1000
	var res string

	switch code {
	case 101:
		res = "cpu"
	case 102:
		res = "memory"
	case 103:
		res = "disk"
	case 104:
		res = "netCard"
	case 201:
		res = "board"
	case 202:
		res = "chip"
	default:
		res = ""
	}
	return res
}

type Base struct {
	DeviceSn  string `json:"deviceSn"`
	IP        string `json:"ip"`
	DevicType string `json:"deviceType"`
}
