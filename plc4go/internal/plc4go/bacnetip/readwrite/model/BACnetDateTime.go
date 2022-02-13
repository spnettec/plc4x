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
type BACnetDateTime struct {
	OpeningTag *BACnetOpeningTag
	DateValue  *BACnetApplicationTagDate
	TimeValue  *BACnetApplicationTagTime
	ClosingTag *BACnetClosingTag
}

// The corresponding interface
type IBACnetDateTime interface {
	// GetOpeningTag returns OpeningTag
	GetOpeningTag() *BACnetOpeningTag
	// GetDateValue returns DateValue
	GetDateValue() *BACnetApplicationTagDate
	// GetTimeValue returns TimeValue
	GetTimeValue() *BACnetApplicationTagTime
	// GetClosingTag returns ClosingTag
	GetClosingTag() *BACnetClosingTag
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetDateTime) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetDateTime) GetDateValue() *BACnetApplicationTagDate {
	return m.DateValue
}

func (m *BACnetDateTime) GetTimeValue() *BACnetApplicationTagTime {
	return m.TimeValue
}

func (m *BACnetDateTime) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewBACnetDateTime(openingTag *BACnetOpeningTag, dateValue *BACnetApplicationTagDate, timeValue *BACnetApplicationTagTime, closingTag *BACnetClosingTag) *BACnetDateTime {
	return &BACnetDateTime{OpeningTag: openingTag, DateValue: dateValue, TimeValue: timeValue, ClosingTag: closingTag}
}

func CastBACnetDateTime(structType interface{}) *BACnetDateTime {
	castFunc := func(typ interface{}) *BACnetDateTime {
		if casted, ok := typ.(BACnetDateTime); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetDateTime); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetDateTime) GetTypeName() string {
	return "BACnetDateTime"
}

func (m *BACnetDateTime) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetDateTime) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.LengthInBits()

	// Simple field (dateValue)
	lengthInBits += m.DateValue.LengthInBits()

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.LengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.LengthInBits()

	return lengthInBits
}

func (m *BACnetDateTime) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetDateTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetDateTime, error) {
	if pullErr := readBuffer.PullContext("BACnetDateTime"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType_OPENING_TAG)
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (dateValue)
	if pullErr := readBuffer.PullContext("dateValue"); pullErr != nil {
		return nil, pullErr
	}
	_dateValue, _dateValueErr := BACnetApplicationTagParse(readBuffer)
	if _dateValueErr != nil {
		return nil, errors.Wrap(_dateValueErr, "Error parsing 'dateValue' field")
	}
	dateValue := CastBACnetApplicationTagDate(_dateValue)
	if closeErr := readBuffer.CloseContext("dateValue"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timeValue)
	if pullErr := readBuffer.PullContext("timeValue"); pullErr != nil {
		return nil, pullErr
	}
	_timeValue, _timeValueErr := BACnetApplicationTagParse(readBuffer)
	if _timeValueErr != nil {
		return nil, errors.Wrap(_timeValueErr, "Error parsing 'timeValue' field")
	}
	timeValue := CastBACnetApplicationTagTime(_timeValue)
	if closeErr := readBuffer.CloseContext("timeValue"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType_CLOSING_TAG)
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetDateTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetDateTime(openingTag, dateValue, timeValue, closingTag), nil
}

func (m *BACnetDateTime) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetDateTime"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (dateValue)
	if pushErr := writeBuffer.PushContext("dateValue"); pushErr != nil {
		return pushErr
	}
	_dateValueErr := m.DateValue.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("dateValue"); popErr != nil {
		return popErr
	}
	if _dateValueErr != nil {
		return errors.Wrap(_dateValueErr, "Error serializing 'dateValue' field")
	}

	// Simple Field (timeValue)
	if pushErr := writeBuffer.PushContext("timeValue"); pushErr != nil {
		return pushErr
	}
	_timeValueErr := m.TimeValue.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("timeValue"); popErr != nil {
		return popErr
	}
	if _timeValueErr != nil {
		return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateTime"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetDateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
