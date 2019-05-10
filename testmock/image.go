// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/go-containerregistry/pkg/v1 (interfaces: Image)

// Package testmock is a generated GoMock package.
package testmock

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	types "github.com/google/go-containerregistry/pkg/v1/types"
	reflect "reflect"
)

// GGCRImage is a mock of Image interface
type GGCRImage struct {
	ctrl     *gomock.Controller
	recorder *GGCRImageMockRecorder
}

// GGCRImageMockRecorder is the mock recorder for GGCRImage
type GGCRImageMockRecorder struct {
	mock *GGCRImage
}

// NewGGCRImage creates a new mock instance
func NewGGCRImage(ctrl *gomock.Controller) *GGCRImage {
	mock := &GGCRImage{ctrl: ctrl}
	mock.recorder = &GGCRImageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *GGCRImage) EXPECT() *GGCRImageMockRecorder {
	return m.recorder
}

// ConfigFile mocks base method
func (m *GGCRImage) ConfigFile() (*v1.ConfigFile, error) {
	ret := m.ctrl.Call(m, "ConfigFile")
	ret0, _ := ret[0].(*v1.ConfigFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigFile indicates an expected call of ConfigFile
func (mr *GGCRImageMockRecorder) ConfigFile() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigFile", reflect.TypeOf((*GGCRImage)(nil).ConfigFile))
}

// ConfigName mocks base method
func (m *GGCRImage) ConfigName() (v1.Hash, error) {
	ret := m.ctrl.Call(m, "ConfigName")
	ret0, _ := ret[0].(v1.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigName indicates an expected call of ConfigName
func (mr *GGCRImageMockRecorder) ConfigName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigName", reflect.TypeOf((*GGCRImage)(nil).ConfigName))
}

// Digest mocks base method
func (m *GGCRImage) Digest() (v1.Hash, error) {
	ret := m.ctrl.Call(m, "Digest")
	ret0, _ := ret[0].(v1.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Digest indicates an expected call of Digest
func (mr *GGCRImageMockRecorder) Digest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Digest", reflect.TypeOf((*GGCRImage)(nil).Digest))
}

// LayerByDiffID mocks base method
func (m *GGCRImage) LayerByDiffID(arg0 v1.Hash) (v1.Layer, error) {
	ret := m.ctrl.Call(m, "LayerByDiffID", arg0)
	ret0, _ := ret[0].(v1.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LayerByDiffID indicates an expected call of LayerByDiffID
func (mr *GGCRImageMockRecorder) LayerByDiffID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerByDiffID", reflect.TypeOf((*GGCRImage)(nil).LayerByDiffID), arg0)
}

// LayerByDigest mocks base method
func (m *GGCRImage) LayerByDigest(arg0 v1.Hash) (v1.Layer, error) {
	ret := m.ctrl.Call(m, "LayerByDigest", arg0)
	ret0, _ := ret[0].(v1.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LayerByDigest indicates an expected call of LayerByDigest
func (mr *GGCRImageMockRecorder) LayerByDigest(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerByDigest", reflect.TypeOf((*GGCRImage)(nil).LayerByDigest), arg0)
}

// Layers mocks base method
func (m *GGCRImage) Layers() ([]v1.Layer, error) {
	ret := m.ctrl.Call(m, "Layers")
	ret0, _ := ret[0].([]v1.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Layers indicates an expected call of Layers
func (mr *GGCRImageMockRecorder) Layers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Layers", reflect.TypeOf((*GGCRImage)(nil).Layers))
}

// Manifest mocks base method
func (m *GGCRImage) Manifest() (*v1.Manifest, error) {
	ret := m.ctrl.Call(m, "Manifest")
	ret0, _ := ret[0].(*v1.Manifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Manifest indicates an expected call of Manifest
func (mr *GGCRImageMockRecorder) Manifest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manifest", reflect.TypeOf((*GGCRImage)(nil).Manifest))
}

// MediaType mocks base method
func (m *GGCRImage) MediaType() (types.MediaType, error) {
	ret := m.ctrl.Call(m, "MediaType")
	ret0, _ := ret[0].(types.MediaType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MediaType indicates an expected call of MediaType
func (mr *GGCRImageMockRecorder) MediaType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MediaType", reflect.TypeOf((*GGCRImage)(nil).MediaType))
}

// RawConfigFile mocks base method
func (m *GGCRImage) RawConfigFile() ([]byte, error) {
	ret := m.ctrl.Call(m, "RawConfigFile")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawConfigFile indicates an expected call of RawConfigFile
func (mr *GGCRImageMockRecorder) RawConfigFile() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawConfigFile", reflect.TypeOf((*GGCRImage)(nil).RawConfigFile))
}

// RawManifest mocks base method
func (m *GGCRImage) RawManifest() ([]byte, error) {
	ret := m.ctrl.Call(m, "RawManifest")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawManifest indicates an expected call of RawManifest
func (mr *GGCRImageMockRecorder) RawManifest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawManifest", reflect.TypeOf((*GGCRImage)(nil).RawManifest))
}
