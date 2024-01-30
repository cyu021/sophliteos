package initialization

import (
	"algoliteos/config"
	"algoliteos/database"
	"algoliteos/global"
	"algoliteos/logger"
	"algoliteos/mvc"
	"time"
)

func InitBase() {
	// 加载配置
	config.LoadConfig()

	// 日志处理
	logger.InitLogging(global.SystemConf.Log.Path, "algo.log", global.SystemConf.Log.Level)

	// 初始化sqlite
	database.InitDB()

	global.TimeOut, _ = time.ParseDuration("30s")
	global.OtaTimeOut, _ = time.ParseDuration("3m")
	global.BlockAllRequests = false

	logger.Info(mvc.StructPrint(global.SystemConf))

}
