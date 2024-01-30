package app

import (
	"net/http"
	"sophliteos/logger"
	"sophliteos/pkg/dto/request"
	containerRes "sophliteos/pkg/dto/response"
	"sophliteos/pkg/handle"
	"time"

	"github.com/gin-gonic/gin"
)

type ContainerApi struct{}

func (b *ContainerApi) ListImage(c *gin.Context) {
	res, err := containerService.ListImages()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Success(res))

}

func (b *ContainerApi) ListPulling(c *gin.Context) {
	res, err := containerService.ListPulling()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Success(res))

}

func (b *ContainerApi) PullImage(c *gin.Context) {
	var req request.ImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := containerService.PullImage(req.Image); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, err.Error()))
		return
	}

	c.JSON(http.StatusOK, handle.OkWithMsg("镜像正在拉取，请点击镜像拉取列表查看"))
}

func (b *ContainerApi) CheckImagePulled(c *gin.Context) {
	var req request.ImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(struct {
		Code int `json:"status"`
	}{
		Code: containerService.CheckImagePulled(req.Image),
	}))
}

func (b *ContainerApi) InspectContainer(c *gin.Context) {
	var req request.ContainerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if res, err := containerService.InspectContainer(req.Name); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, err.Error()))
	} else {
		c.JSON(http.StatusOK, handle.Success(res))
	}
}

func (b *ContainerApi) RemoveImage(c *gin.Context) {
	var req request.ImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := containerService.RemoveImage(req.Image); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, err.Error()))
		return
	}

	logger.Info("删除镜像：%s", req.Image)

	c.JSON(http.StatusOK, handle.Ok())
}

func (b *ContainerApi) AddRegistries(c *gin.Context) {
	var req request.SecurityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := containerService.AddInsecureRegistries(req.InsecureRegistry); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, err.Error()))
		return
	}

	logger.Info("添加仓库：%s", req.InsecureRegistry)

	c.JSON(http.StatusOK, handle.Ok())
}

func (b *ContainerApi) RemoveRegistries(c *gin.Context) {
	var req request.SecurityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := containerService.RemoveInsecureRegistries(req.InsecureRegistry); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设置失败"))
		return
	}

	logger.Info("删除仓库：%s", req.InsecureRegistry)

	c.JSON(http.StatusOK, handle.Ok())
}

func (b *ContainerApi) ListRegistry(c *gin.Context) {
	insecures, err := containerService.ListInsecureRegistries()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}
	var res []containerRes.RegistryType

	for _, ins := range insecures {
		res = append(res, containerRes.RegistryType{Addr: ins})
	}

	c.JSON(http.StatusOK, handle.Success(res))
}

func (b *ContainerApi) StopContainer(c *gin.Context) {
	var req request.ContainerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := containerService.StopContainer(req.Name, 30*time.Second); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, err.Error()))
		return
	}

	logger.Info("停止容器：%s", req.Name)

	c.JSON(http.StatusOK, handle.Ok())
}

func (b *ContainerApi) StartContainer(c *gin.Context) {
	var req request.ContainerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := containerService.StartContainer(req.Name, 30*time.Second); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, err.Error()))
		return
	}

	logger.Info("启动容器：%s", req.Name)

	c.JSON(http.StatusOK, handle.Ok())
}

func (b *ContainerApi) RemoveContainer(c *gin.Context) {
	var req request.ContainerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := containerService.RemoveContainer(req.Name); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, err.Error()))
		return
	}

	logger.Info("删除容器：%s", req.Name)

	c.JSON(http.StatusOK, handle.Ok())
}

func (b *ContainerApi) ListContainer(c *gin.Context) {
	res, err := containerService.ListContainers()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Success(res))

}

func (b *ContainerApi) ListRegistryImageTags(c *gin.Context) {
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

func (b *ContainerApi) ListRegistryImages(c *gin.Context) {
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

func (b *ContainerApi) ListRegistryTags(c *gin.Context) {
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
