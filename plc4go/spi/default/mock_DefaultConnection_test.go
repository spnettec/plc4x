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

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	plc4go "github.com/apache/plc4x/plc4go/pkg/api"

	spi "github.com/apache/plc4x/plc4go/spi"

	time "time"

	transports "github.com/apache/plc4x/plc4go/spi/transports"

	utils "github.com/apache/plc4x/plc4go/spi/utils"
)

// MockDefaultConnection is an autogenerated mock type for the DefaultConnection type
type MockDefaultConnection struct {
	mock.Mock
}

type MockDefaultConnection_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDefaultConnection) EXPECT() *MockDefaultConnection_Expecter {
	return &MockDefaultConnection_Expecter{mock: &_m.Mock}
}

// BlockingClose provides a mock function with given fields:
func (_m *MockDefaultConnection) BlockingClose() {
	_m.Called()
}

// MockDefaultConnection_BlockingClose_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockingClose'
type MockDefaultConnection_BlockingClose_Call struct {
	*mock.Call
}

// BlockingClose is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) BlockingClose() *MockDefaultConnection_BlockingClose_Call {
	return &MockDefaultConnection_BlockingClose_Call{Call: _e.mock.On("BlockingClose")}
}

func (_c *MockDefaultConnection_BlockingClose_Call) Run(run func()) *MockDefaultConnection_BlockingClose_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_BlockingClose_Call) Return() *MockDefaultConnection_BlockingClose_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockDefaultConnection_BlockingClose_Call) RunAndReturn(run func()) *MockDefaultConnection_BlockingClose_Call {
	_c.Call.Return(run)
	return _c
}

// BrowseRequestBuilder provides a mock function with given fields:
func (_m *MockDefaultConnection) BrowseRequestBuilder() model.PlcBrowseRequestBuilder {
	ret := _m.Called()

	var r0 model.PlcBrowseRequestBuilder
	if rf, ok := ret.Get(0).(func() model.PlcBrowseRequestBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcBrowseRequestBuilder)
		}
	}

	return r0
}

// MockDefaultConnection_BrowseRequestBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BrowseRequestBuilder'
type MockDefaultConnection_BrowseRequestBuilder_Call struct {
	*mock.Call
}

// BrowseRequestBuilder is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) BrowseRequestBuilder() *MockDefaultConnection_BrowseRequestBuilder_Call {
	return &MockDefaultConnection_BrowseRequestBuilder_Call{Call: _e.mock.On("BrowseRequestBuilder")}
}

func (_c *MockDefaultConnection_BrowseRequestBuilder_Call) Run(run func()) *MockDefaultConnection_BrowseRequestBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_BrowseRequestBuilder_Call) Return(_a0 model.PlcBrowseRequestBuilder) *MockDefaultConnection_BrowseRequestBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_BrowseRequestBuilder_Call) RunAndReturn(run func() model.PlcBrowseRequestBuilder) *MockDefaultConnection_BrowseRequestBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockDefaultConnection) Close() <-chan plc4go.PlcConnectionCloseResult {
	ret := _m.Called()

	var r0 <-chan plc4go.PlcConnectionCloseResult
	if rf, ok := ret.Get(0).(func() <-chan plc4go.PlcConnectionCloseResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan plc4go.PlcConnectionCloseResult)
		}
	}

	return r0
}

// MockDefaultConnection_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockDefaultConnection_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) Close() *MockDefaultConnection_Close_Call {
	return &MockDefaultConnection_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockDefaultConnection_Close_Call) Run(run func()) *MockDefaultConnection_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_Close_Call) Return(_a0 <-chan plc4go.PlcConnectionCloseResult) *MockDefaultConnection_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_Close_Call) RunAndReturn(run func() <-chan plc4go.PlcConnectionCloseResult) *MockDefaultConnection_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Connect provides a mock function with given fields:
func (_m *MockDefaultConnection) Connect() <-chan plc4go.PlcConnectionConnectResult {
	ret := _m.Called()

	var r0 <-chan plc4go.PlcConnectionConnectResult
	if rf, ok := ret.Get(0).(func() <-chan plc4go.PlcConnectionConnectResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan plc4go.PlcConnectionConnectResult)
		}
	}

	return r0
}

// MockDefaultConnection_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockDefaultConnection_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) Connect() *MockDefaultConnection_Connect_Call {
	return &MockDefaultConnection_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockDefaultConnection_Connect_Call) Run(run func()) *MockDefaultConnection_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_Connect_Call) Return(_a0 <-chan plc4go.PlcConnectionConnectResult) *MockDefaultConnection_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_Connect_Call) RunAndReturn(run func() <-chan plc4go.PlcConnectionConnectResult) *MockDefaultConnection_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// ConnectWithContext provides a mock function with given fields: ctx
