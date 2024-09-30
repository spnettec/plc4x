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

// SysexCommandSamplingInterval is the corresponding interface of SysexCommandSamplingInterval
type SysexCommandSamplingInterval interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SysexCommand
	// IsSysexCommandSamplingInterval is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandSamplingInterval()
	// CreateBuilder creates a SysexCommandSamplingIntervalBuilder
	CreateSysexCommandSamplingIntervalBuilder() SysexCommandSamplingIntervalBuilder
}

// _SysexCommandSamplingInterval is the data-structure of this message
type _SysexCommandSamplingInterval struct {
	SysexCommandContract
}

var _ SysexCommandSamplingInterval = (*_SysexCommandSamplingInterval)(nil)
var _ SysexCommandRequirements = (*_SysexCommandSamplingInterval)(nil)

// NewSysexCommandSamplingInterval factory function for _SysexCommandSamplingInterval
func NewSysexCommandSamplingInterval() *_SysexCommandSamplingInterval {
	_result := &_SysexCommandSamplingInterval{
		SysexCommandContract: NewSysexCommand(),
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SysexCommandSamplingIntervalBuilder is a builder for SysexCommandSamplingInterval
type SysexCommandSamplingIntervalBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SysexCommandSamplingIntervalBuilder
	// Build builds the SysexCommandSamplingInterval or returns an error if something is wrong
	Build() (SysexCommandSamplingInterval, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SysexCommandSamplingInterval
}

// NewSysexCommandSamplingIntervalBuilder() creates a SysexCommandSamplingIntervalBuilder
func NewSysexCommandSamplingIntervalBuilder() SysexCommandSamplingIntervalBuilder {
	return &_SysexCommandSamplingIntervalBuilder{_SysexCommandSamplingInterval: new(_SysexCommandSamplingInterval)}
}

type _SysexCommandSamplingIntervalBuilder struct {
	*_SysexCommandSamplingInterval

	parentBuilder *_SysexCommandBuilder

	err *utils.MultiError
}

var _ (SysexCommandSamplingIntervalBuilder) = (*_SysexCommandSamplingIntervalBuilder)(nil)

func (b *_SysexCommandSamplingIntervalBuilder) setParent(contract SysexCommandContract) {
	b.SysexCommandContract = contract
}

func (b *_SysexCommandSamplingIntervalBuilder) WithMandatoryFields() SysexCommandSamplingIntervalBuilder {
	return b
}

func (b *_SysexCommandSamplingIntervalBuilder) Build() (SysexCommandSamplingInterval, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SysexCommandSamplingInterval.deepCopy(), nil
}

func (b *_SysexCommandSamplingIntervalBuilder) MustBuild() SysexCommandSamplingInterval {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SysexCommandSamplingIntervalBuilder) Done() SysexCommandBuilder {
	return b.parentBuilder
}

func (b *_SysexCommandSamplingIntervalBuilder) buildForSysexCommand() (SysexCommand, error) {
	return b.Build()
}

func (b *_SysexCommandSamplingIntervalBuilder) DeepCopy() any {
	_copy := b.CreateSysexCommandSamplingIntervalBuilder().(*_SysexCommandSamplingIntervalBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSysexCommandSamplingIntervalBuilder creates a SysexCommandSamplingIntervalBuilder
func (b *_SysexCommandSamplingInterval) CreateSysexCommandSamplingIntervalBuilder() SysexCommandSamplingIntervalBuilder {
	if b == nil {
		return NewSysexCommandSamplingIntervalBuilder()
	}
	return &_SysexCommandSamplingIntervalBuilder{_SysexCommandSamplingInterval: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandSamplingInterval) GetCommandType() uint8 {
	return 0x7A
}

func (m *_SysexCommandSamplingInterval) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandSamplingInterval) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

// Deprecated: use the interface for direct cast
func CastSysexCommandSamplingInterval(structType any) SysexCommandSamplingInterval {
	if casted, ok := structType.(SysexCommandSamplingInterval); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandSamplingInterval); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandSamplingInterval) GetTypeName() string {
	return "SysexCommandSamplingInterval"
}

func (m *_SysexCommandSamplingInterval) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandSamplingInterval) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandSamplingInterval) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandSamplingInterval SysexCommandSamplingInterval, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandSamplingInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandSamplingInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandSamplingInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandSamplingInterval")
	}

	return m, nil
}

func (m *_SysexCommandSamplingInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandSamplingInterval) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSamplingInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandSamplingInterval")
		}

		if popErr := writeBuffer.PopContext("SysexCommandSamplingInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandSamplingInterval")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandSamplingInterval) IsSysexCommandSamplingInterval() {}

func (m *_SysexCommandSamplingInterval) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SysexCommandSamplingInterval) deepCopy() *_SysexCommandSamplingInterval {
	if m == nil {
		return nil
	}
	_SysexCommandSamplingIntervalCopy := &_SysexCommandSamplingInterval{
		m.SysexCommandContract.(*_SysexCommand).deepCopy(),
	}
	m.SysexCommandContract.(*_SysexCommand)._SubType = m
	return _SysexCommandSamplingIntervalCopy
}

func (m *_SysexCommandSamplingInterval) String() string {
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
