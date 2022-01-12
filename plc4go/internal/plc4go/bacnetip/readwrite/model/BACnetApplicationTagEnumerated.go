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
type BACnetApplicationTagEnumerated struct {
	*BACnetApplicationTag
	Data []int8
}

// The corresponding interface
type IBACnetApplicationTagEnumerated interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetApplicationTagEnumerated) TagNumber() uint8 {
	return 0x9
}

func (m *BACnetApplicationTagEnumerated) InitializeParent(parent *BACnetApplicationTag, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isBoolean bool, isConstructed bool, isPrimitiveAndNotBoolean bool, actualLength uint32) {
	m.TagClass = tagClass
	m.LengthValueType = lengthValueType
	m.ExtTagNumber = extTagNumber
	m.ExtLength = extLength
	m.ExtExtLength = extExtLength
	m.ExtExtExtLength = extExtExtLength
}

func NewBACnetApplicationTagEnumerated(data []int8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isBoolean bool, isConstructed bool, isPrimitiveAndNotBoolean bool, actualLength uint32) *BACnetApplicationTag {
	child := &BACnetApplicationTagEnumerated{
		Data:                 data,
		BACnetApplicationTag: NewBACnetApplicationTag(tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength, actualTagNumber, isBoolean, isConstructed, isPrimitiveAndNotBoolean, actualLength),
	}
	child.Child = child
	return child.BACnetApplicationTag
}

func CastBACnetApplicationTagEnumerated(structType interface{}) *BACnetApplicationTagEnumerated {
	castFunc := func(typ interface{}) *BACnetApplicationTagEnumerated {
		if casted, ok := typ.(BACnetApplicationTagEnumerated); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetApplicationTagEnumerated); ok {
			return casted
		}
		if casted, ok := typ.(BACnetApplicationTag); ok {
			return CastBACnetApplicationTagEnumerated(casted.Child)
		}
		if casted, ok := typ.(*BACnetApplicationTag); ok {
			return CastBACnetApplicationTagEnumerated(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetApplicationTagEnumerated) GetTypeName() string {
	return "BACnetApplicationTagEnumerated"
}

func (m *BACnetApplicationTagEnumerated) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetApplicationTagEnumerated) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *BACnetApplicationTagEnumerated) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetApplicationTagEnumeratedParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetApplicationTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTagEnumerated"); pullErr != nil {
		return nil, pullErr
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	data := make([]int8, 0)
	{
		_dataLength := actualLength
		_dataEndPos := readBuffer.GetPos() + uint16(_dataLength)
		for readBuffer.GetPos() < _dataEndPos {
			_item, _err := readBuffer.ReadInt8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data = append(data, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagEnumerated"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagEnumerated{
		Data:                 data,
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child.BACnetApplicationTag, nil
}

func (m *BACnetApplicationTagEnumerated) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagEnumerated"); pushErr != nil {
			return pushErr
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

		if popErr := writeBuffer.PopContext("BACnetApplicationTagEnumerated"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
