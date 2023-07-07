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

// Code generated by mockery v2.30.17. DO NOT EDIT.

package knxnetip

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	readwritemodel "github.com/apache/plc4x/plc4go/protocols/knxnetip/readwrite/model"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockDeviceTag is an autogenerated mock type for the DeviceTag type
type MockDeviceTag struct {
	mock.Mock
}

type MockDeviceTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDeviceTag) EXPECT() *MockDeviceTag_Expecter {
	return &MockDeviceTag_Expecter{mock: &_m.Mock}
}

// GetAddressString provides a mock function with given fields:
func (_m *MockDeviceTag) GetAddressString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockDeviceTag_GetAddressString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddressString'
type MockDeviceTag_GetAddressString_Call struct {
	*mock.Call
}

// GetAddressString is a helper method to define mock.On call
func (_e *MockDeviceTag_Expecter) GetAddressString() *MockDeviceTag_GetAddressString_Call {
	return &MockDeviceTag_GetAddressString_Call{Call: _e.mock.On("GetAddressString")}
}

func (_c *MockDeviceTag_GetAddressString_Call) Run(run func()) *MockDeviceTag_GetAddressString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDeviceTag_GetAddressString_Call) Return(_a0 string) *MockDeviceTag_GetAddressString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeviceTag_GetAddressString_Call) RunAndReturn(run func() string) *MockDeviceTag_GetAddressString_Call {
	_c.Call.Return(run)
	return _c
}

// GetArrayInfo provides a mock function with given fields:
func (_m *MockDeviceTag) GetArrayInfo() []model.ArrayInfo {
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

// MockDeviceTag_GetArrayInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArrayInfo'
type MockDeviceTag_GetArrayInfo_Call struct {
	*mock.Call
}

// GetArrayInfo is a helper method to define mock.On call
func (_e *MockDeviceTag_Expecter) GetArrayInfo() *MockDeviceTag_GetArrayInfo_Call {
	return &MockDeviceTag_GetArrayInfo_Call{Call: _e.mock.On("GetArrayInfo")}
}

func (_c *MockDeviceTag_GetArrayInfo_Call) Run(run func()) *MockDeviceTag_GetArrayInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDeviceTag_GetArrayInfo_Call) Return(_a0 []model.ArrayInfo) *MockDeviceTag_GetArrayInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeviceTag_GetArrayInfo_Call) RunAndReturn(run func() []model.ArrayInfo) *MockDeviceTag_GetArrayInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetValueType provides a mock function with given fields:
func (_m *MockDeviceTag) GetValueType() values.PlcValueType {
	ret := _m.Called()

	var r0 values.PlcValueType
	if rf, ok := ret.Get(0).(func() values.PlcValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(values.PlcValueType)
	}

	return r0
}

// MockDeviceTag_GetValueType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValueType'
type MockDeviceTag_GetValueType_Call struct {
	*mock.Call
}

// GetValueType is a helper method to define mock.On call
func (_e *MockDeviceTag_Expecter) GetValueType() *MockDeviceTag_GetValueType_Call {
	return &MockDeviceTag_GetValueType_Call{Call: _e.mock.On("GetValueType")}
}

func (_c *MockDeviceTag_GetValueType_Call) Run(run func()) *MockDeviceTag_GetValueType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDeviceTag_GetValueType_Call) Return(_a0 values.PlcValueType) *MockDeviceTag_GetValueType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeviceTag_GetValueType_Call) RunAndReturn(run func() values.PlcValueType) *MockDeviceTag_GetValueType_Call {
	_c.Call.Return(run)
	return _c
}

// toKnxAddress provides a mock function with given fields:
func (_m *MockDeviceTag) toKnxAddress() readwritemodel.KnxAddress {
	ret := _m.Called()

	var r0 readwritemodel.KnxAddress
	if rf, ok := ret.Get(0).(func() readwritemodel.KnxAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(readwritemodel.KnxAddress)
		}
	}

	return r0
}

// MockDeviceTag_toKnxAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'toKnxAddress'
type MockDeviceTag_toKnxAddress_Call struct {
	*mock.Call
}

// toKnxAddress is a helper method to define mock.On call
func (_e *MockDeviceTag_Expecter) toKnxAddress() *MockDeviceTag_toKnxAddress_Call {
	return &MockDeviceTag_toKnxAddress_Call{Call: _e.mock.On("toKnxAddress")}
}

func (_c *MockDeviceTag_toKnxAddress_Call) Run(run func()) *MockDeviceTag_toKnxAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDeviceTag_toKnxAddress_Call) Return(_a0 readwritemodel.KnxAddress) *MockDeviceTag_toKnxAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeviceTag_toKnxAddress_Call) RunAndReturn(run func() readwritemodel.KnxAddress) *MockDeviceTag_toKnxAddress_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDeviceTag creates a new instance of MockDeviceTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDeviceTag(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDeviceTag {
	mock := &MockDeviceTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
