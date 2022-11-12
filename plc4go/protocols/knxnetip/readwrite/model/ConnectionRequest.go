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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// ConnectionRequest is the corresponding interface of ConnectionRequest
type ConnectionRequest interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetHpaiDiscoveryEndpoint returns HpaiDiscoveryEndpoint (property field)
	GetHpaiDiscoveryEndpoint() HPAIDiscoveryEndpoint
	// GetHpaiDataEndpoint returns HpaiDataEndpoint (property field)
	GetHpaiDataEndpoint() HPAIDataEndpoint
	// GetConnectionRequestInformation returns ConnectionRequestInformation (property field)
	GetConnectionRequestInformation() ConnectionRequestInformation
}

// ConnectionRequestExactly can be used when we want exactly this type and not a type which fulfills ConnectionRequest.
// This is useful for switch cases.
type ConnectionRequestExactly interface {
	ConnectionRequest
	isConnectionRequest() bool
}

// _ConnectionRequest is the data-structure of this message
type _ConnectionRequest struct {
	*_KnxNetIpMessage
        HpaiDiscoveryEndpoint HPAIDiscoveryEndpoint
        HpaiDataEndpoint HPAIDataEndpoint
        ConnectionRequestInformation ConnectionRequestInformation
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionRequest)  GetMsgType() uint16 {
return 0x0205}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionRequest) InitializeParent(parent KnxNetIpMessage ) {}

func (m *_ConnectionRequest)  GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionRequest) GetHpaiDiscoveryEndpoint() HPAIDiscoveryEndpoint {
	return m.HpaiDiscoveryEndpoint
}

func (m *_ConnectionRequest) GetHpaiDataEndpoint() HPAIDataEndpoint {
	return m.HpaiDataEndpoint
}

