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

// Code generated by mockery v2.42.2. DO NOT EDIT.

package bacnetip

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	mock "github.com/stretchr/testify/mock"

	spi "github.com/apache/plc4x/plc4go/spi"

	utils "github.com/apache/plc4x/plc4go/spi/utils"
)

// MockBVLPDU is an autogenerated mock type for the BVLPDU type
type MockBVLPDU struct {
	mock.Mock
}

type MockBVLPDU_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBVLPDU) EXPECT() *MockBVLPDU_Expecter {
	return &MockBVLPDU_Expecter{mock: &_m.Mock}
}

// Decode provides a mock function with given fields: pdu
func (_m *MockBVLPDU) Decode(pdu Arg) error {
	ret := _m.Called(pdu)

	if len(ret) == 0 {
		panic("no return value specified for Decode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Arg) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBVLPDU_Decode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decode'
type MockBVLPDU_Decode_Call struct {
	*mock.Call
}

// Decode is a helper method to define mock.On call
//   - pdu Arg
func (_e *MockBVLPDU_Expecter) Decode(pdu interface{}) *MockBVLPDU_Decode_Call {
	return &MockBVLPDU_Decode_Call{Call: _e.mock.On("Decode", pdu)}
}

func (_c *MockBVLPDU_Decode_Call) Run(run func(pdu Arg)) *MockBVLPDU_Decode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Arg))
	})
	return _c
}

func (_c *MockBVLPDU_Decode_Call) Return(_a0 error) *MockBVLPDU_Decode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_Decode_Call) RunAndReturn(run func(Arg) error) *MockBVLPDU_Decode_Call {
	_c.Call.Return(run)
	return _c
}

// Encode provides a mock function with given fields: pdu
func (_m *MockBVLPDU) Encode(pdu Arg) error {
	ret := _m.Called(pdu)

	if len(ret) == 0 {
		panic("no return value specified for Encode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Arg) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBVLPDU_Encode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encode'
type MockBVLPDU_Encode_Call struct {
	*mock.Call
}

// Encode is a helper method to define mock.On call
//   - pdu Arg
func (_e *MockBVLPDU_Expecter) Encode(pdu interface{}) *MockBVLPDU_Encode_Call {
	return &MockBVLPDU_Encode_Call{Call: _e.mock.On("Encode", pdu)}
}

func (_c *MockBVLPDU_Encode_Call) Run(run func(pdu Arg)) *MockBVLPDU_Encode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Arg))
	})
	return _c
}

func (_c *MockBVLPDU_Encode_Call) Return(_a0 error) *MockBVLPDU_Encode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_Encode_Call) RunAndReturn(run func(Arg) error) *MockBVLPDU_Encode_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields:
func (_m *MockBVLPDU) Get() (byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

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

// MockBVLPDU_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockBVLPDU_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) Get() *MockBVLPDU_Get_Call {
	return &MockBVLPDU_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *MockBVLPDU_Get_Call) Run(run func()) *MockBVLPDU_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_Get_Call) Return(_a0 byte, _a1 error) *MockBVLPDU_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBVLPDU_Get_Call) RunAndReturn(run func() (byte, error)) *MockBVLPDU_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetBvlcFunction provides a mock function with given fields:
func (_m *MockBVLPDU) GetBvlcFunction() uint8 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBvlcFunction")
	}

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}

// MockBVLPDU_GetBvlcFunction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBvlcFunction'
type MockBVLPDU_GetBvlcFunction_Call struct {
	*mock.Call
}

// GetBvlcFunction is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) GetBvlcFunction() *MockBVLPDU_GetBvlcFunction_Call {
	return &MockBVLPDU_GetBvlcFunction_Call{Call: _e.mock.On("GetBvlcFunction")}
}

func (_c *MockBVLPDU_GetBvlcFunction_Call) Run(run func()) *MockBVLPDU_GetBvlcFunction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_GetBvlcFunction_Call) Return(_a0 uint8) *MockBVLPDU_GetBvlcFunction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_GetBvlcFunction_Call) RunAndReturn(run func() uint8) *MockBVLPDU_GetBvlcFunction_Call {
	_c.Call.Return(run)
	return _c
}

// GetBvlcPayloadLength provides a mock function with given fields:
func (_m *MockBVLPDU) GetBvlcPayloadLength() uint16 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBvlcPayloadLength")
	}

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockBVLPDU_GetBvlcPayloadLength_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBvlcPayloadLength'
type MockBVLPDU_GetBvlcPayloadLength_Call struct {
	*mock.Call
}

