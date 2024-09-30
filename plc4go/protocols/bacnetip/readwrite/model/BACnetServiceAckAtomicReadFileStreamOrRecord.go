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

// BACnetServiceAckAtomicReadFileStreamOrRecord is the corresponding interface of BACnetServiceAckAtomicReadFileStreamOrRecord
type BACnetServiceAckAtomicReadFileStreamOrRecord interface {
	BACnetServiceAckAtomicReadFileStreamOrRecordContract
	BACnetServiceAckAtomicReadFileStreamOrRecordRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetServiceAckAtomicReadFileStreamOrRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckAtomicReadFileStreamOrRecord()
	// CreateBuilder creates a BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	CreateBACnetServiceAckAtomicReadFileStreamOrRecordBuilder() BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
}

// BACnetServiceAckAtomicReadFileStreamOrRecordContract provides a set of functions which can be overwritten by a sub struct
type BACnetServiceAckAtomicReadFileStreamOrRecordContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetServiceAckAtomicReadFileStreamOrRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckAtomicReadFileStreamOrRecord()
	// CreateBuilder creates a BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	CreateBACnetServiceAckAtomicReadFileStreamOrRecordBuilder() BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
}

// BACnetServiceAckAtomicReadFileStreamOrRecordRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetServiceAckAtomicReadFileStreamOrRecordRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetServiceAckAtomicReadFileStreamOrRecord is the data-structure of this message
type _BACnetServiceAckAtomicReadFileStreamOrRecord struct {
	_SubType        BACnetServiceAckAtomicReadFileStreamOrRecord
	PeekedTagHeader BACnetTagHeader
	OpeningTag      BACnetOpeningTag
	ClosingTag      BACnetClosingTag
}

var _ BACnetServiceAckAtomicReadFileStreamOrRecordContract = (*_BACnetServiceAckAtomicReadFileStreamOrRecord)(nil)

// NewBACnetServiceAckAtomicReadFileStreamOrRecord factory function for _BACnetServiceAckAtomicReadFileStreamOrRecord
func NewBACnetServiceAckAtomicReadFileStreamOrRecord(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetServiceAckAtomicReadFileStreamOrRecord {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetServiceAckAtomicReadFileStreamOrRecord must not be nil")
	}
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetServiceAckAtomicReadFileStreamOrRecord must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetServiceAckAtomicReadFileStreamOrRecord must not be nil")
	}
	return &_BACnetServiceAckAtomicReadFileStreamOrRecord{PeekedTagHeader: peekedTagHeader, OpeningTag: openingTag, ClosingTag: closingTag}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetServiceAckAtomicReadFileStreamOrRecordBuilder is a builder for BACnetServiceAckAtomicReadFileStreamOrRecord
