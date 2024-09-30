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

// VariantUInt16 is the corresponding interface of VariantUInt16
type VariantUInt16 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []uint16
	// IsVariantUInt16 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantUInt16()
	// CreateBuilder creates a VariantUInt16Builder
	CreateVariantUInt16Builder() VariantUInt16Builder
}

// _VariantUInt16 is the data-structure of this message
type _VariantUInt16 struct {
	VariantContract
	ArrayLength *int32
	Value       []uint16
}

var _ VariantUInt16 = (*_VariantUInt16)(nil)
var _ VariantRequirements = (*_VariantUInt16)(nil)

// NewVariantUInt16 factory function for _VariantUInt16
func NewVariantUInt16(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []uint16) *_VariantUInt16 {
	_result := &_VariantUInt16{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VariantUInt16Builder is a builder for VariantUInt16
type VariantUInt16Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []uint16) VariantUInt16Builder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantUInt16Builder
	// WithValue adds Value (property field)
	WithValue(...uint16) VariantUInt16Builder
	// Build builds the VariantUInt16 or returns an error if something is wrong
	Build() (VariantUInt16, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantUInt16
}

// NewVariantUInt16Builder() creates a VariantUInt16Builder
func NewVariantUInt16Builder() VariantUInt16Builder {
	return &_VariantUInt16Builder{_VariantUInt16: new(_VariantUInt16)}
}

type _VariantUInt16Builder struct {
	*_VariantUInt16

	parentBuilder *_VariantBuilder

	err *utils.MultiError
}

var _ (VariantUInt16Builder) = (*_VariantUInt16Builder)(nil)

func (b *_VariantUInt16Builder) setParent(contract VariantContract) {
	b.VariantContract = contract
}

func (b *_VariantUInt16Builder) WithMandatoryFields(value []uint16) VariantUInt16Builder {
	return b.WithValue(value...)
}

func (b *_VariantUInt16Builder) WithOptionalArrayLength(arrayLength int32) VariantUInt16Builder {
	b.ArrayLength = &arrayLength
	return b
}

func (b *_VariantUInt16Builder) WithValue(value ...uint16) VariantUInt16Builder {
	b.Value = value
	return b
}

func (b *_VariantUInt16Builder) Build() (VariantUInt16, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._VariantUInt16.deepCopy(), nil
}

func (b *_VariantUInt16Builder) MustBuild() VariantUInt16 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_VariantUInt16Builder) Done() VariantBuilder {
	return b.parentBuilder
}

func (b *_VariantUInt16Builder) buildForVariant() (Variant, error) {
	return b.Build()
}

func (b *_VariantUInt16Builder) DeepCopy() any {
	_copy := b.CreateVariantUInt16Builder().(*_VariantUInt16Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateVariantUInt16Builder creates a VariantUInt16Builder
func (b *_VariantUInt16) CreateVariantUInt16Builder() VariantUInt16Builder {
	if b == nil {
		return NewVariantUInt16Builder()
	}
	return &_VariantUInt16Builder{_VariantUInt16: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantUInt16) GetVariantType() uint8 {
	return uint8(5)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantUInt16) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantUInt16) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantUInt16) GetValue() []uint16 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantUInt16(structType any) VariantUInt16 {
	if casted, ok := structType.(VariantUInt16); ok {
		return casted
	}
	if casted, ok := structType.(*VariantUInt16); ok {
		return *casted
	}
	return nil
}

func (m *_VariantUInt16) GetTypeName() string {
	return "VariantUInt16"
}

func (m *_VariantUInt16) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 16 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantUInt16) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantUInt16) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantUInt16 VariantUInt16, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantUInt16"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantUInt16")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[uint16](ctx, "value", ReadUnsignedShort(readBuffer, uint8(16)), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantUInt16"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantUInt16")
	}

	return m, nil
}

func (m *_VariantUInt16) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantUInt16) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantUInt16"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantUInt16")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "value", m.GetValue(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantUInt16"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantUInt16")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantUInt16) IsVariantUInt16() {}

func (m *_VariantUInt16) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantUInt16) deepCopy() *_VariantUInt16 {
	if m == nil {
		return nil
	}
	_VariantUInt16Copy := &_VariantUInt16{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[uint16, uint16](m.Value),
	}
	m.VariantContract.(*_Variant)._SubType = m
	return _VariantUInt16Copy
}

func (m *_VariantUInt16) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
