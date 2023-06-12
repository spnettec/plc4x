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

package testutils

import (
	bufio "bufio"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockTestTransportInstance is an autogenerated mock type for the TestTransportInstance type
type MockTestTransportInstance struct {
	mock.Mock
}

type MockTestTransportInstance_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTestTransportInstance) EXPECT() *MockTestTransportInstance_Expecter {
	return &MockTestTransportInstance_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockTestTransportInstance) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestTransportInstance_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockTestTransportInstance_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockTestTransportInstance_Expecter) Close() *MockTestTransportInstance_Close_Call {
	return &MockTestTransportInstance_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockTestTransportInstance_Close_Call) Run(run func()) *MockTestTransportInstance_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestTransportInstance_Close_Call) Return(_a0 error) *MockTestTransportInstance_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestTransportInstance_Close_Call) RunAndReturn(run func() error) *MockTestTransportInstance_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Connect provides a mock function with given fields:
func (_m *MockTestTransportInstance) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestTransportInstance_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockTestTransportInstance_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockTestTransportInstance_Expecter) Connect() *MockTestTransportInstance_Connect_Call {
	return &MockTestTransportInstance_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockTestTransportInstance_Connect_Call) Run(run func()) *MockTestTransportInstance_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestTransportInstance_Connect_Call) Return(_a0 error) *MockTestTransportInstance_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestTransportInstance_Connect_Call) RunAndReturn(run func() error) *MockTestTransportInstance_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// ConnectWithContext provides a mock function with given fields: ctx
func (_m *MockTestTransportInstance) ConnectWithContext(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestTransportInstance_ConnectWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectWithContext'
type MockTestTransportInstance_ConnectWithContext_Call struct {
	*mock.Call
}

// ConnectWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockTestTransportInstance_Expecter) ConnectWithContext(ctx interface{}) *MockTestTransportInstance_ConnectWithContext_Call {
	return &MockTestTransportInstance_ConnectWithContext_Call{Call: _e.mock.On("ConnectWithContext", ctx)}
}

func (_c *MockTestTransportInstance_ConnectWithContext_Call) Run(run func(ctx context.Context)) *MockTestTransportInstance_ConnectWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockTestTransportInstance_ConnectWithContext_Call) Return(_a0 error) *MockTestTransportInstance_ConnectWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestTransportInstance_ConnectWithContext_Call) RunAndReturn(run func(context.Context) error) *MockTestTransportInstance_ConnectWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// DrainWriteBuffer provides a mock function with given fields: numBytes
func (_m *MockTestTransportInstance) DrainWriteBuffer(numBytes uint32) []byte {
	ret := _m.Called(numBytes)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(uint32) []byte); ok {
		r0 = rf(numBytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockTestTransportInstance_DrainWriteBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DrainWriteBuffer'
type MockTestTransportInstance_DrainWriteBuffer_Call struct {
	*mock.Call
}

// DrainWriteBuffer is a helper method to define mock.On call
//   - numBytes uint32
func (_e *MockTestTransportInstance_Expecter) DrainWriteBuffer(numBytes interface{}) *MockTestTransportInstance_DrainWriteBuffer_Call {
	return &MockTestTransportInstance_DrainWriteBuffer_Call{Call: _e.mock.On("DrainWriteBuffer", numBytes)}
}

func (_c *MockTestTransportInstance_DrainWriteBuffer_Call) Run(run func(numBytes uint32)) *MockTestTransportInstance_DrainWriteBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockTestTransportInstance_DrainWriteBuffer_Call) Return(_a0 []byte) *MockTestTransportInstance_DrainWriteBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestTransportInstance_DrainWriteBuffer_Call) RunAndReturn(run func(uint32) []byte) *MockTestTransportInstance_DrainWriteBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// FillBuffer provides a mock function with given fields: until
func (_m *MockTestTransportInstance) FillBuffer(until func(uint, byte, *bufio.Reader) bool) error {
	ret := _m.Called(until)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(uint, byte, *bufio.Reader) bool) error); ok {
		r0 = rf(until)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestTransportInstance_FillBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FillBuffer'
type MockTestTransportInstance_FillBuffer_Call struct {
	*mock.Call
}

// FillBuffer is a helper method to define mock.On call
//   - until func(uint , byte , *bufio.Reader) bool
func (_e *MockTestTransportInstance_Expecter) FillBuffer(until interface{}) *MockTestTransportInstance_FillBuffer_Call {
	return &MockTestTransportInstance_FillBuffer_Call{Call: _e.mock.On("FillBuffer", until)}
}

func (_c *MockTestTransportInstance_FillBuffer_Call) Run(run func(until func(uint, byte, *bufio.Reader) bool)) *MockTestTransportInstance_FillBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(uint, byte, *bufio.Reader) bool))
	})
	return _c
}

