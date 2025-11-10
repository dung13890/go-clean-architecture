//go:generate mockgen -source=$GOFILE -destination=mock/user_mock.go
package entity

import (
	"time"
)

// User entity
type User struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	RoleID    uint       `json:"role_id"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
