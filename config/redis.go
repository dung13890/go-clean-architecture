package config

import (
	"sync"

	"go-app/pkg/logger"

	"github.com/spf13/viper"
)

var (
	onceRedis sync.Once
	redisConf Redis
)

// Redis config struct
type Redis struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     int    `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
}

// GetRedisConfig Unmarshal Redis Config from env
func GetRedisConfig() Redis {
	onceRedis.Do(func() {
		if err := viper.Unmarshal(&redisConf); err != nil {
			logger.Error().Fatal(err)
		}
	})

	return redisConf
}
