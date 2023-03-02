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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// CBusPointToPointToMultiPointCommand is the corresponding interface of CBusPointToPointToMultiPointCommand
type CBusPointToPointToMultiPointCommand interface {
	utils.LengthAware
	utils.Serializable
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetNetworkRoute returns NetworkRoute (property field)
	GetNetworkRoute() NetworkRoute
	// GetPeekedApplication returns PeekedApplication (property field)
	GetPeekedApplication() byte
}

// CBusPointToPointToMultiPointCommandExactly can be used when we want exactly this type and not a type which fulfills CBusPointToPointToMultiPointCommand.
// This is useful for switch cases.
type CBusPointToPointToMultiPointCommandExactly interface {
	CBusPointToPointToMultiPointCommand
	isCBusPointToPointToMultiPointCommand() bool
}

// _CBusPointToPointToMultiPointCommand is the data-structure of this message
type _CBusPointToPointToMultiPointCommand struct {
	_CBusPointToPointToMultiPointCommandChildRequirements
        BridgeAddress BridgeAddress
        NetworkRoute NetworkRoute
        PeekedApplication byte

	// Arguments.
	CBusOptions CBusOptions
}

type _CBusPointToPointToMultiPointCommandChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedApplication() byte
}


type CBusPointToPointToMultiPointCommandParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CBusPointToPointToMultiPointCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type CBusPointToPointToMultiPointCommandChild interface {
	utils.Serializable
InitializeParent(parent CBusPointToPointToMultiPointCommand , bridgeAddress BridgeAddress , networkRoute NetworkRoute , peekedApplication byte )
	GetParent() *CBusPointToPointToMultiPointCommand

	GetTypeName() string
	CBusPointToPointToMultiPointCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointToMultiPointCommand) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_CBusPointToPointToMultiPointCommand) GetNetworkRoute() NetworkRoute {
	return m.NetworkRoute
}

func (m *_CBusPointToPointToMultiPointCommand) GetPeekedApplication() byte {
	return m.PeekedApplication
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCBusPointToPointToMultiPointCommand factory function for _CBusPointToPointToMultiPointCommand
func NewCBusPointToPointToMultiPointCommand( bridgeAddress BridgeAddress , networkRoute NetworkRoute , peekedApplication byte , cBusOptions CBusOptions ) *_CBusPointToPointToMultiPointCommand {
return &_CBusPointToPointToMultiPointCommand{ BridgeAddress: bridgeAddress , NetworkRoute: networkRoute , PeekedApplication: peekedApplication , CBusOptions: cBusOptions }
}

// Deprecated: use the interface for direct cast
func CastCBusPointToPointToMultiPointCommand(structType interface{}) CBusPointToPointToMultiPointCommand {
    if casted, ok := structType.(CBusPointToPointToMultiPointCommand); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointToMultiPointCommand); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointToMultiPointCommand) GetTypeName() string {
	return "CBusPointToPointToMultiPointCommand"
}


func (m *_CBusPointToPointToMultiPointCommand) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (bridgeAddress)
	lengthInBits += m.BridgeAddress.GetLengthInBits(ctx)

	// Simple field (networkRoute)
	lengthInBits += m.NetworkRoute.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusPointToPointToMultiPointCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusPointToPointToMultiPointCommandParse(theBytes []byte, cBusOptions CBusOptions) (CBusPointToPointToMultiPointCommand, error) {
	return CBusPointToPointToMultiPointCommandParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusPointToPointToMultiPointCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (CBusPointToPointToMultiPointCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointToMultiPointCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointToMultiPointCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bridgeAddress)
	if pullErr := readBuffer.PullContext("bridgeAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bridgeAddress")
	}
