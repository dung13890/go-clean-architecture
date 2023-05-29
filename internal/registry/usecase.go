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
		AuthModule: authUC.NewUsecase(
			repo.AuthModule,
			svc.JWTSvc,
			svc.ThrottleSvc,
		),
	}
}
