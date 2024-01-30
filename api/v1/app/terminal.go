package app

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sophliteos/api/v1/system"
	"sophliteos/global"
	"sophliteos/logger"

	"sophliteos/pkg/dto"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/service"
	"sophliteos/pkg/utils/cmd"
	"sophliteos/pkg/utils/ssh"
	"sophliteos/pkg/utils/terminal"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type TerminalApi struct{}

func (b *TerminalApi) WsSsh(c *gin.Context) {
	logger.Info("ssh websocket start")

	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("gin context http handler failed, err: %v", err)
		return
	}
	defer wsConn.Close()

	ip := c.Query("ip")
	if wshandleError(wsConn, errors.WithMessage(err, "invalid param id in request")) {
		return
	}
	cols, err := strconv.Atoi(c.DefaultQuery("cols", "80"))
	if wshandleError(wsConn, errors.WithMessage(err, "invalid param cols in request")) {
		return
	}
	rows, err := strconv.Atoi(c.DefaultQuery("rows", "40"))
	if wshandleError(wsConn, errors.WithMessage(err, "invalid param rows in request")) {
		return
	}

	var connInfo ssh.ConnInfo
	if core := c.Query("core"); core != "1" {
		// 从数据库获取主机信息
		host, err := hostService.GetHostInfo(ip)
		if wshandleError(wsConn, errors.WithMessage(err, "load host info by id failed")) {
			return
		}
		connInfo = ssh.ConnInfo{
			User:       host.User,
			Addr:       host.Addr,
			Port:       int(host.Port),
			AuthMode:   "password",
			Password:   host.Password,
			PrivateKey: []byte(""),
		}
	} else {
		// SE6核心板直连
		connInfo = ssh.ConnInfo{
			User:       "linaro",
			Addr:       ip,
			Port:       22,
			AuthMode:   "password",
			Password:   "linaro",
			PrivateKey: []byte(""),
		}
	}

	client, err := connInfo.NewClient()
	if wshandleError(wsConn, errors.WithMessage(err, "failed to set up the connection. Please check the host information")) {
		return
	}
	defer client.Close()

	sws, err := terminal.NewLogicSshWsSession(cols, rows, true, connInfo.Client, wsConn)
	if wshandleError(wsConn, err) {
		return
	}
	defer sws.Close()

	logger.Info("ssh connect success")

	quitChan := make(chan bool, 3)
	sws.Start(quitChan)
	go sws.Wait(quitChan)

	<-quitChan

	logger.Info("ssh连接结束，ip(%s),浏览器ip(%s) ", ip, c.ClientIP())

	if wshandleError(wsConn, err) {
		return
	}
}

// 测试主机连接
func (b *TerminalApi) TestByInfo(c *gin.Context) {
	var req dto.HostAddr
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	host, err := hostService.GetHostInfo(req.Addr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, req.Addr+"主机不存在"))
		return
	}
	connInfo := ssh.ConnInfo{
		User:       host.User,
		Addr:       host.Addr,
		Port:       int(host.Port),
		AuthMode:   "password",
		Password:   host.Password,
		PrivateKey: []byte(""),
	}

	connStatus := hostService.TestByInfo(connInfo)
	c.JSON(http.StatusOK, handle.Success(connStatus))
}

// SE6设备核心板列表
func (b *TerminalApi) GetSE6Core(c *gin.Context) {
	if len(global.SE6Cores) == 0 {
		system.GetArmResource(c)
	}
	c.JSON(http.StatusOK, handle.Success(global.SE6Cores))
}

// 获取快速命令列表
func (b *TerminalApi) ListCommand(c *gin.Context) {
	list, err := hostService.List()
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	c.JSON(http.StatusOK, handle.Success(list))
}

