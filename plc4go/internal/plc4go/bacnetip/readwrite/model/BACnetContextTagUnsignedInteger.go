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
type BACnetContextTagUnsignedInteger struct {
	*BACnetContextTag
	ValueUint8  *uint8
	ValueUint16 *uint16
	ValueUint24 *uint32
	ValueUint32 *uint32
	IsUint8     bool
	IsUint16    bool
	IsUint24    bool
	IsUint32    bool
	ActualValue uint32
}

// The corresponding interface
type IBACnetContextTagUnsignedInteger interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagUnsignedInteger) DataType() BACnetDataType {
	return BACnetDataType_UNSIGNED_INTEGER
}

func (m *BACnetContextTagUnsignedInteger) InitializeParent(parent *BACnetContextTag, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, actualLength uint32) {
	m.TagNumber = tagNumber
	m.TagClass = tagClass
	m.LengthValueType = lengthValueType
	m.ExtTagNumber = extTagNumber
	m.ExtLength = extLength
	m.ExtExtLength = extExtLength
	m.ExtExtExtLength = extExtExtLength
}

func NewBACnetContextTagUnsignedInteger(valueUint8 *uint8, valueUint16 *uint16, valueUint24 *uint32, valueUint32 *uint32, isUint8 bool, isUint16 bool, isUint24 bool, isUint32 bool, actualValue uint32, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, actualLength uint32) *BACnetContextTag {
	child := &BACnetContextTagUnsignedInteger{
		ValueUint8:       valueUint8,
		ValueUint16:      valueUint16,
		ValueUint24:      valueUint24,
		ValueUint32:      valueUint32,
		IsUint8:          isUint8,
		IsUint16:         isUint16,
		IsUint24:         isUint24,
		IsUint32:         isUint32,
		ActualValue:      actualValue,
		BACnetContextTag: NewBACnetContextTag(tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength, actualTagNumber, actualLength),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagUnsignedInteger(structType interface{}) *BACnetContextTagUnsignedInteger {
	castFunc := func(typ interface{}) *BACnetContextTagUnsignedInteger {
		if casted, ok := typ.(BACnetContextTagUnsignedInteger); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTagUnsignedInteger); ok {
			return casted
		}
		if casted, ok := typ.(BACnetContextTag); ok {
			return CastBACnetContextTagUnsignedInteger(casted.Child)
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return CastBACnetContextTagUnsignedInteger(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTagUnsignedInteger) GetTypeName() string {
	return "BACnetContextTagUnsignedInteger"
}

func (m *BACnetContextTagUnsignedInteger) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetContextTagUnsignedInteger) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint8)
	if m.ValueUint8 != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint16)
	if m.ValueUint16 != nil {
		lengthInBits += 16
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint24)
	if m.ValueUint24 != nil {
		lengthInBits += 24
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint32)
	if m.ValueUint32 != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTagUnsignedInteger) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetContextTagUnsignedIntegerParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, actualLength uint32) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagUnsignedInteger"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_isUint8 := bool((actualLength) == (1))
	isUint8 := bool(_isUint8)

	// Optional Field (valueUint8) (Can be skipped, if a given expression evaluates to false)
	var valueUint8 *uint8 = nil
	if isUint8 {
		_val, _err := readBuffer.ReadUint8("valueUint8", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint8' field")
		}
		valueUint8 = &_val
	}

	// Virtual field
	_isUint16 := bool((actualLength) == (2))
	isUint16 := bool(_isUint16)

	// Optional Field (valueUint16) (Can be skipped, if a given expression evaluates to false)
	var valueUint16 *uint16 = nil
	if isUint16 {
		_val, _err := readBuffer.ReadUint16("valueUint16", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint16' field")
		}
		valueUint16 = &_val
	}

	// Virtual field
	_isUint24 := bool((actualLength) == (3))
	isUint24 := bool(_isUint24)

	// Optional Field (valueUint24) (Can be skipped, if a given expression evaluates to false)
	var valueUint24 *uint32 = nil
	if isUint24 {
		_val, _err := readBuffer.ReadUint32("valueUint24", 24)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint24' field")
		}
		valueUint24 = &_val
	}

	// Virtual field
	_isUint32 := bool((actualLength) == (4))
	isUint32 := bool(_isUint32)

	// Optional Field (valueUint32) (Can be skipped, if a given expression evaluates to false)
	var valueUint32 *uint32 = nil
	if isUint32 {
		_val, _err := readBuffer.ReadUint32("valueUint32", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint32' field")
		}
		valueUint32 = &_val
	}

	// Virtual field
	_actualValue := utils.InlineIf(isUint8, func() interface{} { return uint32((*valueUint8)) }, func() interface{} {
		return uint32(uint32(utils.InlineIf(isUint16, func() interface{} { return uint32((*valueUint16)) }, func() interface{} {
			return uint32(uint32(utils.InlineIf(isUint24, func() interface{} { return uint32((*valueUint24)) }, func() interface{} {
				return uint32(uint32(utils.InlineIf(isUint32, func() interface{} { return uint32((*valueUint32)) }, func() interface{} { return uint32(uint32(0)) }).(uint32)))
			}).(uint32)))
		}).(uint32)))
	}).(uint32)
	actualValue := uint32(_actualValue)

	if closeErr := readBuffer.CloseContext("BACnetContextTagUnsignedInteger"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagUnsignedInteger{
		ValueUint8:       valueUint8,
		ValueUint16:      valueUint16,
		ValueUint24:      valueUint24,
		ValueUint32:      valueUint32,
		IsUint8:          isUint8,
		IsUint16:         isUint16,
		IsUint24:         isUint24,
		IsUint32:         isUint32,
		ActualValue:      actualValue,
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagUnsignedInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagUnsignedInteger"); pushErr != nil {
			return pushErr
		}
		// Virtual field
		if _isUint8Err := writeBuffer.WriteVirtual("isUint8", m.IsUint8); _isUint8Err != nil {
			return errors.Wrap(_isUint8Err, "Error serializing 'isUint8' field")
		}

		// Optional Field (valueUint8) (Can be skipped, if the value is null)
		var valueUint8 *uint8 = nil
		if m.ValueUint8 != nil {
			valueUint8 = m.ValueUint8
			_valueUint8Err := writeBuffer.WriteUint8("valueUint8", 8, *(valueUint8))
			if _valueUint8Err != nil {
				return errors.Wrap(_valueUint8Err, "Error serializing 'valueUint8' field")
			}
		}
		// Virtual field
		if _isUint16Err := writeBuffer.WriteVirtual("isUint16", m.IsUint16); _isUint16Err != nil {
			return errors.Wrap(_isUint16Err, "Error serializing 'isUint16' field")
		}

		// Optional Field (valueUint16) (Can be skipped, if the value is null)
		var valueUint16 *uint16 = nil
		if m.ValueUint16 != nil {
			valueUint16 = m.ValueUint16
			_valueUint16Err := writeBuffer.WriteUint16("valueUint16", 16, *(valueUint16))
			if _valueUint16Err != nil {
				return errors.Wrap(_valueUint16Err, "Error serializing 'valueUint16' field")
			}
		}
		// Virtual field
		if _isUint24Err := writeBuffer.WriteVirtual("isUint24", m.IsUint24); _isUint24Err != nil {
			return errors.Wrap(_isUint24Err, "Error serializing 'isUint24' field")
		}

		// Optional Field (valueUint24) (Can be skipped, if the value is null)
		var valueUint24 *uint32 = nil
		if m.ValueUint24 != nil {
			valueUint24 = m.ValueUint24
			_valueUint24Err := writeBuffer.WriteUint32("valueUint24", 24, *(valueUint24))
			if _valueUint24Err != nil {
				return errors.Wrap(_valueUint24Err, "Error serializing 'valueUint24' field")
			}
		}
		// Virtual field
		if _isUint32Err := writeBuffer.WriteVirtual("isUint32", m.IsUint32); _isUint32Err != nil {
			return errors.Wrap(_isUint32Err, "Error serializing 'isUint32' field")
		}

		// Optional Field (valueUint32) (Can be skipped, if the value is null)
		var valueUint32 *uint32 = nil
		if m.ValueUint32 != nil {
			valueUint32 = m.ValueUint32
			_valueUint32Err := writeBuffer.WriteUint32("valueUint32", 32, *(valueUint32))
			if _valueUint32Err != nil {
				return errors.Wrap(_valueUint32Err, "Error serializing 'valueUint32' field")
			}
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.ActualValue); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagUnsignedInteger"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagUnsignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
