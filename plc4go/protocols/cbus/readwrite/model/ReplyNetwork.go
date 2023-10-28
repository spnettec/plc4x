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

// ReplyNetwork is the corresponding interface of ReplyNetwork
type ReplyNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNetworkRoute returns NetworkRoute (property field)
	GetNetworkRoute() NetworkRoute
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
}

// ReplyNetworkExactly can be used when we want exactly this type and not a type which fulfills ReplyNetwork.
// This is useful for switch cases.
type ReplyNetworkExactly interface {
	ReplyNetwork
	isReplyNetwork() bool
}

// _ReplyNetwork is the data-structure of this message
type _ReplyNetwork struct {
	NetworkRoute NetworkRoute
	UnitAddress  UnitAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyNetwork) GetNetworkRoute() NetworkRoute {
	return m.NetworkRoute
}

func (m *_ReplyNetwork) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReplyNetwork factory function for _ReplyNetwork
func NewReplyNetwork(networkRoute NetworkRoute, unitAddress UnitAddress) *_ReplyNetwork {
	return &_ReplyNetwork{NetworkRoute: networkRoute, UnitAddress: unitAddress}
}

// Deprecated: use the interface for direct cast
func CastReplyNetwork(structType any) ReplyNetwork {
	if casted, ok := structType.(ReplyNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyNetwork) GetTypeName() string {
	return "ReplyNetwork"
}

func (m *_ReplyNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (networkRoute)
	lengthInBits += m.NetworkRoute.GetLengthInBits(ctx)

	// Simple field (unitAddress)
	lengthInBits += m.UnitAddress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReplyNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ReplyNetworkParse(ctx context.Context, theBytes []byte) (ReplyNetwork, error) {
	return ReplyNetworkParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ReplyNetworkParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ReplyNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ReplyNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkRoute)
	if pullErr := readBuffer.PullContext("networkRoute"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkRoute")
	}
	_networkRoute, _networkRouteErr := NetworkRouteParseWithBuffer(ctx, readBuffer)
	if _networkRouteErr != nil {
		return nil, errors.Wrap(_networkRouteErr, "Error parsing 'networkRoute' field of ReplyNetwork")
	}
	networkRoute := _networkRoute.(NetworkRoute)
	if closeErr := readBuffer.CloseContext("networkRoute"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkRoute")
	}

	// Simple Field (unitAddress)
	if pullErr := readBuffer.PullContext("unitAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unitAddress")
	}
	_unitAddress, _unitAddressErr := UnitAddressParseWithBuffer(ctx, readBuffer)
	if _unitAddressErr != nil {
		return nil, errors.Wrap(_unitAddressErr, "Error parsing 'unitAddress' field of ReplyNetwork")
	}
	unitAddress := _unitAddress.(UnitAddress)
	if closeErr := readBuffer.CloseContext("unitAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unitAddress")
	}

	if closeErr := readBuffer.CloseContext("ReplyNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyNetwork")
	}

	// Create the instance
	return &_ReplyNetwork{
		NetworkRoute: networkRoute,
		UnitAddress:  unitAddress,
	}, nil
}

func (m *_ReplyNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReplyNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ReplyNetwork"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ReplyNetwork")
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

	// Simple Field (unitAddress)
	if pushErr := writeBuffer.PushContext("unitAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for unitAddress")
	}
	_unitAddressErr := writeBuffer.WriteSerializable(ctx, m.GetUnitAddress())
	if popErr := writeBuffer.PopContext("unitAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for unitAddress")
	}
	if _unitAddressErr != nil {
		return errors.Wrap(_unitAddressErr, "Error serializing 'unitAddress' field")
	}

	if popErr := writeBuffer.PopContext("ReplyNetwork"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ReplyNetwork")
	}
	return nil
}

func (m *_ReplyNetwork) isReplyNetwork() bool {
	return true
}

func (m *_ReplyNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
