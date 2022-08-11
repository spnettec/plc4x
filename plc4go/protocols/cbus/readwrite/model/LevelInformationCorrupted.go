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

// LevelInformationCorrupted is the corresponding interface of LevelInformationCorrupted
type LevelInformationCorrupted interface {
	utils.LengthAware
	utils.Serializable
	LevelInformation
	// GetCorruptedNibble1 returns CorruptedNibble1 (property field)
	GetCorruptedNibble1() uint8
	// GetCorruptedNibble2 returns CorruptedNibble2 (property field)
	GetCorruptedNibble2() uint8
	// GetCorruptedNibble3 returns CorruptedNibble3 (property field)
	GetCorruptedNibble3() uint8
	// GetCorruptedNibble4 returns CorruptedNibble4 (property field)
	GetCorruptedNibble4() uint8
}

// LevelInformationCorruptedExactly can be used when we want exactly this type and not a type which fulfills LevelInformationCorrupted.
// This is useful for switch cases.
type LevelInformationCorruptedExactly interface {
	LevelInformationCorrupted
	isLevelInformationCorrupted() bool
}

// _LevelInformationCorrupted is the data-structure of this message
type _LevelInformationCorrupted struct {
	*_LevelInformation
	CorruptedNibble1 uint8
	CorruptedNibble2 uint8
	CorruptedNibble3 uint8
	CorruptedNibble4 uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LevelInformationCorrupted) InitializeParent(parent LevelInformation, raw uint16) {
	m.Raw = raw
}

func (m *_LevelInformationCorrupted) GetParent() LevelInformation {
	return m._LevelInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LevelInformationCorrupted) GetCorruptedNibble1() uint8 {
	return m.CorruptedNibble1
}

func (m *_LevelInformationCorrupted) GetCorruptedNibble2() uint8 {
	return m.CorruptedNibble2
}

func (m *_LevelInformationCorrupted) GetCorruptedNibble3() uint8 {
	return m.CorruptedNibble3
}

func (m *_LevelInformationCorrupted) GetCorruptedNibble4() uint8 {
	return m.CorruptedNibble4
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLevelInformationCorrupted factory function for _LevelInformationCorrupted
func NewLevelInformationCorrupted(corruptedNibble1 uint8, corruptedNibble2 uint8, corruptedNibble3 uint8, corruptedNibble4 uint8, raw uint16) *_LevelInformationCorrupted {
	_result := &_LevelInformationCorrupted{
		CorruptedNibble1:  corruptedNibble1,
		CorruptedNibble2:  corruptedNibble2,
		CorruptedNibble3:  corruptedNibble3,
		CorruptedNibble4:  corruptedNibble4,
		_LevelInformation: NewLevelInformation(raw),
	}
	_result._LevelInformation._LevelInformationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLevelInformationCorrupted(structType interface{}) LevelInformationCorrupted {
	if casted, ok := structType.(LevelInformationCorrupted); ok {
		return casted
	}
	if casted, ok := structType.(*LevelInformationCorrupted); ok {
		return *casted
	}
	return nil
}

func (m *_LevelInformationCorrupted) GetTypeName() string {
	return "LevelInformationCorrupted"
}

func (m *_LevelInformationCorrupted) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_LevelInformationCorrupted) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (corruptedNibble1)
	lengthInBits += 4

	// Simple field (corruptedNibble2)
	lengthInBits += 4

	// Simple field (corruptedNibble3)
	lengthInBits += 4

	// Simple field (corruptedNibble4)
	lengthInBits += 4

	return lengthInBits
}

