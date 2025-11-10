package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	httpHD "go-app/internal/delivery/http"
	"go-app/internal/infrastructure/config"
	"go-app/internal/infrastructure/constant"
	"go-app/internal/infrastructure/database"
	"go-app/internal/infrastructure/redis"
	"go-app/internal/registry"
	"go-app/pkg/errors"
	"go-app/pkg/logger"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

	if err := run(conf); err != nil {
		logger.Error(err)
	}
}

// run Application
func run(conf config.AppConfig) error {
	dbConf := config.GetDBConfig()
	var db *gorm.DB
	var err error
	for {
		db, err = database.NewGormDB(dbConf)
		if err != nil {
			logger.Infof("Wait for starting db: %v", err)
			time.Sleep(constant.ConnectWaitDuration)
		} else {
			break
		}
	}

	rdb := redis.New(config.GetRedisConfig())
	e := echo.New()

	reg := registry.NewRegistry(db, rdb)
	httpHD.NewHTTPHandler(e, reg.JWTSvc, reg)

	s := &http.Server{
		Handler:     e,
		Addr:        conf.AppHost,
		ReadTimeout: constant.ConnectReadTimeout,
	}

	go func() {
		logger.Infof(
			"Start http server: %v, location: %v",
			conf.AppHost,
			time.Now().Location().String(),
		)
		if err := s.ListenAndServe(); err != nil {
			logger.Error(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	logger.Infof("Signal: %d, received", <-quit)
	ctx, cancel := context.WithTimeout(context.Background(), constant.ConnectTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	return nil
}
