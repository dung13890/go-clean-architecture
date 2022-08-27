package config

import (
	"go-app/pkg/logger"

	"github.com/spf13/viper"
)

// Database config struct
type Database struct {
	Connection string `mapstructure:"DB_CONNECTION"`
	Host       string `mapstructure:"DB_HOST"`
	Port       int    `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_DBNAME"`
	User       string `mapstructure:"DB_USER"`
	Password   string `mapstructure:"DB_PASSWORD"`
	SSLMode    string `mapstructure:"DB_SSLMODE" default:"disable"`
	Debug      bool   `mapstructure:"DB_DEBUG" default:"false"`
}

// GetDBConfig Unmarshal Database Config from env
func GetDBConfig() Database {
	c := Database{}
	if err := viper.Unmarshal(&c); err != nil {
		logger.Error().Fatal(err)
	}

	return c
}