_bridgeAddress, _bridgeAddressErr := BridgeAddressParseWithBuffer(ctx, readBuffer)
	if _bridgeAddressErr != nil {
		return nil, errors.Wrap(_bridgeAddressErr, "Error parsing 'bridgeAddress' field of CBusPointToPointToMultiPointCommand")
	}
	bridgeAddress := _bridgeAddress.(BridgeAddress)
	if closeErr := readBuffer.CloseContext("bridgeAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bridgeAddress")
	}

	// Simple Field (networkRoute)
	if pullErr := readBuffer.PullContext("networkRoute"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkRoute")
	}
_networkRoute, _networkRouteErr := NetworkRouteParseWithBuffer(ctx, readBuffer)
	if _networkRouteErr != nil {
		return nil, errors.Wrap(_networkRouteErr, "Error parsing 'networkRoute' field of CBusPointToPointToMultiPointCommand")
	}
	networkRoute := _networkRoute.(NetworkRoute)
	if closeErr := readBuffer.CloseContext("networkRoute"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkRoute")
	}

				// Peek Field (peekedApplication)
				currentPos = positionAware.GetPos()
				peekedApplication, _err := readBuffer.ReadByte("peekedApplication")
				if _err != nil {
					return nil, errors.Wrap(_err, "Error parsing 'peekedApplication' field of CBusPointToPointToMultiPointCommand")
				}

				readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CBusPointToPointToMultiPointCommandChildSerializeRequirement interface {
		CBusPointToPointToMultiPointCommand
		InitializeParent(CBusPointToPointToMultiPointCommand,  BridgeAddress, NetworkRoute, byte)
		GetParent() CBusPointToPointToMultiPointCommand
	}
	var _childTemp interface{}
	var _child CBusPointToPointToMultiPointCommandChildSerializeRequirement
	var typeSwitchError error
	switch {
case peekedApplication == 0xFF : // CBusPointToPointToMultiPointCommandStatus
		_childTemp, typeSwitchError = CBusPointToPointToMultiPointCommandStatusParseWithBuffer(ctx, readBuffer, cBusOptions)
case 0==0 : // CBusPointToPointToMultiPointCommandNormal
		_childTemp, typeSwitchError = CBusPointToPointToMultiPointCommandNormalParseWithBuffer(ctx, readBuffer, cBusOptions)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedApplication=%v]", peekedApplication)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CBusPointToPointToMultiPointCommand")
	}
	_child = _childTemp.(CBusPointToPointToMultiPointCommandChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("CBusPointToPointToMultiPointCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointToMultiPointCommand")
	}

	// Finish initializing
_child.InitializeParent(_child , bridgeAddress , networkRoute , peekedApplication )
	return _child, nil
}

func (pm *_CBusPointToPointToMultiPointCommand) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CBusPointToPointToMultiPointCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("CBusPointToPointToMultiPointCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusPointToPointToMultiPointCommand")
	}

	// Simple Field (bridgeAddress)
	if pushErr := writeBuffer.PushContext("bridgeAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for bridgeAddress")
	}
	_bridgeAddressErr := writeBuffer.WriteSerializable(ctx, m.GetBridgeAddress())
	if popErr := writeBuffer.PopContext("bridgeAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for bridgeAddress")
	}
	if _bridgeAddressErr != nil {
		return errors.Wrap(_bridgeAddressErr, "Error serializing 'bridgeAddress' field")
	}

	// Simple Field (networkRoute)
	if pushErr := writeBuffer.PushContext("networkRoute"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for networkRoute")
	}
	_networkRouteErr := writeBuffer.WriteSerializable(ctx, m.GetNetworkRoute())
	if popErr := writeBuffer.PopContext("networkRoute"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for networkRoute")
	}
	if _networkRouteErr != nil {
		return errors.Wrap(_networkRouteErr, "Error serializing 'networkRoute' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CBusPointToPointToMultiPointCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusPointToPointToMultiPointCommand")
	}
	return nil
}


////
// Arguments Getter

func (m *_CBusPointToPointToMultiPointCommand) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}
//
////

func (m *_CBusPointToPointToMultiPointCommand) isCBusPointToPointToMultiPointCommand() bool {
	return true
}

func (m *_CBusPointToPointToMultiPointCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



