package v1

import (
	"sophliteos/api/v1/app"
	"sophliteos/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	AppGroup       app.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
