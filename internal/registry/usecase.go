package registry

import (
	authUC "go-app/internal/modules/auth/usecase"
)

// Usecase registry
type Usecase struct {
	AuthModule *authUC.Usecase
}

// NewUsecase implements from interface for modules
func NewUsecase(repo *Repository, svc *Service) *Usecase {
	return &Usecase{
		AuthModule: &authUC.Usecase{
			RoleUC: authUC.NewRoleUsecase(
				repo.AuthModule.RoleR,
			),
			UserUC: authUC.NewUserUsecase(
				repo.AuthModule.UserR,
			),
			AuthUC: authUC.NewAuthUsecase(
				svc.JWTSvc,
				svc.ThrottleSvc,
				repo.AuthModule.UserR,
				repo.AuthModule.PasswordR,
			),
		},
	}
}
