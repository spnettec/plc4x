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

// BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble is the data-structure of this message
type BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble struct {
	*BACnetFaultParameterFaultOutOfRangeMinNormalValue
	DoubleValue *BACnetApplicationTagDouble

	// Arguments.
	TagNumber uint8
}

// IBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
type IBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble interface {
	IBACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() *BACnetApplicationTagDouble
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) InitializeParent(parent *BACnetFaultParameterFaultOutOfRangeMinNormalValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValue.OpeningTag = openingTag
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValue.PeekedTagHeader = peekedTagHeader
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValue.ClosingTag = closingTag
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetParent() *BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetDoubleValue() *BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble factory function for BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble(doubleValue *BACnetApplicationTagDouble, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble {
	_result := &BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble{
		DoubleValue: doubleValue,
		BACnetFaultParameterFaultOutOfRangeMinNormalValue: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble(structType interface{}) *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return CastBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return CastBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueDoubleParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doubleValue)
	if pullErr := readBuffer.PullContext("doubleValue"); pullErr != nil {
		return nil, pullErr
	}
	_doubleValue, _doubleValueErr := BACnetApplicationTagParse(readBuffer)
	if _doubleValueErr != nil {
		return nil, errors.Wrap(_doubleValueErr, "Error parsing 'doubleValue' field")
	}
	doubleValue := CastBACnetApplicationTagDouble(_doubleValue)
	if closeErr := readBuffer.CloseContext("doubleValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble{
		DoubleValue: CastBACnetApplicationTagDouble(doubleValue),
		BACnetFaultParameterFaultOutOfRangeMinNormalValue: &BACnetFaultParameterFaultOutOfRangeMinNormalValue{},
	}
	_child.BACnetFaultParameterFaultOutOfRangeMinNormalValue.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); pushErr != nil {
			return pushErr
		}

		// Simple Field (doubleValue)
		if pushErr := writeBuffer.PushContext("doubleValue"); pushErr != nil {
			return pushErr
		}
		_doubleValueErr := m.DoubleValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("doubleValue"); popErr != nil {
			return popErr
		}
		if _doubleValueErr != nil {
			return errors.Wrap(_doubleValueErr, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
