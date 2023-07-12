// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hetznercloud/cli/internal/state (interfaces: ActionWaiter,TokenEnsurer)

// Package state is a generated GoMock package.
package state

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hcloud "github.com/hetznercloud/hcloud-go/v2/hcloud"
	cobra "github.com/spf13/cobra"
)

// MockActionWaiter is a mock of ActionWaiter interface.
type MockActionWaiter struct {
	ctrl     *gomock.Controller
	recorder *MockActionWaiterMockRecorder
}

// MockActionWaiterMockRecorder is the mock recorder for MockActionWaiter.
type MockActionWaiterMockRecorder struct {
	mock *MockActionWaiter
}

// NewMockActionWaiter creates a new mock instance.
func NewMockActionWaiter(ctrl *gomock.Controller) *MockActionWaiter {
	mock := &MockActionWaiter{ctrl: ctrl}
	mock.recorder = &MockActionWaiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActionWaiter) EXPECT() *MockActionWaiterMockRecorder {
	return m.recorder
}

// ActionProgress mocks base method.
func (m *MockActionWaiter) ActionProgress(arg0 context.Context, arg1 *hcloud.Action) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionProgress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActionProgress indicates an expected call of ActionProgress.
func (mr *MockActionWaiterMockRecorder) ActionProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionProgress", reflect.TypeOf((*MockActionWaiter)(nil).ActionProgress), arg0, arg1)
}

// WaitForActions mocks base method.
func (m *MockActionWaiter) WaitForActions(arg0 context.Context, arg1 []*hcloud.Action) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForActions", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForActions indicates an expected call of WaitForActions.
func (mr *MockActionWaiterMockRecorder) WaitForActions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForActions", reflect.TypeOf((*MockActionWaiter)(nil).WaitForActions), arg0, arg1)
}

// MockTokenEnsurer is a mock of TokenEnsurer interface.
type MockTokenEnsurer struct {
	ctrl     *gomock.Controller
	recorder *MockTokenEnsurerMockRecorder
}

// MockTokenEnsurerMockRecorder is the mock recorder for MockTokenEnsurer.
type MockTokenEnsurerMockRecorder struct {
	mock *MockTokenEnsurer
}

// NewMockTokenEnsurer creates a new mock instance.
func NewMockTokenEnsurer(ctrl *gomock.Controller) *MockTokenEnsurer {
	mock := &MockTokenEnsurer{ctrl: ctrl}
	mock.recorder = &MockTokenEnsurerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenEnsurer) EXPECT() *MockTokenEnsurerMockRecorder {
	return m.recorder
}

// EnsureToken mocks base method.
func (m *MockTokenEnsurer) EnsureToken(arg0 *cobra.Command, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureToken indicates an expected call of EnsureToken.
func (mr *MockTokenEnsurerMockRecorder) EnsureToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureToken", reflect.TypeOf((*MockTokenEnsurer)(nil).EnsureToken), arg0, arg1)
}
