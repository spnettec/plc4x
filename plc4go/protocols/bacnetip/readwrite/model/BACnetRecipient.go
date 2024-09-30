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

// BACnetRecipient is the corresponding interface of BACnetRecipient
type BACnetRecipient interface {
	BACnetRecipientContract
	BACnetRecipientRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetRecipient is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRecipient()
	// CreateBuilder creates a BACnetRecipientBuilder
	CreateBACnetRecipientBuilder() BACnetRecipientBuilder
}

// BACnetRecipientContract provides a set of functions which can be overwritten by a sub struct
type BACnetRecipientContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetRecipient is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRecipient()
	// CreateBuilder creates a BACnetRecipientBuilder
	CreateBACnetRecipientBuilder() BACnetRecipientBuilder
}

// BACnetRecipientRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetRecipientRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetRecipient is the data-structure of this message
type _BACnetRecipient struct {
	_SubType        BACnetRecipient
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetRecipientContract = (*_BACnetRecipient)(nil)

// NewBACnetRecipient factory function for _BACnetRecipient
func NewBACnetRecipient(peekedTagHeader BACnetTagHeader) *_BACnetRecipient {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetRecipient must not be nil")
	}
	return &_BACnetRecipient{PeekedTagHeader: peekedTagHeader}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetRecipientBuilder is a builder for BACnetRecipient
type BACnetRecipientBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetRecipientBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetRecipientBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetRecipientBuilder
	// AsBACnetRecipientDevice converts this build to a subType of BACnetRecipient. It is always possible to return to current builder using Done()
	AsBACnetRecipientDevice() interface {
		BACnetRecipientDeviceBuilder
		Done() BACnetRecipientBuilder
	}
	// AsBACnetRecipientAddress converts this build to a subType of BACnetRecipient. It is always possible to return to current builder using Done()
	AsBACnetRecipientAddress() interface {
		BACnetRecipientAddressBuilder
		Done() BACnetRecipientBuilder
	}
	// Build builds the BACnetRecipient or returns an error if something is wrong
	PartialBuild() (BACnetRecipientContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetRecipientContract
	// Build builds the BACnetRecipient or returns an error if something is wrong
	Build() (BACnetRecipient, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetRecipient
}

// NewBACnetRecipientBuilder() creates a BACnetRecipientBuilder
func NewBACnetRecipientBuilder() BACnetRecipientBuilder {
	return &_BACnetRecipientBuilder{_BACnetRecipient: new(_BACnetRecipient)}
}

type _BACnetRecipientChildBuilder interface {
	utils.Copyable
	setParent(BACnetRecipientContract)
	buildForBACnetRecipient() (BACnetRecipient, error)
}

type _BACnetRecipientBuilder struct {
	*_BACnetRecipient

	childBuilder _BACnetRecipientChildBuilder

	err *utils.MultiError
}

var _ (BACnetRecipientBuilder) = (*_BACnetRecipientBuilder)(nil)

func (b *_BACnetRecipientBuilder) WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetRecipientBuilder {
	return b.WithPeekedTagHeader(peekedTagHeader)
}

func (b *_BACnetRecipientBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetRecipientBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetRecipientBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetRecipientBuilder {
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

func (b *_BACnetRecipientBuilder) PartialBuild() (BACnetRecipientContract, error) {
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetRecipient.deepCopy(), nil
}

func (b *_BACnetRecipientBuilder) PartialMustBuild() BACnetRecipientContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetRecipientBuilder) AsBACnetRecipientDevice() interface {
	BACnetRecipientDeviceBuilder
	Done() BACnetRecipientBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetRecipientDeviceBuilder
		Done() BACnetRecipientBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetRecipientDeviceBuilder().(*_BACnetRecipientDeviceBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetRecipientBuilder) AsBACnetRecipientAddress() interface {
	BACnetRecipientAddressBuilder
	Done() BACnetRecipientBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetRecipientAddressBuilder
		Done() BACnetRecipientBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetRecipientAddressBuilder().(*_BACnetRecipientAddressBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetRecipientBuilder) Build() (BACnetRecipient, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetRecipient()
}

func (b *_BACnetRecipientBuilder) MustBuild() BACnetRecipient {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetRecipientBuilder) DeepCopy() any {
	_copy := b.CreateBACnetRecipientBuilder().(*_BACnetRecipientBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetRecipientChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetRecipientBuilder creates a BACnetRecipientBuilder
func (b *_BACnetRecipient) CreateBACnetRecipientBuilder() BACnetRecipientBuilder {
	if b == nil {
		return NewBACnetRecipientBuilder()
	}
	return &_BACnetRecipientBuilder{_BACnetRecipient: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipient) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetRecipient) GetPeekedTagNumber() uint8 {
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
func CastBACnetRecipient(structType any) BACnetRecipient {
	if casted, ok := structType.(BACnetRecipient); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipient); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipient) GetTypeName() string {
	return "BACnetRecipient"
}

func (m *_BACnetRecipient) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetRecipient) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetRecipientParse[T BACnetRecipient](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetRecipientParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetRecipientParseWithBufferProducer[T BACnetRecipient]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetRecipientParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetRecipientParseWithBuffer[T BACnetRecipient](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetRecipient{}).parse(ctx, readBuffer)
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

func (m *_BACnetRecipient) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetRecipient BACnetRecipient, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipient"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipient")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetRecipient
	switch {
	case peekedTagNumber == uint8(0): // BACnetRecipientDevice
		if _child, err = new(_BACnetRecipientDevice).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetRecipientDevice for type-switch of BACnetRecipient")
		}
	case peekedTagNumber == uint8(1): // BACnetRecipientAddress
		if _child, err = new(_BACnetRecipientAddress).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetRecipientAddress for type-switch of BACnetRecipient")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipient"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipient")
	}

	return _child, nil
}

func (pm *_BACnetRecipient) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetRecipient, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRecipient"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRecipient")
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

	if popErr := writeBuffer.PopContext("BACnetRecipient"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRecipient")
	}
	return nil
}

func (m *_BACnetRecipient) IsBACnetRecipient() {}

func (m *_BACnetRecipient) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetRecipient) deepCopy() *_BACnetRecipient {
	if m == nil {
		return nil
	}
	_BACnetRecipientCopy := &_BACnetRecipient{
		nil, // will be set by child
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
	}
	return _BACnetRecipientCopy
}
