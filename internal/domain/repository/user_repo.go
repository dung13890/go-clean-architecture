//go:generate mockgen -source=$GOFILE -destination=mock/user_repo_mock.go
package repository

import (
	"context"

	"go-app/internal/domain/entity"
)

// UserRepository represent the User's repository contract
type UserRepository interface {
	Fetch(context.Context) ([]entity.User, error)
	Find(ctx context.Context, id uint) (*entity.User, error)
	Store(ctx context.Context, u *entity.User) error
	FindByQuery(ctx context.Context, q entity.User) (*entity.User, error)
	CheckExists(ctx context.Context, q entity.User, id *uint) (bool, error)
	Update(ctx context.Context, u *entity.User) error
	Delete(ctx context.Context, id uint) error
}
