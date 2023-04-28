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
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserDataItem is the corresponding interface of S7PayloadUserDataItem
type S7PayloadUserDataItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCpuFunctionGroup returns CpuFunctionGroup (discriminator field)
	GetCpuFunctionGroup() uint8
	// GetCpuFunctionType returns CpuFunctionType (discriminator field)
	GetCpuFunctionType() uint8
	// GetCpuSubfunction returns CpuSubfunction (discriminator field)
	GetCpuSubfunction() uint8
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
	// GetDataLength returns DataLength (property field)
	GetDataLength() uint16
}

// S7PayloadUserDataItemExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItem.
// This is useful for switch cases.
type S7PayloadUserDataItemExactly interface {
	S7PayloadUserDataItem
	isS7PayloadUserDataItem() bool
}

// _S7PayloadUserDataItem is the data-structure of this message
type _S7PayloadUserDataItem struct {
	_S7PayloadUserDataItemChildRequirements
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	DataLength    uint16
}

type _S7PayloadUserDataItemChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCpuFunctionGroup() uint8
	GetCpuFunctionType() uint8
	GetCpuSubfunction() uint8
	GetDataLength() uint16
}

type S7PayloadUserDataItemParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7PayloadUserDataItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type S7PayloadUserDataItemChild interface {
	utils.Serializable
	InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16)
	GetParent() *S7PayloadUserDataItem

	GetTypeName() string
	S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItem) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_S7PayloadUserDataItem) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

