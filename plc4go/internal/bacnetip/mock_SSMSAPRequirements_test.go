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

package bacnetip

import (
	model "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	mock "github.com/stretchr/testify/mock"
)

// MockSSMSAPRequirements is an autogenerated mock type for the SSMSAPRequirements type
type MockSSMSAPRequirements struct {
	mock.Mock
}

type MockSSMSAPRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSSMSAPRequirements) EXPECT() *MockSSMSAPRequirements_Expecter {
	return &MockSSMSAPRequirements_Expecter{mock: &_m.Mock}
}

// Confirmation provides a mock function with given fields: pdu
func (_m *MockSSMSAPRequirements) Confirmation(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSSMSAPRequirements_Confirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Confirmation'
type MockSSMSAPRequirements_Confirmation_Call struct {
	*mock.Call
}

// Confirmation is a helper method to define mock.On call
//   - pdu _PDU
func (_e *MockSSMSAPRequirements_Expecter) Confirmation(pdu interface{}) *MockSSMSAPRequirements_Confirmation_Call {
	return &MockSSMSAPRequirements_Confirmation_Call{Call: _e.mock.On("Confirmation", pdu)}
}

func (_c *MockSSMSAPRequirements_Confirmation_Call) Run(run func(pdu _PDU)) *MockSSMSAPRequirements_Confirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *MockSSMSAPRequirements_Confirmation_Call) Return(_a0 error) *MockSSMSAPRequirements_Confirmation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_Confirmation_Call) RunAndReturn(run func(_PDU) error) *MockSSMSAPRequirements_Confirmation_Call {
	_c.Call.Return(run)
	return _c
}

// GetApplicationTimeout provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetApplicationTimeout() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// MockSSMSAPRequirements_GetApplicationTimeout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetApplicationTimeout'
type MockSSMSAPRequirements_GetApplicationTimeout_Call struct {
	*mock.Call
}

// GetApplicationTimeout is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetApplicationTimeout() *MockSSMSAPRequirements_GetApplicationTimeout_Call {
	return &MockSSMSAPRequirements_GetApplicationTimeout_Call{Call: _e.mock.On("GetApplicationTimeout")}
}

func (_c *MockSSMSAPRequirements_GetApplicationTimeout_Call) Run(run func()) *MockSSMSAPRequirements_GetApplicationTimeout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetApplicationTimeout_Call) Return(_a0 uint) *MockSSMSAPRequirements_GetApplicationTimeout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetApplicationTimeout_Call) RunAndReturn(run func() uint) *MockSSMSAPRequirements_GetApplicationTimeout_Call {
	_c.Call.Return(run)
	return _c
}

// GetClientTransactions provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetClientTransactions() []*ClientSSM {
	ret := _m.Called()

	var r0 []*ClientSSM
	if rf, ok := ret.Get(0).(func() []*ClientSSM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ClientSSM)
		}
	}

	return r0
}

// MockSSMSAPRequirements_GetClientTransactions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClientTransactions'
type MockSSMSAPRequirements_GetClientTransactions_Call struct {
	*mock.Call
}

// GetClientTransactions is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetClientTransactions() *MockSSMSAPRequirements_GetClientTransactions_Call {
	return &MockSSMSAPRequirements_GetClientTransactions_Call{Call: _e.mock.On("GetClientTransactions")}
}

func (_c *MockSSMSAPRequirements_GetClientTransactions_Call) Run(run func()) *MockSSMSAPRequirements_GetClientTransactions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetClientTransactions_Call) Return(_a0 []*ClientSSM) *MockSSMSAPRequirements_GetClientTransactions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetClientTransactions_Call) RunAndReturn(run func() []*ClientSSM) *MockSSMSAPRequirements_GetClientTransactions_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultAPDUSegmentTimeout provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetDefaultAPDUSegmentTimeout() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// MockSSMSAPRequirements_GetDefaultAPDUSegmentTimeout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultAPDUSegmentTimeout'
type MockSSMSAPRequirements_GetDefaultAPDUSegmentTimeout_Call struct {
	*mock.Call
}

