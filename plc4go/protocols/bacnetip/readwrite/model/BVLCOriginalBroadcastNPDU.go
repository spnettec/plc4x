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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BVLCOriginalBroadcastNPDU is the corresponding interface of BVLCOriginalBroadcastNPDU
type BVLCOriginalBroadcastNPDU interface {
	utils.LengthAware
	utils.Serializable
	BVLC
	// GetNpdu returns Npdu (property field)
	GetNpdu() NPDU
}

// BVLCOriginalBroadcastNPDUExactly can be used when we want exactly this type and not a type which fulfills BVLCOriginalBroadcastNPDU.
// This is useful for switch cases.
type BVLCOriginalBroadcastNPDUExactly interface {
	BVLCOriginalBroadcastNPDU
	isBVLCOriginalBroadcastNPDU() bool
}

// _BVLCOriginalBroadcastNPDU is the data-structure of this message
type _BVLCOriginalBroadcastNPDU struct {
	*_BVLC
        Npdu NPDU

	// Arguments.
	BvlcPayloadLength uint16
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCOriginalBroadcastNPDU)  GetBvlcFunction() uint8 {
return 0x0B}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCOriginalBroadcastNPDU) InitializeParent(parent BVLC ) {}

func (m *_BVLCOriginalBroadcastNPDU)  GetParent() BVLC {
	return m._BVLC
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCOriginalBroadcastNPDU) GetNpdu() NPDU {
	return m.Npdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBVLCOriginalBroadcastNPDU factory function for _BVLCOriginalBroadcastNPDU
func NewBVLCOriginalBroadcastNPDU( npdu NPDU , bvlcPayloadLength uint16 ) *_BVLCOriginalBroadcastNPDU {
	_result := &_BVLCOriginalBroadcastNPDU{
		Npdu: npdu,
    	_BVLC: NewBVLC(),
	}
	_result._BVLC._BVLCChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCOriginalBroadcastNPDU(structType interface{}) BVLCOriginalBroadcastNPDU {
    if casted, ok := structType.(BVLCOriginalBroadcastNPDU); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCOriginalBroadcastNPDU); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCOriginalBroadcastNPDU) GetTypeName() string {
	return "BVLCOriginalBroadcastNPDU"
}

func (m *_BVLCOriginalBroadcastNPDU) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BVLCOriginalBroadcastNPDU) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (npdu)
	lengthInBits += m.Npdu.GetLengthInBits()

	return lengthInBits
}


func (m *_BVLCOriginalBroadcastNPDU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCOriginalBroadcastNPDUParse(readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (BVLCOriginalBroadcastNPDU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCOriginalBroadcastNPDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCOriginalBroadcastNPDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (npdu)
	if pullErr := readBuffer.PullContext("npdu"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for npdu")
	}
_npdu, _npduErr := NPDUParse(readBuffer , uint16( bvlcPayloadLength ) )
	if _npduErr != nil {
		return nil, errors.Wrap(_npduErr, "Error parsing 'npdu' field of BVLCOriginalBroadcastNPDU")
	}
	npdu := _npdu.(NPDU)
	if closeErr := readBuffer.CloseContext("npdu"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for npdu")
	}

	if closeErr := readBuffer.CloseContext("BVLCOriginalBroadcastNPDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCOriginalBroadcastNPDU")
	}

	// Create a partially initialized instance
	_child := &_BVLCOriginalBroadcastNPDU{
		_BVLC: &_BVLC{
		},
		Npdu: npdu,
	}
	_child._BVLC._BVLCChildRequirements = _child
	return _child, nil
}

func (m *_BVLCOriginalBroadcastNPDU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCOriginalBroadcastNPDU) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCOriginalBroadcastNPDU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCOriginalBroadcastNPDU")
		}

	// Simple Field (npdu)
	if pushErr := writeBuffer.PushContext("npdu"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for npdu")
	}
	_npduErr := writeBuffer.WriteSerializable(m.GetNpdu())
	if popErr := writeBuffer.PopContext("npdu"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for npdu")
	}
	if _npduErr != nil {
		return errors.Wrap(_npduErr, "Error serializing 'npdu' field")
	}

		if popErr := writeBuffer.PopContext("BVLCOriginalBroadcastNPDU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCOriginalBroadcastNPDU")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


////
// Arguments Getter

func (m *_BVLCOriginalBroadcastNPDU) GetBvlcPayloadLength() uint16 {
	return m.BvlcPayloadLength
}
//
////

func (m *_BVLCOriginalBroadcastNPDU) isBVLCOriginalBroadcastNPDU() bool {
	return true
}

func (m *_BVLCOriginalBroadcastNPDU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



