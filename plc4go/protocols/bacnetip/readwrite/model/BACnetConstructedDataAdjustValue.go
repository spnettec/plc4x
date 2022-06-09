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

// BACnetConstructedDataAdjustValue is the data-structure of this message
type BACnetConstructedDataAdjustValue struct {
	*BACnetConstructedData
	AdjustValue *BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAdjustValue is the corresponding interface of BACnetConstructedDataAdjustValue
type IBACnetConstructedDataAdjustValue interface {
	IBACnetConstructedData
	// GetAdjustValue returns AdjustValue (property field)
	GetAdjustValue() *BACnetApplicationTagSignedInteger
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

func (m *BACnetConstructedDataAdjustValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAdjustValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ADJUST_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAdjustValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAdjustValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAdjustValue) GetAdjustValue() *BACnetApplicationTagSignedInteger {
	return m.AdjustValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAdjustValue factory function for BACnetConstructedDataAdjustValue
func NewBACnetConstructedDataAdjustValue(adjustValue *BACnetApplicationTagSignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAdjustValue {
	_result := &BACnetConstructedDataAdjustValue{
		AdjustValue:           adjustValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAdjustValue(structType interface{}) *BACnetConstructedDataAdjustValue {
	if casted, ok := structType.(BACnetConstructedDataAdjustValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAdjustValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAdjustValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAdjustValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAdjustValue) GetTypeName() string {
	return "BACnetConstructedDataAdjustValue"
}

func (m *BACnetConstructedDataAdjustValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAdjustValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (adjustValue)
	lengthInBits += m.AdjustValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAdjustValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAdjustValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAdjustValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAdjustValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (adjustValue)
	if pullErr := readBuffer.PullContext("adjustValue"); pullErr != nil {
		return nil, pullErr
	}
	_adjustValue, _adjustValueErr := BACnetApplicationTagParse(readBuffer)
	if _adjustValueErr != nil {
		return nil, errors.Wrap(_adjustValueErr, "Error parsing 'adjustValue' field")
	}
	adjustValue := CastBACnetApplicationTagSignedInteger(_adjustValue)
	if closeErr := readBuffer.CloseContext("adjustValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAdjustValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAdjustValue{
		AdjustValue:           CastBACnetApplicationTagSignedInteger(adjustValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAdjustValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAdjustValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (adjustValue)
		if pushErr := writeBuffer.PushContext("adjustValue"); pushErr != nil {
			return pushErr
		}
		_adjustValueErr := m.AdjustValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("adjustValue"); popErr != nil {
			return popErr
		}
		if _adjustValueErr != nil {
			return errors.Wrap(_adjustValueErr, "Error serializing 'adjustValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAdjustValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAdjustValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
