package main

import (
	"context"
	"time"

	"go-app/internal/adapter/repository"
	"go-app/internal/domain/entity"
	"go-app/internal/infrastructure/config"
	"go-app/internal/infrastructure/database"
	"go-app/pkg/errors"
	"go-app/pkg/logger"

	"github.com/spf13/viper"
)

var pathJSON = "db/seeds/data.json"

type seedData struct {
	Roles []entity.Role `json:"roles"`
	Users []entity.User `json:"users"`
}

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

	dbConfig := config.GetDBConfig()
	if err := seed(dbConfig); err != nil {
		logger.Error(err)
	}
}

// seed is function that seed data
func seed(dbConfig config.Database) error {
	db, err := database.NewGormDB(dbConfig)
	if err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	// Registry Repository
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	viper.SetConfigFile(pathJSON)
	if err = viper.ReadInConfig(); err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	data := seedData{}
	if err := viper.Unmarshal(&data); err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}
	// Seed userRepo
	for i := range data.Roles {
		if err := roleRepo.Store(context.Background(), &data.Roles[i]); err != nil {
			return errors.ErrInternalServerError.Wrap(err)
		}
	}

	for j := range data.Users {
		if err := userRepo.Store(context.Background(), &data.Users[j]); err != nil {
			return errors.ErrInternalServerError.Wrap(err)
		}
	}

	return nil
}
