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

// BACnetUnconfirmedServiceRequestUTCTimeSynchronization is the data-structure of this message
type BACnetUnconfirmedServiceRequestUTCTimeSynchronization struct {
	*BACnetUnconfirmedServiceRequest
	SynchronizedDate *BACnetApplicationTagDate
	SynchronizedTime *BACnetApplicationTagTime

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetUnconfirmedServiceRequestUTCTimeSynchronization is the corresponding interface of BACnetUnconfirmedServiceRequestUTCTimeSynchronization
type IBACnetUnconfirmedServiceRequestUTCTimeSynchronization interface {
	IBACnetUnconfirmedServiceRequest
	// GetSynchronizedDate returns SynchronizedDate (property field)
	GetSynchronizedDate() *BACnetApplicationTagDate
	// GetSynchronizedTime returns SynchronizedTime (property field)
	GetSynchronizedTime() *BACnetApplicationTagTime
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

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UTC_TIME_SYNCHRONIZATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetSynchronizedDate() *BACnetApplicationTagDate {
	return m.SynchronizedDate
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetSynchronizedTime() *BACnetApplicationTagTime {
	return m.SynchronizedTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUTCTimeSynchronization factory function for BACnetUnconfirmedServiceRequestUTCTimeSynchronization
func NewBACnetUnconfirmedServiceRequestUTCTimeSynchronization(synchronizedDate *BACnetApplicationTagDate, synchronizedTime *BACnetApplicationTagTime, serviceRequestLength uint16) *BACnetUnconfirmedServiceRequestUTCTimeSynchronization {
	_result := &BACnetUnconfirmedServiceRequestUTCTimeSynchronization{
		SynchronizedDate:                synchronizedDate,
		SynchronizedTime:                synchronizedTime,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetUnconfirmedServiceRequestUTCTimeSynchronization(structType interface{}) *BACnetUnconfirmedServiceRequestUTCTimeSynchronization {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUTCTimeSynchronization); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUTCTimeSynchronization); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUTCTimeSynchronization(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUTCTimeSynchronization(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUTCTimeSynchronization"
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (synchronizedDate)
	lengthInBits += m.SynchronizedDate.GetLengthInBits()

	// Simple field (synchronizedTime)
	lengthInBits += m.SynchronizedTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUTCTimeSynchronizationParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetUnconfirmedServiceRequestUTCTimeSynchronization, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (synchronizedDate)
	if pullErr := readBuffer.PullContext("synchronizedDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for synchronizedDate")
	}
	_synchronizedDate, _synchronizedDateErr := BACnetApplicationTagParse(readBuffer)
	if _synchronizedDateErr != nil {
		return nil, errors.Wrap(_synchronizedDateErr, "Error parsing 'synchronizedDate' field")
	}
	synchronizedDate := CastBACnetApplicationTagDate(_synchronizedDate)
	if closeErr := readBuffer.CloseContext("synchronizedDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for synchronizedDate")
	}

	// Simple Field (synchronizedTime)
	if pullErr := readBuffer.PullContext("synchronizedTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for synchronizedTime")
	}
	_synchronizedTime, _synchronizedTimeErr := BACnetApplicationTagParse(readBuffer)
	if _synchronizedTimeErr != nil {
		return nil, errors.Wrap(_synchronizedTimeErr, "Error parsing 'synchronizedTime' field")
	}
	synchronizedTime := CastBACnetApplicationTagTime(_synchronizedTime)
	if closeErr := readBuffer.CloseContext("synchronizedTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for synchronizedTime")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUTCTimeSynchronization{
		SynchronizedDate:                CastBACnetApplicationTagDate(synchronizedDate),
		SynchronizedTime:                CastBACnetApplicationTagTime(synchronizedTime),
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
		}

		// Simple Field (synchronizedDate)
		if pushErr := writeBuffer.PushContext("synchronizedDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for synchronizedDate")
		}
		_synchronizedDateErr := m.SynchronizedDate.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("synchronizedDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for synchronizedDate")
		}
		if _synchronizedDateErr != nil {
			return errors.Wrap(_synchronizedDateErr, "Error serializing 'synchronizedDate' field")
		}

		// Simple Field (synchronizedTime)
		if pushErr := writeBuffer.PushContext("synchronizedTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for synchronizedTime")
		}
		_synchronizedTimeErr := m.SynchronizedTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("synchronizedTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for synchronizedTime")
		}
		if _synchronizedTimeErr != nil {
			return errors.Wrap(_synchronizedTimeErr, "Error serializing 'synchronizedTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
