package system

import "sophliteos/pkg/service"

type ApiGroup struct {
	BaseApi
	ResourceApi
	BasicApi
	PasswordApi
	IpApi
	AlarmApi
	LogApi
	OtaApi
	VersionApi
	UpgradeApi
	SsmUpgradeApi
	DownApi
}

var hostService = new(service.HostService)
