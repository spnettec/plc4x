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

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AmsString is the corresponding interface of AmsString
type AmsString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetText returns Text (property field)
	GetText() string
	// IsAmsString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAmsString()
	// CreateBuilder creates a AmsStringBuilder
	CreateAmsStringBuilder() AmsStringBuilder
}

// _AmsString is the data-structure of this message
type _AmsString struct {
	Text string
	// Reserved Fields
	reservedField0 *uint8
}

var _ AmsString = (*_AmsString)(nil)

// NewAmsString factory function for _AmsString
func NewAmsString(text string) *_AmsString {
	return &_AmsString{Text: text}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AmsStringBuilder is a builder for AmsString
type AmsStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(text string) AmsStringBuilder
	// WithText adds Text (property field)
	WithText(string) AmsStringBuilder
	// Build builds the AmsString or returns an error if something is wrong
	Build() (AmsString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AmsString
}

// NewAmsStringBuilder() creates a AmsStringBuilder
func NewAmsStringBuilder() AmsStringBuilder {
	return &_AmsStringBuilder{_AmsString: new(_AmsString)}
}

type _AmsStringBuilder struct {
	*_AmsString

	err *utils.MultiError
}

var _ (AmsStringBuilder) = (*_AmsStringBuilder)(nil)

func (b *_AmsStringBuilder) WithMandatoryFields(text string) AmsStringBuilder {
	return b.WithText(text)
}

func (b *_AmsStringBuilder) WithText(text string) AmsStringBuilder {
	b.Text = text
	return b
}

func (b *_AmsStringBuilder) Build() (AmsString, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AmsString.deepCopy(), nil
}

func (b *_AmsStringBuilder) MustBuild() AmsString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AmsStringBuilder) DeepCopy() any {
	_copy := b.CreateAmsStringBuilder().(*_AmsStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAmsStringBuilder creates a AmsStringBuilder
func (b *_AmsString) CreateAmsStringBuilder() AmsStringBuilder {
	if b == nil {
		return NewAmsStringBuilder()
	}
	return &_AmsStringBuilder{_AmsString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsString) GetText() string {
	return m.Text
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAmsString(structType any) AmsString {
	if casted, ok := structType.(AmsString); ok {
		return casted
	}
	if casted, ok := structType.(*AmsString); ok {
		return *casted
	}
	return nil
}

func (m *_AmsString) GetTypeName() string {
	return "AmsString"
}

func (m *_AmsString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (strLen)
	lengthInBits += 16

	// Simple field (text)
	lengthInBits += uint16(int32(int32(8)) * int32((int32(uint16(uint16(len(m.GetText())))+uint16(uint16(1))) - int32(int32(1)))))

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AmsString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AmsStringParse(ctx context.Context, theBytes []byte) (AmsString, error) {
	return AmsStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AmsStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AmsString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AmsString, error) {
		return AmsStringParseWithBuffer(ctx, readBuffer)
	}
}

func AmsStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AmsString, error) {
	v, err := (&_AmsString{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AmsString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__amsString AmsString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AmsString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	strLen, err := ReadImplicitField[uint16](ctx, "strLen", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'strLen' field"))
	}
	_ = strLen

	text, err := ReadSimpleField(ctx, "text", ReadString(readBuffer, uint32(int32(int32(8))*int32((int32(strLen)-int32(int32(1)))))), codegen.WithEncoding("UTF-8"))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'text' field"))
	}
	m.Text = text

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("AmsString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsString")
	}

	return m, nil
}

func (m *_AmsString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AmsString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AmsString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsString")
	}
	strLen := uint16(uint16(uint16(len(m.GetText()))) + uint16(uint16(1)))
	if err := WriteImplicitField(ctx, "strLen", strLen, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'strLen' field")
	}

	if err := WriteSimpleField[string](ctx, "text", m.GetText(), WriteString(writeBuffer, int32(int32(int32(8))*int32((int32(uint16(uint16(len(m.GetText())))+uint16(uint16(1)))-int32(int32(1)))))), codegen.WithEncoding("UTF-8")); err != nil {
		return errors.Wrap(err, "Error serializing 'text' field")
	}

	if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if popErr := writeBuffer.PopContext("AmsString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsString")
	}
	return nil
}

func (m *_AmsString) IsAmsString() {}

func (m *_AmsString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AmsString) deepCopy() *_AmsString {
	if m == nil {
		return nil
	}
	_AmsStringCopy := &_AmsString{
		m.Text,
		m.reservedField0,
	}
	return _AmsStringCopy
}

func (m *_AmsString) String() string {
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
