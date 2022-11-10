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

// NLMRejectRouterToNetwork is the corresponding interface of NLMRejectRouterToNetwork
type NLMRejectRouterToNetwork interface {
	utils.LengthAware
	utils.Serializable
	NLM
	// GetRejectReason returns RejectReason (property field)
	GetRejectReason() NLMRejectRouterToNetworkRejectReason
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
}

// NLMRejectRouterToNetworkExactly can be used when we want exactly this type and not a type which fulfills NLMRejectRouterToNetwork.
// This is useful for switch cases.
type NLMRejectRouterToNetworkExactly interface {
	NLMRejectRouterToNetwork
	isNLMRejectRouterToNetwork() bool
}

// _NLMRejectRouterToNetwork is the data-structure of this message
type _NLMRejectRouterToNetwork struct {
	*_NLM
	RejectReason              NLMRejectRouterToNetworkRejectReason
	DestinationNetworkAddress uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMRejectRouterToNetwork) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMRejectRouterToNetwork) InitializeParent(parent NLM) {}

func (m *_NLMRejectRouterToNetwork) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMRejectRouterToNetwork) GetRejectReason() NLMRejectRouterToNetworkRejectReason {
	return m.RejectReason
}

func (m *_NLMRejectRouterToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMRejectRouterToNetwork factory function for _NLMRejectRouterToNetwork
func NewNLMRejectRouterToNetwork(rejectReason NLMRejectRouterToNetworkRejectReason, destinationNetworkAddress uint16, apduLength uint16) *_NLMRejectRouterToNetwork {
	_result := &_NLMRejectRouterToNetwork{
		RejectReason:              rejectReason,
		DestinationNetworkAddress: destinationNetworkAddress,
		_NLM:                      NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMRejectRouterToNetwork(structType interface{}) NLMRejectRouterToNetwork {
	if casted, ok := structType.(NLMRejectRouterToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMRejectRouterToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMRejectRouterToNetwork) GetTypeName() string {
	return "NLMRejectRouterToNetwork"
}

func (m *_NLMRejectRouterToNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NLMRejectRouterToNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (rejectReason)
	lengthInBits += 8

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *_NLMRejectRouterToNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMRejectRouterToNetworkParse(theBytes []byte, apduLength uint16) (NLMRejectRouterToNetwork, error) {
	return NLMRejectRouterToNetworkParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), apduLength) // TODO: get endianness from mspec
}

func NLMRejectRouterToNetworkParseWithBuffer(readBuffer utils.ReadBuffer, apduLength uint16) (NLMRejectRouterToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMRejectRouterToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMRejectRouterToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (rejectReason)
	if pullErr := readBuffer.PullContext("rejectReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for rejectReason")
	}
	_rejectReason, _rejectReasonErr := NLMRejectRouterToNetworkRejectReasonParseWithBuffer(readBuffer)
	if _rejectReasonErr != nil {
		return nil, errors.Wrap(_rejectReasonErr, "Error parsing 'rejectReason' field of NLMRejectRouterToNetwork")
	}
	rejectReason := _rejectReason
	if closeErr := readBuffer.CloseContext("rejectReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for rejectReason")
	}

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field of NLMRejectRouterToNetwork")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	if closeErr := readBuffer.CloseContext("NLMRejectRouterToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMRejectRouterToNetwork")
	}

	// Create a partially initialized instance
	_child := &_NLMRejectRouterToNetwork{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		RejectReason:              rejectReason,
		DestinationNetworkAddress: destinationNetworkAddress,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMRejectRouterToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMRejectRouterToNetwork) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMRejectRouterToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMRejectRouterToNetwork")
		}

		// Simple Field (rejectReason)
		if pushErr := writeBuffer.PushContext("rejectReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for rejectReason")
		}
		_rejectReasonErr := writeBuffer.WriteSerializable(m.GetRejectReason())
		if popErr := writeBuffer.PopContext("rejectReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for rejectReason")
		}
		if _rejectReasonErr != nil {
			return errors.Wrap(_rejectReasonErr, "Error serializing 'rejectReason' field")
		}

		// Simple Field (destinationNetworkAddress)
		destinationNetworkAddress := uint16(m.GetDestinationNetworkAddress())
		_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, (destinationNetworkAddress))
		if _destinationNetworkAddressErr != nil {
			return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
		}

		if popErr := writeBuffer.PopContext("NLMRejectRouterToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMRejectRouterToNetwork")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_NLMRejectRouterToNetwork) isNLMRejectRouterToNetwork() bool {
	return true
}

func (m *_NLMRejectRouterToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
