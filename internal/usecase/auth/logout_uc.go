package auth

import (
	"context"

	"go-app/pkg/errors"
)

// Logout is function used to logout
func (uc *Usecase) Logout(ctx context.Context, token any) error {
	if err := uc.jwtSvc.Invalidate(ctx, token); err != nil {
		return errors.Throw(err)
	}

	return nil
}
