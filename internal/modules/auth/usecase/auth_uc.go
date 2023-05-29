package usecase

import (
	"context"
	"fmt"
	"go-app/internal/constants"
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"go-app/pkg/logger"
	"go-app/pkg/sendmail"
	"go-app/pkg/utils"
	"time"
)

var (
	tokenLength = 10
)

// AuthUsecase ...
type AuthUsecase struct {
	jwtSvc      domain.JWTService
	throttleSvc domain.ThrottleService
	repo        domain.UserRepository
	pwRepo      domain.PasswordResetRepository
}

// NewAuthUsecase will create new an userUsecase object representation of domain.AuthUsecase interface
func NewAuthUsecase(
	jwtSvc domain.JWTService,
	throttleSvc domain.ThrottleService,
	repo domain.UserRepository,
	pwRepo domain.PasswordResetRepository,
) *AuthUsecase {
	return &AuthUsecase{
		jwtSvc:      jwtSvc,
		throttleSvc: throttleSvc,
		repo:        repo,
		pwRepo:      pwRepo,
	}
}

// Register is function used to register user
func (uc AuthUsecase) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := uc.repo.Store(ctx, user); err != nil {
		return nil, errors.ErrBadRequest.Wrap(err)
	}

	return user, nil
}

// Login is function uses to log in
func (uc AuthUsecase) Login(ctx context.Context, u *domain.User, ip string) (string, int64, error) {
	// Check throttle login
	blocked, err := uc.throttleSvc.Blocked(ctx, u.Email, ip)
	if err != nil {
		return "", 0, errors.Throw(err)
	}
	if blocked {
		return "", 0, errors.ErrAuthThrottleLogin.Trace()
	}

	userByEmail := domain.User{Email: u.Email}
	user, err := uc.repo.FindByQuery(ctx, userByEmail)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound.Trace()) {
			_ = uc.throttleSvc.Incr(ctx, u.Email, ip)
			return "", 0, errors.ErrAuthLoginFailed.Trace()
		}
		return "", 0, errors.Throw(err)
	}

	if !utils.ComparePassword(u.Password, user.Password) {
		_ = uc.throttleSvc.Incr(ctx, u.Email, ip)
		return "", 0, errors.ErrAuthLoginFailed.Trace()
	}

	token, exp, err := uc.jwtSvc.GenerateToken(ctx, user)
	if err != nil {
		return "", 0, errors.Throw(err)
	}
	if err := uc.throttleSvc.Clear(ctx, u.Email, ip); err != nil {
		return "", 0, errors.Throw(err)
	}
	*u = *user

	return token, exp, nil
}

// Logout is function used to logout
func (uc AuthUsecase) Logout(ctx context.Context, token any) error {
	if err := uc.jwtSvc.Invalidate(ctx, token); err != nil {
		return errors.Throw(err)
	}

	return nil
}

// ForgotPassword is function used to forgot password
func (uc AuthUsecase) ForgotPassword(ctx context.Context, email string) error {
	userByEmail := domain.User{Email: email}
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
	// Send email with go routine
	sendmail := sendmail.NewEmail()
	sendmail.SetConfig(
		"Reset Password", // subject
		fmt.Sprintf(
			"Your token to reset password is %s, this token will be expired in %d minutes.",
			token,
			constants.TokenResetPasswordLifetime/time.Minute,
		), // body
		[]string{email}, // to
	)

	go func() {
		if err := sendmail.Send(); err != nil {
			logger.Debug().Printf("Send Email Error: %v", err)
		}
	}()

	return nil
}

// ChangePassword is function used to change password
func (uc AuthUsecase) ChangePassword(ctx context.Context, u *domain.User, confirmPW, pw string) error {
	user, err := uc.repo.Find(ctx, int(u.ID))
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

// ResetPassword is function used to reset password
func (uc AuthUsecase) ResetPassword(ctx context.Context, token, pw string) error {
	email, err := uc.pwRepo.FindEmailByToken(ctx, token)
	if err != nil {
		return errors.Throw(err)
	}
	// Find user by email
	userByEmail := domain.User{Email: email}
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