// GetBvlcPayloadLength is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) GetBvlcPayloadLength() *MockBVLPDU_GetBvlcPayloadLength_Call {
	return &MockBVLPDU_GetBvlcPayloadLength_Call{Call: _e.mock.On("GetBvlcPayloadLength")}
}

func (_c *MockBVLPDU_GetBvlcPayloadLength_Call) Run(run func()) *MockBVLPDU_GetBvlcPayloadLength_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_GetBvlcPayloadLength_Call) Return(_a0 uint16) *MockBVLPDU_GetBvlcPayloadLength_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_GetBvlcPayloadLength_Call) RunAndReturn(run func() uint16) *MockBVLPDU_GetBvlcPayloadLength_Call {
	_c.Call.Return(run)
	return _c
}

// GetData provides a mock function with given fields: dlen
func (_m *MockBVLPDU) GetData(dlen int) ([]byte, error) {
	ret := _m.Called(dlen)

	if len(ret) == 0 {
		panic("no return value specified for GetData")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]byte, error)); ok {
		return rf(dlen)
	}
	if rf, ok := ret.Get(0).(func(int) []byte); ok {
		r0 = rf(dlen)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(dlen)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBVLPDU_GetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetData'
type MockBVLPDU_GetData_Call struct {
	*mock.Call
}

// GetData is a helper method to define mock.On call
//   - dlen int
func (_e *MockBVLPDU_Expecter) GetData(dlen interface{}) *MockBVLPDU_GetData_Call {
	return &MockBVLPDU_GetData_Call{Call: _e.mock.On("GetData", dlen)}
}

func (_c *MockBVLPDU_GetData_Call) Run(run func(dlen int)) *MockBVLPDU_GetData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockBVLPDU_GetData_Call) Return(_a0 []byte, _a1 error) *MockBVLPDU_GetData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBVLPDU_GetData_Call) RunAndReturn(run func(int) ([]byte, error)) *MockBVLPDU_GetData_Call {
	_c.Call.Return(run)
	return _c
}

// GetExpectingReply provides a mock function with given fields:
func (_m *MockBVLPDU) GetExpectingReply() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExpectingReply")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockBVLPDU_GetExpectingReply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExpectingReply'
type MockBVLPDU_GetExpectingReply_Call struct {
	*mock.Call
}

// GetExpectingReply is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) GetExpectingReply() *MockBVLPDU_GetExpectingReply_Call {
	return &MockBVLPDU_GetExpectingReply_Call{Call: _e.mock.On("GetExpectingReply")}
}

func (_c *MockBVLPDU_GetExpectingReply_Call) Run(run func()) *MockBVLPDU_GetExpectingReply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_GetExpectingReply_Call) Return(_a0 bool) *MockBVLPDU_GetExpectingReply_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_GetExpectingReply_Call) RunAndReturn(run func() bool) *MockBVLPDU_GetExpectingReply_Call {
	_c.Call.Return(run)
	return _c
}

// GetLengthInBits provides a mock function with given fields: ctx
func (_m *MockBVLPDU) GetLengthInBits(ctx context.Context) uint16 {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLengthInBits")
	}

	var r0 uint16
	if rf, ok := ret.Get(0).(func(context.Context) uint16); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockBVLPDU_GetLengthInBits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLengthInBits'
type MockBVLPDU_GetLengthInBits_Call struct {
	*mock.Call
}

// GetLengthInBits is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockBVLPDU_Expecter) GetLengthInBits(ctx interface{}) *MockBVLPDU_GetLengthInBits_Call {
	return &MockBVLPDU_GetLengthInBits_Call{Call: _e.mock.On("GetLengthInBits", ctx)}
}

func (_c *MockBVLPDU_GetLengthInBits_Call) Run(run func(ctx context.Context)) *MockBVLPDU_GetLengthInBits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBVLPDU_GetLengthInBits_Call) Return(_a0 uint16) *MockBVLPDU_GetLengthInBits_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_GetLengthInBits_Call) RunAndReturn(run func(context.Context) uint16) *MockBVLPDU_GetLengthInBits_Call {
	_c.Call.Return(run)
	return _c
}

