package factory

import (
	"go-app/config"
	// Role
	roleHttp "go-app/internal/modules/role/delivery/http"
	roleRepository "go-app/internal/modules/role/repository"
	roleUsecase "go-app/internal/modules/role/usecase"

	// User
	userHttp "go-app/internal/modules/user/delivery/http"
	userRepository "go-app/internal/modules/user/repository"
	userUsecase "go-app/internal/modules/user/usecase"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

// NewFactory Create repo, usecase, http object
func NewFactory(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted with a custom function to allow origins
	// and with the GET, PUT, POST or DELETE methods allowed.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: corsAllowOrigin,
		AllowMethods:    []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Role Factory
	rRepo := roleRepository.NewRepository(db)
	rUc := roleUsecase.NewUsecase(rRepo)
	roleHttp.NewHandler(e, rUc)

	// User Factory
	uRepo := userRepository.NewRepository(db)
	uUc := userUsecase.NewUsecase(uRepo)
	userHttp.NewHandler(e, uUc)
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
