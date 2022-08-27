package usecase

import (
	"context"
	"go-app/internal/domain"
	"go-app/pkg/errors"
)

// RoleUsecase ...
type RoleUsecase struct {
	repo domain.RoleRepository
}

// NewUsecase will create new RoleUsecase object
func NewUsecase(rp domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{
		repo: rp,
	}
}

// Fetch will fetch content from repo
func (uc *RoleUsecase) Fetch(c context.Context) ([]domain.Role, error) {
	items, _ := uc.repo.Fetch(c)

	return items, nil
}

// Find will find content from repo
func (uc *RoleUsecase) Find(c context.Context, id int) (*domain.Role, error) {
	item, err := uc.repo.Find(c, id)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return item, nil
}

// Store will create content from repo
func (uc *RoleUsecase) Store(c context.Context, role *domain.Role) error {
	if err := uc.repo.Store(c, role); err != nil {
		return errors.Wrap(err)
	}

	return nil
}
