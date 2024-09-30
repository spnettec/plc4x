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

// HVACStartTime is the corresponding interface of HVACStartTime
type HVACStartTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
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
	// IsHVACStartTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHVACStartTime()
	// CreateBuilder creates a HVACStartTimeBuilder
	CreateHVACStartTimeBuilder() HVACStartTimeBuilder
}

// _HVACStartTime is the data-structure of this message
type _HVACStartTime struct {
	MinutesSinceSunday12AM uint16
}

var _ HVACStartTime = (*_HVACStartTime)(nil)

// NewHVACStartTime factory function for _HVACStartTime
func NewHVACStartTime(minutesSinceSunday12AM uint16) *_HVACStartTime {
	return &_HVACStartTime{MinutesSinceSunday12AM: minutesSinceSunday12AM}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HVACStartTimeBuilder is a builder for HVACStartTime
type HVACStartTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minutesSinceSunday12AM uint16) HVACStartTimeBuilder
	// WithMinutesSinceSunday12AM adds MinutesSinceSunday12AM (property field)
	WithMinutesSinceSunday12AM(uint16) HVACStartTimeBuilder
	// Build builds the HVACStartTime or returns an error if something is wrong
	Build() (HVACStartTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HVACStartTime
}

// NewHVACStartTimeBuilder() creates a HVACStartTimeBuilder
func NewHVACStartTimeBuilder() HVACStartTimeBuilder {
	return &_HVACStartTimeBuilder{_HVACStartTime: new(_HVACStartTime)}
}

type _HVACStartTimeBuilder struct {
	*_HVACStartTime

	err *utils.MultiError
}

var _ (HVACStartTimeBuilder) = (*_HVACStartTimeBuilder)(nil)

func (b *_HVACStartTimeBuilder) WithMandatoryFields(minutesSinceSunday12AM uint16) HVACStartTimeBuilder {
	return b.WithMinutesSinceSunday12AM(minutesSinceSunday12AM)
}

func (b *_HVACStartTimeBuilder) WithMinutesSinceSunday12AM(minutesSinceSunday12AM uint16) HVACStartTimeBuilder {
	b.MinutesSinceSunday12AM = minutesSinceSunday12AM
	return b
}

func (b *_HVACStartTimeBuilder) Build() (HVACStartTime, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HVACStartTime.deepCopy(), nil
}

func (b *_HVACStartTimeBuilder) MustBuild() HVACStartTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HVACStartTimeBuilder) DeepCopy() any {
	_copy := b.CreateHVACStartTimeBuilder().(*_HVACStartTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHVACStartTimeBuilder creates a HVACStartTimeBuilder
func (b *_HVACStartTime) CreateHVACStartTimeBuilder() HVACStartTimeBuilder {
	if b == nil {
		return NewHVACStartTimeBuilder()
	}
	return &_HVACStartTimeBuilder{_HVACStartTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACStartTime) GetMinutesSinceSunday12AM() uint16 {
	return m.MinutesSinceSunday12AM
}

///////////////////////
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

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
	lengthInBits += 16

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

func HVACStartTimeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACStartTime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACStartTime, error) {
		return HVACStartTimeParseWithBuffer(ctx, readBuffer)
	}
}

func HVACStartTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACStartTime, error) {
	v, err := (&_HVACStartTime{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_HVACStartTime) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__hVACStartTime HVACStartTime, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACStartTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACStartTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minutesSinceSunday12AM, err := ReadSimpleField(ctx, "minutesSinceSunday12AM", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minutesSinceSunday12AM' field"))
	}
	m.MinutesSinceSunday12AM = minutesSinceSunday12AM

	hoursSinceSunday12AM, err := ReadVirtualField[float32](ctx, "hoursSinceSunday12AM", (*float32)(nil), float32(minutesSinceSunday12AM)/float32(float32(60)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hoursSinceSunday12AM' field"))
	}
	_ = hoursSinceSunday12AM

	daysSinceSunday12AM, err := ReadVirtualField[float32](ctx, "daysSinceSunday12AM", (*float32)(nil), float32(hoursSinceSunday12AM)/float32(float32(24)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'daysSinceSunday12AM' field"))
	}
	_ = daysSinceSunday12AM

	dayOfWeek, err := ReadVirtualField[uint8](ctx, "dayOfWeek", (*uint8)(nil), uint8(daysSinceSunday12AM)+uint8(uint8(1)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dayOfWeek' field"))
	}
	_ = dayOfWeek

	hour, err := ReadVirtualField[uint8](ctx, "hour", (*uint8)(nil), uint8(hoursSinceSunday12AM)%uint8(uint8(24)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hour' field"))
	}
	_ = hour

	minute, err := ReadVirtualField[uint8](ctx, "minute", (*uint8)(nil), uint8(minutesSinceSunday12AM)%uint8(uint8(60)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minute' field"))
	}
	_ = minute

	if closeErr := readBuffer.CloseContext("HVACStartTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACStartTime")
	}

	return m, nil
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
	if pushErr := writeBuffer.PushContext("HVACStartTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACStartTime")
	}

	if err := WriteSimpleField[uint16](ctx, "minutesSinceSunday12AM", m.GetMinutesSinceSunday12AM(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'minutesSinceSunday12AM' field")
	}
	// Virtual field
	hoursSinceSunday12AM := m.GetHoursSinceSunday12AM()
	_ = hoursSinceSunday12AM
	if _hoursSinceSunday12AMErr := writeBuffer.WriteVirtual(ctx, "hoursSinceSunday12AM", m.GetHoursSinceSunday12AM()); _hoursSinceSunday12AMErr != nil {
		return errors.Wrap(_hoursSinceSunday12AMErr, "Error serializing 'hoursSinceSunday12AM' field")
	}
	// Virtual field
	daysSinceSunday12AM := m.GetDaysSinceSunday12AM()
	_ = daysSinceSunday12AM
	if _daysSinceSunday12AMErr := writeBuffer.WriteVirtual(ctx, "daysSinceSunday12AM", m.GetDaysSinceSunday12AM()); _daysSinceSunday12AMErr != nil {
		return errors.Wrap(_daysSinceSunday12AMErr, "Error serializing 'daysSinceSunday12AM' field")
	}
	// Virtual field
	dayOfWeek := m.GetDayOfWeek()
	_ = dayOfWeek
	if _dayOfWeekErr := writeBuffer.WriteVirtual(ctx, "dayOfWeek", m.GetDayOfWeek()); _dayOfWeekErr != nil {
		return errors.Wrap(_dayOfWeekErr, "Error serializing 'dayOfWeek' field")
	}
	// Virtual field
	hour := m.GetHour()
	_ = hour
	if _hourErr := writeBuffer.WriteVirtual(ctx, "hour", m.GetHour()); _hourErr != nil {
		return errors.Wrap(_hourErr, "Error serializing 'hour' field")
	}
	// Virtual field
	minute := m.GetMinute()
	_ = minute
	if _minuteErr := writeBuffer.WriteVirtual(ctx, "minute", m.GetMinute()); _minuteErr != nil {
		return errors.Wrap(_minuteErr, "Error serializing 'minute' field")
	}

	if popErr := writeBuffer.PopContext("HVACStartTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACStartTime")
	}
	return nil
}

func (m *_HVACStartTime) IsHVACStartTime() {}

func (m *_HVACStartTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HVACStartTime) deepCopy() *_HVACStartTime {
	if m == nil {
		return nil
	}
	_HVACStartTimeCopy := &_HVACStartTime{
		m.MinutesSinceSunday12AM,
	}
	return _HVACStartTimeCopy
}

func (m *_HVACStartTime) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
