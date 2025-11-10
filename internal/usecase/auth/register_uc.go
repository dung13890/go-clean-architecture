package auth

import (
	"context"

	"go-app/internal/domain/entity"
	"go-app/pkg/errors"
)

// Register is function used to register user
func (uc *Usecase) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	// 1. Store user to database
	if err := uc.repo.Store(ctx, user); err != nil {
		return nil, errors.ErrBadRequest.Wrap(err)
	}

	return user, nil
}
