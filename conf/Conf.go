package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	//日志
	log "github.com/sirupsen/logrus"
	//处理各种类型的配置需求和格式
	"github.com/spf13/viper"
	"os"
)

type MysqlConfig struct {
	Host     string `yaml:"Host"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
}
type RedisConfig struct {
	RedisAddr string `yaml:"RedisAddr"`
	RedisPass string `yaml:"RedisPass"`
	DB        int    `yaml:"DB"`
	PoolSize  int    `yaml:"PoolSize"`
}
type AliConfig struct {
	Endpoint        string `yaml:"Endpoint" `
	AccessKeyId     string `yaml:"AccessKeyId"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
}

var MysqlConf MysqlConfig
var RedisConf RedisConfig
var AliConf AliConfig

//初始化优先加载
func init() {
	appName := "App"
	//设置要读取的配置文件目录
	viper.AddConfigPath("./resource")
	configEnv := os.Getenv("GO_ENV")
	//默认读取App.yml
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
	err := viper.Unmarshal(&MysqlConf)
	if err != nil {
		log.Panicf("unable to decode into mysql struct,%v\n", err)
	}
	err = viper.Unmarshal(&RedisConf)
	if err != nil {
		log.Panicf("unable to decode into redis struct,%v\n", err)
	}
	err = viper.Unmarshal(&AliConf)
	if err != nil {
		log.Panicf("unable to decode into aliyun struct,%v\n", err)
	}
	//实时读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("config file changed:", e.Name)
	})
}
