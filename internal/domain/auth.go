package domain

import (
	"context"
)

// AuthUsecase represent the user's repository contract
type AuthUsecase interface {
	Register(ctx context.Context, u *User) (*User, error)
	Login(ctx context.Context, u *User) (*User, string, error)
}
