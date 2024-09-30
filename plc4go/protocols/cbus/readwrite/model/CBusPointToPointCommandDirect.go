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

// CBusPointToPointCommandDirect is the corresponding interface of CBusPointToPointCommandDirect
type CBusPointToPointCommandDirect interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CBusPointToPointCommand
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
	// IsCBusPointToPointCommandDirect is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusPointToPointCommandDirect()
	// CreateBuilder creates a CBusPointToPointCommandDirectBuilder
	CreateCBusPointToPointCommandDirectBuilder() CBusPointToPointCommandDirectBuilder
}

// _CBusPointToPointCommandDirect is the data-structure of this message
type _CBusPointToPointCommandDirect struct {
	CBusPointToPointCommandContract
	UnitAddress UnitAddress
	// Reserved Fields
	reservedField0 *uint8
}

var _ CBusPointToPointCommandDirect = (*_CBusPointToPointCommandDirect)(nil)
var _ CBusPointToPointCommandRequirements = (*_CBusPointToPointCommandDirect)(nil)

// NewCBusPointToPointCommandDirect factory function for _CBusPointToPointCommandDirect
func NewCBusPointToPointCommandDirect(bridgeAddressCountPeek uint16, calData CALData, unitAddress UnitAddress, cBusOptions CBusOptions) *_CBusPointToPointCommandDirect {
	if unitAddress == nil {
		panic("unitAddress of type UnitAddress for CBusPointToPointCommandDirect must not be nil")
	}
	_result := &_CBusPointToPointCommandDirect{
		CBusPointToPointCommandContract: NewCBusPointToPointCommand(bridgeAddressCountPeek, calData, cBusOptions),
		UnitAddress:                     unitAddress,
	}
	_result.CBusPointToPointCommandContract.(*_CBusPointToPointCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusPointToPointCommandDirectBuilder is a builder for CBusPointToPointCommandDirect
type CBusPointToPointCommandDirectBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unitAddress UnitAddress) CBusPointToPointCommandDirectBuilder
	// WithUnitAddress adds UnitAddress (property field)
	WithUnitAddress(UnitAddress) CBusPointToPointCommandDirectBuilder
	// WithUnitAddressBuilder adds UnitAddress (property field) which is build by the builder
	WithUnitAddressBuilder(func(UnitAddressBuilder) UnitAddressBuilder) CBusPointToPointCommandDirectBuilder
	// Build builds the CBusPointToPointCommandDirect or returns an error if something is wrong
	Build() (CBusPointToPointCommandDirect, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusPointToPointCommandDirect
}

// NewCBusPointToPointCommandDirectBuilder() creates a CBusPointToPointCommandDirectBuilder
func NewCBusPointToPointCommandDirectBuilder() CBusPointToPointCommandDirectBuilder {
	return &_CBusPointToPointCommandDirectBuilder{_CBusPointToPointCommandDirect: new(_CBusPointToPointCommandDirect)}
}

type _CBusPointToPointCommandDirectBuilder struct {
	*_CBusPointToPointCommandDirect

	parentBuilder *_CBusPointToPointCommandBuilder

	err *utils.MultiError
}

var _ (CBusPointToPointCommandDirectBuilder) = (*_CBusPointToPointCommandDirectBuilder)(nil)

func (b *_CBusPointToPointCommandDirectBuilder) setParent(contract CBusPointToPointCommandContract) {
	b.CBusPointToPointCommandContract = contract
}

func (b *_CBusPointToPointCommandDirectBuilder) WithMandatoryFields(unitAddress UnitAddress) CBusPointToPointCommandDirectBuilder {
	return b.WithUnitAddress(unitAddress)
}

func (b *_CBusPointToPointCommandDirectBuilder) WithUnitAddress(unitAddress UnitAddress) CBusPointToPointCommandDirectBuilder {
	b.UnitAddress = unitAddress
	return b
}

func (b *_CBusPointToPointCommandDirectBuilder) WithUnitAddressBuilder(builderSupplier func(UnitAddressBuilder) UnitAddressBuilder) CBusPointToPointCommandDirectBuilder {
	builder := builderSupplier(b.UnitAddress.CreateUnitAddressBuilder())
	var err error
	b.UnitAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "UnitAddressBuilder failed"))
	}
	return b
}

