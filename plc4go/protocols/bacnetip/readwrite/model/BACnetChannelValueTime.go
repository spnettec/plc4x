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

// BACnetChannelValueTime is the corresponding interface of BACnetChannelValueTime
type BACnetChannelValueTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
}

// BACnetChannelValueTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueTime.
// This is useful for switch cases.
type BACnetChannelValueTimeExactly interface {
	BACnetChannelValueTime
	isBACnetChannelValueTime() bool
}

// _BACnetChannelValueTime is the data-structure of this message
type _BACnetChannelValueTime struct {
	*_BACnetChannelValue
	TimeValue BACnetApplicationTagTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueTime) InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueTime) GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueTime) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueTime factory function for _BACnetChannelValueTime
func NewBACnetChannelValueTime(timeValue BACnetApplicationTagTime, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueTime {
	_result := &_BACnetChannelValueTime{
		TimeValue:           timeValue,
		_BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueTime(structType any) BACnetChannelValueTime {
	if casted, ok := structType.(BACnetChannelValueTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueTime) GetTypeName() string {
	return "BACnetChannelValueTime"
}

func (m *_BACnetChannelValueTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetChannelValueTimeParse(theBytes []byte) (BACnetChannelValueTime, error) {
	return BACnetChannelValueTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetChannelValueTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetChannelValueTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeValue)
	if pullErr := readBuffer.PullContext("timeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeValue")
	}
	_timeValue, _timeValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _timeValueErr != nil {
		return nil, errors.Wrap(_timeValueErr, "Error parsing 'timeValue' field of BACnetChannelValueTime")
	}
	timeValue := _timeValue.(BACnetApplicationTagTime)
	if closeErr := readBuffer.CloseContext("timeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueTime{
		_BACnetChannelValue: &_BACnetChannelValue{},
		TimeValue:           timeValue,
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueTime")
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

		if popErr := writeBuffer.PopContext("BACnetChannelValueTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueTime) isBACnetChannelValueTime() bool {
	return true
}

func (m *_BACnetChannelValueTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
