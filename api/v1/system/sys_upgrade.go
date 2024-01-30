package system

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"sophliteos/global"
	"sophliteos/logger"

	"sophliteos/pkg/handle"
	"sophliteos/pkg/service"
	"sophliteos/pkg/utils/cmd"

	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/go-update"
)

type UpgradeApi struct{}

var filename string

func (b *UpgradeApi) Upgrade(c *gin.Context) {
	var err error

	filename, err = saveFile(c.Request, "/data/sophliteos/")
	if err != nil {
		logger.Error("update failed", err)
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "操作失败"))
		return
	}

	if filename == "algoliteos-linux_arm64.tgz" || filename == "algoliteos-linux_amd64.tgz" {
		if err := upgradeAlgo(); err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "操作失败"))
		} else {
			c.JSON(http.StatusOK, handle.OkWithMsg("算法插件升级成功"))
			service.SaveOptLog(c.Request, "软件升级", "算法业务插件升级")
		}
		return
	}

	if !archIsOk(filename) {
		logger.Error("升级包类型错误")
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "升级包类型错误"))
		return
	}

	err = upgradeLiteOs()
	if err != nil {
		logger.Error("upgrade failed", err)
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "操作失败"))
		return
	}
	global.BlockAllRequests = true
	service.SaveOptLog(c.Request, "软件升级", "LiteOS升级")

	// 重新执行更新后的程序
	go restartUpgradedProgram()
	c.JSON(http.StatusOK, handle.OkWithMsg("升级成功，LiteOS正在重启，请一分钟后刷新页面重新进入"))
}

func archIsOk(name string) bool {
	switch global.Arch {
	case "riscv64":
		return name == "sophliteos-linux_riscv64.tgz"
	case "amd64":
		return name == "sophliteos-linux_amd64.tgz"
	case "arm64":
		return name == "sophliteos-linux_arm64.tgz"
	default:
		return false
	}
}

func upgradeLiteOs() error {

	command := exec.Command("tar", "-xzf", filename, "-C", "/data/sophliteos/")
	command.Dir = "/data/sophliteos"

	// 执行命令
	err := command.Run()
	if err != nil {
		logger.Error("tar failed", err)
	}

	script := "/data/sophliteos/upgrade.sh"
	_, err = cmd.ExecScript(script, "/data/sophliteos")

	if err != nil {
		logger.Error("script failed", err)
		return err
	}
	// 读取升级文件
	updatePath := "/data/sophliteos/sophliteos"
	updateFile, err := os.Open(updatePath)
	cmdPath := os.Args[0]
	if err != nil {
		return err
	}
	defer updateFile.Close()

	// 执行自更新操作
	err = update.Apply(updateFile, update.Options{
		TargetPath: cmdPath,
	})
	if err != nil {
		if rollbackErr := update.RollbackError(err); rollbackErr != nil {
			logger.Error("Failed to rollback from bad update: %v", rollbackErr)
		}
		return err
	}

	logger.Info("sophliteos self upgrade successful!")
	return nil
}

func restartUpgradedProgram() {
	time.Sleep(5 * time.Second)
	// 启动新进程
	if err := syscall.Exec(os.Args[0], os.Args, os.Environ()); err != nil {
		logger.Error("Failed to restart: %v", err)
	}

	cmd := exec.Command("rm", "-f", filename)
	cmd.Dir = "/data/sophliteos"

	// 执行命令
	err := cmd.Run()
	if err != nil {
		logger.Error("tar rm failed", err)
	}

	// 退出当前进程
	os.Exit(0)
}

func upgradeAlgo() error {
	if err := os.MkdirAll("/data/sophliteos/algo", 0755); err != nil {
		logger.Error("Failed to create directory", err)
		return err
	}
	os.Chmod("/data/sophliteos/algo", 0755)

	cmd := exec.Command("tar", "-xzf", filename, "-C", "/data/sophliteos/algo")
	cmd.Dir = "/data/sophliteos"

	// 执行命令
	err := cmd.Run()
	if err != nil {
		logger.Error("tar failed", err)
	}

	script := "/data/sophliteos/algo/upgrade.sh"
	// 检查脚本文件是否存在
	_, err = os.Stat(script)
	if err != nil {
		logger.Error("Script file not found:", err)
		return err
	}
	cmd = exec.Command("sudo", "/bin/bash", script)
	cmd.Dir = "/data/sophliteos/algo"
	err = cmd.Run()
	if err != nil {
		logger.Error("script failed", err)
		return err
	}

	logger.Info("algoliteos upgrade successful!")
	return nil
}

// 文件上传控制
func saveFile(request *http.Request, dir string) (string, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		logger.Error("Failed to create directory", err)
		return "", err
	}
	os.Chmod(dir, 0755)

	file, handler, err := request.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	f, err := os.OpenFile(dir+handler.Filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(handler.Filename)
		_ = request.MultipartForm.RemoveAll()
	}()
	return handler.Filename, err
}
