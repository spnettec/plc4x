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
type MACAddress struct {
	Addr []byte
}

// The corresponding interface
type IMACAddress interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewMACAddress(addr []byte) *MACAddress {
	return &MACAddress{Addr: addr}
}

func CastMACAddress(structType interface{}) *MACAddress {
	castFunc := func(typ interface{}) *MACAddress {
		if casted, ok := typ.(MACAddress); ok {
			return &casted
		}
		if casted, ok := typ.(*MACAddress); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MACAddress) GetTypeName() string {
	return "MACAddress"
}

func (m *MACAddress) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *MACAddress) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Addr) > 0 {
		lengthInBits += 8 * uint16(len(m.Addr))
	}

	return lengthInBits
}

func (m *MACAddress) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MACAddressParse(readBuffer utils.ReadBuffer) (*MACAddress, error) {
	if pullErr := readBuffer.PullContext("MACAddress"); pullErr != nil {
		return nil, pullErr
	}
	// Byte Array field (addr)
	numberOfBytes := int(uint16(6))
	addr, _readArrayErr := readBuffer.ReadByteArray("addr", numberOfBytes)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'addr' field")
	}

	if closeErr := readBuffer.CloseContext("MACAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewMACAddress(addr), nil
}

func (m *MACAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("MACAddress"); pushErr != nil {
		return pushErr
	}

	// Array Field (addr)
	if m.Addr != nil {
		// Byte Array field (addr)
		_writeArrayErr := writeBuffer.WriteByteArray("addr", m.Addr)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'addr' field")
		}
	}

	if popErr := writeBuffer.PopContext("MACAddress"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *MACAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
