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
type ModbusPDUWriteMultipleHoldingRegistersRequest struct {
	*ModbusPDU
	StartingAddress uint16
	Quantity        uint16
	Value           []byte
}

// The corresponding interface
type IModbusPDUWriteMultipleHoldingRegistersRequest interface {
	// GetStartingAddress returns StartingAddress
	GetStartingAddress() uint16
	// GetQuantity returns Quantity
	GetQuantity() uint16
	// GetValue returns Value
	GetValue() []byte
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
func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) FunctionFlag() uint8 {
	return 0x10
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) GetFunctionFlag() uint8 {
	return 0x10
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) Response() bool {
	return bool(false)
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) GetResponse() bool {
	return bool(false)
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) GetQuantity() uint16 {
	return m.Quantity
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) GetValue() []byte {
	return m.Value
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewModbusPDUWriteMultipleHoldingRegistersRequest factory function for ModbusPDUWriteMultipleHoldingRegistersRequest
func NewModbusPDUWriteMultipleHoldingRegistersRequest(startingAddress uint16, quantity uint16, value []byte) *ModbusPDU {
	child := &ModbusPDUWriteMultipleHoldingRegistersRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		Value:           value,
		ModbusPDU:       NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUWriteMultipleHoldingRegistersRequest(structType interface{}) *ModbusPDUWriteMultipleHoldingRegistersRequest {
	castFunc := func(typ interface{}) *ModbusPDUWriteMultipleHoldingRegistersRequest {
		if casted, ok := typ.(ModbusPDUWriteMultipleHoldingRegistersRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUWriteMultipleHoldingRegistersRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUWriteMultipleHoldingRegistersRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUWriteMultipleHoldingRegistersRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) GetTypeName() string {
	return "ModbusPDUWriteMultipleHoldingRegistersRequest"
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUWriteMultipleHoldingRegistersRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUWriteMultipleHoldingRegistersRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (startingAddress)
	_startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field")
	}
	startingAddress := _startingAddress

	// Simple Field (quantity)
	_quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field")
	}
	quantity := _quantity

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(byteCount)
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteMultipleHoldingRegistersRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteMultipleHoldingRegistersRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		Value:           value,
		ModbusPDU:       &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteMultipleHoldingRegistersRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.StartingAddress)
		_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.Quantity)
		_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(len(m.GetValue())))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (value)
		if m.Value != nil {
			// Byte Array field (value)
			_writeArrayErr := writeBuffer.WriteByteArray("value", m.Value)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'value' field")
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteMultipleHoldingRegistersRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUWriteMultipleHoldingRegistersRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
