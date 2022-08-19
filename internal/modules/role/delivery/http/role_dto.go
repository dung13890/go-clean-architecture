package http

import (
	"time"
)

// RoleResponse is struct used for role
type RoleResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ErrorResponse is struc when error
type ErrorResponse struct {
	Message string `json:"message"`
}

// StatusResponse is struc when success
type StatusResponse struct {
	Status bool `json:"status"`
}
