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

// BACnetConstructedDataCarPosition is the corresponding interface of BACnetConstructedDataCarPosition
type BACnetConstructedDataCarPosition interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCarPosition returns CarPosition (property field)
	GetCarPosition() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataCarPositionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCarPosition.
// This is useful for switch cases.
type BACnetConstructedDataCarPositionExactly interface {
	BACnetConstructedDataCarPosition
	isBACnetConstructedDataCarPosition() bool
}

// _BACnetConstructedDataCarPosition is the data-structure of this message
type _BACnetConstructedDataCarPosition struct {
	*_BACnetConstructedData
	CarPosition BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarPosition) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCarPosition) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_POSITION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarPosition) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCarPosition) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarPosition) GetCarPosition() BACnetApplicationTagUnsignedInteger {
	return m.CarPosition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarPosition) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetCarPosition())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCarPosition factory function for _BACnetConstructedDataCarPosition
func NewBACnetConstructedDataCarPosition(carPosition BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCarPosition {
	_result := &_BACnetConstructedDataCarPosition{
		CarPosition:            carPosition,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarPosition(structType interface{}) BACnetConstructedDataCarPosition {
	if casted, ok := structType.(BACnetConstructedDataCarPosition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarPosition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarPosition) GetTypeName() string {
	return "BACnetConstructedDataCarPosition"
}

func (m *_BACnetConstructedDataCarPosition) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCarPosition) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (carPosition)
	lengthInBits += m.CarPosition.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCarPosition) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCarPositionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarPosition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarPosition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarPosition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (carPosition)
	if pullErr := readBuffer.PullContext("carPosition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for carPosition")
	}
	_carPosition, _carPositionErr := BACnetApplicationTagParse(readBuffer)
	if _carPositionErr != nil {
		return nil, errors.Wrap(_carPositionErr, "Error parsing 'carPosition' field of BACnetConstructedDataCarPosition")
	}
	carPosition := _carPosition.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("carPosition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for carPosition")
	}

	// Virtual field
	_actualValue := carPosition
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarPosition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarPosition")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCarPosition{
		CarPosition: carPosition,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCarPosition) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarPosition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarPosition")
		}

		// Simple Field (carPosition)
		if pushErr := writeBuffer.PushContext("carPosition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for carPosition")
		}
		_carPositionErr := writeBuffer.WriteSerializable(m.GetCarPosition())
		if popErr := writeBuffer.PopContext("carPosition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for carPosition")
		}
		if _carPositionErr != nil {
			return errors.Wrap(_carPositionErr, "Error serializing 'carPosition' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarPosition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarPosition")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCarPosition) isBACnetConstructedDataCarPosition() bool {
	return true
}

func (m *_BACnetConstructedDataCarPosition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
