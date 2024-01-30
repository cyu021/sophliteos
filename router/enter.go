package router

import (
	"sophliteos/router/app"
	"sophliteos/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	App    app.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
