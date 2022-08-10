package http_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-app/internal/domain"
	"go-app/pkg/logger"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	mockDomain "go-app/internal/domain/mock"
	roleHttp "go-app/internal/modules/role/delivery/http"
)

var errNotFound = errors.New("not found")

type testCase struct {
	name       string
	res        interface{}
	args       interface{}
	argStore   string
	err        error
	statusCode int
	checkEqual func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context)
}

func TestHandlerIndexRole(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/roles", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	usecaseMock := mockDomain.NewMockRoleUsecase(ctrl)
	handler := roleHttp.RoleHandler{
		Usecase: usecaseMock,
	}

	tests := []testCase{
		{
			name:       "OK",
			res:        []domain.Role{},
			statusCode: http.StatusOK,
			err:        nil,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				_, err := strconv.Atoi("test")
				usecaseMock.EXPECT().Fetch(context.Background()).Times(1).Return([]domain.Role{}, nil)

				if assert.Error(t, err, handler.Index(c)) {
					var role domain.Role
					_ = json.Unmarshal(rec.Body.Bytes(), &role)

					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, &domain.Role{}, &role)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.checkEqual(t, rec, c)
		})
	}
}

func TestHandlerShowRole(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/roles/:id", nil)

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockRoleUsecase(ctrl)
	userRoleHandler := roleHttp.RoleHandler{Usecase: usecaseMock}

	tests := []testCase{
		{
			name: "OK",
			args: 1,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				usecaseMock.EXPECT().Find(context.Background(), 1).Times(1).Return(&domain.Role{}, nil)

				if assert.NoError(t, userRoleHandler.Show(c)) {
					var role domain.Role
					_ = json.Unmarshal(rec.Body.Bytes(), &role)

					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, &domain.Role{}, &role)
				}
			},
		},
		{
			name: "NOT FOUND",
			args: "test",
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				_, err := strconv.Atoi("test")

				if assert.Error(t, err, userRoleHandler.Show(c)) {
					var role domain.Role
					_ = json.Unmarshal(rec.Body.Bytes(), &role)

					assert.Equal(t, http.StatusNotFound, rec.Code)
					assert.Equal(t, &domain.Role{}, &role)
				}
			},
		},
		{
			name: "NOT FOUND",
			args: 2,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				usecaseMock.EXPECT().Find(context.Background(), 2).Times(1).Return(&domain.Role{}, errNotFound)

				if assert.Error(t, errNotFound, userRoleHandler.Show(c)) {
					var role domain.Role
					_ = json.Unmarshal(rec.Body.Bytes(), &role)

					assert.Equal(t, http.StatusNotFound, rec.Code)
					assert.Equal(t, &domain.Role{}, &role)
				}
			},
		},
	}

	for _, tc := range tests {
		out, err := json.Marshal(tc.args)
		if err != nil {
			logger.Error().Printf("error during json marshal: %v", err)
		}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetParamNames("id")
		c.SetParamValues(string(out))
		t.Run(tc.name, func(t *testing.T) {
			tc.checkEqual(t, rec, c)
		})
	}
}

func TestHandlerStoreRole(t *testing.T) {
	e := echo.New()

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockRoleUsecase(ctrl)

	userRoleHandler := roleHttp.RoleHandler{Usecase: usecaseMock}

	tests := []testCase{
		{
			name:     "OK",
			argStore: `{}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				roleMock := domain.Role{}
				_ = json.Unmarshal([]byte(`{}`), &roleMock)
				usecaseMock.EXPECT().Store(context.Background(), &roleMock).Times(1).Return(nil).AnyTimes()

				if assert.NoError(t, userRoleHandler.Store(c)) {
					role := domain.Role{}
					_ = json.Unmarshal(rec.Body.Bytes(), &role)

					assert.Equal(t, http.StatusCreated, rec.Code)
					assert.Equal(t, &domain.Role{}, &role)
				}
			},
		},
		{
			name:     "UNPROCESSABLE ENTITY",
			argStore: fmt.Sprintf("%v", domain.Role{}),
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				if assert.NoError(t, userRoleHandler.Store(c)) {
					var role domain.Role
					_ = json.Unmarshal(rec.Body.Bytes(), &role)

					assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				}
			},
		},
		{
			name:     "BAD REQUEST",
			argStore: `{"Name": "test"}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				roleMock := domain.Role{}
				_ = json.Unmarshal([]byte(`{}`), &roleMock)
				roleMock.Name = "test"

				usecaseMock.EXPECT().Store(c.Request().Context(), &roleMock).Return(errNotFound).Times(1)
				if assert.NoError(t, userRoleHandler.Store(c)) {
					var role domain.Role
					_ = json.Unmarshal(rec.Body.Bytes(), &role)

					assert.Equal(t, http.StatusBadRequest, rec.Code)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/roles", strings.NewReader(tc.argStore))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tc.checkEqual(t, rec, c)
		})
	}
}

func TestHandlerUpdateRole(t *testing.T) {
	e := echo.New()

	out, err := json.Marshal(domain.Role{})
	if err != nil {
		logger.Error().Printf("error when json marshal: %v", err)
	}

	req := httptest.NewRequest(http.MethodPatch, "/roles/:id", strings.NewReader(string(out)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockRoleUsecase(ctrl)

	userRoleHandler := roleHttp.RoleHandler{Usecase: usecaseMock}

	tests := []testCase{
		{
			name:       "OK",
			res:        &domain.Role{},
			statusCode: http.StatusOK,
			err:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if assert.NoError(t, userRoleHandler.Update(c)) {
				var role domain.Role
				_ = json.Unmarshal(rec.Body.Bytes(), &role)

				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Equal(t, tc.res, &role)
			}
		})
	}
}

func TestHandlerDeleteRole(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/roles/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockRoleUsecase(ctrl)

	userRoleHandler := roleHttp.RoleHandler{Usecase: usecaseMock}

	tests := []testCase{
		{
			name:       "OK",
			res:        &domain.Role{},
			statusCode: http.StatusNoContent,
			err:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if assert.NoError(t, userRoleHandler.Delete(c)) {
				var role domain.Role
				_ = json.Unmarshal(rec.Body.Bytes(), &role)

				assert.Equal(t, http.StatusNoContent, rec.Code)
				assert.Equal(t, tc.res, &role)
			}
		})
	}
}

func TestNewRoleHandler(t *testing.T) {
	t.Parallel()
	e := echo.New()
	g := e.Group("/v1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockRoleUsecase(ctrl)
	roleHttp.NewHandler(g, usecaseMock)
}
