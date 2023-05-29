package usecase

import (
	"go-app/internal/domain"
	"go-app/internal/modules/auth/repository"
)

// Usecase for module auth
type Usecase struct {
	RoleUC domain.RoleUsecase
	UserUC domain.UserUsecase
	AuthUC domain.AuthUsecase
}

// NewUsecase implements from interface
func NewUsecase(
	repo *repository.Repository,
	jwtSvc domain.JWTService,
	thrSvc domain.ThrottleService,
) *Usecase {
	return &Usecase{
		RoleUC: NewRoleUsecase(repo.RoleR),
		UserUC: NewUserUsecase(repo.UserR),
		AuthUC: NewAuthUsecase(jwtSvc, thrSvc, repo.UserR, repo.PasswordR),
	}
}
