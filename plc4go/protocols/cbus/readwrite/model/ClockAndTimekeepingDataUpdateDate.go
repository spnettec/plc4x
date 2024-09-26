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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ClockAndTimekeepingDataUpdateDate is the corresponding interface of ClockAndTimekeepingDataUpdateDate
type ClockAndTimekeepingDataUpdateDate interface {
	fmt.Stringer
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
	// IsClockAndTimekeepingDataUpdateDate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsClockAndTimekeepingDataUpdateDate()
}

// _ClockAndTimekeepingDataUpdateDate is the data-structure of this message
type _ClockAndTimekeepingDataUpdateDate struct {
	ClockAndTimekeepingDataContract
	YearHigh  byte
	YearLow   byte
	Month     uint8
	Day       uint8
	DayOfWeek uint8
}

var _ ClockAndTimekeepingDataUpdateDate = (*_ClockAndTimekeepingDataUpdateDate)(nil)
var _ ClockAndTimekeepingDataRequirements = (*_ClockAndTimekeepingDataUpdateDate)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ClockAndTimekeepingDataUpdateDate) GetParent() ClockAndTimekeepingDataContract {
	return m.ClockAndTimekeepingDataContract
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
func NewClockAndTimekeepingDataUpdateDate(commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte, yearHigh byte, yearLow byte, month uint8, day uint8, dayOfWeek uint8) *_ClockAndTimekeepingDataUpdateDate {
	_result := &_ClockAndTimekeepingDataUpdateDate{
		ClockAndTimekeepingDataContract: NewClockAndTimekeepingData(commandTypeContainer, argument),
		YearHigh:                        yearHigh,
		YearLow:                         yearLow,
		Month:                           month,
		Day:                             day,
		DayOfWeek:                       dayOfWeek,
	}
	_result.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastClockAndTimekeepingDataUpdateDate(structType any) ClockAndTimekeepingDataUpdateDate {
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

func (m *_ClockAndTimekeepingDataUpdateDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData).getLengthInBits(ctx))

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

func (m *_ClockAndTimekeepingDataUpdateDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ClockAndTimekeepingDataUpdateDate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ClockAndTimekeepingData) (__clockAndTimekeepingDataUpdateDate ClockAndTimekeepingDataUpdateDate, err error) {
	m.ClockAndTimekeepingDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ClockAndTimekeepingDataUpdateDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClockAndTimekeepingDataUpdateDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	yearHigh, err := ReadSimpleField(ctx, "yearHigh", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'yearHigh' field"))
	}
	m.YearHigh = yearHigh

	yearLow, err := ReadSimpleField(ctx, "yearLow", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'yearLow' field"))
	}
	m.YearLow = yearLow

	month, err := ReadSimpleField(ctx, "month", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'month' field"))
	}
	m.Month = month

	day, err := ReadSimpleField(ctx, "day", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'day' field"))
	}
	m.Day = day

	dayOfWeek, err := ReadSimpleField(ctx, "dayOfWeek", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dayOfWeek' field"))
	}
	m.DayOfWeek = dayOfWeek

	if closeErr := readBuffer.CloseContext("ClockAndTimekeepingDataUpdateDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClockAndTimekeepingDataUpdateDate")
	}

	return m, nil
}

func (m *_ClockAndTimekeepingDataUpdateDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ClockAndTimekeepingDataUpdateDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ClockAndTimekeepingDataUpdateDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ClockAndTimekeepingDataUpdateDate")
		}

		if err := WriteSimpleField[byte](ctx, "yearHigh", m.GetYearHigh(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'yearHigh' field")
		}

		if err := WriteSimpleField[byte](ctx, "yearLow", m.GetYearLow(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'yearLow' field")
		}

		if err := WriteSimpleField[uint8](ctx, "month", m.GetMonth(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'month' field")
		}

		if err := WriteSimpleField[uint8](ctx, "day", m.GetDay(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'day' field")
		}

		if err := WriteSimpleField[uint8](ctx, "dayOfWeek", m.GetDayOfWeek(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'dayOfWeek' field")
		}

		if popErr := writeBuffer.PopContext("ClockAndTimekeepingDataUpdateDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ClockAndTimekeepingDataUpdateDate")
		}
		return nil
	}
	return m.ClockAndTimekeepingDataContract.(*_ClockAndTimekeepingData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ClockAndTimekeepingDataUpdateDate) IsClockAndTimekeepingDataUpdateDate() {}

func (m *_ClockAndTimekeepingDataUpdateDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
