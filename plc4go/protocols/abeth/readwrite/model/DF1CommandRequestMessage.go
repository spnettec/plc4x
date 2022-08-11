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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// DF1CommandRequestMessage is the corresponding interface of DF1CommandRequestMessage
type DF1CommandRequestMessage interface {
	utils.LengthAware
	utils.Serializable
	DF1RequestMessage
	// GetCommand returns Command (property field)
	GetCommand() DF1RequestCommand
}

// DF1CommandRequestMessageExactly can be used when we want exactly this type and not a type which fulfills DF1CommandRequestMessage.
// This is useful for switch cases.
type DF1CommandRequestMessageExactly interface {
	DF1CommandRequestMessage
	isDF1CommandRequestMessage() bool
}

// _DF1CommandRequestMessage is the data-structure of this message
type _DF1CommandRequestMessage struct {
	*_DF1RequestMessage
	Command DF1RequestCommand
}

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

func (m *_DF1CommandRequestMessage) InitializeParent(parent DF1RequestMessage, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) {
	m.DestinationAddress = destinationAddress
	m.SourceAddress = sourceAddress
	m.Status = status
	m.TransactionCounter = transactionCounter
}

func (m *_DF1CommandRequestMessage) GetParent() DF1RequestMessage {
	return m._DF1RequestMessage
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
func NewDF1CommandRequestMessage(command DF1RequestCommand, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) *_DF1CommandRequestMessage {
	_result := &_DF1CommandRequestMessage{
		Command:            command,
		_DF1RequestMessage: NewDF1RequestMessage(destinationAddress, sourceAddress, status, transactionCounter),
	}
	_result._DF1RequestMessage._DF1RequestMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1CommandRequestMessage(structType interface{}) DF1CommandRequestMessage {
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

func (m *_DF1CommandRequestMessage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_DF1CommandRequestMessage) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits()

	return lengthInBits
}

func (m *_DF1CommandRequestMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1CommandRequestMessageParse(readBuffer utils.ReadBuffer) (DF1CommandRequestMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1CommandRequestMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1CommandRequestMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for command")
	}
	_command, _commandErr := DF1RequestCommandParse(readBuffer)
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field of DF1CommandRequestMessage")
	}
	command := _command.(DF1RequestCommand)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for command")
	}

	if closeErr := readBuffer.CloseContext("DF1CommandRequestMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1CommandRequestMessage")
	}

	// Create a partially initialized instance
	_child := &_DF1CommandRequestMessage{
		_DF1RequestMessage: &_DF1RequestMessage{},
		Command:            command,
	}
	_child._DF1RequestMessage._DF1RequestMessageChildRequirements = _child
	return _child, nil
}

func (m *_DF1CommandRequestMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1CommandRequestMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1CommandRequestMessage")
		}

		// Simple Field (command)
		if pushErr := writeBuffer.PushContext("command"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for command")
		}
		_commandErr := writeBuffer.WriteSerializable(m.GetCommand())
		if popErr := writeBuffer.PopContext("command"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for command")
		}
		if _commandErr != nil {
			return errors.Wrap(_commandErr, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("DF1CommandRequestMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1CommandRequestMessage")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_DF1CommandRequestMessage) isDF1CommandRequestMessage() bool {
	return true
}

func (m *_DF1CommandRequestMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
