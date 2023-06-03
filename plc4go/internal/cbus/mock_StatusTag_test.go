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

// Code generated by mockery v2.28.2. DO NOT EDIT.

package cbus

import (
	apimodel "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	model "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockStatusTag is an autogenerated mock type for the StatusTag type
type MockStatusTag struct {
	mock.Mock
}

type MockStatusTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStatusTag) EXPECT() *MockStatusTag_Expecter {
	return &MockStatusTag_Expecter{mock: &_m.Mock}
}

// GetAddressString provides a mock function with given fields:
func (_m *MockStatusTag) GetAddressString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockStatusTag_GetAddressString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddressString'
type MockStatusTag_GetAddressString_Call struct {
	*mock.Call
}

// GetAddressString is a helper method to define mock.On call
func (_e *MockStatusTag_Expecter) GetAddressString() *MockStatusTag_GetAddressString_Call {
	return &MockStatusTag_GetAddressString_Call{Call: _e.mock.On("GetAddressString")}
}

func (_c *MockStatusTag_GetAddressString_Call) Run(run func()) *MockStatusTag_GetAddressString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStatusTag_GetAddressString_Call) Return(_a0 string) *MockStatusTag_GetAddressString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStatusTag_GetAddressString_Call) RunAndReturn(run func() string) *MockStatusTag_GetAddressString_Call {
	_c.Call.Return(run)
	return _c
}

// GetApplication provides a mock function with given fields:
func (_m *MockStatusTag) GetApplication() model.ApplicationIdContainer {
	ret := _m.Called()

	var r0 model.ApplicationIdContainer
	if rf, ok := ret.Get(0).(func() model.ApplicationIdContainer); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.ApplicationIdContainer)
	}

	return r0
}

// MockStatusTag_GetApplication_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetApplication'
type MockStatusTag_GetApplication_Call struct {
	*mock.Call
}

// GetApplication is a helper method to define mock.On call
func (_e *MockStatusTag_Expecter) GetApplication() *MockStatusTag_GetApplication_Call {
	return &MockStatusTag_GetApplication_Call{Call: _e.mock.On("GetApplication")}
}

func (_c *MockStatusTag_GetApplication_Call) Run(run func()) *MockStatusTag_GetApplication_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStatusTag_GetApplication_Call) Return(_a0 model.ApplicationIdContainer) *MockStatusTag_GetApplication_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStatusTag_GetApplication_Call) RunAndReturn(run func() model.ApplicationIdContainer) *MockStatusTag_GetApplication_Call {
	_c.Call.Return(run)
	return _c
}

// GetArrayInfo provides a mock function with given fields:
func (_m *MockStatusTag) GetArrayInfo() []apimodel.ArrayInfo {
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

// MockStatusTag_GetArrayInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArrayInfo'
type MockStatusTag_GetArrayInfo_Call struct {
	*mock.Call
}

// GetArrayInfo is a helper method to define mock.On call
func (_e *MockStatusTag_Expecter) GetArrayInfo() *MockStatusTag_GetArrayInfo_Call {
	return &MockStatusTag_GetArrayInfo_Call{Call: _e.mock.On("GetArrayInfo")}
}

func (_c *MockStatusTag_GetArrayInfo_Call) Run(run func()) *MockStatusTag_GetArrayInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStatusTag_GetArrayInfo_Call) Return(_a0 []apimodel.ArrayInfo) *MockStatusTag_GetArrayInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStatusTag_GetArrayInfo_Call) RunAndReturn(run func() []apimodel.ArrayInfo) *MockStatusTag_GetArrayInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetBridgeAddresses provides a mock function with given fields:
func (_m *MockStatusTag) GetBridgeAddresses() []model.BridgeAddress {
	ret := _m.Called()

	var r0 []model.BridgeAddress
	if rf, ok := ret.Get(0).(func() []model.BridgeAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.BridgeAddress)
		}
	}

	return r0
}

// MockStatusTag_GetBridgeAddresses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBridgeAddresses'
type MockStatusTag_GetBridgeAddresses_Call struct {
	*mock.Call
}

// GetBridgeAddresses is a helper method to define mock.On call
func (_e *MockStatusTag_Expecter) GetBridgeAddresses() *MockStatusTag_GetBridgeAddresses_Call {
	return &MockStatusTag_GetBridgeAddresses_Call{Call: _e.mock.On("GetBridgeAddresses")}
}

func (_c *MockStatusTag_GetBridgeAddresses_Call) Run(run func()) *MockStatusTag_GetBridgeAddresses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStatusTag_GetBridgeAddresses_Call) Return(_a0 []model.BridgeAddress) *MockStatusTag_GetBridgeAddresses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStatusTag_GetBridgeAddresses_Call) RunAndReturn(run func() []model.BridgeAddress) *MockStatusTag_GetBridgeAddresses_Call {
	_c.Call.Return(run)
	return _c
}

