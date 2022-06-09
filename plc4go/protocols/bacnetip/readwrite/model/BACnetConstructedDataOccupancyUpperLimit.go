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

// BACnetConstructedDataOccupancyUpperLimit is the data-structure of this message
type BACnetConstructedDataOccupancyUpperLimit struct {
	*BACnetConstructedData
	OccupancyUpperLimit *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataOccupancyUpperLimit is the corresponding interface of BACnetConstructedDataOccupancyUpperLimit
type IBACnetConstructedDataOccupancyUpperLimit interface {
	IBACnetConstructedData
	// GetOccupancyUpperLimit returns OccupancyUpperLimit (property field)
	GetOccupancyUpperLimit() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataOccupancyUpperLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataOccupancyUpperLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_UPPER_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataOccupancyUpperLimit) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataOccupancyUpperLimit) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataOccupancyUpperLimit) GetOccupancyUpperLimit() *BACnetApplicationTagUnsignedInteger {
	return m.OccupancyUpperLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOccupancyUpperLimit factory function for BACnetConstructedDataOccupancyUpperLimit
func NewBACnetConstructedDataOccupancyUpperLimit(occupancyUpperLimit *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataOccupancyUpperLimit {
	_result := &BACnetConstructedDataOccupancyUpperLimit{
		OccupancyUpperLimit:   occupancyUpperLimit,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataOccupancyUpperLimit(structType interface{}) *BACnetConstructedDataOccupancyUpperLimit {
	if casted, ok := structType.(BACnetConstructedDataOccupancyUpperLimit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyUpperLimit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataOccupancyUpperLimit(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataOccupancyUpperLimit(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataOccupancyUpperLimit) GetTypeName() string {
	return "BACnetConstructedDataOccupancyUpperLimit"
}

func (m *BACnetConstructedDataOccupancyUpperLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataOccupancyUpperLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (occupancyUpperLimit)
	lengthInBits += m.OccupancyUpperLimit.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataOccupancyUpperLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOccupancyUpperLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataOccupancyUpperLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyUpperLimit"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (occupancyUpperLimit)
	if pullErr := readBuffer.PullContext("occupancyUpperLimit"); pullErr != nil {
		return nil, pullErr
	}
	_occupancyUpperLimit, _occupancyUpperLimitErr := BACnetApplicationTagParse(readBuffer)
	if _occupancyUpperLimitErr != nil {
		return nil, errors.Wrap(_occupancyUpperLimitErr, "Error parsing 'occupancyUpperLimit' field")
	}
	occupancyUpperLimit := CastBACnetApplicationTagUnsignedInteger(_occupancyUpperLimit)
	if closeErr := readBuffer.CloseContext("occupancyUpperLimit"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyUpperLimit"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataOccupancyUpperLimit{
		OccupancyUpperLimit:   CastBACnetApplicationTagUnsignedInteger(occupancyUpperLimit),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataOccupancyUpperLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyUpperLimit"); pushErr != nil {
			return pushErr
		}

		// Simple Field (occupancyUpperLimit)
		if pushErr := writeBuffer.PushContext("occupancyUpperLimit"); pushErr != nil {
			return pushErr
		}
		_occupancyUpperLimitErr := m.OccupancyUpperLimit.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("occupancyUpperLimit"); popErr != nil {
			return popErr
		}
		if _occupancyUpperLimitErr != nil {
			return errors.Wrap(_occupancyUpperLimitErr, "Error serializing 'occupancyUpperLimit' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyUpperLimit"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataOccupancyUpperLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
