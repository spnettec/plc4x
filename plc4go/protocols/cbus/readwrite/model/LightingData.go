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


// LightingData is the corresponding interface of LightingData
type LightingData interface {
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() LightingCommandTypeContainer
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() LightingCommandType
}

// LightingDataExactly can be used when we want exactly this type and not a type which fulfills LightingData.
// This is useful for switch cases.
type LightingDataExactly interface {
	LightingData
	isLightingData() bool
}

// _LightingData is the data-structure of this message
type _LightingData struct {
	_LightingDataChildRequirements
        CommandTypeContainer LightingCommandTypeContainer
}

type _LightingDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}


type LightingDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child LightingData, serializeChildFunction func() error) error
	GetTypeName() string
}

type LightingDataChild interface {
	utils.Serializable
InitializeParent(parent LightingData , commandTypeContainer LightingCommandTypeContainer )
	GetParent() *LightingData

	GetTypeName() string
	LightingData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingData) GetCommandTypeContainer() LightingCommandTypeContainer {
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

func (m *_LightingData) GetCommandType() LightingCommandType {
	return CastLightingCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewLightingData factory function for _LightingData
func NewLightingData( commandTypeContainer LightingCommandTypeContainer ) *_LightingData {
return &_LightingData{ CommandTypeContainer: commandTypeContainer }
}

// Deprecated: use the interface for direct cast
func CastLightingData(structType interface{}) LightingData {
    if casted, ok := structType.(LightingData); ok {
		return casted
	}
	if casted, ok := structType.(*LightingData); ok {
		return *casted
	}
	return nil
}

func (m *_LightingData) GetTypeName() string {
	return "LightingData"
}



func (m *_LightingData) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_LightingData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LightingDataParse(readBuffer utils.ReadBuffer) (LightingData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if (!(KnowsLightingCommandTypeContainer(readBuffer))) {
		return nil, errors.WithStack(utils.ParseAssertError{"no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
_commandTypeContainer, _commandTypeContainerErr := LightingCommandTypeContainerParse(readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of LightingData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := LightingCommandType(_commandType)
	_ = commandType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type LightingDataChildSerializeRequirement interface {
		LightingData
		InitializeParent(LightingData,  LightingCommandTypeContainer)
		GetParent() LightingData
	}
	var _childTemp interface{}
	var _child LightingDataChildSerializeRequirement
	var typeSwitchError error
	switch {
case commandType == LightingCommandType_OFF : // LightingDataOff
		_childTemp, typeSwitchError = LightingDataOffParse(readBuffer, )
case commandType == LightingCommandType_ON : // LightingDataOn
		_childTemp, typeSwitchError = LightingDataOnParse(readBuffer, )
case commandType == LightingCommandType_RAMP_TO_LEVEL : // LightingDataRampToLevel
		_childTemp, typeSwitchError = LightingDataRampToLevelParse(readBuffer, )
case commandType == LightingCommandType_TERMINATE_RAMP : // LightingDataTerminateRamp
		_childTemp, typeSwitchError = LightingDataTerminateRampParse(readBuffer, )
case commandType == LightingCommandType_LABEL : // LightingDataLabel
		_childTemp, typeSwitchError = LightingDataLabelParse(readBuffer, commandTypeContainer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of LightingData")
	}
	_child = _childTemp.(LightingDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("LightingData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingData")
	}

	// Finish initializing
_child.InitializeParent(_child , commandTypeContainer )
	return _child, nil
}

func (pm *_LightingData) SerializeParent(writeBuffer utils.WriteBuffer, child LightingData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("LightingData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LightingData")
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

	if popErr := writeBuffer.PopContext("LightingData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LightingData")
	}
	return nil
}


func (m *_LightingData) isLightingData() bool {
	return true
}

func (m *_LightingData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



