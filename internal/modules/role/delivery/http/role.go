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

type errorResponse struct {
	Message string `json:"message"`
}

// NewHandler will initialize the roles/ resources endpoint
func NewHandler(e *echo.Echo, uc domain.RoleUsecase) {
	handler := &RoleHandler{
		Usecase: uc,
	}

	g := e.Group("/api")
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

	return c.JSON(http.StatusOK, roles)
}

// Show will Find data
func (hl *RoleHandler) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, &errorResponse{Message: err.Error()})
	}
	ctx := c.Request().Context()
	role, err := hl.Usecase.Find(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &errorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, role)
}

// Store will create data
func (hl *RoleHandler) Store(c echo.Context) error {
	role := &domain.Role{}
	if err := c.Bind(role); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &errorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, role); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, role)
}

// Update will update data
func (hl *RoleHandler) Update(c echo.Context) error {
	role := &domain.Role{}

	return c.JSON(http.StatusOK, role)
}

// Delete will delete data
func (hl *RoleHandler) Delete(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
