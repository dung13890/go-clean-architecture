package service

import (
	"context"
	"time"

	"go-app/internal/adapter/cache"
	"go-app/internal/domain/entity"
	"go-app/internal/domain/service"
	"go-app/internal/infrastructure/config"
	"go-app/internal/infrastructure/constant"
	"go-app/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

// jWTService is a struct that represent the jwt's service
type jWTService struct {
	cm cache.Client
}

// NewJWTService will create new an jwtService object representation of service.JWTService interface
func NewJWTService(cm cache.Client) service.JWTService {
	return &jWTService{
		cm: cm,
	}
}

// GenerateToken is a function to generate the jwt token
func (*jWTService) GenerateToken(_ context.Context, user *entity.User) (string, int64, error) {
	// Create token and store to claims
	now := time.Now()
	expirationTime := now.Add(constant.TokenLifetime)

	// JwtKey is secret key fow singed
	jwtKey := []byte(config.GetAppConfig().AppJWTKey)
	exp := jwt.NewNumericDate(expirationTime)

	// Generate token
	cls := &entity.Claims{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, cls).SignedString(jwtKey)
	if err != nil {
		return "", int64(0), errors.ErrBadRequest.Wrap(err)
	}

	return token, exp.Unix(), nil
}

// Invalidate is a function to invalidate the jwt token
func (svc *jWTService) Invalidate(ctx context.Context, token any) error {
	claims, err := convertToClaims(token)
	if err != nil {
		return errors.Throw(err)
	}

	if err := svc.cm.Set(ctx, claims.RegisteredClaims.ID, true, constant.TokenLifetime); err != nil {
		return errors.Throw(err)
	}

	return nil
}

// Decode is a function to convert the jwt token to user
func (svc *jWTService) Decode(ctx context.Context, token any) (*entity.User, error) {
	claims, err := convertToClaims(token)
	if err != nil {
		return nil, errors.Throw(err)
	}

	// Check JWT token invalid
	if _, err := svc.cm.Get(ctx, claims.RegisteredClaims.ID); err == nil {
		return nil, errors.ErrJWTRevoke.Trace()
	}

	user := &entity.User{
		ID:     claims.ID,
		Name:   claims.Name,
		Email:  claims.Email,
		RoleID: claims.RoleID,
	}

	return user, nil
}

// convertToClaims is a function to convert the token to claims
func convertToClaims(token any) (*entity.Claims, error) {
	tk, ok := token.(*jwt.Token)
	if !ok {
		return nil, errors.ErrJWTInvalidCredentials.Trace()
	}
	claims, ok := tk.Claims.(*entity.Claims)
	if !ok {
		return nil, errors.ErrJWTInvalidClaims.Trace()
	}

	return claims, nil
}
