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
type LRawReq struct {
	Parent *CEMI
}

// The corresponding interface
type ILRawReq interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LRawReq) MessageCode() uint8 {
	return 0x10
}

func (m *LRawReq) InitializeParent(parent *CEMI) {
}

func NewLRawReq() *CEMI {
	child := &LRawReq{
		Parent: NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastLRawReq(structType interface{}) *LRawReq {
	castFunc := func(typ interface{}) *LRawReq {
		if casted, ok := typ.(LRawReq); ok {
			return &casted
		}
		if casted, ok := typ.(*LRawReq); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastLRawReq(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastLRawReq(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *LRawReq) GetTypeName() string {
	return "LRawReq"
}

func (m *LRawReq) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *LRawReq) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *LRawReq) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func LRawReqParse(readBuffer utils.ReadBuffer, size uint16) (*CEMI, error) {
	if pullErr := readBuffer.PullContext("LRawReq"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("LRawReq"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &LRawReq{
		Parent: &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *LRawReq) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LRawReq"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("LRawReq"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *LRawReq) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
