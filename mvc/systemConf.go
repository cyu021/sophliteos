package mvc

type SystemConf struct {
	AlgoPort      string         `mapstructure:"algoport"`
	Log           LogConfig      `mapstructure:"log"`
	Db            DbConfig       `mapstructure:"db"`
	MediaHost     string         `mapstructure:"mediahost"`
	Picture       PictureConfig  `mapstructure:"picture"`
	AlgorithmHost string         `mapstructure:"algorithmhost"`
	UploadHost    string         `mapstructure:"uploadhost"`
	LocalHost     string         `mapstructure:"localhost"`
	Abilities     map[int]string `mapstructure:"abilities" yaml:"abilities"`
}

type LogConfig struct {
	Path  string `mapstructure:"path"`
	Level string `mapstructure:"level"`
}

type DbConfig struct {
	Path     string `mapstructure:"path"`
	SaveDays int    `mapstructure:"savedays"`
}

type PictureConfig struct {
	Dir     string `mapstructure:"dir"`
	MaxSize int64  `mapstructure:"maxsize"`
	Quality int64  `mapstructure:"quality"`
}
