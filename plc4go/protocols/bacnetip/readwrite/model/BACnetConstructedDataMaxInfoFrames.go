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

// BACnetConstructedDataMaxInfoFrames is the data-structure of this message
type BACnetConstructedDataMaxInfoFrames struct {
	*BACnetConstructedData
	MaxInfoFrames *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMaxInfoFrames is the corresponding interface of BACnetConstructedDataMaxInfoFrames
type IBACnetConstructedDataMaxInfoFrames interface {
	IBACnetConstructedData
	// GetMaxInfoFrames returns MaxInfoFrames (property field)
	GetMaxInfoFrames() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataMaxInfoFrames) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMaxInfoFrames) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_INFO_FRAMES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMaxInfoFrames) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMaxInfoFrames) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMaxInfoFrames) GetMaxInfoFrames() *BACnetApplicationTagUnsignedInteger {
	return m.MaxInfoFrames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaxInfoFrames factory function for BACnetConstructedDataMaxInfoFrames
func NewBACnetConstructedDataMaxInfoFrames(maxInfoFrames *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMaxInfoFrames {
	_result := &BACnetConstructedDataMaxInfoFrames{
		MaxInfoFrames:         maxInfoFrames,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMaxInfoFrames(structType interface{}) *BACnetConstructedDataMaxInfoFrames {
	if casted, ok := structType.(BACnetConstructedDataMaxInfoFrames); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxInfoFrames); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaxInfoFrames(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaxInfoFrames(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMaxInfoFrames) GetTypeName() string {
	return "BACnetConstructedDataMaxInfoFrames"
}

func (m *BACnetConstructedDataMaxInfoFrames) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMaxInfoFrames) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxInfoFrames)
	lengthInBits += m.MaxInfoFrames.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMaxInfoFrames) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaxInfoFramesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMaxInfoFrames, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxInfoFrames")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxInfoFrames)
	if pullErr := readBuffer.PullContext("maxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxInfoFrames")
	}
	_maxInfoFrames, _maxInfoFramesErr := BACnetApplicationTagParse(readBuffer)
	if _maxInfoFramesErr != nil {
		return nil, errors.Wrap(_maxInfoFramesErr, "Error parsing 'maxInfoFrames' field")
	}
	maxInfoFrames := CastBACnetApplicationTagUnsignedInteger(_maxInfoFrames)
	if closeErr := readBuffer.CloseContext("maxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxInfoFrames")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxInfoFrames")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMaxInfoFrames{
		MaxInfoFrames:         CastBACnetApplicationTagUnsignedInteger(maxInfoFrames),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMaxInfoFrames) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxInfoFrames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxInfoFrames")
		}

		// Simple Field (maxInfoFrames)
		if pushErr := writeBuffer.PushContext("maxInfoFrames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxInfoFrames")
		}
		_maxInfoFramesErr := m.MaxInfoFrames.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("maxInfoFrames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxInfoFrames")
		}
		if _maxInfoFramesErr != nil {
			return errors.Wrap(_maxInfoFramesErr, "Error serializing 'maxInfoFrames' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxInfoFrames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxInfoFrames")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMaxInfoFrames) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
