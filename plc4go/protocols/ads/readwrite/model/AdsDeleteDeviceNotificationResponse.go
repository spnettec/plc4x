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

// AdsDeleteDeviceNotificationResponse is the corresponding interface of AdsDeleteDeviceNotificationResponse
type AdsDeleteDeviceNotificationResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// IsAdsDeleteDeviceNotificationResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDeleteDeviceNotificationResponse()
	// CreateBuilder creates a AdsDeleteDeviceNotificationResponseBuilder
	CreateAdsDeleteDeviceNotificationResponseBuilder() AdsDeleteDeviceNotificationResponseBuilder
}

// _AdsDeleteDeviceNotificationResponse is the data-structure of this message
type _AdsDeleteDeviceNotificationResponse struct {
	AmsPacketContract
	Result ReturnCode
}

var _ AdsDeleteDeviceNotificationResponse = (*_AdsDeleteDeviceNotificationResponse)(nil)
var _ AmsPacketRequirements = (*_AdsDeleteDeviceNotificationResponse)(nil)

// NewAdsDeleteDeviceNotificationResponse factory function for _AdsDeleteDeviceNotificationResponse
func NewAdsDeleteDeviceNotificationResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32, result ReturnCode) *_AdsDeleteDeviceNotificationResponse {
	_result := &_AdsDeleteDeviceNotificationResponse{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		Result:            result,
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsDeleteDeviceNotificationResponseBuilder is a builder for AdsDeleteDeviceNotificationResponse
type AdsDeleteDeviceNotificationResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(result ReturnCode) AdsDeleteDeviceNotificationResponseBuilder
	// WithResult adds Result (property field)
	WithResult(ReturnCode) AdsDeleteDeviceNotificationResponseBuilder
	// Build builds the AdsDeleteDeviceNotificationResponse or returns an error if something is wrong
	Build() (AdsDeleteDeviceNotificationResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsDeleteDeviceNotificationResponse
}

// NewAdsDeleteDeviceNotificationResponseBuilder() creates a AdsDeleteDeviceNotificationResponseBuilder
func NewAdsDeleteDeviceNotificationResponseBuilder() AdsDeleteDeviceNotificationResponseBuilder {
	return &_AdsDeleteDeviceNotificationResponseBuilder{_AdsDeleteDeviceNotificationResponse: new(_AdsDeleteDeviceNotificationResponse)}
}

type _AdsDeleteDeviceNotificationResponseBuilder struct {
	*_AdsDeleteDeviceNotificationResponse

	parentBuilder *_AmsPacketBuilder

	err *utils.MultiError
}

var _ (AdsDeleteDeviceNotificationResponseBuilder) = (*_AdsDeleteDeviceNotificationResponseBuilder)(nil)

func (b *_AdsDeleteDeviceNotificationResponseBuilder) setParent(contract AmsPacketContract) {
	b.AmsPacketContract = contract
}

func (b *_AdsDeleteDeviceNotificationResponseBuilder) WithMandatoryFields(result ReturnCode) AdsDeleteDeviceNotificationResponseBuilder {
	return b.WithResult(result)
}

func (b *_AdsDeleteDeviceNotificationResponseBuilder) WithResult(result ReturnCode) AdsDeleteDeviceNotificationResponseBuilder {
	b.Result = result
	return b
}

func (b *_AdsDeleteDeviceNotificationResponseBuilder) Build() (AdsDeleteDeviceNotificationResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsDeleteDeviceNotificationResponse.deepCopy(), nil
}

func (b *_AdsDeleteDeviceNotificationResponseBuilder) MustBuild() AdsDeleteDeviceNotificationResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AdsDeleteDeviceNotificationResponseBuilder) Done() AmsPacketBuilder {
	return b.parentBuilder
}

func (b *_AdsDeleteDeviceNotificationResponseBuilder) buildForAmsPacket() (AmsPacket, error) {
	return b.Build()
}

func (b *_AdsDeleteDeviceNotificationResponseBuilder) DeepCopy() any {
	_copy := b.CreateAdsDeleteDeviceNotificationResponseBuilder().(*_AdsDeleteDeviceNotificationResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsDeleteDeviceNotificationResponseBuilder creates a AdsDeleteDeviceNotificationResponseBuilder
func (b *_AdsDeleteDeviceNotificationResponse) CreateAdsDeleteDeviceNotificationResponseBuilder() AdsDeleteDeviceNotificationResponseBuilder {
	if b == nil {
		return NewAdsDeleteDeviceNotificationResponseBuilder()
	}
	return &_AdsDeleteDeviceNotificationResponseBuilder{_AdsDeleteDeviceNotificationResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) GetCommandId() CommandId {
	return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
}

func (m *_AdsDeleteDeviceNotificationResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) GetResult() ReturnCode {
	return m.Result
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDeleteDeviceNotificationResponse(structType any) AdsDeleteDeviceNotificationResponse {
	if casted, ok := structType.(AdsDeleteDeviceNotificationResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDeleteDeviceNotificationResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDeleteDeviceNotificationResponse) GetTypeName() string {
	return "AdsDeleteDeviceNotificationResponse"
}

func (m *_AdsDeleteDeviceNotificationResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsDeleteDeviceNotificationResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDeleteDeviceNotificationResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsDeleteDeviceNotificationResponse AdsDeleteDeviceNotificationResponse, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDeleteDeviceNotificationResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDeleteDeviceNotificationResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	result, err := ReadEnumField[ReturnCode](ctx, "result", "ReturnCode", ReadEnum(ReturnCodeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'result' field"))
	}
	m.Result = result

	if closeErr := readBuffer.CloseContext("AdsDeleteDeviceNotificationResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDeleteDeviceNotificationResponse")
	}

	return m, nil
}

func (m *_AdsDeleteDeviceNotificationResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDeleteDeviceNotificationResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeleteDeviceNotificationResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDeleteDeviceNotificationResponse")
		}

		if err := WriteSimpleEnumField[ReturnCode](ctx, "result", "ReturnCode", m.GetResult(), WriteEnum[ReturnCode, uint32](ReturnCode.GetValue, ReturnCode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'result' field")
		}

		if popErr := writeBuffer.PopContext("AdsDeleteDeviceNotificationResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDeleteDeviceNotificationResponse")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDeleteDeviceNotificationResponse) IsAdsDeleteDeviceNotificationResponse() {}

func (m *_AdsDeleteDeviceNotificationResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDeleteDeviceNotificationResponse) deepCopy() *_AdsDeleteDeviceNotificationResponse {
	if m == nil {
		return nil
	}
	_AdsDeleteDeviceNotificationResponseCopy := &_AdsDeleteDeviceNotificationResponse{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
		m.Result,
	}
	m.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsDeleteDeviceNotificationResponseCopy
}

func (m *_AdsDeleteDeviceNotificationResponse) String() string {
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
