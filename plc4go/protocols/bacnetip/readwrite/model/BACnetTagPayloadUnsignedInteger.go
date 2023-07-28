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


// BACnetTagPayloadUnsignedInteger is the corresponding interface of BACnetTagPayloadUnsignedInteger
type BACnetTagPayloadUnsignedInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetValueUint8 returns ValueUint8 (property field)
	GetValueUint8() *uint8
	// GetValueUint16 returns ValueUint16 (property field)
	GetValueUint16() *uint16
	// GetValueUint24 returns ValueUint24 (property field)
	GetValueUint24() *uint32
	// GetValueUint32 returns ValueUint32 (property field)
	GetValueUint32() *uint32
	// GetValueUint40 returns ValueUint40 (property field)
	GetValueUint40() *uint64
	// GetValueUint48 returns ValueUint48 (property field)
	GetValueUint48() *uint64
	// GetValueUint56 returns ValueUint56 (property field)
	GetValueUint56() *uint64
	// GetValueUint64 returns ValueUint64 (property field)
	GetValueUint64() *uint64
	// GetIsUint8 returns IsUint8 (virtual field)
	GetIsUint8() bool
	// GetIsUint16 returns IsUint16 (virtual field)
	GetIsUint16() bool
	// GetIsUint24 returns IsUint24 (virtual field)
	GetIsUint24() bool
	// GetIsUint32 returns IsUint32 (virtual field)
	GetIsUint32() bool
	// GetIsUint40 returns IsUint40 (virtual field)
	GetIsUint40() bool
	// GetIsUint48 returns IsUint48 (virtual field)
	GetIsUint48() bool
	// GetIsUint56 returns IsUint56 (virtual field)
	GetIsUint56() bool
	// GetIsUint64 returns IsUint64 (virtual field)
	GetIsUint64() bool
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
}

// BACnetTagPayloadUnsignedIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetTagPayloadUnsignedInteger.
// This is useful for switch cases.
type BACnetTagPayloadUnsignedIntegerExactly interface {
	BACnetTagPayloadUnsignedInteger
	isBACnetTagPayloadUnsignedInteger() bool
}

