package algorithm

import (
	v1 "algoliteos/api/v1"
	"algoliteos/global"
	"algoliteos/middleware"

	"github.com/gin-gonic/gin"
)

type TestAlarmRouter struct{}

func (s *TestAlarmRouter) InitTestAlarmRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	router := Router.Group("algorithm", middleware.TimeoutMiddleware(global.TimeOut))
	api := v1.ApiGroupApp.AlgoGroup.TestApi
	{
		router.POST("upload/test", api.TestAlarm)

	}

	return router
}
