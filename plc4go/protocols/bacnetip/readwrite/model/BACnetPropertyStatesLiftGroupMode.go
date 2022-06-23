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

// BACnetPropertyStatesLiftGroupMode is the corresponding interface of BACnetPropertyStatesLiftGroupMode
type BACnetPropertyStatesLiftGroupMode interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLiftGroupMode returns LiftGroupMode (property field)
	GetLiftGroupMode() BACnetLiftGroupModeTagged
}

// BACnetPropertyStatesLiftGroupModeExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLiftGroupMode.
// This is useful for switch cases.
type BACnetPropertyStatesLiftGroupModeExactly interface {
	BACnetPropertyStatesLiftGroupMode
	isBACnetPropertyStatesLiftGroupMode() bool
}

// _BACnetPropertyStatesLiftGroupMode is the data-structure of this message
type _BACnetPropertyStatesLiftGroupMode struct {
	*_BACnetPropertyStates
	LiftGroupMode BACnetLiftGroupModeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLiftGroupMode) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLiftGroupMode) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLiftGroupMode) GetLiftGroupMode() BACnetLiftGroupModeTagged {
	return m.LiftGroupMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLiftGroupMode factory function for _BACnetPropertyStatesLiftGroupMode
func NewBACnetPropertyStatesLiftGroupMode(liftGroupMode BACnetLiftGroupModeTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLiftGroupMode {
	_result := &_BACnetPropertyStatesLiftGroupMode{
		LiftGroupMode:         liftGroupMode,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLiftGroupMode(structType interface{}) BACnetPropertyStatesLiftGroupMode {
	if casted, ok := structType.(BACnetPropertyStatesLiftGroupMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftGroupMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLiftGroupMode) GetTypeName() string {
	return "BACnetPropertyStatesLiftGroupMode"
}

func (m *_BACnetPropertyStatesLiftGroupMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesLiftGroupMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (liftGroupMode)
	lengthInBits += m.LiftGroupMode.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesLiftGroupMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLiftGroupModeParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLiftGroupMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftGroupMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLiftGroupMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (liftGroupMode)
	if pullErr := readBuffer.PullContext("liftGroupMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for liftGroupMode")
	}
	_liftGroupMode, _liftGroupModeErr := BACnetLiftGroupModeTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _liftGroupModeErr != nil {
		return nil, errors.Wrap(_liftGroupModeErr, "Error parsing 'liftGroupMode' field")
	}
	liftGroupMode := _liftGroupMode.(BACnetLiftGroupModeTagged)
	if closeErr := readBuffer.CloseContext("liftGroupMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for liftGroupMode")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftGroupMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLiftGroupMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLiftGroupMode{
		LiftGroupMode:         liftGroupMode,
		_BACnetPropertyStates: &_BACnetPropertyStates{},
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLiftGroupMode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftGroupMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLiftGroupMode")
		}

		// Simple Field (liftGroupMode)
		if pushErr := writeBuffer.PushContext("liftGroupMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for liftGroupMode")
		}
		_liftGroupModeErr := writeBuffer.WriteSerializable(m.GetLiftGroupMode())
		if popErr := writeBuffer.PopContext("liftGroupMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for liftGroupMode")
		}
		if _liftGroupModeErr != nil {
			return errors.Wrap(_liftGroupModeErr, "Error serializing 'liftGroupMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftGroupMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLiftGroupMode")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLiftGroupMode) isBACnetPropertyStatesLiftGroupMode() bool {
	return true
}

func (m *_BACnetPropertyStatesLiftGroupMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