type BACnetServiceAckAtomicReadFileStreamOrRecordBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	// AsBACnetServiceAckAtomicReadFileStream converts this build to a subType of BACnetServiceAckAtomicReadFileStreamOrRecord. It is always possible to return to current builder using Done()
	AsBACnetServiceAckAtomicReadFileStream() interface {
		BACnetServiceAckAtomicReadFileStreamBuilder
		Done() BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	}
	// AsBACnetServiceAckAtomicReadFileRecord converts this build to a subType of BACnetServiceAckAtomicReadFileStreamOrRecord. It is always possible to return to current builder using Done()
	AsBACnetServiceAckAtomicReadFileRecord() interface {
		BACnetServiceAckAtomicReadFileRecordBuilder
		Done() BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	}
	// Build builds the BACnetServiceAckAtomicReadFileStreamOrRecord or returns an error if something is wrong
	PartialBuild() (BACnetServiceAckAtomicReadFileStreamOrRecordContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetServiceAckAtomicReadFileStreamOrRecordContract
	// Build builds the BACnetServiceAckAtomicReadFileStreamOrRecord or returns an error if something is wrong
	Build() (BACnetServiceAckAtomicReadFileStreamOrRecord, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetServiceAckAtomicReadFileStreamOrRecord
}

// NewBACnetServiceAckAtomicReadFileStreamOrRecordBuilder() creates a BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
func NewBACnetServiceAckAtomicReadFileStreamOrRecordBuilder() BACnetServiceAckAtomicReadFileStreamOrRecordBuilder {
	return &_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder{_BACnetServiceAckAtomicReadFileStreamOrRecord: new(_BACnetServiceAckAtomicReadFileStreamOrRecord)}
}

type _BACnetServiceAckAtomicReadFileStreamOrRecordChildBuilder interface {
	utils.Copyable
	setParent(BACnetServiceAckAtomicReadFileStreamOrRecordContract)
	buildForBACnetServiceAckAtomicReadFileStreamOrRecord() (BACnetServiceAckAtomicReadFileStreamOrRecord, error)
}

type _BACnetServiceAckAtomicReadFileStreamOrRecordBuilder struct {
	*_BACnetServiceAckAtomicReadFileStreamOrRecord

	childBuilder _BACnetServiceAckAtomicReadFileStreamOrRecordChildBuilder

	err *utils.MultiError
}

var _ (BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) = (*_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder)(nil)

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) WithMandatoryFields(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder {
	return b.WithPeekedTagHeader(peekedTagHeader).WithOpeningTag(openingTag).WithClosingTag(closingTag)
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder {
	builder := builderSupplier(b.PeekedTagHeader.CreateBACnetTagHeaderBuilder())
	var err error
	b.PeekedTagHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetServiceAckAtomicReadFileStreamOrRecordBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) PartialBuild() (BACnetServiceAckAtomicReadFileStreamOrRecordContract, error) {
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetServiceAckAtomicReadFileStreamOrRecord.deepCopy(), nil
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) PartialMustBuild() BACnetServiceAckAtomicReadFileStreamOrRecordContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) AsBACnetServiceAckAtomicReadFileStream() interface {
	BACnetServiceAckAtomicReadFileStreamBuilder
	Done() BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetServiceAckAtomicReadFileStreamBuilder
		Done() BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetServiceAckAtomicReadFileStreamBuilder().(*_BACnetServiceAckAtomicReadFileStreamBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) AsBACnetServiceAckAtomicReadFileRecord() interface {
	BACnetServiceAckAtomicReadFileRecordBuilder
	Done() BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetServiceAckAtomicReadFileRecordBuilder
		Done() BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetServiceAckAtomicReadFileRecordBuilder().(*_BACnetServiceAckAtomicReadFileRecordBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) Build() (BACnetServiceAckAtomicReadFileStreamOrRecord, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetServiceAckAtomicReadFileStreamOrRecord()
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) MustBuild() BACnetServiceAckAtomicReadFileStreamOrRecord {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder) DeepCopy() any {
	_copy := b.CreateBACnetServiceAckAtomicReadFileStreamOrRecordBuilder().(*_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetServiceAckAtomicReadFileStreamOrRecordChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetServiceAckAtomicReadFileStreamOrRecordBuilder creates a BACnetServiceAckAtomicReadFileStreamOrRecordBuilder
func (b *_BACnetServiceAckAtomicReadFileStreamOrRecord) CreateBACnetServiceAckAtomicReadFileStreamOrRecordBuilder() BACnetServiceAckAtomicReadFileStreamOrRecordBuilder {
	if b == nil {
		return NewBACnetServiceAckAtomicReadFileStreamOrRecordBuilder()
	}
	return &_BACnetServiceAckAtomicReadFileStreamOrRecordBuilder{_BACnetServiceAckAtomicReadFileStreamOrRecord: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicReadFileStreamOrRecord(structType any) BACnetServiceAckAtomicReadFileStreamOrRecord {
	if casted, ok := structType.(BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFileStreamOrRecord"
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckAtomicReadFileStreamOrRecordParse[T BACnetServiceAckAtomicReadFileStreamOrRecord](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBufferProducer[T BACnetServiceAckAtomicReadFileStreamOrRecord]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBuffer[T BACnetServiceAckAtomicReadFileStreamOrRecord](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetServiceAckAtomicReadFileStreamOrRecord{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetServiceAckAtomicReadFileStreamOrRecord BACnetServiceAckAtomicReadFileStreamOrRecord, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetServiceAckAtomicReadFileStreamOrRecord
	switch {
	case peekedTagNumber == 0x0: // BACnetServiceAckAtomicReadFileStream
		if _child, err = new(_BACnetServiceAckAtomicReadFileStream).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckAtomicReadFileStream for type-switch of BACnetServiceAckAtomicReadFileStreamOrRecord")
		}
	case peekedTagNumber == 0x1: // BACnetServiceAckAtomicReadFileRecord
		if _child, err = new(_BACnetServiceAckAtomicReadFileRecord).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckAtomicReadFileRecord for type-switch of BACnetServiceAckAtomicReadFileStreamOrRecord")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}

	return _child, nil
}

func (pm *_BACnetServiceAckAtomicReadFileStreamOrRecord) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetServiceAckAtomicReadFileStreamOrRecord, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) IsBACnetServiceAckAtomicReadFileStreamOrRecord() {
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) deepCopy() *_BACnetServiceAckAtomicReadFileStreamOrRecord {
	if m == nil {
		return nil
	}
	_BACnetServiceAckAtomicReadFileStreamOrRecordCopy := &_BACnetServiceAckAtomicReadFileStreamOrRecord{
		nil, // will be set by child
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	return _BACnetServiceAckAtomicReadFileStreamOrRecordCopy
}
