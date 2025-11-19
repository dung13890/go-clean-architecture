package presenter

import (
	"go-app/internal/delivery/http/dto"
	"go-app/internal/domain/entity"
)

// ConvertUserEntityToResponse DTO http purpose
func ConvertUserEntityToResponse(user *entity.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ConvertUserRequestToEntity DTO http purpose
func ConvertUserRequestToEntity(u *dto.UserRequest) *entity.User {
	return &entity.User{
		Name:     u.Name,
		Email:    u.Email,
		RoleID:   u.RoleID,
		Password: u.Password,
	}
}