func (_m *MockDefaultConnection) ConnectWithContext(ctx context.Context) <-chan plc4go.PlcConnectionConnectResult {
	ret := _m.Called(ctx)

	var r0 <-chan plc4go.PlcConnectionConnectResult
	if rf, ok := ret.Get(0).(func(context.Context) <-chan plc4go.PlcConnectionConnectResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan plc4go.PlcConnectionConnectResult)
		}
	}

	return r0
}

// MockDefaultConnection_ConnectWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectWithContext'
type MockDefaultConnection_ConnectWithContext_Call struct {
	*mock.Call
}

// ConnectWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockDefaultConnection_Expecter) ConnectWithContext(ctx interface{}) *MockDefaultConnection_ConnectWithContext_Call {
	return &MockDefaultConnection_ConnectWithContext_Call{Call: _e.mock.On("ConnectWithContext", ctx)}
}

func (_c *MockDefaultConnection_ConnectWithContext_Call) Run(run func(ctx context.Context)) *MockDefaultConnection_ConnectWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockDefaultConnection_ConnectWithContext_Call) Return(_a0 <-chan plc4go.PlcConnectionConnectResult) *MockDefaultConnection_ConnectWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_ConnectWithContext_Call) RunAndReturn(run func(context.Context) <-chan plc4go.PlcConnectionConnectResult) *MockDefaultConnection_ConnectWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetMetadata provides a mock function with given fields:
func (_m *MockDefaultConnection) GetMetadata() model.PlcConnectionMetadata {
	ret := _m.Called()

	var r0 model.PlcConnectionMetadata
	if rf, ok := ret.Get(0).(func() model.PlcConnectionMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcConnectionMetadata)
		}
	}

	return r0
}

// MockDefaultConnection_GetMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetadata'
type MockDefaultConnection_GetMetadata_Call struct {
	*mock.Call
}

// GetMetadata is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) GetMetadata() *MockDefaultConnection_GetMetadata_Call {
	return &MockDefaultConnection_GetMetadata_Call{Call: _e.mock.On("GetMetadata")}
}

func (_c *MockDefaultConnection_GetMetadata_Call) Run(run func()) *MockDefaultConnection_GetMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_GetMetadata_Call) Return(_a0 model.PlcConnectionMetadata) *MockDefaultConnection_GetMetadata_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_GetMetadata_Call) RunAndReturn(run func() model.PlcConnectionMetadata) *MockDefaultConnection_GetMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlcTagHandler provides a mock function with given fields:
func (_m *MockDefaultConnection) GetPlcTagHandler() spi.PlcTagHandler {
	ret := _m.Called()

	var r0 spi.PlcTagHandler
	if rf, ok := ret.Get(0).(func() spi.PlcTagHandler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.PlcTagHandler)
		}
	}

	return r0
}

// MockDefaultConnection_GetPlcTagHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlcTagHandler'
type MockDefaultConnection_GetPlcTagHandler_Call struct {
	*mock.Call
}

// GetPlcTagHandler is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) GetPlcTagHandler() *MockDefaultConnection_GetPlcTagHandler_Call {
	return &MockDefaultConnection_GetPlcTagHandler_Call{Call: _e.mock.On("GetPlcTagHandler")}
}

func (_c *MockDefaultConnection_GetPlcTagHandler_Call) Run(run func()) *MockDefaultConnection_GetPlcTagHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_GetPlcTagHandler_Call) Return(_a0 spi.PlcTagHandler) *MockDefaultConnection_GetPlcTagHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_GetPlcTagHandler_Call) RunAndReturn(run func() spi.PlcTagHandler) *MockDefaultConnection_GetPlcTagHandler_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlcValueHandler provides a mock function with given fields:
func (_m *MockDefaultConnection) GetPlcValueHandler() spi.PlcValueHandler {
	ret := _m.Called()

	var r0 spi.PlcValueHandler
	if rf, ok := ret.Get(0).(func() spi.PlcValueHandler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.PlcValueHandler)
		}
	}

	return r0
}

// MockDefaultConnection_GetPlcValueHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlcValueHandler'
type MockDefaultConnection_GetPlcValueHandler_Call struct {
	*mock.Call
}

