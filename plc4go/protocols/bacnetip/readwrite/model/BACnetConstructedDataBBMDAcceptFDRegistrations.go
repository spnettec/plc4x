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

// BACnetConstructedDataBBMDAcceptFDRegistrations is the data-structure of this message
type BACnetConstructedDataBBMDAcceptFDRegistrations struct {
	*BACnetConstructedData
	BbmdAcceptFDRegistrations *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBBMDAcceptFDRegistrations is the corresponding interface of BACnetConstructedDataBBMDAcceptFDRegistrations
type IBACnetConstructedDataBBMDAcceptFDRegistrations interface {
	IBACnetConstructedData
	// GetBbmdAcceptFDRegistrations returns BbmdAcceptFDRegistrations (property field)
	GetBbmdAcceptFDRegistrations() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BBMD_ACCEPT_FD_REGISTRATIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) GetBbmdAcceptFDRegistrations() *BACnetApplicationTagBoolean {
	return m.BbmdAcceptFDRegistrations
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBBMDAcceptFDRegistrations factory function for BACnetConstructedDataBBMDAcceptFDRegistrations
func NewBACnetConstructedDataBBMDAcceptFDRegistrations(bbmdAcceptFDRegistrations *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBBMDAcceptFDRegistrations {
	_result := &BACnetConstructedDataBBMDAcceptFDRegistrations{
		BbmdAcceptFDRegistrations: bbmdAcceptFDRegistrations,
		BACnetConstructedData:     NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBBMDAcceptFDRegistrations(structType interface{}) *BACnetConstructedDataBBMDAcceptFDRegistrations {
	if casted, ok := structType.(BACnetConstructedDataBBMDAcceptFDRegistrations); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBBMDAcceptFDRegistrations); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBBMDAcceptFDRegistrations(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBBMDAcceptFDRegistrations(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) GetTypeName() string {
	return "BACnetConstructedDataBBMDAcceptFDRegistrations"
}

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bbmdAcceptFDRegistrations)
	lengthInBits += m.BbmdAcceptFDRegistrations.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBBMDAcceptFDRegistrationsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBBMDAcceptFDRegistrations, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBBMDAcceptFDRegistrations")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bbmdAcceptFDRegistrations)
	if pullErr := readBuffer.PullContext("bbmdAcceptFDRegistrations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bbmdAcceptFDRegistrations")
	}
	_bbmdAcceptFDRegistrations, _bbmdAcceptFDRegistrationsErr := BACnetApplicationTagParse(readBuffer)
	if _bbmdAcceptFDRegistrationsErr != nil {
		return nil, errors.Wrap(_bbmdAcceptFDRegistrationsErr, "Error parsing 'bbmdAcceptFDRegistrations' field")
	}
	bbmdAcceptFDRegistrations := CastBACnetApplicationTagBoolean(_bbmdAcceptFDRegistrations)
	if closeErr := readBuffer.CloseContext("bbmdAcceptFDRegistrations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bbmdAcceptFDRegistrations")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBBMDAcceptFDRegistrations")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBBMDAcceptFDRegistrations{
		BbmdAcceptFDRegistrations: CastBACnetApplicationTagBoolean(bbmdAcceptFDRegistrations),
		BACnetConstructedData:     &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBBMDAcceptFDRegistrations")
		}

		// Simple Field (bbmdAcceptFDRegistrations)
		if pushErr := writeBuffer.PushContext("bbmdAcceptFDRegistrations"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bbmdAcceptFDRegistrations")
		}
		_bbmdAcceptFDRegistrationsErr := m.BbmdAcceptFDRegistrations.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("bbmdAcceptFDRegistrations"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bbmdAcceptFDRegistrations")
		}
		if _bbmdAcceptFDRegistrationsErr != nil {
			return errors.Wrap(_bbmdAcceptFDRegistrationsErr, "Error serializing 'bbmdAcceptFDRegistrations' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBBMDAcceptFDRegistrations")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBBMDAcceptFDRegistrations) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
