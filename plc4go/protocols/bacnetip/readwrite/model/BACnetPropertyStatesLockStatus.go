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

// BACnetPropertyStatesLockStatus is the data-structure of this message
type BACnetPropertyStatesLockStatus struct {
	*BACnetPropertyStates
	LockStatus *BACnetLockStatusTagged
}

// IBACnetPropertyStatesLockStatus is the corresponding interface of BACnetPropertyStatesLockStatus
type IBACnetPropertyStatesLockStatus interface {
	IBACnetPropertyStates
	// GetLockStatus returns LockStatus (property field)
	GetLockStatus() *BACnetLockStatusTagged
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

func (m *BACnetPropertyStatesLockStatus) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesLockStatus) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesLockStatus) GetLockStatus() *BACnetLockStatusTagged {
	return m.LockStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLockStatus factory function for BACnetPropertyStatesLockStatus
func NewBACnetPropertyStatesLockStatus(lockStatus *BACnetLockStatusTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesLockStatus {
	_result := &BACnetPropertyStatesLockStatus{
		LockStatus:           lockStatus,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesLockStatus(structType interface{}) *BACnetPropertyStatesLockStatus {
	if casted, ok := structType.(BACnetPropertyStatesLockStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLockStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesLockStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesLockStatus(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesLockStatus) GetTypeName() string {
	return "BACnetPropertyStatesLockStatus"
}

func (m *BACnetPropertyStatesLockStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesLockStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lockStatus)
	lengthInBits += m.LockStatus.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesLockStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLockStatusParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesLockStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLockStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lockStatus)
	if pullErr := readBuffer.PullContext("lockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lockStatus")
	}
	_lockStatus, _lockStatusErr := BACnetLockStatusTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _lockStatusErr != nil {
		return nil, errors.Wrap(_lockStatusErr, "Error parsing 'lockStatus' field")
	}
	lockStatus := CastBACnetLockStatusTagged(_lockStatus)
	if closeErr := readBuffer.CloseContext("lockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lockStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLockStatus")
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesLockStatus{
		LockStatus:           CastBACnetLockStatusTagged(lockStatus),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesLockStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLockStatus")
		}

		// Simple Field (lockStatus)
		if pushErr := writeBuffer.PushContext("lockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lockStatus")
		}
		_lockStatusErr := writeBuffer.WriteSerializable(m.LockStatus)
		if popErr := writeBuffer.PopContext("lockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lockStatus")
		}
		if _lockStatusErr != nil {
			return errors.Wrap(_lockStatusErr, "Error serializing 'lockStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLockStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesLockStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
