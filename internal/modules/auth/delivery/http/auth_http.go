package http

import (
	"go-app/internal/constants"
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthHandler represent the httphandler
type AuthHandler struct {
	Usecase domain.AuthUsecase
}

// Login for user
func (hl *AuthHandler) Login(c echo.Context) error {
	userReq := new(UserLoginRequest)
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := c.Validate(userReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	ctx := c.Request().Context()
	user := convertLoginRequestToEntity(userReq)
	tokenStr, exp, err := hl.Usecase.Login(ctx, user, c.RealIP())
	if err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, convertUserToLoginResponse(*user, tokenStr, exp))
}

// Logout for user
func (hl *AuthHandler) Logout(c echo.Context) error {
	token := c.Get("user")
	ctx := c.Request().Context()
	if err := hl.Usecase.Logout(ctx, token); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}

// Me for user
func (_ *AuthHandler) Me(c echo.Context) error {
	user, _ := c.Get(constants.GuardJWT).(*domain.User)

	return c.JSON(http.StatusOK, convertUserEntityToResponse(user))
}

// Register for user
func (hl *AuthHandler) Register(c echo.Context) error {
	userReq := &UserRegisterRequest{}
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := c.Validate(userReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	ctx := c.Request().Context()
	user, err := hl.Usecase.Register(ctx, convertRegisterRequestToEntity(userReq))
	if err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusCreated, convertUserEntityToResponse(user))
}

// ChangePassword will return status when change password success
func (hl *AuthHandler) ChangePassword(c echo.Context) error {
	userReq := &UserChangePasswordRequest{}
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := c.Validate(userReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	user, ok := c.Get(constants.GuardJWT).(*domain.User)
	if !ok {
		return errors.ErrBadRequest.Trace()
	}

	ctx := c.Request().Context()
	if err := hl.Usecase.ChangePassword(ctx, user, userReq.ConfirmPassword, userReq.Password); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}

// ForgotPassword will return status when verify email success
func (hl *AuthHandler) ForgotPassword(c echo.Context) error {
	userReq := &UserForgotRequest{}
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := c.Validate(userReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	ctx := c.Request().Context()
	if err := hl.Usecase.ForgotPassword(ctx, userReq.Email); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}

// ResetPassword use token from email to reset password
func (hl *AuthHandler) ResetPassword(c echo.Context) error {
	userReq := &UserResetPasswordRequest{}
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := c.Validate(userReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	ctx := c.Request().Context()
	if err := hl.Usecase.ResetPassword(ctx, userReq.Token, userReq.Password); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}
