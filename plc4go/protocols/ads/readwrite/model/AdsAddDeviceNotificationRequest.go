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

// AdsAddDeviceNotificationRequest is the corresponding interface of AdsAddDeviceNotificationRequest
type AdsAddDeviceNotificationRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetIndexGroup returns IndexGroup (property field)
	GetIndexGroup() uint32
	// GetIndexOffset returns IndexOffset (property field)
	GetIndexOffset() uint32
	// GetLength returns Length (property field)
	GetLength() uint32
	// GetTransmissionMode returns TransmissionMode (property field)
	GetTransmissionMode() AdsTransMode
	// GetMaxDelayInMs returns MaxDelayInMs (property field)
	GetMaxDelayInMs() uint32
	// GetCycleTimeInMs returns CycleTimeInMs (property field)
	GetCycleTimeInMs() uint32
}

// AdsAddDeviceNotificationRequestExactly can be used when we want exactly this type and not a type which fulfills AdsAddDeviceNotificationRequest.
// This is useful for switch cases.
type AdsAddDeviceNotificationRequestExactly interface {
	AdsAddDeviceNotificationRequest
	isAdsAddDeviceNotificationRequest() bool
}

// _AdsAddDeviceNotificationRequest is the data-structure of this message
type _AdsAddDeviceNotificationRequest struct {
	*_AmsPacket
	IndexGroup       uint32
	IndexOffset      uint32
	Length           uint32
	TransmissionMode AdsTransMode
	MaxDelayInMs     uint32
	CycleTimeInMs    uint32
	// Reserved Fields
	reservedField0 *uint64
	reservedField1 *uint64
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsAddDeviceNotificationRequest) GetCommandId() CommandId {
	return CommandId_ADS_ADD_DEVICE_NOTIFICATION
}

func (m *_AdsAddDeviceNotificationRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsAddDeviceNotificationRequest) InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) {
	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsAddDeviceNotificationRequest) GetParent() AmsPacket {
	return m._AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsAddDeviceNotificationRequest) GetIndexGroup() uint32 {
	return m.IndexGroup
}

func (m *_AdsAddDeviceNotificationRequest) GetIndexOffset() uint32 {
	return m.IndexOffset
}

func (m *_AdsAddDeviceNotificationRequest) GetLength() uint32 {
	return m.Length
}

func (m *_AdsAddDeviceNotificationRequest) GetTransmissionMode() AdsTransMode {
	return m.TransmissionMode
}

func (m *_AdsAddDeviceNotificationRequest) GetMaxDelayInMs() uint32 {
	return m.MaxDelayInMs
}

