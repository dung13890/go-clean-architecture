package repository

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"go-app/pkg/utils"

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
func (rp *RoleRepository) Fetch(c context.Context) ([]domain.Role, error) {
	roles := []domain.Role{}
	if err := rp.DB.Find(&roles).Error; err != nil {
		return roles, errors.Wrap(err)
	}

	return roles, nil
}

// Find will find content from db
func (rp *RoleRepository) Find(c context.Context, id int) (*domain.Role, error) {
	role := domain.Role{}
	if err := rp.DB.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return &role, nil
}

// Store will create data to db
func (rp *RoleRepository) Store(c context.Context, role *domain.Role) error {
	role.Slug = utils.Slugify(role.Name)
	if err := rp.DB.Create(role).Error; err != nil {
		return errors.Wrap(err)
	}

	return nil
}