func (m *_LevelInformationCorrupted) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LevelInformationCorruptedParse(readBuffer utils.ReadBuffer) (LevelInformationCorrupted, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LevelInformationCorrupted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LevelInformationCorrupted")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (corruptedNibble1)
	_corruptedNibble1, _corruptedNibble1Err := readBuffer.ReadUint8("corruptedNibble1", 4)
	if _corruptedNibble1Err != nil {
		return nil, errors.Wrap(_corruptedNibble1Err, "Error parsing 'corruptedNibble1' field of LevelInformationCorrupted")
	}
	corruptedNibble1 := _corruptedNibble1

	// Simple Field (corruptedNibble2)
	_corruptedNibble2, _corruptedNibble2Err := readBuffer.ReadUint8("corruptedNibble2", 4)
	if _corruptedNibble2Err != nil {
		return nil, errors.Wrap(_corruptedNibble2Err, "Error parsing 'corruptedNibble2' field of LevelInformationCorrupted")
	}
	corruptedNibble2 := _corruptedNibble2

	// Simple Field (corruptedNibble3)
	_corruptedNibble3, _corruptedNibble3Err := readBuffer.ReadUint8("corruptedNibble3", 4)
	if _corruptedNibble3Err != nil {
		return nil, errors.Wrap(_corruptedNibble3Err, "Error parsing 'corruptedNibble3' field of LevelInformationCorrupted")
	}
	corruptedNibble3 := _corruptedNibble3

	// Simple Field (corruptedNibble4)
	_corruptedNibble4, _corruptedNibble4Err := readBuffer.ReadUint8("corruptedNibble4", 4)
	if _corruptedNibble4Err != nil {
		return nil, errors.Wrap(_corruptedNibble4Err, "Error parsing 'corruptedNibble4' field of LevelInformationCorrupted")
	}
	corruptedNibble4 := _corruptedNibble4

	if closeErr := readBuffer.CloseContext("LevelInformationCorrupted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LevelInformationCorrupted")
	}

	// Create a partially initialized instance
	_child := &_LevelInformationCorrupted{
		_LevelInformation: &_LevelInformation{},
		CorruptedNibble1:  corruptedNibble1,
		CorruptedNibble2:  corruptedNibble2,
		CorruptedNibble3:  corruptedNibble3,
		CorruptedNibble4:  corruptedNibble4,
	}
	_child._LevelInformation._LevelInformationChildRequirements = _child
	return _child, nil
}

func (m *_LevelInformationCorrupted) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LevelInformationCorrupted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LevelInformationCorrupted")
		}

		// Simple Field (corruptedNibble1)
		corruptedNibble1 := uint8(m.GetCorruptedNibble1())
		_corruptedNibble1Err := writeBuffer.WriteUint8("corruptedNibble1", 4, (corruptedNibble1))
		if _corruptedNibble1Err != nil {
			return errors.Wrap(_corruptedNibble1Err, "Error serializing 'corruptedNibble1' field")
		}

		// Simple Field (corruptedNibble2)
		corruptedNibble2 := uint8(m.GetCorruptedNibble2())
		_corruptedNibble2Err := writeBuffer.WriteUint8("corruptedNibble2", 4, (corruptedNibble2))
		if _corruptedNibble2Err != nil {
			return errors.Wrap(_corruptedNibble2Err, "Error serializing 'corruptedNibble2' field")
		}

		// Simple Field (corruptedNibble3)
		corruptedNibble3 := uint8(m.GetCorruptedNibble3())
		_corruptedNibble3Err := writeBuffer.WriteUint8("corruptedNibble3", 4, (corruptedNibble3))
		if _corruptedNibble3Err != nil {
			return errors.Wrap(_corruptedNibble3Err, "Error serializing 'corruptedNibble3' field")
		}

		// Simple Field (corruptedNibble4)
		corruptedNibble4 := uint8(m.GetCorruptedNibble4())
		_corruptedNibble4Err := writeBuffer.WriteUint8("corruptedNibble4", 4, (corruptedNibble4))
		if _corruptedNibble4Err != nil {
			return errors.Wrap(_corruptedNibble4Err, "Error serializing 'corruptedNibble4' field")
		}

		if popErr := writeBuffer.PopContext("LevelInformationCorrupted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LevelInformationCorrupted")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_LevelInformationCorrupted) isLevelInformationCorrupted() bool {
	return true
}

func (m *_LevelInformationCorrupted) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
