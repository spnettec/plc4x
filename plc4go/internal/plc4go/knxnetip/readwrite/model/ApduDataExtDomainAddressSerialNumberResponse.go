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

// The data-structure of this message
type ApduDataExtDomainAddressSerialNumberResponse struct {
	*ApduDataExt
}

// The corresponding interface
type IApduDataExtDomainAddressSerialNumberResponse interface {
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtDomainAddressSerialNumberResponse) ExtApciType() uint8 {
	return 0x2D
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) GetExtApciType() uint8 {
	return 0x2D
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) InitializeParent(parent *ApduDataExt) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewApduDataExtDomainAddressSerialNumberResponse() *ApduDataExt {
	child := &ApduDataExtDomainAddressSerialNumberResponse{
		ApduDataExt: NewApduDataExt(),
	}
	child.Child = child
	return child.ApduDataExt
}

func CastApduDataExtDomainAddressSerialNumberResponse(structType interface{}) *ApduDataExtDomainAddressSerialNumberResponse {
	castFunc := func(typ interface{}) *ApduDataExtDomainAddressSerialNumberResponse {
		if casted, ok := typ.(ApduDataExtDomainAddressSerialNumberResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtDomainAddressSerialNumberResponse); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtDomainAddressSerialNumberResponse(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtDomainAddressSerialNumberResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) GetTypeName() string {
	return "ApduDataExtDomainAddressSerialNumberResponse"
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtDomainAddressSerialNumberResponseParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressSerialNumberResponse"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressSerialNumberResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtDomainAddressSerialNumberResponse{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child.ApduDataExt, nil
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressSerialNumberResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressSerialNumberResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtDomainAddressSerialNumberResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
