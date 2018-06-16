// Code generated by MockGen. DO NOT EDIT.
// Source: jail/cell.go

// Package jail is a generated GoMock package.
package jail

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockManager) Call(chatID, this, args string) string {
	ret := m.ctrl.Call(m, "Call", chatID, this, args)
	ret0, _ := ret[0].(string)
	return ret0
}

// Call indicates an expected call of Call
func (mr *MockManagerMockRecorder) Call(chatID, this, args interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockManager)(nil).Call), chatID, this, args)
}

// CreateCell mocks base method
func (m *MockManager) CreateCell(chatID string) (JSCell, error) {
	ret := m.ctrl.Call(m, "CreateCell", chatID)
	ret0, _ := ret[0].(JSCell)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCell indicates an expected call of CreateCell
func (mr *MockManagerMockRecorder) CreateCell(chatID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCell", reflect.TypeOf((*MockManager)(nil).CreateCell), chatID)
}

// Parse mocks base method
func (m *MockManager) Parse(chatID, js string) string {
	ret := m.ctrl.Call(m, "Parse", chatID, js)
	ret0, _ := ret[0].(string)
	return ret0
}

// Parse indicates an expected call of Parse
func (mr *MockManagerMockRecorder) Parse(chatID, js interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockManager)(nil).Parse), chatID, js)
}

// CreateAndInitCell mocks base method
func (m *MockManager) CreateAndInitCell(chatID string, code ...string) string {
	varargs := []interface{}{chatID}
	for _, a := range code {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAndInitCell", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateAndInitCell indicates an expected call of CreateAndInitCell
func (mr *MockManagerMockRecorder) CreateAndInitCell(chatID interface{}, code ...interface{}) *gomock.Call {
	varargs := append([]interface{}{chatID}, code...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAndInitCell", reflect.TypeOf((*MockManager)(nil).CreateAndInitCell), varargs...)
}

// Cell mocks base method
func (m *MockManager) Cell(chatID string) (JSCell, error) {
	ret := m.ctrl.Call(m, "Cell", chatID)
	ret0, _ := ret[0].(JSCell)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cell indicates an expected call of Cell
func (mr *MockManagerMockRecorder) Cell(chatID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cell", reflect.TypeOf((*MockManager)(nil).Cell), chatID)
}

// Execute mocks base method
func (m *MockManager) Execute(chatID, code string) string {
	ret := m.ctrl.Call(m, "Execute", chatID, code)
	ret0, _ := ret[0].(string)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockManagerMockRecorder) Execute(chatID, code interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockManager)(nil).Execute), chatID, code)
}

// SetBaseJS mocks base method
func (m *MockManager) SetBaseJS(js string) {
	m.ctrl.Call(m, "SetBaseJS", js)
}

// SetBaseJS indicates an expected call of SetBaseJS
func (mr *MockManagerMockRecorder) SetBaseJS(js interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBaseJS", reflect.TypeOf((*MockManager)(nil).SetBaseJS), js)
}

// Stop mocks base method
func (m *MockManager) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockManagerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockManager)(nil).Stop))
}

// MockJSCell is a mock of JSCell interface
type MockJSCell struct {
	ctrl     *gomock.Controller
	recorder *MockJSCellMockRecorder
}

// MockJSCellMockRecorder is the mock recorder for MockJSCell
type MockJSCellMockRecorder struct {
	mock *MockJSCell
}

// NewMockJSCell creates a new mock instance
func NewMockJSCell(ctrl *gomock.Controller) *MockJSCell {
	mock := &MockJSCell{ctrl: ctrl}
	mock.recorder = &MockJSCellMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJSCell) EXPECT() *MockJSCellMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockJSCell) Set(arg0 string, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockJSCellMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockJSCell)(nil).Set), arg0, arg1)
}

// Get mocks base method
func (m *MockJSCell) Get(arg0 string) (JSValue, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(JSValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockJSCellMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockJSCell)(nil).Get), arg0)
}

// Run mocks base method
func (m *MockJSCell) Run(arg0 interface{}) (JSValue, error) {
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(JSValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockJSCellMockRecorder) Run(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockJSCell)(nil).Run), arg0)
}

// Call mocks base method
func (m *MockJSCell) Call(item string, this interface{}, args ...interface{}) (JSValue, error) {
	varargs := []interface{}{item, this}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Call", varargs...)
	ret0, _ := ret[0].(JSValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockJSCellMockRecorder) Call(item, this interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{item, this}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockJSCell)(nil).Call), varargs...)
}

// Stop mocks base method
func (m *MockJSCell) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockJSCellMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockJSCell)(nil).Stop))
}
