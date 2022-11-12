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


// KnxGroupAddress2Level is the corresponding interface of KnxGroupAddress2Level
type KnxGroupAddress2Level interface {
	utils.LengthAware
	utils.Serializable
	KnxGroupAddress
	// GetMainGroup returns MainGroup (property field)
	GetMainGroup() uint8
	// GetSubGroup returns SubGroup (property field)
	GetSubGroup() uint16
}

// KnxGroupAddress2LevelExactly can be used when we want exactly this type and not a type which fulfills KnxGroupAddress2Level.
// This is useful for switch cases.
type KnxGroupAddress2LevelExactly interface {
	KnxGroupAddress2Level
	isKnxGroupAddress2Level() bool
}

// _KnxGroupAddress2Level is the data-structure of this message
type _KnxGroupAddress2Level struct {
	*_KnxGroupAddress
        MainGroup uint8
        SubGroup uint16
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxGroupAddress2Level)  GetNumLevels() uint8 {
return uint8(2)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxGroupAddress2Level) InitializeParent(parent KnxGroupAddress ) {}

func (m *_KnxGroupAddress2Level)  GetParent() KnxGroupAddress {
	return m._KnxGroupAddress
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxGroupAddress2Level) GetMainGroup() uint8 {
	return m.MainGroup
}

func (m *_KnxGroupAddress2Level) GetSubGroup() uint16 {
	return m.SubGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewKnxGroupAddress2Level factory function for _KnxGroupAddress2Level
func NewKnxGroupAddress2Level( mainGroup uint8 , subGroup uint16 ) *_KnxGroupAddress2Level {
	_result := &_KnxGroupAddress2Level{
		MainGroup: mainGroup,
		SubGroup: subGroup,
    	_KnxGroupAddress: NewKnxGroupAddress(),
	}
	_result._KnxGroupAddress._KnxGroupAddressChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxGroupAddress2Level(structType interface{}) KnxGroupAddress2Level {
    if casted, ok := structType.(KnxGroupAddress2Level); ok {
		return casted
	}
	if casted, ok := structType.(*KnxGroupAddress2Level); ok {
		return *casted
	}
	return nil
}

func (m *_KnxGroupAddress2Level) GetTypeName() string {
	return "KnxGroupAddress2Level"
}

func (m *_KnxGroupAddress2Level) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_KnxGroupAddress2Level) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (mainGroup)
	lengthInBits += 5;

	// Simple field (subGroup)
	lengthInBits += 11;

	return lengthInBits
}


func (m *_KnxGroupAddress2Level) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxGroupAddress2LevelParse(theBytes []byte, numLevels uint8) (KnxGroupAddress2Level, error) {
	return KnxGroupAddress2LevelParseWithBuffer(utils.NewReadBufferByteBased(theBytes), numLevels)
}

func KnxGroupAddress2LevelParseWithBuffer(readBuffer utils.ReadBuffer, numLevels uint8) (KnxGroupAddress2Level, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxGroupAddress2Level"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxGroupAddress2Level")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (mainGroup)
_mainGroup, _mainGroupErr := readBuffer.ReadUint8("mainGroup", 5)
	if _mainGroupErr != nil {
		return nil, errors.Wrap(_mainGroupErr, "Error parsing 'mainGroup' field of KnxGroupAddress2Level")
	}
	mainGroup := _mainGroup

	// Simple Field (subGroup)
_subGroup, _subGroupErr := readBuffer.ReadUint16("subGroup", 11)
	if _subGroupErr != nil {
		return nil, errors.Wrap(_subGroupErr, "Error parsing 'subGroup' field of KnxGroupAddress2Level")
	}
	subGroup := _subGroup

	if closeErr := readBuffer.CloseContext("KnxGroupAddress2Level"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxGroupAddress2Level")
	}

	// Create a partially initialized instance
	_child := &_KnxGroupAddress2Level{
		_KnxGroupAddress: &_KnxGroupAddress{
		},
		MainGroup: mainGroup,
		SubGroup: subGroup,
	}
	_child._KnxGroupAddress._KnxGroupAddressChildRequirements = _child
	return _child, nil
}

func (m *_KnxGroupAddress2Level) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxGroupAddress2Level) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxGroupAddress2Level"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxGroupAddress2Level")
		}

	// Simple Field (mainGroup)
	mainGroup := uint8(m.GetMainGroup())
	_mainGroupErr := writeBuffer.WriteUint8("mainGroup", 5, (mainGroup))
	if _mainGroupErr != nil {
		return errors.Wrap(_mainGroupErr, "Error serializing 'mainGroup' field")
	}

	// Simple Field (subGroup)
	subGroup := uint16(m.GetSubGroup())
	_subGroupErr := writeBuffer.WriteUint16("subGroup", 11, (subGroup))
	if _subGroupErr != nil {
		return errors.Wrap(_subGroupErr, "Error serializing 'subGroup' field")
	}

		if popErr := writeBuffer.PopContext("KnxGroupAddress2Level"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxGroupAddress2Level")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_KnxGroupAddress2Level) isKnxGroupAddress2Level() bool {
	return true
}

func (m *_KnxGroupAddress2Level) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



