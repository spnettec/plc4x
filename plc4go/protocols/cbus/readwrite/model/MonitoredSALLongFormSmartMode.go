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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// MonitoredSALLongFormSmartMode is the corresponding interface of MonitoredSALLongFormSmartMode
type MonitoredSALLongFormSmartMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MonitoredSAL
	// GetTerminatingByte returns TerminatingByte (property field)
	GetTerminatingByte() uint32
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetReservedByte returns ReservedByte (property field)
	GetReservedByte() *byte
	// GetReplyNetwork returns ReplyNetwork (property field)
	GetReplyNetwork() ReplyNetwork
	// GetSalData returns SalData (property field)
	GetSalData() SALData
	// GetIsUnitAddress returns IsUnitAddress (virtual field)
	GetIsUnitAddress() bool
}

// MonitoredSALLongFormSmartModeExactly can be used when we want exactly this type and not a type which fulfills MonitoredSALLongFormSmartMode.
// This is useful for switch cases.
type MonitoredSALLongFormSmartModeExactly interface {
	MonitoredSALLongFormSmartMode
	isMonitoredSALLongFormSmartMode() bool
}

// _MonitoredSALLongFormSmartMode is the data-structure of this message
type _MonitoredSALLongFormSmartMode struct {
	*_MonitoredSAL
	TerminatingByte uint32
	UnitAddress     UnitAddress
	BridgeAddress   BridgeAddress
	Application     ApplicationIdContainer
	ReservedByte    *byte
	ReplyNetwork    ReplyNetwork
	SalData         SALData
	// Reserved Fields
	reservedField0 *byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredSALLongFormSmartMode) InitializeParent(parent MonitoredSAL, salType byte) {
	m.SalType = salType
}

func (m *_MonitoredSALLongFormSmartMode) GetParent() MonitoredSAL {
	return m._MonitoredSAL
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredSALLongFormSmartMode) GetTerminatingByte() uint32 {
	return m.TerminatingByte
}

func (m *_MonitoredSALLongFormSmartMode) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

func (m *_MonitoredSALLongFormSmartMode) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_MonitoredSALLongFormSmartMode) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_MonitoredSALLongFormSmartMode) GetReservedByte() *byte {
	return m.ReservedByte
}

func (m *_MonitoredSALLongFormSmartMode) GetReplyNetwork() ReplyNetwork {
	return m.ReplyNetwork
}

