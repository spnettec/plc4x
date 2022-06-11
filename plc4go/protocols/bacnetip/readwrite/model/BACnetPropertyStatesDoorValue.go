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

// BACnetPropertyStatesDoorValue is the data-structure of this message
type BACnetPropertyStatesDoorValue struct {
	*BACnetPropertyStates
	DoorValue *BACnetDoorValueTagged
}

// IBACnetPropertyStatesDoorValue is the corresponding interface of BACnetPropertyStatesDoorValue
type IBACnetPropertyStatesDoorValue interface {
	IBACnetPropertyStates
	// GetDoorValue returns DoorValue (property field)
	GetDoorValue() *BACnetDoorValueTagged
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

func (m *BACnetPropertyStatesDoorValue) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesDoorValue) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesDoorValue) GetDoorValue() *BACnetDoorValueTagged {
	return m.DoorValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesDoorValue factory function for BACnetPropertyStatesDoorValue
func NewBACnetPropertyStatesDoorValue(doorValue *BACnetDoorValueTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesDoorValue {
	_result := &BACnetPropertyStatesDoorValue{
		DoorValue:            doorValue,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesDoorValue(structType interface{}) *BACnetPropertyStatesDoorValue {
	if casted, ok := structType.(BACnetPropertyStatesDoorValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesDoorValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesDoorValue(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesDoorValue) GetTypeName() string {
	return "BACnetPropertyStatesDoorValue"
}

func (m *BACnetPropertyStatesDoorValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesDoorValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorValue)
	lengthInBits += m.DoorValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesDoorValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesDoorValueParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesDoorValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesDoorValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorValue)
	if pullErr := readBuffer.PullContext("doorValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorValue")
	}
	_doorValue, _doorValueErr := BACnetDoorValueTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _doorValueErr != nil {
		return nil, errors.Wrap(_doorValueErr, "Error parsing 'doorValue' field")
	}
	doorValue := CastBACnetDoorValueTagged(_doorValue)
	if closeErr := readBuffer.CloseContext("doorValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesDoorValue")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesDoorValue{
		DoorValue:            CastBACnetDoorValueTagged(doorValue),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesDoorValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesDoorValue")
		}

		// Simple Field (doorValue)
		if pushErr := writeBuffer.PushContext("doorValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorValue")
		}
		_doorValueErr := m.DoorValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("doorValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorValue")
		}
		if _doorValueErr != nil {
			return errors.Wrap(_doorValueErr, "Error serializing 'doorValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesDoorValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesDoorValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