// GetPlcValueHandler is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) GetPlcValueHandler() *MockDefaultConnection_GetPlcValueHandler_Call {
	return &MockDefaultConnection_GetPlcValueHandler_Call{Call: _e.mock.On("GetPlcValueHandler")}
}

func (_c *MockDefaultConnection_GetPlcValueHandler_Call) Run(run func()) *MockDefaultConnection_GetPlcValueHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_GetPlcValueHandler_Call) Return(_a0 spi.PlcValueHandler) *MockDefaultConnection_GetPlcValueHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_GetPlcValueHandler_Call) RunAndReturn(run func() spi.PlcValueHandler) *MockDefaultConnection_GetPlcValueHandler_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransportInstance provides a mock function with given fields:
func (_m *MockDefaultConnection) GetTransportInstance() transports.TransportInstance {
	ret := _m.Called()

	var r0 transports.TransportInstance
	if rf, ok := ret.Get(0).(func() transports.TransportInstance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transports.TransportInstance)
		}
	}

	return r0
}

// MockDefaultConnection_GetTransportInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransportInstance'
type MockDefaultConnection_GetTransportInstance_Call struct {
	*mock.Call
}

// GetTransportInstance is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) GetTransportInstance() *MockDefaultConnection_GetTransportInstance_Call {
	return &MockDefaultConnection_GetTransportInstance_Call{Call: _e.mock.On("GetTransportInstance")}
}

func (_c *MockDefaultConnection_GetTransportInstance_Call) Run(run func()) *MockDefaultConnection_GetTransportInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_GetTransportInstance_Call) Return(_a0 transports.TransportInstance) *MockDefaultConnection_GetTransportInstance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_GetTransportInstance_Call) RunAndReturn(run func() transports.TransportInstance) *MockDefaultConnection_GetTransportInstance_Call {
	_c.Call.Return(run)
	return _c
}

// GetTtl provides a mock function with given fields:
func (_m *MockDefaultConnection) GetTtl() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// MockDefaultConnection_GetTtl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTtl'
type MockDefaultConnection_GetTtl_Call struct {
	*mock.Call
}

// GetTtl is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) GetTtl() *MockDefaultConnection_GetTtl_Call {
	return &MockDefaultConnection_GetTtl_Call{Call: _e.mock.On("GetTtl")}
}

func (_c *MockDefaultConnection_GetTtl_Call) Run(run func()) *MockDefaultConnection_GetTtl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_GetTtl_Call) Return(_a0 time.Duration) *MockDefaultConnection_GetTtl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_GetTtl_Call) RunAndReturn(run func() time.Duration) *MockDefaultConnection_GetTtl_Call {
	_c.Call.Return(run)
	return _c
}

// IsConnected provides a mock function with given fields:
func (_m *MockDefaultConnection) IsConnected() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockDefaultConnection_IsConnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsConnected'
type MockDefaultConnection_IsConnected_Call struct {
	*mock.Call
}

// IsConnected is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) IsConnected() *MockDefaultConnection_IsConnected_Call {
	return &MockDefaultConnection_IsConnected_Call{Call: _e.mock.On("IsConnected")}
}

func (_c *MockDefaultConnection_IsConnected_Call) Run(run func()) *MockDefaultConnection_IsConnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_IsConnected_Call) Return(_a0 bool) *MockDefaultConnection_IsConnected_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_IsConnected_Call) RunAndReturn(run func() bool) *MockDefaultConnection_IsConnected_Call {
	_c.Call.Return(run)
	return _c
}

// Ping provides a mock function with given fields:
func (_m *MockDefaultConnection) Ping() <-chan plc4go.PlcConnectionPingResult {
	ret := _m.Called()

	var r0 <-chan plc4go.PlcConnectionPingResult
	if rf, ok := ret.Get(0).(func() <-chan plc4go.PlcConnectionPingResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan plc4go.PlcConnectionPingResult)
		}
	}

	return r0
}

// MockDefaultConnection_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type MockDefaultConnection_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) Ping() *MockDefaultConnection_Ping_Call {
	return &MockDefaultConnection_Ping_Call{Call: _e.mock.On("Ping")}
}

func (_c *MockDefaultConnection_Ping_Call) Run(run func()) *MockDefaultConnection_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_Ping_Call) Return(_a0 <-chan plc4go.PlcConnectionPingResult) *MockDefaultConnection_Ping_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_Ping_Call) RunAndReturn(run func() <-chan plc4go.PlcConnectionPingResult) *MockDefaultConnection_Ping_Call {
	_c.Call.Return(run)
	return _c
}