// GetLengthInBytes provides a mock function with given fields: ctx
func (_m *MockBVLPDU) GetLengthInBytes(ctx context.Context) uint16 {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLengthInBytes")
	}

	var r0 uint16
	if rf, ok := ret.Get(0).(func(context.Context) uint16); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockBVLPDU_GetLengthInBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLengthInBytes'
type MockBVLPDU_GetLengthInBytes_Call struct {
	*mock.Call
}

// GetLengthInBytes is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockBVLPDU_Expecter) GetLengthInBytes(ctx interface{}) *MockBVLPDU_GetLengthInBytes_Call {
	return &MockBVLPDU_GetLengthInBytes_Call{Call: _e.mock.On("GetLengthInBytes", ctx)}
}

func (_c *MockBVLPDU_GetLengthInBytes_Call) Run(run func(ctx context.Context)) *MockBVLPDU_GetLengthInBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBVLPDU_GetLengthInBytes_Call) Return(_a0 uint16) *MockBVLPDU_GetLengthInBytes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_GetLengthInBytes_Call) RunAndReturn(run func(context.Context) uint16) *MockBVLPDU_GetLengthInBytes_Call {
	_c.Call.Return(run)
	return _c
}

// GetLong provides a mock function with given fields:
func (_m *MockBVLPDU) GetLong() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLong")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBVLPDU_GetLong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLong'
type MockBVLPDU_GetLong_Call struct {
	*mock.Call
}

// GetLong is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) GetLong() *MockBVLPDU_GetLong_Call {
	return &MockBVLPDU_GetLong_Call{Call: _e.mock.On("GetLong")}
}

func (_c *MockBVLPDU_GetLong_Call) Run(run func()) *MockBVLPDU_GetLong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_GetLong_Call) Return(_a0 int64, _a1 error) *MockBVLPDU_GetLong_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBVLPDU_GetLong_Call) RunAndReturn(run func() (int64, error)) *MockBVLPDU_GetLong_Call {
	_c.Call.Return(run)
	return _c
}

// GetNetworkPriority provides a mock function with given fields:
func (_m *MockBVLPDU) GetNetworkPriority() model.NPDUNetworkPriority {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkPriority")
	}

	var r0 model.NPDUNetworkPriority
	if rf, ok := ret.Get(0).(func() model.NPDUNetworkPriority); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.NPDUNetworkPriority)
	}

	return r0
}

// MockBVLPDU_GetNetworkPriority_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNetworkPriority'
type MockBVLPDU_GetNetworkPriority_Call struct {
	*mock.Call
}

// GetNetworkPriority is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) GetNetworkPriority() *MockBVLPDU_GetNetworkPriority_Call {
	return &MockBVLPDU_GetNetworkPriority_Call{Call: _e.mock.On("GetNetworkPriority")}
}

func (_c *MockBVLPDU_GetNetworkPriority_Call) Run(run func()) *MockBVLPDU_GetNetworkPriority_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_GetNetworkPriority_Call) Return(_a0 model.NPDUNetworkPriority) *MockBVLPDU_GetNetworkPriority_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_GetNetworkPriority_Call) RunAndReturn(run func() model.NPDUNetworkPriority) *MockBVLPDU_GetNetworkPriority_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUDestination provides a mock function with given fields:
func (_m *MockBVLPDU) GetPDUDestination() *Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPDUDestination")
	}

	var r0 *Address
	if rf, ok := ret.Get(0).(func() *Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Address)
		}
	}

	return r0
}

// MockBVLPDU_GetPDUDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUDestination'
type MockBVLPDU_GetPDUDestination_Call struct {
	*mock.Call
}

// GetPDUDestination is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) GetPDUDestination() *MockBVLPDU_GetPDUDestination_Call {
	return &MockBVLPDU_GetPDUDestination_Call{Call: _e.mock.On("GetPDUDestination")}
}

func (_c *MockBVLPDU_GetPDUDestination_Call) Run(run func()) *MockBVLPDU_GetPDUDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_GetPDUDestination_Call) Return(_a0 *Address) *MockBVLPDU_GetPDUDestination_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_GetPDUDestination_Call) RunAndReturn(run func() *Address) *MockBVLPDU_GetPDUDestination_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUSource provides a mock function with given fields:
func (_m *MockBVLPDU) GetPDUSource() *Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPDUSource")
	}

	var r0 *Address
	if rf, ok := ret.Get(0).(func() *Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Address)
		}
	}

	return r0
}

// MockBVLPDU_GetPDUSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUSource'
type MockBVLPDU_GetPDUSource_Call struct {
	*mock.Call
}