// GetDefaultAPDUSegmentTimeout is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetDefaultAPDUSegmentTimeout() *MockSSMSAPRequirements_GetDefaultAPDUSegmentTimeout_Call {
	return &MockSSMSAPRequirements_GetDefaultAPDUSegmentTimeout_Call{Call: _e.mock.On("GetDefaultAPDUSegmentTimeout")}
}

func (_c *MockSSMSAPRequirements_GetDefaultAPDUSegmentTimeout_Call) Run(run func()) *MockSSMSAPRequirements_GetDefaultAPDUSegmentTimeout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetDefaultAPDUSegmentTimeout_Call) Return(_a0 uint) *MockSSMSAPRequirements_GetDefaultAPDUSegmentTimeout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetDefaultAPDUSegmentTimeout_Call) RunAndReturn(run func() uint) *MockSSMSAPRequirements_GetDefaultAPDUSegmentTimeout_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultAPDUTimeout provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetDefaultAPDUTimeout() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// MockSSMSAPRequirements_GetDefaultAPDUTimeout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultAPDUTimeout'
type MockSSMSAPRequirements_GetDefaultAPDUTimeout_Call struct {
	*mock.Call
}

// GetDefaultAPDUTimeout is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetDefaultAPDUTimeout() *MockSSMSAPRequirements_GetDefaultAPDUTimeout_Call {
	return &MockSSMSAPRequirements_GetDefaultAPDUTimeout_Call{Call: _e.mock.On("GetDefaultAPDUTimeout")}
}

func (_c *MockSSMSAPRequirements_GetDefaultAPDUTimeout_Call) Run(run func()) *MockSSMSAPRequirements_GetDefaultAPDUTimeout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetDefaultAPDUTimeout_Call) Return(_a0 uint) *MockSSMSAPRequirements_GetDefaultAPDUTimeout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetDefaultAPDUTimeout_Call) RunAndReturn(run func() uint) *MockSSMSAPRequirements_GetDefaultAPDUTimeout_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultMaxSegmentsAccepted provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetDefaultMaxSegmentsAccepted() model.MaxSegmentsAccepted {
	ret := _m.Called()

	var r0 model.MaxSegmentsAccepted
	if rf, ok := ret.Get(0).(func() model.MaxSegmentsAccepted); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.MaxSegmentsAccepted)
	}

	return r0
}

// MockSSMSAPRequirements_GetDefaultMaxSegmentsAccepted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultMaxSegmentsAccepted'
type MockSSMSAPRequirements_GetDefaultMaxSegmentsAccepted_Call struct {
	*mock.Call
}

// GetDefaultMaxSegmentsAccepted is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetDefaultMaxSegmentsAccepted() *MockSSMSAPRequirements_GetDefaultMaxSegmentsAccepted_Call {
	return &MockSSMSAPRequirements_GetDefaultMaxSegmentsAccepted_Call{Call: _e.mock.On("GetDefaultMaxSegmentsAccepted")}
}

func (_c *MockSSMSAPRequirements_GetDefaultMaxSegmentsAccepted_Call) Run(run func()) *MockSSMSAPRequirements_GetDefaultMaxSegmentsAccepted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetDefaultMaxSegmentsAccepted_Call) Return(_a0 model.MaxSegmentsAccepted) *MockSSMSAPRequirements_GetDefaultMaxSegmentsAccepted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetDefaultMaxSegmentsAccepted_Call) RunAndReturn(run func() model.MaxSegmentsAccepted) *MockSSMSAPRequirements_GetDefaultMaxSegmentsAccepted_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultMaximumApduLengthAccepted provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetDefaultMaximumApduLengthAccepted() model.MaxApduLengthAccepted {
	ret := _m.Called()

	var r0 model.MaxApduLengthAccepted
	if rf, ok := ret.Get(0).(func() model.MaxApduLengthAccepted); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.MaxApduLengthAccepted)
	}

	return r0
}

