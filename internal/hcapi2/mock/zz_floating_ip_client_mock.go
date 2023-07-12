// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hetznercloud/cli/internal/hcapi2 (interfaces: FloatingIPClient)

// Package hcapi2_mock is a generated GoMock package.
package hcapi2_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hcloud "github.com/hetznercloud/hcloud-go/v2/hcloud"
)

// MockFloatingIPClient is a mock of FloatingIPClient interface.
type MockFloatingIPClient struct {
	ctrl     *gomock.Controller
	recorder *MockFloatingIPClientMockRecorder
}

// MockFloatingIPClientMockRecorder is the mock recorder for MockFloatingIPClient.
type MockFloatingIPClientMockRecorder struct {
	mock *MockFloatingIPClient
}

// NewMockFloatingIPClient creates a new mock instance.
func NewMockFloatingIPClient(ctrl *gomock.Controller) *MockFloatingIPClient {
	mock := &MockFloatingIPClient{ctrl: ctrl}
	mock.recorder = &MockFloatingIPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFloatingIPClient) EXPECT() *MockFloatingIPClientMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockFloatingIPClient) All(arg0 context.Context) ([]*hcloud.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", arg0)
	ret0, _ := ret[0].([]*hcloud.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockFloatingIPClientMockRecorder) All(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockFloatingIPClient)(nil).All), arg0)
}

// AllWithOpts mocks base method.
func (m *MockFloatingIPClient) AllWithOpts(arg0 context.Context, arg1 hcloud.FloatingIPListOpts) ([]*hcloud.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllWithOpts", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllWithOpts indicates an expected call of AllWithOpts.
func (mr *MockFloatingIPClientMockRecorder) AllWithOpts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllWithOpts", reflect.TypeOf((*MockFloatingIPClient)(nil).AllWithOpts), arg0, arg1)
}

// Assign mocks base method.
func (m *MockFloatingIPClient) Assign(arg0 context.Context, arg1 *hcloud.FloatingIP, arg2 *hcloud.Server) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Assign", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Assign indicates an expected call of Assign.
func (mr *MockFloatingIPClientMockRecorder) Assign(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Assign", reflect.TypeOf((*MockFloatingIPClient)(nil).Assign), arg0, arg1, arg2)
}

// ChangeDNSPtr mocks base method.
func (m *MockFloatingIPClient) ChangeDNSPtr(arg0 context.Context, arg1 *hcloud.FloatingIP, arg2 string, arg3 *string) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeDNSPtr", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeDNSPtr indicates an expected call of ChangeDNSPtr.
func (mr *MockFloatingIPClientMockRecorder) ChangeDNSPtr(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeDNSPtr", reflect.TypeOf((*MockFloatingIPClient)(nil).ChangeDNSPtr), arg0, arg1, arg2, arg3)
}

// ChangeProtection mocks base method.
func (m *MockFloatingIPClient) ChangeProtection(arg0 context.Context, arg1 *hcloud.FloatingIP, arg2 hcloud.FloatingIPChangeProtectionOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeProtection", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeProtection indicates an expected call of ChangeProtection.
func (mr *MockFloatingIPClientMockRecorder) ChangeProtection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeProtection", reflect.TypeOf((*MockFloatingIPClient)(nil).ChangeProtection), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockFloatingIPClient) Create(arg0 context.Context, arg1 hcloud.FloatingIPCreateOpts) (hcloud.FloatingIPCreateResult, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(hcloud.FloatingIPCreateResult)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockFloatingIPClientMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFloatingIPClient)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockFloatingIPClient) Delete(arg0 context.Context, arg1 *hcloud.FloatingIP) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockFloatingIPClientMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFloatingIPClient)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockFloatingIPClient) Get(arg0 context.Context, arg1 string) (*hcloud.FloatingIP, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.FloatingIP)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockFloatingIPClientMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFloatingIPClient)(nil).Get), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockFloatingIPClient) GetByID(arg0 context.Context, arg1 int64) (*hcloud.FloatingIP, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.FloatingIP)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockFloatingIPClientMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockFloatingIPClient)(nil).GetByID), arg0, arg1)
}

// GetByName mocks base method.
func (m *MockFloatingIPClient) GetByName(arg0 context.Context, arg1 string) (*hcloud.FloatingIP, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.FloatingIP)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByName indicates an expected call of GetByName.
func (mr *MockFloatingIPClientMockRecorder) GetByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockFloatingIPClient)(nil).GetByName), arg0, arg1)
}

// LabelKeys mocks base method.
func (m *MockFloatingIPClient) LabelKeys(arg0 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LabelKeys", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// LabelKeys indicates an expected call of LabelKeys.
func (mr *MockFloatingIPClientMockRecorder) LabelKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LabelKeys", reflect.TypeOf((*MockFloatingIPClient)(nil).LabelKeys), arg0)
}

// List mocks base method.
func (m *MockFloatingIPClient) List(arg0 context.Context, arg1 hcloud.FloatingIPListOpts) ([]*hcloud.FloatingIP, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.FloatingIP)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockFloatingIPClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFloatingIPClient)(nil).List), arg0, arg1)
}

// Names mocks base method.
func (m *MockFloatingIPClient) Names() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Names")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Names indicates an expected call of Names.
func (mr *MockFloatingIPClientMockRecorder) Names() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Names", reflect.TypeOf((*MockFloatingIPClient)(nil).Names))
}

// Unassign mocks base method.
func (m *MockFloatingIPClient) Unassign(arg0 context.Context, arg1 *hcloud.FloatingIP) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unassign", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Unassign indicates an expected call of Unassign.
func (mr *MockFloatingIPClientMockRecorder) Unassign(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unassign", reflect.TypeOf((*MockFloatingIPClient)(nil).Unassign), arg0, arg1)
}

// Update mocks base method.
func (m *MockFloatingIPClient) Update(arg0 context.Context, arg1 *hcloud.FloatingIP, arg2 hcloud.FloatingIPUpdateOpts) (*hcloud.FloatingIP, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.FloatingIP)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockFloatingIPClientMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockFloatingIPClient)(nil).Update), arg0, arg1, arg2)
}
