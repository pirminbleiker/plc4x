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

// Code generated by mockery v2.27.1. DO NOT EDIT.

package interceptors

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	spi "github.com/apache/plc4x/plc4go/spi"
)

// MockPlcReadRequest is an autogenerated mock type for the PlcReadRequest type
type MockPlcReadRequest struct {
	mock.Mock
}

type MockPlcReadRequest_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcReadRequest) EXPECT() *MockPlcReadRequest_Expecter {
	return &MockPlcReadRequest_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields:
func (_m *MockPlcReadRequest) Execute() <-chan model.PlcReadRequestResult {
	ret := _m.Called()

	var r0 <-chan model.PlcReadRequestResult
	if rf, ok := ret.Get(0).(func() <-chan model.PlcReadRequestResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan model.PlcReadRequestResult)
		}
	}

	return r0
}

// MockPlcReadRequest_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockPlcReadRequest_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockPlcReadRequest_Expecter) Execute() *MockPlcReadRequest_Execute_Call {
	return &MockPlcReadRequest_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockPlcReadRequest_Execute_Call) Run(run func()) *MockPlcReadRequest_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadRequest_Execute_Call) Return(_a0 <-chan model.PlcReadRequestResult) *MockPlcReadRequest_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequest_Execute_Call) RunAndReturn(run func() <-chan model.PlcReadRequestResult) *MockPlcReadRequest_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteWithContext provides a mock function with given fields: ctx
func (_m *MockPlcReadRequest) ExecuteWithContext(ctx context.Context) <-chan model.PlcReadRequestResult {
	ret := _m.Called(ctx)

	var r0 <-chan model.PlcReadRequestResult
	if rf, ok := ret.Get(0).(func(context.Context) <-chan model.PlcReadRequestResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan model.PlcReadRequestResult)
		}
	}

	return r0
}

// MockPlcReadRequest_ExecuteWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteWithContext'
type MockPlcReadRequest_ExecuteWithContext_Call struct {
	*mock.Call
}

// ExecuteWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPlcReadRequest_Expecter) ExecuteWithContext(ctx interface{}) *MockPlcReadRequest_ExecuteWithContext_Call {
	return &MockPlcReadRequest_ExecuteWithContext_Call{Call: _e.mock.On("ExecuteWithContext", ctx)}
}

func (_c *MockPlcReadRequest_ExecuteWithContext_Call) Run(run func(ctx context.Context)) *MockPlcReadRequest_ExecuteWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPlcReadRequest_ExecuteWithContext_Call) Return(_a0 <-chan model.PlcReadRequestResult) *MockPlcReadRequest_ExecuteWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequest_ExecuteWithContext_Call) RunAndReturn(run func(context.Context) <-chan model.PlcReadRequestResult) *MockPlcReadRequest_ExecuteWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetReadRequestInterceptor provides a mock function with given fields:
func (_m *MockPlcReadRequest) GetReadRequestInterceptor() ReadRequestInterceptor {
	ret := _m.Called()

	var r0 ReadRequestInterceptor
	if rf, ok := ret.Get(0).(func() ReadRequestInterceptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ReadRequestInterceptor)
		}
	}

	return r0
}

// MockPlcReadRequest_GetReadRequestInterceptor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReadRequestInterceptor'
type MockPlcReadRequest_GetReadRequestInterceptor_Call struct {
	*mock.Call
}

// GetReadRequestInterceptor is a helper method to define mock.On call
func (_e *MockPlcReadRequest_Expecter) GetReadRequestInterceptor() *MockPlcReadRequest_GetReadRequestInterceptor_Call {
	return &MockPlcReadRequest_GetReadRequestInterceptor_Call{Call: _e.mock.On("GetReadRequestInterceptor")}
}

func (_c *MockPlcReadRequest_GetReadRequestInterceptor_Call) Run(run func()) *MockPlcReadRequest_GetReadRequestInterceptor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadRequest_GetReadRequestInterceptor_Call) Return(_a0 ReadRequestInterceptor) *MockPlcReadRequest_GetReadRequestInterceptor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequest_GetReadRequestInterceptor_Call) RunAndReturn(run func() ReadRequestInterceptor) *MockPlcReadRequest_GetReadRequestInterceptor_Call {
	_c.Call.Return(run)
	return _c
}

// GetReader provides a mock function with given fields:
func (_m *MockPlcReadRequest) GetReader() spi.PlcReader {
	ret := _m.Called()

	var r0 spi.PlcReader
	if rf, ok := ret.Get(0).(func() spi.PlcReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.PlcReader)
		}
	}

	return r0
}

