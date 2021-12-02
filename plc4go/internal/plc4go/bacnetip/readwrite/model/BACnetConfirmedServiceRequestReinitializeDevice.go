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
type BACnetConfirmedServiceRequestReinitializeDevice struct {
	*BACnetConfirmedServiceRequest
	ReinitializedStateOfDevice *BACnetComplexTagDeviceState
	Password                   *BACnetComplexTagCharacterString
}

// The corresponding interface
type IBACnetConfirmedServiceRequestReinitializeDevice interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestReinitializeDevice) ServiceChoice() uint8 {
	return 0x14
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func NewBACnetConfirmedServiceRequestReinitializeDevice(reinitializedStateOfDevice *BACnetComplexTagDeviceState, password *BACnetComplexTagCharacterString) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestReinitializeDevice{
		ReinitializedStateOfDevice:    reinitializedStateOfDevice,
		Password:                      password,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(),
	}
	child.Child = child
	return child.BACnetConfirmedServiceRequest
}

func CastBACnetConfirmedServiceRequestReinitializeDevice(structType interface{}) *BACnetConfirmedServiceRequestReinitializeDevice {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestReinitializeDevice {
		if casted, ok := typ.(BACnetConfirmedServiceRequestReinitializeDevice); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestReinitializeDevice); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestReinitializeDevice(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestReinitializeDevice(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReinitializeDevice"
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (reinitializedStateOfDevice)
	lengthInBits += m.ReinitializedStateOfDevice.LengthInBits()

	// Optional Field (password)
	if m.Password != nil {
		lengthInBits += (*m.Password).LengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestReinitializeDevice) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestReinitializeDeviceParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReinitializeDevice"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (reinitializedStateOfDevice)
	if pullErr := readBuffer.PullContext("reinitializedStateOfDevice"); pullErr != nil {
		return nil, pullErr
	}
	_reinitializedStateOfDevice, _reinitializedStateOfDeviceErr := BACnetComplexTagParse(readBuffer, uint8(0), BACnetDataType_BACNET_DEVICE_STATE)
	if _reinitializedStateOfDeviceErr != nil {
		return nil, errors.Wrap(_reinitializedStateOfDeviceErr, "Error parsing 'reinitializedStateOfDevice' field")
	}
	reinitializedStateOfDevice := CastBACnetComplexTagDeviceState(_reinitializedStateOfDevice)
	if closeErr := readBuffer.CloseContext("reinitializedStateOfDevice"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (password) (Can be skipped, if a given expression evaluates to false)
	var password *BACnetComplexTagCharacterString = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("password"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetComplexTagParse(readBuffer, uint8(1), BACnetDataType_CHARACTER_STRING)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'password' field")
		case _err == utils.ParseAssertError:
			readBuffer.SetPos(currentPos)
		default:
			password = CastBACnetComplexTagCharacterString(_val)
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
		ReinitializedStateOfDevice:    CastBACnetComplexTagDeviceState(reinitializedStateOfDevice),
		Password:                      CastBACnetComplexTagCharacterString(password),
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
		var password *BACnetComplexTagCharacterString = nil
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
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
