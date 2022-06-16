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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// NLMDisconnectConnectionToNetwork is the corresponding interface of NLMDisconnectConnectionToNetwork
type NLMDisconnectConnectionToNetwork interface {
	NLM
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _NLMDisconnectConnectionToNetwork is the data-structure of this message
type _NLMDisconnectConnectionToNetwork struct {
	*_NLM
	DestinationNetworkAddress uint16

	// Arguments.
	ApduLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMDisconnectConnectionToNetwork) GetMessageType() uint8 {
	return 0x09
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMDisconnectConnectionToNetwork) InitializeParent(parent NLM, vendorId *BACnetVendorId) {
	m.VendorId = vendorId
}

func (m *_NLMDisconnectConnectionToNetwork) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMDisconnectConnectionToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMDisconnectConnectionToNetwork factory function for _NLMDisconnectConnectionToNetwork
func NewNLMDisconnectConnectionToNetwork(destinationNetworkAddress uint16, vendorId *BACnetVendorId, apduLength uint16) *_NLMDisconnectConnectionToNetwork {
	_result := &_NLMDisconnectConnectionToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		_NLM:                      NewNLM(vendorId, apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMDisconnectConnectionToNetwork(structType interface{}) NLMDisconnectConnectionToNetwork {
	if casted, ok := structType.(NLMDisconnectConnectionToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMDisconnectConnectionToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMDisconnectConnectionToNetwork) GetTypeName() string {
	return "NLMDisconnectConnectionToNetwork"
}

func (m *_NLMDisconnectConnectionToNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NLMDisconnectConnectionToNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *_NLMDisconnectConnectionToNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMDisconnectConnectionToNetworkParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (NLMDisconnectConnectionToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMDisconnectConnectionToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMDisconnectConnectionToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	if closeErr := readBuffer.CloseContext("NLMDisconnectConnectionToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMDisconnectConnectionToNetwork")
	}

	// Create a partially initialized instance
	_child := &_NLMDisconnectConnectionToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		_NLM:                      &_NLM{},
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMDisconnectConnectionToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMDisconnectConnectionToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMDisconnectConnectionToNetwork")
		}

		// Simple Field (destinationNetworkAddress)
		destinationNetworkAddress := uint16(m.GetDestinationNetworkAddress())
		_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, (destinationNetworkAddress))
		if _destinationNetworkAddressErr != nil {
			return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
		}

		if popErr := writeBuffer.PopContext("NLMDisconnectConnectionToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMDisconnectConnectionToNetwork")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_NLMDisconnectConnectionToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
