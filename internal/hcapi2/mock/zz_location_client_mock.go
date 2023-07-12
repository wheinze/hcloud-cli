// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hetznercloud/cli/internal/hcapi2 (interfaces: LocationClient)

// Package hcapi2_mock is a generated GoMock package.
package hcapi2_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hcloud "github.com/hetznercloud/hcloud-go/v2/hcloud"
)

// MockLocationClient is a mock of LocationClient interface.
type MockLocationClient struct {
	ctrl     *gomock.Controller
	recorder *MockLocationClientMockRecorder
}

// MockLocationClientMockRecorder is the mock recorder for MockLocationClient.
type MockLocationClientMockRecorder struct {
	mock *MockLocationClient
}

// NewMockLocationClient creates a new mock instance.
func NewMockLocationClient(ctrl *gomock.Controller) *MockLocationClient {
	mock := &MockLocationClient{ctrl: ctrl}
	mock.recorder = &MockLocationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocationClient) EXPECT() *MockLocationClientMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockLocationClient) All(arg0 context.Context) ([]*hcloud.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", arg0)
	ret0, _ := ret[0].([]*hcloud.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockLocationClientMockRecorder) All(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockLocationClient)(nil).All), arg0)
}

// AllWithOpts mocks base method.
func (m *MockLocationClient) AllWithOpts(arg0 context.Context, arg1 hcloud.LocationListOpts) ([]*hcloud.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllWithOpts", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllWithOpts indicates an expected call of AllWithOpts.
func (mr *MockLocationClientMockRecorder) AllWithOpts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllWithOpts", reflect.TypeOf((*MockLocationClient)(nil).AllWithOpts), arg0, arg1)
}

// Get mocks base method.
func (m *MockLocationClient) Get(arg0 context.Context, arg1 string) (*hcloud.Location, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Location)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockLocationClientMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLocationClient)(nil).Get), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockLocationClient) GetByID(arg0 context.Context, arg1 int64) (*hcloud.Location, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Location)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockLocationClientMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockLocationClient)(nil).GetByID), arg0, arg1)
}

// GetByName mocks base method.
func (m *MockLocationClient) GetByName(arg0 context.Context, arg1 string) (*hcloud.Location, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Location)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByName indicates an expected call of GetByName.
func (mr *MockLocationClientMockRecorder) GetByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockLocationClient)(nil).GetByName), arg0, arg1)
}

// List mocks base method.
func (m *MockLocationClient) List(arg0 context.Context, arg1 hcloud.LocationListOpts) ([]*hcloud.Location, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.Location)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockLocationClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLocationClient)(nil).List), arg0, arg1)
}

// Names mocks base method.
func (m *MockLocationClient) Names() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Names")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Names indicates an expected call of Names.
func (mr *MockLocationClientMockRecorder) Names() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Names", reflect.TypeOf((*MockLocationClient)(nil).Names))
}

// NetworkZones mocks base method.
func (m *MockLocationClient) NetworkZones() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkZones")
	ret0, _ := ret[0].([]string)
	return ret0
}

// NetworkZones indicates an expected call of NetworkZones.
func (mr *MockLocationClientMockRecorder) NetworkZones() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkZones", reflect.TypeOf((*MockLocationClient)(nil).NetworkZones))
}
