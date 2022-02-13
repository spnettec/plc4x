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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type CipRRData struct {
	*EipPacket
	Exchange *CipExchange
}

// The corresponding interface
type ICipRRData interface {
	// GetExchange returns Exchange
	GetExchange() *CipExchange
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
func (m *CipRRData) Command() uint16 {
	return 0x006F
}

func (m *CipRRData) GetCommand() uint16 {
	return 0x006F
}

func (m *CipRRData) InitializeParent(parent *EipPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.EipPacket.SessionHandle = sessionHandle
	m.EipPacket.Status = status
	m.EipPacket.SenderContext = senderContext
	m.EipPacket.Options = options
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CipRRData) GetExchange() *CipExchange {
	return m.Exchange
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewCipRRData(exchange *CipExchange, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *EipPacket {
	child := &CipRRData{
		Exchange:  exchange,
		EipPacket: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	child.Child = child
	return child.EipPacket
}

func CastCipRRData(structType interface{}) *CipRRData {
	castFunc := func(typ interface{}) *CipRRData {
		if casted, ok := typ.(CipRRData); ok {
			return &casted
		}
		if casted, ok := typ.(*CipRRData); ok {
			return casted
		}
		if casted, ok := typ.(EipPacket); ok {
			return CastCipRRData(casted.Child)
		}
		if casted, ok := typ.(*EipPacket); ok {
			return CastCipRRData(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CipRRData) GetTypeName() string {
	return "CipRRData"
}

func (m *CipRRData) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *CipRRData) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (exchange)
	lengthInBits += m.Exchange.LengthInBits()

	return lengthInBits
}

func (m *CipRRData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CipRRDataParse(readBuffer utils.ReadBuffer, len uint16) (*EipPacket, error) {
	if pullErr := readBuffer.PullContext("CipRRData"); pullErr != nil {
		return nil, pullErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint32("reserved", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint32(0x00000000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint32(0x00000000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (exchange)
	if pullErr := readBuffer.PullContext("exchange"); pullErr != nil {
		return nil, pullErr
	}
	_exchange, _exchangeErr := CipExchangeParse(readBuffer, uint16(uint16(len)-uint16(uint16(6))))
	if _exchangeErr != nil {
		return nil, errors.Wrap(_exchangeErr, "Error parsing 'exchange' field")
	}
	exchange := CastCipExchange(_exchange)
	if closeErr := readBuffer.CloseContext("exchange"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("CipRRData"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CipRRData{
		Exchange:  CastCipExchange(exchange),
		EipPacket: &EipPacket{},
	}
	_child.EipPacket.Child = _child
	return _child.EipPacket, nil
}

func (m *CipRRData) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipRRData"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint32("reserved", 32, uint32(0x00000000))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint16("reserved", 16, uint16(0x0000))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (exchange)
		if pushErr := writeBuffer.PushContext("exchange"); pushErr != nil {
			return pushErr
		}
		_exchangeErr := m.Exchange.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("exchange"); popErr != nil {
			return popErr
		}
		if _exchangeErr != nil {
			return errors.Wrap(_exchangeErr, "Error serializing 'exchange' field")
		}

		if popErr := writeBuffer.PopContext("CipRRData"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CipRRData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
