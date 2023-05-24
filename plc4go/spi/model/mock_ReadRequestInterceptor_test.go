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

package model

import (
	context "context"

	apimodel "github.com/apache/plc4x/plc4go/pkg/api/model"

	mock "github.com/stretchr/testify/mock"
)

// MockReadRequestInterceptor is an autogenerated mock type for the ReadRequestInterceptor type
type MockReadRequestInterceptor struct {
	mock.Mock
}

type MockReadRequestInterceptor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReadRequestInterceptor) EXPECT() *MockReadRequestInterceptor_Expecter {
	return &MockReadRequestInterceptor_Expecter{mock: &_m.Mock}
}

// InterceptReadRequest provides a mock function with given fields: ctx, readRequest
func (_m *MockReadRequestInterceptor) InterceptReadRequest(ctx context.Context, readRequest apimodel.PlcReadRequest) []apimodel.PlcReadRequest {
	ret := _m.Called(ctx, readRequest)

	var r0 []apimodel.PlcReadRequest
	if rf, ok := ret.Get(0).(func(context.Context, apimodel.PlcReadRequest) []apimodel.PlcReadRequest); ok {
		r0 = rf(ctx, readRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]apimodel.PlcReadRequest)
		}
	}

	return r0
}

// MockReadRequestInterceptor_InterceptReadRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InterceptReadRequest'
type MockReadRequestInterceptor_InterceptReadRequest_Call struct {
	*mock.Call
}

// InterceptReadRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - readRequest apimodel.PlcReadRequest
func (_e *MockReadRequestInterceptor_Expecter) InterceptReadRequest(ctx interface{}, readRequest interface{}) *MockReadRequestInterceptor_InterceptReadRequest_Call {
	return &MockReadRequestInterceptor_InterceptReadRequest_Call{Call: _e.mock.On("InterceptReadRequest", ctx, readRequest)}
}

func (_c *MockReadRequestInterceptor_InterceptReadRequest_Call) Run(run func(ctx context.Context, readRequest apimodel.PlcReadRequest)) *MockReadRequestInterceptor_InterceptReadRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(apimodel.PlcReadRequest))
	})
	return _c
}

func (_c *MockReadRequestInterceptor_InterceptReadRequest_Call) Return(_a0 []apimodel.PlcReadRequest) *MockReadRequestInterceptor_InterceptReadRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReadRequestInterceptor_InterceptReadRequest_Call) RunAndReturn(run func(context.Context, apimodel.PlcReadRequest) []apimodel.PlcReadRequest) *MockReadRequestInterceptor_InterceptReadRequest_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessReadResponses provides a mock function with given fields: ctx, readRequest, readResults
func (_m *MockReadRequestInterceptor) ProcessReadResponses(ctx context.Context, readRequest apimodel.PlcReadRequest, readResults []apimodel.PlcReadRequestResult) apimodel.PlcReadRequestResult {
	ret := _m.Called(ctx, readRequest, readResults)

	var r0 apimodel.PlcReadRequestResult
	if rf, ok := ret.Get(0).(func(context.Context, apimodel.PlcReadRequest, []apimodel.PlcReadRequestResult) apimodel.PlcReadRequestResult); ok {
		r0 = rf(ctx, readRequest, readResults)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apimodel.PlcReadRequestResult)
		}
	}

	return r0
}

// MockReadRequestInterceptor_ProcessReadResponses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessReadResponses'
type MockReadRequestInterceptor_ProcessReadResponses_Call struct {
	*mock.Call
}

// ProcessReadResponses is a helper method to define mock.On call
//   - ctx context.Context
//   - readRequest apimodel.PlcReadRequest
//   - readResults []apimodel.PlcReadRequestResult
func (_e *MockReadRequestInterceptor_Expecter) ProcessReadResponses(ctx interface{}, readRequest interface{}, readResults interface{}) *MockReadRequestInterceptor_ProcessReadResponses_Call {
	return &MockReadRequestInterceptor_ProcessReadResponses_Call{Call: _e.mock.On("ProcessReadResponses", ctx, readRequest, readResults)}
}

func (_c *MockReadRequestInterceptor_ProcessReadResponses_Call) Run(run func(ctx context.Context, readRequest apimodel.PlcReadRequest, readResults []apimodel.PlcReadRequestResult)) *MockReadRequestInterceptor_ProcessReadResponses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(apimodel.PlcReadRequest), args[2].([]apimodel.PlcReadRequestResult))
	})
	return _c
}

func (_c *MockReadRequestInterceptor_ProcessReadResponses_Call) Return(_a0 apimodel.PlcReadRequestResult) *MockReadRequestInterceptor_ProcessReadResponses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReadRequestInterceptor_ProcessReadResponses_Call) RunAndReturn(run func(context.Context, apimodel.PlcReadRequest, []apimodel.PlcReadRequestResult) apimodel.PlcReadRequestResult) *MockReadRequestInterceptor_ProcessReadResponses_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockReadRequestInterceptor interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockReadRequestInterceptor creates a new instance of MockReadRequestInterceptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockReadRequestInterceptor(t mockConstructorTestingTNewMockReadRequestInterceptor) *MockReadRequestInterceptor {
	mock := &MockReadRequestInterceptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
