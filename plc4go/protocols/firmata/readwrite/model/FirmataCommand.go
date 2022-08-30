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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// FirmataCommand is the corresponding interface of FirmataCommand
type FirmataCommand interface {
	utils.LengthAware
	utils.Serializable
	// GetCommandCode returns CommandCode (discriminator field)
	GetCommandCode() uint8
}

// FirmataCommandExactly can be used when we want exactly this type and not a type which fulfills FirmataCommand.
// This is useful for switch cases.
type FirmataCommandExactly interface {
	FirmataCommand
	isFirmataCommand() bool
}

// _FirmataCommand is the data-structure of this message
type _FirmataCommand struct {
	_FirmataCommandChildRequirements

	// Arguments.
	Response bool
}

type _FirmataCommandChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetCommandCode() uint8
}


type FirmataCommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child FirmataCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type FirmataCommandChild interface {
	utils.Serializable
InitializeParent(parent FirmataCommand )
	GetParent() *FirmataCommand

	GetTypeName() string
	FirmataCommand
}


// NewFirmataCommand factory function for _FirmataCommand
func NewFirmataCommand( response bool ) *_FirmataCommand {
return &_FirmataCommand{ Response: response }
}

// Deprecated: use the interface for direct cast
func CastFirmataCommand(structType interface{}) FirmataCommand {
    if casted, ok := structType.(FirmataCommand); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataCommand); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataCommand) GetTypeName() string {
	return "FirmataCommand"
}



func (m *_FirmataCommand) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandCode)
	lengthInBits += 4;

	return lengthInBits
}

func (m *_FirmataCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataCommandParse(readBuffer utils.ReadBuffer, response bool) (FirmataCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode, _commandCodeErr := readBuffer.ReadUint8("commandCode", 4)
	if _commandCodeErr != nil {
		return nil, errors.Wrap(_commandCodeErr, "Error parsing 'commandCode' field of FirmataCommand")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type FirmataCommandChildSerializeRequirement interface {
		FirmataCommand
		InitializeParent(FirmataCommand )
		GetParent() FirmataCommand
	}
	var _childTemp interface{}
	var _child FirmataCommandChildSerializeRequirement
	var typeSwitchError error
	switch {
case commandCode == 0x0 : // FirmataCommandSysex
		_childTemp, typeSwitchError = FirmataCommandSysexParse(readBuffer, response)
case commandCode == 0x4 : // FirmataCommandSetPinMode
		_childTemp, typeSwitchError = FirmataCommandSetPinModeParse(readBuffer, response)
case commandCode == 0x5 : // FirmataCommandSetDigitalPinValue
		_childTemp, typeSwitchError = FirmataCommandSetDigitalPinValueParse(readBuffer, response)
case commandCode == 0x9 : // FirmataCommandProtocolVersion
		_childTemp, typeSwitchError = FirmataCommandProtocolVersionParse(readBuffer, response)
case commandCode == 0xF : // FirmataCommandSystemReset
		_childTemp, typeSwitchError = FirmataCommandSystemResetParse(readBuffer, response)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandCode=%v]", commandCode)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of FirmataCommand")
	}
	_child = _childTemp.(FirmataCommandChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("FirmataCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataCommand")
	}

	// Finish initializing
_child.InitializeParent(_child )
	return _child, nil
}

func (pm *_FirmataCommand) SerializeParent(writeBuffer utils.WriteBuffer, child FirmataCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("FirmataCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for FirmataCommand")
	}

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode := uint8(child.GetCommandCode())
	_commandCodeErr := writeBuffer.WriteUint8("commandCode", 4, (commandCode))

	if _commandCodeErr != nil {
		return errors.Wrap(_commandCodeErr, "Error serializing 'commandCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("FirmataCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for FirmataCommand")
	}
	return nil
}


////
// Arguments Getter

func (m *_FirmataCommand) GetResponse() bool {
	return m.Response
}
//
////

func (m *_FirmataCommand) isFirmataCommand() bool {
	return true
}

func (m *_FirmataCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



