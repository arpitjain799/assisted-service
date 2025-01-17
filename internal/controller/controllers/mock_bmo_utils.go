// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/controller/controllers (interfaces: BMOUtils)

// Package controllers is a generated GoMock package.
package controllers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBMOUtils is a mock of BMOUtils interface.
type MockBMOUtils struct {
	ctrl     *gomock.Controller
	recorder *MockBMOUtilsMockRecorder
}

// MockBMOUtilsMockRecorder is the mock recorder for MockBMOUtils.
type MockBMOUtilsMockRecorder struct {
	mock *MockBMOUtils
}

// NewMockBMOUtils creates a new mock instance.
func NewMockBMOUtils(ctrl *gomock.Controller) *MockBMOUtils {
	mock := &MockBMOUtils{ctrl: ctrl}
	mock.recorder = &MockBMOUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBMOUtils) EXPECT() *MockBMOUtilsMockRecorder {
	return m.recorder
}

// ConvergedFlowAvailable mocks base method.
func (m *MockBMOUtils) ConvergedFlowAvailable() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvergedFlowAvailable")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ConvergedFlowAvailable indicates an expected call of ConvergedFlowAvailable.
func (mr *MockBMOUtilsMockRecorder) ConvergedFlowAvailable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvergedFlowAvailable", reflect.TypeOf((*MockBMOUtils)(nil).ConvergedFlowAvailable))
}

// GetIronicServiceURLS mocks base method.
func (m *MockBMOUtils) GetIronicServiceURLS() (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIronicServiceURLS")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIronicServiceURLS indicates an expected call of GetIronicServiceURLS.
func (mr *MockBMOUtilsMockRecorder) GetIronicServiceURLS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIronicServiceURLS", reflect.TypeOf((*MockBMOUtils)(nil).GetIronicServiceURLS))
}
