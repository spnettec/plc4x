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

// TunnelingRequest is the corresponding interface of TunnelingRequest
type TunnelingRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetTunnelingRequestDataBlock returns TunnelingRequestDataBlock (property field)
	GetTunnelingRequestDataBlock() TunnelingRequestDataBlock
	// GetCemi returns Cemi (property field)
	GetCemi() CEMI
	// IsTunnelingRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTunnelingRequest()
	// CreateBuilder creates a TunnelingRequestBuilder
	CreateTunnelingRequestBuilder() TunnelingRequestBuilder
}

// _TunnelingRequest is the data-structure of this message
type _TunnelingRequest struct {
	KnxNetIpMessageContract
	TunnelingRequestDataBlock TunnelingRequestDataBlock
	Cemi                      CEMI

	// Arguments.
	TotalLength uint16
}

var _ TunnelingRequest = (*_TunnelingRequest)(nil)
var _ KnxNetIpMessageRequirements = (*_TunnelingRequest)(nil)

// NewTunnelingRequest factory function for _TunnelingRequest
func NewTunnelingRequest(tunnelingRequestDataBlock TunnelingRequestDataBlock, cemi CEMI, totalLength uint16) *_TunnelingRequest {
	if tunnelingRequestDataBlock == nil {
		panic("tunnelingRequestDataBlock of type TunnelingRequestDataBlock for TunnelingRequest must not be nil")
	}
	if cemi == nil {
		panic("cemi of type CEMI for TunnelingRequest must not be nil")
	}
	_result := &_TunnelingRequest{
		KnxNetIpMessageContract:   NewKnxNetIpMessage(),
		TunnelingRequestDataBlock: tunnelingRequestDataBlock,
		Cemi:                      cemi,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TunnelingRequestBuilder is a builder for TunnelingRequest
type TunnelingRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(tunnelingRequestDataBlock TunnelingRequestDataBlock, cemi CEMI) TunnelingRequestBuilder
	// WithTunnelingRequestDataBlock adds TunnelingRequestDataBlock (property field)
	WithTunnelingRequestDataBlock(TunnelingRequestDataBlock) TunnelingRequestBuilder
	// WithTunnelingRequestDataBlockBuilder adds TunnelingRequestDataBlock (property field) which is build by the builder
	WithTunnelingRequestDataBlockBuilder(func(TunnelingRequestDataBlockBuilder) TunnelingRequestDataBlockBuilder) TunnelingRequestBuilder
	// WithCemi adds Cemi (property field)
	WithCemi(CEMI) TunnelingRequestBuilder
	// WithCemiBuilder adds Cemi (property field) which is build by the builder
	WithCemiBuilder(func(CEMIBuilder) CEMIBuilder) TunnelingRequestBuilder
	// Build builds the TunnelingRequest or returns an error if something is wrong
	Build() (TunnelingRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TunnelingRequest
}

// NewTunnelingRequestBuilder() creates a TunnelingRequestBuilder
func NewTunnelingRequestBuilder() TunnelingRequestBuilder {
	return &_TunnelingRequestBuilder{_TunnelingRequest: new(_TunnelingRequest)}
}

type _TunnelingRequestBuilder struct {
	*_TunnelingRequest

	parentBuilder *_KnxNetIpMessageBuilder

	err *utils.MultiError
}

var _ (TunnelingRequestBuilder) = (*_TunnelingRequestBuilder)(nil)

func (b *_TunnelingRequestBuilder) setParent(contract KnxNetIpMessageContract) {
	b.KnxNetIpMessageContract = contract
}

func (b *_TunnelingRequestBuilder) WithMandatoryFields(tunnelingRequestDataBlock TunnelingRequestDataBlock, cemi CEMI) TunnelingRequestBuilder {
	return b.WithTunnelingRequestDataBlock(tunnelingRequestDataBlock).WithCemi(cemi)
}

func (b *_TunnelingRequestBuilder) WithTunnelingRequestDataBlock(tunnelingRequestDataBlock TunnelingRequestDataBlock) TunnelingRequestBuilder {
	b.TunnelingRequestDataBlock = tunnelingRequestDataBlock
	return b
}

func (b *_TunnelingRequestBuilder) WithTunnelingRequestDataBlockBuilder(builderSupplier func(TunnelingRequestDataBlockBuilder) TunnelingRequestDataBlockBuilder) TunnelingRequestBuilder {
	builder := builderSupplier(b.TunnelingRequestDataBlock.CreateTunnelingRequestDataBlockBuilder())
	var err error
	b.TunnelingRequestDataBlock, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "TunnelingRequestDataBlockBuilder failed"))
	}
	return b
}

func (b *_TunnelingRequestBuilder) WithCemi(cemi CEMI) TunnelingRequestBuilder {
	b.Cemi = cemi
	return b
}

