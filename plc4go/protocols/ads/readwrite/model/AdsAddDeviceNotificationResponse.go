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

// AdsAddDeviceNotificationResponse is the corresponding interface of AdsAddDeviceNotificationResponse
type AdsAddDeviceNotificationResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetNotificationHandle returns NotificationHandle (property field)
	GetNotificationHandle() uint32
	// IsAdsAddDeviceNotificationResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsAddDeviceNotificationResponse()
}

// _AdsAddDeviceNotificationResponse is the data-structure of this message
type _AdsAddDeviceNotificationResponse struct {
	AmsPacketContract
	Result             ReturnCode
	NotificationHandle uint32
}

var _ AdsAddDeviceNotificationResponse = (*_AdsAddDeviceNotificationResponse)(nil)
var _ AmsPacketRequirements = (*_AdsAddDeviceNotificationResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsAddDeviceNotificationResponse) GetCommandId() CommandId {
	return CommandId_ADS_ADD_DEVICE_NOTIFICATION
}

func (m *_AdsAddDeviceNotificationResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsAddDeviceNotificationResponse) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsAddDeviceNotificationResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *_AdsAddDeviceNotificationResponse) GetNotificationHandle() uint32 {
	return m.NotificationHandle
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsAddDeviceNotificationResponse factory function for _AdsAddDeviceNotificationResponse
func NewAdsAddDeviceNotificationResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32, result ReturnCode, notificationHandle uint32) *_AdsAddDeviceNotificationResponse {
	_result := &_AdsAddDeviceNotificationResponse{
		AmsPacketContract:  NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		Result:             result,
		NotificationHandle: notificationHandle,
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsAddDeviceNotificationResponse(structType any) AdsAddDeviceNotificationResponse {
	if casted, ok := structType.(AdsAddDeviceNotificationResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsAddDeviceNotificationResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsAddDeviceNotificationResponse) GetTypeName() string {
	return "AdsAddDeviceNotificationResponse"
}

func (m *_AdsAddDeviceNotificationResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	// Simple field (notificationHandle)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsAddDeviceNotificationResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsAddDeviceNotificationResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsAddDeviceNotificationResponse AdsAddDeviceNotificationResponse, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsAddDeviceNotificationResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsAddDeviceNotificationResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	result, err := ReadEnumField[ReturnCode](ctx, "result", "ReturnCode", ReadEnum(ReturnCodeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'result' field"))
	}
	m.Result = result

	notificationHandle, err := ReadSimpleField(ctx, "notificationHandle", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationHandle' field"))
	}
	m.NotificationHandle = notificationHandle

	if closeErr := readBuffer.CloseContext("AdsAddDeviceNotificationResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsAddDeviceNotificationResponse")
	}

	return m, nil
}

func (m *_AdsAddDeviceNotificationResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsAddDeviceNotificationResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsAddDeviceNotificationResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsAddDeviceNotificationResponse")
		}

		if err := WriteSimpleEnumField[ReturnCode](ctx, "result", "ReturnCode", m.GetResult(), WriteEnum[ReturnCode, uint32](ReturnCode.GetValue, ReturnCode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'result' field")
		}

		if err := WriteSimpleField[uint32](ctx, "notificationHandle", m.GetNotificationHandle(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationHandle' field")
		}

		if popErr := writeBuffer.PopContext("AdsAddDeviceNotificationResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsAddDeviceNotificationResponse")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsAddDeviceNotificationResponse) IsAdsAddDeviceNotificationResponse() {}

func (m *_AdsAddDeviceNotificationResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
