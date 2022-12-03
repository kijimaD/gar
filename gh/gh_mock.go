// Code generated by MockGen. DO NOT EDIT.
// Source: gh.go

// Package gh is a generated GoMock package.
package gh

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v48/github"
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

// PRCommits mocks base method.
func (m *MockclientI) PRCommits() []*github.RepositoryCommit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PRCommits")
	ret0, _ := ret[0].([]*github.RepositoryCommit)
	return ret0
}

// PRCommits indicates an expected call of PRCommits.
func (mr *MockclientIMockRecorder) PRCommits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PRCommits", reflect.TypeOf((*MockclientI)(nil).PRCommits))
}

// PRDetail mocks base method.
func (m *MockclientI) PRDetail() *github.PullRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PRDetail")
	ret0, _ := ret[0].(*github.PullRequest)
	return ret0
}

// PRDetail indicates an expected call of PRDetail.
func (mr *MockclientIMockRecorder) PRDetail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PRDetail", reflect.TypeOf((*MockclientI)(nil).PRDetail))
}

// Reply mocks base method.
func (m *MockclientI) Reply() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reply")
}

// Reply indicates an expected call of Reply.
func (mr *MockclientIMockRecorder) Reply() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reply", reflect.TypeOf((*MockclientI)(nil).Reply))
}
