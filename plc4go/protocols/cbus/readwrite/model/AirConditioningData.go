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

// AirConditioningData is the corresponding interface of AirConditioningData
type AirConditioningData interface {
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() AirConditioningCommandTypeContainer
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() AirConditioningCommandType
}

// AirConditioningDataExactly can be used when we want exactly this type and not a type which fulfills AirConditioningData.
// This is useful for switch cases.
type AirConditioningDataExactly interface {
	AirConditioningData
	isAirConditioningData() bool
}

// _AirConditioningData is the data-structure of this message
type _AirConditioningData struct {
	_AirConditioningDataChildRequirements
	CommandTypeContainer AirConditioningCommandTypeContainer
}

type _AirConditioningDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type AirConditioningDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child AirConditioningData, serializeChildFunction func() error) error
	GetTypeName() string
}

type AirConditioningDataChild interface {
	utils.Serializable
	InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer)
	GetParent() *AirConditioningData

	GetTypeName() string
	AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningData) GetCommandTypeContainer() AirConditioningCommandTypeContainer {
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

func (m *_AirConditioningData) GetCommandType() AirConditioningCommandType {
	return CastAirConditioningCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningData factory function for _AirConditioningData
func NewAirConditioningData(commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningData {
	return &_AirConditioningData{CommandTypeContainer: commandTypeContainer}
}

// Deprecated: use the interface for direct cast
func CastAirConditioningData(structType interface{}) AirConditioningData {
	if casted, ok := structType.(AirConditioningData); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningData); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningData) GetTypeName() string {
	return "AirConditioningData"
}

func (m *_AirConditioningData) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_AirConditioningData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AirConditioningDataParse(readBuffer utils.ReadBuffer) (AirConditioningData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsAirConditioningCommandTypeContainer(readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{"no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := AirConditioningCommandTypeContainerParse(readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of AirConditioningData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := AirConditioningCommandType(_commandType)
	_ = commandType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type AirConditioningDataChildSerializeRequirement interface {
		AirConditioningData
		InitializeParent(AirConditioningData, AirConditioningCommandTypeContainer)
		GetParent() AirConditioningData
	}
	var _childTemp interface{}
	var _child AirConditioningDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == AirConditioningCommandType_HVAC_SCHEDULE_ENTRY: // AirConditioningDataHvacScheduleEntry
		_childTemp, typeSwitchError = AirConditioningDataHvacScheduleEntryParse(readBuffer)
	case commandType == AirConditioningCommandType_HUMIDITY_SCHEDULE_ENTRY: // AirConditioningDataHvacScheduleEntry
		_childTemp, typeSwitchError = AirConditioningDataHvacScheduleEntryParse(readBuffer)
	case commandType == AirConditioningCommandType_REFRESH: // AirConditioningDataRefresh
		_childTemp, typeSwitchError = AirConditioningDataRefreshParse(readBuffer)
	case commandType == AirConditioningCommandType_ZONE_HVAC_PLANT_STATUS: // AirConditioningDataZoneHvacPlantStatus
		_childTemp, typeSwitchError = AirConditioningDataZoneHvacPlantStatusParse(readBuffer)
	case commandType == AirConditioningCommandType_ZONE_HUMIDITY_PLANT_STATUS: // AirConditioningDataZoneHumidityPlantStatus
		_childTemp, typeSwitchError = AirConditioningDataZoneHumidityPlantStatusParse(readBuffer)
	case commandType == AirConditioningCommandType_ZONE_TEMPERATURE: // AirConditioningDataZoneTemperature
		_childTemp, typeSwitchError = AirConditioningDataZoneTemperatureParse(readBuffer)
	case commandType == AirConditioningCommandType_ZONE_HUMIDITY: // AirConditioningDataZoneHumidity
		_childTemp, typeSwitchError = AirConditioningDataZoneHumidityParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_ZONE_GROUP_OFF: // AirConditioningDataSetZoneGroupOff
		_childTemp, typeSwitchError = AirConditioningDataSetZoneGroupOffParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_ZONE_GROUP_ON: // AirConditioningDataSetZoneGroupOn
		_childTemp, typeSwitchError = AirConditioningDataSetZoneGroupOnParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_ZONE_HVAC_MODE: // AirConditioningDataSetZoneHvacMode
		_childTemp, typeSwitchError = AirConditioningDataSetZoneHvacModeParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_PLANT_HVAC_LEVEL: // AirConditioningDataSetPlantHvacLevel
		_childTemp, typeSwitchError = AirConditioningDataSetPlantHvacLevelParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_ZONE_HUMIDITY_MODE: // AirConditioningDataSetZoneHumidityMode
		_childTemp, typeSwitchError = AirConditioningDataSetZoneHumidityModeParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_PLANT_HUMIDITY_LEVEL: // AirConditioningDataSetPlantHumidityLevel
		_childTemp, typeSwitchError = AirConditioningDataSetPlantHumidityLevelParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_HVAC_UPPER_GUARD_LIMIT: // AirConditioningDataSetHvacUpperGuardLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHvacUpperGuardLimitParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_HVAC_LOWER_GUARD_LIMIT: // AirConditioningDataSetHvacLowerGuardLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHvacLowerGuardLimitParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_HVAC_SETBACK_LIMIT: // AirConditioningDataSetHvacSetbackLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHvacSetbackLimitParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_HUMIDITY_UPPER_GUARD_LIMIT: // AirConditioningDataSetHumidityUpperGuardLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHumidityUpperGuardLimitParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_HUMIDITY_LOWER_GUARD_LIMIT: // AirConditioningDataSetHumidityLowerGuardLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHumidityLowerGuardLimitParse(readBuffer)
	case commandType == AirConditioningCommandType_SET_HUMIDITY_SETBACK_LIMIT: // AirConditioningDataSetHumiditySetbackLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHumiditySetbackLimitParse(readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of AirConditioningData")
	}
	_child = _childTemp.(AirConditioningDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("AirConditioningData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer)
	return _child, nil
}

func (pm *_AirConditioningData) SerializeParent(writeBuffer utils.WriteBuffer, child AirConditioningData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AirConditioningData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AirConditioningData")
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

	if popErr := writeBuffer.PopContext("AirConditioningData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AirConditioningData")
	}
	return nil
}

func (m *_AirConditioningData) isAirConditioningData() bool {
	return true
}

func (m *_AirConditioningData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
