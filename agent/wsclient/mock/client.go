// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/wsclient (interfaces: ClientServer)

// Package mock_wsclient is a generated GoMock package.
package mock_wsclient

import (
	reflect "reflect"
	time "time"

	wsclient "github.com/aws/amazon-ecs-agent/agent/wsclient"
	wsconn "github.com/aws/amazon-ecs-agent/agent/wsclient/wsconn"
	gomock "github.com/golang/mock/gomock"
)

// MockClientServer is a mock of ClientServer interface
type MockClientServer struct {
	ctrl     *gomock.Controller
	recorder *MockClientServerMockRecorder
}

// MockClientServerMockRecorder is the mock recorder for MockClientServer
type MockClientServerMockRecorder struct {
	mock *MockClientServer
}

// NewMockClientServer creates a new mock instance
func NewMockClientServer(ctrl *gomock.Controller) *MockClientServer {
	mock := &MockClientServer{ctrl: ctrl}
	mock.recorder = &MockClientServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientServer) EXPECT() *MockClientServerMockRecorder {
	return m.recorder
}

// AddRequestHandler mocks base method
func (m *MockClientServer) AddRequestHandler(arg0 wsclient.RequestHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddRequestHandler", arg0)
}

// AddRequestHandler indicates an expected call of AddRequestHandler
func (mr *MockClientServerMockRecorder) AddRequestHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRequestHandler", reflect.TypeOf((*MockClientServer)(nil).AddRequestHandler), arg0)
}

// Close mocks base method
func (m *MockClientServer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClientServerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientServer)(nil).Close))
}

// Connect mocks base method
func (m *MockClientServer) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockClientServerMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClientServer)(nil).Connect))
}

// Disconnect mocks base method
func (m *MockClientServer) Disconnect(arg0 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Disconnect", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockClientServerMockRecorder) Disconnect(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockClientServer)(nil).Disconnect), arg0...)
}

// IsConnected mocks base method
func (m *MockClientServer) IsConnected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConnected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConnected indicates an expected call of IsConnected
func (mr *MockClientServerMockRecorder) IsConnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnected", reflect.TypeOf((*MockClientServer)(nil).IsConnected))
}

// MakeRequest mocks base method
func (m *MockClientServer) MakeRequest(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRequest", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeRequest indicates an expected call of MakeRequest
func (mr *MockClientServerMockRecorder) MakeRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRequest", reflect.TypeOf((*MockClientServer)(nil).MakeRequest), arg0)
}

// Serve mocks base method
func (m *MockClientServer) Serve() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serve")
	ret0, _ := ret[0].(error)
	return ret0
}

// Serve indicates an expected call of Serve
func (mr *MockClientServerMockRecorder) Serve() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockClientServer)(nil).Serve))
}

// SetAnyRequestHandler mocks base method
func (m *MockClientServer) SetAnyRequestHandler(arg0 wsclient.RequestHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnyRequestHandler", arg0)
}

// SetAnyRequestHandler indicates an expected call of SetAnyRequestHandler
func (mr *MockClientServerMockRecorder) SetAnyRequestHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnyRequestHandler", reflect.TypeOf((*MockClientServer)(nil).SetAnyRequestHandler), arg0)
}

// SetConnection mocks base method
func (m *MockClientServer) SetConnection(arg0 wsconn.WebsocketConn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConnection", arg0)
}

// SetConnection indicates an expected call of SetConnection
func (mr *MockClientServerMockRecorder) SetConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConnection", reflect.TypeOf((*MockClientServer)(nil).SetConnection), arg0)
}

// SetReadDeadline mocks base method
func (m *MockClientServer) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline
func (mr *MockClientServerMockRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockClientServer)(nil).SetReadDeadline), arg0)
}

// WriteMessage mocks base method
func (m *MockClientServer) WriteMessage(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage
func (mr *MockClientServerMockRecorder) WriteMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockClientServer)(nil).WriteMessage), arg0)
}
