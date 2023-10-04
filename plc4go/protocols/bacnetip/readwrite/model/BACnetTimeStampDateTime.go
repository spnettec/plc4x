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

// BACnetTimeStampDateTime is the corresponding interface of BACnetTimeStampDateTime
type BACnetTimeStampDateTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimeStamp
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() BACnetDateTimeEnclosed
}

// BACnetTimeStampDateTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetTimeStampDateTime.
// This is useful for switch cases.
type BACnetTimeStampDateTimeExactly interface {
	BACnetTimeStampDateTime
	isBACnetTimeStampDateTime() bool
}

// _BACnetTimeStampDateTime is the data-structure of this message
type _BACnetTimeStampDateTime struct {
	*_BACnetTimeStamp
	DateTimeValue BACnetDateTimeEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimeStampDateTime) InitializeParent(parent BACnetTimeStamp, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimeStampDateTime) GetParent() BACnetTimeStamp {
	return m._BACnetTimeStamp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStampDateTime) GetDateTimeValue() BACnetDateTimeEnclosed {
	return m.DateTimeValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampDateTime factory function for _BACnetTimeStampDateTime
func NewBACnetTimeStampDateTime(dateTimeValue BACnetDateTimeEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetTimeStampDateTime {
	_result := &_BACnetTimeStampDateTime{
		DateTimeValue:    dateTimeValue,
		_BACnetTimeStamp: NewBACnetTimeStamp(peekedTagHeader),
	}
	_result._BACnetTimeStamp._BACnetTimeStampChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimeStampDateTime(structType any) BACnetTimeStampDateTime {
	if casted, ok := structType.(BACnetTimeStampDateTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeStampDateTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeStampDateTime) GetTypeName() string {
	return "BACnetTimeStampDateTime"
}

func (m *_BACnetTimeStampDateTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeStampDateTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimeStampDateTimeParse(ctx context.Context, theBytes []byte) (BACnetTimeStampDateTime, error) {
	return BACnetTimeStampDateTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTimeStampDateTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimeStampDateTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetTimeStampDateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampDateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateTimeValue)
	if pullErr := readBuffer.PullContext("dateTimeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateTimeValue")
	}
	_dateTimeValue, _dateTimeValueErr := BACnetDateTimeEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _dateTimeValueErr != nil {
		return nil, errors.Wrap(_dateTimeValueErr, "Error parsing 'dateTimeValue' field of BACnetTimeStampDateTime")
	}
	dateTimeValue := _dateTimeValue.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("dateTimeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateTimeValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStampDateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampDateTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimeStampDateTime{
		_BACnetTimeStamp: &_BACnetTimeStamp{},
		DateTimeValue:    dateTimeValue,
	}
	_child._BACnetTimeStamp._BACnetTimeStampChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimeStampDateTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeStampDateTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimeStampDateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampDateTime")
		}

		// Simple Field (dateTimeValue)
		if pushErr := writeBuffer.PushContext("dateTimeValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateTimeValue")
		}
		_dateTimeValueErr := writeBuffer.WriteSerializable(ctx, m.GetDateTimeValue())
		if popErr := writeBuffer.PopContext("dateTimeValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateTimeValue")
		}
		if _dateTimeValueErr != nil {
			return errors.Wrap(_dateTimeValueErr, "Error serializing 'dateTimeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimeStampDateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimeStampDateTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimeStampDateTime) isBACnetTimeStampDateTime() bool {
	return true
}

func (m *_BACnetTimeStampDateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
