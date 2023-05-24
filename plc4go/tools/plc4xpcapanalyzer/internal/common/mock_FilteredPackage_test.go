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

package common

import (
	gopacket "github.com/gopacket/gopacket"
	mock "github.com/stretchr/testify/mock"
)

// MockFilteredPackage is an autogenerated mock type for the FilteredPackage type
type MockFilteredPackage struct {
	mock.Mock
}

type MockFilteredPackage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFilteredPackage) EXPECT() *MockFilteredPackage_Expecter {
	return &MockFilteredPackage_Expecter{mock: &_m.Mock}
}

// ApplicationLayer provides a mock function with given fields:
func (_m *MockFilteredPackage) ApplicationLayer() gopacket.ApplicationLayer {
	ret := _m.Called()

	var r0 gopacket.ApplicationLayer
	if rf, ok := ret.Get(0).(func() gopacket.ApplicationLayer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gopacket.ApplicationLayer)
		}
	}

	return r0
}

// MockFilteredPackage_ApplicationLayer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplicationLayer'
type MockFilteredPackage_ApplicationLayer_Call struct {
	*mock.Call
}

// ApplicationLayer is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) ApplicationLayer() *MockFilteredPackage_ApplicationLayer_Call {
	return &MockFilteredPackage_ApplicationLayer_Call{Call: _e.mock.On("ApplicationLayer")}
}

func (_c *MockFilteredPackage_ApplicationLayer_Call) Run(run func()) *MockFilteredPackage_ApplicationLayer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_ApplicationLayer_Call) Return(_a0 gopacket.ApplicationLayer) *MockFilteredPackage_ApplicationLayer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_ApplicationLayer_Call) RunAndReturn(run func() gopacket.ApplicationLayer) *MockFilteredPackage_ApplicationLayer_Call {
	_c.Call.Return(run)
	return _c
}

// Data provides a mock function with given fields:
func (_m *MockFilteredPackage) Data() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockFilteredPackage_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type MockFilteredPackage_Data_Call struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) Data() *MockFilteredPackage_Data_Call {
	return &MockFilteredPackage_Data_Call{Call: _e.mock.On("Data")}
}

func (_c *MockFilteredPackage_Data_Call) Run(run func()) *MockFilteredPackage_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_Data_Call) Return(_a0 []byte) *MockFilteredPackage_Data_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_Data_Call) RunAndReturn(run func() []byte) *MockFilteredPackage_Data_Call {
	_c.Call.Return(run)
	return _c
}

// Dump provides a mock function with given fields:
func (_m *MockFilteredPackage) Dump() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockFilteredPackage_Dump_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Dump'
type MockFilteredPackage_Dump_Call struct {
	*mock.Call
}

// Dump is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) Dump() *MockFilteredPackage_Dump_Call {
	return &MockFilteredPackage_Dump_Call{Call: _e.mock.On("Dump")}
}

func (_c *MockFilteredPackage_Dump_Call) Run(run func()) *MockFilteredPackage_Dump_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_Dump_Call) Return(_a0 string) *MockFilteredPackage_Dump_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_Dump_Call) RunAndReturn(run func() string) *MockFilteredPackage_Dump_Call {
	_c.Call.Return(run)
	return _c
}

// ErrorLayer provides a mock function with given fields:
func (_m *MockFilteredPackage) ErrorLayer() gopacket.ErrorLayer {
	ret := _m.Called()

	var r0 gopacket.ErrorLayer
	if rf, ok := ret.Get(0).(func() gopacket.ErrorLayer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gopacket.ErrorLayer)
		}
	}

	return r0
}

// MockFilteredPackage_ErrorLayer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ErrorLayer'
type MockFilteredPackage_ErrorLayer_Call struct {
	*mock.Call
}

// ErrorLayer is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) ErrorLayer() *MockFilteredPackage_ErrorLayer_Call {
	return &MockFilteredPackage_ErrorLayer_Call{Call: _e.mock.On("ErrorLayer")}
}

func (_c *MockFilteredPackage_ErrorLayer_Call) Run(run func()) *MockFilteredPackage_ErrorLayer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_ErrorLayer_Call) Return(_a0 gopacket.ErrorLayer) *MockFilteredPackage_ErrorLayer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_ErrorLayer_Call) RunAndReturn(run func() gopacket.ErrorLayer) *MockFilteredPackage_ErrorLayer_Call {
	_c.Call.Return(run)
	return _c
}

// FilterReason provides a mock function with given fields:
func (_m *MockFilteredPackage) FilterReason() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFilteredPackage_FilterReason_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FilterReason'
type MockFilteredPackage_FilterReason_Call struct {
	*mock.Call
}

// FilterReason is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) FilterReason() *MockFilteredPackage_FilterReason_Call {
	return &MockFilteredPackage_FilterReason_Call{Call: _e.mock.On("FilterReason")}
}

func (_c *MockFilteredPackage_FilterReason_Call) Run(run func()) *MockFilteredPackage_FilterReason_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_FilterReason_Call) Return(_a0 error) *MockFilteredPackage_FilterReason_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_FilterReason_Call) RunAndReturn(run func() error) *MockFilteredPackage_FilterReason_Call {
	_c.Call.Return(run)
	return _c
}

