package app

import (
	"context"

	"go-app/config"
	"go-app/internal/domain"
	authModule "go-app/internal/modules/auth/repository"
	"go-app/pkg/database"
	"go-app/pkg/errors"

	"github.com/spf13/viper"
)

var pathJSON = "db/seeds/data.json"

type seedData struct {
	Roles []domain.Role `json:"roles"`
	Users []domain.User `json:"users"`
}

// Seed is function that seed data
func Seed(dbConfig config.Database) error {
	db, err := database.NewGormDB(dbConfig)
	if err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	authMD := &authModule.Repository{
		RoleR:     authModule.NewRoleRepository(db),
		UserR:     authModule.NewUserRepository(db),
		PasswordR: authModule.NewPasswordResetRepository(db),
	}

	viper.SetConfigFile(pathJSON)
	if err = viper.ReadInConfig(); err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	data := seedData{}
	if err := viper.Unmarshal(&data); err != nil {
		return errors.ErrInternalServerError.Wrap(err)
	}

	if err := data.seedAuth(authMD); err != nil {
		return errors.Throw(err)
	}

	return nil
}

// seedAuth is function that seed auth data
func (seed *seedData) seedAuth(md *authModule.Repository) error {
	for i := range seed.Roles {
		if err := md.RoleR.Store(context.Background(), &seed.Roles[i]); err != nil {
			return errors.ErrInternalServerError.Wrap(err)
		}
	}

	for j := range seed.Users {
		if err := md.UserR.Store(context.Background(), &seed.Users[j]); err != nil {
			return errors.ErrInternalServerError.Wrap(err)
		}
	}

	return nil
}
