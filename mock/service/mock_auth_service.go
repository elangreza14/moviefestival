// Code generated by MockGen. DO NOT EDIT.
// Source: auth_service.go
//
// Generated by this command:
//
//	mockgen -source auth_service.go -destination ../../mock/service/mock_auth_service.go -package mockservice
//

// Package mockservice is a generated GoMock package.
package mockservice

import (
	context "context"
	reflect "reflect"

	domain "github.com/elangreza14/moviefestival/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockuserRepo is a mock of userRepo interface.
type MockuserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockuserRepoMockRecorder
	isgomock struct{}
}

// MockuserRepoMockRecorder is the mock recorder for MockuserRepo.
type MockuserRepoMockRecorder struct {
	mock *MockuserRepo
}

// NewMockuserRepo creates a new mock instance.
func NewMockuserRepo(ctrl *gomock.Controller) *MockuserRepo {
	mock := &MockuserRepo{ctrl: ctrl}
	mock.recorder = &MockuserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserRepo) EXPECT() *MockuserRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockuserRepo) Create(ctx context.Context, entities ...domain.User) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range entities {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockuserRepoMockRecorder) Create(ctx any, entities ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, entities...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockuserRepo)(nil).Create), varargs...)
}

// Get mocks base method.
func (m *MockuserRepo) Get(ctx context.Context, by string, val any, columns ...string) (*domain.User, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, by, val}
	for _, a := range columns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockuserRepoMockRecorder) Get(ctx, by, val any, columns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, by, val}, columns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockuserRepo)(nil).Get), varargs...)
}

// MocktokenRepo is a mock of tokenRepo interface.
type MocktokenRepo struct {
	ctrl     *gomock.Controller
	recorder *MocktokenRepoMockRecorder
	isgomock struct{}
}

// MocktokenRepoMockRecorder is the mock recorder for MocktokenRepo.
type MocktokenRepoMockRecorder struct {
	mock *MocktokenRepo
}

// NewMocktokenRepo creates a new mock instance.
func NewMocktokenRepo(ctrl *gomock.Controller) *MocktokenRepo {
	mock := &MocktokenRepo{ctrl: ctrl}
	mock.recorder = &MocktokenRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktokenRepo) EXPECT() *MocktokenRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MocktokenRepo) Create(ctx context.Context, entities ...domain.Token) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range entities {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MocktokenRepoMockRecorder) Create(ctx any, entities ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, entities...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MocktokenRepo)(nil).Create), varargs...)
}

// Get mocks base method.
func (m *MocktokenRepo) Get(ctx context.Context, by string, val any, columns ...string) (*domain.Token, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, by, val}
	for _, a := range columns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*domain.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MocktokenRepoMockRecorder) Get(ctx, by, val any, columns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, by, val}, columns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MocktokenRepo)(nil).Get), varargs...)
}