func (_c *MockTestTransportInstance_FillBuffer_Call) Return(_a0 error) *MockTestTransportInstance_FillBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestTransportInstance_FillBuffer_Call) RunAndReturn(run func(func(uint, byte, *bufio.Reader) bool) error) *MockTestTransportInstance_FillBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// FillReadBuffer provides a mock function with given fields: data
func (_m *MockTestTransportInstance) FillReadBuffer(data []byte) {
	_m.Called(data)
}

// MockTestTransportInstance_FillReadBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FillReadBuffer'
type MockTestTransportInstance_FillReadBuffer_Call struct {
	*mock.Call
}

// FillReadBuffer is a helper method to define mock.On call
//   - data []byte
func (_e *MockTestTransportInstance_Expecter) FillReadBuffer(data interface{}) *MockTestTransportInstance_FillReadBuffer_Call {
	return &MockTestTransportInstance_FillReadBuffer_Call{Call: _e.mock.On("FillReadBuffer", data)}
}

func (_c *MockTestTransportInstance_FillReadBuffer_Call) Run(run func(data []byte)) *MockTestTransportInstance_FillReadBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockTestTransportInstance_FillReadBuffer_Call) Return() *MockTestTransportInstance_FillReadBuffer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTestTransportInstance_FillReadBuffer_Call) RunAndReturn(run func([]byte)) *MockTestTransportInstance_FillReadBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// GetNumBytesAvailableInBuffer provides a mock function with given fields:
func (_m *MockTestTransportInstance) GetNumBytesAvailableInBuffer() (uint32, error) {
	ret := _m.Called()

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint32, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTestTransportInstance_GetNumBytesAvailableInBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNumBytesAvailableInBuffer'
type MockTestTransportInstance_GetNumBytesAvailableInBuffer_Call struct {
	*mock.Call
}

// GetNumBytesAvailableInBuffer is a helper method to define mock.On call
func (_e *MockTestTransportInstance_Expecter) GetNumBytesAvailableInBuffer() *MockTestTransportInstance_GetNumBytesAvailableInBuffer_Call {
	return &MockTestTransportInstance_GetNumBytesAvailableInBuffer_Call{Call: _e.mock.On("GetNumBytesAvailableInBuffer")}
}

func (_c *MockTestTransportInstance_GetNumBytesAvailableInBuffer_Call) Run(run func()) *MockTestTransportInstance_GetNumBytesAvailableInBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestTransportInstance_GetNumBytesAvailableInBuffer_Call) Return(_a0 uint32, _a1 error) *MockTestTransportInstance_GetNumBytesAvailableInBuffer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTestTransportInstance_GetNumBytesAvailableInBuffer_Call) RunAndReturn(run func() (uint32, error)) *MockTestTransportInstance_GetNumBytesAvailableInBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// GetNumDrainableBytes provides a mock function with given fields:
func (_m *MockTestTransportInstance) GetNumDrainableBytes() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// MockTestTransportInstance_GetNumDrainableBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNumDrainableBytes'
type MockTestTransportInstance_GetNumDrainableBytes_Call struct {
	*mock.Call
}

// GetNumDrainableBytes is a helper method to define mock.On call
func (_e *MockTestTransportInstance_Expecter) GetNumDrainableBytes() *MockTestTransportInstance_GetNumDrainableBytes_Call {
	return &MockTestTransportInstance_GetNumDrainableBytes_Call{Call: _e.mock.On("GetNumDrainableBytes")}
}

func (_c *MockTestTransportInstance_GetNumDrainableBytes_Call) Run(run func()) *MockTestTransportInstance_GetNumDrainableBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestTransportInstance_GetNumDrainableBytes_Call) Return(_a0 uint32) *MockTestTransportInstance_GetNumDrainableBytes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestTransportInstance_GetNumDrainableBytes_Call) RunAndReturn(run func() uint32) *MockTestTransportInstance_GetNumDrainableBytes_Call {
	_c.Call.Return(run)
	return _c
}

// IsConnected provides a mock function with given fields:
func (_m *MockTestTransportInstance) IsConnected() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockTestTransportInstance_IsConnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsConnected'
type MockTestTransportInstance_IsConnected_Call struct {
	*mock.Call
}

