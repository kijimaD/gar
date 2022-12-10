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

// GetCommentByID mocks base method.
func (m *MockclientI) GetCommentByID(commentID int64) *github.PullRequestComment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentByID", commentID)
	ret0, _ := ret[0].(*github.PullRequestComment)
	return ret0
}

// GetCommentByID indicates an expected call of GetCommentByID.
func (mr *MockclientIMockRecorder) GetCommentByID(commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentByID", reflect.TypeOf((*MockclientI)(nil).GetCommentByID), commentID)
}

// GetCommentList mocks base method.
func (m *MockclientI) GetCommentList() []*github.PullRequestComment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentList")
	ret0, _ := ret[0].([]*github.PullRequestComment)
	return ret0
}

// GetCommentList indicates an expected call of GetCommentList.
func (mr *MockclientIMockRecorder) GetCommentList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentList", reflect.TypeOf((*MockclientI)(nil).GetCommentList))
}

// GetPR mocks base method.
func (m *MockclientI) GetPR() PR {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPR")
	ret0, _ := ret[0].(PR)
	return ret0
}

// GetPR indicates an expected call of GetPR.
func (mr *MockclientIMockRecorder) GetPR() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPR", reflect.TypeOf((*MockclientI)(nil).GetPR))
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

// SendReply mocks base method.
func (m *MockclientI) SendReply(arg0 Reply) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendReply", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendReply indicates an expected call of SendReply.
func (mr *MockclientIMockRecorder) SendReply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendReply", reflect.TypeOf((*MockclientI)(nil).SendReply), arg0)
}
