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

package cbus

import (
	"context"
	apiModel "github.com/apache/plc4x/plc4go/pkg/api/model"
	apiValues "github.com/apache/plc4x/plc4go/pkg/api/values"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
	spiModel "github.com/apache/plc4x/plc4go/spi/model"
	"github.com/apache/plc4x/plc4go/spi/options"
	"github.com/apache/plc4x/plc4go/spi/testutils"
	"github.com/apache/plc4x/plc4go/spi/transactions"
	"github.com/apache/plc4x/plc4go/spi/transports/test"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewReader(t *testing.T) {
	type args struct {
		tpduGenerator *AlphaGenerator
		messageCodec  *MessageCodec
		tm            transactions.RequestTransactionManager
	}
	tests := []struct {
		name string
		args args
		want *Reader
	}{
		{
			name: "create a new one",
			want: &Reader{
				alphaGenerator: nil,
				messageCodec:   nil,
				tm:             nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := testutils.ProduceTestingLogger(t)
			reader := NewReader(tt.args.tpduGenerator, tt.args.messageCodec, tt.args.tm, options.WithCustomLogger(logger))
			tt.want.log = logger
			assert.Equalf(t, tt.want, reader, "NewReader(%v, %v, %v)", tt.args.tpduGenerator, tt.args.messageCodec, tt.args.tm)
		})
	}
}

func TestReader_Read(t *testing.T) {
	type fields struct {
		alphaGenerator *AlphaGenerator
		messageCodec   *MessageCodec
		tm             transactions.RequestTransactionManager
	}
	type args struct {
		ctx         context.Context
		readRequest apiModel.PlcReadRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantAsserter func(t *testing.T, results <-chan apiModel.PlcReadRequestResult) bool
	}{
		{
			name: "read and bail",
			args: args{
				ctx: testutils.TestContext(t),
				readRequest: spiModel.NewDefaultPlcReadRequest(nil, func() []string {
					return strings.Split(strings.Repeat("asd,", 40), ",")
				}(), nil, nil),
			},
			wantAsserter: func(t *testing.T, results <-chan apiModel.PlcReadRequestResult) bool {
				timer := time.NewTimer(2 * time.Second)
				defer timer.Stop()
				select {
				case <-timer.C:
					t.Fail()
				case result := <-results:
					assert.NotNil(t, result.GetErr())
				}
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Reader{
				alphaGenerator: tt.fields.alphaGenerator,
				messageCodec:   tt.fields.messageCodec,
				tm:             tt.fields.tm,
			}
			assert.Truef(t, tt.wantAsserter(t, m.Read(tt.args.ctx, tt.args.readRequest)), "Read(%v, %v)", tt.args.ctx, tt.args.readRequest)
		})
	}
}

func TestReader_readSync(t *testing.T) {
	type fields struct {
		alphaGenerator *AlphaGenerator
		messageCodec   *MessageCodec
		tm             transactions.RequestTransactionManager
	}
	type args struct {
		ctx         context.Context
		readRequest apiModel.PlcReadRequest
		result      chan apiModel.PlcReadRequestResult
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		setup           func(t *testing.T, fields *fields)
		resultEvaluator func(t *testing.T, results chan apiModel.PlcReadRequestResult) bool
	}{
		{
			name: "too many tags",
			args: args{
				ctx: testutils.TestContext(t),
				readRequest: spiModel.NewDefaultPlcReadRequest(nil, func() []string {
					return strings.Split(strings.Repeat("asd,", 40), ",")
				}(), nil, nil),
				result: make(chan apiModel.PlcReadRequestResult, 1),
			},
			resultEvaluator: func(t *testing.T, results chan apiModel.PlcReadRequestResult) bool {
				timer := time.NewTimer(2 * time.Second)
				defer timer.Stop()
				select {
				case <-timer.C:
					t.Fail()
				case result := <-results:
					assert.NotNil(t, result.GetErr())
				}
				return true
			},
		},
		{
			name: "unmapped tag",
			args: args{
				ctx: testutils.TestContext(t),
				readRequest: spiModel.NewDefaultPlcReadRequest(
					map[string]apiModel.PlcTag{
						"asd": nil,
					},
					[]string{
						"asd",
					},
					nil,
					nil,
				),
				result: make(chan apiModel.PlcReadRequestResult, 1),
			},
			setup: func(t *testing.T, fields *fields) {
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				transport := test.NewTransport(_options...)
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec
				fields.tm = transactions.NewRequestTransactionManager(10, _options...)
			},
			resultEvaluator: func(t *testing.T, results chan apiModel.PlcReadRequestResult) bool {
				timer := time.NewTimer(2 * time.Second)
				defer timer.Stop()
				select {
				case <-timer.C:
					t.Fail()
				case result := <-results:
					assert.NotNil(t, result.GetErr())
				}
				return true
			},
		},
		{
			name: "read something without any tag",
			args: args{
				ctx: testutils.TestContext(t),
				readRequest: spiModel.NewDefaultPlcReadRequest(
					map[string]apiModel.PlcTag{},
					[]string{},
					nil,
					nil,
				),
				result: make(chan apiModel.PlcReadRequestResult, 1),
			},
			resultEvaluator: func(t *testing.T, results chan apiModel.PlcReadRequestResult) bool {
				timer := time.NewTimer(2 * time.Second)
				defer timer.Stop()
				select {
				case <-timer.C:
					t.Fail()
				case result := <-results:
					assert.Nil(t, result.GetErr())
					assert.NotNil(t, result.GetResponse())
				}
				return true
			},
		},
		{
			name: "read identify type",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					t.Cleanup(cancel)
					return timeout
				}(),
				readRequest: spiModel.NewDefaultPlcReadRequest(
					map[string]apiModel.PlcTag{
						"blub": NewCALIdentifyTag(readWriteModel.NewUnitAddress(2), nil, readWriteModel.Attribute_Type, 1),
					},
					[]string{
						"blub",
					},
					nil,
					nil,
				),
				result: make(chan apiModel.PlcReadRequestResult, 1),
			},
			setup: func(t *testing.T, fields *fields) {
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				fields.tm = transactions.NewRequestTransactionManager(10, _options...)
				transport := test.NewTransport()
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				type MockState uint8
				const (
					INITIAL MockState = iota
					DONE
				)
				currentState := atomic.Value{}
				currentState.Store(INITIAL)
				stateChangeMutex := sync.Mutex{}
				transportInstance.(*test.TransportInstance).SetWriteInterceptor(func(transportInstance *test.TransportInstance, data []byte) {
					stateChangeMutex.Lock()
					defer stateChangeMutex.Unlock()
					switch currentState.Load().(MockState) {
					case INITIAL:
						t.Log("Dispatching read response")
						transportInstance.FillReadBuffer([]byte("g.890150435F434E49454421\r\n"))
						currentState.Store(DONE)
					case DONE:
						t.Log("Done")
					}
				})
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec
			},
			resultEvaluator: func(t *testing.T, results chan apiModel.PlcReadRequestResult) bool {
				timer := time.NewTimer(2 * time.Second)
				defer timer.Stop()
				select {
				case <-timer.C:
					t.Fail()
				case result := <-results:
					assert.Nil(t, result.GetErr())
					response := result.GetResponse()
					assert.NotNil(t, response)
					value := response.GetValue("blub")
					assert.NotNil(t, value)
					assert.Equal(t, "PC_CNIED", value.GetString())
				}
				return true
			},
		},
		{
			name: "read identify type aborted",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithCancel(context.Background())
					cancel()
					return timeout
				}(),
				readRequest: spiModel.NewDefaultPlcReadRequest(
					map[string]apiModel.PlcTag{
						"blub": NewCALIdentifyTag(readWriteModel.NewUnitAddress(2), nil, readWriteModel.Attribute_Type, 1),
					},
					[]string{
						"blub",
					},
					nil,
					nil,
				),
				result: make(chan apiModel.PlcReadRequestResult, 1),
			},
			setup: func(t *testing.T, fields *fields) {
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				fields.tm = transactions.NewRequestTransactionManager(10, _options...)

				transport := test.NewTransport()
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec
			},
			resultEvaluator: func(t *testing.T, results chan apiModel.PlcReadRequestResult) bool {
				timer := time.NewTimer(2 * time.Second)
				defer timer.Stop()
				select {
				case <-timer.C:
					t.Fail()
				case result := <-results:
					assert.NotNil(t, result.GetErr())
				}
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(t, &tt.fields)
			}
			m := &Reader{
				alphaGenerator: tt.fields.alphaGenerator,
				messageCodec:   tt.fields.messageCodec,
				tm:             tt.fields.tm,
			}
			m.readSync(tt.args.ctx, tt.args.readRequest, tt.args.result)
			assert.True(t, tt.resultEvaluator(t, tt.args.result))
		})
	}
}

