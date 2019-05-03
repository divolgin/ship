// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/replicatedhq/ship/pkg/lifecycle/render/helm (interfaces: Commands)

// Package helm is a generated GoMock package.
package helm

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	chartutil "k8s.io/helm/pkg/chartutil"
)

// MockCommands is a mock of Commands interface
type MockCommands struct {
	ctrl     *gomock.Controller
	recorder *MockCommandsMockRecorder
}

// MockCommandsMockRecorder is the mock recorder for MockCommands
type MockCommandsMockRecorder struct {
	mock *MockCommands
}

// NewMockCommands creates a new mock instance
func NewMockCommands(ctrl *gomock.Controller) *MockCommands {
	mock := &MockCommands{ctrl: ctrl}
	mock.recorder = &MockCommandsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommands) EXPECT() *MockCommandsMockRecorder {
	return m.recorder
}

// Fetch mocks base method
func (m *MockCommands) Fetch(arg0, arg1, arg2, arg3, arg4 string) error {
	ret := m.ctrl.Call(m, "Fetch", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fetch indicates an expected call of Fetch
func (mr *MockCommandsMockRecorder) Fetch(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockCommands)(nil).Fetch), arg0, arg1, arg2, arg3, arg4)
}

// Init mocks base method
func (m *MockCommands) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockCommandsMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockCommands)(nil).Init))
}

// MaybeDependencyUpdate mocks base method
func (m *MockCommands) MaybeDependencyUpdate(arg0 string, arg1 chartutil.Requirements) error {
	ret := m.ctrl.Call(m, "MaybeDependencyUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MaybeDependencyUpdate indicates an expected call of MaybeDependencyUpdate
func (mr *MockCommandsMockRecorder) MaybeDependencyUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeDependencyUpdate", reflect.TypeOf((*MockCommands)(nil).MaybeDependencyUpdate), arg0, arg1)
}

// RepoAdd mocks base method
func (m *MockCommands) RepoAdd(arg0, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "RepoAdd", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RepoAdd indicates an expected call of RepoAdd
func (mr *MockCommandsMockRecorder) RepoAdd(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepoAdd", reflect.TypeOf((*MockCommands)(nil).RepoAdd), arg0, arg1, arg2)
}

// Template mocks base method
func (m *MockCommands) Template(arg0 string, arg1 []string) error {
	ret := m.ctrl.Call(m, "Template", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Template indicates an expected call of Template
func (mr *MockCommandsMockRecorder) Template(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockCommands)(nil).Template), arg0, arg1)
}
