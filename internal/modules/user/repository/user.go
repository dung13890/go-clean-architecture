package repository

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"time"

	"gorm.io/gorm"
)

// UserRepository ...
type UserRepository struct {
	*gorm.DB
}

// NewRepository will create new postgres object
func NewRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{DB: db}
}

// Fetch will fetch content from db
func (rp *UserRepository) Fetch(c context.Context) ([]domain.User, error) {
	users := []domain.User{}
	if err := rp.DB.Debug().Find(&users).Error; err != nil {
		return users, errors.Wrap(err)
	}

	return users, nil
}

// Find will find content from db
func (rp *UserRepository) Find(c context.Context, id int) (*domain.User, error) {
	user := domain.User{}
	if err := rp.DB.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return &user, nil
}

// Store will create data to db
func (rp *UserRepository) Store(c context.Context, user *domain.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := rp.DB.Debug().Create(user).Error; err != nil {
		return errors.Wrap(err)
	}

	return nil
}
