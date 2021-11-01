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
type BACnetComplexTagOctetString struct {
	TheString         string
	ActualLengthInBit uint16
	Parent            *BACnetComplexTag
}

// The corresponding interface
type IBACnetComplexTagOctetString interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetComplexTagOctetString) DataType() BACnetDataType {
	return BACnetDataType_OCTET_STRING
}

func (m *BACnetComplexTagOctetString) InitializeParent(parent *BACnetComplexTag, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isPrimitiveAndNotBoolean bool, actualLength uint32) {
	m.Parent.TagNumber = tagNumber
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
	m.Parent.ExtExtLength = extExtLength
	m.Parent.ExtExtExtLength = extExtExtLength
}

func NewBACnetComplexTagOctetString(theString string, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetComplexTag {
	child := &BACnetComplexTagOctetString{
		TheString: theString,
		Parent:    NewBACnetComplexTag(tagNumber, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetComplexTagOctetString(structType interface{}) *BACnetComplexTagOctetString {
	castFunc := func(typ interface{}) *BACnetComplexTagOctetString {
		if casted, ok := typ.(BACnetComplexTagOctetString); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetComplexTagOctetString); ok {
			return casted
		}
		if casted, ok := typ.(BACnetComplexTag); ok {
			return CastBACnetComplexTagOctetString(casted.Child)
		}
		if casted, ok := typ.(*BACnetComplexTag); ok {
			return CastBACnetComplexTagOctetString(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetComplexTagOctetString) GetTypeName() string {
	return "BACnetComplexTagOctetString"
}

func (m *BACnetComplexTagOctetString) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetComplexTagOctetString) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Simple field (theString)
	lengthInBits += uint16(m.ActualLengthInBit)

	return lengthInBits
}

func (m *BACnetComplexTagOctetString) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetComplexTagOctetStringParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetComplexTag, error) {
	if pullErr := readBuffer.PullContext("BACnetComplexTagOctetString"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	actualLengthInBit := uint16(actualLength) * uint16(uint16(8))

	// Simple Field (theString)
	theString, _theStringErr := readBuffer.ReadString("theString", uint32(-1))
	if _theStringErr != nil {
		return nil, errors.Wrap(_theStringErr, "Error parsing 'theString' field")
	}

	if closeErr := readBuffer.CloseContext("BACnetComplexTagOctetString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetComplexTagOctetString{
		TheString: theString,
		Parent:    &BACnetComplexTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetComplexTagOctetString) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetComplexTagOctetString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (theString)
		theString := string(m.TheString)
		_theStringErr := writeBuffer.WriteString("theString", uint8(-1), "ASCII", (theString))
		if _theStringErr != nil {
			return errors.Wrap(_theStringErr, "Error serializing 'theString' field")
		}

		if popErr := writeBuffer.PopContext("BACnetComplexTagOctetString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetComplexTagOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
