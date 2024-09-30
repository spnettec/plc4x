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

// BACnetAccessZoneOccupancyStateTagged is the corresponding interface of BACnetAccessZoneOccupancyStateTagged
type BACnetAccessZoneOccupancyStateTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessZoneOccupancyState
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetAccessZoneOccupancyStateTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccessZoneOccupancyStateTagged()
	// CreateBuilder creates a BACnetAccessZoneOccupancyStateTaggedBuilder
	CreateBACnetAccessZoneOccupancyStateTaggedBuilder() BACnetAccessZoneOccupancyStateTaggedBuilder
}

// _BACnetAccessZoneOccupancyStateTagged is the data-structure of this message
type _BACnetAccessZoneOccupancyStateTagged struct {
	Header           BACnetTagHeader
	Value            BACnetAccessZoneOccupancyState
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAccessZoneOccupancyStateTagged = (*_BACnetAccessZoneOccupancyStateTagged)(nil)

// NewBACnetAccessZoneOccupancyStateTagged factory function for _BACnetAccessZoneOccupancyStateTagged
func NewBACnetAccessZoneOccupancyStateTagged(header BACnetTagHeader, value BACnetAccessZoneOccupancyState, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetAccessZoneOccupancyStateTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAccessZoneOccupancyStateTagged must not be nil")
	}
	return &_BACnetAccessZoneOccupancyStateTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAccessZoneOccupancyStateTaggedBuilder is a builder for BACnetAccessZoneOccupancyStateTagged
type BACnetAccessZoneOccupancyStateTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAccessZoneOccupancyState, proprietaryValue uint32) BACnetAccessZoneOccupancyStateTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAccessZoneOccupancyStateTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessZoneOccupancyStateTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAccessZoneOccupancyState) BACnetAccessZoneOccupancyStateTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetAccessZoneOccupancyStateTaggedBuilder
	// Build builds the BACnetAccessZoneOccupancyStateTagged or returns an error if something is wrong
	Build() (BACnetAccessZoneOccupancyStateTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAccessZoneOccupancyStateTagged
}

// NewBACnetAccessZoneOccupancyStateTaggedBuilder() creates a BACnetAccessZoneOccupancyStateTaggedBuilder
func NewBACnetAccessZoneOccupancyStateTaggedBuilder() BACnetAccessZoneOccupancyStateTaggedBuilder {
	return &_BACnetAccessZoneOccupancyStateTaggedBuilder{_BACnetAccessZoneOccupancyStateTagged: new(_BACnetAccessZoneOccupancyStateTagged)}
}

type _BACnetAccessZoneOccupancyStateTaggedBuilder struct {
	*_BACnetAccessZoneOccupancyStateTagged

	err *utils.MultiError
}

var _ (BACnetAccessZoneOccupancyStateTaggedBuilder) = (*_BACnetAccessZoneOccupancyStateTaggedBuilder)(nil)

func (b *_BACnetAccessZoneOccupancyStateTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAccessZoneOccupancyState, proprietaryValue uint32) BACnetAccessZoneOccupancyStateTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetAccessZoneOccupancyStateTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAccessZoneOccupancyStateTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAccessZoneOccupancyStateTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessZoneOccupancyStateTaggedBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetAccessZoneOccupancyStateTaggedBuilder) WithValue(value BACnetAccessZoneOccupancyState) BACnetAccessZoneOccupancyStateTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAccessZoneOccupancyStateTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetAccessZoneOccupancyStateTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetAccessZoneOccupancyStateTaggedBuilder) Build() (BACnetAccessZoneOccupancyStateTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAccessZoneOccupancyStateTagged.deepCopy(), nil
}

func (b *_BACnetAccessZoneOccupancyStateTaggedBuilder) MustBuild() BACnetAccessZoneOccupancyStateTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAccessZoneOccupancyStateTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAccessZoneOccupancyStateTaggedBuilder().(*_BACnetAccessZoneOccupancyStateTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAccessZoneOccupancyStateTaggedBuilder creates a BACnetAccessZoneOccupancyStateTaggedBuilder
func (b *_BACnetAccessZoneOccupancyStateTagged) CreateBACnetAccessZoneOccupancyStateTaggedBuilder() BACnetAccessZoneOccupancyStateTaggedBuilder {
	if b == nil {
		return NewBACnetAccessZoneOccupancyStateTaggedBuilder()
	}
	return &_BACnetAccessZoneOccupancyStateTaggedBuilder{_BACnetAccessZoneOccupancyStateTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessZoneOccupancyStateTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessZoneOccupancyStateTagged) GetValue() BACnetAccessZoneOccupancyState {
	return m.Value
}

func (m *_BACnetAccessZoneOccupancyStateTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetAccessZoneOccupancyStateTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAccessZoneOccupancyStateTagged(structType any) BACnetAccessZoneOccupancyStateTagged {
	if casted, ok := structType.(BACnetAccessZoneOccupancyStateTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessZoneOccupancyStateTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessZoneOccupancyStateTagged) GetTypeName() string {
	return "BACnetAccessZoneOccupancyStateTagged"
}

func (m *_BACnetAccessZoneOccupancyStateTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetAccessZoneOccupancyStateTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessZoneOccupancyStateTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccessZoneOccupancyStateTagged, error) {
	return BACnetAccessZoneOccupancyStateTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccessZoneOccupancyStateTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessZoneOccupancyStateTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessZoneOccupancyStateTagged, error) {
		return BACnetAccessZoneOccupancyStateTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAccessZoneOccupancyStateTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessZoneOccupancyStateTagged, error) {
	v, err := (&_BACnetAccessZoneOccupancyStateTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccessZoneOccupancyStateTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAccessZoneOccupancyStateTagged BACnetAccessZoneOccupancyStateTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessZoneOccupancyStateTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessZoneOccupancyStateTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetAccessZoneOccupancyState](ctx, "value", readBuffer, EnsureType[BACnetAccessZoneOccupancyState](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetAccessZoneOccupancyStateTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessZoneOccupancyStateTagged")
	}

	return m, nil
}

func (m *_BACnetAccessZoneOccupancyStateTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessZoneOccupancyStateTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessZoneOccupancyStateTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessZoneOccupancyStateTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAccessZoneOccupancyState](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessZoneOccupancyStateTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessZoneOccupancyStateTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccessZoneOccupancyStateTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessZoneOccupancyStateTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccessZoneOccupancyStateTagged) IsBACnetAccessZoneOccupancyStateTagged() {}

func (m *_BACnetAccessZoneOccupancyStateTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAccessZoneOccupancyStateTagged) deepCopy() *_BACnetAccessZoneOccupancyStateTagged {
	if m == nil {
		return nil
	}
	_BACnetAccessZoneOccupancyStateTaggedCopy := &_BACnetAccessZoneOccupancyStateTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAccessZoneOccupancyStateTaggedCopy
}

func (m *_BACnetAccessZoneOccupancyStateTagged) String() string {
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
