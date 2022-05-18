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
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtDomainAddressSerialNumberRead is the data-structure of this message
type ApduDataExtDomainAddressSerialNumberRead struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtDomainAddressSerialNumberRead is the corresponding interface of ApduDataExtDomainAddressSerialNumberRead
type IApduDataExtDomainAddressSerialNumberRead interface {
	IApduDataExt
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

func (m *ApduDataExtDomainAddressSerialNumberRead) GetExtApciType() uint8 {
	return 0x2C
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtDomainAddressSerialNumberRead) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtDomainAddressSerialNumberRead) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtDomainAddressSerialNumberRead factory function for ApduDataExtDomainAddressSerialNumberRead
func NewApduDataExtDomainAddressSerialNumberRead(length uint8) *ApduDataExtDomainAddressSerialNumberRead {
	_result := &ApduDataExtDomainAddressSerialNumberRead{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtDomainAddressSerialNumberRead(structType interface{}) *ApduDataExtDomainAddressSerialNumberRead {
	if casted, ok := structType.(ApduDataExtDomainAddressSerialNumberRead); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressSerialNumberRead); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtDomainAddressSerialNumberRead(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtDomainAddressSerialNumberRead(casted.Child)
	}
	return nil
}

func (m *ApduDataExtDomainAddressSerialNumberRead) GetTypeName() string {
	return "ApduDataExtDomainAddressSerialNumberRead"
}

func (m *ApduDataExtDomainAddressSerialNumberRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtDomainAddressSerialNumberRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtDomainAddressSerialNumberRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtDomainAddressSerialNumberReadParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtDomainAddressSerialNumberRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressSerialNumberRead"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressSerialNumberRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtDomainAddressSerialNumberRead{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtDomainAddressSerialNumberRead) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressSerialNumberRead"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressSerialNumberRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtDomainAddressSerialNumberRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
