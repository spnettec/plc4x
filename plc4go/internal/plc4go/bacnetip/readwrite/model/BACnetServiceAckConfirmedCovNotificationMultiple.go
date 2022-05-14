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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetServiceAckConfirmedCovNotificationMultiple is the data-structure of this message
type BACnetServiceAckConfirmedCovNotificationMultiple struct {
	*BACnetServiceAck

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetServiceAckConfirmedCovNotificationMultiple is the corresponding interface of BACnetServiceAckConfirmedCovNotificationMultiple
type IBACnetServiceAckConfirmedCovNotificationMultiple interface {
	IBACnetServiceAck
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

func (m *BACnetServiceAckConfirmedCovNotificationMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetServiceAckConfirmedCovNotificationMultiple) InitializeParent(parent *BACnetServiceAck) {
}

func (m *BACnetServiceAckConfirmedCovNotificationMultiple) GetParent() *BACnetServiceAck {
	return m.BACnetServiceAck
}

// NewBACnetServiceAckConfirmedCovNotificationMultiple factory function for BACnetServiceAckConfirmedCovNotificationMultiple
func NewBACnetServiceAckConfirmedCovNotificationMultiple(serviceRequestLength uint16) *BACnetServiceAckConfirmedCovNotificationMultiple {
	_result := &BACnetServiceAckConfirmedCovNotificationMultiple{
		BACnetServiceAck: NewBACnetServiceAck(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetServiceAckConfirmedCovNotificationMultiple(structType interface{}) *BACnetServiceAckConfirmedCovNotificationMultiple {
	if casted, ok := structType.(BACnetServiceAckConfirmedCovNotificationMultiple); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckConfirmedCovNotificationMultiple); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAck); ok {
		return CastBACnetServiceAckConfirmedCovNotificationMultiple(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return CastBACnetServiceAckConfirmedCovNotificationMultiple(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckConfirmedCovNotificationMultiple) GetTypeName() string {
	return "BACnetServiceAckConfirmedCovNotificationMultiple"
}

func (m *BACnetServiceAckConfirmedCovNotificationMultiple) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckConfirmedCovNotificationMultiple) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetServiceAckConfirmedCovNotificationMultiple) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckConfirmedCovNotificationMultipleParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetServiceAckConfirmedCovNotificationMultiple, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckConfirmedCovNotificationMultiple"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, utils.ParseValidationError{"TODO: implement me"}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckConfirmedCovNotificationMultiple"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckConfirmedCovNotificationMultiple{
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child, nil
}

func (m *BACnetServiceAckConfirmedCovNotificationMultiple) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckConfirmedCovNotificationMultiple"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckConfirmedCovNotificationMultiple"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckConfirmedCovNotificationMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
