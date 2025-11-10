package dto

import (
	"time"
)

// UserRequest is struct used for create user
type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	RoleID   uint   `json:"role_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserResponse is struct used for user
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	RoleID    uint      `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
