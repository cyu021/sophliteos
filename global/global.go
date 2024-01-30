package global

import (
	"algoliteos/mvc"
	"time"
)

var (
	TimeOut          time.Duration
	OtaTimeOut       time.Duration
	BlockAllRequests bool
	SystemConf       mvc.SystemConf
)
