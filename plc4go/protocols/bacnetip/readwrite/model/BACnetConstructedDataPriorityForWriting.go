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

// BACnetConstructedDataPriorityForWriting is the data-structure of this message
type BACnetConstructedDataPriorityForWriting struct {
	*BACnetConstructedData
	PriorityForWriting *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataPriorityForWriting is the corresponding interface of BACnetConstructedDataPriorityForWriting
type IBACnetConstructedDataPriorityForWriting interface {
	IBACnetConstructedData
	// GetPriorityForWriting returns PriorityForWriting (property field)
	GetPriorityForWriting() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataPriorityForWriting) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataPriorityForWriting) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRIORITY_FOR_WRITING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataPriorityForWriting) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataPriorityForWriting) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataPriorityForWriting) GetPriorityForWriting() *BACnetApplicationTagUnsignedInteger {
	return m.PriorityForWriting
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPriorityForWriting factory function for BACnetConstructedDataPriorityForWriting
func NewBACnetConstructedDataPriorityForWriting(priorityForWriting *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataPriorityForWriting {
	_result := &BACnetConstructedDataPriorityForWriting{
		PriorityForWriting:    priorityForWriting,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataPriorityForWriting(structType interface{}) *BACnetConstructedDataPriorityForWriting {
	if casted, ok := structType.(BACnetConstructedDataPriorityForWriting); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPriorityForWriting); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataPriorityForWriting(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataPriorityForWriting(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataPriorityForWriting) GetTypeName() string {
	return "BACnetConstructedDataPriorityForWriting"
}

func (m *BACnetConstructedDataPriorityForWriting) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataPriorityForWriting) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (priorityForWriting)
	lengthInBits += m.PriorityForWriting.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataPriorityForWriting) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPriorityForWritingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataPriorityForWriting, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPriorityForWriting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPriorityForWriting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (priorityForWriting)
	if pullErr := readBuffer.PullContext("priorityForWriting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priorityForWriting")
	}
	_priorityForWriting, _priorityForWritingErr := BACnetApplicationTagParse(readBuffer)
	if _priorityForWritingErr != nil {
		return nil, errors.Wrap(_priorityForWritingErr, "Error parsing 'priorityForWriting' field")
	}
	priorityForWriting := CastBACnetApplicationTagUnsignedInteger(_priorityForWriting)
	if closeErr := readBuffer.CloseContext("priorityForWriting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priorityForWriting")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPriorityForWriting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPriorityForWriting")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataPriorityForWriting{
		PriorityForWriting:    CastBACnetApplicationTagUnsignedInteger(priorityForWriting),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataPriorityForWriting) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPriorityForWriting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPriorityForWriting")
		}

		// Simple Field (priorityForWriting)
		if pushErr := writeBuffer.PushContext("priorityForWriting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priorityForWriting")
		}
		_priorityForWritingErr := writeBuffer.WriteSerializable(m.PriorityForWriting)
		if popErr := writeBuffer.PopContext("priorityForWriting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priorityForWriting")
		}
		if _priorityForWritingErr != nil {
			return errors.Wrap(_priorityForWritingErr, "Error serializing 'priorityForWriting' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPriorityForWriting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPriorityForWriting")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataPriorityForWriting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
