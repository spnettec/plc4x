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

// CIPEncapsulationReadRequest is the corresponding interface of CIPEncapsulationReadRequest
type CIPEncapsulationReadRequest interface {
	CIPEncapsulationPacket
	// GetRequest returns Request (property field)
	GetRequest() DF1RequestMessage
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _CIPEncapsulationReadRequest is the data-structure of this message
type _CIPEncapsulationReadRequest struct {
	*_CIPEncapsulationPacket
	Request DF1RequestMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CIPEncapsulationReadRequest) GetCommandType() uint16 {
	return 0x0107
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CIPEncapsulationReadRequest) InitializeParent(parent CIPEncapsulationPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_CIPEncapsulationReadRequest) GetParent() CIPEncapsulationPacket {
	return m._CIPEncapsulationPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPEncapsulationReadRequest) GetRequest() DF1RequestMessage {
	return m.Request
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCIPEncapsulationReadRequest factory function for _CIPEncapsulationReadRequest
func NewCIPEncapsulationReadRequest(request DF1RequestMessage, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *_CIPEncapsulationReadRequest {
	_result := &_CIPEncapsulationReadRequest{
		Request:                 request,
		_CIPEncapsulationPacket: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
	}
	_result._CIPEncapsulationPacket._CIPEncapsulationPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCIPEncapsulationReadRequest(structType interface{}) CIPEncapsulationReadRequest {
	if casted, ok := structType.(CIPEncapsulationReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CIPEncapsulationReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CIPEncapsulationReadRequest) GetTypeName() string {
	return "CIPEncapsulationReadRequest"
}

func (m *_CIPEncapsulationReadRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CIPEncapsulationReadRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (request)
	lengthInBits += m.Request.GetLengthInBits()

	return lengthInBits
}

func (m *_CIPEncapsulationReadRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CIPEncapsulationReadRequestParse(readBuffer utils.ReadBuffer) (CIPEncapsulationReadRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (request)
	if pullErr := readBuffer.PullContext("request"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for request")
	}
	_request, _requestErr := DF1RequestMessageParse(readBuffer)
	if _requestErr != nil {
		return nil, errors.Wrap(_requestErr, "Error parsing 'request' field")
	}
	request := _request.(DF1RequestMessage)
	if closeErr := readBuffer.CloseContext("request"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for request")
	}

	if closeErr := readBuffer.CloseContext("CIPEncapsulationReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationReadRequest")
	}

	// Create a partially initialized instance
	_child := &_CIPEncapsulationReadRequest{
		Request:                 request,
		_CIPEncapsulationPacket: &_CIPEncapsulationPacket{},
	}
	_child._CIPEncapsulationPacket._CIPEncapsulationPacketChildRequirements = _child
	return _child, nil
}

func (m *_CIPEncapsulationReadRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationReadRequest")
		}

		// Simple Field (request)
		if pushErr := writeBuffer.PushContext("request"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for request")
		}
		_requestErr := writeBuffer.WriteSerializable(m.GetRequest())
		if popErr := writeBuffer.PopContext("request"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for request")
		}
		if _requestErr != nil {
			return errors.Wrap(_requestErr, "Error serializing 'request' field")
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CIPEncapsulationReadRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CIPEncapsulationReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