func (m *_MonitoredSALLongFormSmartMode) GetSalData() SALData {
	return m.SalData
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MonitoredSALLongFormSmartMode) GetIsUnitAddress() bool {
	ctx := context.Background()
	_ = ctx
	unitAddress := m.UnitAddress
	_ = unitAddress
	bridgeAddress := m.BridgeAddress
	_ = bridgeAddress
	reservedByte := m.ReservedByte
	_ = reservedByte
	replyNetwork := m.ReplyNetwork
	_ = replyNetwork
	salData := m.SalData
	_ = salData
	return bool(bool((m.GetTerminatingByte() & 0xff) == (0x00)))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredSALLongFormSmartMode factory function for _MonitoredSALLongFormSmartMode
func NewMonitoredSALLongFormSmartMode(terminatingByte uint32, unitAddress UnitAddress, bridgeAddress BridgeAddress, application ApplicationIdContainer, reservedByte *byte, replyNetwork ReplyNetwork, salData SALData, salType byte, cBusOptions CBusOptions) *_MonitoredSALLongFormSmartMode {
	_result := &_MonitoredSALLongFormSmartMode{
		TerminatingByte: terminatingByte,
		UnitAddress:     unitAddress,
		BridgeAddress:   bridgeAddress,
		Application:     application,
		ReservedByte:    reservedByte,
		ReplyNetwork:    replyNetwork,
		SalData:         salData,
		_MonitoredSAL:   NewMonitoredSAL(salType, cBusOptions),
	}
	_result._MonitoredSAL._MonitoredSALChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredSALLongFormSmartMode(structType any) MonitoredSALLongFormSmartMode {
	if casted, ok := structType.(MonitoredSALLongFormSmartMode); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredSALLongFormSmartMode); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredSALLongFormSmartMode) GetTypeName() string {
	return "MonitoredSALLongFormSmartMode"
}

func (m *_MonitoredSALLongFormSmartMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Optional Field (unitAddress)
	if m.UnitAddress != nil {
		lengthInBits += m.UnitAddress.GetLengthInBits(ctx)
	}

	// Optional Field (bridgeAddress)
	if m.BridgeAddress != nil {
		lengthInBits += m.BridgeAddress.GetLengthInBits(ctx)
	}

	// Simple field (application)
	lengthInBits += 8

	// Optional Field (reservedByte)
	if m.ReservedByte != nil {
		lengthInBits += 8
	}

	// Optional Field (replyNetwork)
	if m.ReplyNetwork != nil {
		lengthInBits += m.ReplyNetwork.GetLengthInBits(ctx)
	}

	// Optional Field (salData)
	if m.SalData != nil {
		lengthInBits += m.SalData.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_MonitoredSALLongFormSmartMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredSALLongFormSmartModeParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (MonitoredSALLongFormSmartMode, error) {
	return MonitoredSALLongFormSmartModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func MonitoredSALLongFormSmartModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (MonitoredSALLongFormSmartMode, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MonitoredSALLongFormSmartMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredSALLongFormSmartMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *byte
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of MonitoredSALLongFormSmartMode")
		}
		if reserved != byte(0x05) {
			log.Info().Fields(map[string]any{
				"expected value": byte(0x05),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Peek Field (terminatingByte)
	currentPos = positionAware.GetPos()
	terminatingByte, _err := readBuffer.ReadUint32("terminatingByte", 24)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'terminatingByte' field of MonitoredSALLongFormSmartMode")
	}

	readBuffer.Reset(currentPos)

	// Virtual field
	_isUnitAddress := bool((terminatingByte & 0xff) == (0x00))
	isUnitAddress := bool(_isUnitAddress)
	_ = isUnitAddress

	// Optional Field (unitAddress) (Can be skipped, if a given expression evaluates to false)
	var unitAddress UnitAddress = nil
	if isUnitAddress {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("unitAddress"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for unitAddress")
		}
		_val, _err := UnitAddressParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'unitAddress' field of MonitoredSALLongFormSmartMode")
		default:
			unitAddress = _val.(UnitAddress)
			if closeErr := readBuffer.CloseContext("unitAddress"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for unitAddress")
			}
		}
	}

	// Optional Field (bridgeAddress) (Can be skipped, if a given expression evaluates to false)
	var bridgeAddress BridgeAddress = nil
	if !(isUnitAddress) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("bridgeAddress"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for bridgeAddress")
		}
		_val, _err := BridgeAddressParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'bridgeAddress' field of MonitoredSALLongFormSmartMode")
		default:
			bridgeAddress = _val.(BridgeAddress)
			if closeErr := readBuffer.CloseContext("bridgeAddress"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for bridgeAddress")
			}
		}
	}

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
	_application, _applicationErr := ApplicationIdContainerParseWithBuffer(ctx, readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of MonitoredSALLongFormSmartMode")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	// Optional Field (reservedByte) (Can be skipped, if a given expression evaluates to false)
	var reservedByte *byte = nil
	if isUnitAddress {
		_val, _err := readBuffer.ReadByte("reservedByte")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reservedByte' field of MonitoredSALLongFormSmartMode")
		}
		reservedByte = &_val
	}

	// Validation
	if !(bool(bool(isUnitAddress) && bool(bool((*reservedByte) == (0x00)))) || bool(!(isUnitAddress))) {
		return nil, errors.WithStack(utils.ParseValidationError{"invalid unit address"})
	}

	// Optional Field (replyNetwork) (Can be skipped, if a given expression evaluates to false)
	var replyNetwork ReplyNetwork = nil
	if !(isUnitAddress) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("replyNetwork"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for replyNetwork")
		}
		_val, _err := ReplyNetworkParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'replyNetwork' field of MonitoredSALLongFormSmartMode")
		default:
			replyNetwork = _val.(ReplyNetwork)
			if closeErr := readBuffer.CloseContext("replyNetwork"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for replyNetwork")
			}
		}
	}

	// Optional Field (salData) (Can be skipped, if a given expression evaluates to false)
	var salData SALData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("salData"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for salData")
		}
		_val, _err := SALDataParseWithBuffer(ctx, readBuffer, application.ApplicationId())
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'salData' field of MonitoredSALLongFormSmartMode")
		default:
			salData = _val.(SALData)
			if closeErr := readBuffer.CloseContext("salData"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for salData")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("MonitoredSALLongFormSmartMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredSALLongFormSmartMode")
	}

	// Create a partially initialized instance
	_child := &_MonitoredSALLongFormSmartMode{
		_MonitoredSAL: &_MonitoredSAL{
			CBusOptions: cBusOptions,
		},
		TerminatingByte: terminatingByte,
		UnitAddress:     unitAddress,
		BridgeAddress:   bridgeAddress,
		Application:     application,
		ReservedByte:    reservedByte,
		ReplyNetwork:    replyNetwork,
		SalData:         salData,
		reservedField0:  reservedField0,
	}
	_child._MonitoredSAL._MonitoredSALChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredSALLongFormSmartMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredSALLongFormSmartMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredSALLongFormSmartMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredSALLongFormSmartMode")
		}

		// Reserved Field (reserved)
		{
			var reserved byte = byte(0x05)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": byte(0x05),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteByte("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}
		// Virtual field
		isUnitAddress := m.GetIsUnitAddress()
		_ = isUnitAddress
		if _isUnitAddressErr := writeBuffer.WriteVirtual(ctx, "isUnitAddress", m.GetIsUnitAddress()); _isUnitAddressErr != nil {
			return errors.Wrap(_isUnitAddressErr, "Error serializing 'isUnitAddress' field")
		}

		// Optional Field (unitAddress) (Can be skipped, if the value is null)
		var unitAddress UnitAddress = nil
		if m.GetUnitAddress() != nil {
			if pushErr := writeBuffer.PushContext("unitAddress"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for unitAddress")
			}
			unitAddress = m.GetUnitAddress()
			_unitAddressErr := writeBuffer.WriteSerializable(ctx, unitAddress)
			if popErr := writeBuffer.PopContext("unitAddress"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for unitAddress")
			}
			if _unitAddressErr != nil {
				return errors.Wrap(_unitAddressErr, "Error serializing 'unitAddress' field")
			}
		}

		// Optional Field (bridgeAddress) (Can be skipped, if the value is null)
		var bridgeAddress BridgeAddress = nil
		if m.GetBridgeAddress() != nil {
			if pushErr := writeBuffer.PushContext("bridgeAddress"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for bridgeAddress")
			}
			bridgeAddress = m.GetBridgeAddress()
			_bridgeAddressErr := writeBuffer.WriteSerializable(ctx, bridgeAddress)
			if popErr := writeBuffer.PopContext("bridgeAddress"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for bridgeAddress")
			}
			if _bridgeAddressErr != nil {
				return errors.Wrap(_bridgeAddressErr, "Error serializing 'bridgeAddress' field")
			}
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for application")
		}
		_applicationErr := writeBuffer.WriteSerializable(ctx, m.GetApplication())
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for application")
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Optional Field (reservedByte) (Can be skipped, if the value is null)
		var reservedByte *byte = nil
		if m.GetReservedByte() != nil {
			reservedByte = m.GetReservedByte()
			_reservedByteErr := writeBuffer.WriteByte("reservedByte", *(reservedByte))
			if _reservedByteErr != nil {
				return errors.Wrap(_reservedByteErr, "Error serializing 'reservedByte' field")
			}
		}

		// Optional Field (replyNetwork) (Can be skipped, if the value is null)
		var replyNetwork ReplyNetwork = nil
		if m.GetReplyNetwork() != nil {
			if pushErr := writeBuffer.PushContext("replyNetwork"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for replyNetwork")
			}
			replyNetwork = m.GetReplyNetwork()
			_replyNetworkErr := writeBuffer.WriteSerializable(ctx, replyNetwork)
			if popErr := writeBuffer.PopContext("replyNetwork"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for replyNetwork")
			}
			if _replyNetworkErr != nil {
				return errors.Wrap(_replyNetworkErr, "Error serializing 'replyNetwork' field")
			}
		}

		// Optional Field (salData) (Can be skipped, if the value is null)
		var salData SALData = nil
		if m.GetSalData() != nil {
			if pushErr := writeBuffer.PushContext("salData"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for salData")
			}
			salData = m.GetSalData()
			_salDataErr := writeBuffer.WriteSerializable(ctx, salData)
			if popErr := writeBuffer.PopContext("salData"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for salData")
			}
			if _salDataErr != nil {
				return errors.Wrap(_salDataErr, "Error serializing 'salData' field")
			}
		}

		if popErr := writeBuffer.PopContext("MonitoredSALLongFormSmartMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredSALLongFormSmartMode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredSALLongFormSmartMode) isMonitoredSALLongFormSmartMode() bool {
	return true
}

func (m *_MonitoredSALLongFormSmartMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
