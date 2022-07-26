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

// RoleRepository represent the role's usecases
type RoleRepository interface {
	Fetch(context.Context) ([]Role, error)
	Find(ctx context.Context, id int) (*Role, error)
	Store(ctx context.Context, u *Role) error
}

// RoleUsecase represent the role's repository contract
type RoleUsecase interface {
	Fetch(context.Context) ([]Role, error)
	Find(ctx context.Context, id int) (*Role, error)
	Store(ctx context.Context, u *Role) error
}
