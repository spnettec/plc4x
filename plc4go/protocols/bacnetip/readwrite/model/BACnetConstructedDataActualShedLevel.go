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

// BACnetConstructedDataActualShedLevel is the data-structure of this message
type BACnetConstructedDataActualShedLevel struct {
	*BACnetConstructedData
	ActualShedLevel *BACnetShedLevel

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataActualShedLevel is the corresponding interface of BACnetConstructedDataActualShedLevel
type IBACnetConstructedDataActualShedLevel interface {
	IBACnetConstructedData
	// GetActualShedLevel returns ActualShedLevel (property field)
	GetActualShedLevel() *BACnetShedLevel
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

func (m *BACnetConstructedDataActualShedLevel) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataActualShedLevel) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTUAL_SHED_LEVEL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataActualShedLevel) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataActualShedLevel) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataActualShedLevel) GetActualShedLevel() *BACnetShedLevel {
	return m.ActualShedLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataActualShedLevel factory function for BACnetConstructedDataActualShedLevel
func NewBACnetConstructedDataActualShedLevel(actualShedLevel *BACnetShedLevel, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataActualShedLevel {
	_result := &BACnetConstructedDataActualShedLevel{
		ActualShedLevel:       actualShedLevel,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataActualShedLevel(structType interface{}) *BACnetConstructedDataActualShedLevel {
	if casted, ok := structType.(BACnetConstructedDataActualShedLevel); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActualShedLevel); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataActualShedLevel(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataActualShedLevel(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataActualShedLevel) GetTypeName() string {
	return "BACnetConstructedDataActualShedLevel"
}

func (m *BACnetConstructedDataActualShedLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataActualShedLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (actualShedLevel)
	lengthInBits += m.ActualShedLevel.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataActualShedLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataActualShedLevelParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataActualShedLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActualShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActualShedLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (actualShedLevel)
	if pullErr := readBuffer.PullContext("actualShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for actualShedLevel")
	}
	_actualShedLevel, _actualShedLevelErr := BACnetShedLevelParse(readBuffer)
	if _actualShedLevelErr != nil {
		return nil, errors.Wrap(_actualShedLevelErr, "Error parsing 'actualShedLevel' field")
	}
	actualShedLevel := CastBACnetShedLevel(_actualShedLevel)
	if closeErr := readBuffer.CloseContext("actualShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for actualShedLevel")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActualShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActualShedLevel")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataActualShedLevel{
		ActualShedLevel:       CastBACnetShedLevel(actualShedLevel),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataActualShedLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActualShedLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActualShedLevel")
		}

		// Simple Field (actualShedLevel)
		if pushErr := writeBuffer.PushContext("actualShedLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for actualShedLevel")
		}
		_actualShedLevelErr := m.ActualShedLevel.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("actualShedLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for actualShedLevel")
		}
		if _actualShedLevelErr != nil {
			return errors.Wrap(_actualShedLevelErr, "Error serializing 'actualShedLevel' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActualShedLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActualShedLevel")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataActualShedLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