// MockSSMSAPRequirements_GetDefaultMaximumApduLengthAccepted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultMaximumApduLengthAccepted'
type MockSSMSAPRequirements_GetDefaultMaximumApduLengthAccepted_Call struct {
	*mock.Call
}

// GetDefaultMaximumApduLengthAccepted is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetDefaultMaximumApduLengthAccepted() *MockSSMSAPRequirements_GetDefaultMaximumApduLengthAccepted_Call {
	return &MockSSMSAPRequirements_GetDefaultMaximumApduLengthAccepted_Call{Call: _e.mock.On("GetDefaultMaximumApduLengthAccepted")}
}

func (_c *MockSSMSAPRequirements_GetDefaultMaximumApduLengthAccepted_Call) Run(run func()) *MockSSMSAPRequirements_GetDefaultMaximumApduLengthAccepted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetDefaultMaximumApduLengthAccepted_Call) Return(_a0 model.MaxApduLengthAccepted) *MockSSMSAPRequirements_GetDefaultMaximumApduLengthAccepted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetDefaultMaximumApduLengthAccepted_Call) RunAndReturn(run func() model.MaxApduLengthAccepted) *MockSSMSAPRequirements_GetDefaultMaximumApduLengthAccepted_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultSegmentationSupported provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetDefaultSegmentationSupported() model.BACnetSegmentation {
	ret := _m.Called()

	var r0 model.BACnetSegmentation
	if rf, ok := ret.Get(0).(func() model.BACnetSegmentation); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.BACnetSegmentation)
	}

	return r0
}

// MockSSMSAPRequirements_GetDefaultSegmentationSupported_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultSegmentationSupported'
type MockSSMSAPRequirements_GetDefaultSegmentationSupported_Call struct {
	*mock.Call
}

// GetDefaultSegmentationSupported is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetDefaultSegmentationSupported() *MockSSMSAPRequirements_GetDefaultSegmentationSupported_Call {
	return &MockSSMSAPRequirements_GetDefaultSegmentationSupported_Call{Call: _e.mock.On("GetDefaultSegmentationSupported")}
}

func (_c *MockSSMSAPRequirements_GetDefaultSegmentationSupported_Call) Run(run func()) *MockSSMSAPRequirements_GetDefaultSegmentationSupported_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetDefaultSegmentationSupported_Call) Return(_a0 model.BACnetSegmentation) *MockSSMSAPRequirements_GetDefaultSegmentationSupported_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetDefaultSegmentationSupported_Call) RunAndReturn(run func() model.BACnetSegmentation) *MockSSMSAPRequirements_GetDefaultSegmentationSupported_Call {
	_c.Call.Return(run)
	return _c
}

// GetDeviceInfoCache provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetDeviceInfoCache() *DeviceInfoCache {
	ret := _m.Called()

	var r0 *DeviceInfoCache
	if rf, ok := ret.Get(0).(func() *DeviceInfoCache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeviceInfoCache)
		}
	}

	return r0
}

// MockSSMSAPRequirements_GetDeviceInfoCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDeviceInfoCache'
type MockSSMSAPRequirements_GetDeviceInfoCache_Call struct {
	*mock.Call
}

// GetDeviceInfoCache is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetDeviceInfoCache() *MockSSMSAPRequirements_GetDeviceInfoCache_Call {
	return &MockSSMSAPRequirements_GetDeviceInfoCache_Call{Call: _e.mock.On("GetDeviceInfoCache")}
}

func (_c *MockSSMSAPRequirements_GetDeviceInfoCache_Call) Run(run func()) *MockSSMSAPRequirements_GetDeviceInfoCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetDeviceInfoCache_Call) Return(_a0 *DeviceInfoCache) *MockSSMSAPRequirements_GetDeviceInfoCache_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetDeviceInfoCache_Call) RunAndReturn(run func() *DeviceInfoCache) *MockSSMSAPRequirements_GetDeviceInfoCache_Call {
	_c.Call.Return(run)
	return _c
}

