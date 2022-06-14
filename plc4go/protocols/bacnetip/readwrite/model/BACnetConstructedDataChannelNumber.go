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

// BACnetConstructedDataChannelNumber is the data-structure of this message
type BACnetConstructedDataChannelNumber struct {
	*BACnetConstructedData
	ChannelNumber *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataChannelNumber is the corresponding interface of BACnetConstructedDataChannelNumber
type IBACnetConstructedDataChannelNumber interface {
	IBACnetConstructedData
	// GetChannelNumber returns ChannelNumber (property field)
	GetChannelNumber() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataChannelNumber) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataChannelNumber) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CHANNEL_NUMBER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataChannelNumber) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataChannelNumber) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataChannelNumber) GetChannelNumber() *BACnetApplicationTagUnsignedInteger {
	return m.ChannelNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataChannelNumber factory function for BACnetConstructedDataChannelNumber
func NewBACnetConstructedDataChannelNumber(channelNumber *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataChannelNumber {
	_result := &BACnetConstructedDataChannelNumber{
		ChannelNumber:         channelNumber,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataChannelNumber(structType interface{}) *BACnetConstructedDataChannelNumber {
	if casted, ok := structType.(BACnetConstructedDataChannelNumber); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChannelNumber); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataChannelNumber(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataChannelNumber(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataChannelNumber) GetTypeName() string {
	return "BACnetConstructedDataChannelNumber"
}

func (m *BACnetConstructedDataChannelNumber) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataChannelNumber) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (channelNumber)
	lengthInBits += m.ChannelNumber.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataChannelNumber) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataChannelNumberParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataChannelNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChannelNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChannelNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (channelNumber)
	if pullErr := readBuffer.PullContext("channelNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for channelNumber")
	}
	_channelNumber, _channelNumberErr := BACnetApplicationTagParse(readBuffer)
	if _channelNumberErr != nil {
		return nil, errors.Wrap(_channelNumberErr, "Error parsing 'channelNumber' field")
	}
	channelNumber := CastBACnetApplicationTagUnsignedInteger(_channelNumber)
	if closeErr := readBuffer.CloseContext("channelNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for channelNumber")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChannelNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChannelNumber")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataChannelNumber{
		ChannelNumber:         CastBACnetApplicationTagUnsignedInteger(channelNumber),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataChannelNumber) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChannelNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChannelNumber")
		}

		// Simple Field (channelNumber)
		if pushErr := writeBuffer.PushContext("channelNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for channelNumber")
		}
		_channelNumberErr := writeBuffer.WriteSerializable(m.ChannelNumber)
		if popErr := writeBuffer.PopContext("channelNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for channelNumber")
		}
		if _channelNumberErr != nil {
			return errors.Wrap(_channelNumberErr, "Error serializing 'channelNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChannelNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChannelNumber")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataChannelNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
