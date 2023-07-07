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

package model

import mock "github.com/stretchr/testify/mock"

// MockPlcSubscriptionResponse is an autogenerated mock type for the PlcSubscriptionResponse type
type MockPlcSubscriptionResponse struct {
	mock.Mock
}

type MockPlcSubscriptionResponse_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcSubscriptionResponse) EXPECT() *MockPlcSubscriptionResponse_Expecter {
	return &MockPlcSubscriptionResponse_Expecter{mock: &_m.Mock}
}

// GetRequest provides a mock function with given fields:
func (_m *MockPlcSubscriptionResponse) GetRequest() PlcSubscriptionRequest {
	ret := _m.Called()

	var r0 PlcSubscriptionRequest
	if rf, ok := ret.Get(0).(func() PlcSubscriptionRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcSubscriptionRequest)
		}
	}

	return r0
}

// MockPlcSubscriptionResponse_GetRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequest'
type MockPlcSubscriptionResponse_GetRequest_Call struct {
	*mock.Call
}

// GetRequest is a helper method to define mock.On call
func (_e *MockPlcSubscriptionResponse_Expecter) GetRequest() *MockPlcSubscriptionResponse_GetRequest_Call {
	return &MockPlcSubscriptionResponse_GetRequest_Call{Call: _e.mock.On("GetRequest")}
}

func (_c *MockPlcSubscriptionResponse_GetRequest_Call) Run(run func()) *MockPlcSubscriptionResponse_GetRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionResponse_GetRequest_Call) Return(_a0 PlcSubscriptionRequest) *MockPlcSubscriptionResponse_GetRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionResponse_GetRequest_Call) RunAndReturn(run func() PlcSubscriptionRequest) *MockPlcSubscriptionResponse_GetRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetResponseCode provides a mock function with given fields: name
func (_m *MockPlcSubscriptionResponse) GetResponseCode(name string) PlcResponseCode {
	ret := _m.Called(name)

	var r0 PlcResponseCode
	if rf, ok := ret.Get(0).(func(string) PlcResponseCode); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(PlcResponseCode)
	}

	return r0
}

// MockPlcSubscriptionResponse_GetResponseCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResponseCode'
type MockPlcSubscriptionResponse_GetResponseCode_Call struct {
	*mock.Call
}

// GetResponseCode is a helper method to define mock.On call
//   - name string
func (_e *MockPlcSubscriptionResponse_Expecter) GetResponseCode(name interface{}) *MockPlcSubscriptionResponse_GetResponseCode_Call {
	return &MockPlcSubscriptionResponse_GetResponseCode_Call{Call: _e.mock.On("GetResponseCode", name)}
}

func (_c *MockPlcSubscriptionResponse_GetResponseCode_Call) Run(run func(name string)) *MockPlcSubscriptionResponse_GetResponseCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcSubscriptionResponse_GetResponseCode_Call) Return(_a0 PlcResponseCode) *MockPlcSubscriptionResponse_GetResponseCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionResponse_GetResponseCode_Call) RunAndReturn(run func(string) PlcResponseCode) *MockPlcSubscriptionResponse_GetResponseCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubscriptionHandle provides a mock function with given fields: name
func (_m *MockPlcSubscriptionResponse) GetSubscriptionHandle(name string) (PlcSubscriptionHandle, error) {
	ret := _m.Called(name)

	var r0 PlcSubscriptionHandle
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (PlcSubscriptionHandle, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) PlcSubscriptionHandle); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcSubscriptionHandle)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlcSubscriptionResponse_GetSubscriptionHandle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptionHandle'
type MockPlcSubscriptionResponse_GetSubscriptionHandle_Call struct {
	*mock.Call
}

// GetSubscriptionHandle is a helper method to define mock.On call
//   - name string
func (_e *MockPlcSubscriptionResponse_Expecter) GetSubscriptionHandle(name interface{}) *MockPlcSubscriptionResponse_GetSubscriptionHandle_Call {
	return &MockPlcSubscriptionResponse_GetSubscriptionHandle_Call{Call: _e.mock.On("GetSubscriptionHandle", name)}
}

func (_c *MockPlcSubscriptionResponse_GetSubscriptionHandle_Call) Run(run func(name string)) *MockPlcSubscriptionResponse_GetSubscriptionHandle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcSubscriptionResponse_GetSubscriptionHandle_Call) Return(_a0 PlcSubscriptionHandle, _a1 error) *MockPlcSubscriptionResponse_GetSubscriptionHandle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlcSubscriptionResponse_GetSubscriptionHandle_Call) RunAndReturn(run func(string) (PlcSubscriptionHandle, error)) *MockPlcSubscriptionResponse_GetSubscriptionHandle_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubscriptionHandles provides a mock function with given fields:
func (_m *MockPlcSubscriptionResponse) GetSubscriptionHandles() []PlcSubscriptionHandle {
	ret := _m.Called()

	var r0 []PlcSubscriptionHandle
	if rf, ok := ret.Get(0).(func() []PlcSubscriptionHandle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]PlcSubscriptionHandle)
		}
	}

	return r0
}

// MockPlcSubscriptionResponse_GetSubscriptionHandles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptionHandles'
type MockPlcSubscriptionResponse_GetSubscriptionHandles_Call struct {
	*mock.Call
}

// GetSubscriptionHandles is a helper method to define mock.On call
func (_e *MockPlcSubscriptionResponse_Expecter) GetSubscriptionHandles() *MockPlcSubscriptionResponse_GetSubscriptionHandles_Call {
	return &MockPlcSubscriptionResponse_GetSubscriptionHandles_Call{Call: _e.mock.On("GetSubscriptionHandles")}
}

func (_c *MockPlcSubscriptionResponse_GetSubscriptionHandles_Call) Run(run func()) *MockPlcSubscriptionResponse_GetSubscriptionHandles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionResponse_GetSubscriptionHandles_Call) Return(_a0 []PlcSubscriptionHandle) *MockPlcSubscriptionResponse_GetSubscriptionHandles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionResponse_GetSubscriptionHandles_Call) RunAndReturn(run func() []PlcSubscriptionHandle) *MockPlcSubscriptionResponse_GetSubscriptionHandles_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagNames provides a mock function with given fields:
func (_m *MockPlcSubscriptionResponse) GetTagNames() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockPlcSubscriptionResponse_GetTagNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagNames'
type MockPlcSubscriptionResponse_GetTagNames_Call struct {
	*mock.Call
}

// GetTagNames is a helper method to define mock.On call
func (_e *MockPlcSubscriptionResponse_Expecter) GetTagNames() *MockPlcSubscriptionResponse_GetTagNames_Call {
	return &MockPlcSubscriptionResponse_GetTagNames_Call{Call: _e.mock.On("GetTagNames")}
}

func (_c *MockPlcSubscriptionResponse_GetTagNames_Call) Run(run func()) *MockPlcSubscriptionResponse_GetTagNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionResponse_GetTagNames_Call) Return(_a0 []string) *MockPlcSubscriptionResponse_GetTagNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionResponse_GetTagNames_Call) RunAndReturn(run func() []string) *MockPlcSubscriptionResponse_GetTagNames_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcSubscriptionResponse creates a new instance of MockPlcSubscriptionResponse. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcSubscriptionResponse(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcSubscriptionResponse {
	mock := &MockPlcSubscriptionResponse{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
