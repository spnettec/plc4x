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

// BACnetPriorityValueTime is the corresponding interface of BACnetPriorityValueTime
type BACnetPriorityValueTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
}

// BACnetPriorityValueTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueTime.
// This is useful for switch cases.
type BACnetPriorityValueTimeExactly interface {
	BACnetPriorityValueTime
	isBACnetPriorityValueTime() bool
}

// _BACnetPriorityValueTime is the data-structure of this message
type _BACnetPriorityValueTime struct {
	*_BACnetPriorityValue
	TimeValue BACnetApplicationTagTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueTime) InitializeParent(parent BACnetPriorityValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueTime) GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueTime) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueTime factory function for _BACnetPriorityValueTime
func NewBACnetPriorityValueTime(timeValue BACnetApplicationTagTime, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueTime {
	_result := &_BACnetPriorityValueTime{
		TimeValue:            timeValue,
		_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueTime(structType interface{}) BACnetPriorityValueTime {
	if casted, ok := structType.(BACnetPriorityValueTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueTime) GetTypeName() string {
	return "BACnetPriorityValueTime"
}

func (m *_BACnetPriorityValueTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPriorityValueTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPriorityValueTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueTimeParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeValue)
	if pullErr := readBuffer.PullContext("timeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeValue")
	}
	_timeValue, _timeValueErr := BACnetApplicationTagParse(readBuffer)
	if _timeValueErr != nil {
		return nil, errors.Wrap(_timeValueErr, "Error parsing 'timeValue' field of BACnetPriorityValueTime")
	}
	timeValue := _timeValue.(BACnetApplicationTagTime)
	if closeErr := readBuffer.CloseContext("timeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueTime{
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		TimeValue: timeValue,
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueTime")
		}

		// Simple Field (timeValue)
		if pushErr := writeBuffer.PushContext("timeValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeValue")
		}
		_timeValueErr := writeBuffer.WriteSerializable(m.GetTimeValue())
		if popErr := writeBuffer.PopContext("timeValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeValue")
		}
		if _timeValueErr != nil {
			return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueTime) isBACnetPriorityValueTime() bool {
	return true
}

func (m *_BACnetPriorityValueTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
