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
)

// Code generated by code-generation. DO NOT EDIT.

// CBusCommand is the corresponding interface of CBusCommand
type CBusCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() CBusHeader
	// GetIsDeviceManagement returns IsDeviceManagement (virtual field)
	GetIsDeviceManagement() bool
	// GetDestinationAddressType returns DestinationAddressType (virtual field)
	GetDestinationAddressType() DestinationAddressType
}

// CBusCommandExactly can be used when we want exactly this type and not a type which fulfills CBusCommand.
// This is useful for switch cases.
type CBusCommandExactly interface {
	CBusCommand
	isCBusCommand() bool
}

// _CBusCommand is the data-structure of this message
type _CBusCommand struct {
	_CBusCommandChildRequirements
	Header CBusHeader

	// Arguments.
	CBusOptions CBusOptions
}

type _CBusCommandChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetDestinationAddressType() DestinationAddressType
	GetIsDeviceManagement() bool
}

type CBusCommandParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CBusCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type CBusCommandChild interface {
	utils.Serializable
	InitializeParent(parent CBusCommand, header CBusHeader)
	GetParent() *CBusCommand

	GetTypeName() string
	CBusCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusCommand) GetHeader() CBusHeader {
	return m.Header
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_CBusCommand) GetIsDeviceManagement() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetHeader().GetDp())
}

func (m *_CBusCommand) GetDestinationAddressType() DestinationAddressType {
	ctx := context.Background()
	_ = ctx
	return CastDestinationAddressType(m.GetHeader().GetDestinationAddressType())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusCommand factory function for _CBusCommand
func NewCBusCommand(header CBusHeader, cBusOptions CBusOptions) *_CBusCommand {
	return &_CBusCommand{Header: header, CBusOptions: cBusOptions}
}

// Deprecated: use the interface for direct cast
func CastCBusCommand(structType any) CBusCommand {
	if casted, ok := structType.(CBusCommand); ok {
		return casted
	}
	if casted, ok := structType.(*CBusCommand); ok {
		return *casted
	}
	return nil
}

func (m *_CBusCommand) GetTypeName() string {
	return "CBusCommand"
}

func (m *_CBusCommand) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_CBusCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusCommandParse(theBytes []byte, cBusOptions CBusOptions) (CBusCommand, error) {
	return CBusCommandParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (CBusCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := CBusHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of CBusCommand")
	}
	header := _header.(CBusHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Virtual field
	_isDeviceManagement := header.GetDp()
	isDeviceManagement := bool(_isDeviceManagement)
	_ = isDeviceManagement

	// Virtual field
	_destinationAddressType := header.GetDestinationAddressType()
	destinationAddressType := DestinationAddressType(_destinationAddressType)
	_ = destinationAddressType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CBusCommandChildSerializeRequirement interface {
		CBusCommand
		InitializeParent(CBusCommand, CBusHeader)
		GetParent() CBusCommand
	}
	var _childTemp any
	var _child CBusCommandChildSerializeRequirement
	var typeSwitchError error
	switch {
	case 0 == 0 && isDeviceManagement == bool(true): // CBusCommandDeviceManagement
		_childTemp, typeSwitchError = CBusCommandDeviceManagementParseWithBuffer(ctx, readBuffer, cBusOptions)
	case destinationAddressType == DestinationAddressType_PointToPointToMultiPoint: // CBusCommandPointToPointToMultiPoint
		_childTemp, typeSwitchError = CBusCommandPointToPointToMultiPointParseWithBuffer(ctx, readBuffer, cBusOptions)
	case destinationAddressType == DestinationAddressType_PointToMultiPoint: // CBusCommandPointToMultiPoint
		_childTemp, typeSwitchError = CBusCommandPointToMultiPointParseWithBuffer(ctx, readBuffer, cBusOptions)
	case destinationAddressType == DestinationAddressType_PointToPoint: // CBusCommandPointToPoint
		_childTemp, typeSwitchError = CBusCommandPointToPointParseWithBuffer(ctx, readBuffer, cBusOptions)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [destinationAddressType=%v, isDeviceManagement=%v]", destinationAddressType, isDeviceManagement)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CBusCommand")
	}
	_child = _childTemp.(CBusCommandChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("CBusCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusCommand")
	}

	// Finish initializing
	_child.InitializeParent(_child, header)
	return _child, nil
}

func (pm *_CBusCommand) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CBusCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CBusCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusCommand")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}
	// Virtual field
	if _isDeviceManagementErr := writeBuffer.WriteVirtual(ctx, "isDeviceManagement", m.GetIsDeviceManagement()); _isDeviceManagementErr != nil {
		return errors.Wrap(_isDeviceManagementErr, "Error serializing 'isDeviceManagement' field")
	}
	// Virtual field
	if _destinationAddressTypeErr := writeBuffer.WriteVirtual(ctx, "destinationAddressType", m.GetDestinationAddressType()); _destinationAddressTypeErr != nil {
		return errors.Wrap(_destinationAddressTypeErr, "Error serializing 'destinationAddressType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CBusCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusCommand")
	}
	return nil
}

////
// Arguments Getter

func (m *_CBusCommand) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_CBusCommand) isCBusCommand() bool {
	return true
}

func (m *_CBusCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
