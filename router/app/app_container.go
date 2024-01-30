package app

import (
	v1 "sophliteos/api/v1"

	"github.com/gin-gonic/gin"
)

type ContainerRouter struct{}

func (s *ContainerRouter) InitContainerRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	containerRouter := Router.Group("api/container")
	baseApi := v1.ApiGroupApp.AppGroup.ContainerApi
	{
		containerRouter.GET("/image/list", baseApi.ListImage)
		containerRouter.GET("/image/listpulling", baseApi.ListPulling)
		containerRouter.POST("/image/pull", baseApi.PullImage)
		containerRouter.POST("/image/delete", baseApi.RemoveImage)
		containerRouter.POST("/image/checkPull", baseApi.CheckImagePulled)

		containerRouter.GET("/list", baseApi.ListContainer)
		containerRouter.POST("/delete", baseApi.RemoveContainer)
		containerRouter.POST("/stop", baseApi.StopContainer)
		containerRouter.POST("/start", baseApi.StartContainer)
		containerRouter.POST("/inspect", baseApi.InspectContainer)

		containerRouter.GET("/insecure/list", baseApi.ListRegistry)
		containerRouter.POST("/insecure/add", baseApi.AddRegistries)
		containerRouter.POST("/insecure/delete", baseApi.RemoveRegistries)
		containerRouter.POST("/insecure/imagetaglist", baseApi.ListRegistryImageTags)
		containerRouter.POST("/insecure/imagelist", baseApi.ListRegistryImages)
		containerRouter.POST("/insecure/taglist", baseApi.ListRegistryTags)
	}

	return containerRouter
}

// SE6核心板docker接口
func (s *ContainerRouter) InitCoreContainerRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	containerRouter := Router.Group("api/container/core")
	baseApi := v1.ApiGroupApp.AppGroup.CoreContainerApi
	{
		containerRouter.GET("/image/list", baseApi.ListImage)

		containerRouter.POST("/image/delete", baseApi.RemoveImage)

		containerRouter.GET("/list", baseApi.ListContainer)
		containerRouter.POST("/delete", baseApi.RemoveContainer)
		containerRouter.POST("/stop", baseApi.StopContainer)
		containerRouter.POST("/start", baseApi.StartContainer)
		containerRouter.POST("/inspect", baseApi.InspectContainer)

		containerRouter.GET("/insecure/list", baseApi.ListRegistry)
		containerRouter.POST("/insecure/add", baseApi.AddRegistries)
		containerRouter.POST("/insecure/delete", baseApi.RemoveRegistries)
		containerRouter.POST("/insecure/imagetaglist", baseApi.ListRegistryImageTags)
		containerRouter.POST("/insecure/imagelist", baseApi.ListRegistryImages)
		containerRouter.POST("/insecure/taglist", baseApi.ListRegistryTags)
	}

	return containerRouter
}
