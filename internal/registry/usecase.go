package registry

import (
	"go-app/internal/domain"
	roleUC "go-app/internal/modules/role/usecase"
	userUC "go-app/internal/modules/user/usecase"
)

// Usecase registry
type Usecase struct {
	RoleUsecase domain.RoleUsecase
	UserUsecase domain.UserUsecase
}

// NewUsecase implements from interface
func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{
		RoleUsecase: roleUC.NewUsecase(repo.RoleRepository),
		UserUsecase: userUC.NewUsecase(repo.UserRepository),
	}
}
