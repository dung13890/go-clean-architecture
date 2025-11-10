package repository

import (
	"context"
	"go-app/internal/domain/entity"
	"go-app/internal/domain/repository"
	"go-app/pkg/errors"

	"gorm.io/gorm"
)

// userRepository ...
type userRepository struct {
	*gorm.DB
}

// NewUserRepository will implement of repository.UserRepository interface
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		DB: db,
	}
}

// Fetch will fetch content from db
func (rp *userRepository) Fetch(ctx context.Context) ([]entity.User, error) {
	dao := []User{}

	if err := rp.DB.WithContext(ctx).Find(&dao).Error; err != nil {
		return nil, errors.ErrUnexpectedDBError.Wrap(err)
	}
	users := []entity.User{}

	for i := range dao {
		u := convertUserToEntity(&dao[i])
		users = append(users, *u)
	}

	return users, nil
}

// Find will find content from db
func (rp *userRepository) Find(ctx context.Context, id uint) (*entity.User, error) {
	dao := User{}
	if err := rp.DB.WithContext(ctx).First(&dao, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound.Wrap(err)
		}
		return nil, errors.ErrUnexpectedDBError.Wrap(err)
	}

	return convertUserToEntity(&dao), nil
}

// Store will create data to db
func (rp *userRepository) Store(ctx context.Context, user *entity.User) error {
	dao := convertUserToDao(user)
	if err := rp.DB.WithContext(ctx).Create(&dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}
	*user = *convertUserToEntity(dao)

	return nil
}

// FindByQuery is a function that returns a users filtered by query
func (rp *userRepository) FindByQuery(ctx context.Context, q entity.User) (*entity.User, error) {
	dao := convertUserToDao(&q)
	if err := rp.DB.WithContext(ctx).First(&dao, &dao).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound.Wrap(err)
		}
		return nil, errors.ErrUnexpectedDBError.Wrap(err)
	}

	return convertUserToEntity(dao), nil
}

// CheckExists is a function that returns a users filtered by query
func (rp *userRepository) CheckExists(ctx context.Context, q entity.User, id *uint) (bool, error) {
	dao := convertUserToDao(&q)
	var exists bool
	subQuery := rp.DB.WithContext(ctx).
		Model(&User{}).
		Select("count(*) > 0").
		Where(&dao)
	if id != nil {
		subQuery = subQuery.Where("id <> ?", id)
	}
	if err := subQuery.Find(&exists).Error; err != nil {
		return false, errors.ErrUnexpectedDBError.Wrap(err)
	}

	return exists, nil
}

// Update will update data to db
func (rp *userRepository) Update(ctx context.Context, user *entity.User) error {
	dao := convertUserToDao(user)
	if err := rp.DB.WithContext(ctx).Save(&dao).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}
	*user = *convertUserToEntity(dao)

	return nil
}

// Delete will delete data from db
func (rp *userRepository) Delete(ctx context.Context, id uint) error {
	dao := User{}
	if err := rp.DB.WithContext(ctx).Delete(&dao, id).Error; err != nil {
		return errors.ErrUnexpectedDBError.Wrap(err)
	}

	return nil
}
