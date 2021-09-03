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
const BACnetServiceAckReadProperty_OBJECTIDENTIFIERHEADER uint8 = 0x0C
const BACnetServiceAckReadProperty_PROPERTYIDENTIFIERHEADER uint8 = 0x03
const BACnetServiceAckReadProperty_OPENINGTAG uint8 = 0x3E
const BACnetServiceAckReadProperty_CLOSINGTAG uint8 = 0x3F

// The data-structure of this message
type BACnetServiceAckReadProperty struct {
	ObjectType               uint16
	ObjectInstanceNumber     uint32
	PropertyIdentifierLength uint8
	PropertyIdentifier       []int8
	Value                    *BACnetTag
	Parent                   *BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckReadProperty interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckReadProperty) ServiceChoice() uint8 {
	return 0x0C
}

func (m *BACnetServiceAckReadProperty) InitializeParent(parent *BACnetServiceAck) {
}

func NewBACnetServiceAckReadProperty(objectType uint16, objectInstanceNumber uint32, propertyIdentifierLength uint8, propertyIdentifier []int8, value *BACnetTag) *BACnetServiceAck {
	child := &BACnetServiceAckReadProperty{
		ObjectType:               objectType,
		ObjectInstanceNumber:     objectInstanceNumber,
		PropertyIdentifierLength: propertyIdentifierLength,
		PropertyIdentifier:       propertyIdentifier,
		Value:                    value,
		Parent:                   NewBACnetServiceAck(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetServiceAckReadProperty(structType interface{}) *BACnetServiceAckReadProperty {
	castFunc := func(typ interface{}) *BACnetServiceAckReadProperty {
		if casted, ok := typ.(BACnetServiceAckReadProperty); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckReadProperty); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAck); ok {
			return CastBACnetServiceAckReadProperty(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return CastBACnetServiceAckReadProperty(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckReadProperty) GetTypeName() string {
	return "BACnetServiceAckReadProperty"
}

func (m *BACnetServiceAckReadProperty) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetServiceAckReadProperty) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Const Field (objectIdentifierHeader)
	lengthInBits += 8

	// Simple field (objectType)
	lengthInBits += 10

	// Simple field (objectInstanceNumber)
	lengthInBits += 22

	// Const Field (propertyIdentifierHeader)
	lengthInBits += 5

	// Simple field (propertyIdentifierLength)
	lengthInBits += 3

	// Array field
	if len(m.PropertyIdentifier) > 0 {
		lengthInBits += 8 * uint16(len(m.PropertyIdentifier))
	}

	// Const Field (openingTag)
	lengthInBits += 8

	// Simple field (value)
	lengthInBits += m.Value.LengthInBits()

	// Const Field (closingTag)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetServiceAckReadProperty) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckReadPropertyParse(readBuffer utils.ReadBuffer) (*BACnetServiceAck, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadProperty"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (objectIdentifierHeader)
	objectIdentifierHeader, _objectIdentifierHeaderErr := readBuffer.ReadUint8("objectIdentifierHeader", 8)
	if _objectIdentifierHeaderErr != nil {
		return nil, errors.Wrap(_objectIdentifierHeaderErr, "Error parsing 'objectIdentifierHeader' field")
	}
	if objectIdentifierHeader != BACnetServiceAckReadProperty_OBJECTIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetServiceAckReadProperty_OBJECTIDENTIFIERHEADER) + " but got " + fmt.Sprintf("%d", objectIdentifierHeader))
	}

	// Simple Field (objectType)
	objectType, _objectTypeErr := readBuffer.ReadUint16("objectType", 10)
	if _objectTypeErr != nil {
		return nil, errors.Wrap(_objectTypeErr, "Error parsing 'objectType' field")
	}

	// Simple Field (objectInstanceNumber)
	objectInstanceNumber, _objectInstanceNumberErr := readBuffer.ReadUint32("objectInstanceNumber", 22)
	if _objectInstanceNumberErr != nil {
		return nil, errors.Wrap(_objectInstanceNumberErr, "Error parsing 'objectInstanceNumber' field")
	}

	// Const Field (propertyIdentifierHeader)
	propertyIdentifierHeader, _propertyIdentifierHeaderErr := readBuffer.ReadUint8("propertyIdentifierHeader", 5)
	if _propertyIdentifierHeaderErr != nil {
		return nil, errors.Wrap(_propertyIdentifierHeaderErr, "Error parsing 'propertyIdentifierHeader' field")
	}
	if propertyIdentifierHeader != BACnetServiceAckReadProperty_PROPERTYIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetServiceAckReadProperty_PROPERTYIDENTIFIERHEADER) + " but got " + fmt.Sprintf("%d", propertyIdentifierHeader))
	}

	// Simple Field (propertyIdentifierLength)
	propertyIdentifierLength, _propertyIdentifierLengthErr := readBuffer.ReadUint8("propertyIdentifierLength", 3)
	if _propertyIdentifierLengthErr != nil {
		return nil, errors.Wrap(_propertyIdentifierLengthErr, "Error parsing 'propertyIdentifierLength' field")
	}

	// Array field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	propertyIdentifier := make([]int8, propertyIdentifierLength)
	for curItem := uint16(0); curItem < uint16(propertyIdentifierLength); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'propertyIdentifier' field")
		}
		propertyIdentifier[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("propertyIdentifier", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Const Field (openingTag)
	openingTag, _openingTagErr := readBuffer.ReadUint8("openingTag", 8)
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	if openingTag != BACnetServiceAckReadProperty_OPENINGTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetServiceAckReadProperty_OPENINGTAG) + " but got " + fmt.Sprintf("%d", openingTag))
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
	if closingTag != BACnetServiceAckReadProperty_CLOSINGTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetServiceAckReadProperty_CLOSINGTAG) + " but got " + fmt.Sprintf("%d", closingTag))
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadProperty"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckReadProperty{
		ObjectType:               objectType,
		ObjectInstanceNumber:     objectInstanceNumber,
		PropertyIdentifierLength: propertyIdentifierLength,
		PropertyIdentifier:       propertyIdentifier,
		Value:                    value,
		Parent:                   &BACnetServiceAck{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetServiceAckReadProperty) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadProperty"); pushErr != nil {
			return pushErr
		}

		// Const Field (objectIdentifierHeader)
		_objectIdentifierHeaderErr := writeBuffer.WriteUint8("objectIdentifierHeader", 8, 0x0C)
		if _objectIdentifierHeaderErr != nil {
			return errors.Wrap(_objectIdentifierHeaderErr, "Error serializing 'objectIdentifierHeader' field")
		}

		// Simple Field (objectType)
		objectType := uint16(m.ObjectType)
		_objectTypeErr := writeBuffer.WriteUint16("objectType", 10, (objectType))
		if _objectTypeErr != nil {
			return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
		}

		// Simple Field (objectInstanceNumber)
		objectInstanceNumber := uint32(m.ObjectInstanceNumber)
		_objectInstanceNumberErr := writeBuffer.WriteUint32("objectInstanceNumber", 22, (objectInstanceNumber))
		if _objectInstanceNumberErr != nil {
			return errors.Wrap(_objectInstanceNumberErr, "Error serializing 'objectInstanceNumber' field")
		}

		// Const Field (propertyIdentifierHeader)
		_propertyIdentifierHeaderErr := writeBuffer.WriteUint8("propertyIdentifierHeader", 5, 0x03)
		if _propertyIdentifierHeaderErr != nil {
			return errors.Wrap(_propertyIdentifierHeaderErr, "Error serializing 'propertyIdentifierHeader' field")
		}

		// Simple Field (propertyIdentifierLength)
		propertyIdentifierLength := uint8(m.PropertyIdentifierLength)
		_propertyIdentifierLengthErr := writeBuffer.WriteUint8("propertyIdentifierLength", 3, (propertyIdentifierLength))
		if _propertyIdentifierLengthErr != nil {
			return errors.Wrap(_propertyIdentifierLengthErr, "Error serializing 'propertyIdentifierLength' field")
		}

		// Array Field (propertyIdentifier)
		if m.PropertyIdentifier != nil {
			if pushErr := writeBuffer.PushContext("propertyIdentifier", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.PropertyIdentifier {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'propertyIdentifier' field")
				}
			}
			if popErr := writeBuffer.PopContext("propertyIdentifier", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		// Const Field (openingTag)
		_openingTagErr := writeBuffer.WriteUint8("openingTag", 8, 0x3E)
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
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
		_closingTagErr := writeBuffer.WriteUint8("closingTag", 8, 0x3F)
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadProperty"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckReadProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
