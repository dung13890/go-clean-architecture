package http

import (
	"go-app/internal/domain"
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

// UserForgotRequest is request for forgot password
type UserForgotRequest struct {
	Email string `json:"email" validate:"required"`
}

// UserChangePasswordRequest is request for change password
type UserChangePasswordRequest struct {
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

// UserResetPasswordRequest is request for reset password
type UserResetPasswordRequest struct {
	Token    string `json:"token" validate:"required"`
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

// StatusResponse is struct when success
type StatusResponse struct {
	Status bool `json:"status"`
}

// convertUserToLoginResponse DTO
func convertUserToLoginResponse(user domain.User, tokenStr string, exp int64) UserLoginResponse {
	return UserLoginResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RoleID: user.RoleID,
		Auth: AuthResponse{
			AccessToken: tokenStr,
			ExpiresAt:   exp,
		},
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
