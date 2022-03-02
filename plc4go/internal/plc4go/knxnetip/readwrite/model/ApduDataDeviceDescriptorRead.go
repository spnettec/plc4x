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
type ApduDataDeviceDescriptorRead struct {
	*ApduData
	DescriptorType uint8

	// Arguments.
	DataLength uint8
}

// The corresponding interface
type IApduDataDeviceDescriptorRead interface {
	// GetDescriptorType returns DescriptorType
	GetDescriptorType() uint8
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
func (m *ApduDataDeviceDescriptorRead) ApciType() uint8 {
	return 0xC
}

func (m *ApduDataDeviceDescriptorRead) GetApciType() uint8 {
	return 0xC
}

func (m *ApduDataDeviceDescriptorRead) InitializeParent(parent *ApduData) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ApduDataDeviceDescriptorRead) GetDescriptorType() uint8 {
	return m.DescriptorType
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewApduDataDeviceDescriptorRead factory function for ApduDataDeviceDescriptorRead
func NewApduDataDeviceDescriptorRead(descriptorType uint8, dataLength uint8) *ApduData {
	child := &ApduDataDeviceDescriptorRead{
		DescriptorType: descriptorType,
		ApduData:       NewApduData(dataLength),
	}
	child.Child = child
	return child.ApduData
}

func CastApduDataDeviceDescriptorRead(structType interface{}) *ApduDataDeviceDescriptorRead {
	castFunc := func(typ interface{}) *ApduDataDeviceDescriptorRead {
		if casted, ok := typ.(ApduDataDeviceDescriptorRead); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataDeviceDescriptorRead); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataDeviceDescriptorRead(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataDeviceDescriptorRead(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataDeviceDescriptorRead) GetTypeName() string {
	return "ApduDataDeviceDescriptorRead"
}

func (m *ApduDataDeviceDescriptorRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataDeviceDescriptorRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (descriptorType)
	lengthInBits += 6

	return lengthInBits
}

func (m *ApduDataDeviceDescriptorRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataDeviceDescriptorReadParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataDeviceDescriptorRead"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (descriptorType)
	_descriptorType, _descriptorTypeErr := readBuffer.ReadUint8("descriptorType", 6)
	if _descriptorTypeErr != nil {
		return nil, errors.Wrap(_descriptorTypeErr, "Error parsing 'descriptorType' field")
	}
	descriptorType := _descriptorType

	if closeErr := readBuffer.CloseContext("ApduDataDeviceDescriptorRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataDeviceDescriptorRead{
		DescriptorType: descriptorType,
		ApduData:       &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child.ApduData, nil
}

func (m *ApduDataDeviceDescriptorRead) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataDeviceDescriptorRead"); pushErr != nil {
			return pushErr
		}

		// Simple Field (descriptorType)
		descriptorType := uint8(m.DescriptorType)
		_descriptorTypeErr := writeBuffer.WriteUint8("descriptorType", 6, (descriptorType))
		if _descriptorTypeErr != nil {
			return errors.Wrap(_descriptorTypeErr, "Error serializing 'descriptorType' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataDeviceDescriptorRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataDeviceDescriptorRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
