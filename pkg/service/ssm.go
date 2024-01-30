package service

import (
	"encoding/json"
	"net/http"

	"sophliteos/config"
	"sophliteos/logger"
	"sophliteos/pkg/dto"
	http2 "sophliteos/pkg/utils/httpclient"

	"github.com/mitchellh/mapstructure"
)

const (
	ssmPrefix = "/bitmain/v1/ssm"

	ssmDevicePrefix = ssmPrefix + "/software/device"
	// SSM系统登录接口
	ssmSystemLoginUri = ssmPrefix + "/login"
	// SSM设备登录接口
	ssmDeviceLoginUri = ssmDevicePrefix + "/login"
	// 查询算力设备算力信息
	ssmDeviceResourceUri = ssmDevicePrefix + "/resource/list?all=0"
	// 查询算力设备基础信息
	ssmDeviceBasicUri = ssmDevicePrefix + "/basic"
	// IP查询
	ssmIpList = ssmPrefix + "/hardware/ip"
	// IP修改
	ssmIpSetUri = ssmPrefix + "/hardware/ip"
	// Name修改
	ssmBasicSetUri = ssmDevicePrefix + "/configure/basic"
	// 阈值告警
	ssmAlarmUri = ssmDevicePrefix + "/configure/alarm"

	ssmSubscribePrefix = ssmPrefix + "/software/notify"
	// 查询告警，订阅告警
	ssmAlarmSubscribeUri = ssmSubscribePrefix + "/subscribe"
	// 取消告警订阅
	ssmAlarmUnSubscribeUri = ssmSubscribePrefix + "/unsubscribe"
	// 回调地址
	ssmAlarmCallbackUri = "/api/device/alarm"
	// 上传ota
	ssmUploadUri = ssmPrefix + "/file/ota"
	// 升级Uri
	ssmUpgradeUri = ssmPrefix + "/workflow/upgrade"
	// 回滚uri
	ssmRollbackUri = ssmPrefix + "/workflow/rollback"
	// 关机
	ssmCoreBoardShutdownUri = ssmPrefix + "/hardware/devices/down"
	// 重启
	ssmCoreBoardRebootUri = ssmPrefix + "/hardware/devices/reset"
	// scp文件传输
	ssmCoreBoardScpUri = ssmPrefix + "/hardware/devices/scp"
	// 执行命令
	ssmCoreBoardExecUri = ssmPrefix + "/hardware/devices/exec"
	// 执行命令
	ssmIpTablesUri = ssmPrefix + "/hardware/nat"
	// 配置修改
	ssmBasicModUri = ssmDevicePrefix + "/configure/serviceAddress"
)

func getUrlHeader(uri string) (string, map[string]string, error) {

	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	server := v.GetString("ssm.server")

	conf.Unlock()

	header := make(map[string]string)

	return server + uri, header, nil
}

// 系统登录
func SystemLogin(ssmServer, username, password string) (string, error) {
	login, _ := json.Marshal(dto.LoginRequest{
		UserName: username,
		Password: password,
	})
	result, err := NewSsmRequestWithHeaders(ssmServer+ssmSystemLoginUri, http.MethodPost, map[string]string{}, login)
	if err != nil {
		return "", err
	}
	var res dto.SystemLoginResponse
	_ = mapstructure.Decode(result.Result, &res)
	return res.Token, nil
}

// 查询控制板算力信息
func GetCtrlResource() (dto.CtrlResource, error) {
	var err error
	result, err := NewSsmRequest(ssmDeviceResourceUri, http.MethodGet, []byte{})
	if err != nil {
		return dto.CtrlResource{}, err
	}
	var resource dto.CtrlResource
	// logger.Info("%v", result.Result.([]interface{})[0].(map[string]interface{}))
	jsonData, _ := json.Marshal(result.Result.([]interface{})[0])

	if err = json.Unmarshal(jsonData, &resource); err != nil {
		err = mapstructure.Decode(result.Result, &resource)
	}
	return resource, err
}

