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

// The data-structure of this message
type BACnetUnconfirmedServiceRequestUTCTimeSynchronization struct {
	*BACnetUnconfirmedServiceRequest

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestUTCTimeSynchronization interface {
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) ServiceChoice() uint8 {
	return 0x09
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetServiceChoice() uint8 {
	return 0x09
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUTCTimeSynchronization factory function for BACnetUnconfirmedServiceRequestUTCTimeSynchronization
func NewBACnetUnconfirmedServiceRequestUTCTimeSynchronization(len uint16) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestUTCTimeSynchronization{
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetUnconfirmedServiceRequest
}

func CastBACnetUnconfirmedServiceRequestUTCTimeSynchronization(structType interface{}) *BACnetUnconfirmedServiceRequestUTCTimeSynchronization {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestUTCTimeSynchronization {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestUTCTimeSynchronization); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestUTCTimeSynchronization); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestUTCTimeSynchronization(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestUTCTimeSynchronization(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUTCTimeSynchronization"
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUTCTimeSynchronizationParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUTCTimeSynchronization{
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child.BACnetUnconfirmedServiceRequest, nil
}

func (m *BACnetUnconfirmedServiceRequestUTCTimeSynchronization) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); popErr != nil {
			return popErr
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
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
