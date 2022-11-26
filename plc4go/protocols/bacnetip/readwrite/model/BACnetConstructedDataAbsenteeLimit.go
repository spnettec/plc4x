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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataAbsenteeLimit is the corresponding interface of BACnetConstructedDataAbsenteeLimit
type BACnetConstructedDataAbsenteeLimit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAbsenteeLimit returns AbsenteeLimit (property field)
	GetAbsenteeLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAbsenteeLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAbsenteeLimit.
// This is useful for switch cases.
type BACnetConstructedDataAbsenteeLimitExactly interface {
	BACnetConstructedDataAbsenteeLimit
	isBACnetConstructedDataAbsenteeLimit() bool
}

// _BACnetConstructedDataAbsenteeLimit is the data-structure of this message
type _BACnetConstructedDataAbsenteeLimit struct {
	*_BACnetConstructedData
	AbsenteeLimit BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAbsenteeLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAbsenteeLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ABSENTEE_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAbsenteeLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAbsenteeLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAbsenteeLimit) GetAbsenteeLimit() BACnetApplicationTagUnsignedInteger {
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

func (m *_BACnetConstructedDataAbsenteeLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetAbsenteeLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAbsenteeLimit factory function for _BACnetConstructedDataAbsenteeLimit
func NewBACnetConstructedDataAbsenteeLimit(absenteeLimit BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAbsenteeLimit {
	_result := &_BACnetConstructedDataAbsenteeLimit{
		AbsenteeLimit:          absenteeLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAbsenteeLimit(structType interface{}) BACnetConstructedDataAbsenteeLimit {
	if casted, ok := structType.(BACnetConstructedDataAbsenteeLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAbsenteeLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAbsenteeLimit) GetTypeName() string {
	return "BACnetConstructedDataAbsenteeLimit"
}

func (m *_BACnetConstructedDataAbsenteeLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAbsenteeLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (absenteeLimit)
	lengthInBits += m.AbsenteeLimit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAbsenteeLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAbsenteeLimitParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAbsenteeLimit, error) {
	return BACnetConstructedDataAbsenteeLimitParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAbsenteeLimitParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAbsenteeLimit, error) {
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
	_absenteeLimit, _absenteeLimitErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _absenteeLimitErr != nil {
		return nil, errors.Wrap(_absenteeLimitErr, "Error parsing 'absenteeLimit' field of BACnetConstructedDataAbsenteeLimit")
	}
	absenteeLimit := _absenteeLimit.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("absenteeLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for absenteeLimit")
	}

	// Virtual field
	_actualValue := absenteeLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAbsenteeLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAbsenteeLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAbsenteeLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AbsenteeLimit: absenteeLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAbsenteeLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAbsenteeLimit) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
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
		_absenteeLimitErr := writeBuffer.WriteSerializable(m.GetAbsenteeLimit())
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

func (m *_BACnetConstructedDataAbsenteeLimit) isBACnetConstructedDataAbsenteeLimit() bool {
	return true
}

func (m *_BACnetConstructedDataAbsenteeLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