// IsConnected is a helper method to define mock.On call
func (_e *MockTestTransportInstance_Expecter) IsConnected() *MockTestTransportInstance_IsConnected_Call {
	return &MockTestTransportInstance_IsConnected_Call{Call: _e.mock.On("IsConnected")}
}

func (_c *MockTestTransportInstance_IsConnected_Call) Run(run func()) *MockTestTransportInstance_IsConnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestTransportInstance_IsConnected_Call) Return(_a0 bool) *MockTestTransportInstance_IsConnected_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestTransportInstance_IsConnected_Call) RunAndReturn(run func() bool) *MockTestTransportInstance_IsConnected_Call {
	_c.Call.Return(run)
	return _c
}

// PeekReadableBytes provides a mock function with given fields: numBytes
func (_m *MockTestTransportInstance) PeekReadableBytes(numBytes uint32) ([]byte, error) {
	ret := _m.Called(numBytes)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32) ([]byte, error)); ok {
		return rf(numBytes)
	}
	if rf, ok := ret.Get(0).(func(uint32) []byte); ok {
		r0 = rf(numBytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(numBytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTestTransportInstance_PeekReadableBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PeekReadableBytes'
type MockTestTransportInstance_PeekReadableBytes_Call struct {
	*mock.Call
}

// PeekReadableBytes is a helper method to define mock.On call
//   - numBytes uint32
func (_e *MockTestTransportInstance_Expecter) PeekReadableBytes(numBytes interface{}) *MockTestTransportInstance_PeekReadableBytes_Call {
	return &MockTestTransportInstance_PeekReadableBytes_Call{Call: _e.mock.On("PeekReadableBytes", numBytes)}
}

func (_c *MockTestTransportInstance_PeekReadableBytes_Call) Run(run func(numBytes uint32)) *MockTestTransportInstance_PeekReadableBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockTestTransportInstance_PeekReadableBytes_Call) Return(_a0 []byte, _a1 error) *MockTestTransportInstance_PeekReadableBytes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTestTransportInstance_PeekReadableBytes_Call) RunAndReturn(run func(uint32) ([]byte, error)) *MockTestTransportInstance_PeekReadableBytes_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: numBytes
func (_m *MockTestTransportInstance) Read(numBytes uint32) ([]byte, error) {
	ret := _m.Called(numBytes)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32) ([]byte, error)); ok {
		return rf(numBytes)
	}
	if rf, ok := ret.Get(0).(func(uint32) []byte); ok {
		r0 = rf(numBytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(numBytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTestTransportInstance_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockTestTransportInstance_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - numBytes uint32
func (_e *MockTestTransportInstance_Expecter) Read(numBytes interface{}) *MockTestTransportInstance_Read_Call {
	return &MockTestTransportInstance_Read_Call{Call: _e.mock.On("Read", numBytes)}
}

func (_c *MockTestTransportInstance_Read_Call) Run(run func(numBytes uint32)) *MockTestTransportInstance_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockTestTransportInstance_Read_Call) Return(_a0 []byte, _a1 error) *MockTestTransportInstance_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTestTransportInstance_Read_Call) RunAndReturn(run func(uint32) ([]byte, error)) *MockTestTransportInstance_Read_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockTestTransportInstance) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTestTransportInstance_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockTestTransportInstance_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockTestTransportInstance_Expecter) String() *MockTestTransportInstance_String_Call {
	return &MockTestTransportInstance_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockTestTransportInstance_String_Call) Run(run func()) *MockTestTransportInstance_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestTransportInstance_String_Call) Return(_a0 string) *MockTestTransportInstance_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestTransportInstance_String_Call) RunAndReturn(run func() string) *MockTestTransportInstance_String_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: data
func (_m *MockTestTransportInstance) Write(data []byte) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestTransportInstance_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type MockTestTransportInstance_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - data []byte
func (_e *MockTestTransportInstance_Expecter) Write(data interface{}) *MockTestTransportInstance_Write_Call {
	return &MockTestTransportInstance_Write_Call{Call: _e.mock.On("Write", data)}
}

func (_c *MockTestTransportInstance_Write_Call) Run(run func(data []byte)) *MockTestTransportInstance_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockTestTransportInstance_Write_Call) Return(_a0 error) *MockTestTransportInstance_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestTransportInstance_Write_Call) RunAndReturn(run func([]byte) error) *MockTestTransportInstance_Write_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockTestTransportInstance interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTestTransportInstance creates a new instance of MockTestTransportInstance. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTestTransportInstance(t mockConstructorTestingTNewMockTestTransportInstance) *MockTestTransportInstance {
	mock := &MockTestTransportInstance{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
