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
type BACnetConfirmedServiceRequest struct {

	// Arguments.
	Len   uint16
	Child IBACnetConfirmedServiceRequestChild
}

// The corresponding interface
type IBACnetConfirmedServiceRequest interface {
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetConfirmedServiceRequestParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetConfirmedServiceRequest, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetConfirmedServiceRequestChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetConfirmedServiceRequest)
	GetTypeName() string
	IBACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequest factory function for BACnetConfirmedServiceRequest
func NewBACnetConfirmedServiceRequest(len uint16) *BACnetConfirmedServiceRequest {
	return &BACnetConfirmedServiceRequest{Len: len}
}

func CastBACnetConfirmedServiceRequest(structType interface{}) *BACnetConfirmedServiceRequest {
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return casted
	}
	return nil
}

func (m *BACnetConfirmedServiceRequest) GetTypeName() string {
	return "BACnetConfirmedServiceRequest"
}

func (m *BACnetConfirmedServiceRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetConfirmedServiceRequest) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := readBuffer.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetConfirmedServiceRequest
	var typeSwitchError error
	switch {
	case serviceChoice == 0x00: // BACnetConfirmedServiceRequestAcknowledgeAlarm
		_parent, typeSwitchError = BACnetConfirmedServiceRequestAcknowledgeAlarmParse(readBuffer, len)
	case serviceChoice == 0x01: // BACnetConfirmedServiceRequestConfirmedCOVNotification
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedCOVNotificationParse(readBuffer, len)
	case serviceChoice == 0x02: // BACnetConfirmedServiceRequestConfirmedEventNotification
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedEventNotificationParse(readBuffer, len)
	case serviceChoice == 0x04: // BACnetConfirmedServiceRequestGetEnrollmentSummary
		_parent, typeSwitchError = BACnetConfirmedServiceRequestGetEnrollmentSummaryParse(readBuffer, len)
	case serviceChoice == 0x05: // BACnetConfirmedServiceRequestSubscribeCOV
		_parent, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVParse(readBuffer, len)
	case serviceChoice == 0x06: // BACnetConfirmedServiceRequestAtomicReadFile
		_parent, typeSwitchError = BACnetConfirmedServiceRequestAtomicReadFileParse(readBuffer, len)
	case serviceChoice == 0x07: // BACnetConfirmedServiceRequestAtomicWriteFile
		_parent, typeSwitchError = BACnetConfirmedServiceRequestAtomicWriteFileParse(readBuffer, len)
	case serviceChoice == 0x08: // BACnetConfirmedServiceRequestAddListElement
		_parent, typeSwitchError = BACnetConfirmedServiceRequestAddListElementParse(readBuffer, len)
	case serviceChoice == 0x09: // BACnetConfirmedServiceRequestRemoveListElement
		_parent, typeSwitchError = BACnetConfirmedServiceRequestRemoveListElementParse(readBuffer, len)
	case serviceChoice == 0x0A: // BACnetConfirmedServiceRequestCreateObject
		_parent, typeSwitchError = BACnetConfirmedServiceRequestCreateObjectParse(readBuffer, len)
	case serviceChoice == 0x0B: // BACnetConfirmedServiceRequestDeleteObject
		_parent, typeSwitchError = BACnetConfirmedServiceRequestDeleteObjectParse(readBuffer, len)
	case serviceChoice == 0x0C: // BACnetConfirmedServiceRequestReadProperty
		_parent, typeSwitchError = BACnetConfirmedServiceRequestReadPropertyParse(readBuffer, len)
	case serviceChoice == 0x0E: // BACnetConfirmedServiceRequestReadPropertyMultiple
		_parent, typeSwitchError = BACnetConfirmedServiceRequestReadPropertyMultipleParse(readBuffer, len)
	case serviceChoice == 0x0F: // BACnetConfirmedServiceRequestWriteProperty
		_parent, typeSwitchError = BACnetConfirmedServiceRequestWritePropertyParse(readBuffer, len)
	case serviceChoice == 0x10: // BACnetConfirmedServiceRequestWritePropertyMultiple
		_parent, typeSwitchError = BACnetConfirmedServiceRequestWritePropertyMultipleParse(readBuffer, len)
	case serviceChoice == 0x11: // BACnetConfirmedServiceRequestDeviceCommunicationControl
		_parent, typeSwitchError = BACnetConfirmedServiceRequestDeviceCommunicationControlParse(readBuffer, len)
	case serviceChoice == 0x12: // BACnetConfirmedServiceRequestConfirmedPrivateTransfer
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedPrivateTransferParse(readBuffer, len)
	case serviceChoice == 0x13: // BACnetConfirmedServiceRequestConfirmedTextMessage
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedTextMessageParse(readBuffer, len)
	case serviceChoice == 0x14: // BACnetConfirmedServiceRequestReinitializeDevice
		_parent, typeSwitchError = BACnetConfirmedServiceRequestReinitializeDeviceParse(readBuffer, len)
	case serviceChoice == 0x15: // BACnetConfirmedServiceRequestVTOpen
		_parent, typeSwitchError = BACnetConfirmedServiceRequestVTOpenParse(readBuffer, len)
	case serviceChoice == 0x16: // BACnetConfirmedServiceRequestVTClose
		_parent, typeSwitchError = BACnetConfirmedServiceRequestVTCloseParse(readBuffer, len)
	case serviceChoice == 0x17: // BACnetConfirmedServiceRequestVTData
		_parent, typeSwitchError = BACnetConfirmedServiceRequestVTDataParse(readBuffer, len)
	case serviceChoice == 0x18: // BACnetConfirmedServiceRequestRemovedAuthenticate
		_parent, typeSwitchError = BACnetConfirmedServiceRequestRemovedAuthenticateParse(readBuffer, len)
	case serviceChoice == 0x19: // BACnetConfirmedServiceRequestRemovedRequestKey
		_parent, typeSwitchError = BACnetConfirmedServiceRequestRemovedRequestKeyParse(readBuffer, len)
	case serviceChoice == 0x0D: // BACnetConfirmedServiceRequestRemovedReadPropertyConditional
		_parent, typeSwitchError = BACnetConfirmedServiceRequestRemovedReadPropertyConditionalParse(readBuffer, len)
	case serviceChoice == 0x1A: // BACnetConfirmedServiceRequestReadRange
		_parent, typeSwitchError = BACnetConfirmedServiceRequestReadRangeParse(readBuffer, len)
	case serviceChoice == 0x1B: // BACnetConfirmedServiceRequestLifeSafetyOperation
		_parent, typeSwitchError = BACnetConfirmedServiceRequestLifeSafetyOperationParse(readBuffer, len)
	case serviceChoice == 0x1C: // BACnetConfirmedServiceRequestSubscribeCOVProperty
		_parent, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVPropertyParse(readBuffer, len)
	case serviceChoice == 0x1D: // BACnetConfirmedServiceRequestGetEventInformation
		_parent, typeSwitchError = BACnetConfirmedServiceRequestGetEventInformationParse(readBuffer, len)
	case serviceChoice == 0x1E: // BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
		_parent, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleParse(readBuffer, len)
	case serviceChoice == 0x1F: // BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleParse(readBuffer, len)
	case true: // BACnetConfirmedServiceRequestConfirmedUnknown
		_parent, typeSwitchError = BACnetConfirmedServiceRequestConfirmedUnknownParse(readBuffer, len)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *BACnetConfirmedServiceRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetConfirmedServiceRequest) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetConfirmedServiceRequest, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequest"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := uint8(child.GetServiceChoice())
	_serviceChoiceErr := writeBuffer.WriteUint8("serviceChoice", 8, (serviceChoice))

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequest"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConfirmedServiceRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
