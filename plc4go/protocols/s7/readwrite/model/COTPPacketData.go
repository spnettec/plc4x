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


// COTPPacketData is the corresponding interface of COTPPacketData
type COTPPacketData interface {
	utils.LengthAware
	utils.Serializable
	COTPPacket
	// GetEot returns Eot (property field)
	GetEot() bool
	// GetTpduRef returns TpduRef (property field)
	GetTpduRef() uint8
}

// COTPPacketDataExactly can be used when we want exactly this type and not a type which fulfills COTPPacketData.
// This is useful for switch cases.
type COTPPacketDataExactly interface {
	COTPPacketData
	isCOTPPacketData() bool
}

// _COTPPacketData is the data-structure of this message
type _COTPPacketData struct {
	*_COTPPacket
        Eot bool
        TpduRef uint8
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketData)  GetTpduCode() uint8 {
return 0xF0}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketData) InitializeParent(parent COTPPacket , parameters []COTPParameter , payload S7Message ) {	m.Parameters = parameters
	m.Payload = payload
}

func (m *_COTPPacketData)  GetParent() COTPPacket {
	return m._COTPPacket
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketData) GetEot() bool {
	return m.Eot
}

func (m *_COTPPacketData) GetTpduRef() uint8 {
	return m.TpduRef
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCOTPPacketData factory function for _COTPPacketData
func NewCOTPPacketData( eot bool , tpduRef uint8 , parameters []COTPParameter , payload S7Message , cotpLen uint16 ) *_COTPPacketData {
	_result := &_COTPPacketData{
		Eot: eot,
		TpduRef: tpduRef,
    	_COTPPacket: NewCOTPPacket(parameters, payload, cotpLen),
	}
	_result._COTPPacket._COTPPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPPacketData(structType interface{}) COTPPacketData {
    if casted, ok := structType.(COTPPacketData); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketData); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketData) GetTypeName() string {
	return "COTPPacketData"
}

func (m *_COTPPacketData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_COTPPacketData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eot)
	lengthInBits += 1;

	// Simple field (tpduRef)
	lengthInBits += 7;

	return lengthInBits
}


func (m *_COTPPacketData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPPacketDataParse(theBytes []byte, cotpLen uint16) (COTPPacketData, error) {
	return COTPPacketDataParseWithBuffer(utils.NewReadBufferByteBased(theBytes), cotpLen)
}

func COTPPacketDataParseWithBuffer(readBuffer utils.ReadBuffer, cotpLen uint16) (COTPPacketData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eot)
_eot, _eotErr := readBuffer.ReadBit("eot")
	if _eotErr != nil {
		return nil, errors.Wrap(_eotErr, "Error parsing 'eot' field of COTPPacketData")
	}
	eot := _eot

	// Simple Field (tpduRef)
_tpduRef, _tpduRefErr := readBuffer.ReadUint8("tpduRef", 7)
	if _tpduRefErr != nil {
		return nil, errors.Wrap(_tpduRefErr, "Error parsing 'tpduRef' field of COTPPacketData")
	}
	tpduRef := _tpduRef

	if closeErr := readBuffer.CloseContext("COTPPacketData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketData")
	}

	// Create a partially initialized instance
	_child := &_COTPPacketData{
		_COTPPacket: &_COTPPacket{
			CotpLen: cotpLen,
		},
		Eot: eot,
		TpduRef: tpduRef,
	}
	_child._COTPPacket._COTPPacketChildRequirements = _child
	return _child, nil
}

func (m *_COTPPacketData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPPacketData) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketData")
		}

	// Simple Field (eot)
	eot := bool(m.GetEot())
	_eotErr := writeBuffer.WriteBit("eot", (eot))
	if _eotErr != nil {
		return errors.Wrap(_eotErr, "Error serializing 'eot' field")
	}

	// Simple Field (tpduRef)
	tpduRef := uint8(m.GetTpduRef())
	_tpduRefErr := writeBuffer.WriteUint8("tpduRef", 7, (tpduRef))
	if _tpduRefErr != nil {
		return errors.Wrap(_tpduRefErr, "Error serializing 'tpduRef' field")
	}

		if popErr := writeBuffer.PopContext("COTPPacketData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketData")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_COTPPacketData) isCOTPPacketData() bool {
	return true
}

func (m *_COTPPacketData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



