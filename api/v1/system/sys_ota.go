package system

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"sophliteos/global"
	"sophliteos/logger"

	"sophliteos/pkg/buserr"
	"sophliteos/pkg/dto"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/service"
	"sophliteos/pkg/utils/cmd"

	"github.com/gin-gonic/gin"
)

type OtaApi struct{}

const (
	Ctrl = "ctrl"
	Core = "core"
)

var (
	ctrlFileName string
	coreFileName string
	ctrlFileMd5  string
	coreFileMd5  string
)

func init() {

}

func (b *OtaApi) OtaFileChunked(c *gin.Context) {
	chunkIndex := c.Request.FormValue("chunkIndex") // 分片的索引
	totalChunks := c.Request.FormValue("totalChunks")
	chunkName := c.Request.FormValue("fileName")
	md5Value := strings.ToLower(c.Request.FormValue("md5"))
	module := c.Request.FormValue("module")

	index, _ := strconv.Atoi(chunkIndex)
	total, _ := strconv.Atoi(totalChunks)

	switch module {
	case "crtl":
		if !strings.Contains(chunkName, "sdcard") {
			logger.Error("文件错误")
			c.JSON(http.StatusOK, handle.Fail(-2, "升级文件上传错误，请上传名称为sdcard.tgz的卡刷包"))
			return
		}
	case "core":
		if !strings.Contains(chunkName, "tftp") {
			logger.Error("文件错误")
			c.JSON(http.StatusOK, handle.Fail(-2, "升级文件上传错误，请上传名称为tftp.tgz的卡刷包"))
			return
		}
	}

	// 创建存储分片的目录
	chunksDir := filepath.Join("/data/sophliteos", "upload")
	os.MkdirAll(chunksDir, os.ModePerm)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, handle.Fail(-1, "file error"))
		return
	}
	// 分片文件的存储路径
	chunkFilePath := filepath.Join(chunksDir, chunkName+"-"+chunkIndex)

	// 保存分片文件
	if err := c.SaveUploadedFile(file, chunkFilePath); err != nil {
		c.JSON(http.StatusOK, handle.Fail(-1, "SaveUploadedFile error"))
		return
	}

	if total == index+1 {
		ctrlFileMd5, err = MergeChunked(total, md5Value, module, chunkName)
		if err != nil {
			c.JSON(http.StatusOK, handle.Fail(-1, "文件上传失败"+err.Error()))
			return
		}
		service.SaveOptLog(c.Request, "系统升级", "升级包上传")
		c.JSON(http.StatusOK, handle.Success(ctrlFileMd5))
	} else {
		c.JSON(http.StatusOK, handle.Ok())
	}

}

func MergeChunked(total int, md5Value, module, name string) (string, error) {

	err := os.RemoveAll("/data/ota")
	if err != nil {
		logger.Error("rm failed %v", err)
	}

	// 最终文件的路径
	finalFilePath := filepath.Join("/data/ota/", name)
	os.MkdirAll(filepath.Dir(finalFilePath), os.ModePerm)

	// 创建最终文件
	finalFile, err := os.Create(finalFilePath)
	if err != nil {
		logger.Error("创建最终文件失败： %v", err)
		return "", err
	}
	defer finalFile.Close()

	// 合并所有分片
	for i := 0; i < total; i++ {
		chunkFilePath := filepath.Join("/data/sophliteos/upload", name) + "-" + strconv.Itoa(i)
		chunkFile, err := os.Open(chunkFilePath)
		if err != nil {
			logger.Error("合并分片失败： %v", err)
			return "", err
		}

		// 将分片内容写入最终文件
		if _, err := io.Copy(finalFile, chunkFile); err != nil {
			chunkFile.Close()
			logger.Error("分片内容写入最终文件失败： %v", err)
			return "", err
		}
		chunkFile.Close()

		// 删除已经合并的分片文件
		os.Remove(chunkFilePath)
	}

	md5String, err := calculateFileMD5("/data/ota/" + name)
	if err != nil {
		logger.Error("md5计算失败： %v", err)
		return "", err
	}
	if md5String != md5Value {
		logger.Error("md5值不一致")
		return "", errors.New("md5 error")
	}

	switch module {
	case "crtl":
		ctrlFileName = name
		ctrlFileMd5 = md5String
	case "core":
		go DeCompressTftpFile(name)
	}
	return md5String, nil

}

