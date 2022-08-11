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

// BACnetConstructedDataAccumulatorHighLimit is the corresponding interface of BACnetConstructedDataAccumulatorHighLimit
type BACnetConstructedDataAccumulatorHighLimit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAccumulatorHighLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccumulatorHighLimit.
// This is useful for switch cases.
type BACnetConstructedDataAccumulatorHighLimitExactly interface {
	BACnetConstructedDataAccumulatorHighLimit
	isBACnetConstructedDataAccumulatorHighLimit() bool
}

// _BACnetConstructedDataAccumulatorHighLimit is the data-structure of this message
type _BACnetConstructedDataAccumulatorHighLimit struct {
	*_BACnetConstructedData
	HighLimit BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccumulatorHighLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetHighLimit() BACnetApplicationTagUnsignedInteger {
	return m.HighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccumulatorHighLimit factory function for _BACnetConstructedDataAccumulatorHighLimit
func NewBACnetConstructedDataAccumulatorHighLimit(highLimit BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccumulatorHighLimit {
	_result := &_BACnetConstructedDataAccumulatorHighLimit{
		HighLimit:              highLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccumulatorHighLimit(structType interface{}) BACnetConstructedDataAccumulatorHighLimit {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorHighLimit"
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccumulatorHighLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccumulatorHighLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccumulatorHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (highLimit)
	if pullErr := readBuffer.PullContext("highLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for highLimit")
	}
	_highLimit, _highLimitErr := BACnetApplicationTagParse(readBuffer)
	if _highLimitErr != nil {
		return nil, errors.Wrap(_highLimitErr, "Error parsing 'highLimit' field of BACnetConstructedDataAccumulatorHighLimit")
	}
	highLimit := _highLimit.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("highLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for highLimit")
	}

	// Virtual field
	_actualValue := highLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccumulatorHighLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccumulatorHighLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		HighLimit: highLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccumulatorHighLimit")
		}

		// Simple Field (highLimit)
		if pushErr := writeBuffer.PushContext("highLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for highLimit")
		}
		_highLimitErr := writeBuffer.WriteSerializable(m.GetHighLimit())
		if popErr := writeBuffer.PopContext("highLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for highLimit")
		}
		if _highLimitErr != nil {
			return errors.Wrap(_highLimitErr, "Error serializing 'highLimit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccumulatorHighLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) isBACnetConstructedDataAccumulatorHighLimit() bool {
	return true
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
