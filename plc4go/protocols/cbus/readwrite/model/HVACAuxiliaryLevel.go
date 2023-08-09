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


// HVACAuxiliaryLevel is the corresponding interface of HVACAuxiliaryLevel
type HVACAuxiliaryLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetFanMode returns FanMode (property field)
	GetFanMode() bool
	// GetMode returns Mode (property field)
	GetMode() uint8
	// GetIsFanModeAutomatic returns IsFanModeAutomatic (virtual field)
	GetIsFanModeAutomatic() bool
	// GetIsFanModeContinuous returns IsFanModeContinuous (virtual field)
	GetIsFanModeContinuous() bool
	// GetIsFanSpeedAtDefaultSpeed returns IsFanSpeedAtDefaultSpeed (virtual field)
	GetIsFanSpeedAtDefaultSpeed() bool
	// GetSpeedSettings returns SpeedSettings (virtual field)
	GetSpeedSettings() uint8
}

// HVACAuxiliaryLevelExactly can be used when we want exactly this type and not a type which fulfills HVACAuxiliaryLevel.
// This is useful for switch cases.
type HVACAuxiliaryLevelExactly interface {
	HVACAuxiliaryLevel
	isHVACAuxiliaryLevel() bool
}

// _HVACAuxiliaryLevel is the data-structure of this message
type _HVACAuxiliaryLevel struct {
        FanMode bool
        Mode uint8
	// Reserved Fields
	reservedField0 *bool
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACAuxiliaryLevel) GetFanMode() bool {
	return m.FanMode
}

func (m *_HVACAuxiliaryLevel) GetMode() uint8 {
	return m.Mode
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACAuxiliaryLevel) GetIsFanModeAutomatic() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetFanMode()))
}

func (m *_HVACAuxiliaryLevel) GetIsFanModeContinuous() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetFanMode())
}

func (m *_HVACAuxiliaryLevel) GetIsFanSpeedAtDefaultSpeed() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMode()) == (0x00)))
}

func (m *_HVACAuxiliaryLevel) GetSpeedSettings() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetMode())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewHVACAuxiliaryLevel factory function for _HVACAuxiliaryLevel
func NewHVACAuxiliaryLevel( fanMode bool , mode uint8 ) *_HVACAuxiliaryLevel {
return &_HVACAuxiliaryLevel{ FanMode: fanMode , Mode: mode }
}

// Deprecated: use the interface for direct cast
func CastHVACAuxiliaryLevel(structType any) HVACAuxiliaryLevel {
    if casted, ok := structType.(HVACAuxiliaryLevel); ok {
		return casted
	}
	if casted, ok := structType.(*HVACAuxiliaryLevel); ok {
		return *casted
	}
	return nil
}

func (m *_HVACAuxiliaryLevel) GetTypeName() string {
	return "HVACAuxiliaryLevel"
}

func (m *_HVACAuxiliaryLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (fanMode)
	lengthInBits += 1;

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (mode)
	lengthInBits += 6;

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_HVACAuxiliaryLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACAuxiliaryLevelParse(ctx context.Context, theBytes []byte) (HVACAuxiliaryLevel, error) {
	return HVACAuxiliaryLevelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACAuxiliaryLevelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACAuxiliaryLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HVACAuxiliaryLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACAuxiliaryLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of HVACAuxiliaryLevel")
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

	// Simple Field (fanMode)
_fanMode, _fanModeErr := readBuffer.ReadBit("fanMode")
	if _fanModeErr != nil {
		return nil, errors.Wrap(_fanModeErr, "Error parsing 'fanMode' field of HVACAuxiliaryLevel")
	}
	fanMode := _fanMode

	// Virtual field
	_isFanModeAutomatic := !(fanMode)
	isFanModeAutomatic := bool(_isFanModeAutomatic)
	_ = isFanModeAutomatic

	// Virtual field
	_isFanModeContinuous := fanMode
	isFanModeContinuous := bool(_isFanModeContinuous)
	_ = isFanModeContinuous

	// Simple Field (mode)
_mode, _modeErr := readBuffer.ReadUint8("mode", 6)
	if _modeErr != nil {
		return nil, errors.Wrap(_modeErr, "Error parsing 'mode' field of HVACAuxiliaryLevel")
	}
	mode := _mode

	// Virtual field
	_isFanSpeedAtDefaultSpeed := bool((mode) == (0x00))
	isFanSpeedAtDefaultSpeed := bool(_isFanSpeedAtDefaultSpeed)
	_ = isFanSpeedAtDefaultSpeed

	// Virtual field
	_speedSettings := mode
	speedSettings := uint8(_speedSettings)
	_ = speedSettings

	if closeErr := readBuffer.CloseContext("HVACAuxiliaryLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACAuxiliaryLevel")
	}

	// Create the instance
	return &_HVACAuxiliaryLevel{
			FanMode: fanMode,
			Mode: mode,
			reservedField0: reservedField0,
		}, nil
}

