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
type IdentifyReplyCommandGAVPhysicalAddresses struct {
	*IdentifyReplyCommand
	Values []byte
}

// The corresponding interface
type IIdentifyReplyCommandGAVPhysicalAddresses interface {
	// GetValues returns Values
	GetValues() []byte
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
func (m *IdentifyReplyCommandGAVPhysicalAddresses) Attribute() Attribute {
	return Attribute_GAVPhysicalAddresses
}

func (m *IdentifyReplyCommandGAVPhysicalAddresses) GetAttribute() Attribute {
	return Attribute_GAVPhysicalAddresses
}

func (m *IdentifyReplyCommandGAVPhysicalAddresses) InitializeParent(parent *IdentifyReplyCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *IdentifyReplyCommandGAVPhysicalAddresses) GetValues() []byte {
	return m.Values
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandGAVPhysicalAddresses factory function for IdentifyReplyCommandGAVPhysicalAddresses
func NewIdentifyReplyCommandGAVPhysicalAddresses(values []byte) *IdentifyReplyCommand {
	child := &IdentifyReplyCommandGAVPhysicalAddresses{
		Values:               values,
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	child.Child = child
	return child.IdentifyReplyCommand
}

func CastIdentifyReplyCommandGAVPhysicalAddresses(structType interface{}) *IdentifyReplyCommandGAVPhysicalAddresses {
	castFunc := func(typ interface{}) *IdentifyReplyCommandGAVPhysicalAddresses {
		if casted, ok := typ.(IdentifyReplyCommandGAVPhysicalAddresses); ok {
			return &casted
		}
		if casted, ok := typ.(*IdentifyReplyCommandGAVPhysicalAddresses); ok {
			return casted
		}
		if casted, ok := typ.(IdentifyReplyCommand); ok {
			return CastIdentifyReplyCommandGAVPhysicalAddresses(casted.Child)
		}
		if casted, ok := typ.(*IdentifyReplyCommand); ok {
			return CastIdentifyReplyCommandGAVPhysicalAddresses(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *IdentifyReplyCommandGAVPhysicalAddresses) GetTypeName() string {
	return "IdentifyReplyCommandGAVPhysicalAddresses"
}

func (m *IdentifyReplyCommandGAVPhysicalAddresses) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandGAVPhysicalAddresses) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Values) > 0 {
		lengthInBits += 8 * uint16(len(m.Values))
	}

	return lengthInBits
}

func (m *IdentifyReplyCommandGAVPhysicalAddresses) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandGAVPhysicalAddressesParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommand, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandGAVPhysicalAddresses"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos
	// Byte Array field (values)
	numberOfBytesvalues := int(uint16(16))
	values, _readArrayErr := readBuffer.ReadByteArray("values", numberOfBytesvalues)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'values' field")
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandGAVPhysicalAddresses"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandGAVPhysicalAddresses{
		Values:               values,
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child.IdentifyReplyCommand, nil
}

func (m *IdentifyReplyCommandGAVPhysicalAddresses) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandGAVPhysicalAddresses"); pushErr != nil {
			return pushErr
		}

		// Array Field (values)
		if m.Values != nil {
			// Byte Array field (values)
			_writeArrayErr := writeBuffer.WriteByteArray("values", m.Values)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'values' field")
			}
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandGAVPhysicalAddresses"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandGAVPhysicalAddresses) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
