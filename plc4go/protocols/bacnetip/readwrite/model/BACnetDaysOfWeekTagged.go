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

// BACnetDaysOfWeekTagged is the data-structure of this message
type BACnetDaysOfWeekTagged struct {
	Header  *BACnetTagHeader
	Payload *BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

// IBACnetDaysOfWeekTagged is the corresponding interface of BACnetDaysOfWeekTagged
type IBACnetDaysOfWeekTagged interface {
	// GetHeader returns Header (property field)
	GetHeader() *BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadBitString
	// GetMonday returns Monday (virtual field)
	GetMonday() bool
	// GetTuesday returns Tuesday (virtual field)
	GetTuesday() bool
	// GetWednesday returns Wednesday (virtual field)
	GetWednesday() bool
	// GetThursday returns Thursday (virtual field)
	GetThursday() bool
	// GetFriday returns Friday (virtual field)
	GetFriday() bool
	// GetSaturday returns Saturday (virtual field)
	GetSaturday() bool
	// GetSunday returns Sunday (virtual field)
	GetSunday() bool
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

func (m *BACnetDaysOfWeekTagged) GetHeader() *BACnetTagHeader {
	return m.Header
}

func (m *BACnetDaysOfWeekTagged) GetPayload() *BACnetTagPayloadBitString {
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

func (m *BACnetDaysOfWeekTagged) GetMonday() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (0))), func() interface{} { return bool(m.GetPayload().GetData()[0]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *BACnetDaysOfWeekTagged) GetTuesday() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (1))), func() interface{} { return bool(m.GetPayload().GetData()[1]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *BACnetDaysOfWeekTagged) GetWednesday() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (2))), func() interface{} { return bool(m.GetPayload().GetData()[2]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *BACnetDaysOfWeekTagged) GetThursday() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (3))), func() interface{} { return bool(m.GetPayload().GetData()[3]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *BACnetDaysOfWeekTagged) GetFriday() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (4))), func() interface{} { return bool(m.GetPayload().GetData()[4]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *BACnetDaysOfWeekTagged) GetSaturday() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (5))), func() interface{} { return bool(m.GetPayload().GetData()[5]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *BACnetDaysOfWeekTagged) GetSunday() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (6))), func() interface{} { return bool(m.GetPayload().GetData()[6]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDaysOfWeekTagged factory function for BACnetDaysOfWeekTagged
func NewBACnetDaysOfWeekTagged(header *BACnetTagHeader, payload *BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *BACnetDaysOfWeekTagged {
	return &BACnetDaysOfWeekTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

func CastBACnetDaysOfWeekTagged(structType interface{}) *BACnetDaysOfWeekTagged {
	if casted, ok := structType.(BACnetDaysOfWeekTagged); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetDaysOfWeekTagged); ok {
		return casted
	}
	return nil
}

func (m *BACnetDaysOfWeekTagged) GetTypeName() string {
	return "BACnetDaysOfWeekTagged"
}

func (m *BACnetDaysOfWeekTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetDaysOfWeekTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetDaysOfWeekTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDaysOfWeekTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (*BACnetDaysOfWeekTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDaysOfWeekTagged"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, pullErr
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
	}
	header := CastBACnetTagHeader(_header)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, closeErr
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, utils.ParseValidationError{"tag class doesn't match"}
	}

	// Validation
	if !(bool(bool(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool(bool(bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, utils.ParseAssertError{"tagnumber doesn't match"}
	}

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, pullErr
	}
	_payload, _payloadErr := BACnetTagPayloadBitStringParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadBitString(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_monday := utils.InlineIf(bool(bool((len(payload.GetData())) > (0))), func() interface{} { return bool(payload.GetData()[0]) }, func() interface{} { return bool(bool(false)) }).(bool)
	monday := bool(_monday)
	_ = monday

	// Virtual field
	_tuesday := utils.InlineIf(bool(bool((len(payload.GetData())) > (1))), func() interface{} { return bool(payload.GetData()[1]) }, func() interface{} { return bool(bool(false)) }).(bool)
	tuesday := bool(_tuesday)
	_ = tuesday

	// Virtual field
	_wednesday := utils.InlineIf(bool(bool((len(payload.GetData())) > (2))), func() interface{} { return bool(payload.GetData()[2]) }, func() interface{} { return bool(bool(false)) }).(bool)
	wednesday := bool(_wednesday)
	_ = wednesday

	// Virtual field
	_thursday := utils.InlineIf(bool(bool((len(payload.GetData())) > (3))), func() interface{} { return bool(payload.GetData()[3]) }, func() interface{} { return bool(bool(false)) }).(bool)
	thursday := bool(_thursday)
	_ = thursday

	// Virtual field
	_friday := utils.InlineIf(bool(bool((len(payload.GetData())) > (4))), func() interface{} { return bool(payload.GetData()[4]) }, func() interface{} { return bool(bool(false)) }).(bool)
	friday := bool(_friday)
	_ = friday

	// Virtual field
	_saturday := utils.InlineIf(bool(bool((len(payload.GetData())) > (5))), func() interface{} { return bool(payload.GetData()[5]) }, func() interface{} { return bool(bool(false)) }).(bool)
	saturday := bool(_saturday)
	_ = saturday

	// Virtual field
	_sunday := utils.InlineIf(bool(bool((len(payload.GetData())) > (6))), func() interface{} { return bool(payload.GetData()[6]) }, func() interface{} { return bool(bool(false)) }).(bool)
	sunday := bool(_sunday)
	_ = sunday

	if closeErr := readBuffer.CloseContext("BACnetDaysOfWeekTagged"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetDaysOfWeekTagged(header, payload, tagNumber, tagClass), nil
}

func (m *BACnetDaysOfWeekTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDaysOfWeekTagged"); pushErr != nil {
		return pushErr
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return pushErr
	}
	_headerErr := m.Header.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return popErr
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
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
	if _mondayErr := writeBuffer.WriteVirtual("monday", m.GetMonday()); _mondayErr != nil {
		return errors.Wrap(_mondayErr, "Error serializing 'monday' field")
	}
	// Virtual field
	if _tuesdayErr := writeBuffer.WriteVirtual("tuesday", m.GetTuesday()); _tuesdayErr != nil {
		return errors.Wrap(_tuesdayErr, "Error serializing 'tuesday' field")
	}
	// Virtual field
	if _wednesdayErr := writeBuffer.WriteVirtual("wednesday", m.GetWednesday()); _wednesdayErr != nil {
		return errors.Wrap(_wednesdayErr, "Error serializing 'wednesday' field")
	}
	// Virtual field
	if _thursdayErr := writeBuffer.WriteVirtual("thursday", m.GetThursday()); _thursdayErr != nil {
		return errors.Wrap(_thursdayErr, "Error serializing 'thursday' field")
	}
	// Virtual field
	if _fridayErr := writeBuffer.WriteVirtual("friday", m.GetFriday()); _fridayErr != nil {
		return errors.Wrap(_fridayErr, "Error serializing 'friday' field")
	}
	// Virtual field
	if _saturdayErr := writeBuffer.WriteVirtual("saturday", m.GetSaturday()); _saturdayErr != nil {
		return errors.Wrap(_saturdayErr, "Error serializing 'saturday' field")
	}
	// Virtual field
	if _sundayErr := writeBuffer.WriteVirtual("sunday", m.GetSunday()); _sundayErr != nil {
		return errors.Wrap(_sundayErr, "Error serializing 'sunday' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDaysOfWeekTagged"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetDaysOfWeekTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
