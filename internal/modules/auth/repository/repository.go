package repository

import (
	"go-app/internal/domain"
)

// Repository for module auth
type Repository struct {
	RoleR     domain.RoleRepository
	UserR     domain.UserRepository
	PasswordR domain.PasswordResetRepository
}
