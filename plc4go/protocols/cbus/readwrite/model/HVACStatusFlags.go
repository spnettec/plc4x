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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// HVACStatusFlags is the corresponding interface of HVACStatusFlags
type HVACStatusFlags interface {
	utils.LengthAware
	utils.Serializable
	// GetExpansion returns Expansion (property field)
	GetExpansion() bool
	// GetError returns Error (property field)
	GetError() bool
	// GetBusy returns Busy (property field)
	GetBusy() bool
	// GetDamperState returns DamperState (property field)
	GetDamperState() bool
	// GetFanActive returns FanActive (property field)
	GetFanActive() bool
	// GetHeatingPlant returns HeatingPlant (property field)
	GetHeatingPlant() bool
	// GetCoolingPlant returns CoolingPlant (property field)
	GetCoolingPlant() bool
	// GetIsDamperStateClosed returns IsDamperStateClosed (virtual field)
	GetIsDamperStateClosed() bool
	// GetIsDamperStateOpen returns IsDamperStateOpen (virtual field)
	GetIsDamperStateOpen() bool
}

// HVACStatusFlagsExactly can be used when we want exactly this type and not a type which fulfills HVACStatusFlags.
// This is useful for switch cases.
type HVACStatusFlagsExactly interface {
	HVACStatusFlags
	isHVACStatusFlags() bool
}

// _HVACStatusFlags is the data-structure of this message
type _HVACStatusFlags struct {
	Expansion    bool
	Error        bool
	Busy         bool
	DamperState  bool
	FanActive    bool
	HeatingPlant bool
	CoolingPlant bool
	// Reserved Fields
	reservedField0 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACStatusFlags) GetExpansion() bool {
	return m.Expansion
}

func (m *_HVACStatusFlags) GetError() bool {
	return m.Error
}

func (m *_HVACStatusFlags) GetBusy() bool {
	return m.Busy
}

func (m *_HVACStatusFlags) GetDamperState() bool {
	return m.DamperState
}

func (m *_HVACStatusFlags) GetFanActive() bool {
	return m.FanActive
}

func (m *_HVACStatusFlags) GetHeatingPlant() bool {
	return m.HeatingPlant
}

func (m *_HVACStatusFlags) GetCoolingPlant() bool {
	return m.CoolingPlant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACStatusFlags) GetIsDamperStateClosed() bool {
	return bool(!(m.GetDamperState()))
}

func (m *_HVACStatusFlags) GetIsDamperStateOpen() bool {
	return bool(m.GetDamperState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHVACStatusFlags factory function for _HVACStatusFlags
func NewHVACStatusFlags(expansion bool, error bool, busy bool, damperState bool, fanActive bool, heatingPlant bool, coolingPlant bool) *_HVACStatusFlags {
	return &_HVACStatusFlags{Expansion: expansion, Error: error, Busy: busy, DamperState: damperState, FanActive: fanActive, HeatingPlant: heatingPlant, CoolingPlant: coolingPlant}
}

// Deprecated: use the interface for direct cast
func CastHVACStatusFlags(structType interface{}) HVACStatusFlags {
	if casted, ok := structType.(HVACStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*HVACStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_HVACStatusFlags) GetTypeName() string {
	return "HVACStatusFlags"
}

func (m *_HVACStatusFlags) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_HVACStatusFlags) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (expansion)
	lengthInBits += 1

	// Simple field (error)
	lengthInBits += 1

	// Simple field (busy)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (damperState)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (fanActive)
	lengthInBits += 1

	// Simple field (heatingPlant)
	lengthInBits += 1

	// Simple field (coolingPlant)
	lengthInBits += 1

	return lengthInBits
}

func (m *_HVACStatusFlags) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func HVACStatusFlagsParse(readBuffer utils.ReadBuffer) (HVACStatusFlags, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (expansion)
	_expansion, _expansionErr := readBuffer.ReadBit("expansion")
	if _expansionErr != nil {
		return nil, errors.Wrap(_expansionErr, "Error parsing 'expansion' field of HVACStatusFlags")
	}
	expansion := _expansion

	// Simple Field (error)
	_error, _errorErr := readBuffer.ReadBit("error")
	if _errorErr != nil {
		return nil, errors.Wrap(_errorErr, "Error parsing 'error' field of HVACStatusFlags")
	}
	error := _error

	// Simple Field (busy)
	_busy, _busyErr := readBuffer.ReadBit("busy")
	if _busyErr != nil {
		return nil, errors.Wrap(_busyErr, "Error parsing 'busy' field of HVACStatusFlags")
	}
	busy := _busy

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of HVACStatusFlags")
		}
		if reserved != bool(false) {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (damperState)
	_damperState, _damperStateErr := readBuffer.ReadBit("damperState")
	if _damperStateErr != nil {
		return nil, errors.Wrap(_damperStateErr, "Error parsing 'damperState' field of HVACStatusFlags")
	}
	damperState := _damperState

	// Virtual field
	_isDamperStateClosed := !(damperState)
	isDamperStateClosed := bool(_isDamperStateClosed)
	_ = isDamperStateClosed

	// Virtual field
	_isDamperStateOpen := damperState
	isDamperStateOpen := bool(_isDamperStateOpen)
	_ = isDamperStateOpen

	// Simple Field (fanActive)
	_fanActive, _fanActiveErr := readBuffer.ReadBit("fanActive")
	if _fanActiveErr != nil {
		return nil, errors.Wrap(_fanActiveErr, "Error parsing 'fanActive' field of HVACStatusFlags")
	}
	fanActive := _fanActive

	// Simple Field (heatingPlant)
	_heatingPlant, _heatingPlantErr := readBuffer.ReadBit("heatingPlant")
	if _heatingPlantErr != nil {
		return nil, errors.Wrap(_heatingPlantErr, "Error parsing 'heatingPlant' field of HVACStatusFlags")
	}
	heatingPlant := _heatingPlant

	// Simple Field (coolingPlant)
	_coolingPlant, _coolingPlantErr := readBuffer.ReadBit("coolingPlant")
	if _coolingPlantErr != nil {
		return nil, errors.Wrap(_coolingPlantErr, "Error parsing 'coolingPlant' field of HVACStatusFlags")
	}
	coolingPlant := _coolingPlant

	if closeErr := readBuffer.CloseContext("HVACStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACStatusFlags")
	}

	// Create the instance
	return &_HVACStatusFlags{
		Expansion:      expansion,
		Error:          error,
		Busy:           busy,
		DamperState:    damperState,
		FanActive:      fanActive,
		HeatingPlant:   heatingPlant,
		CoolingPlant:   coolingPlant,
		reservedField0: reservedField0,
	}, nil
}

