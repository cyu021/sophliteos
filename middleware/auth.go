package middleware

import (
	"net/http"
	"sophliteos/logger"
	"sophliteos/pkg/buserr"
	"sophliteos/pkg/repo"
	"sophliteos/pkg/service"
	"sophliteos/pkg/utils/i18n"

	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_token := service.Token(c.Request)
		if _token != "" {
			user := service.GetUser(_token)
			if user != nil {
				now := time.Now()
				if now.Before(user.ExpireTime) {
					if user.ExpireTime.After(user.ExpireTime.Add(time.Minute * 10)) {
						user.ExpireTime = now.Add(time.Hour * 2)
						repo.UpdateUser(user)
					}
				}
				if service.IsMultiPartRequest(c.Request) {
					err := c.Request.ParseMultipartForm(32 << 20)
					if err != nil {
						logger.Error("multipart/form-data请求读取失败：%s %s", c.Request.RequestURI, err.Error())
						c.AbortWithStatusJSON(http.StatusInternalServerError, i18n.GetString(service.GetLang(c.Request), buserr.Err))
						return
					}
				}
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, i18n.GetString(service.GetLang(c.Request), buserr.InvalidToken))
	}
}
