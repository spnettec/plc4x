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

// DF1CommandResponseMessageProtectedTypedLogicalRead is the corresponding interface of DF1CommandResponseMessageProtectedTypedLogicalRead
type DF1CommandResponseMessageProtectedTypedLogicalRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	DF1ResponseMessage
	// GetData returns Data (property field)
	GetData() []uint8
	// IsDF1CommandResponseMessageProtectedTypedLogicalRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1CommandResponseMessageProtectedTypedLogicalRead()
}

// _DF1CommandResponseMessageProtectedTypedLogicalRead is the data-structure of this message
type _DF1CommandResponseMessageProtectedTypedLogicalRead struct {
	DF1ResponseMessageContract
	Data []uint8
}

var _ DF1CommandResponseMessageProtectedTypedLogicalRead = (*_DF1CommandResponseMessageProtectedTypedLogicalRead)(nil)
var _ DF1ResponseMessageRequirements = (*_DF1CommandResponseMessageProtectedTypedLogicalRead)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetCommandCode() uint8 {
	return 0x4F
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetParent() DF1ResponseMessageContract {
	return m.DF1ResponseMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetData() []uint8 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1CommandResponseMessageProtectedTypedLogicalRead factory function for _DF1CommandResponseMessageProtectedTypedLogicalRead
func NewDF1CommandResponseMessageProtectedTypedLogicalRead(destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16, data []uint8, payloadLength uint16) *_DF1CommandResponseMessageProtectedTypedLogicalRead {
	_result := &_DF1CommandResponseMessageProtectedTypedLogicalRead{
		DF1ResponseMessageContract: NewDF1ResponseMessage(destinationAddress, sourceAddress, status, transactionCounter, payloadLength),
		Data:                       data,
	}
	_result.DF1ResponseMessageContract.(*_DF1ResponseMessage)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1CommandResponseMessageProtectedTypedLogicalRead(structType any) DF1CommandResponseMessageProtectedTypedLogicalRead {
	if casted, ok := structType.(DF1CommandResponseMessageProtectedTypedLogicalRead); ok {
		return casted
	}
	if casted, ok := structType.(*DF1CommandResponseMessageProtectedTypedLogicalRead); ok {
		return *casted
	}
	return nil
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetTypeName() string {
	return "DF1CommandResponseMessageProtectedTypedLogicalRead"
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DF1ResponseMessageContract.(*_DF1ResponseMessage).getLengthInBits(ctx))

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DF1ResponseMessage, payloadLength uint16) (__dF1CommandResponseMessageProtectedTypedLogicalRead DF1CommandResponseMessageProtectedTypedLogicalRead, err error) {
	m.DF1ResponseMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1CommandResponseMessageProtectedTypedLogicalRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := ReadLengthArrayField[uint8](ctx, "data", ReadUnsignedByte(readBuffer, uint8(8)), int(int32(payloadLength)-int32(int32(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1CommandResponseMessageProtectedTypedLogicalRead")
	}

	return m, nil
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1CommandResponseMessageProtectedTypedLogicalRead")
		}

		if err := WriteSimpleTypeArrayField(ctx, "data", m.GetData(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1CommandResponseMessageProtectedTypedLogicalRead")
		}
		return nil
	}
	return m.DF1ResponseMessageContract.(*_DF1ResponseMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) IsDF1CommandResponseMessageProtectedTypedLogicalRead() {
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
