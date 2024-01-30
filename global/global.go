package global

import (
	"sophliteos/pkg/dto"
	"sophliteos/pkg/dto/response"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

var (
	TimeOut          time.Duration
	OtaTimeOut       time.Duration
	Version          dto.BuildInfo
	BlockAllRequests bool
	DeviceType       string
	SSmLists         response.SsmList
	SdkVersion       string
	AlgoFlag         bool
	SE6Cores         []SE6Core

	Arch   string
	VALID  *validator.Validate
	System SystemConf
	DB     *gorm.DB
)

type SystemConf struct {
	Port           string `mapstructure:"port"`
	SSL            string `mapstructure:"ssl"`
	DbFile         string `mapstructure:"db_file"`
	DbPath         string `mapstructure:"db_path"`
	LogPath        string `mapstructure:"log_path"`
	DataDir        string `mapstructure:"data_dir"`
	TmpDir         string `mapstructure:"tmp_dir"`
	Cache          string `mapstructure:"cache"`
	Backup         string `mapstructure:"backup"`
	EncryptKey     string `mapstructure:"encrypt_key"`
	BaseDir        string `mapstructure:"base_dir"`
	Mode           string `mapstructure:"mode"`
	RepoUrl        string `mapstructure:"repo_url"`
	Version        string `mapstructure:"version"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Entrance       string `mapstructure:"entrance"`
	IsDemo         bool   `mapstructure:"is_demo"`
	AppRepo        string `mapstructure:"app_repo"`
	ChangeUserInfo bool   `mapstructure:"change_user_info"`
	OneDriveID     string `mapstructure:"one_drive_id"`
	OneDriveSc     string `mapstructure:"one_drive_sc"`
}

type SE6Core struct {
	Ip       string `json:"ip"`
	DeviceSn string `json:"deviceSn"`
	Number   int    `json:"number"`
}
