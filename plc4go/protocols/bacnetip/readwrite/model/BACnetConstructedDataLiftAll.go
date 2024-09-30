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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataLiftAll is the corresponding interface of BACnetConstructedDataLiftAll
type BACnetConstructedDataLiftAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataLiftAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLiftAll()
	// CreateBuilder creates a BACnetConstructedDataLiftAllBuilder
	CreateBACnetConstructedDataLiftAllBuilder() BACnetConstructedDataLiftAllBuilder
}

// _BACnetConstructedDataLiftAll is the data-structure of this message
type _BACnetConstructedDataLiftAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataLiftAll = (*_BACnetConstructedDataLiftAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLiftAll)(nil)

// NewBACnetConstructedDataLiftAll factory function for _BACnetConstructedDataLiftAll
func NewBACnetConstructedDataLiftAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLiftAll {
	_result := &_BACnetConstructedDataLiftAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLiftAllBuilder is a builder for BACnetConstructedDataLiftAll
type BACnetConstructedDataLiftAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataLiftAllBuilder
	// Build builds the BACnetConstructedDataLiftAll or returns an error if something is wrong
	Build() (BACnetConstructedDataLiftAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLiftAll
}

// NewBACnetConstructedDataLiftAllBuilder() creates a BACnetConstructedDataLiftAllBuilder
func NewBACnetConstructedDataLiftAllBuilder() BACnetConstructedDataLiftAllBuilder {
	return &_BACnetConstructedDataLiftAllBuilder{_BACnetConstructedDataLiftAll: new(_BACnetConstructedDataLiftAll)}
}

type _BACnetConstructedDataLiftAllBuilder struct {
	*_BACnetConstructedDataLiftAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLiftAllBuilder) = (*_BACnetConstructedDataLiftAllBuilder)(nil)

func (b *_BACnetConstructedDataLiftAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLiftAllBuilder) WithMandatoryFields() BACnetConstructedDataLiftAllBuilder {
	return b
}

func (b *_BACnetConstructedDataLiftAllBuilder) Build() (BACnetConstructedDataLiftAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLiftAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataLiftAllBuilder) MustBuild() BACnetConstructedDataLiftAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLiftAllBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLiftAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLiftAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLiftAllBuilder().(*_BACnetConstructedDataLiftAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLiftAllBuilder creates a BACnetConstructedDataLiftAllBuilder
func (b *_BACnetConstructedDataLiftAll) CreateBACnetConstructedDataLiftAllBuilder() BACnetConstructedDataLiftAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataLiftAllBuilder()
	}
	return &_BACnetConstructedDataLiftAllBuilder{_BACnetConstructedDataLiftAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLiftAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFT
}

func (m *_BACnetConstructedDataLiftAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLiftAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLiftAll(structType any) BACnetConstructedDataLiftAll {
	if casted, ok := structType.(BACnetConstructedDataLiftAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLiftAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLiftAll) GetTypeName() string {
	return "BACnetConstructedDataLiftAll"
}

func (m *_BACnetConstructedDataLiftAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataLiftAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLiftAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLiftAll BACnetConstructedDataLiftAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLiftAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLiftAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLiftAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLiftAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLiftAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLiftAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLiftAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLiftAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLiftAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLiftAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLiftAll) IsBACnetConstructedDataLiftAll() {}

func (m *_BACnetConstructedDataLiftAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLiftAll) deepCopy() *_BACnetConstructedDataLiftAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLiftAllCopy := &_BACnetConstructedDataLiftAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLiftAllCopy
}

func (m *_BACnetConstructedDataLiftAll) String() string {
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
