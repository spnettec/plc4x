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
	"github.com/rs/zerolog/log"
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// ConnectionResponse is the corresponding interface of ConnectionResponse
type ConnectionResponse interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetStatus returns Status (property field)
	GetStatus() Status
	// GetHpaiDataEndpoint returns HpaiDataEndpoint (property field)
	GetHpaiDataEndpoint() HPAIDataEndpoint
	// GetConnectionResponseDataBlock returns ConnectionResponseDataBlock (property field)
	GetConnectionResponseDataBlock() ConnectionResponseDataBlock
}

// ConnectionResponseExactly can be used when we want exactly this type and not a type which fulfills ConnectionResponse.
// This is useful for switch cases.
type ConnectionResponseExactly interface {
	ConnectionResponse
	isConnectionResponse() bool
}

// _ConnectionResponse is the data-structure of this message
type _ConnectionResponse struct {
	*_KnxNetIpMessage
        CommunicationChannelId uint8
        Status Status
        HpaiDataEndpoint HPAIDataEndpoint
        ConnectionResponseDataBlock ConnectionResponseDataBlock
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionResponse)  GetMsgType() uint16 {
return 0x0206}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionResponse) InitializeParent(parent KnxNetIpMessage ) {}

func (m *_ConnectionResponse)  GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionResponse) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_ConnectionResponse) GetStatus() Status {
	return m.Status
}

func (m *_ConnectionResponse) GetHpaiDataEndpoint() HPAIDataEndpoint {
	return m.HpaiDataEndpoint
}

func (m *_ConnectionResponse) GetConnectionResponseDataBlock() ConnectionResponseDataBlock {
	return m.ConnectionResponseDataBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewConnectionResponse factory function for _ConnectionResponse
func NewConnectionResponse( communicationChannelId uint8 , status Status , hpaiDataEndpoint HPAIDataEndpoint , connectionResponseDataBlock ConnectionResponseDataBlock ) *_ConnectionResponse {
	_result := &_ConnectionResponse{
		CommunicationChannelId: communicationChannelId,
		Status: status,
		HpaiDataEndpoint: hpaiDataEndpoint,
		ConnectionResponseDataBlock: connectionResponseDataBlock,
    	_KnxNetIpMessage: NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectionResponse(structType interface{}) ConnectionResponse {
    if casted, ok := structType.(ConnectionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionResponse) GetTypeName() string {
	return "ConnectionResponse"
}

func (m *_ConnectionResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ConnectionResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (communicationChannelId)
	lengthInBits += 8;

	// Simple field (status)
	lengthInBits += 8

	// Optional Field (hpaiDataEndpoint)
	if m.HpaiDataEndpoint != nil {
		lengthInBits += m.HpaiDataEndpoint.GetLengthInBits()
	}

	// Optional Field (connectionResponseDataBlock)
	if m.ConnectionResponseDataBlock != nil {
		lengthInBits += m.ConnectionResponseDataBlock.GetLengthInBits()
	}

	return lengthInBits
}


func (m *_ConnectionResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConnectionResponseParse(readBuffer utils.ReadBuffer) (ConnectionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (communicationChannelId)
_communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field of ConnectionResponse")
	}
	communicationChannelId := _communicationChannelId

	// Simple Field (status)
	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for status")
	}
_status, _statusErr := StatusParse(readBuffer)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of ConnectionResponse")
	}
	status := _status
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for status")
	}

	// Optional Field (hpaiDataEndpoint) (Can be skipped, if a given expression evaluates to false)
	var hpaiDataEndpoint HPAIDataEndpoint = nil
	if bool((status) == (Status_NO_ERROR)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("hpaiDataEndpoint"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for hpaiDataEndpoint")
		}
_val, _err := HPAIDataEndpointParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'hpaiDataEndpoint' field of ConnectionResponse")
		default:
			hpaiDataEndpoint = _val.(HPAIDataEndpoint)
			if closeErr := readBuffer.CloseContext("hpaiDataEndpoint"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for hpaiDataEndpoint")
			}
		}
	}

	// Optional Field (connectionResponseDataBlock) (Can be skipped, if a given expression evaluates to false)
	var connectionResponseDataBlock ConnectionResponseDataBlock = nil
	if bool((status) == (Status_NO_ERROR)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("connectionResponseDataBlock"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for connectionResponseDataBlock")
		}
_val, _err := ConnectionResponseDataBlockParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'connectionResponseDataBlock' field of ConnectionResponse")
		default:
			connectionResponseDataBlock = _val.(ConnectionResponseDataBlock)
			if closeErr := readBuffer.CloseContext("connectionResponseDataBlock"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for connectionResponseDataBlock")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("ConnectionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionResponse")
	}

	// Create a partially initialized instance
	_child := &_ConnectionResponse{
		_KnxNetIpMessage: &_KnxNetIpMessage{
		},
		CommunicationChannelId: communicationChannelId,
		Status: status,
		HpaiDataEndpoint: hpaiDataEndpoint,
		ConnectionResponseDataBlock: connectionResponseDataBlock,
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_ConnectionResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionResponse")
		}

	// Simple Field (communicationChannelId)
	communicationChannelId := uint8(m.GetCommunicationChannelId())
	_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, (communicationChannelId))
	if _communicationChannelIdErr != nil {
		return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
	}

	// Simple Field (status)
	if pushErr := writeBuffer.PushContext("status"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for status")
	}
	_statusErr := writeBuffer.WriteSerializable(m.GetStatus())
	if popErr := writeBuffer.PopContext("status"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for status")
	}
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	// Optional Field (hpaiDataEndpoint) (Can be skipped, if the value is null)
	var hpaiDataEndpoint HPAIDataEndpoint = nil
	if m.GetHpaiDataEndpoint() != nil {
		if pushErr := writeBuffer.PushContext("hpaiDataEndpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hpaiDataEndpoint")
		}
		hpaiDataEndpoint = m.GetHpaiDataEndpoint()
		_hpaiDataEndpointErr := writeBuffer.WriteSerializable(hpaiDataEndpoint)
		if popErr := writeBuffer.PopContext("hpaiDataEndpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hpaiDataEndpoint")
		}
		if _hpaiDataEndpointErr != nil {
			return errors.Wrap(_hpaiDataEndpointErr, "Error serializing 'hpaiDataEndpoint' field")
		}
	}

	// Optional Field (connectionResponseDataBlock) (Can be skipped, if the value is null)
	var connectionResponseDataBlock ConnectionResponseDataBlock = nil
	if m.GetConnectionResponseDataBlock() != nil {
		if pushErr := writeBuffer.PushContext("connectionResponseDataBlock"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for connectionResponseDataBlock")
		}
		connectionResponseDataBlock = m.GetConnectionResponseDataBlock()
		_connectionResponseDataBlockErr := writeBuffer.WriteSerializable(connectionResponseDataBlock)
		if popErr := writeBuffer.PopContext("connectionResponseDataBlock"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for connectionResponseDataBlock")
		}
		if _connectionResponseDataBlockErr != nil {
			return errors.Wrap(_connectionResponseDataBlockErr, "Error serializing 'connectionResponseDataBlock' field")
		}
	}

		if popErr := writeBuffer.PopContext("ConnectionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_ConnectionResponse) isConnectionResponse() bool {
	return true
}

func (m *_ConnectionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



