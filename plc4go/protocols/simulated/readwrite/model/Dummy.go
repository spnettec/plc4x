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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Dummy is the corresponding interface of Dummy
type Dummy interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetDummy returns Dummy (property field)
	GetDummy() uint16
	// IsDummy is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDummy()
	// CreateBuilder creates a DummyBuilder
	CreateDummyBuilder() DummyBuilder
}

// _Dummy is the data-structure of this message
type _Dummy struct {
	Dummy uint16
}

var _ Dummy = (*_Dummy)(nil)

// NewDummy factory function for _Dummy
func NewDummy(dummy uint16) *_Dummy {
	return &_Dummy{Dummy: dummy}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DummyBuilder is a builder for Dummy
type DummyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dummy uint16) DummyBuilder
	// WithDummy adds Dummy (property field)
	WithDummy(uint16) DummyBuilder
	// Build builds the Dummy or returns an error if something is wrong
	Build() (Dummy, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Dummy
}

// NewDummyBuilder() creates a DummyBuilder
func NewDummyBuilder() DummyBuilder {
	return &_DummyBuilder{_Dummy: new(_Dummy)}
}

type _DummyBuilder struct {
	*_Dummy

	err *utils.MultiError
}

var _ (DummyBuilder) = (*_DummyBuilder)(nil)

func (b *_DummyBuilder) WithMandatoryFields(dummy uint16) DummyBuilder {
	return b.WithDummy(dummy)
}

func (b *_DummyBuilder) WithDummy(dummy uint16) DummyBuilder {
	b.Dummy = dummy
	return b
}

func (b *_DummyBuilder) Build() (Dummy, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._Dummy.deepCopy(), nil
}

func (b *_DummyBuilder) MustBuild() Dummy {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DummyBuilder) DeepCopy() any {
	_copy := b.CreateDummyBuilder().(*_DummyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDummyBuilder creates a DummyBuilder
func (b *_Dummy) CreateDummyBuilder() DummyBuilder {
	if b == nil {
		return NewDummyBuilder()
	}
	return &_DummyBuilder{_Dummy: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Dummy) GetDummy() uint16 {
	return m.Dummy
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDummy(structType any) Dummy {
	if casted, ok := structType.(Dummy); ok {
		return casted
	}
	if casted, ok := structType.(*Dummy); ok {
		return *casted
	}
	return nil
}

func (m *_Dummy) GetTypeName() string {
	return "Dummy"
}

func (m *_Dummy) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dummy)
	lengthInBits += 16

	return lengthInBits
}

func (m *_Dummy) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DummyParse(ctx context.Context, theBytes []byte) (Dummy, error) {
	return DummyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func DummyParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (Dummy, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Dummy, error) {
		return DummyParseWithBuffer(ctx, readBuffer)
	}
}

func DummyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Dummy, error) {
	v, err := (&_Dummy{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_Dummy) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__dummy Dummy, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Dummy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Dummy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dummy, err := ReadSimpleField(ctx, "dummy", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dummy' field"))
	}
	m.Dummy = dummy

	if closeErr := readBuffer.CloseContext("Dummy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Dummy")
	}

	return m, nil
}

func (m *_Dummy) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Dummy) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Dummy"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Dummy")
	}

	if err := WriteSimpleField[uint16](ctx, "dummy", m.GetDummy(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'dummy' field")
	}

	if popErr := writeBuffer.PopContext("Dummy"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Dummy")
	}
	return nil
}

func (m *_Dummy) IsDummy() {}

func (m *_Dummy) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Dummy) deepCopy() *_Dummy {
	if m == nil {
		return nil
	}
	_DummyCopy := &_Dummy{
		m.Dummy,
	}
	return _DummyCopy
}

func (m *_Dummy) String() string {
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
