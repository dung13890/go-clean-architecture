package http_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"

	mockDomain "go-app/internal/domain/mock"
	authHttp "go-app/internal/modules/auth/delivery/http"
	"go-app/internal/modules/auth/usecase"
)

func TestNewAuthHandler(t *testing.T) {
	t.Parallel()
	e := echo.New()
	g := e.Group("/v1")
	ga := e.Group("")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := &usecase.Usecase{
		RoleUC: mockDomain.NewMockRoleUsecase(ctrl),
		UserUC: mockDomain.NewMockUserUsecase(ctrl),
		AuthUC: mockDomain.NewMockAuthUsecase(ctrl),
	}
	authHttp.NewHandler(g, ga, uc)
}
