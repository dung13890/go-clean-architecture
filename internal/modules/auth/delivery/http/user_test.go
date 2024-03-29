package http_test

import (
	"context"
	"encoding/json"
	"go-app/internal/domain"
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
	userHttp "go-app/internal/modules/auth/delivery/http"
)

type userTestCase struct {
	name       string
	res        interface{}
	args       interface{}
	argStore   string
	err        error
	statusCode int
	checkEqual func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context)
}

func TestHandlerIndexUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	usecaseMock := mockDomain.NewMockUserUsecase(ctrl)
	handler := userHttp.UserHandler{
		Usecase: usecaseMock,
	}

	tests := []userTestCase{
		{
			name:       "OK",
			res:        []domain.User{},
			statusCode: http.StatusOK,
			err:        nil,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				_, err := strconv.Atoi("test")
				usecaseMock.
					EXPECT().
					Fetch(context.Background()).
					Times(1).
					Return([]domain.User{
						{
							Name: "name",
						},
					}, nil)

				if assert.Error(t, err, handler.Index(c)) {
					var usersDto []userHttp.UserResponse
					_ = json.Unmarshal(rec.Body.Bytes(), &usersDto)

					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, []userHttp.UserResponse{
						{
							Name: "name",
						},
					}, usersDto)
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

func TestHandlerShowUser(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)

	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockUserUsecase(ctrl)

	userHandler := userHttp.UserHandler{Usecase: usecaseMock}

	tests := []userTestCase{
		{
			name: "OK",
			args: 1,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				usecaseMock.EXPECT().Find(context.Background(), 1).Times(1).
					Return(&domain.User{}, nil)

				if assert.NoError(t, userHandler.Show(c)) {
					user := &userHttp.UserResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), user)

					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, &userHttp.UserResponse{}, user)
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

func TestHandlerStoreUser(t *testing.T) {
	e := echo.New()

	e.Validator = validate.NewValidate()
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockUserUsecase(ctrl)

	userHandler := userHttp.UserHandler{Usecase: usecaseMock}

	tests := []userTestCase{
		{
			name:     "OK",
			argStore: `{"name":"user", "email":"user@example.com","role_id":1, "password":"abc@123"}`,
			checkEqual: func(t *testing.T, rec *httptest.ResponseRecorder, c echo.Context) {
				t.Helper()
				userMock := domain.User{}
				_ = json.Unmarshal([]byte(`{"name":"user", "email":"user@example.com",
								"role_id":1, "password":"abc@123"}`), &userMock)
				usecaseMock.EXPECT().Store(context.Background(), &userMock).Times(1).
					Return(nil).AnyTimes()

				if assert.NoError(t, userHandler.Store(c)) {
					statusResponse := &userHttp.StatusResponse{}
					_ = json.Unmarshal(rec.Body.Bytes(), statusResponse)

					assert.Equal(t, http.StatusCreated, rec.Code)
					assert.Equal(t, &userHttp.StatusResponse{Status: true}, statusResponse)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tc.argStore))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tc.checkEqual(t, rec, c)
		})
	}
}
