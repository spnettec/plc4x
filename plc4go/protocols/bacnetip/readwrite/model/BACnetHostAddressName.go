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

// BACnetHostAddressName is the data-structure of this message
type BACnetHostAddressName struct {
	*BACnetHostAddress
	Name *BACnetContextTagCharacterString
}

// IBACnetHostAddressName is the corresponding interface of BACnetHostAddressName
type IBACnetHostAddressName interface {
	IBACnetHostAddress
	// GetName returns Name (property field)
	GetName() *BACnetContextTagCharacterString
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

func (m *BACnetHostAddressName) InitializeParent(parent *BACnetHostAddress, peekedTagHeader *BACnetTagHeader) {
	m.BACnetHostAddress.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetHostAddressName) GetParent() *BACnetHostAddress {
	return m.BACnetHostAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetHostAddressName) GetName() *BACnetContextTagCharacterString {
	return m.Name
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostAddressName factory function for BACnetHostAddressName
func NewBACnetHostAddressName(name *BACnetContextTagCharacterString, peekedTagHeader *BACnetTagHeader) *BACnetHostAddressName {
	_result := &BACnetHostAddressName{
		Name:              name,
		BACnetHostAddress: NewBACnetHostAddress(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetHostAddressName(structType interface{}) *BACnetHostAddressName {
	if casted, ok := structType.(BACnetHostAddressName); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetHostAddressName); ok {
		return casted
	}
	if casted, ok := structType.(BACnetHostAddress); ok {
		return CastBACnetHostAddressName(casted.Child)
	}
	if casted, ok := structType.(*BACnetHostAddress); ok {
		return CastBACnetHostAddressName(casted.Child)
	}
	return nil
}

func (m *BACnetHostAddressName) GetTypeName() string {
	return "BACnetHostAddressName"
}

func (m *BACnetHostAddressName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetHostAddressName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetHostAddressName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetHostAddressNameParse(readBuffer utils.ReadBuffer) (*BACnetHostAddressName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddressName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddressName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (name)
	if pullErr := readBuffer.PullContext("name"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for name")
	}
	_name, _nameErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field")
	}
	name := CastBACnetContextTagCharacterString(_name)
	if closeErr := readBuffer.CloseContext("name"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for name")
	}

	if closeErr := readBuffer.CloseContext("BACnetHostAddressName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddressName")
	}

	// Create a partially initialized instance
	_child := &BACnetHostAddressName{
		Name:              CastBACnetContextTagCharacterString(name),
		BACnetHostAddress: &BACnetHostAddress{},
	}
	_child.BACnetHostAddress.Child = _child
	return _child, nil
}

func (m *BACnetHostAddressName) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetHostAddressName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetHostAddressName")
		}

		// Simple Field (name)
		if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for name")
		}
		_nameErr := m.Name.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("name"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for name")
		}
		if _nameErr != nil {
			return errors.Wrap(_nameErr, "Error serializing 'name' field")
		}

		if popErr := writeBuffer.PopContext("BACnetHostAddressName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetHostAddressName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetHostAddressName) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
