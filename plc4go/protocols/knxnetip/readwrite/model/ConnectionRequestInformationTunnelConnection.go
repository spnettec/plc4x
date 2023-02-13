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

// ConnectionRequestInformationTunnelConnection is the corresponding interface of ConnectionRequestInformationTunnelConnection
type ConnectionRequestInformationTunnelConnection interface {
	utils.LengthAware
	utils.Serializable
	ConnectionRequestInformation
	// GetKnxLayer returns KnxLayer (property field)
	GetKnxLayer() KnxLayer
}

// ConnectionRequestInformationTunnelConnectionExactly can be used when we want exactly this type and not a type which fulfills ConnectionRequestInformationTunnelConnection.
// This is useful for switch cases.
type ConnectionRequestInformationTunnelConnectionExactly interface {
	ConnectionRequestInformationTunnelConnection
	isConnectionRequestInformationTunnelConnection() bool
}

// _ConnectionRequestInformationTunnelConnection is the data-structure of this message
type _ConnectionRequestInformationTunnelConnection struct {
	*_ConnectionRequestInformation
	KnxLayer KnxLayer
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionRequestInformationTunnelConnection) GetConnectionType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionRequestInformationTunnelConnection) InitializeParent(parent ConnectionRequestInformation) {
}

func (m *_ConnectionRequestInformationTunnelConnection) GetParent() ConnectionRequestInformation {
	return m._ConnectionRequestInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionRequestInformationTunnelConnection) GetKnxLayer() KnxLayer {
	return m.KnxLayer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConnectionRequestInformationTunnelConnection factory function for _ConnectionRequestInformationTunnelConnection
func NewConnectionRequestInformationTunnelConnection(knxLayer KnxLayer) *_ConnectionRequestInformationTunnelConnection {
	_result := &_ConnectionRequestInformationTunnelConnection{
		KnxLayer:                      knxLayer,
		_ConnectionRequestInformation: NewConnectionRequestInformation(),
	}
	_result._ConnectionRequestInformation._ConnectionRequestInformationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectionRequestInformationTunnelConnection(structType interface{}) ConnectionRequestInformationTunnelConnection {
	if casted, ok := structType.(ConnectionRequestInformationTunnelConnection); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionRequestInformationTunnelConnection); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionRequestInformationTunnelConnection) GetTypeName() string {
	return "ConnectionRequestInformationTunnelConnection"
}

func (m *_ConnectionRequestInformationTunnelConnection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (knxLayer)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ConnectionRequestInformationTunnelConnection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConnectionRequestInformationTunnelConnectionParse(theBytes []byte) (ConnectionRequestInformationTunnelConnection, error) {
	return ConnectionRequestInformationTunnelConnectionParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ConnectionRequestInformationTunnelConnectionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectionRequestInformationTunnelConnection, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequestInformationTunnelConnection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequestInformationTunnelConnection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (knxLayer)
	if pullErr := readBuffer.PullContext("knxLayer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for knxLayer")
	}
	_knxLayer, _knxLayerErr := KnxLayerParseWithBuffer(ctx, readBuffer)
	if _knxLayerErr != nil {
		return nil, errors.Wrap(_knxLayerErr, "Error parsing 'knxLayer' field of ConnectionRequestInformationTunnelConnection")
	}
	knxLayer := _knxLayer
	if closeErr := readBuffer.CloseContext("knxLayer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for knxLayer")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ConnectionRequestInformationTunnelConnection")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("ConnectionRequestInformationTunnelConnection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequestInformationTunnelConnection")
	}

	// Create a partially initialized instance
	_child := &_ConnectionRequestInformationTunnelConnection{
		_ConnectionRequestInformation: &_ConnectionRequestInformation{},
		KnxLayer:                      knxLayer,
		reservedField0:                reservedField0,
	}
	_child._ConnectionRequestInformation._ConnectionRequestInformationChildRequirements = _child
	return _child, nil
}

func (m *_ConnectionRequestInformationTunnelConnection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionRequestInformationTunnelConnection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionRequestInformationTunnelConnection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionRequestInformationTunnelConnection")
		}

		// Simple Field (knxLayer)
		if pushErr := writeBuffer.PushContext("knxLayer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for knxLayer")
		}
		_knxLayerErr := writeBuffer.WriteSerializable(ctx, m.GetKnxLayer())
		if popErr := writeBuffer.PopContext("knxLayer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for knxLayer")
		}
		if _knxLayerErr != nil {
			return errors.Wrap(_knxLayerErr, "Error serializing 'knxLayer' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("ConnectionRequestInformationTunnelConnection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionRequestInformationTunnelConnection")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionRequestInformationTunnelConnection) isConnectionRequestInformationTunnelConnection() bool {
	return true
}

func (m *_ConnectionRequestInformationTunnelConnection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
