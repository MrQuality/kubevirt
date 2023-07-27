// Automatically generated by MockGen. DO NOT EDIT!
// Source: libvirt.go

package cli

import (
	gomock "github.com/golang/mock/gomock"
	libvirt "libvirt.org/go/libvirt"

	api "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
	stats "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/stats"
)

// Mock of Connection interface
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *_MockConnectionRecorder
}

// Recorder for MockConnection (not exported)
type _MockConnectionRecorder struct {
	mock *MockConnection
}

func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &_MockConnectionRecorder{mock}
	return mock
}

func (_m *MockConnection) EXPECT() *_MockConnectionRecorder {
	return _m.recorder
}

func (_m *MockConnection) LookupDomainByName(name string) (VirDomain, error) {
	ret := _m.ctrl.Call(_m, "LookupDomainByName", name)
	ret0, _ := ret[0].(VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) LookupDomainByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LookupDomainByName", arg0)
}

func (_m *MockConnection) DomainDefineXML(xml string) (VirDomain, error) {
	ret := _m.ctrl.Call(_m, "DomainDefineXML", xml)
	ret0, _ := ret[0].(VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) DomainDefineXML(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DomainDefineXML", arg0)
}

func (_m *MockConnection) Close() (int, error) {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockConnection) DomainEventLifecycleRegister(callback libvirt.DomainEventLifecycleCallback) error {
	ret := _m.ctrl.Call(_m, "DomainEventLifecycleRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnectionRecorder) DomainEventLifecycleRegister(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DomainEventLifecycleRegister", arg0)
}

func (_m *MockConnection) DomainEventDeviceAddedRegister(callback libvirt.DomainEventDeviceAddedCallback) error {
	ret := _m.ctrl.Call(_m, "DomainEventDeviceAddedRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnectionRecorder) DomainEventDeviceAddedRegister(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DomainEventDeviceAddedRegister", arg0)
}

func (_m *MockConnection) DomainEventDeviceRemovedRegister(callback libvirt.DomainEventDeviceRemovedCallback) error {
	ret := _m.ctrl.Call(_m, "DomainEventDeviceRemovedRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnectionRecorder) DomainEventDeviceRemovedRegister(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DomainEventDeviceRemovedRegister", arg0)
}

func (_m *MockConnection) AgentEventLifecycleRegister(callback libvirt.DomainEventAgentLifecycleCallback) error {
	ret := _m.ctrl.Call(_m, "AgentEventLifecycleRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnectionRecorder) AgentEventLifecycleRegister(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AgentEventLifecycleRegister", arg0)
}

func (_m *MockConnection) VolatileDomainEventDeviceRemovedRegister(domain VirDomain, callback libvirt.DomainEventDeviceRemovedCallback) (int, error) {
	ret := _m.ctrl.Call(_m, "VolatileDomainEventDeviceRemovedRegister", domain, callback)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) VolatileDomainEventDeviceRemovedRegister(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VolatileDomainEventDeviceRemovedRegister", arg0, arg1)
}

func (_m *MockConnection) DomainEventMemoryDeviceSizeChangeRegister(callback libvirt.DomainEventMemoryDeviceSizeChangeCallback) error {
	ret := _m.ctrl.Call(_m, "DomainEventMemoryDeviceSizeChangeRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnectionRecorder) DomainEventMemoryDeviceSizeChangeRegister(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DomainEventMemoryDeviceSizeChangeRegister", arg0)
}

func (_m *MockConnection) DomainEventDeregister(registrationID int) error {
	ret := _m.ctrl.Call(_m, "DomainEventDeregister", registrationID)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnectionRecorder) DomainEventDeregister(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DomainEventDeregister", arg0)
}

func (_m *MockConnection) ListAllDomains(flags libvirt.ConnectListAllDomainsFlags) ([]VirDomain, error) {
	ret := _m.ctrl.Call(_m, "ListAllDomains", flags)
	ret0, _ := ret[0].([]VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) ListAllDomains(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAllDomains", arg0)
}

func (_m *MockConnection) NewStream(flags libvirt.StreamFlags) (Stream, error) {
	ret := _m.ctrl.Call(_m, "NewStream", flags)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) NewStream(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewStream", arg0)
}

func (_m *MockConnection) SetReconnectChan(reconnect chan bool) {
	_m.ctrl.Call(_m, "SetReconnectChan", reconnect)
}

func (_mr *_MockConnectionRecorder) SetReconnectChan(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetReconnectChan", arg0)
}

func (_m *MockConnection) QemuAgentCommand(command string, domainName string) (string, error) {
	ret := _m.ctrl.Call(_m, "QemuAgentCommand", command, domainName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) QemuAgentCommand(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QemuAgentCommand", arg0, arg1)
}

func (_m *MockConnection) GetAllDomainStats(statsTypes libvirt.DomainStatsTypes, flags libvirt.ConnectGetAllDomainStatsFlags) ([]libvirt.DomainStats, error) {
	ret := _m.ctrl.Call(_m, "GetAllDomainStats", statsTypes, flags)
	ret0, _ := ret[0].([]libvirt.DomainStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) GetAllDomainStats(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllDomainStats", arg0, arg1)
}

func (_m *MockConnection) GetDomainStats(statsTypes libvirt.DomainStatsTypes, l *stats.DomainJobInfo, flags libvirt.ConnectGetAllDomainStatsFlags) ([]*stats.DomainStats, error) {
	ret := _m.ctrl.Call(_m, "GetDomainStats", statsTypes, l, flags)
	ret0, _ := ret[0].([]*stats.DomainStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) GetDomainStats(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDomainStats", arg0, arg1, arg2)
}

func (_m *MockConnection) GetQemuVersion() (string, error) {
	ret := _m.ctrl.Call(_m, "GetQemuVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) GetQemuVersion() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetQemuVersion")
}

func (_m *MockConnection) GetSEVInfo() (*api.SEVNodeParameters, error) {
	ret := _m.ctrl.Call(_m, "GetSEVInfo")
	ret0, _ := ret[0].(*api.SEVNodeParameters)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) GetSEVInfo() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSEVInfo")
}

// Mock of Stream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *_MockStreamRecorder
}

// Recorder for MockStream (not exported)
type _MockStreamRecorder struct {
	mock *MockStream
}

func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &_MockStreamRecorder{mock}
	return mock
}

func (_m *MockStream) EXPECT() *_MockStreamRecorder {
	return _m.recorder
}

func (_m *MockStream) Read(p []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStreamRecorder) Read(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0)
}

func (_m *MockStream) Write(p []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStreamRecorder) Write(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0)
}