// ReadRequestBuilder provides a mock function with given fields:
func (_m *MockDefaultConnection) ReadRequestBuilder() model.PlcReadRequestBuilder {
	ret := _m.Called()

	var r0 model.PlcReadRequestBuilder
	if rf, ok := ret.Get(0).(func() model.PlcReadRequestBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcReadRequestBuilder)
		}
	}

	return r0
}

// MockDefaultConnection_ReadRequestBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadRequestBuilder'
type MockDefaultConnection_ReadRequestBuilder_Call struct {
	*mock.Call
}

// ReadRequestBuilder is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) ReadRequestBuilder() *MockDefaultConnection_ReadRequestBuilder_Call {
	return &MockDefaultConnection_ReadRequestBuilder_Call{Call: _e.mock.On("ReadRequestBuilder")}
}

func (_c *MockDefaultConnection_ReadRequestBuilder_Call) Run(run func()) *MockDefaultConnection_ReadRequestBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_ReadRequestBuilder_Call) Return(_a0 model.PlcReadRequestBuilder) *MockDefaultConnection_ReadRequestBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_ReadRequestBuilder_Call) RunAndReturn(run func() model.PlcReadRequestBuilder) *MockDefaultConnection_ReadRequestBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// Serialize provides a mock function with given fields:
func (_m *MockDefaultConnection) Serialize() ([]byte, error) {
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

// MockDefaultConnection_Serialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serialize'
type MockDefaultConnection_Serialize_Call struct {
	*mock.Call
}

// Serialize is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) Serialize() *MockDefaultConnection_Serialize_Call {
	return &MockDefaultConnection_Serialize_Call{Call: _e.mock.On("Serialize")}
}

func (_c *MockDefaultConnection_Serialize_Call) Run(run func()) *MockDefaultConnection_Serialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_Serialize_Call) Return(_a0 []byte, _a1 error) *MockDefaultConnection_Serialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDefaultConnection_Serialize_Call) RunAndReturn(run func() ([]byte, error)) *MockDefaultConnection_Serialize_Call {
	_c.Call.Return(run)
	return _c
}

// SerializeWithWriteBuffer provides a mock function with given fields: ctx, writeBuffer
func (_m *MockDefaultConnection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	ret := _m.Called(ctx, writeBuffer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, utils.WriteBuffer) error); ok {
		r0 = rf(ctx, writeBuffer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultConnection_SerializeWithWriteBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SerializeWithWriteBuffer'
type MockDefaultConnection_SerializeWithWriteBuffer_Call struct {
	*mock.Call
}

// SerializeWithWriteBuffer is a helper method to define mock.On call
//   - ctx context.Context
//   - writeBuffer utils.WriteBuffer
func (_e *MockDefaultConnection_Expecter) SerializeWithWriteBuffer(ctx interface{}, writeBuffer interface{}) *MockDefaultConnection_SerializeWithWriteBuffer_Call {
	return &MockDefaultConnection_SerializeWithWriteBuffer_Call{Call: _e.mock.On("SerializeWithWriteBuffer", ctx, writeBuffer)}
}

func (_c *MockDefaultConnection_SerializeWithWriteBuffer_Call) Run(run func(ctx context.Context, writeBuffer utils.WriteBuffer)) *MockDefaultConnection_SerializeWithWriteBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(utils.WriteBuffer))
	})
	return _c
}

func (_c *MockDefaultConnection_SerializeWithWriteBuffer_Call) Return(_a0 error) *MockDefaultConnection_SerializeWithWriteBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_SerializeWithWriteBuffer_Call) RunAndReturn(run func(context.Context, utils.WriteBuffer) error) *MockDefaultConnection_SerializeWithWriteBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// SetConnected provides a mock function with given fields: connected
func (_m *MockDefaultConnection) SetConnected(connected bool) {
	_m.Called(connected)
}

// MockDefaultConnection_SetConnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConnected'
type MockDefaultConnection_SetConnected_Call struct {
	*mock.Call
}

// SetConnected is a helper method to define mock.On call
//   - connected bool
func (_e *MockDefaultConnection_Expecter) SetConnected(connected interface{}) *MockDefaultConnection_SetConnected_Call {
	return &MockDefaultConnection_SetConnected_Call{Call: _e.mock.On("SetConnected", connected)}
}

func (_c *MockDefaultConnection_SetConnected_Call) Run(run func(connected bool)) *MockDefaultConnection_SetConnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockDefaultConnection_SetConnected_Call) Return() *MockDefaultConnection_SetConnected_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockDefaultConnection_SetConnected_Call) RunAndReturn(run func(bool)) *MockDefaultConnection_SetConnected_Call {
	_c.Call.Return(run)
	return _c
}