// GetPDUSource is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) GetPDUSource() *MockBVLPDU_GetPDUSource_Call {
	return &MockBVLPDU_GetPDUSource_Call{Call: _e.mock.On("GetPDUSource")}
}

func (_c *MockBVLPDU_GetPDUSource_Call) Run(run func()) *MockBVLPDU_GetPDUSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_GetPDUSource_Call) Return(_a0 *Address) *MockBVLPDU_GetPDUSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_GetPDUSource_Call) RunAndReturn(run func() *Address) *MockBVLPDU_GetPDUSource_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUUserData provides a mock function with given fields:
func (_m *MockBVLPDU) GetPDUUserData() spi.Message {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPDUUserData")
	}

	var r0 spi.Message
	if rf, ok := ret.Get(0).(func() spi.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.Message)
		}
	}

	return r0
}

// MockBVLPDU_GetPDUUserData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUUserData'
type MockBVLPDU_GetPDUUserData_Call struct {
	*mock.Call
}

// GetPDUUserData is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) GetPDUUserData() *MockBVLPDU_GetPDUUserData_Call {
	return &MockBVLPDU_GetPDUUserData_Call{Call: _e.mock.On("GetPDUUserData")}
}

func (_c *MockBVLPDU_GetPDUUserData_Call) Run(run func()) *MockBVLPDU_GetPDUUserData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_GetPDUUserData_Call) Return(_a0 spi.Message) *MockBVLPDU_GetPDUUserData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_GetPDUUserData_Call) RunAndReturn(run func() spi.Message) *MockBVLPDU_GetPDUUserData_Call {
	_c.Call.Return(run)
	return _c
}

// GetPduData provides a mock function with given fields:
func (_m *MockBVLPDU) GetPduData() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPduData")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockBVLPDU_GetPduData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPduData'
type MockBVLPDU_GetPduData_Call struct {
	*mock.Call
}

// GetPduData is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) GetPduData() *MockBVLPDU_GetPduData_Call {
	return &MockBVLPDU_GetPduData_Call{Call: _e.mock.On("GetPduData")}
}

func (_c *MockBVLPDU_GetPduData_Call) Run(run func()) *MockBVLPDU_GetPduData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_GetPduData_Call) Return(_a0 []byte) *MockBVLPDU_GetPduData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_GetPduData_Call) RunAndReturn(run func() []byte) *MockBVLPDU_GetPduData_Call {
	_c.Call.Return(run)
	return _c
}

// GetShort provides a mock function with given fields:
func (_m *MockBVLPDU) GetShort() (int16, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetShort")
	}

	var r0 int16
	var r1 error
	if rf, ok := ret.Get(0).(func() (int16, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int16)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBVLPDU_GetShort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShort'
type MockBVLPDU_GetShort_Call struct {
	*mock.Call
}

// GetShort is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) GetShort() *MockBVLPDU_GetShort_Call {
	return &MockBVLPDU_GetShort_Call{Call: _e.mock.On("GetShort")}
}

func (_c *MockBVLPDU_GetShort_Call) Run(run func()) *MockBVLPDU_GetShort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_GetShort_Call) Return(_a0 int16, _a1 error) *MockBVLPDU_GetShort_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBVLPDU_GetShort_Call) RunAndReturn(run func() (int16, error)) *MockBVLPDU_GetShort_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: _a0
func (_m *MockBVLPDU) Put(_a0 byte) {
	_m.Called(_a0)
}

// MockBVLPDU_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockBVLPDU_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - _a0 byte
func (_e *MockBVLPDU_Expecter) Put(_a0 interface{}) *MockBVLPDU_Put_Call {
	return &MockBVLPDU_Put_Call{Call: _e.mock.On("Put", _a0)}
}

func (_c *MockBVLPDU_Put_Call) Run(run func(_a0 byte)) *MockBVLPDU_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(byte))
	})
	return _c
}

func (_c *MockBVLPDU_Put_Call) Return() *MockBVLPDU_Put_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBVLPDU_Put_Call) RunAndReturn(run func(byte)) *MockBVLPDU_Put_Call {
	_c.Call.Return(run)
	return _c
}

// PutData provides a mock function with given fields: _a0
func (_m *MockBVLPDU) PutData(_a0 ...byte) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockBVLPDU_PutData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutData'
type MockBVLPDU_PutData_Call struct {
	*mock.Call
}