// 查询算力控制板基础信息
func GetCtrlBasic() (dto.CtrlBasic, string, error) {
	var ctrl dto.CtrlBasic
	var res dto.SsmResult
	result, err := NewSsmRequest(ssmDeviceBasicUri, http.MethodGet, []byte{})
	if err != nil {
		return ctrl, "", err
	}
	err = mapstructure.Decode(result.Result, &ctrl)
	mapstructure.Decode(result, &res)
	return ctrl, res.DeviceSn, err
}

// ipTables
func GetIpTables() ([]string, error) {
	var list []string
	result, err := NewSsmRequest(ssmIpTablesUri, http.MethodGet, []byte{})
	err = mapstructure.Decode(result.Result, &list)
	return list, err
}

// ipTables增加
func AddIpTable(table dto.AddTable) (dto.SsmResult, error) {
	data, _ := json.Marshal(table)
	return NewSsmRequest(ssmIpTablesUri, http.MethodPost, data)
}

// ipTables删除
func DeleteIpTable(num string) (dto.SsmResult, error) {
	return NewSsmRequest(ssmIpTablesUri+"/PREROUTING-"+num, http.MethodDelete, []byte{})
}

// IP修改
func SetIP(ip dto.IPSettings) (dto.SsmResult, error) {
	data, _ := json.Marshal(ip)
	return NewSsmRequest(ssmIpSetUri, http.MethodPost, data)
}

// 基本信息修改
func SetBasic(name dto.BasicSettings) (dto.SsmResult, error) {
	data, _ := json.Marshal(name)
	return NewSsmRequest(ssmBasicSetUri, http.MethodPost, data)
}

// 修改basic
func ModBasic(data []byte) (dto.SsmResult, error) {
	return NewSsmRequest(ssmBasicModUri, http.MethodPost, data)
}

// IP查询
func GetIP() (dto.SsmResult, error) {
	return NewSsmRequest(ssmIpList, http.MethodGet, []byte{})
}

