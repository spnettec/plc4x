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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataDescription is the data-structure of this message
type BACnetConstructedDataDescription struct {
	*BACnetConstructedData
	Description *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDescription is the corresponding interface of BACnetConstructedDataDescription
type IBACnetConstructedDataDescription interface {
	IBACnetConstructedData
	// GetDescription returns Description (property field)
	GetDescription() *BACnetApplicationTagCharacterString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataDescription) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDescription) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DESCRIPTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDescription) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDescription) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDescription) GetDescription() *BACnetApplicationTagCharacterString {
	return m.Description
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDescription factory function for BACnetConstructedDataDescription
func NewBACnetConstructedDataDescription(description *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDescription {
	_result := &BACnetConstructedDataDescription{
		Description:           description,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDescription(structType interface{}) *BACnetConstructedDataDescription {
	if casted, ok := structType.(BACnetConstructedDataDescription); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDescription); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDescription(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDescription(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDescription) GetTypeName() string {
	return "BACnetConstructedDataDescription"
}

func (m *BACnetConstructedDataDescription) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDescription) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDescription) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDescriptionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDescription, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (description)
	if pullErr := readBuffer.PullContext("description"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for description")
	}
	_description, _descriptionErr := BACnetApplicationTagParse(readBuffer)
	if _descriptionErr != nil {
		return nil, errors.Wrap(_descriptionErr, "Error parsing 'description' field")
	}
	description := CastBACnetApplicationTagCharacterString(_description)
	if closeErr := readBuffer.CloseContext("description"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for description")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDescription")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDescription{
		Description:           CastBACnetApplicationTagCharacterString(description),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDescription) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDescription")
		}

		// Simple Field (description)
		if pushErr := writeBuffer.PushContext("description"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for description")
		}
		_descriptionErr := m.Description.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("description"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for description")
		}
		if _descriptionErr != nil {
			return errors.Wrap(_descriptionErr, "Error serializing 'description' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDescription")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
