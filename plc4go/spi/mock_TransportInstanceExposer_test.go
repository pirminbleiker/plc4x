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

package spi

import (
	transports "github.com/apache/plc4x/plc4go/spi/transports"
	mock "github.com/stretchr/testify/mock"
)

// MockTransportInstanceExposer is an autogenerated mock type for the TransportInstanceExposer type
type MockTransportInstanceExposer struct {
	mock.Mock
}

type MockTransportInstanceExposer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransportInstanceExposer) EXPECT() *MockTransportInstanceExposer_Expecter {
	return &MockTransportInstanceExposer_Expecter{mock: &_m.Mock}
}

// GetTransportInstance provides a mock function with given fields:
func (_m *MockTransportInstanceExposer) GetTransportInstance() transports.TransportInstance {
	ret := _m.Called()

	var r0 transports.TransportInstance
	if rf, ok := ret.Get(0).(func() transports.TransportInstance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transports.TransportInstance)
		}
	}

	return r0
}

// MockTransportInstanceExposer_GetTransportInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransportInstance'
type MockTransportInstanceExposer_GetTransportInstance_Call struct {
	*mock.Call
}

// GetTransportInstance is a helper method to define mock.On call
func (_e *MockTransportInstanceExposer_Expecter) GetTransportInstance() *MockTransportInstanceExposer_GetTransportInstance_Call {
	return &MockTransportInstanceExposer_GetTransportInstance_Call{Call: _e.mock.On("GetTransportInstance")}
}

func (_c *MockTransportInstanceExposer_GetTransportInstance_Call) Run(run func()) *MockTransportInstanceExposer_GetTransportInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransportInstanceExposer_GetTransportInstance_Call) Return(_a0 transports.TransportInstance) *MockTransportInstanceExposer_GetTransportInstance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransportInstanceExposer_GetTransportInstance_Call) RunAndReturn(run func() transports.TransportInstance) *MockTransportInstanceExposer_GetTransportInstance_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockTransportInstanceExposer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTransportInstanceExposer creates a new instance of MockTransportInstanceExposer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTransportInstanceExposer(t mockConstructorTestingTNewMockTransportInstanceExposer) *MockTransportInstanceExposer {
	mock := &MockTransportInstanceExposer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
