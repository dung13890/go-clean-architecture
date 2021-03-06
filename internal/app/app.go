package app

import (
	"context"
	"go-app/config"
	"go-app/internal/factory"
	"go-app/pkg/errors"
	"go-app/pkg/logger"
	"go-app/pkg/postgres"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	connectTimeout      = time.Second * 10
	connectWaitDuration = time.Second * 5
	connectReadTimeout  = time.Second * 30
)

// Run Application
func Run(conf config.AppConfig) error {
	dbConf := config.GetDBConfig()
	var db *gorm.DB
	var err error
	for {
		db, err = postgres.NewGormDB(dbConf)
		if err != nil {
			logger.Info().Printf("Wait for starting db: %v", err)
			time.Sleep(connectWaitDuration)
		} else {
			break
		}
	}

	e := echo.New()
	factory.NewFactory(e, db)

	s := &http.Server{
		Handler:     e,
		Addr:        conf.AppHost,
		ReadTimeout: connectReadTimeout,
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
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return errors.Wrap(err)
	}

	return nil
}
