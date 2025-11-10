package http

import (
	"net/http"
	"strconv"

	"go-app/internal/delivery/http/dto"
	"go-app/internal/delivery/http/mapper"
	"go-app/internal/usecase/role"
	"go-app/pkg/errors"

	"github.com/labstack/echo/v4"
)

// roleHandler represent the http handler
type roleHandler struct {
	usecase *role.Usecase
}

// NewRoleHandler will create new an roleHandler object
func NewRoleHandler(usecase *role.Usecase) *roleHandler {
	return &roleHandler{
		usecase: usecase,
	}
}

// Index will fetch data
func (hl *roleHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	roles, err := hl.usecase.Fetch(ctx)
	if err != nil {
		return errors.Throw(err)
	}
	rolesRes := make([]dto.RoleResponse, 0)
	for i := range roles {
		role := mapper.ConvertRoleEntityToResponse(&roles[i])
		rolesRes = append(rolesRes, role)
	}

	return c.JSON(http.StatusOK, rolesRes)
}

// Show will Find data
func (hl *roleHandler) Show(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return errors.Throw(err)
	}
	ctx := c.Request().Context()
	role, err := hl.usecase.Find(ctx, uint(id))
	if err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, mapper.ConvertRoleEntityToResponse(role))
}

// Store will create data
func (hl *roleHandler) Store(c echo.Context) error {
	roleReq := new(dto.RoleRequest)
	if err := c.Bind(roleReq); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	err := c.Validate(roleReq)
	if err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	role := mapper.ConvertRoleRequestToEntity(roleReq)

	ctx := c.Request().Context()
	if err := hl.usecase.Store(ctx, role); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusCreated, dto.StatusResponse{Status: true})
}

// Update will update data
func (hl *roleHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	roleReq := new(dto.RoleRequest)
	if err := c.Bind(roleReq); err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	err = c.Validate(roleReq)
	if err != nil {
		return errors.ErrUnprocessableEntity.Wrap(err)
	}

	role := mapper.ConvertRoleRequestToEntity(roleReq)

	ctx := c.Request().Context()
	if err := hl.usecase.Update(ctx, uint(id), role); err != nil {
		return errors.Throw(err)
	}

	return c.JSON(http.StatusOK, dto.StatusResponse{Status: true})
}

// Delete will delete data
func (hl *roleHandler) Delete(c echo.Context) error {
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