func (m *_HVACStatusFlags) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("HVACStatusFlags"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACStatusFlags")
	}

	// Simple Field (expansion)
	expansion := bool(m.GetExpansion())
	_expansionErr := writeBuffer.WriteBit("expansion", (expansion))
	if _expansionErr != nil {
		return errors.Wrap(_expansionErr, "Error serializing 'expansion' field")
	}

	// Simple Field (error)
	error := bool(m.GetError())
	_errorErr := writeBuffer.WriteBit("error", (error))
	if _errorErr != nil {
		return errors.Wrap(_errorErr, "Error serializing 'error' field")
	}

	// Simple Field (busy)
	busy := bool(m.GetBusy())
	_busyErr := writeBuffer.WriteBit("busy", (busy))
	if _busyErr != nil {
		return errors.Wrap(_busyErr, "Error serializing 'busy' field")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (damperState)
	damperState := bool(m.GetDamperState())
	_damperStateErr := writeBuffer.WriteBit("damperState", (damperState))
	if _damperStateErr != nil {
		return errors.Wrap(_damperStateErr, "Error serializing 'damperState' field")
	}
	// Virtual field
	if _isDamperStateClosedErr := writeBuffer.WriteVirtual("isDamperStateClosed", m.GetIsDamperStateClosed()); _isDamperStateClosedErr != nil {
		return errors.Wrap(_isDamperStateClosedErr, "Error serializing 'isDamperStateClosed' field")
	}
	// Virtual field
	if _isDamperStateOpenErr := writeBuffer.WriteVirtual("isDamperStateOpen", m.GetIsDamperStateOpen()); _isDamperStateOpenErr != nil {
		return errors.Wrap(_isDamperStateOpenErr, "Error serializing 'isDamperStateOpen' field")
	}

	// Simple Field (fanActive)
	fanActive := bool(m.GetFanActive())
	_fanActiveErr := writeBuffer.WriteBit("fanActive", (fanActive))
	if _fanActiveErr != nil {
		return errors.Wrap(_fanActiveErr, "Error serializing 'fanActive' field")
	}

	// Simple Field (heatingPlant)
	heatingPlant := bool(m.GetHeatingPlant())
	_heatingPlantErr := writeBuffer.WriteBit("heatingPlant", (heatingPlant))
	if _heatingPlantErr != nil {
		return errors.Wrap(_heatingPlantErr, "Error serializing 'heatingPlant' field")
	}

	// Simple Field (coolingPlant)
	coolingPlant := bool(m.GetCoolingPlant())
	_coolingPlantErr := writeBuffer.WriteBit("coolingPlant", (coolingPlant))
	if _coolingPlantErr != nil {
		return errors.Wrap(_coolingPlantErr, "Error serializing 'coolingPlant' field")
	}

	if popErr := writeBuffer.PopContext("HVACStatusFlags"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACStatusFlags")
	}
	return nil
}

func (m *_HVACStatusFlags) isHVACStatusFlags() bool {
	return true
}

func (m *_HVACStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
