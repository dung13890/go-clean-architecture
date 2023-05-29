//go:generate mockgen -source=$GOFILE -destination=mock/auth_mock.go

package domain

import (
	"context"
)

// AuthUsecase represent the auth's usecase contract
type AuthUsecase interface {
	// Register new user
	Register(ctx context.Context, u *User) (*User, error)
	// Login user
	Login(ctx context.Context, u *User, ip string) (string, int64, error)
	// Logout user
	Logout(ctx context.Context, token any) error
	// ChangePassword user
	ChangePassword(ctx context.Context, u *User, confirmPW, pw string) error
	// ForgotPassword user
	ForgotPassword(ctx context.Context, email string) error
	// ResetPassword user
	ResetPassword(ctx context.Context, token, pw string) error
}
