package http

import (
	"go-app/internal/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// RoleHandler represent the httphandler
type RoleHandler struct {
	Usecase domain.RoleUsecase
}

// NewHandler will initialize the roles/ resources endpoint
func NewHandler(g *echo.Group, uc domain.RoleUsecase) {
	handler := &RoleHandler{
		Usecase: uc,
	}

	g.GET("/roles", handler.Index)
	g.GET("/roles/:id", handler.Show)
	g.POST("/roles", handler.Store)
	g.PATCH("/roles/:id", handler.Update)
	g.DELETE("/roles/:id", handler.Delete)
}

// Index will fetch data
func (hl *RoleHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	roles, _ := hl.Usecase.Fetch(ctx)

	rolesDto := []RoleResponse{}
	for _, role := range roles {
		rolesDto = append(rolesDto, RoleResponse{
			ID:        role.ID,
			Name:      role.Name,
			Slug:      role.Slug,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, rolesDto)
}

// Show will Find data
func (hl *RoleHandler) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Message: err.Error()})
	}
	ctx := c.Request().Context()
	role, err := hl.Usecase.Find(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, RoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		Slug:      role.Slug,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	})
}

// Store will create data
func (hl *RoleHandler) Store(c echo.Context) error {
	role := &domain.Role{}
	if err := c.Bind(role); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, role); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, StatusResponse{Status: true})
}

// Update will update data
func (hl *RoleHandler) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}

// Delete will delete data
func (hl *RoleHandler) Delete(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
