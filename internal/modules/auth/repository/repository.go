package repository

import (
	"go-app/internal/domain"

	"gorm.io/gorm"
)

// Repository for module auth
type Repository struct {
	RoleR     domain.RoleRepository
	UserR     domain.UserRepository
	PasswordR domain.PasswordResetRepository
}

// NewRepository will create new postgres object
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		RoleR:     &RoleRepository{DB: db},
		UserR:     &UserRepository{DB: db},
		PasswordR: &PasswordResetRepository{DB: db},
	}
}