// PutData is a helper method to define mock.On call
//   - _a0 ...byte
func (_e *MockBVLPDU_Expecter) PutData(_a0 ...interface{}) *MockBVLPDU_PutData_Call {
	return &MockBVLPDU_PutData_Call{Call: _e.mock.On("PutData",
		append([]interface{}{}, _a0...)...)}
}

func (_c *MockBVLPDU_PutData_Call) Run(run func(_a0 ...byte)) *MockBVLPDU_PutData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]byte, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(byte)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockBVLPDU_PutData_Call) Return() *MockBVLPDU_PutData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBVLPDU_PutData_Call) RunAndReturn(run func(...byte)) *MockBVLPDU_PutData_Call {
	_c.Call.Return(run)
	return _c
}

// PutLong provides a mock function with given fields: _a0
func (_m *MockBVLPDU) PutLong(_a0 uint32) {
	_m.Called(_a0)
}

// MockBVLPDU_PutLong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutLong'
type MockBVLPDU_PutLong_Call struct {
	*mock.Call
}

// PutLong is a helper method to define mock.On call
//   - _a0 uint32
func (_e *MockBVLPDU_Expecter) PutLong(_a0 interface{}) *MockBVLPDU_PutLong_Call {
	return &MockBVLPDU_PutLong_Call{Call: _e.mock.On("PutLong", _a0)}
}

func (_c *MockBVLPDU_PutLong_Call) Run(run func(_a0 uint32)) *MockBVLPDU_PutLong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockBVLPDU_PutLong_Call) Return() *MockBVLPDU_PutLong_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBVLPDU_PutLong_Call) RunAndReturn(run func(uint32)) *MockBVLPDU_PutLong_Call {
	_c.Call.Return(run)
	return _c
}

// PutShort provides a mock function with given fields: _a0
func (_m *MockBVLPDU) PutShort(_a0 uint16) {
	_m.Called(_a0)
}

// MockBVLPDU_PutShort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutShort'
type MockBVLPDU_PutShort_Call struct {
	*mock.Call
}

// PutShort is a helper method to define mock.On call
//   - _a0 uint16
func (_e *MockBVLPDU_Expecter) PutShort(_a0 interface{}) *MockBVLPDU_PutShort_Call {
	return &MockBVLPDU_PutShort_Call{Call: _e.mock.On("PutShort", _a0)}
}

func (_c *MockBVLPDU_PutShort_Call) Run(run func(_a0 uint16)) *MockBVLPDU_PutShort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint16))
	})
	return _c
}

func (_c *MockBVLPDU_PutShort_Call) Return() *MockBVLPDU_PutShort_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBVLPDU_PutShort_Call) RunAndReturn(run func(uint16)) *MockBVLPDU_PutShort_Call {
	_c.Call.Return(run)
	return _c
}

// Serialize provides a mock function with given fields:
func (_m *MockBVLPDU) Serialize() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Serialize")
	}

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

// MockBVLPDU_Serialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serialize'
type MockBVLPDU_Serialize_Call struct {
	*mock.Call
}

// Serialize is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) Serialize() *MockBVLPDU_Serialize_Call {
	return &MockBVLPDU_Serialize_Call{Call: _e.mock.On("Serialize")}
}

func (_c *MockBVLPDU_Serialize_Call) Run(run func()) *MockBVLPDU_Serialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_Serialize_Call) Return(_a0 []byte, _a1 error) *MockBVLPDU_Serialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBVLPDU_Serialize_Call) RunAndReturn(run func() ([]byte, error)) *MockBVLPDU_Serialize_Call {
	_c.Call.Return(run)
	return _c
}

// SerializeWithWriteBuffer provides a mock function with given fields: ctx, writeBuffer
func (_m *MockBVLPDU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	ret := _m.Called(ctx, writeBuffer)

	if len(ret) == 0 {
		panic("no return value specified for SerializeWithWriteBuffer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, utils.WriteBuffer) error); ok {
		r0 = rf(ctx, writeBuffer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBVLPDU_SerializeWithWriteBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SerializeWithWriteBuffer'
type MockBVLPDU_SerializeWithWriteBuffer_Call struct {
	*mock.Call
}

// SerializeWithWriteBuffer is a helper method to define mock.On call
//   - ctx context.Context
//   - writeBuffer utils.WriteBuffer
func (_e *MockBVLPDU_Expecter) SerializeWithWriteBuffer(ctx interface{}, writeBuffer interface{}) *MockBVLPDU_SerializeWithWriteBuffer_Call {
	return &MockBVLPDU_SerializeWithWriteBuffer_Call{Call: _e.mock.On("SerializeWithWriteBuffer", ctx, writeBuffer)}
}

func (_c *MockBVLPDU_SerializeWithWriteBuffer_Call) Run(run func(ctx context.Context, writeBuffer utils.WriteBuffer)) *MockBVLPDU_SerializeWithWriteBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(utils.WriteBuffer))
	})
	return _c
}

