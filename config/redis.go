package config

import (
	"github.com/spf13/viper"
	"time"
)

//redis 配置
type redisConfig struct {
	Host string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}

func initRedisConfig() *redisConfig {
	viper.SetDefault("REDIS.HOST", "127.0.0.1:6379")
	viper.SetDefault("REDIS.PASSWORD", "")
	viper.SetDefault("REDIS.MAX_IDLE", 30)
	viper.SetDefault("REDIS.MAX_ACTIVE", 30)
	viper.SetDefault("REDIS.IDLE_TIMEOUT", 200)

	return &redisConfig{
		Host:viper.GetString("REDIS.HOST"),
		Password: viper.GetString("REDIS.PASSWORD"),
		MaxIdle:    viper.GetInt("REDIS.MAX_IDLE"),
		MaxActive: viper.GetInt("REDIS.MAX_ACTIVE"),
		IdleTimeout:    viper.GetDuration("REDIS.IDLE_TIMEOUT"),
	}
}