func (m *_S7PayloadUserDataItem) GetDataLength() uint16 {
	return m.DataLength
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItem factory function for _S7PayloadUserDataItem
func NewS7PayloadUserDataItem(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItem {
	return &_S7PayloadUserDataItem{ReturnCode: returnCode, TransportSize: transportSize, DataLength: dataLength}
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItem(structType any) S7PayloadUserDataItem {
	if casted, ok := structType.(S7PayloadUserDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItem) GetTypeName() string {
	return "S7PayloadUserDataItem"
}

func (m *_S7PayloadUserDataItem) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Simple field (dataLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *_S7PayloadUserDataItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemParse(theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItem, error) {
	return S7PayloadUserDataItemParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for returnCode")
	}
	_returnCode, _returnCodeErr := DataTransportErrorCodeParseWithBuffer(ctx, readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field of S7PayloadUserDataItem")
	}
	returnCode := _returnCode
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for returnCode")
	}

	// Simple Field (transportSize)
	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportSize")
	}
	_transportSize, _transportSizeErr := DataTransportSizeParseWithBuffer(ctx, readBuffer)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field of S7PayloadUserDataItem")
	}
	transportSize := _transportSize
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportSize")
	}

	// Simple Field (dataLength)
	_dataLength, _dataLengthErr := readBuffer.ReadUint16("dataLength", 16)
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field of S7PayloadUserDataItem")
	}
	dataLength := _dataLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7PayloadUserDataItemChildSerializeRequirement interface {
		S7PayloadUserDataItem
		InitializeParent(S7PayloadUserDataItem, DataTransportErrorCode, DataTransportSize, uint16)
		GetParent() S7PayloadUserDataItem
	}
	var _childTemp any
	var _child S7PayloadUserDataItemChildSerializeRequirement
	var typeSwitchError error
	switch {
	case cpuFunctionGroup == 0x02 && cpuFunctionType == 0x00 && cpuSubfunction == 0x01: // S7PayloadUserDataItemCyclicServicesPush
		_childTemp, typeSwitchError = S7PayloadUserDataItemCyclicServicesPushParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x02 && cpuFunctionType == 0x00 && cpuSubfunction == 0x05: // S7PayloadUserDataItemCyclicServicesChangeDrivenPush
		_childTemp, typeSwitchError = S7PayloadUserDataItemCyclicServicesChangeDrivenPushParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x02 && cpuFunctionType == 0x04 && cpuSubfunction == 0x01: // S7PayloadUserDataItemCyclicServicesSubscribeRequest
		_childTemp, typeSwitchError = S7PayloadUserDataItemCyclicServicesSubscribeRequestParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x02 && cpuFunctionType == 0x04 && cpuSubfunction == 0x04: // S7PayloadUserDataItemCyclicServicesUnsubscribeRequest
		_childTemp, typeSwitchError = S7PayloadUserDataItemCyclicServicesUnsubscribeRequestParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x02 && cpuFunctionType == 0x08 && cpuSubfunction == 0x01: // S7PayloadUserDataItemCyclicServicesSubscribeResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCyclicServicesSubscribeResponseParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x02 && cpuFunctionType == 0x08 && cpuSubfunction == 0x04: // S7PayloadUserDataItemCyclicServicesUnsubscribeResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCyclicServicesUnsubscribeResponseParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x02 && cpuFunctionType == 0x08 && cpuSubfunction == 0x05 && dataLength == 0x00: // S7PayloadUserDataItemCyclicServicesErrorResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCyclicServicesErrorResponseParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x02 && cpuFunctionType == 0x08 && cpuSubfunction == 0x05: // S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponseParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x00 && cpuSubfunction == 0x03: // S7PayloadDiagnosticMessage
		_childTemp, typeSwitchError = S7PayloadDiagnosticMessageParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x00 && cpuSubfunction == 0x05: // S7PayloadAlarm8
		_childTemp, typeSwitchError = S7PayloadAlarm8ParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x00 && cpuSubfunction == 0x06: // S7PayloadNotify
		_childTemp, typeSwitchError = S7PayloadNotifyParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x00 && cpuSubfunction == 0x0c: // S7PayloadAlarmAckInd
		_childTemp, typeSwitchError = S7PayloadAlarmAckIndParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x00 && cpuSubfunction == 0x11: // S7PayloadAlarmSQ
		_childTemp, typeSwitchError = S7PayloadAlarmSQParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x00 && cpuSubfunction == 0x12: // S7PayloadAlarmS
		_childTemp, typeSwitchError = S7PayloadAlarmSParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x00 && cpuSubfunction == 0x13: // S7PayloadAlarmSC
		_childTemp, typeSwitchError = S7PayloadAlarmSCParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x00 && cpuSubfunction == 0x16: // S7PayloadNotify8
		_childTemp, typeSwitchError = S7PayloadNotify8ParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x04 && cpuSubfunction == 0x01 && dataLength == 0x00: // S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequestParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x04 && cpuSubfunction == 0x01: // S7PayloadUserDataItemCpuFunctionReadSzlRequest
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionReadSzlRequestParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x08 && cpuSubfunction == 0x01: // S7PayloadUserDataItemCpuFunctionReadSzlResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionReadSzlResponseParseWithBuffer(ctx, readBuffer, dataLength, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x04 && cpuSubfunction == 0x02: // S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequestParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x08 && cpuSubfunction == 0x02 && dataLength == 0x00: // S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x08 && cpuSubfunction == 0x02 && dataLength == 0x02: // S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x08 && cpuSubfunction == 0x02 && dataLength == 0x05: // S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x04 && cpuSubfunction == 0x0b: // S7PayloadUserDataItemCpuFunctionAlarmAckRequest
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionAlarmAckRequestParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x08 && cpuSubfunction == 0x0b && dataLength == 0x00: // S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponseParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x08 && cpuSubfunction == 0x0b: // S7PayloadUserDataItemCpuFunctionAlarmAckResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionAlarmAckResponseParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x04 && cpuSubfunction == 0x13: // S7PayloadUserDataItemCpuFunctionAlarmQueryRequest
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionAlarmQueryRequestParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	case cpuFunctionGroup == 0x04 && cpuFunctionType == 0x08 && cpuSubfunction == 0x13: // S7PayloadUserDataItemCpuFunctionAlarmQueryResponse
		_childTemp, typeSwitchError = S7PayloadUserDataItemCpuFunctionAlarmQueryResponseParseWithBuffer(ctx, readBuffer, dataLength, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [cpuFunctionGroup=%v, cpuFunctionType=%v, cpuSubfunction=%v, dataLength=%v]", cpuFunctionGroup, cpuFunctionType, cpuSubfunction, dataLength)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of S7PayloadUserDataItem")
	}
	_child = _childTemp.(S7PayloadUserDataItemChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItem")
	}

	// Finish initializing
	_child.InitializeParent(_child, returnCode, transportSize, dataLength)
	return _child, nil
}

func (pm *_S7PayloadUserDataItem) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7PayloadUserDataItem, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("S7PayloadUserDataItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItem")
	}

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for returnCode")
	}
	_returnCodeErr := writeBuffer.WriteSerializable(ctx, m.GetReturnCode())
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for returnCode")
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Simple Field (transportSize)
	if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for transportSize")
	}
	_transportSizeErr := writeBuffer.WriteSerializable(ctx, m.GetTransportSize())
	if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for transportSize")
	}
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Simple Field (dataLength)
	dataLength := uint16(m.GetDataLength())
	_dataLengthErr := writeBuffer.WriteUint16("dataLength", 16, (dataLength))
	if _dataLengthErr != nil {
		return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7PayloadUserDataItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItem")
	}
	return nil
}

func (m *_S7PayloadUserDataItem) isS7PayloadUserDataItem() bool {
	return true
}

func (m *_S7PayloadUserDataItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
