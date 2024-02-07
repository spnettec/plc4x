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

// ErrorResponse is the corresponding interface of ErrorResponse
type ErrorResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
}

// ErrorResponseExactly can be used when we want exactly this type and not a type which fulfills ErrorResponse.
// This is useful for switch cases.
type ErrorResponseExactly interface {
	ErrorResponse
	isErrorResponse() bool
}

// _ErrorResponse is the data-structure of this message
type _ErrorResponse struct {
	*_AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorResponse) GetCommandId() CommandId {
	return 0
}

func (m *_ErrorResponse) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorResponse) InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) {
	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_ErrorResponse) GetParent() AmsPacket {
	return m._AmsPacket
}

// NewErrorResponse factory function for _ErrorResponse
func NewErrorResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_ErrorResponse {
	_result := &_ErrorResponse{
		_AmsPacket: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastErrorResponse(structType any) ErrorResponse {
	if casted, ok := structType.(ErrorResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorResponse) GetTypeName() string {
	return "ErrorResponse"
}

func (m *_ErrorResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ErrorResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorResponseParse(ctx context.Context, theBytes []byte) (ErrorResponse, error) {
	return ErrorResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ErrorResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ErrorResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorResponse")
	}

	// Create a partially initialized instance
	_child := &_ErrorResponse{
		_AmsPacket: &_AmsPacket{},
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_ErrorResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorResponse")
		}

		if popErr := writeBuffer.PopContext("ErrorResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorResponse) isErrorResponse() bool {
	return true
}

func (m *_ErrorResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
