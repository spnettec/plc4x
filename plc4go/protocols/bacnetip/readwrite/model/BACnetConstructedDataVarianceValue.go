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

// BACnetConstructedDataVarianceValue is the data-structure of this message
type BACnetConstructedDataVarianceValue struct {
	*BACnetConstructedData
	VarianceValue *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataVarianceValue is the corresponding interface of BACnetConstructedDataVarianceValue
type IBACnetConstructedDataVarianceValue interface {
	IBACnetConstructedData
	// GetVarianceValue returns VarianceValue (property field)
	GetVarianceValue() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataVarianceValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataVarianceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VARIANCE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataVarianceValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataVarianceValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataVarianceValue) GetVarianceValue() *BACnetApplicationTagReal {
	return m.VarianceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVarianceValue factory function for BACnetConstructedDataVarianceValue
func NewBACnetConstructedDataVarianceValue(varianceValue *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataVarianceValue {
	_result := &BACnetConstructedDataVarianceValue{
		VarianceValue:         varianceValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataVarianceValue(structType interface{}) *BACnetConstructedDataVarianceValue {
	if casted, ok := structType.(BACnetConstructedDataVarianceValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVarianceValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataVarianceValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataVarianceValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataVarianceValue) GetTypeName() string {
	return "BACnetConstructedDataVarianceValue"
}

func (m *BACnetConstructedDataVarianceValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataVarianceValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (varianceValue)
	lengthInBits += m.VarianceValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataVarianceValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataVarianceValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataVarianceValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVarianceValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (varianceValue)
	if pullErr := readBuffer.PullContext("varianceValue"); pullErr != nil {
		return nil, pullErr
	}
	_varianceValue, _varianceValueErr := BACnetApplicationTagParse(readBuffer)
	if _varianceValueErr != nil {
		return nil, errors.Wrap(_varianceValueErr, "Error parsing 'varianceValue' field")
	}
	varianceValue := CastBACnetApplicationTagReal(_varianceValue)
	if closeErr := readBuffer.CloseContext("varianceValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVarianceValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataVarianceValue{
		VarianceValue:         CastBACnetApplicationTagReal(varianceValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataVarianceValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVarianceValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (varianceValue)
		if pushErr := writeBuffer.PushContext("varianceValue"); pushErr != nil {
			return pushErr
		}
		_varianceValueErr := m.VarianceValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("varianceValue"); popErr != nil {
			return popErr
		}
		if _varianceValueErr != nil {
			return errors.Wrap(_varianceValueErr, "Error serializing 'varianceValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVarianceValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataVarianceValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
