package main

import (
	"go-app/config"
	"go-app/pkg/logger"
	"go-app/pkg/postgres"
)

func main() {
	logger.InitLogger()
	if err := config.LoadConfig(); err != nil {
		logger.Error().Fatal(err)
	}

	db := config.GetDBConfig()

	if err := postgres.Migrate(db); err != nil {
		logger.Error().Fatal(err)
	}
}
