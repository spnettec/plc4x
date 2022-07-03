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

// BACnetConstructedDataUserInformationReference is the corresponding interface of BACnetConstructedDataUserInformationReference
type BACnetConstructedDataUserInformationReference interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUserInformationReference returns UserInformationReference (property field)
	GetUserInformationReference() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataUserInformationReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUserInformationReference.
// This is useful for switch cases.
type BACnetConstructedDataUserInformationReferenceExactly interface {
	BACnetConstructedDataUserInformationReference
	isBACnetConstructedDataUserInformationReference() bool
}

// _BACnetConstructedDataUserInformationReference is the data-structure of this message
type _BACnetConstructedDataUserInformationReference struct {
	*_BACnetConstructedData
	UserInformationReference BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUserInformationReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUserInformationReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_USER_INFORMATION_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUserInformationReference) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUserInformationReference) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUserInformationReference) GetUserInformationReference() BACnetApplicationTagCharacterString {
	return m.UserInformationReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUserInformationReference) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetUserInformationReference())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUserInformationReference factory function for _BACnetConstructedDataUserInformationReference
func NewBACnetConstructedDataUserInformationReference(userInformationReference BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUserInformationReference {
	_result := &_BACnetConstructedDataUserInformationReference{
		UserInformationReference: userInformationReference,
		_BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUserInformationReference(structType interface{}) BACnetConstructedDataUserInformationReference {
	if casted, ok := structType.(BACnetConstructedDataUserInformationReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUserInformationReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUserInformationReference) GetTypeName() string {
	return "BACnetConstructedDataUserInformationReference"
}

func (m *_BACnetConstructedDataUserInformationReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataUserInformationReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (userInformationReference)
	lengthInBits += m.UserInformationReference.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUserInformationReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUserInformationReferenceParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUserInformationReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUserInformationReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUserInformationReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (userInformationReference)
	if pullErr := readBuffer.PullContext("userInformationReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userInformationReference")
	}
	_userInformationReference, _userInformationReferenceErr := BACnetApplicationTagParse(readBuffer)
	if _userInformationReferenceErr != nil {
		return nil, errors.Wrap(_userInformationReferenceErr, "Error parsing 'userInformationReference' field of BACnetConstructedDataUserInformationReference")
	}
	userInformationReference := _userInformationReference.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("userInformationReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userInformationReference")
	}

	// Virtual field
	_actualValue := userInformationReference
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUserInformationReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUserInformationReference")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUserInformationReference{
		UserInformationReference: userInformationReference,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUserInformationReference) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUserInformationReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUserInformationReference")
		}

		// Simple Field (userInformationReference)
		if pushErr := writeBuffer.PushContext("userInformationReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userInformationReference")
		}
		_userInformationReferenceErr := writeBuffer.WriteSerializable(m.GetUserInformationReference())
		if popErr := writeBuffer.PopContext("userInformationReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userInformationReference")
		}
		if _userInformationReferenceErr != nil {
			return errors.Wrap(_userInformationReferenceErr, "Error serializing 'userInformationReference' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUserInformationReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUserInformationReference")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUserInformationReference) isBACnetConstructedDataUserInformationReference() bool {
	return true
}

func (m *_BACnetConstructedDataUserInformationReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
