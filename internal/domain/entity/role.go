//go:generate mockgen -source=$GOFILE -destination=mock/role_mock.go
package entity

import (
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