func (b *OtaApi) OtaFile(c *gin.Context) {
	// 参数判断
	md5Value := strings.ToLower(c.Request.FormValue("md5"))
	module := c.Request.FormValue("module")
	if module != Ctrl && module != Core {
		c.JSON(http.StatusOK, handle.Fail(buserr.UpgradeParamErr, "param error"))
		return
	}

	err := os.RemoveAll("/data/ota")
	if err != nil {
		logger.Error("rm failed %v", err)
	}

	var otaFile string
	switch module {
	case Ctrl:
		otaFile, err = saveFile(c.Request, "/data/ota/")
		if err != nil {
			logger.Error("update failed %v", err)
			c.JSON(http.StatusOK, handle.FailWithMsg(buserr.UpgradeErr, "文件上传失败"))
			return
		}
		ctrlFileName = otaFile
	case Core:
		otaFile, err = saveFile(c.Request, "/data/ota/")
		if err != nil {
			logger.Error("update failed %v", err)
			c.JSON(http.StatusOK, handle.FailWithMsg(buserr.UpgradeErr, "文件上传失败"))
			return
		}
		coreFileName = otaFile
	}

	md5String, err := calculateFileMD5("/data/ota/" + otaFile)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(buserr.UpgradeErr, "文件上传失败"))
		return
	}

	logger.Info("文件名:%s", otaFile)
	logger.Info("初始文件MD5值:%s", md5Value)
	logger.Info("接受文件MD5值:%s", md5String)

	if md5String != md5Value {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "文件上传失败:MD5值不一致"))
		coreFileName = ""
		ctrlFileName = ""
		return
	}
	switch module {
	case Core:
		coreFileMd5 = md5String
	case Ctrl:
		ctrlFileMd5 = md5String
	}
	service.SaveOptLog(c.Request, "系统升级", "升级包上传")

	c.JSON(http.StatusOK, handle.Success(md5String))

}

func DeCompressTftpFile(name string) error {
	cmdStr := "tar -xzf /data/ota/" + name + " -C /recovery/tftp"
	out, err := cmd.ExecWithTimeOut(cmdStr, 10*time.Minute)
	if err != nil {
		logger.Error("tar file err:%v; out :%s", err, out)
		return err
	}

	cmdStr = "mv /recovery/tftp/tftp/* /recovery/tftp/"
	out, err = cmd.ExecWithTimeOut(cmdStr, 10*time.Minute)
	if err != nil {
		logger.Error("tar file err:%v; out :%s", err, out)
		return err
	}

	cmdStr = "chown -R linaro:linaro /recovery/tftp"
	out, err = cmd.ExecWithTimeOut(cmdStr, 1*time.Minute)
	if err != nil {
		logger.Error("tar file err:%v; out :%s", err, out)
		return err
	}

	if err := TftpFileCheck(); err != nil {
		logger.Error("%v", err)
	}
	return nil
}

func TftpFileCheck() error {
	// 读取 md5 文件内容
	md5File, err := os.Open("/recovery/tftp/md5.txt")
	if err != nil {
		logger.Error("Error opening md5 file:%v", err)
		return err
	}
	defer md5File.Close()

	md5Map := make(map[string]string)
	scanner := bufio.NewScanner(md5File)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 2 {
			md5Map[fields[1]] = fields[0]
		}
	}

	if err := scanner.Err(); err != nil {
		logger.Error("Error reading md5 file:", err)
		return err
	}

	for path, value := range md5Map {
		calculatedMD5, err := calculateFileMD5("/recovery/tftp/" + path)
		if err != nil {
			logger.Error("Error calculating MD5 for file")
			return err
		}

		// 检查 MD5 值是否一致
		if calculatedMD5 != value {
			return errors.New("md5 not ok")
		}
	}
	coreFileName = "tftp.tgz"
	coreFileMd5 = "ok"
	logger.Info("tftp file is ok")
	return nil
}

