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


// NetworkRoute is the corresponding interface of NetworkRoute
type NetworkRoute interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNetworkPCI returns NetworkPCI (property field)
	GetNetworkPCI() NetworkProtocolControlInformation
	// GetAdditionalBridgeAddresses returns AdditionalBridgeAddresses (property field)
	GetAdditionalBridgeAddresses() []BridgeAddress
}

// NetworkRouteExactly can be used when we want exactly this type and not a type which fulfills NetworkRoute.
// This is useful for switch cases.
type NetworkRouteExactly interface {
	NetworkRoute
	isNetworkRoute() bool
}

// _NetworkRoute is the data-structure of this message
type _NetworkRoute struct {
        NetworkPCI NetworkProtocolControlInformation
        AdditionalBridgeAddresses []BridgeAddress
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NetworkRoute) GetNetworkPCI() NetworkProtocolControlInformation {
	return m.NetworkPCI
}

func (m *_NetworkRoute) GetAdditionalBridgeAddresses() []BridgeAddress {
	return m.AdditionalBridgeAddresses
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewNetworkRoute factory function for _NetworkRoute
func NewNetworkRoute( networkPCI NetworkProtocolControlInformation , additionalBridgeAddresses []BridgeAddress ) *_NetworkRoute {
return &_NetworkRoute{ NetworkPCI: networkPCI , AdditionalBridgeAddresses: additionalBridgeAddresses }
}

// Deprecated: use the interface for direct cast
func CastNetworkRoute(structType any) NetworkRoute {
    if casted, ok := structType.(NetworkRoute); ok {
		return casted
	}
	if casted, ok := structType.(*NetworkRoute); ok {
		return *casted
	}
	return nil
}

func (m *_NetworkRoute) GetTypeName() string {
	return "NetworkRoute"
}

func (m *_NetworkRoute) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (networkPCI)
	lengthInBits += m.NetworkPCI.GetLengthInBits(ctx)

	// Array field
	if len(m.AdditionalBridgeAddresses) > 0 {
		for _curItem, element := range m.AdditionalBridgeAddresses {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.AdditionalBridgeAddresses), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_NetworkRoute) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NetworkRouteParse(ctx context.Context, theBytes []byte) (NetworkRoute, error) {
	return NetworkRouteParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NetworkRouteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NetworkRoute, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NetworkRoute"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NetworkRoute")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkPCI)
	if pullErr := readBuffer.PullContext("networkPCI"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkPCI")
	}
_networkPCI, _networkPCIErr := NetworkProtocolControlInformationParseWithBuffer(ctx, readBuffer)
	if _networkPCIErr != nil {
		return nil, errors.Wrap(_networkPCIErr, "Error parsing 'networkPCI' field of NetworkRoute")
	}
	networkPCI := _networkPCI.(NetworkProtocolControlInformation)
	if closeErr := readBuffer.CloseContext("networkPCI"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkPCI")
	}

	// Array field (additionalBridgeAddresses)
	if pullErr := readBuffer.PullContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for additionalBridgeAddresses")
	}
	// Count array
	additionalBridgeAddresses := make([]BridgeAddress, utils.Max(uint16(networkPCI.GetStackDepth()) - uint16(uint16(1)), 0))
	// This happens when the size is set conditional to 0
	if len(additionalBridgeAddresses) == 0 {
		additionalBridgeAddresses = nil
	}
	{
		_numItems := uint16(utils.Max(uint16(networkPCI.GetStackDepth()) - uint16(uint16(1)), 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := BridgeAddressParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'additionalBridgeAddresses' field of NetworkRoute")
			}
			additionalBridgeAddresses[_curItem] = _item.(BridgeAddress)
		}
	}
	if closeErr := readBuffer.CloseContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for additionalBridgeAddresses")
	}

	if closeErr := readBuffer.CloseContext("NetworkRoute"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NetworkRoute")
	}

	// Create the instance
	return &_NetworkRoute{
			NetworkPCI: networkPCI,
			AdditionalBridgeAddresses: additionalBridgeAddresses,
		}, nil
}

func (m *_NetworkRoute) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NetworkRoute) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("NetworkRoute"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NetworkRoute")
	}

	// Simple Field (networkPCI)
	if pushErr := writeBuffer.PushContext("networkPCI"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for networkPCI")
	}
	_networkPCIErr := writeBuffer.WriteSerializable(ctx, m.GetNetworkPCI())
	if popErr := writeBuffer.PopContext("networkPCI"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for networkPCI")
	}
	if _networkPCIErr != nil {
		return errors.Wrap(_networkPCIErr, "Error serializing 'networkPCI' field")
	}

	// Array Field (additionalBridgeAddresses)
	if pushErr := writeBuffer.PushContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for additionalBridgeAddresses")
	}
	for _curItem, _element := range m.GetAdditionalBridgeAddresses() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetAdditionalBridgeAddresses()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'additionalBridgeAddresses' field")
		}
	}
	if popErr := writeBuffer.PopContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for additionalBridgeAddresses")
	}

	if popErr := writeBuffer.PopContext("NetworkRoute"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NetworkRoute")
	}
	return nil
}


func (m *_NetworkRoute) isNetworkRoute() bool {
	return true
}

func (m *_NetworkRoute) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



