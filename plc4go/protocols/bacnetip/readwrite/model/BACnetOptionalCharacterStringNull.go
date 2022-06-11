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

// BACnetOptionalCharacterStringNull is the data-structure of this message
type BACnetOptionalCharacterStringNull struct {
	*BACnetOptionalCharacterString
	NullValue *BACnetApplicationTagNull
}

// IBACnetOptionalCharacterStringNull is the corresponding interface of BACnetOptionalCharacterStringNull
type IBACnetOptionalCharacterStringNull interface {
	IBACnetOptionalCharacterString
	// GetNullValue returns NullValue (property field)
	GetNullValue() *BACnetApplicationTagNull
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

func (m *BACnetOptionalCharacterStringNull) InitializeParent(parent *BACnetOptionalCharacterString, peekedTagHeader *BACnetTagHeader) {
	m.BACnetOptionalCharacterString.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetOptionalCharacterStringNull) GetParent() *BACnetOptionalCharacterString {
	return m.BACnetOptionalCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetOptionalCharacterStringNull) GetNullValue() *BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalCharacterStringNull factory function for BACnetOptionalCharacterStringNull
func NewBACnetOptionalCharacterStringNull(nullValue *BACnetApplicationTagNull, peekedTagHeader *BACnetTagHeader) *BACnetOptionalCharacterStringNull {
	_result := &BACnetOptionalCharacterStringNull{
		NullValue:                     nullValue,
		BACnetOptionalCharacterString: NewBACnetOptionalCharacterString(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetOptionalCharacterStringNull(structType interface{}) *BACnetOptionalCharacterStringNull {
	if casted, ok := structType.(BACnetOptionalCharacterStringNull); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetOptionalCharacterStringNull); ok {
		return casted
	}
	if casted, ok := structType.(BACnetOptionalCharacterString); ok {
		return CastBACnetOptionalCharacterStringNull(casted.Child)
	}
	if casted, ok := structType.(*BACnetOptionalCharacterString); ok {
		return CastBACnetOptionalCharacterStringNull(casted.Child)
	}
	return nil
}

func (m *BACnetOptionalCharacterStringNull) GetTypeName() string {
	return "BACnetOptionalCharacterStringNull"
}

func (m *BACnetOptionalCharacterStringNull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetOptionalCharacterStringNull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetOptionalCharacterStringNull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetOptionalCharacterStringNullParse(readBuffer utils.ReadBuffer) (*BACnetOptionalCharacterStringNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalCharacterStringNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalCharacterStringNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nullValue)
	if pullErr := readBuffer.PullContext("nullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nullValue")
	}
	_nullValue, _nullValueErr := BACnetApplicationTagParse(readBuffer)
	if _nullValueErr != nil {
		return nil, errors.Wrap(_nullValueErr, "Error parsing 'nullValue' field")
	}
	nullValue := CastBACnetApplicationTagNull(_nullValue)
	if closeErr := readBuffer.CloseContext("nullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nullValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalCharacterStringNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalCharacterStringNull")
	}

	// Create a partially initialized instance
	_child := &BACnetOptionalCharacterStringNull{
		NullValue:                     CastBACnetApplicationTagNull(nullValue),
		BACnetOptionalCharacterString: &BACnetOptionalCharacterString{},
	}
	_child.BACnetOptionalCharacterString.Child = _child
	return _child, nil
}

func (m *BACnetOptionalCharacterStringNull) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalCharacterStringNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalCharacterStringNull")
		}

		// Simple Field (nullValue)
		if pushErr := writeBuffer.PushContext("nullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nullValue")
		}
		_nullValueErr := m.NullValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("nullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nullValue")
		}
		if _nullValueErr != nil {
			return errors.Wrap(_nullValueErr, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalCharacterStringNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalCharacterStringNull")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetOptionalCharacterStringNull) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
