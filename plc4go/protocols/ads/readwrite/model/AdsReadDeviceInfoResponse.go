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

// AdsReadDeviceInfoResponse is the corresponding interface of AdsReadDeviceInfoResponse
type AdsReadDeviceInfoResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetMajorVersion returns MajorVersion (property field)
	GetMajorVersion() uint8
	// GetMinorVersion returns MinorVersion (property field)
	GetMinorVersion() uint8
	// GetVersion returns Version (property field)
	GetVersion() uint16
	// GetDevice returns Device (property field)
	GetDevice() []byte
	// IsAdsReadDeviceInfoResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsReadDeviceInfoResponse()
}

// _AdsReadDeviceInfoResponse is the data-structure of this message
type _AdsReadDeviceInfoResponse struct {
	AmsPacketContract
	Result       ReturnCode
	MajorVersion uint8
	MinorVersion uint8
	Version      uint16
	Device       []byte
}

var _ AdsReadDeviceInfoResponse = (*_AdsReadDeviceInfoResponse)(nil)
var _ AmsPacketRequirements = (*_AdsReadDeviceInfoResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadDeviceInfoResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ_DEVICE_INFO
}

func (m *_AdsReadDeviceInfoResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadDeviceInfoResponse) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadDeviceInfoResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *_AdsReadDeviceInfoResponse) GetMajorVersion() uint8 {
	return m.MajorVersion
}

func (m *_AdsReadDeviceInfoResponse) GetMinorVersion() uint8 {
	return m.MinorVersion
}

func (m *_AdsReadDeviceInfoResponse) GetVersion() uint16 {
	return m.Version
}

func (m *_AdsReadDeviceInfoResponse) GetDevice() []byte {
	return m.Device
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadDeviceInfoResponse factory function for _AdsReadDeviceInfoResponse
func NewAdsReadDeviceInfoResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32, result ReturnCode, majorVersion uint8, minorVersion uint8, version uint16, device []byte) *_AdsReadDeviceInfoResponse {
	_result := &_AdsReadDeviceInfoResponse{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		Result:            result,
		MajorVersion:      majorVersion,
		MinorVersion:      minorVersion,
		Version:           version,
		Device:            device,
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadDeviceInfoResponse(structType any) AdsReadDeviceInfoResponse {
	if casted, ok := structType.(AdsReadDeviceInfoResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadDeviceInfoResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadDeviceInfoResponse) GetTypeName() string {
	return "AdsReadDeviceInfoResponse"
}

func (m *_AdsReadDeviceInfoResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	// Simple field (version)
	lengthInBits += 16

	// Array field
	if len(m.Device) > 0 {
		lengthInBits += 8 * uint16(len(m.Device))
	}

	return lengthInBits
}

func (m *_AdsReadDeviceInfoResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsReadDeviceInfoResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsReadDeviceInfoResponse AdsReadDeviceInfoResponse, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadDeviceInfoResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadDeviceInfoResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	result, err := ReadEnumField[ReturnCode](ctx, "result", "ReturnCode", ReadEnum(ReturnCodeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'result' field"))
	}
	m.Result = result

	majorVersion, err := ReadSimpleField(ctx, "majorVersion", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'majorVersion' field"))
	}
	m.MajorVersion = majorVersion

	minorVersion, err := ReadSimpleField(ctx, "minorVersion", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minorVersion' field"))
	}
	m.MinorVersion = minorVersion

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	device, err := readBuffer.ReadByteArray("device", int(int32(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'device' field"))
	}
	m.Device = device

	if closeErr := readBuffer.CloseContext("AdsReadDeviceInfoResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadDeviceInfoResponse")
	}

	return m, nil
}

func (m *_AdsReadDeviceInfoResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsReadDeviceInfoResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadDeviceInfoResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadDeviceInfoResponse")
		}

		if err := WriteSimpleEnumField[ReturnCode](ctx, "result", "ReturnCode", m.GetResult(), WriteEnum[ReturnCode, uint32](ReturnCode.GetValue, ReturnCode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'result' field")
		}

		if err := WriteSimpleField[uint8](ctx, "majorVersion", m.GetMajorVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'majorVersion' field")
		}

		if err := WriteSimpleField[uint8](ctx, "minorVersion", m.GetMinorVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'minorVersion' field")
		}

		if err := WriteSimpleField[uint16](ctx, "version", m.GetVersion(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if err := WriteByteArrayField(ctx, "device", m.GetDevice(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'device' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadDeviceInfoResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadDeviceInfoResponse")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsReadDeviceInfoResponse) IsAdsReadDeviceInfoResponse() {}

func (m *_AdsReadDeviceInfoResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
