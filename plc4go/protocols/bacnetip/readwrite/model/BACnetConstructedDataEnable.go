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

// BACnetConstructedDataEnable is the corresponding interface of BACnetConstructedDataEnable
type BACnetConstructedDataEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEnable returns Enable (property field)
	GetEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataEnable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEnable()
	// CreateBuilder creates a BACnetConstructedDataEnableBuilder
	CreateBACnetConstructedDataEnableBuilder() BACnetConstructedDataEnableBuilder
}

// _BACnetConstructedDataEnable is the data-structure of this message
type _BACnetConstructedDataEnable struct {
	BACnetConstructedDataContract
	Enable BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataEnable = (*_BACnetConstructedDataEnable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEnable)(nil)

// NewBACnetConstructedDataEnable factory function for _BACnetConstructedDataEnable
func NewBACnetConstructedDataEnable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, enable BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEnable {
	if enable == nil {
		panic("enable of type BACnetApplicationTagBoolean for BACnetConstructedDataEnable must not be nil")
	}
	_result := &_BACnetConstructedDataEnable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Enable:                        enable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEnableBuilder is a builder for BACnetConstructedDataEnable
type BACnetConstructedDataEnableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(enable BACnetApplicationTagBoolean) BACnetConstructedDataEnableBuilder
	// WithEnable adds Enable (property field)
	WithEnable(BACnetApplicationTagBoolean) BACnetConstructedDataEnableBuilder
	// WithEnableBuilder adds Enable (property field) which is build by the builder
	WithEnableBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataEnableBuilder
	// Build builds the BACnetConstructedDataEnable or returns an error if something is wrong
	Build() (BACnetConstructedDataEnable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEnable
}

// NewBACnetConstructedDataEnableBuilder() creates a BACnetConstructedDataEnableBuilder
func NewBACnetConstructedDataEnableBuilder() BACnetConstructedDataEnableBuilder {
	return &_BACnetConstructedDataEnableBuilder{_BACnetConstructedDataEnable: new(_BACnetConstructedDataEnable)}
}

type _BACnetConstructedDataEnableBuilder struct {
	*_BACnetConstructedDataEnable

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEnableBuilder) = (*_BACnetConstructedDataEnableBuilder)(nil)

func (b *_BACnetConstructedDataEnableBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataEnableBuilder) WithMandatoryFields(enable BACnetApplicationTagBoolean) BACnetConstructedDataEnableBuilder {
	return b.WithEnable(enable)
}

func (b *_BACnetConstructedDataEnableBuilder) WithEnable(enable BACnetApplicationTagBoolean) BACnetConstructedDataEnableBuilder {
	b.Enable = enable
	return b
}

func (b *_BACnetConstructedDataEnableBuilder) WithEnableBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataEnableBuilder {
	builder := builderSupplier(b.Enable.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.Enable, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataEnableBuilder) Build() (BACnetConstructedDataEnable, error) {
	if b.Enable == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'enable' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEnable.deepCopy(), nil
}

func (b *_BACnetConstructedDataEnableBuilder) MustBuild() BACnetConstructedDataEnable {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataEnableBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEnableBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEnableBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEnableBuilder().(*_BACnetConstructedDataEnableBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEnableBuilder creates a BACnetConstructedDataEnableBuilder
func (b *_BACnetConstructedDataEnable) CreateBACnetConstructedDataEnableBuilder() BACnetConstructedDataEnableBuilder {
	if b == nil {
		return NewBACnetConstructedDataEnableBuilder()
	}
	return &_BACnetConstructedDataEnableBuilder{_BACnetConstructedDataEnable: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEnable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEnable) GetEnable() BACnetApplicationTagBoolean {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEnable) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEnable(structType any) BACnetConstructedDataEnable {
	if casted, ok := structType.(BACnetConstructedDataEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEnable) GetTypeName() string {
	return "BACnetConstructedDataEnable"
}

func (m *_BACnetConstructedDataEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (enable)
	lengthInBits += m.Enable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEnable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEnable BACnetConstructedDataEnable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	enable, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "enable", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enable' field"))
	}
	m.Enable = enable

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), enable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEnable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEnable")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "enable", m.GetEnable(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEnable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEnable) IsBACnetConstructedDataEnable() {}

func (m *_BACnetConstructedDataEnable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEnable) deepCopy() *_BACnetConstructedDataEnable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEnableCopy := &_BACnetConstructedDataEnable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Enable.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEnableCopy
}

func (m *_BACnetConstructedDataEnable) String() string {
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
