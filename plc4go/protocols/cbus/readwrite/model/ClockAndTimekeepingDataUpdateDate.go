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

// ClockAndTimekeepingDataUpdateDate is the corresponding interface of ClockAndTimekeepingDataUpdateDate
type ClockAndTimekeepingDataUpdateDate interface {
	utils.LengthAware
	utils.Serializable
	ClockAndTimekeepingData
	// GetYearHigh returns YearHigh (property field)
	GetYearHigh() byte
	// GetYearLow returns YearLow (property field)
	GetYearLow() byte
	// GetMonth returns Month (property field)
	GetMonth() uint8
	// GetDay returns Day (property field)
	GetDay() uint8
	// GetDayOfWeek returns DayOfWeek (property field)
	GetDayOfWeek() uint8
}

// ClockAndTimekeepingDataUpdateDateExactly can be used when we want exactly this type and not a type which fulfills ClockAndTimekeepingDataUpdateDate.
// This is useful for switch cases.
type ClockAndTimekeepingDataUpdateDateExactly interface {
	ClockAndTimekeepingDataUpdateDate
	isClockAndTimekeepingDataUpdateDate() bool
}

// _ClockAndTimekeepingDataUpdateDate is the data-structure of this message
type _ClockAndTimekeepingDataUpdateDate struct {
	*_ClockAndTimekeepingData
	YearHigh  byte
	YearLow   byte
	Month     uint8
	Day       uint8
	DayOfWeek uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ClockAndTimekeepingDataUpdateDate) InitializeParent(parent ClockAndTimekeepingData, commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_ClockAndTimekeepingDataUpdateDate) GetParent() ClockAndTimekeepingData {
	return m._ClockAndTimekeepingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ClockAndTimekeepingDataUpdateDate) GetYearHigh() byte {
	return m.YearHigh
}

func (m *_ClockAndTimekeepingDataUpdateDate) GetYearLow() byte {
	return m.YearLow
}

func (m *_ClockAndTimekeepingDataUpdateDate) GetMonth() uint8 {
	return m.Month
}

func (m *_ClockAndTimekeepingDataUpdateDate) GetDay() uint8 {
	return m.Day
}

