// Code generated by MockGen. DO NOT EDIT.
// Source: tar.go

// Package mocks is a generated GoMock package.
package mocks

import (
	tar "archive/tar"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTarGenerator is a mock of TarGenerator interface.
type MockTarGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockTarGeneratorMockRecorder
}

// MockTarGeneratorMockRecorder is the mock recorder for MockTarGenerator.
type MockTarGeneratorMockRecorder struct {
	mock *MockTarGenerator
}

// NewMockTarGenerator creates a new mock instance.
func NewMockTarGenerator(ctrl *gomock.Controller) *MockTarGenerator {
	mock := &MockTarGenerator{ctrl: ctrl}
	mock.recorder = &MockTarGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTarGenerator) EXPECT() *MockTarGeneratorMockRecorder {
	return m.recorder
}

// WriteTo mocks base method.
func (m *MockTarGenerator) WriteTo(ctx context.Context, writer *tar.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", ctx, writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTo indicates an expected call of WriteTo.
func (mr *MockTarGeneratorMockRecorder) WriteTo(ctx, writer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockTarGenerator)(nil).WriteTo), ctx, writer)
}
