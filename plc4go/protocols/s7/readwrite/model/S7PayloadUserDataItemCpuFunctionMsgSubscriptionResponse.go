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

package model


import (
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse interface {
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
}

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseExactly interface {
	S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
	isS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse() bool
}

// _S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse struct {
	*_S7PayloadUserDataItem
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse)  GetCpuFunctionType() uint8 {
return 0x08}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse)  GetCpuSubfunction() uint8 {
return 0x02}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse)  GetDataLength() uint16 {
return 0x00}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) InitializeParent(parent S7PayloadUserDataItem , returnCode DataTransportErrorCode , transportSize DataTransportSize ) {	m.ReturnCode = returnCode
	m.TransportSize = transportSize
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse)  GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}


// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse factory function for _S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse( returnCode DataTransportErrorCode , transportSize DataTransportSize ) *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse{
    	_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse(structType interface{}) S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse {
    if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{
		},
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) isS7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



