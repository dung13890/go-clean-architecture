package repository

import (
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"go-app/pkg/utils"

	"gorm.io/gorm"
)

// User DAO model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleID   int    `json:"role_id"`
	Password string `json:"Password"`
}

// BeforeSave hooks
func (dao *User) BeforeSave(tx *gorm.DB) error {
	hashPW, err := utils.GeneratePassword(dao.Password)
	if err != nil {
		return errors.Wrap(err)
	}
	dao.Password = hashPW

	return nil
}

// convertToEntity .-
func convertToEntity(dao *User) *domain.User {
	e := &domain.User{
		ID:        dao.ID,
		Name:      dao.Name,
		Email:     dao.Email,
		RoleID:    dao.RoleID,
		Password:  dao.Password,
		CreatedAt: dao.CreatedAt,
		UpdatedAt: dao.UpdatedAt,
	}

	return e
}

// convertToDao .-
func convertToDao(entity *domain.User) *User {
	d := &User{
		Model: gorm.Model{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
		Name:     entity.Name,
		Email:    entity.Email,
		RoleID:   entity.RoleID,
		Password: entity.Password,
	}

	return d
}
