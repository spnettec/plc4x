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

// SecurityDataZoneShort is the corresponding interface of SecurityDataZoneShort
type SecurityDataZoneShort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
	// IsSecurityDataZoneShort is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataZoneShort()
	// CreateBuilder creates a SecurityDataZoneShortBuilder
	CreateSecurityDataZoneShortBuilder() SecurityDataZoneShortBuilder
}

// _SecurityDataZoneShort is the data-structure of this message
type _SecurityDataZoneShort struct {
	SecurityDataContract
	ZoneNumber uint8
}

var _ SecurityDataZoneShort = (*_SecurityDataZoneShort)(nil)
var _ SecurityDataRequirements = (*_SecurityDataZoneShort)(nil)

// NewSecurityDataZoneShort factory function for _SecurityDataZoneShort
func NewSecurityDataZoneShort(commandTypeContainer SecurityCommandTypeContainer, argument byte, zoneNumber uint8) *_SecurityDataZoneShort {
	_result := &_SecurityDataZoneShort{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ZoneNumber:           zoneNumber,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataZoneShortBuilder is a builder for SecurityDataZoneShort
type SecurityDataZoneShortBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneNumber uint8) SecurityDataZoneShortBuilder
	// WithZoneNumber adds ZoneNumber (property field)
	WithZoneNumber(uint8) SecurityDataZoneShortBuilder
	// Build builds the SecurityDataZoneShort or returns an error if something is wrong
	Build() (SecurityDataZoneShort, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataZoneShort
}

// NewSecurityDataZoneShortBuilder() creates a SecurityDataZoneShortBuilder
func NewSecurityDataZoneShortBuilder() SecurityDataZoneShortBuilder {
	return &_SecurityDataZoneShortBuilder{_SecurityDataZoneShort: new(_SecurityDataZoneShort)}
}

type _SecurityDataZoneShortBuilder struct {
	*_SecurityDataZoneShort

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataZoneShortBuilder) = (*_SecurityDataZoneShortBuilder)(nil)

func (b *_SecurityDataZoneShortBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
}

func (b *_SecurityDataZoneShortBuilder) WithMandatoryFields(zoneNumber uint8) SecurityDataZoneShortBuilder {
	return b.WithZoneNumber(zoneNumber)
}

func (b *_SecurityDataZoneShortBuilder) WithZoneNumber(zoneNumber uint8) SecurityDataZoneShortBuilder {
	b.ZoneNumber = zoneNumber
	return b
}

func (b *_SecurityDataZoneShortBuilder) Build() (SecurityDataZoneShort, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataZoneShort.deepCopy(), nil
}

func (b *_SecurityDataZoneShortBuilder) MustBuild() SecurityDataZoneShort {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SecurityDataZoneShortBuilder) Done() SecurityDataBuilder {
	return b.parentBuilder
}

func (b *_SecurityDataZoneShortBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataZoneShortBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataZoneShortBuilder().(*_SecurityDataZoneShortBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataZoneShortBuilder creates a SecurityDataZoneShortBuilder
func (b *_SecurityDataZoneShort) CreateSecurityDataZoneShortBuilder() SecurityDataZoneShortBuilder {
	if b == nil {
		return NewSecurityDataZoneShortBuilder()
	}
	return &_SecurityDataZoneShortBuilder{_SecurityDataZoneShort: b.deepCopy()}
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

func (m *_SecurityDataZoneShort) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneShort) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneShort(structType any) SecurityDataZoneShort {
	if casted, ok := structType.(SecurityDataZoneShort); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneShort); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneShort) GetTypeName() string {
	return "SecurityDataZoneShort"
}

func (m *_SecurityDataZoneShort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (zoneNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityDataZoneShort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataZoneShort) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataZoneShort SecurityDataZoneShort, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataZoneShort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneShort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneNumber, err := ReadSimpleField(ctx, "zoneNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneNumber' field"))
	}
	m.ZoneNumber = zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataZoneShort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneShort")
	}

	return m, nil
}

func (m *_SecurityDataZoneShort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataZoneShort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneShort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneShort")
		}

		if err := WriteSimpleField[uint8](ctx, "zoneNumber", m.GetZoneNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneNumber' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataZoneShort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneShort")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataZoneShort) IsSecurityDataZoneShort() {}

func (m *_SecurityDataZoneShort) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataZoneShort) deepCopy() *_SecurityDataZoneShort {
	if m == nil {
		return nil
	}
	_SecurityDataZoneShortCopy := &_SecurityDataZoneShort{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
		m.ZoneNumber,
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataZoneShortCopy
}

func (m *_SecurityDataZoneShort) String() string {
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
