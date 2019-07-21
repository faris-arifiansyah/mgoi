// Automatically generated by MockGen. DO NOT EDIT!

package mgoimock

import (
	time "time"

	. "github.com/faris-arifiansyah/mgoi"
	"github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
	gomock "github.com/golang/mock/gomock"
)

// Mock of SessionManager interface
type MockSessionManager struct {
	ctrl     *gomock.Controller
	recorder *_MockSessionManagerRecorder
}

// Recorder for MockSessionManager (not exported)
type _MockSessionManagerRecorder struct {
	mock *MockSessionManager
}

func NewMockSessionManager(ctrl *gomock.Controller) *MockSessionManager {
	mock := &MockSessionManager{ctrl: ctrl}
	mock.recorder = &_MockSessionManagerRecorder{mock}
	return mock
}

func (_m *MockSessionManager) EXPECT() *_MockSessionManagerRecorder {
	return _m.recorder
}

func (_m *MockSessionManager) BuildInfo() (mgo.BuildInfo, error) {
	ret := _m.ctrl.Call(_m, "BuildInfo")
	ret0, _ := ret[0].(mgo.BuildInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSessionManagerRecorder) BuildInfo() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BuildInfo")
}

func (_m *MockSessionManager) Clone() SessionManager {
	ret := _m.ctrl.Call(_m, "Clone")
	ret0, _ := ret[0].(SessionManager)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) Clone() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Clone")
}

func (_m *MockSessionManager) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockSessionManagerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockSessionManager) Copy() SessionManager {
	ret := _m.ctrl.Call(_m, "Copy")
	ret0, _ := ret[0].(SessionManager)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) Copy() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Copy")
}

func (_m *MockSessionManager) DB(name string) DatabaseManager {
	ret := _m.ctrl.Call(_m, "DB", name)
	ret0, _ := ret[0].(DatabaseManager)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) DB(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DB", arg0)
}

func (_m *MockSessionManager) DatabaseNames() ([]string, error) {
	ret := _m.ctrl.Call(_m, "DatabaseNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSessionManagerRecorder) DatabaseNames() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseNames")
}

func (_m *MockSessionManager) EnsureSafe(safe *mgo.Safe) {
	_m.ctrl.Call(_m, "EnsureSafe", safe)
}

func (_mr *_MockSessionManagerRecorder) EnsureSafe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EnsureSafe", arg0)
}

func (_m *MockSessionManager) FindRef(ref *mgo.DBRef) QueryManager {
	ret := _m.ctrl.Call(_m, "FindRef", ref)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) FindRef(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindRef", arg0)
}

func (_m *MockSessionManager) Fsync(async bool) error {
	ret := _m.ctrl.Call(_m, "Fsync", async)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) Fsync(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fsync", arg0)
}

func (_m *MockSessionManager) FsyncLock() error {
	ret := _m.ctrl.Call(_m, "FsyncLock")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) FsyncLock() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FsyncLock")
}

func (_m *MockSessionManager) FsyncUnlock() error {
	ret := _m.ctrl.Call(_m, "FsyncUnlock")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) FsyncUnlock() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FsyncUnlock")
}

func (_m *MockSessionManager) LiveServers() []string {
	ret := _m.ctrl.Call(_m, "LiveServers")
	ret0, _ := ret[0].([]string)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) LiveServers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LiveServers")
}

func (_m *MockSessionManager) Login(cred *mgo.Credential) error {
	ret := _m.ctrl.Call(_m, "Login", cred)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) Login(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Login", arg0)
}

func (_m *MockSessionManager) LogoutAll() {
	_m.ctrl.Call(_m, "LogoutAll")
}

func (_mr *_MockSessionManagerRecorder) LogoutAll() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogoutAll")
}

func (_m *MockSessionManager) Mode() mgo.Mode {
	ret := _m.ctrl.Call(_m, "Mode")
	ret0, _ := ret[0].(mgo.Mode)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) Mode() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Mode")
}

