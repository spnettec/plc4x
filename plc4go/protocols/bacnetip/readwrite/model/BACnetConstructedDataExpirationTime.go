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

// BACnetConstructedDataExpirationTime is the data-structure of this message
type BACnetConstructedDataExpirationTime struct {
	*BACnetConstructedData
	ExpirationTime *BACnetDateTime

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataExpirationTime is the corresponding interface of BACnetConstructedDataExpirationTime
type IBACnetConstructedDataExpirationTime interface {
	IBACnetConstructedData
	// GetExpirationTime returns ExpirationTime (property field)
	GetExpirationTime() *BACnetDateTime
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

func (m *BACnetConstructedDataExpirationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataExpirationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXPIRATION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataExpirationTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataExpirationTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataExpirationTime) GetExpirationTime() *BACnetDateTime {
	return m.ExpirationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataExpirationTime factory function for BACnetConstructedDataExpirationTime
func NewBACnetConstructedDataExpirationTime(expirationTime *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataExpirationTime {
	_result := &BACnetConstructedDataExpirationTime{
		ExpirationTime:        expirationTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataExpirationTime(structType interface{}) *BACnetConstructedDataExpirationTime {
	if casted, ok := structType.(BACnetConstructedDataExpirationTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExpirationTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataExpirationTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataExpirationTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataExpirationTime) GetTypeName() string {
	return "BACnetConstructedDataExpirationTime"
}

func (m *BACnetConstructedDataExpirationTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataExpirationTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (expirationTime)
	lengthInBits += m.ExpirationTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataExpirationTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataExpirationTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataExpirationTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExpirationTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (expirationTime)
	if pullErr := readBuffer.PullContext("expirationTime"); pullErr != nil {
		return nil, pullErr
	}
	_expirationTime, _expirationTimeErr := BACnetDateTimeParse(readBuffer)
	if _expirationTimeErr != nil {
		return nil, errors.Wrap(_expirationTimeErr, "Error parsing 'expirationTime' field")
	}
	expirationTime := CastBACnetDateTime(_expirationTime)
	if closeErr := readBuffer.CloseContext("expirationTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExpirationTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataExpirationTime{
		ExpirationTime:        CastBACnetDateTime(expirationTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataExpirationTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExpirationTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (expirationTime)
		if pushErr := writeBuffer.PushContext("expirationTime"); pushErr != nil {
			return pushErr
		}
		_expirationTimeErr := m.ExpirationTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("expirationTime"); popErr != nil {
			return popErr
		}
		if _expirationTimeErr != nil {
			return errors.Wrap(_expirationTimeErr, "Error serializing 'expirationTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExpirationTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataExpirationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
