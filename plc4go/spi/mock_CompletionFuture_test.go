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

package spi

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockCompletionFuture is an autogenerated mock type for the CompletionFuture type
type MockCompletionFuture struct {
	mock.Mock
}

type MockCompletionFuture_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCompletionFuture) EXPECT() *MockCompletionFuture_Expecter {
	return &MockCompletionFuture_Expecter{mock: &_m.Mock}
}

// AwaitCompletion provides a mock function with given fields: ctx
func (_m *MockCompletionFuture) AwaitCompletion(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCompletionFuture_AwaitCompletion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AwaitCompletion'
type MockCompletionFuture_AwaitCompletion_Call struct {
	*mock.Call
}

// AwaitCompletion is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCompletionFuture_Expecter) AwaitCompletion(ctx interface{}) *MockCompletionFuture_AwaitCompletion_Call {
	return &MockCompletionFuture_AwaitCompletion_Call{Call: _e.mock.On("AwaitCompletion", ctx)}
}

func (_c *MockCompletionFuture_AwaitCompletion_Call) Run(run func(ctx context.Context)) *MockCompletionFuture_AwaitCompletion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCompletionFuture_AwaitCompletion_Call) Return(_a0 error) *MockCompletionFuture_AwaitCompletion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompletionFuture_AwaitCompletion_Call) RunAndReturn(run func(context.Context) error) *MockCompletionFuture_AwaitCompletion_Call {
	_c.Call.Return(run)
	return _c
}

// Cancel provides a mock function with given fields: interrupt, err
func (_m *MockCompletionFuture) Cancel(interrupt bool, err error) {
	_m.Called(interrupt, err)
}

// MockCompletionFuture_Cancel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cancel'
type MockCompletionFuture_Cancel_Call struct {
	*mock.Call
}

// Cancel is a helper method to define mock.On call
//   - interrupt bool
//   - err error
func (_e *MockCompletionFuture_Expecter) Cancel(interrupt interface{}, err interface{}) *MockCompletionFuture_Cancel_Call {
	return &MockCompletionFuture_Cancel_Call{Call: _e.mock.On("Cancel", interrupt, err)}
}

func (_c *MockCompletionFuture_Cancel_Call) Run(run func(interrupt bool, err error)) *MockCompletionFuture_Cancel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool), args[1].(error))
	})
	return _c
}

func (_c *MockCompletionFuture_Cancel_Call) Return() *MockCompletionFuture_Cancel_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCompletionFuture_Cancel_Call) RunAndReturn(run func(bool, error)) *MockCompletionFuture_Cancel_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCompletionFuture creates a new instance of MockCompletionFuture. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCompletionFuture(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCompletionFuture {
	mock := &MockCompletionFuture{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
