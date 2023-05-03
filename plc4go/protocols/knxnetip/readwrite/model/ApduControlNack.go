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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduControlNack is the corresponding interface of ApduControlNack
type ApduControlNack interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduControl
}

// ApduControlNackExactly can be used when we want exactly this type and not a type which fulfills ApduControlNack.
// This is useful for switch cases.
type ApduControlNackExactly interface {
	ApduControlNack
	isApduControlNack() bool
}

// _ApduControlNack is the data-structure of this message
type _ApduControlNack struct {
	*_ApduControl
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduControlNack) GetControlType() uint8 {
	return 0x3
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduControlNack) InitializeParent(parent ApduControl) {}

func (m *_ApduControlNack) GetParent() ApduControl {
	return m._ApduControl
}

// NewApduControlNack factory function for _ApduControlNack
func NewApduControlNack() *_ApduControlNack {
	_result := &_ApduControlNack{
		_ApduControl: NewApduControl(),
	}
	_result._ApduControl._ApduControlChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduControlNack(structType any) ApduControlNack {
	if casted, ok := structType.(ApduControlNack); ok {
		return casted
	}
	if casted, ok := structType.(*ApduControlNack); ok {
		return *casted
	}
	return nil
}

func (m *_ApduControlNack) GetTypeName() string {
	return "ApduControlNack"
}

func (m *_ApduControlNack) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduControlNack) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduControlNackParse(theBytes []byte) (ApduControlNack, error) {
	return ApduControlNackParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ApduControlNackParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ApduControlNack, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduControlNack"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControlNack")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduControlNack"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControlNack")
	}

	// Create a partially initialized instance
	_child := &_ApduControlNack{
		_ApduControl: &_ApduControl{},
	}
	_child._ApduControl._ApduControlChildRequirements = _child
	return _child, nil
}

func (m *_ApduControlNack) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduControlNack) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlNack"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduControlNack")
		}

		if popErr := writeBuffer.PopContext("ApduControlNack"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduControlNack")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduControlNack) isApduControlNack() bool {
	return true
}

func (m *_ApduControlNack) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
