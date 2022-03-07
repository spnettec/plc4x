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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetConfirmedServiceRequestReinitializeDevice struct {
	*BACnetConfirmedServiceRequest
	ReinitializedStateOfDevice *BACnetContextTagDeviceState
	Password                   *BACnetContextTagCharacterString

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetConfirmedServiceRequestReinitializeDevice interface {
	IBACnetConfirmedServiceRequest
	// GetReinitializedStateOfDevice returns ReinitializedStateOfDevice (property field)
	GetReinitializedStateOfDevice() *BACnetContextTagDeviceState
	// GetPassword returns Password (property field)
	GetPassword() *BACnetContextTagCharacterString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestReinitializeDevice) ServiceChoice() uint8 {
	return 0x14
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetServiceChoice() uint8 {
	return 0x14
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetReinitializedStateOfDevice() *BACnetContextTagDeviceState {
	return m.ReinitializedStateOfDevice
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetPassword() *BACnetContextTagCharacterString {
	return m.Password
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReinitializeDevice factory function for BACnetConfirmedServiceRequestReinitializeDevice
func NewBACnetConfirmedServiceRequestReinitializeDevice(reinitializedStateOfDevice *BACnetContextTagDeviceState, password *BACnetContextTagCharacterString, len uint16) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestReinitializeDevice{
		ReinitializedStateOfDevice:    reinitializedStateOfDevice,
		Password:                      password,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetConfirmedServiceRequest
}

func CastBACnetConfirmedServiceRequestReinitializeDevice(structType interface{}) *BACnetConfirmedServiceRequestReinitializeDevice {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReinitializeDevice); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReinitializeDevice); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestReinitializeDevice(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestReinitializeDevice(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReinitializeDevice"
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (reinitializedStateOfDevice)
	lengthInBits += m.ReinitializedStateOfDevice.GetLengthInBits()

	// Optional Field (password)
	if m.Password != nil {
		lengthInBits += (*m.Password).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReinitializeDeviceParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReinitializeDevice"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (reinitializedStateOfDevice)
	if pullErr := readBuffer.PullContext("reinitializedStateOfDevice"); pullErr != nil {
		return nil, pullErr
	}
	_reinitializedStateOfDevice, _reinitializedStateOfDeviceErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_DEVICE_STATE))
	if _reinitializedStateOfDeviceErr != nil {
		return nil, errors.Wrap(_reinitializedStateOfDeviceErr, "Error parsing 'reinitializedStateOfDevice' field")
	}
	reinitializedStateOfDevice := CastBACnetContextTagDeviceState(_reinitializedStateOfDevice)
	if closeErr := readBuffer.CloseContext("reinitializedStateOfDevice"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (password) (Can be skipped, if a given expression evaluates to false)
	var password *BACnetContextTagCharacterString = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("password"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_CHARACTER_STRING)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'password' field")
		default:
			password = CastBACnetContextTagCharacterString(_val)
			if closeErr := readBuffer.CloseContext("password"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReinitializeDevice"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestReinitializeDevice{
		ReinitializedStateOfDevice:    CastBACnetContextTagDeviceState(reinitializedStateOfDevice),
		Password:                      CastBACnetContextTagCharacterString(password),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child.BACnetConfirmedServiceRequest, nil
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReinitializeDevice"); pushErr != nil {
			return pushErr
		}

		// Simple Field (reinitializedStateOfDevice)
		if pushErr := writeBuffer.PushContext("reinitializedStateOfDevice"); pushErr != nil {
			return pushErr
		}
		_reinitializedStateOfDeviceErr := m.ReinitializedStateOfDevice.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("reinitializedStateOfDevice"); popErr != nil {
			return popErr
		}
		if _reinitializedStateOfDeviceErr != nil {
			return errors.Wrap(_reinitializedStateOfDeviceErr, "Error serializing 'reinitializedStateOfDevice' field")
		}

		// Optional Field (password) (Can be skipped, if the value is null)
		var password *BACnetContextTagCharacterString = nil
		if m.Password != nil {
			if pushErr := writeBuffer.PushContext("password"); pushErr != nil {
				return pushErr
			}
			password = m.Password
			_passwordErr := password.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("password"); popErr != nil {
				return popErr
			}
			if _passwordErr != nil {
				return errors.Wrap(_passwordErr, "Error serializing 'password' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReinitializeDevice"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
