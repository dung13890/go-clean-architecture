package config

import (
	"go-app/pkg/logger"

	"github.com/spf13/viper"
)

// Email config struct
type Email struct {
	Host     string `mapstructure:"MAIL_HOST"`
	Port     int    `mapstructure:"MAIL_PORT"`
	Username string `mapstructure:"MAIL_USERNAME"`
	Password string `mapstructure:"MAIL_PASSWORD"`
	From     string `mapstructure:"MAIL_FROM_ADDRESS"`
}

// GetEmailConfig Unmarshal Email Config from env
func GetEmailConfig() Email {
	c := Email{}
	if err := viper.Unmarshal(&c); err != nil {
		logger.Error().Fatal(err)
	}

	return c
}
