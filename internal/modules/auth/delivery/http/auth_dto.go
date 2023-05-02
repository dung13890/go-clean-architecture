package http

import (
	"go-app/internal/domain"
	"time"
)

// UserLoginRequest is request for log in
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserRegisterRequest is request for register
type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	RoleID   int    `json:"role_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

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

// UserResponse is struct used for register
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

type StatusResponse struct {
	Status bool `json:"status"`
}

// convertUserToLoginResponse DTO
func convertUserToLoginResponse(claims domain.Claims, tokenStr string) UserLoginResponse {
	return UserLoginResponse{
		ID:     claims.ID,
		Name:   claims.Name,
		Email:  claims.Email,
		RoleID: claims.RoleID,
		Auth: AuthResponse{
			AccessToken: tokenStr,
			ExpiresAt:   claims.ExpiresAt.Unix(),
		},
	}
}

// convertUserToUserResponse DTO
func convertUserToUserResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// convertLoginRequestToEntity DTO
func convertLoginRequestToEntity(userReq *UserLoginRequest) *domain.User {
	return &domain.User{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
}

// convertRegisterRequestToEntity DTO
func convertRegisterRequestToEntity(userReq *UserRegisterRequest) *domain.User {
	return &domain.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		RoleID:   userReq.RoleID,
		Password: userReq.Password,
	}
}
