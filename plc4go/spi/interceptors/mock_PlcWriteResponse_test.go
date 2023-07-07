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

package interceptors

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcWriteResponse is an autogenerated mock type for the PlcWriteResponse type
type MockPlcWriteResponse struct {
	mock.Mock
}

type MockPlcWriteResponse_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcWriteResponse) EXPECT() *MockPlcWriteResponse_Expecter {
	return &MockPlcWriteResponse_Expecter{mock: &_m.Mock}
}

// GetRequest provides a mock function with given fields:
func (_m *MockPlcWriteResponse) GetRequest() model.PlcWriteRequest {
	ret := _m.Called()

	var r0 model.PlcWriteRequest
	if rf, ok := ret.Get(0).(func() model.PlcWriteRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcWriteRequest)
		}
	}

	return r0
}

// MockPlcWriteResponse_GetRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequest'
type MockPlcWriteResponse_GetRequest_Call struct {
	*mock.Call
}

// GetRequest is a helper method to define mock.On call
func (_e *MockPlcWriteResponse_Expecter) GetRequest() *MockPlcWriteResponse_GetRequest_Call {
	return &MockPlcWriteResponse_GetRequest_Call{Call: _e.mock.On("GetRequest")}
}

func (_c *MockPlcWriteResponse_GetRequest_Call) Run(run func()) *MockPlcWriteResponse_GetRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteResponse_GetRequest_Call) Return(_a0 model.PlcWriteRequest) *MockPlcWriteResponse_GetRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteResponse_GetRequest_Call) RunAndReturn(run func() model.PlcWriteRequest) *MockPlcWriteResponse_GetRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetResponseCode provides a mock function with given fields: tagName
func (_m *MockPlcWriteResponse) GetResponseCode(tagName string) model.PlcResponseCode {
	ret := _m.Called(tagName)

	var r0 model.PlcResponseCode
	if rf, ok := ret.Get(0).(func(string) model.PlcResponseCode); ok {
		r0 = rf(tagName)
	} else {
		r0 = ret.Get(0).(model.PlcResponseCode)
	}

	return r0
}

// MockPlcWriteResponse_GetResponseCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResponseCode'
type MockPlcWriteResponse_GetResponseCode_Call struct {
	*mock.Call
}

// GetResponseCode is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcWriteResponse_Expecter) GetResponseCode(tagName interface{}) *MockPlcWriteResponse_GetResponseCode_Call {
	return &MockPlcWriteResponse_GetResponseCode_Call{Call: _e.mock.On("GetResponseCode", tagName)}
}

func (_c *MockPlcWriteResponse_GetResponseCode_Call) Run(run func(tagName string)) *MockPlcWriteResponse_GetResponseCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcWriteResponse_GetResponseCode_Call) Return(_a0 model.PlcResponseCode) *MockPlcWriteResponse_GetResponseCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteResponse_GetResponseCode_Call) RunAndReturn(run func(string) model.PlcResponseCode) *MockPlcWriteResponse_GetResponseCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagNames provides a mock function with given fields:
func (_m *MockPlcWriteResponse) GetTagNames() []string {
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

// MockPlcWriteResponse_GetTagNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagNames'
type MockPlcWriteResponse_GetTagNames_Call struct {
	*mock.Call
}

// GetTagNames is a helper method to define mock.On call
func (_e *MockPlcWriteResponse_Expecter) GetTagNames() *MockPlcWriteResponse_GetTagNames_Call {
	return &MockPlcWriteResponse_GetTagNames_Call{Call: _e.mock.On("GetTagNames")}
}

func (_c *MockPlcWriteResponse_GetTagNames_Call) Run(run func()) *MockPlcWriteResponse_GetTagNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteResponse_GetTagNames_Call) Return(_a0 []string) *MockPlcWriteResponse_GetTagNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteResponse_GetTagNames_Call) RunAndReturn(run func() []string) *MockPlcWriteResponse_GetTagNames_Call {
	_c.Call.Return(run)
	return _c
}

// IsAPlcMessage provides a mock function with given fields:
func (_m *MockPlcWriteResponse) IsAPlcMessage() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcWriteResponse_IsAPlcMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAPlcMessage'
type MockPlcWriteResponse_IsAPlcMessage_Call struct {
	*mock.Call
}

// IsAPlcMessage is a helper method to define mock.On call
func (_e *MockPlcWriteResponse_Expecter) IsAPlcMessage() *MockPlcWriteResponse_IsAPlcMessage_Call {
	return &MockPlcWriteResponse_IsAPlcMessage_Call{Call: _e.mock.On("IsAPlcMessage")}
}

func (_c *MockPlcWriteResponse_IsAPlcMessage_Call) Run(run func()) *MockPlcWriteResponse_IsAPlcMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteResponse_IsAPlcMessage_Call) Return(_a0 bool) *MockPlcWriteResponse_IsAPlcMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteResponse_IsAPlcMessage_Call) RunAndReturn(run func() bool) *MockPlcWriteResponse_IsAPlcMessage_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcWriteResponse) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcWriteResponse_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcWriteResponse_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcWriteResponse_Expecter) String() *MockPlcWriteResponse_String_Call {
	return &MockPlcWriteResponse_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcWriteResponse_String_Call) Run(run func()) *MockPlcWriteResponse_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteResponse_String_Call) Return(_a0 string) *MockPlcWriteResponse_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteResponse_String_Call) RunAndReturn(run func() string) *MockPlcWriteResponse_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcWriteResponse creates a new instance of MockPlcWriteResponse. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcWriteResponse(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcWriteResponse {
	mock := &MockPlcWriteResponse{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
