package usecase

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
)

// UserUsecase ...
type UserUsecase struct {
	repo domain.UserRepository
}

// NewUsecase will create new UserUsecase object
func NewUsecase(rp domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		repo: rp,
	}
}

// Fetch will fetch content from repo
func (uc *UserUsecase) Fetch(c context.Context) ([]domain.User, error) {
	items, _ := uc.repo.Fetch(c)

	return items, nil
}

// Find will find content from repo
func (uc *UserUsecase) Find(c context.Context, id int) (*domain.User, error) {
	item, err := uc.repo.Find(c, id)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return item, nil
}

// Store will create content from repo
func (uc *UserUsecase) Store(c context.Context, user *domain.User) error {
	if err := uc.repo.Store(c, user); err != nil {
		return errors.Wrap(err)
	}

	return nil
}
