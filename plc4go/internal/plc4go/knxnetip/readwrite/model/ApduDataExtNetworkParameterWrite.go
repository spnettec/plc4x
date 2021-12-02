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
type ApduDataExtNetworkParameterWrite struct {
	*ApduDataExt
}

// The corresponding interface
type IApduDataExtNetworkParameterWrite interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtNetworkParameterWrite) ExtApciType() uint8 {
	return 0x24
}

func (m *ApduDataExtNetworkParameterWrite) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtNetworkParameterWrite() *ApduDataExt {
	child := &ApduDataExtNetworkParameterWrite{
		ApduDataExt: NewApduDataExt(),
	}
	child.Child = child
	return child.ApduDataExt
}

func CastApduDataExtNetworkParameterWrite(structType interface{}) *ApduDataExtNetworkParameterWrite {
	castFunc := func(typ interface{}) *ApduDataExtNetworkParameterWrite {
		if casted, ok := typ.(ApduDataExtNetworkParameterWrite); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtNetworkParameterWrite); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtNetworkParameterWrite(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtNetworkParameterWrite(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtNetworkParameterWrite) GetTypeName() string {
	return "ApduDataExtNetworkParameterWrite"
}

func (m *ApduDataExtNetworkParameterWrite) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExtNetworkParameterWrite) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtNetworkParameterWrite) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtNetworkParameterWriteParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtNetworkParameterWrite"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtNetworkParameterWrite"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtNetworkParameterWrite{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child.ApduDataExt, nil
}

func (m *ApduDataExtNetworkParameterWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtNetworkParameterWrite"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtNetworkParameterWrite"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtNetworkParameterWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
