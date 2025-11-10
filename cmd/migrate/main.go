package main

import (
	"go-app/internal/infrastructure/config"
	"go-app/internal/infrastructure/database"
	"go-app/pkg/logger"
	"time"
)

func main() {
	logger.Init()
	if err := config.LoadConfig(); err != nil {
		logger.Error(err)
	}

	conf := config.GetAppConfig()

	// Set timezone
	loc, err := time.LoadLocation(conf.AppTimeZone)
	if err != nil {
		logger.Error(err)
	}
	time.Local = loc

	db := config.GetDBConfig()

	if err := database.Migrate(db); err != nil {
		logger.Error(err)
	}
}
