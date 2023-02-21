package conf

import (
	//日志
	log "github.com/sirupsen/logrus"
	//处理各种类型的配置需求和格式
	"github.com/spf13/viper"
	"os"
)

type DbConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
}

type Struct struct {
	Mysql DbConfig
}

var Config Struct

//初始化优先加载
func init() {
	appName := "app"
	//设置要读取的配置文件目录
	viper.AddConfigPath("./resource")
	configEnv := os.Getenv("GO_ENV")
	//默认读取app.yml
	if configEnv != "" {
		appName += "-" + configEnv
	}
	//设置配置文件名
	viper.SetConfigName(appName)
	//读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file %s\n", err)
	}
	//读取配置，反序列化到结构体
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Panicf("unable to decode into struct,%v\n", err)
	}
}
