package http

import (
	"net/http"

	"go-app/internal/delivery/http/dto"
	"go-app/internal/delivery/http/mapper"
	"go-app/internal/domain/entity"
	"go-app/internal/infrastructure/constant"
	"go-app/internal/usecase"
	"go-app/pkg/errors"

	"github.com/labstack/echo/v4"
)

// authHandler represent the httphandler
type authHandler struct {
	usecase *usecase.AuthUsecase
}

// NewAuthHandler will create new an authHandler object
func NewAuthHandler(usecase *usecase.AuthUsecase) *authHandler {
	return &authHandler{
		usecase: usecase,
	}
}

// Login for user
func (hl *authHandler) Login(c echo.Context) error {
	userReq := new(dto.UserLoginRequest)
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := c.Validate(userReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	ctx := c.Request().Context()
	user := mapper.ConvertLoginRequestToEntity(userReq)
	tokenStr, exp, err := hl.usecase.Login(ctx, user, c.RealIP())
	if err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, mapper.ConvertUserToLoginResponse(*user, tokenStr, exp))
}

// Logout for user
func (hl *authHandler) Logout(c echo.Context) error {
	token := c.Get("user")
	ctx := c.Request().Context()
	if err := hl.usecase.Logout(ctx, token); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, dto.StatusResponse{Status: true})
}

// Me for user
func (*authHandler) Me(c echo.Context) error {
	user, _ := c.Get(constant.GuardJWT).(*entity.User)

	return c.JSON(http.StatusOK, mapper.ConvertUserEntityToResponse(user))
}

// Register for user
func (hl *authHandler) Register(c echo.Context) error {
	userReq := &dto.UserRegisterRequest{}
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := c.Validate(userReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	ctx := c.Request().Context()
	user, err := hl.usecase.Register(ctx, mapper.ConvertRegisterRequestToEntity(userReq))
	if err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusCreated, mapper.ConvertUserEntityToResponse(user))
}

// ChangePassword will return status when change password success
func (hl *authHandler) ChangePassword(c echo.Context) error {
	userReq := &dto.UserChangePasswordRequest{}
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := c.Validate(userReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	user, ok := c.Get(constant.GuardJWT).(*entity.User)
	if !ok {
		return errors.ErrBadRequest.Trace()
	}

	ctx := c.Request().Context()
	if err := hl.usecase.ChangePassword(ctx, user, userReq.ConfirmPassword, userReq.Password); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, dto.StatusResponse{Status: true})
}

// ForgotPassword will return status when verify email success
func (hl *authHandler) ForgotPassword(c echo.Context) error {
	userReq := &dto.UserForgotRequest{}
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := c.Validate(userReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	ctx := c.Request().Context()
	if err := hl.usecase.ForgotPassword(ctx, userReq.Email); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, dto.StatusResponse{Status: true})
}

// ResetPassword use token from email to reset password
func (hl *authHandler) ResetPassword(c echo.Context) error {
	userReq := &dto.UserResetPasswordRequest{}
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	if err := c.Validate(userReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	ctx := c.Request().Context()
	if err := hl.usecase.ResetPassword(ctx, userReq.Token, userReq.Password); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, dto.StatusResponse{Status: true})
}
