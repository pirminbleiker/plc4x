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

package s7

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	readwritemodel "github.com/apache/plc4x/plc4go/protocols/s7/readwrite/model"

	utils "github.com/apache/plc4x/plc4go/spi/utils"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockPlcTag is an autogenerated mock type for the PlcTag type
type MockPlcTag struct {
	mock.Mock
}

type MockPlcTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcTag) EXPECT() *MockPlcTag_Expecter {
	return &MockPlcTag_Expecter{mock: &_m.Mock}
}

// GetAddressString provides a mock function with given fields:
func (_m *MockPlcTag) GetAddressString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcTag_GetAddressString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddressString'
type MockPlcTag_GetAddressString_Call struct {
	*mock.Call
}

// GetAddressString is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetAddressString() *MockPlcTag_GetAddressString_Call {
	return &MockPlcTag_GetAddressString_Call{Call: _e.mock.On("GetAddressString")}
}

func (_c *MockPlcTag_GetAddressString_Call) Run(run func()) *MockPlcTag_GetAddressString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetAddressString_Call) Return(_a0 string) *MockPlcTag_GetAddressString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetAddressString_Call) RunAndReturn(run func() string) *MockPlcTag_GetAddressString_Call {
	_c.Call.Return(run)
	return _c
}

// GetArrayInfo provides a mock function with given fields:
func (_m *MockPlcTag) GetArrayInfo() []model.ArrayInfo {
	ret := _m.Called()

	var r0 []model.ArrayInfo
	if rf, ok := ret.Get(0).(func() []model.ArrayInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ArrayInfo)
		}
	}

	return r0
}

// MockPlcTag_GetArrayInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArrayInfo'
type MockPlcTag_GetArrayInfo_Call struct {
	*mock.Call
}

// GetArrayInfo is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetArrayInfo() *MockPlcTag_GetArrayInfo_Call {
	return &MockPlcTag_GetArrayInfo_Call{Call: _e.mock.On("GetArrayInfo")}
}

func (_c *MockPlcTag_GetArrayInfo_Call) Run(run func()) *MockPlcTag_GetArrayInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetArrayInfo_Call) Return(_a0 []model.ArrayInfo) *MockPlcTag_GetArrayInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetArrayInfo_Call) RunAndReturn(run func() []model.ArrayInfo) *MockPlcTag_GetArrayInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetBitOffset provides a mock function with given fields:
func (_m *MockPlcTag) GetBitOffset() uint8 {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}

// MockPlcTag_GetBitOffset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBitOffset'
type MockPlcTag_GetBitOffset_Call struct {
	*mock.Call
}

// GetBitOffset is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetBitOffset() *MockPlcTag_GetBitOffset_Call {
	return &MockPlcTag_GetBitOffset_Call{Call: _e.mock.On("GetBitOffset")}
}

func (_c *MockPlcTag_GetBitOffset_Call) Run(run func()) *MockPlcTag_GetBitOffset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetBitOffset_Call) Return(_a0 uint8) *MockPlcTag_GetBitOffset_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetBitOffset_Call) RunAndReturn(run func() uint8) *MockPlcTag_GetBitOffset_Call {
	_c.Call.Return(run)
	return _c
}

// GetBlockNumber provides a mock function with given fields:
func (_m *MockPlcTag) GetBlockNumber() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockPlcTag_GetBlockNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockNumber'
type MockPlcTag_GetBlockNumber_Call struct {
	*mock.Call
}

// GetBlockNumber is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetBlockNumber() *MockPlcTag_GetBlockNumber_Call {
	return &MockPlcTag_GetBlockNumber_Call{Call: _e.mock.On("GetBlockNumber")}
}

func (_c *MockPlcTag_GetBlockNumber_Call) Run(run func()) *MockPlcTag_GetBlockNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetBlockNumber_Call) Return(_a0 uint16) *MockPlcTag_GetBlockNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetBlockNumber_Call) RunAndReturn(run func() uint16) *MockPlcTag_GetBlockNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetByteOffset provides a mock function with given fields:
func (_m *MockPlcTag) GetByteOffset() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockPlcTag_GetByteOffset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByteOffset'
type MockPlcTag_GetByteOffset_Call struct {
	*mock.Call
}

// GetByteOffset is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetByteOffset() *MockPlcTag_GetByteOffset_Call {
	return &MockPlcTag_GetByteOffset_Call{Call: _e.mock.On("GetByteOffset")}
}

func (_c *MockPlcTag_GetByteOffset_Call) Run(run func()) *MockPlcTag_GetByteOffset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetByteOffset_Call) Return(_a0 uint16) *MockPlcTag_GetByteOffset_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetByteOffset_Call) RunAndReturn(run func() uint16) *MockPlcTag_GetByteOffset_Call {
	_c.Call.Return(run)
	return _c
}

// GetDataType provides a mock function with given fields:
func (_m *MockPlcTag) GetDataType() readwritemodel.TransportSize {
	ret := _m.Called()

	var r0 readwritemodel.TransportSize
	if rf, ok := ret.Get(0).(func() readwritemodel.TransportSize); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(readwritemodel.TransportSize)
	}

	return r0
}

// MockPlcTag_GetDataType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataType'
type MockPlcTag_GetDataType_Call struct {
	*mock.Call
}

// GetDataType is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetDataType() *MockPlcTag_GetDataType_Call {
	return &MockPlcTag_GetDataType_Call{Call: _e.mock.On("GetDataType")}
}

func (_c *MockPlcTag_GetDataType_Call) Run(run func()) *MockPlcTag_GetDataType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetDataType_Call) Return(_a0 readwritemodel.TransportSize) *MockPlcTag_GetDataType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetDataType_Call) RunAndReturn(run func() readwritemodel.TransportSize) *MockPlcTag_GetDataType_Call {
	_c.Call.Return(run)
	return _c
}

