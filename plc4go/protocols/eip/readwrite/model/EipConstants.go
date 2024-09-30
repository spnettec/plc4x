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

// Constant values.
const EipConstants_EIPUDPDISCOVERYDEFAULTPORT uint16 = uint16(44818)
const EipConstants_EIPTCPDEFAULTPORT uint16 = uint16(44818)

// EipConstants is the corresponding interface of EipConstants
type EipConstants interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsEipConstants is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEipConstants()
	// CreateBuilder creates a EipConstantsBuilder
	CreateEipConstantsBuilder() EipConstantsBuilder
}

// _EipConstants is the data-structure of this message
type _EipConstants struct {
}

var _ EipConstants = (*_EipConstants)(nil)

// NewEipConstants factory function for _EipConstants
func NewEipConstants() *_EipConstants {
	return &_EipConstants{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EipConstantsBuilder is a builder for EipConstants
type EipConstantsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() EipConstantsBuilder
	// Build builds the EipConstants or returns an error if something is wrong
	Build() (EipConstants, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EipConstants
}

// NewEipConstantsBuilder() creates a EipConstantsBuilder
func NewEipConstantsBuilder() EipConstantsBuilder {
	return &_EipConstantsBuilder{_EipConstants: new(_EipConstants)}
}

type _EipConstantsBuilder struct {
	*_EipConstants

	err *utils.MultiError
}

var _ (EipConstantsBuilder) = (*_EipConstantsBuilder)(nil)

func (b *_EipConstantsBuilder) WithMandatoryFields() EipConstantsBuilder {
	return b
}

func (b *_EipConstantsBuilder) Build() (EipConstants, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EipConstants.deepCopy(), nil
}

func (b *_EipConstantsBuilder) MustBuild() EipConstants {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EipConstantsBuilder) DeepCopy() any {
	_copy := b.CreateEipConstantsBuilder().(*_EipConstantsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEipConstantsBuilder creates a EipConstantsBuilder
func (b *_EipConstants) CreateEipConstantsBuilder() EipConstantsBuilder {
	if b == nil {
		return NewEipConstantsBuilder()
	}
	return &_EipConstantsBuilder{_EipConstants: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_EipConstants) GetEipUdpDiscoveryDefaultPort() uint16 {
	return EipConstants_EIPUDPDISCOVERYDEFAULTPORT
}

func (m *_EipConstants) GetEipTcpDefaultPort() uint16 {
	return EipConstants_EIPTCPDEFAULTPORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEipConstants(structType any) EipConstants {
	if casted, ok := structType.(EipConstants); ok {
		return casted
	}
	if casted, ok := structType.(*EipConstants); ok {
		return *casted
	}
	return nil
}

func (m *_EipConstants) GetTypeName() string {
	return "EipConstants"
}

func (m *_EipConstants) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (eipUdpDiscoveryDefaultPort)
	lengthInBits += 16

	// Const Field (eipTcpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_EipConstants) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EipConstantsParse(ctx context.Context, theBytes []byte) (EipConstants, error) {
	return EipConstantsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func EipConstantsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (EipConstants, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (EipConstants, error) {
		return EipConstantsParseWithBuffer(ctx, readBuffer)
	}
}

func EipConstantsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (EipConstants, error) {
	v, err := (&_EipConstants{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_EipConstants) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__eipConstants EipConstants, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	eipUdpDiscoveryDefaultPort, err := ReadConstField[uint16](ctx, "eipUdpDiscoveryDefaultPort", ReadUnsignedShort(readBuffer, uint8(16)), EipConstants_EIPUDPDISCOVERYDEFAULTPORT)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eipUdpDiscoveryDefaultPort' field"))
	}
	_ = eipUdpDiscoveryDefaultPort

	eipTcpDefaultPort, err := ReadConstField[uint16](ctx, "eipTcpDefaultPort", ReadUnsignedShort(readBuffer, uint8(16)), EipConstants_EIPTCPDEFAULTPORT)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eipTcpDefaultPort' field"))
	}
	_ = eipTcpDefaultPort

	if closeErr := readBuffer.CloseContext("EipConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipConstants")
	}

	return m, nil
}

func (m *_EipConstants) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EipConstants) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("EipConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for EipConstants")
	}

	if err := WriteConstField(ctx, "eipUdpDiscoveryDefaultPort", EipConstants_EIPUDPDISCOVERYDEFAULTPORT, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'eipUdpDiscoveryDefaultPort' field")
	}

	if err := WriteConstField(ctx, "eipTcpDefaultPort", EipConstants_EIPTCPDEFAULTPORT, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'eipTcpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("EipConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for EipConstants")
	}
	return nil
}

func (m *_EipConstants) IsEipConstants() {}

func (m *_EipConstants) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EipConstants) deepCopy() *_EipConstants {
	if m == nil {
		return nil
	}
	_EipConstantsCopy := &_EipConstants{}
	return _EipConstantsCopy
}

func (m *_EipConstants) String() string {
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
