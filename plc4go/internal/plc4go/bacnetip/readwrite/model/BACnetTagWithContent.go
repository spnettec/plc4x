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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const BACnetTagWithContent_OPENTAG uint8 = 0x2e
const BACnetTagWithContent_CLOSINGTAG uint8 = 0x2f

// The data-structure of this message
type BACnetTagWithContent struct {
	TagNumber          uint8
	TagClass           TagClass
	LengthValueType    uint8
	ExtTagNumber       *uint8
	ExtLength          *uint8
	PropertyIdentifier []uint8
	Value              *BACnetTag
}

// The corresponding interface
type IBACnetTagWithContent interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewBACnetTagWithContent(tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, propertyIdentifier []uint8, value *BACnetTag) *BACnetTagWithContent {
	return &BACnetTagWithContent{TagNumber: tagNumber, TagClass: tagClass, LengthValueType: lengthValueType, ExtTagNumber: extTagNumber, ExtLength: extLength, PropertyIdentifier: propertyIdentifier, Value: value}
}

func CastBACnetTagWithContent(structType interface{}) *BACnetTagWithContent {
	castFunc := func(typ interface{}) *BACnetTagWithContent {
		if casted, ok := typ.(BACnetTagWithContent); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagWithContent); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagWithContent) GetTypeName() string {
	return "BACnetTagWithContent"
}

func (m *BACnetTagWithContent) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTagWithContent) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (tagNumber)
	lengthInBits += 4

	// Simple field (tagClass)
	lengthInBits += 1

	// Simple field (lengthValueType)
	lengthInBits += 3

	// Optional Field (extTagNumber)
	if m.ExtTagNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (extLength)
	if m.ExtLength != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.PropertyIdentifier) > 0 {
		lengthInBits += 8 * uint16(len(m.PropertyIdentifier))
	}

	// Const Field (openTag)
	lengthInBits += 8

	// Simple field (value)
	lengthInBits += m.Value.LengthInBits()

	// Const Field (closingTag)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetTagWithContent) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagWithContentParse(readBuffer utils.ReadBuffer) (*BACnetTagWithContent, error) {
	if pullErr := readBuffer.PullContext("BACnetTagWithContent"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (tagNumber)
	tagNumber, _tagNumberErr := readBuffer.ReadUint8("tagNumber", 4)
	if _tagNumberErr != nil {
		return nil, errors.Wrap(_tagNumberErr, "Error parsing 'tagNumber' field")
	}

	// Simple Field (tagClass)
	if pullErr := readBuffer.PullContext("tagClass"); pullErr != nil {
		return nil, pullErr
	}
	tagClass, _tagClassErr := TagClassParse(readBuffer)
	if _tagClassErr != nil {
		return nil, errors.Wrap(_tagClassErr, "Error parsing 'tagClass' field")
	}
	if closeErr := readBuffer.CloseContext("tagClass"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (lengthValueType)
	lengthValueType, _lengthValueTypeErr := readBuffer.ReadUint8("lengthValueType", 3)
	if _lengthValueTypeErr != nil {
		return nil, errors.Wrap(_lengthValueTypeErr, "Error parsing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if a given expression evaluates to false)
	var extTagNumber *uint8 = nil
	if bool((tagNumber) == (15)) {
		_val, _err := readBuffer.ReadUint8("extTagNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extTagNumber' field")
		}
		extTagNumber = &_val
	}

	// Optional Field (extLength) (Can be skipped, if a given expression evaluates to false)
	var extLength *uint8 = nil
	if bool((lengthValueType) == (5)) {
		_val, _err := readBuffer.ReadUint8("extLength", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extLength' field")
		}
		extLength = &_val
	}

	// Array field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	propertyIdentifier := make([]uint8, 0)
	_propertyIdentifierLength := utils.InlineIf(bool(bool((lengthValueType) == (5))), func() interface{} { return uint16((*extLength)) }, func() interface{} { return uint16(lengthValueType) }).(uint16)
	_propertyIdentifierEndPos := readBuffer.GetPos() + uint16(_propertyIdentifierLength)
	for readBuffer.GetPos() < _propertyIdentifierEndPos {
		_item, _err := readBuffer.ReadUint8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'propertyIdentifier' field")
		}
		propertyIdentifier = append(propertyIdentifier, _item)
	}
	if closeErr := readBuffer.CloseContext("propertyIdentifier", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Const Field (openTag)
	openTag, _openTagErr := readBuffer.ReadUint8("openTag", 8)
	if _openTagErr != nil {
		return nil, errors.Wrap(_openTagErr, "Error parsing 'openTag' field")
	}
	if openTag != BACnetTagWithContent_OPENTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetTagWithContent_OPENTAG) + " but got " + fmt.Sprintf("%d", openTag))
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, pullErr
	}
	value, _valueErr := BACnetTagParse(readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, closeErr
	}

	// Const Field (closingTag)
	closingTag, _closingTagErr := readBuffer.ReadUint8("closingTag", 8)
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	if closingTag != BACnetTagWithContent_CLOSINGTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetTagWithContent_CLOSINGTAG) + " but got " + fmt.Sprintf("%d", closingTag))
	}

	if closeErr := readBuffer.CloseContext("BACnetTagWithContent"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagWithContent(tagNumber, tagClass, lengthValueType, extTagNumber, extLength, propertyIdentifier, value), nil
}

