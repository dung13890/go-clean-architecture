package registry

import (
	"go-app/config"
	authHttp "go-app/internal/modules/auth/delivery/http"
	"go-app/pkg/errors"
	"go-app/pkg/logger"
	"go-app/pkg/validate"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewHTTPHandler registry http
func NewHTTPHandler(e *echo.Echo, uc *Usecase, svc *Service) {
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

	authGroup := g.Group("")
	authGroup.Use(authHttp.SetupJWT())
	authGroup.Use(authHttp.Authenticated(svc.JWTSvc))

	authHttp.NewHandler(g, authGroup, uc.AuthModule)
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
			logger.Debug().Printf("%+v", err)
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
