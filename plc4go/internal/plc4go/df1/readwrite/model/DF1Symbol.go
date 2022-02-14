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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const DF1Symbol_MESSAGESTART uint8 = 0x10

// The data-structure of this message
type DF1Symbol struct {
	Child IDF1SymbolChild
}

// The corresponding interface
type IDF1Symbol interface {
	// SymbolType returns SymbolType
	SymbolType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IDF1SymbolParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IDF1Symbol, serializeChildFunction func() error) error
	GetTypeName() string
}

type IDF1SymbolChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *DF1Symbol)
	GetTypeName() string
	IDF1Symbol
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewDF1Symbol factory function for DF1Symbol
func NewDF1Symbol() *DF1Symbol {
	return &DF1Symbol{}
}

func CastDF1Symbol(structType interface{}) *DF1Symbol {
	castFunc := func(typ interface{}) *DF1Symbol {
		if casted, ok := typ.(DF1Symbol); ok {
			return &casted
		}
		if casted, ok := typ.(*DF1Symbol); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DF1Symbol) GetTypeName() string {
	return "DF1Symbol"
}

func (m *DF1Symbol) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DF1Symbol) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *DF1Symbol) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (messageStart)
	lengthInBits += 8
	// Discriminator Field (symbolType)
	lengthInBits += 8

	return lengthInBits
}

func (m *DF1Symbol) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1SymbolParse(readBuffer utils.ReadBuffer) (*DF1Symbol, error) {
	if pullErr := readBuffer.PullContext("DF1Symbol"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (messageStart)
	messageStart, _messageStartErr := readBuffer.ReadUint8("messageStart", 8)
	if _messageStartErr != nil {
		return nil, errors.Wrap(_messageStartErr, "Error parsing 'messageStart' field")
	}
	if messageStart != DF1Symbol_MESSAGESTART {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", DF1Symbol_MESSAGESTART) + " but got " + fmt.Sprintf("%d", messageStart))
	}

	// Discriminator Field (symbolType) (Used as input to a switch field)
	symbolType, _symbolTypeErr := readBuffer.ReadUint8("symbolType", 8)
	if _symbolTypeErr != nil {
		return nil, errors.Wrap(_symbolTypeErr, "Error parsing 'symbolType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *DF1Symbol
	var typeSwitchError error
	switch {
	case symbolType == 0x02: // DF1SymbolMessageFrame
		_parent, typeSwitchError = DF1SymbolMessageFrameParse(readBuffer)
	case symbolType == 0x06: // DF1SymbolMessageFrameACK
		_parent, typeSwitchError = DF1SymbolMessageFrameACKParse(readBuffer)
	case symbolType == 0x15: // DF1SymbolMessageFrameNAK
		_parent, typeSwitchError = DF1SymbolMessageFrameNAKParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("DF1Symbol"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *DF1Symbol) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *DF1Symbol) SerializeParent(writeBuffer utils.WriteBuffer, child IDF1Symbol, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("DF1Symbol"); pushErr != nil {
		return pushErr
	}

	// Const Field (messageStart)
	_messageStartErr := writeBuffer.WriteUint8("messageStart", 8, 0x10)
	if _messageStartErr != nil {
		return errors.Wrap(_messageStartErr, "Error serializing 'messageStart' field")
	}

	// Discriminator Field (symbolType) (Used as input to a switch field)
	symbolType := uint8(child.SymbolType())
	_symbolTypeErr := writeBuffer.WriteUint8("symbolType", 8, (symbolType))

	if _symbolTypeErr != nil {
		return errors.Wrap(_symbolTypeErr, "Error serializing 'symbolType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1Symbol"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *DF1Symbol) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
