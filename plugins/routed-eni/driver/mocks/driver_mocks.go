// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-cni-k8s/plugins/routed-eni/driver (interfaces: NetworkAPIs)

// Package mock_driver is a generated GoMock package.
package mock_driver

import (
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNetworkAPIs is a mock of NetworkAPIs interface
type MockNetworkAPIs struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkAPIsMockRecorder
}

// MockNetworkAPIsMockRecorder is the mock recorder for MockNetworkAPIs
type MockNetworkAPIsMockRecorder struct {
	mock *MockNetworkAPIs
}

// NewMockNetworkAPIs creates a new mock instance
func NewMockNetworkAPIs(ctrl *gomock.Controller) *MockNetworkAPIs {
	mock := &MockNetworkAPIs{ctrl: ctrl}
	mock.recorder = &MockNetworkAPIsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkAPIs) EXPECT() *MockNetworkAPIsMockRecorder {
	return m.recorder
}

// SetupNS mocks base method
func (m *MockNetworkAPIs) SetupNS(arg0, arg1, arg2 string, arg3 *net.IPNet, arg4 int, arg5 []string, arg6 bool) error {
	ret := m.ctrl.Call(m, "SetupNS", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupNS indicates an expected call of SetupNS
func (mr *MockNetworkAPIsMockRecorder) SetupNS(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupNS", reflect.TypeOf((*MockNetworkAPIs)(nil).SetupNS), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// TeardownNS mocks base method
func (m *MockNetworkAPIs) TeardownNS(arg0 *net.IPNet, arg1 int) error {
	ret := m.ctrl.Call(m, "TeardownNS", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TeardownNS indicates an expected call of TeardownNS
func (mr *MockNetworkAPIsMockRecorder) TeardownNS(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeardownNS", reflect.TypeOf((*MockNetworkAPIs)(nil).TeardownNS), arg0, arg1)
}
