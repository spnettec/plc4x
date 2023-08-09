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

package transactions

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRequestTransaction is an autogenerated mock type for the RequestTransaction type
type MockRequestTransaction struct {
	mock.Mock
}

type MockRequestTransaction_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRequestTransaction) EXPECT() *MockRequestTransaction_Expecter {
	return &MockRequestTransaction_Expecter{mock: &_m.Mock}
}

// AwaitCompletion provides a mock function with given fields: ctx
func (_m *MockRequestTransaction) AwaitCompletion(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRequestTransaction_AwaitCompletion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AwaitCompletion'
type MockRequestTransaction_AwaitCompletion_Call struct {
	*mock.Call
}

// AwaitCompletion is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRequestTransaction_Expecter) AwaitCompletion(ctx interface{}) *MockRequestTransaction_AwaitCompletion_Call {
	return &MockRequestTransaction_AwaitCompletion_Call{Call: _e.mock.On("AwaitCompletion", ctx)}
}

func (_c *MockRequestTransaction_AwaitCompletion_Call) Run(run func(ctx context.Context)) *MockRequestTransaction_AwaitCompletion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRequestTransaction_AwaitCompletion_Call) Return(_a0 error) *MockRequestTransaction_AwaitCompletion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestTransaction_AwaitCompletion_Call) RunAndReturn(run func(context.Context) error) *MockRequestTransaction_AwaitCompletion_Call {
	_c.Call.Return(run)
	return _c
}

// EndRequest provides a mock function with given fields:
func (_m *MockRequestTransaction) EndRequest() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRequestTransaction_EndRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EndRequest'
type MockRequestTransaction_EndRequest_Call struct {
	*mock.Call
}

// EndRequest is a helper method to define mock.On call
func (_e *MockRequestTransaction_Expecter) EndRequest() *MockRequestTransaction_EndRequest_Call {
	return &MockRequestTransaction_EndRequest_Call{Call: _e.mock.On("EndRequest")}
}

func (_c *MockRequestTransaction_EndRequest_Call) Run(run func()) *MockRequestTransaction_EndRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRequestTransaction_EndRequest_Call) Return(_a0 error) *MockRequestTransaction_EndRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestTransaction_EndRequest_Call) RunAndReturn(run func() error) *MockRequestTransaction_EndRequest_Call {
	_c.Call.Return(run)
	return _c
}

// FailRequest provides a mock function with given fields: err
func (_m *MockRequestTransaction) FailRequest(err error) error {
	ret := _m.Called(err)

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRequestTransaction_FailRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailRequest'
type MockRequestTransaction_FailRequest_Call struct {
	*mock.Call
}

// FailRequest is a helper method to define mock.On call
//   - err error
func (_e *MockRequestTransaction_Expecter) FailRequest(err interface{}) *MockRequestTransaction_FailRequest_Call {
	return &MockRequestTransaction_FailRequest_Call{Call: _e.mock.On("FailRequest", err)}
}

func (_c *MockRequestTransaction_FailRequest_Call) Run(run func(err error)) *MockRequestTransaction_FailRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *MockRequestTransaction_FailRequest_Call) Return(_a0 error) *MockRequestTransaction_FailRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestTransaction_FailRequest_Call) RunAndReturn(run func(error) error) *MockRequestTransaction_FailRequest_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompleted provides a mock function with given fields:
func (_m *MockRequestTransaction) IsCompleted() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockRequestTransaction_IsCompleted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompleted'
type MockRequestTransaction_IsCompleted_Call struct {
	*mock.Call
}

// IsCompleted is a helper method to define mock.On call
func (_e *MockRequestTransaction_Expecter) IsCompleted() *MockRequestTransaction_IsCompleted_Call {
	return &MockRequestTransaction_IsCompleted_Call{Call: _e.mock.On("IsCompleted")}
}

func (_c *MockRequestTransaction_IsCompleted_Call) Run(run func()) *MockRequestTransaction_IsCompleted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRequestTransaction_IsCompleted_Call) Return(_a0 bool) *MockRequestTransaction_IsCompleted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestTransaction_IsCompleted_Call) RunAndReturn(run func() bool) *MockRequestTransaction_IsCompleted_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockRequestTransaction) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockRequestTransaction_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockRequestTransaction_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockRequestTransaction_Expecter) String() *MockRequestTransaction_String_Call {
	return &MockRequestTransaction_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockRequestTransaction_String_Call) Run(run func()) *MockRequestTransaction_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRequestTransaction_String_Call) Return(_a0 string) *MockRequestTransaction_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestTransaction_String_Call) RunAndReturn(run func() string) *MockRequestTransaction_String_Call {
	_c.Call.Return(run)
	return _c
}

// Submit provides a mock function with given fields: operation
func (_m *MockRequestTransaction) Submit(operation RequestTransactionRunnable) {
	_m.Called(operation)
}

// MockRequestTransaction_Submit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Submit'
type MockRequestTransaction_Submit_Call struct {
	*mock.Call
}

// Submit is a helper method to define mock.On call
//   - operation RequestTransactionRunnable
func (_e *MockRequestTransaction_Expecter) Submit(operation interface{}) *MockRequestTransaction_Submit_Call {
	return &MockRequestTransaction_Submit_Call{Call: _e.mock.On("Submit", operation)}
}

func (_c *MockRequestTransaction_Submit_Call) Run(run func(operation RequestTransactionRunnable)) *MockRequestTransaction_Submit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(RequestTransactionRunnable))
	})
	return _c
}

func (_c *MockRequestTransaction_Submit_Call) Return() *MockRequestTransaction_Submit_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRequestTransaction_Submit_Call) RunAndReturn(run func(RequestTransactionRunnable)) *MockRequestTransaction_Submit_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRequestTransaction creates a new instance of MockRequestTransaction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRequestTransaction(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRequestTransaction {
	mock := &MockRequestTransaction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