func (m *_AdsAddDeviceNotificationRequest) GetCycleTimeInMs() uint32 {
	return m.CycleTimeInMs
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsAddDeviceNotificationRequest factory function for _AdsAddDeviceNotificationRequest
func NewAdsAddDeviceNotificationRequest(indexGroup uint32, indexOffset uint32, length uint32, transmissionMode AdsTransMode, maxDelayInMs uint32, cycleTimeInMs uint32, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsAddDeviceNotificationRequest {
	_result := &_AdsAddDeviceNotificationRequest{
		IndexGroup:       indexGroup,
		IndexOffset:      indexOffset,
		Length:           length,
		TransmissionMode: transmissionMode,
		MaxDelayInMs:     maxDelayInMs,
		CycleTimeInMs:    cycleTimeInMs,
		_AmsPacket:       NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsAddDeviceNotificationRequest(structType any) AdsAddDeviceNotificationRequest {
	if casted, ok := structType.(AdsAddDeviceNotificationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsAddDeviceNotificationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsAddDeviceNotificationRequest) GetTypeName() string {
	return "AdsAddDeviceNotificationRequest"
}

func (m *_AdsAddDeviceNotificationRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (length)
	lengthInBits += 32

	// Simple field (transmissionMode)
	lengthInBits += 32

	// Simple field (maxDelayInMs)
	lengthInBits += 32

	// Simple field (cycleTimeInMs)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 64

	// Reserved Field (reserved)
	lengthInBits += 64

	return lengthInBits
}

func (m *_AdsAddDeviceNotificationRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsAddDeviceNotificationRequestParse(theBytes []byte) (AdsAddDeviceNotificationRequest, error) {
	return AdsAddDeviceNotificationRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AdsAddDeviceNotificationRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsAddDeviceNotificationRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsAddDeviceNotificationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsAddDeviceNotificationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (indexGroup)
	_indexGroup, _indexGroupErr := readBuffer.ReadUint32("indexGroup", 32)
	if _indexGroupErr != nil {
		return nil, errors.Wrap(_indexGroupErr, "Error parsing 'indexGroup' field of AdsAddDeviceNotificationRequest")
	}
	indexGroup := _indexGroup

	// Simple Field (indexOffset)
	_indexOffset, _indexOffsetErr := readBuffer.ReadUint32("indexOffset", 32)
	if _indexOffsetErr != nil {
		return nil, errors.Wrap(_indexOffsetErr, "Error parsing 'indexOffset' field of AdsAddDeviceNotificationRequest")
	}
	indexOffset := _indexOffset

	// Simple Field (length)
	_length, _lengthErr := readBuffer.ReadUint32("length", 32)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AdsAddDeviceNotificationRequest")
	}
	length := _length

	// Simple Field (transmissionMode)
	if pullErr := readBuffer.PullContext("transmissionMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transmissionMode")
	}
	_transmissionMode, _transmissionModeErr := AdsTransModeParseWithBuffer(ctx, readBuffer)
	if _transmissionModeErr != nil {
		return nil, errors.Wrap(_transmissionModeErr, "Error parsing 'transmissionMode' field of AdsAddDeviceNotificationRequest")
	}
	transmissionMode := _transmissionMode
	if closeErr := readBuffer.CloseContext("transmissionMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transmissionMode")
	}

	// Simple Field (maxDelayInMs)
	_maxDelayInMs, _maxDelayInMsErr := readBuffer.ReadUint32("maxDelayInMs", 32)
	if _maxDelayInMsErr != nil {
		return nil, errors.Wrap(_maxDelayInMsErr, "Error parsing 'maxDelayInMs' field of AdsAddDeviceNotificationRequest")
	}
	maxDelayInMs := _maxDelayInMs

	// Simple Field (cycleTimeInMs)
	_cycleTimeInMs, _cycleTimeInMsErr := readBuffer.ReadUint32("cycleTimeInMs", 32)
	if _cycleTimeInMsErr != nil {
		return nil, errors.Wrap(_cycleTimeInMsErr, "Error parsing 'cycleTimeInMs' field of AdsAddDeviceNotificationRequest")
	}
	cycleTimeInMs := _cycleTimeInMs

	var reservedField0 *uint64
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint64("reserved", 64)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of AdsAddDeviceNotificationRequest")
		}
		if reserved != uint64(0x0000) {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": uint64(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	var reservedField1 *uint64
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint64("reserved", 64)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of AdsAddDeviceNotificationRequest")
		}
		if reserved != uint64(0x0000) {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": uint64(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("AdsAddDeviceNotificationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsAddDeviceNotificationRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsAddDeviceNotificationRequest{
		_AmsPacket:       &_AmsPacket{},
		IndexGroup:       indexGroup,
		IndexOffset:      indexOffset,
		Length:           length,
		TransmissionMode: transmissionMode,
		MaxDelayInMs:     maxDelayInMs,
		CycleTimeInMs:    cycleTimeInMs,
		reservedField0:   reservedField0,
		reservedField1:   reservedField1,
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsAddDeviceNotificationRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsAddDeviceNotificationRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsAddDeviceNotificationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsAddDeviceNotificationRequest")
		}

		// Simple Field (indexGroup)
		indexGroup := uint32(m.GetIndexGroup())
		_indexGroupErr := writeBuffer.WriteUint32("indexGroup", 32, (indexGroup))
		if _indexGroupErr != nil {
			return errors.Wrap(_indexGroupErr, "Error serializing 'indexGroup' field")
		}

		// Simple Field (indexOffset)
		indexOffset := uint32(m.GetIndexOffset())
		_indexOffsetErr := writeBuffer.WriteUint32("indexOffset", 32, (indexOffset))
		if _indexOffsetErr != nil {
			return errors.Wrap(_indexOffsetErr, "Error serializing 'indexOffset' field")
		}

		// Simple Field (length)
		length := uint32(m.GetLength())
		_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Simple Field (transmissionMode)
		if pushErr := writeBuffer.PushContext("transmissionMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transmissionMode")
		}
		_transmissionModeErr := writeBuffer.WriteSerializable(ctx, m.GetTransmissionMode())
		if popErr := writeBuffer.PopContext("transmissionMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transmissionMode")
		}
		if _transmissionModeErr != nil {
			return errors.Wrap(_transmissionModeErr, "Error serializing 'transmissionMode' field")
		}

		// Simple Field (maxDelayInMs)
		maxDelayInMs := uint32(m.GetMaxDelayInMs())
		_maxDelayInMsErr := writeBuffer.WriteUint32("maxDelayInMs", 32, (maxDelayInMs))
		if _maxDelayInMsErr != nil {
			return errors.Wrap(_maxDelayInMsErr, "Error serializing 'maxDelayInMs' field")
		}

		// Simple Field (cycleTimeInMs)
		cycleTimeInMs := uint32(m.GetCycleTimeInMs())
		_cycleTimeInMsErr := writeBuffer.WriteUint32("cycleTimeInMs", 32, (cycleTimeInMs))
		if _cycleTimeInMsErr != nil {
			return errors.Wrap(_cycleTimeInMsErr, "Error serializing 'cycleTimeInMs' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint64 = uint64(0x0000)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]any{
					"expected value": uint64(0x0000),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint64("reserved", 64, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			var reserved uint64 = uint64(0x0000)
			if m.reservedField1 != nil {
				Plc4xModelLog.Info().Fields(map[string]any{
					"expected value": uint64(0x0000),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := writeBuffer.WriteUint64("reserved", 64, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("AdsAddDeviceNotificationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsAddDeviceNotificationRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsAddDeviceNotificationRequest) isAdsAddDeviceNotificationRequest() bool {
	return true
}

func (m *_AdsAddDeviceNotificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
