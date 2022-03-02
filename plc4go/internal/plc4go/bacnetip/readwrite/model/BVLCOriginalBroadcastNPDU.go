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
type BVLCOriginalBroadcastNPDU struct {
	*BVLC
	Npdu *NPDU

	// Arguments.
	BvlcPayloadLength uint16
}

// The corresponding interface
type IBVLCOriginalBroadcastNPDU interface {
	// GetNpdu returns Npdu
	GetNpdu() *NPDU
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
func (m *BVLCOriginalBroadcastNPDU) BvlcFunction() uint8 {
	return 0x0B
}

func (m *BVLCOriginalBroadcastNPDU) GetBvlcFunction() uint8 {
	return 0x0B
}

func (m *BVLCOriginalBroadcastNPDU) InitializeParent(parent *BVLC) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BVLCOriginalBroadcastNPDU) GetNpdu() *NPDU {
	return m.Npdu
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBVLCOriginalBroadcastNPDU factory function for BVLCOriginalBroadcastNPDU
func NewBVLCOriginalBroadcastNPDU(npdu *NPDU, bvlcPayloadLength uint16) *BVLC {
	child := &BVLCOriginalBroadcastNPDU{
		Npdu: npdu,
		BVLC: NewBVLC(),
	}
	child.Child = child
	return child.BVLC
}

func CastBVLCOriginalBroadcastNPDU(structType interface{}) *BVLCOriginalBroadcastNPDU {
	castFunc := func(typ interface{}) *BVLCOriginalBroadcastNPDU {
		if casted, ok := typ.(BVLCOriginalBroadcastNPDU); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCOriginalBroadcastNPDU); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCOriginalBroadcastNPDU(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCOriginalBroadcastNPDU(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCOriginalBroadcastNPDU) GetTypeName() string {
	return "BVLCOriginalBroadcastNPDU"
}

func (m *BVLCOriginalBroadcastNPDU) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BVLCOriginalBroadcastNPDU) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (npdu)
	lengthInBits += m.Npdu.GetLengthInBits()

	return lengthInBits
}

func (m *BVLCOriginalBroadcastNPDU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCOriginalBroadcastNPDUParse(readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCOriginalBroadcastNPDU"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (npdu)
	if pullErr := readBuffer.PullContext("npdu"); pullErr != nil {
		return nil, pullErr
	}
	_npdu, _npduErr := NPDUParse(readBuffer, uint16(bvlcPayloadLength))
	if _npduErr != nil {
		return nil, errors.Wrap(_npduErr, "Error parsing 'npdu' field")
	}
	npdu := CastNPDU(_npdu)
	if closeErr := readBuffer.CloseContext("npdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BVLCOriginalBroadcastNPDU"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCOriginalBroadcastNPDU{
		Npdu: CastNPDU(npdu),
		BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child.BVLC, nil
}

func (m *BVLCOriginalBroadcastNPDU) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCOriginalBroadcastNPDU"); pushErr != nil {
			return pushErr
		}

		// Simple Field (npdu)
		if pushErr := writeBuffer.PushContext("npdu"); pushErr != nil {
			return pushErr
		}
		_npduErr := m.Npdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("npdu"); popErr != nil {
			return popErr
		}
		if _npduErr != nil {
			return errors.Wrap(_npduErr, "Error serializing 'npdu' field")
		}

		if popErr := writeBuffer.PopContext("BVLCOriginalBroadcastNPDU"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCOriginalBroadcastNPDU) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
