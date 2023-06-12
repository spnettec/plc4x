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

// Code generated by mockery v2.29.0. DO NOT EDIT.

package cbus

import (
	model "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
	mock "github.com/stretchr/testify/mock"
)

// MockCalTag is an autogenerated mock type for the CalTag type
type MockCalTag struct {
	mock.Mock
}

type MockCalTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCalTag) EXPECT() *MockCalTag_Expecter {
	return &MockCalTag_Expecter{mock: &_m.Mock}
}

// GetBridgeAddresses provides a mock function with given fields:
func (_m *MockCalTag) GetBridgeAddresses() []model.BridgeAddress {
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

// MockCalTag_GetBridgeAddresses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBridgeAddresses'
type MockCalTag_GetBridgeAddresses_Call struct {
	*mock.Call
}

// GetBridgeAddresses is a helper method to define mock.On call
func (_e *MockCalTag_Expecter) GetBridgeAddresses() *MockCalTag_GetBridgeAddresses_Call {
	return &MockCalTag_GetBridgeAddresses_Call{Call: _e.mock.On("GetBridgeAddresses")}
}

func (_c *MockCalTag_GetBridgeAddresses_Call) Run(run func()) *MockCalTag_GetBridgeAddresses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCalTag_GetBridgeAddresses_Call) Return(_a0 []model.BridgeAddress) *MockCalTag_GetBridgeAddresses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCalTag_GetBridgeAddresses_Call) RunAndReturn(run func() []model.BridgeAddress) *MockCalTag_GetBridgeAddresses_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnitAddress provides a mock function with given fields:
func (_m *MockCalTag) GetUnitAddress() model.UnitAddress {
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

// MockCalTag_GetUnitAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnitAddress'
type MockCalTag_GetUnitAddress_Call struct {
	*mock.Call
}

// GetUnitAddress is a helper method to define mock.On call
func (_e *MockCalTag_Expecter) GetUnitAddress() *MockCalTag_GetUnitAddress_Call {
	return &MockCalTag_GetUnitAddress_Call{Call: _e.mock.On("GetUnitAddress")}
}

func (_c *MockCalTag_GetUnitAddress_Call) Run(run func()) *MockCalTag_GetUnitAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCalTag_GetUnitAddress_Call) Return(_a0 model.UnitAddress) *MockCalTag_GetUnitAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCalTag_GetUnitAddress_Call) RunAndReturn(run func() model.UnitAddress) *MockCalTag_GetUnitAddress_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockCalTag interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCalTag creates a new instance of MockCalTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCalTag(t mockConstructorTestingTNewMockCalTag) *MockCalTag {
	mock := &MockCalTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
