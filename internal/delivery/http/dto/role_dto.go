package dto

import (
	"time"
)

// RoleRequest is request for create
type RoleRequest struct {
	Name string `json:"name" validate:"required"`
}

// RoleResponse is struct used for role
type RoleResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
