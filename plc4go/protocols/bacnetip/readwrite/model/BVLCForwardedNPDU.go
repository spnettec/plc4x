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

// BVLCForwardedNPDU is the corresponding interface of BVLCForwardedNPDU
type BVLCForwardedNPDU interface {
	utils.LengthAware
	utils.Serializable
	BVLC
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
	// GetNpdu returns Npdu (property field)
	GetNpdu() NPDU
}

// BVLCForwardedNPDUExactly can be used when we want exactly this type and not a type which fulfills BVLCForwardedNPDU.
// This is useful for switch cases.
type BVLCForwardedNPDUExactly interface {
	BVLCForwardedNPDU
	isBVLCForwardedNPDU() bool
}

// _BVLCForwardedNPDU is the data-structure of this message
type _BVLCForwardedNPDU struct {
	*_BVLC
	Ip   []uint8
	Port uint16
	Npdu NPDU

	// Arguments.
	BvlcPayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCForwardedNPDU) GetBvlcFunction() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCForwardedNPDU) InitializeParent(parent BVLC) {}

func (m *_BVLCForwardedNPDU) GetParent() BVLC {
	return m._BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCForwardedNPDU) GetIp() []uint8 {
	return m.Ip
}

func (m *_BVLCForwardedNPDU) GetPort() uint16 {
	return m.Port
}

func (m *_BVLCForwardedNPDU) GetNpdu() NPDU {
	return m.Npdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCForwardedNPDU factory function for _BVLCForwardedNPDU
func NewBVLCForwardedNPDU(ip []uint8, port uint16, npdu NPDU, bvlcPayloadLength uint16) *_BVLCForwardedNPDU {
	_result := &_BVLCForwardedNPDU{
		Ip:    ip,
		Port:  port,
		Npdu:  npdu,
		_BVLC: NewBVLC(),
	}
	_result._BVLC._BVLCChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCForwardedNPDU(structType interface{}) BVLCForwardedNPDU {
	if casted, ok := structType.(BVLCForwardedNPDU); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCForwardedNPDU); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCForwardedNPDU) GetTypeName() string {
	return "BVLCForwardedNPDU"
}

func (m *_BVLCForwardedNPDU) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BVLCForwardedNPDU) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Simple field (npdu)
	lengthInBits += m.Npdu.GetLengthInBits()

	return lengthInBits
}

func (m *_BVLCForwardedNPDU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCForwardedNPDUParse(readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (BVLCForwardedNPDU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCForwardedNPDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCForwardedNPDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (ip)
	if pullErr := readBuffer.PullContext("ip", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ip")
	}
	// Count array
	ip := make([]uint8, uint16(4))
	// This happens when the size is set conditional to 0
	if len(ip) == 0 {
		ip = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'ip' field of BVLCForwardedNPDU")
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
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field of BVLCForwardedNPDU")
	}
	port := _port

	// Simple Field (npdu)
	if pullErr := readBuffer.PullContext("npdu"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for npdu")
	}
	_npdu, _npduErr := NPDUParse(readBuffer, uint16(uint16(bvlcPayloadLength)-uint16(uint16(6))))
	if _npduErr != nil {
		return nil, errors.Wrap(_npduErr, "Error parsing 'npdu' field of BVLCForwardedNPDU")
	}
	npdu := _npdu.(NPDU)
	if closeErr := readBuffer.CloseContext("npdu"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for npdu")
	}

	if closeErr := readBuffer.CloseContext("BVLCForwardedNPDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCForwardedNPDU")
	}

	// Create a partially initialized instance
	_child := &_BVLCForwardedNPDU{
		Ip:    ip,
		Port:  port,
		Npdu:  npdu,
		_BVLC: &_BVLC{},
	}
	_child._BVLC._BVLCChildRequirements = _child
	return _child, nil
}

func (m *_BVLCForwardedNPDU) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCForwardedNPDU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCForwardedNPDU")
		}

		// Array Field (ip)
		if pushErr := writeBuffer.PushContext("ip", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ip")
		}
		for _, _element := range m.GetIp() {
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'ip' field")
			}
		}
		if popErr := writeBuffer.PopContext("ip", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ip")
		}

		// Simple Field (port)
		port := uint16(m.GetPort())
		_portErr := writeBuffer.WriteUint16("port", 16, (port))
		if _portErr != nil {
			return errors.Wrap(_portErr, "Error serializing 'port' field")
		}

		// Simple Field (npdu)
		if pushErr := writeBuffer.PushContext("npdu"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for npdu")
		}
		_npduErr := writeBuffer.WriteSerializable(m.GetNpdu())
		if popErr := writeBuffer.PopContext("npdu"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for npdu")
		}
		if _npduErr != nil {
			return errors.Wrap(_npduErr, "Error serializing 'npdu' field")
		}

		if popErr := writeBuffer.PopContext("BVLCForwardedNPDU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCForwardedNPDU")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BVLCForwardedNPDU) isBVLCForwardedNPDU() bool {
	return true
}

func (m *_BVLCForwardedNPDU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
