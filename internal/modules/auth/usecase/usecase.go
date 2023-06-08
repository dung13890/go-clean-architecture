package usecase

import (
	"go-app/internal/domain"
)

// Usecase for module auth
type Usecase struct {
	RoleUC domain.RoleUsecase
	UserUC domain.UserUsecase
	AuthUC domain.AuthUsecase
}
