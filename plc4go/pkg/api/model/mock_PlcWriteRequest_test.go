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

package model

import (
	context "context"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcWriteRequest is an autogenerated mock type for the PlcWriteRequest type
type MockPlcWriteRequest struct {
	mock.Mock
}

type MockPlcWriteRequest_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcWriteRequest) EXPECT() *MockPlcWriteRequest_Expecter {
	return &MockPlcWriteRequest_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields:
func (_m *MockPlcWriteRequest) Execute() <-chan PlcWriteRequestResult {
	ret := _m.Called()

	var r0 <-chan PlcWriteRequestResult
	if rf, ok := ret.Get(0).(func() <-chan PlcWriteRequestResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcWriteRequestResult)
		}
	}

	return r0
}

// MockPlcWriteRequest_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockPlcWriteRequest_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockPlcWriteRequest_Expecter) Execute() *MockPlcWriteRequest_Execute_Call {
	return &MockPlcWriteRequest_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockPlcWriteRequest_Execute_Call) Run(run func()) *MockPlcWriteRequest_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequest_Execute_Call) Return(_a0 <-chan PlcWriteRequestResult) *MockPlcWriteRequest_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_Execute_Call) RunAndReturn(run func() <-chan PlcWriteRequestResult) *MockPlcWriteRequest_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteWithContext provides a mock function with given fields: ctx
func (_m *MockPlcWriteRequest) ExecuteWithContext(ctx context.Context) <-chan PlcWriteRequestResult {
	ret := _m.Called(ctx)

	var r0 <-chan PlcWriteRequestResult
	if rf, ok := ret.Get(0).(func(context.Context) <-chan PlcWriteRequestResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcWriteRequestResult)
		}
	}

	return r0
}

// MockPlcWriteRequest_ExecuteWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteWithContext'
type MockPlcWriteRequest_ExecuteWithContext_Call struct {
	*mock.Call
}

// ExecuteWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPlcWriteRequest_Expecter) ExecuteWithContext(ctx interface{}) *MockPlcWriteRequest_ExecuteWithContext_Call {
	return &MockPlcWriteRequest_ExecuteWithContext_Call{Call: _e.mock.On("ExecuteWithContext", ctx)}
}

func (_c *MockPlcWriteRequest_ExecuteWithContext_Call) Run(run func(ctx context.Context)) *MockPlcWriteRequest_ExecuteWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPlcWriteRequest_ExecuteWithContext_Call) Return(_a0 <-chan PlcWriteRequestResult) *MockPlcWriteRequest_ExecuteWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_ExecuteWithContext_Call) RunAndReturn(run func(context.Context) <-chan PlcWriteRequestResult) *MockPlcWriteRequest_ExecuteWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetTag provides a mock function with given fields: tagName
func (_m *MockPlcWriteRequest) GetTag(tagName string) PlcTag {
	ret := _m.Called(tagName)

	var r0 PlcTag
	if rf, ok := ret.Get(0).(func(string) PlcTag); ok {
		r0 = rf(tagName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcTag)
		}
	}

	return r0
}

// MockPlcWriteRequest_GetTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTag'
type MockPlcWriteRequest_GetTag_Call struct {
	*mock.Call
}

// GetTag is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcWriteRequest_Expecter) GetTag(tagName interface{}) *MockPlcWriteRequest_GetTag_Call {
	return &MockPlcWriteRequest_GetTag_Call{Call: _e.mock.On("GetTag", tagName)}
}

func (_c *MockPlcWriteRequest_GetTag_Call) Run(run func(tagName string)) *MockPlcWriteRequest_GetTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcWriteRequest_GetTag_Call) Return(_a0 PlcTag) *MockPlcWriteRequest_GetTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_GetTag_Call) RunAndReturn(run func(string) PlcTag) *MockPlcWriteRequest_GetTag_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagNames provides a mock function with given fields:
func (_m *MockPlcWriteRequest) GetTagNames() []string {
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

// MockPlcWriteRequest_GetTagNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagNames'
type MockPlcWriteRequest_GetTagNames_Call struct {
	*mock.Call
}

// GetTagNames is a helper method to define mock.On call
func (_e *MockPlcWriteRequest_Expecter) GetTagNames() *MockPlcWriteRequest_GetTagNames_Call {
	return &MockPlcWriteRequest_GetTagNames_Call{Call: _e.mock.On("GetTagNames")}
}

func (_c *MockPlcWriteRequest_GetTagNames_Call) Run(run func()) *MockPlcWriteRequest_GetTagNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequest_GetTagNames_Call) Return(_a0 []string) *MockPlcWriteRequest_GetTagNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_GetTagNames_Call) RunAndReturn(run func() []string) *MockPlcWriteRequest_GetTagNames_Call {
	_c.Call.Return(run)
	return _c
}

// GetValue provides a mock function with given fields: tagName
func (_m *MockPlcWriteRequest) GetValue(tagName string) values.PlcValue {
	ret := _m.Called(tagName)

	var r0 values.PlcValue
	if rf, ok := ret.Get(0).(func(string) values.PlcValue); ok {
		r0 = rf(tagName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(values.PlcValue)
		}
	}

	return r0
}

// MockPlcWriteRequest_GetValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValue'
type MockPlcWriteRequest_GetValue_Call struct {
	*mock.Call
}

// GetValue is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcWriteRequest_Expecter) GetValue(tagName interface{}) *MockPlcWriteRequest_GetValue_Call {
	return &MockPlcWriteRequest_GetValue_Call{Call: _e.mock.On("GetValue", tagName)}
}

func (_c *MockPlcWriteRequest_GetValue_Call) Run(run func(tagName string)) *MockPlcWriteRequest_GetValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcWriteRequest_GetValue_Call) Return(_a0 values.PlcValue) *MockPlcWriteRequest_GetValue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_GetValue_Call) RunAndReturn(run func(string) values.PlcValue) *MockPlcWriteRequest_GetValue_Call {
	_c.Call.Return(run)
	return _c
}

// IsAPlcMessage provides a mock function with given fields:
func (_m *MockPlcWriteRequest) IsAPlcMessage() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcWriteRequest_IsAPlcMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAPlcMessage'
type MockPlcWriteRequest_IsAPlcMessage_Call struct {
	*mock.Call
}

// IsAPlcMessage is a helper method to define mock.On call
func (_e *MockPlcWriteRequest_Expecter) IsAPlcMessage() *MockPlcWriteRequest_IsAPlcMessage_Call {
	return &MockPlcWriteRequest_IsAPlcMessage_Call{Call: _e.mock.On("IsAPlcMessage")}
}

func (_c *MockPlcWriteRequest_IsAPlcMessage_Call) Run(run func()) *MockPlcWriteRequest_IsAPlcMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequest_IsAPlcMessage_Call) Return(_a0 bool) *MockPlcWriteRequest_IsAPlcMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_IsAPlcMessage_Call) RunAndReturn(run func() bool) *MockPlcWriteRequest_IsAPlcMessage_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcWriteRequest) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcWriteRequest_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcWriteRequest_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcWriteRequest_Expecter) String() *MockPlcWriteRequest_String_Call {
	return &MockPlcWriteRequest_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcWriteRequest_String_Call) Run(run func()) *MockPlcWriteRequest_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequest_String_Call) Return(_a0 string) *MockPlcWriteRequest_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_String_Call) RunAndReturn(run func() string) *MockPlcWriteRequest_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcWriteRequest creates a new instance of MockPlcWriteRequest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcWriteRequest(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcWriteRequest {
	mock := &MockPlcWriteRequest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
