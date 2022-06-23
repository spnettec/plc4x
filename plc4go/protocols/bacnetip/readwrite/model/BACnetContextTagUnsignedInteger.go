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

// BACnetContextTagUnsignedInteger is the corresponding interface of BACnetContextTagUnsignedInteger
type BACnetContextTagUnsignedInteger interface {
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
}

// BACnetContextTagUnsignedIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagUnsignedInteger.
// This is useful for switch cases.
type BACnetContextTagUnsignedIntegerExactly interface {
	BACnetContextTagUnsignedInteger
	isBACnetContextTagUnsignedInteger() bool
}

// _BACnetContextTagUnsignedInteger is the data-structure of this message
type _BACnetContextTagUnsignedInteger struct {
	*_BACnetContextTag
	Payload BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagUnsignedInteger) GetDataType() BACnetDataType {
	return BACnetDataType_UNSIGNED_INTEGER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagUnsignedInteger) InitializeParent(parent BACnetContextTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetContextTagUnsignedInteger) GetParent() BACnetContextTag {
	return m._BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagUnsignedInteger) GetPayload() BACnetTagPayloadUnsignedInteger {
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

func (m *_BACnetContextTagUnsignedInteger) GetActualValue() uint64 {
	return uint64(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagUnsignedInteger factory function for _BACnetContextTagUnsignedInteger
func NewBACnetContextTagUnsignedInteger(payload BACnetTagPayloadUnsignedInteger, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagUnsignedInteger {
	_result := &_BACnetContextTagUnsignedInteger{
		Payload:           payload,
		_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagUnsignedInteger(structType interface{}) BACnetContextTagUnsignedInteger {
	if casted, ok := structType.(BACnetContextTagUnsignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagUnsignedInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagUnsignedInteger) GetTypeName() string {
	return "BACnetContextTagUnsignedInteger"
}

func (m *_BACnetContextTagUnsignedInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetContextTagUnsignedInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagUnsignedInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagUnsignedIntegerParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, header BACnetTagHeader) (BACnetContextTagUnsignedInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagUnsignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagUnsignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadUnsignedIntegerParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := _payload.(BACnetTagPayloadUnsignedInteger)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_actualValue := payload.GetActualValue()
	actualValue := uint64(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagUnsignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagUnsignedInteger")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagUnsignedInteger{
		Payload: payload,
		_BACnetContextTag: &_BACnetContextTag{
			TagNumberArgument: tagNumberArgument,
		},
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagUnsignedInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagUnsignedInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagUnsignedInteger")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
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

		if popErr := writeBuffer.PopContext("BACnetContextTagUnsignedInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagUnsignedInteger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetContextTagUnsignedInteger) isBACnetContextTagUnsignedInteger() bool {
	return true
}

func (m *_BACnetContextTagUnsignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
