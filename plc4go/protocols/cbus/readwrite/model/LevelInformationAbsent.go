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


// LevelInformationAbsent is the corresponding interface of LevelInformationAbsent
type LevelInformationAbsent interface {
	utils.LengthAware
	utils.Serializable
	LevelInformation
}

// LevelInformationAbsentExactly can be used when we want exactly this type and not a type which fulfills LevelInformationAbsent.
// This is useful for switch cases.
type LevelInformationAbsentExactly interface {
	LevelInformationAbsent
	isLevelInformationAbsent() bool
}

// _LevelInformationAbsent is the data-structure of this message
type _LevelInformationAbsent struct {
	*_LevelInformation
	// Reserved Fields
	reservedField0 *uint16
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LevelInformationAbsent) InitializeParent(parent LevelInformation , raw uint16 ) {	m.Raw = raw
}

func (m *_LevelInformationAbsent)  GetParent() LevelInformation {
	return m._LevelInformation
}


// NewLevelInformationAbsent factory function for _LevelInformationAbsent
func NewLevelInformationAbsent( raw uint16 ) *_LevelInformationAbsent {
	_result := &_LevelInformationAbsent{
    	_LevelInformation: NewLevelInformation(raw),
	}
	_result._LevelInformation._LevelInformationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLevelInformationAbsent(structType interface{}) LevelInformationAbsent {
    if casted, ok := structType.(LevelInformationAbsent); ok {
		return casted
	}
	if casted, ok := structType.(*LevelInformationAbsent); ok {
		return *casted
	}
	return nil
}

func (m *_LevelInformationAbsent) GetTypeName() string {
	return "LevelInformationAbsent"
}

func (m *_LevelInformationAbsent) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_LevelInformationAbsent) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 16

	return lengthInBits
}


func (m *_LevelInformationAbsent) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LevelInformationAbsentParse(readBuffer utils.ReadBuffer) (LevelInformationAbsent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LevelInformationAbsent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LevelInformationAbsent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint16
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of LevelInformationAbsent")
		}
		if reserved != uint16(0x0000) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("LevelInformationAbsent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LevelInformationAbsent")
	}

	// Create a partially initialized instance
	_child := &_LevelInformationAbsent{
		_LevelInformation: &_LevelInformation{
		},
		reservedField0: reservedField0,
	}
	_child._LevelInformation._LevelInformationChildRequirements = _child
	return _child, nil
}

func (m *_LevelInformationAbsent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LevelInformationAbsent) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LevelInformationAbsent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LevelInformationAbsent")
		}

	// Reserved Field (reserved)
	{
		var reserved uint16 = uint16(0x0000)
		if m.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteUint16("reserved", 16, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

		if popErr := writeBuffer.PopContext("LevelInformationAbsent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LevelInformationAbsent")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_LevelInformationAbsent) isLevelInformationAbsent() bool {
	return true
}

func (m *_LevelInformationAbsent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



