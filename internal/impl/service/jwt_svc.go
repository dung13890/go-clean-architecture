package service

import (
	"context"
	"go-app/config"
	"go-app/internal/constants"
	"go-app/internal/domain"
	"go-app/pkg/cache"
	"go-app/pkg/errors"
	"go-app/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTService is a struct that represent the jwt's service
type JWTService struct {
	cm cache.Client
}

// Claims is struct claims for jwt
type Claims struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID int    `json:"role_id"`
	jwt.RegisteredClaims
}

// NewJWTService will create new an jwtService object representation of domain.JWTService interface
func NewJWTService(cm cache.Client) *JWTService {
	return &JWTService{
		cm: cm,
	}
}

// GenerateToken is a function to generate the jwt token
func (_ *JWTService) GenerateToken(_ context.Context, user *domain.User) (string, int64, error) {
	// Create token and store to claims
	now := time.Now()
	expirationTime := now.Add(constants.TokenLifetime)

	// JwtKey is secret key fow singed
	jwtKey := []byte(config.GetAppConfig().AppJWTKey)
	exp := jwt.NewNumericDate(expirationTime)

	// Generate token
	cls := &Claims{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: exp,
			ID:        utils.GenerateUUID(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, cls).SignedString(jwtKey)
	if err != nil {
		return "", int64(0), errors.ErrBadRequest.Wrap(err)
	}

	return token, exp.Unix(), nil
}

// Invalidate is a function to invalidate the jwt token
func (svc *JWTService) Invalidate(ctx context.Context, token any) error {
	claims, err := convertToClaims(token)
	if err != nil {
		return errors.Throw(err)
	}

	if err := svc.cm.Set(ctx, claims.RegisteredClaims.ID, true, constants.TokenLifetime); err != nil {
		return errors.Throw(err)
	}

	return nil
}

// Decode is a function to convert the jwt token to user
func (svc *JWTService) Decode(ctx context.Context, token any) (*domain.User, error) {
	claims, err := convertToClaims(token)
	if err != nil {
		return nil, errors.Throw(err)
	}

	// Check JWT token invalid
	if _, err := svc.cm.Get(ctx, claims.RegisteredClaims.ID); err == nil {
		return nil, errors.ErrJWTRevoke.Trace()
	}

	user := &domain.User{
		ID:     claims.ID,
		Name:   claims.Name,
		Email:  claims.Email,
		RoleID: claims.RoleID,
	}

	return user, nil
}

// convertToClaims is a function to convert the token to claims
func convertToClaims(token any) (*Claims, error) {
	tk, ok := token.(*jwt.Token)
	if !ok {
		return nil, errors.ErrJWTInvalidCredentials.Trace()
	}
	claims, ok := tk.Claims.(*Claims)
	if !ok {
		return nil, errors.ErrJWTInvalidClaims.Trace()
	}

	return claims, nil
}
