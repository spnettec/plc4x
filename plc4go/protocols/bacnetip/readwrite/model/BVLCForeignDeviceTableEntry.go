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

// BVLCForeignDeviceTableEntry is the data-structure of this message
type BVLCForeignDeviceTableEntry struct {
	Ip                         []uint8
	Port                       uint16
	Ttl                        uint16
	SecondRemainingBeforePurge uint16
}

// IBVLCForeignDeviceTableEntry is the corresponding interface of BVLCForeignDeviceTableEntry
type IBVLCForeignDeviceTableEntry interface {
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
	// GetTtl returns Ttl (property field)
	GetTtl() uint16
	// GetSecondRemainingBeforePurge returns SecondRemainingBeforePurge (property field)
	GetSecondRemainingBeforePurge() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BVLCForeignDeviceTableEntry) GetIp() []uint8 {
	return m.Ip
}

func (m *BVLCForeignDeviceTableEntry) GetPort() uint16 {
	return m.Port
}

func (m *BVLCForeignDeviceTableEntry) GetTtl() uint16 {
	return m.Ttl
}

func (m *BVLCForeignDeviceTableEntry) GetSecondRemainingBeforePurge() uint16 {
	return m.SecondRemainingBeforePurge
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCForeignDeviceTableEntry factory function for BVLCForeignDeviceTableEntry
func NewBVLCForeignDeviceTableEntry(ip []uint8, port uint16, ttl uint16, secondRemainingBeforePurge uint16) *BVLCForeignDeviceTableEntry {
	return &BVLCForeignDeviceTableEntry{Ip: ip, Port: port, Ttl: ttl, SecondRemainingBeforePurge: secondRemainingBeforePurge}
}

func CastBVLCForeignDeviceTableEntry(structType interface{}) *BVLCForeignDeviceTableEntry {
	if casted, ok := structType.(BVLCForeignDeviceTableEntry); ok {
		return &casted
	}
	if casted, ok := structType.(*BVLCForeignDeviceTableEntry); ok {
		return casted
	}
	return nil
}

func (m *BVLCForeignDeviceTableEntry) GetTypeName() string {
	return "BVLCForeignDeviceTableEntry"
}

func (m *BVLCForeignDeviceTableEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BVLCForeignDeviceTableEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Simple field (ttl)
	lengthInBits += 16

	// Simple field (secondRemainingBeforePurge)
	lengthInBits += 16

	return lengthInBits
}

func (m *BVLCForeignDeviceTableEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCForeignDeviceTableEntryParse(readBuffer utils.ReadBuffer) (*BVLCForeignDeviceTableEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCForeignDeviceTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCForeignDeviceTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (ip)
	if pullErr := readBuffer.PullContext("ip", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ip")
	}
	// Count array
	ip := make([]uint8, uint16(4))
	{
		for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'ip' field")
			}
			ip[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("ip", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ip")
	}

	// Simple Field (port)
	_port, _portErr := readBuffer.ReadUint16("port", 16)
	if _portErr != nil {
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field")
	}
	port := _port

	// Simple Field (ttl)
	_ttl, _ttlErr := readBuffer.ReadUint16("ttl", 16)
	if _ttlErr != nil {
		return nil, errors.Wrap(_ttlErr, "Error parsing 'ttl' field")
	}
	ttl := _ttl

	// Simple Field (secondRemainingBeforePurge)
	_secondRemainingBeforePurge, _secondRemainingBeforePurgeErr := readBuffer.ReadUint16("secondRemainingBeforePurge", 16)
	if _secondRemainingBeforePurgeErr != nil {
		return nil, errors.Wrap(_secondRemainingBeforePurgeErr, "Error parsing 'secondRemainingBeforePurge' field")
	}
	secondRemainingBeforePurge := _secondRemainingBeforePurge

	if closeErr := readBuffer.CloseContext("BVLCForeignDeviceTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCForeignDeviceTableEntry")
	}

	// Create the instance
	return NewBVLCForeignDeviceTableEntry(ip, port, ttl, secondRemainingBeforePurge), nil
}

func (m *BVLCForeignDeviceTableEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BVLCForeignDeviceTableEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BVLCForeignDeviceTableEntry")
	}

	// Array Field (ip)
	if m.Ip != nil {
		if pushErr := writeBuffer.PushContext("ip", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ip")
		}
		for _, _element := range m.Ip {
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'ip' field")
			}
		}
		if popErr := writeBuffer.PopContext("ip", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ip")
		}
	}

	// Simple Field (port)
	port := uint16(m.Port)
	_portErr := writeBuffer.WriteUint16("port", 16, (port))
	if _portErr != nil {
		return errors.Wrap(_portErr, "Error serializing 'port' field")
	}

	// Simple Field (ttl)
	ttl := uint16(m.Ttl)
	_ttlErr := writeBuffer.WriteUint16("ttl", 16, (ttl))
	if _ttlErr != nil {
		return errors.Wrap(_ttlErr, "Error serializing 'ttl' field")
	}

	// Simple Field (secondRemainingBeforePurge)
	secondRemainingBeforePurge := uint16(m.SecondRemainingBeforePurge)
	_secondRemainingBeforePurgeErr := writeBuffer.WriteUint16("secondRemainingBeforePurge", 16, (secondRemainingBeforePurge))
	if _secondRemainingBeforePurgeErr != nil {
		return errors.Wrap(_secondRemainingBeforePurgeErr, "Error serializing 'secondRemainingBeforePurge' field")
	}

	if popErr := writeBuffer.PopContext("BVLCForeignDeviceTableEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BVLCForeignDeviceTableEntry")
	}
	return nil
}

func (m *BVLCForeignDeviceTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
