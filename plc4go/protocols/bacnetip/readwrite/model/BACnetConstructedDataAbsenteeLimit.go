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

// BACnetConstructedDataAbsenteeLimit is the data-structure of this message
type BACnetConstructedDataAbsenteeLimit struct {
	*BACnetConstructedData
	AbsenteeLimit *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAbsenteeLimit is the corresponding interface of BACnetConstructedDataAbsenteeLimit
type IBACnetConstructedDataAbsenteeLimit interface {
	IBACnetConstructedData
	// GetAbsenteeLimit returns AbsenteeLimit (property field)
	GetAbsenteeLimit() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataAbsenteeLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAbsenteeLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ABSENTEE_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAbsenteeLimit) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAbsenteeLimit) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAbsenteeLimit) GetAbsenteeLimit() *BACnetApplicationTagUnsignedInteger {
	return m.AbsenteeLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataAbsenteeLimit) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetAbsenteeLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAbsenteeLimit factory function for BACnetConstructedDataAbsenteeLimit
func NewBACnetConstructedDataAbsenteeLimit(absenteeLimit *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAbsenteeLimit {
	_result := &BACnetConstructedDataAbsenteeLimit{
		AbsenteeLimit:         absenteeLimit,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAbsenteeLimit(structType interface{}) *BACnetConstructedDataAbsenteeLimit {
	if casted, ok := structType.(BACnetConstructedDataAbsenteeLimit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAbsenteeLimit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAbsenteeLimit(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAbsenteeLimit(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAbsenteeLimit) GetTypeName() string {
	return "BACnetConstructedDataAbsenteeLimit"
}

func (m *BACnetConstructedDataAbsenteeLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAbsenteeLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (absenteeLimit)
	lengthInBits += m.AbsenteeLimit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataAbsenteeLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAbsenteeLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAbsenteeLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAbsenteeLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAbsenteeLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (absenteeLimit)
	if pullErr := readBuffer.PullContext("absenteeLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for absenteeLimit")
	}
	_absenteeLimit, _absenteeLimitErr := BACnetApplicationTagParse(readBuffer)
	if _absenteeLimitErr != nil {
		return nil, errors.Wrap(_absenteeLimitErr, "Error parsing 'absenteeLimit' field")
	}
	absenteeLimit := CastBACnetApplicationTagUnsignedInteger(_absenteeLimit)
	if closeErr := readBuffer.CloseContext("absenteeLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for absenteeLimit")
	}

	// Virtual field
	_actualValue := absenteeLimit
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAbsenteeLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAbsenteeLimit")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAbsenteeLimit{
		AbsenteeLimit:         CastBACnetApplicationTagUnsignedInteger(absenteeLimit),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAbsenteeLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAbsenteeLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAbsenteeLimit")
		}

		// Simple Field (absenteeLimit)
		if pushErr := writeBuffer.PushContext("absenteeLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for absenteeLimit")
		}
		_absenteeLimitErr := writeBuffer.WriteSerializable(m.AbsenteeLimit)
		if popErr := writeBuffer.PopContext("absenteeLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for absenteeLimit")
		}
		if _absenteeLimitErr != nil {
			return errors.Wrap(_absenteeLimitErr, "Error serializing 'absenteeLimit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAbsenteeLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAbsenteeLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAbsenteeLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
