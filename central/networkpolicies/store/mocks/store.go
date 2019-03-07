// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/networkpolicies/store (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddNetworkPolicy mocks base method
func (m *MockStore) AddNetworkPolicy(arg0 *storage.NetworkPolicy) error {
	ret := m.ctrl.Call(m, "AddNetworkPolicy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNetworkPolicy indicates an expected call of AddNetworkPolicy
func (mr *MockStoreMockRecorder) AddNetworkPolicy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNetworkPolicy", reflect.TypeOf((*MockStore)(nil).AddNetworkPolicy), arg0)
}

// CountMatchingNetworkPolicies mocks base method
func (m *MockStore) CountMatchingNetworkPolicies(arg0, arg1 string) (int, error) {
	ret := m.ctrl.Call(m, "CountMatchingNetworkPolicies", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountMatchingNetworkPolicies indicates an expected call of CountMatchingNetworkPolicies
func (mr *MockStoreMockRecorder) CountMatchingNetworkPolicies(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountMatchingNetworkPolicies", reflect.TypeOf((*MockStore)(nil).CountMatchingNetworkPolicies), arg0, arg1)
}

// CountNetworkPolicies mocks base method
func (m *MockStore) CountNetworkPolicies() (int, error) {
	ret := m.ctrl.Call(m, "CountNetworkPolicies")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountNetworkPolicies indicates an expected call of CountNetworkPolicies
func (mr *MockStoreMockRecorder) CountNetworkPolicies() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountNetworkPolicies", reflect.TypeOf((*MockStore)(nil).CountNetworkPolicies))
}

// GetNetworkPolicies mocks base method
func (m *MockStore) GetNetworkPolicies(arg0, arg1 string) ([]*storage.NetworkPolicy, error) {
	ret := m.ctrl.Call(m, "GetNetworkPolicies", arg0, arg1)
	ret0, _ := ret[0].([]*storage.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkPolicies indicates an expected call of GetNetworkPolicies
func (mr *MockStoreMockRecorder) GetNetworkPolicies(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkPolicies", reflect.TypeOf((*MockStore)(nil).GetNetworkPolicies), arg0, arg1)
}

// GetNetworkPolicy mocks base method
func (m *MockStore) GetNetworkPolicy(arg0 string) (*storage.NetworkPolicy, bool, error) {
	ret := m.ctrl.Call(m, "GetNetworkPolicy", arg0)
	ret0, _ := ret[0].(*storage.NetworkPolicy)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNetworkPolicy indicates an expected call of GetNetworkPolicy
func (mr *MockStoreMockRecorder) GetNetworkPolicy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkPolicy", reflect.TypeOf((*MockStore)(nil).GetNetworkPolicy), arg0)
}

// RemoveNetworkPolicy mocks base method
func (m *MockStore) RemoveNetworkPolicy(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveNetworkPolicy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNetworkPolicy indicates an expected call of RemoveNetworkPolicy
func (mr *MockStoreMockRecorder) RemoveNetworkPolicy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNetworkPolicy", reflect.TypeOf((*MockStore)(nil).RemoveNetworkPolicy), arg0)
}

// UpdateNetworkPolicy mocks base method
func (m *MockStore) UpdateNetworkPolicy(arg0 *storage.NetworkPolicy) error {
	ret := m.ctrl.Call(m, "UpdateNetworkPolicy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNetworkPolicy indicates an expected call of UpdateNetworkPolicy
func (mr *MockStoreMockRecorder) UpdateNetworkPolicy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNetworkPolicy", reflect.TypeOf((*MockStore)(nil).UpdateNetworkPolicy), arg0)
}
