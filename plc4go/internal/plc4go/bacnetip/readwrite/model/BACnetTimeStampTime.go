/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetTimeStampTime struct {
	*BACnetTimeStamp
	TimeValue *BACnetContextTagTime

	// Arguments.
	TagNumber uint8
}

// The corresponding interface
type IBACnetTimeStampTime interface {
	// GetTimeValue returns TimeValue
	GetTimeValue() *BACnetContextTagTime
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetTimeStampTime) PeekedTagNumber() uint8 {
	return uint8(0)
}

func (m *BACnetTimeStampTime) GetPeekedTagNumber() uint8 {
	return uint8(0)
}

func (m *BACnetTimeStampTime) InitializeParent(parent *BACnetTimeStamp, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetTimeStamp.OpeningTag = openingTag
	m.BACnetTimeStamp.PeekedTagHeader = peekedTagHeader
	m.BACnetTimeStamp.ClosingTag = closingTag
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetTimeStampTime) GetTimeValue() *BACnetContextTagTime {
	return m.TimeValue
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetTimeStampTime factory function for BACnetTimeStampTime
func NewBACnetTimeStampTime(timeValue *BACnetContextTagTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetTimeStamp {
	child := &BACnetTimeStampTime{
		TimeValue:       timeValue,
		BACnetTimeStamp: NewBACnetTimeStamp(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	child.Child = child
	return child.BACnetTimeStamp
}

func CastBACnetTimeStampTime(structType interface{}) *BACnetTimeStampTime {
	castFunc := func(typ interface{}) *BACnetTimeStampTime {
		if casted, ok := typ.(BACnetTimeStampTime); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTimeStampTime); ok {
			return casted
		}
		if casted, ok := typ.(BACnetTimeStamp); ok {
			return CastBACnetTimeStampTime(casted.Child)
		}
		if casted, ok := typ.(*BACnetTimeStamp); ok {
			return CastBACnetTimeStampTime(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTimeStampTime) GetTypeName() string {
	return "BACnetTimeStampTime"
}

func (m *BACnetTimeStampTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTimeStampTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetTimeStampTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimeStampTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetTimeStamp, error) {
	if pullErr := readBuffer.PullContext("BACnetTimeStampTime"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (timeValue)
	if pullErr := readBuffer.PullContext("timeValue"); pullErr != nil {
		return nil, pullErr
	}
	_timeValue, _timeValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType_TIME)
	if _timeValueErr != nil {
		return nil, errors.Wrap(_timeValueErr, "Error parsing 'timeValue' field")
	}
	timeValue := CastBACnetContextTagTime(_timeValue)
	if closeErr := readBuffer.CloseContext("timeValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStampTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetTimeStampTime{
		TimeValue:       CastBACnetContextTagTime(timeValue),
		BACnetTimeStamp: &BACnetTimeStamp{},
	}
	_child.BACnetTimeStamp.Child = _child
	return _child.BACnetTimeStamp, nil
}

func (m *BACnetTimeStampTime) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimeStampTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (timeValue)
		if pushErr := writeBuffer.PushContext("timeValue"); pushErr != nil {
			return pushErr
		}
		_timeValueErr := m.TimeValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeValue"); popErr != nil {
			return popErr
		}
		if _timeValueErr != nil {
			return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimeStampTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetTimeStampTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
