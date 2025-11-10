package mapper

import (
	"go-app/internal/delivery/http/dto"
	"go-app/internal/domain/entity"
)

// ConvertUserToLoginResponse DTO
func ConvertUserToLoginResponse(user entity.User, tokenStr string, exp int64) dto.UserLoginResponse {
	return dto.UserLoginResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RoleID: user.RoleID,
		Auth: dto.AuthResponse{
			AccessToken: tokenStr,
			ExpiresAt:   exp,
		},
	}
}

// ConvertLoginRequestToEntity DTO
func ConvertLoginRequestToEntity(userReq *dto.UserLoginRequest) *entity.User {
	return &entity.User{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
}

// ConvertRegisterRequestToEntity DTO
func ConvertRegisterRequestToEntity(userReq *dto.UserRegisterRequest) *entity.User {
	return &entity.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		RoleID:   userReq.RoleID,
		Password: userReq.Password,
	}
}
