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

// ApduDataExtReadRouterMemoryRequest is the corresponding interface of ApduDataExtReadRouterMemoryRequest
type ApduDataExtReadRouterMemoryRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtReadRouterMemoryRequestExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtReadRouterMemoryRequest.
// This is useful for switch cases.
type ApduDataExtReadRouterMemoryRequestExactly interface {
	ApduDataExtReadRouterMemoryRequest
	isApduDataExtReadRouterMemoryRequest() bool
}

// _ApduDataExtReadRouterMemoryRequest is the data-structure of this message
type _ApduDataExtReadRouterMemoryRequest struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtReadRouterMemoryRequest) GetExtApciType() uint8 {
	return 0x08
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtReadRouterMemoryRequest) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtReadRouterMemoryRequest) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtReadRouterMemoryRequest factory function for _ApduDataExtReadRouterMemoryRequest
func NewApduDataExtReadRouterMemoryRequest(length uint8) *_ApduDataExtReadRouterMemoryRequest {
	_result := &_ApduDataExtReadRouterMemoryRequest{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtReadRouterMemoryRequest(structType any) ApduDataExtReadRouterMemoryRequest {
	if casted, ok := structType.(ApduDataExtReadRouterMemoryRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtReadRouterMemoryRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtReadRouterMemoryRequest) GetTypeName() string {
	return "ApduDataExtReadRouterMemoryRequest"
}

func (m *_ApduDataExtReadRouterMemoryRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtReadRouterMemoryRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtReadRouterMemoryRequestParse(ctx context.Context, theBytes []byte, length uint8) (ApduDataExtReadRouterMemoryRequest, error) {
	return ApduDataExtReadRouterMemoryRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtReadRouterMemoryRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtReadRouterMemoryRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduDataExtReadRouterMemoryRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtReadRouterMemoryRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRouterMemoryRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtReadRouterMemoryRequest")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtReadRouterMemoryRequest{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtReadRouterMemoryRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtReadRouterMemoryRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRouterMemoryRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtReadRouterMemoryRequest")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRouterMemoryRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtReadRouterMemoryRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtReadRouterMemoryRequest) isApduDataExtReadRouterMemoryRequest() bool {
	return true
}

func (m *_ApduDataExtReadRouterMemoryRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
