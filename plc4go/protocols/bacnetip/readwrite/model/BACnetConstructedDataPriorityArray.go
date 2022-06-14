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

// BACnetConstructedDataPriorityArray is the data-structure of this message
type BACnetConstructedDataPriorityArray struct {
	*BACnetConstructedData
	PriorityArray *BACnetPriorityArray

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataPriorityArray is the corresponding interface of BACnetConstructedDataPriorityArray
type IBACnetConstructedDataPriorityArray interface {
	IBACnetConstructedData
	// GetPriorityArray returns PriorityArray (property field)
	GetPriorityArray() *BACnetPriorityArray
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

func (m *BACnetConstructedDataPriorityArray) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataPriorityArray) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRIORITY_ARRAY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataPriorityArray) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataPriorityArray) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataPriorityArray) GetPriorityArray() *BACnetPriorityArray {
	return m.PriorityArray
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPriorityArray factory function for BACnetConstructedDataPriorityArray
func NewBACnetConstructedDataPriorityArray(priorityArray *BACnetPriorityArray, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataPriorityArray {
	_result := &BACnetConstructedDataPriorityArray{
		PriorityArray:         priorityArray,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataPriorityArray(structType interface{}) *BACnetConstructedDataPriorityArray {
	if casted, ok := structType.(BACnetConstructedDataPriorityArray); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPriorityArray); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataPriorityArray(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataPriorityArray(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataPriorityArray) GetTypeName() string {
	return "BACnetConstructedDataPriorityArray"
}

func (m *BACnetConstructedDataPriorityArray) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataPriorityArray) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (priorityArray)
	lengthInBits += m.PriorityArray.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataPriorityArray) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPriorityArrayParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataPriorityArray, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPriorityArray"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPriorityArray")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (priorityArray)
	if pullErr := readBuffer.PullContext("priorityArray"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priorityArray")
	}
	_priorityArray, _priorityArrayErr := BACnetPriorityArrayParse(readBuffer, BACnetObjectType(objectTypeArgument), uint8(tagNumber), arrayIndexArgument)
	if _priorityArrayErr != nil {
		return nil, errors.Wrap(_priorityArrayErr, "Error parsing 'priorityArray' field")
	}
	priorityArray := CastBACnetPriorityArray(_priorityArray)
	if closeErr := readBuffer.CloseContext("priorityArray"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priorityArray")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPriorityArray"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPriorityArray")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataPriorityArray{
		PriorityArray:         CastBACnetPriorityArray(priorityArray),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataPriorityArray) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPriorityArray"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPriorityArray")
		}

		// Simple Field (priorityArray)
		if pushErr := writeBuffer.PushContext("priorityArray"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priorityArray")
		}
		_priorityArrayErr := writeBuffer.WriteSerializable(m.PriorityArray)
		if popErr := writeBuffer.PopContext("priorityArray"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priorityArray")
		}
		if _priorityArrayErr != nil {
			return errors.Wrap(_priorityArrayErr, "Error serializing 'priorityArray' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPriorityArray"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPriorityArray")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataPriorityArray) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
