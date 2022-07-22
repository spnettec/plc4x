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

// MeasurementData is the corresponding interface of MeasurementData
type MeasurementData interface {
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() MeasurementCommandTypeContainer
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() MeasurementCommandType
}

// MeasurementDataExactly can be used when we want exactly this type and not a type which fulfills MeasurementData.
// This is useful for switch cases.
type MeasurementDataExactly interface {
	MeasurementData
	isMeasurementData() bool
}

// _MeasurementData is the data-structure of this message
type _MeasurementData struct {
	_MeasurementDataChildRequirements
	CommandTypeContainer MeasurementCommandTypeContainer
}

type _MeasurementDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type MeasurementDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child MeasurementData, serializeChildFunction func() error) error
	GetTypeName() string
}

type MeasurementDataChild interface {
	utils.Serializable
	InitializeParent(parent MeasurementData, commandTypeContainer MeasurementCommandTypeContainer)
	GetParent() *MeasurementData

	GetTypeName() string
	MeasurementData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeasurementData) GetCommandTypeContainer() MeasurementCommandTypeContainer {
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

func (m *_MeasurementData) GetCommandType() MeasurementCommandType {
	return CastMeasurementCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMeasurementData factory function for _MeasurementData
func NewMeasurementData(commandTypeContainer MeasurementCommandTypeContainer) *_MeasurementData {
	return &_MeasurementData{CommandTypeContainer: commandTypeContainer}
}

// Deprecated: use the interface for direct cast
func CastMeasurementData(structType interface{}) MeasurementData {
	if casted, ok := structType.(MeasurementData); ok {
		return casted
	}
	if casted, ok := structType.(*MeasurementData); ok {
		return *casted
	}
	return nil
}

func (m *_MeasurementData) GetTypeName() string {
	return "MeasurementData"
}

func (m *_MeasurementData) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MeasurementData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MeasurementDataParse(readBuffer utils.ReadBuffer) (MeasurementData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeasurementData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeasurementData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsMeasurementCommandTypeContainer(readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{"no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := MeasurementCommandTypeContainerParse(readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of MeasurementData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := MeasurementCommandType(_commandType)
	_ = commandType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type MeasurementDataChildSerializeRequirement interface {
		MeasurementData
		InitializeParent(MeasurementData, MeasurementCommandTypeContainer)
		GetParent() MeasurementData
	}
	var _childTemp interface{}
	var _child MeasurementDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == MeasurementCommandType_MEASUREMENT_EVENT: // MeasurementDataChannelMeasurementData
		_childTemp, typeSwitchError = MeasurementDataChannelMeasurementDataParse(readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of MeasurementData")
	}
	_child = _childTemp.(MeasurementDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("MeasurementData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeasurementData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer)
	return _child, nil
}

func (pm *_MeasurementData) SerializeParent(writeBuffer utils.WriteBuffer, child MeasurementData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("MeasurementData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MeasurementData")
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

	if popErr := writeBuffer.PopContext("MeasurementData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MeasurementData")
	}
	return nil
}

func (m *_MeasurementData) isMeasurementData() bool {
	return true
}

func (m *_MeasurementData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
