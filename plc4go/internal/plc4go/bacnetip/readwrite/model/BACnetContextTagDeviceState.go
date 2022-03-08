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
type BACnetContextTagDeviceState struct {
	*BACnetContextTag
	State BACnetDeviceState

	// Arguments.
	TagNumberArgument        uint8
	IsNotOpeningOrClosingTag bool
}

// The corresponding interface
type IBACnetContextTagDeviceState interface {
	IBACnetContextTag
	// GetState returns State (property field)
	GetState() BACnetDeviceState
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
func (m *BACnetContextTagDeviceState) GetDataType() BACnetDataType {
	return BACnetDataType_BACNET_DEVICE_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetContextTagDeviceState) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

func (m *BACnetContextTagDeviceState) GetParent() *BACnetContextTag {
	return m.BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetContextTagDeviceState) GetState() BACnetDeviceState {
	return m.State
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagDeviceState factory function for BACnetContextTagDeviceState
func NewBACnetContextTagDeviceState(state BACnetDeviceState, header *BACnetTagHeader, tagNumberArgument uint8, isNotOpeningOrClosingTag bool) *BACnetContextTagDeviceState {
	_result := &BACnetContextTagDeviceState{
		State:            state,
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetContextTagDeviceState(structType interface{}) *BACnetContextTagDeviceState {
	if casted, ok := structType.(BACnetContextTagDeviceState); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetContextTagDeviceState); ok {
		return casted
	}
	if casted, ok := structType.(BACnetContextTag); ok {
		return CastBACnetContextTagDeviceState(casted.Child)
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return CastBACnetContextTagDeviceState(casted.Child)
	}
	return nil
}

func (m *BACnetContextTagDeviceState) GetTypeName() string {
	return "BACnetContextTagDeviceState"
}

func (m *BACnetContextTagDeviceState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagDeviceState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (state)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetContextTagDeviceState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagDeviceStateParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool) (*BACnetContextTagDeviceState, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagDeviceState"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Simple Field (state)
	if pullErr := readBuffer.PullContext("state"); pullErr != nil {
		return nil, pullErr
	}
	_state, _stateErr := BACnetDeviceStateParse(readBuffer)
	if _stateErr != nil {
		return nil, errors.Wrap(_stateErr, "Error parsing 'state' field")
	}
	state := _state
	if closeErr := readBuffer.CloseContext("state"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagDeviceState"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagDeviceState{
		State:            state,
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child, nil
}

func (m *BACnetContextTagDeviceState) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagDeviceState"); pushErr != nil {
			return pushErr
		}

		// Simple Field (state)
		if pushErr := writeBuffer.PushContext("state"); pushErr != nil {
			return pushErr
		}
		_stateErr := m.State.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("state"); popErr != nil {
			return popErr
		}
		if _stateErr != nil {
			return errors.Wrap(_stateErr, "Error serializing 'state' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagDeviceState"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagDeviceState) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
