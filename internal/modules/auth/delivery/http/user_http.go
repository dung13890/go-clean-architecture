package http

import (
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserHandler represent the http handler
type UserHandler struct {
	Usecase domain.UserUsecase
}

// Index will fetch data
func (hl *UserHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := hl.Usecase.Fetch(ctx)
	if err != nil {
		return errors.Throw(err)
	}
	usersRes := make([]UserResponse, 0)
	for i := range users {
		user := convertUserEntityToResponse(&users[i])
		usersRes = append(usersRes, user)
	}

	return c.JSON(http.StatusOK, usersRes)
}

// Show will Find data
func (hl *UserHandler) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}
	ctx := c.Request().Context()
	user, err := hl.Usecase.Find(ctx, id)
	if err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, convertUserEntityToResponse(user))
}

// Store will create data
func (hl *UserHandler) Store(c echo.Context) error {
	userReq := new(UserRequest)
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	err := c.Validate(userReq)
	if err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	user := convertUserRequestToEntity(userReq)

	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, user); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusCreated, StatusResponse{Status: true})
}

// Update will update data
func (hl *UserHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	userReq := new(UserRequest)
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	err = c.Validate(userReq)
	if err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	user := convertUserRequestToEntity(userReq)
	ctx := c.Request().Context()
	if err := hl.Usecase.Update(ctx, id, user); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}

// Delete will delete data
func (hl *UserHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	ctx := c.Request().Context()
	if err := hl.Usecase.Delete(ctx, id); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}
