//go:generate mockgen -source=$GOFILE -destination=mock/role_repo_mock.go
package repository

import (
	"context"

	"go-app/internal/domain/entity"
)

// RoleRepository represent the Role's repository contract
type RoleRepository interface {
	Fetch(context.Context) ([]entity.Role, error)
	Find(ctx context.Context, id uint) (*entity.Role, error)
	CheckExists(ctx context.Context, q entity.Role, id *uint) (bool, error)
	Store(ctx context.Context, u *entity.Role) error
	Update(ctx context.Context, u *entity.Role) error
	Delete(ctx context.Context, id uint) error
}
