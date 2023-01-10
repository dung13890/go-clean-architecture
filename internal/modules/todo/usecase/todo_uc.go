package usecase

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
)

// TodoUsecase ...
type TodoUsecase struct {
	repo domain.TodoRepository
}

// NewUsecase will create new TodoUsecase object
func NewUsecase(rp domain.TodoRepository) domain.TodoUsecase {
	return &TodoUsecase{
		repo: rp,
	}
}

// Fetch will fetch content from repo
func (uc *TodoUsecase) Fetch(c context.Context) ([]domain.Todo, error) {
	items, _ := uc.repo.Fetch(c)

	return items, nil
}

// Find will find content from repo
func (uc *TodoUsecase) Find(c context.Context, id int) (*domain.Todo, error) {
	item, err := uc.repo.Find(c, id)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return item, nil
}

// Store will create content from repo
func (uc *TodoUsecase) Store(c context.Context, todo *domain.Todo) error {
	if err := uc.repo.Store(c, todo); err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// Edit will update content from repo
func (uc *TodoUsecase) Edit(c context.Context, todo *domain.Todo) error {
	if err := uc.repo.Edit(c, todo); err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// Remove will delete content from repo
func (uc *TodoUsecase) Remove(c context.Context, id int) error {
	if err := uc.repo.Remove(c, id); err != nil {
		return errors.Wrap(err)
	}

	return nil
}
