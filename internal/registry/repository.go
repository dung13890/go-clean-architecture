package registry

import (
	authRepo "go-app/internal/modules/auth/repository"

	"gorm.io/gorm"
)

// Repository registry
type Repository struct {
	AuthModule *authRepo.Repository
}

// NewRepository implements from interface for modules
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		AuthModule: &authRepo.Repository{
			RoleR:     authRepo.NewRoleRepository(db),
			UserR:     authRepo.NewUserRepository(db),
			PasswordR: authRepo.NewPasswordResetRepository(db),
		},
	}
}