func (m *_ConnectionRequest) GetConnectionRequestInformation() ConnectionRequestInformation {
	return m.ConnectionRequestInformation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewConnectionRequest factory function for _ConnectionRequest
func NewConnectionRequest( hpaiDiscoveryEndpoint HPAIDiscoveryEndpoint , hpaiDataEndpoint HPAIDataEndpoint , connectionRequestInformation ConnectionRequestInformation ) *_ConnectionRequest {
	_result := &_ConnectionRequest{
		HpaiDiscoveryEndpoint: hpaiDiscoveryEndpoint,
		HpaiDataEndpoint: hpaiDataEndpoint,
		ConnectionRequestInformation: connectionRequestInformation,
    	_KnxNetIpMessage: NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectionRequest(structType interface{}) ConnectionRequest {
    if casted, ok := structType.(ConnectionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionRequest) GetTypeName() string {
	return "ConnectionRequest"
}

func (m *_ConnectionRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ConnectionRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (hpaiDiscoveryEndpoint)
	lengthInBits += m.HpaiDiscoveryEndpoint.GetLengthInBits()

	// Simple field (hpaiDataEndpoint)
	lengthInBits += m.HpaiDataEndpoint.GetLengthInBits()

	// Simple field (connectionRequestInformation)
	lengthInBits += m.ConnectionRequestInformation.GetLengthInBits()

	return lengthInBits
}


func (m *_ConnectionRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConnectionRequestParse(theBytes []byte) (ConnectionRequest, error) {
	return ConnectionRequestParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func ConnectionRequestParseWithBuffer(readBuffer utils.ReadBuffer) (ConnectionRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (hpaiDiscoveryEndpoint)
	if pullErr := readBuffer.PullContext("hpaiDiscoveryEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hpaiDiscoveryEndpoint")
	}
_hpaiDiscoveryEndpoint, _hpaiDiscoveryEndpointErr := HPAIDiscoveryEndpointParseWithBuffer(readBuffer)
	if _hpaiDiscoveryEndpointErr != nil {
		return nil, errors.Wrap(_hpaiDiscoveryEndpointErr, "Error parsing 'hpaiDiscoveryEndpoint' field of ConnectionRequest")
	}
	hpaiDiscoveryEndpoint := _hpaiDiscoveryEndpoint.(HPAIDiscoveryEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiDiscoveryEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hpaiDiscoveryEndpoint")
	}

	// Simple Field (hpaiDataEndpoint)
	if pullErr := readBuffer.PullContext("hpaiDataEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hpaiDataEndpoint")
	}
_hpaiDataEndpoint, _hpaiDataEndpointErr := HPAIDataEndpointParseWithBuffer(readBuffer)
	if _hpaiDataEndpointErr != nil {
		return nil, errors.Wrap(_hpaiDataEndpointErr, "Error parsing 'hpaiDataEndpoint' field of ConnectionRequest")
	}
	hpaiDataEndpoint := _hpaiDataEndpoint.(HPAIDataEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiDataEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hpaiDataEndpoint")
	}

	// Simple Field (connectionRequestInformation)
	if pullErr := readBuffer.PullContext("connectionRequestInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for connectionRequestInformation")
	}
_connectionRequestInformation, _connectionRequestInformationErr := ConnectionRequestInformationParseWithBuffer(readBuffer)
	if _connectionRequestInformationErr != nil {
		return nil, errors.Wrap(_connectionRequestInformationErr, "Error parsing 'connectionRequestInformation' field of ConnectionRequest")
	}
	connectionRequestInformation := _connectionRequestInformation.(ConnectionRequestInformation)
	if closeErr := readBuffer.CloseContext("connectionRequestInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for connectionRequestInformation")
	}

	if closeErr := readBuffer.CloseContext("ConnectionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequest")
	}

	// Create a partially initialized instance
	_child := &_ConnectionRequest{
		_KnxNetIpMessage: &_KnxNetIpMessage{
		},
		HpaiDiscoveryEndpoint: hpaiDiscoveryEndpoint,
		HpaiDataEndpoint: hpaiDataEndpoint,
		ConnectionRequestInformation: connectionRequestInformation,
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_ConnectionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionRequest")
		}

	// Simple Field (hpaiDiscoveryEndpoint)
	if pushErr := writeBuffer.PushContext("hpaiDiscoveryEndpoint"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for hpaiDiscoveryEndpoint")
	}
	_hpaiDiscoveryEndpointErr := writeBuffer.WriteSerializable(m.GetHpaiDiscoveryEndpoint())
	if popErr := writeBuffer.PopContext("hpaiDiscoveryEndpoint"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for hpaiDiscoveryEndpoint")
	}
	if _hpaiDiscoveryEndpointErr != nil {
		return errors.Wrap(_hpaiDiscoveryEndpointErr, "Error serializing 'hpaiDiscoveryEndpoint' field")
	}

	// Simple Field (hpaiDataEndpoint)
	if pushErr := writeBuffer.PushContext("hpaiDataEndpoint"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for hpaiDataEndpoint")
	}
	_hpaiDataEndpointErr := writeBuffer.WriteSerializable(m.GetHpaiDataEndpoint())
	if popErr := writeBuffer.PopContext("hpaiDataEndpoint"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for hpaiDataEndpoint")
	}
	if _hpaiDataEndpointErr != nil {
		return errors.Wrap(_hpaiDataEndpointErr, "Error serializing 'hpaiDataEndpoint' field")
	}

	// Simple Field (connectionRequestInformation)
	if pushErr := writeBuffer.PushContext("connectionRequestInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for connectionRequestInformation")
	}
	_connectionRequestInformationErr := writeBuffer.WriteSerializable(m.GetConnectionRequestInformation())
	if popErr := writeBuffer.PopContext("connectionRequestInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for connectionRequestInformation")
	}
	if _connectionRequestInformationErr != nil {
		return errors.Wrap(_connectionRequestInformationErr, "Error serializing 'connectionRequestInformation' field")
	}

		if popErr := writeBuffer.PopContext("ConnectionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_ConnectionRequest) isConnectionRequest() bool {
	return true
}

func (m *_ConnectionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



