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

// DeviceConfigurationAck is the corresponding interface of DeviceConfigurationAck
type DeviceConfigurationAck interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetDeviceConfigurationAckDataBlock returns DeviceConfigurationAckDataBlock (property field)
	GetDeviceConfigurationAckDataBlock() DeviceConfigurationAckDataBlock
}

// DeviceConfigurationAckExactly can be used when we want exactly this type and not a type which fulfills DeviceConfigurationAck.
// This is useful for switch cases.
type DeviceConfigurationAckExactly interface {
	DeviceConfigurationAck
	isDeviceConfigurationAck() bool
}

// _DeviceConfigurationAck is the data-structure of this message
type _DeviceConfigurationAck struct {
	*_KnxNetIpMessage
	DeviceConfigurationAckDataBlock DeviceConfigurationAckDataBlock
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeviceConfigurationAck) GetMsgType() uint16 {
	return 0x0311
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeviceConfigurationAck) InitializeParent(parent KnxNetIpMessage) {}

func (m *_DeviceConfigurationAck) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationAck) GetDeviceConfigurationAckDataBlock() DeviceConfigurationAckDataBlock {
	return m.DeviceConfigurationAckDataBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeviceConfigurationAck factory function for _DeviceConfigurationAck
func NewDeviceConfigurationAck(deviceConfigurationAckDataBlock DeviceConfigurationAckDataBlock) *_DeviceConfigurationAck {
	_result := &_DeviceConfigurationAck{
		DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
		_KnxNetIpMessage:                NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationAck(structType interface{}) DeviceConfigurationAck {
	if casted, ok := structType.(DeviceConfigurationAck); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationAck); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationAck) GetTypeName() string {
	return "DeviceConfigurationAck"
}

func (m *_DeviceConfigurationAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_DeviceConfigurationAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (deviceConfigurationAckDataBlock)
	lengthInBits += m.DeviceConfigurationAckDataBlock.GetLengthInBits()

	return lengthInBits
}

func (m *_DeviceConfigurationAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DeviceConfigurationAckParse(readBuffer utils.ReadBuffer) (DeviceConfigurationAck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceConfigurationAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deviceConfigurationAckDataBlock)
	if pullErr := readBuffer.PullContext("deviceConfigurationAckDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deviceConfigurationAckDataBlock")
	}
	_deviceConfigurationAckDataBlock, _deviceConfigurationAckDataBlockErr := DeviceConfigurationAckDataBlockParse(readBuffer)
	if _deviceConfigurationAckDataBlockErr != nil {
		return nil, errors.Wrap(_deviceConfigurationAckDataBlockErr, "Error parsing 'deviceConfigurationAckDataBlock' field of DeviceConfigurationAck")
	}
	deviceConfigurationAckDataBlock := _deviceConfigurationAckDataBlock.(DeviceConfigurationAckDataBlock)
	if closeErr := readBuffer.CloseContext("deviceConfigurationAckDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deviceConfigurationAckDataBlock")
	}

	if closeErr := readBuffer.CloseContext("DeviceConfigurationAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationAck")
	}

	// Create a partially initialized instance
	_child := &_DeviceConfigurationAck{
		_KnxNetIpMessage:                &_KnxNetIpMessage{},
		DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_DeviceConfigurationAck) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeviceConfigurationAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationAck")
		}

		// Simple Field (deviceConfigurationAckDataBlock)
		if pushErr := writeBuffer.PushContext("deviceConfigurationAckDataBlock"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceConfigurationAckDataBlock")
		}
		_deviceConfigurationAckDataBlockErr := writeBuffer.WriteSerializable(m.GetDeviceConfigurationAckDataBlock())
		if popErr := writeBuffer.PopContext("deviceConfigurationAckDataBlock"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceConfigurationAckDataBlock")
		}
		if _deviceConfigurationAckDataBlockErr != nil {
			return errors.Wrap(_deviceConfigurationAckDataBlockErr, "Error serializing 'deviceConfigurationAckDataBlock' field")
		}

		if popErr := writeBuffer.PopContext("DeviceConfigurationAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeviceConfigurationAck")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_DeviceConfigurationAck) isDeviceConfigurationAck() bool {
	return true
}

func (m *_DeviceConfigurationAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
