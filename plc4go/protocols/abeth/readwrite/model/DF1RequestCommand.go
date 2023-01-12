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


// DF1RequestCommand is the corresponding interface of DF1RequestCommand
type DF1RequestCommand interface {
	utils.LengthAware
	utils.Serializable
	// GetFunctionCode returns FunctionCode (discriminator field)
	GetFunctionCode() uint8
}

// DF1RequestCommandExactly can be used when we want exactly this type and not a type which fulfills DF1RequestCommand.
// This is useful for switch cases.
type DF1RequestCommandExactly interface {
	DF1RequestCommand
	isDF1RequestCommand() bool
}

// _DF1RequestCommand is the data-structure of this message
type _DF1RequestCommand struct {
	_DF1RequestCommandChildRequirements
}

type _DF1RequestCommandChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetFunctionCode() uint8
}


type DF1RequestCommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child DF1RequestCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type DF1RequestCommandChild interface {
	utils.Serializable
InitializeParent(parent DF1RequestCommand )
	GetParent() *DF1RequestCommand

	GetTypeName() string
	DF1RequestCommand
}


// NewDF1RequestCommand factory function for _DF1RequestCommand
func NewDF1RequestCommand( ) *_DF1RequestCommand {
return &_DF1RequestCommand{ }
}

// Deprecated: use the interface for direct cast
func CastDF1RequestCommand(structType interface{}) DF1RequestCommand {
    if casted, ok := structType.(DF1RequestCommand); ok {
		return casted
	}
	if casted, ok := structType.(*DF1RequestCommand); ok {
		return *casted
	}
	return nil
}

func (m *_DF1RequestCommand) GetTypeName() string {
	return "DF1RequestCommand"
}



func (m *_DF1RequestCommand) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (functionCode)
	lengthInBits += 8;

	return lengthInBits
}

func (m *_DF1RequestCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1RequestCommandParse(theBytes []byte) (DF1RequestCommand, error) {
	return DF1RequestCommandParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func DF1RequestCommandParseWithBuffer(readBuffer utils.ReadBuffer) (DF1RequestCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1RequestCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1RequestCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (functionCode) (Used as input to a switch field)
	functionCode, _functionCodeErr := readBuffer.ReadUint8("functionCode", 8)
	if _functionCodeErr != nil {
		return nil, errors.Wrap(_functionCodeErr, "Error parsing 'functionCode' field of DF1RequestCommand")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type DF1RequestCommandChildSerializeRequirement interface {
		DF1RequestCommand
		InitializeParent(DF1RequestCommand )
		GetParent() DF1RequestCommand
	}
	var _childTemp interface{}
	var _child DF1RequestCommandChildSerializeRequirement
	var typeSwitchError error
	switch {
case functionCode == 0xA2 : // DF1RequestProtectedTypedLogicalRead
		_childTemp, typeSwitchError = DF1RequestProtectedTypedLogicalReadParseWithBuffer(readBuffer, )
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [functionCode=%v]", functionCode)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of DF1RequestCommand")
	}
	_child = _childTemp.(DF1RequestCommandChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("DF1RequestCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1RequestCommand")
	}

	// Finish initializing
_child.InitializeParent(_child )
	return _child, nil
}

func (pm *_DF1RequestCommand) SerializeParent(writeBuffer utils.WriteBuffer, child DF1RequestCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("DF1RequestCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DF1RequestCommand")
	}

	// Discriminator Field (functionCode) (Used as input to a switch field)
	functionCode := uint8(child.GetFunctionCode())
	_functionCodeErr := writeBuffer.WriteUint8("functionCode", 8, (functionCode))

	if _functionCodeErr != nil {
		return errors.Wrap(_functionCodeErr, "Error serializing 'functionCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1RequestCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DF1RequestCommand")
	}
	return nil
}


func (m *_DF1RequestCommand) isDF1RequestCommand() bool {
	return true
}

func (m *_DF1RequestCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



