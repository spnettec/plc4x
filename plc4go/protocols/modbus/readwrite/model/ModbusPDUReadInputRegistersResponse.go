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

// ModbusPDUReadInputRegistersResponse is the corresponding interface of ModbusPDUReadInputRegistersResponse
type ModbusPDUReadInputRegistersResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
	// IsModbusPDUReadInputRegistersResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadInputRegistersResponse()
	// CreateBuilder creates a ModbusPDUReadInputRegistersResponseBuilder
	CreateModbusPDUReadInputRegistersResponseBuilder() ModbusPDUReadInputRegistersResponseBuilder
}

// _ModbusPDUReadInputRegistersResponse is the data-structure of this message
type _ModbusPDUReadInputRegistersResponse struct {
	ModbusPDUContract
	Value []byte
}

var _ ModbusPDUReadInputRegistersResponse = (*_ModbusPDUReadInputRegistersResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadInputRegistersResponse)(nil)

// NewModbusPDUReadInputRegistersResponse factory function for _ModbusPDUReadInputRegistersResponse
func NewModbusPDUReadInputRegistersResponse(value []byte) *_ModbusPDUReadInputRegistersResponse {
	_result := &_ModbusPDUReadInputRegistersResponse{
		ModbusPDUContract: NewModbusPDU(),
		Value:             value,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUReadInputRegistersResponseBuilder is a builder for ModbusPDUReadInputRegistersResponse
type ModbusPDUReadInputRegistersResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []byte) ModbusPDUReadInputRegistersResponseBuilder
	// WithValue adds Value (property field)
	WithValue(...byte) ModbusPDUReadInputRegistersResponseBuilder
	// Build builds the ModbusPDUReadInputRegistersResponse or returns an error if something is wrong
	Build() (ModbusPDUReadInputRegistersResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUReadInputRegistersResponse
}

// NewModbusPDUReadInputRegistersResponseBuilder() creates a ModbusPDUReadInputRegistersResponseBuilder
func NewModbusPDUReadInputRegistersResponseBuilder() ModbusPDUReadInputRegistersResponseBuilder {
	return &_ModbusPDUReadInputRegistersResponseBuilder{_ModbusPDUReadInputRegistersResponse: new(_ModbusPDUReadInputRegistersResponse)}
}

type _ModbusPDUReadInputRegistersResponseBuilder struct {
	*_ModbusPDUReadInputRegistersResponse

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUReadInputRegistersResponseBuilder) = (*_ModbusPDUReadInputRegistersResponseBuilder)(nil)

func (b *_ModbusPDUReadInputRegistersResponseBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
}

func (b *_ModbusPDUReadInputRegistersResponseBuilder) WithMandatoryFields(value []byte) ModbusPDUReadInputRegistersResponseBuilder {
	return b.WithValue(value...)
}

func (b *_ModbusPDUReadInputRegistersResponseBuilder) WithValue(value ...byte) ModbusPDUReadInputRegistersResponseBuilder {
	b.Value = value
	return b
}

func (b *_ModbusPDUReadInputRegistersResponseBuilder) Build() (ModbusPDUReadInputRegistersResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUReadInputRegistersResponse.deepCopy(), nil
}

func (b *_ModbusPDUReadInputRegistersResponseBuilder) MustBuild() ModbusPDUReadInputRegistersResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ModbusPDUReadInputRegistersResponseBuilder) Done() ModbusPDUBuilder {
	return b.parentBuilder
}

func (b *_ModbusPDUReadInputRegistersResponseBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUReadInputRegistersResponseBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUReadInputRegistersResponseBuilder().(*_ModbusPDUReadInputRegistersResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUReadInputRegistersResponseBuilder creates a ModbusPDUReadInputRegistersResponseBuilder
func (b *_ModbusPDUReadInputRegistersResponse) CreateModbusPDUReadInputRegistersResponseBuilder() ModbusPDUReadInputRegistersResponseBuilder {
	if b == nil {
		return NewModbusPDUReadInputRegistersResponseBuilder()
	}
	return &_ModbusPDUReadInputRegistersResponseBuilder{_ModbusPDUReadInputRegistersResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadInputRegistersResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadInputRegistersResponse) GetFunctionFlag() uint8 {
	return 0x04
}

func (m *_ModbusPDUReadInputRegistersResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadInputRegistersResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadInputRegistersResponse) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadInputRegistersResponse(structType any) ModbusPDUReadInputRegistersResponse {
	if casted, ok := structType.(ModbusPDUReadInputRegistersResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadInputRegistersResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadInputRegistersResponse) GetTypeName() string {
	return "ModbusPDUReadInputRegistersResponse"
}

func (m *_ModbusPDUReadInputRegistersResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadInputRegistersResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadInputRegistersResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadInputRegistersResponse ModbusPDUReadInputRegistersResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadInputRegistersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadInputRegistersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	value, err := readBuffer.ReadByteArray("value", int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ModbusPDUReadInputRegistersResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadInputRegistersResponse")
	}

	return m, nil
}

func (m *_ModbusPDUReadInputRegistersResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadInputRegistersResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadInputRegistersResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadInputRegistersResponse")
		}
		byteCount := uint8(uint8(len(m.GetValue())))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteByteArrayField(ctx, "value", m.GetValue(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadInputRegistersResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadInputRegistersResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadInputRegistersResponse) IsModbusPDUReadInputRegistersResponse() {}

func (m *_ModbusPDUReadInputRegistersResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadInputRegistersResponse) deepCopy() *_ModbusPDUReadInputRegistersResponse {
	if m == nil {
		return nil
	}
	_ModbusPDUReadInputRegistersResponseCopy := &_ModbusPDUReadInputRegistersResponse{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.Value),
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReadInputRegistersResponseCopy
}

func (m *_ModbusPDUReadInputRegistersResponse) String() string {
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
