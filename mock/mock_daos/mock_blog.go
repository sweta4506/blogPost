// Code generated by MockGen. DO NOT EDIT.
// Source: ./blog.go

// Package mock_blog is a generated GoMock package.
package mock_blog

import (
	dtos "blogPost/dtos"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIBlogPost is a mock of IBlogPost interface.
type MockIBlogPost struct {
	ctrl     *gomock.Controller
	recorder *MockIBlogPostMockRecorder
}

// MockIBlogPostMockRecorder is the mock recorder for MockIBlogPost.
type MockIBlogPostMockRecorder struct {
	mock *MockIBlogPost
}

// NewMockIBlogPost creates a new mock instance.
func NewMockIBlogPost(ctrl *gomock.Controller) *MockIBlogPost {
	mock := &MockIBlogPost{ctrl: ctrl}
	mock.recorder = &MockIBlogPostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBlogPost) EXPECT() *MockIBlogPostMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockIBlogPost) CreatePost(req dtos.Post) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", req)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockIBlogPostMockRecorder) CreatePost(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockIBlogPost)(nil).CreatePost), req)
}

// DeletePost mocks base method.
func (m *MockIBlogPost) DeletePost(postID int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockIBlogPostMockRecorder) DeletePost(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockIBlogPost)(nil).DeletePost), postID)
}

// ReadPost mocks base method.
func (m *MockIBlogPost) ReadPost(postID int32) (*dtos.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPost", postID)
	ret0, _ := ret[0].(*dtos.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPost indicates an expected call of ReadPost.
func (mr *MockIBlogPostMockRecorder) ReadPost(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPost", reflect.TypeOf((*MockIBlogPost)(nil).ReadPost), postID)
}

// UpdatePost mocks base method.
func (m *MockIBlogPost) UpdatePost(req *dtos.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockIBlogPostMockRecorder) UpdatePost(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockIBlogPost)(nil).UpdatePost), req)
}