// GetLocalDevice provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetLocalDevice() *LocalDeviceObject {
	ret := _m.Called()

	var r0 *LocalDeviceObject
	if rf, ok := ret.Get(0).(func() *LocalDeviceObject); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LocalDeviceObject)
		}
	}

	return r0
}

// MockSSMSAPRequirements_GetLocalDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLocalDevice'
type MockSSMSAPRequirements_GetLocalDevice_Call struct {
	*mock.Call
}

// GetLocalDevice is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetLocalDevice() *MockSSMSAPRequirements_GetLocalDevice_Call {
	return &MockSSMSAPRequirements_GetLocalDevice_Call{Call: _e.mock.On("GetLocalDevice")}
}

func (_c *MockSSMSAPRequirements_GetLocalDevice_Call) Run(run func()) *MockSSMSAPRequirements_GetLocalDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetLocalDevice_Call) Return(_a0 *LocalDeviceObject) *MockSSMSAPRequirements_GetLocalDevice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetLocalDevice_Call) RunAndReturn(run func() *LocalDeviceObject) *MockSSMSAPRequirements_GetLocalDevice_Call {
	_c.Call.Return(run)
	return _c
}

// GetProposedWindowSize provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetProposedWindowSize() uint8 {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}

// MockSSMSAPRequirements_GetProposedWindowSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProposedWindowSize'
type MockSSMSAPRequirements_GetProposedWindowSize_Call struct {
	*mock.Call
}

// GetProposedWindowSize is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetProposedWindowSize() *MockSSMSAPRequirements_GetProposedWindowSize_Call {
	return &MockSSMSAPRequirements_GetProposedWindowSize_Call{Call: _e.mock.On("GetProposedWindowSize")}
}

func (_c *MockSSMSAPRequirements_GetProposedWindowSize_Call) Run(run func()) *MockSSMSAPRequirements_GetProposedWindowSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetProposedWindowSize_Call) Return(_a0 uint8) *MockSSMSAPRequirements_GetProposedWindowSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetProposedWindowSize_Call) RunAndReturn(run func() uint8) *MockSSMSAPRequirements_GetProposedWindowSize_Call {
	_c.Call.Return(run)
	return _c
}

// GetServerTransactions provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) GetServerTransactions() []*ServerSSM {
	ret := _m.Called()

	var r0 []*ServerSSM
	if rf, ok := ret.Get(0).(func() []*ServerSSM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ServerSSM)
		}
	}

	return r0
}

// MockSSMSAPRequirements_GetServerTransactions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetServerTransactions'
type MockSSMSAPRequirements_GetServerTransactions_Call struct {
	*mock.Call
}

// GetServerTransactions is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) GetServerTransactions() *MockSSMSAPRequirements_GetServerTransactions_Call {
	return &MockSSMSAPRequirements_GetServerTransactions_Call{Call: _e.mock.On("GetServerTransactions")}
}

func (_c *MockSSMSAPRequirements_GetServerTransactions_Call) Run(run func()) *MockSSMSAPRequirements_GetServerTransactions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_GetServerTransactions_Call) Return(_a0 []*ServerSSM) *MockSSMSAPRequirements_GetServerTransactions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_GetServerTransactions_Call) RunAndReturn(run func() []*ServerSSM) *MockSSMSAPRequirements_GetServerTransactions_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveClientTransaction provides a mock function with given fields: _a0
func (_m *MockSSMSAPRequirements) RemoveClientTransaction(_a0 *ClientSSM) {
	_m.Called(_a0)
}

// MockSSMSAPRequirements_RemoveClientTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveClientTransaction'
type MockSSMSAPRequirements_RemoveClientTransaction_Call struct {
	*mock.Call
}

