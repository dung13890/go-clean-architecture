package repository

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"

	"gorm.io/gorm"
)

// TodoRepository ...
type TodoRepository struct {
	*gorm.DB
}

// NewRepository will create new postgres object
func NewRepository(db *gorm.DB) domain.TodoRepository {
	return &TodoRepository{DB: db}
}

// Fetch will fetch content from db
func (rp *TodoRepository) Fetch(c context.Context) ([]domain.Todo, error) {
	todos := []domain.Todo{}

	if err := rp.DB.Find(&todos).Error; err != nil {
		return todos, errors.Wrap(err)
	}

	return todos, nil
}

// Find will find content from db
func (rp *TodoRepository) Find(c context.Context, id int) (*domain.Todo, error) {
	todo := domain.Todo{}

	if err := rp.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, errors.Wrap(err)
	}

	return &todo, nil
}

// Store will create data to db
func (rp *TodoRepository) Store(c context.Context, todo *domain.Todo) error {
	todo.Status = false

	if err := rp.DB.Create(todo).Error; err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// Edit will update data to db
func (rp *TodoRepository) Edit(c context.Context, todo *domain.Todo) error {
	if err := rp.DB.Save(todo).Error; err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// Remove Todo with Id in DB
func (rp *TodoRepository) Remove(c context.Context, id int) error {
	todo, err := rp.Find(c, id)
	if err != nil {
		return errors.Wrap(err)
	}

	if err := rp.DB.Delete(&todo, id).Error; err != nil {
		return errors.Wrap(err)
	}

	return nil
}
