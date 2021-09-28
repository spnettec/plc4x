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
type BACnetComplexTagBitString struct {
	UnusedBits uint8
	Data       []int8
	Parent     *BACnetComplexTag
}

// The corresponding interface
type IBACnetComplexTagBitString interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetComplexTagBitString) DataType() BACnetDataType {
	return BACnetDataType_BIT_STRING
}

func (m *BACnetComplexTagBitString) InitializeParent(parent *BACnetComplexTag, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isPrimitiveAndNotBoolean bool, actualLength uint32) {
	m.Parent.TagNumber = tagNumber
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
	m.Parent.ExtExtLength = extExtLength
	m.Parent.ExtExtExtLength = extExtExtLength
}

func NewBACnetComplexTagBitString(unusedBits uint8, data []int8, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetComplexTag {
	child := &BACnetComplexTagBitString{
		UnusedBits: unusedBits,
		Data:       data,
		Parent:     NewBACnetComplexTag(tagNumber, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetComplexTagBitString(structType interface{}) *BACnetComplexTagBitString {
	castFunc := func(typ interface{}) *BACnetComplexTagBitString {
		if casted, ok := typ.(BACnetComplexTagBitString); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetComplexTagBitString); ok {
			return casted
		}
		if casted, ok := typ.(BACnetComplexTag); ok {
			return CastBACnetComplexTagBitString(casted.Child)
		}
		if casted, ok := typ.(*BACnetComplexTag); ok {
			return CastBACnetComplexTagBitString(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetComplexTagBitString) GetTypeName() string {
	return "BACnetComplexTagBitString"
}

func (m *BACnetComplexTagBitString) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetComplexTagBitString) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (unusedBits)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *BACnetComplexTagBitString) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetComplexTagBitStringParse(readBuffer utils.ReadBuffer, lengthValueType uint8, extLength uint8) (*BACnetComplexTag, error) {
	if pullErr := readBuffer.PullContext("BACnetComplexTagBitString"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (unusedBits)
	unusedBits, _unusedBitsErr := readBuffer.ReadUint8("unusedBits", 8)
	if _unusedBitsErr != nil {
		return nil, errors.Wrap(_unusedBitsErr, "Error parsing 'unusedBits' field")
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	data := make([]int8, 0)
	_dataLength := utils.InlineIf(bool(bool((lengthValueType) == (5))), func() interface{} { return uint16(uint16(uint16(extLength) - uint16(uint16(1)))) }, func() interface{} { return uint16(uint16(uint16(lengthValueType) - uint16(uint16(1)))) }).(uint16)
	_dataEndPos := readBuffer.GetPos() + uint16(_dataLength)
	for readBuffer.GetPos() < _dataEndPos {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data = append(data, _item)
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetComplexTagBitString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetComplexTagBitString{
		UnusedBits: unusedBits,
		Data:       data,
		Parent:     &BACnetComplexTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetComplexTagBitString) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetComplexTagBitString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (unusedBits)
		unusedBits := uint8(m.UnusedBits)
		_unusedBitsErr := writeBuffer.WriteUint8("unusedBits", 8, (unusedBits))
		if _unusedBitsErr != nil {
			return errors.Wrap(_unusedBitsErr, "Error serializing 'unusedBits' field")
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetComplexTagBitString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetComplexTagBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
