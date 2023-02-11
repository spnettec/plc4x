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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// AdsAddDeviceNotificationResponse is the corresponding interface of AdsAddDeviceNotificationResponse
type AdsAddDeviceNotificationResponse interface {
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetNotificationHandle returns NotificationHandle (property field)
	GetNotificationHandle() uint32
}

// AdsAddDeviceNotificationResponseExactly can be used when we want exactly this type and not a type which fulfills AdsAddDeviceNotificationResponse.
// This is useful for switch cases.
type AdsAddDeviceNotificationResponseExactly interface {
	AdsAddDeviceNotificationResponse
	isAdsAddDeviceNotificationResponse() bool
}

// _AdsAddDeviceNotificationResponse is the data-structure of this message
type _AdsAddDeviceNotificationResponse struct {
	*_AmsPacket
        Result ReturnCode
        NotificationHandle uint32
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsAddDeviceNotificationResponse)  GetCommandId() CommandId {
return CommandId_ADS_ADD_DEVICE_NOTIFICATION}

func (m *_AdsAddDeviceNotificationResponse)  GetResponse() bool {
return bool(true)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsAddDeviceNotificationResponse) InitializeParent(parent AmsPacket , targetAmsNetId AmsNetId , targetAmsPort uint16 , sourceAmsNetId AmsNetId , sourceAmsPort uint16 , errorCode uint32 , invokeId uint32 ) {	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsAddDeviceNotificationResponse)  GetParent() AmsPacket {
	return m._AmsPacket
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
func NewAdsAddDeviceNotificationResponse( result ReturnCode , notificationHandle uint32 , targetAmsNetId AmsNetId , targetAmsPort uint16 , sourceAmsNetId AmsNetId , sourceAmsPort uint16 , errorCode uint32 , invokeId uint32 ) *_AdsAddDeviceNotificationResponse {
	_result := &_AdsAddDeviceNotificationResponse{
		Result: result,
		NotificationHandle: notificationHandle,
    	_AmsPacket: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsAddDeviceNotificationResponse(structType interface{}) AdsAddDeviceNotificationResponse {
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
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	// Simple field (notificationHandle)
	lengthInBits += 32;

	return lengthInBits
}


func (m *_AdsAddDeviceNotificationResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsAddDeviceNotificationResponseParse(theBytes []byte) (AdsAddDeviceNotificationResponse, error) {
	return AdsAddDeviceNotificationResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AdsAddDeviceNotificationResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsAddDeviceNotificationResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsAddDeviceNotificationResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsAddDeviceNotificationResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for result")
	}
_result, _resultErr := ReturnCodeParseWithBuffer(ctx, readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field of AdsAddDeviceNotificationResponse")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for result")
	}

	// Simple Field (notificationHandle)
_notificationHandle, _notificationHandleErr := readBuffer.ReadUint32("notificationHandle", 32)
	if _notificationHandleErr != nil {
		return nil, errors.Wrap(_notificationHandleErr, "Error parsing 'notificationHandle' field of AdsAddDeviceNotificationResponse")
	}
	notificationHandle := _notificationHandle

	if closeErr := readBuffer.CloseContext("AdsAddDeviceNotificationResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsAddDeviceNotificationResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsAddDeviceNotificationResponse{
		_AmsPacket: &_AmsPacket{
		},
		Result: result,
		NotificationHandle: notificationHandle,
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
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
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsAddDeviceNotificationResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsAddDeviceNotificationResponse")
		}

	// Simple Field (result)
	if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for result")
	}
	_resultErr := writeBuffer.WriteSerializable(ctx, m.GetResult())
	if popErr := writeBuffer.PopContext("result"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for result")
	}
	if _resultErr != nil {
		return errors.Wrap(_resultErr, "Error serializing 'result' field")
	}

	// Simple Field (notificationHandle)
	notificationHandle := uint32(m.GetNotificationHandle())
	_notificationHandleErr := writeBuffer.WriteUint32("notificationHandle", 32, (notificationHandle))
	if _notificationHandleErr != nil {
		return errors.Wrap(_notificationHandleErr, "Error serializing 'notificationHandle' field")
	}

		if popErr := writeBuffer.PopContext("AdsAddDeviceNotificationResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsAddDeviceNotificationResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_AdsAddDeviceNotificationResponse) isAdsAddDeviceNotificationResponse() bool {
	return true
}

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



