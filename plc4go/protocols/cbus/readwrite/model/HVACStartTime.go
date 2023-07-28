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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// HVACStartTime is the corresponding interface of HVACStartTime
type HVACStartTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMinutesSinceSunday12AM returns MinutesSinceSunday12AM (property field)
	GetMinutesSinceSunday12AM() uint16
	// GetHoursSinceSunday12AM returns HoursSinceSunday12AM (virtual field)
	GetHoursSinceSunday12AM() float32
	// GetDaysSinceSunday12AM returns DaysSinceSunday12AM (virtual field)
	GetDaysSinceSunday12AM() float32
	// GetDayOfWeek returns DayOfWeek (virtual field)
	GetDayOfWeek() uint8
	// GetHour returns Hour (virtual field)
	GetHour() uint8
	// GetMinute returns Minute (virtual field)
	GetMinute() uint8
}

// HVACStartTimeExactly can be used when we want exactly this type and not a type which fulfills HVACStartTime.
// This is useful for switch cases.
type HVACStartTimeExactly interface {
	HVACStartTime
	isHVACStartTime() bool
}

// _HVACStartTime is the data-structure of this message
type _HVACStartTime struct {
        MinutesSinceSunday12AM uint16
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACStartTime) GetMinutesSinceSunday12AM() uint16 {
	return m.MinutesSinceSunday12AM
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACStartTime) GetHoursSinceSunday12AM() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(float32(m.GetMinutesSinceSunday12AM()) / float32(float32(60)))
}

func (m *_HVACStartTime) GetDaysSinceSunday12AM() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(float32(m.GetHoursSinceSunday12AM()) / float32(float32(24)))
}

func (m *_HVACStartTime) GetDayOfWeek() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(uint8(m.GetDaysSinceSunday12AM()) + uint8(uint8(1)))
}

func (m *_HVACStartTime) GetHour() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(uint8(m.GetHoursSinceSunday12AM()) % uint8(uint8(24)))
}

func (m *_HVACStartTime) GetMinute() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(uint8(m.GetMinutesSinceSunday12AM()) % uint8(uint8(60)))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewHVACStartTime factory function for _HVACStartTime
func NewHVACStartTime( minutesSinceSunday12AM uint16 ) *_HVACStartTime {
return &_HVACStartTime{ MinutesSinceSunday12AM: minutesSinceSunday12AM }
}

// Deprecated: use the interface for direct cast
func CastHVACStartTime(structType any) HVACStartTime {
    if casted, ok := structType.(HVACStartTime); ok {
		return casted
	}
	if casted, ok := structType.(*HVACStartTime); ok {
		return *casted
	}
	return nil
}

func (m *_HVACStartTime) GetTypeName() string {
	return "HVACStartTime"
}

func (m *_HVACStartTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (minutesSinceSunday12AM)
	lengthInBits += 16;

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_HVACStartTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACStartTimeParse(ctx context.Context, theBytes []byte) (HVACStartTime, error) {
	return HVACStartTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACStartTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACStartTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HVACStartTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACStartTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minutesSinceSunday12AM)
_minutesSinceSunday12AM, _minutesSinceSunday12AMErr := readBuffer.ReadUint16("minutesSinceSunday12AM", 16)
	if _minutesSinceSunday12AMErr != nil {
		return nil, errors.Wrap(_minutesSinceSunday12AMErr, "Error parsing 'minutesSinceSunday12AM' field of HVACStartTime")
	}
	minutesSinceSunday12AM := _minutesSinceSunday12AM

	// Virtual field
	_hoursSinceSunday12AM := float32(minutesSinceSunday12AM) / float32(float32(60))
	hoursSinceSunday12AM := float32(_hoursSinceSunday12AM)
	_ = hoursSinceSunday12AM

	// Virtual field
	_daysSinceSunday12AM := float32(hoursSinceSunday12AM) / float32(float32(24))
	daysSinceSunday12AM := float32(_daysSinceSunday12AM)
	_ = daysSinceSunday12AM

	// Virtual field
	_dayOfWeek := uint8(daysSinceSunday12AM) + uint8(uint8(1))
	dayOfWeek := uint8(_dayOfWeek)
	_ = dayOfWeek

	// Virtual field
	_hour := uint8(hoursSinceSunday12AM) % uint8(uint8(24))
	hour := uint8(_hour)
	_ = hour

	// Virtual field
	_minute := uint8(minutesSinceSunday12AM) % uint8(uint8(60))
	minute := uint8(_minute)
	_ = minute

	if closeErr := readBuffer.CloseContext("HVACStartTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACStartTime")
	}

	// Create the instance
	return &_HVACStartTime{
			MinutesSinceSunday12AM: minutesSinceSunday12AM,
		}, nil
}

func (m *_HVACStartTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACStartTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("HVACStartTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACStartTime")
	}

	// Simple Field (minutesSinceSunday12AM)
	minutesSinceSunday12AM := uint16(m.GetMinutesSinceSunday12AM())
	_minutesSinceSunday12AMErr := writeBuffer.WriteUint16("minutesSinceSunday12AM", 16, (minutesSinceSunday12AM))
	if _minutesSinceSunday12AMErr != nil {
		return errors.Wrap(_minutesSinceSunday12AMErr, "Error serializing 'minutesSinceSunday12AM' field")
	}
	// Virtual field
	if _hoursSinceSunday12AMErr := writeBuffer.WriteVirtual(ctx, "hoursSinceSunday12AM", m.GetHoursSinceSunday12AM()); _hoursSinceSunday12AMErr != nil {
		return errors.Wrap(_hoursSinceSunday12AMErr, "Error serializing 'hoursSinceSunday12AM' field")
	}
	// Virtual field
	if _daysSinceSunday12AMErr := writeBuffer.WriteVirtual(ctx, "daysSinceSunday12AM", m.GetDaysSinceSunday12AM()); _daysSinceSunday12AMErr != nil {
		return errors.Wrap(_daysSinceSunday12AMErr, "Error serializing 'daysSinceSunday12AM' field")
	}
	// Virtual field
	if _dayOfWeekErr := writeBuffer.WriteVirtual(ctx, "dayOfWeek", m.GetDayOfWeek()); _dayOfWeekErr != nil {
		return errors.Wrap(_dayOfWeekErr, "Error serializing 'dayOfWeek' field")
	}
	// Virtual field
	if _hourErr := writeBuffer.WriteVirtual(ctx, "hour", m.GetHour()); _hourErr != nil {
		return errors.Wrap(_hourErr, "Error serializing 'hour' field")
	}
	// Virtual field
	if _minuteErr := writeBuffer.WriteVirtual(ctx, "minute", m.GetMinute()); _minuteErr != nil {
		return errors.Wrap(_minuteErr, "Error serializing 'minute' field")
	}

	if popErr := writeBuffer.PopContext("HVACStartTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACStartTime")
	}
	return nil
}


func (m *_HVACStartTime) isHVACStartTime() bool {
	return true
}

func (m *_HVACStartTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