// RemoveClientTransaction is a helper method to define mock.On call
//   - _a0 *ClientSSM
func (_e *MockSSMSAPRequirements_Expecter) RemoveClientTransaction(_a0 interface{}) *MockSSMSAPRequirements_RemoveClientTransaction_Call {
	return &MockSSMSAPRequirements_RemoveClientTransaction_Call{Call: _e.mock.On("RemoveClientTransaction", _a0)}
}

func (_c *MockSSMSAPRequirements_RemoveClientTransaction_Call) Run(run func(_a0 *ClientSSM)) *MockSSMSAPRequirements_RemoveClientTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*ClientSSM))
	})
	return _c
}

func (_c *MockSSMSAPRequirements_RemoveClientTransaction_Call) Return() *MockSSMSAPRequirements_RemoveClientTransaction_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSSMSAPRequirements_RemoveClientTransaction_Call) RunAndReturn(run func(*ClientSSM)) *MockSSMSAPRequirements_RemoveClientTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveServerTransaction provides a mock function with given fields: _a0
func (_m *MockSSMSAPRequirements) RemoveServerTransaction(_a0 *ServerSSM) {
	_m.Called(_a0)
}

// MockSSMSAPRequirements_RemoveServerTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveServerTransaction'
type MockSSMSAPRequirements_RemoveServerTransaction_Call struct {
	*mock.Call
}

// RemoveServerTransaction is a helper method to define mock.On call
//   - _a0 *ServerSSM
func (_e *MockSSMSAPRequirements_Expecter) RemoveServerTransaction(_a0 interface{}) *MockSSMSAPRequirements_RemoveServerTransaction_Call {
	return &MockSSMSAPRequirements_RemoveServerTransaction_Call{Call: _e.mock.On("RemoveServerTransaction", _a0)}
}

func (_c *MockSSMSAPRequirements_RemoveServerTransaction_Call) Run(run func(_a0 *ServerSSM)) *MockSSMSAPRequirements_RemoveServerTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*ServerSSM))
	})
	return _c
}

func (_c *MockSSMSAPRequirements_RemoveServerTransaction_Call) Return() *MockSSMSAPRequirements_RemoveServerTransaction_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSSMSAPRequirements_RemoveServerTransaction_Call) RunAndReturn(run func(*ServerSSM)) *MockSSMSAPRequirements_RemoveServerTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// Request provides a mock function with given fields: pdu
func (_m *MockSSMSAPRequirements) Request(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSSMSAPRequirements_Request_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Request'
type MockSSMSAPRequirements_Request_Call struct {
	*mock.Call
}

// Request is a helper method to define mock.On call
//   - pdu _PDU
func (_e *MockSSMSAPRequirements_Expecter) Request(pdu interface{}) *MockSSMSAPRequirements_Request_Call {
	return &MockSSMSAPRequirements_Request_Call{Call: _e.mock.On("Request", pdu)}
}

func (_c *MockSSMSAPRequirements_Request_Call) Run(run func(pdu _PDU)) *MockSSMSAPRequirements_Request_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *MockSSMSAPRequirements_Request_Call) Return(_a0 error) *MockSSMSAPRequirements_Request_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_Request_Call) RunAndReturn(run func(_PDU) error) *MockSSMSAPRequirements_Request_Call {
	_c.Call.Return(run)
	return _c
}

