package http

import (
	"go-app/internal/domain"
	"go-app/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthHandler represent the httphandler
type AuthHandler struct {
	Usecase domain.UserUsecase
}

// NewAuthHandler will initialize the Auth endpoint
func NewAuthHandler(g *echo.Group, uc domain.UserUsecase) {
	handler := &AuthHandler{
		Usecase: uc,
	}

	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
}

// Login for user
func (hl *AuthHandler) Login(c echo.Context) error {
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &errorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	query := domain.UserQueryParam{Email: user.Email}
	UserDb, err := hl.Usecase.FindByQuery(ctx, query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: "Invalidate Email"})
	}

	if !utils.ComparePassword(user.Password, UserDb.Password) {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: "Invalidate Password"})
	}

	tokenStr, err := utils.GenerateToken(UserDb)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &errorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserLoginResponse{
		UserID:      UserDb.ID,
		Email:       UserDb.Email,
		RoleID:      UserDb.RoleID,
		AccessToken: tokenStr,
	})
}

// Register for user
func (hl *AuthHandler) Register(c echo.Context) error {
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &errorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	passHash, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: err.Error()})
	}

	user.Password = passHash
	err = hl.Usecase.Store(ctx, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