// GetMemoryArea provides a mock function with given fields:
func (_m *MockPlcTag) GetMemoryArea() readwritemodel.MemoryArea {
	ret := _m.Called()

	var r0 readwritemodel.MemoryArea
	if rf, ok := ret.Get(0).(func() readwritemodel.MemoryArea); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(readwritemodel.MemoryArea)
	}

	return r0
}

// MockPlcTag_GetMemoryArea_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMemoryArea'
type MockPlcTag_GetMemoryArea_Call struct {
	*mock.Call
}

// GetMemoryArea is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetMemoryArea() *MockPlcTag_GetMemoryArea_Call {
	return &MockPlcTag_GetMemoryArea_Call{Call: _e.mock.On("GetMemoryArea")}
}

func (_c *MockPlcTag_GetMemoryArea_Call) Run(run func()) *MockPlcTag_GetMemoryArea_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetMemoryArea_Call) Return(_a0 readwritemodel.MemoryArea) *MockPlcTag_GetMemoryArea_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetMemoryArea_Call) RunAndReturn(run func() readwritemodel.MemoryArea) *MockPlcTag_GetMemoryArea_Call {
	_c.Call.Return(run)
	return _c
}

// GetNumElements provides a mock function with given fields:
func (_m *MockPlcTag) GetNumElements() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockPlcTag_GetNumElements_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNumElements'
type MockPlcTag_GetNumElements_Call struct {
	*mock.Call
}

// GetNumElements is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetNumElements() *MockPlcTag_GetNumElements_Call {
	return &MockPlcTag_GetNumElements_Call{Call: _e.mock.On("GetNumElements")}
}

func (_c *MockPlcTag_GetNumElements_Call) Run(run func()) *MockPlcTag_GetNumElements_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetNumElements_Call) Return(_a0 uint16) *MockPlcTag_GetNumElements_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetNumElements_Call) RunAndReturn(run func() uint16) *MockPlcTag_GetNumElements_Call {
	_c.Call.Return(run)
	return _c
}

// GetValueType provides a mock function with given fields:
func (_m *MockPlcTag) GetValueType() values.PlcValueType {
	ret := _m.Called()

	var r0 values.PlcValueType
	if rf, ok := ret.Get(0).(func() values.PlcValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(values.PlcValueType)
	}

	return r0
}

// MockPlcTag_GetValueType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValueType'
type MockPlcTag_GetValueType_Call struct {
	*mock.Call
}

// GetValueType is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) GetValueType() *MockPlcTag_GetValueType_Call {
	return &MockPlcTag_GetValueType_Call{Call: _e.mock.On("GetValueType")}
}

func (_c *MockPlcTag_GetValueType_Call) Run(run func()) *MockPlcTag_GetValueType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_GetValueType_Call) Return(_a0 values.PlcValueType) *MockPlcTag_GetValueType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_GetValueType_Call) RunAndReturn(run func() values.PlcValueType) *MockPlcTag_GetValueType_Call {
	_c.Call.Return(run)
	return _c
}

// Serialize provides a mock function with given fields:
func (_m *MockPlcTag) Serialize() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlcTag_Serialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serialize'
type MockPlcTag_Serialize_Call struct {
	*mock.Call
}

// Serialize is a helper method to define mock.On call
func (_e *MockPlcTag_Expecter) Serialize() *MockPlcTag_Serialize_Call {
	return &MockPlcTag_Serialize_Call{Call: _e.mock.On("Serialize")}
}

func (_c *MockPlcTag_Serialize_Call) Run(run func()) *MockPlcTag_Serialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcTag_Serialize_Call) Return(_a0 []byte, _a1 error) *MockPlcTag_Serialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlcTag_Serialize_Call) RunAndReturn(run func() ([]byte, error)) *MockPlcTag_Serialize_Call {
	_c.Call.Return(run)
	return _c
}

// SerializeWithWriteBuffer provides a mock function with given fields: ctx, writeBuffer
func (_m *MockPlcTag) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	ret := _m.Called(ctx, writeBuffer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, utils.WriteBuffer) error); ok {
		r0 = rf(ctx, writeBuffer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcTag_SerializeWithWriteBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SerializeWithWriteBuffer'
type MockPlcTag_SerializeWithWriteBuffer_Call struct {
	*mock.Call
}

// SerializeWithWriteBuffer is a helper method to define mock.On call
//   - ctx context.Context
//   - writeBuffer utils.WriteBuffer
func (_e *MockPlcTag_Expecter) SerializeWithWriteBuffer(ctx interface{}, writeBuffer interface{}) *MockPlcTag_SerializeWithWriteBuffer_Call {
	return &MockPlcTag_SerializeWithWriteBuffer_Call{Call: _e.mock.On("SerializeWithWriteBuffer", ctx, writeBuffer)}
}

func (_c *MockPlcTag_SerializeWithWriteBuffer_Call) Run(run func(ctx context.Context, writeBuffer utils.WriteBuffer)) *MockPlcTag_SerializeWithWriteBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(utils.WriteBuffer))
	})
	return _c
}

func (_c *MockPlcTag_SerializeWithWriteBuffer_Call) Return(_a0 error) *MockPlcTag_SerializeWithWriteBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcTag_SerializeWithWriteBuffer_Call) RunAndReturn(run func(context.Context, utils.WriteBuffer) error) *MockPlcTag_SerializeWithWriteBuffer_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcTag interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcTag creates a new instance of MockPlcTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcTag(t mockConstructorTestingTNewMockPlcTag) *MockPlcTag {
	mock := &MockPlcTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
