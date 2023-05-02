package repository

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"

	"gorm.io/gorm"
)

// RoleRepository ...
type RoleRepository struct {
	*gorm.DB
}

// NewRepository will create new postgres object
func NewRepository(db *gorm.DB) domain.RoleRepository {
	return &RoleRepository{DB: db}
}

// Fetch will fetch content from db
func (rp *RoleRepository) Fetch(ctx context.Context) ([]domain.Role, error) {
	dao := []Role{}
	if err := rp.DB.WithContext(ctx).Find(&dao).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	roles := []domain.Role{}

	for i := range dao {
		r := convertToEntity(&dao[i])
		roles = append(roles, *r)
	}

	return roles, nil
}

// Find will find content from db
func (rp *RoleRepository) Find(ctx context.Context, id int) (*domain.Role, error) {
	dao := Role{}
	if err := rp.DB.WithContext(ctx).First(&dao, id).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return convertToEntity(&dao), nil
}

// Store will create data to db
func (rp *RoleRepository) Store(ctx context.Context, role *domain.Role) error {
	dao := convertToDao(role)
	if err := rp.DB.WithContext(ctx).Create(&dao).Error; err != nil {
		return errors.Wrap(err)
	}

	*role = *convertToEntity(dao)

	return nil
}
