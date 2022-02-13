/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ConnectionStateResponse struct {
	*KnxNetIpMessage
	CommunicationChannelId uint8
	Status                 Status
}

// The corresponding interface
type IConnectionStateResponse interface {
	// GetCommunicationChannelId returns CommunicationChannelId
	GetCommunicationChannelId() uint8
	// GetStatus returns Status
	GetStatus() Status
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ConnectionStateResponse) MsgType() uint16 {
	return 0x0208
}

func (m *ConnectionStateResponse) GetMsgType() uint16 {
	return 0x0208
}

func (m *ConnectionStateResponse) InitializeParent(parent *KnxNetIpMessage) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ConnectionStateResponse) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *ConnectionStateResponse) GetStatus() Status {
	return m.Status
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewConnectionStateResponse(communicationChannelId uint8, status Status) *KnxNetIpMessage {
	child := &ConnectionStateResponse{
		CommunicationChannelId: communicationChannelId,
		Status:                 status,
		KnxNetIpMessage:        NewKnxNetIpMessage(),
	}
	child.Child = child
	return child.KnxNetIpMessage
}

func CastConnectionStateResponse(structType interface{}) *ConnectionStateResponse {
	castFunc := func(typ interface{}) *ConnectionStateResponse {
		if casted, ok := typ.(ConnectionStateResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ConnectionStateResponse); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastConnectionStateResponse(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastConnectionStateResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ConnectionStateResponse) GetTypeName() string {
	return "ConnectionStateResponse"
}

func (m *ConnectionStateResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ConnectionStateResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *ConnectionStateResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ConnectionStateResponseParse(readBuffer utils.ReadBuffer) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("ConnectionStateResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (communicationChannelId)
	_communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field")
	}
	communicationChannelId := _communicationChannelId

	// Simple Field (status)
	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, pullErr
	}
	_status, _statusErr := StatusParse(readBuffer)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	status := _status
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ConnectionStateResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ConnectionStateResponse{
		CommunicationChannelId: communicationChannelId,
		Status:                 status,
		KnxNetIpMessage:        &KnxNetIpMessage{},
	}
	_child.KnxNetIpMessage.Child = _child
	return _child.KnxNetIpMessage, nil
}

func (m *ConnectionStateResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionStateResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (communicationChannelId)
		communicationChannelId := uint8(m.CommunicationChannelId)
		_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, (communicationChannelId))
		if _communicationChannelIdErr != nil {
			return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
		}

		// Simple Field (status)
		if pushErr := writeBuffer.PushContext("status"); pushErr != nil {
			return pushErr
		}
		_statusErr := m.Status.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("status"); popErr != nil {
			return popErr
		}
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		if popErr := writeBuffer.PopContext("ConnectionStateResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ConnectionStateResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
