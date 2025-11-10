package auth

import (
	"context"

	"go-app/internal/domain/entity"
	"go-app/pkg/errors"
)

// ResetPassword is function used to reset password
func (uc *Usecase) ResetPassword(ctx context.Context, token, pw string) error {
	email, err := uc.pwRepo.FindEmailByToken(ctx, token)
	if err != nil {
		return errors.Throw(err)
	}
	// Find user by email
	userByEmail := entity.User{Email: email}
	user, err := uc.repo.FindByQuery(ctx, userByEmail)
	if err != nil {
		return errors.Throw(err)
	}

	user.Password = pw
	if err := uc.repo.Update(ctx, user); err != nil {
		return errors.Throw(err)
	}

	// Revoke token
	if err := uc.pwRepo.Delete(ctx, email, token); err != nil {
		return errors.Throw(err)
	}

	return nil
}
