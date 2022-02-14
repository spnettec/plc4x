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
type BACnetConstructedData struct {
	OpeningTag *BACnetOpeningTag
	ClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber                  uint8
	PropertyIdentifierArgument BACnetContextTagPropertyIdentifier
	Child                      IBACnetConstructedDataChild
}

// The corresponding interface
type IBACnetConstructedData interface {
	// ObjectType returns ObjectType
	ObjectType() BACnetObjectType
	// PropertyIdentifierEnum returns PropertyIdentifierEnum
	PropertyIdentifierEnum() BACnetPropertyIdentifier
	// GetOpeningTag returns OpeningTag
	GetOpeningTag() *BACnetOpeningTag
	// GetClosingTag returns ClosingTag
	GetClosingTag() *BACnetClosingTag
	// GetPropertyIdentifierEnum returns PropertyIdentifierEnum
	GetPropertyIdentifierEnum() BACnetPropertyIdentifier
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetConstructedDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetConstructedData, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetConstructedDataChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag)
	GetTypeName() string
	IBACnetConstructedData
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetConstructedData) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetConstructedData) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetConstructedData) GetPropertyIdentifierEnum() BACnetPropertyIdentifier {
	return m.PropertyIdentifierArgument.GetPropertyIdentifier()
}

// NewBACnetConstructedData factory function for BACnetConstructedData
func NewBACnetConstructedData(openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8, propertyIdentifierArgument BACnetContextTagPropertyIdentifier) *BACnetConstructedData {
	return &BACnetConstructedData{OpeningTag: openingTag, ClosingTag: closingTag, TagNumber: tagNumber, PropertyIdentifierArgument: propertyIdentifierArgument}
}

func CastBACnetConstructedData(structType interface{}) *BACnetConstructedData {
	castFunc := func(typ interface{}) *BACnetConstructedData {
		if casted, ok := typ.(BACnetConstructedData); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConstructedData); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConstructedData) GetTypeName() string {
	return "BACnetConstructedData"
}

func (m *BACnetConstructedData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedData) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetConstructedData) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, propertyIdentifierArgument *BACnetContextTagPropertyIdentifier) (*BACnetConstructedData, error) {
	if pullErr := readBuffer.PullContext("BACnetConstructedData"); pullErr != nil {
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

	// Virtual field
	_propertyIdentifierEnum := propertyIdentifierArgument.GetPropertyIdentifier()
	propertyIdentifierEnum := BACnetPropertyIdentifier(_propertyIdentifierEnum)
	_ = propertyIdentifierEnum

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetConstructedData
	var typeSwitchError error
	switch {
	case objectType == BACnetObjectType_COMMAND: // BACnetConstructedDataCommand
		_parent, typeSwitchError = BACnetConstructedDataCommandParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LIFE_SAFETY_ZONE: // BACnetConstructedDataLifeSafetyZone
		_parent, typeSwitchError = BACnetConstructedDataLifeSafetyZoneParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true && propertyIdentifierEnum == BACnetPropertyIdentifier_EVENT_TIME_STAMPS: // BACnetConstructedDataEventTimestamps
		_parent, typeSwitchError = BACnetConstructedDataEventTimestampsParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true: // BACnetConstructedDataUnspecified
		_parent, typeSwitchError = BACnetConstructedDataUnspecifiedParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
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

	if closeErr := readBuffer.CloseContext("BACnetConstructedData"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, openingTag, closingTag)
	return _parent, nil
}

func (m *BACnetConstructedData) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetConstructedData) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetConstructedData, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetConstructedData"); pushErr != nil {
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
	// Virtual field
	if _propertyIdentifierEnumErr := writeBuffer.WriteVirtual("propertyIdentifierEnum", m.GetPropertyIdentifierEnum()); _propertyIdentifierEnumErr != nil {
		return errors.Wrap(_propertyIdentifierEnumErr, "Error serializing 'propertyIdentifierEnum' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
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

	if popErr := writeBuffer.PopContext("BACnetConstructedData"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConstructedData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
