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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaMessageResponse is the corresponding interface of OpcuaMessageResponse
type OpcuaMessageResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MessagePDU
	// GetSecurityHeader returns SecurityHeader (property field)
	GetSecurityHeader() SecurityHeader
	// GetMessage returns Message (property field)
	GetMessage() Payload
	// IsOpcuaMessageResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpcuaMessageResponse()
}

// _OpcuaMessageResponse is the data-structure of this message
type _OpcuaMessageResponse struct {
	MessagePDUContract
	SecurityHeader SecurityHeader
	Message        Payload

	// Arguments.
	TotalLength uint32
}

var _ OpcuaMessageResponse = (*_OpcuaMessageResponse)(nil)
var _ MessagePDURequirements = (*_OpcuaMessageResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaMessageResponse) GetMessageType() string {
	return "MSG"
}

func (m *_OpcuaMessageResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaMessageResponse) GetParent() MessagePDUContract {
	return m.MessagePDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaMessageResponse) GetSecurityHeader() SecurityHeader {
	return m.SecurityHeader
}

func (m *_OpcuaMessageResponse) GetMessage() Payload {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpcuaMessageResponse factory function for _OpcuaMessageResponse
func NewOpcuaMessageResponse(chunk ChunkType, securityHeader SecurityHeader, message Payload, totalLength uint32) *_OpcuaMessageResponse {
	if securityHeader == nil {
		panic("securityHeader of type SecurityHeader for OpcuaMessageResponse must not be nil")
	}
	if message == nil {
		panic("message of type Payload for OpcuaMessageResponse must not be nil")
	}
	_result := &_OpcuaMessageResponse{
		MessagePDUContract: NewMessagePDU(chunk),
		SecurityHeader:     securityHeader,
		Message:            message,
	}
	_result.MessagePDUContract.(*_MessagePDU)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpcuaMessageResponse(structType any) OpcuaMessageResponse {
	if casted, ok := structType.(OpcuaMessageResponse); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaMessageResponse); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaMessageResponse) GetTypeName() string {
	return "OpcuaMessageResponse"
}

func (m *_OpcuaMessageResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MessagePDUContract.(*_MessagePDU).getLengthInBits(ctx))

	// Simple field (securityHeader)
	lengthInBits += m.SecurityHeader.GetLengthInBits(ctx)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaMessageResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpcuaMessageResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MessagePDU, totalLength uint32, response bool) (__opcuaMessageResponse OpcuaMessageResponse, err error) {
	m.MessagePDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaMessageResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaMessageResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	securityHeader, err := ReadSimpleField[SecurityHeader](ctx, "securityHeader", ReadComplex[SecurityHeader](SecurityHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityHeader' field"))
	}
	m.SecurityHeader = securityHeader

	message, err := ReadSimpleField[Payload](ctx, "message", ReadComplex[Payload](PayloadParseWithBufferProducer[Payload]((bool)(bool(false)), (uint32)(uint32(uint32(totalLength)-uint32(securityHeader.GetLengthInBytes(ctx)))-uint32(uint32(16)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	if closeErr := readBuffer.CloseContext("OpcuaMessageResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaMessageResponse")
	}

	return m, nil
}

func (m *_OpcuaMessageResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaMessageResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaMessageResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaMessageResponse")
		}

		if err := WriteSimpleField[SecurityHeader](ctx, "securityHeader", m.GetSecurityHeader(), WriteComplex[SecurityHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityHeader' field")
		}

		if err := WriteSimpleField[Payload](ctx, "message", m.GetMessage(), WriteComplex[Payload](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaMessageResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaMessageResponse")
		}
		return nil
	}
	return m.MessagePDUContract.(*_MessagePDU).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_OpcuaMessageResponse) GetTotalLength() uint32 {
	return m.TotalLength
}

//
////

func (m *_OpcuaMessageResponse) IsOpcuaMessageResponse() {}

func (m *_OpcuaMessageResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
