package auth

import (
	"context"

	"go-app/internal/domain/entity"
	"go-app/pkg/errors"
	"go-app/pkg/utils"
)

// ChangePassword is function used to change password
func (uc Usecase) ChangePassword(ctx context.Context, u *entity.User, confirmPW, pw string) error {
	user, err := uc.repo.Find(ctx, u.ID)
	if err != nil {
		return errors.Throw(err)
	}
	if !utils.ComparePassword(confirmPW, user.Password) {
		return errors.ErrAuthInvalidateConfirmPass.Trace()
	}

	user.Password = pw
	if err := uc.repo.Update(ctx, user); err != nil {
		return errors.Throw(err)
	}

	return nil
}
