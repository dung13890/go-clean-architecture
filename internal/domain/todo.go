//go:generate mockgen -source=$GOFILE -destination=mock/todo_mock.go

package domain

import (
	"context"
	"time"
)

// Todo entity
type Todo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TodoRepository represent the todo's usecases
type TodoRepository interface {
	Fetch(context.Context) ([]Todo, error)
	Find(ctx context.Context, id int) (*Todo, error)
	Store(ctx context.Context, t *Todo) error
	Edit(ctx context.Context, t *Todo) error
	Remove(ctx context.Context, id int) error
}

// TodoUsecase represent the todo's repository contract
type TodoUsecase interface {
	Fetch(context.Context) ([]Todo, error)
	Find(ctx context.Context, id int) (*Todo, error)
	Store(ctx context.Context, t *Todo) error
	Edit(ctx context.Context, t *Todo) error
	Remove(ctx context.Context, id int) error
}