// GetStartingGroupAddressLabel provides a mock function with given fields:
func (_m *MockStatusTag) GetStartingGroupAddressLabel() *byte {
	ret := _m.Called()

	var r0 *byte
	if rf, ok := ret.Get(0).(func() *byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*byte)
		}
	}

	return r0
}

// MockStatusTag_GetStartingGroupAddressLabel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStartingGroupAddressLabel'
type MockStatusTag_GetStartingGroupAddressLabel_Call struct {
	*mock.Call
}

// GetStartingGroupAddressLabel is a helper method to define mock.On call
func (_e *MockStatusTag_Expecter) GetStartingGroupAddressLabel() *MockStatusTag_GetStartingGroupAddressLabel_Call {
	return &MockStatusTag_GetStartingGroupAddressLabel_Call{Call: _e.mock.On("GetStartingGroupAddressLabel")}
}

func (_c *MockStatusTag_GetStartingGroupAddressLabel_Call) Run(run func()) *MockStatusTag_GetStartingGroupAddressLabel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStatusTag_GetStartingGroupAddressLabel_Call) Return(_a0 *byte) *MockStatusTag_GetStartingGroupAddressLabel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStatusTag_GetStartingGroupAddressLabel_Call) RunAndReturn(run func() *byte) *MockStatusTag_GetStartingGroupAddressLabel_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatusRequestType provides a mock function with given fields:
func (_m *MockStatusTag) GetStatusRequestType() StatusRequestType {
	ret := _m.Called()

	var r0 StatusRequestType
	if rf, ok := ret.Get(0).(func() StatusRequestType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(StatusRequestType)
	}

	return r0
}

// MockStatusTag_GetStatusRequestType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatusRequestType'
type MockStatusTag_GetStatusRequestType_Call struct {
	*mock.Call
}

// GetStatusRequestType is a helper method to define mock.On call
func (_e *MockStatusTag_Expecter) GetStatusRequestType() *MockStatusTag_GetStatusRequestType_Call {
	return &MockStatusTag_GetStatusRequestType_Call{Call: _e.mock.On("GetStatusRequestType")}
}

func (_c *MockStatusTag_GetStatusRequestType_Call) Run(run func()) *MockStatusTag_GetStatusRequestType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStatusTag_GetStatusRequestType_Call) Return(_a0 StatusRequestType) *MockStatusTag_GetStatusRequestType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStatusTag_GetStatusRequestType_Call) RunAndReturn(run func() StatusRequestType) *MockStatusTag_GetStatusRequestType_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagType provides a mock function with given fields:
func (_m *MockStatusTag) GetTagType() TagType {
	ret := _m.Called()

	var r0 TagType
	if rf, ok := ret.Get(0).(func() TagType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(TagType)
	}

	return r0
}

// MockStatusTag_GetTagType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagType'
type MockStatusTag_GetTagType_Call struct {
	*mock.Call
}

// GetTagType is a helper method to define mock.On call
func (_e *MockStatusTag_Expecter) GetTagType() *MockStatusTag_GetTagType_Call {
	return &MockStatusTag_GetTagType_Call{Call: _e.mock.On("GetTagType")}
}

func (_c *MockStatusTag_GetTagType_Call) Run(run func()) *MockStatusTag_GetTagType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStatusTag_GetTagType_Call) Return(_a0 TagType) *MockStatusTag_GetTagType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStatusTag_GetTagType_Call) RunAndReturn(run func() TagType) *MockStatusTag_GetTagType_Call {
	_c.Call.Return(run)
	return _c
}

// GetValueType provides a mock function with given fields:
func (_m *MockStatusTag) GetValueType() values.PlcValueType {
	ret := _m.Called()

	var r0 values.PlcValueType
	if rf, ok := ret.Get(0).(func() values.PlcValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(values.PlcValueType)
	}

	return r0
}

// MockStatusTag_GetValueType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValueType'
type MockStatusTag_GetValueType_Call struct {
	*mock.Call
}

// GetValueType is a helper method to define mock.On call
func (_e *MockStatusTag_Expecter) GetValueType() *MockStatusTag_GetValueType_Call {
	return &MockStatusTag_GetValueType_Call{Call: _e.mock.On("GetValueType")}
}

func (_c *MockStatusTag_GetValueType_Call) Run(run func()) *MockStatusTag_GetValueType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStatusTag_GetValueType_Call) Return(_a0 values.PlcValueType) *MockStatusTag_GetValueType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStatusTag_GetValueType_Call) RunAndReturn(run func() values.PlcValueType) *MockStatusTag_GetValueType_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockStatusTag interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockStatusTag creates a new instance of MockStatusTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockStatusTag(t mockConstructorTestingTNewMockStatusTag) *MockStatusTag {
	mock := &MockStatusTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
