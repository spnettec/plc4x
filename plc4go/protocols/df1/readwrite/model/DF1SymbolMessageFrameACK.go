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
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// DF1SymbolMessageFrameACK is the corresponding interface of DF1SymbolMessageFrameACK
type DF1SymbolMessageFrameACK interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	DF1Symbol
}

// DF1SymbolMessageFrameACKExactly can be used when we want exactly this type and not a type which fulfills DF1SymbolMessageFrameACK.
// This is useful for switch cases.
type DF1SymbolMessageFrameACKExactly interface {
	DF1SymbolMessageFrameACK
	isDF1SymbolMessageFrameACK() bool
}

// _DF1SymbolMessageFrameACK is the data-structure of this message
type _DF1SymbolMessageFrameACK struct {
	*_DF1Symbol
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1SymbolMessageFrameACK)  GetSymbolType() uint8 {
return 0x06}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1SymbolMessageFrameACK) InitializeParent(parent DF1Symbol ) {}

func (m *_DF1SymbolMessageFrameACK)  GetParent() DF1Symbol {
	return m._DF1Symbol
}


// NewDF1SymbolMessageFrameACK factory function for _DF1SymbolMessageFrameACK
func NewDF1SymbolMessageFrameACK( ) *_DF1SymbolMessageFrameACK {
	_result := &_DF1SymbolMessageFrameACK{
    	_DF1Symbol: NewDF1Symbol(),
	}
	_result._DF1Symbol._DF1SymbolChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1SymbolMessageFrameACK(structType any) DF1SymbolMessageFrameACK {
    if casted, ok := structType.(DF1SymbolMessageFrameACK); ok {
		return casted
	}
	if casted, ok := structType.(*DF1SymbolMessageFrameACK); ok {
		return *casted
	}
	return nil
}

func (m *_DF1SymbolMessageFrameACK) GetTypeName() string {
	return "DF1SymbolMessageFrameACK"
}

func (m *_DF1SymbolMessageFrameACK) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}


func (m *_DF1SymbolMessageFrameACK) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DF1SymbolMessageFrameACKParse(ctx context.Context, theBytes []byte) (DF1SymbolMessageFrameACK, error) {
	return DF1SymbolMessageFrameACKParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func DF1SymbolMessageFrameACKParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DF1SymbolMessageFrameACK, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DF1SymbolMessageFrameACK"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1SymbolMessageFrameACK")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DF1SymbolMessageFrameACK"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1SymbolMessageFrameACK")
	}

	// Create a partially initialized instance
	_child := &_DF1SymbolMessageFrameACK{
		_DF1Symbol: &_DF1Symbol{
		},
	}
	_child._DF1Symbol._DF1SymbolChildRequirements = _child
	return _child, nil
}

func (m *_DF1SymbolMessageFrameACK) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1SymbolMessageFrameACK) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1SymbolMessageFrameACK"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1SymbolMessageFrameACK")
		}

		if popErr := writeBuffer.PopContext("DF1SymbolMessageFrameACK"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1SymbolMessageFrameACK")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_DF1SymbolMessageFrameACK) isDF1SymbolMessageFrameACK() bool {
	return true
}

func (m *_DF1SymbolMessageFrameACK) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



