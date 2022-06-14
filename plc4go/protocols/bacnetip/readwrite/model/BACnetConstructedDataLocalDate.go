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

// BACnetConstructedDataLocalDate is the data-structure of this message
type BACnetConstructedDataLocalDate struct {
	*BACnetConstructedData
	LocalDate *BACnetApplicationTagDate

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLocalDate is the corresponding interface of BACnetConstructedDataLocalDate
type IBACnetConstructedDataLocalDate interface {
	IBACnetConstructedData
	// GetLocalDate returns LocalDate (property field)
	GetLocalDate() *BACnetApplicationTagDate
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagDate
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

func (m *BACnetConstructedDataLocalDate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLocalDate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCAL_DATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLocalDate) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLocalDate) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLocalDate) GetLocalDate() *BACnetApplicationTagDate {
	return m.LocalDate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataLocalDate) GetActualValue() *BACnetApplicationTagDate {
	return CastBACnetApplicationTagDate(m.GetLocalDate())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLocalDate factory function for BACnetConstructedDataLocalDate
func NewBACnetConstructedDataLocalDate(localDate *BACnetApplicationTagDate, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLocalDate {
	_result := &BACnetConstructedDataLocalDate{
		LocalDate:             localDate,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLocalDate(structType interface{}) *BACnetConstructedDataLocalDate {
	if casted, ok := structType.(BACnetConstructedDataLocalDate); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLocalDate); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLocalDate(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLocalDate(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLocalDate) GetTypeName() string {
	return "BACnetConstructedDataLocalDate"
}

func (m *BACnetConstructedDataLocalDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLocalDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (localDate)
	lengthInBits += m.LocalDate.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataLocalDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLocalDateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLocalDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLocalDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLocalDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (localDate)
	if pullErr := readBuffer.PullContext("localDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for localDate")
	}
	_localDate, _localDateErr := BACnetApplicationTagParse(readBuffer)
	if _localDateErr != nil {
		return nil, errors.Wrap(_localDateErr, "Error parsing 'localDate' field")
	}
	localDate := CastBACnetApplicationTagDate(_localDate)
	if closeErr := readBuffer.CloseContext("localDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for localDate")
	}

	// Virtual field
	_actualValue := localDate
	actualValue := CastBACnetApplicationTagDate(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLocalDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLocalDate")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLocalDate{
		LocalDate:             CastBACnetApplicationTagDate(localDate),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLocalDate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLocalDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLocalDate")
		}

		// Simple Field (localDate)
		if pushErr := writeBuffer.PushContext("localDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for localDate")
		}
		_localDateErr := writeBuffer.WriteSerializable(m.LocalDate)
		if popErr := writeBuffer.PopContext("localDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for localDate")
		}
		if _localDateErr != nil {
			return errors.Wrap(_localDateErr, "Error serializing 'localDate' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLocalDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLocalDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLocalDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
