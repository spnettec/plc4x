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

// CipConnectionManagerResponse is the corresponding interface of CipConnectionManagerResponse
type CipConnectionManagerResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetOtConnectionId returns OtConnectionId (property field)
	GetOtConnectionId() uint32
	// GetToConnectionId returns ToConnectionId (property field)
	GetToConnectionId() uint32
	// GetConnectionSerialNumber returns ConnectionSerialNumber (property field)
	GetConnectionSerialNumber() uint16
	// GetOriginatorVendorId returns OriginatorVendorId (property field)
	GetOriginatorVendorId() uint16
	// GetOriginatorSerialNumber returns OriginatorSerialNumber (property field)
	GetOriginatorSerialNumber() uint32
	// GetOtApi returns OtApi (property field)
	GetOtApi() uint32
	// GetToApi returns ToApi (property field)
	GetToApi() uint32
}

// CipConnectionManagerResponseExactly can be used when we want exactly this type and not a type which fulfills CipConnectionManagerResponse.
// This is useful for switch cases.
type CipConnectionManagerResponseExactly interface {
	CipConnectionManagerResponse
	isCipConnectionManagerResponse() bool
}

// _CipConnectionManagerResponse is the data-structure of this message
type _CipConnectionManagerResponse struct {
	*_CipService
	OtConnectionId         uint32
	ToConnectionId         uint32
	ConnectionSerialNumber uint16
	OriginatorVendorId     uint16
	OriginatorSerialNumber uint32
	OtApi                  uint32
	ToApi                  uint32
	// Reserved Fields
	reservedField0 *uint32
	reservedField1 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectionManagerResponse) GetService() uint8 {
	return 0x5B
}

func (m *_CipConnectionManagerResponse) GetResponse() bool {
	return bool(true)
}

func (m *_CipConnectionManagerResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectionManagerResponse) InitializeParent(parent CipService) {}

func (m *_CipConnectionManagerResponse) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectionManagerResponse) GetOtConnectionId() uint32 {
	return m.OtConnectionId
}

func (m *_CipConnectionManagerResponse) GetToConnectionId() uint32 {
	return m.ToConnectionId
}

func (m *_CipConnectionManagerResponse) GetConnectionSerialNumber() uint16 {
	return m.ConnectionSerialNumber
}

func (m *_CipConnectionManagerResponse) GetOriginatorVendorId() uint16 {
	return m.OriginatorVendorId
}

func (m *_CipConnectionManagerResponse) GetOriginatorSerialNumber() uint32 {
	return m.OriginatorSerialNumber
}

func (m *_CipConnectionManagerResponse) GetOtApi() uint32 {
	return m.OtApi
}

