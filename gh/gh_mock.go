// Code generated by MockGen. DO NOT EDIT.
// Source: gh.go

// Package gh is a generated GoMock package.
package gh

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockclientI is a mock of clientI interface.
type MockclientI struct {
	ctrl     *gomock.Controller
	recorder *MockclientIMockRecorder
}

// MockclientIMockRecorder is the mock recorder for MockclientI.
type MockclientIMockRecorder struct {
	mock *MockclientI
}

// NewMockclientI creates a new mock instance.
func NewMockclientI(ctrl *gomock.Controller) *MockclientI {
	mock := &MockclientI{ctrl: ctrl}
	mock.recorder = &MockclientIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockclientI) EXPECT() *MockclientIMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockclientI) List() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "List")
}

// List indicates an expected call of List.
func (mr *MockclientIMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockclientI)(nil).List))
}