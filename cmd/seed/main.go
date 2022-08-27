package main

import (
	"go-app/config"
	"go-app/internal/app"
	"go-app/pkg/logger"
)

func main() {
	logger.InitLogger()
	if err := config.LoadConfig(); err != nil {
		logger.Error().Fatal(err)
	}

	dbConfig := config.GetDBConfig()
	if err := app.Seed(dbConfig); err != nil {
		logger.Error().Fatal(err)
	}
}
