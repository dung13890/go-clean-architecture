// Code generated by MockGen. DO NOT EDIT.
// Source: jwt_svc.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	domain "go-app/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJWTService is a mock of JWTService interface.
type MockJWTService struct {
	ctrl     *gomock.Controller
	recorder *MockJWTServiceMockRecorder
}

// MockJWTServiceMockRecorder is the mock recorder for MockJWTService.
type MockJWTServiceMockRecorder struct {
	mock *MockJWTService
}

// NewMockJWTService creates a new mock instance.
func NewMockJWTService(ctrl *gomock.Controller) *MockJWTService {
	mock := &MockJWTService{ctrl: ctrl}
	mock.recorder = &MockJWTServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTService) EXPECT() *MockJWTServiceMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockJWTService) Decode(ctx context.Context, token any) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", ctx, token)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockJWTServiceMockRecorder) Decode(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockJWTService)(nil).Decode), ctx, token)
}

// GenerateToken mocks base method.
func (m *MockJWTService) GenerateToken(ctx context.Context, user *domain.User) (string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", ctx, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockJWTServiceMockRecorder) GenerateToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockJWTService)(nil).GenerateToken), ctx, user)
}

// Invalidate mocks base method.
func (m *MockJWTService) Invalidate(ctx context.Context, token any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invalidate", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Invalidate indicates an expected call of Invalidate.
func (mr *MockJWTServiceMockRecorder) Invalidate(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockJWTService)(nil).Invalidate), ctx, token)
}
