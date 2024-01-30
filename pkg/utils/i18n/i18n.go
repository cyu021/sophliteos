package i18n

import (
	"fmt"
	"sophliteos/pkg/buserr"
	"sophliteos/pkg/utils/validation"
)

type Lang map[interface{}]string

var langMap map[string]Lang
var en Lang
var zh Lang

const En = "en"
const Zh = "zh_CN"

var dicts = []interface{}{
	buserr.Ok, "ok", "操作成功",
	buserr.Err, "server internal error", "服务器内部错误",
	buserr.NotImplemented, "not implemented", "不支持的请求",
	buserr.SetAlarmErr, "configure alarm error", "设置告警失败",
	buserr.SetIpErr, "configure ip error", "设置IP失败",
	buserr.SetDeviceInfoErr, "configure device info error", "设置设备信息失败",
	buserr.UpgradeParamErr, "upgrade request param error", "升级请求参数错误",
	buserr.UpgradeErr, "upgrade request error", "升级请求失败",
	buserr.UpgradeTaskNotFound, "task not found", "任务不存在",
	buserr.RollbackErr, "rollback request", "回滚请求失败",
	buserr.InvalidUsernameOrPassword, "invalid user or password", "用户名或密码错误",
	buserr.InvalidToken, "invalid token", "无效的token",
	buserr.PwdNotEqErr, "wrong password error", "原密码错误",
	buserr.PwdValidErr, "Low password strength, 8 to 12 characters in length, and it should contain at least 3 types of uppercase and " +
		"lowercase letters, numbers, and characteristic characters", "密码强度低，长度8到16位，包含大小写字母，数字，特征字符至少3种",
	validation.NotNil, "can't be nil", "不能为空",
	validation.UnknownFormat, "unknown format", "格式错误",
}

func init() {
	langMap = make(map[string]Lang)
	en := make(Lang)
	zh := make(Lang)
	langMap[En] = en
	langMap[Zh] = zh
	langMap["zh"] = zh
	size := len(dicts)
	for i := 0; i < size; i = i + 3 {
		en[dicts[i]] = dicts[i+1].(string)
		zh[dicts[i]] = dicts[i+2].(string)
	}
}

// 多语言获取
func GetString(lang string, code interface{}) string {
	if item, ok := langMap[lang]; ok {
		if val, ok := item[code]; ok {
			return val
		}
	}
	return fmt.Sprintf("%v", code)
}

// 多语言获取
func GetLang(lang string) Lang {
	switch lang {
	case En:
		return en
	default:
		return zh
	}
}

// 设置多语言
func SetString(lang string, code interface{}, value string) {
	if item, ok := langMap[lang]; ok {
		item[code] = value
	}
}
