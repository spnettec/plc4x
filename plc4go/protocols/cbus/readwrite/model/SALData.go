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

// SALData is the corresponding interface of SALData
type SALData interface {
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() SALCommandTypeContainer
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() SALCommandType
}

// SALDataExactly can be used when we want exactly this type and not a type which fulfills SALData.
// This is useful for switch cases.
type SALDataExactly interface {
	SALData
	isSALData() bool
}

// _SALData is the data-structure of this message
type _SALData struct {
	_SALDataChildRequirements
	CommandTypeContainer SALCommandTypeContainer
}

type _SALDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type SALDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child SALData, serializeChildFunction func() error) error
	GetTypeName() string
}

type SALDataChild interface {
	utils.Serializable
	InitializeParent(parent SALData, commandTypeContainer SALCommandTypeContainer)
	GetParent() *SALData

	GetTypeName() string
	SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALData) GetCommandTypeContainer() SALCommandTypeContainer {
	return m.CommandTypeContainer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SALData) GetCommandType() SALCommandType {
	return CastSALCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALData factory function for _SALData
func NewSALData(commandTypeContainer SALCommandTypeContainer) *_SALData {
	return &_SALData{CommandTypeContainer: commandTypeContainer}
}

// Deprecated: use the interface for direct cast
func CastSALData(structType interface{}) SALData {
	if casted, ok := structType.(SALData); ok {
		return casted
	}
	if casted, ok := structType.(*SALData); ok {
		return *casted
	}
	return nil
}

func (m *_SALData) GetTypeName() string {
	return "SALData"
}

func (m *_SALData) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_SALData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataParse(readBuffer utils.ReadBuffer) (SALData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := SALCommandTypeContainerParse(readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of SALData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := SALCommandType(_commandType)
	_ = commandType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type SALDataChildSerializeRequirement interface {
		SALData
		InitializeParent(SALData, SALCommandTypeContainer)
		GetParent() SALData
	}
	var _childTemp interface{}
	var _child SALDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == SALCommandType_OFF: // SALDataOff
		_childTemp, typeSwitchError = SALDataOffParse(readBuffer)
	case commandType == SALCommandType_ON: // SALDataOn
		_childTemp, typeSwitchError = SALDataOnParse(readBuffer)
	case commandType == SALCommandType_RAMP_TO_LEVEL: // SALDataRampToLevel
		_childTemp, typeSwitchError = SALDataRampToLevelParse(readBuffer)
	case commandType == SALCommandType_TERMINATE_RAMP: // SALDataTerminateRamp
		_childTemp, typeSwitchError = SALDataTerminateRampParse(readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of SALData")
	}
	_child = _childTemp.(SALDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("SALData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer)
	return _child, nil
}

func (pm *_SALData) SerializeParent(writeBuffer utils.WriteBuffer, child SALData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("SALData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SALData")
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandTypeContainer")
	}
	_commandTypeContainerErr := writeBuffer.WriteSerializable(m.GetCommandTypeContainer())
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandTypeContainer")
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	if _commandTypeErr := writeBuffer.WriteVirtual("commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("SALData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SALData")
	}
	return nil
}

func (m *_SALData) isSALData() bool {
	return true
}

func (m *_SALData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
