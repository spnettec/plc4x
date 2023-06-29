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

// Code generated by mockery v2.30.16. DO NOT EDIT.

package model

import (
	values "github.com/apache/plc4x/plc4go/pkg/api/values"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcBrowseItem is an autogenerated mock type for the PlcBrowseItem type
type MockPlcBrowseItem struct {
	mock.Mock
}

type MockPlcBrowseItem_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcBrowseItem) EXPECT() *MockPlcBrowseItem_Expecter {
	return &MockPlcBrowseItem_Expecter{mock: &_m.Mock}
}

// GetChildren provides a mock function with given fields:
func (_m *MockPlcBrowseItem) GetChildren() map[string]PlcBrowseItem {
	ret := _m.Called()

	var r0 map[string]PlcBrowseItem
	if rf, ok := ret.Get(0).(func() map[string]PlcBrowseItem); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]PlcBrowseItem)
		}
	}

	return r0
}

// MockPlcBrowseItem_GetChildren_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChildren'
type MockPlcBrowseItem_GetChildren_Call struct {
	*mock.Call
}

// GetChildren is a helper method to define mock.On call
func (_e *MockPlcBrowseItem_Expecter) GetChildren() *MockPlcBrowseItem_GetChildren_Call {
	return &MockPlcBrowseItem_GetChildren_Call{Call: _e.mock.On("GetChildren")}
}

func (_c *MockPlcBrowseItem_GetChildren_Call) Run(run func()) *MockPlcBrowseItem_GetChildren_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseItem_GetChildren_Call) Return(_a0 map[string]PlcBrowseItem) *MockPlcBrowseItem_GetChildren_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseItem_GetChildren_Call) RunAndReturn(run func() map[string]PlcBrowseItem) *MockPlcBrowseItem_GetChildren_Call {
	_c.Call.Return(run)
	return _c
}

// GetName provides a mock function with given fields:
func (_m *MockPlcBrowseItem) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcBrowseItem_GetName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetName'
type MockPlcBrowseItem_GetName_Call struct {
	*mock.Call
}

// GetName is a helper method to define mock.On call
func (_e *MockPlcBrowseItem_Expecter) GetName() *MockPlcBrowseItem_GetName_Call {
	return &MockPlcBrowseItem_GetName_Call{Call: _e.mock.On("GetName")}
}

func (_c *MockPlcBrowseItem_GetName_Call) Run(run func()) *MockPlcBrowseItem_GetName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseItem_GetName_Call) Return(_a0 string) *MockPlcBrowseItem_GetName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseItem_GetName_Call) RunAndReturn(run func() string) *MockPlcBrowseItem_GetName_Call {
	_c.Call.Return(run)
	return _c
}

// GetOptions provides a mock function with given fields:
func (_m *MockPlcBrowseItem) GetOptions() map[string]values.PlcValue {
	ret := _m.Called()

	var r0 map[string]values.PlcValue
	if rf, ok := ret.Get(0).(func() map[string]values.PlcValue); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]values.PlcValue)
		}
	}

	return r0
}

// MockPlcBrowseItem_GetOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOptions'
type MockPlcBrowseItem_GetOptions_Call struct {
	*mock.Call
}

// GetOptions is a helper method to define mock.On call
func (_e *MockPlcBrowseItem_Expecter) GetOptions() *MockPlcBrowseItem_GetOptions_Call {
	return &MockPlcBrowseItem_GetOptions_Call{Call: _e.mock.On("GetOptions")}
}

func (_c *MockPlcBrowseItem_GetOptions_Call) Run(run func()) *MockPlcBrowseItem_GetOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseItem_GetOptions_Call) Return(_a0 map[string]values.PlcValue) *MockPlcBrowseItem_GetOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseItem_GetOptions_Call) RunAndReturn(run func() map[string]values.PlcValue) *MockPlcBrowseItem_GetOptions_Call {
	_c.Call.Return(run)
	return _c
}

