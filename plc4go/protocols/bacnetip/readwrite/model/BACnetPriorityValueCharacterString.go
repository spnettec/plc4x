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

// BACnetPriorityValueCharacterString is the corresponding interface of BACnetPriorityValueCharacterString
type BACnetPriorityValueCharacterString interface {
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetCharacterStringValue returns CharacterStringValue (property field)
	GetCharacterStringValue() BACnetApplicationTagCharacterString
}

// BACnetPriorityValueCharacterStringExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueCharacterString.
// This is useful for switch cases.
type BACnetPriorityValueCharacterStringExactly interface {
	BACnetPriorityValueCharacterString
	isBACnetPriorityValueCharacterString() bool
}

// _BACnetPriorityValueCharacterString is the data-structure of this message
type _BACnetPriorityValueCharacterString struct {
	*_BACnetPriorityValue
	CharacterStringValue BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueCharacterString) InitializeParent(parent BACnetPriorityValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueCharacterString) GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueCharacterString) GetCharacterStringValue() BACnetApplicationTagCharacterString {
	return m.CharacterStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueCharacterString factory function for _BACnetPriorityValueCharacterString
func NewBACnetPriorityValueCharacterString(characterStringValue BACnetApplicationTagCharacterString, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueCharacterString {
	_result := &_BACnetPriorityValueCharacterString{
		CharacterStringValue: characterStringValue,
		_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueCharacterString(structType interface{}) BACnetPriorityValueCharacterString {
	if casted, ok := structType.(BACnetPriorityValueCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueCharacterString) GetTypeName() string {
	return "BACnetPriorityValueCharacterString"
}

func (m *_BACnetPriorityValueCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPriorityValueCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (characterStringValue)
	lengthInBits += m.CharacterStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPriorityValueCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueCharacterStringParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (characterStringValue)
	if pullErr := readBuffer.PullContext("characterStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for characterStringValue")
	}
	_characterStringValue, _characterStringValueErr := BACnetApplicationTagParse(readBuffer)
	if _characterStringValueErr != nil {
		return nil, errors.Wrap(_characterStringValueErr, "Error parsing 'characterStringValue' field of BACnetPriorityValueCharacterString")
	}
	characterStringValue := _characterStringValue.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("characterStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for characterStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueCharacterString")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueCharacterString{
		CharacterStringValue: characterStringValue,
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueCharacterString")
		}

		// Simple Field (characterStringValue)
		if pushErr := writeBuffer.PushContext("characterStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for characterStringValue")
		}
		_characterStringValueErr := writeBuffer.WriteSerializable(m.GetCharacterStringValue())
		if popErr := writeBuffer.PopContext("characterStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for characterStringValue")
		}
		if _characterStringValueErr != nil {
			return errors.Wrap(_characterStringValueErr, "Error serializing 'characterStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueCharacterString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueCharacterString) isBACnetPriorityValueCharacterString() bool {
	return true
}

func (m *_BACnetPriorityValueCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
