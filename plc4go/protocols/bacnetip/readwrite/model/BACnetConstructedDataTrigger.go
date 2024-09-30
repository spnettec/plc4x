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

// BACnetConstructedDataTrigger is the corresponding interface of BACnetConstructedDataTrigger
type BACnetConstructedDataTrigger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetTrigger returns Trigger (property field)
	GetTrigger() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataTrigger is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTrigger()
	// CreateBuilder creates a BACnetConstructedDataTriggerBuilder
	CreateBACnetConstructedDataTriggerBuilder() BACnetConstructedDataTriggerBuilder
}

// _BACnetConstructedDataTrigger is the data-structure of this message
type _BACnetConstructedDataTrigger struct {
	BACnetConstructedDataContract
	Trigger BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataTrigger = (*_BACnetConstructedDataTrigger)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTrigger)(nil)

// NewBACnetConstructedDataTrigger factory function for _BACnetConstructedDataTrigger
func NewBACnetConstructedDataTrigger(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, trigger BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrigger {
	if trigger == nil {
		panic("trigger of type BACnetApplicationTagBoolean for BACnetConstructedDataTrigger must not be nil")
	}
	_result := &_BACnetConstructedDataTrigger{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Trigger:                       trigger,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTriggerBuilder is a builder for BACnetConstructedDataTrigger
type BACnetConstructedDataTriggerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(trigger BACnetApplicationTagBoolean) BACnetConstructedDataTriggerBuilder
	// WithTrigger adds Trigger (property field)
	WithTrigger(BACnetApplicationTagBoolean) BACnetConstructedDataTriggerBuilder
	// WithTriggerBuilder adds Trigger (property field) which is build by the builder
	WithTriggerBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataTriggerBuilder
	// Build builds the BACnetConstructedDataTrigger or returns an error if something is wrong
	Build() (BACnetConstructedDataTrigger, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTrigger
}

// NewBACnetConstructedDataTriggerBuilder() creates a BACnetConstructedDataTriggerBuilder
func NewBACnetConstructedDataTriggerBuilder() BACnetConstructedDataTriggerBuilder {
	return &_BACnetConstructedDataTriggerBuilder{_BACnetConstructedDataTrigger: new(_BACnetConstructedDataTrigger)}
}

type _BACnetConstructedDataTriggerBuilder struct {
	*_BACnetConstructedDataTrigger

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTriggerBuilder) = (*_BACnetConstructedDataTriggerBuilder)(nil)

func (b *_BACnetConstructedDataTriggerBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataTriggerBuilder) WithMandatoryFields(trigger BACnetApplicationTagBoolean) BACnetConstructedDataTriggerBuilder {
	return b.WithTrigger(trigger)
}

func (b *_BACnetConstructedDataTriggerBuilder) WithTrigger(trigger BACnetApplicationTagBoolean) BACnetConstructedDataTriggerBuilder {
	b.Trigger = trigger
	return b
}

func (b *_BACnetConstructedDataTriggerBuilder) WithTriggerBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataTriggerBuilder {
	builder := builderSupplier(b.Trigger.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.Trigger, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataTriggerBuilder) Build() (BACnetConstructedDataTrigger, error) {
	if b.Trigger == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'trigger' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTrigger.deepCopy(), nil
}

func (b *_BACnetConstructedDataTriggerBuilder) MustBuild() BACnetConstructedDataTrigger {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataTriggerBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTriggerBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTriggerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTriggerBuilder().(*_BACnetConstructedDataTriggerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTriggerBuilder creates a BACnetConstructedDataTriggerBuilder
func (b *_BACnetConstructedDataTrigger) CreateBACnetConstructedDataTriggerBuilder() BACnetConstructedDataTriggerBuilder {
	if b == nil {
		return NewBACnetConstructedDataTriggerBuilder()
	}
	return &_BACnetConstructedDataTriggerBuilder{_BACnetConstructedDataTrigger: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrigger) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTrigger) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRIGGER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrigger) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTrigger) GetTrigger() BACnetApplicationTagBoolean {
	return m.Trigger
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTrigger) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetTrigger())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrigger(structType any) BACnetConstructedDataTrigger {
	if casted, ok := structType.(BACnetConstructedDataTrigger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrigger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrigger) GetTypeName() string {
	return "BACnetConstructedDataTrigger"
}

func (m *_BACnetConstructedDataTrigger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (trigger)
	lengthInBits += m.Trigger.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTrigger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTrigger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTrigger BACnetConstructedDataTrigger, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrigger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrigger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	trigger, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "trigger", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'trigger' field"))
	}
	m.Trigger = trigger

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), trigger)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrigger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrigger")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTrigger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTrigger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrigger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrigger")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "trigger", m.GetTrigger(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'trigger' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrigger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrigger")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrigger) IsBACnetConstructedDataTrigger() {}

func (m *_BACnetConstructedDataTrigger) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTrigger) deepCopy() *_BACnetConstructedDataTrigger {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTriggerCopy := &_BACnetConstructedDataTrigger{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Trigger.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTriggerCopy
}

func (m *_BACnetConstructedDataTrigger) String() string {
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
