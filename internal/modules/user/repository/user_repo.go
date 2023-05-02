package repository

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"

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
func (rp *UserRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	dao := []User{}

	if err := rp.DB.WithContext(ctx).Find(&dao).Error; err != nil {
		return nil, errors.Wrap(err)
	}
	users := []domain.User{}

	for i := range dao {
		u := convertToEntity(&dao[i])
		users = append(users, *u)
	}

	return users, nil
}

// Find will find content from db
func (rp *UserRepository) Find(ctx context.Context, id int) (*domain.User, error) {
	dao := User{}
	if err := rp.DB.WithContext(ctx).First(&dao, id).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return convertToEntity(&dao), nil
}

// Store will create data to db
func (rp *UserRepository) Store(ctx context.Context, user *domain.User) error {
	dao := convertToDao(user)
	if err := rp.DB.WithContext(ctx).Create(&dao).Error; err != nil {
		return errors.Wrap(err)
	}
	*user = *convertToEntity(dao)

	return nil
}

// FindByQuery is a function that returns a users filtered by query
func (rp *UserRepository) FindByQuery(ctx context.Context, q domain.User) (*domain.User, error) {
	dao := convertToDao(&q)
	if err := rp.DB.WithContext(ctx).First(&dao, &dao).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return convertToEntity(dao), nil
}
