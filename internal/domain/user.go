package domain

import (
	"context"
	"time"
)

// User entity
type User struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	RoleID    int        `json:"role_id"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// UserQueryParam for search
type UserQueryParam struct {
	Email string `json:"email"`
}

// UserRepository represent the user's usecases
type UserRepository interface {
	Fetch(context.Context) ([]User, error)
	Find(ctx context.Context, id int) (*User, error)
	Store(ctx context.Context, u *User) error
	Search(ctx context.Context, q UserQueryParam) ([]User, error)
	FindByQuery(ctx context.Context, q UserQueryParam) (*User, error)
}

// UserUsecase represent the user's repository contract
type UserUsecase interface {
	Fetch(context.Context) ([]User, error)
	Find(ctx context.Context, id int) (*User, error)
	Store(ctx context.Context, u *User) error
	Search(ctx context.Context, q UserQueryParam) ([]User, error)
	FindByQuery(ctx context.Context, q UserQueryParam) (*User, error)
}
