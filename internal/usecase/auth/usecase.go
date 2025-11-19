package auth

import (
	"go-app/internal/domain/gateway"
	"go-app/internal/domain/repository"
)

var (
	tokenLength = 10
)

// Usecase ...
type Usecase struct {
	jwtSvc      gateway.JWTService
	throttleSvc gateway.ThrottleService
	mailSvc     gateway.MailService
	repo        repository.UserRepository
	pwRepo      repository.PasswordResetRepository
}

// NewUsecase will create new an userUsecase object representation of domain.Usecase interface
func NewUsecase(
	jwtSvc gateway.JWTService,
	throttleSvc gateway.ThrottleService,
	mailSvc gateway.MailService,
	repo repository.UserRepository,
	pwRepo repository.PasswordResetRepository,
) *Usecase {
	return &Usecase{
		jwtSvc:      jwtSvc,
		throttleSvc: throttleSvc,
		mailSvc:     mailSvc,
		repo:        repo,
		pwRepo:      pwRepo,
	}
}
