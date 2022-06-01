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

// BACnetOptionalUnsignedValue is the data-structure of this message
type BACnetOptionalUnsignedValue struct {
	*BACnetOptionalUnsigned
	UnsignedValue *BACnetApplicationTagUnsignedInteger
}

// IBACnetOptionalUnsignedValue is the corresponding interface of BACnetOptionalUnsignedValue
type IBACnetOptionalUnsignedValue interface {
	IBACnetOptionalUnsigned
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetOptionalUnsignedValue) InitializeParent(parent *BACnetOptionalUnsigned, peekedTagHeader *BACnetTagHeader) {
	m.BACnetOptionalUnsigned.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetOptionalUnsignedValue) GetParent() *BACnetOptionalUnsigned {
	return m.BACnetOptionalUnsigned
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetOptionalUnsignedValue) GetUnsignedValue() *BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalUnsignedValue factory function for BACnetOptionalUnsignedValue
func NewBACnetOptionalUnsignedValue(unsignedValue *BACnetApplicationTagUnsignedInteger, peekedTagHeader *BACnetTagHeader) *BACnetOptionalUnsignedValue {
	_result := &BACnetOptionalUnsignedValue{
		UnsignedValue:          unsignedValue,
		BACnetOptionalUnsigned: NewBACnetOptionalUnsigned(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetOptionalUnsignedValue(structType interface{}) *BACnetOptionalUnsignedValue {
	if casted, ok := structType.(BACnetOptionalUnsignedValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetOptionalUnsignedValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetOptionalUnsigned); ok {
		return CastBACnetOptionalUnsignedValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetOptionalUnsigned); ok {
		return CastBACnetOptionalUnsignedValue(casted.Child)
	}
	return nil
}

func (m *BACnetOptionalUnsignedValue) GetTypeName() string {
	return "BACnetOptionalUnsignedValue"
}

func (m *BACnetOptionalUnsignedValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetOptionalUnsignedValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetOptionalUnsignedValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetOptionalUnsignedValueParse(readBuffer utils.ReadBuffer) (*BACnetOptionalUnsignedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalUnsignedValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, pullErr
	}
	_unsignedValue, _unsignedValueErr := BACnetApplicationTagParse(readBuffer)
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field")
	}
	unsignedValue := CastBACnetApplicationTagUnsignedInteger(_unsignedValue)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalUnsignedValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetOptionalUnsignedValue{
		UnsignedValue:          CastBACnetApplicationTagUnsignedInteger(unsignedValue),
		BACnetOptionalUnsigned: &BACnetOptionalUnsigned{},
	}
	_child.BACnetOptionalUnsigned.Child = _child
	return _child, nil
}

func (m *BACnetOptionalUnsignedValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalUnsignedValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (unsignedValue)
		if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
			return pushErr
		}
		_unsignedValueErr := m.UnsignedValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
			return popErr
		}
		if _unsignedValueErr != nil {
			return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalUnsignedValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetOptionalUnsignedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
