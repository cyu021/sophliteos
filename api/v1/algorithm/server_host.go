package algorithm

import (
	"algoliteos/config"
	"algoliteos/global"
	"algoliteos/logger"
	"algoliteos/mvc"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type HostApi struct{}

var registerEntryID cron.EntryID
var cornRegister *cron.Cron

func init() {
	cornRegister = cron.New(cron.WithSeconds())
	var err error
	registerEntryID, err = cornRegister.AddFunc("0/5 * * * * ?", func() { //每5秒 执行任务
		register()
	})
	if err != nil {
		fmt.Println("err:", err)
	}

	cornRegister.Start()
}

// 注册到sophliteos服务
func register() {
	var req struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}
	logger.Info("尝试注册sophliteos服务")

	data := mvc.NewRequestWithHeaders("127.0.0.1:8080/api/register", "GET", nil, nil)
	json.Unmarshal(data, &req)

	if req.Msg != "ok" {
		return
	}
	logger.Info("注册到sophliteos服务成功")
	cornRegister.Remove(registerEntryID)
}

// 接收sophliteos服务的检测
func (b *HostApi) GetRegisterHost(c *gin.Context) {
	c.JSON(http.StatusOK, mvc.Ok())
}

// 获取告警ip
func (b *HostApi) GetAlarmIP(c *gin.Context) {
	host, _, err := mvc.ExtractIP(global.SystemConf.UploadHost)
	if err != nil {
		c.JSON(http.StatusOK, mvc.Fail(-1, "地址错误"))
	}
	c.JSON(http.StatusOK, mvc.Success(host))
}

// 配置告警ip
func (b *HostApi) ModAlarmIP(c *gin.Context) {
	var req mvc.UploadIp
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	global.SystemConf.UploadHost = req.Ip + ":8081"

	err := config.SaveConfig()
	if err != nil {
		c.JSON(http.StatusOK, mvc.Fail(-1, "配置失败"))
		return
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

// 配置算法服务IP和端口
func (b *HostApi) AddServerHost(c *gin.Context) {
	var serverHost mvc.AlgorithmHost
	if err := c.ShouldBindJSON(&serverHost); err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误"))
		return
	}

	global.SystemConf.AlgorithmHost = serverHost.Ip + ":" + strconv.Itoa(serverHost.Port)

	err := config.SaveConfig()
	if err != nil {
		c.JSON(http.StatusOK, mvc.Fail(-1, "配置失败"))
		return
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

// 获取算法服务ip和端口，接口
func (b *HostApi) GetServerHost(c *gin.Context) {

	parts := strings.Split(global.SystemConf.AlgorithmHost, ":")
	port, _ := strconv.Atoi(parts[1])
	algoHost := mvc.AlgorithmHost{
		Ip:   parts[0],
		Port: port,
	}
	c.JSON(http.StatusOK, mvc.Success(algoHost))
}
