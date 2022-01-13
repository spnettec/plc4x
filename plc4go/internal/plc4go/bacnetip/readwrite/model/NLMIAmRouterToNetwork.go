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
type NLMIAmRouterToNetwork struct {
	*NLM
	DestinationNetworkAddress []uint16
}

// The corresponding interface
type INLMIAmRouterToNetwork interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *NLMIAmRouterToNetwork) MessageType() uint8 {
	return 0x01
}

func (m *NLMIAmRouterToNetwork) InitializeParent(parent *NLM, vendorId *uint16) {
	m.VendorId = vendorId
}

func NewNLMIAmRouterToNetwork(destinationNetworkAddress []uint16, vendorId *uint16) *NLM {
	child := &NLMIAmRouterToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		NLM:                       NewNLM(vendorId),
	}
	child.Child = child
	return child.NLM
}

func CastNLMIAmRouterToNetwork(structType interface{}) *NLMIAmRouterToNetwork {
	castFunc := func(typ interface{}) *NLMIAmRouterToNetwork {
		if casted, ok := typ.(NLMIAmRouterToNetwork); ok {
			return &casted
		}
		if casted, ok := typ.(*NLMIAmRouterToNetwork); ok {
			return casted
		}
		if casted, ok := typ.(NLM); ok {
			return CastNLMIAmRouterToNetwork(casted.Child)
		}
		if casted, ok := typ.(*NLM); ok {
			return CastNLMIAmRouterToNetwork(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *NLMIAmRouterToNetwork) GetTypeName() string {
	return "NLMIAmRouterToNetwork"
}

func (m *NLMIAmRouterToNetwork) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *NLMIAmRouterToNetwork) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Array field
	if len(m.DestinationNetworkAddress) > 0 {
		lengthInBits += 16 * uint16(len(m.DestinationNetworkAddress))
	}

	return lengthInBits
}

func (m *NLMIAmRouterToNetwork) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func NLMIAmRouterToNetworkParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (*NLM, error) {
	if pullErr := readBuffer.PullContext("NLMIAmRouterToNetwork"); pullErr != nil {
		return nil, pullErr
	}

	// Array field (destinationNetworkAddress)
	if pullErr := readBuffer.PullContext("destinationNetworkAddress", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	destinationNetworkAddress := make([]uint16, 0)
	{
		_destinationNetworkAddressLength := uint16(apduLength) - uint16(uint16(utils.InlineIf(bool(bool(bool(bool((messageType) >= (128)))) && bool(bool(bool((messageType) <= (255))))), func() interface{} { return uint16(uint16(3)) }, func() interface{} { return uint16(uint16(1)) }).(uint16)))
		_destinationNetworkAddressEndPos := readBuffer.GetPos() + uint16(_destinationNetworkAddressLength)
		for readBuffer.GetPos() < _destinationNetworkAddressEndPos {
			_item, _err := readBuffer.ReadUint16("", 16)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'destinationNetworkAddress' field")
			}
			destinationNetworkAddress = append(destinationNetworkAddress, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("destinationNetworkAddress", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("NLMIAmRouterToNetwork"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &NLMIAmRouterToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		NLM:                       &NLM{},
	}
	_child.NLM.Child = _child
	return _child.NLM, nil
}

func (m *NLMIAmRouterToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMIAmRouterToNetwork"); pushErr != nil {
			return pushErr
		}

		// Array Field (destinationNetworkAddress)
		if m.DestinationNetworkAddress != nil {
			if pushErr := writeBuffer.PushContext("destinationNetworkAddress", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.DestinationNetworkAddress {
				_elementErr := writeBuffer.WriteUint16("", 16, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'destinationNetworkAddress' field")
				}
			}
			if popErr := writeBuffer.PopContext("destinationNetworkAddress", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("NLMIAmRouterToNetwork"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NLMIAmRouterToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
