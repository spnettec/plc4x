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

// AdsReadResponse is the corresponding interface of AdsReadResponse
type AdsReadResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetData returns Data (property field)
	GetData() []byte
	// IsAdsReadResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsReadResponse()
	// CreateBuilder creates a AdsReadResponseBuilder
	CreateAdsReadResponseBuilder() AdsReadResponseBuilder
}

// _AdsReadResponse is the data-structure of this message
type _AdsReadResponse struct {
	AmsPacketContract
	Result ReturnCode
	Data   []byte
}

var _ AdsReadResponse = (*_AdsReadResponse)(nil)
var _ AmsPacketRequirements = (*_AdsReadResponse)(nil)

// NewAdsReadResponse factory function for _AdsReadResponse
func NewAdsReadResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32, result ReturnCode, data []byte) *_AdsReadResponse {
	_result := &_AdsReadResponse{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		Result:            result,
		Data:              data,
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsReadResponseBuilder is a builder for AdsReadResponse
type AdsReadResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(result ReturnCode, data []byte) AdsReadResponseBuilder
	// WithResult adds Result (property field)
	WithResult(ReturnCode) AdsReadResponseBuilder
	// WithData adds Data (property field)
	WithData(...byte) AdsReadResponseBuilder
	// Build builds the AdsReadResponse or returns an error if something is wrong
	Build() (AdsReadResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsReadResponse
}

// NewAdsReadResponseBuilder() creates a AdsReadResponseBuilder
func NewAdsReadResponseBuilder() AdsReadResponseBuilder {
	return &_AdsReadResponseBuilder{_AdsReadResponse: new(_AdsReadResponse)}
}

type _AdsReadResponseBuilder struct {
	*_AdsReadResponse

	parentBuilder *_AmsPacketBuilder

	err *utils.MultiError
}

var _ (AdsReadResponseBuilder) = (*_AdsReadResponseBuilder)(nil)

func (b *_AdsReadResponseBuilder) setParent(contract AmsPacketContract) {
	b.AmsPacketContract = contract
}

func (b *_AdsReadResponseBuilder) WithMandatoryFields(result ReturnCode, data []byte) AdsReadResponseBuilder {
	return b.WithResult(result).WithData(data...)
}

func (b *_AdsReadResponseBuilder) WithResult(result ReturnCode) AdsReadResponseBuilder {
	b.Result = result
	return b
}

func (b *_AdsReadResponseBuilder) WithData(data ...byte) AdsReadResponseBuilder {
	b.Data = data
	return b
}

func (b *_AdsReadResponseBuilder) Build() (AdsReadResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsReadResponse.deepCopy(), nil
}

func (b *_AdsReadResponseBuilder) MustBuild() AdsReadResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AdsReadResponseBuilder) Done() AmsPacketBuilder {
	return b.parentBuilder
}

func (b *_AdsReadResponseBuilder) buildForAmsPacket() (AmsPacket, error) {
	return b.Build()
}

func (b *_AdsReadResponseBuilder) DeepCopy() any {
	_copy := b.CreateAdsReadResponseBuilder().(*_AdsReadResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsReadResponseBuilder creates a AdsReadResponseBuilder
func (b *_AdsReadResponse) CreateAdsReadResponseBuilder() AdsReadResponseBuilder {
	if b == nil {
		return NewAdsReadResponseBuilder()
	}
	return &_AdsReadResponseBuilder{_AdsReadResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ
}

func (m *_AdsReadResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadResponse) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *_AdsReadResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsReadResponse(structType any) AdsReadResponse {
	if casted, ok := structType.(AdsReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadResponse) GetTypeName() string {
	return "AdsReadResponse"
}

func (m *_AdsReadResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsReadResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsReadResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsReadResponse AdsReadResponse, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	result, err := ReadEnumField[ReturnCode](ctx, "result", "ReturnCode", ReadEnum(ReturnCodeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'result' field"))
	}
	m.Result = result

	length, err := ReadImplicitField[uint32](ctx, "length", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	data, err := readBuffer.ReadByteArray("data", int(length))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("AdsReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadResponse")
	}

	return m, nil
}

func (m *_AdsReadResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsReadResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadResponse")
		}

		if err := WriteSimpleEnumField[ReturnCode](ctx, "result", "ReturnCode", m.GetResult(), WriteEnum[ReturnCode, uint32](ReturnCode.GetValue, ReturnCode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'result' field")
		}
		length := uint32(uint32(len(m.GetData())))
		if err := WriteImplicitField(ctx, "length", length, WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadResponse")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsReadResponse) IsAdsReadResponse() {}

func (m *_AdsReadResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsReadResponse) deepCopy() *_AdsReadResponse {
	if m == nil {
		return nil
	}
	_AdsReadResponseCopy := &_AdsReadResponse{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
		m.Result,
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsReadResponseCopy
}

func (m *_AdsReadResponse) String() string {
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
