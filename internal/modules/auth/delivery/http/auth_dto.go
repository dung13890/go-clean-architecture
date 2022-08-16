package http

// UserLoginResponse is struct used for log in
type UserLoginResponse struct {
	UserID      uint   `json:"user_id"`
	Email       string `json:"email"`
	RoleID      int    `json:"role_id"`
	AccessToken string `json:"access_token"`
}
