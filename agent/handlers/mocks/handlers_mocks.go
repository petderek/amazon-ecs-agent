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
// Source: github.com/aws/amazon-ecs-agent/agent/handlers/utils (interfaces: DockerStateResolver)

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	dockerstate "github.com/aws/amazon-ecs-agent/agent/engine/dockerstate"
	gomock "github.com/golang/mock/gomock"
)

// MockDockerStateResolver is a mock of DockerStateResolver interface
type MockDockerStateResolver struct {
	ctrl     *gomock.Controller
	recorder *MockDockerStateResolverMockRecorder
}

// MockDockerStateResolverMockRecorder is the mock recorder for MockDockerStateResolver
type MockDockerStateResolverMockRecorder struct {
	mock *MockDockerStateResolver
}

// NewMockDockerStateResolver creates a new mock instance
func NewMockDockerStateResolver(ctrl *gomock.Controller) *MockDockerStateResolver {
	mock := &MockDockerStateResolver{ctrl: ctrl}
	mock.recorder = &MockDockerStateResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDockerStateResolver) EXPECT() *MockDockerStateResolverMockRecorder {
	return m.recorder
}

// State mocks base method
func (m *MockDockerStateResolver) State() dockerstate.TaskEngineState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(dockerstate.TaskEngineState)
	return ret0
}

// State indicates an expected call of State
func (mr *MockDockerStateResolverMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockDockerStateResolver)(nil).State))
}
