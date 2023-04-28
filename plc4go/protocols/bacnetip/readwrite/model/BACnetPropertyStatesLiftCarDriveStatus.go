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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesLiftCarDriveStatus is the corresponding interface of BACnetPropertyStatesLiftCarDriveStatus
type BACnetPropertyStatesLiftCarDriveStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLiftCarDriveStatus returns LiftCarDriveStatus (property field)
	GetLiftCarDriveStatus() BACnetLiftCarDriveStatusTagged
}

// BACnetPropertyStatesLiftCarDriveStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLiftCarDriveStatus.
// This is useful for switch cases.
type BACnetPropertyStatesLiftCarDriveStatusExactly interface {
	BACnetPropertyStatesLiftCarDriveStatus
	isBACnetPropertyStatesLiftCarDriveStatus() bool
}

// _BACnetPropertyStatesLiftCarDriveStatus is the data-structure of this message
type _BACnetPropertyStatesLiftCarDriveStatus struct {
	*_BACnetPropertyStates
	LiftCarDriveStatus BACnetLiftCarDriveStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLiftCarDriveStatus) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLiftCarDriveStatus) GetLiftCarDriveStatus() BACnetLiftCarDriveStatusTagged {
	return m.LiftCarDriveStatus
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLiftCarDriveStatus factory function for _BACnetPropertyStatesLiftCarDriveStatus
func NewBACnetPropertyStatesLiftCarDriveStatus(liftCarDriveStatus BACnetLiftCarDriveStatusTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLiftCarDriveStatus {
	_result := &_BACnetPropertyStatesLiftCarDriveStatus{
		LiftCarDriveStatus:    liftCarDriveStatus,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLiftCarDriveStatus(structType any) BACnetPropertyStatesLiftCarDriveStatus {
	if casted, ok := structType.(BACnetPropertyStatesLiftCarDriveStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftCarDriveStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) GetTypeName() string {
	return "BACnetPropertyStatesLiftCarDriveStatus"
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (liftCarDriveStatus)
	lengthInBits += m.LiftCarDriveStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesLiftCarDriveStatusParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesLiftCarDriveStatus, error) {
	return BACnetPropertyStatesLiftCarDriveStatusParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesLiftCarDriveStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLiftCarDriveStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftCarDriveStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLiftCarDriveStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (liftCarDriveStatus)
	if pullErr := readBuffer.PullContext("liftCarDriveStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for liftCarDriveStatus")
	}
	_liftCarDriveStatus, _liftCarDriveStatusErr := BACnetLiftCarDriveStatusTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _liftCarDriveStatusErr != nil {
		return nil, errors.Wrap(_liftCarDriveStatusErr, "Error parsing 'liftCarDriveStatus' field of BACnetPropertyStatesLiftCarDriveStatus")
	}
	liftCarDriveStatus := _liftCarDriveStatus.(BACnetLiftCarDriveStatusTagged)
	if closeErr := readBuffer.CloseContext("liftCarDriveStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for liftCarDriveStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftCarDriveStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLiftCarDriveStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLiftCarDriveStatus{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		LiftCarDriveStatus:    liftCarDriveStatus,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftCarDriveStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLiftCarDriveStatus")
		}

		// Simple Field (liftCarDriveStatus)
		if pushErr := writeBuffer.PushContext("liftCarDriveStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for liftCarDriveStatus")
		}
		_liftCarDriveStatusErr := writeBuffer.WriteSerializable(ctx, m.GetLiftCarDriveStatus())
		if popErr := writeBuffer.PopContext("liftCarDriveStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for liftCarDriveStatus")
		}
		if _liftCarDriveStatusErr != nil {
			return errors.Wrap(_liftCarDriveStatusErr, "Error serializing 'liftCarDriveStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftCarDriveStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLiftCarDriveStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) isBACnetPropertyStatesLiftCarDriveStatus() bool {
	return true
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
