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

// BACnetPropertyStatesLifeSafetyState is the corresponding interface of BACnetPropertyStatesLifeSafetyState
type BACnetPropertyStatesLifeSafetyState interface {
	BACnetPropertyStates
	// GetLifeSafetyState returns LifeSafetyState (property field)
	GetLifeSafetyState() BACnetLifeSafetyStateTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetPropertyStatesLifeSafetyState is the data-structure of this message
type _BACnetPropertyStatesLifeSafetyState struct {
	*_BACnetPropertyStates
	LifeSafetyState BACnetLifeSafetyStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLifeSafetyState) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLifeSafetyState) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLifeSafetyState) GetLifeSafetyState() BACnetLifeSafetyStateTagged {
	return m.LifeSafetyState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLifeSafetyState factory function for _BACnetPropertyStatesLifeSafetyState
func NewBACnetPropertyStatesLifeSafetyState(lifeSafetyState BACnetLifeSafetyStateTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLifeSafetyState {
	_result := &_BACnetPropertyStatesLifeSafetyState{
		LifeSafetyState:       lifeSafetyState,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLifeSafetyState(structType interface{}) BACnetPropertyStatesLifeSafetyState {
	if casted, ok := structType.(BACnetPropertyStatesLifeSafetyState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLifeSafetyState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLifeSafetyState) GetTypeName() string {
	return "BACnetPropertyStatesLifeSafetyState"
}

func (m *_BACnetPropertyStatesLifeSafetyState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesLifeSafetyState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lifeSafetyState)
	lengthInBits += m.LifeSafetyState.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesLifeSafetyState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLifeSafetyStateParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLifeSafetyState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLifeSafetyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLifeSafetyState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lifeSafetyState)
	if pullErr := readBuffer.PullContext("lifeSafetyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lifeSafetyState")
	}
	_lifeSafetyState, _lifeSafetyStateErr := BACnetLifeSafetyStateTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _lifeSafetyStateErr != nil {
		return nil, errors.Wrap(_lifeSafetyStateErr, "Error parsing 'lifeSafetyState' field")
	}
	lifeSafetyState := _lifeSafetyState.(BACnetLifeSafetyStateTagged)
	if closeErr := readBuffer.CloseContext("lifeSafetyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lifeSafetyState")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLifeSafetyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLifeSafetyState")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLifeSafetyState{
		LifeSafetyState:       lifeSafetyState,
		_BACnetPropertyStates: &_BACnetPropertyStates{},
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLifeSafetyState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLifeSafetyState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLifeSafetyState")
		}

		// Simple Field (lifeSafetyState)
		if pushErr := writeBuffer.PushContext("lifeSafetyState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lifeSafetyState")
		}
		_lifeSafetyStateErr := writeBuffer.WriteSerializable(m.GetLifeSafetyState())
		if popErr := writeBuffer.PopContext("lifeSafetyState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lifeSafetyState")
		}
		if _lifeSafetyStateErr != nil {
			return errors.Wrap(_lifeSafetyStateErr, "Error serializing 'lifeSafetyState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLifeSafetyState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLifeSafetyState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLifeSafetyState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