func (b *_CBusPointToPointCommandDirectBuilder) Build() (CBusPointToPointCommandDirect, error) {
	if b.UnitAddress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'unitAddress' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CBusPointToPointCommandDirect.deepCopy(), nil
}

func (b *_CBusPointToPointCommandDirectBuilder) MustBuild() CBusPointToPointCommandDirect {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CBusPointToPointCommandDirectBuilder) Done() CBusPointToPointCommandBuilder {
	return b.parentBuilder
}

func (b *_CBusPointToPointCommandDirectBuilder) buildForCBusPointToPointCommand() (CBusPointToPointCommand, error) {
	return b.Build()
}

func (b *_CBusPointToPointCommandDirectBuilder) DeepCopy() any {
	_copy := b.CreateCBusPointToPointCommandDirectBuilder().(*_CBusPointToPointCommandDirectBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCBusPointToPointCommandDirectBuilder creates a CBusPointToPointCommandDirectBuilder
func (b *_CBusPointToPointCommandDirect) CreateCBusPointToPointCommandDirectBuilder() CBusPointToPointCommandDirectBuilder {
	if b == nil {
		return NewCBusPointToPointCommandDirectBuilder()
	}
	return &_CBusPointToPointCommandDirectBuilder{_CBusPointToPointCommandDirect: b.deepCopy()}
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

func (m *_CBusPointToPointCommandDirect) GetParent() CBusPointToPointCommandContract {
	return m.CBusPointToPointCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointCommandDirect) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusPointToPointCommandDirect(structType any) CBusPointToPointCommandDirect {
	if casted, ok := structType.(CBusPointToPointCommandDirect); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointCommandDirect); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointCommandDirect) GetTypeName() string {
	return "CBusPointToPointCommandDirect"
}

func (m *_CBusPointToPointCommandDirect) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).getLengthInBits(ctx))

	// Simple field (unitAddress)
	lengthInBits += m.UnitAddress.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CBusPointToPointCommandDirect) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CBusPointToPointCommandDirect) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CBusPointToPointCommand, cBusOptions CBusOptions) (__cBusPointToPointCommandDirect CBusPointToPointCommandDirect, err error) {
	m.CBusPointToPointCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointCommandDirect"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointCommandDirect")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unitAddress, err := ReadSimpleField[UnitAddress](ctx, "unitAddress", ReadComplex[UnitAddress](UnitAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitAddress' field"))
	}
	m.UnitAddress = unitAddress

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("CBusPointToPointCommandDirect"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointCommandDirect")
	}

	return m, nil
}

func (m *_CBusPointToPointCommandDirect) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToPointCommandDirect) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToPointCommandDirect"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToPointCommandDirect")
		}

		if err := WriteSimpleField[UnitAddress](ctx, "unitAddress", m.GetUnitAddress(), WriteComplex[UnitAddress](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitAddress' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if popErr := writeBuffer.PopContext("CBusPointToPointCommandDirect"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToPointCommandDirect")
		}
		return nil
	}
	return m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusPointToPointCommandDirect) IsCBusPointToPointCommandDirect() {}

func (m *_CBusPointToPointCommandDirect) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusPointToPointCommandDirect) deepCopy() *_CBusPointToPointCommandDirect {
	if m == nil {
		return nil
	}
	_CBusPointToPointCommandDirectCopy := &_CBusPointToPointCommandDirect{
		m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).deepCopy(),
		m.UnitAddress.DeepCopy().(UnitAddress),
		m.reservedField0,
	}
	m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand)._SubType = m
	return _CBusPointToPointCommandDirectCopy
}

func (m *_CBusPointToPointCommandDirect) String() string {
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
