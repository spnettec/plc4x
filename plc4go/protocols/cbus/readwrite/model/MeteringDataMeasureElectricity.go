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

// MeteringDataMeasureElectricity is the corresponding interface of MeteringDataMeasureElectricity
type MeteringDataMeasureElectricity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MeteringData
	// IsMeteringDataMeasureElectricity is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMeteringDataMeasureElectricity()
	// CreateBuilder creates a MeteringDataMeasureElectricityBuilder
	CreateMeteringDataMeasureElectricityBuilder() MeteringDataMeasureElectricityBuilder
}

// _MeteringDataMeasureElectricity is the data-structure of this message
type _MeteringDataMeasureElectricity struct {
	MeteringDataContract
}

var _ MeteringDataMeasureElectricity = (*_MeteringDataMeasureElectricity)(nil)
var _ MeteringDataRequirements = (*_MeteringDataMeasureElectricity)(nil)

// NewMeteringDataMeasureElectricity factory function for _MeteringDataMeasureElectricity
func NewMeteringDataMeasureElectricity(commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataMeasureElectricity {
	_result := &_MeteringDataMeasureElectricity{
		MeteringDataContract: NewMeteringData(commandTypeContainer, argument),
	}
	_result.MeteringDataContract.(*_MeteringData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MeteringDataMeasureElectricityBuilder is a builder for MeteringDataMeasureElectricity
type MeteringDataMeasureElectricityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MeteringDataMeasureElectricityBuilder
	// Build builds the MeteringDataMeasureElectricity or returns an error if something is wrong
	Build() (MeteringDataMeasureElectricity, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MeteringDataMeasureElectricity
}

// NewMeteringDataMeasureElectricityBuilder() creates a MeteringDataMeasureElectricityBuilder
func NewMeteringDataMeasureElectricityBuilder() MeteringDataMeasureElectricityBuilder {
	return &_MeteringDataMeasureElectricityBuilder{_MeteringDataMeasureElectricity: new(_MeteringDataMeasureElectricity)}
}

type _MeteringDataMeasureElectricityBuilder struct {
	*_MeteringDataMeasureElectricity

	parentBuilder *_MeteringDataBuilder

	err *utils.MultiError
}

var _ (MeteringDataMeasureElectricityBuilder) = (*_MeteringDataMeasureElectricityBuilder)(nil)

func (b *_MeteringDataMeasureElectricityBuilder) setParent(contract MeteringDataContract) {
	b.MeteringDataContract = contract
}

func (b *_MeteringDataMeasureElectricityBuilder) WithMandatoryFields() MeteringDataMeasureElectricityBuilder {
	return b
}

func (b *_MeteringDataMeasureElectricityBuilder) Build() (MeteringDataMeasureElectricity, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MeteringDataMeasureElectricity.deepCopy(), nil
}

func (b *_MeteringDataMeasureElectricityBuilder) MustBuild() MeteringDataMeasureElectricity {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MeteringDataMeasureElectricityBuilder) Done() MeteringDataBuilder {
	return b.parentBuilder
}

func (b *_MeteringDataMeasureElectricityBuilder) buildForMeteringData() (MeteringData, error) {
	return b.Build()
}

func (b *_MeteringDataMeasureElectricityBuilder) DeepCopy() any {
	_copy := b.CreateMeteringDataMeasureElectricityBuilder().(*_MeteringDataMeasureElectricityBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMeteringDataMeasureElectricityBuilder creates a MeteringDataMeasureElectricityBuilder
func (b *_MeteringDataMeasureElectricity) CreateMeteringDataMeasureElectricityBuilder() MeteringDataMeasureElectricityBuilder {
	if b == nil {
		return NewMeteringDataMeasureElectricityBuilder()
	}
	return &_MeteringDataMeasureElectricityBuilder{_MeteringDataMeasureElectricity: b.deepCopy()}
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

func (m *_MeteringDataMeasureElectricity) GetParent() MeteringDataContract {
	return m.MeteringDataContract
}

// Deprecated: use the interface for direct cast
func CastMeteringDataMeasureElectricity(structType any) MeteringDataMeasureElectricity {
	if casted, ok := structType.(MeteringDataMeasureElectricity); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataMeasureElectricity); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataMeasureElectricity) GetTypeName() string {
	return "MeteringDataMeasureElectricity"
}

func (m *_MeteringDataMeasureElectricity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MeteringDataContract.(*_MeteringData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MeteringDataMeasureElectricity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MeteringDataMeasureElectricity) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MeteringData) (__meteringDataMeasureElectricity MeteringDataMeasureElectricity, err error) {
	m.MeteringDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataMeasureElectricity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataMeasureElectricity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MeteringDataMeasureElectricity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataMeasureElectricity")
	}

	return m, nil
}

func (m *_MeteringDataMeasureElectricity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataMeasureElectricity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataMeasureElectricity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataMeasureElectricity")
		}

		if popErr := writeBuffer.PopContext("MeteringDataMeasureElectricity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataMeasureElectricity")
		}
		return nil
	}
	return m.MeteringDataContract.(*_MeteringData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataMeasureElectricity) IsMeteringDataMeasureElectricity() {}

func (m *_MeteringDataMeasureElectricity) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MeteringDataMeasureElectricity) deepCopy() *_MeteringDataMeasureElectricity {
	if m == nil {
		return nil
	}
	_MeteringDataMeasureElectricityCopy := &_MeteringDataMeasureElectricity{
		m.MeteringDataContract.(*_MeteringData).deepCopy(),
	}
	m.MeteringDataContract.(*_MeteringData)._SubType = m
	return _MeteringDataMeasureElectricityCopy
}

func (m *_MeteringDataMeasureElectricity) String() string {
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
