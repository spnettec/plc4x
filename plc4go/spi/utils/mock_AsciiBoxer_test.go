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

// Code generated by mockery v2.30.1. DO NOT EDIT.

package utils

import mock "github.com/stretchr/testify/mock"

// MockAsciiBoxer is an autogenerated mock type for the AsciiBoxer type
type MockAsciiBoxer struct {
	mock.Mock
}

type MockAsciiBoxer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAsciiBoxer) EXPECT() *MockAsciiBoxer_Expecter {
	return &MockAsciiBoxer_Expecter{mock: &_m.Mock}
}

// Box provides a mock function with given fields: _a0, _a1
func (_m *MockAsciiBoxer) Box(_a0 string, _a1 int) AsciiBox {
	ret := _m.Called(_a0, _a1)

	var r0 AsciiBox
	if rf, ok := ret.Get(0).(func(string, int) AsciiBox); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(AsciiBox)
	}

	return r0
}

// MockAsciiBoxer_Box_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Box'
type MockAsciiBoxer_Box_Call struct {
	*mock.Call
}

// Box is a helper method to define mock.On call
//   - _a0 string
//   - _a1 int
func (_e *MockAsciiBoxer_Expecter) Box(_a0 interface{}, _a1 interface{}) *MockAsciiBoxer_Box_Call {
	return &MockAsciiBoxer_Box_Call{Call: _e.mock.On("Box", _a0, _a1)}
}

func (_c *MockAsciiBoxer_Box_Call) Run(run func(_a0 string, _a1 int)) *MockAsciiBoxer_Box_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *MockAsciiBoxer_Box_Call) Return(_a0 AsciiBox) *MockAsciiBoxer_Box_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAsciiBoxer_Box_Call) RunAndReturn(run func(string, int) AsciiBox) *MockAsciiBoxer_Box_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAsciiBoxer creates a new instance of MockAsciiBoxer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAsciiBoxer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAsciiBoxer {
	mock := &MockAsciiBoxer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
