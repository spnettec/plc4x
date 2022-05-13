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

// BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple is the data-structure of this message
type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple struct {
	*BACnetUnconfirmedServiceRequest

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
type IBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple interface {
	IBACnetUnconfirmedServiceRequest
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

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

// NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple factory function for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
func NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(serviceRequestLength uint16) *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
	_result := &BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(structType interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, utils.ParseAssertError{"TODO: implement me"}
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
