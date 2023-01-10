package http

import (
	"go-app/internal/domain"
	"time"

	"github.com/labstack/echo/v4"
)

// TodosResponse is array of todo response
type TodosResponse []TodoResponse

// TodoRequest is request for create
type TodoRequest struct {
	Title  string `json:"title" validate:"required"`
	Status bool   `json:"status" validate:"required"`
}

// TodoResponse is struct used for todo
type TodoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StatusResponse is struct when success
type StatusResponse struct {
	Status bool `json:"status"`
}

// ErrorResponse is struct when error
type ErrorResponse struct {
	Message string `json:"message"`
}

// ConvertTodoToResponse DTO
func ConvertTodoToResponse(todo *domain.Todo) TodoResponse {
	return TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Status:    todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

// ConvertTodosToResponse DTO
func ConvertTodosToResponse(todos []domain.Todo) TodosResponse {
	todosRes := make(TodosResponse, 0)

	for _, t := range todos {
		todoRes := TodoResponse{
			ID:        t.ID,
			Title:     t.Title,
			Status:    t.Status,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		}

		todosRes = append(todosRes, todoRes)
	}

	return todosRes
}

// ConvertRequestToEntity DTO
func ConvertRequestToEntity(todo *TodoRequest) *domain.Todo {
	return &domain.Todo{
		Title:  todo.Title,
		Status: todo.Status,
	}
}

// CheckBindAndValidate DTO
func CheckBindAndValidate(c echo.Context) (*TodoRequest, error) {
	todoReq := new(TodoRequest)

	if err := c.Bind(todoReq); err != nil {
		return nil, err
	}

	if err := c.Validate(todoReq); err != nil {
		return nil, err
	}

	return todoReq, nil
}
