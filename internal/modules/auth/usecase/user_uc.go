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

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

// Fetch will fetch content from repo
func (uc *UserUsecase) Fetch(c context.Context) ([]domain.User, error) {
	items, err := uc.repo.Fetch(c)
	if err != nil {
		return nil, errors.Throw(err)
	}

	return items, nil
}

// Find will find content from repo
func (uc *UserUsecase) Find(c context.Context, id int) (*domain.User, error) {
	item, err := uc.repo.Find(c, id)
	if err != nil {
		return nil, errors.Throw(err)
	}

	return item, nil
}

// Store will create content from repo
func (uc *UserUsecase) Store(c context.Context, user *domain.User) error {
	if err := uc.repo.Store(c, user); err != nil {
		return errors.Throw(err)
	}

	return nil
}

// FindByQuery is a function that returns a user filtered by query
func (uc *UserUsecase) FindByQuery(ctx context.Context, q domain.User) (*domain.User, error) {
	item, err := uc.repo.FindByQuery(ctx, q)
	if err != nil {
		return nil, errors.Throw(err)
	}

	return item, nil
}

// Update will update content from repo
func (uc *UserUsecase) Update(ctx context.Context, id int, u *domain.User) error {
	// Check exist by email
	userByEmail := domain.User{Email: u.Email}
	exists, err := uc.repo.CheckExists(ctx, userByEmail, &id)
	if err != nil {
		return errors.Throw(err)
	}
	if exists {
		return errors.ErrUserExistsByEmail.Trace()
	}
	u.ID = uint(id)
	if err := uc.repo.Update(ctx, u); err != nil {
		return errors.Throw(err)
	}

	return nil
}

// Delete will delete content from repo
func (uc *UserUsecase) Delete(c context.Context, id int) error {
	if err := uc.repo.Delete(c, id); err != nil {
		return errors.Throw(err)
	}

	return nil
}
