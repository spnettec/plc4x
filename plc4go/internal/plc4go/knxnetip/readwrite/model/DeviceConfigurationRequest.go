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
type DeviceConfigurationRequest struct {
	*KnxNetIpMessage
	DeviceConfigurationRequestDataBlock *DeviceConfigurationRequestDataBlock
	Cemi                                *CEMI
}

// The corresponding interface
type IDeviceConfigurationRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *DeviceConfigurationRequest) MsgType() uint16 {
	return 0x0310
}

func (m *DeviceConfigurationRequest) InitializeParent(parent *KnxNetIpMessage) {}

func NewDeviceConfigurationRequest(deviceConfigurationRequestDataBlock *DeviceConfigurationRequestDataBlock, cemi *CEMI) *KnxNetIpMessage {
	child := &DeviceConfigurationRequest{
		DeviceConfigurationRequestDataBlock: deviceConfigurationRequestDataBlock,
		Cemi:                                cemi,
		KnxNetIpMessage:                     NewKnxNetIpMessage(),
	}
	child.Child = child
	return child.KnxNetIpMessage
}

func CastDeviceConfigurationRequest(structType interface{}) *DeviceConfigurationRequest {
	castFunc := func(typ interface{}) *DeviceConfigurationRequest {
		if casted, ok := typ.(DeviceConfigurationRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*DeviceConfigurationRequest); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastDeviceConfigurationRequest(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastDeviceConfigurationRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DeviceConfigurationRequest) GetTypeName() string {
	return "DeviceConfigurationRequest"
}

func (m *DeviceConfigurationRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DeviceConfigurationRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (deviceConfigurationRequestDataBlock)
	lengthInBits += m.DeviceConfigurationRequestDataBlock.LengthInBits()

	// Simple field (cemi)
	lengthInBits += m.Cemi.LengthInBits()

	return lengthInBits
}

func (m *DeviceConfigurationRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DeviceConfigurationRequestParse(readBuffer utils.ReadBuffer, totalLength uint16) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("DeviceConfigurationRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (deviceConfigurationRequestDataBlock)
	if pullErr := readBuffer.PullContext("deviceConfigurationRequestDataBlock"); pullErr != nil {
		return nil, pullErr
	}
	_deviceConfigurationRequestDataBlock, _deviceConfigurationRequestDataBlockErr := DeviceConfigurationRequestDataBlockParse(readBuffer)
	if _deviceConfigurationRequestDataBlockErr != nil {
		return nil, errors.Wrap(_deviceConfigurationRequestDataBlockErr, "Error parsing 'deviceConfigurationRequestDataBlock' field")
	}
	deviceConfigurationRequestDataBlock := CastDeviceConfigurationRequestDataBlock(_deviceConfigurationRequestDataBlock)
	if closeErr := readBuffer.CloseContext("deviceConfigurationRequestDataBlock"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (cemi)
	if pullErr := readBuffer.PullContext("cemi"); pullErr != nil {
		return nil, pullErr
	}
	_cemi, _cemiErr := CEMIParse(readBuffer, uint16(totalLength)-uint16(uint16(uint16(uint16(6))+uint16(deviceConfigurationRequestDataBlock.LengthInBytes()))))
	if _cemiErr != nil {
		return nil, errors.Wrap(_cemiErr, "Error parsing 'cemi' field")
	}
	cemi := CastCEMI(_cemi)
	if closeErr := readBuffer.CloseContext("cemi"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("DeviceConfigurationRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DeviceConfigurationRequest{
		DeviceConfigurationRequestDataBlock: CastDeviceConfigurationRequestDataBlock(deviceConfigurationRequestDataBlock),
		Cemi:                                CastCEMI(cemi),
		KnxNetIpMessage:                     &KnxNetIpMessage{},
	}
	_child.KnxNetIpMessage.Child = _child
	return _child.KnxNetIpMessage, nil
}

func (m *DeviceConfigurationRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeviceConfigurationRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (deviceConfigurationRequestDataBlock)
		if pushErr := writeBuffer.PushContext("deviceConfigurationRequestDataBlock"); pushErr != nil {
			return pushErr
		}
		_deviceConfigurationRequestDataBlockErr := m.DeviceConfigurationRequestDataBlock.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deviceConfigurationRequestDataBlock"); popErr != nil {
			return popErr
		}
		if _deviceConfigurationRequestDataBlockErr != nil {
			return errors.Wrap(_deviceConfigurationRequestDataBlockErr, "Error serializing 'deviceConfigurationRequestDataBlock' field")
		}

		// Simple Field (cemi)
		if pushErr := writeBuffer.PushContext("cemi"); pushErr != nil {
			return pushErr
		}
		_cemiErr := m.Cemi.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("cemi"); popErr != nil {
			return popErr
		}
		if _cemiErr != nil {
			return errors.Wrap(_cemiErr, "Error serializing 'cemi' field")
		}

		if popErr := writeBuffer.PopContext("DeviceConfigurationRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *DeviceConfigurationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
