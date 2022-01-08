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
type BACnetContextTagTime struct {
	*BACnetContextTag
	Hour                 int8
	Minute               int8
	Second               int8
	Fractional           int8
	Wildcard             int8
	HourIsWildcard       bool
	MinuteIsWildcard     bool
	SecondIsWildcard     bool
	FractionalIsWildcard bool
}

// The corresponding interface
type IBACnetContextTagTime interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagTime) DataType() BACnetDataType {
	return BACnetDataType_TIME
}

func (m *BACnetContextTagTime) InitializeParent(parent *BACnetContextTag, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, actualLength uint32) {
	m.TagNumber = tagNumber
	m.TagClass = tagClass
	m.LengthValueType = lengthValueType
	m.ExtTagNumber = extTagNumber
	m.ExtLength = extLength
	m.ExtExtLength = extExtLength
	m.ExtExtExtLength = extExtExtLength
}

func NewBACnetContextTagTime(hour int8, minute int8, second int8, fractional int8, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetContextTag {
	child := &BACnetContextTagTime{
		Hour:             hour,
		Minute:           minute,
		Second:           second,
		Fractional:       fractional,
		BACnetContextTag: NewBACnetContextTag(tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagTime(structType interface{}) *BACnetContextTagTime {
	castFunc := func(typ interface{}) *BACnetContextTagTime {
		if casted, ok := typ.(BACnetContextTagTime); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTagTime); ok {
			return casted
		}
		if casted, ok := typ.(BACnetContextTag); ok {
			return CastBACnetContextTagTime(casted.Child)
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return CastBACnetContextTagTime(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTagTime) GetTypeName() string {
	return "BACnetContextTagTime"
}

func (m *BACnetContextTagTime) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetContextTagTime) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Simple field (hour)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (minute)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (second)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (fractional)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTagTime) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetContextTagTimeParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagTime"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_wildcard := 0xFF
	wildcard := int8(_wildcard)

	// Simple Field (hour)
	_hour, _hourErr := readBuffer.ReadInt8("hour", 8)
	if _hourErr != nil {
		return nil, errors.Wrap(_hourErr, "Error parsing 'hour' field")
	}
	hour := _hour

	// Virtual field
	_hourIsWildcard := bool((hour) == (wildcard))
	hourIsWildcard := bool(_hourIsWildcard)

	// Simple Field (minute)
	_minute, _minuteErr := readBuffer.ReadInt8("minute", 8)
	if _minuteErr != nil {
		return nil, errors.Wrap(_minuteErr, "Error parsing 'minute' field")
	}
	minute := _minute

	// Virtual field
	_minuteIsWildcard := bool((minute) == (wildcard))
	minuteIsWildcard := bool(_minuteIsWildcard)

	// Simple Field (second)
	_second, _secondErr := readBuffer.ReadInt8("second", 8)
	if _secondErr != nil {
		return nil, errors.Wrap(_secondErr, "Error parsing 'second' field")
	}
	second := _second

	// Virtual field
	_secondIsWildcard := bool((second) == (wildcard))
	secondIsWildcard := bool(_secondIsWildcard)

	// Simple Field (fractional)
	_fractional, _fractionalErr := readBuffer.ReadInt8("fractional", 8)
	if _fractionalErr != nil {
		return nil, errors.Wrap(_fractionalErr, "Error parsing 'fractional' field")
	}
	fractional := _fractional

	// Virtual field
	_fractionalIsWildcard := bool((fractional) == (wildcard))
	fractionalIsWildcard := bool(_fractionalIsWildcard)

	if closeErr := readBuffer.CloseContext("BACnetContextTagTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagTime{
		Hour:                 hour,
		Minute:               minute,
		Second:               second,
		Fractional:           fractional,
		Wildcard:             wildcard,
		HourIsWildcard:       hourIsWildcard,
		MinuteIsWildcard:     minuteIsWildcard,
		SecondIsWildcard:     secondIsWildcard,
		FractionalIsWildcard: fractionalIsWildcard,
		BACnetContextTag:     &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagTime) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagTime"); pushErr != nil {
			return pushErr
		}
		// Virtual field
		if _wildcardErr := writeBuffer.WriteVirtual("wildcard", m.Wildcard); _wildcardErr != nil {
			return errors.Wrap(_wildcardErr, "Error serializing 'wildcard' field")
		}

		// Simple Field (hour)
		hour := int8(m.Hour)
		_hourErr := writeBuffer.WriteInt8("hour", 8, (hour))
		if _hourErr != nil {
			return errors.Wrap(_hourErr, "Error serializing 'hour' field")
		}
		// Virtual field
		if _hourIsWildcardErr := writeBuffer.WriteVirtual("hourIsWildcard", m.HourIsWildcard); _hourIsWildcardErr != nil {
			return errors.Wrap(_hourIsWildcardErr, "Error serializing 'hourIsWildcard' field")
		}

		// Simple Field (minute)
		minute := int8(m.Minute)
		_minuteErr := writeBuffer.WriteInt8("minute", 8, (minute))
		if _minuteErr != nil {
			return errors.Wrap(_minuteErr, "Error serializing 'minute' field")
		}
		// Virtual field
		if _minuteIsWildcardErr := writeBuffer.WriteVirtual("minuteIsWildcard", m.MinuteIsWildcard); _minuteIsWildcardErr != nil {
			return errors.Wrap(_minuteIsWildcardErr, "Error serializing 'minuteIsWildcard' field")
		}

		// Simple Field (second)
		second := int8(m.Second)
		_secondErr := writeBuffer.WriteInt8("second", 8, (second))
		if _secondErr != nil {
			return errors.Wrap(_secondErr, "Error serializing 'second' field")
		}
		// Virtual field
		if _secondIsWildcardErr := writeBuffer.WriteVirtual("secondIsWildcard", m.SecondIsWildcard); _secondIsWildcardErr != nil {
			return errors.Wrap(_secondIsWildcardErr, "Error serializing 'secondIsWildcard' field")
		}

		// Simple Field (fractional)
		fractional := int8(m.Fractional)
		_fractionalErr := writeBuffer.WriteInt8("fractional", 8, (fractional))
		if _fractionalErr != nil {
			return errors.Wrap(_fractionalErr, "Error serializing 'fractional' field")
		}
		// Virtual field
		if _fractionalIsWildcardErr := writeBuffer.WriteVirtual("fractionalIsWildcard", m.FractionalIsWildcard); _fractionalIsWildcardErr != nil {
			return errors.Wrap(_fractionalIsWildcardErr, "Error serializing 'fractionalIsWildcard' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
