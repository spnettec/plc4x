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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtDomainAddressSelectiveRead is the corresponding interface of ApduDataExtDomainAddressSelectiveRead
type ApduDataExtDomainAddressSelectiveRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtDomainAddressSelectiveReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtDomainAddressSelectiveRead.
// This is useful for switch cases.
type ApduDataExtDomainAddressSelectiveReadExactly interface {
	ApduDataExtDomainAddressSelectiveRead
	isApduDataExtDomainAddressSelectiveRead() bool
}

// _ApduDataExtDomainAddressSelectiveRead is the data-structure of this message
type _ApduDataExtDomainAddressSelectiveRead struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressSelectiveRead) GetExtApciType() uint8 {
	return 0x23
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressSelectiveRead) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtDomainAddressSelectiveRead) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtDomainAddressSelectiveRead factory function for _ApduDataExtDomainAddressSelectiveRead
func NewApduDataExtDomainAddressSelectiveRead(length uint8) *_ApduDataExtDomainAddressSelectiveRead {
	_result := &_ApduDataExtDomainAddressSelectiveRead{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressSelectiveRead(structType any) ApduDataExtDomainAddressSelectiveRead {
	if casted, ok := structType.(ApduDataExtDomainAddressSelectiveRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressSelectiveRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressSelectiveRead) GetTypeName() string {
	return "ApduDataExtDomainAddressSelectiveRead"
}

func (m *_ApduDataExtDomainAddressSelectiveRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressSelectiveRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtDomainAddressSelectiveReadParse(ctx context.Context, theBytes []byte, length uint8) (ApduDataExtDomainAddressSelectiveRead, error) {
	return ApduDataExtDomainAddressSelectiveReadParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtDomainAddressSelectiveReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtDomainAddressSelectiveRead, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressSelectiveRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressSelectiveRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressSelectiveRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressSelectiveRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtDomainAddressSelectiveRead{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtDomainAddressSelectiveRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtDomainAddressSelectiveRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressSelectiveRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressSelectiveRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressSelectiveRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressSelectiveRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressSelectiveRead) isApduDataExtDomainAddressSelectiveRead() bool {
	return true
}

func (m *_ApduDataExtDomainAddressSelectiveRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
