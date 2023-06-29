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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetResult returns Result (property field)
	GetResult() uint8
	// GetReserved01 returns Reserved01 (property field)
	GetReserved01() uint8
	// GetAlarmType returns AlarmType (property field)
	GetAlarmType() AlarmType
	// GetReserved02 returns Reserved02 (property field)
	GetReserved02() uint8
	// GetReserved03 returns Reserved03 (property field)
	GetReserved03() uint8
}

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseExactly interface {
	S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
	isS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse() bool
}

// _S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse struct {
	*_S7PayloadUserDataItem
	Result     uint8
	Reserved01 uint8
	AlarmType  AlarmType
	Reserved02 uint8
	Reserved03 uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetCpuSubfunction() uint8 {
	return 0x02
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetResult() uint8 {
	return m.Result
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetReserved01() uint8 {
	return m.Reserved01
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetAlarmType() AlarmType {
	return m.AlarmType
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetReserved02() uint8 {
	return m.Reserved02
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetReserved03() uint8 {
	return m.Reserved03
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse factory function for _S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse(result uint8, reserved01 uint8, alarmType AlarmType, reserved02 uint8, reserved03 uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse{
		Result:                 result,
		Reserved01:             reserved01,
		AlarmType:              alarmType,
		Reserved02:             reserved02,
		Reserved03:             reserved03,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse(structType any) S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 8

	// Simple field (reserved01)
	lengthInBits += 8

	// Simple field (alarmType)
	lengthInBits += 8

	// Simple field (reserved02)
	lengthInBits += 8

	// Simple field (reserved03)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseParse(ctx context.Context, theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse, error) {
	return S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	_result, _resultErr := readBuffer.ReadUint8("result", 8)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
	}
	result := _result

	// Simple Field (reserved01)
	_reserved01, _reserved01Err := readBuffer.ReadUint8("reserved01", 8)
	if _reserved01Err != nil {
		return nil, errors.Wrap(_reserved01Err, "Error parsing 'reserved01' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
	}
	reserved01 := _reserved01

	// Simple Field (alarmType)
	if pullErr := readBuffer.PullContext("alarmType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmType")
	}
	_alarmType, _alarmTypeErr := AlarmTypeParseWithBuffer(ctx, readBuffer)
	if _alarmTypeErr != nil {
		return nil, errors.Wrap(_alarmTypeErr, "Error parsing 'alarmType' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
	}
	alarmType := _alarmType
	if closeErr := readBuffer.CloseContext("alarmType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmType")
	}

	// Simple Field (reserved02)
	_reserved02, _reserved02Err := readBuffer.ReadUint8("reserved02", 8)
	if _reserved02Err != nil {
		return nil, errors.Wrap(_reserved02Err, "Error parsing 'reserved02' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
	}
	reserved02 := _reserved02

	// Simple Field (reserved03)
	_reserved03, _reserved03Err := readBuffer.ReadUint8("reserved03", 8)
	if _reserved03Err != nil {
		return nil, errors.Wrap(_reserved03Err, "Error parsing 'reserved03' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
	}
	reserved03 := _reserved03

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		Result:                 result,
		Reserved01:             reserved01,
		AlarmType:              alarmType,
		Reserved02:             reserved02,
		Reserved03:             reserved03,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
		}

		// Simple Field (result)
		result := uint8(m.GetResult())
		_resultErr := writeBuffer.WriteUint8("result", 8, (result))
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (reserved01)
		reserved01 := uint8(m.GetReserved01())
		_reserved01Err := writeBuffer.WriteUint8("reserved01", 8, (reserved01))
		if _reserved01Err != nil {
			return errors.Wrap(_reserved01Err, "Error serializing 'reserved01' field")
		}

		// Simple Field (alarmType)
		if pushErr := writeBuffer.PushContext("alarmType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alarmType")
		}
		_alarmTypeErr := writeBuffer.WriteSerializable(ctx, m.GetAlarmType())
		if popErr := writeBuffer.PopContext("alarmType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alarmType")
		}
		if _alarmTypeErr != nil {
			return errors.Wrap(_alarmTypeErr, "Error serializing 'alarmType' field")
		}

		// Simple Field (reserved02)
		reserved02 := uint8(m.GetReserved02())
		_reserved02Err := writeBuffer.WriteUint8("reserved02", 8, (reserved02))
		if _reserved02Err != nil {
			return errors.Wrap(_reserved02Err, "Error serializing 'reserved02' field")
		}

		// Simple Field (reserved03)
		reserved03 := uint8(m.GetReserved03())
		_reserved03Err := writeBuffer.WriteUint8("reserved03", 8, (reserved03))
		if _reserved03Err != nil {
			return errors.Wrap(_reserved03Err, "Error serializing 'reserved03' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) isS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
