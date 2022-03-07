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
type IdentifyReplyCommandType struct {
	*IdentifyReplyCommand
	UnitType string
}

// The corresponding interface
type IIdentifyReplyCommandType interface {
	IIdentifyReplyCommand
	// GetUnitType returns UnitType (property field)
	GetUnitType() string
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
func (m *IdentifyReplyCommandType) Attribute() Attribute {
	return Attribute_Type
}

func (m *IdentifyReplyCommandType) GetAttribute() Attribute {
	return Attribute_Type
}

func (m *IdentifyReplyCommandType) InitializeParent(parent *IdentifyReplyCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *IdentifyReplyCommandType) GetUnitType() string {
	return m.UnitType
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandType factory function for IdentifyReplyCommandType
func NewIdentifyReplyCommandType(unitType string) *IdentifyReplyCommand {
	child := &IdentifyReplyCommandType{
		UnitType:             unitType,
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	child.Child = child
	return child.IdentifyReplyCommand
}

func CastIdentifyReplyCommandType(structType interface{}) *IdentifyReplyCommandType {
	if casted, ok := structType.(IdentifyReplyCommandType); ok {
		return &casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandType); ok {
		return casted
	}
	if casted, ok := structType.(IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandType(casted.Child)
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandType(casted.Child)
	}
	return nil
}

func (m *IdentifyReplyCommandType) GetTypeName() string {
	return "IdentifyReplyCommandType"
}

func (m *IdentifyReplyCommandType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (unitType)
	lengthInBits += 64

	return lengthInBits
}

func (m *IdentifyReplyCommandType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandTypeParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommand, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (unitType)
	_unitType, _unitTypeErr := readBuffer.ReadString("unitType", uint32(64))
	if _unitTypeErr != nil {
		return nil, errors.Wrap(_unitTypeErr, "Error parsing 'unitType' field")
	}
	unitType := _unitType

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandType{
		UnitType:             unitType,
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child.IdentifyReplyCommand, nil
}

func (m *IdentifyReplyCommandType) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandType"); pushErr != nil {
			return pushErr
		}

		// Simple Field (unitType)
		unitType := string(m.UnitType)
		_unitTypeErr := writeBuffer.WriteString("unitType", uint32(64), "UTF-8", (unitType))
		if _unitTypeErr != nil {
			return errors.Wrap(_unitTypeErr, "Error serializing 'unitType' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
