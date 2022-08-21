package registry

import (
	"go-app/config"
	"go-app/internal/domain"
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

	authGroup := g.Group("")
	authHttp.NewHandler(authGroup, uc.AuthUsecase)

	DefaultJWTConfig := middleware.JWTConfig{
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
		Claims:      &domain.Claims{},
		SigningKey:  []byte(config.GetAppConfig().AppJWTKey),
	}
	g.Use(middleware.JWTWithConfig(DefaultJWTConfig))

	// CORS restricted with a custom function to allow origins
	// and with the GET, PUT, POST or DELETE methods allowed.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: corsAllowOrigin,
		AllowMethods:    []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	roleHttp.NewHandler(g, uc.RoleUsecase)
	userHttp.NewHandler(g, uc.UserUsecase)
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
