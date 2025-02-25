// Code generated by MockGen. DO NOT EDIT.
// Source: zip.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/zip.go -source zip.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	zip "archive/zip"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockZipGenerator is a mock of ZipGenerator interface.
type MockZipGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockZipGeneratorMockRecorder
}

// MockZipGeneratorMockRecorder is the mock recorder for MockZipGenerator.
type MockZipGeneratorMockRecorder struct {
	mock *MockZipGenerator
}

// NewMockZipGenerator creates a new mock instance.
func NewMockZipGenerator(ctrl *gomock.Controller) *MockZipGenerator {
	mock := &MockZipGenerator{ctrl: ctrl}
	mock.recorder = &MockZipGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockZipGenerator) EXPECT() *MockZipGeneratorMockRecorder {
	return m.recorder
}

// WriteTo mocks base method.
func (m *MockZipGenerator) WriteTo(ctx context.Context, writer *zip.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", ctx, writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTo indicates an expected call of WriteTo.
func (mr *MockZipGeneratorMockRecorder) WriteTo(ctx, writer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockZipGenerator)(nil).WriteTo), ctx, writer)
}
