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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// DisconnectResponse is the corresponding interface of DisconnectResponse
type DisconnectResponse interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetStatus returns Status (property field)
	GetStatus() Status
}

// DisconnectResponseExactly can be used when we want exactly this type and not a type which fulfills DisconnectResponse.
// This is useful for switch cases.
type DisconnectResponseExactly interface {
	DisconnectResponse
	isDisconnectResponse() bool
}

// _DisconnectResponse is the data-structure of this message
type _DisconnectResponse struct {
	*_KnxNetIpMessage
        CommunicationChannelId uint8
        Status Status
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DisconnectResponse)  GetMsgType() uint16 {
return 0x020A}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DisconnectResponse) InitializeParent(parent KnxNetIpMessage ) {}

func (m *_DisconnectResponse)  GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DisconnectResponse) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_DisconnectResponse) GetStatus() Status {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewDisconnectResponse factory function for _DisconnectResponse
func NewDisconnectResponse( communicationChannelId uint8 , status Status ) *_DisconnectResponse {
	_result := &_DisconnectResponse{
		CommunicationChannelId: communicationChannelId,
		Status: status,
    	_KnxNetIpMessage: NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDisconnectResponse(structType interface{}) DisconnectResponse {
    if casted, ok := structType.(DisconnectResponse); ok {
		return casted
	}
	if casted, ok := structType.(*DisconnectResponse); ok {
		return *casted
	}
	return nil
}

func (m *_DisconnectResponse) GetTypeName() string {
	return "DisconnectResponse"
}

func (m *_DisconnectResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (communicationChannelId)
	lengthInBits += 8;

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}


func (m *_DisconnectResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DisconnectResponseParse(theBytes []byte) (DisconnectResponse, error) {
	return DisconnectResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func DisconnectResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DisconnectResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DisconnectResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DisconnectResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (communicationChannelId)
_communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field of DisconnectResponse")
	}
	communicationChannelId := _communicationChannelId

	// Simple Field (status)
	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for status")
	}
_status, _statusErr := StatusParseWithBuffer(ctx, readBuffer)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of DisconnectResponse")
	}
	status := _status
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for status")
	}

	if closeErr := readBuffer.CloseContext("DisconnectResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DisconnectResponse")
	}

	// Create a partially initialized instance
	_child := &_DisconnectResponse{
		_KnxNetIpMessage: &_KnxNetIpMessage{
		},
		CommunicationChannelId: communicationChannelId,
		Status: status,
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_DisconnectResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DisconnectResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DisconnectResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DisconnectResponse")
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
	_statusErr := writeBuffer.WriteSerializable(ctx, m.GetStatus())
	if popErr := writeBuffer.PopContext("status"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for status")
	}
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

		if popErr := writeBuffer.PopContext("DisconnectResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DisconnectResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_DisconnectResponse) isDisconnectResponse() bool {
	return true
}

func (m *_DisconnectResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



