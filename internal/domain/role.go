//go:generate mockgen -source=$GOFILE -destination=mock/role_mock.go

package domain

import (
	"context"
	"time"
)

// Role entity
type Role struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// RoleRepository represent the Role's repository contract
type RoleRepository interface {
	Fetch(context.Context) ([]Role, error)
	Find(ctx context.Context, id int) (*Role, error)
	CheckExists(ctx context.Context, q Role, id *int) (bool, error)
	Store(ctx context.Context, u *Role) error
	Update(ctx context.Context, u *Role) error
	Delete(ctx context.Context, id int) error
}

// RoleUsecase represent the Role's usecase contract
type RoleUsecase interface {
	Fetch(context.Context) ([]Role, error)
	Find(ctx context.Context, id int) (*Role, error)
	Store(ctx context.Context, u *Role) error
	Update(ctx context.Context, id int, u *Role) error
	Delete(ctx context.Context, id int) error
}
