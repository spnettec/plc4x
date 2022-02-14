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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetNotificationParametersChangeOfValueNewValueChangedValue struct {
	*BACnetNotificationParametersChangeOfValueNewValue
	ChangedValue *BACnetContextTagReal

	// Arguments.
	TagNumber uint8
}

// The corresponding interface
type IBACnetNotificationParametersChangeOfValueNewValueChangedValue interface {
	// GetChangedValue returns ChangedValue
	GetChangedValue() *BACnetContextTagReal
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) PeekedTagNumber() uint8 {
	return uint8(1)
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetPeekedTagNumber() uint8 {
	return uint8(1)
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) InitializeParent(parent *BACnetNotificationParametersChangeOfValueNewValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParametersChangeOfValueNewValue.OpeningTag = openingTag
	m.BACnetNotificationParametersChangeOfValueNewValue.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParametersChangeOfValueNewValue.ClosingTag = closingTag
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetChangedValue() *BACnetContextTagReal {
	return m.ChangedValue
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfValueNewValueChangedValue factory function for BACnetNotificationParametersChangeOfValueNewValueChangedValue
func NewBACnetNotificationParametersChangeOfValueNewValueChangedValue(changedValue *BACnetContextTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetNotificationParametersChangeOfValueNewValue {
	child := &BACnetNotificationParametersChangeOfValueNewValueChangedValue{
		ChangedValue: changedValue,
		BACnetNotificationParametersChangeOfValueNewValue: NewBACnetNotificationParametersChangeOfValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	child.Child = child
	return child.BACnetNotificationParametersChangeOfValueNewValue
}

func CastBACnetNotificationParametersChangeOfValueNewValueChangedValue(structType interface{}) *BACnetNotificationParametersChangeOfValueNewValueChangedValue {
	castFunc := func(typ interface{}) *BACnetNotificationParametersChangeOfValueNewValueChangedValue {
		if casted, ok := typ.(BACnetNotificationParametersChangeOfValueNewValueChangedValue); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetNotificationParametersChangeOfValueNewValueChangedValue); ok {
			return casted
		}
		if casted, ok := typ.(BACnetNotificationParametersChangeOfValueNewValue); ok {
			return CastBACnetNotificationParametersChangeOfValueNewValueChangedValue(casted.Child)
		}
		if casted, ok := typ.(*BACnetNotificationParametersChangeOfValueNewValue); ok {
			return CastBACnetNotificationParametersChangeOfValueNewValueChangedValue(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValueNewValueChangedValue"
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (changedValue)
	lengthInBits += m.ChangedValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfValueNewValueChangedValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, peekedTagNumber uint8) (*BACnetNotificationParametersChangeOfValueNewValue, error) {
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (changedValue)
	if pullErr := readBuffer.PullContext("changedValue"); pullErr != nil {
		return nil, pullErr
	}
	_changedValue, _changedValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType_REAL)
	if _changedValueErr != nil {
		return nil, errors.Wrap(_changedValueErr, "Error parsing 'changedValue' field")
	}
	changedValue := CastBACnetContextTagReal(_changedValue)
	if closeErr := readBuffer.CloseContext("changedValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfValueNewValueChangedValue{
		ChangedValue: CastBACnetContextTagReal(changedValue),
		BACnetNotificationParametersChangeOfValueNewValue: &BACnetNotificationParametersChangeOfValueNewValue{},
	}
	_child.BACnetNotificationParametersChangeOfValueNewValue.Child = _child
	return _child.BACnetNotificationParametersChangeOfValueNewValue, nil
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (changedValue)
		if pushErr := writeBuffer.PushContext("changedValue"); pushErr != nil {
			return pushErr
		}
		_changedValueErr := m.ChangedValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("changedValue"); popErr != nil {
			return popErr
		}
		if _changedValueErr != nil {
			return errors.Wrap(_changedValueErr, "Error serializing 'changedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
