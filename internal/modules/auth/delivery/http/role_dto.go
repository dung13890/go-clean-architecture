package http

import (
	"go-app/internal/domain"
	"time"
)

// RoleRequest is request for create
type RoleRequest struct {
	Name string `json:"name" validate:"required"`
}

// RoleResponse is struct used for role
type RoleResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// convertRoleEntityToResponse DTO
func convertRoleEntityToResponse(role *domain.Role) RoleResponse {
	return RoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		Slug:      role.Slug,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// convertRoleRequestToEntity DTO
func convertRoleRequestToEntity(role *RoleRequest) *domain.Role {
	return &domain.Role{
		Name: role.Name,
	}
}
