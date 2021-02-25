package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"guoke-assistant-go/constant"
	"log"
	"os"
)

type AppConfig struct {
	Name 			string
	Hostname 		string
	Port 			string
	Release			bool
	Magic			string
}

type AdminConfig struct {
	Uid			int
	Account		string
	Pwd			string
}

type MySQLConfig struct {
	Host		string
	Database	string
	Username	string
	Password	string
	Port		string
	Charset		string
}

type RedisConfig struct {
	Host	string
	Port	string
	AUTH	string
}

type ESConfig struct {
	Address	string
}

type CosConfig struct {
	AppId		string
	SecretId	string
	SecretKey	string
	Bucket		string
	Region		string
}

type WeChatConfig struct {
	AppId		string
	AppSecret	string
}

type BotConfig struct {
	Feishu		string
	Dingding	string
}

type LogConfig struct {
	LogfilePath		string
	LogfileName		string
}

type Config struct {
	App		AppConfig
	Admin	AdminConfig
	Mysql	MySQLConfig
	Redis	RedisConfig
	ES		ESConfig
	Cos		CosConfig
	WeChat	WeChatConfig
	Bot		BotConfig
	Log		LogConfig
}

var (
	allConf		Config
	AppConf		*AppConfig
	AdminConf	*AdminConfig
	MysqlConf	*MySQLConfig
	RedisConf	*RedisConfig
	ESConf		*ESConfig
	CosConf		*CosConfig
	WeChatConf	*WeChatConfig
	BotConf		*BotConfig
	LogConf		*LogConfig
)

func init() {
	confFile := constant.ConfigFilePathDevelop
	for i := 0; i < 5; i++ {
		if _, err := os.Stat(confFile); err == nil {
			break
		} else {
			confFile = "../" + confFile
		}
	}
	InitConfig(confFile)
}

func InitConfig(fileName string) {
	viper.SetConfigFile(fileName)
	updateConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置发生改变：%s", e.Name)
		updateConfig()
	})
}

func updateConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("加载全局配置文件出错：%s\n", err)
	}
	if err := viper.Unmarshal(&allConf); err != nil {
		log.Printf("配置绑定出错！")
	}
	AppConf		= &allConf.App
	AdminConf	= &allConf.Admin
	MysqlConf	= &allConf.Mysql
	RedisConf	= &allConf.Redis
	ESConf		= &allConf.ES
	CosConf		= &allConf.Cos
	WeChatConf	= &allConf.WeChat
	BotConf		= &allConf.Bot
	LogConf		= &allConf.Log
}
