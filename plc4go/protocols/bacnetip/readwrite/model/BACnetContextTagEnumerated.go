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

// BACnetContextTagEnumerated is the corresponding interface of BACnetContextTagEnumerated
type BACnetContextTagEnumerated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadEnumerated
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint32
	// IsBACnetContextTagEnumerated is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagEnumerated()
	// CreateBuilder creates a BACnetContextTagEnumeratedBuilder
	CreateBACnetContextTagEnumeratedBuilder() BACnetContextTagEnumeratedBuilder
}

// _BACnetContextTagEnumerated is the data-structure of this message
type _BACnetContextTagEnumerated struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadEnumerated
}

var _ BACnetContextTagEnumerated = (*_BACnetContextTagEnumerated)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagEnumerated)(nil)

// NewBACnetContextTagEnumerated factory function for _BACnetContextTagEnumerated
func NewBACnetContextTagEnumerated(header BACnetTagHeader, payload BACnetTagPayloadEnumerated, tagNumberArgument uint8) *_BACnetContextTagEnumerated {
	if payload == nil {
		panic("payload of type BACnetTagPayloadEnumerated for BACnetContextTagEnumerated must not be nil")
	}
	_result := &_BACnetContextTagEnumerated{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetContextTagEnumeratedBuilder is a builder for BACnetContextTagEnumerated
type BACnetContextTagEnumeratedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadEnumerated) BACnetContextTagEnumeratedBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadEnumerated) BACnetContextTagEnumeratedBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadEnumeratedBuilder) BACnetTagPayloadEnumeratedBuilder) BACnetContextTagEnumeratedBuilder
	// Build builds the BACnetContextTagEnumerated or returns an error if something is wrong
	Build() (BACnetContextTagEnumerated, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagEnumerated
}

// NewBACnetContextTagEnumeratedBuilder() creates a BACnetContextTagEnumeratedBuilder
func NewBACnetContextTagEnumeratedBuilder() BACnetContextTagEnumeratedBuilder {
	return &_BACnetContextTagEnumeratedBuilder{_BACnetContextTagEnumerated: new(_BACnetContextTagEnumerated)}
}

type _BACnetContextTagEnumeratedBuilder struct {
	*_BACnetContextTagEnumerated

	parentBuilder *_BACnetContextTagBuilder

	err *utils.MultiError
}

var _ (BACnetContextTagEnumeratedBuilder) = (*_BACnetContextTagEnumeratedBuilder)(nil)

func (b *_BACnetContextTagEnumeratedBuilder) setParent(contract BACnetContextTagContract) {
	b.BACnetContextTagContract = contract
}

func (b *_BACnetContextTagEnumeratedBuilder) WithMandatoryFields(payload BACnetTagPayloadEnumerated) BACnetContextTagEnumeratedBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetContextTagEnumeratedBuilder) WithPayload(payload BACnetTagPayloadEnumerated) BACnetContextTagEnumeratedBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetContextTagEnumeratedBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadEnumeratedBuilder) BACnetTagPayloadEnumeratedBuilder) BACnetContextTagEnumeratedBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadEnumeratedBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadEnumeratedBuilder failed"))
	}
	return b
}

func (b *_BACnetContextTagEnumeratedBuilder) Build() (BACnetContextTagEnumerated, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetContextTagEnumerated.deepCopy(), nil
}

func (b *_BACnetContextTagEnumeratedBuilder) MustBuild() BACnetContextTagEnumerated {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetContextTagEnumeratedBuilder) Done() BACnetContextTagBuilder {
	return b.parentBuilder
}

func (b *_BACnetContextTagEnumeratedBuilder) buildForBACnetContextTag() (BACnetContextTag, error) {
	return b.Build()
}

func (b *_BACnetContextTagEnumeratedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetContextTagEnumeratedBuilder().(*_BACnetContextTagEnumeratedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetContextTagEnumeratedBuilder creates a BACnetContextTagEnumeratedBuilder
func (b *_BACnetContextTagEnumerated) CreateBACnetContextTagEnumeratedBuilder() BACnetContextTagEnumeratedBuilder {
	if b == nil {
		return NewBACnetContextTagEnumeratedBuilder()
	}
	return &_BACnetContextTagEnumeratedBuilder{_BACnetContextTagEnumerated: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagEnumerated) GetDataType() BACnetDataType {
	return BACnetDataType_ENUMERATED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagEnumerated) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagEnumerated) GetPayload() BACnetTagPayloadEnumerated {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTagEnumerated) GetActualValue() uint32 {
	ctx := context.Background()
	_ = ctx
	return uint32(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTagEnumerated(structType any) BACnetContextTagEnumerated {
	if casted, ok := structType.(BACnetContextTagEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagEnumerated) GetTypeName() string {
	return "BACnetContextTagEnumerated"
}

func (m *_BACnetContextTagEnumerated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagEnumerated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagEnumerated) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagEnumerated BACnetContextTagEnumerated, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadEnumerated](ctx, "payload", ReadComplex[BACnetTagPayloadEnumerated](BACnetTagPayloadEnumeratedParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[uint32](ctx, "actualValue", (*uint32)(nil), payload.GetActualValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagEnumerated")
	}

	return m, nil
}

func (m *_BACnetContextTagEnumerated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagEnumerated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagEnumerated")
		}

		if err := WriteSimpleField[BACnetTagPayloadEnumerated](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadEnumerated](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagEnumerated")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagEnumerated) IsBACnetContextTagEnumerated() {}

func (m *_BACnetContextTagEnumerated) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTagEnumerated) deepCopy() *_BACnetContextTagEnumerated {
	if m == nil {
		return nil
	}
	_BACnetContextTagEnumeratedCopy := &_BACnetContextTagEnumerated{
		m.BACnetContextTagContract.(*_BACnetContextTag).deepCopy(),
		m.Payload.DeepCopy().(BACnetTagPayloadEnumerated),
	}
	m.BACnetContextTagContract.(*_BACnetContextTag)._SubType = m
	return _BACnetContextTagEnumeratedCopy
}

func (m *_BACnetContextTagEnumerated) String() string {
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