func (_c *MockBVLPDU_SerializeWithWriteBuffer_Call) Return(_a0 error) *MockBVLPDU_SerializeWithWriteBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_SerializeWithWriteBuffer_Call) RunAndReturn(run func(context.Context, utils.WriteBuffer) error) *MockBVLPDU_SerializeWithWriteBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUDestination provides a mock function with given fields: _a0
func (_m *MockBVLPDU) SetPDUDestination(_a0 *Address) {
	_m.Called(_a0)
}

// MockBVLPDU_SetPDUDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUDestination'
type MockBVLPDU_SetPDUDestination_Call struct {
	*mock.Call
}

// SetPDUDestination is a helper method to define mock.On call
//   - _a0 *Address
func (_e *MockBVLPDU_Expecter) SetPDUDestination(_a0 interface{}) *MockBVLPDU_SetPDUDestination_Call {
	return &MockBVLPDU_SetPDUDestination_Call{Call: _e.mock.On("SetPDUDestination", _a0)}
}

func (_c *MockBVLPDU_SetPDUDestination_Call) Run(run func(_a0 *Address)) *MockBVLPDU_SetPDUDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Address))
	})
	return _c
}

func (_c *MockBVLPDU_SetPDUDestination_Call) Return() *MockBVLPDU_SetPDUDestination_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBVLPDU_SetPDUDestination_Call) RunAndReturn(run func(*Address)) *MockBVLPDU_SetPDUDestination_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUSource provides a mock function with given fields: source
func (_m *MockBVLPDU) SetPDUSource(source *Address) {
	_m.Called(source)
}

// MockBVLPDU_SetPDUSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUSource'
type MockBVLPDU_SetPDUSource_Call struct {
	*mock.Call
}

// SetPDUSource is a helper method to define mock.On call
//   - source *Address
func (_e *MockBVLPDU_Expecter) SetPDUSource(source interface{}) *MockBVLPDU_SetPDUSource_Call {
	return &MockBVLPDU_SetPDUSource_Call{Call: _e.mock.On("SetPDUSource", source)}
}

func (_c *MockBVLPDU_SetPDUSource_Call) Run(run func(source *Address)) *MockBVLPDU_SetPDUSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Address))
	})
	return _c
}

func (_c *MockBVLPDU_SetPDUSource_Call) Return() *MockBVLPDU_SetPDUSource_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBVLPDU_SetPDUSource_Call) RunAndReturn(run func(*Address)) *MockBVLPDU_SetPDUSource_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUUserData provides a mock function with given fields: _a0
func (_m *MockBVLPDU) SetPDUUserData(_a0 spi.Message) {
	_m.Called(_a0)
}

// MockBVLPDU_SetPDUUserData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUUserData'
type MockBVLPDU_SetPDUUserData_Call struct {
	*mock.Call
}

// SetPDUUserData is a helper method to define mock.On call
//   - _a0 spi.Message
func (_e *MockBVLPDU_Expecter) SetPDUUserData(_a0 interface{}) *MockBVLPDU_SetPDUUserData_Call {
	return &MockBVLPDU_SetPDUUserData_Call{Call: _e.mock.On("SetPDUUserData", _a0)}
}

func (_c *MockBVLPDU_SetPDUUserData_Call) Run(run func(_a0 spi.Message)) *MockBVLPDU_SetPDUUserData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spi.Message))
	})
	return _c
}

func (_c *MockBVLPDU_SetPDUUserData_Call) Return() *MockBVLPDU_SetPDUUserData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBVLPDU_SetPDUUserData_Call) RunAndReturn(run func(spi.Message)) *MockBVLPDU_SetPDUUserData_Call {
	_c.Call.Return(run)
	return _c
}

// SetPduData provides a mock function with given fields: _a0
func (_m *MockBVLPDU) SetPduData(_a0 []byte) {
	_m.Called(_a0)
}

