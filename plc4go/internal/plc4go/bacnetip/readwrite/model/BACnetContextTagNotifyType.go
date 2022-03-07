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
type BACnetContextTagNotifyType struct {
	*BACnetContextTag
	Value BACnetNotifyType

	// Arguments.
	TagNumberArgument        uint8
	IsNotOpeningOrClosingTag bool
	ActualLength             uint32
}

// The corresponding interface
type IBACnetContextTagNotifyType interface {
	IBACnetContextTag
	// GetValue returns Value (property field)
	GetValue() BACnetNotifyType
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
func (m *BACnetContextTagNotifyType) DataType() BACnetDataType {
	return BACnetDataType_NOTIFY_TYPE
}

func (m *BACnetContextTagNotifyType) GetDataType() BACnetDataType {
	return BACnetDataType_NOTIFY_TYPE
}

func (m *BACnetContextTagNotifyType) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagNotifyType) GetValue() BACnetNotifyType {
	return m.Value
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetContextTagNotifyType factory function for BACnetContextTagNotifyType
func NewBACnetContextTagNotifyType(value BACnetNotifyType, header *BACnetTagHeader, tagNumberArgument uint8, isNotOpeningOrClosingTag bool, actualLength uint32) *BACnetContextTag {
	child := &BACnetContextTagNotifyType{
		Value:            value,
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagNotifyType(structType interface{}) *BACnetContextTagNotifyType {
	if casted, ok := structType.(BACnetContextTagNotifyType); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetContextTagNotifyType); ok {
		return casted
	}
	if casted, ok := structType.(BACnetContextTag); ok {
		return CastBACnetContextTagNotifyType(casted.Child)
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return CastBACnetContextTagNotifyType(casted.Child)
	}
	return nil
}

func (m *BACnetContextTagNotifyType) GetTypeName() string {
	return "BACnetContextTagNotifyType"
}

func (m *BACnetContextTagNotifyType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagNotifyType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (value)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetContextTagNotifyType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagNotifyTypeParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool, actualLength uint32) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagNotifyType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, pullErr
	}
	_value, _valueErr := BACnetNotifyTypeParse(readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagNotifyType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagNotifyType{
		Value:            value,
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagNotifyType) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagNotifyType"); pushErr != nil {
			return pushErr
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return pushErr
		}
		_valueErr := m.Value.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return popErr
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagNotifyType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagNotifyType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
