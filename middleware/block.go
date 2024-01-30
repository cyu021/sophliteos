package middleware

import (
	"net/http"
	"sophliteos/global"
	"sophliteos/pkg/buserr"
	"sophliteos/pkg/handle"

	"github.com/gin-gonic/gin"
)

func BlockerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.BlockAllRequests {
			c.JSON(http.StatusServiceUnavailable, handle.FailWithMsg(buserr.Upgradeing, "服务器升级中，暂不可用"))
			// c.File("/var/lib/sophliteos/dist/updating.html")
			c.Abort()
		}
	}
}
