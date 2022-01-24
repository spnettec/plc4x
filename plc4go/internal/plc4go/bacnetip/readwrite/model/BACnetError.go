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
type BACnetError struct {
	ErrorClass *BACnetApplicationTagEnumerated
	ErrorCode  *BACnetApplicationTagEnumerated
	Child      IBACnetErrorChild
}

// The corresponding interface
type IBACnetError interface {
	ServiceChoice() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetErrorParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetError, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetErrorChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetError, errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated)
	GetTypeName() string
	IBACnetError
}

func NewBACnetError(errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) *BACnetError {
	return &BACnetError{ErrorClass: errorClass, ErrorCode: errorCode}
}

func CastBACnetError(structType interface{}) *BACnetError {
	castFunc := func(typ interface{}) *BACnetError {
		if casted, ok := typ.(BACnetError); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetError); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetError) GetTypeName() string {
	return "BACnetError"
}

func (m *BACnetError) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetError) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetError) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	// Simple field (errorClass)
	lengthInBits += m.ErrorClass.LengthInBits()

	// Simple field (errorCode)
	lengthInBits += m.ErrorCode.LengthInBits()

	return lengthInBits
}

func (m *BACnetError) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetErrorParse(readBuffer utils.ReadBuffer) (*BACnetError, error) {
	if pullErr := readBuffer.PullContext("BACnetError"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := readBuffer.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetError
	var typeSwitchError error
	switch {
	case serviceChoice == 0x00: // BACnetErrorAcknowledgeAlarm
		_parent, typeSwitchError = BACnetErrorAcknowledgeAlarmParse(readBuffer)
	case serviceChoice == 0x03: // BACnetErrorGetAlarmSummary
		_parent, typeSwitchError = BACnetErrorGetAlarmSummaryParse(readBuffer)
	case serviceChoice == 0x02: // BACnetErrorConfirmedEventNotification
		_parent, typeSwitchError = BACnetErrorConfirmedEventNotificationParse(readBuffer)
	case serviceChoice == 0x04: // BACnetErrorGetEnrollmentSummary
		_parent, typeSwitchError = BACnetErrorGetEnrollmentSummaryParse(readBuffer)
	case serviceChoice == 0x05: // BACnetErrorDeviceCommunicationProtocol
		_parent, typeSwitchError = BACnetErrorDeviceCommunicationProtocolParse(readBuffer)
	case serviceChoice == 0x1D: // BACnetErrorGetEventInformation
		_parent, typeSwitchError = BACnetErrorGetEventInformationParse(readBuffer)
	case serviceChoice == 0x06: // BACnetErrorAtomicReadFile
		_parent, typeSwitchError = BACnetErrorAtomicReadFileParse(readBuffer)
	case serviceChoice == 0x07: // BACnetErrorAtomicWriteFile
		_parent, typeSwitchError = BACnetErrorAtomicWriteFileParse(readBuffer)
	case serviceChoice == 0x0A: // BACnetErrorCreateObject
		_parent, typeSwitchError = BACnetErrorCreateObjectParse(readBuffer)
	case serviceChoice == 0x0C: // BACnetErrorReadProperty
		_parent, typeSwitchError = BACnetErrorReadPropertyParse(readBuffer)
	case serviceChoice == 0x0E: // BACnetErrorReadPropertyMultiple
		_parent, typeSwitchError = BACnetErrorReadPropertyMultipleParse(readBuffer)
	case serviceChoice == 0x0F: // BACnetErrorWriteProperty
		_parent, typeSwitchError = BACnetErrorWritePropertyParse(readBuffer)
	case serviceChoice == 0x1A: // BACnetErrorReadRange
		_parent, typeSwitchError = BACnetErrorReadRangeParse(readBuffer)
	case serviceChoice == 0x11: // BACnetErrorDeviceCommunicationProtocol
		_parent, typeSwitchError = BACnetErrorDeviceCommunicationProtocolParse(readBuffer)
	case serviceChoice == 0x12: // BACnetErrorConfirmedPrivateTransfer
		_parent, typeSwitchError = BACnetErrorConfirmedPrivateTransferParse(readBuffer)
	case serviceChoice == 0x14: // BACnetErrorPasswordFailure
		_parent, typeSwitchError = BACnetErrorPasswordFailureParse(readBuffer)
	case serviceChoice == 0x15: // BACnetErrorVTOpen
		_parent, typeSwitchError = BACnetErrorVTOpenParse(readBuffer)
	case serviceChoice == 0x17: // BACnetErrorVTData
		_parent, typeSwitchError = BACnetErrorVTDataParse(readBuffer)
	case serviceChoice == 0x18: // BACnetErrorRemovedAuthenticate
		_parent, typeSwitchError = BACnetErrorRemovedAuthenticateParse(readBuffer)
	case serviceChoice == 0x0D: // BACnetErrorRemovedReadPropertyConditional
		_parent, typeSwitchError = BACnetErrorRemovedReadPropertyConditionalParse(readBuffer)
	case true: // BACnetErrorUnknown
		_parent, typeSwitchError = BACnetErrorUnknownParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (errorClass)
	if pullErr := readBuffer.PullContext("errorClass"); pullErr != nil {
		return nil, pullErr
	}
	_errorClass, _errorClassErr := BACnetApplicationTagParse(readBuffer)
	if _errorClassErr != nil {
		return nil, errors.Wrap(_errorClassErr, "Error parsing 'errorClass' field")
	}
	errorClass := CastBACnetApplicationTagEnumerated(_errorClass)
	if closeErr := readBuffer.CloseContext("errorClass"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (errorCode)
	if pullErr := readBuffer.PullContext("errorCode"); pullErr != nil {
		return nil, pullErr
	}
	_errorCode, _errorCodeErr := BACnetApplicationTagParse(readBuffer)
	if _errorCodeErr != nil {
		return nil, errors.Wrap(_errorCodeErr, "Error parsing 'errorCode' field")
	}
	errorCode := CastBACnetApplicationTagEnumerated(_errorCode)
	if closeErr := readBuffer.CloseContext("errorCode"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetError"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, errorClass, errorCode)
	return _parent, nil
}

func (m *BACnetError) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetError) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetError, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetError"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := uint8(child.ServiceChoice())
	_serviceChoiceErr := writeBuffer.WriteUint8("serviceChoice", 8, (serviceChoice))

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (errorClass)
	if pushErr := writeBuffer.PushContext("errorClass"); pushErr != nil {
		return pushErr
	}
	_errorClassErr := m.ErrorClass.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("errorClass"); popErr != nil {
		return popErr
	}
	if _errorClassErr != nil {
		return errors.Wrap(_errorClassErr, "Error serializing 'errorClass' field")
	}

	// Simple Field (errorCode)
	if pushErr := writeBuffer.PushContext("errorCode"); pushErr != nil {
		return pushErr
	}
	_errorCodeErr := m.ErrorCode.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("errorCode"); popErr != nil {
		return popErr
	}
	if _errorCodeErr != nil {
		return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
	}

	if popErr := writeBuffer.PopContext("BACnetError"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetError) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
