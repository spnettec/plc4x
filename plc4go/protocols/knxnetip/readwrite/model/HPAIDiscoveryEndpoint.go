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

// HPAIDiscoveryEndpoint is the corresponding interface of HPAIDiscoveryEndpoint
type HPAIDiscoveryEndpoint interface {
	utils.LengthAware
	utils.Serializable
	// GetHostProtocolCode returns HostProtocolCode (property field)
	GetHostProtocolCode() HostProtocolCode
	// GetIpAddress returns IpAddress (property field)
	GetIpAddress() IPAddress
	// GetIpPort returns IpPort (property field)
	GetIpPort() uint16
}

// HPAIDiscoveryEndpointExactly can be used when we want exactly this type and not a type which fulfills HPAIDiscoveryEndpoint.
// This is useful for switch cases.
type HPAIDiscoveryEndpointExactly interface {
	HPAIDiscoveryEndpoint
	isHPAIDiscoveryEndpoint() bool
}

// _HPAIDiscoveryEndpoint is the data-structure of this message
type _HPAIDiscoveryEndpoint struct {
	HostProtocolCode HostProtocolCode
	IpAddress        IPAddress
	IpPort           uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HPAIDiscoveryEndpoint) GetHostProtocolCode() HostProtocolCode {
	return m.HostProtocolCode
}

func (m *_HPAIDiscoveryEndpoint) GetIpAddress() IPAddress {
	return m.IpAddress
}

func (m *_HPAIDiscoveryEndpoint) GetIpPort() uint16 {
	return m.IpPort
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHPAIDiscoveryEndpoint factory function for _HPAIDiscoveryEndpoint
func NewHPAIDiscoveryEndpoint(hostProtocolCode HostProtocolCode, ipAddress IPAddress, ipPort uint16) *_HPAIDiscoveryEndpoint {
	return &_HPAIDiscoveryEndpoint{HostProtocolCode: hostProtocolCode, IpAddress: ipAddress, IpPort: ipPort}
}

// Deprecated: use the interface for direct cast
func CastHPAIDiscoveryEndpoint(structType interface{}) HPAIDiscoveryEndpoint {
	if casted, ok := structType.(HPAIDiscoveryEndpoint); ok {
		return casted
	}
	if casted, ok := structType.(*HPAIDiscoveryEndpoint); ok {
		return *casted
	}
	return nil
}

func (m *_HPAIDiscoveryEndpoint) GetTypeName() string {
	return "HPAIDiscoveryEndpoint"
}

func (m *_HPAIDiscoveryEndpoint) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_HPAIDiscoveryEndpoint) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (hostProtocolCode)
	lengthInBits += 8

	// Simple field (ipAddress)
	lengthInBits += m.IpAddress.GetLengthInBits()

	// Simple field (ipPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_HPAIDiscoveryEndpoint) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func HPAIDiscoveryEndpointParse(theBytes []byte) (HPAIDiscoveryEndpoint, error) {
	return HPAIDiscoveryEndpointParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func HPAIDiscoveryEndpointParseWithBuffer(readBuffer utils.ReadBuffer) (HPAIDiscoveryEndpoint, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HPAIDiscoveryEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HPAIDiscoveryEndpoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field of HPAIDiscoveryEndpoint")
	}

	// Simple Field (hostProtocolCode)
	if pullErr := readBuffer.PullContext("hostProtocolCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hostProtocolCode")
	}
	_hostProtocolCode, _hostProtocolCodeErr := HostProtocolCodeParseWithBuffer(readBuffer)
	if _hostProtocolCodeErr != nil {
		return nil, errors.Wrap(_hostProtocolCodeErr, "Error parsing 'hostProtocolCode' field of HPAIDiscoveryEndpoint")
	}
	hostProtocolCode := _hostProtocolCode
	if closeErr := readBuffer.CloseContext("hostProtocolCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hostProtocolCode")
	}

	// Simple Field (ipAddress)
	if pullErr := readBuffer.PullContext("ipAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipAddress")
	}
	_ipAddress, _ipAddressErr := IPAddressParseWithBuffer(readBuffer)
	if _ipAddressErr != nil {
		return nil, errors.Wrap(_ipAddressErr, "Error parsing 'ipAddress' field of HPAIDiscoveryEndpoint")
	}
	ipAddress := _ipAddress.(IPAddress)
	if closeErr := readBuffer.CloseContext("ipAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipAddress")
	}

	// Simple Field (ipPort)
	_ipPort, _ipPortErr := readBuffer.ReadUint16("ipPort", 16)
	if _ipPortErr != nil {
		return nil, errors.Wrap(_ipPortErr, "Error parsing 'ipPort' field of HPAIDiscoveryEndpoint")
	}
	ipPort := _ipPort

	if closeErr := readBuffer.CloseContext("HPAIDiscoveryEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HPAIDiscoveryEndpoint")
	}

	// Create the instance
	return &_HPAIDiscoveryEndpoint{
		HostProtocolCode: hostProtocolCode,
		IpAddress:        ipAddress,
		IpPort:           ipPort,
	}, nil
}

func (m *_HPAIDiscoveryEndpoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HPAIDiscoveryEndpoint) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("HPAIDiscoveryEndpoint"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HPAIDiscoveryEndpoint")
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.GetLengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (hostProtocolCode)
	if pushErr := writeBuffer.PushContext("hostProtocolCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for hostProtocolCode")
	}
	_hostProtocolCodeErr := writeBuffer.WriteSerializable(m.GetHostProtocolCode())
	if popErr := writeBuffer.PopContext("hostProtocolCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for hostProtocolCode")
	}
	if _hostProtocolCodeErr != nil {
		return errors.Wrap(_hostProtocolCodeErr, "Error serializing 'hostProtocolCode' field")
	}

	// Simple Field (ipAddress)
	if pushErr := writeBuffer.PushContext("ipAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ipAddress")
	}
	_ipAddressErr := writeBuffer.WriteSerializable(m.GetIpAddress())
	if popErr := writeBuffer.PopContext("ipAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ipAddress")
	}
	if _ipAddressErr != nil {
		return errors.Wrap(_ipAddressErr, "Error serializing 'ipAddress' field")
	}

	// Simple Field (ipPort)
	ipPort := uint16(m.GetIpPort())
	_ipPortErr := writeBuffer.WriteUint16("ipPort", 16, (ipPort))
	if _ipPortErr != nil {
		return errors.Wrap(_ipPortErr, "Error serializing 'ipPort' field")
	}

	if popErr := writeBuffer.PopContext("HPAIDiscoveryEndpoint"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HPAIDiscoveryEndpoint")
	}
	return nil
}

func (m *_HPAIDiscoveryEndpoint) isHPAIDiscoveryEndpoint() bool {
	return true
}

func (m *_HPAIDiscoveryEndpoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
