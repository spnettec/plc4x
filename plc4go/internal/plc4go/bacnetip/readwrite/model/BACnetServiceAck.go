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
type BACnetServiceAck struct {
	Child IBACnetServiceAckChild
}

// The corresponding interface
type IBACnetServiceAck interface {
	// ServiceChoice returns ServiceChoice
	ServiceChoice() uint8
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetServiceAckParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetServiceAck, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetServiceAckChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetServiceAck)
	GetTypeName() string
	IBACnetServiceAck
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewBACnetServiceAck() *BACnetServiceAck {
	return &BACnetServiceAck{}
}

func CastBACnetServiceAck(structType interface{}) *BACnetServiceAck {
	castFunc := func(typ interface{}) *BACnetServiceAck {
		if casted, ok := typ.(BACnetServiceAck); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAck) GetTypeName() string {
	return "BACnetServiceAck"
}

func (m *BACnetServiceAck) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetServiceAck) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetServiceAck) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetServiceAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckParse(readBuffer utils.ReadBuffer) (*BACnetServiceAck, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAck"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := readBuffer.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetServiceAck
	var typeSwitchError error
	switch {
	case serviceChoice == 0x03: // BACnetServiceAckGetAlarmSummary
		_parent, typeSwitchError = BACnetServiceAckGetAlarmSummaryParse(readBuffer)
	case serviceChoice == 0x04: // BACnetServiceAckGetEnrollmentSummary
		_parent, typeSwitchError = BACnetServiceAckGetEnrollmentSummaryParse(readBuffer)
	case serviceChoice == 0x1D: // BACnetServiceAckGetEventInformation
		_parent, typeSwitchError = BACnetServiceAckGetEventInformationParse(readBuffer)
	case serviceChoice == 0x06: // BACnetServiceAckAtomicReadFile
		_parent, typeSwitchError = BACnetServiceAckAtomicReadFileParse(readBuffer)
	case serviceChoice == 0x07: // BACnetServiceAckAtomicWriteFile
		_parent, typeSwitchError = BACnetServiceAckAtomicWriteFileParse(readBuffer)
	case serviceChoice == 0x0A: // BACnetServiceAckCreateObject
		_parent, typeSwitchError = BACnetServiceAckCreateObjectParse(readBuffer)
	case serviceChoice == 0x0C: // BACnetServiceAckReadProperty
		_parent, typeSwitchError = BACnetServiceAckReadPropertyParse(readBuffer)
	case serviceChoice == 0x0E: // BACnetServiceAckReadPropertyMultiple
		_parent, typeSwitchError = BACnetServiceAckReadPropertyMultipleParse(readBuffer)
	case serviceChoice == 0x1A: // BACnetServiceAckReadRange
		_parent, typeSwitchError = BACnetServiceAckReadRangeParse(readBuffer)
	case serviceChoice == 0x12: // BACnetServiceAckConfirmedPrivateTransfer
		_parent, typeSwitchError = BACnetServiceAckConfirmedPrivateTransferParse(readBuffer)
	case serviceChoice == 0x15: // BACnetServiceAckVTOpen
		_parent, typeSwitchError = BACnetServiceAckVTOpenParse(readBuffer)
	case serviceChoice == 0x17: // BACnetServiceAckVTData
		_parent, typeSwitchError = BACnetServiceAckVTDataParse(readBuffer)
	case serviceChoice == 0x18: // BACnetServiceAckRemovedAuthenticate
		_parent, typeSwitchError = BACnetServiceAckRemovedAuthenticateParse(readBuffer)
	case serviceChoice == 0x0D: // BACnetServiceAckRemovedReadPropertyConditional
		_parent, typeSwitchError = BACnetServiceAckRemovedReadPropertyConditionalParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAck"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *BACnetServiceAck) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetServiceAck) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetServiceAck, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetServiceAck"); pushErr != nil {
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

	if popErr := writeBuffer.PopContext("BACnetServiceAck"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetServiceAck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
