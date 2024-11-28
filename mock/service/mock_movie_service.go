// Code generated by MockGen. DO NOT EDIT.
// Source: movie_service.go
//
// Generated by this command:
//
//	mockgen -source movie_service.go -destination ../../mock/service/mock_movie_service.go -package mockservice
//

// Package mockservice is a generated GoMock package.
package mockservice

import (
	context "context"
	reflect "reflect"

	domain "github.com/elangreza14/moviefestival/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockmovieRepo is a mock of movieRepo interface.
type MockmovieRepo struct {
	ctrl     *gomock.Controller
	recorder *MockmovieRepoMockRecorder
	isgomock struct{}
}

// MockmovieRepoMockRecorder is the mock recorder for MockmovieRepo.
type MockmovieRepoMockRecorder struct {
	mock *MockmovieRepo
}

// NewMockmovieRepo creates a new mock instance.
func NewMockmovieRepo(ctrl *gomock.Controller) *MockmovieRepo {
	mock := &MockmovieRepo{ctrl: ctrl}
	mock.recorder = &MockmovieRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmovieRepo) EXPECT() *MockmovieRepoMockRecorder {
	return m.recorder
}

// CreateMovieTX mocks base method.
func (m *MockmovieRepo) CreateMovieTX(ctx context.Context, movie domain.Movie, artists, genres []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMovieTX", ctx, movie, artists, genres)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMovieTX indicates an expected call of CreateMovieTX.
func (mr *MockmovieRepoMockRecorder) CreateMovieTX(ctx, movie, artists, genres any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMovieTX", reflect.TypeOf((*MockmovieRepo)(nil).CreateMovieTX), ctx, movie, artists, genres)
}

// Get mocks base method.
func (m *MockmovieRepo) Get(ctx context.Context, by string, val any, columns ...string) (*domain.Movie, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, by, val}
	for _, a := range columns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*domain.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockmovieRepoMockRecorder) Get(ctx, by, val any, columns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, by, val}, columns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockmovieRepo)(nil).Get), varargs...)
}

// GetAll mocks base method.
func (m *MockmovieRepo) GetAll(ctx context.Context) ([]domain.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]domain.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockmovieRepoMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockmovieRepo)(nil).GetAll), ctx)
}

// GetMovieDetail mocks base method.
func (m *MockmovieRepo) GetMovieDetail(ctx context.Context, movieID int) (*domain.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieDetail", ctx, movieID)
	ret0, _ := ret[0].(*domain.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieDetail indicates an expected call of GetMovieDetail.
func (mr *MockmovieRepoMockRecorder) GetMovieDetail(ctx, movieID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieDetail", reflect.TypeOf((*MockmovieRepo)(nil).GetMovieDetail), ctx, movieID)
}

// GetMoviesWithPaginationAndSearch mocks base method.
func (m *MockmovieRepo) GetMoviesWithPaginationAndSearch(ctx context.Context, search, searchBy, orderBy, orderDirection string, page, pageSize int) ([]domain.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMoviesWithPaginationAndSearch", ctx, search, searchBy, orderBy, orderDirection, page, pageSize)
	ret0, _ := ret[0].([]domain.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMoviesWithPaginationAndSearch indicates an expected call of GetMoviesWithPaginationAndSearch.
func (mr *MockmovieRepoMockRecorder) GetMoviesWithPaginationAndSearch(ctx, search, searchBy, orderBy, orderDirection, page, pageSize any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMoviesWithPaginationAndSearch", reflect.TypeOf((*MockmovieRepo)(nil).GetMoviesWithPaginationAndSearch), ctx, search, searchBy, orderBy, orderDirection, page, pageSize)
}

// UpdateMovieTX mocks base method.
func (m *MockmovieRepo) UpdateMovieTX(ctx context.Context, movie domain.Movie, artists, genres []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovieTX", ctx, movie, artists, genres)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMovieTX indicates an expected call of UpdateMovieTX.
func (mr *MockmovieRepoMockRecorder) UpdateMovieTX(ctx, movie, artists, genres any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovieTX", reflect.TypeOf((*MockmovieRepo)(nil).UpdateMovieTX), ctx, movie, artists, genres)
}

// MockmovieViewRepo is a mock of movieViewRepo interface.
type MockmovieViewRepo struct {
	ctrl     *gomock.Controller
	recorder *MockmovieViewRepoMockRecorder
	isgomock struct{}
}

// MockmovieViewRepoMockRecorder is the mock recorder for MockmovieViewRepo.
type MockmovieViewRepoMockRecorder struct {
	mock *MockmovieViewRepo
}

// NewMockmovieViewRepo creates a new mock instance.
func NewMockmovieViewRepo(ctrl *gomock.Controller) *MockmovieViewRepo {
	mock := &MockmovieViewRepo{ctrl: ctrl}
	mock.recorder = &MockmovieViewRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmovieViewRepo) EXPECT() *MockmovieViewRepoMockRecorder {
	return m.recorder
}

// AddMovieViewTX mocks base method.
func (m *MockmovieViewRepo) AddMovieViewTX(ctx context.Context, movieID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMovieViewTX", ctx, movieID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMovieViewTX indicates an expected call of AddMovieViewTX.
func (mr *MockmovieViewRepoMockRecorder) AddMovieViewTX(ctx, movieID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMovieViewTX", reflect.TypeOf((*MockmovieViewRepo)(nil).AddMovieViewTX), ctx, movieID)
}
