// Code generated by MockGen. DO NOT EDIT.
// Source: snapshot.go

package snapshot_test

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return _m.recorder
}

// Debug mocks base method
func (_m *MockLogger) Debug(format string, args ...interface{}) {
	_s := []interface{}{format}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Debug", _s...)
}

// Debug indicates an expected call of Debug
func (_mr *MockLoggerMockRecorder) Debug(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), _s...)
}
