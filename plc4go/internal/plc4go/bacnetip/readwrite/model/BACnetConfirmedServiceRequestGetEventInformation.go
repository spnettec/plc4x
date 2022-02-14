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
type BACnetConfirmedServiceRequestGetEventInformation struct {
	*BACnetConfirmedServiceRequest

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetConfirmedServiceRequestGetEventInformation interface {
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
func (m *BACnetConfirmedServiceRequestGetEventInformation) ServiceChoice() uint8 {
	return 0x1D
}

func (m *BACnetConfirmedServiceRequestGetEventInformation) GetServiceChoice() uint8 {
	return 0x1D
}

func (m *BACnetConfirmedServiceRequestGetEventInformation) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestGetEventInformation factory function for BACnetConfirmedServiceRequestGetEventInformation
func NewBACnetConfirmedServiceRequestGetEventInformation(len uint16) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestGetEventInformation{
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetConfirmedServiceRequest
}

func CastBACnetConfirmedServiceRequestGetEventInformation(structType interface{}) *BACnetConfirmedServiceRequestGetEventInformation {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestGetEventInformation {
		if casted, ok := typ.(BACnetConfirmedServiceRequestGetEventInformation); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestGetEventInformation); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestGetEventInformation(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestGetEventInformation(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestGetEventInformation) GetTypeName() string {
	return "BACnetConfirmedServiceRequestGetEventInformation"
}

func (m *BACnetConfirmedServiceRequestGetEventInformation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestGetEventInformation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestGetEventInformation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestGetEventInformationParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestGetEventInformation"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestGetEventInformation"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestGetEventInformation{
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child.BACnetConfirmedServiceRequest, nil
}

func (m *BACnetConfirmedServiceRequestGetEventInformation) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestGetEventInformation"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestGetEventInformation"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestGetEventInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
