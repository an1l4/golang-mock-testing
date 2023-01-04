// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/an1l4/mocktest/iuser (interfaces: IUserInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserInterface is a mock of IUserInterface interface.
type MockIUserInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIUserInterfaceMockRecorder
}

// MockIUserInterfaceMockRecorder is the mock recorder for MockIUserInterface.
type MockIUserInterfaceMockRecorder struct {
	mock *MockIUserInterface
}

// NewMockIUserInterface creates a new mock instance.
func NewMockIUserInterface(ctrl *gomock.Controller) *MockIUserInterface {
	mock := &MockIUserInterface{ctrl: ctrl}
	mock.recorder = &MockIUserInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserInterface) EXPECT() *MockIUserInterfaceMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockIUserInterface) AddUser(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockIUserInterfaceMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockIUserInterface)(nil).AddUser), arg0, arg1)
}
