package role

import (
	"context"

	"go-app/internal/domain/entity"
	"go-app/internal/domain/repository"
	"go-app/pkg/errors"
)

// Usecase ...
type Usecase struct {
	repo repository.RoleRepository
}

// NewUsecase will create new an Usecase object representation of entity.Usecase interface
func NewUsecase(repo repository.RoleRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

// Fetch will fetch content from repo
func (uc *Usecase) Fetch(c context.Context) ([]entity.Role, error) {
	items, err := uc.repo.Fetch(c)
	if err != nil {
		return nil, errors.Throw(err)
	}

	return items, nil
}

// Find will find content from repo
func (uc *Usecase) Find(c context.Context, id uint) (*entity.Role, error) {
	item, err := uc.repo.Find(c, id)
	if err != nil {
		return nil, errors.Throw(err)
	}

	return item, nil
}

// Store will create content from repo
func (uc *Usecase) Store(c context.Context, role *entity.Role) error {
	if err := uc.repo.Store(c, role); err != nil {
		return errors.Throw(err)
	}

	return nil
}

// Update will update content from repo
func (uc *Usecase) Update(ctx context.Context, id uint, r *entity.Role) error {
	// Check exist by name
	exists, err := uc.repo.CheckExists(ctx, *r, &id)
	if err != nil {
		return errors.Throw(err)
	}
	if exists {
		return errors.ErrRoleExists.Trace()
	}

	r.ID = id
	if err := uc.repo.Update(ctx, r); err != nil {
		return errors.Throw(err)
	}

	return nil
}

// Delete will delete content from repo
func (uc *Usecase) Delete(c context.Context, id uint) error {
	if err := uc.repo.Delete(c, id); err != nil {
		return errors.Throw(err)
	}

	return nil
}