// MockPlcReadRequest_GetReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReader'
type MockPlcReadRequest_GetReader_Call struct {
	*mock.Call
}

// GetReader is a helper method to define mock.On call
func (_e *MockPlcReadRequest_Expecter) GetReader() *MockPlcReadRequest_GetReader_Call {
	return &MockPlcReadRequest_GetReader_Call{Call: _e.mock.On("GetReader")}
}

func (_c *MockPlcReadRequest_GetReader_Call) Run(run func()) *MockPlcReadRequest_GetReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadRequest_GetReader_Call) Return(_a0 spi.PlcReader) *MockPlcReadRequest_GetReader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequest_GetReader_Call) RunAndReturn(run func() spi.PlcReader) *MockPlcReadRequest_GetReader_Call {
	_c.Call.Return(run)
	return _c
}

// GetTag provides a mock function with given fields: tagName
func (_m *MockPlcReadRequest) GetTag(tagName string) model.PlcTag {
	ret := _m.Called(tagName)

	var r0 model.PlcTag
	if rf, ok := ret.Get(0).(func(string) model.PlcTag); ok {
		r0 = rf(tagName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcTag)
		}
	}

	return r0
}

// MockPlcReadRequest_GetTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTag'
type MockPlcReadRequest_GetTag_Call struct {
	*mock.Call
}

// GetTag is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcReadRequest_Expecter) GetTag(tagName interface{}) *MockPlcReadRequest_GetTag_Call {
	return &MockPlcReadRequest_GetTag_Call{Call: _e.mock.On("GetTag", tagName)}
}

func (_c *MockPlcReadRequest_GetTag_Call) Run(run func(tagName string)) *MockPlcReadRequest_GetTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcReadRequest_GetTag_Call) Return(_a0 model.PlcTag) *MockPlcReadRequest_GetTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequest_GetTag_Call) RunAndReturn(run func(string) model.PlcTag) *MockPlcReadRequest_GetTag_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagNames provides a mock function with given fields:
func (_m *MockPlcReadRequest) GetTagNames() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockPlcReadRequest_GetTagNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagNames'
type MockPlcReadRequest_GetTagNames_Call struct {
	*mock.Call
}

// GetTagNames is a helper method to define mock.On call
func (_e *MockPlcReadRequest_Expecter) GetTagNames() *MockPlcReadRequest_GetTagNames_Call {
	return &MockPlcReadRequest_GetTagNames_Call{Call: _e.mock.On("GetTagNames")}
}

func (_c *MockPlcReadRequest_GetTagNames_Call) Run(run func()) *MockPlcReadRequest_GetTagNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadRequest_GetTagNames_Call) Return(_a0 []string) *MockPlcReadRequest_GetTagNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequest_GetTagNames_Call) RunAndReturn(run func() []string) *MockPlcReadRequest_GetTagNames_Call {
	_c.Call.Return(run)
	return _c
}

// IsAPlcMessage provides a mock function with given fields:
func (_m *MockPlcReadRequest) IsAPlcMessage() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcReadRequest_IsAPlcMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAPlcMessage'
type MockPlcReadRequest_IsAPlcMessage_Call struct {
	*mock.Call
}

// IsAPlcMessage is a helper method to define mock.On call
func (_e *MockPlcReadRequest_Expecter) IsAPlcMessage() *MockPlcReadRequest_IsAPlcMessage_Call {
	return &MockPlcReadRequest_IsAPlcMessage_Call{Call: _e.mock.On("IsAPlcMessage")}
}

func (_c *MockPlcReadRequest_IsAPlcMessage_Call) Run(run func()) *MockPlcReadRequest_IsAPlcMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadRequest_IsAPlcMessage_Call) Return(_a0 bool) *MockPlcReadRequest_IsAPlcMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequest_IsAPlcMessage_Call) RunAndReturn(run func() bool) *MockPlcReadRequest_IsAPlcMessage_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcReadRequest) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcReadRequest_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcReadRequest_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcReadRequest_Expecter) String() *MockPlcReadRequest_String_Call {
	return &MockPlcReadRequest_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcReadRequest_String_Call) Run(run func()) *MockPlcReadRequest_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadRequest_String_Call) Return(_a0 string) *MockPlcReadRequest_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequest_String_Call) RunAndReturn(run func() string) *MockPlcReadRequest_String_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcReadRequest interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcReadRequest creates a new instance of MockPlcReadRequest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcReadRequest(t mockConstructorTestingTNewMockPlcReadRequest) *MockPlcReadRequest {
	mock := &MockPlcReadRequest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
