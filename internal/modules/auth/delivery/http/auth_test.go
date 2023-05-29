package http_test

import (
	"context"
	"encoding/json"
	"errors"
	"go-app/internal/domain"
	pkgErrors "go-app/pkg/errors"
	"go-app/pkg/validate"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	mockDomain "go-app/internal/domain/mock"
	authHttp "go-app/internal/modules/auth/delivery/http"
)

var (
	errAuthNotFound           = errors.New("not found")
	errAuthRegisterInvalidate = errors.New(strings.Join([]string{
		"Name must have a value!;Email must have a value!",
		"RoleID must have a value!",
		"Password must have a value!",
	}, ";"))
	errAuthLoginInvalidate = errors.New(strings.Join([]string{
		"Email must have a value!",
		"Password must have a value!",
	}, ";"))
)

type authTestCase struct {
	name       string
	argStore   string
	checkEqual func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context)
}

func TestHandlerRegisterUser(t *testing.T) {
	e := echo.New()

	e.Validator = validate.NewValidate()
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockAuthUsecase(ctrl)

	authHandler := authHttp.AuthHandler{Usecase: usecaseMock}

	tests := []authTestCase{
		{
			name:     "UNPROCESSABLE ENTITY",
			argStore: `[]`,

			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()

				err := authHandler.Register(c)
				assert.Equal(t, err.Error(), pkgErrors.ErrBadRequest.Error())
			},
		},
		{
			name:     "OK",
			argStore: `{"email": "email@example.com", "name": "user", "role_id": 1, "password": "user"}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				userMock := &domain.User{}
				_ = json.Unmarshal([]byte(`{"email": "email@example.com", "name": "user", "role_id": 1,
					"password": "user"}`), &userMock)
				usecaseMock.EXPECT().Register(context.Background(), userMock).Times(1).
					Return(&domain.User{}, nil).AnyTimes()

				if assert.NoError(t, authHandler.Register(c)) {
					userResponse := &authHttp.UserResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), userResponse)

					assert.Equal(t, http.StatusCreated, rec.Code)
					assert.Equal(t, &authHttp.UserResponse{}, userResponse)
				}
			},
		},
		{
			name:     "Validate",
			argStore: `{}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				userMock := &domain.User{}
				_ = json.Unmarshal([]byte(`{}`), userMock)

				usecaseMock.EXPECT().Register(c.Request().Context(), userMock).Return(&domain.User{},
					errAuthRegisterInvalidate).Times(1).AnyTimes()

				if assert.Error(t, errAuthRegisterInvalidate, authHandler.Register(c)) {
					var bErr *pkgErrors.BaseError
					if errors.As(errAuthRegisterInvalidate, &bErr) {
						assert.ErrorIs(t, errAuthRegisterInvalidate, bErr.Unwrap())
					}
				}
			},
		},
		{
			name:     "BAD REQUEST",
			argStore: `{"email": "email1@example.com", "name": "user1", "role_id": 2, "password": "user1"}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				userMock := &domain.User{}
				_ = json.Unmarshal([]byte(`{"email": "email1@example.com", "name": "user1",
								"role_id": 2, "password": "user1"}`), userMock)

				usecaseMock.EXPECT().Register(c.Request().Context(), userMock).Return(&domain.User{},
					errAuthRegisterInvalidate).Times(1).AnyTimes()

				if assert.Error(t, errAuthRegisterInvalidate, authHandler.Register(c)) {
					var bErr *pkgErrors.BaseError
					if errors.As(errAuthRegisterInvalidate, &bErr) {
						assert.ErrorIs(t, errAuthRegisterInvalidate, bErr.Unwrap())
					}
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(tc.argStore))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tc.checkEqual(t, rec, c)
		})
	}
}

func TestHandlerLoginUser(t *testing.T) {
	e := echo.New()

	e.Validator = validate.NewValidate()
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockAuthUsecase(ctrl)

	authHandler := authHttp.AuthHandler{Usecase: usecaseMock}

	tests := []authTestCase{
		{
			name:     "UNPROCESSABLE ENTITY",
			argStore: `[]`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				err := authHandler.Login(c)
				assert.Equal(t, err.Error(), pkgErrors.ErrBadRequest.Error())
			},
		},
		{
			name:     "OK",
			argStore: `{"email": "user1@example", "password": "Abc1233"}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				now := time.Now()
				t.Helper()
				userMock := &domain.User{}
				_ = json.Unmarshal([]byte(`{"email": "user1@example", "password": "Abc1233"}`),
					userMock)
				usecaseMock.EXPECT().Login(context.Background(), userMock, gomock.Any()).Times(1).
					Return("string_token", now.Unix(), nil).AnyTimes()

				if assert.NoError(t, authHandler.Login(c)) {
					userResponse := &authHttp.UserLoginResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), userResponse)

					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, &authHttp.UserLoginResponse{
						Email: "user1@example",
						Auth: authHttp.AuthResponse{
							AccessToken: "string_token",
							ExpiresAt:   now.Unix(),
						},
					}, userResponse)
				}
			},
		},
		{
			name:     "Validate",
			argStore: `{}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				userMock := &domain.User{}
				_ = json.Unmarshal([]byte(`{}`), userMock)

				usecaseMock.
					EXPECT().
					Login(c.Request().Context(), userMock, gomock.Any()).
					Return("", int64(0), errAuthLoginInvalidate).Times(1).AnyTimes()

				if assert.Error(t, errAuthLoginInvalidate, authHandler.Login(c)) {
					var bErr *pkgErrors.BaseError
					if errors.As(errAuthLoginInvalidate, &bErr) {
						assert.ErrorIs(t, errAuthLoginInvalidate, bErr.Unwrap())
					}
				}
			},
		},
		{
			name:     "BAD REQUEST",
			argStore: `{"email": "user2@example", "password": "Abc1234"}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				userMock := &domain.User{}
				_ = json.Unmarshal([]byte(`{"email": "user2@example", "password": "Abc1234"}`),
					userMock)

				usecaseMock.
					EXPECT().
					Login(c.Request().Context(), userMock, gomock.Any()).
					Return("", int64(0), errAuthNotFound).Times(1).AnyTimes()

				if assert.Error(t, errAuthNotFound, authHandler.Login(c)) {
					var bErr *pkgErrors.BaseError
					if errors.As(errAuthNotFound, &bErr) {
						assert.ErrorIs(t, errAuthNotFound, bErr.Unwrap())
					}
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(tc.argStore))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tc.checkEqual(t, rec, c)
		})
	}
}
