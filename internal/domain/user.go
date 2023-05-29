//go:generate mockgen -source=$GOFILE -destination=mock/user_mock.go

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

// UserRepository represent the User's repository contract
type UserRepository interface {
	Fetch(context.Context) ([]User, error)
	Find(ctx context.Context, id int) (*User, error)
	Store(ctx context.Context, u *User) error
	FindByQuery(ctx context.Context, q User) (*User, error)
	CheckExists(ctx context.Context, q User, id *int) (bool, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id int) error
}

// UserUsecase represent the User's usecase contract
type UserUsecase interface {
	Fetch(context.Context) ([]User, error)
	Find(ctx context.Context, id int) (*User, error)
	Store(ctx context.Context, u *User) error
	FindByQuery(ctx context.Context, q User) (*User, error)
	Update(ctx context.Context, id int, u *User) error
	Delete(ctx context.Context, id int) error
}
