// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iboware/bot-detect/pkg/detector (interfaces: Detector)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDetector is a mock of Detector interface.
type MockDetector struct {
	ctrl     *gomock.Controller
	recorder *MockDetectorMockRecorder
}

// MockDetectorMockRecorder is the mock recorder for MockDetector.
type MockDetectorMockRecorder struct {
	mock *MockDetector
}

// NewMockDetector creates a new mock instance.
func NewMockDetector(ctrl *gomock.Controller) *MockDetector {
	mock := &MockDetector{ctrl: ctrl}
	mock.recorder = &MockDetectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDetector) EXPECT() *MockDetectorMockRecorder {
	return m.recorder
}

// Detect mocks base method.
func (m *MockDetector) Detect(arg0 context.Context, arg1 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Detect", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Detect indicates an expected call of Detect.
func (mr *MockDetectorMockRecorder) Detect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detect", reflect.TypeOf((*MockDetector)(nil).Detect), arg0, arg1)
}
