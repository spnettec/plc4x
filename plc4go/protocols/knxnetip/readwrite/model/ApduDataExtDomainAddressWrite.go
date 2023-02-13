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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtDomainAddressWrite is the corresponding interface of ApduDataExtDomainAddressWrite
type ApduDataExtDomainAddressWrite interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtDomainAddressWriteExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtDomainAddressWrite.
// This is useful for switch cases.
type ApduDataExtDomainAddressWriteExactly interface {
	ApduDataExtDomainAddressWrite
	isApduDataExtDomainAddressWrite() bool
}

// _ApduDataExtDomainAddressWrite is the data-structure of this message
type _ApduDataExtDomainAddressWrite struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressWrite) GetExtApciType() uint8 {
	return 0x20
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressWrite) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtDomainAddressWrite) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtDomainAddressWrite factory function for _ApduDataExtDomainAddressWrite
func NewApduDataExtDomainAddressWrite(length uint8) *_ApduDataExtDomainAddressWrite {
	_result := &_ApduDataExtDomainAddressWrite{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressWrite(structType interface{}) ApduDataExtDomainAddressWrite {
	if casted, ok := structType.(ApduDataExtDomainAddressWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressWrite) GetTypeName() string {
	return "ApduDataExtDomainAddressWrite"
}

func (m *_ApduDataExtDomainAddressWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtDomainAddressWriteParse(theBytes []byte, length uint8) (ApduDataExtDomainAddressWrite, error) {
	return ApduDataExtDomainAddressWriteParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtDomainAddressWriteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtDomainAddressWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressWrite")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtDomainAddressWrite{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtDomainAddressWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtDomainAddressWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressWrite")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressWrite) isApduDataExtDomainAddressWrite() bool {
	return true
}

func (m *_ApduDataExtDomainAddressWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