// IsFilteredPackage provides a mock function with given fields:
func (_m *MockFilteredPackage) IsFilteredPackage() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockFilteredPackage_IsFilteredPackage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsFilteredPackage'
type MockFilteredPackage_IsFilteredPackage_Call struct {
	*mock.Call
}

// IsFilteredPackage is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) IsFilteredPackage() *MockFilteredPackage_IsFilteredPackage_Call {
	return &MockFilteredPackage_IsFilteredPackage_Call{Call: _e.mock.On("IsFilteredPackage")}
}

func (_c *MockFilteredPackage_IsFilteredPackage_Call) Run(run func()) *MockFilteredPackage_IsFilteredPackage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_IsFilteredPackage_Call) Return(_a0 bool) *MockFilteredPackage_IsFilteredPackage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_IsFilteredPackage_Call) RunAndReturn(run func() bool) *MockFilteredPackage_IsFilteredPackage_Call {
	_c.Call.Return(run)
	return _c
}

// Layer provides a mock function with given fields: _a0
func (_m *MockFilteredPackage) Layer(_a0 gopacket.LayerType) gopacket.Layer {
	ret := _m.Called(_a0)

	var r0 gopacket.Layer
	if rf, ok := ret.Get(0).(func(gopacket.LayerType) gopacket.Layer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gopacket.Layer)
		}
	}

	return r0
}

// MockFilteredPackage_Layer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Layer'
type MockFilteredPackage_Layer_Call struct {
	*mock.Call
}

// Layer is a helper method to define mock.On call
//   - _a0 gopacket.LayerType
func (_e *MockFilteredPackage_Expecter) Layer(_a0 interface{}) *MockFilteredPackage_Layer_Call {
	return &MockFilteredPackage_Layer_Call{Call: _e.mock.On("Layer", _a0)}
}

func (_c *MockFilteredPackage_Layer_Call) Run(run func(_a0 gopacket.LayerType)) *MockFilteredPackage_Layer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(gopacket.LayerType))
	})
	return _c
}

func (_c *MockFilteredPackage_Layer_Call) Return(_a0 gopacket.Layer) *MockFilteredPackage_Layer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_Layer_Call) RunAndReturn(run func(gopacket.LayerType) gopacket.Layer) *MockFilteredPackage_Layer_Call {
	_c.Call.Return(run)
	return _c
}

// LayerClass provides a mock function with given fields: _a0
func (_m *MockFilteredPackage) LayerClass(_a0 gopacket.LayerClass) gopacket.Layer {
	ret := _m.Called(_a0)

	var r0 gopacket.Layer
	if rf, ok := ret.Get(0).(func(gopacket.LayerClass) gopacket.Layer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gopacket.Layer)
		}
	}

	return r0
}

// MockFilteredPackage_LayerClass_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LayerClass'
type MockFilteredPackage_LayerClass_Call struct {
	*mock.Call
}

// LayerClass is a helper method to define mock.On call
//   - _a0 gopacket.LayerClass
func (_e *MockFilteredPackage_Expecter) LayerClass(_a0 interface{}) *MockFilteredPackage_LayerClass_Call {
	return &MockFilteredPackage_LayerClass_Call{Call: _e.mock.On("LayerClass", _a0)}
}

func (_c *MockFilteredPackage_LayerClass_Call) Run(run func(_a0 gopacket.LayerClass)) *MockFilteredPackage_LayerClass_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(gopacket.LayerClass))
	})
	return _c
}

func (_c *MockFilteredPackage_LayerClass_Call) Return(_a0 gopacket.Layer) *MockFilteredPackage_LayerClass_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_LayerClass_Call) RunAndReturn(run func(gopacket.LayerClass) gopacket.Layer) *MockFilteredPackage_LayerClass_Call {
	_c.Call.Return(run)
	return _c
}

// Layers provides a mock function with given fields:
func (_m *MockFilteredPackage) Layers() []gopacket.Layer {
	ret := _m.Called()

	var r0 []gopacket.Layer
	if rf, ok := ret.Get(0).(func() []gopacket.Layer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gopacket.Layer)
		}
	}

	return r0
}

// MockFilteredPackage_Layers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Layers'
type MockFilteredPackage_Layers_Call struct {
	*mock.Call
}

// Layers is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) Layers() *MockFilteredPackage_Layers_Call {
	return &MockFilteredPackage_Layers_Call{Call: _e.mock.On("Layers")}
}

func (_c *MockFilteredPackage_Layers_Call) Run(run func()) *MockFilteredPackage_Layers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_Layers_Call) Return(_a0 []gopacket.Layer) *MockFilteredPackage_Layers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_Layers_Call) RunAndReturn(run func() []gopacket.Layer) *MockFilteredPackage_Layers_Call {
	_c.Call.Return(run)
	return _c
}

// LinkLayer provides a mock function with given fields:
func (_m *MockFilteredPackage) LinkLayer() gopacket.LinkLayer {
	ret := _m.Called()

	var r0 gopacket.LinkLayer
	if rf, ok := ret.Get(0).(func() gopacket.LinkLayer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gopacket.LinkLayer)
		}
	}

	return r0
}

