package app

import (
	v1 "sophliteos/api/v1"

	"github.com/gin-gonic/gin"
)

type TerminalRouter struct{}

func (s *TerminalRouter) InitTerminalRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	terminalRouter := Router.Group("api/terminals")
	baseApi := v1.ApiGroupApp.AppGroup.TerminalApi
	{

		terminalRouter.GET("/cores/list", baseApi.GetSE6Core)

		terminalRouter.POST("test", baseApi.TestByInfo)

		terminalRouter.GET("/command", baseApi.ListCommand)
		terminalRouter.POST("/command", baseApi.CreateCommand)
		terminalRouter.POST("/command/del", baseApi.DeleteCommand)
		terminalRouter.POST("/command/update", baseApi.UpdateCommand)

		terminalRouter.POST("host/add", baseApi.CreateHost)
		terminalRouter.GET("host/list", baseApi.GetHosts)
		terminalRouter.POST("host/del", baseApi.DeleteHost)
		terminalRouter.POST("host/update", baseApi.UpdateHost)
	}
	return terminalRouter
}

func (s *TerminalRouter) InitSshRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	terminalRouter := Router.Group("api/terminals")
	baseApi := v1.ApiGroupApp.AppGroup.TerminalApi
	{
		terminalRouter.GET("docker/ssh", baseApi.ContainerWsSsh)
		terminalRouter.GET("ssh", baseApi.WsSsh)
	}
	return terminalRouter
}
