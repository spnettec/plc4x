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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetApplicationTagSignedInteger struct {
	*BACnetApplicationTag
	Payload *BACnetTagPayloadSignedInteger
}

// The corresponding interface
type IBACnetApplicationTagSignedInteger interface {
	IBACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
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

func (m *BACnetApplicationTagSignedInteger) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader) {
	m.BACnetApplicationTag.Header = header
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetApplicationTagSignedInteger) GetPayload() *BACnetTagPayloadSignedInteger {
	return m.Payload
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetApplicationTagSignedInteger) GetActualValue() uint64 {
	return m.GetPayload().GetActualValue()
}

// NewBACnetApplicationTagSignedInteger factory function for BACnetApplicationTagSignedInteger
func NewBACnetApplicationTagSignedInteger(payload *BACnetTagPayloadSignedInteger, header *BACnetTagHeader) *BACnetApplicationTag {
	child := &BACnetApplicationTagSignedInteger{
		Payload:              payload,
		BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	child.Child = child
	return child.BACnetApplicationTag
}

func CastBACnetApplicationTagSignedInteger(structType interface{}) *BACnetApplicationTagSignedInteger {
	if casted, ok := structType.(BACnetApplicationTagSignedInteger); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetApplicationTagSignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(BACnetApplicationTag); ok {
		return CastBACnetApplicationTagSignedInteger(casted.Child)
	}
	if casted, ok := structType.(*BACnetApplicationTag); ok {
		return CastBACnetApplicationTagSignedInteger(casted.Child)
	}
	return nil
}

func (m *BACnetApplicationTagSignedInteger) GetTypeName() string {
	return "BACnetApplicationTagSignedInteger"
}

func (m *BACnetApplicationTagSignedInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetApplicationTagSignedInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetApplicationTagSignedInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetApplicationTagSignedIntegerParse(readBuffer utils.ReadBuffer, header *BACnetTagHeader) (*BACnetApplicationTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTagSignedInteger"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, pullErr
	}
	_payload, _payloadErr := BACnetTagPayloadSignedIntegerParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadSignedInteger(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_actualValue := payload.GetActualValue()
	actualValue := uint64(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagSignedInteger"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagSignedInteger{
		Payload:              CastBACnetTagPayloadSignedInteger(payload),
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child.BACnetApplicationTag, nil
}

func (m *BACnetApplicationTagSignedInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagSignedInteger"); pushErr != nil {
			return pushErr
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return pushErr
		}
		_payloadErr := m.Payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return popErr
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagSignedInteger"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagSignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
