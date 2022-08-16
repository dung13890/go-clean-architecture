package usecase

import (
	"context"
	"errors"
	"go-app/internal/domain"
	iErrors "go-app/pkg/errors"
	"go-app/pkg/utils"
)

// AuthUsecase ...
type AuthUsecase struct {
	repo domain.UserRepository
}

var errInvalidatePass = errors.New("invalidate Password")

// Register is function used to register user
func (a AuthUsecase) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	passHash, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = passHash
	err = a.repo.Store(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login is function uses to log in
func (a AuthUsecase) Login(ctx context.Context, u *domain.User) (*domain.User, string, error) {
	tokenStr := ""

	user, err := a.repo.FindByQuery(ctx, domain.UserQueryParam{Email: u.Email})
	if err != nil {
		return nil, "", err
	}

	if !utils.ComparePassword(u.Password, user.Password) {
		return nil, "", iErrors.Wrap(errInvalidatePass)
	}

	tokenStr, err = utils.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, tokenStr, nil
}

// NewAuthUsecase will create new AuthUsecase object
func NewAuthUsecase(rp domain.UserRepository) domain.AuthUsecase {
	return &AuthUsecase{
		repo: rp,
	}
}
