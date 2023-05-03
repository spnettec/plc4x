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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const S7Message_PROTOCOLID uint8 = 0x32

// S7Message is the corresponding interface of S7Message
type S7Message interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetTpduReference returns TpduReference (property field)
	GetTpduReference() uint16
	// GetParameter returns Parameter (property field)
	GetParameter() S7Parameter
	// GetPayload returns Payload (property field)
	GetPayload() S7Payload
}

// S7MessageExactly can be used when we want exactly this type and not a type which fulfills S7Message.
// This is useful for switch cases.
type S7MessageExactly interface {
	S7Message
	isS7Message() bool
}

// _S7Message is the data-structure of this message
type _S7Message struct {
	_S7MessageChildRequirements
	TpduReference uint16
	Parameter     S7Parameter
	Payload       S7Payload
	// Reserved Fields
	reservedField0 *uint16
}

type _S7MessageChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetMessageType() uint8
}

type S7MessageParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7Message, serializeChildFunction func() error) error
	GetTypeName() string
}

type S7MessageChild interface {
	utils.Serializable
	InitializeParent(parent S7Message, tpduReference uint16, parameter S7Parameter, payload S7Payload)
	GetParent() *S7Message

	GetTypeName() string
	S7Message
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7Message) GetTpduReference() uint16 {
	return m.TpduReference
}

func (m *_S7Message) GetParameter() S7Parameter {
	return m.Parameter
}

func (m *_S7Message) GetPayload() S7Payload {
	return m.Payload
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7Message) GetProtocolId() uint8 {
	return S7Message_PROTOCOLID
}

///////////////////////-4
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7Message factory function for _S7Message
func NewS7Message(tpduReference uint16, parameter S7Parameter, payload S7Payload) *_S7Message {
	return &_S7Message{TpduReference: tpduReference, Parameter: parameter, Payload: payload}
}

// Deprecated: use the interface for direct cast
func CastS7Message(structType any) S7Message {
	if casted, ok := structType.(S7Message); ok {
		return casted
	}
	if casted, ok := structType.(*S7Message); ok {
		return *casted
	}
	return nil
}

func (m *_S7Message) GetTypeName() string {
	return "S7Message"
}

func (m *_S7Message) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (protocolId)
	lengthInBits += 8
	// Discriminator Field (messageType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (tpduReference)
	lengthInBits += 16

	// Implicit Field (parameterLength)
	lengthInBits += 16

	// Implicit Field (payloadLength)
	lengthInBits += 16

	// Optional Field (parameter)
	if m.Parameter != nil {
		lengthInBits += m.Parameter.GetLengthInBits(ctx)
	}

	// Optional Field (payload)
	if m.Payload != nil {
		lengthInBits += m.Payload.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_S7Message) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7MessageParse(theBytes []byte) (S7Message, error) {
	return S7MessageParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func S7MessageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (S7Message, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7Message"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7Message")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (protocolId)
	protocolId, _protocolIdErr := readBuffer.ReadUint8("protocolId", 8)
	if _protocolIdErr != nil {
		return nil, errors.Wrap(_protocolIdErr, "Error parsing 'protocolId' field of S7Message")
	}
	if protocolId != S7Message_PROTOCOLID {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7Message_PROTOCOLID) + " but got " + fmt.Sprintf("%d", protocolId))
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType, _messageTypeErr := readBuffer.ReadUint8("messageType", 8)
	if _messageTypeErr != nil {
		return nil, errors.Wrap(_messageTypeErr, "Error parsing 'messageType' field of S7Message")
	}

	var reservedField0 *uint16
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of S7Message")
		}
		if reserved != uint16(0x0000) {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (tpduReference)
	_tpduReference, _tpduReferenceErr := readBuffer.ReadUint16("tpduReference", 16)
	if _tpduReferenceErr != nil {
		return nil, errors.Wrap(_tpduReferenceErr, "Error parsing 'tpduReference' field of S7Message")
	}
	tpduReference := _tpduReference

	// Implicit Field (parameterLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	parameterLength, _parameterLengthErr := readBuffer.ReadUint16("parameterLength", 16)
	_ = parameterLength
	if _parameterLengthErr != nil {
		return nil, errors.Wrap(_parameterLengthErr, "Error parsing 'parameterLength' field of S7Message")
	}

	// Implicit Field (payloadLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	payloadLength, _payloadLengthErr := readBuffer.ReadUint16("payloadLength", 16)
	_ = payloadLength
	if _payloadLengthErr != nil {
		return nil, errors.Wrap(_payloadLengthErr, "Error parsing 'payloadLength' field of S7Message")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7MessageChildSerializeRequirement interface {
		S7Message
		InitializeParent(S7Message, uint16, S7Parameter, S7Payload)
		GetParent() S7Message
	}
	var _childTemp any
	var _child S7MessageChildSerializeRequirement
	var typeSwitchError error
	switch {
	case messageType == 0x01: // S7MessageRequest
		_childTemp, typeSwitchError = S7MessageRequestParseWithBuffer(ctx, readBuffer)
	case messageType == 0x02: // S7MessageResponse
		_childTemp, typeSwitchError = S7MessageResponseParseWithBuffer(ctx, readBuffer)
	case messageType == 0x03: // S7MessageResponseData
		_childTemp, typeSwitchError = S7MessageResponseDataParseWithBuffer(ctx, readBuffer)
	case messageType == 0x07: // S7MessageUserData
		_childTemp, typeSwitchError = S7MessageUserDataParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [messageType=%v]", messageType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of S7Message")
	}
	_child = _childTemp.(S7MessageChildSerializeRequirement)

	// Optional Field (parameter) (Can be skipped, if a given expression evaluates to false)
	var parameter S7Parameter = nil
	if bool((parameterLength) > (0)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("parameter"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for parameter")
		}
		_val, _err := S7ParameterParseWithBuffer(ctx, readBuffer, messageType)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'parameter' field of S7Message")
		default:
			parameter = _val.(S7Parameter)
			if closeErr := readBuffer.CloseContext("parameter"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for parameter")
			}
		}
	}

	// Optional Field (payload) (Can be skipped, if a given expression evaluates to false)
	var payload S7Payload = nil
	if bool((payloadLength) > (0)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for payload")
		}
		_val, _err := S7PayloadParseWithBuffer(ctx, readBuffer, messageType, (parameter))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'payload' field of S7Message")
		default:
			payload = _val.(S7Payload)
			if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for payload")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("S7Message"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7Message")
	}

	// Finish initializing
	_child.InitializeParent(_child, tpduReference, parameter, payload)
	_child.GetParent().(*_S7Message).reservedField0 = reservedField0
	return _child, nil
}

