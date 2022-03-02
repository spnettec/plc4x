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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type CALDataRequestReset struct {
	*CALData
}

// The corresponding interface
type ICALDataRequestReset interface {
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
func (m *CALDataRequestReset) CommandType() CALCommandType {
	return CALCommandType_RESET
}

func (m *CALDataRequestReset) GetCommandType() CALCommandType {
	return CALCommandType_RESET
}

func (m *CALDataRequestReset) InitializeParent(parent *CALData, commandTypeContainer CALCommandTypeContainer) {
	m.CALData.CommandTypeContainer = commandTypeContainer
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCALDataRequestReset factory function for CALDataRequestReset
func NewCALDataRequestReset(commandTypeContainer CALCommandTypeContainer) *CALData {
	child := &CALDataRequestReset{
		CALData: NewCALData(commandTypeContainer),
	}
	child.Child = child
	return child.CALData
}

func CastCALDataRequestReset(structType interface{}) *CALDataRequestReset {
	castFunc := func(typ interface{}) *CALDataRequestReset {
		if casted, ok := typ.(CALDataRequestReset); ok {
			return &casted
		}
		if casted, ok := typ.(*CALDataRequestReset); ok {
			return casted
		}
		if casted, ok := typ.(CALData); ok {
			return CastCALDataRequestReset(casted.Child)
		}
		if casted, ok := typ.(*CALData); ok {
			return CastCALDataRequestReset(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CALDataRequestReset) GetTypeName() string {
	return "CALDataRequestReset"
}

func (m *CALDataRequestReset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALDataRequestReset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *CALDataRequestReset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataRequestResetParse(readBuffer utils.ReadBuffer) (*CALData, error) {
	if pullErr := readBuffer.PullContext("CALDataRequestReset"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CALDataRequestReset"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CALDataRequestReset{
		CALData: &CALData{},
	}
	_child.CALData.Child = _child
	return _child.CALData, nil
}

func (m *CALDataRequestReset) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataRequestReset"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("CALDataRequestReset"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CALDataRequestReset) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
