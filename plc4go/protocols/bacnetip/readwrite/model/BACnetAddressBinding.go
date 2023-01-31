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


// BACnetAddressBinding is the corresponding interface of BACnetAddressBinding
type BACnetAddressBinding interface {
	utils.LengthAware
	utils.Serializable
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier
	// GetDeviceAddress returns DeviceAddress (property field)
	GetDeviceAddress() BACnetAddress
}

// BACnetAddressBindingExactly can be used when we want exactly this type and not a type which fulfills BACnetAddressBinding.
// This is useful for switch cases.
type BACnetAddressBindingExactly interface {
	BACnetAddressBinding
	isBACnetAddressBinding() bool
}

// _BACnetAddressBinding is the data-structure of this message
type _BACnetAddressBinding struct {
        DeviceIdentifier BACnetApplicationTagObjectIdentifier
        DeviceAddress BACnetAddress
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAddressBinding) GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetAddressBinding) GetDeviceAddress() BACnetAddress {
	return m.DeviceAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetAddressBinding factory function for _BACnetAddressBinding
func NewBACnetAddressBinding( deviceIdentifier BACnetApplicationTagObjectIdentifier , deviceAddress BACnetAddress ) *_BACnetAddressBinding {
return &_BACnetAddressBinding{ DeviceIdentifier: deviceIdentifier , DeviceAddress: deviceAddress }
}

// Deprecated: use the interface for direct cast
func CastBACnetAddressBinding(structType interface{}) BACnetAddressBinding {
    if casted, ok := structType.(BACnetAddressBinding); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAddressBinding); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAddressBinding) GetTypeName() string {
	return "BACnetAddressBinding"
}

func (m *_BACnetAddressBinding) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetAddressBinding) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.GetLengthInBits()

	// Simple field (deviceAddress)
	lengthInBits += m.DeviceAddress.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetAddressBinding) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAddressBindingParse(theBytes []byte) (BACnetAddressBinding, error) {
	return BACnetAddressBindingParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetAddressBindingParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetAddressBinding, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAddressBinding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAddressBinding")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deviceIdentifier)
	if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deviceIdentifier")
	}
_deviceIdentifier, _deviceIdentifierErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _deviceIdentifierErr != nil {
		return nil, errors.Wrap(_deviceIdentifierErr, "Error parsing 'deviceIdentifier' field of BACnetAddressBinding")
	}
	deviceIdentifier := _deviceIdentifier.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deviceIdentifier")
	}

	// Simple Field (deviceAddress)
	if pullErr := readBuffer.PullContext("deviceAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deviceAddress")
	}
_deviceAddress, _deviceAddressErr := BACnetAddressParseWithBuffer(readBuffer)
	if _deviceAddressErr != nil {
		return nil, errors.Wrap(_deviceAddressErr, "Error parsing 'deviceAddress' field of BACnetAddressBinding")
	}
	deviceAddress := _deviceAddress.(BACnetAddress)
	if closeErr := readBuffer.CloseContext("deviceAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deviceAddress")
	}

	if closeErr := readBuffer.CloseContext("BACnetAddressBinding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAddressBinding")
	}

	// Create the instance
	return &_BACnetAddressBinding{
			DeviceIdentifier: deviceIdentifier,
			DeviceAddress: deviceAddress,
		}, nil
}

func (m *_BACnetAddressBinding) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAddressBinding) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetAddressBinding"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAddressBinding")
	}

	// Simple Field (deviceIdentifier)
	if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for deviceIdentifier")
	}
	_deviceIdentifierErr := writeBuffer.WriteSerializable(m.GetDeviceIdentifier())
	if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for deviceIdentifier")
	}
	if _deviceIdentifierErr != nil {
		return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
	}

	// Simple Field (deviceAddress)
	if pushErr := writeBuffer.PushContext("deviceAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for deviceAddress")
	}
	_deviceAddressErr := writeBuffer.WriteSerializable(m.GetDeviceAddress())
	if popErr := writeBuffer.PopContext("deviceAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for deviceAddress")
	}
	if _deviceAddressErr != nil {
		return errors.Wrap(_deviceAddressErr, "Error serializing 'deviceAddress' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAddressBinding"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAddressBinding")
	}
	return nil
}


func (m *_BACnetAddressBinding) isBACnetAddressBinding() bool {
	return true
}

func (m *_BACnetAddressBinding) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



