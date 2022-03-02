/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type COTPPacket struct {
	Parameters []*COTPParameter
	Payload    *S7Message

	// Arguments.
	CotpLen uint16
	Child   ICOTPPacketChild
}

// The corresponding interface
type ICOTPPacket interface {
	// TpduCode returns TpduCode
	TpduCode() uint8
	// GetParameters returns Parameters
	GetParameters() []*COTPParameter
	// GetPayload returns Payload
	GetPayload() *S7Message
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ICOTPPacketParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ICOTPPacket, serializeChildFunction func() error) error
	GetTypeName() string
}

type ICOTPPacketChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message)
	GetTypeName() string
	ICOTPPacket
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *COTPPacket) GetParameters() []*COTPParameter {
	return m.Parameters
}

func (m *COTPPacket) GetPayload() *S7Message {
	return m.Payload
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCOTPPacket factory function for COTPPacket
func NewCOTPPacket(parameters []*COTPParameter, payload *S7Message, cotpLen uint16) *COTPPacket {
	return &COTPPacket{Parameters: parameters, Payload: payload, CotpLen: cotpLen}
}

func CastCOTPPacket(structType interface{}) *COTPPacket {
	castFunc := func(typ interface{}) *COTPPacket {
		if casted, ok := typ.(COTPPacket); ok {
			return &casted
		}
		if casted, ok := typ.(*COTPPacket); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *COTPPacket) GetTypeName() string {
	return "COTPPacket"
}

func (m *COTPPacket) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *COTPPacket) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *COTPPacket) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (headerLength)
	lengthInBits += 8
	// Discriminator Field (tpduCode)
	lengthInBits += 8

	// Array field
	if len(m.Parameters) > 0 {
		for _, element := range m.Parameters {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Optional Field (payload)
	if m.Payload != nil {
		lengthInBits += (*m.Payload).GetLengthInBits()
	}

	return lengthInBits
}

func (m *COTPPacket) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPPacketParse(readBuffer utils.ReadBuffer, cotpLen uint16) (*COTPPacket, error) {
	if pullErr := readBuffer.PullContext("COTPPacket"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos
	var startPos = readBuffer.GetPos()
	var curPos uint16

	// Implicit Field (headerLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	headerLength, _headerLengthErr := readBuffer.ReadUint8("headerLength", 8)
	_ = headerLength
	if _headerLengthErr != nil {
		return nil, errors.Wrap(_headerLengthErr, "Error parsing 'headerLength' field")
	}

	// Discriminator Field (tpduCode) (Used as input to a switch field)
	tpduCode, _tpduCodeErr := readBuffer.ReadUint8("tpduCode", 8)
	if _tpduCodeErr != nil {
		return nil, errors.Wrap(_tpduCodeErr, "Error parsing 'tpduCode' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *COTPPacket
	var typeSwitchError error
	switch {
	case tpduCode == 0xF0: // COTPPacketData
		_parent, typeSwitchError = COTPPacketDataParse(readBuffer, cotpLen)
	case tpduCode == 0xE0: // COTPPacketConnectionRequest
		_parent, typeSwitchError = COTPPacketConnectionRequestParse(readBuffer, cotpLen)
	case tpduCode == 0xD0: // COTPPacketConnectionResponse
		_parent, typeSwitchError = COTPPacketConnectionResponseParse(readBuffer, cotpLen)
	case tpduCode == 0x80: // COTPPacketDisconnectRequest
		_parent, typeSwitchError = COTPPacketDisconnectRequestParse(readBuffer, cotpLen)
	case tpduCode == 0xC0: // COTPPacketDisconnectResponse
		_parent, typeSwitchError = COTPPacketDisconnectResponseParse(readBuffer, cotpLen)
	case tpduCode == 0x70: // COTPPacketTpduError
		_parent, typeSwitchError = COTPPacketTpduErrorParse(readBuffer, cotpLen)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Array field (parameters)
	if pullErr := readBuffer.PullContext("parameters", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	curPos = readBuffer.GetPos() - startPos
	// Length array
	parameters := make([]*COTPParameter, 0)
	{
		_parametersLength := uint16(uint16(uint16(headerLength)+uint16(uint16(1)))) - uint16(curPos)
		_parametersEndPos := readBuffer.GetPos() + uint16(_parametersLength)
		for readBuffer.GetPos() < _parametersEndPos {
			_item, _err := COTPParameterParse(readBuffer, uint8(uint8(uint8(headerLength)+uint8(uint8(1))))-uint8(curPos))
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'parameters' field")
			}
			parameters = append(parameters, _item)
			curPos = readBuffer.GetPos() - startPos
		}
	}
	if closeErr := readBuffer.CloseContext("parameters", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (payload) (Can be skipped, if a given expression evaluates to false)
	curPos = readBuffer.GetPos() - startPos
	var payload *S7Message = nil
	if bool((curPos) < (cotpLen)) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := S7MessageParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'payload' field")
		default:
			payload = CastS7Message(_val)
			if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("COTPPacket"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, parameters, payload)
	return _parent, nil
}

func (m *COTPPacket) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *COTPPacket) SerializeParent(writeBuffer utils.WriteBuffer, child ICOTPPacket, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("COTPPacket"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	headerLength := uint8(uint8(uint8(m.GetLengthInBytes())) - uint8(uint8(uint8(uint8(utils.InlineIf(bool(bool((m.GetPayload()) != (nil))), func() interface{} { return uint8((*m.GetPayload()).GetLengthInBytes()) }, func() interface{} { return uint8(uint8(0)) }).(uint8)))+uint8(uint8(1)))))
	_headerLengthErr := writeBuffer.WriteUint8("headerLength", 8, (headerLength))
	if _headerLengthErr != nil {
		return errors.Wrap(_headerLengthErr, "Error serializing 'headerLength' field")
	}

	// Discriminator Field (tpduCode) (Used as input to a switch field)
	tpduCode := uint8(child.TpduCode())
	_tpduCodeErr := writeBuffer.WriteUint8("tpduCode", 8, (tpduCode))

	if _tpduCodeErr != nil {
		return errors.Wrap(_tpduCodeErr, "Error serializing 'tpduCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Array Field (parameters)
	if m.Parameters != nil {
		if pushErr := writeBuffer.PushContext("parameters", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.Parameters {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'parameters' field")
			}
		}
		if popErr := writeBuffer.PopContext("parameters", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	// Optional Field (payload) (Can be skipped, if the value is null)
	var payload *S7Message = nil
	if m.Payload != nil {
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return pushErr
		}
		payload = m.Payload
		_payloadErr := payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return popErr
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
	}

	if popErr := writeBuffer.PopContext("COTPPacket"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *COTPPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
