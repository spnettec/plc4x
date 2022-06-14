/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataUserExternalIdentifier is the data-structure of this message
type BACnetConstructedDataUserExternalIdentifier struct {
	*BACnetConstructedData
	UserExternalIdentifier *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataUserExternalIdentifier is the corresponding interface of BACnetConstructedDataUserExternalIdentifier
type IBACnetConstructedDataUserExternalIdentifier interface {
	IBACnetConstructedData
	// GetUserExternalIdentifier returns UserExternalIdentifier (property field)
	GetUserExternalIdentifier() *BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataUserExternalIdentifier) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataUserExternalIdentifier) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_USER_EXTERNAL_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataUserExternalIdentifier) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataUserExternalIdentifier) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataUserExternalIdentifier) GetUserExternalIdentifier() *BACnetApplicationTagCharacterString {
	return m.UserExternalIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataUserExternalIdentifier) GetActualValue() *BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetUserExternalIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUserExternalIdentifier factory function for BACnetConstructedDataUserExternalIdentifier
func NewBACnetConstructedDataUserExternalIdentifier(userExternalIdentifier *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataUserExternalIdentifier {
	_result := &BACnetConstructedDataUserExternalIdentifier{
		UserExternalIdentifier: userExternalIdentifier,
		BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataUserExternalIdentifier(structType interface{}) *BACnetConstructedDataUserExternalIdentifier {
	if casted, ok := structType.(BACnetConstructedDataUserExternalIdentifier); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUserExternalIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataUserExternalIdentifier(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataUserExternalIdentifier(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataUserExternalIdentifier) GetTypeName() string {
	return "BACnetConstructedDataUserExternalIdentifier"
}

func (m *BACnetConstructedDataUserExternalIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataUserExternalIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (userExternalIdentifier)
	lengthInBits += m.UserExternalIdentifier.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataUserExternalIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUserExternalIdentifierParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataUserExternalIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUserExternalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUserExternalIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (userExternalIdentifier)
	if pullErr := readBuffer.PullContext("userExternalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userExternalIdentifier")
	}
	_userExternalIdentifier, _userExternalIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _userExternalIdentifierErr != nil {
		return nil, errors.Wrap(_userExternalIdentifierErr, "Error parsing 'userExternalIdentifier' field")
	}
	userExternalIdentifier := CastBACnetApplicationTagCharacterString(_userExternalIdentifier)
	if closeErr := readBuffer.CloseContext("userExternalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userExternalIdentifier")
	}

	// Virtual field
	_actualValue := userExternalIdentifier
	actualValue := CastBACnetApplicationTagCharacterString(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUserExternalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUserExternalIdentifier")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataUserExternalIdentifier{
		UserExternalIdentifier: CastBACnetApplicationTagCharacterString(userExternalIdentifier),
		BACnetConstructedData:  &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataUserExternalIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUserExternalIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUserExternalIdentifier")
		}

		// Simple Field (userExternalIdentifier)
		if pushErr := writeBuffer.PushContext("userExternalIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userExternalIdentifier")
		}
		_userExternalIdentifierErr := writeBuffer.WriteSerializable(m.UserExternalIdentifier)
		if popErr := writeBuffer.PopContext("userExternalIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userExternalIdentifier")
		}
		if _userExternalIdentifierErr != nil {
			return errors.Wrap(_userExternalIdentifierErr, "Error serializing 'userExternalIdentifier' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUserExternalIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUserExternalIdentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataUserExternalIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
