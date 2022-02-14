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
type BACnetServiceAckVTOpen struct {
	*BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckVTOpen interface {
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
func (m *BACnetServiceAckVTOpen) ServiceChoice() uint8 {
	return 0x15
}

func (m *BACnetServiceAckVTOpen) GetServiceChoice() uint8 {
	return 0x15
}

func (m *BACnetServiceAckVTOpen) InitializeParent(parent *BACnetServiceAck) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetServiceAckVTOpen factory function for BACnetServiceAckVTOpen
func NewBACnetServiceAckVTOpen() *BACnetServiceAck {
	child := &BACnetServiceAckVTOpen{
		BACnetServiceAck: NewBACnetServiceAck(),
	}
	child.Child = child
	return child.BACnetServiceAck
}

func CastBACnetServiceAckVTOpen(structType interface{}) *BACnetServiceAckVTOpen {
	castFunc := func(typ interface{}) *BACnetServiceAckVTOpen {
		if casted, ok := typ.(BACnetServiceAckVTOpen); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckVTOpen); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAck); ok {
			return CastBACnetServiceAckVTOpen(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return CastBACnetServiceAckVTOpen(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckVTOpen) GetTypeName() string {
	return "BACnetServiceAckVTOpen"
}

func (m *BACnetServiceAckVTOpen) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckVTOpen) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetServiceAckVTOpen) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckVTOpenParse(readBuffer utils.ReadBuffer) (*BACnetServiceAck, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckVTOpen"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckVTOpen"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckVTOpen{
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child.BACnetServiceAck, nil
}

func (m *BACnetServiceAckVTOpen) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckVTOpen"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckVTOpen"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckVTOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
