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

// Code generated by mockery v2.32.4. DO NOT EDIT.

package pool

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockExecutor is an autogenerated mock type for the Executor type
type MockExecutor struct {
	mock.Mock
}

type MockExecutor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExecutor) EXPECT() *MockExecutor_Expecter {
	return &MockExecutor_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockExecutor) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExecutor_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockExecutor_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockExecutor_Expecter) Close() *MockExecutor_Close_Call {
	return &MockExecutor_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockExecutor_Close_Call) Run(run func()) *MockExecutor_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExecutor_Close_Call) Return(_a0 error) *MockExecutor_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExecutor_Close_Call) RunAndReturn(run func() error) *MockExecutor_Close_Call {
	_c.Call.Return(run)
	return _c
}

// IsRunning provides a mock function with given fields:
func (_m *MockExecutor) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockExecutor_IsRunning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRunning'
type MockExecutor_IsRunning_Call struct {
	*mock.Call
}

// IsRunning is a helper method to define mock.On call
func (_e *MockExecutor_Expecter) IsRunning() *MockExecutor_IsRunning_Call {
	return &MockExecutor_IsRunning_Call{Call: _e.mock.On("IsRunning")}
}

func (_c *MockExecutor_IsRunning_Call) Run(run func()) *MockExecutor_IsRunning_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExecutor_IsRunning_Call) Return(_a0 bool) *MockExecutor_IsRunning_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExecutor_IsRunning_Call) RunAndReturn(run func() bool) *MockExecutor_IsRunning_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockExecutor) Start() {
	_m.Called()
}

// MockExecutor_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockExecutor_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockExecutor_Expecter) Start() *MockExecutor_Start_Call {
	return &MockExecutor_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockExecutor_Start_Call) Run(run func()) *MockExecutor_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExecutor_Start_Call) Return() *MockExecutor_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockExecutor_Start_Call) RunAndReturn(run func()) *MockExecutor_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockExecutor) Stop() {
	_m.Called()
}

// MockExecutor_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockExecutor_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockExecutor_Expecter) Stop() *MockExecutor_Stop_Call {
	return &MockExecutor_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockExecutor_Stop_Call) Run(run func()) *MockExecutor_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExecutor_Stop_Call) Return() *MockExecutor_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockExecutor_Stop_Call) RunAndReturn(run func()) *MockExecutor_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// Submit provides a mock function with given fields: ctx, workItemId, runnable
func (_m *MockExecutor) Submit(ctx context.Context, workItemId int32, runnable Runnable) CompletionFuture {
	ret := _m.Called(ctx, workItemId, runnable)

	var r0 CompletionFuture
	if rf, ok := ret.Get(0).(func(context.Context, int32, Runnable) CompletionFuture); ok {
		r0 = rf(ctx, workItemId, runnable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(CompletionFuture)
		}
	}

	return r0
}

// MockExecutor_Submit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Submit'
type MockExecutor_Submit_Call struct {
	*mock.Call
}

// Submit is a helper method to define mock.On call
//   - ctx context.Context
//   - workItemId int32
//   - runnable Runnable
func (_e *MockExecutor_Expecter) Submit(ctx interface{}, workItemId interface{}, runnable interface{}) *MockExecutor_Submit_Call {
	return &MockExecutor_Submit_Call{Call: _e.mock.On("Submit", ctx, workItemId, runnable)}
}

func (_c *MockExecutor_Submit_Call) Run(run func(ctx context.Context, workItemId int32, runnable Runnable)) *MockExecutor_Submit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(Runnable))
	})
	return _c
}

func (_c *MockExecutor_Submit_Call) Return(_a0 CompletionFuture) *MockExecutor_Submit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExecutor_Submit_Call) RunAndReturn(run func(context.Context, int32, Runnable) CompletionFuture) *MockExecutor_Submit_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExecutor creates a new instance of MockExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExecutor {
	mock := &MockExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