func (b *_TunnelingRequestBuilder) WithCemiBuilder(builderSupplier func(CEMIBuilder) CEMIBuilder) TunnelingRequestBuilder {
	builder := builderSupplier(b.Cemi.CreateCEMIBuilder())
	var err error
	b.Cemi, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CEMIBuilder failed"))
	}
	return b
}

func (b *_TunnelingRequestBuilder) Build() (TunnelingRequest, error) {
	if b.TunnelingRequestDataBlock == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'tunnelingRequestDataBlock' not set"))
	}
	if b.Cemi == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'cemi' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TunnelingRequest.deepCopy(), nil
}

func (b *_TunnelingRequestBuilder) MustBuild() TunnelingRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_TunnelingRequestBuilder) Done() KnxNetIpMessageBuilder {
	return b.parentBuilder
}

func (b *_TunnelingRequestBuilder) buildForKnxNetIpMessage() (KnxNetIpMessage, error) {
	return b.Build()
}

func (b *_TunnelingRequestBuilder) DeepCopy() any {
	_copy := b.CreateTunnelingRequestBuilder().(*_TunnelingRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTunnelingRequestBuilder creates a TunnelingRequestBuilder
func (b *_TunnelingRequest) CreateTunnelingRequestBuilder() TunnelingRequestBuilder {
	if b == nil {
		return NewTunnelingRequestBuilder()
	}
	return &_TunnelingRequestBuilder{_TunnelingRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TunnelingRequest) GetMsgType() uint16 {
	return 0x0420
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TunnelingRequest) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TunnelingRequest) GetTunnelingRequestDataBlock() TunnelingRequestDataBlock {
	return m.TunnelingRequestDataBlock
}

func (m *_TunnelingRequest) GetCemi() CEMI {
	return m.Cemi
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTunnelingRequest(structType any) TunnelingRequest {
	if casted, ok := structType.(TunnelingRequest); ok {
		return casted
	}
	if casted, ok := structType.(*TunnelingRequest); ok {
		return *casted
	}
	return nil
}

func (m *_TunnelingRequest) GetTypeName() string {
	return "TunnelingRequest"
}

func (m *_TunnelingRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (tunnelingRequestDataBlock)
	lengthInBits += m.TunnelingRequestDataBlock.GetLengthInBits(ctx)

	// Simple field (cemi)
	lengthInBits += m.Cemi.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_TunnelingRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TunnelingRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage, totalLength uint16) (__tunnelingRequest TunnelingRequest, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TunnelingRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TunnelingRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	tunnelingRequestDataBlock, err := ReadSimpleField[TunnelingRequestDataBlock](ctx, "tunnelingRequestDataBlock", ReadComplex[TunnelingRequestDataBlock](TunnelingRequestDataBlockParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tunnelingRequestDataBlock' field"))
	}
	m.TunnelingRequestDataBlock = tunnelingRequestDataBlock

	cemi, err := ReadSimpleField[CEMI](ctx, "cemi", ReadComplex[CEMI](CEMIParseWithBufferProducer[CEMI]((uint16)(uint16(totalLength)-uint16((uint16(uint16(6))+uint16(tunnelingRequestDataBlock.GetLengthInBytes(ctx)))))), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cemi' field"))
	}
	m.Cemi = cemi

	if closeErr := readBuffer.CloseContext("TunnelingRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TunnelingRequest")
	}

	return m, nil
}

func (m *_TunnelingRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TunnelingRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TunnelingRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TunnelingRequest")
		}

		if err := WriteSimpleField[TunnelingRequestDataBlock](ctx, "tunnelingRequestDataBlock", m.GetTunnelingRequestDataBlock(), WriteComplex[TunnelingRequestDataBlock](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'tunnelingRequestDataBlock' field")
		}

		if err := WriteSimpleField[CEMI](ctx, "cemi", m.GetCemi(), WriteComplex[CEMI](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'cemi' field")
		}

		if popErr := writeBuffer.PopContext("TunnelingRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TunnelingRequest")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_TunnelingRequest) GetTotalLength() uint16 {
	return m.TotalLength
}

//
////

func (m *_TunnelingRequest) IsTunnelingRequest() {}

func (m *_TunnelingRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TunnelingRequest) deepCopy() *_TunnelingRequest {
	if m == nil {
		return nil
	}
	_TunnelingRequestCopy := &_TunnelingRequest{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		m.TunnelingRequestDataBlock.DeepCopy().(TunnelingRequestDataBlock),
		m.Cemi.DeepCopy().(CEMI),
		m.TotalLength,
	}
	m.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _TunnelingRequestCopy
}

func (m *_TunnelingRequest) String() string {
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
