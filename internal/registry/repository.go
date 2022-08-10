package registry

import (
	"go-app/internal/domain"
	roleRepo "go-app/internal/modules/role/repository"
	userRepo "go-app/internal/modules/user/repository"

	"gorm.io/gorm"
)

// Repository registry
type Repository struct {
	RoleRepository domain.RoleRepository
	UserRepository domain.UserRepository
}

// NewRepository implements from interface
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		RoleRepository: roleRepo.NewRepository(db),
		UserRepository: userRepo.NewRepository(db),
	}
}
