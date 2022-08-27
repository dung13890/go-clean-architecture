package app

import (
	"context"
	"go-app/config"
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"go-app/pkg/postgres"

	"github.com/spf13/viper"

	roleRepository "go-app/internal/modules/role/repository"
	userRepository "go-app/internal/modules/user/repository"
)

var pathJSON = "cmd/seed/data.json"

type seedData struct {
	Roles []domain.Role `json:"roles"`
	Users []domain.User `json:"users"`
}

// Seed is function that seed data
func Seed(dbConfig config.Database) error {
	db, err := postgres.NewGormDB(dbConfig)
	if err != nil {
		return errors.Wrap(err)
	}

	userRepo := userRepository.NewRepository(db)
	roleRepo := roleRepository.NewRepository(db)

	viper.SetConfigFile(pathJSON)
	err = viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err)
	}

	data := seedData{}
	err = viper.Unmarshal(&data)
	if err != nil {
		return errors.Wrap(err)
	}

	for _, r := range data.Roles {
		err = roleRepo.Store(context.Background(), &r)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	for _, u := range data.Users {
		err = userRepo.Store(context.Background(), &u)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	return nil
}
