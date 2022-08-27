package http_test

import (
	"context"
	"encoding/json"
	"errors"
	"go-app/internal/domain"
	"go-app/pkg/validate"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	mockDomain "go-app/internal/domain/mock"
	authHttp "go-app/internal/modules/auth/delivery/http"
)

var (
	errNotFound           = errors.New("not found")
	errRegisterInvalidate = errors.New(strings.Join([]string{
		"Name must have a value!;Email must have a value!",
		"RoleID must have a value!",
		"Password must have a value!",
	}, ";"))
	errLoginInvalidate = errors.New(strings.Join([]string{
		"Email must have a value!",
		"Password must have a value!",
	}, ";"))
)

type testCase struct {
	name       string
	argStore   string
	checkEqual func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context)
}

func TestNewAuthHandler(t *testing.T) {
	t.Parallel()
	e := echo.New()
	g := e.Group("/v1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockAuthUsecase(ctrl)
	authHttp.NewHandler(g, usecaseMock)
}

func TestHandlerRegisterUser(t *testing.T) {
	e := echo.New()

	e.Validator = validate.NewValidate()
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockAuthUsecase(ctrl)

	authHandler := authHttp.AuthHandler{Usecase: usecaseMock}

	tests := []testCase{
		{
			name:     "UNPROCESSABLE ENTITY",
			argStore: `[]`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				if assert.NoError(t, authHandler.Register(c)) {
					errorResponse := &authHttp.ErrorResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), errorResponse)

					assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				}
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
					userResponse := &authHttp.UserRegisterResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), userResponse)

					assert.Equal(t, http.StatusCreated, rec.Code)
					assert.Equal(t, &authHttp.UserRegisterResponse{}, userResponse)
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
					errRegisterInvalidate).Times(1).AnyTimes()
				if assert.NoError(t, authHandler.Register(c)) {
					errorResponse := &authHttp.ErrorResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), errorResponse)

					assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
					assert.Equal(t, &authHttp.ErrorResponse{Message: errRegisterInvalidate.Error()},
						errorResponse)
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
					errRegisterInvalidate).Times(1).AnyTimes()
				if assert.NoError(t, authHandler.Register(c)) {
					errorResponse := &authHttp.ErrorResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), errorResponse)

					assert.Equal(t, http.StatusBadRequest, rec.Code)
					assert.Equal(t, &authHttp.ErrorResponse{Message: errRegisterInvalidate.Error()},
						errorResponse)
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

	tests := []testCase{
		{
			name:     "UNPROCESSABLE ENTITY",
			argStore: `[]`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				if assert.NoError(t, authHandler.Login(c)) {
					errorResponse := &authHttp.ErrorResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), errorResponse)

					assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
				}
			},
		},
		{
			name:     "OK",
			argStore: `{"email": "user1@example", "password": "Abc1233"}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				userMock := &domain.User{}
				_ = json.Unmarshal([]byte(`{"email": "user1@example", "password": "Abc1233"}`),
					userMock)
				usecaseMock.EXPECT().Login(context.Background(), userMock).Times(1).
					Return(&domain.Claims{}, "string_token", nil).AnyTimes()

				if assert.NoError(t, authHandler.Login(c)) {
					userResponse := &authHttp.UserLoginResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), userResponse)

					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, &authHttp.UserLoginResponse{
						Auth: authHttp.AuthResponse{
							AccessToken: "string_token",
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

				usecaseMock.EXPECT().Login(c.Request().Context(), userMock).Return(&domain.Claims{}, "",
					errLoginInvalidate).Times(1).AnyTimes()
				if assert.NoError(t, authHandler.Login(c)) {
					errorResponse := &authHttp.ErrorResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), errorResponse)

					assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
					assert.Equal(t, &authHttp.ErrorResponse{Message: errLoginInvalidate.Error()},
						errorResponse)
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

				usecaseMock.EXPECT().Login(c.Request().Context(), userMock).Return(&domain.Claims{},
					"", errNotFound).Times(1).AnyTimes()
				if assert.NoError(t, authHandler.Login(c)) {
					errorResponse := &authHttp.ErrorResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), errorResponse)

					assert.Equal(t, http.StatusBadRequest, rec.Code)
					assert.Equal(t, &authHttp.ErrorResponse{Message: errNotFound.Error()},
						errorResponse)
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