func TestReader_sendMessageOverTheWire(t *testing.T) {
	type fields struct {
		alphaGenerator *AlphaGenerator
		messageCodec   *MessageCodec
		tm             transactions.RequestTransactionManager
	}
	type args struct {
		ctx             context.Context
		transaction     transactions.RequestTransaction
		messageToSend   readWriteModel.CBusMessage
		addResponseCode func(t *testing.T) func(name string, responseCode apiModel.PlcResponseCode)
		tagName         string
		addPlcValue     func(t *testing.T) func(name string, plcValue apiValues.PlcValue)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		setup  func(t *testing.T, fields *fields, args *args, ch chan struct{})
	}{
		{
			name: "Send message empty message",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					t.Cleanup(cancel)
					return timeout
				}(),
				messageToSend: nil,
				addResponseCode: func(t *testing.T) func(name string, responseCode apiModel.PlcResponseCode) {
					return func(name string, responseCode apiModel.PlcResponseCode) {
						t.Logf("Got response code %s for %s", responseCode, name)
						assert.Equal(t, "horst", name)
						assert.Equal(t, apiModel.PlcResponseCode_INTERNAL_ERROR, responseCode)
					}
				},
				tagName: "horst",
				addPlcValue: func(t *testing.T) func(name string, plcValue apiValues.PlcValue) {
					return func(name string, plcValue apiValues.PlcValue) {
						t.Logf("Got response %s for %s", plcValue, name)
					}
				},
			},
			setup: func(t *testing.T, fields *fields, args *args, ch chan struct{}) {
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				transport := test.NewTransport(_options...)
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec

				transaction := NewMockRequestTransaction(t)
				expect := transaction.EXPECT()
				expect.FailRequest(mock.Anything).Return(errors.New("no I say")).Run(func(_ error) {
					close(ch)
				})
				args.transaction = transaction
			},
		},
		{
			name: "Send message which responds with message to client",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					t.Cleanup(cancel)
					return timeout
				}(),
				messageToSend: readWriteModel.NewCBusMessageToServer(
					readWriteModel.NewRequestReset(
						readWriteModel.RequestType_RESET,
						nil,
						0,
						nil,
						readWriteModel.RequestType_EMPTY,
						nil,
						nil,
						readWriteModel.RequestType_EMPTY,
						readWriteModel.NewRequestTermination(),
						nil,
					),
					nil,
					nil,
				),
				addResponseCode: func(t *testing.T) func(name string, responseCode apiModel.PlcResponseCode) {
					return func(name string, responseCode apiModel.PlcResponseCode) {
						t.Logf("Got response code %s for %s", responseCode, name)
						assert.Equal(t, "horst", name)
						assert.Equal(t, apiModel.PlcResponseCode_REQUEST_TIMEOUT, responseCode)
					}
				},
				tagName: "horst",
				addPlcValue: func(t *testing.T) func(name string, plcValue apiValues.PlcValue) {
					return func(name string, plcValue apiValues.PlcValue) {
						t.Logf("Got response %s for %s", plcValue, name)
					}
				},
			},
			setup: func(t *testing.T, fields *fields, args *args, ch chan struct{}) {
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)
				transaction := NewMockRequestTransaction(t)
				expect := transaction.EXPECT()
				expect.FailRequest(mock.Anything).Return(errors.New("Nope")).Run(func(_ error) {
					close(ch)
				})
				args.transaction = transaction

				transport := test.NewTransport(_options...)
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				type MockState uint8
				const (
					INITIAL MockState = iota
					DONE
				)
				currentState := atomic.Value{}
				currentState.Store(INITIAL)
				stateChangeMutex := sync.Mutex{}
				transportInstance.(*test.TransportInstance).SetWriteInterceptor(func(transportInstance *test.TransportInstance, data []byte) {
					stateChangeMutex.Lock()
					defer stateChangeMutex.Unlock()
					switch currentState.Load().(MockState) {
					case INITIAL:
						t.Log("Dispatching read response")
						transportInstance.FillReadBuffer([]byte("@1A2001\r@"))
						currentState.Store(DONE)
					case DONE:
						t.Log("Done")
					}
				})
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec
			},
		},
		{
			name: "Send message which responds with server error",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					t.Cleanup(cancel)
					return timeout
				}(),
				messageToSend: readWriteModel.NewCBusMessageToServer(
					readWriteModel.NewRequestReset(
						readWriteModel.RequestType_RESET,
						nil,
						0,
						nil,
						readWriteModel.RequestType_EMPTY,
						nil,
						nil,
						readWriteModel.RequestType_EMPTY,
						readWriteModel.NewRequestTermination(),
						nil,
					),
					nil,
					nil,
				),
				addResponseCode: func(t *testing.T) func(name string, responseCode apiModel.PlcResponseCode) {
					return func(name string, responseCode apiModel.PlcResponseCode) {
						t.Logf("Got response code %s for %s", responseCode, name)
						assert.Equal(t, "horst", name)
						assert.Equal(t, apiModel.PlcResponseCode_INVALID_DATA, responseCode)
					}
				},
				tagName: "horst",
				addPlcValue: func(t *testing.T) func(name string, plcValue apiValues.PlcValue) {
					return func(name string, plcValue apiValues.PlcValue) {
						t.Logf("Got response %s for %s", plcValue, name)
					}
				},
			},
			setup: func(t *testing.T, fields *fields, args *args, ch chan struct{}) {
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				transport := test.NewTransport(_options...)
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				type MockState uint8
				const (
					INITIAL MockState = iota
					DONE
				)
				currentState := atomic.Value{}
				currentState.Store(INITIAL)
				stateChangeMutex := sync.Mutex{}
				transportInstance.(*test.TransportInstance).SetWriteInterceptor(func(transportInstance *test.TransportInstance, data []byte) {
					stateChangeMutex.Lock()
					defer stateChangeMutex.Unlock()
					switch currentState.Load().(MockState) {
					case INITIAL:
						t.Log("Dispatching read response")
						transportInstance.FillReadBuffer([]byte("!"))
						currentState.Store(DONE)
					case DONE:
						t.Log("Done")
					}
				})
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec

				transaction := NewMockRequestTransaction(t)
				expect := transaction.EXPECT()
				expect.EndRequest().Return(nil).Run(func() {
					close(ch)
				})
				args.transaction = transaction
			},
		},
		{
			name: "Send message which responds with too many retransmissions",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					t.Cleanup(cancel)
					return timeout
				}(),
				messageToSend: readWriteModel.NewCBusMessageToServer(
					readWriteModel.NewRequestDirectCommandAccess(
						readWriteModel.NewCALDataIdentify(
							readWriteModel.Attribute_CurrentSenseLevels,
							readWriteModel.CALCommandTypeContainer_CALCommandIdentify,
							nil,
							nil,
						),
						readWriteModel.NewAlpha('g'),
						readWriteModel.RequestType_DIRECT_COMMAND,
						nil,
						nil,
						readWriteModel.RequestType_EMPTY,
						readWriteModel.NewRequestTermination(),
						nil,
					),
					nil,
					nil,
				),
				addResponseCode: func(t *testing.T) func(name string, responseCode apiModel.PlcResponseCode) {
					return func(name string, responseCode apiModel.PlcResponseCode) {
						t.Logf("Got response code %s for %s", responseCode, name)
						assert.Equal(t, "horst", name)
						assert.Equal(t, apiModel.PlcResponseCode_REMOTE_ERROR, responseCode)
					}
				},
				tagName: "horst",
				addPlcValue: func(t *testing.T) func(name string, plcValue apiValues.PlcValue) {
					return func(name string, plcValue apiValues.PlcValue) {
						t.Logf("Got response %s for %s", plcValue, name)
					}
				},
			},
			setup: func(t *testing.T, fields *fields, args *args, ch chan struct{}) {
				transaction := NewMockRequestTransaction(t)
				expect := transaction.EXPECT()
				expect.EndRequest().Return(nil).Run(func() {
					close(ch)
				})
				args.transaction = transaction
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				transport := test.NewTransport(_options...)
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				type MockState uint8
				const (
					INITIAL MockState = iota
					DONE
				)
				currentState := atomic.Value{}
				currentState.Store(INITIAL)
				stateChangeMutex := sync.Mutex{}
				transportInstance.(*test.TransportInstance).SetWriteInterceptor(func(transportInstance *test.TransportInstance, data []byte) {
					stateChangeMutex.Lock()
					defer stateChangeMutex.Unlock()
					switch currentState.Load().(MockState) {
					case INITIAL:
						t.Log("Dispatching read response")
						transportInstance.FillReadBuffer([]byte("g#\r\n"))
						currentState.Store(DONE)
					case DONE:
						t.Log("Done")
					}
				})
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec
			},
		},
		{
			name: "Send message which responds with corruption",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					t.Cleanup(cancel)
					return timeout
				}(),
				messageToSend: readWriteModel.NewCBusMessageToServer(
					readWriteModel.NewRequestDirectCommandAccess(
						readWriteModel.NewCALDataIdentify(
							readWriteModel.Attribute_CurrentSenseLevels,
							readWriteModel.CALCommandTypeContainer_CALCommandIdentify,
							nil,
							nil,
						),
						readWriteModel.NewAlpha('g'),
						readWriteModel.RequestType_DIRECT_COMMAND,
						nil,
						nil,
						readWriteModel.RequestType_EMPTY,
						readWriteModel.NewRequestTermination(),
						nil,
					),
					nil,
					nil,
				),
				addResponseCode: func(t *testing.T) func(name string, responseCode apiModel.PlcResponseCode) {
					return func(name string, responseCode apiModel.PlcResponseCode) {
						t.Logf("Got response code %s for %s", responseCode, name)
						assert.Equal(t, "horst", name)
						assert.Equal(t, apiModel.PlcResponseCode_INVALID_DATA, responseCode)
					}
				},
				tagName: "horst",
				addPlcValue: func(t *testing.T) func(name string, plcValue apiValues.PlcValue) {
					return func(name string, plcValue apiValues.PlcValue) {
						t.Logf("Got response %s for %s", plcValue, name)
					}
				},
			},
			setup: func(t *testing.T, fields *fields, args *args, ch chan struct{}) {
				transaction := NewMockRequestTransaction(t)
				expect := transaction.EXPECT()
				expect.EndRequest().Return(nil).Run(func() {
					close(ch)
				})
				args.transaction = transaction
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				transport := test.NewTransport(_options...)
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				type MockState uint8
				const (
					INITIAL MockState = iota
					DONE
				)
				currentState := atomic.Value{}
				currentState.Store(INITIAL)
				stateChangeMutex := sync.Mutex{}
				transportInstance.(*test.TransportInstance).SetWriteInterceptor(func(transportInstance *test.TransportInstance, data []byte) {
					stateChangeMutex.Lock()
					defer stateChangeMutex.Unlock()
					switch currentState.Load().(MockState) {
					case INITIAL:
						t.Log("Dispatching read response")
						transportInstance.FillReadBuffer([]byte("g$\r\n"))
						currentState.Store(DONE)
					case DONE:
						t.Log("Done")
					}
				})
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec
			},
		},
		{
			name: "Send message which responds with sync loss",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					t.Cleanup(cancel)
					return timeout
				}(),
				messageToSend: readWriteModel.NewCBusMessageToServer(
					readWriteModel.NewRequestDirectCommandAccess(
						readWriteModel.NewCALDataIdentify(
							readWriteModel.Attribute_CurrentSenseLevels,
							readWriteModel.CALCommandTypeContainer_CALCommandIdentify,
							nil,
							nil,
						),
						readWriteModel.NewAlpha('g'),
						readWriteModel.RequestType_DIRECT_COMMAND,
						nil,
						nil,
						readWriteModel.RequestType_EMPTY,
						readWriteModel.NewRequestTermination(),
						nil,
					),
					nil,
					nil,
				),
				addResponseCode: func(t *testing.T) func(name string, responseCode apiModel.PlcResponseCode) {
					return func(name string, responseCode apiModel.PlcResponseCode) {
						t.Logf("Got response code %s for %s", responseCode, name)
						assert.Equal(t, "horst", name)
						assert.Equal(t, apiModel.PlcResponseCode_REMOTE_BUSY, responseCode)
					}
				},
				tagName: "horst",
				addPlcValue: func(t *testing.T) func(name string, plcValue apiValues.PlcValue) {
					return func(name string, plcValue apiValues.PlcValue) {
						t.Logf("Got response %s for %s", plcValue, name)
					}
				},
			},
			setup: func(t *testing.T, fields *fields, args *args, ch chan struct{}) {
				transaction := NewMockRequestTransaction(t)
				expect := transaction.EXPECT()
				expect.EndRequest().Return(nil).Run(func() {
					close(ch)
				})
				args.transaction = transaction
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				transport := test.NewTransport(_options...)
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				type MockState uint8
				const (
					INITIAL MockState = iota
					DONE
				)
				currentState := atomic.Value{}
				currentState.Store(INITIAL)
				stateChangeMutex := sync.Mutex{}
				transportInstance.(*test.TransportInstance).SetWriteInterceptor(func(transportInstance *test.TransportInstance, data []byte) {
					stateChangeMutex.Lock()
					defer stateChangeMutex.Unlock()
					switch currentState.Load().(MockState) {
					case INITIAL:
						t.Log("Dispatching read response")
						transportInstance.FillReadBuffer([]byte("g%\r\n"))
						currentState.Store(DONE)
					case DONE:
						t.Log("Done")
					}
				})
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec
			},
		},
		{
			name: "Send message which responds with too long",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					t.Cleanup(cancel)
					return timeout
				}(),
				messageToSend: readWriteModel.NewCBusMessageToServer(
					readWriteModel.NewRequestDirectCommandAccess(
						readWriteModel.NewCALDataIdentify(
							readWriteModel.Attribute_CurrentSenseLevels,
							readWriteModel.CALCommandTypeContainer_CALCommandIdentify,
							nil,
							nil,
						),
						readWriteModel.NewAlpha('g'),
						readWriteModel.RequestType_DIRECT_COMMAND,
						nil,
						nil,
						readWriteModel.RequestType_EMPTY,
						readWriteModel.NewRequestTermination(),
						nil,
					),
					nil,
					nil,
				),
				addResponseCode: func(t *testing.T) func(name string, responseCode apiModel.PlcResponseCode) {
					return func(name string, responseCode apiModel.PlcResponseCode) {
						t.Logf("Got response code %s for %s", responseCode, name)
						assert.Equal(t, "horst", name)
						assert.Equal(t, apiModel.PlcResponseCode_INVALID_DATA, responseCode)
					}
				},
				tagName: "horst",
				addPlcValue: func(t *testing.T) func(name string, plcValue apiValues.PlcValue) {
					return func(name string, plcValue apiValues.PlcValue) {
						t.Logf("Got response %s for %s", plcValue, name)
					}
				},
			},
			setup: func(t *testing.T, fields *fields, args *args, ch chan struct{}) {
				transaction := NewMockRequestTransaction(t)
				expect := transaction.EXPECT()
				expect.EndRequest().Return(nil).Run(func() {
					close(ch)
				})
				args.transaction = transaction
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				transport := test.NewTransport(_options...)
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				type MockState uint8
				const (
					INITIAL MockState = iota
					DONE
				)
				currentState := atomic.Value{}
				currentState.Store(INITIAL)
				stateChangeMutex := sync.Mutex{}
				transportInstance.(*test.TransportInstance).SetWriteInterceptor(func(transportInstance *test.TransportInstance, data []byte) {
					stateChangeMutex.Lock()
					defer stateChangeMutex.Unlock()
					switch currentState.Load().(MockState) {
					case INITIAL:
						t.Log("Dispatching read response")
						transportInstance.FillReadBuffer([]byte("g'\r\n"))
						currentState.Store(DONE)
					case DONE:
						t.Log("Done")
					}
				})
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec
			},
		},
		{
			name: "Send message which responds with confirm only",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					t.Cleanup(cancel)
					return timeout
				}(),
				messageToSend: readWriteModel.NewCBusMessageToServer(
					readWriteModel.NewRequestDirectCommandAccess(
						readWriteModel.NewCALDataIdentify(
							readWriteModel.Attribute_CurrentSenseLevels,
							readWriteModel.CALCommandTypeContainer_CALCommandIdentify,
							nil,
							nil,
						),
						readWriteModel.NewAlpha('g'),
						readWriteModel.RequestType_DIRECT_COMMAND,
						nil,
						nil,
						readWriteModel.RequestType_EMPTY,
						readWriteModel.NewRequestTermination(),
						nil,
					),
					nil,
					nil,
				),
				addResponseCode: func(t *testing.T) func(name string, responseCode apiModel.PlcResponseCode) {
					return func(name string, responseCode apiModel.PlcResponseCode) {
						t.Logf("Got response code %s for %s", responseCode, name)
						assert.Equal(t, "horst", name)
						assert.Equal(t, apiModel.PlcResponseCode_NOT_FOUND, responseCode)
					}
				},
				tagName: "horst",
				addPlcValue: func(t *testing.T) func(name string, plcValue apiValues.PlcValue) {
					return func(name string, plcValue apiValues.PlcValue) {
						t.Logf("Got response %s for %s", plcValue, name)
					}
				},
			},
			setup: func(t *testing.T, fields *fields, args *args, ch chan struct{}) {
				transaction := NewMockRequestTransaction(t)
				expect := transaction.EXPECT()
				expect.EndRequest().Return(nil).Run(func() {
					close(ch)
				})
				args.transaction = transaction
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				transport := test.NewTransport(_options...)
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				type MockState uint8
				const (
					INITIAL MockState = iota
					DONE
				)
				currentState := atomic.Value{}
				currentState.Store(INITIAL)
				stateChangeMutex := sync.Mutex{}
				transportInstance.(*test.TransportInstance).SetWriteInterceptor(func(transportInstance *test.TransportInstance, data []byte) {
					stateChangeMutex.Lock()
					defer stateChangeMutex.Unlock()
					switch currentState.Load().(MockState) {
					case INITIAL:
						t.Log("Dispatching read response")
						transportInstance.FillReadBuffer([]byte("g.\r\n"))
						currentState.Store(DONE)
					case DONE:
						t.Log("Done")
					}
				})
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec
			},
		},
		{
			name: "Send message which responds with ok",
			fields: fields{
				alphaGenerator: &AlphaGenerator{currentAlpha: 'g'},
			},
			args: args{
				ctx: func() context.Context {
					timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					t.Cleanup(cancel)
					return timeout
				}(),
				messageToSend: readWriteModel.NewCBusMessageToServer(
					readWriteModel.NewRequestDirectCommandAccess(
						readWriteModel.NewCALDataIdentify(
							readWriteModel.Attribute_CurrentSenseLevels,
							readWriteModel.CALCommandTypeContainer_CALCommandIdentify,
							nil,
							nil,
						),
						readWriteModel.NewAlpha('g'),
						readWriteModel.RequestType_DIRECT_COMMAND,
						nil,
						nil,
						readWriteModel.RequestType_EMPTY,
						readWriteModel.NewRequestTermination(),
						nil,
					),
					nil,
					nil,
				),
				addResponseCode: func(t *testing.T) func(name string, responseCode apiModel.PlcResponseCode) {
					return func(name string, responseCode apiModel.PlcResponseCode) {
						t.Logf("Got response code %s for %s", responseCode, name)
						assert.Equal(t, "horst", name)
						assert.Equal(t, apiModel.PlcResponseCode_OK, responseCode)
					}
				},
				tagName: "horst",
				addPlcValue: func(t *testing.T) func(name string, plcValue apiValues.PlcValue) {
					return func(name string, plcValue apiValues.PlcValue) {
						t.Logf("Got response %s for %s", plcValue, name)
					}
				},
			},
			setup: func(t *testing.T, fields *fields, args *args, ch chan struct{}) {
				transaction := NewMockRequestTransaction(t)
				expect := transaction.EXPECT()
				expect.EndRequest().Return(nil).Run(func() {
					close(ch)
				})
				args.transaction = transaction
				_options := testutils.EnrichOptionsWithOptionsForTesting(t)

				transport := test.NewTransport(_options...)
				transportUrl := url.URL{Scheme: "test"}
				transportInstance, err := transport.CreateTransportInstance(transportUrl, nil, _options...)
				require.NoError(t, err)
				type MockState uint8
				const (
					INITIAL MockState = iota
					DONE
				)
				currentState := atomic.Value{}
				currentState.Store(INITIAL)
				stateChangeMutex := sync.Mutex{}
				transportInstance.(*test.TransportInstance).SetWriteInterceptor(func(transportInstance *test.TransportInstance, data []byte) {
					stateChangeMutex.Lock()
					defer stateChangeMutex.Unlock()
					switch currentState.Load().(MockState) {
					case INITIAL:
						t.Log("Dispatching read response")
						transportInstance.FillReadBuffer([]byte("g.890150435F434E49454421\r\n"))
						currentState.Store(DONE)
					case DONE:
						t.Log("Done")
					}
				})
				codec := NewMessageCodec(transportInstance, _options...)
				require.NoError(t, codec.Connect())
				t.Cleanup(func() {
					assert.NoError(t, codec.Disconnect())
				})
				fields.messageCodec = codec
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan struct{})
			if tt.setup != nil {
				tt.setup(t, &tt.fields, &tt.args, ch)
			}
			m := &Reader{
				alphaGenerator: tt.fields.alphaGenerator,
				messageCodec:   tt.fields.messageCodec,
				tm:             tt.fields.tm,
			}
			m.sendMessageOverTheWire(tt.args.ctx, tt.args.transaction, tt.args.messageToSend, tt.args.addResponseCode(t), tt.args.tagName, tt.args.addPlcValue(t))
			t.Log("Waiting now")
			timer := time.NewTimer(3 * time.Second)
			defer utils.CleanupTimer(timer)
			select {
			case <-ch:
				t.Log("Done waiting")
			case <-timer.C:
				t.Error("Timeout")
			}
		})
	}
}
