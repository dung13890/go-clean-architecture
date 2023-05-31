package config

import (
	"sync"

	"go-app/pkg/logger"

	"github.com/spf13/viper"
)

var (
	onceEmail sync.Once
	mailConf  Email
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
	onceEmail.Do(func() {
		if err := viper.Unmarshal(&mailConf); err != nil {
			logger.Error().Fatal(err)
		}
	})

	return mailConf
}