func (m *_ClockAndTimekeepingDataUpdateDate) GetDayOfWeek() uint8 {
	return m.DayOfWeek
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewClockAndTimekeepingDataUpdateDate factory function for _ClockAndTimekeepingDataUpdateDate
func NewClockAndTimekeepingDataUpdateDate(yearHigh byte, yearLow byte, month uint8, day uint8, dayOfWeek uint8, commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) *_ClockAndTimekeepingDataUpdateDate {
	_result := &_ClockAndTimekeepingDataUpdateDate{
		YearHigh:                 yearHigh,
		YearLow:                  yearLow,
		Month:                    month,
		Day:                      day,
		DayOfWeek:                dayOfWeek,
		_ClockAndTimekeepingData: NewClockAndTimekeepingData(commandTypeContainer, argument),
	}
	_result._ClockAndTimekeepingData._ClockAndTimekeepingDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastClockAndTimekeepingDataUpdateDate(structType interface{}) ClockAndTimekeepingDataUpdateDate {
	if casted, ok := structType.(ClockAndTimekeepingDataUpdateDate); ok {
		return casted
	}
	if casted, ok := structType.(*ClockAndTimekeepingDataUpdateDate); ok {
		return *casted
	}
	return nil
}

func (m *_ClockAndTimekeepingDataUpdateDate) GetTypeName() string {
	return "ClockAndTimekeepingDataUpdateDate"
}

func (m *_ClockAndTimekeepingDataUpdateDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ClockAndTimekeepingDataUpdateDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (yearHigh)
	lengthInBits += 8

	// Simple field (yearLow)
	lengthInBits += 8

	// Simple field (month)
	lengthInBits += 8

	// Simple field (day)
	lengthInBits += 8

	// Simple field (dayOfWeek)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ClockAndTimekeepingDataUpdateDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ClockAndTimekeepingDataUpdateDateParse(readBuffer utils.ReadBuffer) (ClockAndTimekeepingDataUpdateDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ClockAndTimekeepingDataUpdateDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClockAndTimekeepingDataUpdateDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (yearHigh)
	_yearHigh, _yearHighErr := readBuffer.ReadByte("yearHigh")
	if _yearHighErr != nil {
		return nil, errors.Wrap(_yearHighErr, "Error parsing 'yearHigh' field of ClockAndTimekeepingDataUpdateDate")
	}
	yearHigh := _yearHigh

	// Simple Field (yearLow)
	_yearLow, _yearLowErr := readBuffer.ReadByte("yearLow")
	if _yearLowErr != nil {
		return nil, errors.Wrap(_yearLowErr, "Error parsing 'yearLow' field of ClockAndTimekeepingDataUpdateDate")
	}
	yearLow := _yearLow

	// Simple Field (month)
	_month, _monthErr := readBuffer.ReadUint8("month", 8)
	if _monthErr != nil {
		return nil, errors.Wrap(_monthErr, "Error parsing 'month' field of ClockAndTimekeepingDataUpdateDate")
	}
	month := _month

	// Simple Field (day)
	_day, _dayErr := readBuffer.ReadUint8("day", 8)
	if _dayErr != nil {
		return nil, errors.Wrap(_dayErr, "Error parsing 'day' field of ClockAndTimekeepingDataUpdateDate")
	}
	day := _day

	// Simple Field (dayOfWeek)
	_dayOfWeek, _dayOfWeekErr := readBuffer.ReadUint8("dayOfWeek", 8)
	if _dayOfWeekErr != nil {
		return nil, errors.Wrap(_dayOfWeekErr, "Error parsing 'dayOfWeek' field of ClockAndTimekeepingDataUpdateDate")
	}
	dayOfWeek := _dayOfWeek

	if closeErr := readBuffer.CloseContext("ClockAndTimekeepingDataUpdateDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClockAndTimekeepingDataUpdateDate")
	}

	// Create a partially initialized instance
	_child := &_ClockAndTimekeepingDataUpdateDate{
		YearHigh:                 yearHigh,
		YearLow:                  yearLow,
		Month:                    month,
		Day:                      day,
		DayOfWeek:                dayOfWeek,
		_ClockAndTimekeepingData: &_ClockAndTimekeepingData{},
	}
	_child._ClockAndTimekeepingData._ClockAndTimekeepingDataChildRequirements = _child
	return _child, nil
}

func (m *_ClockAndTimekeepingDataUpdateDate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ClockAndTimekeepingDataUpdateDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ClockAndTimekeepingDataUpdateDate")
		}

		// Simple Field (yearHigh)
		yearHigh := byte(m.GetYearHigh())
		_yearHighErr := writeBuffer.WriteByte("yearHigh", (yearHigh))
		if _yearHighErr != nil {
			return errors.Wrap(_yearHighErr, "Error serializing 'yearHigh' field")
		}

		// Simple Field (yearLow)
		yearLow := byte(m.GetYearLow())
		_yearLowErr := writeBuffer.WriteByte("yearLow", (yearLow))
		if _yearLowErr != nil {
			return errors.Wrap(_yearLowErr, "Error serializing 'yearLow' field")
		}

		// Simple Field (month)
		month := uint8(m.GetMonth())
		_monthErr := writeBuffer.WriteUint8("month", 8, (month))
		if _monthErr != nil {
			return errors.Wrap(_monthErr, "Error serializing 'month' field")
		}

		// Simple Field (day)
		day := uint8(m.GetDay())
		_dayErr := writeBuffer.WriteUint8("day", 8, (day))
		if _dayErr != nil {
			return errors.Wrap(_dayErr, "Error serializing 'day' field")
		}

		// Simple Field (dayOfWeek)
		dayOfWeek := uint8(m.GetDayOfWeek())
		_dayOfWeekErr := writeBuffer.WriteUint8("dayOfWeek", 8, (dayOfWeek))
		if _dayOfWeekErr != nil {
			return errors.Wrap(_dayOfWeekErr, "Error serializing 'dayOfWeek' field")
		}

		if popErr := writeBuffer.PopContext("ClockAndTimekeepingDataUpdateDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ClockAndTimekeepingDataUpdateDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ClockAndTimekeepingDataUpdateDate) isClockAndTimekeepingDataUpdateDate() bool {
	return true
}

func (m *_ClockAndTimekeepingDataUpdateDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
