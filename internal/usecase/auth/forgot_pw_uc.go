package auth

import (
	"context"
	"fmt"
	"time"

	"go-app/internal/domain/entity"
	"go-app/internal/infrastructure/constant"
	"go-app/pkg/errors"
	"go-app/pkg/logger"
	"go-app/pkg/utils"
)

// ForgotPassword is function used to forgot password
func (uc *Usecase) ForgotPassword(ctx context.Context, email string) error {
	userByEmail := entity.User{Email: email}
	exists, err := uc.repo.CheckExists(ctx, userByEmail, nil)
	if err != nil {
		return errors.Throw(err)
	}
	if !exists {
		return errors.ErrAuthInvalidateEmail.Trace()
	}

	// Create token and store to database
	token, err := utils.RandString(tokenLength)
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := uc.pwRepo.StoreOrUpdate(ctx, email, token); err != nil {
		return errors.Throw(err)
	}

	bodyMail := fmt.Sprintf(
		"Your token to reset password is %s, this token will be expired in %d minutes.",
		token,
		constant.TokenResetPasswordLifetime/time.Minute,
	)
	// Send email with go routine
	go func() {
		if err := uc.mailSvc.Send(ctx, "Reset Password", bodyMail, []string{email}); err != nil {
			logger.Debugf("Send Email Error: %v", err)
		}
	}()

	return nil
}
