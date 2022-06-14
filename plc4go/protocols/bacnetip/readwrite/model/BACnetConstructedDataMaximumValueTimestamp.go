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

// BACnetConstructedDataMaximumValueTimestamp is the data-structure of this message
type BACnetConstructedDataMaximumValueTimestamp struct {
	*BACnetConstructedData
	MaximumValueTimestamp *BACnetDateTime

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMaximumValueTimestamp is the corresponding interface of BACnetConstructedDataMaximumValueTimestamp
type IBACnetConstructedDataMaximumValueTimestamp interface {
	IBACnetConstructedData
	// GetMaximumValueTimestamp returns MaximumValueTimestamp (property field)
	GetMaximumValueTimestamp() *BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetDateTime
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

func (m *BACnetConstructedDataMaximumValueTimestamp) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMaximumValueTimestamp) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAXIMUM_VALUE_TIMESTAMP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMaximumValueTimestamp) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMaximumValueTimestamp) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMaximumValueTimestamp) GetMaximumValueTimestamp() *BACnetDateTime {
	return m.MaximumValueTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataMaximumValueTimestamp) GetActualValue() *BACnetDateTime {
	return CastBACnetDateTime(m.GetMaximumValueTimestamp())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaximumValueTimestamp factory function for BACnetConstructedDataMaximumValueTimestamp
func NewBACnetConstructedDataMaximumValueTimestamp(maximumValueTimestamp *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMaximumValueTimestamp {
	_result := &BACnetConstructedDataMaximumValueTimestamp{
		MaximumValueTimestamp: maximumValueTimestamp,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMaximumValueTimestamp(structType interface{}) *BACnetConstructedDataMaximumValueTimestamp {
	if casted, ok := structType.(BACnetConstructedDataMaximumValueTimestamp); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaximumValueTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaximumValueTimestamp(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaximumValueTimestamp(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMaximumValueTimestamp) GetTypeName() string {
	return "BACnetConstructedDataMaximumValueTimestamp"
}

func (m *BACnetConstructedDataMaximumValueTimestamp) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMaximumValueTimestamp) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maximumValueTimestamp)
	lengthInBits += m.MaximumValueTimestamp.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataMaximumValueTimestamp) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaximumValueTimestampParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMaximumValueTimestamp, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaximumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaximumValueTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maximumValueTimestamp)
	if pullErr := readBuffer.PullContext("maximumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maximumValueTimestamp")
	}
	_maximumValueTimestamp, _maximumValueTimestampErr := BACnetDateTimeParse(readBuffer)
	if _maximumValueTimestampErr != nil {
		return nil, errors.Wrap(_maximumValueTimestampErr, "Error parsing 'maximumValueTimestamp' field")
	}
	maximumValueTimestamp := CastBACnetDateTime(_maximumValueTimestamp)
	if closeErr := readBuffer.CloseContext("maximumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maximumValueTimestamp")
	}

	// Virtual field
	_actualValue := maximumValueTimestamp
	actualValue := CastBACnetDateTime(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaximumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaximumValueTimestamp")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMaximumValueTimestamp{
		MaximumValueTimestamp: CastBACnetDateTime(maximumValueTimestamp),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMaximumValueTimestamp) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaximumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaximumValueTimestamp")
		}

		// Simple Field (maximumValueTimestamp)
		if pushErr := writeBuffer.PushContext("maximumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maximumValueTimestamp")
		}
		_maximumValueTimestampErr := writeBuffer.WriteSerializable(m.MaximumValueTimestamp)
		if popErr := writeBuffer.PopContext("maximumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maximumValueTimestamp")
		}
		if _maximumValueTimestampErr != nil {
			return errors.Wrap(_maximumValueTimestampErr, "Error serializing 'maximumValueTimestamp' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaximumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaximumValueTimestamp")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMaximumValueTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
