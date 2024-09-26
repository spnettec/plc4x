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

// BACnetTimeStampDateTime is the corresponding interface of BACnetTimeStampDateTime
type BACnetTimeStampDateTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimeStamp
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() BACnetDateTimeEnclosed
	// IsBACnetTimeStampDateTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimeStampDateTime()
}

// _BACnetTimeStampDateTime is the data-structure of this message
type _BACnetTimeStampDateTime struct {
	BACnetTimeStampContract
	DateTimeValue BACnetDateTimeEnclosed
}

var _ BACnetTimeStampDateTime = (*_BACnetTimeStampDateTime)(nil)
var _ BACnetTimeStampRequirements = (*_BACnetTimeStampDateTime)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimeStampDateTime) GetParent() BACnetTimeStampContract {
	return m.BACnetTimeStampContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStampDateTime) GetDateTimeValue() BACnetDateTimeEnclosed {
	return m.DateTimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampDateTime factory function for _BACnetTimeStampDateTime
func NewBACnetTimeStampDateTime(peekedTagHeader BACnetTagHeader, dateTimeValue BACnetDateTimeEnclosed) *_BACnetTimeStampDateTime {
	if dateTimeValue == nil {
		panic("dateTimeValue of type BACnetDateTimeEnclosed for BACnetTimeStampDateTime must not be nil")
	}
	_result := &_BACnetTimeStampDateTime{
		BACnetTimeStampContract: NewBACnetTimeStamp(peekedTagHeader),
		DateTimeValue:           dateTimeValue,
	}
	_result.BACnetTimeStampContract.(*_BACnetTimeStamp)._SubType = _result
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
	lengthInBits := uint16(m.BACnetTimeStampContract.(*_BACnetTimeStamp).getLengthInBits(ctx))

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeStampDateTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimeStampDateTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimeStamp) (__bACnetTimeStampDateTime BACnetTimeStampDateTime, err error) {
	m.BACnetTimeStampContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStampDateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampDateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateTimeValue, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "dateTimeValue", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateTimeValue' field"))
	}
	m.DateTimeValue = dateTimeValue

	if closeErr := readBuffer.CloseContext("BACnetTimeStampDateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampDateTime")
	}

	return m, nil
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

		if err := WriteSimpleField[BACnetDateTimeEnclosed](ctx, "dateTimeValue", m.GetDateTimeValue(), WriteComplex[BACnetDateTimeEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dateTimeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimeStampDateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimeStampDateTime")
		}
		return nil
	}
	return m.BACnetTimeStampContract.(*_BACnetTimeStamp).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimeStampDateTime) IsBACnetTimeStampDateTime() {}

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
