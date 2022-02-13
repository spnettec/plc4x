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
type BACnetStatusFlags struct {
	RawBits      *BACnetContextTagBitString
	InAlarm      bool
	Fault        bool
	Overriden    bool
	OutOfService bool
}

// The corresponding interface
type IBACnetStatusFlags interface {
	// GetRawBits returns RawBits
	GetRawBits() *BACnetContextTagBitString
	// GetInAlarm returns InAlarm
	GetInAlarm() bool
	// GetFault returns Fault
	GetFault() bool
	// GetOverriden returns Overriden
	GetOverriden() bool
	// GetOutOfService returns OutOfService
	GetOutOfService() bool
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetStatusFlags) GetRawBits() *BACnetContextTagBitString {
	return m.RawBits
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetStatusFlags) GetInAlarm() bool {
	// TODO: calculation should happen here instead accessing the stored field
	return m.InAlarm
}

func (m *BACnetStatusFlags) GetFault() bool {
	// TODO: calculation should happen here instead accessing the stored field
	return m.Fault
}

func (m *BACnetStatusFlags) GetOverriden() bool {
	// TODO: calculation should happen here instead accessing the stored field
	return m.Overriden
}

func (m *BACnetStatusFlags) GetOutOfService() bool {
	// TODO: calculation should happen here instead accessing the stored field
	return m.OutOfService
}

func NewBACnetStatusFlags(rawBits *BACnetContextTagBitString, inAlarm bool, fault bool, overriden bool, outOfService bool) *BACnetStatusFlags {
	return &BACnetStatusFlags{RawBits: rawBits, InAlarm: inAlarm, Fault: fault, Overriden: overriden, OutOfService: outOfService}
}

func CastBACnetStatusFlags(structType interface{}) *BACnetStatusFlags {
	castFunc := func(typ interface{}) *BACnetStatusFlags {
		if casted, ok := typ.(BACnetStatusFlags); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetStatusFlags); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetStatusFlags) GetTypeName() string {
	return "BACnetStatusFlags"
}

func (m *BACnetStatusFlags) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetStatusFlags) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (rawBits)
	lengthInBits += m.RawBits.LengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetStatusFlags) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetStatusFlagsParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetStatusFlags, error) {
	if pullErr := readBuffer.PullContext("BACnetStatusFlags"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (rawBits)
	if pullErr := readBuffer.PullContext("rawBits"); pullErr != nil {
		return nil, pullErr
	}
	_rawBits, _rawBitsErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType_BIT_STRING)
	if _rawBitsErr != nil {
		return nil, errors.Wrap(_rawBitsErr, "Error parsing 'rawBits' field")
	}
	rawBits := CastBACnetContextTagBitString(_rawBits)
	if closeErr := readBuffer.CloseContext("rawBits"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_inAlarm := rawBits.Payload.Data[0]
	inAlarm := bool(_inAlarm)

	// Virtual field
	_fault := rawBits.Payload.Data[1]
	fault := bool(_fault)

	// Virtual field
	_overriden := rawBits.Payload.Data[2]
	overriden := bool(_overriden)

	// Virtual field
	_outOfService := rawBits.Payload.Data[3]
	outOfService := bool(_outOfService)

	if closeErr := readBuffer.CloseContext("BACnetStatusFlags"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetStatusFlags(rawBits, inAlarm, fault, overriden, outOfService), nil
}

func (m *BACnetStatusFlags) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetStatusFlags"); pushErr != nil {
		return pushErr
	}

	// Simple Field (rawBits)
	if pushErr := writeBuffer.PushContext("rawBits"); pushErr != nil {
		return pushErr
	}
	_rawBitsErr := m.RawBits.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("rawBits"); popErr != nil {
		return popErr
	}
	if _rawBitsErr != nil {
		return errors.Wrap(_rawBitsErr, "Error serializing 'rawBits' field")
	}
	// Virtual field
	if _inAlarmErr := writeBuffer.WriteVirtual("inAlarm", m.InAlarm); _inAlarmErr != nil {
		return errors.Wrap(_inAlarmErr, "Error serializing 'inAlarm' field")
	}
	// Virtual field
	if _faultErr := writeBuffer.WriteVirtual("fault", m.Fault); _faultErr != nil {
		return errors.Wrap(_faultErr, "Error serializing 'fault' field")
	}
	// Virtual field
	if _overridenErr := writeBuffer.WriteVirtual("overriden", m.Overriden); _overridenErr != nil {
		return errors.Wrap(_overridenErr, "Error serializing 'overriden' field")
	}
	// Virtual field
	if _outOfServiceErr := writeBuffer.WriteVirtual("outOfService", m.OutOfService); _outOfServiceErr != nil {
		return errors.Wrap(_outOfServiceErr, "Error serializing 'outOfService' field")
	}

	if popErr := writeBuffer.PopContext("BACnetStatusFlags"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
