package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	configFilePath = "./config.yaml"
	configFileType = "yml"
)

var (
	// AppConfig 应用配置
	AppConfig *appConfig
	// DBConfig 数据库配置
	DBConfig *dbConfig
	// Redis 配置
	RedisConfig *redisConfig
)

func InitConfig()  {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err!= nil {
		panic(fmt.Sprintf("读取配置文件失败，请检查: %v", err))
	}

	// 初始化 app 配置
	AppConfig = initAppConfig()
	// 初始化数据库配置
	DBConfig = initDBConfig()
	// 初始化 redis 配置
	RedisConfig = initRedisConfig()

	// 热更新配置文件
	watchConfig()
}

// 监控配置文件变化
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(ev fsnotify.Event) {
		// 配置文件更新了
		fmt.Printf("\n\n --------------- Config file changed: %s --------------- \n\n", ev.Name)
	})
}