// SubscriptionRequestBuilder provides a mock function with given fields:
func (_m *MockDefaultConnection) SubscriptionRequestBuilder() model.PlcSubscriptionRequestBuilder {
	ret := _m.Called()

	var r0 model.PlcSubscriptionRequestBuilder
	if rf, ok := ret.Get(0).(func() model.PlcSubscriptionRequestBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcSubscriptionRequestBuilder)
		}
	}

	return r0
}

// MockDefaultConnection_SubscriptionRequestBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscriptionRequestBuilder'
type MockDefaultConnection_SubscriptionRequestBuilder_Call struct {
	*mock.Call
}

// SubscriptionRequestBuilder is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) SubscriptionRequestBuilder() *MockDefaultConnection_SubscriptionRequestBuilder_Call {
	return &MockDefaultConnection_SubscriptionRequestBuilder_Call{Call: _e.mock.On("SubscriptionRequestBuilder")}
}

func (_c *MockDefaultConnection_SubscriptionRequestBuilder_Call) Run(run func()) *MockDefaultConnection_SubscriptionRequestBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_SubscriptionRequestBuilder_Call) Return(_a0 model.PlcSubscriptionRequestBuilder) *MockDefaultConnection_SubscriptionRequestBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_SubscriptionRequestBuilder_Call) RunAndReturn(run func() model.PlcSubscriptionRequestBuilder) *MockDefaultConnection_SubscriptionRequestBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// UnsubscriptionRequestBuilder provides a mock function with given fields:
func (_m *MockDefaultConnection) UnsubscriptionRequestBuilder() model.PlcUnsubscriptionRequestBuilder {
	ret := _m.Called()

	var r0 model.PlcUnsubscriptionRequestBuilder
	if rf, ok := ret.Get(0).(func() model.PlcUnsubscriptionRequestBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcUnsubscriptionRequestBuilder)
		}
	}

	return r0
}

// MockDefaultConnection_UnsubscriptionRequestBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnsubscriptionRequestBuilder'
type MockDefaultConnection_UnsubscriptionRequestBuilder_Call struct {
	*mock.Call
}

// UnsubscriptionRequestBuilder is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) UnsubscriptionRequestBuilder() *MockDefaultConnection_UnsubscriptionRequestBuilder_Call {
	return &MockDefaultConnection_UnsubscriptionRequestBuilder_Call{Call: _e.mock.On("UnsubscriptionRequestBuilder")}
}

func (_c *MockDefaultConnection_UnsubscriptionRequestBuilder_Call) Run(run func()) *MockDefaultConnection_UnsubscriptionRequestBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_UnsubscriptionRequestBuilder_Call) Return(_a0 model.PlcUnsubscriptionRequestBuilder) *MockDefaultConnection_UnsubscriptionRequestBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_UnsubscriptionRequestBuilder_Call) RunAndReturn(run func() model.PlcUnsubscriptionRequestBuilder) *MockDefaultConnection_UnsubscriptionRequestBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// WriteRequestBuilder provides a mock function with given fields:
func (_m *MockDefaultConnection) WriteRequestBuilder() model.PlcWriteRequestBuilder {
	ret := _m.Called()

	var r0 model.PlcWriteRequestBuilder
	if rf, ok := ret.Get(0).(func() model.PlcWriteRequestBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcWriteRequestBuilder)
		}
	}

	return r0
}

// MockDefaultConnection_WriteRequestBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteRequestBuilder'
type MockDefaultConnection_WriteRequestBuilder_Call struct {
	*mock.Call
}

// WriteRequestBuilder is a helper method to define mock.On call
func (_e *MockDefaultConnection_Expecter) WriteRequestBuilder() *MockDefaultConnection_WriteRequestBuilder_Call {
	return &MockDefaultConnection_WriteRequestBuilder_Call{Call: _e.mock.On("WriteRequestBuilder")}
}

func (_c *MockDefaultConnection_WriteRequestBuilder_Call) Run(run func()) *MockDefaultConnection_WriteRequestBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnection_WriteRequestBuilder_Call) Return(_a0 model.PlcWriteRequestBuilder) *MockDefaultConnection_WriteRequestBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnection_WriteRequestBuilder_Call) RunAndReturn(run func() model.PlcWriteRequestBuilder) *MockDefaultConnection_WriteRequestBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDefaultConnection creates a new instance of MockDefaultConnection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDefaultConnection(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDefaultConnection {
	mock := &MockDefaultConnection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
