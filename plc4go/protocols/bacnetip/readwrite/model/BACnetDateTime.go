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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetDateTime is the corresponding interface of BACnetDateTime
type BACnetDateTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
}

// BACnetDateTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetDateTime.
// This is useful for switch cases.
type BACnetDateTimeExactly interface {
	BACnetDateTime
	isBACnetDateTime() bool
}

// _BACnetDateTime is the data-structure of this message
type _BACnetDateTime struct {
	DateValue BACnetApplicationTagDate
	TimeValue BACnetApplicationTagTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDateTime) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

func (m *_BACnetDateTime) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDateTime factory function for _BACnetDateTime
func NewBACnetDateTime(dateValue BACnetApplicationTagDate, timeValue BACnetApplicationTagTime) *_BACnetDateTime {
	return &_BACnetDateTime{DateValue: dateValue, TimeValue: timeValue}
}

// Deprecated: use the interface for direct cast
func CastBACnetDateTime(structType any) BACnetDateTime {
	if casted, ok := structType.(BACnetDateTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDateTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDateTime) GetTypeName() string {
	return "BACnetDateTime"
}

func (m *_BACnetDateTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits(ctx)

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDateTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDateTimeParse(theBytes []byte) (BACnetDateTime, error) {
	return BACnetDateTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetDateTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateValue)
	if pullErr := readBuffer.PullContext("dateValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateValue")
	}
	_dateValue, _dateValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _dateValueErr != nil {
		return nil, errors.Wrap(_dateValueErr, "Error parsing 'dateValue' field of BACnetDateTime")
	}
	dateValue := _dateValue.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("dateValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateValue")
	}

	// Simple Field (timeValue)
	if pullErr := readBuffer.PullContext("timeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeValue")
	}
	_timeValue, _timeValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _timeValueErr != nil {
		return nil, errors.Wrap(_timeValueErr, "Error parsing 'timeValue' field of BACnetDateTime")
	}
	timeValue := _timeValue.(BACnetApplicationTagTime)
	if closeErr := readBuffer.CloseContext("timeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetDateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateTime")
	}

	// Create the instance
	return &_BACnetDateTime{
		DateValue: dateValue,
		TimeValue: timeValue,
	}, nil
}

func (m *_BACnetDateTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDateTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDateTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateTime")
	}

	// Simple Field (dateValue)
	if pushErr := writeBuffer.PushContext("dateValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for dateValue")
	}
	_dateValueErr := writeBuffer.WriteSerializable(ctx, m.GetDateValue())
	if popErr := writeBuffer.PopContext("dateValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for dateValue")
	}
	if _dateValueErr != nil {
		return errors.Wrap(_dateValueErr, "Error serializing 'dateValue' field")
	}

	// Simple Field (timeValue)
	if pushErr := writeBuffer.PushContext("timeValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeValue")
	}
	_timeValueErr := writeBuffer.WriteSerializable(ctx, m.GetTimeValue())
	if popErr := writeBuffer.PopContext("timeValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeValue")
	}
	if _timeValueErr != nil {
		return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateTime")
	}
	return nil
}

func (m *_BACnetDateTime) isBACnetDateTime() bool {
	return true
}

func (m *_BACnetDateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
