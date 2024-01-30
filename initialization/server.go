package initialization

import (
	"algoliteos/global"
	"algoliteos/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func InitServer(router *gin.Engine) server {
	logger.Info("Starting HTTP service at %s", global.SystemConf.AlgoPort)

	return &http.Server{
		Addr:         ":" + global.SystemConf.AlgoPort,
		Handler:      router,
		ReadTimeout:  global.OtaTimeOut,
		WriteTimeout: global.OtaTimeOut,
	}
}
