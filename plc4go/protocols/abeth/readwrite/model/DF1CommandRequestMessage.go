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

// DF1CommandRequestMessage is the corresponding interface of DF1CommandRequestMessage
type DF1CommandRequestMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	DF1RequestMessage
	// GetCommand returns Command (property field)
	GetCommand() DF1RequestCommand
	// IsDF1CommandRequestMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1CommandRequestMessage()
}

// _DF1CommandRequestMessage is the data-structure of this message
type _DF1CommandRequestMessage struct {
	DF1RequestMessageContract
	Command DF1RequestCommand
}

var _ DF1CommandRequestMessage = (*_DF1CommandRequestMessage)(nil)
var _ DF1RequestMessageRequirements = (*_DF1CommandRequestMessage)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1CommandRequestMessage) GetCommandCode() uint8 {
	return 0x0F
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1CommandRequestMessage) GetParent() DF1RequestMessageContract {
	return m.DF1RequestMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1CommandRequestMessage) GetCommand() DF1RequestCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1CommandRequestMessage factory function for _DF1CommandRequestMessage
func NewDF1CommandRequestMessage(destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16, command DF1RequestCommand) *_DF1CommandRequestMessage {
	if command == nil {
		panic("command of type DF1RequestCommand for DF1CommandRequestMessage must not be nil")
	}
	_result := &_DF1CommandRequestMessage{
		DF1RequestMessageContract: NewDF1RequestMessage(destinationAddress, sourceAddress, status, transactionCounter),
		Command:                   command,
	}
	_result.DF1RequestMessageContract.(*_DF1RequestMessage)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1CommandRequestMessage(structType any) DF1CommandRequestMessage {
	if casted, ok := structType.(DF1CommandRequestMessage); ok {
		return casted
	}
	if casted, ok := structType.(*DF1CommandRequestMessage); ok {
		return *casted
	}
	return nil
}

func (m *_DF1CommandRequestMessage) GetTypeName() string {
	return "DF1CommandRequestMessage"
}

func (m *_DF1CommandRequestMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DF1RequestMessageContract.(*_DF1RequestMessage).getLengthInBits(ctx))

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DF1CommandRequestMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DF1CommandRequestMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DF1RequestMessage) (__dF1CommandRequestMessage DF1CommandRequestMessage, err error) {
	m.DF1RequestMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1CommandRequestMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1CommandRequestMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	command, err := ReadSimpleField[DF1RequestCommand](ctx, "command", ReadComplex[DF1RequestCommand](DF1RequestCommandParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}
	m.Command = command

	if closeErr := readBuffer.CloseContext("DF1CommandRequestMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1CommandRequestMessage")
	}

	return m, nil
}

func (m *_DF1CommandRequestMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1CommandRequestMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1CommandRequestMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1CommandRequestMessage")
		}

		if err := WriteSimpleField[DF1RequestCommand](ctx, "command", m.GetCommand(), WriteComplex[DF1RequestCommand](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("DF1CommandRequestMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1CommandRequestMessage")
		}
		return nil
	}
	return m.DF1RequestMessageContract.(*_DF1RequestMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1CommandRequestMessage) IsDF1CommandRequestMessage() {}

func (m *_DF1CommandRequestMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
