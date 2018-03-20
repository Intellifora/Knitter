// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ZTE/Knitter/pkg/db-accessor (interfaces: DbAccessor)

/*
Copyright 2018 ZTE Corporation. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockdbaccessor

import (
	client "github.com/coreos/etcd/client"
	gomock "github.com/golang/mock/gomock"
)

// Mock of DbAccessor interface
type MockDbAccessor struct {
	ctrl     *gomock.Controller
	recorder *MockDbAccessorRecorder
}

// Recorder for MockDbAccessor (not exported)
type MockDbAccessorRecorder struct {
	mock *MockDbAccessor
}

func NewMockDbAccessor(ctrl *gomock.Controller) *MockDbAccessor {
	mock := &MockDbAccessor{ctrl: ctrl}
	mock.recorder = &MockDbAccessorRecorder{mock}
	return mock
}

func (_m *MockDbAccessor) EXPECT() *MockDbAccessorRecorder {
	return _m.recorder
}

func (_m *MockDbAccessor) DeleteDir(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteDir", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *MockDbAccessorRecorder) DeleteDir(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteDir", arg0)
}

func (_m *MockDbAccessor) DeleteLeaf(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteLeaf", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *MockDbAccessorRecorder) DeleteLeaf(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteLeaf", arg0)
}

func (_m *MockDbAccessor) Lock(_param0 string) bool {
	ret := _m.ctrl.Call(_m, "Lock", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *MockDbAccessorRecorder) Lock(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Lock", arg0)
}

func (_m *MockDbAccessor) ReadDir(_param0 string) ([]*client.Node, error) {
	ret := _m.ctrl.Call(_m, "ReadDir", _param0)
	ret0, _ := ret[0].([]*client.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *MockDbAccessorRecorder) ReadDir(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadDir", arg0)
}

func (_m *MockDbAccessor) ReadLeaf(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "ReadLeaf", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *MockDbAccessorRecorder) ReadLeaf(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadLeaf", arg0)
}

func (_m *MockDbAccessor) SaveLeaf(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "SaveLeaf", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *MockDbAccessorRecorder) SaveLeaf(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveLeaf", arg0, arg1)
}

func (_m *MockDbAccessor) Unlock(_param0 string) bool {
	ret := _m.ctrl.Call(_m, "Unlock", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *MockDbAccessorRecorder) Unlock(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unlock", arg0)
}

func (_m *MockDbAccessor) WatcherDir(_param0 string) (*client.Response, error) {
	ret := _m.ctrl.Call(_m, "WatcherDir", _param0)
	ret0, _ := ret[0].(*client.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *MockDbAccessorRecorder) WatcherDir(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WatcherDir", arg0)
}