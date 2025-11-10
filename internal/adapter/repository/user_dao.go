package repository

import (
	"go-app/internal/domain/entity"
	"go-app/pkg/errors"
	"go-app/pkg/utils"

	"gorm.io/gorm"
)

// User DAO model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleID   uint   `json:"role_id"`
	Password string `json:"Password"`
}

// BeforeSave hooks
func (dao *User) BeforeSave(_ *gorm.DB) error {
	hashPW, err := utils.GeneratePassword(dao.Password)
	if err != nil {
		return errors.ErrBadGateway.Wrap(err)
	}
	dao.Password = hashPW

	return nil
}

// convertUserToEntity .-
func convertUserToEntity(dao *User) *entity.User {
	e := &entity.User{
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

// convertUserToDao .-
func convertUserToDao(entity *entity.User) *User {
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
