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

// BACnetConstructedDataInstallationID is the data-structure of this message
type BACnetConstructedDataInstallationID struct {
	*BACnetConstructedData
	InstallationId *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataInstallationID is the corresponding interface of BACnetConstructedDataInstallationID
type IBACnetConstructedDataInstallationID interface {
	IBACnetConstructedData
	// GetInstallationId returns InstallationId (property field)
	GetInstallationId() *BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataInstallationID) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataInstallationID) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INSTALLATION_ID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataInstallationID) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataInstallationID) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataInstallationID) GetInstallationId() *BACnetApplicationTagUnsignedInteger {
	return m.InstallationId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataInstallationID) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetInstallationId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInstallationID factory function for BACnetConstructedDataInstallationID
func NewBACnetConstructedDataInstallationID(installationId *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataInstallationID {
	_result := &BACnetConstructedDataInstallationID{
		InstallationId:        installationId,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataInstallationID(structType interface{}) *BACnetConstructedDataInstallationID {
	if casted, ok := structType.(BACnetConstructedDataInstallationID); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInstallationID); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataInstallationID(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataInstallationID(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataInstallationID) GetTypeName() string {
	return "BACnetConstructedDataInstallationID"
}

func (m *BACnetConstructedDataInstallationID) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataInstallationID) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (installationId)
	lengthInBits += m.InstallationId.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataInstallationID) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInstallationIDParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataInstallationID, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInstallationID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInstallationID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (installationId)
	if pullErr := readBuffer.PullContext("installationId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for installationId")
	}
	_installationId, _installationIdErr := BACnetApplicationTagParse(readBuffer)
	if _installationIdErr != nil {
		return nil, errors.Wrap(_installationIdErr, "Error parsing 'installationId' field")
	}
	installationId := CastBACnetApplicationTagUnsignedInteger(_installationId)
	if closeErr := readBuffer.CloseContext("installationId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for installationId")
	}

	// Virtual field
	_actualValue := installationId
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInstallationID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInstallationID")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataInstallationID{
		InstallationId:        CastBACnetApplicationTagUnsignedInteger(installationId),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataInstallationID) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInstallationID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInstallationID")
		}

		// Simple Field (installationId)
		if pushErr := writeBuffer.PushContext("installationId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for installationId")
		}
		_installationIdErr := writeBuffer.WriteSerializable(m.InstallationId)
		if popErr := writeBuffer.PopContext("installationId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for installationId")
		}
		if _installationIdErr != nil {
			return errors.Wrap(_installationIdErr, "Error serializing 'installationId' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInstallationID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInstallationID")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataInstallationID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
