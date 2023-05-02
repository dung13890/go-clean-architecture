package http

import (
	"go-app/config"
	"go-app/internal/constants"
	"go-app/internal/domain"
	"go-app/pkg/errors"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Authenticate() echo.MiddlewareFunc {
	jwtConf := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.Claims)
		},
		SigningKey: []byte(config.GetAppConfig().AppJWTKey),
	}

	return echojwt.WithConfig(jwtConf)
}

func SetUserFromClaims() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return errors.New("JWT token missing or invalid")
			}

			claims, ok := token.Claims.(*domain.Claims)
			if !ok {
				return errors.New("Failed to cast claims as jwt.MapClaims")
			}

			user := &domain.User{
				ID:     claims.ID,
				Name:   claims.Name,
				Email:  claims.Email,
				RoleID: claims.RoleID,
			}

			c.Set(constants.GuardJWT, user)

			return next(c)
		}
	}
}
