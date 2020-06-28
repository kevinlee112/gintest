package config

import (
	"github.com/spf13/viper"
	"time"
)

// 应用程序配置
type appConfig struct {
	// 应用名称
	Name string
	// 运行模式: debug, release, test
	RunMode string
	// 运行 addr
	Addr string

	// 超时时间
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	//日志文件路径
	LogDir string
}

func initAppConfig() *appConfig {
	viper.SetDefault("App.NAME", "gin_api")
	viper.SetDefault("APP.RUNMODE", "release")
	viper.SetDefault("APP.ADDR", ":8080")
	viper.SetDefault("APP.ADDR", ":8080")
	viper.SetDefault("APP.ADDR", ":8080")
	viper.SetDefault("APP.READTIMEOUT", 120)
	viper.SetDefault("APP.WRITETIMEOUT", 120)
	viper.SetDefault("APP.LOGDIR", "log/" + viper.GetString("APP.NAME"))

	return &appConfig{
		Name:viper.GetString("APP.NAME"),
		RunMode: viper.GetString("APP.RUNMODE"),
		Addr:    viper.GetString("APP.ADDR"),
		ReadTimeout: viper.GetDuration("APP.READTIMEOUT"),
		WriteTimeout:    viper.GetDuration("APP.WRITETIMEOUT"),
		LogDir:    viper.GetString("APP.LOGDIR"),
	}
}
