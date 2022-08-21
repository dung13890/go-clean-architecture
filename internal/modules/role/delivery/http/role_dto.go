package http

import (
	"go-app/internal/domain"
	"time"
)

// RolesResponse is array of role response
type RolesResponse []RoleResponse

// RoleRequest is request for create
type RoleRequest struct {
	Name string `json:"name" validate:"required"`
	Slug string `json:"slug" validate:"required"`
}

// RoleResponse is struct used for role
type RoleResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
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

// ConvertRoleToResponse DTO
func ConvertRoleToResponse(role *domain.Role) RoleResponse {
	return RoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		Slug:      role.Slug,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// ConvertRolesToResponse DTO
func ConvertRolesToResponse(roles []domain.Role) RolesResponse {
	rolesRes := make(RolesResponse, 0)

	for _, r := range roles {
		roleRes := RoleResponse{
			ID:        r.ID,
			Name:      r.Name,
			Slug:      r.Slug,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		}

		rolesRes = append(rolesRes, roleRes)
	}

	return rolesRes
}

// ConvertRequestToEntity DTO
func ConvertRequestToEntity(role *RoleRequest) *domain.Role {
	return &domain.Role{
		Name: role.Name,
		Slug: role.Slug,
	}
}
