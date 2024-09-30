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

// UnitAddress is the corresponding interface of UnitAddress
type UnitAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetAddress returns Address (property field)
	GetAddress() byte
	// IsUnitAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUnitAddress()
	// CreateBuilder creates a UnitAddressBuilder
	CreateUnitAddressBuilder() UnitAddressBuilder
}

// _UnitAddress is the data-structure of this message
type _UnitAddress struct {
	Address byte
}

var _ UnitAddress = (*_UnitAddress)(nil)

// NewUnitAddress factory function for _UnitAddress
func NewUnitAddress(address byte) *_UnitAddress {
	return &_UnitAddress{Address: address}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// UnitAddressBuilder is a builder for UnitAddress
type UnitAddressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(address byte) UnitAddressBuilder
	// WithAddress adds Address (property field)
	WithAddress(byte) UnitAddressBuilder
	// Build builds the UnitAddress or returns an error if something is wrong
	Build() (UnitAddress, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() UnitAddress
}

// NewUnitAddressBuilder() creates a UnitAddressBuilder
func NewUnitAddressBuilder() UnitAddressBuilder {
	return &_UnitAddressBuilder{_UnitAddress: new(_UnitAddress)}
}

type _UnitAddressBuilder struct {
	*_UnitAddress

	err *utils.MultiError
}

var _ (UnitAddressBuilder) = (*_UnitAddressBuilder)(nil)

func (b *_UnitAddressBuilder) WithMandatoryFields(address byte) UnitAddressBuilder {
	return b.WithAddress(address)
}

func (b *_UnitAddressBuilder) WithAddress(address byte) UnitAddressBuilder {
	b.Address = address
	return b
}

func (b *_UnitAddressBuilder) Build() (UnitAddress, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._UnitAddress.deepCopy(), nil
}

func (b *_UnitAddressBuilder) MustBuild() UnitAddress {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_UnitAddressBuilder) DeepCopy() any {
	_copy := b.CreateUnitAddressBuilder().(*_UnitAddressBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateUnitAddressBuilder creates a UnitAddressBuilder
func (b *_UnitAddress) CreateUnitAddressBuilder() UnitAddressBuilder {
	if b == nil {
		return NewUnitAddressBuilder()
	}
	return &_UnitAddressBuilder{_UnitAddress: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UnitAddress) GetAddress() byte {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastUnitAddress(structType any) UnitAddress {
	if casted, ok := structType.(UnitAddress); ok {
		return casted
	}
	if casted, ok := structType.(*UnitAddress); ok {
		return *casted
	}
	return nil
}

func (m *_UnitAddress) GetTypeName() string {
	return "UnitAddress"
}

func (m *_UnitAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (address)
	lengthInBits += 8

	return lengthInBits
}

func (m *_UnitAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UnitAddressParse(ctx context.Context, theBytes []byte) (UnitAddress, error) {
	return UnitAddressParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func UnitAddressParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (UnitAddress, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (UnitAddress, error) {
		return UnitAddressParseWithBuffer(ctx, readBuffer)
	}
}

func UnitAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (UnitAddress, error) {
	v, err := (&_UnitAddress{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_UnitAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__unitAddress UnitAddress, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UnitAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UnitAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	if closeErr := readBuffer.CloseContext("UnitAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UnitAddress")
	}

	return m, nil
}

func (m *_UnitAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UnitAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("UnitAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for UnitAddress")
	}

	if err := WriteSimpleField[byte](ctx, "address", m.GetAddress(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'address' field")
	}

	if popErr := writeBuffer.PopContext("UnitAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for UnitAddress")
	}
	return nil
}

func (m *_UnitAddress) IsUnitAddress() {}

func (m *_UnitAddress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_UnitAddress) deepCopy() *_UnitAddress {
	if m == nil {
		return nil
	}
	_UnitAddressCopy := &_UnitAddress{
		m.Address,
	}
	return _UnitAddressCopy
}

func (m *_UnitAddress) String() string {
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
