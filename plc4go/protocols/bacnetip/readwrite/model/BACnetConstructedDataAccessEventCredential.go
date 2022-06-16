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

// BACnetConstructedDataAccessEventCredential is the corresponding interface of BACnetConstructedDataAccessEventCredential
type BACnetConstructedDataAccessEventCredential interface {
	BACnetConstructedData
	// GetAccessEventCredential returns AccessEventCredential (property field)
	GetAccessEventCredential() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConstructedDataAccessEventCredential is the data-structure of this message
type _BACnetConstructedDataAccessEventCredential struct {
	*_BACnetConstructedData
	AccessEventCredential BACnetDeviceObjectReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessEventCredential) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessEventCredential) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT_CREDENTIAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessEventCredential) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessEventCredential) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventCredential) GetAccessEventCredential() BACnetDeviceObjectReference {
	return m.AccessEventCredential
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventCredential) GetActualValue() BACnetDeviceObjectReference {
	return CastBACnetDeviceObjectReference(m.GetAccessEventCredential())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessEventCredential factory function for _BACnetConstructedDataAccessEventCredential
func NewBACnetConstructedDataAccessEventCredential(accessEventCredential BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessEventCredential {
	_result := &_BACnetConstructedDataAccessEventCredential{
		AccessEventCredential:  accessEventCredential,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessEventCredential(structType interface{}) BACnetConstructedDataAccessEventCredential {
	if casted, ok := structType.(BACnetConstructedDataAccessEventCredential); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEventCredential); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessEventCredential) GetTypeName() string {
	return "BACnetConstructedDataAccessEventCredential"
}

func (m *_BACnetConstructedDataAccessEventCredential) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAccessEventCredential) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accessEventCredential)
	lengthInBits += m.AccessEventCredential.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessEventCredential) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccessEventCredentialParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessEventCredential, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEventCredential"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessEventCredential")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessEventCredential)
	if pullErr := readBuffer.PullContext("accessEventCredential"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEventCredential")
	}
	_accessEventCredential, _accessEventCredentialErr := BACnetDeviceObjectReferenceParse(readBuffer)
	if _accessEventCredentialErr != nil {
		return nil, errors.Wrap(_accessEventCredentialErr, "Error parsing 'accessEventCredential' field")
	}
	accessEventCredential := _accessEventCredential.(BACnetDeviceObjectReference)
	if closeErr := readBuffer.CloseContext("accessEventCredential"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEventCredential")
	}

	// Virtual field
	_actualValue := accessEventCredential
	actualValue := _actualValue.(BACnetDeviceObjectReference)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEventCredential"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessEventCredential")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessEventCredential{
		AccessEventCredential:  accessEventCredential,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessEventCredential) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEventCredential"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessEventCredential")
		}

		// Simple Field (accessEventCredential)
		if pushErr := writeBuffer.PushContext("accessEventCredential"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEventCredential")
		}
		_accessEventCredentialErr := writeBuffer.WriteSerializable(m.GetAccessEventCredential())
		if popErr := writeBuffer.PopContext("accessEventCredential"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEventCredential")
		}
		if _accessEventCredentialErr != nil {
			return errors.Wrap(_accessEventCredentialErr, "Error serializing 'accessEventCredential' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEventCredential"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessEventCredential")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessEventCredential) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