// _BACnetTagPayloadUnsignedInteger is the data-structure of this message
type _BACnetTagPayloadUnsignedInteger struct {
        ValueUint8 *uint8
        ValueUint16 *uint16
        ValueUint24 *uint32
        ValueUint32 *uint32
        ValueUint40 *uint64
        ValueUint48 *uint64
        ValueUint56 *uint64
        ValueUint64 *uint64

	// Arguments.
	ActualLength uint32
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint8() *uint8 {
	return m.ValueUint8
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint16() *uint16 {
	return m.ValueUint16
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint24() *uint32 {
	return m.ValueUint24
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint32() *uint32 {
	return m.ValueUint32
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint40() *uint64 {
	return m.ValueUint40
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint48() *uint64 {
	return m.ValueUint48
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint56() *uint64 {
	return m.ValueUint56
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint64() *uint64 {
	return m.ValueUint64
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint8() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	valueUint40 := m.ValueUint40
	_ = valueUint40
	valueUint48 := m.ValueUint48
	_ = valueUint48
	valueUint56 := m.ValueUint56
	_ = valueUint56
	valueUint64 := m.ValueUint64
	_ = valueUint64
	return bool(bool((m.ActualLength) == ((1))))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint16() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	valueUint40 := m.ValueUint40
	_ = valueUint40
	valueUint48 := m.ValueUint48
	_ = valueUint48
	valueUint56 := m.ValueUint56
	_ = valueUint56
	valueUint64 := m.ValueUint64
	_ = valueUint64
	return bool(bool((m.ActualLength) == ((2))))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint24() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	valueUint40 := m.ValueUint40
	_ = valueUint40
	valueUint48 := m.ValueUint48
	_ = valueUint48
	valueUint56 := m.ValueUint56
	_ = valueUint56
	valueUint64 := m.ValueUint64
	_ = valueUint64
	return bool(bool((m.ActualLength) == ((3))))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint32() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	valueUint40 := m.ValueUint40
	_ = valueUint40
	valueUint48 := m.ValueUint48
	_ = valueUint48
	valueUint56 := m.ValueUint56
	_ = valueUint56
	valueUint64 := m.ValueUint64
	_ = valueUint64
	return bool(bool((m.ActualLength) == ((4))))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint40() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	valueUint40 := m.ValueUint40
	_ = valueUint40
	valueUint48 := m.ValueUint48
	_ = valueUint48
	valueUint56 := m.ValueUint56
	_ = valueUint56
	valueUint64 := m.ValueUint64
	_ = valueUint64
	return bool(bool((m.ActualLength) == ((5))))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint48() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	valueUint40 := m.ValueUint40
	_ = valueUint40
	valueUint48 := m.ValueUint48
	_ = valueUint48
	valueUint56 := m.ValueUint56
	_ = valueUint56
	valueUint64 := m.ValueUint64
	_ = valueUint64
	return bool(bool((m.ActualLength) == ((6))))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint56() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	valueUint40 := m.ValueUint40
	_ = valueUint40
	valueUint48 := m.ValueUint48
	_ = valueUint48
	valueUint56 := m.ValueUint56
	_ = valueUint56
	valueUint64 := m.ValueUint64
	_ = valueUint64
	return bool(bool((m.ActualLength) == ((7))))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint64() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	valueUint40 := m.ValueUint40
	_ = valueUint40
	valueUint48 := m.ValueUint48
	_ = valueUint48
	valueUint56 := m.ValueUint56
	_ = valueUint56
	valueUint64 := m.ValueUint64
	_ = valueUint64
	return bool(bool((m.ActualLength) == ((8))))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetActualValue() uint64 {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.ValueUint8
	_ = valueUint8
	valueUint16 := m.ValueUint16
	_ = valueUint16
	valueUint24 := m.ValueUint24
	_ = valueUint24
	valueUint32 := m.ValueUint32
	_ = valueUint32
	valueUint40 := m.ValueUint40
	_ = valueUint40
	valueUint48 := m.ValueUint48
	_ = valueUint48
	valueUint56 := m.ValueUint56
	_ = valueUint56
	valueUint64 := m.ValueUint64
	_ = valueUint64
	return uint64(utils.InlineIf(m.GetIsUint8(), func() any {return uint64((*m.GetValueUint8()))}, func() any {return uint64((utils.InlineIf(m.GetIsUint16(), func() any {return uint64((*m.GetValueUint16()))}, func() any {return uint64((utils.InlineIf(m.GetIsUint24(), func() any {return uint64((*m.GetValueUint24()))}, func() any {return uint64((utils.InlineIf(m.GetIsUint32(), func() any {return uint64((*m.GetValueUint32()))}, func() any {return uint64((utils.InlineIf(m.GetIsUint40(), func() any {return uint64((*m.GetValueUint40()))}, func() any {return uint64((utils.InlineIf(m.GetIsUint48(), func() any {return uint64((*m.GetValueUint48()))}, func() any {return uint64((utils.InlineIf(m.GetIsUint56(), func() any {return uint64((*m.GetValueUint56()))}, func() any {return uint64((*m.GetValueUint64()))}).(uint64)))}).(uint64)))}).(uint64)))}).(uint64)))}).(uint64)))}).(uint64)))}).(uint64))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetTagPayloadUnsignedInteger factory function for _BACnetTagPayloadUnsignedInteger
func NewBACnetTagPayloadUnsignedInteger( valueUint8 *uint8 , valueUint16 *uint16 , valueUint24 *uint32 , valueUint32 *uint32 , valueUint40 *uint64 , valueUint48 *uint64 , valueUint56 *uint64 , valueUint64 *uint64 , actualLength uint32 ) *_BACnetTagPayloadUnsignedInteger {
return &_BACnetTagPayloadUnsignedInteger{ ValueUint8: valueUint8 , ValueUint16: valueUint16 , ValueUint24: valueUint24 , ValueUint32: valueUint32 , ValueUint40: valueUint40 , ValueUint48: valueUint48 , ValueUint56: valueUint56 , ValueUint64: valueUint64 , ActualLength: actualLength }
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadUnsignedInteger(structType any) BACnetTagPayloadUnsignedInteger {
    if casted, ok := structType.(BACnetTagPayloadUnsignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadUnsignedInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadUnsignedInteger) GetTypeName() string {
	return "BACnetTagPayloadUnsignedInteger"
}

func (m *_BACnetTagPayloadUnsignedInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

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

	// Optional Field (valueUint40)
	if m.ValueUint40 != nil {
		lengthInBits += 40
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint48)
	if m.ValueUint48 != nil {
		lengthInBits += 48
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint56)
	if m.ValueUint56 != nil {
		lengthInBits += 56
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint64)
	if m.ValueUint64 != nil {
		lengthInBits += 64
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetTagPayloadUnsignedInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadUnsignedIntegerParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetTagPayloadUnsignedInteger, error) {
	return BACnetTagPayloadUnsignedIntegerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetTagPayloadUnsignedIntegerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadUnsignedInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetTagPayloadUnsignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadUnsignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_isUint8 := bool((actualLength) == ((1)))
	isUint8 := bool(_isUint8)
	_ = isUint8

	// Optional Field (valueUint8) (Can be skipped, if a given expression evaluates to false)
	var valueUint8 *uint8 = nil
	if isUint8 {
		_val, _err := readBuffer.ReadUint8("valueUint8", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint8' field of BACnetTagPayloadUnsignedInteger")
		}
		valueUint8 = &_val
	}

	// Virtual field
	_isUint16 := bool((actualLength) == ((2)))
	isUint16 := bool(_isUint16)
	_ = isUint16

	// Optional Field (valueUint16) (Can be skipped, if a given expression evaluates to false)
	var valueUint16 *uint16 = nil
	if isUint16 {
		_val, _err := readBuffer.ReadUint16("valueUint16", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint16' field of BACnetTagPayloadUnsignedInteger")
		}
		valueUint16 = &_val
	}

	// Virtual field
	_isUint24 := bool((actualLength) == ((3)))
	isUint24 := bool(_isUint24)
	_ = isUint24

	// Optional Field (valueUint24) (Can be skipped, if a given expression evaluates to false)
	var valueUint24 *uint32 = nil
	if isUint24 {
		_val, _err := readBuffer.ReadUint32("valueUint24", 24)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint24' field of BACnetTagPayloadUnsignedInteger")
		}
		valueUint24 = &_val
	}

	// Virtual field
	_isUint32 := bool((actualLength) == ((4)))
	isUint32 := bool(_isUint32)
	_ = isUint32

	// Optional Field (valueUint32) (Can be skipped, if a given expression evaluates to false)
	var valueUint32 *uint32 = nil
	if isUint32 {
		_val, _err := readBuffer.ReadUint32("valueUint32", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint32' field of BACnetTagPayloadUnsignedInteger")
		}
		valueUint32 = &_val
	}

	// Virtual field
	_isUint40 := bool((actualLength) == ((5)))
	isUint40 := bool(_isUint40)
	_ = isUint40

	// Optional Field (valueUint40) (Can be skipped, if a given expression evaluates to false)
	var valueUint40 *uint64 = nil
	if isUint40 {
		_val, _err := readBuffer.ReadUint64("valueUint40", 40)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint40' field of BACnetTagPayloadUnsignedInteger")
		}
		valueUint40 = &_val
	}

	// Virtual field
	_isUint48 := bool((actualLength) == ((6)))
	isUint48 := bool(_isUint48)
	_ = isUint48

	// Optional Field (valueUint48) (Can be skipped, if a given expression evaluates to false)
	var valueUint48 *uint64 = nil
	if isUint48 {
		_val, _err := readBuffer.ReadUint64("valueUint48", 48)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint48' field of BACnetTagPayloadUnsignedInteger")
		}
		valueUint48 = &_val
	}

	// Virtual field
	_isUint56 := bool((actualLength) == ((7)))
	isUint56 := bool(_isUint56)
	_ = isUint56

	// Optional Field (valueUint56) (Can be skipped, if a given expression evaluates to false)
	var valueUint56 *uint64 = nil
	if isUint56 {
		_val, _err := readBuffer.ReadUint64("valueUint56", 56)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint56' field of BACnetTagPayloadUnsignedInteger")
		}
		valueUint56 = &_val
	}

	// Virtual field
	_isUint64 := bool((actualLength) == ((8)))
	isUint64 := bool(_isUint64)
	_ = isUint64

	// Optional Field (valueUint64) (Can be skipped, if a given expression evaluates to false)
	var valueUint64 *uint64 = nil
	if isUint64 {
		_val, _err := readBuffer.ReadUint64("valueUint64", 64)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint64' field of BACnetTagPayloadUnsignedInteger")
		}
		valueUint64 = &_val
	}

	// Validation
	if (!(bool(bool(bool(bool(bool(bool(bool(isUint8) || bool(isUint16)) || bool(isUint24)) || bool(isUint32)) || bool(isUint40)) || bool(isUint48)) || bool(isUint56)) || bool(isUint64))) {
		return nil, errors.WithStack(utils.ParseValidationError{"unmapped integer length"})
	}

	// Virtual field
	_actualValue := utils.InlineIf(isUint8, func() any {return uint64((*valueUint8))}, func() any {return uint64((utils.InlineIf(isUint16, func() any {return uint64((*valueUint16))}, func() any {return uint64((utils.InlineIf(isUint24, func() any {return uint64((*valueUint24))}, func() any {return uint64((utils.InlineIf(isUint32, func() any {return uint64((*valueUint32))}, func() any {return uint64((utils.InlineIf(isUint40, func() any {return uint64((*valueUint40))}, func() any {return uint64((utils.InlineIf(isUint48, func() any {return uint64((*valueUint48))}, func() any {return uint64((utils.InlineIf(isUint56, func() any {return uint64((*valueUint56))}, func() any {return uint64((*valueUint64))}).(uint64)))}).(uint64)))}).(uint64)))}).(uint64)))}).(uint64)))}).(uint64)))}).(uint64)
	actualValue := uint64(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadUnsignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadUnsignedInteger")
	}

	// Create the instance
	return &_BACnetTagPayloadUnsignedInteger{
            ActualLength: actualLength,
			ValueUint8: valueUint8,
			ValueUint16: valueUint16,
			ValueUint24: valueUint24,
			ValueUint32: valueUint32,
			ValueUint40: valueUint40,
			ValueUint48: valueUint48,
			ValueUint56: valueUint56,
			ValueUint64: valueUint64,
		}, nil
}

func (m *_BACnetTagPayloadUnsignedInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadUnsignedInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("BACnetTagPayloadUnsignedInteger"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadUnsignedInteger")
	}
	// Virtual field
	if _isUint8Err := writeBuffer.WriteVirtual(ctx, "isUint8", m.GetIsUint8()); _isUint8Err != nil {
		return errors.Wrap(_isUint8Err, "Error serializing 'isUint8' field")
	}

	// Optional Field (valueUint8) (Can be skipped, if the value is null)
	var valueUint8 *uint8 = nil
	if m.GetValueUint8() != nil {
		valueUint8 = m.GetValueUint8()
		_valueUint8Err := writeBuffer.WriteUint8("valueUint8", 8, *(valueUint8))
		if _valueUint8Err != nil {
			return errors.Wrap(_valueUint8Err, "Error serializing 'valueUint8' field")
		}
	}
	// Virtual field
	if _isUint16Err := writeBuffer.WriteVirtual(ctx, "isUint16", m.GetIsUint16()); _isUint16Err != nil {
		return errors.Wrap(_isUint16Err, "Error serializing 'isUint16' field")
	}

	// Optional Field (valueUint16) (Can be skipped, if the value is null)
	var valueUint16 *uint16 = nil
	if m.GetValueUint16() != nil {
		valueUint16 = m.GetValueUint16()
		_valueUint16Err := writeBuffer.WriteUint16("valueUint16", 16, *(valueUint16))
		if _valueUint16Err != nil {
			return errors.Wrap(_valueUint16Err, "Error serializing 'valueUint16' field")
		}
	}
	// Virtual field
	if _isUint24Err := writeBuffer.WriteVirtual(ctx, "isUint24", m.GetIsUint24()); _isUint24Err != nil {
		return errors.Wrap(_isUint24Err, "Error serializing 'isUint24' field")
	}

	// Optional Field (valueUint24) (Can be skipped, if the value is null)
	var valueUint24 *uint32 = nil
	if m.GetValueUint24() != nil {
		valueUint24 = m.GetValueUint24()
		_valueUint24Err := writeBuffer.WriteUint32("valueUint24", 24, *(valueUint24))
		if _valueUint24Err != nil {
			return errors.Wrap(_valueUint24Err, "Error serializing 'valueUint24' field")
		}
	}
	// Virtual field
	if _isUint32Err := writeBuffer.WriteVirtual(ctx, "isUint32", m.GetIsUint32()); _isUint32Err != nil {
		return errors.Wrap(_isUint32Err, "Error serializing 'isUint32' field")
	}

	// Optional Field (valueUint32) (Can be skipped, if the value is null)
	var valueUint32 *uint32 = nil
	if m.GetValueUint32() != nil {
		valueUint32 = m.GetValueUint32()
		_valueUint32Err := writeBuffer.WriteUint32("valueUint32", 32, *(valueUint32))
		if _valueUint32Err != nil {
			return errors.Wrap(_valueUint32Err, "Error serializing 'valueUint32' field")
		}
	}
	// Virtual field
	if _isUint40Err := writeBuffer.WriteVirtual(ctx, "isUint40", m.GetIsUint40()); _isUint40Err != nil {
		return errors.Wrap(_isUint40Err, "Error serializing 'isUint40' field")
	}

	// Optional Field (valueUint40) (Can be skipped, if the value is null)
	var valueUint40 *uint64 = nil
	if m.GetValueUint40() != nil {
		valueUint40 = m.GetValueUint40()
		_valueUint40Err := writeBuffer.WriteUint64("valueUint40", 40, *(valueUint40))
		if _valueUint40Err != nil {
			return errors.Wrap(_valueUint40Err, "Error serializing 'valueUint40' field")
		}
	}
	// Virtual field
	if _isUint48Err := writeBuffer.WriteVirtual(ctx, "isUint48", m.GetIsUint48()); _isUint48Err != nil {
		return errors.Wrap(_isUint48Err, "Error serializing 'isUint48' field")
	}

	// Optional Field (valueUint48) (Can be skipped, if the value is null)
	var valueUint48 *uint64 = nil
	if m.GetValueUint48() != nil {
		valueUint48 = m.GetValueUint48()
		_valueUint48Err := writeBuffer.WriteUint64("valueUint48", 48, *(valueUint48))
		if _valueUint48Err != nil {
			return errors.Wrap(_valueUint48Err, "Error serializing 'valueUint48' field")
		}
	}
	// Virtual field
	if _isUint56Err := writeBuffer.WriteVirtual(ctx, "isUint56", m.GetIsUint56()); _isUint56Err != nil {
		return errors.Wrap(_isUint56Err, "Error serializing 'isUint56' field")
	}

	// Optional Field (valueUint56) (Can be skipped, if the value is null)
	var valueUint56 *uint64 = nil
	if m.GetValueUint56() != nil {
		valueUint56 = m.GetValueUint56()
		_valueUint56Err := writeBuffer.WriteUint64("valueUint56", 56, *(valueUint56))
		if _valueUint56Err != nil {
			return errors.Wrap(_valueUint56Err, "Error serializing 'valueUint56' field")
		}
	}
	// Virtual field
	if _isUint64Err := writeBuffer.WriteVirtual(ctx, "isUint64", m.GetIsUint64()); _isUint64Err != nil {
		return errors.Wrap(_isUint64Err, "Error serializing 'isUint64' field")
	}

	// Optional Field (valueUint64) (Can be skipped, if the value is null)
	var valueUint64 *uint64 = nil
	if m.GetValueUint64() != nil {
		valueUint64 = m.GetValueUint64()
		_valueUint64Err := writeBuffer.WriteUint64("valueUint64", 64, *(valueUint64))
		if _valueUint64Err != nil {
			return errors.Wrap(_valueUint64Err, "Error serializing 'valueUint64' field")
		}
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadUnsignedInteger"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadUnsignedInteger")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetTagPayloadUnsignedInteger) GetActualLength() uint32 {
	return m.ActualLength
}
//
////

func (m *_BACnetTagPayloadUnsignedInteger) isBACnetTagPayloadUnsignedInteger() bool {
	return true
}

func (m *_BACnetTagPayloadUnsignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



