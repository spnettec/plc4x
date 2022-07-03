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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// S7MessageResponse is the corresponding interface of S7MessageResponse
type S7MessageResponse interface {
	utils.LengthAware
	utils.Serializable
	S7Message
	// GetErrorClass returns ErrorClass (property field)
	GetErrorClass() uint8
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() uint8
}

// S7MessageResponseExactly can be used when we want exactly this type and not a type which fulfills S7MessageResponse.
// This is useful for switch cases.
type S7MessageResponseExactly interface {
	S7MessageResponse
	isS7MessageResponse() bool
}

// _S7MessageResponse is the data-structure of this message
type _S7MessageResponse struct {
	*_S7Message
	ErrorClass uint8
	ErrorCode  uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7MessageResponse) GetMessageType() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7MessageResponse) InitializeParent(parent S7Message, tpduReference uint16, parameter S7Parameter, payload S7Payload) {
	m.TpduReference = tpduReference
	m.Parameter = parameter
	m.Payload = payload
}

func (m *_S7MessageResponse) GetParent() S7Message {
	return m._S7Message
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7MessageResponse) GetErrorClass() uint8 {
	return m.ErrorClass
}

func (m *_S7MessageResponse) GetErrorCode() uint8 {
	return m.ErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7MessageResponse factory function for _S7MessageResponse
func NewS7MessageResponse(errorClass uint8, errorCode uint8, tpduReference uint16, parameter S7Parameter, payload S7Payload) *_S7MessageResponse {
	_result := &_S7MessageResponse{
		ErrorClass: errorClass,
		ErrorCode:  errorCode,
		_S7Message: NewS7Message(tpduReference, parameter, payload),
	}
	_result._S7Message._S7MessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7MessageResponse(structType interface{}) S7MessageResponse {
	if casted, ok := structType.(S7MessageResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7MessageResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7MessageResponse) GetTypeName() string {
	return "S7MessageResponse"
}

func (m *_S7MessageResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7MessageResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (errorClass)
	lengthInBits += 8

	// Simple field (errorCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7MessageResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7MessageResponseParse(readBuffer utils.ReadBuffer) (S7MessageResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7MessageResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7MessageResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (errorClass)
	_errorClass, _errorClassErr := readBuffer.ReadUint8("errorClass", 8)
	if _errorClassErr != nil {
		return nil, errors.Wrap(_errorClassErr, "Error parsing 'errorClass' field of S7MessageResponse")
	}
	errorClass := _errorClass

	// Simple Field (errorCode)
	_errorCode, _errorCodeErr := readBuffer.ReadUint8("errorCode", 8)
	if _errorCodeErr != nil {
		return nil, errors.Wrap(_errorCodeErr, "Error parsing 'errorCode' field of S7MessageResponse")
	}
	errorCode := _errorCode

	if closeErr := readBuffer.CloseContext("S7MessageResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7MessageResponse")
	}

	// Create a partially initialized instance
	_child := &_S7MessageResponse{
		ErrorClass: errorClass,
		ErrorCode:  errorCode,
		_S7Message: &_S7Message{},
	}
	_child._S7Message._S7MessageChildRequirements = _child
	return _child, nil
}

func (m *_S7MessageResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7MessageResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7MessageResponse")
		}

		// Simple Field (errorClass)
		errorClass := uint8(m.GetErrorClass())
		_errorClassErr := writeBuffer.WriteUint8("errorClass", 8, (errorClass))
		if _errorClassErr != nil {
			return errors.Wrap(_errorClassErr, "Error serializing 'errorClass' field")
		}

		// Simple Field (errorCode)
		errorCode := uint8(m.GetErrorCode())
		_errorCodeErr := writeBuffer.WriteUint8("errorCode", 8, (errorCode))
		if _errorCodeErr != nil {
			return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
		}

		if popErr := writeBuffer.PopContext("S7MessageResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7MessageResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7MessageResponse) isS7MessageResponse() bool {
	return true
}

func (m *_S7MessageResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
