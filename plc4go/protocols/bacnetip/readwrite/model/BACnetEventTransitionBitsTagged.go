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

// BACnetEventTransitionBitsTagged is the data-structure of this message
type BACnetEventTransitionBitsTagged struct {
	Header  *BACnetTagHeader
	Payload *BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

// IBACnetEventTransitionBitsTagged is the corresponding interface of BACnetEventTransitionBitsTagged
type IBACnetEventTransitionBitsTagged interface {
	// GetHeader returns Header (property field)
	GetHeader() *BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadBitString
	// GetToOffnormal returns ToOffnormal (virtual field)
	GetToOffnormal() bool
	// GetToFault returns ToFault (virtual field)
	GetToFault() bool
	// GetToNormal returns ToNormal (virtual field)
	GetToNormal() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventTransitionBitsTagged) GetHeader() *BACnetTagHeader {
	return m.Header
}

func (m *BACnetEventTransitionBitsTagged) GetPayload() *BACnetTagPayloadBitString {
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

func (m *BACnetEventTransitionBitsTagged) GetToOffnormal() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (0))), func() interface{} { return bool(m.GetPayload().GetData()[0]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *BACnetEventTransitionBitsTagged) GetToFault() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (1))), func() interface{} { return bool(m.GetPayload().GetData()[1]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *BACnetEventTransitionBitsTagged) GetToNormal() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (2))), func() interface{} { return bool(m.GetPayload().GetData()[2]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventTransitionBitsTagged factory function for BACnetEventTransitionBitsTagged
func NewBACnetEventTransitionBitsTagged(header *BACnetTagHeader, payload *BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *BACnetEventTransitionBitsTagged {
	return &BACnetEventTransitionBitsTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

func CastBACnetEventTransitionBitsTagged(structType interface{}) *BACnetEventTransitionBitsTagged {
	if casted, ok := structType.(BACnetEventTransitionBitsTagged); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventTransitionBitsTagged); ok {
		return casted
	}
	return nil
}

func (m *BACnetEventTransitionBitsTagged) GetTypeName() string {
	return "BACnetEventTransitionBitsTagged"
}

func (m *BACnetEventTransitionBitsTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventTransitionBitsTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetEventTransitionBitsTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventTransitionBitsTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (*BACnetEventTransitionBitsTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventTransitionBitsTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventTransitionBitsTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
	}
	header := CastBACnetTagHeader(_header)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool(bool(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool(bool(bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadBitStringParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadBitString(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_toOffnormal := utils.InlineIf(bool(bool((len(payload.GetData())) > (0))), func() interface{} { return bool(payload.GetData()[0]) }, func() interface{} { return bool(bool(false)) }).(bool)
	toOffnormal := bool(_toOffnormal)
	_ = toOffnormal

	// Virtual field
	_toFault := utils.InlineIf(bool(bool((len(payload.GetData())) > (1))), func() interface{} { return bool(payload.GetData()[1]) }, func() interface{} { return bool(bool(false)) }).(bool)
	toFault := bool(_toFault)
	_ = toFault

	// Virtual field
	_toNormal := utils.InlineIf(bool(bool((len(payload.GetData())) > (2))), func() interface{} { return bool(payload.GetData()[2]) }, func() interface{} { return bool(bool(false)) }).(bool)
	toNormal := bool(_toNormal)
	_ = toNormal

	if closeErr := readBuffer.CloseContext("BACnetEventTransitionBitsTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventTransitionBitsTagged")
	}

	// Create the instance
	return NewBACnetEventTransitionBitsTagged(header, payload, tagNumber, tagClass), nil
}

func (m *BACnetEventTransitionBitsTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventTransitionBitsTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventTransitionBitsTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := m.Header.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
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
	if _toOffnormalErr := writeBuffer.WriteVirtual("toOffnormal", m.GetToOffnormal()); _toOffnormalErr != nil {
		return errors.Wrap(_toOffnormalErr, "Error serializing 'toOffnormal' field")
	}
	// Virtual field
	if _toFaultErr := writeBuffer.WriteVirtual("toFault", m.GetToFault()); _toFaultErr != nil {
		return errors.Wrap(_toFaultErr, "Error serializing 'toFault' field")
	}
	// Virtual field
	if _toNormalErr := writeBuffer.WriteVirtual("toNormal", m.GetToNormal()); _toNormalErr != nil {
		return errors.Wrap(_toNormalErr, "Error serializing 'toNormal' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventTransitionBitsTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventTransitionBitsTagged")
	}
	return nil
}

func (m *BACnetEventTransitionBitsTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