func (_m *MockStream) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStreamRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockStream) UnderlyingStream() *libvirt.Stream {
	ret := _m.ctrl.Call(_m, "UnderlyingStream")
	ret0, _ := ret[0].(*libvirt.Stream)
	return ret0
}

func (_mr *_MockStreamRecorder) UnderlyingStream() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnderlyingStream")
}

// Mock of VirDomain interface
type MockVirDomain struct {
	ctrl     *gomock.Controller
	recorder *_MockVirDomainRecorder
}

// Recorder for MockVirDomain (not exported)
type _MockVirDomainRecorder struct {
	mock *MockVirDomain
}

func NewMockVirDomain(ctrl *gomock.Controller) *MockVirDomain {
	mock := &MockVirDomain{ctrl: ctrl}
	mock.recorder = &_MockVirDomainRecorder{mock}
	return mock
}

func (_m *MockVirDomain) EXPECT() *_MockVirDomainRecorder {
	return _m.recorder
}

func (_m *MockVirDomain) GetState() (libvirt.DomainState, int, error) {
	ret := _m.ctrl.Call(_m, "GetState")
	ret0, _ := ret[0].(libvirt.DomainState)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockVirDomainRecorder) GetState() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetState")
}

func (_m *MockVirDomain) Create() error {
	ret := _m.ctrl.Call(_m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Create() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create")
}

func (_m *MockVirDomain) CreateWithFlags(flags libvirt.DomainCreateFlags) error {
	ret := _m.ctrl.Call(_m, "CreateWithFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) CreateWithFlags(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateWithFlags", arg0)
}

func (_m *MockVirDomain) Suspend() error {
	ret := _m.ctrl.Call(_m, "Suspend")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Suspend() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Suspend")
}

func (_m *MockVirDomain) Resume() error {
	ret := _m.ctrl.Call(_m, "Resume")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Resume() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Resume")
}

func (_m *MockVirDomain) BlockResize(disk string, size uint64, flags libvirt.DomainBlockResizeFlags) error {
	ret := _m.ctrl.Call(_m, "BlockResize", disk, size, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) BlockResize(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BlockResize", arg0, arg1, arg2)
}

func (_m *MockVirDomain) GetBlockInfo(disk string, flags uint32) (*libvirt.DomainBlockInfo, error) {
	ret := _m.ctrl.Call(_m, "GetBlockInfo", disk, flags)
	ret0, _ := ret[0].(*libvirt.DomainBlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetBlockInfo(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBlockInfo", arg0, arg1)
}

func (_m *MockVirDomain) AttachDevice(xml string) error {
	ret := _m.ctrl.Call(_m, "AttachDevice", xml)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) AttachDevice(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AttachDevice", arg0)
}

func (_m *MockVirDomain) AttachDeviceFlags(xml string, flags libvirt.DomainDeviceModifyFlags) error {
	ret := _m.ctrl.Call(_m, "AttachDeviceFlags", xml, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) AttachDeviceFlags(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AttachDeviceFlags", arg0, arg1)
}

func (_m *MockVirDomain) DetachDevice(xml string) error {
	ret := _m.ctrl.Call(_m, "DetachDevice", xml)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) DetachDevice(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DetachDevice", arg0)
}

func (_m *MockVirDomain) DetachDeviceFlags(xml string, flags libvirt.DomainDeviceModifyFlags) error {
	ret := _m.ctrl.Call(_m, "DetachDeviceFlags", xml, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) DetachDeviceFlags(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DetachDeviceFlags", arg0, arg1)
}

func (_m *MockVirDomain) DestroyFlags(flags libvirt.DomainDestroyFlags) error {
	ret := _m.ctrl.Call(_m, "DestroyFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) DestroyFlags(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DestroyFlags", arg0)
}

func (_m *MockVirDomain) ShutdownFlags(flags libvirt.DomainShutdownFlags) error {
	ret := _m.ctrl.Call(_m, "ShutdownFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) ShutdownFlags(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShutdownFlags", arg0)
}

func (_m *MockVirDomain) Reboot(flags libvirt.DomainRebootFlagValues) error {
	ret := _m.ctrl.Call(_m, "Reboot", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Reboot(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reboot", arg0)
}

func (_m *MockVirDomain) UndefineFlags(flags libvirt.DomainUndefineFlagsValues) error {
	ret := _m.ctrl.Call(_m, "UndefineFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) UndefineFlags(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UndefineFlags", arg0)
}

func (_m *MockVirDomain) GetName() (string, error) {
	ret := _m.ctrl.Call(_m, "GetName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetName")
}

func (_m *MockVirDomain) GetUUIDString() (string, error) {
	ret := _m.ctrl.Call(_m, "GetUUIDString")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetUUIDString() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUUIDString")
}

func (_m *MockVirDomain) GetXMLDesc(flags libvirt.DomainXMLFlags) (string, error) {
	ret := _m.ctrl.Call(_m, "GetXMLDesc", flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetXMLDesc(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetXMLDesc", arg0)
}

func (_m *MockVirDomain) GetMetadata(tipus libvirt.DomainMetadataType, uri string, flags libvirt.DomainModificationImpact) (string, error) {
	ret := _m.ctrl.Call(_m, "GetMetadata", tipus, uri, flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetadata", arg0, arg1, arg2)
}

func (_m *MockVirDomain) OpenConsole(devname string, stream *libvirt.Stream, flags libvirt.DomainConsoleFlags) error {
	ret := _m.ctrl.Call(_m, "OpenConsole", devname, stream, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) OpenConsole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OpenConsole", arg0, arg1, arg2)
}

func (_m *MockVirDomain) MigrateToURI3(_param0 string, _param1 *libvirt.DomainMigrateParameters, _param2 libvirt.DomainMigrateFlags) error {
	ret := _m.ctrl.Call(_m, "MigrateToURI3", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) MigrateToURI3(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MigrateToURI3", arg0, arg1, arg2)
}

func (_m *MockVirDomain) MigrateStartPostCopy(flags uint32) error {
	ret := _m.ctrl.Call(_m, "MigrateStartPostCopy", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) MigrateStartPostCopy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MigrateStartPostCopy", arg0)
}

func (_m *MockVirDomain) MemoryStats(nrStats uint32, flags uint32) ([]libvirt.DomainMemoryStat, error) {
	ret := _m.ctrl.Call(_m, "MemoryStats", nrStats, flags)
	ret0, _ := ret[0].([]libvirt.DomainMemoryStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) MemoryStats(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MemoryStats", arg0, arg1)
}

func (_m *MockVirDomain) GetJobStats(flags libvirt.DomainGetJobStatsFlags) (*libvirt.DomainJobInfo, error) {
	ret := _m.ctrl.Call(_m, "GetJobStats", flags)
	ret0, _ := ret[0].(*libvirt.DomainJobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetJobStats(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJobStats", arg0)
}

func (_m *MockVirDomain) GetJobInfo() (*libvirt.DomainJobInfo, error) {
	ret := _m.ctrl.Call(_m, "GetJobInfo")
	ret0, _ := ret[0].(*libvirt.DomainJobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetJobInfo() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJobInfo")
}

func (_m *MockVirDomain) GetDiskErrors(flags uint32) ([]libvirt.DomainDiskError, error) {
	ret := _m.ctrl.Call(_m, "GetDiskErrors", flags)
	ret0, _ := ret[0].([]libvirt.DomainDiskError)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetDiskErrors(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDiskErrors", arg0)
}

func (_m *MockVirDomain) SetTime(secs int64, nsecs uint, flags libvirt.DomainSetTimeFlags) error {
	ret := _m.ctrl.Call(_m, "SetTime", secs, nsecs, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) SetTime(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTime", arg0, arg1, arg2)
}

func (_m *MockVirDomain) AuthorizedSSHKeysGet(user string, flags libvirt.DomainAuthorizedSSHKeysFlags) ([]string, error) {
	ret := _m.ctrl.Call(_m, "AuthorizedSSHKeysGet", user, flags)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) AuthorizedSSHKeysGet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AuthorizedSSHKeysGet", arg0, arg1)
}

func (_m *MockVirDomain) AuthorizedSSHKeysSet(user string, keys []string, flags libvirt.DomainAuthorizedSSHKeysFlags) error {
	ret := _m.ctrl.Call(_m, "AuthorizedSSHKeysSet", user, keys, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) AuthorizedSSHKeysSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AuthorizedSSHKeysSet", arg0, arg1, arg2)
}

func (_m *MockVirDomain) AbortJob() error {
	ret := _m.ctrl.Call(_m, "AbortJob")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) AbortJob() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AbortJob")
}

func (_m *MockVirDomain) Free() error {
	ret := _m.ctrl.Call(_m, "Free")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Free() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Free")
}

func (_m *MockVirDomain) CoreDumpWithFormat(to string, format libvirt.DomainCoreDumpFormat, flags libvirt.DomainCoreDumpFlags) error {
	ret := _m.ctrl.Call(_m, "CoreDumpWithFormat", to, format, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) CoreDumpWithFormat(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CoreDumpWithFormat", arg0, arg1, arg2)
}

func (_m *MockVirDomain) PinVcpuFlags(vcpu uint, cpuMap []bool, flags libvirt.DomainModificationImpact) error {
	ret := _m.ctrl.Call(_m, "PinVcpuFlags", vcpu, cpuMap, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) PinVcpuFlags(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PinVcpuFlags", arg0, arg1, arg2)
}

func (_m *MockVirDomain) PinEmulator(cpumap []bool, flags libvirt.DomainModificationImpact) error {
	ret := _m.ctrl.Call(_m, "PinEmulator", cpumap, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) PinEmulator(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PinEmulator", arg0, arg1)
}

func (_m *MockVirDomain) SetVcpusFlags(vcpu uint, flags libvirt.DomainVcpuFlags) error {
	ret := _m.ctrl.Call(_m, "SetVcpusFlags", vcpu, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) SetVcpusFlags(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetVcpusFlags", arg0, arg1)
}

func (_m *MockVirDomain) GetLaunchSecurityInfo(flags uint32) (*libvirt.DomainLaunchSecurityParameters, error) {
	ret := _m.ctrl.Call(_m, "GetLaunchSecurityInfo", flags)
	ret0, _ := ret[0].(*libvirt.DomainLaunchSecurityParameters)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetLaunchSecurityInfo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLaunchSecurityInfo", arg0)
}

func (_m *MockVirDomain) SetLaunchSecurityState(params *libvirt.DomainLaunchSecurityStateParameters, flags uint32) error {
	ret := _m.ctrl.Call(_m, "SetLaunchSecurityState", params, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) SetLaunchSecurityState(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetLaunchSecurityState", arg0, arg1)
}
