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
type BACnetContextTagSignedInteger struct {
	*BACnetContextTag
	ValueInt8   *int8
	ValueInt16  *int16
	ValueInt32  *int32
	ValueInt64  *int64
	IsInt8      bool
	IsInt16     bool
	IsInt32     bool
	IsInt64     bool
	ActualValue uint64
}

// The corresponding interface
type IBACnetContextTagSignedInteger interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagSignedInteger) DataType() BACnetDataType {
	return BACnetDataType_SIGNED_INTEGER
}

func (m *BACnetContextTagSignedInteger) InitializeParent(parent *BACnetContextTag, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, actualLength uint32) {
	m.TagNumber = tagNumber
	m.TagClass = tagClass
	m.LengthValueType = lengthValueType
	m.ExtTagNumber = extTagNumber
	m.ExtLength = extLength
	m.ExtExtLength = extExtLength
	m.ExtExtExtLength = extExtExtLength
}

func NewBACnetContextTagSignedInteger(valueInt8 *int8, valueInt16 *int16, valueInt32 *int32, valueInt64 *int64, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetContextTag {
	child := &BACnetContextTagSignedInteger{
		ValueInt8:        valueInt8,
		ValueInt16:       valueInt16,
		ValueInt32:       valueInt32,
		ValueInt64:       valueInt64,
		BACnetContextTag: NewBACnetContextTag(tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagSignedInteger(structType interface{}) *BACnetContextTagSignedInteger {
	castFunc := func(typ interface{}) *BACnetContextTagSignedInteger {
		if casted, ok := typ.(BACnetContextTagSignedInteger); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTagSignedInteger); ok {
			return casted
		}
		if casted, ok := typ.(BACnetContextTag); ok {
			return CastBACnetContextTagSignedInteger(casted.Child)
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return CastBACnetContextTagSignedInteger(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTagSignedInteger) GetTypeName() string {
	return "BACnetContextTagSignedInteger"
}

func (m *BACnetContextTagSignedInteger) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetContextTagSignedInteger) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt8)
	if m.ValueInt8 != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt16)
	if m.ValueInt16 != nil {
		lengthInBits += 16
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt32)
	if m.ValueInt32 != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt64)
	if m.ValueInt64 != nil {
		lengthInBits += 64
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTagSignedInteger) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetContextTagSignedIntegerParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, actualLength uint32) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagSignedInteger"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_isInt8 := bool((actualLength) == (1))
	isInt8 := bool(_isInt8)

	// Optional Field (valueInt8) (Can be skipped, if a given expression evaluates to false)
	var valueInt8 *int8 = nil
	if isInt8 {
		_val, _err := readBuffer.ReadInt8("valueInt8", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt8' field")
		}
		valueInt8 = &_val
	}

	// Virtual field
	_isInt16 := bool((actualLength) == (2))
	isInt16 := bool(_isInt16)

	// Optional Field (valueInt16) (Can be skipped, if a given expression evaluates to false)
	var valueInt16 *int16 = nil
	if isInt16 {
		_val, _err := readBuffer.ReadInt16("valueInt16", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt16' field")
		}
		valueInt16 = &_val
	}

	// Virtual field
	_isInt32 := bool((actualLength) == (3))
	isInt32 := bool(_isInt32)

	// Optional Field (valueInt32) (Can be skipped, if a given expression evaluates to false)
	var valueInt32 *int32 = nil
	if isInt32 {
		_val, _err := readBuffer.ReadInt32("valueInt32", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt32' field")
		}
		valueInt32 = &_val
	}

	// Virtual field
	_isInt64 := bool((actualLength) == (4))
	isInt64 := bool(_isInt64)

	// Optional Field (valueInt64) (Can be skipped, if a given expression evaluates to false)
	var valueInt64 *int64 = nil
	if isInt64 {
		_val, _err := readBuffer.ReadInt64("valueInt64", 64)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt64' field")
		}
		valueInt64 = &_val
	}

	// Virtual field
	_actualValue := utils.InlineIf(isInt8, func() interface{} { return uint64((*valueInt8)) }, func() interface{} {
		return uint64(uint64(utils.InlineIf(isInt16, func() interface{} { return uint64((*valueInt16)) }, func() interface{} {
			return uint64(uint64(utils.InlineIf(isInt64, func() interface{} { return uint64((*valueInt64)) }, func() interface{} { return uint64(uint64(0)) }).(uint64)))
		}).(uint64)))
	}).(uint64)
	actualValue := uint64(_actualValue)

	if closeErr := readBuffer.CloseContext("BACnetContextTagSignedInteger"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagSignedInteger{
		ValueInt8:        valueInt8,
		ValueInt16:       valueInt16,
		ValueInt32:       valueInt32,
		ValueInt64:       valueInt64,
		IsInt8:           isInt8,
		IsInt16:          isInt16,
		IsInt32:          isInt32,
		IsInt64:          isInt64,
		ActualValue:      actualValue,
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagSignedInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagSignedInteger"); pushErr != nil {
			return pushErr
		}
		// Virtual field
		if _isInt8Err := writeBuffer.WriteVirtual("isInt8", m.IsInt8); _isInt8Err != nil {
			return errors.Wrap(_isInt8Err, "Error serializing 'isInt8' field")
		}

		// Optional Field (valueInt8) (Can be skipped, if the value is null)
		var valueInt8 *int8 = nil
		if m.ValueInt8 != nil {
			valueInt8 = m.ValueInt8
			_valueInt8Err := writeBuffer.WriteInt8("valueInt8", 8, *(valueInt8))
			if _valueInt8Err != nil {
				return errors.Wrap(_valueInt8Err, "Error serializing 'valueInt8' field")
			}
		}
		// Virtual field
		if _isInt16Err := writeBuffer.WriteVirtual("isInt16", m.IsInt16); _isInt16Err != nil {
			return errors.Wrap(_isInt16Err, "Error serializing 'isInt16' field")
		}

		// Optional Field (valueInt16) (Can be skipped, if the value is null)
		var valueInt16 *int16 = nil
		if m.ValueInt16 != nil {
			valueInt16 = m.ValueInt16
			_valueInt16Err := writeBuffer.WriteInt16("valueInt16", 16, *(valueInt16))
			if _valueInt16Err != nil {
				return errors.Wrap(_valueInt16Err, "Error serializing 'valueInt16' field")
			}
		}
		// Virtual field
		if _isInt32Err := writeBuffer.WriteVirtual("isInt32", m.IsInt32); _isInt32Err != nil {
			return errors.Wrap(_isInt32Err, "Error serializing 'isInt32' field")
		}

		// Optional Field (valueInt32) (Can be skipped, if the value is null)
		var valueInt32 *int32 = nil
		if m.ValueInt32 != nil {
			valueInt32 = m.ValueInt32
			_valueInt32Err := writeBuffer.WriteInt32("valueInt32", 32, *(valueInt32))
			if _valueInt32Err != nil {
				return errors.Wrap(_valueInt32Err, "Error serializing 'valueInt32' field")
			}
		}
		// Virtual field
		if _isInt64Err := writeBuffer.WriteVirtual("isInt64", m.IsInt64); _isInt64Err != nil {
			return errors.Wrap(_isInt64Err, "Error serializing 'isInt64' field")
		}

		// Optional Field (valueInt64) (Can be skipped, if the value is null)
		var valueInt64 *int64 = nil
		if m.ValueInt64 != nil {
			valueInt64 = m.ValueInt64
			_valueInt64Err := writeBuffer.WriteInt64("valueInt64", 64, *(valueInt64))
			if _valueInt64Err != nil {
				return errors.Wrap(_valueInt64Err, "Error serializing 'valueInt64' field")
			}
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.ActualValue); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagSignedInteger"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagSignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