// MockBVLPDU_SetPduData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPduData'
type MockBVLPDU_SetPduData_Call struct {
	*mock.Call
}

// SetPduData is a helper method to define mock.On call
//   - _a0 []byte
func (_e *MockBVLPDU_Expecter) SetPduData(_a0 interface{}) *MockBVLPDU_SetPduData_Call {
	return &MockBVLPDU_SetPduData_Call{Call: _e.mock.On("SetPduData", _a0)}
}

func (_c *MockBVLPDU_SetPduData_Call) Run(run func(_a0 []byte)) *MockBVLPDU_SetPduData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockBVLPDU_SetPduData_Call) Return() *MockBVLPDU_SetPduData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBVLPDU_SetPduData_Call) RunAndReturn(run func([]byte)) *MockBVLPDU_SetPduData_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockBVLPDU) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockBVLPDU_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockBVLPDU_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) String() *MockBVLPDU_String_Call {
	return &MockBVLPDU_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockBVLPDU_String_Call) Run(run func()) *MockBVLPDU_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_String_Call) Return(_a0 string) *MockBVLPDU_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_String_Call) RunAndReturn(run func() string) *MockBVLPDU_String_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: pci
func (_m *MockBVLPDU) Update(pci Arg) error {
	ret := _m.Called(pci)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Arg) error); ok {
		r0 = rf(pci)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBVLPDU_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockBVLPDU_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - pci Arg
func (_e *MockBVLPDU_Expecter) Update(pci interface{}) *MockBVLPDU_Update_Call {
	return &MockBVLPDU_Update_Call{Call: _e.mock.On("Update", pci)}
}

func (_c *MockBVLPDU_Update_Call) Run(run func(pci Arg)) *MockBVLPDU_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Arg))
	})
	return _c
}

func (_c *MockBVLPDU_Update_Call) Return(_a0 error) *MockBVLPDU_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_Update_Call) RunAndReturn(run func(Arg) error) *MockBVLPDU_Update_Call {
	_c.Call.Return(run)
	return _c
}

// getBVLC provides a mock function with given fields:
func (_m *MockBVLPDU) getBVLC() model.BVLC {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getBVLC")
	}

	var r0 model.BVLC
	if rf, ok := ret.Get(0).(func() model.BVLC); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.BVLC)
		}
	}

	return r0
}

// MockBVLPDU_getBVLC_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getBVLC'
type MockBVLPDU_getBVLC_Call struct {
	*mock.Call
}

// getBVLC is a helper method to define mock.On call
func (_e *MockBVLPDU_Expecter) getBVLC() *MockBVLPDU_getBVLC_Call {
	return &MockBVLPDU_getBVLC_Call{Call: _e.mock.On("getBVLC")}
}

func (_c *MockBVLPDU_getBVLC_Call) Run(run func()) *MockBVLPDU_getBVLC_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBVLPDU_getBVLC_Call) Return(_a0 model.BVLC) *MockBVLPDU_getBVLC_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBVLPDU_getBVLC_Call) RunAndReturn(run func() model.BVLC) *MockBVLPDU_getBVLC_Call {
	_c.Call.Return(run)
	return _c
}

// setBVLC provides a mock function with given fields: _a0
func (_m *MockBVLPDU) setBVLC(_a0 model.BVLC) {
	_m.Called(_a0)
}

// MockBVLPDU_setBVLC_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setBVLC'
type MockBVLPDU_setBVLC_Call struct {
	*mock.Call
}

// setBVLC is a helper method to define mock.On call
//   - _a0 model.BVLC
func (_e *MockBVLPDU_Expecter) setBVLC(_a0 interface{}) *MockBVLPDU_setBVLC_Call {
	return &MockBVLPDU_setBVLC_Call{Call: _e.mock.On("setBVLC", _a0)}
}

func (_c *MockBVLPDU_setBVLC_Call) Run(run func(_a0 model.BVLC)) *MockBVLPDU_setBVLC_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.BVLC))
	})
	return _c
}

func (_c *MockBVLPDU_setBVLC_Call) Return() *MockBVLPDU_setBVLC_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBVLPDU_setBVLC_Call) RunAndReturn(run func(model.BVLC)) *MockBVLPDU_setBVLC_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBVLPDU creates a new instance of MockBVLPDU. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBVLPDU(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBVLPDU {
	mock := &MockBVLPDU{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
