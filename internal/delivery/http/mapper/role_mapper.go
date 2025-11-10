package mapper

import (
	"go-app/internal/delivery/http/dto"
	"go-app/internal/domain/entity"
)

// ConvertRoleEntityToResponse DTO
func ConvertRoleEntityToResponse(role *entity.Role) dto.RoleResponse {
	return dto.RoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		Slug:      role.Slug,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// ConvertRoleRequestToEntity DTO
func ConvertRoleRequestToEntity(role *dto.RoleRequest) *entity.Role {
	return &entity.Role{
		Name: role.Name,
	}
}
