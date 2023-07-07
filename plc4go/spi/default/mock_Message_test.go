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

package _default

import (
	context "context"

	utils "github.com/apache/plc4x/plc4go/spi/utils"
	mock "github.com/stretchr/testify/mock"
)

// MockMessage is an autogenerated mock type for the Message type
type MockMessage struct {
	mock.Mock
}

type MockMessage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMessage) EXPECT() *MockMessage_Expecter {
	return &MockMessage_Expecter{mock: &_m.Mock}
}

// GetLengthInBits provides a mock function with given fields: ctx
func (_m *MockMessage) GetLengthInBits(ctx context.Context) uint16 {
	ret := _m.Called(ctx)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(context.Context) uint16); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockMessage_GetLengthInBits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLengthInBits'
type MockMessage_GetLengthInBits_Call struct {
	*mock.Call
}

// GetLengthInBits is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockMessage_Expecter) GetLengthInBits(ctx interface{}) *MockMessage_GetLengthInBits_Call {
	return &MockMessage_GetLengthInBits_Call{Call: _e.mock.On("GetLengthInBits", ctx)}
}

func (_c *MockMessage_GetLengthInBits_Call) Run(run func(ctx context.Context)) *MockMessage_GetLengthInBits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMessage_GetLengthInBits_Call) Return(_a0 uint16) *MockMessage_GetLengthInBits_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessage_GetLengthInBits_Call) RunAndReturn(run func(context.Context) uint16) *MockMessage_GetLengthInBits_Call {
	_c.Call.Return(run)
	return _c
}

// GetLengthInBytes provides a mock function with given fields: ctx
func (_m *MockMessage) GetLengthInBytes(ctx context.Context) uint16 {
	ret := _m.Called(ctx)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(context.Context) uint16); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockMessage_GetLengthInBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLengthInBytes'
type MockMessage_GetLengthInBytes_Call struct {
	*mock.Call
}

// GetLengthInBytes is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockMessage_Expecter) GetLengthInBytes(ctx interface{}) *MockMessage_GetLengthInBytes_Call {
	return &MockMessage_GetLengthInBytes_Call{Call: _e.mock.On("GetLengthInBytes", ctx)}
}

func (_c *MockMessage_GetLengthInBytes_Call) Run(run func(ctx context.Context)) *MockMessage_GetLengthInBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMessage_GetLengthInBytes_Call) Return(_a0 uint16) *MockMessage_GetLengthInBytes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessage_GetLengthInBytes_Call) RunAndReturn(run func(context.Context) uint16) *MockMessage_GetLengthInBytes_Call {
	_c.Call.Return(run)
	return _c
}

// Serialize provides a mock function with given fields:
func (_m *MockMessage) Serialize() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMessage_Serialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serialize'
type MockMessage_Serialize_Call struct {
	*mock.Call
}

// Serialize is a helper method to define mock.On call
func (_e *MockMessage_Expecter) Serialize() *MockMessage_Serialize_Call {
	return &MockMessage_Serialize_Call{Call: _e.mock.On("Serialize")}
}

func (_c *MockMessage_Serialize_Call) Run(run func()) *MockMessage_Serialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMessage_Serialize_Call) Return(_a0 []byte, _a1 error) *MockMessage_Serialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMessage_Serialize_Call) RunAndReturn(run func() ([]byte, error)) *MockMessage_Serialize_Call {
	_c.Call.Return(run)
	return _c
}

// SerializeWithWriteBuffer provides a mock function with given fields: ctx, writeBuffer
func (_m *MockMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	ret := _m.Called(ctx, writeBuffer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, utils.WriteBuffer) error); ok {
		r0 = rf(ctx, writeBuffer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMessage_SerializeWithWriteBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SerializeWithWriteBuffer'
type MockMessage_SerializeWithWriteBuffer_Call struct {
	*mock.Call
}

// SerializeWithWriteBuffer is a helper method to define mock.On call
//   - ctx context.Context
//   - writeBuffer utils.WriteBuffer
func (_e *MockMessage_Expecter) SerializeWithWriteBuffer(ctx interface{}, writeBuffer interface{}) *MockMessage_SerializeWithWriteBuffer_Call {
	return &MockMessage_SerializeWithWriteBuffer_Call{Call: _e.mock.On("SerializeWithWriteBuffer", ctx, writeBuffer)}
}

func (_c *MockMessage_SerializeWithWriteBuffer_Call) Run(run func(ctx context.Context, writeBuffer utils.WriteBuffer)) *MockMessage_SerializeWithWriteBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(utils.WriteBuffer))
	})
	return _c
}

func (_c *MockMessage_SerializeWithWriteBuffer_Call) Return(_a0 error) *MockMessage_SerializeWithWriteBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessage_SerializeWithWriteBuffer_Call) RunAndReturn(run func(context.Context, utils.WriteBuffer) error) *MockMessage_SerializeWithWriteBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMessage creates a new instance of MockMessage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMessage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMessage {
	mock := &MockMessage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