// SapConfirmation provides a mock function with given fields: pdu
func (_m *MockSSMSAPRequirements) SapConfirmation(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSSMSAPRequirements_SapConfirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapConfirmation'
type MockSSMSAPRequirements_SapConfirmation_Call struct {
	*mock.Call
}

// SapConfirmation is a helper method to define mock.On call
//   - pdu _PDU
func (_e *MockSSMSAPRequirements_Expecter) SapConfirmation(pdu interface{}) *MockSSMSAPRequirements_SapConfirmation_Call {
	return &MockSSMSAPRequirements_SapConfirmation_Call{Call: _e.mock.On("SapConfirmation", pdu)}
}

func (_c *MockSSMSAPRequirements_SapConfirmation_Call) Run(run func(pdu _PDU)) *MockSSMSAPRequirements_SapConfirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *MockSSMSAPRequirements_SapConfirmation_Call) Return(_a0 error) *MockSSMSAPRequirements_SapConfirmation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_SapConfirmation_Call) RunAndReturn(run func(_PDU) error) *MockSSMSAPRequirements_SapConfirmation_Call {
	_c.Call.Return(run)
	return _c
}

// SapIndication provides a mock function with given fields: pdu
func (_m *MockSSMSAPRequirements) SapIndication(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSSMSAPRequirements_SapIndication_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapIndication'
type MockSSMSAPRequirements_SapIndication_Call struct {
	*mock.Call
}

// SapIndication is a helper method to define mock.On call
//   - pdu _PDU
func (_e *MockSSMSAPRequirements_Expecter) SapIndication(pdu interface{}) *MockSSMSAPRequirements_SapIndication_Call {
	return &MockSSMSAPRequirements_SapIndication_Call{Call: _e.mock.On("SapIndication", pdu)}
}

func (_c *MockSSMSAPRequirements_SapIndication_Call) Run(run func(pdu _PDU)) *MockSSMSAPRequirements_SapIndication_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *MockSSMSAPRequirements_SapIndication_Call) Return(_a0 error) *MockSSMSAPRequirements_SapIndication_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_SapIndication_Call) RunAndReturn(run func(_PDU) error) *MockSSMSAPRequirements_SapIndication_Call {
	_c.Call.Return(run)
	return _c
}

// SapRequest provides a mock function with given fields: pdu
func (_m *MockSSMSAPRequirements) SapRequest(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSSMSAPRequirements_SapRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapRequest'
type MockSSMSAPRequirements_SapRequest_Call struct {
	*mock.Call
}

// SapRequest is a helper method to define mock.On call
//   - pdu _PDU
func (_e *MockSSMSAPRequirements_Expecter) SapRequest(pdu interface{}) *MockSSMSAPRequirements_SapRequest_Call {
	return &MockSSMSAPRequirements_SapRequest_Call{Call: _e.mock.On("SapRequest", pdu)}
}

func (_c *MockSSMSAPRequirements_SapRequest_Call) Run(run func(pdu _PDU)) *MockSSMSAPRequirements_SapRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *MockSSMSAPRequirements_SapRequest_Call) Return(_a0 error) *MockSSMSAPRequirements_SapRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_SapRequest_Call) RunAndReturn(run func(_PDU) error) *MockSSMSAPRequirements_SapRequest_Call {
	_c.Call.Return(run)
	return _c
}

// SapResponse provides a mock function with given fields: pdu
func (_m *MockSSMSAPRequirements) SapResponse(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSSMSAPRequirements_SapResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapResponse'
type MockSSMSAPRequirements_SapResponse_Call struct {
	*mock.Call
}

// SapResponse is a helper method to define mock.On call
//   - pdu _PDU
func (_e *MockSSMSAPRequirements_Expecter) SapResponse(pdu interface{}) *MockSSMSAPRequirements_SapResponse_Call {
	return &MockSSMSAPRequirements_SapResponse_Call{Call: _e.mock.On("SapResponse", pdu)}
}

func (_c *MockSSMSAPRequirements_SapResponse_Call) Run(run func(pdu _PDU)) *MockSSMSAPRequirements_SapResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *MockSSMSAPRequirements_SapResponse_Call) Return(_a0 error) *MockSSMSAPRequirements_SapResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_SapResponse_Call) RunAndReturn(run func(_PDU) error) *MockSSMSAPRequirements_SapResponse_Call {
	_c.Call.Return(run)
	return _c
}

