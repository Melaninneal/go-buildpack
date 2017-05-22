// Automatically generated by MockGen. DO NOT EDIT!
// Source: finalize.go

package finalize_test

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
)

// Mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *_MockCommandRecorder
}

// Recorder for MockCommand (not exported)
type _MockCommandRecorder struct {
	mock *MockCommand
}

func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &_MockCommandRecorder{mock}
	return mock
}

func (_m *MockCommand) EXPECT() *_MockCommandRecorder {
	return _m.recorder
}

func (_m *MockCommand) Execute(_param0 string, _param1 io.Writer, _param2 io.Writer, _param3 string, _param4 ...string) error {
	_s := []interface{}{_param0, _param1, _param2, _param3}
	for _, _x := range _param4 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Execute", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCommandRecorder) Execute(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Execute", _s...)
}

// Mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *_MockStagerRecorder
}

// Recorder for MockStager (not exported)
type _MockStagerRecorder struct {
	mock *MockStager
}

func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &_MockStagerRecorder{mock}
	return mock
}

func (_m *MockStager) EXPECT() *_MockStagerRecorder {
	return _m.recorder
}

func (_m *MockStager) BuildDir() string {
	ret := _m.ctrl.Call(_m, "BuildDir")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockStagerRecorder) BuildDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BuildDir")
}

func (_m *MockStager) ClearDepDir() error {
	ret := _m.ctrl.Call(_m, "ClearDepDir")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStagerRecorder) ClearDepDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClearDepDir")
}

func (_m *MockStager) DepDir() string {
	ret := _m.ctrl.Call(_m, "DepDir")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockStagerRecorder) DepDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DepDir")
}

func (_m *MockStager) DepsIdx() string {
	ret := _m.ctrl.Call(_m, "DepsIdx")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockStagerRecorder) DepsIdx() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DepsIdx")
}

func (_m *MockStager) WriteProfileD(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "WriteProfileD", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStagerRecorder) WriteProfileD(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteProfileD", arg0, arg1)
}
