package repository

import (
	"go-app/internal/domain"
	"go-app/pkg/utils"

	"gorm.io/gorm"
)

// Role DAO model
type Role struct {
	gorm.Model
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// BeforeSave hooks
func (dao *Role) BeforeSave(tx *gorm.DB) error {
	dao.Slug = utils.Slugify(dao.Name)

	return nil
}

// convertToEntity .-
func convertToEntity(dao *Role) *domain.Role {
	e := &domain.Role{
		ID:        dao.ID,
		Name:      dao.Name,
		Slug:      dao.Slug,
		CreatedAt: dao.CreatedAt,
		UpdatedAt: dao.UpdatedAt,
	}

	return e
}

// convertToDao .-
func convertToDao(entity *domain.Role) *Role {
	d := &Role{
		Model: gorm.Model{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
		Name: entity.Name,
		Slug: entity.Slug,
	}

	return d
}
