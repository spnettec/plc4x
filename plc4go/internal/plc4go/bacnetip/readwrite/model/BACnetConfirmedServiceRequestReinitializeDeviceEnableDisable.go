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
type BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable struct {
	RawData             *BACnetContextTagEnumerated
	IsEnable            bool
	IsDisable           bool
	IsDisableInitiation bool
}

// The corresponding interface
type IBACnetConfirmedServiceRequestReinitializeDeviceEnableDisable interface {
	// GetRawData returns RawData
	GetRawData() *BACnetContextTagEnumerated
	// GetIsEnable returns IsEnable
	GetIsEnable() bool
	// GetIsDisable returns IsDisable
	GetIsDisable() bool
	// GetIsDisableInitiation returns IsDisableInitiation
	GetIsDisableInitiation() bool
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
func (m *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) GetRawData() *BACnetContextTagEnumerated {
	return m.RawData
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) GetIsEnable() bool {
	// TODO: calculation should happen here instead accessing the stored field
	return m.IsEnable
}

func (m *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) GetIsDisable() bool {
	// TODO: calculation should happen here instead accessing the stored field
	return m.IsDisable
}

func (m *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) GetIsDisableInitiation() bool {
	// TODO: calculation should happen here instead accessing the stored field
	return m.IsDisableInitiation
}

func NewBACnetConfirmedServiceRequestReinitializeDeviceEnableDisable(rawData *BACnetContextTagEnumerated, isEnable bool, isDisable bool, isDisableInitiation bool) *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable {
	return &BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable{RawData: rawData, IsEnable: isEnable, IsDisable: isDisable, IsDisableInitiation: isDisableInitiation}
}

func CastBACnetConfirmedServiceRequestReinitializeDeviceEnableDisable(structType interface{}) *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable {
		if casted, ok := typ.(BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable"
}

func (m *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (rawData)
	if m.RawData != nil {
		lengthInBits += (*m.RawData).LengthInBits()
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable"); pullErr != nil {
		return nil, pullErr
	}

	// Optional Field (rawData) (Can be skipped, if a given expression evaluates to false)
	var rawData *BACnetContextTagEnumerated = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("rawData"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, tagNumber, BACnetDataType_ENUMERATED)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'rawData' field")
		default:
			rawData = CastBACnetContextTagEnumerated(_val)
			if closeErr := readBuffer.CloseContext("rawData"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Virtual field
	_isEnable := bool(bool((rawData) != (nil))) && bool(bool(((*rawData).Payload.ActualValue) == (0)))
	isEnable := bool(_isEnable)

	// Virtual field
	_isDisable := bool(bool((rawData) != (nil))) && bool(bool(((*rawData).Payload.ActualValue) == (1)))
	isDisable := bool(_isDisable)

	// Virtual field
	_isDisableInitiation := bool(bool((rawData) != (nil))) && bool(bool(((*rawData).Payload.ActualValue) == (2)))
	isDisableInitiation := bool(_isDisableInitiation)

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetConfirmedServiceRequestReinitializeDeviceEnableDisable(rawData, isEnable, isDisable, isDisableInitiation), nil
}

func (m *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable"); pushErr != nil {
		return pushErr
	}

	// Optional Field (rawData) (Can be skipped, if the value is null)
	var rawData *BACnetContextTagEnumerated = nil
	if m.RawData != nil {
		if pushErr := writeBuffer.PushContext("rawData"); pushErr != nil {
			return pushErr
		}
		rawData = m.RawData
		_rawDataErr := rawData.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("rawData"); popErr != nil {
			return popErr
		}
		if _rawDataErr != nil {
			return errors.Wrap(_rawDataErr, "Error serializing 'rawData' field")
		}
	}
	// Virtual field
	if _isEnableErr := writeBuffer.WriteVirtual("isEnable", m.IsEnable); _isEnableErr != nil {
		return errors.Wrap(_isEnableErr, "Error serializing 'isEnable' field")
	}
	// Virtual field
	if _isDisableErr := writeBuffer.WriteVirtual("isDisable", m.IsDisable); _isDisableErr != nil {
		return errors.Wrap(_isDisableErr, "Error serializing 'isDisable' field")
	}
	// Virtual field
	if _isDisableInitiationErr := writeBuffer.WriteVirtual("isDisableInitiation", m.IsDisableInitiation); _isDisableInitiationErr != nil {
		return errors.Wrap(_isDisableInitiationErr, "Error serializing 'isDisableInitiation' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
