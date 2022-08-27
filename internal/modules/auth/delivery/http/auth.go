package http

import (
	"go-app/internal/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthHandler represent the httphandler
type AuthHandler struct {
	Usecase domain.AuthUsecase
}

// NewHandler will initialize the Auth endpoint
func NewHandler(g *echo.Group, uc domain.AuthUsecase) {
	handler := &AuthHandler{
		Usecase: uc,
	}

	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
}

// Login for user
func (hl *AuthHandler) Login(c echo.Context) error {
	userReq := new(UserLoginRequest)
	if err := c.Bind(userReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	if err := c.Validate(userReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	claims, tokenStr, err := hl.Usecase.Login(ctx, ConvertLoginRequestToEntity(userReq))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ConvertUserToLoginResponse(*claims, tokenStr))
}

// Register for user
func (hl *AuthHandler) Register(c echo.Context) error {
	userReq := &UserRegisterRequest{}
	if err := c.Bind(userReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	if err := c.Validate(userReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ErrorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	user, err := hl.Usecase.Register(ctx, ConvertRegisterRequestToeEntity(userReq))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, ConvertUserToRegisterResponse(user))
}
