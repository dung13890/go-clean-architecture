package http

import (
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// RoleHandler represent the http handler
type RoleHandler struct {
	Usecase domain.RoleUsecase
}

// Index will fetch data
func (hl *RoleHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	roles, err := hl.Usecase.Fetch(ctx)
	if err != nil {
		return errors.Throw(err)
	}
	rolesRes := make([]RoleResponse, 0)
	for i := range roles {
		role := convertRoleEntityToResponse(&roles[i])
		rolesRes = append(rolesRes, role)
	}

	return c.JSON(http.StatusOK, rolesRes)
}

// Show will Find data
func (hl *RoleHandler) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.Throw(err)
	}
	ctx := c.Request().Context()
	role, err := hl.Usecase.Find(ctx, id)
	if err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, convertRoleEntityToResponse(role))
}

// Store will create data
func (hl *RoleHandler) Store(c echo.Context) error {
	roleReq := new(RoleRequest)
	if err := c.Bind(roleReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	err := c.Validate(roleReq)
	if err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	role := convertRoleRequestToEntity(roleReq)

	ctx := c.Request().Context()
	if err := hl.Usecase.Store(ctx, role); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusCreated, StatusResponse{Status: true})
}

// Update will update data
func (hl *RoleHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	roleReq := new(RoleRequest)
	if err := c.Bind(roleReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	err = c.Validate(roleReq)
	if err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	role := convertRoleRequestToEntity(roleReq)

	ctx := c.Request().Context()
	if err := hl.Usecase.Update(ctx, id, role); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, StatusResponse{Status: true})
}

// Delete will delete data
func (hl *RoleHandler) Delete(c echo.Context) error {
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
