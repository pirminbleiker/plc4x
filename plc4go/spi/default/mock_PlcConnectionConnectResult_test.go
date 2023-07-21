/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.32.0. DO NOT EDIT.

package _default

import (
	plc4go "github.com/apache/plc4x/plc4go/pkg/api"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcConnectionConnectResult is an autogenerated mock type for the PlcConnectionConnectResult type
type MockPlcConnectionConnectResult struct {
	mock.Mock
}

type MockPlcConnectionConnectResult_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcConnectionConnectResult) EXPECT() *MockPlcConnectionConnectResult_Expecter {
	return &MockPlcConnectionConnectResult_Expecter{mock: &_m.Mock}
}

// GetConnection provides a mock function with given fields:
func (_m *MockPlcConnectionConnectResult) GetConnection() plc4go.PlcConnection {
	ret := _m.Called()

	var r0 plc4go.PlcConnection
	if rf, ok := ret.Get(0).(func() plc4go.PlcConnection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(plc4go.PlcConnection)
		}
	}

	return r0
}

// MockPlcConnectionConnectResult_GetConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnection'
type MockPlcConnectionConnectResult_GetConnection_Call struct {
	*mock.Call
}

// GetConnection is a helper method to define mock.On call
func (_e *MockPlcConnectionConnectResult_Expecter) GetConnection() *MockPlcConnectionConnectResult_GetConnection_Call {
	return &MockPlcConnectionConnectResult_GetConnection_Call{Call: _e.mock.On("GetConnection")}
}

func (_c *MockPlcConnectionConnectResult_GetConnection_Call) Run(run func()) *MockPlcConnectionConnectResult_GetConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionConnectResult_GetConnection_Call) Return(_a0 plc4go.PlcConnection) *MockPlcConnectionConnectResult_GetConnection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionConnectResult_GetConnection_Call) RunAndReturn(run func() plc4go.PlcConnection) *MockPlcConnectionConnectResult_GetConnection_Call {
	_c.Call.Return(run)
	return _c
}

// GetErr provides a mock function with given fields:
func (_m *MockPlcConnectionConnectResult) GetErr() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcConnectionConnectResult_GetErr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetErr'
type MockPlcConnectionConnectResult_GetErr_Call struct {
	*mock.Call
}

// GetErr is a helper method to define mock.On call
func (_e *MockPlcConnectionConnectResult_Expecter) GetErr() *MockPlcConnectionConnectResult_GetErr_Call {
	return &MockPlcConnectionConnectResult_GetErr_Call{Call: _e.mock.On("GetErr")}
}

func (_c *MockPlcConnectionConnectResult_GetErr_Call) Run(run func()) *MockPlcConnectionConnectResult_GetErr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionConnectResult_GetErr_Call) Return(_a0 error) *MockPlcConnectionConnectResult_GetErr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionConnectResult_GetErr_Call) RunAndReturn(run func() error) *MockPlcConnectionConnectResult_GetErr_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcConnectionConnectResult) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcConnectionConnectResult_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcConnectionConnectResult_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcConnectionConnectResult_Expecter) String() *MockPlcConnectionConnectResult_String_Call {
	return &MockPlcConnectionConnectResult_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcConnectionConnectResult_String_Call) Run(run func()) *MockPlcConnectionConnectResult_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionConnectResult_String_Call) Return(_a0 string) *MockPlcConnectionConnectResult_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionConnectResult_String_Call) RunAndReturn(run func() string) *MockPlcConnectionConnectResult_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcConnectionConnectResult creates a new instance of MockPlcConnectionConnectResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcConnectionConnectResult(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcConnectionConnectResult {
	mock := &MockPlcConnectionConnectResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
