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

// Code generated by mockery v2.23.4. DO NOT EDIT.

package cbus

import (
	apimodel "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	model "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"

	time "time"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockMMIMonitorTag is an autogenerated mock type for the MMIMonitorTag type
type MockMMIMonitorTag struct {
	mock.Mock
}

type MockMMIMonitorTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMMIMonitorTag) EXPECT() *MockMMIMonitorTag_Expecter {
	return &MockMMIMonitorTag_Expecter{mock: &_m.Mock}
}

// GetAddressString provides a mock function with given fields:
func (_m *MockMMIMonitorTag) GetAddressString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockMMIMonitorTag_GetAddressString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddressString'
type MockMMIMonitorTag_GetAddressString_Call struct {
	*mock.Call
}

// GetAddressString is a helper method to define mock.On call
func (_e *MockMMIMonitorTag_Expecter) GetAddressString() *MockMMIMonitorTag_GetAddressString_Call {
	return &MockMMIMonitorTag_GetAddressString_Call{Call: _e.mock.On("GetAddressString")}
}

func (_c *MockMMIMonitorTag_GetAddressString_Call) Run(run func()) *MockMMIMonitorTag_GetAddressString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMMIMonitorTag_GetAddressString_Call) Return(_a0 string) *MockMMIMonitorTag_GetAddressString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMMIMonitorTag_GetAddressString_Call) RunAndReturn(run func() string) *MockMMIMonitorTag_GetAddressString_Call {
	_c.Call.Return(run)
	return _c
}

// GetApplication provides a mock function with given fields:
func (_m *MockMMIMonitorTag) GetApplication() *model.ApplicationIdContainer {
	ret := _m.Called()

	var r0 *model.ApplicationIdContainer
	if rf, ok := ret.Get(0).(func() *model.ApplicationIdContainer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ApplicationIdContainer)
		}
	}

	return r0
}

// MockMMIMonitorTag_GetApplication_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetApplication'
type MockMMIMonitorTag_GetApplication_Call struct {
	*mock.Call
}

// GetApplication is a helper method to define mock.On call
func (_e *MockMMIMonitorTag_Expecter) GetApplication() *MockMMIMonitorTag_GetApplication_Call {
	return &MockMMIMonitorTag_GetApplication_Call{Call: _e.mock.On("GetApplication")}
}

func (_c *MockMMIMonitorTag_GetApplication_Call) Run(run func()) *MockMMIMonitorTag_GetApplication_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMMIMonitorTag_GetApplication_Call) Return(_a0 *model.ApplicationIdContainer) *MockMMIMonitorTag_GetApplication_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMMIMonitorTag_GetApplication_Call) RunAndReturn(run func() *model.ApplicationIdContainer) *MockMMIMonitorTag_GetApplication_Call {
	_c.Call.Return(run)
	return _c
}

// GetArrayInfo provides a mock function with given fields:
func (_m *MockMMIMonitorTag) GetArrayInfo() []apimodel.ArrayInfo {
	ret := _m.Called()

	var r0 []apimodel.ArrayInfo
	if rf, ok := ret.Get(0).(func() []apimodel.ArrayInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]apimodel.ArrayInfo)
		}
	}

	return r0
}

// MockMMIMonitorTag_GetArrayInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArrayInfo'
type MockMMIMonitorTag_GetArrayInfo_Call struct {
	*mock.Call
}

// GetArrayInfo is a helper method to define mock.On call
func (_e *MockMMIMonitorTag_Expecter) GetArrayInfo() *MockMMIMonitorTag_GetArrayInfo_Call {
	return &MockMMIMonitorTag_GetArrayInfo_Call{Call: _e.mock.On("GetArrayInfo")}
}

func (_c *MockMMIMonitorTag_GetArrayInfo_Call) Run(run func()) *MockMMIMonitorTag_GetArrayInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMMIMonitorTag_GetArrayInfo_Call) Return(_a0 []apimodel.ArrayInfo) *MockMMIMonitorTag_GetArrayInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMMIMonitorTag_GetArrayInfo_Call) RunAndReturn(run func() []apimodel.ArrayInfo) *MockMMIMonitorTag_GetArrayInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetDuration provides a mock function with given fields:
func (_m *MockMMIMonitorTag) GetDuration() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// MockMMIMonitorTag_GetDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDuration'
type MockMMIMonitorTag_GetDuration_Call struct {
	*mock.Call
}

// GetDuration is a helper method to define mock.On call
func (_e *MockMMIMonitorTag_Expecter) GetDuration() *MockMMIMonitorTag_GetDuration_Call {
	return &MockMMIMonitorTag_GetDuration_Call{Call: _e.mock.On("GetDuration")}
}

func (_c *MockMMIMonitorTag_GetDuration_Call) Run(run func()) *MockMMIMonitorTag_GetDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMMIMonitorTag_GetDuration_Call) Return(_a0 time.Duration) *MockMMIMonitorTag_GetDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMMIMonitorTag_GetDuration_Call) RunAndReturn(run func() time.Duration) *MockMMIMonitorTag_GetDuration_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlcSubscriptionType provides a mock function with given fields:
func (_m *MockMMIMonitorTag) GetPlcSubscriptionType() apimodel.PlcSubscriptionType {
	ret := _m.Called()

	var r0 apimodel.PlcSubscriptionType
	if rf, ok := ret.Get(0).(func() apimodel.PlcSubscriptionType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(apimodel.PlcSubscriptionType)
	}

	return r0
}

