package http

import (
	"go-app/internal/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserHandler represent the httphandler
type UserHandler struct {
	Usecase domain.UserUsecase
}

// NewHandler will initialize the users/ resources endpoint
func NewHandler(g *echo.Group, uc domain.UserUsecase) {
	handler := &UserHandler{
		Usecase: uc,
	}

	g.GET("/users", handler.Index)
	g.GET("/users/:id", handler.Show)
	g.POST("/users", handler.Store)
	g.PATCH("/users/:id", handler.Update)
	g.DELETE("/users/:id", handler.Delete)
}

// Index will fetch data
func (hl *UserHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	users, _ := hl.Usecase.Fetch(ctx)

	return c.JSON(http.StatusOK, ConvertUsersToResponse(users))
}

// Show will Find data
func (hl *UserHandler) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Message: err.Error()})
	}
	ctx := c.Request().Context()
	user, err := hl.Usecase.Find(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ConvertUserToResponse(user))
}

// Store will create data
func (hl *UserHandler) Store(c echo.Context) error {
	userReq := new(UserRequest)
	if err := c.Bind(userReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	err := c.Validate(userReq)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	user := ConvertRequestToEntity(userReq)

	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, user); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, StatusResponse{Status: true})
}

// Update will update data
func (hl *UserHandler) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}

// Delete will delete data
func (hl *UserHandler) Delete(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
