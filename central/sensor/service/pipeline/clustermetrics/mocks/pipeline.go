// Code generated by MockGen. DO NOT EDIT.
// Source: pipeline.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/pipeline.go -source pipeline.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	central "github.com/stackrox/rox/generated/internalapi/central"
	gomock "go.uber.org/mock/gomock"
)

// MockMetricsStore is a mock of MetricsStore interface.
type MockMetricsStore struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsStoreMockRecorder
}

// MockMetricsStoreMockRecorder is the mock recorder for MockMetricsStore.
type MockMetricsStoreMockRecorder struct {
	mock *MockMetricsStore
}

// NewMockMetricsStore creates a new mock instance.
func NewMockMetricsStore(ctrl *gomock.Controller) *MockMetricsStore {
	mock := &MockMetricsStore{ctrl: ctrl}
	mock.recorder = &MockMetricsStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricsStore) EXPECT() *MockMetricsStoreMockRecorder {
	return m.recorder
}

// Set mocks base method.
func (m *MockMetricsStore) Set(arg0 string, arg1 *central.ClusterMetrics) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set.
func (mr *MockMetricsStoreMockRecorder) Set(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockMetricsStore)(nil).Set), arg0, arg1)
}
