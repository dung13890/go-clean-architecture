package http

import (
	"time"
)

// UserLoginResponse is struct used for log in
type UserLoginResponse struct {
	UserID      uint   `json:"user_id"`
	Email       string `json:"email"`
	RoleID      int    `json:"role_id"`
	AccessToken string `json:"access_token"`
}

// UserRegisterResponse is struct used for register
type UserRegisterResponse struct {
	UserID    uint      `json:"user_id"`
	Email     string    `json:"email"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
