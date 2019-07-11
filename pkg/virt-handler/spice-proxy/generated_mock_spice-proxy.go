// Automatically generated by MockGen. DO NOT EDIT!
// Source: spice-proxy.go

package spice_proxy

import (
	gomock "github.com/golang/mock/gomock"

	v1 "kubevirt.io/client-go/api/v1"
)

// Mock of UnixToTcpProxyManager interface
type MockUnixToTcpProxyManager struct {
	ctrl     *gomock.Controller
	recorder *_MockUnixToTcpProxyManagerRecorder
}

// Recorder for MockUnixToTcpProxyManager (not exported)
type _MockUnixToTcpProxyManagerRecorder struct {
	mock *MockUnixToTcpProxyManager
}

func NewMockUnixToTcpProxyManager(ctrl *gomock.Controller) *MockUnixToTcpProxyManager {
	mock := &MockUnixToTcpProxyManager{ctrl: ctrl}
	mock.recorder = &_MockUnixToTcpProxyManagerRecorder{mock}
	return mock
}

func (_m *MockUnixToTcpProxyManager) EXPECT() *_MockUnixToTcpProxyManagerRecorder {
	return _m.recorder
}

func (_m *MockUnixToTcpProxyManager) StartListener(vmNamespace string, vmName string, vmi *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "StartListener", vmNamespace, vmName, vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUnixToTcpProxyManagerRecorder) StartListener(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartListener", arg0, arg1, arg2)
}

func (_m *MockUnixToTcpProxyManager) GetPort(vmNamespace string, vmName string) (int32, error) {
	ret := _m.ctrl.Call(_m, "GetPort", vmNamespace, vmName)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockUnixToTcpProxyManagerRecorder) GetPort(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPort", arg0, arg1)
}

func (_m *MockUnixToTcpProxyManager) StopListener(vmNamespace string, vmName string) {
	_m.ctrl.Call(_m, "StopListener", vmNamespace, vmName)
}

func (_mr *_MockUnixToTcpProxyManagerRecorder) StopListener(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopListener", arg0, arg1)
}
