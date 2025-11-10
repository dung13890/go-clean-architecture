package http

import (
	"net/http"
	"strconv"

	"go-app/internal/delivery/http/dto"
	"go-app/internal/delivery/http/mapper"
	"go-app/internal/usecase/user"
	"go-app/pkg/errors"

	"github.com/labstack/echo/v4"
)

// userHandler represent the http handler
type userHandler struct {
	usecase *user.Usecase
}

// NewUserHandler will create new an userHandler object
func NewUserHandler(usecase *user.Usecase) *userHandler {
	return &userHandler{
		usecase: usecase,
	}
}

// Index will fetch data
func (hl *userHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := hl.usecase.Fetch(ctx)
	if err != nil {
		return errors.Throw(err)
	}
	usersRes := make([]dto.UserResponse, 0)
	for i := range users {
		user := mapper.ConvertUserEntityToResponse(&users[i])
		usersRes = append(usersRes, user)
	}

	return c.JSON(http.StatusOK, usersRes)
}

// Show will Find data
func (hl *userHandler) Show(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}
	ctx := c.Request().Context()
	user, err := hl.usecase.Find(ctx, uint(id))
	if err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, mapper.ConvertUserEntityToResponse(user))
}

// Store will create data
func (hl *userHandler) Store(c echo.Context) error {
	userReq := new(dto.UserRequest)
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	err := c.Validate(userReq)
	if err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	user := mapper.ConvertUserRequestToEntity(userReq)

	ctx := c.Request().Context()
	if err := hl.usecase.Store(ctx, user); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusCreated, dto.StatusResponse{Status: true})
}

// Update will update data
func (hl *userHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	userReq := new(dto.UserRequest)
	if err := c.Bind(userReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	err = c.Validate(userReq)
	if err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	user := mapper.ConvertUserRequestToEntity(userReq)
	ctx := c.Request().Context()
	if err := hl.usecase.Update(ctx, uint(id), user); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, dto.StatusResponse{Status: true})
}

// Delete will delete data
func (hl *userHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	ctx := c.Request().Context()
	if err := hl.usecase.Delete(ctx, uint(id)); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, dto.StatusResponse{Status: true})
}
