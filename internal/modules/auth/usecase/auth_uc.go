package usecase

import (
	"context"
	"errors"
	"go-app/config"
	"go-app/internal/constants"
	"go-app/internal/domain"
	pkgErrors "go-app/pkg/errors"
	"go-app/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// AuthUsecase ...
type AuthUsecase struct {
	repo domain.UserRepository
}

var errInvalidatePass = errors.New("invalidate Password")

// NewUsecase will create new AuthUsecase object
func NewUsecase(rp domain.UserRepository) domain.AuthUsecase {
	return &AuthUsecase{
		repo: rp,
	}
}

// Register is function used to register user
func (uc AuthUsecase) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := uc.repo.Store(ctx, user); err != nil {
		return nil, pkgErrors.Wrap(err)
	}

	return user, nil
}

// Login is function uses to log in
func (uc AuthUsecase) Login(ctx context.Context, u *domain.User) (*domain.Claims, string, error) {
	userByEmail := domain.User{Email: u.Email}
	user, err := uc.repo.FindByQuery(ctx, userByEmail)
	if err != nil {
		return nil, "", pkgErrors.Wrap(err)
	}

	if !utils.ComparePassword(u.Password, user.Password) {
		return nil, "", pkgErrors.Wrap(errInvalidatePass)
	}

	// Create token and store to claims
	now := time.Now()
	expirationTime := now.Add(constants.TokenLifetime)

	// JwtKey is secret key fow singed
	jwtKey := []byte(config.GetAppConfig().AppJWTKey)

	claims := &domain.Claims{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return nil, "", pkgErrors.Wrap(err)
	}

	return claims, tokenString, nil
}
