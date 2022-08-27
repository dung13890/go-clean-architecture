package http

import (
	"go-app/internal/domain"
	"time"
)

// UsersResponse is array of user response
type UsersResponse []UserResponse

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

// ConvertUserToResponse DTO
func ConvertUserToResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ConvertUsersToResponse DTO
func ConvertUsersToResponse(users []domain.User) UsersResponse {
	usersRes := make(UsersResponse, 0)

	for _, u := range users {
		userRes := UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			RoleID:    u.RoleID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}

		usersRes = append(usersRes, userRes)
	}

	return usersRes
}

// ConvertRequestToEntity DTO
func ConvertRequestToEntity(u *UserRequest) *domain.User {
	return &domain.User{
		Name:     u.Name,
		Email:    u.Email,
		RoleID:   u.RoleID,
		Password: u.Password,
	}
}
