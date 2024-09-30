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

// BACnetConstructedDataDerivativeConstant is the corresponding interface of BACnetConstructedDataDerivativeConstant
type BACnetConstructedDataDerivativeConstant interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDerivativeConstant returns DerivativeConstant (property field)
	GetDerivativeConstant() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataDerivativeConstant is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDerivativeConstant()
	// CreateBuilder creates a BACnetConstructedDataDerivativeConstantBuilder
	CreateBACnetConstructedDataDerivativeConstantBuilder() BACnetConstructedDataDerivativeConstantBuilder
}

// _BACnetConstructedDataDerivativeConstant is the data-structure of this message
type _BACnetConstructedDataDerivativeConstant struct {
	BACnetConstructedDataContract
	DerivativeConstant BACnetApplicationTagReal
}

var _ BACnetConstructedDataDerivativeConstant = (*_BACnetConstructedDataDerivativeConstant)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDerivativeConstant)(nil)

// NewBACnetConstructedDataDerivativeConstant factory function for _BACnetConstructedDataDerivativeConstant
func NewBACnetConstructedDataDerivativeConstant(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, derivativeConstant BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDerivativeConstant {
	if derivativeConstant == nil {
		panic("derivativeConstant of type BACnetApplicationTagReal for BACnetConstructedDataDerivativeConstant must not be nil")
	}
	_result := &_BACnetConstructedDataDerivativeConstant{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DerivativeConstant:            derivativeConstant,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDerivativeConstantBuilder is a builder for BACnetConstructedDataDerivativeConstant
type BACnetConstructedDataDerivativeConstantBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(derivativeConstant BACnetApplicationTagReal) BACnetConstructedDataDerivativeConstantBuilder
	// WithDerivativeConstant adds DerivativeConstant (property field)
	WithDerivativeConstant(BACnetApplicationTagReal) BACnetConstructedDataDerivativeConstantBuilder
	// WithDerivativeConstantBuilder adds DerivativeConstant (property field) which is build by the builder
	WithDerivativeConstantBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataDerivativeConstantBuilder
	// Build builds the BACnetConstructedDataDerivativeConstant or returns an error if something is wrong
	Build() (BACnetConstructedDataDerivativeConstant, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDerivativeConstant
}

// NewBACnetConstructedDataDerivativeConstantBuilder() creates a BACnetConstructedDataDerivativeConstantBuilder
func NewBACnetConstructedDataDerivativeConstantBuilder() BACnetConstructedDataDerivativeConstantBuilder {
	return &_BACnetConstructedDataDerivativeConstantBuilder{_BACnetConstructedDataDerivativeConstant: new(_BACnetConstructedDataDerivativeConstant)}
}

type _BACnetConstructedDataDerivativeConstantBuilder struct {
	*_BACnetConstructedDataDerivativeConstant

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDerivativeConstantBuilder) = (*_BACnetConstructedDataDerivativeConstantBuilder)(nil)

func (b *_BACnetConstructedDataDerivativeConstantBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataDerivativeConstantBuilder) WithMandatoryFields(derivativeConstant BACnetApplicationTagReal) BACnetConstructedDataDerivativeConstantBuilder {
	return b.WithDerivativeConstant(derivativeConstant)
}

func (b *_BACnetConstructedDataDerivativeConstantBuilder) WithDerivativeConstant(derivativeConstant BACnetApplicationTagReal) BACnetConstructedDataDerivativeConstantBuilder {
	b.DerivativeConstant = derivativeConstant
	return b
}

func (b *_BACnetConstructedDataDerivativeConstantBuilder) WithDerivativeConstantBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataDerivativeConstantBuilder {
	builder := builderSupplier(b.DerivativeConstant.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.DerivativeConstant, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDerivativeConstantBuilder) Build() (BACnetConstructedDataDerivativeConstant, error) {
	if b.DerivativeConstant == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'derivativeConstant' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDerivativeConstant.deepCopy(), nil
}

func (b *_BACnetConstructedDataDerivativeConstantBuilder) MustBuild() BACnetConstructedDataDerivativeConstant {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataDerivativeConstantBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDerivativeConstantBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDerivativeConstantBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDerivativeConstantBuilder().(*_BACnetConstructedDataDerivativeConstantBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDerivativeConstantBuilder creates a BACnetConstructedDataDerivativeConstantBuilder
func (b *_BACnetConstructedDataDerivativeConstant) CreateBACnetConstructedDataDerivativeConstantBuilder() BACnetConstructedDataDerivativeConstantBuilder {
	if b == nil {
		return NewBACnetConstructedDataDerivativeConstantBuilder()
	}
	return &_BACnetConstructedDataDerivativeConstantBuilder{_BACnetConstructedDataDerivativeConstant: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstant) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDerivativeConstant) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DERIVATIVE_CONSTANT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDerivativeConstant) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstant) GetDerivativeConstant() BACnetApplicationTagReal {
	return m.DerivativeConstant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstant) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetDerivativeConstant())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDerivativeConstant(structType any) BACnetConstructedDataDerivativeConstant {
	if casted, ok := structType.(BACnetConstructedDataDerivativeConstant); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDerivativeConstant); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDerivativeConstant) GetTypeName() string {
	return "BACnetConstructedDataDerivativeConstant"
}

func (m *_BACnetConstructedDataDerivativeConstant) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (derivativeConstant)
	lengthInBits += m.DerivativeConstant.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDerivativeConstant) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDerivativeConstant) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDerivativeConstant BACnetConstructedDataDerivativeConstant, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDerivativeConstant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDerivativeConstant")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	derivativeConstant, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "derivativeConstant", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'derivativeConstant' field"))
	}
	m.DerivativeConstant = derivativeConstant

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), derivativeConstant)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDerivativeConstant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDerivativeConstant")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDerivativeConstant) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDerivativeConstant) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDerivativeConstant"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDerivativeConstant")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "derivativeConstant", m.GetDerivativeConstant(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'derivativeConstant' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDerivativeConstant"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDerivativeConstant")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDerivativeConstant) IsBACnetConstructedDataDerivativeConstant() {}

func (m *_BACnetConstructedDataDerivativeConstant) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDerivativeConstant) deepCopy() *_BACnetConstructedDataDerivativeConstant {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDerivativeConstantCopy := &_BACnetConstructedDataDerivativeConstant{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.DerivativeConstant.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDerivativeConstantCopy
}

func (m *_BACnetConstructedDataDerivativeConstant) String() string {
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
