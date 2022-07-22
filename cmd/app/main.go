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

	conf := config.GetAppConfig()

	if err := app.Run(conf); err != nil {
		logger.Error().Fatal(err)
	}
}
