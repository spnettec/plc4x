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

// COTPParameterCalledTsap is the corresponding interface of COTPParameterCalledTsap
type COTPParameterCalledTsap interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	COTPParameter
	// GetTsapId returns TsapId (property field)
	GetTsapId() uint16
	// IsCOTPParameterCalledTsap is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPParameterCalledTsap()
	// CreateBuilder creates a COTPParameterCalledTsapBuilder
	CreateCOTPParameterCalledTsapBuilder() COTPParameterCalledTsapBuilder
}

// _COTPParameterCalledTsap is the data-structure of this message
type _COTPParameterCalledTsap struct {
	COTPParameterContract
	TsapId uint16
}

var _ COTPParameterCalledTsap = (*_COTPParameterCalledTsap)(nil)
var _ COTPParameterRequirements = (*_COTPParameterCalledTsap)(nil)

// NewCOTPParameterCalledTsap factory function for _COTPParameterCalledTsap
func NewCOTPParameterCalledTsap(tsapId uint16, rest uint8) *_COTPParameterCalledTsap {
	_result := &_COTPParameterCalledTsap{
		COTPParameterContract: NewCOTPParameter(rest),
		TsapId:                tsapId,
	}
	_result.COTPParameterContract.(*_COTPParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// COTPParameterCalledTsapBuilder is a builder for COTPParameterCalledTsap
type COTPParameterCalledTsapBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(tsapId uint16) COTPParameterCalledTsapBuilder
	// WithTsapId adds TsapId (property field)
	WithTsapId(uint16) COTPParameterCalledTsapBuilder
	// Build builds the COTPParameterCalledTsap or returns an error if something is wrong
	Build() (COTPParameterCalledTsap, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() COTPParameterCalledTsap
}

// NewCOTPParameterCalledTsapBuilder() creates a COTPParameterCalledTsapBuilder
func NewCOTPParameterCalledTsapBuilder() COTPParameterCalledTsapBuilder {
	return &_COTPParameterCalledTsapBuilder{_COTPParameterCalledTsap: new(_COTPParameterCalledTsap)}
}

type _COTPParameterCalledTsapBuilder struct {
	*_COTPParameterCalledTsap

	parentBuilder *_COTPParameterBuilder

	err *utils.MultiError
}

var _ (COTPParameterCalledTsapBuilder) = (*_COTPParameterCalledTsapBuilder)(nil)

func (b *_COTPParameterCalledTsapBuilder) setParent(contract COTPParameterContract) {
	b.COTPParameterContract = contract
}

func (b *_COTPParameterCalledTsapBuilder) WithMandatoryFields(tsapId uint16) COTPParameterCalledTsapBuilder {
	return b.WithTsapId(tsapId)
}

func (b *_COTPParameterCalledTsapBuilder) WithTsapId(tsapId uint16) COTPParameterCalledTsapBuilder {
	b.TsapId = tsapId
	return b
}

func (b *_COTPParameterCalledTsapBuilder) Build() (COTPParameterCalledTsap, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._COTPParameterCalledTsap.deepCopy(), nil
}

func (b *_COTPParameterCalledTsapBuilder) MustBuild() COTPParameterCalledTsap {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_COTPParameterCalledTsapBuilder) Done() COTPParameterBuilder {
	return b.parentBuilder
}

func (b *_COTPParameterCalledTsapBuilder) buildForCOTPParameter() (COTPParameter, error) {
	return b.Build()
}

func (b *_COTPParameterCalledTsapBuilder) DeepCopy() any {
	_copy := b.CreateCOTPParameterCalledTsapBuilder().(*_COTPParameterCalledTsapBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCOTPParameterCalledTsapBuilder creates a COTPParameterCalledTsapBuilder
func (b *_COTPParameterCalledTsap) CreateCOTPParameterCalledTsapBuilder() COTPParameterCalledTsapBuilder {
	if b == nil {
		return NewCOTPParameterCalledTsapBuilder()
	}
	return &_COTPParameterCalledTsapBuilder{_COTPParameterCalledTsap: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPParameterCalledTsap) GetParameterType() uint8 {
	return 0xC2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPParameterCalledTsap) GetParent() COTPParameterContract {
	return m.COTPParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPParameterCalledTsap) GetTsapId() uint16 {
	return m.TsapId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCOTPParameterCalledTsap(structType any) COTPParameterCalledTsap {
	if casted, ok := structType.(COTPParameterCalledTsap); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameterCalledTsap); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameterCalledTsap) GetTypeName() string {
	return "COTPParameterCalledTsap"
}

func (m *_COTPParameterCalledTsap) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.COTPParameterContract.(*_COTPParameter).getLengthInBits(ctx))

	// Simple field (tsapId)
	lengthInBits += 16

	return lengthInBits
}

func (m *_COTPParameterCalledTsap) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_COTPParameterCalledTsap) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_COTPParameter, rest uint8) (__cOTPParameterCalledTsap COTPParameterCalledTsap, err error) {
	m.COTPParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameterCalledTsap"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameterCalledTsap")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	tsapId, err := ReadSimpleField(ctx, "tsapId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tsapId' field"))
	}
	m.TsapId = tsapId

	if closeErr := readBuffer.CloseContext("COTPParameterCalledTsap"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameterCalledTsap")
	}

	return m, nil
}

func (m *_COTPParameterCalledTsap) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPParameterCalledTsap) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterCalledTsap"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPParameterCalledTsap")
		}

		if err := WriteSimpleField[uint16](ctx, "tsapId", m.GetTsapId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'tsapId' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterCalledTsap"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPParameterCalledTsap")
		}
		return nil
	}
	return m.COTPParameterContract.(*_COTPParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPParameterCalledTsap) IsCOTPParameterCalledTsap() {}

func (m *_COTPParameterCalledTsap) DeepCopy() any {
	return m.deepCopy()
}

func (m *_COTPParameterCalledTsap) deepCopy() *_COTPParameterCalledTsap {
	if m == nil {
		return nil
	}
	_COTPParameterCalledTsapCopy := &_COTPParameterCalledTsap{
		m.COTPParameterContract.(*_COTPParameter).deepCopy(),
		m.TsapId,
	}
	m.COTPParameterContract.(*_COTPParameter)._SubType = m
	return _COTPParameterCalledTsapCopy
}

func (m *_COTPParameterCalledTsap) String() string {
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
