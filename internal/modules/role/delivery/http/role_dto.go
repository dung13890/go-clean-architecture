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

// ErrorResponse is struct when error
type ErrorResponse struct {
	Message string `json:"message"`
}

// StatusResponse is struct when success
type StatusResponse struct {
	Status bool `json:"status"`
}

// convertEntityToResponse DTO
func convertEntityToResponse(role *domain.Role) RoleResponse {
	return RoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		Slug:      role.Slug,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// convertRequestToEntity DTO
func convertRequestToEntity(role *RoleRequest) *domain.Role {
	return &domain.Role{
		Name: role.Name,
	}
}
