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

// BACnetConstructedDataDirectReading is the data-structure of this message
type BACnetConstructedDataDirectReading struct {
	*BACnetConstructedData
	DirectReading *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDirectReading is the corresponding interface of BACnetConstructedDataDirectReading
type IBACnetConstructedDataDirectReading interface {
	IBACnetConstructedData
	// GetDirectReading returns DirectReading (property field)
	GetDirectReading() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataDirectReading) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDirectReading) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DIRECT_READING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDirectReading) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDirectReading) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDirectReading) GetDirectReading() *BACnetApplicationTagReal {
	return m.DirectReading
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDirectReading factory function for BACnetConstructedDataDirectReading
func NewBACnetConstructedDataDirectReading(directReading *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDirectReading {
	_result := &BACnetConstructedDataDirectReading{
		DirectReading:         directReading,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDirectReading(structType interface{}) *BACnetConstructedDataDirectReading {
	if casted, ok := structType.(BACnetConstructedDataDirectReading); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDirectReading); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDirectReading(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDirectReading(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDirectReading) GetTypeName() string {
	return "BACnetConstructedDataDirectReading"
}

func (m *BACnetConstructedDataDirectReading) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDirectReading) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (directReading)
	lengthInBits += m.DirectReading.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDirectReading) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDirectReadingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDirectReading, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDirectReading"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDirectReading")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (directReading)
	if pullErr := readBuffer.PullContext("directReading"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for directReading")
	}
	_directReading, _directReadingErr := BACnetApplicationTagParse(readBuffer)
	if _directReadingErr != nil {
		return nil, errors.Wrap(_directReadingErr, "Error parsing 'directReading' field")
	}
	directReading := CastBACnetApplicationTagReal(_directReading)
	if closeErr := readBuffer.CloseContext("directReading"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for directReading")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDirectReading"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDirectReading")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDirectReading{
		DirectReading:         CastBACnetApplicationTagReal(directReading),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDirectReading) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDirectReading"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDirectReading")
		}

		// Simple Field (directReading)
		if pushErr := writeBuffer.PushContext("directReading"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for directReading")
		}
		_directReadingErr := m.DirectReading.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("directReading"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for directReading")
		}
		if _directReadingErr != nil {
			return errors.Wrap(_directReadingErr, "Error serializing 'directReading' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDirectReading"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDirectReading")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDirectReading) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
