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

// BACnetEventLogRecordLogDatumTimeChange is the corresponding interface of BACnetEventLogRecordLogDatumTimeChange
type BACnetEventLogRecordLogDatumTimeChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventLogRecordLogDatum
	// GetTimeChange returns TimeChange (property field)
	GetTimeChange() BACnetContextTagReal
}

// BACnetEventLogRecordLogDatumTimeChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetEventLogRecordLogDatumTimeChange.
// This is useful for switch cases.
type BACnetEventLogRecordLogDatumTimeChangeExactly interface {
	BACnetEventLogRecordLogDatumTimeChange
	isBACnetEventLogRecordLogDatumTimeChange() bool
}

// _BACnetEventLogRecordLogDatumTimeChange is the data-structure of this message
type _BACnetEventLogRecordLogDatumTimeChange struct {
	*_BACnetEventLogRecordLogDatum
	TimeChange BACnetContextTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventLogRecordLogDatumTimeChange) InitializeParent(parent BACnetEventLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) GetParent() BACnetEventLogRecordLogDatum {
	return m._BACnetEventLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventLogRecordLogDatumTimeChange) GetTimeChange() BACnetContextTagReal {
	return m.TimeChange
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventLogRecordLogDatumTimeChange factory function for _BACnetEventLogRecordLogDatumTimeChange
func NewBACnetEventLogRecordLogDatumTimeChange(timeChange BACnetContextTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventLogRecordLogDatumTimeChange {
	_result := &_BACnetEventLogRecordLogDatumTimeChange{
		TimeChange:                    timeChange,
		_BACnetEventLogRecordLogDatum: NewBACnetEventLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetEventLogRecordLogDatum._BACnetEventLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventLogRecordLogDatumTimeChange(structType any) BACnetEventLogRecordLogDatumTimeChange {
	if casted, ok := structType.(BACnetEventLogRecordLogDatumTimeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatumTimeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) GetTypeName() string {
	return "BACnetEventLogRecordLogDatumTimeChange"
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timeChange)
	lengthInBits += m.TimeChange.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventLogRecordLogDatumTimeChangeParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventLogRecordLogDatumTimeChange, error) {
	return BACnetEventLogRecordLogDatumTimeChangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventLogRecordLogDatumTimeChangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventLogRecordLogDatumTimeChange, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetEventLogRecordLogDatumTimeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventLogRecordLogDatumTimeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeChange)
	if pullErr := readBuffer.PullContext("timeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeChange")
	}
	_timeChange, _timeChangeErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_REAL))
	if _timeChangeErr != nil {
		return nil, errors.Wrap(_timeChangeErr, "Error parsing 'timeChange' field of BACnetEventLogRecordLogDatumTimeChange")
	}
	timeChange := _timeChange.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("timeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeChange")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecordLogDatumTimeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventLogRecordLogDatumTimeChange")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventLogRecordLogDatumTimeChange{
		_BACnetEventLogRecordLogDatum: &_BACnetEventLogRecordLogDatum{
			TagNumber: tagNumber,
		},
		TimeChange: timeChange,
	}
	_child._BACnetEventLogRecordLogDatum._BACnetEventLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventLogRecordLogDatumTimeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventLogRecordLogDatumTimeChange")
		}

		// Simple Field (timeChange)
		if pushErr := writeBuffer.PushContext("timeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeChange")
		}
		_timeChangeErr := writeBuffer.WriteSerializable(ctx, m.GetTimeChange())
		if popErr := writeBuffer.PopContext("timeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeChange")
		}
		if _timeChangeErr != nil {
			return errors.Wrap(_timeChangeErr, "Error serializing 'timeChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventLogRecordLogDatumTimeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventLogRecordLogDatumTimeChange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) isBACnetEventLogRecordLogDatumTimeChange() bool {
	return true
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