func (_m *MockSessionManager) New() SessionManager {
	ret := _m.ctrl.Call(_m, "New")
	ret0, _ := ret[0].(SessionManager)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) New() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "New")
}

func (_m *MockSessionManager) Ping() error {
	ret := _m.ctrl.Call(_m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) Ping() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ping")
}

func (_m *MockSessionManager) Refresh() {
	_m.ctrl.Call(_m, "Refresh")
}

func (_mr *_MockSessionManagerRecorder) Refresh() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Refresh")
}

func (_m *MockSessionManager) ResetIndexCache() {
	_m.ctrl.Call(_m, "ResetIndexCache")
}

func (_mr *_MockSessionManagerRecorder) ResetIndexCache() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResetIndexCache")
}

func (_m *MockSessionManager) Run(cmd interface{}, result interface{}) error {
	ret := _m.ctrl.Call(_m, "Run", cmd, result)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0, arg1)
}

func (_m *MockSessionManager) Safe() *mgo.Safe {
	ret := _m.ctrl.Call(_m, "Safe")
	ret0, _ := ret[0].(*mgo.Safe)
	return ret0
}

func (_mr *_MockSessionManagerRecorder) Safe() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Safe")
}

func (_m *MockSessionManager) SelectServers(tags ...bson.D) {
	_s := []interface{}{}
	for _, _x := range tags {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "SelectServers", _s...)
}

func (_mr *_MockSessionManagerRecorder) SelectServers(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SelectServers", arg0...)
}

func (_m *MockSessionManager) SetBatch(n int) {
	_m.ctrl.Call(_m, "SetBatch", n)
}

func (_mr *_MockSessionManagerRecorder) SetBatch(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetBatch", arg0)
}

func (_m *MockSessionManager) SetBypassValidation(bypass bool) {
	_m.ctrl.Call(_m, "SetBypassValidation", bypass)
}

func (_mr *_MockSessionManagerRecorder) SetBypassValidation(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetBypassValidation", arg0)
}

func (_m *MockSessionManager) SetCursorTimeout(d time.Duration) {
	_m.ctrl.Call(_m, "SetCursorTimeout", d)
}

func (_mr *_MockSessionManagerRecorder) SetCursorTimeout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetCursorTimeout", arg0)
}

func (_m *MockSessionManager) SetMode(consistency mgo.Mode, refresh bool) {
	_m.ctrl.Call(_m, "SetMode", consistency, refresh)
}

func (_mr *_MockSessionManagerRecorder) SetMode(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetMode", arg0, arg1)
}

func (_m *MockSessionManager) SetPoolLimit(limit int) {
	_m.ctrl.Call(_m, "SetPoolLimit", limit)
}

func (_mr *_MockSessionManagerRecorder) SetPoolLimit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetPoolLimit", arg0)
}

func (_m *MockSessionManager) SetPrefetch(p float64) {
	_m.ctrl.Call(_m, "SetPrefetch", p)
}

func (_mr *_MockSessionManagerRecorder) SetPrefetch(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetPrefetch", arg0)
}

func (_m *MockSessionManager) SetSafe(safe *mgo.Safe) {
	_m.ctrl.Call(_m, "SetSafe", safe)
}

func (_mr *_MockSessionManagerRecorder) SetSafe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSafe", arg0)
}

func (_m *MockSessionManager) SetSocketTimeout(d time.Duration) {
	_m.ctrl.Call(_m, "SetSocketTimeout", d)
}

func (_mr *_MockSessionManagerRecorder) SetSocketTimeout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSocketTimeout", arg0)
}

func (_m *MockSessionManager) SetSyncTimeout(d time.Duration) {
	_m.ctrl.Call(_m, "SetSyncTimeout", d)
}

func (_mr *_MockSessionManagerRecorder) SetSyncTimeout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSyncTimeout", arg0)
}
