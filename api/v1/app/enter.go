package app

import "sophliteos/pkg/service"

type ApiGroup struct {
	FileApi
	ContainerApi
	TerminalApi
	CoreContainerApi
}

var fileService = new(service.FileService)
var hostService = new(service.HostService)
var containerService = new(service.ContainerService)

func init() {
	containerService.Init()
}
