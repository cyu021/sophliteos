package system

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"sophliteos/global"
	"sophliteos/logger"

	"sophliteos/pkg/buserr"
	"sophliteos/pkg/dto"
	"sophliteos/pkg/dto/response"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/service"
	"sophliteos/pkg/utils/ssh"

	"github.com/gin-gonic/gin"
)

type SsmUpgradeApi struct{}

var ssmFileName string

func (b *SsmUpgradeApi) SsmList(c *gin.Context) {
	var res []response.SsmVersion
	ctrl := global.SSmLists.CtrlSsm
	sv := response.SsmVersion{
		ChipIndex: -1,
		DeviceSn:  ctrl.DeviceSn,
		Ip:        ctrl.Ip,
		Version:   ctrl.Version,
	}
	res = append(res, sv)

	if global.DeviceType == "SE6" || global.DeviceType == "SE8" {
		res = append(res, global.SSmLists.CoreSsm...)
	}

	c.JSON(http.StatusOK, handle.Success(res))
}

func (b *SsmUpgradeApi) Upgrade(c *gin.Context) {
	var err error
	var sns string
	var pwd, user []byte

	module := c.Request.FormValue("module")
	logger.Info("module:%s\n", module)
	if module != Ctrl && module != Core {
		c.JSON(http.StatusOK, handle.Fail(buserr.UpgradeParamErr, "param error"))
		return
	}

	ssmFileName, err = saveFile(c.Request, "/data/ssm/")
	if err != nil {
		logger.Error("save file failed:%v", err)
		c.JSON(http.StatusOK, handle.FailWithMsg(buserr.UpgradeErr, "操作失败"))
		return
	}

	if !ssmArchIsOk(ssmFileName) {
		logger.Error("升级包: %s类型错误", ssmFileName)
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "升级包类型错误"))
		return
	}

	if module == Core {
		sns = c.Request.FormValue("sns")
		user, _ = base64.StdEncoding.DecodeString(c.Request.FormValue("user"))
		pwd, _ = base64.StdEncoding.DecodeString(c.Request.FormValue("pwd"))

		user, _ = base64.StdEncoding.DecodeString(string(user))
		pwd, _ = base64.StdEncoding.DecodeString(string(pwd))
	}

	logger.Info("file name:%s", ssmFileName)
	logger.Info("sns:%s\n  user:%s\n  pwd:%s\n", sns, string(user), string(pwd))

	global.BlockAllRequests = true

	if module == Ctrl {
		err = installSsmCtrl()
		if err != nil {
			logger.Error("ctrl update ssm failed:%v", err)
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "操作失败"))
			global.BlockAllRequests = false
			return
		}
		service.SaveOptLog(c.Request, "软件升级", "控制板SSM升级")
		c.JSON(http.StatusOK, handle.OkWithMsg("升级成功"))
	} else {
		if !checkPassword(string(user), string(pwd)) {
			logger.Error("用户名或密码错误")
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "用户名或密码错误"))
			global.BlockAllRequests = false
			return
		}

		c.JSON(http.StatusOK, handle.OkWithMsg("核心板正在升级ssm，预计5分钟升级完成"))
		err = scpInstallSsm(sns, string(user), string(pwd))
		if err != nil {
			logger.Error("core update ssm failed:%v", err)
			c.JSON(http.StatusOK, handle.FailWithMsg(buserr.UpgradeErr, "操作失败"))
			global.BlockAllRequests = false
			return
		}
		service.SaveOptLog(c.Request, "软件升级", "核心板SSM升级")
	}

	global.BlockAllRequests = false

}

func installSsmCtrl() error {
	cmd := exec.Command("tar", "-xzf", ssmFileName, "-C", "/data/ssm/")
	cmd.Dir = "/data/ssm"

	fmt.Println("Executing command:", cmd.String())

	// 执行命令
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("unzip file %s failed:%v  out:%s", ssmFileName, err, out)
	}

	defer func() {
		err = os.RemoveAll("/data/ssm/install")
		if err != nil {
			logger.Info("deleting directory:", err)
		}
	}()

	script := "/data/ssm/install/update-bmssm.sh"
	// 检查脚本文件是否存在
	_, err = os.Stat(script)
	if err != nil {
		logger.Error("Script file not found:%v", err)
		return err
	}
	cmd = exec.Command("sudo", "/bin/bash", script)
	cmd.Dir = "/data/ssm/install"
	err = cmd.Run()
	if err != nil {
		logger.Error("run upgrade script failed:%v", err)
		return err
	}

	logger.Info("ssm upgrade successful!")
	return nil
}

func scpInstallSsm(sns, user, pwd string) error {

	res, err := service.ExecCore(dto.ExecCorex{
		Cmd:   "mkdir -p /data/ssm && sudo chown linaro /data/ssm",
		User:  "linaro",
		Pwd:   "linaro",
		DevId: sns,
	})
	logger.Info("ExecCore:%v, err:%v", res, err)

	res, err = service.ScpFile(dto.Scp2Core{
		SrcCtrl:  "/data/ssm/" + ssmFileName,
		DstCore:  "/data/ssm/",
		CtrlUser: user,
		CtrlPwd:  pwd,
		CoreUser: "linaro",
		CorePwd:  "linaro",
	})

	logger.Info("ScpFile return:%v; err:%v", res, err)
	if err != nil {
		logger.Error("ssm scp file  failed:%v", err)
		return err
	}

	command := fmt.Sprintf("cd /data/ssm && sudo tar -xzvf %s", ssmFileName)

	res, err = service.ExecCore(dto.ExecCorex{
		Cmd:   command,
		User:  "linaro",
		Pwd:   "linaro",
		DevId: sns,
	})
	logger.Info("core boards upgrade ssm:%v\n err:%v", res, err)

	command = "cd /data/ssm/install && sudo ./update-bmssm.sh"

	res, err = service.ExecCore(dto.ExecCorex{
		Cmd:   command,
		User:  "linaro",
		Pwd:   "linaro",
		DevId: sns,
	})
	logger.Info("core boards upgrade ssm:%v\n err:%v", res, err)

	return err
}

func ssmArchIsOk(name string) bool {

	switch global.Arch {
	case "riscv64":
		return name == "bmssm-riscv64-v1.2.0.tgz"
	case "amd64":
		return name == "bmssm-pcie-v1.2.0.tgz"
	case "arm64":
		return name == "bmssm-arm64-v1.2.0.tgz"
	default:
		return false
	}
}

func checkPassword(userName, pwd string) bool {
	connInfo := ssh.ConnInfo{
		User:       userName,
		Addr:       "127.0.0.1",
		Port:       22,
		AuthMode:   "password",
		Password:   pwd,
		PrivateKey: []byte(""),
	}

	return hostService.TestByInfo(connInfo)
}
