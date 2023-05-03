/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const DF1Symbol_MESSAGESTART uint8 = 0x10

// DF1Symbol is the corresponding interface of DF1Symbol
type DF1Symbol interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetSymbolType returns SymbolType (discriminator field)
	GetSymbolType() uint8
}

// DF1SymbolExactly can be used when we want exactly this type and not a type which fulfills DF1Symbol.
// This is useful for switch cases.
type DF1SymbolExactly interface {
	DF1Symbol
	isDF1Symbol() bool
}

// _DF1Symbol is the data-structure of this message
type _DF1Symbol struct {
	_DF1SymbolChildRequirements
}

type _DF1SymbolChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetSymbolType() uint8
}

type DF1SymbolParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DF1Symbol, serializeChildFunction func() error) error
	GetTypeName() string
}

type DF1SymbolChild interface {
	utils.Serializable
	InitializeParent(parent DF1Symbol)
	GetParent() *DF1Symbol

	GetTypeName() string
	DF1Symbol
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_DF1Symbol) GetMessageStart() uint8 {
	return DF1Symbol_MESSAGESTART
}

///////////////////////-4
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1Symbol factory function for _DF1Symbol
func NewDF1Symbol() *_DF1Symbol {
	return &_DF1Symbol{}
}

// Deprecated: use the interface for direct cast
func CastDF1Symbol(structType any) DF1Symbol {
	if casted, ok := structType.(DF1Symbol); ok {
		return casted
	}
	if casted, ok := structType.(*DF1Symbol); ok {
		return *casted
	}
	return nil
}

func (m *_DF1Symbol) GetTypeName() string {
	return "DF1Symbol"
}

func (m *_DF1Symbol) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (messageStart)
	lengthInBits += 8
	// Discriminator Field (symbolType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DF1Symbol) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DF1SymbolParse(theBytes []byte) (DF1Symbol, error) {
	return DF1SymbolParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func DF1SymbolParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DF1Symbol, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1Symbol"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1Symbol")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (messageStart)
	messageStart, _messageStartErr := readBuffer.ReadUint8("messageStart", 8)
	if _messageStartErr != nil {
		return nil, errors.Wrap(_messageStartErr, "Error parsing 'messageStart' field of DF1Symbol")
	}
	if messageStart != DF1Symbol_MESSAGESTART {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", DF1Symbol_MESSAGESTART) + " but got " + fmt.Sprintf("%d", messageStart))
	}

	// Discriminator Field (symbolType) (Used as input to a switch field)
	symbolType, _symbolTypeErr := readBuffer.ReadUint8("symbolType", 8)
	if _symbolTypeErr != nil {
		return nil, errors.Wrap(_symbolTypeErr, "Error parsing 'symbolType' field of DF1Symbol")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type DF1SymbolChildSerializeRequirement interface {
		DF1Symbol
		InitializeParent(DF1Symbol)
		GetParent() DF1Symbol
	}
	var _childTemp any
	var _child DF1SymbolChildSerializeRequirement
	var typeSwitchError error
	switch {
	case symbolType == 0x02: // DF1SymbolMessageFrame
		_childTemp, typeSwitchError = DF1SymbolMessageFrameParseWithBuffer(ctx, readBuffer)
	case symbolType == 0x06: // DF1SymbolMessageFrameACK
		_childTemp, typeSwitchError = DF1SymbolMessageFrameACKParseWithBuffer(ctx, readBuffer)
	case symbolType == 0x15: // DF1SymbolMessageFrameNAK
		_childTemp, typeSwitchError = DF1SymbolMessageFrameNAKParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [symbolType=%v]", symbolType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of DF1Symbol")
	}
	_child = _childTemp.(DF1SymbolChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("DF1Symbol"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1Symbol")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_DF1Symbol) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DF1Symbol, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("DF1Symbol"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DF1Symbol")
	}

	// Const Field (messageStart)
	_messageStartErr := writeBuffer.WriteUint8("messageStart", 8, 0x10)
	if _messageStartErr != nil {
		return errors.Wrap(_messageStartErr, "Error serializing 'messageStart' field")
	}

	// Discriminator Field (symbolType) (Used as input to a switch field)
	symbolType := uint8(child.GetSymbolType())
	_symbolTypeErr := writeBuffer.WriteUint8("symbolType", 8, (symbolType))

	if _symbolTypeErr != nil {
		return errors.Wrap(_symbolTypeErr, "Error serializing 'symbolType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1Symbol"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DF1Symbol")
	}
	return nil
}

func (m *_DF1Symbol) isDF1Symbol() bool {
	return true
}

func (m *_DF1Symbol) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
