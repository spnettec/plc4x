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
type DF1SymbolMessageFrameNAK struct {
	*DF1Symbol
}

// The corresponding interface
type IDF1SymbolMessageFrameNAK interface {
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
func (m *DF1SymbolMessageFrameNAK) SymbolType() uint8 {
	return 0x15
}

func (m *DF1SymbolMessageFrameNAK) GetSymbolType() uint8 {
	return 0x15
}

func (m *DF1SymbolMessageFrameNAK) InitializeParent(parent *DF1Symbol) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewDF1SymbolMessageFrameNAK() *DF1Symbol {
	child := &DF1SymbolMessageFrameNAK{
		DF1Symbol: NewDF1Symbol(),
	}
	child.Child = child
	return child.DF1Symbol
}

func CastDF1SymbolMessageFrameNAK(structType interface{}) *DF1SymbolMessageFrameNAK {
	castFunc := func(typ interface{}) *DF1SymbolMessageFrameNAK {
		if casted, ok := typ.(DF1SymbolMessageFrameNAK); ok {
			return &casted
		}
		if casted, ok := typ.(*DF1SymbolMessageFrameNAK); ok {
			return casted
		}
		if casted, ok := typ.(DF1Symbol); ok {
			return CastDF1SymbolMessageFrameNAK(casted.Child)
		}
		if casted, ok := typ.(*DF1Symbol); ok {
			return CastDF1SymbolMessageFrameNAK(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DF1SymbolMessageFrameNAK) GetTypeName() string {
	return "DF1SymbolMessageFrameNAK"
}

func (m *DF1SymbolMessageFrameNAK) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DF1SymbolMessageFrameNAK) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *DF1SymbolMessageFrameNAK) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DF1SymbolMessageFrameNAKParse(readBuffer utils.ReadBuffer) (*DF1Symbol, error) {
	if pullErr := readBuffer.PullContext("DF1SymbolMessageFrameNAK"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("DF1SymbolMessageFrameNAK"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DF1SymbolMessageFrameNAK{
		DF1Symbol: &DF1Symbol{},
	}
	_child.DF1Symbol.Child = _child
	return _child.DF1Symbol, nil
}

func (m *DF1SymbolMessageFrameNAK) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1SymbolMessageFrameNAK"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("DF1SymbolMessageFrameNAK"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *DF1SymbolMessageFrameNAK) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
