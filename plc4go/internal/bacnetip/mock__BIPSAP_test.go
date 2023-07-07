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

package bacnetip

import mock "github.com/stretchr/testify/mock"

// mock_BIPSAP is an autogenerated mock type for the _BIPSAP type
type mock_BIPSAP struct {
	mock.Mock
}

type mock_BIPSAP_Expecter struct {
	mock *mock.Mock
}

func (_m *mock_BIPSAP) EXPECT() *mock_BIPSAP_Expecter {
	return &mock_BIPSAP_Expecter{mock: &_m.Mock}
}

// Confirmation provides a mock function with given fields: pdu
func (_m *mock_BIPSAP) Confirmation(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_BIPSAP_Confirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Confirmation'
type mock_BIPSAP_Confirmation_Call struct {
	*mock.Call
}

// Confirmation is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_BIPSAP_Expecter) Confirmation(pdu interface{}) *mock_BIPSAP_Confirmation_Call {
	return &mock_BIPSAP_Confirmation_Call{Call: _e.mock.On("Confirmation", pdu)}
}

func (_c *mock_BIPSAP_Confirmation_Call) Run(run func(pdu _PDU)) *mock_BIPSAP_Confirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_BIPSAP_Confirmation_Call) Return(_a0 error) *mock_BIPSAP_Confirmation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_BIPSAP_Confirmation_Call) RunAndReturn(run func(_PDU) error) *mock_BIPSAP_Confirmation_Call {
	_c.Call.Return(run)
	return _c
}

// Request provides a mock function with given fields: pdu
func (_m *mock_BIPSAP) Request(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_BIPSAP_Request_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Request'
type mock_BIPSAP_Request_Call struct {
	*mock.Call
}

// Request is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_BIPSAP_Expecter) Request(pdu interface{}) *mock_BIPSAP_Request_Call {
	return &mock_BIPSAP_Request_Call{Call: _e.mock.On("Request", pdu)}
}

func (_c *mock_BIPSAP_Request_Call) Run(run func(pdu _PDU)) *mock_BIPSAP_Request_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_BIPSAP_Request_Call) Return(_a0 error) *mock_BIPSAP_Request_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_BIPSAP_Request_Call) RunAndReturn(run func(_PDU) error) *mock_BIPSAP_Request_Call {
	_c.Call.Return(run)
	return _c
}

// SapConfirmation provides a mock function with given fields: pdu
func (_m *mock_BIPSAP) SapConfirmation(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_BIPSAP_SapConfirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapConfirmation'
type mock_BIPSAP_SapConfirmation_Call struct {
	*mock.Call
}

// SapConfirmation is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_BIPSAP_Expecter) SapConfirmation(pdu interface{}) *mock_BIPSAP_SapConfirmation_Call {
	return &mock_BIPSAP_SapConfirmation_Call{Call: _e.mock.On("SapConfirmation", pdu)}
}

func (_c *mock_BIPSAP_SapConfirmation_Call) Run(run func(pdu _PDU)) *mock_BIPSAP_SapConfirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_BIPSAP_SapConfirmation_Call) Return(_a0 error) *mock_BIPSAP_SapConfirmation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_BIPSAP_SapConfirmation_Call) RunAndReturn(run func(_PDU) error) *mock_BIPSAP_SapConfirmation_Call {
	_c.Call.Return(run)
	return _c
}

// SapIndication provides a mock function with given fields: pdu
func (_m *mock_BIPSAP) SapIndication(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_BIPSAP_SapIndication_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapIndication'
type mock_BIPSAP_SapIndication_Call struct {
	*mock.Call
}

// SapIndication is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_BIPSAP_Expecter) SapIndication(pdu interface{}) *mock_BIPSAP_SapIndication_Call {
	return &mock_BIPSAP_SapIndication_Call{Call: _e.mock.On("SapIndication", pdu)}
}

func (_c *mock_BIPSAP_SapIndication_Call) Run(run func(pdu _PDU)) *mock_BIPSAP_SapIndication_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_BIPSAP_SapIndication_Call) Return(_a0 error) *mock_BIPSAP_SapIndication_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_BIPSAP_SapIndication_Call) RunAndReturn(run func(_PDU) error) *mock_BIPSAP_SapIndication_Call {
	_c.Call.Return(run)
	return _c
}

// SapRequest provides a mock function with given fields: pdu
func (_m *mock_BIPSAP) SapRequest(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_BIPSAP_SapRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapRequest'
type mock_BIPSAP_SapRequest_Call struct {
	*mock.Call
}

// SapRequest is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_BIPSAP_Expecter) SapRequest(pdu interface{}) *mock_BIPSAP_SapRequest_Call {
	return &mock_BIPSAP_SapRequest_Call{Call: _e.mock.On("SapRequest", pdu)}
}

