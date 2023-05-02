package http

import (
	"go-app/internal/domain"
	"time"
)

// UserRequest is struct used for create user
type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	RoleID   int    `json:"role_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserResponse is struct used for user
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ErrorResponse is struct when error
type ErrorResponse struct {
	Message string `json:"message"`
}

// StatusResponse is struct when success
type StatusResponse struct {
	Status bool `json:"status"`
}

// convertEntityToResponse DTO
func convertEntityToResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// convertRequestToEntity DTO
func convertRequestToEntity(u *UserRequest) *domain.User {
	return &domain.User{
		Name:     u.Name,
		Email:    u.Email,
		RoleID:   u.RoleID,
		Password: u.Password,
	}
}
