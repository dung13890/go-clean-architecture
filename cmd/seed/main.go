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

	dbConfig := config.GetDBConfig()
	if err := postgres.Seed(dbConfig); err != nil {
		logger.Error().Fatal(err)
	}
}
