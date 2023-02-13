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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// DF1SymbolMessageFrameNAK is the corresponding interface of DF1SymbolMessageFrameNAK
type DF1SymbolMessageFrameNAK interface {
	utils.LengthAware
	utils.Serializable
	DF1Symbol
}

// DF1SymbolMessageFrameNAKExactly can be used when we want exactly this type and not a type which fulfills DF1SymbolMessageFrameNAK.
// This is useful for switch cases.
type DF1SymbolMessageFrameNAKExactly interface {
	DF1SymbolMessageFrameNAK
	isDF1SymbolMessageFrameNAK() bool
}

// _DF1SymbolMessageFrameNAK is the data-structure of this message
type _DF1SymbolMessageFrameNAK struct {
	*_DF1Symbol
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1SymbolMessageFrameNAK) GetSymbolType() uint8 {
	return 0x15
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1SymbolMessageFrameNAK) InitializeParent(parent DF1Symbol) {}

func (m *_DF1SymbolMessageFrameNAK) GetParent() DF1Symbol {
	return m._DF1Symbol
}

// NewDF1SymbolMessageFrameNAK factory function for _DF1SymbolMessageFrameNAK
func NewDF1SymbolMessageFrameNAK() *_DF1SymbolMessageFrameNAK {
	_result := &_DF1SymbolMessageFrameNAK{
		_DF1Symbol: NewDF1Symbol(),
	}
	_result._DF1Symbol._DF1SymbolChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1SymbolMessageFrameNAK(structType interface{}) DF1SymbolMessageFrameNAK {
	if casted, ok := structType.(DF1SymbolMessageFrameNAK); ok {
		return casted
	}
	if casted, ok := structType.(*DF1SymbolMessageFrameNAK); ok {
		return *casted
	}
	return nil
}

func (m *_DF1SymbolMessageFrameNAK) GetTypeName() string {
	return "DF1SymbolMessageFrameNAK"
}

func (m *_DF1SymbolMessageFrameNAK) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_DF1SymbolMessageFrameNAK) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DF1SymbolMessageFrameNAKParse(theBytes []byte) (DF1SymbolMessageFrameNAK, error) {
	return DF1SymbolMessageFrameNAKParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func DF1SymbolMessageFrameNAKParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DF1SymbolMessageFrameNAK, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1SymbolMessageFrameNAK"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1SymbolMessageFrameNAK")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DF1SymbolMessageFrameNAK"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1SymbolMessageFrameNAK")
	}

	// Create a partially initialized instance
	_child := &_DF1SymbolMessageFrameNAK{
		_DF1Symbol: &_DF1Symbol{},
	}
	_child._DF1Symbol._DF1SymbolChildRequirements = _child
	return _child, nil
}

func (m *_DF1SymbolMessageFrameNAK) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1SymbolMessageFrameNAK) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1SymbolMessageFrameNAK"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1SymbolMessageFrameNAK")
		}

		if popErr := writeBuffer.PopContext("DF1SymbolMessageFrameNAK"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1SymbolMessageFrameNAK")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1SymbolMessageFrameNAK) isDF1SymbolMessageFrameNAK() bool {
	return true
}

func (m *_DF1SymbolMessageFrameNAK) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
