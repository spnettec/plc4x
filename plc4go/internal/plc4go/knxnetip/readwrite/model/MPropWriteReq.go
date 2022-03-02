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
type MPropWriteReq struct {
	*CEMI

	// Arguments.
	Size uint16
}

// The corresponding interface
type IMPropWriteReq interface {
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
func (m *MPropWriteReq) MessageCode() uint8 {
	return 0xF6
}

func (m *MPropWriteReq) GetMessageCode() uint8 {
	return 0xF6
}

func (m *MPropWriteReq) InitializeParent(parent *CEMI) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewMPropWriteReq factory function for MPropWriteReq
func NewMPropWriteReq(size uint16) *CEMI {
	child := &MPropWriteReq{
		CEMI: NewCEMI(size),
	}
	child.Child = child
	return child.CEMI
}

func CastMPropWriteReq(structType interface{}) *MPropWriteReq {
	castFunc := func(typ interface{}) *MPropWriteReq {
		if casted, ok := typ.(MPropWriteReq); ok {
			return &casted
		}
		if casted, ok := typ.(*MPropWriteReq); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastMPropWriteReq(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastMPropWriteReq(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MPropWriteReq) GetTypeName() string {
	return "MPropWriteReq"
}

func (m *MPropWriteReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *MPropWriteReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *MPropWriteReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MPropWriteReqParse(readBuffer utils.ReadBuffer, size uint16) (*CEMI, error) {
	if pullErr := readBuffer.PullContext("MPropWriteReq"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MPropWriteReq"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MPropWriteReq{
		CEMI: &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child.CEMI, nil
}

func (m *MPropWriteReq) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropWriteReq"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("MPropWriteReq"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MPropWriteReq) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