func (b *TerminalApi) CreateCommand(c *gin.Context) {
	var req dto.CommandInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	if err := hostService.Create(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 删除快速命令
func (b *TerminalApi) DeleteCommand(c *gin.Context) {
	var req dto.CommandInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	if err := global.VALID.Struct(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	if err := hostService.Delete(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 更新快速命令
func (b *TerminalApi) UpdateCommand(c *gin.Context) {
	var req dto.CommandInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	if err := global.VALID.Struct(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	if err := hostService.Update(req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 主机列表
func (b *TerminalApi) GetHosts(c *gin.Context) {

	hosts, err := hostService.ListHost()
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Success(hosts))
}

// 创建主机
func (b *TerminalApi) CreateHost(c *gin.Context) {
	var req dto.HostOperate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	if err := global.VALID.Struct(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	var pwd string
	var err error
	if pwd, err = hostService.EncryptHost(req.Password); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	connInfo := ssh.ConnInfo{
		User:       req.User,
		Addr:       req.Addr,
		Port:       int(req.Port),
		AuthMode:   "password",
		Password:   pwd,
		PrivateKey: []byte(""),
	}
	if hostService.TestByInfo(connInfo) == false {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "添加失败:"+"用户名或密码错误"))
		return
	}

	err = hostService.CreateHost(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	service.SaveOptLog(c.Request, "终端", "添加主机："+req.Addr)
	c.JSON(http.StatusOK, handle.Ok())
}

// 删除主机
func (b *TerminalApi) DeleteHost(c *gin.Context) {
	var req dto.HostOperate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	if err := global.VALID.Struct(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	err := hostService.DeleteHost(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	service.SaveOptLog(c.Request, "终端", "删除主机："+req.Addr)
	c.JSON(http.StatusOK, handle.Ok())
}

// 修改主机
func (b *TerminalApi) UpdateHost(c *gin.Context) {
	var req dto.HostOperate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	if err := global.VALID.Struct(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	var pwd string
	var err error
	if pwd, err = hostService.EncryptHost(req.Password); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	connInfo := ssh.ConnInfo{
		User:       req.User,
		Addr:       req.Addr,
		Port:       int(req.Port),
		AuthMode:   "password",
		Password:   pwd,
		PrivateKey: []byte(""),
	}
	if hostService.TestByInfo(connInfo) == false {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "修改失败:"+"用户名或密码错误"))
		return
	}

	err = hostService.UpdateHost(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

func (b *TerminalApi) ContainerWsSsh(c *gin.Context) {
	logger.Info("docker websocket start")

	var Docker string
	switch global.Arch {
	case "riscv64":
		Docker = "podman"
	default:
		Docker = "docker"
	}
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("gin context http handler failed, err: %v", err)
		return
	}
	defer wsConn.Close()

	containerID := c.Query("containerid")
	command := c.Query("command")
	user := c.Query("user")
	if len(command) == 0 || len(containerID) == 0 {
		if wshandleError(wsConn, errors.New("error param of command or containerID")) {
			logger.Error(("command or containerID error"))
			return
		}
	}
	cols, err := strconv.Atoi(c.DefaultQuery("cols", "80"))
	if wshandleError(wsConn, errors.WithMessage(err, "invalid param cols in request")) {
		logger.Error(("invalid param cols in request"))
		return
	}
	rows, err := strconv.Atoi(c.DefaultQuery("rows", "40"))
	if wshandleError(wsConn, errors.WithMessage(err, "invalid param rows in request")) {
		logger.Error(("invalid param rows in request"))
		return
	}

	cmds := []string{"exec", containerID, command}
	if len(user) != 0 {
		cmds = []string{"exec", "-u", user, containerID, command}
	}
	if cmd.CheckIllegal(user, containerID, command) {
		if wshandleError(wsConn, errors.New("  The command contains illegal characters.")) {
			logger.Error(("The command contains illegal characters."))
			return
		}
	}

	stdout, err := cmd.ExecWithCheck(Docker, cmds...)
	if wshandleError(wsConn, errors.WithMessage(err, stdout)) {
		logger.Error(("error"))
		return
	}

	commands := fmt.Sprintf("%s exec -it %s %s", Docker, containerID, command)
	if len(user) != 0 {
		commands = fmt.Sprintf("%s exec -it -u %s %s %s", Docker, user, containerID, command)
	}
	pidMap := loadMapFromDockerTop(containerID)
	slave, err := terminal.NewCommand(commands)
	if wshandleError(wsConn, err) {
		logger.Error(("error"))
		return
	}
	defer killBash(containerID, command, pidMap)
	defer slave.Close()

	tty, err := terminal.NewLocalWsSession(cols, rows, wsConn, slave)
	if wshandleError(wsConn, err) {
		logger.Error(("error"))
		return
	}

	logger.Info("docker ssh connect success")

	quitChan := make(chan bool, 3)
	tty.Start(quitChan)
	go slave.Wait(quitChan)

	<-quitChan

	logger.Info("docker ssh连接结束，容器id(%s) ", containerID)
	if wshandleError(wsConn, err) {
		return
	}
}

func wshandleError(ws *websocket.Conn, err error) bool {
	if err != nil {
		// logger.Error("handler ws faled:, err: %v", err)
		dt := time.Now().Add(time.Second)
		// 发生错误,强制关闭连接
		if ctlerr := ws.WriteControl(websocket.CloseMessage, []byte(err.Error()), dt); ctlerr != nil {
			wsData, err := json.Marshal(terminal.WsMsg{
				Type: terminal.WsMsgCmd,
				Data: base64.StdEncoding.EncodeToString([]byte(err.Error())),
			})
			if err != nil {
				_ = ws.WriteMessage(websocket.TextMessage, []byte("{\"type\":\"cmd\",\"data\":\"failed to encoding to json\"}"))
			} else {
				_ = ws.WriteMessage(websocket.TextMessage, wsData)
			}
		}
		return true
	}
	return false
}

func loadMapFromDockerTop(containerID string) map[string]string {
	pidMap := make(map[string]string)
	sudo := cmd.SudoHandleCmd()

	var Docker string
	switch global.Arch {
	case "riscv64":
		Docker = "podman"
	default:
		Docker = "docker"
	}

	stdout, err := cmd.Execf("%s %s top %s -eo pid,command ", sudo, Docker, containerID)
	if err != nil {
		return pidMap
	}
	lines := strings.Split(stdout, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		pidMap[parts[0]] = parts[1]
	}
	return pidMap
}

func killBash(containerID, comm string, pidMap map[string]string) {
	sudo := cmd.SudoHandleCmd()
	newPidMap := loadMapFromDockerTop(containerID)
	for pid, command := range newPidMap {
		isOld := false
		for pid2 := range pidMap {
			if pid == pid2 {
				isOld = true
				break
			}
		}
		if !isOld && command == comm {
			_, _ = cmd.Execf("%s kill -9 %s", sudo, pid)
		}
	}
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
