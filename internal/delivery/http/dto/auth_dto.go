package dto

// UserLoginRequest is request for log in
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserRegisterRequest is request for register
type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	RoleID   uint   `json:"role_id" validate:"required"`
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
	RoleID uint         `json:"role_id"`
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
