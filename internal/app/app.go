package app

import (
	"context"
	"go-app/config"
	"go-app/internal/constants"
	"go-app/internal/registry"
	"go-app/pkg/cache"
	"go-app/pkg/database"
	"go-app/pkg/errors"
	"go-app/pkg/logger"
	"go-app/pkg/redis"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Run Application
func Run(conf config.AppConfig) error {
	dbConf := config.GetDBConfig()
	var db *gorm.DB
	var err error
	for {
		db, err = database.NewGormDB(dbConf)
		if err != nil {
			logger.Info().Printf("Wait for starting db: %v", err)
			time.Sleep(constants.ConnectWaitDuration)
		} else {
			break
		}
	}

	rdb := redis.New(config.GetRedisConfig())
	cacheManager := cache.NewRedisStore(rdb)

	e := echo.New()

	repo := registry.NewRepository(db)
	service := registry.NewService(cacheManager)
	usecase := registry.NewUsecase(repo, service)
	registry.NewHTTPHandler(e, usecase, service)

	s := &http.Server{
		Handler:     e,
		Addr:        conf.AppHost,
		ReadTimeout: constants.ConnectReadTimeout,
	}

	go func() {
		logger.Info().Printf(
			"Start http server: %v, location: %v",
			conf.AppHost,
			time.Now().Location().String(),
		)
		if err := s.ListenAndServe(); err != nil {
			logger.Error().Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	logger.Info().Printf("Signal: %d, received", <-quit)
	ctx, cancel := context.WithTimeout(context.Background(), constants.ConnectTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	return nil
}