// MockFilteredPackage_LinkLayer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LinkLayer'
type MockFilteredPackage_LinkLayer_Call struct {
	*mock.Call
}

// LinkLayer is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) LinkLayer() *MockFilteredPackage_LinkLayer_Call {
	return &MockFilteredPackage_LinkLayer_Call{Call: _e.mock.On("LinkLayer")}
}

func (_c *MockFilteredPackage_LinkLayer_Call) Run(run func()) *MockFilteredPackage_LinkLayer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_LinkLayer_Call) Return(_a0 gopacket.LinkLayer) *MockFilteredPackage_LinkLayer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_LinkLayer_Call) RunAndReturn(run func() gopacket.LinkLayer) *MockFilteredPackage_LinkLayer_Call {
	_c.Call.Return(run)
	return _c
}

// Metadata provides a mock function with given fields:
func (_m *MockFilteredPackage) Metadata() *gopacket.PacketMetadata {
	ret := _m.Called()

	var r0 *gopacket.PacketMetadata
	if rf, ok := ret.Get(0).(func() *gopacket.PacketMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gopacket.PacketMetadata)
		}
	}

	return r0
}

// MockFilteredPackage_Metadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Metadata'
type MockFilteredPackage_Metadata_Call struct {
	*mock.Call
}

// Metadata is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) Metadata() *MockFilteredPackage_Metadata_Call {
	return &MockFilteredPackage_Metadata_Call{Call: _e.mock.On("Metadata")}
}

func (_c *MockFilteredPackage_Metadata_Call) Run(run func()) *MockFilteredPackage_Metadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_Metadata_Call) Return(_a0 *gopacket.PacketMetadata) *MockFilteredPackage_Metadata_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_Metadata_Call) RunAndReturn(run func() *gopacket.PacketMetadata) *MockFilteredPackage_Metadata_Call {
	_c.Call.Return(run)
	return _c
}

// NetworkLayer provides a mock function with given fields:
func (_m *MockFilteredPackage) NetworkLayer() gopacket.NetworkLayer {
	ret := _m.Called()

	var r0 gopacket.NetworkLayer
	if rf, ok := ret.Get(0).(func() gopacket.NetworkLayer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gopacket.NetworkLayer)
		}
	}

	return r0
}

// MockFilteredPackage_NetworkLayer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NetworkLayer'
type MockFilteredPackage_NetworkLayer_Call struct {
	*mock.Call
}

// NetworkLayer is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) NetworkLayer() *MockFilteredPackage_NetworkLayer_Call {
	return &MockFilteredPackage_NetworkLayer_Call{Call: _e.mock.On("NetworkLayer")}
}

func (_c *MockFilteredPackage_NetworkLayer_Call) Run(run func()) *MockFilteredPackage_NetworkLayer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_NetworkLayer_Call) Return(_a0 gopacket.NetworkLayer) *MockFilteredPackage_NetworkLayer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_NetworkLayer_Call) RunAndReturn(run func() gopacket.NetworkLayer) *MockFilteredPackage_NetworkLayer_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockFilteredPackage) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockFilteredPackage_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockFilteredPackage_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) String() *MockFilteredPackage_String_Call {
	return &MockFilteredPackage_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockFilteredPackage_String_Call) Run(run func()) *MockFilteredPackage_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_String_Call) Return(_a0 string) *MockFilteredPackage_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_String_Call) RunAndReturn(run func() string) *MockFilteredPackage_String_Call {
	_c.Call.Return(run)
	return _c
}

// TransportLayer provides a mock function with given fields:
func (_m *MockFilteredPackage) TransportLayer() gopacket.TransportLayer {
	ret := _m.Called()

	var r0 gopacket.TransportLayer
	if rf, ok := ret.Get(0).(func() gopacket.TransportLayer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gopacket.TransportLayer)
		}
	}

	return r0
}

// MockFilteredPackage_TransportLayer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransportLayer'
type MockFilteredPackage_TransportLayer_Call struct {
	*mock.Call
}

// TransportLayer is a helper method to define mock.On call
func (_e *MockFilteredPackage_Expecter) TransportLayer() *MockFilteredPackage_TransportLayer_Call {
	return &MockFilteredPackage_TransportLayer_Call{Call: _e.mock.On("TransportLayer")}
}

func (_c *MockFilteredPackage_TransportLayer_Call) Run(run func()) *MockFilteredPackage_TransportLayer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilteredPackage_TransportLayer_Call) Return(_a0 gopacket.TransportLayer) *MockFilteredPackage_TransportLayer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilteredPackage_TransportLayer_Call) RunAndReturn(run func() gopacket.TransportLayer) *MockFilteredPackage_TransportLayer_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockFilteredPackage interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockFilteredPackage creates a new instance of MockFilteredPackage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockFilteredPackage(t mockConstructorTestingTNewMockFilteredPackage) *MockFilteredPackage {
	mock := &MockFilteredPackage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
