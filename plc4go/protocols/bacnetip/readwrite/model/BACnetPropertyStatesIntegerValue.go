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

// BACnetPropertyStatesIntegerValue is the corresponding interface of BACnetPropertyStatesIntegerValue
type BACnetPropertyStatesIntegerValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetContextTagSignedInteger
	// IsBACnetPropertyStatesIntegerValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesIntegerValue()
	// CreateBuilder creates a BACnetPropertyStatesIntegerValueBuilder
	CreateBACnetPropertyStatesIntegerValueBuilder() BACnetPropertyStatesIntegerValueBuilder
}

// _BACnetPropertyStatesIntegerValue is the data-structure of this message
type _BACnetPropertyStatesIntegerValue struct {
	BACnetPropertyStatesContract
	IntegerValue BACnetContextTagSignedInteger
}

var _ BACnetPropertyStatesIntegerValue = (*_BACnetPropertyStatesIntegerValue)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesIntegerValue)(nil)

// NewBACnetPropertyStatesIntegerValue factory function for _BACnetPropertyStatesIntegerValue
func NewBACnetPropertyStatesIntegerValue(peekedTagHeader BACnetTagHeader, integerValue BACnetContextTagSignedInteger) *_BACnetPropertyStatesIntegerValue {
	if integerValue == nil {
		panic("integerValue of type BACnetContextTagSignedInteger for BACnetPropertyStatesIntegerValue must not be nil")
	}
	_result := &_BACnetPropertyStatesIntegerValue{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		IntegerValue:                 integerValue,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesIntegerValueBuilder is a builder for BACnetPropertyStatesIntegerValue
type BACnetPropertyStatesIntegerValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(integerValue BACnetContextTagSignedInteger) BACnetPropertyStatesIntegerValueBuilder
	// WithIntegerValue adds IntegerValue (property field)
	WithIntegerValue(BACnetContextTagSignedInteger) BACnetPropertyStatesIntegerValueBuilder
	// WithIntegerValueBuilder adds IntegerValue (property field) which is build by the builder
	WithIntegerValueBuilder(func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetPropertyStatesIntegerValueBuilder
	// Build builds the BACnetPropertyStatesIntegerValue or returns an error if something is wrong
	Build() (BACnetPropertyStatesIntegerValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesIntegerValue
}

// NewBACnetPropertyStatesIntegerValueBuilder() creates a BACnetPropertyStatesIntegerValueBuilder
func NewBACnetPropertyStatesIntegerValueBuilder() BACnetPropertyStatesIntegerValueBuilder {
	return &_BACnetPropertyStatesIntegerValueBuilder{_BACnetPropertyStatesIntegerValue: new(_BACnetPropertyStatesIntegerValue)}
}

type _BACnetPropertyStatesIntegerValueBuilder struct {
	*_BACnetPropertyStatesIntegerValue

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesIntegerValueBuilder) = (*_BACnetPropertyStatesIntegerValueBuilder)(nil)

func (b *_BACnetPropertyStatesIntegerValueBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesIntegerValueBuilder) WithMandatoryFields(integerValue BACnetContextTagSignedInteger) BACnetPropertyStatesIntegerValueBuilder {
	return b.WithIntegerValue(integerValue)
}

func (b *_BACnetPropertyStatesIntegerValueBuilder) WithIntegerValue(integerValue BACnetContextTagSignedInteger) BACnetPropertyStatesIntegerValueBuilder {
	b.IntegerValue = integerValue
	return b
}

func (b *_BACnetPropertyStatesIntegerValueBuilder) WithIntegerValueBuilder(builderSupplier func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetPropertyStatesIntegerValueBuilder {
	builder := builderSupplier(b.IntegerValue.CreateBACnetContextTagSignedIntegerBuilder())
	var err error
	b.IntegerValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesIntegerValueBuilder) Build() (BACnetPropertyStatesIntegerValue, error) {
	if b.IntegerValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'integerValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesIntegerValue.deepCopy(), nil
}

func (b *_BACnetPropertyStatesIntegerValueBuilder) MustBuild() BACnetPropertyStatesIntegerValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesIntegerValueBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesIntegerValueBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesIntegerValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesIntegerValueBuilder().(*_BACnetPropertyStatesIntegerValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesIntegerValueBuilder creates a BACnetPropertyStatesIntegerValueBuilder
func (b *_BACnetPropertyStatesIntegerValue) CreateBACnetPropertyStatesIntegerValueBuilder() BACnetPropertyStatesIntegerValueBuilder {
	if b == nil {
		return NewBACnetPropertyStatesIntegerValueBuilder()
	}
	return &_BACnetPropertyStatesIntegerValueBuilder{_BACnetPropertyStatesIntegerValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesIntegerValue) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesIntegerValue) GetIntegerValue() BACnetContextTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesIntegerValue(structType any) BACnetPropertyStatesIntegerValue {
	if casted, ok := structType.(BACnetPropertyStatesIntegerValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesIntegerValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesIntegerValue) GetTypeName() string {
	return "BACnetPropertyStatesIntegerValue"
}

func (m *_BACnetPropertyStatesIntegerValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesIntegerValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesIntegerValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesIntegerValue BACnetPropertyStatesIntegerValue, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesIntegerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesIntegerValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(peekedTagNumber), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesIntegerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesIntegerValue")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesIntegerValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesIntegerValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesIntegerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesIntegerValue")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesIntegerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesIntegerValue")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesIntegerValue) IsBACnetPropertyStatesIntegerValue() {}

func (m *_BACnetPropertyStatesIntegerValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesIntegerValue) deepCopy() *_BACnetPropertyStatesIntegerValue {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesIntegerValueCopy := &_BACnetPropertyStatesIntegerValue{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.IntegerValue.DeepCopy().(BACnetContextTagSignedInteger),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesIntegerValueCopy
}

func (m *_BACnetPropertyStatesIntegerValue) String() string {
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
