package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sophliteos/logger"
	"sophliteos/pkg/dto/request"
	"sophliteos/pkg/dto/response"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/service"

	"github.com/gin-gonic/gin"
)

const (
	ssmPrefix = "/bitmain/v1/ssm"
	// 容器列表
	listContain = ssmPrefix + "/docker/container"

	// 镜像
	Image = ssmPrefix + "/docker/image"

	// 仓库
	Registries = ssmPrefix + "/docker/conf/insecure_registry"
)

type CoreContainerApi struct{}

// 镜像列表
func (b *CoreContainerApi) ListImage(c *gin.Context) {
	ip := c.Query("ip")

	res, err := service.SsmCoreRequest(ip, Image, "GET", nil)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}

	var images []response.CoreImageInfo
	jsonData, _ := json.Marshal(res.Result)

	if err = json.Unmarshal(jsonData, &images); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}

	c.JSON(http.StatusOK, handle.Success(images))

}

// 删除镜像
func (b *CoreContainerApi) RemoveImage(c *gin.Context) {

	var req request.CoreImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	_, err := service.SsmCoreRequest(req.Ip, Image+"/"+req.Image, "DELETE", nil)
	if err != nil {
		logger.Error("请求失败:" + err.Error())
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Ok())
}

func (b *CoreContainerApi) InspectContainer(c *gin.Context) {
	ip := c.Query("ip")
	var req request.CoreContainerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if res, err := service.SsmCoreRequest(ip, listContain+"/"+req.Name, "GET", nil); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
	} else {
		var containerInfo []response.CoreInspectJson
		jsonData, _ := json.Marshal(res.Result)

		if err = json.Unmarshal(jsonData, &containerInfo); err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
			return
		}

		c.JSON(http.StatusOK, handle.Success(containerInfo))
	}

}

// 添加私仓
func (b *CoreContainerApi) AddRegistries(c *gin.Context) {
	var req request.CoreSecurityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	body := fmt.Sprintf(`{"registry": "%s"}`, req.InsecureRegistry)

	res, err := service.SsmCoreRequest(req.Ip, Registries, "POST", []byte(body))
	if err != nil {
		logger.Error("请求失败:" + err.Error())
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"))
		return
	}

	if res.Code != 0 {
		logger.Error("请求失败:" + err.Error())
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+res.Msg))
		return
	}

	c.JSON(http.StatusOK, handle.Ok())
}

// 删除私仓
func (b *CoreContainerApi) RemoveRegistries(c *gin.Context) {
	var req request.CoreSecurityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	body := fmt.Sprintf(`{"registry": "%s"}`, req.InsecureRegistry)

	res, err := service.SsmCoreRequest(req.Ip, Registries, "DELETE", []byte(body))
	if err != nil {
		logger.Error("请求失败:" + err.Error())
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"))
		return
	}

	if res.Code != 0 {
		logger.Error("请求失败:" + err.Error())
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+res.Msg))
		return
	}

	c.JSON(http.StatusOK, handle.Ok())
}

// 私有仓库列表
func (b *CoreContainerApi) ListRegistry(c *gin.Context) {

	ip := c.Query("ip")
	res, err := service.SsmCoreRequest(ip, Registries, "GET", nil)
	if err != nil {
		logger.Error("请求失败:" + err.Error())
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}

	var registries []string
	jsonData, _ := json.Marshal(res.Result)

	if err = json.Unmarshal(jsonData, &registries); err != nil {
		logger.Error("请求失败:" + err.Error())
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}

	c.JSON(http.StatusOK, handle.Success(registries))
}

// 停止容器
func (b *CoreContainerApi) StopContainer(c *gin.Context) {
	var req request.CoreContainerOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	req.Operation = "stop"

	if postData, err := json.Marshal(req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	} else {
		if res, err := service.SsmCoreRequest(req.Ip, listContain, "PATCH", postData); err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "内部错误"))
			return
		} else {
			if res.Code == 0 {
				c.JSON(http.StatusOK, handle.Ok())
				return
			} else {
				c.JSON(http.StatusOK, handle.FailWithMsg(-1, "停止失败"))
				return
			}
		}
	}
}

// 启动容器
func (b *CoreContainerApi) StartContainer(c *gin.Context) {
	var req request.CoreContainerOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	req.Operation = "start"

	if postData, err := json.Marshal(req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	} else {
		if res, err := service.SsmCoreRequest(req.Ip, listContain, "PATCH", postData); err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "启动失败"))
			return
		} else {
			if res.Code == 0 {
				c.JSON(http.StatusOK, handle.Ok())
				return
			} else {
				c.JSON(http.StatusOK, handle.FailWithMsg(-1, "启动失败"))
				return
			}
		}
	}
}

// 删除容器
func (b *CoreContainerApi) RemoveContainer(c *gin.Context) {

	var req request.CoreContainerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if postData, err := json.Marshal(req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	} else {
		if res, err := service.SsmCoreRequest(req.Ip, listContain, "DELETE", postData); err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "内部错误"))
			return
		} else {
			if res.Code == 0 {
				c.JSON(http.StatusOK, handle.Ok())
				return
			} else {
				c.JSON(http.StatusOK, handle.FailWithMsg(-1, "启动失败"))
				return
			}
		}
	}
}

// 容器列表
func (b *CoreContainerApi) ListContainer(c *gin.Context) {
	ip := c.Query("ip")

	res, err := service.SsmCoreRequest(ip, listContain, "GET", nil)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}

	var images []response.CoreListContainerJson
	jsonData, _ := json.Marshal(res.Result)

	if err = json.Unmarshal(jsonData, &images); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}

	c.JSON(http.StatusOK, handle.Success(images))
}

func (b *CoreContainerApi) ListRegistryImageTags(c *gin.Context) {
	var req request.SecurityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	res, err := containerService.ListInsecureImageTags(req.InsecureRegistry)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}
	c.JSON(http.StatusOK, handle.Success(res))
}

func (b *CoreContainerApi) ListRegistryImages(c *gin.Context) {
	var req request.SecurityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	res, err := containerService.ListInsecureImages(req.InsecureRegistry)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}
	c.JSON(http.StatusOK, handle.Success(res))
}

func (b *CoreContainerApi) ListRegistryTags(c *gin.Context) {
	var req request.SecurityImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	res, err := containerService.ListInsecureTags(req.InsecureRegistry, req.Image)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}
	c.JSON(http.StatusOK, handle.Success(res))
}
