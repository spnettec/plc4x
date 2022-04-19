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

// ModbusPDUWriteSingleCoilRequest is the data-structure of this message
type ModbusPDUWriteSingleCoilRequest struct {
	*ModbusPDU
	Address uint16
	Value   uint16
}

// IModbusPDUWriteSingleCoilRequest is the corresponding interface of ModbusPDUWriteSingleCoilRequest
type IModbusPDUWriteSingleCoilRequest interface {
	IModbusPDU
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetValue returns Value (property field)
	GetValue() uint16
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

func (m *ModbusPDUWriteSingleCoilRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUWriteSingleCoilRequest) GetFunctionFlag() uint8 {
	return 0x05
}

func (m *ModbusPDUWriteSingleCoilRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUWriteSingleCoilRequest) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUWriteSingleCoilRequest) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ModbusPDUWriteSingleCoilRequest) GetAddress() uint16 {
	return m.Address
}

func (m *ModbusPDUWriteSingleCoilRequest) GetValue() uint16 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteSingleCoilRequest factory function for ModbusPDUWriteSingleCoilRequest
func NewModbusPDUWriteSingleCoilRequest(address uint16, value uint16) *ModbusPDUWriteSingleCoilRequest {
	_result := &ModbusPDUWriteSingleCoilRequest{
		Address:   address,
		Value:     value,
		ModbusPDU: NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUWriteSingleCoilRequest(structType interface{}) *ModbusPDUWriteSingleCoilRequest {
	if casted, ok := structType.(ModbusPDUWriteSingleCoilRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUWriteSingleCoilRequest); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUWriteSingleCoilRequest(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUWriteSingleCoilRequest(casted.Child)
	}
	return nil
}

func (m *ModbusPDUWriteSingleCoilRequest) GetTypeName() string {
	return "ModbusPDUWriteSingleCoilRequest"
}

func (m *ModbusPDUWriteSingleCoilRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUWriteSingleCoilRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (address)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUWriteSingleCoilRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUWriteSingleCoilRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUWriteSingleCoilRequest, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUWriteSingleCoilRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}
	address := _address

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadUint16("value", 16)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteSingleCoilRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteSingleCoilRequest{
		Address:   address,
		Value:     value,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUWriteSingleCoilRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteSingleCoilRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (address)
		address := uint16(m.Address)
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (value)
		value := uint16(m.Value)
		_valueErr := writeBuffer.WriteUint16("value", 16, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteSingleCoilRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUWriteSingleCoilRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
