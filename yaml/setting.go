package setting

import (
	"github.com/spf13/viper"
)

/*
File name    : setting.go
Author       : miaoyc
Create date  : 2023/06/13 11:50
Description  : 配置信息读取
*/

type GlobalConfig struct {
	Auth    Auth    `yaml:"auth"`
	Data    Data    `yaml:"data"`
	Upgrade Upgrade `yaml:"upgrade"`
}

type Auth struct {
	Mid   string `yaml:"mid"`
	Token string `yaml:"token"`
}

type Data struct {
	DownloadPath string `yaml:"downloadPath"`
}

type Upgrade struct {
	HttpProxy string `yaml:"httpProxy"`
}

type FilePath struct {
	PackageVerPath string
	TmpPath        string
	EncryptPath    string
	DownloadPath   string
}

var (
	GlobalConf GlobalConfig
)

func getConf(configFile string) {
	vip := viper.New()
	vip.SetConfigType("yaml")
	vip.SetConfigFile(configFile)
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	err := vip.Unmarshal(&GlobalConf)
	if err != nil {
		panic(err)
	}
}

// Setup initialize the configuration instance
func Setup(configFile string) {
	getConf(configFile)
}
