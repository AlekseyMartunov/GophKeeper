// Code generated by MockGen. DO NOT EDIT.
// Source: pairhandler.go

// Package mock_pairhandlers is a generated GoMock package.
package mock_pairhandlers

import (
	"GophKeeper/app/internal/entity/pairs"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mocklogger is a mock of logger interface.
type Mocklogger struct {
	ctrl     *gomock.Controller
	recorder *MockloggerMockRecorder
}

// MockloggerMockRecorder is the mock recorder for Mocklogger.
type MockloggerMockRecorder struct {
	mock *Mocklogger
}

// NewMocklogger creates a new mock instance.
func NewMocklogger(ctrl *gomock.Controller) *Mocklogger {
	mock := &Mocklogger{ctrl: ctrl}
	mock.recorder = &MockloggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocklogger) EXPECT() *MockloggerMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *Mocklogger) Error(e error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", e)
}

// Error indicates an expected call of Error.
func (mr *MockloggerMockRecorder) Error(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*Mocklogger)(nil).Error), e)
}

// Info mocks base method.
func (m *Mocklogger) Info(s string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Info", s)
}

// Info indicates an expected call of Info.
func (mr *MockloggerMockRecorder) Info(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*Mocklogger)(nil).Info), s)
}

// MockpairService is a mock of pairService interface.
type MockpairService struct {
	ctrl     *gomock.Controller
	recorder *MockpairServiceMockRecorder
}

// MockpairServiceMockRecorder is the mock recorder for MockpairService.
type MockpairServiceMockRecorder struct {
	mock *MockpairService
}

// NewMockpairService creates a new mock instance.
func NewMockpairService(ctrl *gomock.Controller) *MockpairService {
	mock := &MockpairService{ctrl: ctrl}
	mock.recorder = &MockpairServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpairService) EXPECT() *MockpairServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockpairService) Delete(ctx context.Context, pairName string, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, pairName, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockpairServiceMockRecorder) Delete(ctx, pairName, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockpairService)(nil).Delete), ctx, pairName, userID)
}

// Get mocks base method.
func (m *MockpairService) Get(ctx context.Context, pairName string, userID int) (pairs.Pair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, pairName, userID)
	ret0, _ := ret[0].(pairs.Pair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockpairServiceMockRecorder) Get(ctx, pairName, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockpairService)(nil).Get), ctx, pairName, userID)
}

// GetAll mocks base method.
func (m *MockpairService) GetAll(ctx context.Context, userID int) ([]pairs.Pair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, userID)
	ret0, _ := ret[0].([]pairs.Pair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockpairServiceMockRecorder) GetAll(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockpairService)(nil).GetAll), ctx, userID)
}

// Save mocks base method.
func (m *MockpairService) Save(ctx context.Context, pair pairs.Pair) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, pair)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockpairServiceMockRecorder) Save(ctx, pair interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockpairService)(nil).Save), ctx, pair)
}
