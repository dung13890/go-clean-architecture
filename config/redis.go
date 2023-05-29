package config

import (
	"go-app/pkg/logger"

	"github.com/spf13/viper"
)

// Redis config struct
type Redis struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     int    `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
}

// GetRedisConfig Unmarshal Redis Config from env
func GetRedisConfig() Redis {
	c := Redis{}
	if err := viper.Unmarshal(&c); err != nil {
		logger.Error().Fatal(err)
	}

	return c
}
