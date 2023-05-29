package http

import (
	"go-app/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

// NewHandler will initialize the auth module endpoints
func NewHandler(
	guest *echo.Group,
	auth *echo.Group,
	uc *usecase.Usecase,
) {
	// authHandler will initialize the Auth endpoint
	authHandler := &AuthHandler{
		Usecase: uc.AuthUC,
	}

	// userHandler will initialize the users/ resources endpoint
	userHandler := &UserHandler{
		Usecase: uc.UserUC,
	}

	// roleHandler will initialize the roles/ resources endpoint
	roleHandler := &RoleHandler{
		Usecase: uc.RoleUC,
	}

	// Authenticated routes
	guest.POST("/login", authHandler.Login)
	guest.POST("/register", authHandler.Register)
	guest.POST("/forgot-password", authHandler.ForgotPassword)
	guest.POST("/reset-password", authHandler.ResetPassword)

	auth.POST("/logout", authHandler.Logout)
	auth.POST("/change-password", authHandler.ChangePassword)
	auth.GET("/me", authHandler.Me)

	// User routes
	auth.GET("/users", userHandler.Index)
	auth.GET("/users/:id", userHandler.Show)
	auth.POST("/users", userHandler.Store)
	auth.PATCH("/users/:id", userHandler.Update)
	auth.DELETE("/users/:id", userHandler.Delete)

	// Role routes
	auth.GET("/roles", roleHandler.Index)
	auth.GET("/roles/:id", roleHandler.Show)
	auth.POST("/roles", roleHandler.Store)
	auth.PATCH("/roles/:id", roleHandler.Update)
	auth.DELETE("/roles/:id", roleHandler.Delete)
}
