//go:generate mockgen -source=$GOFILE -destination=mock/password_reset_mock.go

package domain

import (
	"context"
)

// PasswordResetRepository represent the passwordreset's repository contract
type PasswordResetRepository interface {
	StoreOrUpdate(ctx context.Context, email, token string) error
	FindEmailByToken(ctx context.Context, token string) (string, error)
	Delete(ctx context.Context, email, token string) error
}
