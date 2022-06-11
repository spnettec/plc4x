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

// BACnetApplicationTagEnumerated is the data-structure of this message
type BACnetApplicationTagEnumerated struct {
	*BACnetApplicationTag
	Payload *BACnetTagPayloadEnumerated
}

// IBACnetApplicationTagEnumerated is the corresponding interface of BACnetApplicationTagEnumerated
type IBACnetApplicationTagEnumerated interface {
	IBACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadEnumerated
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint32
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

func (m *BACnetApplicationTagEnumerated) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader) {
	m.BACnetApplicationTag.Header = header
}

func (m *BACnetApplicationTagEnumerated) GetParent() *BACnetApplicationTag {
	return m.BACnetApplicationTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetApplicationTagEnumerated) GetPayload() *BACnetTagPayloadEnumerated {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetApplicationTagEnumerated) GetActualValue() uint32 {
	return uint32(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagEnumerated factory function for BACnetApplicationTagEnumerated
func NewBACnetApplicationTagEnumerated(payload *BACnetTagPayloadEnumerated, header *BACnetTagHeader) *BACnetApplicationTagEnumerated {
	_result := &BACnetApplicationTagEnumerated{
		Payload:              payload,
		BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	_result.Child = _result
	return _result
}

func CastBACnetApplicationTagEnumerated(structType interface{}) *BACnetApplicationTagEnumerated {
	if casted, ok := structType.(BACnetApplicationTagEnumerated); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetApplicationTagEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(BACnetApplicationTag); ok {
		return CastBACnetApplicationTagEnumerated(casted.Child)
	}
	if casted, ok := structType.(*BACnetApplicationTag); ok {
		return CastBACnetApplicationTagEnumerated(casted.Child)
	}
	return nil
}

func (m *BACnetApplicationTagEnumerated) GetTypeName() string {
	return "BACnetApplicationTagEnumerated"
}

func (m *BACnetApplicationTagEnumerated) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetApplicationTagEnumerated) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetApplicationTagEnumerated) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetApplicationTagEnumeratedParse(readBuffer utils.ReadBuffer, header *BACnetTagHeader) (*BACnetApplicationTagEnumerated, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadEnumeratedParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadEnumerated(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_actualValue := payload.GetActualValue()
	actualValue := uint32(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagEnumerated")
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagEnumerated{
		Payload:              CastBACnetTagPayloadEnumerated(payload),
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child, nil
}

func (m *BACnetApplicationTagEnumerated) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagEnumerated")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := m.Payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagEnumerated")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
