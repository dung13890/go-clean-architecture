package http

import (
	"time"
)

// UserLoginResponse is struct used for log in
type UserLoginResponse struct {
	ID     uint         `json:"id"`
	Name   string       `json:"name"`
	Email  string       `json:"email"`
	RoleID int          `json:"role_id"`
	Auth   AuthResponse `json:"auth"`
}

// AuthResponse is struct used for token
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

// UserRegisterResponse is struct used for register
type UserRegisterResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ErrorResponse is struc when error
type ErrorResponse struct {
	Message string `json:"message"`
}