// MockMMIMonitorTag_GetPlcSubscriptionType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlcSubscriptionType'
type MockMMIMonitorTag_GetPlcSubscriptionType_Call struct {
	*mock.Call
}

// GetPlcSubscriptionType is a helper method to define mock.On call
func (_e *MockMMIMonitorTag_Expecter) GetPlcSubscriptionType() *MockMMIMonitorTag_GetPlcSubscriptionType_Call {
	return &MockMMIMonitorTag_GetPlcSubscriptionType_Call{Call: _e.mock.On("GetPlcSubscriptionType")}
}

func (_c *MockMMIMonitorTag_GetPlcSubscriptionType_Call) Run(run func()) *MockMMIMonitorTag_GetPlcSubscriptionType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMMIMonitorTag_GetPlcSubscriptionType_Call) Return(_a0 apimodel.PlcSubscriptionType) *MockMMIMonitorTag_GetPlcSubscriptionType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMMIMonitorTag_GetPlcSubscriptionType_Call) RunAndReturn(run func() apimodel.PlcSubscriptionType) *MockMMIMonitorTag_GetPlcSubscriptionType_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagType provides a mock function with given fields:
func (_m *MockMMIMonitorTag) GetTagType() TagType {
	ret := _m.Called()

	var r0 TagType
	if rf, ok := ret.Get(0).(func() TagType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(TagType)
	}

	return r0
}

// MockMMIMonitorTag_GetTagType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagType'
type MockMMIMonitorTag_GetTagType_Call struct {
	*mock.Call
}

// GetTagType is a helper method to define mock.On call
func (_e *MockMMIMonitorTag_Expecter) GetTagType() *MockMMIMonitorTag_GetTagType_Call {
	return &MockMMIMonitorTag_GetTagType_Call{Call: _e.mock.On("GetTagType")}
}

func (_c *MockMMIMonitorTag_GetTagType_Call) Run(run func()) *MockMMIMonitorTag_GetTagType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMMIMonitorTag_GetTagType_Call) Return(_a0 TagType) *MockMMIMonitorTag_GetTagType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMMIMonitorTag_GetTagType_Call) RunAndReturn(run func() TagType) *MockMMIMonitorTag_GetTagType_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnitAddress provides a mock function with given fields:
func (_m *MockMMIMonitorTag) GetUnitAddress() model.UnitAddress {
	ret := _m.Called()

	var r0 model.UnitAddress
	if rf, ok := ret.Get(0).(func() model.UnitAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.UnitAddress)
		}
	}

	return r0
}

// MockMMIMonitorTag_GetUnitAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnitAddress'
type MockMMIMonitorTag_GetUnitAddress_Call struct {
	*mock.Call
}

// GetUnitAddress is a helper method to define mock.On call
func (_e *MockMMIMonitorTag_Expecter) GetUnitAddress() *MockMMIMonitorTag_GetUnitAddress_Call {
	return &MockMMIMonitorTag_GetUnitAddress_Call{Call: _e.mock.On("GetUnitAddress")}
}

func (_c *MockMMIMonitorTag_GetUnitAddress_Call) Run(run func()) *MockMMIMonitorTag_GetUnitAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMMIMonitorTag_GetUnitAddress_Call) Return(_a0 model.UnitAddress) *MockMMIMonitorTag_GetUnitAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMMIMonitorTag_GetUnitAddress_Call) RunAndReturn(run func() model.UnitAddress) *MockMMIMonitorTag_GetUnitAddress_Call {
	_c.Call.Return(run)
	return _c
}

// GetValueType provides a mock function with given fields:
func (_m *MockMMIMonitorTag) GetValueType() values.PlcValueType {
	ret := _m.Called()

	var r0 values.PlcValueType
	if rf, ok := ret.Get(0).(func() values.PlcValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(values.PlcValueType)
	}

	return r0
}

// MockMMIMonitorTag_GetValueType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValueType'
type MockMMIMonitorTag_GetValueType_Call struct {
	*mock.Call
}

// GetValueType is a helper method to define mock.On call
func (_e *MockMMIMonitorTag_Expecter) GetValueType() *MockMMIMonitorTag_GetValueType_Call {
	return &MockMMIMonitorTag_GetValueType_Call{Call: _e.mock.On("GetValueType")}
}

func (_c *MockMMIMonitorTag_GetValueType_Call) Run(run func()) *MockMMIMonitorTag_GetValueType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMMIMonitorTag_GetValueType_Call) Return(_a0 values.PlcValueType) *MockMMIMonitorTag_GetValueType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMMIMonitorTag_GetValueType_Call) RunAndReturn(run func() values.PlcValueType) *MockMMIMonitorTag_GetValueType_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockMMIMonitorTag) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockMMIMonitorTag_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockMMIMonitorTag_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockMMIMonitorTag_Expecter) String() *MockMMIMonitorTag_String_Call {
	return &MockMMIMonitorTag_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockMMIMonitorTag_String_Call) Run(run func()) *MockMMIMonitorTag_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMMIMonitorTag_String_Call) Return(_a0 string) *MockMMIMonitorTag_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMMIMonitorTag_String_Call) RunAndReturn(run func() string) *MockMMIMonitorTag_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMMIMonitorTag creates a new instance of MockMMIMonitorTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMMIMonitorTag(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMMIMonitorTag {
	mock := &MockMMIMonitorTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
