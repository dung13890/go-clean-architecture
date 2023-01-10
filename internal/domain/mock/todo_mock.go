// Code generated by MockGen. DO NOT EDIT.
// Source: todo.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	domain "go-app/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTodoRepository is a mock of TodoRepository interface.
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository.
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance.
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// Edit mocks base method.
func (m *MockTodoRepository) Edit(ctx context.Context, t *domain.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit.
func (mr *MockTodoRepositoryMockRecorder) Edit(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockTodoRepository)(nil).Edit), ctx, t)
}

// Fetch mocks base method.
func (m *MockTodoRepository) Fetch(arg0 context.Context) ([]domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].([]domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockTodoRepositoryMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockTodoRepository)(nil).Fetch), arg0)
}

// Find mocks base method.
func (m *MockTodoRepository) Find(ctx context.Context, id int) (*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockTodoRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTodoRepository)(nil).Find), ctx, id)
}

// Remove mocks base method.
func (m *MockTodoRepository) Remove(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockTodoRepositoryMockRecorder) Remove(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockTodoRepository)(nil).Remove), ctx, id)
}

// Store mocks base method.
func (m *MockTodoRepository) Store(ctx context.Context, t *domain.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockTodoRepositoryMockRecorder) Store(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockTodoRepository)(nil).Store), ctx, t)
}

// MockTodoUsecase is a mock of TodoUsecase interface.
type MockTodoUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTodoUsecaseMockRecorder
}

// MockTodoUsecaseMockRecorder is the mock recorder for MockTodoUsecase.
type MockTodoUsecaseMockRecorder struct {
	mock *MockTodoUsecase
}

// NewMockTodoUsecase creates a new mock instance.
func NewMockTodoUsecase(ctrl *gomock.Controller) *MockTodoUsecase {
	mock := &MockTodoUsecase{ctrl: ctrl}
	mock.recorder = &MockTodoUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoUsecase) EXPECT() *MockTodoUsecaseMockRecorder {
	return m.recorder
}

// Edit mocks base method.
func (m *MockTodoUsecase) Edit(ctx context.Context, t *domain.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit.
func (mr *MockTodoUsecaseMockRecorder) Edit(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockTodoUsecase)(nil).Edit), ctx, t)
}

// Fetch mocks base method.
func (m *MockTodoUsecase) Fetch(arg0 context.Context) ([]domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].([]domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockTodoUsecaseMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockTodoUsecase)(nil).Fetch), arg0)
}

// Find mocks base method.
func (m *MockTodoUsecase) Find(ctx context.Context, id int) (*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockTodoUsecaseMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTodoUsecase)(nil).Find), ctx, id)
}

// Remove mocks base method.
func (m *MockTodoUsecase) Remove(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockTodoUsecaseMockRecorder) Remove(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockTodoUsecase)(nil).Remove), ctx, id)
}

// Store mocks base method.
func (m *MockTodoUsecase) Store(ctx context.Context, t *domain.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockTodoUsecaseMockRecorder) Store(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockTodoUsecase)(nil).Store), ctx, t)
}
