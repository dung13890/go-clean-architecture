//go:generate mockgen -source=$GOFILE -destination=mock/jwt_svc_mock.go

package domain

import (
	"context"
)

// JWTService is a struct that represent the jwt's service
type JWTService interface {
	GenerateToken(ctx context.Context, user *User) (string, int64, error)
	Invalidate(ctx context.Context, token any) error
	Decode(ctx context.Context, token any) (*User, error)
}
