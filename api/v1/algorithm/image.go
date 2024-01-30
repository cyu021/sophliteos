package algorithm

import (
	"algoliteos/global"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ImageApi struct{}

func (b *ImageApi) GetImage(c *gin.Context) {
	taskId := c.Query("taskId")
	event := c.Query("type")
	fileName := c.Query("fileName")

	var file *os.File
	var err error
	// 打开图片文件
	if event == "" {
		file, err = os.Open(global.SystemConf.Picture.Dir + "/" + taskId + "/" + fileName)
	} else {
		file, err = os.Open(global.SystemConf.Picture.Dir + "/" + taskId + "/" + event + "/" + fileName)
	}

	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to open image")
		return
	}
	defer file.Close()

	// 设置响应头部信息
	c.Header("Content-Type", "image/jpeg")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to send image")
		return
	}
}

func (b *ImageApi) GetVedio(c *gin.Context) {
	pathVedio := c.Query("path")
	c.File(pathVedio)
}
