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

// BACnetConfirmedServiceRequestDeviceCommunicationControl is the data-structure of this message
type BACnetConfirmedServiceRequestDeviceCommunicationControl struct {
	*BACnetConfirmedServiceRequest
	TimeDuration  *BACnetContextTagUnsignedInteger
	EnableDisable *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable
	Password      *BACnetContextTagCharacterString

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestDeviceCommunicationControl is the corresponding interface of BACnetConfirmedServiceRequestDeviceCommunicationControl
type IBACnetConfirmedServiceRequestDeviceCommunicationControl interface {
	IBACnetConfirmedServiceRequest
	// GetTimeDuration returns TimeDuration (property field)
	GetTimeDuration() *BACnetContextTagUnsignedInteger
	// GetEnableDisable returns EnableDisable (property field)
	GetEnableDisable() *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable
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
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) GetTimeDuration() *BACnetContextTagUnsignedInteger {
	return m.TimeDuration
}

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) GetEnableDisable() *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable {
	return m.EnableDisable
}

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) GetPassword() *BACnetContextTagCharacterString {
	return m.Password
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestDeviceCommunicationControl factory function for BACnetConfirmedServiceRequestDeviceCommunicationControl
func NewBACnetConfirmedServiceRequestDeviceCommunicationControl(timeDuration *BACnetContextTagUnsignedInteger, enableDisable *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable, password *BACnetContextTagCharacterString, serviceRequestLength uint16) *BACnetConfirmedServiceRequestDeviceCommunicationControl {
	_result := &BACnetConfirmedServiceRequestDeviceCommunicationControl{
		TimeDuration:                  timeDuration,
		EnableDisable:                 enableDisable,
		Password:                      password,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestDeviceCommunicationControl(structType interface{}) *BACnetConfirmedServiceRequestDeviceCommunicationControl {
	if casted, ok := structType.(BACnetConfirmedServiceRequestDeviceCommunicationControl); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestDeviceCommunicationControl); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestDeviceCommunicationControl(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestDeviceCommunicationControl(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) GetTypeName() string {
	return "BACnetConfirmedServiceRequestDeviceCommunicationControl"
}

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Optional Field (timeDuration)
	if m.TimeDuration != nil {
		lengthInBits += (*m.TimeDuration).GetLengthInBits()
	}

	// Simple field (enableDisable)
	lengthInBits += m.EnableDisable.GetLengthInBits()

	// Optional Field (password)
	if m.Password != nil {
		lengthInBits += (*m.Password).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestDeviceCommunicationControl, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Optional Field (timeDuration) (Can be skipped, if a given expression evaluates to false)
	var timeDuration *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("timeDuration"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'timeDuration' field")
		default:
			timeDuration = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("timeDuration"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (enableDisable)
	if pullErr := readBuffer.PullContext("enableDisable"); pullErr != nil {
		return nil, pullErr
	}
	_enableDisable, _enableDisableErr := BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableParse(readBuffer, uint8(uint8(1)))
	if _enableDisableErr != nil {
		return nil, errors.Wrap(_enableDisableErr, "Error parsing 'enableDisable' field")
	}
	enableDisable := CastBACnetConfirmedServiceRequestReinitializeDeviceEnableDisable(_enableDisable)
	if closeErr := readBuffer.CloseContext("enableDisable"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (password) (Can be skipped, if a given expression evaluates to false)
	var password *BACnetContextTagCharacterString = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("password"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_CHARACTER_STRING)
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

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestDeviceCommunicationControl{
		TimeDuration:                  CastBACnetContextTagUnsignedInteger(timeDuration),
		EnableDisable:                 CastBACnetConfirmedServiceRequestReinitializeDeviceEnableDisable(enableDisable),
		Password:                      CastBACnetContextTagCharacterString(password),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); pushErr != nil {
			return pushErr
		}

		// Optional Field (timeDuration) (Can be skipped, if the value is null)
		var timeDuration *BACnetContextTagUnsignedInteger = nil
		if m.TimeDuration != nil {
			if pushErr := writeBuffer.PushContext("timeDuration"); pushErr != nil {
				return pushErr
			}
			timeDuration = m.TimeDuration
			_timeDurationErr := timeDuration.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("timeDuration"); popErr != nil {
				return popErr
			}
			if _timeDurationErr != nil {
				return errors.Wrap(_timeDurationErr, "Error serializing 'timeDuration' field")
			}
		}

		// Simple Field (enableDisable)
		if pushErr := writeBuffer.PushContext("enableDisable"); pushErr != nil {
			return pushErr
		}
		_enableDisableErr := m.EnableDisable.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("enableDisable"); popErr != nil {
			return popErr
		}
		if _enableDisableErr != nil {
			return errors.Wrap(_enableDisableErr, "Error serializing 'enableDisable' field")
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

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestDeviceCommunicationControl) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
