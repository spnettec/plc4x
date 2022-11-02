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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// FirmataCommandSysex is the corresponding interface of FirmataCommandSysex
type FirmataCommandSysex interface {
	utils.LengthAware
	utils.Serializable
	FirmataCommand
	// GetCommand returns Command (property field)
	GetCommand() SysexCommand
}

// FirmataCommandSysexExactly can be used when we want exactly this type and not a type which fulfills FirmataCommandSysex.
// This is useful for switch cases.
type FirmataCommandSysexExactly interface {
	FirmataCommandSysex
	isFirmataCommandSysex() bool
}

// _FirmataCommandSysex is the data-structure of this message
type _FirmataCommandSysex struct {
	*_FirmataCommand
        Command SysexCommand
	// Reserved Fields
	reservedField0 *uint8
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataCommandSysex)  GetCommandCode() uint8 {
return 0x0}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataCommandSysex) InitializeParent(parent FirmataCommand ) {}

func (m *_FirmataCommandSysex)  GetParent() FirmataCommand {
	return m._FirmataCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataCommandSysex) GetCommand() SysexCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewFirmataCommandSysex factory function for _FirmataCommandSysex
func NewFirmataCommandSysex( command SysexCommand , response bool ) *_FirmataCommandSysex {
	_result := &_FirmataCommandSysex{
		Command: command,
    	_FirmataCommand: NewFirmataCommand(response),
	}
	_result._FirmataCommand._FirmataCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFirmataCommandSysex(structType interface{}) FirmataCommandSysex {
    if casted, ok := structType.(FirmataCommandSysex); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataCommandSysex); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataCommandSysex) GetTypeName() string {
	return "FirmataCommandSysex"
}

func (m *_FirmataCommandSysex) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_FirmataCommandSysex) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits()

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}


func (m *_FirmataCommandSysex) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataCommandSysexParse(readBuffer utils.ReadBuffer, response bool) (FirmataCommandSysex, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataCommandSysex"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataCommandSysex")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for command")
	}
_command, _commandErr := SysexCommandParse(readBuffer , bool( response ) )
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field of FirmataCommandSysex")
	}
	command := _command.(SysexCommand)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for command")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of FirmataCommandSysex")
		}
		if reserved != uint8(0xF7) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0xF7),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("FirmataCommandSysex"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataCommandSysex")
	}

	// Create a partially initialized instance
	_child := &_FirmataCommandSysex{
		_FirmataCommand: &_FirmataCommand{
			Response: response,
		},
		Command: command,
		reservedField0: reservedField0,
	}
	_child._FirmataCommand._FirmataCommandChildRequirements = _child
	return _child, nil
}

func (m *_FirmataCommandSysex) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataCommandSysex) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSysex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataCommandSysex")
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

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0xF7)
		if m.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0xF7),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteUint8("reserved", 8, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

		if popErr := writeBuffer.PopContext("FirmataCommandSysex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataCommandSysex")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_FirmataCommandSysex) isFirmataCommandSysex() bool {
	return true
}

func (m *_FirmataCommandSysex) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