// GetTag provides a mock function with given fields:
func (_m *MockPlcBrowseItem) GetTag() PlcTag {
	ret := _m.Called()

	var r0 PlcTag
	if rf, ok := ret.Get(0).(func() PlcTag); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcTag)
		}
	}

	return r0
}

// MockPlcBrowseItem_GetTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTag'
type MockPlcBrowseItem_GetTag_Call struct {
	*mock.Call
}

// GetTag is a helper method to define mock.On call
func (_e *MockPlcBrowseItem_Expecter) GetTag() *MockPlcBrowseItem_GetTag_Call {
	return &MockPlcBrowseItem_GetTag_Call{Call: _e.mock.On("GetTag")}
}

func (_c *MockPlcBrowseItem_GetTag_Call) Run(run func()) *MockPlcBrowseItem_GetTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseItem_GetTag_Call) Return(_a0 PlcTag) *MockPlcBrowseItem_GetTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseItem_GetTag_Call) RunAndReturn(run func() PlcTag) *MockPlcBrowseItem_GetTag_Call {
	_c.Call.Return(run)
	return _c
}

// IsReadable provides a mock function with given fields:
func (_m *MockPlcBrowseItem) IsReadable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcBrowseItem_IsReadable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsReadable'
type MockPlcBrowseItem_IsReadable_Call struct {
	*mock.Call
}

// IsReadable is a helper method to define mock.On call
func (_e *MockPlcBrowseItem_Expecter) IsReadable() *MockPlcBrowseItem_IsReadable_Call {
	return &MockPlcBrowseItem_IsReadable_Call{Call: _e.mock.On("IsReadable")}
}

func (_c *MockPlcBrowseItem_IsReadable_Call) Run(run func()) *MockPlcBrowseItem_IsReadable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseItem_IsReadable_Call) Return(_a0 bool) *MockPlcBrowseItem_IsReadable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseItem_IsReadable_Call) RunAndReturn(run func() bool) *MockPlcBrowseItem_IsReadable_Call {
	_c.Call.Return(run)
	return _c
}

// IsSubscribable provides a mock function with given fields:
func (_m *MockPlcBrowseItem) IsSubscribable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcBrowseItem_IsSubscribable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSubscribable'
type MockPlcBrowseItem_IsSubscribable_Call struct {
	*mock.Call
}

// IsSubscribable is a helper method to define mock.On call
func (_e *MockPlcBrowseItem_Expecter) IsSubscribable() *MockPlcBrowseItem_IsSubscribable_Call {
	return &MockPlcBrowseItem_IsSubscribable_Call{Call: _e.mock.On("IsSubscribable")}
}

func (_c *MockPlcBrowseItem_IsSubscribable_Call) Run(run func()) *MockPlcBrowseItem_IsSubscribable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseItem_IsSubscribable_Call) Return(_a0 bool) *MockPlcBrowseItem_IsSubscribable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseItem_IsSubscribable_Call) RunAndReturn(run func() bool) *MockPlcBrowseItem_IsSubscribable_Call {
	_c.Call.Return(run)
	return _c
}

// IsWritable provides a mock function with given fields:
func (_m *MockPlcBrowseItem) IsWritable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcBrowseItem_IsWritable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsWritable'
type MockPlcBrowseItem_IsWritable_Call struct {
	*mock.Call
}

// IsWritable is a helper method to define mock.On call
func (_e *MockPlcBrowseItem_Expecter) IsWritable() *MockPlcBrowseItem_IsWritable_Call {
	return &MockPlcBrowseItem_IsWritable_Call{Call: _e.mock.On("IsWritable")}
}

func (_c *MockPlcBrowseItem_IsWritable_Call) Run(run func()) *MockPlcBrowseItem_IsWritable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseItem_IsWritable_Call) Return(_a0 bool) *MockPlcBrowseItem_IsWritable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseItem_IsWritable_Call) RunAndReturn(run func() bool) *MockPlcBrowseItem_IsWritable_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcBrowseItem creates a new instance of MockPlcBrowseItem. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcBrowseItem(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcBrowseItem {
	mock := &MockPlcBrowseItem{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
