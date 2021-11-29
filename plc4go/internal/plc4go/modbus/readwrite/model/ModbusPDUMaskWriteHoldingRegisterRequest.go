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
type ModbusPDUMaskWriteHoldingRegisterRequest struct {
	ReferenceAddress uint16
	AndMask          uint16
	OrMask           uint16
	Parent           *ModbusPDU
}

// The corresponding interface
type IModbusPDUMaskWriteHoldingRegisterRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUMaskWriteHoldingRegisterRequest) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) FunctionFlag() uint8 {
	return 0x16
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) Response() bool {
	return bool(false)
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUMaskWriteHoldingRegisterRequest(referenceAddress uint16, andMask uint16, orMask uint16) *ModbusPDU {
	child := &ModbusPDUMaskWriteHoldingRegisterRequest{
		ReferenceAddress: referenceAddress,
		AndMask:          andMask,
		OrMask:           orMask,
		Parent:           NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUMaskWriteHoldingRegisterRequest(structType interface{}) *ModbusPDUMaskWriteHoldingRegisterRequest {
	castFunc := func(typ interface{}) *ModbusPDUMaskWriteHoldingRegisterRequest {
		if casted, ok := typ.(ModbusPDUMaskWriteHoldingRegisterRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUMaskWriteHoldingRegisterRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUMaskWriteHoldingRegisterRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUMaskWriteHoldingRegisterRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) GetTypeName() string {
	return "ModbusPDUMaskWriteHoldingRegisterRequest"
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (referenceAddress)
	lengthInBits += 16

	// Simple field (andMask)
	lengthInBits += 16

	// Simple field (orMask)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUMaskWriteHoldingRegisterRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUMaskWriteHoldingRegisterRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (referenceAddress)
	referenceAddress, _referenceAddressErr := readBuffer.ReadUint16("referenceAddress", 16)
	if _referenceAddressErr != nil {
		return nil, errors.Wrap(_referenceAddressErr, "Error parsing 'referenceAddress' field")
	}

	// Simple Field (andMask)
	andMask, _andMaskErr := readBuffer.ReadUint16("andMask", 16)
	if _andMaskErr != nil {
		return nil, errors.Wrap(_andMaskErr, "Error parsing 'andMask' field")
	}

	// Simple Field (orMask)
	orMask, _orMaskErr := readBuffer.ReadUint16("orMask", 16)
	if _orMaskErr != nil {
		return nil, errors.Wrap(_orMaskErr, "Error parsing 'orMask' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUMaskWriteHoldingRegisterRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUMaskWriteHoldingRegisterRequest{
		ReferenceAddress: referenceAddress,
		AndMask:          andMask,
		OrMask:           orMask,
		Parent:           &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUMaskWriteHoldingRegisterRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (referenceAddress)
		referenceAddress := uint16(m.ReferenceAddress)
		_referenceAddressErr := writeBuffer.WriteUint16("referenceAddress", 16, (referenceAddress))
		if _referenceAddressErr != nil {
			return errors.Wrap(_referenceAddressErr, "Error serializing 'referenceAddress' field")
		}

		// Simple Field (andMask)
		andMask := uint16(m.AndMask)
		_andMaskErr := writeBuffer.WriteUint16("andMask", 16, (andMask))
		if _andMaskErr != nil {
			return errors.Wrap(_andMaskErr, "Error serializing 'andMask' field")
		}

		// Simple Field (orMask)
		orMask := uint16(m.OrMask)
		_orMaskErr := writeBuffer.WriteUint16("orMask", 16, (orMask))
		if _orMaskErr != nil {
			return errors.Wrap(_orMaskErr, "Error serializing 'orMask' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUMaskWriteHoldingRegisterRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUMaskWriteHoldingRegisterRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
