package auth

import (
	"context"

	"go-app/internal/domain/entity"
	"go-app/pkg/errors"
	"go-app/pkg/utils"
)

// Login is function uses to log in
func (uc *Usecase) Login(ctx context.Context, u *entity.User, ip string) (string, int64, error) {
	// Check throttle login
	if blocked, err := uc.throttleSvc.Blocked(ctx, u.Email, ip); err != nil {
		return "", 0, errors.Throw(err)
	} else if blocked {
		return "", 0, errors.ErrAuthThrottleLogin.Trace()
	}

	// Retrieve user by email
	user, err := uc.repo.FindByQuery(ctx, entity.User{Email: u.Email})
	if err != nil {
		if errors.Is(err, errors.ErrNotFound.Trace()) {
			_ = uc.throttleSvc.Incr(ctx, u.Email, ip)
			return "", 0, errors.ErrAuthLoginFailed.Trace()
		}
		return "", 0, errors.Throw(err)
	}

	// Compare passwords
	if !utils.ComparePassword(u.Password, user.Password) {
		_ = uc.throttleSvc.Incr(ctx, u.Email, ip)
		return "", 0, errors.ErrAuthLoginFailed.Trace()
	}

	// Generate token
	token, exp, err := uc.jwtSvc.GenerateToken(ctx, user)
	if err != nil {
		return "", 0, errors.Throw(err)
	}

	// Clear throttle data
	if err := uc.throttleSvc.Clear(ctx, u.Email, ip); err != nil {
		return "", 0, errors.Throw(err)
	}

	*u = *user
	return token, exp, nil
}