func (pm *_S7Message) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7Message, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("S7Message"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7Message")
	}

	// Const Field (protocolId)
	_protocolIdErr := writeBuffer.WriteUint8("protocolId", 8, 0x32)
	if _protocolIdErr != nil {
		return errors.Wrap(_protocolIdErr, "Error serializing 'protocolId' field")
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := uint8(child.GetMessageType())
	_messageTypeErr := writeBuffer.WriteUint8("messageType", 8, (messageType))

	if _messageTypeErr != nil {
		return errors.Wrap(_messageTypeErr, "Error serializing 'messageType' field")
	}

	// Reserved Field (reserved)
	{
		var reserved uint16 = uint16(0x0000)
		if pm.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *pm.reservedField0
		}
		_err := writeBuffer.WriteUint16("reserved", 16, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (tpduReference)
	tpduReference := uint16(m.GetTpduReference())
	_tpduReferenceErr := writeBuffer.WriteUint16("tpduReference", 16, (tpduReference))
	if _tpduReferenceErr != nil {
		return errors.Wrap(_tpduReferenceErr, "Error serializing 'tpduReference' field")
	}

	// Implicit Field (parameterLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	parameterLength := uint16(utils.InlineIf(bool((m.GetParameter()) != (nil)), func() any { return uint16((m.GetParameter()).GetLengthInBytes(ctx)) }, func() any { return uint16(uint16(0)) }).(uint16))
	_parameterLengthErr := writeBuffer.WriteUint16("parameterLength", 16, (parameterLength))
	if _parameterLengthErr != nil {
		return errors.Wrap(_parameterLengthErr, "Error serializing 'parameterLength' field")
	}

	// Implicit Field (payloadLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	payloadLength := uint16(utils.InlineIf(bool((m.GetPayload()) != (nil)), func() any { return uint16((m.GetPayload()).GetLengthInBytes(ctx)) }, func() any { return uint16(uint16(0)) }).(uint16))
	_payloadLengthErr := writeBuffer.WriteUint16("payloadLength", 16, (payloadLength))
	if _payloadLengthErr != nil {
		return errors.Wrap(_payloadLengthErr, "Error serializing 'payloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Optional Field (parameter) (Can be skipped, if the value is null)
	var parameter S7Parameter = nil
	if m.GetParameter() != nil {
		if pushErr := writeBuffer.PushContext("parameter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for parameter")
		}
		parameter = m.GetParameter()
		_parameterErr := writeBuffer.WriteSerializable(ctx, parameter)
		if popErr := writeBuffer.PopContext("parameter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for parameter")
		}
		if _parameterErr != nil {
			return errors.Wrap(_parameterErr, "Error serializing 'parameter' field")
		}
	}

	// Optional Field (payload) (Can be skipped, if the value is null)
	var payload S7Payload = nil
	if m.GetPayload() != nil {
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		payload = m.GetPayload()
		_payloadErr := writeBuffer.WriteSerializable(ctx, payload)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
	}

	if popErr := writeBuffer.PopContext("S7Message"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7Message")
	}
	return nil
}

func (m *_S7Message) isS7Message() bool {
	return true
}

func (m *_S7Message) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