func (_c *mock_BIPSAP_SapRequest_Call) Run(run func(pdu _PDU)) *mock_BIPSAP_SapRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_BIPSAP_SapRequest_Call) Return(_a0 error) *mock_BIPSAP_SapRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_BIPSAP_SapRequest_Call) RunAndReturn(run func(_PDU) error) *mock_BIPSAP_SapRequest_Call {
	_c.Call.Return(run)
	return _c
}

// SapResponse provides a mock function with given fields: pdu
func (_m *mock_BIPSAP) SapResponse(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_BIPSAP_SapResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapResponse'
type mock_BIPSAP_SapResponse_Call struct {
	*mock.Call
}

// SapResponse is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_BIPSAP_Expecter) SapResponse(pdu interface{}) *mock_BIPSAP_SapResponse_Call {
	return &mock_BIPSAP_SapResponse_Call{Call: _e.mock.On("SapResponse", pdu)}
}

func (_c *mock_BIPSAP_SapResponse_Call) Run(run func(pdu _PDU)) *mock_BIPSAP_SapResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_BIPSAP_SapResponse_Call) Return(_a0 error) *mock_BIPSAP_SapResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_BIPSAP_SapResponse_Call) RunAndReturn(run func(_PDU) error) *mock_BIPSAP_SapResponse_Call {
	_c.Call.Return(run)
	return _c
}

// _setClientPeer provides a mock function with given fields: server
func (_m *mock_BIPSAP) _setClientPeer(server _Server) {
	_m.Called(server)
}

// mock_BIPSAP__setClientPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method '_setClientPeer'
type mock_BIPSAP__setClientPeer_Call struct {
	*mock.Call
}

// _setClientPeer is a helper method to define mock.On call
//   - server _Server
func (_e *mock_BIPSAP_Expecter) _setClientPeer(server interface{}) *mock_BIPSAP__setClientPeer_Call {
	return &mock_BIPSAP__setClientPeer_Call{Call: _e.mock.On("_setClientPeer", server)}
}

func (_c *mock_BIPSAP__setClientPeer_Call) Run(run func(server _Server)) *mock_BIPSAP__setClientPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_Server))
	})
	return _c
}

func (_c *mock_BIPSAP__setClientPeer_Call) Return() *mock_BIPSAP__setClientPeer_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_BIPSAP__setClientPeer_Call) RunAndReturn(run func(_Server)) *mock_BIPSAP__setClientPeer_Call {
	_c.Call.Return(run)
	return _c
}

// _setServiceElement provides a mock function with given fields: serviceElement
func (_m *mock_BIPSAP) _setServiceElement(serviceElement _ApplicationServiceElement) {
	_m.Called(serviceElement)
}

// mock_BIPSAP__setServiceElement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method '_setServiceElement'
type mock_BIPSAP__setServiceElement_Call struct {
	*mock.Call
}

// _setServiceElement is a helper method to define mock.On call
//   - serviceElement _ApplicationServiceElement
func (_e *mock_BIPSAP_Expecter) _setServiceElement(serviceElement interface{}) *mock_BIPSAP__setServiceElement_Call {
	return &mock_BIPSAP__setServiceElement_Call{Call: _e.mock.On("_setServiceElement", serviceElement)}
}

func (_c *mock_BIPSAP__setServiceElement_Call) Run(run func(serviceElement _ApplicationServiceElement)) *mock_BIPSAP__setServiceElement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_ApplicationServiceElement))
	})
	return _c
}

func (_c *mock_BIPSAP__setServiceElement_Call) Return() *mock_BIPSAP__setServiceElement_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_BIPSAP__setServiceElement_Call) RunAndReturn(run func(_ApplicationServiceElement)) *mock_BIPSAP__setServiceElement_Call {
	_c.Call.Return(run)
	return _c
}

// getClientId provides a mock function with given fields:
func (_m *mock_BIPSAP) getClientId() *int {
	ret := _m.Called()

	var r0 *int
	if rf, ok := ret.Get(0).(func() *int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	return r0
}

// mock_BIPSAP_getClientId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getClientId'
type mock_BIPSAP_getClientId_Call struct {
	*mock.Call
}

// getClientId is a helper method to define mock.On call
func (_e *mock_BIPSAP_Expecter) getClientId() *mock_BIPSAP_getClientId_Call {
	return &mock_BIPSAP_getClientId_Call{Call: _e.mock.On("getClientId")}
}

func (_c *mock_BIPSAP_getClientId_Call) Run(run func()) *mock_BIPSAP_getClientId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_BIPSAP_getClientId_Call) Return(_a0 *int) *mock_BIPSAP_getClientId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_BIPSAP_getClientId_Call) RunAndReturn(run func() *int) *mock_BIPSAP_getClientId_Call {
	_c.Call.Return(run)
	return _c
}

// newMock_BIPSAP creates a new instance of mock_BIPSAP. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMock_BIPSAP(t interface {
	mock.TestingT
	Cleanup(func())
}) *mock_BIPSAP {
	mock := &mock_BIPSAP{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