// 告警设置
func SetAlarm(alarmThreshold []byte) ([]dto.SsmResult, error) {
	var result []dto.SsmResult
	// 设置核心板
	err := NewRequest(ssmAlarmUri+"?devices=cores", http.MethodPost, alarmThreshold, &result)
	if err != nil {
		return nil, err
	}
	// 设置控制板
	var cr dto.SsmResult
	err = NewRequest(ssmAlarmUri, http.MethodPost, alarmThreshold, &cr)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 告警设置
func SetAlarmSe5(alarmThreshold []byte) (dto.SsmResult, error) {
	var result dto.SsmResult
	// 设置核心板
	err := NewRequest(ssmAlarmUri+"?devices=cores", http.MethodPost, alarmThreshold, &result)
	if err != nil {
		return dto.SsmResult{}, err
	}
	// 设置控制板
	var cr dto.SsmResult
	err = NewRequest(ssmAlarmUri, http.MethodPost, alarmThreshold, &cr)
	if err != nil {
		return dto.SsmResult{}, err
	}
	return result, nil
}

// 查询订阅
func GetAlarm() (dto.SsmResult, error) {
	return NewSsmRequest(ssmAlarmSubscribeUri+"/"+config.Conf.GetName(), http.MethodGet, []byte{})
}

// 订阅通知
func SubscribeAlarm() (dto.SsmResult, error) {
	subscribe := getSubscribeAlarm()
	return NewSsmRequest(ssmAlarmSubscribeUri, http.MethodPost, subscribe)
}

// 全局订阅对象
func getSubscribeAlarm() []byte {
	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	port := v.GetString("server.port")
	conf.Unlock()
	data, _ := json.Marshal(&dto.AlarmSubscribe{
		Platform:            config.Conf.GetName(),
		SubscribeDetailType: []int{1, 2},
		NotificationURL:     "http://127.0.0.1:" + port + ssmAlarmCallbackUri,
	})
	return data
}

// 取消订阅
func UnSubscribeAlarm() (dto.SsmResult, error) {
	return NewSsmRequest(ssmAlarmUnSubscribeUri, http.MethodPost, getSubscribeAlarm())
}

// 上传OTA升级包
func OtaUpload(otaFile dto.OtaFile) (dto.SsmResult, error) {
	logger.Info("上传%s文件", otaFile.Module)
	result, err := NewSsmMultiFileRequest(ssmUploadUri, http.MethodPost,
		map[string]string{
			"module": otaFile.Module,
		}, otaFile.Filename, otaFile.Body)
	return result, err
}

// 传输文件到核心板
func ScpFile(upgrade dto.Scp2Core) (dto.SsmResult, error) {
	data, _ := json.Marshal(upgrade)
	return NewSsmRequest(ssmCoreBoardScpUri, http.MethodPost, data)
}

// 核心板执行命令或脚本
func ExecCore(exec dto.ExecCorex) (dto.SsmResult, error) {
	data, _ := json.Marshal(exec)
	return NewSsmRequest(ssmCoreBoardExecUri, http.MethodPost, data)
}

// 查看升级文件
func OtaUploadList() (dto.SsmResult, error) {
	return NewSsmRequest(ssmUploadUri, http.MethodGet, []byte{})
}

// OTA升级
func OtaUpgrade(upgrade dto.OtaVersion) (dto.SsmResult, error) {
	data, _ := json.Marshal(upgrade)
	return NewSsmRequest(ssmUpgradeUri, http.MethodPost, data)
}

// OTA升级列表
func OtaUpgradeList() (dto.SsmResult, error) {
	return NewSsmRequest(ssmUpgradeUri, http.MethodGet, []byte{})
}

// OTA升级回滚
func OtaRollback(upgrade dto.OtaVersion) (dto.SsmResult, error) {
	data, _ := json.Marshal(upgrade)
	return NewSsmRequest(ssmRollbackUri, http.MethodPost, data)
}

// 核心板关机重启操作
func CoreOperate(number, operationType int) (dto.SsmResult, error) {
	uri := ""
	if operationType == dto.Reboot {
		uri = ssmCoreBoardRebootUri
	} else if operationType == dto.Shutdown {
		uri = ssmCoreBoardShutdownUri
	} else {
		return dto.SsmResult{
			ErrorCode:    0,
			ErrorMessage: "",
		}, nil
	}
	data, _ := json.Marshal(dto.CoreOpe{
		Id: number,
	})
	return NewSsmRequest(uri, http.MethodPost, data)
}

// 请求验证
func NewSsmRequest(uri string, method string, bytes []byte) (dto.SsmResult, error) {
	url, header, err := getUrlHeader(uri)
	if err != nil {
		var res dto.SsmResult
		return res, err
	}
	return NewSsmRequestWithHeaders(url, method, header, bytes)
}

// 请求验证
func NewSsmMultiFileRequest(uri string, method string, params map[string]string, filename string, bytes []byte) (dto.SsmResult, error) {
	var result dto.SsmResult
	url, header, err := getUrlHeader(uri)
	if err != nil {
		return result, err
	}
	data, err := http2.NewMultiFileRequest(url, method, header, params, filename, bytes)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func NewSsmRequestWithHeaders(url string, method string, header map[string]string, bytes []byte) (dto.SsmResult, error) {
	var result dto.SsmResult
	err := NewRequestWithHeaders(url, method, header, bytes, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// 请求验证
func NewRequestWithHeaders(url string, method string, header map[string]string, bytes []byte, v interface{}) error {
	// 远程调用
	data, err := http2.NewRequest(url, method, header, bytes)
	logger.Debug("url %s method %s header %v request %v",
		url, method, header, string(bytes))
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func NewRequest(uri string, method string, bytes []byte, v interface{}) error {
	// 远程调用
	url, header, err := getUrlHeader(uri)
	if err != nil {
		return nil
	}
	return NewRequestWithHeaders(url, method, header, bytes, v)
}

func SsmCoreRequest(host, uri string, method string, bytes []byte) (dto.SsmResult, error) {
	url := host + ":9779" + uri
	// logger.Info(url)
	return NewSsmRequestWithHeaders(url, method, nil, bytes)
}
