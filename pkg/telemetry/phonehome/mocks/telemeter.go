// Code generated by MockGen. DO NOT EDIT.
// Source: telemeter.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTelemeter is a mock of Telemeter interface.
type MockTelemeter struct {
	ctrl     *gomock.Controller
	recorder *MockTelemeterMockRecorder
}

// MockTelemeterMockRecorder is the mock recorder for MockTelemeter.
type MockTelemeterMockRecorder struct {
	mock *MockTelemeter
}

// NewMockTelemeter creates a new mock instance.
func NewMockTelemeter(ctrl *gomock.Controller) *MockTelemeter {
	mock := &MockTelemeter{ctrl: ctrl}
	mock.recorder = &MockTelemeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelemeter) EXPECT() *MockTelemeterMockRecorder {
	return m.recorder
}

// Group mocks base method.
func (m *MockTelemeter) Group(groupID string, props map[string]any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Group", groupID, props)
}

// Group indicates an expected call of Group.
func (mr *MockTelemeterMockRecorder) Group(groupID, props interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockTelemeter)(nil).Group), groupID, props)
}

// GroupUserAs mocks base method.
func (m *MockTelemeter) GroupUserAs(userID, clientID, clientType, groupID string, props map[string]any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GroupUserAs", userID, clientID, clientType, groupID, props)
}

// GroupUserAs indicates an expected call of GroupUserAs.
func (mr *MockTelemeterMockRecorder) GroupUserAs(userID, clientID, clientType, groupID, props interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupUserAs", reflect.TypeOf((*MockTelemeter)(nil).GroupUserAs), userID, clientID, clientType, groupID, props)
}

// Identify mocks base method.
func (m *MockTelemeter) Identify(props map[string]any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Identify", props)
}

// Identify indicates an expected call of Identify.
func (mr *MockTelemeterMockRecorder) Identify(props interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identify", reflect.TypeOf((*MockTelemeter)(nil).Identify), props)
}

// IdentifyUserAs mocks base method.
func (m *MockTelemeter) IdentifyUserAs(userID, clientID, clientType string, props map[string]any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IdentifyUserAs", userID, clientID, clientType, props)
}

// IdentifyUserAs indicates an expected call of IdentifyUserAs.
func (mr *MockTelemeterMockRecorder) IdentifyUserAs(userID, clientID, clientType, props interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentifyUserAs", reflect.TypeOf((*MockTelemeter)(nil).IdentifyUserAs), userID, clientID, clientType, props)
}

// Stop mocks base method.
func (m *MockTelemeter) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockTelemeterMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTelemeter)(nil).Stop))
}

// Track mocks base method.
func (m *MockTelemeter) Track(event string, props map[string]any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Track", event, props)
}

// Track indicates an expected call of Track.
func (mr *MockTelemeterMockRecorder) Track(event, props interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Track", reflect.TypeOf((*MockTelemeter)(nil).Track), event, props)
}

// TrackUserAs mocks base method.
func (m *MockTelemeter) TrackUserAs(userID, clientID, clientType, event string, props map[string]any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TrackUserAs", userID, clientID, clientType, event, props)
}

// TrackUserAs indicates an expected call of TrackUserAs.
func (mr *MockTelemeterMockRecorder) TrackUserAs(userID, clientID, clientType, event, props interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackUserAs", reflect.TypeOf((*MockTelemeter)(nil).TrackUserAs), userID, clientID, clientType, event, props)
}
