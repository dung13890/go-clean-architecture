package http

import (
	"go-app/internal/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// userHandler represent the httphandler
type userHandler struct {
	Usecase domain.UserUsecase
}

type errorResponse struct {
	Message string `json:"message"`
}

// NewHandler will initialize the users/ resources endpoint
func NewHandler(e *echo.Echo, uc domain.UserUsecase) {
	handler := &userHandler{
		Usecase: uc,
	}

	g := e.Group("/api")
	g.GET("/users", handler.Index)
	g.GET("/users/:id", handler.Show)
	g.POST("/users", handler.Store)
	g.PATCH("/users/:id", handler.Update)
	g.DELETE("/users/:id", handler.Delete)
}

// Index will fetch data
func (hl *userHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	users, _ := hl.Usecase.Fetch(ctx)

	return c.JSON(http.StatusOK, users)
}

// Show will Find data
func (hl *userHandler) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, &errorResponse{Message: err.Error()})
	}
	ctx := c.Request().Context()
	user, err := hl.Usecase.Find(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &errorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

// Store will create data
func (hl *userHandler) Store(c echo.Context) error {
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &errorResponse{Message: err.Error()})
	}
	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, user); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

// Update will update data
func (hl *userHandler) Update(c echo.Context) error {
	user := &domain.User{}

	return c.JSON(http.StatusOK, user)
}

// Delete will delete data
func (hl *userHandler) Delete(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
