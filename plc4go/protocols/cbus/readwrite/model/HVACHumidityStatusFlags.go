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


// HVACHumidityStatusFlags is the corresponding interface of HVACHumidityStatusFlags
type HVACHumidityStatusFlags interface {
	fmt.Stringer
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
	// GetDehumidifyingPlant returns DehumidifyingPlant (property field)
	GetDehumidifyingPlant() bool
	// GetHumidifyingPlant returns HumidifyingPlant (property field)
	GetHumidifyingPlant() bool
	// GetIsDamperStateClosed returns IsDamperStateClosed (virtual field)
	GetIsDamperStateClosed() bool
	// GetIsDamperStateOpen returns IsDamperStateOpen (virtual field)
	GetIsDamperStateOpen() bool
}

// HVACHumidityStatusFlagsExactly can be used when we want exactly this type and not a type which fulfills HVACHumidityStatusFlags.
// This is useful for switch cases.
type HVACHumidityStatusFlagsExactly interface {
	HVACHumidityStatusFlags
	isHVACHumidityStatusFlags() bool
}

// _HVACHumidityStatusFlags is the data-structure of this message
type _HVACHumidityStatusFlags struct {
        Expansion bool
        Error bool
        Busy bool
        DamperState bool
        FanActive bool
        DehumidifyingPlant bool
        HumidifyingPlant bool
	// Reserved Fields
	reservedField0 *bool
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACHumidityStatusFlags) GetExpansion() bool {
	return m.Expansion
}

func (m *_HVACHumidityStatusFlags) GetError() bool {
	return m.Error
}

func (m *_HVACHumidityStatusFlags) GetBusy() bool {
	return m.Busy
}

func (m *_HVACHumidityStatusFlags) GetDamperState() bool {
	return m.DamperState
}

func (m *_HVACHumidityStatusFlags) GetFanActive() bool {
	return m.FanActive
}

func (m *_HVACHumidityStatusFlags) GetDehumidifyingPlant() bool {
	return m.DehumidifyingPlant
}

func (m *_HVACHumidityStatusFlags) GetHumidifyingPlant() bool {
	return m.HumidifyingPlant
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACHumidityStatusFlags) GetIsDamperStateClosed() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetDamperState()))
}

func (m *_HVACHumidityStatusFlags) GetIsDamperStateOpen() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetDamperState())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewHVACHumidityStatusFlags factory function for _HVACHumidityStatusFlags
func NewHVACHumidityStatusFlags( expansion bool , error bool , busy bool , damperState bool , fanActive bool , dehumidifyingPlant bool , humidifyingPlant bool ) *_HVACHumidityStatusFlags {
return &_HVACHumidityStatusFlags{ Expansion: expansion , Error: error , Busy: busy , DamperState: damperState , FanActive: fanActive , DehumidifyingPlant: dehumidifyingPlant , HumidifyingPlant: humidifyingPlant }
}