func (m *_CipConnectionManagerResponse) GetToApi() uint32 {
	return m.ToApi
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipConnectionManagerResponse factory function for _CipConnectionManagerResponse
func NewCipConnectionManagerResponse(otConnectionId uint32, toConnectionId uint32, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, otApi uint32, toApi uint32, serviceLen uint16) *_CipConnectionManagerResponse {
	_result := &_CipConnectionManagerResponse{
		OtConnectionId:         otConnectionId,
		ToConnectionId:         toConnectionId,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		OtApi:                  otApi,
		ToApi:                  toApi,
		_CipService:            NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipConnectionManagerResponse(structType any) CipConnectionManagerResponse {
	if casted, ok := structType.(CipConnectionManagerResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectionManagerResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectionManagerResponse) GetTypeName() string {
	return "CipConnectionManagerResponse"
}

func (m *_CipConnectionManagerResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 24

	// Simple field (otConnectionId)
	lengthInBits += 32

	// Simple field (toConnectionId)
	lengthInBits += 32

	// Simple field (connectionSerialNumber)
	lengthInBits += 16

	// Simple field (originatorVendorId)
	lengthInBits += 16

	// Simple field (originatorSerialNumber)
	lengthInBits += 32

	// Simple field (otApi)
	lengthInBits += 32

	// Simple field (toApi)
	lengthInBits += 32

	// Implicit Field (replySize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipConnectionManagerResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipConnectionManagerResponseParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (CipConnectionManagerResponse, error) {
	return CipConnectionManagerResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func CipConnectionManagerResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (CipConnectionManagerResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CipConnectionManagerResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectionManagerResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint32
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint32("reserved", 24)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipConnectionManagerResponse")
		}
		if reserved != uint32(0x000000) {
			log.Info().Fields(map[string]any{
				"expected value": uint32(0x000000),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (otConnectionId)
	_otConnectionId, _otConnectionIdErr := readBuffer.ReadUint32("otConnectionId", 32)
	if _otConnectionIdErr != nil {
		return nil, errors.Wrap(_otConnectionIdErr, "Error parsing 'otConnectionId' field of CipConnectionManagerResponse")
	}
	otConnectionId := _otConnectionId

	// Simple Field (toConnectionId)
	_toConnectionId, _toConnectionIdErr := readBuffer.ReadUint32("toConnectionId", 32)
	if _toConnectionIdErr != nil {
		return nil, errors.Wrap(_toConnectionIdErr, "Error parsing 'toConnectionId' field of CipConnectionManagerResponse")
	}
	toConnectionId := _toConnectionId

	// Simple Field (connectionSerialNumber)
	_connectionSerialNumber, _connectionSerialNumberErr := readBuffer.ReadUint16("connectionSerialNumber", 16)
	if _connectionSerialNumberErr != nil {
		return nil, errors.Wrap(_connectionSerialNumberErr, "Error parsing 'connectionSerialNumber' field of CipConnectionManagerResponse")
	}
	connectionSerialNumber := _connectionSerialNumber

	// Simple Field (originatorVendorId)
	_originatorVendorId, _originatorVendorIdErr := readBuffer.ReadUint16("originatorVendorId", 16)
	if _originatorVendorIdErr != nil {
		return nil, errors.Wrap(_originatorVendorIdErr, "Error parsing 'originatorVendorId' field of CipConnectionManagerResponse")
	}
	originatorVendorId := _originatorVendorId

	// Simple Field (originatorSerialNumber)
	_originatorSerialNumber, _originatorSerialNumberErr := readBuffer.ReadUint32("originatorSerialNumber", 32)
	if _originatorSerialNumberErr != nil {
		return nil, errors.Wrap(_originatorSerialNumberErr, "Error parsing 'originatorSerialNumber' field of CipConnectionManagerResponse")
	}
	originatorSerialNumber := _originatorSerialNumber

	// Simple Field (otApi)
	_otApi, _otApiErr := readBuffer.ReadUint32("otApi", 32)
	if _otApiErr != nil {
		return nil, errors.Wrap(_otApiErr, "Error parsing 'otApi' field of CipConnectionManagerResponse")
	}
	otApi := _otApi

	// Simple Field (toApi)
	_toApi, _toApiErr := readBuffer.ReadUint32("toApi", 32)
	if _toApiErr != nil {
		return nil, errors.Wrap(_toApiErr, "Error parsing 'toApi' field of CipConnectionManagerResponse")
	}
	toApi := _toApi

	// Implicit Field (replySize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	replySize, _replySizeErr := readBuffer.ReadUint8("replySize", 8)
	_ = replySize
	if _replySizeErr != nil {
		return nil, errors.Wrap(_replySizeErr, "Error parsing 'replySize' field of CipConnectionManagerResponse")
	}

	var reservedField1 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipConnectionManagerResponse")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("CipConnectionManagerResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectionManagerResponse")
	}

	// Create a partially initialized instance
	_child := &_CipConnectionManagerResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		OtConnectionId:         otConnectionId,
		ToConnectionId:         toConnectionId,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		OtApi:                  otApi,
		ToApi:                  toApi,
		reservedField0:         reservedField0,
		reservedField1:         reservedField1,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipConnectionManagerResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectionManagerResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectionManagerResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectionManagerResponse")
		}

		// Reserved Field (reserved)
		{
			var reserved uint32 = uint32(0x000000)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint32(0x000000),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint32("reserved", 24, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (otConnectionId)
		otConnectionId := uint32(m.GetOtConnectionId())
		_otConnectionIdErr := writeBuffer.WriteUint32("otConnectionId", 32, (otConnectionId))
		if _otConnectionIdErr != nil {
			return errors.Wrap(_otConnectionIdErr, "Error serializing 'otConnectionId' field")
		}

		// Simple Field (toConnectionId)
		toConnectionId := uint32(m.GetToConnectionId())
		_toConnectionIdErr := writeBuffer.WriteUint32("toConnectionId", 32, (toConnectionId))
		if _toConnectionIdErr != nil {
			return errors.Wrap(_toConnectionIdErr, "Error serializing 'toConnectionId' field")
		}

		// Simple Field (connectionSerialNumber)
		connectionSerialNumber := uint16(m.GetConnectionSerialNumber())
		_connectionSerialNumberErr := writeBuffer.WriteUint16("connectionSerialNumber", 16, (connectionSerialNumber))
		if _connectionSerialNumberErr != nil {
			return errors.Wrap(_connectionSerialNumberErr, "Error serializing 'connectionSerialNumber' field")
		}

		// Simple Field (originatorVendorId)
		originatorVendorId := uint16(m.GetOriginatorVendorId())
		_originatorVendorIdErr := writeBuffer.WriteUint16("originatorVendorId", 16, (originatorVendorId))
		if _originatorVendorIdErr != nil {
			return errors.Wrap(_originatorVendorIdErr, "Error serializing 'originatorVendorId' field")
		}

		// Simple Field (originatorSerialNumber)
		originatorSerialNumber := uint32(m.GetOriginatorSerialNumber())
		_originatorSerialNumberErr := writeBuffer.WriteUint32("originatorSerialNumber", 32, (originatorSerialNumber))
		if _originatorSerialNumberErr != nil {
			return errors.Wrap(_originatorSerialNumberErr, "Error serializing 'originatorSerialNumber' field")
		}

		// Simple Field (otApi)
		otApi := uint32(m.GetOtApi())
		_otApiErr := writeBuffer.WriteUint32("otApi", 32, (otApi))
		if _otApiErr != nil {
			return errors.Wrap(_otApiErr, "Error serializing 'otApi' field")
		}

		// Simple Field (toApi)
		toApi := uint32(m.GetToApi())
		_toApiErr := writeBuffer.WriteUint32("toApi", 32, (toApi))
		if _toApiErr != nil {
			return errors.Wrap(_toApiErr, "Error serializing 'toApi' field")
		}

		// Implicit Field (replySize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		replySize := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8(uint8(30)))
		_replySizeErr := writeBuffer.WriteUint8("replySize", 8, (replySize))
		if _replySizeErr != nil {
			return errors.Wrap(_replySizeErr, "Error serializing 'replySize' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField1 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("CipConnectionManagerResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectionManagerResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectionManagerResponse) isCipConnectionManagerResponse() bool {
	return true
}

func (m *_CipConnectionManagerResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
