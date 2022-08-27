package repository

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"go-app/pkg/utils"

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
	if err := rp.DB.Find(&users).Error; err != nil {
		return users, errors.Wrap(err)
	}

	return users, nil
}

// Find will find content from db
func (rp *UserRepository) Find(c context.Context, id int) (*domain.User, error) {
	user := domain.User{}
	if err := rp.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return &user, nil
}

// Store will create data to db
func (rp *UserRepository) Store(c context.Context, user *domain.User) error {
	passHash, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return errors.Wrap(err)
	}
	user.Password = passHash
	if err := rp.DB.Create(user).Error; err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// Search is a function that returns a list of users filtered by query
func (rp *UserRepository) Search(ctx context.Context, q domain.UserQueryParam) ([]domain.User, error) {
	var users []domain.User
	if err := rp.DB.Where("email = ? ", q.Email).Find(&users).Error; err != nil {
		return users, errors.Wrap(err)
	}

	return users, nil
}

// FindByQuery is a function that returns a users filtered by query
func (rp *UserRepository) FindByQuery(ctx context.Context, q domain.UserQueryParam) (*domain.User, error) {
	user := &domain.User{}
	if err := rp.DB.Where("email = ? ", q.Email).First(&user).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return user, nil
}
