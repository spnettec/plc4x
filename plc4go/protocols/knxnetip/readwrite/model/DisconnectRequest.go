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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// DisconnectRequest is the corresponding interface of DisconnectRequest
type DisconnectRequest interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetHpaiControlEndpoint returns HpaiControlEndpoint (property field)
	GetHpaiControlEndpoint() HPAIControlEndpoint
}

// DisconnectRequestExactly can be used when we want exactly this type and not a type which fulfills DisconnectRequest.
// This is useful for switch cases.
type DisconnectRequestExactly interface {
	DisconnectRequest
	isDisconnectRequest() bool
}

// _DisconnectRequest is the data-structure of this message
type _DisconnectRequest struct {
	*_KnxNetIpMessage
	CommunicationChannelId uint8
	HpaiControlEndpoint    HPAIControlEndpoint
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DisconnectRequest) GetMsgType() uint16 {
	return 0x0209
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DisconnectRequest) InitializeParent(parent KnxNetIpMessage) {}

func (m *_DisconnectRequest) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DisconnectRequest) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_DisconnectRequest) GetHpaiControlEndpoint() HPAIControlEndpoint {
	return m.HpaiControlEndpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDisconnectRequest factory function for _DisconnectRequest
func NewDisconnectRequest(communicationChannelId uint8, hpaiControlEndpoint HPAIControlEndpoint) *_DisconnectRequest {
	_result := &_DisconnectRequest{
		CommunicationChannelId: communicationChannelId,
		HpaiControlEndpoint:    hpaiControlEndpoint,
		_KnxNetIpMessage:       NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDisconnectRequest(structType interface{}) DisconnectRequest {
	if casted, ok := structType.(DisconnectRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DisconnectRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DisconnectRequest) GetTypeName() string {
	return "DisconnectRequest"
}

func (m *_DisconnectRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_DisconnectRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.GetLengthInBits()

	return lengthInBits
}

func (m *_DisconnectRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DisconnectRequestParse(readBuffer utils.ReadBuffer) (DisconnectRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DisconnectRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DisconnectRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (communicationChannelId)
	_communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field of DisconnectRequest")
	}
	communicationChannelId := _communicationChannelId

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of DisconnectRequest")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (hpaiControlEndpoint)
	if pullErr := readBuffer.PullContext("hpaiControlEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hpaiControlEndpoint")
	}
	_hpaiControlEndpoint, _hpaiControlEndpointErr := HPAIControlEndpointParse(readBuffer)
	if _hpaiControlEndpointErr != nil {
		return nil, errors.Wrap(_hpaiControlEndpointErr, "Error parsing 'hpaiControlEndpoint' field of DisconnectRequest")
	}
	hpaiControlEndpoint := _hpaiControlEndpoint.(HPAIControlEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiControlEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hpaiControlEndpoint")
	}

	if closeErr := readBuffer.CloseContext("DisconnectRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DisconnectRequest")
	}

	// Create a partially initialized instance
	_child := &_DisconnectRequest{
		_KnxNetIpMessage:       &_KnxNetIpMessage{},
		CommunicationChannelId: communicationChannelId,
		HpaiControlEndpoint:    hpaiControlEndpoint,
		reservedField0:         reservedField0,
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_DisconnectRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DisconnectRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DisconnectRequest")
		}

		// Simple Field (communicationChannelId)
		communicationChannelId := uint8(m.GetCommunicationChannelId())
		_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, (communicationChannelId))
		if _communicationChannelIdErr != nil {
			return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (hpaiControlEndpoint)
		if pushErr := writeBuffer.PushContext("hpaiControlEndpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hpaiControlEndpoint")
		}
		_hpaiControlEndpointErr := writeBuffer.WriteSerializable(m.GetHpaiControlEndpoint())
		if popErr := writeBuffer.PopContext("hpaiControlEndpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hpaiControlEndpoint")
		}
		if _hpaiControlEndpointErr != nil {
			return errors.Wrap(_hpaiControlEndpointErr, "Error serializing 'hpaiControlEndpoint' field")
		}

		if popErr := writeBuffer.PopContext("DisconnectRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DisconnectRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_DisconnectRequest) isDisconnectRequest() bool {
	return true
}

func (m *_DisconnectRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