func (b *OtaApi) OtaFileList(c *gin.Context) {
	fileInfo := getFileName()
	c.JSON(http.StatusOK, handle.Success(fileInfo))

}

func (b *OtaApi) OtaUpgrate(c *gin.Context) {
	// 参数判断
	module := c.Request.FormValue("module")

	getFileName()

	cmdFlag := "--target="
	ip := c.Request.FormValue("ip")
	if len(ip) > 0 {
		cmdFlag = cmdFlag + ip
	} else {
		cmdFlag = cmdFlag + "all"
	}

	var otaName string
	switch module {
	case Ctrl:
		otaName = ctrlFileName
		cmdFlag = "/data/ota/local_update.sh md5.txt 1"
	case Core:
		otaName = coreFileName
	default:
		c.JSON(http.StatusOK, handle.Fail(buserr.UpgradeParamErr, "param error"))
		return
	}

	if otaName == "" {
		c.JSON(http.StatusOK, handle.Fail(-1, "升级文件未就绪"))
		return
	}

	if global.DeviceType == "" {
		logger.Error("设备类型获取异常,请刷新页面重试")
		c.JSON(http.StatusOK, handle.Fail(buserr.UpgradeParamErr, "设备类型获取异常"))
		return
	}

	otaInfo := dto.OtaVersion{
		Name:       strings.ToLower(global.DeviceType) + "_" + module + "_upgrade_" + time.Now().Format("20060102150405"),
		Product:    strings.ToUpper(global.DeviceType),
		FileName:   otaName,
		ModuleName: module,
		CmdFlag:    cmdFlag,
	}
	logger.Info("OTA info is :%v", otaInfo)

	_, err := service.OtaUpgrade(otaInfo)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "升级失败"))
		return
	}

	service.SaveOptLog(c.Request, "系统升级", "系统升级")

	c.JSON(http.StatusOK, handle.OkWithMsg("系统开始升级，预计需要5分钟"))
}

func (b *OtaApi) OtaUpgradeList(c *gin.Context) {

	result, err := service.OtaUpgradeList()
	handle.HandleError(err)
	var otaTasks []dto.OtaTask
	res, _ := json.Marshal(result.Result)
	_ = json.Unmarshal(res, &otaTasks)
	c.JSON(http.StatusOK, handle.Success(otaTasks))

}

type OtaFileInfo struct {
	CtrlName string `json:"ctrlName"`
	CtrlMd5  string `json:"ctrlMd5"`
	CoreName string `json:"coreName"`
	CoreMd5  string `json:"coreMd5"`
}

func calculateFileMD5(filePath string) (string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("无法打开文件: %v", err)
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		logger.Error("无法读取文件: %v", err)
		return "", err
	}

	hashInBytes := hash.Sum(nil)
	md5String := hex.EncodeToString(hashInBytes)

	return md5String, nil
}

func getFileName() OtaFileInfo {
	var fileInfo OtaFileInfo
	if ctrlFileName != "" {
		fileInfo.CtrlName = ctrlFileName
		fileInfo.CtrlMd5 = ctrlFileMd5
	}

	if coreFileName != "" {
		fileInfo.CoreName = coreFileName
		fileInfo.CoreMd5 = coreFileMd5
	}
	return fileInfo
}

func uploadFileHandler(c *gin.Context, dir string) (string, error) {
	// 从请求中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	// 保存文件
	err = c.SaveUploadedFile(file, dir+file.Filename)
	if err != nil {
		return "", err
	}

	return file.Filename, nil
}
