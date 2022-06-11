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

// BACnetConstructedDataMinActualValue is the data-structure of this message
type BACnetConstructedDataMinActualValue struct {
	*BACnetConstructedData
	MinActualValue *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMinActualValue is the corresponding interface of BACnetConstructedDataMinActualValue
type IBACnetConstructedDataMinActualValue interface {
	IBACnetConstructedData
	// GetMinActualValue returns MinActualValue (property field)
	GetMinActualValue() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataMinActualValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMinActualValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_ACTUAL_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMinActualValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMinActualValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMinActualValue) GetMinActualValue() *BACnetApplicationTagReal {
	return m.MinActualValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinActualValue factory function for BACnetConstructedDataMinActualValue
func NewBACnetConstructedDataMinActualValue(minActualValue *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMinActualValue {
	_result := &BACnetConstructedDataMinActualValue{
		MinActualValue:        minActualValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMinActualValue(structType interface{}) *BACnetConstructedDataMinActualValue {
	if casted, ok := structType.(BACnetConstructedDataMinActualValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinActualValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinActualValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinActualValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMinActualValue) GetTypeName() string {
	return "BACnetConstructedDataMinActualValue"
}

func (m *BACnetConstructedDataMinActualValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMinActualValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (minActualValue)
	lengthInBits += m.MinActualValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMinActualValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMinActualValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMinActualValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinActualValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinActualValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minActualValue)
	if pullErr := readBuffer.PullContext("minActualValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minActualValue")
	}
	_minActualValue, _minActualValueErr := BACnetApplicationTagParse(readBuffer)
	if _minActualValueErr != nil {
		return nil, errors.Wrap(_minActualValueErr, "Error parsing 'minActualValue' field")
	}
	minActualValue := CastBACnetApplicationTagReal(_minActualValue)
	if closeErr := readBuffer.CloseContext("minActualValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minActualValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinActualValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinActualValue")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMinActualValue{
		MinActualValue:        CastBACnetApplicationTagReal(minActualValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMinActualValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinActualValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinActualValue")
		}

		// Simple Field (minActualValue)
		if pushErr := writeBuffer.PushContext("minActualValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minActualValue")
		}
		_minActualValueErr := m.MinActualValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("minActualValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minActualValue")
		}
		if _minActualValueErr != nil {
			return errors.Wrap(_minActualValueErr, "Error serializing 'minActualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinActualValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinActualValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMinActualValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
