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
type MFuncPropCon struct {
	*CEMI
}

// The corresponding interface
type IMFuncPropCon interface {
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
func (m *MFuncPropCon) MessageCode() uint8 {
	return 0xFA
}

func (m *MFuncPropCon) GetMessageCode() uint8 {
	return 0xFA
}

func (m *MFuncPropCon) InitializeParent(parent *CEMI) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewMFuncPropCon() *CEMI {
	child := &MFuncPropCon{
		CEMI: NewCEMI(),
	}
	child.Child = child
	return child.CEMI
}

func CastMFuncPropCon(structType interface{}) *MFuncPropCon {
	castFunc := func(typ interface{}) *MFuncPropCon {
		if casted, ok := typ.(MFuncPropCon); ok {
			return &casted
		}
		if casted, ok := typ.(*MFuncPropCon); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastMFuncPropCon(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastMFuncPropCon(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MFuncPropCon) GetTypeName() string {
	return "MFuncPropCon"
}

func (m *MFuncPropCon) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *MFuncPropCon) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *MFuncPropCon) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MFuncPropConParse(readBuffer utils.ReadBuffer, size uint16) (*CEMI, error) {
	if pullErr := readBuffer.PullContext("MFuncPropCon"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("MFuncPropCon"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MFuncPropCon{
		CEMI: &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child.CEMI, nil
}

func (m *MFuncPropCon) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropCon"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("MFuncPropCon"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MFuncPropCon) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
