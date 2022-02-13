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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetDeviceObjectPropertyReference struct {
	OpeningTag         *BACnetOpeningTag
	ObjectIdentifier   *BACnetContextTagObjectIdentifier
	PropertyIdentifier *BACnetContextTagPropertyIdentifier
	ArrayIndex         *BACnetContextTagUnsignedInteger
	DeviceIdentifier   *BACnetContextTagObjectIdentifier
	ClosingTag         *BACnetClosingTag
}

// The corresponding interface
type IBACnetDeviceObjectPropertyReference interface {
	// GetOpeningTag returns OpeningTag
	GetOpeningTag() *BACnetOpeningTag
	// GetObjectIdentifier returns ObjectIdentifier
	GetObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier
	GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier
	// GetArrayIndex returns ArrayIndex
	GetArrayIndex() *BACnetContextTagUnsignedInteger
	// GetDeviceIdentifier returns DeviceIdentifier
	GetDeviceIdentifier() *BACnetContextTagObjectIdentifier
	// GetClosingTag returns ClosingTag
	GetClosingTag() *BACnetClosingTag
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetDeviceObjectPropertyReference) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetDeviceObjectPropertyReference) GetObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetDeviceObjectPropertyReference) GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier {
	return m.PropertyIdentifier
}

func (m *BACnetDeviceObjectPropertyReference) GetArrayIndex() *BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *BACnetDeviceObjectPropertyReference) GetDeviceIdentifier() *BACnetContextTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *BACnetDeviceObjectPropertyReference) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewBACnetDeviceObjectPropertyReference(openingTag *BACnetOpeningTag, objectIdentifier *BACnetContextTagObjectIdentifier, propertyIdentifier *BACnetContextTagPropertyIdentifier, arrayIndex *BACnetContextTagUnsignedInteger, deviceIdentifier *BACnetContextTagObjectIdentifier, closingTag *BACnetClosingTag) *BACnetDeviceObjectPropertyReference {
	return &BACnetDeviceObjectPropertyReference{OpeningTag: openingTag, ObjectIdentifier: objectIdentifier, PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, DeviceIdentifier: deviceIdentifier, ClosingTag: closingTag}
}

func CastBACnetDeviceObjectPropertyReference(structType interface{}) *BACnetDeviceObjectPropertyReference {
	castFunc := func(typ interface{}) *BACnetDeviceObjectPropertyReference {
		if casted, ok := typ.(BACnetDeviceObjectPropertyReference); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetDeviceObjectPropertyReference); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetDeviceObjectPropertyReference) GetTypeName() string {
	return "BACnetDeviceObjectPropertyReference"
}

func (m *BACnetDeviceObjectPropertyReference) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetDeviceObjectPropertyReference) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.LengthInBits()

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.LengthInBits()

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.LengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += (*m.ArrayIndex).LengthInBits()
	}

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += (*m.DeviceIdentifier).LengthInBits()
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.LengthInBits()

	return lengthInBits
}

func (m *BACnetDeviceObjectPropertyReference) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetDeviceObjectPropertyReferenceParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetDeviceObjectPropertyReference, error) {
	if pullErr := readBuffer.PullContext("BACnetDeviceObjectPropertyReference"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType_OPENING_TAG)
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetContextTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType_BACNET_PROPERTY_IDENTIFIER)
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := CastBACnetContextTagPropertyIdentifier(_propertyIdentifier)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field")
		default:
			arrayIndex = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (deviceIdentifier) (Can be skipped, if a given expression evaluates to false)
	var deviceIdentifier *BACnetContextTagObjectIdentifier = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceIdentifier' field")
		default:
			deviceIdentifier = CastBACnetContextTagObjectIdentifier(_val)
			if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType_CLOSING_TAG)
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetDeviceObjectPropertyReference"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetDeviceObjectPropertyReference(openingTag, objectIdentifier, propertyIdentifier, arrayIndex, deviceIdentifier, closingTag), nil
}

func (m *BACnetDeviceObjectPropertyReference) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetDeviceObjectPropertyReference"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return pushErr
	}
	_objectIdentifierErr := m.ObjectIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return popErr
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return pushErr
	}
	_propertyIdentifierErr := m.PropertyIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return popErr
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	if m.ArrayIndex != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return pushErr
		}
		arrayIndex = m.ArrayIndex
		_arrayIndexErr := arrayIndex.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return popErr
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Optional Field (deviceIdentifier) (Can be skipped, if the value is null)
	var deviceIdentifier *BACnetContextTagObjectIdentifier = nil
	if m.DeviceIdentifier != nil {
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return pushErr
		}
		deviceIdentifier = m.DeviceIdentifier
		_deviceIdentifierErr := deviceIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return popErr
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDeviceObjectPropertyReference"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetDeviceObjectPropertyReference) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
