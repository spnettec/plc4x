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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityData is the corresponding interface of SecurityData
type SecurityData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() SecurityCommandTypeContainer
	// GetArgument returns Argument (property field)
	GetArgument() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() SecurityCommandType
}

// SecurityDataExactly can be used when we want exactly this type and not a type which fulfills SecurityData.
// This is useful for switch cases.
type SecurityDataExactly interface {
	SecurityData
	isSecurityData() bool
}

// _SecurityData is the data-structure of this message
type _SecurityData struct {
	_SecurityDataChildRequirements
	CommandTypeContainer SecurityCommandTypeContainer
	Argument             byte
}

type _SecurityDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandType() SecurityCommandType
	GetArgument() byte
}

type SecurityDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child SecurityData, serializeChildFunction func() error) error
	GetTypeName() string
}

type SecurityDataChild interface {
	utils.Serializable
	InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte)
	GetParent() *SecurityData

	GetTypeName() string
	SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityData) GetCommandTypeContainer() SecurityCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_SecurityData) GetArgument() byte {
	return m.Argument
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SecurityData) GetCommandType() SecurityCommandType {
	ctx := context.Background()
	_ = ctx
	return CastSecurityCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityData factory function for _SecurityData
func NewSecurityData(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityData {
	return &_SecurityData{CommandTypeContainer: commandTypeContainer, Argument: argument}
}

// Deprecated: use the interface for direct cast
func CastSecurityData(structType any) SecurityData {
	if casted, ok := structType.(SecurityData); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityData); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityData) GetTypeName() string {
	return "SecurityData"
}

func (m *_SecurityData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (argument)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataParse(ctx context.Context, theBytes []byte) (SecurityData, error) {
	return SecurityDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SecurityData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsSecurityCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{"no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := SecurityCommandTypeContainerParseWithBuffer(ctx, readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of SecurityData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := SecurityCommandType(_commandType)
	_ = commandType

	// Simple Field (argument)
	_argument, _argumentErr := readBuffer.ReadByte("argument")
	if _argumentErr != nil {
		return nil, errors.Wrap(_argumentErr, "Error parsing 'argument' field of SecurityData")
	}
	argument := _argument

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type SecurityDataChildSerializeRequirement interface {
		SecurityData
		InitializeParent(SecurityData, SecurityCommandTypeContainer, byte)
		GetParent() SecurityData
	}
	var _childTemp any
	var _child SecurityDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == SecurityCommandType_ON && argument == 0x80: // SecurityDataSystemArmedDisarmed
		_childTemp, typeSwitchError = SecurityDataSystemArmedDisarmedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x80: // SecurityDataSystemDisarmed
		_childTemp, typeSwitchError = SecurityDataSystemDisarmedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x81: // SecurityDataExitDelayStarted
		_childTemp, typeSwitchError = SecurityDataExitDelayStartedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x82: // SecurityDataEntryDelayStarted
		_childTemp, typeSwitchError = SecurityDataEntryDelayStartedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0x83: // SecurityDataAlarmOn
		_childTemp, typeSwitchError = SecurityDataAlarmOnParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x83: // SecurityDataAlarmOff
		_childTemp, typeSwitchError = SecurityDataAlarmOffParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0x84: // SecurityDataTamperOn
		_childTemp, typeSwitchError = SecurityDataTamperOnParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x84: // SecurityDataTamperOff
		_childTemp, typeSwitchError = SecurityDataTamperOffParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0x85: // SecurityDataPanicActivated
		_childTemp, typeSwitchError = SecurityDataPanicActivatedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x85: // SecurityDataPanicCleared
		_childTemp, typeSwitchError = SecurityDataPanicClearedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x86: // SecurityDataZoneUnsealed
		_childTemp, typeSwitchError = SecurityDataZoneUnsealedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x87: // SecurityDataZoneSealed
		_childTemp, typeSwitchError = SecurityDataZoneSealedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x88: // SecurityDataZoneOpen
		_childTemp, typeSwitchError = SecurityDataZoneOpenParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x89: // SecurityDataZoneShort
		_childTemp, typeSwitchError = SecurityDataZoneShortParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x89: // SecurityDataZoneIsolated
		_childTemp, typeSwitchError = SecurityDataZoneIsolatedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0x8B: // SecurityDataLowBatteryDetected
		_childTemp, typeSwitchError = SecurityDataLowBatteryDetectedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x8B: // SecurityDataLowBatteryCorrected
		_childTemp, typeSwitchError = SecurityDataLowBatteryCorrectedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x8C: // SecurityDataLowBatteryCharging
		_childTemp, typeSwitchError = SecurityDataLowBatteryChargingParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x8D: // SecurityDataZoneName
		_childTemp, typeSwitchError = SecurityDataZoneNameParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x8E: // SecurityDataStatusReport1
		_childTemp, typeSwitchError = SecurityDataStatusReport1ParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x8F: // SecurityDataStatusReport2
		_childTemp, typeSwitchError = SecurityDataStatusReport2ParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x90: // SecurityDataPasswordEntryStatus
		_childTemp, typeSwitchError = SecurityDataPasswordEntryStatusParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0x91: // SecurityDataMainsFailure
		_childTemp, typeSwitchError = SecurityDataMainsFailureParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x91: // SecurityDataMainsRestoredOrApplied
		_childTemp, typeSwitchError = SecurityDataMainsRestoredOrAppliedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x92: // SecurityDataArmReadyNotReady
		_childTemp, typeSwitchError = SecurityDataArmReadyNotReadyParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0x93: // SecurityDataCurrentAlarmType
		_childTemp, typeSwitchError = SecurityDataCurrentAlarmTypeParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0x94: // SecurityDataLineCutAlarmRaised
		_childTemp, typeSwitchError = SecurityDataLineCutAlarmRaisedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x94: // SecurityDataLineCutAlarmCleared
		_childTemp, typeSwitchError = SecurityDataLineCutAlarmClearedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0x95: // SecurityDataArmFailedRaised
		_childTemp, typeSwitchError = SecurityDataArmFailedRaisedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x95: // SecurityDataArmFailedCleared
		_childTemp, typeSwitchError = SecurityDataArmFailedClearedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0x96: // SecurityDataFireAlarmRaised
		_childTemp, typeSwitchError = SecurityDataFireAlarmRaisedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x96: // SecurityDataFireAlarmCleared
		_childTemp, typeSwitchError = SecurityDataFireAlarmClearedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0x97: // SecurityDataGasAlarmRaised
		_childTemp, typeSwitchError = SecurityDataGasAlarmRaisedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x97: // SecurityDataGasAlarmCleared
		_childTemp, typeSwitchError = SecurityDataGasAlarmClearedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0x98: // SecurityDataOtherAlarmRaised
		_childTemp, typeSwitchError = SecurityDataOtherAlarmRaisedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0x98: // SecurityDataOtherAlarmCleared
		_childTemp, typeSwitchError = SecurityDataOtherAlarmClearedParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0xA0: // SecurityDataStatus1Request
		_childTemp, typeSwitchError = SecurityDataStatus1RequestParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0xA1: // SecurityDataStatus2Request
		_childTemp, typeSwitchError = SecurityDataStatus2RequestParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0xA2: // SecurityDataArmSystem
		_childTemp, typeSwitchError = SecurityDataArmSystemParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0xA3: // SecurityDataRaiseTamper
		_childTemp, typeSwitchError = SecurityDataRaiseTamperParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF && argument == 0xA3: // SecurityDataDropTamper
		_childTemp, typeSwitchError = SecurityDataDropTamperParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0xA4: // SecurityDataRaiseAlarm
		_childTemp, typeSwitchError = SecurityDataRaiseAlarmParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_EVENT && argument == 0xA5: // SecurityDataEmulatedKeypad
		_childTemp, typeSwitchError = SecurityDataEmulatedKeypadParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_ON && argument == 0xA6: // SecurityDataDisplayMessage
		_childTemp, typeSwitchError = SecurityDataDisplayMessageParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == SecurityCommandType_EVENT && argument == 0xA7: // SecurityDataRequestZoneName
		_childTemp, typeSwitchError = SecurityDataRequestZoneNameParseWithBuffer(ctx, readBuffer)
	case commandType == SecurityCommandType_OFF: // SecurityDataOff
		_childTemp, typeSwitchError = SecurityDataOffParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == SecurityCommandType_ON: // SecurityDataOn
		_childTemp, typeSwitchError = SecurityDataOnParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == SecurityCommandType_EVENT: // SecurityDataEvent
		_childTemp, typeSwitchError = SecurityDataEventParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v, argument=%v]", commandType, argument)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of SecurityData")
	}
	_child = _childTemp.(SecurityDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("SecurityData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer, argument)
	return _child, nil
}

func (pm *_SecurityData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child SecurityData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SecurityData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SecurityData")
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandTypeContainer")
	}
	_commandTypeContainerErr := writeBuffer.WriteSerializable(ctx, m.GetCommandTypeContainer())
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandTypeContainer")
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Simple Field (argument)
	argument := byte(m.GetArgument())
	_argumentErr := writeBuffer.WriteByte("argument", (argument))
	if _argumentErr != nil {
		return errors.Wrap(_argumentErr, "Error serializing 'argument' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("SecurityData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SecurityData")
	}
	return nil
}

func (m *_SecurityData) isSecurityData() bool {
	return true
}

func (m *_SecurityData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
