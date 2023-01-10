package http

import (
	"go-app/internal/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// TodoHandler represent the httphandler
type TodoHandler struct {
	Usecase domain.TodoUsecase
}

// NewHandler will initialize the todos/ resources endpoint
func NewHandler(g *echo.Group, uc domain.TodoUsecase) {
	handler := &TodoHandler{
		Usecase: uc,
	}

	g.GET("/todos", handler.Fetch)
	g.GET("/todos/:id", handler.Find)
	g.POST("/todos", handler.Store)
	g.PATCH("/todos/:id", handler.Update)
	g.DELETE("/todos/:id", handler.Remove)
}

// Fetch will get all todo
func (hl *TodoHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()
	todos, _ := hl.Usecase.Fetch(ctx)

	return c.JSON(http.StatusOK, ConvertTodosToResponse(todos))
}

// Find will get todo by id
func (hl *TodoHandler) Find(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	todo, err := hl.Usecase.Find(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ConvertTodoToResponse(todo))
}

// Store will create new todo
func (hl *TodoHandler) Store(c echo.Context) error {
	todoReq, err := CheckBindAndValidate(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	todo := ConvertRequestToEntity(todoReq)
	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, todo); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, StatusResponse{Status: true})
}

// Update will edit todo by id
func (hl *TodoHandler) Update(c echo.Context) error {
	todoReq, err := CheckBindAndValidate(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	todo := ConvertRequestToEntity(todoReq)
	ctx := c.Request().Context()
	if err := hl.Usecase.Edit(ctx, todo); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}

// Remove will delete todo by id
func (hl *TodoHandler) Remove(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if err := hl.Usecase.Remove(ctx, id); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}
