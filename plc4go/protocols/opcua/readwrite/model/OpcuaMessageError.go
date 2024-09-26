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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaMessageError is the corresponding interface of OpcuaMessageError
type OpcuaMessageError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MessagePDU
	// GetError returns Error (property field)
	GetError() OpcuaStatusCode
	// GetReason returns Reason (property field)
	GetReason() PascalString
	// IsOpcuaMessageError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpcuaMessageError()
}

// _OpcuaMessageError is the data-structure of this message
type _OpcuaMessageError struct {
	MessagePDUContract
	Error  OpcuaStatusCode
	Reason PascalString
}

var _ OpcuaMessageError = (*_OpcuaMessageError)(nil)
var _ MessagePDURequirements = (*_OpcuaMessageError)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaMessageError) GetMessageType() string {
	return "ERR"
}

func (m *_OpcuaMessageError) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaMessageError) GetParent() MessagePDUContract {
	return m.MessagePDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaMessageError) GetError() OpcuaStatusCode {
	return m.Error
}

func (m *_OpcuaMessageError) GetReason() PascalString {
	return m.Reason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpcuaMessageError factory function for _OpcuaMessageError
func NewOpcuaMessageError(chunk ChunkType, error OpcuaStatusCode, reason PascalString) *_OpcuaMessageError {
	if reason == nil {
		panic("reason of type PascalString for OpcuaMessageError must not be nil")
	}
	_result := &_OpcuaMessageError{
		MessagePDUContract: NewMessagePDU(chunk),
		Error:              error,
		Reason:             reason,
	}
	_result.MessagePDUContract.(*_MessagePDU)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpcuaMessageError(structType any) OpcuaMessageError {
	if casted, ok := structType.(OpcuaMessageError); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaMessageError); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaMessageError) GetTypeName() string {
	return "OpcuaMessageError"
}

func (m *_OpcuaMessageError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MessagePDUContract.(*_MessagePDU).getLengthInBits(ctx))

	// Simple field (error)
	lengthInBits += 32

	// Simple field (reason)
	lengthInBits += m.Reason.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaMessageError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpcuaMessageError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MessagePDU, response bool) (__opcuaMessageError OpcuaMessageError, err error) {
	m.MessagePDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaMessageError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaMessageError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	error, err := ReadEnumField[OpcuaStatusCode](ctx, "error", "OpcuaStatusCode", ReadEnum(OpcuaStatusCodeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'error' field"))
	}
	m.Error = error

	reason, err := ReadSimpleField[PascalString](ctx, "reason", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reason' field"))
	}
	m.Reason = reason

	if closeErr := readBuffer.CloseContext("OpcuaMessageError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaMessageError")
	}

	return m, nil
}

func (m *_OpcuaMessageError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaMessageError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaMessageError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaMessageError")
		}

		if err := WriteSimpleEnumField[OpcuaStatusCode](ctx, "error", "OpcuaStatusCode", m.GetError(), WriteEnum[OpcuaStatusCode, uint32](OpcuaStatusCode.GetValue, OpcuaStatusCode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'error' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "reason", m.GetReason(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reason' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaMessageError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaMessageError")
		}
		return nil
	}
	return m.MessagePDUContract.(*_MessagePDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpcuaMessageError) IsOpcuaMessageError() {}

func (m *_OpcuaMessageError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
