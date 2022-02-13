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
type ApduData struct {
	Child IApduDataChild
}

// The corresponding interface
type IApduData interface {
	// ApciType returns ApciType
	ApciType() uint8
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IApduDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IApduData, serializeChildFunction func() error) error
	GetTypeName() string
}

type IApduDataChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *ApduData)
	GetTypeName() string
	IApduData
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewApduData() *ApduData {
	return &ApduData{}
}

func CastApduData(structType interface{}) *ApduData {
	castFunc := func(typ interface{}) *ApduData {
		if casted, ok := typ.(ApduData); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduData); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduData) GetTypeName() string {
	return "ApduData"
}

func (m *ApduData) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduData) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *ApduData) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (apciType)
	lengthInBits += 4

	return lengthInBits
}

func (m *ApduData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduData"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (apciType) (Used as input to a switch field)
	apciType, _apciTypeErr := readBuffer.ReadUint8("apciType", 4)
	if _apciTypeErr != nil {
		return nil, errors.Wrap(_apciTypeErr, "Error parsing 'apciType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ApduData
	var typeSwitchError error
	switch {
	case apciType == 0x0: // ApduDataGroupValueRead
		_parent, typeSwitchError = ApduDataGroupValueReadParse(readBuffer, dataLength)
	case apciType == 0x1: // ApduDataGroupValueResponse
		_parent, typeSwitchError = ApduDataGroupValueResponseParse(readBuffer, dataLength)
	case apciType == 0x2: // ApduDataGroupValueWrite
		_parent, typeSwitchError = ApduDataGroupValueWriteParse(readBuffer, dataLength)
	case apciType == 0x3: // ApduDataIndividualAddressWrite
		_parent, typeSwitchError = ApduDataIndividualAddressWriteParse(readBuffer, dataLength)
	case apciType == 0x4: // ApduDataIndividualAddressRead
		_parent, typeSwitchError = ApduDataIndividualAddressReadParse(readBuffer, dataLength)
	case apciType == 0x5: // ApduDataIndividualAddressResponse
		_parent, typeSwitchError = ApduDataIndividualAddressResponseParse(readBuffer, dataLength)
	case apciType == 0x6: // ApduDataAdcRead
		_parent, typeSwitchError = ApduDataAdcReadParse(readBuffer, dataLength)
	case apciType == 0x7: // ApduDataAdcResponse
		_parent, typeSwitchError = ApduDataAdcResponseParse(readBuffer, dataLength)
	case apciType == 0x8: // ApduDataMemoryRead
		_parent, typeSwitchError = ApduDataMemoryReadParse(readBuffer, dataLength)
	case apciType == 0x9: // ApduDataMemoryResponse
		_parent, typeSwitchError = ApduDataMemoryResponseParse(readBuffer, dataLength)
	case apciType == 0xA: // ApduDataMemoryWrite
		_parent, typeSwitchError = ApduDataMemoryWriteParse(readBuffer, dataLength)
	case apciType == 0xB: // ApduDataUserMessage
		_parent, typeSwitchError = ApduDataUserMessageParse(readBuffer, dataLength)
	case apciType == 0xC: // ApduDataDeviceDescriptorRead
		_parent, typeSwitchError = ApduDataDeviceDescriptorReadParse(readBuffer, dataLength)
	case apciType == 0xD: // ApduDataDeviceDescriptorResponse
		_parent, typeSwitchError = ApduDataDeviceDescriptorResponseParse(readBuffer, dataLength)
	case apciType == 0xE: // ApduDataRestart
		_parent, typeSwitchError = ApduDataRestartParse(readBuffer, dataLength)
	case apciType == 0xF: // ApduDataOther
		_parent, typeSwitchError = ApduDataOtherParse(readBuffer, dataLength)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("ApduData"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ApduData) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *ApduData) SerializeParent(writeBuffer utils.WriteBuffer, child IApduData, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("ApduData"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (apciType) (Used as input to a switch field)
	apciType := uint8(child.ApciType())
	_apciTypeErr := writeBuffer.WriteUint8("apciType", 4, (apciType))

	if _apciTypeErr != nil {
		return errors.Wrap(_apciTypeErr, "Error serializing 'apciType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduData"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ApduData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
