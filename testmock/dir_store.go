// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/lifecycle (interfaces: DirStore)

// Package testmock is a generated GoMock package.
package testmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	buildpack "github.com/buildpacks/lifecycle/buildpack"
)

// MockDirStore is a mock of DirStore interface.
type MockDirStore struct {
	ctrl     *gomock.Controller
	recorder *MockDirStoreMockRecorder
}

// MockDirStoreMockRecorder is the mock recorder for MockDirStore.
type MockDirStoreMockRecorder struct {
	mock *MockDirStore
}

// NewMockDirStore creates a new mock instance.
func NewMockDirStore(ctrl *gomock.Controller) *MockDirStore {
	mock := &MockDirStore{ctrl: ctrl}
	mock.recorder = &MockDirStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirStore) EXPECT() *MockDirStoreMockRecorder {
	return m.recorder
}

// Lookup mocks base method.
func (m *MockDirStore) Lookup(arg0, arg1, arg2 string) (buildpack.Descriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", arg0, arg1, arg2)
	ret0, _ := ret[0].(buildpack.Descriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup.
func (mr *MockDirStoreMockRecorder) Lookup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockDirStore)(nil).Lookup), arg0, arg1, arg2)
}

// LookupBp mocks base method.
func (m *MockDirStore) LookupBp(arg0, arg1 string) (*buildpack.BpDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupBp", arg0, arg1)
	ret0, _ := ret[0].(*buildpack.BpDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupBp indicates an expected call of LookupBp.
func (mr *MockDirStoreMockRecorder) LookupBp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupBp", reflect.TypeOf((*MockDirStore)(nil).LookupBp), arg0, arg1)
}

// LookupExt mocks base method.
func (m *MockDirStore) LookupExt(arg0, arg1 string) (*buildpack.ExtDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupExt", arg0, arg1)
	ret0, _ := ret[0].(*buildpack.ExtDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupExt indicates an expected call of LookupExt.
func (mr *MockDirStoreMockRecorder) LookupExt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupExt", reflect.TypeOf((*MockDirStore)(nil).LookupExt), arg0, arg1)
}
