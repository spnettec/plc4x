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

// Code generated by mockery v2.30.16. DO NOT EDIT.

package interceptors

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// mockReadResponseFactory is an autogenerated mock type for the readResponseFactory type
type mockReadResponseFactory struct {
	mock.Mock
}

type mockReadResponseFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *mockReadResponseFactory) EXPECT() *mockReadResponseFactory_Expecter {
	return &mockReadResponseFactory_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: request, responseCodes, _a2
func (_m *mockReadResponseFactory) Execute(request model.PlcReadRequest, responseCodes map[string]model.PlcResponseCode, _a2 map[string]values.PlcValue) model.PlcReadResponse {
	ret := _m.Called(request, responseCodes, _a2)

	var r0 model.PlcReadResponse
	if rf, ok := ret.Get(0).(func(model.PlcReadRequest, map[string]model.PlcResponseCode, map[string]values.PlcValue) model.PlcReadResponse); ok {
		r0 = rf(request, responseCodes, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcReadResponse)
		}
	}

	return r0
}

// mockReadResponseFactory_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type mockReadResponseFactory_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - request model.PlcReadRequest
//   - responseCodes map[string]model.PlcResponseCode
//   - _a2 map[string]values.PlcValue
func (_e *mockReadResponseFactory_Expecter) Execute(request interface{}, responseCodes interface{}, _a2 interface{}) *mockReadResponseFactory_Execute_Call {
	return &mockReadResponseFactory_Execute_Call{Call: _e.mock.On("Execute", request, responseCodes, _a2)}
}

func (_c *mockReadResponseFactory_Execute_Call) Run(run func(request model.PlcReadRequest, responseCodes map[string]model.PlcResponseCode, _a2 map[string]values.PlcValue)) *mockReadResponseFactory_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.PlcReadRequest), args[1].(map[string]model.PlcResponseCode), args[2].(map[string]values.PlcValue))
	})
	return _c
}

func (_c *mockReadResponseFactory_Execute_Call) Return(_a0 model.PlcReadResponse) *mockReadResponseFactory_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockReadResponseFactory_Execute_Call) RunAndReturn(run func(model.PlcReadRequest, map[string]model.PlcResponseCode, map[string]values.PlcValue) model.PlcReadResponse) *mockReadResponseFactory_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// newMockReadResponseFactory creates a new instance of mockReadResponseFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockReadResponseFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockReadResponseFactory {
	mock := &mockReadResponseFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
