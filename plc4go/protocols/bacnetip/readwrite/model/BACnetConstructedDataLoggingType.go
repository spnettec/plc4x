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

// BACnetConstructedDataLoggingType is the data-structure of this message
type BACnetConstructedDataLoggingType struct {
	*BACnetConstructedData
	LoggingType *BACnetLoggingTypeTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLoggingType is the corresponding interface of BACnetConstructedDataLoggingType
type IBACnetConstructedDataLoggingType interface {
	IBACnetConstructedData
	// GetLoggingType returns LoggingType (property field)
	GetLoggingType() *BACnetLoggingTypeTagged
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

func (m *BACnetConstructedDataLoggingType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLoggingType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOGGING_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLoggingType) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLoggingType) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLoggingType) GetLoggingType() *BACnetLoggingTypeTagged {
	return m.LoggingType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLoggingType factory function for BACnetConstructedDataLoggingType
func NewBACnetConstructedDataLoggingType(loggingType *BACnetLoggingTypeTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLoggingType {
	_result := &BACnetConstructedDataLoggingType{
		LoggingType:           loggingType,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLoggingType(structType interface{}) *BACnetConstructedDataLoggingType {
	if casted, ok := structType.(BACnetConstructedDataLoggingType); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoggingType); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLoggingType(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLoggingType(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLoggingType) GetTypeName() string {
	return "BACnetConstructedDataLoggingType"
}

func (m *BACnetConstructedDataLoggingType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLoggingType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (loggingType)
	lengthInBits += m.LoggingType.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLoggingType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLoggingTypeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLoggingType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoggingType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (loggingType)
	if pullErr := readBuffer.PullContext("loggingType"); pullErr != nil {
		return nil, pullErr
	}
	_loggingType, _loggingTypeErr := BACnetLoggingTypeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _loggingTypeErr != nil {
		return nil, errors.Wrap(_loggingTypeErr, "Error parsing 'loggingType' field")
	}
	loggingType := CastBACnetLoggingTypeTagged(_loggingType)
	if closeErr := readBuffer.CloseContext("loggingType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoggingType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLoggingType{
		LoggingType:           CastBACnetLoggingTypeTagged(loggingType),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLoggingType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoggingType"); pushErr != nil {
			return pushErr
		}

		// Simple Field (loggingType)
		if pushErr := writeBuffer.PushContext("loggingType"); pushErr != nil {
			return pushErr
		}
		_loggingTypeErr := m.LoggingType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("loggingType"); popErr != nil {
			return popErr
		}
		if _loggingTypeErr != nil {
			return errors.Wrap(_loggingTypeErr, "Error serializing 'loggingType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoggingType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLoggingType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
