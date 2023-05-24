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

// Code generated by mockery v2.28.0. DO NOT EDIT.

package interceptors

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcReadRequestResult is an autogenerated mock type for the PlcReadRequestResult type
type MockPlcReadRequestResult struct {
	mock.Mock
}

type MockPlcReadRequestResult_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcReadRequestResult) EXPECT() *MockPlcReadRequestResult_Expecter {
	return &MockPlcReadRequestResult_Expecter{mock: &_m.Mock}
}

// GetErr provides a mock function with given fields:
func (_m *MockPlcReadRequestResult) GetErr() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcReadRequestResult_GetErr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetErr'
type MockPlcReadRequestResult_GetErr_Call struct {
	*mock.Call
}

// GetErr is a helper method to define mock.On call
func (_e *MockPlcReadRequestResult_Expecter) GetErr() *MockPlcReadRequestResult_GetErr_Call {
	return &MockPlcReadRequestResult_GetErr_Call{Call: _e.mock.On("GetErr")}
}

func (_c *MockPlcReadRequestResult_GetErr_Call) Run(run func()) *MockPlcReadRequestResult_GetErr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadRequestResult_GetErr_Call) Return(_a0 error) *MockPlcReadRequestResult_GetErr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequestResult_GetErr_Call) RunAndReturn(run func() error) *MockPlcReadRequestResult_GetErr_Call {
	_c.Call.Return(run)
	return _c
}

// GetRequest provides a mock function with given fields:
func (_m *MockPlcReadRequestResult) GetRequest() model.PlcReadRequest {
	ret := _m.Called()

	var r0 model.PlcReadRequest
	if rf, ok := ret.Get(0).(func() model.PlcReadRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcReadRequest)
		}
	}

	return r0
}

// MockPlcReadRequestResult_GetRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequest'
type MockPlcReadRequestResult_GetRequest_Call struct {
	*mock.Call
}

// GetRequest is a helper method to define mock.On call
func (_e *MockPlcReadRequestResult_Expecter) GetRequest() *MockPlcReadRequestResult_GetRequest_Call {
	return &MockPlcReadRequestResult_GetRequest_Call{Call: _e.mock.On("GetRequest")}
}

func (_c *MockPlcReadRequestResult_GetRequest_Call) Run(run func()) *MockPlcReadRequestResult_GetRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadRequestResult_GetRequest_Call) Return(_a0 model.PlcReadRequest) *MockPlcReadRequestResult_GetRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequestResult_GetRequest_Call) RunAndReturn(run func() model.PlcReadRequest) *MockPlcReadRequestResult_GetRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetResponse provides a mock function with given fields:
func (_m *MockPlcReadRequestResult) GetResponse() model.PlcReadResponse {
	ret := _m.Called()

	var r0 model.PlcReadResponse
	if rf, ok := ret.Get(0).(func() model.PlcReadResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcReadResponse)
		}
	}

	return r0
}

// MockPlcReadRequestResult_GetResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResponse'
type MockPlcReadRequestResult_GetResponse_Call struct {
	*mock.Call
}

// GetResponse is a helper method to define mock.On call
func (_e *MockPlcReadRequestResult_Expecter) GetResponse() *MockPlcReadRequestResult_GetResponse_Call {
	return &MockPlcReadRequestResult_GetResponse_Call{Call: _e.mock.On("GetResponse")}
}

func (_c *MockPlcReadRequestResult_GetResponse_Call) Run(run func()) *MockPlcReadRequestResult_GetResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadRequestResult_GetResponse_Call) Return(_a0 model.PlcReadResponse) *MockPlcReadRequestResult_GetResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequestResult_GetResponse_Call) RunAndReturn(run func() model.PlcReadResponse) *MockPlcReadRequestResult_GetResponse_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcReadRequestResult interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcReadRequestResult creates a new instance of MockPlcReadRequestResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcReadRequestResult(t mockConstructorTestingTNewMockPlcReadRequestResult) *MockPlcReadRequestResult {
	mock := &MockPlcReadRequestResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
