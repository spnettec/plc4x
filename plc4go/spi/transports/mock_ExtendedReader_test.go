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

package transports

import mock "github.com/stretchr/testify/mock"

// MockExtendedReader is an autogenerated mock type for the ExtendedReader type
type MockExtendedReader struct {
	mock.Mock
}

type MockExtendedReader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExtendedReader) EXPECT() *MockExtendedReader_Expecter {
	return &MockExtendedReader_Expecter{mock: &_m.Mock}
}

// Buffered provides a mock function with given fields:
func (_m *MockExtendedReader) Buffered() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockExtendedReader_Buffered_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Buffered'
type MockExtendedReader_Buffered_Call struct {
	*mock.Call
}

// Buffered is a helper method to define mock.On call
func (_e *MockExtendedReader_Expecter) Buffered() *MockExtendedReader_Buffered_Call {
	return &MockExtendedReader_Buffered_Call{Call: _e.mock.On("Buffered")}
}

func (_c *MockExtendedReader_Buffered_Call) Run(run func()) *MockExtendedReader_Buffered_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExtendedReader_Buffered_Call) Return(_a0 int) *MockExtendedReader_Buffered_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtendedReader_Buffered_Call) RunAndReturn(run func() int) *MockExtendedReader_Buffered_Call {
	_c.Call.Return(run)
	return _c
}

// Peek provides a mock function with given fields: _a0
func (_m *MockExtendedReader) Peek(_a0 int) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]byte, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExtendedReader_Peek_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Peek'
type MockExtendedReader_Peek_Call struct {
	*mock.Call
}

// Peek is a helper method to define mock.On call
//   - _a0 int
func (_e *MockExtendedReader_Expecter) Peek(_a0 interface{}) *MockExtendedReader_Peek_Call {
	return &MockExtendedReader_Peek_Call{Call: _e.mock.On("Peek", _a0)}
}

func (_c *MockExtendedReader_Peek_Call) Run(run func(_a0 int)) *MockExtendedReader_Peek_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockExtendedReader_Peek_Call) Return(_a0 []byte, _a1 error) *MockExtendedReader_Peek_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExtendedReader_Peek_Call) RunAndReturn(run func(int) ([]byte, error)) *MockExtendedReader_Peek_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: p
func (_m *MockExtendedReader) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExtendedReader_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockExtendedReader_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - p []byte
func (_e *MockExtendedReader_Expecter) Read(p interface{}) *MockExtendedReader_Read_Call {
	return &MockExtendedReader_Read_Call{Call: _e.mock.On("Read", p)}
}

func (_c *MockExtendedReader_Read_Call) Run(run func(p []byte)) *MockExtendedReader_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockExtendedReader_Read_Call) Return(n int, err error) *MockExtendedReader_Read_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *MockExtendedReader_Read_Call) RunAndReturn(run func([]byte) (int, error)) *MockExtendedReader_Read_Call {
	_c.Call.Return(run)
	return _c
}

// ReadByte provides a mock function with given fields:
func (_m *MockExtendedReader) ReadByte() (byte, error) {
	ret := _m.Called()

	var r0 byte
	var r1 error
	if rf, ok := ret.Get(0).(func() (byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() byte); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(byte)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExtendedReader_ReadByte_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadByte'
type MockExtendedReader_ReadByte_Call struct {
	*mock.Call
}

// ReadByte is a helper method to define mock.On call
func (_e *MockExtendedReader_Expecter) ReadByte() *MockExtendedReader_ReadByte_Call {
	return &MockExtendedReader_ReadByte_Call{Call: _e.mock.On("ReadByte")}
}

func (_c *MockExtendedReader_ReadByte_Call) Run(run func()) *MockExtendedReader_ReadByte_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExtendedReader_ReadByte_Call) Return(_a0 byte, _a1 error) *MockExtendedReader_ReadByte_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExtendedReader_ReadByte_Call) RunAndReturn(run func() (byte, error)) *MockExtendedReader_ReadByte_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExtendedReader creates a new instance of MockExtendedReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExtendedReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExtendedReader {
	mock := &MockExtendedReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
