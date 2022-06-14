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

// BACnetConstructedDataLocalTime is the data-structure of this message
type BACnetConstructedDataLocalTime struct {
	*BACnetConstructedData
	LocalTime *BACnetApplicationTagTime

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLocalTime is the corresponding interface of BACnetConstructedDataLocalTime
type IBACnetConstructedDataLocalTime interface {
	IBACnetConstructedData
	// GetLocalTime returns LocalTime (property field)
	GetLocalTime() *BACnetApplicationTagTime
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

func (m *BACnetConstructedDataLocalTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLocalTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCAL_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLocalTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLocalTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLocalTime) GetLocalTime() *BACnetApplicationTagTime {
	return m.LocalTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLocalTime factory function for BACnetConstructedDataLocalTime
func NewBACnetConstructedDataLocalTime(localTime *BACnetApplicationTagTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLocalTime {
	_result := &BACnetConstructedDataLocalTime{
		LocalTime:             localTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLocalTime(structType interface{}) *BACnetConstructedDataLocalTime {
	if casted, ok := structType.(BACnetConstructedDataLocalTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLocalTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLocalTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLocalTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLocalTime) GetTypeName() string {
	return "BACnetConstructedDataLocalTime"
}

func (m *BACnetConstructedDataLocalTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLocalTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (localTime)
	lengthInBits += m.LocalTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLocalTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLocalTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLocalTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLocalTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLocalTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (localTime)
	if pullErr := readBuffer.PullContext("localTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for localTime")
	}
	_localTime, _localTimeErr := BACnetApplicationTagParse(readBuffer)
	if _localTimeErr != nil {
		return nil, errors.Wrap(_localTimeErr, "Error parsing 'localTime' field")
	}
	localTime := CastBACnetApplicationTagTime(_localTime)
	if closeErr := readBuffer.CloseContext("localTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for localTime")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLocalTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLocalTime")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLocalTime{
		LocalTime:             CastBACnetApplicationTagTime(localTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLocalTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLocalTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLocalTime")
		}

		// Simple Field (localTime)
		if pushErr := writeBuffer.PushContext("localTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for localTime")
		}
		_localTimeErr := writeBuffer.WriteSerializable(m.LocalTime)
		if popErr := writeBuffer.PopContext("localTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for localTime")
		}
		if _localTimeErr != nil {
			return errors.Wrap(_localTimeErr, "Error serializing 'localTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLocalTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLocalTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLocalTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