func (m *BACnetTagWithContent) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTagWithContent"); pushErr != nil {
		return pushErr
	}

	// Simple Field (tagNumber)
	tagNumber := uint8(m.TagNumber)
	_tagNumberErr := writeBuffer.WriteUint8("tagNumber", 4, (tagNumber))
	if _tagNumberErr != nil {
		return errors.Wrap(_tagNumberErr, "Error serializing 'tagNumber' field")
	}

	// Simple Field (tagClass)
	if pushErr := writeBuffer.PushContext("tagClass"); pushErr != nil {
		return pushErr
	}
	_tagClassErr := m.TagClass.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("tagClass"); popErr != nil {
		return popErr
	}
	if _tagClassErr != nil {
		return errors.Wrap(_tagClassErr, "Error serializing 'tagClass' field")
	}

	// Simple Field (lengthValueType)
	lengthValueType := uint8(m.LengthValueType)
	_lengthValueTypeErr := writeBuffer.WriteUint8("lengthValueType", 3, (lengthValueType))
	if _lengthValueTypeErr != nil {
		return errors.Wrap(_lengthValueTypeErr, "Error serializing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if the value is null)
	var extTagNumber *uint8 = nil
	if m.ExtTagNumber != nil {
		extTagNumber = m.ExtTagNumber
		_extTagNumberErr := writeBuffer.WriteUint8("extTagNumber", 8, *(extTagNumber))
		if _extTagNumberErr != nil {
			return errors.Wrap(_extTagNumberErr, "Error serializing 'extTagNumber' field")
		}
	}

	// Optional Field (extLength) (Can be skipped, if the value is null)
	var extLength *uint8 = nil
	if m.ExtLength != nil {
		extLength = m.ExtLength
		_extLengthErr := writeBuffer.WriteUint8("extLength", 8, *(extLength))
		if _extLengthErr != nil {
			return errors.Wrap(_extLengthErr, "Error serializing 'extLength' field")
		}
	}

	// Array Field (propertyIdentifier)
	if m.PropertyIdentifier != nil {
		if pushErr := writeBuffer.PushContext("propertyIdentifier", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.PropertyIdentifier {
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'propertyIdentifier' field")
			}
		}
		if popErr := writeBuffer.PopContext("propertyIdentifier", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	// Const Field (openTag)
	_openTagErr := writeBuffer.WriteUint8("openTag", 8, 0x2e)
	if _openTagErr != nil {
		return errors.Wrap(_openTagErr, "Error serializing 'openTag' field")
	}

	// Simple Field (value)
	if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
		return pushErr
	}
	_valueErr := m.Value.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("value"); popErr != nil {
		return popErr
	}
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	// Const Field (closingTag)
	_closingTagErr := writeBuffer.WriteUint8("closingTag", 8, 0x2f)
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagWithContent"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagWithContent) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
