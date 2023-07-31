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

package simulated

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	readwritemodel "github.com/apache/plc4x/plc4go/protocols/simulated/readwrite/model"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockTag is an autogenerated mock type for the Tag type
type MockTag struct {
	mock.Mock
}

type MockTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTag) EXPECT() *MockTag_Expecter {
	return &MockTag_Expecter{mock: &_m.Mock}
}

// GetAddressString provides a mock function with given fields:
func (_m *MockTag) GetAddressString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTag_GetAddressString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddressString'
type MockTag_GetAddressString_Call struct {
	*mock.Call
}

// GetAddressString is a helper method to define mock.On call
func (_e *MockTag_Expecter) GetAddressString() *MockTag_GetAddressString_Call {
	return &MockTag_GetAddressString_Call{Call: _e.mock.On("GetAddressString")}
}

func (_c *MockTag_GetAddressString_Call) Run(run func()) *MockTag_GetAddressString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTag_GetAddressString_Call) Return(_a0 string) *MockTag_GetAddressString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTag_GetAddressString_Call) RunAndReturn(run func() string) *MockTag_GetAddressString_Call {
	_c.Call.Return(run)
	return _c
}

// GetArrayInfo provides a mock function with given fields:
func (_m *MockTag) GetArrayInfo() []model.ArrayInfo {
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

// MockTag_GetArrayInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArrayInfo'
type MockTag_GetArrayInfo_Call struct {
	*mock.Call
}

// GetArrayInfo is a helper method to define mock.On call
func (_e *MockTag_Expecter) GetArrayInfo() *MockTag_GetArrayInfo_Call {
	return &MockTag_GetArrayInfo_Call{Call: _e.mock.On("GetArrayInfo")}
}

func (_c *MockTag_GetArrayInfo_Call) Run(run func()) *MockTag_GetArrayInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTag_GetArrayInfo_Call) Return(_a0 []model.ArrayInfo) *MockTag_GetArrayInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTag_GetArrayInfo_Call) RunAndReturn(run func() []model.ArrayInfo) *MockTag_GetArrayInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetDataTypeSize provides a mock function with given fields:
func (_m *MockTag) GetDataTypeSize() readwritemodel.SimulatedDataTypeSizes {
	ret := _m.Called()

	var r0 readwritemodel.SimulatedDataTypeSizes
	if rf, ok := ret.Get(0).(func() readwritemodel.SimulatedDataTypeSizes); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(readwritemodel.SimulatedDataTypeSizes)
	}

	return r0
}

// MockTag_GetDataTypeSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataTypeSize'
type MockTag_GetDataTypeSize_Call struct {
	*mock.Call
}

// GetDataTypeSize is a helper method to define mock.On call
func (_e *MockTag_Expecter) GetDataTypeSize() *MockTag_GetDataTypeSize_Call {
	return &MockTag_GetDataTypeSize_Call{Call: _e.mock.On("GetDataTypeSize")}
}

func (_c *MockTag_GetDataTypeSize_Call) Run(run func()) *MockTag_GetDataTypeSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTag_GetDataTypeSize_Call) Return(_a0 readwritemodel.SimulatedDataTypeSizes) *MockTag_GetDataTypeSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTag_GetDataTypeSize_Call) RunAndReturn(run func() readwritemodel.SimulatedDataTypeSizes) *MockTag_GetDataTypeSize_Call {
	_c.Call.Return(run)
	return _c
}

// GetName provides a mock function with given fields:
func (_m *MockTag) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTag_GetName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetName'
type MockTag_GetName_Call struct {
	*mock.Call
}

// GetName is a helper method to define mock.On call
func (_e *MockTag_Expecter) GetName() *MockTag_GetName_Call {
	return &MockTag_GetName_Call{Call: _e.mock.On("GetName")}
}

func (_c *MockTag_GetName_Call) Run(run func()) *MockTag_GetName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTag_GetName_Call) Return(_a0 string) *MockTag_GetName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTag_GetName_Call) RunAndReturn(run func() string) *MockTag_GetName_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagType provides a mock function with given fields:
func (_m *MockTag) GetTagType() TagType {
	ret := _m.Called()

	var r0 TagType
	if rf, ok := ret.Get(0).(func() TagType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(TagType)
	}

	return r0
}

// MockTag_GetTagType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagType'
type MockTag_GetTagType_Call struct {
	*mock.Call
}

// GetTagType is a helper method to define mock.On call
func (_e *MockTag_Expecter) GetTagType() *MockTag_GetTagType_Call {
	return &MockTag_GetTagType_Call{Call: _e.mock.On("GetTagType")}
}

func (_c *MockTag_GetTagType_Call) Run(run func()) *MockTag_GetTagType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTag_GetTagType_Call) Return(_a0 TagType) *MockTag_GetTagType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTag_GetTagType_Call) RunAndReturn(run func() TagType) *MockTag_GetTagType_Call {
	_c.Call.Return(run)
	return _c
}

// GetValueType provides a mock function with given fields:
func (_m *MockTag) GetValueType() values.PlcValueType {
	ret := _m.Called()

	var r0 values.PlcValueType
	if rf, ok := ret.Get(0).(func() values.PlcValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(values.PlcValueType)
	}

	return r0
}

// MockTag_GetValueType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValueType'
type MockTag_GetValueType_Call struct {
	*mock.Call
}

// GetValueType is a helper method to define mock.On call
func (_e *MockTag_Expecter) GetValueType() *MockTag_GetValueType_Call {
	return &MockTag_GetValueType_Call{Call: _e.mock.On("GetValueType")}
}

func (_c *MockTag_GetValueType_Call) Run(run func()) *MockTag_GetValueType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTag_GetValueType_Call) Return(_a0 values.PlcValueType) *MockTag_GetValueType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTag_GetValueType_Call) RunAndReturn(run func() values.PlcValueType) *MockTag_GetValueType_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockTag) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTag_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockTag_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockTag_Expecter) String() *MockTag_String_Call {
	return &MockTag_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockTag_String_Call) Run(run func()) *MockTag_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTag_String_Call) Return(_a0 string) *MockTag_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTag_String_Call) RunAndReturn(run func() string) *MockTag_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTag creates a new instance of MockTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTag(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTag {
	mock := &MockTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