// _setClientPeer provides a mock function with given fields: server
func (_m *MockSSMSAPRequirements) _setClientPeer(server _Server) {
	_m.Called(server)
}

// MockSSMSAPRequirements__setClientPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method '_setClientPeer'
type MockSSMSAPRequirements__setClientPeer_Call struct {
	*mock.Call
}

// _setClientPeer is a helper method to define mock.On call
//   - server _Server
func (_e *MockSSMSAPRequirements_Expecter) _setClientPeer(server interface{}) *MockSSMSAPRequirements__setClientPeer_Call {
	return &MockSSMSAPRequirements__setClientPeer_Call{Call: _e.mock.On("_setClientPeer", server)}
}

func (_c *MockSSMSAPRequirements__setClientPeer_Call) Run(run func(server _Server)) *MockSSMSAPRequirements__setClientPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_Server))
	})
	return _c
}

func (_c *MockSSMSAPRequirements__setClientPeer_Call) Return() *MockSSMSAPRequirements__setClientPeer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSSMSAPRequirements__setClientPeer_Call) RunAndReturn(run func(_Server)) *MockSSMSAPRequirements__setClientPeer_Call {
	_c.Call.Return(run)
	return _c
}

// _setServiceElement provides a mock function with given fields: serviceElement
func (_m *MockSSMSAPRequirements) _setServiceElement(serviceElement _ApplicationServiceElement) {
	_m.Called(serviceElement)
}

// MockSSMSAPRequirements__setServiceElement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method '_setServiceElement'
type MockSSMSAPRequirements__setServiceElement_Call struct {
	*mock.Call
}

// _setServiceElement is a helper method to define mock.On call
//   - serviceElement _ApplicationServiceElement
func (_e *MockSSMSAPRequirements_Expecter) _setServiceElement(serviceElement interface{}) *MockSSMSAPRequirements__setServiceElement_Call {
	return &MockSSMSAPRequirements__setServiceElement_Call{Call: _e.mock.On("_setServiceElement", serviceElement)}
}

func (_c *MockSSMSAPRequirements__setServiceElement_Call) Run(run func(serviceElement _ApplicationServiceElement)) *MockSSMSAPRequirements__setServiceElement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_ApplicationServiceElement))
	})
	return _c
}

func (_c *MockSSMSAPRequirements__setServiceElement_Call) Return() *MockSSMSAPRequirements__setServiceElement_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSSMSAPRequirements__setServiceElement_Call) RunAndReturn(run func(_ApplicationServiceElement)) *MockSSMSAPRequirements__setServiceElement_Call {
	_c.Call.Return(run)
	return _c
}

// getClientId provides a mock function with given fields:
func (_m *MockSSMSAPRequirements) getClientId() *int {
	ret := _m.Called()

	var r0 *int
	if rf, ok := ret.Get(0).(func() *int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	return r0
}

// MockSSMSAPRequirements_getClientId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getClientId'
type MockSSMSAPRequirements_getClientId_Call struct {
	*mock.Call
}

// getClientId is a helper method to define mock.On call
func (_e *MockSSMSAPRequirements_Expecter) getClientId() *MockSSMSAPRequirements_getClientId_Call {
	return &MockSSMSAPRequirements_getClientId_Call{Call: _e.mock.On("getClientId")}
}

func (_c *MockSSMSAPRequirements_getClientId_Call) Run(run func()) *MockSSMSAPRequirements_getClientId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSSMSAPRequirements_getClientId_Call) Return(_a0 *int) *MockSSMSAPRequirements_getClientId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSSMSAPRequirements_getClientId_Call) RunAndReturn(run func() *int) *MockSSMSAPRequirements_getClientId_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockSSMSAPRequirements interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSSMSAPRequirements creates a new instance of MockSSMSAPRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSSMSAPRequirements(t mockConstructorTestingTNewMockSSMSAPRequirements) *MockSSMSAPRequirements {
	mock := &MockSSMSAPRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