// Deprecated: use the interface for direct cast
func CastHVACHumidityStatusFlags(structType any) HVACHumidityStatusFlags {
    if casted, ok := structType.(HVACHumidityStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*HVACHumidityStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_HVACHumidityStatusFlags) GetTypeName() string {
	return "HVACHumidityStatusFlags"
}

func (m *_HVACHumidityStatusFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (expansion)
	lengthInBits += 1;

	// Simple field (error)
	lengthInBits += 1;

	// Simple field (busy)
	lengthInBits += 1;

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (damperState)
	lengthInBits += 1;

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (fanActive)
	lengthInBits += 1;

	// Simple field (dehumidifyingPlant)
	lengthInBits += 1;

	// Simple field (humidifyingPlant)
	lengthInBits += 1;

	return lengthInBits
}


func (m *_HVACHumidityStatusFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACHumidityStatusFlagsParse(ctx context.Context, theBytes []byte) (HVACHumidityStatusFlags, error) {
	return HVACHumidityStatusFlagsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACHumidityStatusFlagsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACHumidityStatusFlags, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HVACHumidityStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACHumidityStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (expansion)
_expansion, _expansionErr := readBuffer.ReadBit("expansion")
	if _expansionErr != nil {
		return nil, errors.Wrap(_expansionErr, "Error parsing 'expansion' field of HVACHumidityStatusFlags")
	}
	expansion := _expansion

	// Simple Field (error)
_error, _errorErr := readBuffer.ReadBit("error")
	if _errorErr != nil {
		return nil, errors.Wrap(_errorErr, "Error parsing 'error' field of HVACHumidityStatusFlags")
	}
	error := _error

	// Simple Field (busy)
_busy, _busyErr := readBuffer.ReadBit("busy")
	if _busyErr != nil {
		return nil, errors.Wrap(_busyErr, "Error parsing 'busy' field of HVACHumidityStatusFlags")
	}
	busy := _busy

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of HVACHumidityStatusFlags")
		}
		if reserved != bool(false) {
			log.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (damperState)
_damperState, _damperStateErr := readBuffer.ReadBit("damperState")
	if _damperStateErr != nil {
		return nil, errors.Wrap(_damperStateErr, "Error parsing 'damperState' field of HVACHumidityStatusFlags")
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
		return nil, errors.Wrap(_fanActiveErr, "Error parsing 'fanActive' field of HVACHumidityStatusFlags")
	}
	fanActive := _fanActive

	// Simple Field (dehumidifyingPlant)
_dehumidifyingPlant, _dehumidifyingPlantErr := readBuffer.ReadBit("dehumidifyingPlant")
	if _dehumidifyingPlantErr != nil {
		return nil, errors.Wrap(_dehumidifyingPlantErr, "Error parsing 'dehumidifyingPlant' field of HVACHumidityStatusFlags")
	}
	dehumidifyingPlant := _dehumidifyingPlant

	// Simple Field (humidifyingPlant)
_humidifyingPlant, _humidifyingPlantErr := readBuffer.ReadBit("humidifyingPlant")
	if _humidifyingPlantErr != nil {
		return nil, errors.Wrap(_humidifyingPlantErr, "Error parsing 'humidifyingPlant' field of HVACHumidityStatusFlags")
	}
	humidifyingPlant := _humidifyingPlant

	if closeErr := readBuffer.CloseContext("HVACHumidityStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACHumidityStatusFlags")
	}

	// Create the instance
	return &_HVACHumidityStatusFlags{
			Expansion: expansion,
			Error: error,
			Busy: busy,
			DamperState: damperState,
			FanActive: fanActive,
			DehumidifyingPlant: dehumidifyingPlant,
			HumidifyingPlant: humidifyingPlant,
			reservedField0: reservedField0,
		}, nil
}

func (m *_HVACHumidityStatusFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACHumidityStatusFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("HVACHumidityStatusFlags"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACHumidityStatusFlags")
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
			log.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value": reserved,
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
	isDamperStateClosed := m.GetIsDamperStateClosed()
	_ = isDamperStateClosed
	if _isDamperStateClosedErr := writeBuffer.WriteVirtual(ctx, "isDamperStateClosed", m.GetIsDamperStateClosed()); _isDamperStateClosedErr != nil {
		return errors.Wrap(_isDamperStateClosedErr, "Error serializing 'isDamperStateClosed' field")
	}
	// Virtual field
	isDamperStateOpen := m.GetIsDamperStateOpen()
	_ = isDamperStateOpen
	if _isDamperStateOpenErr := writeBuffer.WriteVirtual(ctx, "isDamperStateOpen", m.GetIsDamperStateOpen()); _isDamperStateOpenErr != nil {
		return errors.Wrap(_isDamperStateOpenErr, "Error serializing 'isDamperStateOpen' field")
	}

	// Simple Field (fanActive)
	fanActive := bool(m.GetFanActive())
	_fanActiveErr := writeBuffer.WriteBit("fanActive", (fanActive))
	if _fanActiveErr != nil {
		return errors.Wrap(_fanActiveErr, "Error serializing 'fanActive' field")
	}

	// Simple Field (dehumidifyingPlant)
	dehumidifyingPlant := bool(m.GetDehumidifyingPlant())
	_dehumidifyingPlantErr := writeBuffer.WriteBit("dehumidifyingPlant", (dehumidifyingPlant))
	if _dehumidifyingPlantErr != nil {
		return errors.Wrap(_dehumidifyingPlantErr, "Error serializing 'dehumidifyingPlant' field")
	}

	// Simple Field (humidifyingPlant)
	humidifyingPlant := bool(m.GetHumidifyingPlant())
	_humidifyingPlantErr := writeBuffer.WriteBit("humidifyingPlant", (humidifyingPlant))
	if _humidifyingPlantErr != nil {
		return errors.Wrap(_humidifyingPlantErr, "Error serializing 'humidifyingPlant' field")
	}

	if popErr := writeBuffer.PopContext("HVACHumidityStatusFlags"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACHumidityStatusFlags")
	}
	return nil
}


func (m *_HVACHumidityStatusFlags) isHVACHumidityStatusFlags() bool {
	return true
}

func (m *_HVACHumidityStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



