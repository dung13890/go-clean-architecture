package registry

import (
	"go-app/config"
	authHttp "go-app/internal/modules/auth/delivery/http"
	roleHttp "go-app/internal/modules/role/delivery/http"
	userHttp "go-app/internal/modules/user/delivery/http"
	"go-app/pkg/validate"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewHTTPHandler registry http
func NewHTTPHandler(e *echo.Echo, uc *Usecase) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = validate.NewValidate()
	g := e.Group("/api")

	// CORS restricted with a custom function to allow origins
	// and with the GET, PUT, POST or DELETE methods allowed.
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: corsAllowOrigin,
		AllowMethods:    []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	authGroup := g.Group("")
	authGroup.Use(authHttp.Authenticate())
	authGroup.Use(authHttp.SetUserFromClaims())

	authHttp.NewHandler(g, authGroup, uc.AuthUsecase)
	roleHttp.NewHandler(authGroup, uc.RoleUsecase)
	userHttp.NewHandler(authGroup, uc.UserUsecase)
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
