package http

import (
	"net/http"
	"strings"

	"go-app/internal/domain/gateway"
	"go-app/internal/infrastructure/config"
	"go-app/internal/infrastructure/registry"
	"go-app/pkg/errors"
	"go-app/pkg/logger"
	"go-app/pkg/validate"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewHTTPHandler registry http
func NewHTTPHandler(
	e *echo.Echo,
	svc gateway.JWTService,
	registry *registry.Registry,
) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = validate.NewValidate()
	e.HTTPErrorHandler = jsonErrorHandler
	g := e.Group("/api")

	// CORS restricted with a custom function to allow origins
	// and with the GET, PUT, POST or DELETE methods allowed.
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: corsAllowOrigin,
		AllowMethods:    []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Middleware
	au := g.Group("")
	au.Use(setupJWT())
	au.Use(authenticated(svc))

	// Init Handler
	authHandler := NewAuthHandler(registry.AuthUc)
	userHandler := NewUserHandler(registry.UserUc)
	roleHandler := NewRoleHandler(registry.RoleUc)

	// Authenticated routes
	g.POST("/login", authHandler.Login)
	g.POST("/register", authHandler.Register)
	g.POST("/forgot-password", authHandler.ForgotPassword)
	g.POST("/reset-password", authHandler.ResetPassword)

	au.POST("/logout", authHandler.Logout)
	au.POST("/change-password", authHandler.ChangePassword)
	au.GET("/me", authHandler.Me)

	// User routes
	au.GET("/users", userHandler.Index)
	au.GET("/users/:id", userHandler.Show)
	au.POST("/users", userHandler.Store)
	au.PATCH("/users/:id", userHandler.Update)
	au.DELETE("/users/:id", userHandler.Delete)

	// Role routes
	au.GET("/roles", roleHandler.Index)
	au.GET("/roles/:id", roleHandler.Show)
	au.POST("/roles", roleHandler.Store)
	au.PATCH("/roles/:id", roleHandler.Update)
	au.DELETE("/roles/:id", roleHandler.Delete)
}

func corsAllowOrigin(origin string) (bool, error) {
	list := strings.Split(config.GetAppConfig().AllowedOrigin, ",")

	for _, b := range list {
		if b == origin {
			return true, nil
		}
	}

	return false, nil
}

func jsonErrorHandler(err error, ctx echo.Context) {
	status := http.StatusInternalServerError
	responseError := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    status,
		Message: http.StatusText(status),
	}

	var he *echo.HTTPError
	if errors.As(err, &he) {
		status = he.Code
		responseError.Code = status
		if m, ok := he.Message.(string); ok {
			responseError.Message = m
		}
	}

	var be *errors.BaseError
	if errors.As(err, &be) {
		status = be.Status
		responseError.Message = be.Message
		if status == http.StatusUnprocessableEntity {
			if beErr := be.Unwrap(); beErr != nil {
				responseError.Message = beErr.Error()
			}
		}
		responseError.Code = be.Code
	}

	if !ctx.Response().Committed {
		// Logger if status >= 500
		if status >= http.StatusInternalServerError {
			logger.Debugf("%+v", err)
		}
		if ctx.Request().Method == http.MethodHead { // Issue #608
			err = ctx.NoContent(status)
		} else {
			err = ctx.JSON(status, responseError)
		}
		if err != nil {
			ctx.Logger().Error(err)
		}
	}
}
