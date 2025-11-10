package http

import (
	"go-app/internal/domain/entity"
	"go-app/internal/domain/service"
	"go-app/internal/infrastructure/config"
	"go-app/internal/infrastructure/constant"
	"go-app/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// setupJWT .-
func setupJWT() echo.MiddlewareFunc {
	jwtConf := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.Claims)
		},
		SigningKey: []byte(config.GetAppConfig().AppJWTKey),
	}

	return echojwt.WithConfig(jwtConf)
}

// authenticated .-
func authenticated(svc service.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user")
			user, err := svc.Decode(c.Request().Context(), token)
			if err != nil {
				return errors.Throw(err)
			}

			c.Set(constant.GuardJWT, user)

			return next(c)
		}
	}
}
