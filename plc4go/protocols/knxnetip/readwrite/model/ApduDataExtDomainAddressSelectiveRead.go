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

// ApduDataExtDomainAddressSelectiveRead is the corresponding interface of ApduDataExtDomainAddressSelectiveRead
type ApduDataExtDomainAddressSelectiveRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtDomainAddressSelectiveRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtDomainAddressSelectiveRead()
	// CreateBuilder creates a ApduDataExtDomainAddressSelectiveReadBuilder
	CreateApduDataExtDomainAddressSelectiveReadBuilder() ApduDataExtDomainAddressSelectiveReadBuilder
}

// _ApduDataExtDomainAddressSelectiveRead is the data-structure of this message
type _ApduDataExtDomainAddressSelectiveRead struct {
	ApduDataExtContract
}

var _ ApduDataExtDomainAddressSelectiveRead = (*_ApduDataExtDomainAddressSelectiveRead)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtDomainAddressSelectiveRead)(nil)

// NewApduDataExtDomainAddressSelectiveRead factory function for _ApduDataExtDomainAddressSelectiveRead
func NewApduDataExtDomainAddressSelectiveRead(length uint8) *_ApduDataExtDomainAddressSelectiveRead {
	_result := &_ApduDataExtDomainAddressSelectiveRead{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtDomainAddressSelectiveReadBuilder is a builder for ApduDataExtDomainAddressSelectiveRead
type ApduDataExtDomainAddressSelectiveReadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtDomainAddressSelectiveReadBuilder
	// Build builds the ApduDataExtDomainAddressSelectiveRead or returns an error if something is wrong
	Build() (ApduDataExtDomainAddressSelectiveRead, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtDomainAddressSelectiveRead
}

// NewApduDataExtDomainAddressSelectiveReadBuilder() creates a ApduDataExtDomainAddressSelectiveReadBuilder
func NewApduDataExtDomainAddressSelectiveReadBuilder() ApduDataExtDomainAddressSelectiveReadBuilder {
	return &_ApduDataExtDomainAddressSelectiveReadBuilder{_ApduDataExtDomainAddressSelectiveRead: new(_ApduDataExtDomainAddressSelectiveRead)}
}

type _ApduDataExtDomainAddressSelectiveReadBuilder struct {
	*_ApduDataExtDomainAddressSelectiveRead

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtDomainAddressSelectiveReadBuilder) = (*_ApduDataExtDomainAddressSelectiveReadBuilder)(nil)

func (b *_ApduDataExtDomainAddressSelectiveReadBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
}

func (b *_ApduDataExtDomainAddressSelectiveReadBuilder) WithMandatoryFields() ApduDataExtDomainAddressSelectiveReadBuilder {
	return b
}

func (b *_ApduDataExtDomainAddressSelectiveReadBuilder) Build() (ApduDataExtDomainAddressSelectiveRead, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtDomainAddressSelectiveRead.deepCopy(), nil
}

func (b *_ApduDataExtDomainAddressSelectiveReadBuilder) MustBuild() ApduDataExtDomainAddressSelectiveRead {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ApduDataExtDomainAddressSelectiveReadBuilder) Done() ApduDataExtBuilder {
	return b.parentBuilder
}

func (b *_ApduDataExtDomainAddressSelectiveReadBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtDomainAddressSelectiveReadBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtDomainAddressSelectiveReadBuilder().(*_ApduDataExtDomainAddressSelectiveReadBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtDomainAddressSelectiveReadBuilder creates a ApduDataExtDomainAddressSelectiveReadBuilder
func (b *_ApduDataExtDomainAddressSelectiveRead) CreateApduDataExtDomainAddressSelectiveReadBuilder() ApduDataExtDomainAddressSelectiveReadBuilder {
	if b == nil {
		return NewApduDataExtDomainAddressSelectiveReadBuilder()
	}
	return &_ApduDataExtDomainAddressSelectiveReadBuilder{_ApduDataExtDomainAddressSelectiveRead: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressSelectiveRead) GetExtApciType() uint8 {
	return 0x23
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressSelectiveRead) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressSelectiveRead(structType any) ApduDataExtDomainAddressSelectiveRead {
	if casted, ok := structType.(ApduDataExtDomainAddressSelectiveRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressSelectiveRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressSelectiveRead) GetTypeName() string {
	return "ApduDataExtDomainAddressSelectiveRead"
}

func (m *_ApduDataExtDomainAddressSelectiveRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressSelectiveRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtDomainAddressSelectiveRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtDomainAddressSelectiveRead ApduDataExtDomainAddressSelectiveRead, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressSelectiveRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressSelectiveRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressSelectiveRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressSelectiveRead")
	}

	return m, nil
}

func (m *_ApduDataExtDomainAddressSelectiveRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtDomainAddressSelectiveRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressSelectiveRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressSelectiveRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressSelectiveRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressSelectiveRead")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressSelectiveRead) IsApduDataExtDomainAddressSelectiveRead() {}

func (m *_ApduDataExtDomainAddressSelectiveRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtDomainAddressSelectiveRead) deepCopy() *_ApduDataExtDomainAddressSelectiveRead {
	if m == nil {
		return nil
	}
	_ApduDataExtDomainAddressSelectiveReadCopy := &_ApduDataExtDomainAddressSelectiveRead{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtDomainAddressSelectiveReadCopy
}

func (m *_ApduDataExtDomainAddressSelectiveRead) String() string {
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
