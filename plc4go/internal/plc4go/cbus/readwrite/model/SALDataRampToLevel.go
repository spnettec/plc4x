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

// SALDataRampToLevel is the data-structure of this message
type SALDataRampToLevel struct {
	*SALData
	Group byte
	Level byte
}

// ISALDataRampToLevel is the corresponding interface of SALDataRampToLevel
type ISALDataRampToLevel interface {
	ISALData
	// GetGroup returns Group (property field)
	GetGroup() byte
	// GetLevel returns Level (property field)
	GetLevel() byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *SALDataRampToLevel) InitializeParent(parent *SALData, commandTypeContainer SALCommandTypeContainer) {
	m.SALData.CommandTypeContainer = commandTypeContainer
}

func (m *SALDataRampToLevel) GetParent() *SALData {
	return m.SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *SALDataRampToLevel) GetGroup() byte {
	return m.Group
}

func (m *SALDataRampToLevel) GetLevel() byte {
	return m.Level
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataRampToLevel factory function for SALDataRampToLevel
func NewSALDataRampToLevel(group byte, level byte, commandTypeContainer SALCommandTypeContainer) *SALDataRampToLevel {
	_result := &SALDataRampToLevel{
		Group:   group,
		Level:   level,
		SALData: NewSALData(commandTypeContainer),
	}
	_result.Child = _result
	return _result
}

func CastSALDataRampToLevel(structType interface{}) *SALDataRampToLevel {
	if casted, ok := structType.(SALDataRampToLevel); ok {
		return &casted
	}
	if casted, ok := structType.(*SALDataRampToLevel); ok {
		return casted
	}
	if casted, ok := structType.(SALData); ok {
		return CastSALDataRampToLevel(casted.Child)
	}
	if casted, ok := structType.(*SALData); ok {
		return CastSALDataRampToLevel(casted.Child)
	}
	return nil
}

func (m *SALDataRampToLevel) GetTypeName() string {
	return "SALDataRampToLevel"
}

func (m *SALDataRampToLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SALDataRampToLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (group)
	lengthInBits += 8

	// Simple field (level)
	lengthInBits += 8

	return lengthInBits
}

func (m *SALDataRampToLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataRampToLevelParse(readBuffer utils.ReadBuffer) (*SALDataRampToLevel, error) {
	if pullErr := readBuffer.PullContext("SALDataRampToLevel"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (group)
	_group, _groupErr := readBuffer.ReadByte("group")
	if _groupErr != nil {
		return nil, errors.Wrap(_groupErr, "Error parsing 'group' field")
	}
	group := _group

	// Simple Field (level)
	_level, _levelErr := readBuffer.ReadByte("level")
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field")
	}
	level := _level

	if closeErr := readBuffer.CloseContext("SALDataRampToLevel"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SALDataRampToLevel{
		Group:   group,
		Level:   level,
		SALData: &SALData{},
	}
	_child.SALData.Child = _child
	return _child, nil
}

func (m *SALDataRampToLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataRampToLevel"); pushErr != nil {
			return pushErr
		}

		// Simple Field (group)
		group := byte(m.Group)
		_groupErr := writeBuffer.WriteByte("group", (group))
		if _groupErr != nil {
			return errors.Wrap(_groupErr, "Error serializing 'group' field")
		}

		// Simple Field (level)
		level := byte(m.Level)
		_levelErr := writeBuffer.WriteByte("level", (level))
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
		}

		if popErr := writeBuffer.PopContext("SALDataRampToLevel"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SALDataRampToLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
