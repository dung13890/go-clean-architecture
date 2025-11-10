package auth

import (
	"go-app/internal/domain/repository"
	"go-app/internal/domain/service"
)

var (
	tokenLength = 10
)

// Usecase ...
type Usecase struct {
	jwtSvc      service.JWTService
	throttleSvc service.ThrottleService
	mailSvc     service.MailService
	repo        repository.UserRepository
	pwRepo      repository.PasswordResetRepository
}

// NewUsecase will create new an userUsecase object representation of domain.Usecase interface
func NewUsecase(
	jwtSvc service.JWTService,
	throttleSvc service.ThrottleService,
	mailSvc service.MailService,
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