func (m *_HVACAuxiliaryLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACAuxiliaryLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("HVACAuxiliaryLevel"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACAuxiliaryLevel")
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

	// Simple Field (fanMode)
	fanMode := bool(m.GetFanMode())
	_fanModeErr := writeBuffer.WriteBit("fanMode", (fanMode))
	if _fanModeErr != nil {
		return errors.Wrap(_fanModeErr, "Error serializing 'fanMode' field")
	}
	// Virtual field
	isFanModeAutomatic := m.GetIsFanModeAutomatic()
	_ = isFanModeAutomatic
	if _isFanModeAutomaticErr := writeBuffer.WriteVirtual(ctx, "isFanModeAutomatic", m.GetIsFanModeAutomatic()); _isFanModeAutomaticErr != nil {
		return errors.Wrap(_isFanModeAutomaticErr, "Error serializing 'isFanModeAutomatic' field")
	}
	// Virtual field
	isFanModeContinuous := m.GetIsFanModeContinuous()
	_ = isFanModeContinuous
	if _isFanModeContinuousErr := writeBuffer.WriteVirtual(ctx, "isFanModeContinuous", m.GetIsFanModeContinuous()); _isFanModeContinuousErr != nil {
		return errors.Wrap(_isFanModeContinuousErr, "Error serializing 'isFanModeContinuous' field")
	}

	// Simple Field (mode)
	mode := uint8(m.GetMode())
	_modeErr := writeBuffer.WriteUint8("mode", 6, (mode))
	if _modeErr != nil {
		return errors.Wrap(_modeErr, "Error serializing 'mode' field")
	}
	// Virtual field
	isFanSpeedAtDefaultSpeed := m.GetIsFanSpeedAtDefaultSpeed()
	_ = isFanSpeedAtDefaultSpeed
	if _isFanSpeedAtDefaultSpeedErr := writeBuffer.WriteVirtual(ctx, "isFanSpeedAtDefaultSpeed", m.GetIsFanSpeedAtDefaultSpeed()); _isFanSpeedAtDefaultSpeedErr != nil {
		return errors.Wrap(_isFanSpeedAtDefaultSpeedErr, "Error serializing 'isFanSpeedAtDefaultSpeed' field")
	}
	// Virtual field
	speedSettings := m.GetSpeedSettings()
	_ = speedSettings
	if _speedSettingsErr := writeBuffer.WriteVirtual(ctx, "speedSettings", m.GetSpeedSettings()); _speedSettingsErr != nil {
		return errors.Wrap(_speedSettingsErr, "Error serializing 'speedSettings' field")
	}

	if popErr := writeBuffer.PopContext("HVACAuxiliaryLevel"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACAuxiliaryLevel")
	}
	return nil
}


func (m *_HVACAuxiliaryLevel) isHVACAuxiliaryLevel() bool {
	return true
}

func (m *_HVACAuxiliaryLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



