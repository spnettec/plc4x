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

// BACnetPropertyStatesRestartReason is the corresponding interface of BACnetPropertyStatesRestartReason
type BACnetPropertyStatesRestartReason interface {
	BACnetPropertyStates
	// GetRestartReason returns RestartReason (property field)
	GetRestartReason() BACnetRestartReasonTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetPropertyStatesRestartReason is the data-structure of this message
type _BACnetPropertyStatesRestartReason struct {
	*_BACnetPropertyStates
	RestartReason BACnetRestartReasonTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesRestartReason) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesRestartReason) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesRestartReason) GetRestartReason() BACnetRestartReasonTagged {
	return m.RestartReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesRestartReason factory function for _BACnetPropertyStatesRestartReason
func NewBACnetPropertyStatesRestartReason(restartReason BACnetRestartReasonTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesRestartReason {
	_result := &_BACnetPropertyStatesRestartReason{
		RestartReason:         restartReason,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesRestartReason(structType interface{}) BACnetPropertyStatesRestartReason {
	if casted, ok := structType.(BACnetPropertyStatesRestartReason); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesRestartReason); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesRestartReason) GetTypeName() string {
	return "BACnetPropertyStatesRestartReason"
}

func (m *_BACnetPropertyStatesRestartReason) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesRestartReason) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (restartReason)
	lengthInBits += m.RestartReason.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesRestartReason) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesRestartReasonParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesRestartReason, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesRestartReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesRestartReason")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (restartReason)
	if pullErr := readBuffer.PullContext("restartReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for restartReason")
	}
	_restartReason, _restartReasonErr := BACnetRestartReasonTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _restartReasonErr != nil {
		return nil, errors.Wrap(_restartReasonErr, "Error parsing 'restartReason' field")
	}
	restartReason := _restartReason.(BACnetRestartReasonTagged)
	if closeErr := readBuffer.CloseContext("restartReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for restartReason")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesRestartReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesRestartReason")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesRestartReason{
		RestartReason:         restartReason,
		_BACnetPropertyStates: &_BACnetPropertyStates{},
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesRestartReason) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesRestartReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesRestartReason")
		}

		// Simple Field (restartReason)
		if pushErr := writeBuffer.PushContext("restartReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for restartReason")
		}
		_restartReasonErr := writeBuffer.WriteSerializable(m.GetRestartReason())
		if popErr := writeBuffer.PopContext("restartReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for restartReason")
		}
		if _restartReasonErr != nil {
			return errors.Wrap(_restartReasonErr, "Error serializing 'restartReason' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesRestartReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesRestartReason")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesRestartReason) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
