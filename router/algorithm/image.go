package algorithm

import (
	v1 "algoliteos/api/v1"

	"github.com/gin-gonic/gin"
)

type ImageRouter struct{}

func (s *ImageRouter) InitImageRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	router := Router.Group("algorithm")
	api := v1.ApiGroupApp.AlgoGroup.ImageApi
	{
		router.GET("image", api.GetImage)
		router.GET("vedio", api.GetVedio)
	}

	return router
}
