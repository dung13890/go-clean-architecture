//go:generate mockgen -source=$GOFILE -destination=mock/auth_mock.go

package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

// Claims is struct claims for jwt
type Claims struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID int    `json:"role_id"`
	jwt.RegisteredClaims
}

// AuthUsecase represent the user's repository contract
type AuthUsecase interface {
	Register(ctx context.Context, u *User) (*User, error)
	Login(ctx context.Context, u *User) (*Claims, string, error)
}
