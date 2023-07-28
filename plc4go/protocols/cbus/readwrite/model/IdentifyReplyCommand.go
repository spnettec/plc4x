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


// IdentifyReplyCommand is the corresponding interface of IdentifyReplyCommand
type IdentifyReplyCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetAttribute returns Attribute (discriminator field)
	GetAttribute() Attribute
}

// IdentifyReplyCommandExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommand.
// This is useful for switch cases.
type IdentifyReplyCommandExactly interface {
	IdentifyReplyCommand
	isIdentifyReplyCommand() bool
}

// _IdentifyReplyCommand is the data-structure of this message
type _IdentifyReplyCommand struct {
	_IdentifyReplyCommandChildRequirements

	// Arguments.
	NumBytes uint8
}

type _IdentifyReplyCommandChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetAttribute() Attribute
}


type IdentifyReplyCommandParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child IdentifyReplyCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type IdentifyReplyCommandChild interface {
	utils.Serializable
InitializeParent(parent IdentifyReplyCommand )
	GetParent() *IdentifyReplyCommand

	GetTypeName() string
	IdentifyReplyCommand
}


// NewIdentifyReplyCommand factory function for _IdentifyReplyCommand
func NewIdentifyReplyCommand( numBytes uint8 ) *_IdentifyReplyCommand {
return &_IdentifyReplyCommand{ NumBytes: numBytes }
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommand(structType any) IdentifyReplyCommand {
    if casted, ok := structType.(IdentifyReplyCommand); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommand) GetTypeName() string {
	return "IdentifyReplyCommand"
}


func (m *_IdentifyReplyCommand) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_IdentifyReplyCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandParse(ctx context.Context, theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommand, error) {
	return IdentifyReplyCommandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("IdentifyReplyCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type IdentifyReplyCommandChildSerializeRequirement interface {
		IdentifyReplyCommand
		InitializeParent(IdentifyReplyCommand )
		GetParent() IdentifyReplyCommand
	}
	var _childTemp any
	var _child IdentifyReplyCommandChildSerializeRequirement
	var typeSwitchError error
	switch {
case attribute == Attribute_Manufacturer : // IdentifyReplyCommandManufacturer
		_childTemp, typeSwitchError = IdentifyReplyCommandManufacturerParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_Type : // IdentifyReplyCommandType
		_childTemp, typeSwitchError = IdentifyReplyCommandTypeParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_FirmwareVersion : // IdentifyReplyCommandFirmwareVersion
		_childTemp, typeSwitchError = IdentifyReplyCommandFirmwareVersionParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_Summary : // IdentifyReplyCommandSummary
		_childTemp, typeSwitchError = IdentifyReplyCommandSummaryParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_ExtendedDiagnosticSummary : // IdentifyReplyCommandExtendedDiagnosticSummary
		_childTemp, typeSwitchError = IdentifyReplyCommandExtendedDiagnosticSummaryParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_NetworkTerminalLevels : // IdentifyReplyCommandNetworkTerminalLevels
		_childTemp, typeSwitchError = IdentifyReplyCommandNetworkTerminalLevelsParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_TerminalLevel : // IdentifyReplyCommandTerminalLevels
		_childTemp, typeSwitchError = IdentifyReplyCommandTerminalLevelsParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_NetworkVoltage : // IdentifyReplyCommandNetworkVoltage
		_childTemp, typeSwitchError = IdentifyReplyCommandNetworkVoltageParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_GAVValuesCurrent : // IdentifyReplyCommandGAVValuesCurrent
		_childTemp, typeSwitchError = IdentifyReplyCommandGAVValuesCurrentParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_GAVValuesStored : // IdentifyReplyCommandGAVValuesStored
		_childTemp, typeSwitchError = IdentifyReplyCommandGAVValuesStoredParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_GAVPhysicalAddresses : // IdentifyReplyCommandGAVPhysicalAddresses
		_childTemp, typeSwitchError = IdentifyReplyCommandGAVPhysicalAddressesParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_LogicalAssignment : // IdentifyReplyCommandLogicalAssignment
		_childTemp, typeSwitchError = IdentifyReplyCommandLogicalAssignmentParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_Delays : // IdentifyReplyCommandDelays
		_childTemp, typeSwitchError = IdentifyReplyCommandDelaysParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_MinimumLevels : // IdentifyReplyCommandMinimumLevels
		_childTemp, typeSwitchError = IdentifyReplyCommandMinimumLevelsParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_MaximumLevels : // IdentifyReplyCommandMaximumLevels
		_childTemp, typeSwitchError = IdentifyReplyCommandMaximumLevelsParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_CurrentSenseLevels : // IdentifyReplyCommandCurrentSenseLevels
		_childTemp, typeSwitchError = IdentifyReplyCommandCurrentSenseLevelsParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_OutputUnitSummary : // IdentifyReplyCommandOutputUnitSummary
		_childTemp, typeSwitchError = IdentifyReplyCommandOutputUnitSummaryParseWithBuffer(ctx, readBuffer, attribute, numBytes)
case attribute == Attribute_DSIStatus : // IdentifyReplyCommandDSIStatus
		_childTemp, typeSwitchError = IdentifyReplyCommandDSIStatusParseWithBuffer(ctx, readBuffer, attribute, numBytes)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [attribute=%v]", attribute)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of IdentifyReplyCommand")
	}
	_child = _childTemp.(IdentifyReplyCommandChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommand")
	}

	// Finish initializing
_child.InitializeParent(_child )
	return _child, nil
}

func (pm *_IdentifyReplyCommand) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child IdentifyReplyCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("IdentifyReplyCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommand")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("IdentifyReplyCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for IdentifyReplyCommand")
	}
	return nil
}


////
// Arguments Getter

func (m *_IdentifyReplyCommand) GetNumBytes() uint8 {
	return m.NumBytes
}
//
////

func (m *_IdentifyReplyCommand) isIdentifyReplyCommand() bool {
	return true
}

func (m *_IdentifyReplyCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



