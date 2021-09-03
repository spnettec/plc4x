/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type FirmataMessageCommand struct {
	Command *FirmataCommand
	Parent  *FirmataMessage
}

// The corresponding interface
type IFirmataMessageCommand interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *FirmataMessageCommand) MessageType() uint8 {
	return 0xF
}

func (m *FirmataMessageCommand) InitializeParent(parent *FirmataMessage) {
}

func NewFirmataMessageCommand(command *FirmataCommand) *FirmataMessage {
	child := &FirmataMessageCommand{
		Command: command,
		Parent:  NewFirmataMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastFirmataMessageCommand(structType interface{}) *FirmataMessageCommand {
	castFunc := func(typ interface{}) *FirmataMessageCommand {
		if casted, ok := typ.(FirmataMessageCommand); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataMessageCommand); ok {
			return casted
		}
		if casted, ok := typ.(FirmataMessage); ok {
			return CastFirmataMessageCommand(casted.Child)
		}
		if casted, ok := typ.(*FirmataMessage); ok {
			return CastFirmataMessageCommand(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataMessageCommand) GetTypeName() string {
	return "FirmataMessageCommand"
}

func (m *FirmataMessageCommand) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *FirmataMessageCommand) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (command)
	lengthInBits += m.Command.LengthInBits()

	return lengthInBits
}

func (m *FirmataMessageCommand) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func FirmataMessageCommandParse(readBuffer utils.ReadBuffer, response bool) (*FirmataMessage, error) {
	if pullErr := readBuffer.PullContext("FirmataMessageCommand"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, pullErr
	}
	command, _commandErr := FirmataCommandParse(readBuffer, response)
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field")
	}
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("FirmataMessageCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataMessageCommand{
		Command: command,
		Parent:  &FirmataMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *FirmataMessageCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageCommand"); pushErr != nil {
			return pushErr
		}

		// Simple Field (command)
		if pushErr := writeBuffer.PushContext("command"); pushErr != nil {
			return pushErr
		}
		_commandErr := m.Command.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("command"); popErr != nil {
			return popErr
		}
		if _commandErr != nil {
			return errors.Wrap(_commandErr, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageCommand"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *FirmataMessageCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
