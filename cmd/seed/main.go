package main

import (
	"go-app/config"
	"go-app/internal/app"
	"go-app/pkg/logger"
	"time"
)

func main() {
	logger.InitLogger()
	if err := config.LoadConfig(); err != nil {
		logger.Error().Fatal(err)
	}

	conf := config.GetAppConfig()

	// Set timezone
	loc, err := time.LoadLocation(conf.AppTimeZone)
	if err != nil {
		logger.Error().Fatal(err)
	}
	time.Local = loc

	dbConfig := config.GetDBConfig()
	if err := app.Seed(dbConfig); err != nil {
		logger.Error().Fatal(err)
	}
}
