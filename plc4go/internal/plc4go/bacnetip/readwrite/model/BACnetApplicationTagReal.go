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
type BACnetApplicationTagReal struct {
	*BACnetApplicationTag
	Value float32
}

// The corresponding interface
type IBACnetApplicationTagReal interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetApplicationTagReal) TagNumber() uint8 {
	return 0x4
}

func (m *BACnetApplicationTagReal) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader, tagNumber uint8, actualLength uint32) {
	m.Header = header
}

func NewBACnetApplicationTagReal(value float32, header *BACnetTagHeader, tagNumber uint8, actualLength uint32) *BACnetApplicationTag {
	child := &BACnetApplicationTagReal{
		Value:                value,
		BACnetApplicationTag: NewBACnetApplicationTag(header, tagNumber, actualLength),
	}
	child.Child = child
	return child.BACnetApplicationTag
}

func CastBACnetApplicationTagReal(structType interface{}) *BACnetApplicationTagReal {
	castFunc := func(typ interface{}) *BACnetApplicationTagReal {
		if casted, ok := typ.(BACnetApplicationTagReal); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetApplicationTagReal); ok {
			return casted
		}
		if casted, ok := typ.(BACnetApplicationTag); ok {
			return CastBACnetApplicationTagReal(casted.Child)
		}
		if casted, ok := typ.(*BACnetApplicationTag); ok {
			return CastBACnetApplicationTagReal(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetApplicationTagReal) GetTypeName() string {
	return "BACnetApplicationTagReal"
}

func (m *BACnetApplicationTagReal) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetApplicationTagReal) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (value)
	lengthInBits += 32

	return lengthInBits
}

func (m *BACnetApplicationTagReal) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetApplicationTagRealParse(readBuffer utils.ReadBuffer) (*BACnetApplicationTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTagReal"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadFloat32("value", 32)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagReal"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagReal{
		Value:                value,
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child.BACnetApplicationTag, nil
}

func (m *BACnetApplicationTagReal) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagReal"); pushErr != nil {
			return pushErr
		}

		// Simple Field (value)
		value := float32(m.Value)
		_valueErr := writeBuffer.WriteFloat32("value", 32, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagReal"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagReal) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
