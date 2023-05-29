package http_test

import (
	"context"
	"encoding/json"
	"errors"
	"go-app/internal/domain"
	pkgErrors "go-app/pkg/errors"
	"go-app/pkg/logger"
	"go-app/pkg/validate"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	mockDomain "go-app/internal/domain/mock"
	roleHttp "go-app/internal/modules/auth/delivery/http"
)

var (
	errRoleNotFound = errors.New("not found")
)

type roleTestCase struct {
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

	tests := []roleTestCase{
		{
			name:       "OK",
			res:        []domain.Role{},
			statusCode: http.StatusOK,
			err:        nil,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				_, err := strconv.Atoi("test")
				usecaseMock.
					EXPECT().
					Fetch(context.Background()).
					Times(1).
					Return([]domain.Role{
						{
							Name: "name",
						},
					}, nil)

				if assert.Error(t, err, handler.Index(c)) {
					var rolesDto []roleHttp.RoleResponse
					_ = json.Unmarshal(rec.Body.Bytes(), &rolesDto)

					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, []roleHttp.RoleResponse{
						{
							Name: "name",
						},
					}, rolesDto)
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
	roleHandler := roleHttp.RoleHandler{Usecase: usecaseMock}

	tests := []roleTestCase{
		{
			name: "OK",
			args: 1,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				usecaseMock.EXPECT().Find(context.Background(), 1).Times(1).
					Return(&domain.Role{}, nil)

				if assert.NoError(t, roleHandler.Show(c)) {
					role := &roleHttp.RoleResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), role)

					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, &roleHttp.RoleResponse{}, role)
				}
			},
		},
		{
			name: "NOT FOUND",
			args: "test",
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				_, err := strconv.Atoi("test")

				if assert.Error(t, err, roleHandler.Show(c)) {
					httpErr := pkgErrors.ErrBadRequest.Wrap(err)
					var bErr *pkgErrors.BaseError
					if errors.As(httpErr, &bErr) {
						assert.ErrorIs(t, httpErr, bErr.Unwrap())
					}
				}
			},
		},
		{
			name: "BAD REQUEST",
			args: 2,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				usecaseMock.EXPECT().Find(context.Background(), 2).Times(1).
					Return(&domain.Role{}, errRoleNotFound)

				if assert.Error(t, errRoleNotFound, roleHandler.Show(c)) {
					err := pkgErrors.ErrBadRequest.Wrap(errRoleNotFound)
					var bErr *pkgErrors.BaseError
					if errors.As(err, &bErr) {
						assert.ErrorIs(t, err, bErr.Unwrap())
					}
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
	e.Validator = validate.NewValidate()

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockRoleUsecase(ctrl)

	roleHandler := roleHttp.RoleHandler{Usecase: usecaseMock}

	tests := []roleTestCase{
		{
			name:     "OK",
			argStore: `{"name":"admin"}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				roleMock := domain.Role{}
				_ = json.Unmarshal([]byte(`{"name":"admin"}`), &roleMock)
				usecaseMock.EXPECT().Store(context.Background(), &roleMock).
					Times(1).Return(nil).AnyTimes()

				if assert.NoError(t, roleHandler.Store(c)) {
					statusResponse := &roleHttp.StatusResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), statusResponse)

					assert.Equal(t, http.StatusCreated, rec.Code)
					assert.Equal(t, &roleHttp.StatusResponse{Status: true}, statusResponse)
				}
			},
		},
		{
			name:     "Err Binding",
			argStore: "[]",
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				roleMock := domain.Role{}
				err := json.Unmarshal([]byte(`[]`), &roleMock)
				if assert.Error(t, err, roleHandler.Store(c)) {
					var bErr *pkgErrors.BaseError
					if errors.As(err, &bErr) {
						assert.ErrorIs(t, err, bErr.Unwrap())
					}
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
