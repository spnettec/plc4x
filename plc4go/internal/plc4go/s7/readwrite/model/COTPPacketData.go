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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type COTPPacketData struct {
	*COTPPacket
	Eot     bool
	TpduRef uint8

	// Arguments.
	CotpLen uint16
}

// The corresponding interface
type ICOTPPacketData interface {
	// GetEot returns Eot
	GetEot() bool
	// GetTpduRef returns TpduRef
	GetTpduRef() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPPacketData) TpduCode() uint8 {
	return 0xF0
}

func (m *COTPPacketData) GetTpduCode() uint8 {
	return 0xF0
}

func (m *COTPPacketData) InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message) {
	m.COTPPacket.Parameters = parameters
	m.COTPPacket.Payload = payload
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *COTPPacketData) GetEot() bool {
	return m.Eot
}

func (m *COTPPacketData) GetTpduRef() uint8 {
	return m.TpduRef
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCOTPPacketData factory function for COTPPacketData
func NewCOTPPacketData(eot bool, tpduRef uint8, parameters []*COTPParameter, payload *S7Message, cotpLen uint16) *COTPPacket {
	child := &COTPPacketData{
		Eot:        eot,
		TpduRef:    tpduRef,
		COTPPacket: NewCOTPPacket(parameters, payload, cotpLen),
	}
	child.Child = child
	return child.COTPPacket
}

func CastCOTPPacketData(structType interface{}) *COTPPacketData {
	castFunc := func(typ interface{}) *COTPPacketData {
		if casted, ok := typ.(COTPPacketData); ok {
			return &casted
		}
		if casted, ok := typ.(*COTPPacketData); ok {
			return casted
		}
		if casted, ok := typ.(COTPPacket); ok {
			return CastCOTPPacketData(casted.Child)
		}
		if casted, ok := typ.(*COTPPacket); ok {
			return CastCOTPPacketData(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *COTPPacketData) GetTypeName() string {
	return "COTPPacketData"
}

func (m *COTPPacketData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *COTPPacketData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eot)
	lengthInBits += 1

	// Simple field (tpduRef)
	lengthInBits += 7

	return lengthInBits
}

func (m *COTPPacketData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPPacketDataParse(readBuffer utils.ReadBuffer, cotpLen uint16) (*COTPPacket, error) {
	if pullErr := readBuffer.PullContext("COTPPacketData"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (eot)
	_eot, _eotErr := readBuffer.ReadBit("eot")
	if _eotErr != nil {
		return nil, errors.Wrap(_eotErr, "Error parsing 'eot' field")
	}
	eot := _eot

	// Simple Field (tpduRef)
	_tpduRef, _tpduRefErr := readBuffer.ReadUint8("tpduRef", 7)
	if _tpduRefErr != nil {
		return nil, errors.Wrap(_tpduRefErr, "Error parsing 'tpduRef' field")
	}
	tpduRef := _tpduRef

	if closeErr := readBuffer.CloseContext("COTPPacketData"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &COTPPacketData{
		Eot:        eot,
		TpduRef:    tpduRef,
		COTPPacket: &COTPPacket{},
	}
	_child.COTPPacket.Child = _child
	return _child.COTPPacket, nil
}

func (m *COTPPacketData) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketData"); pushErr != nil {
			return pushErr
		}

		// Simple Field (eot)
		eot := bool(m.Eot)
		_eotErr := writeBuffer.WriteBit("eot", (eot))
		if _eotErr != nil {
			return errors.Wrap(_eotErr, "Error serializing 'eot' field")
		}

		// Simple Field (tpduRef)
		tpduRef := uint8(m.TpduRef)
		_tpduRefErr := writeBuffer.WriteUint8("tpduRef", 7, (tpduRef))
		if _tpduRefErr != nil {
			return errors.Wrap(_tpduRefErr, "Error serializing 'tpduRef' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketData"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *COTPPacketData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
