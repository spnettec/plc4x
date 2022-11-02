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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetLightingCommand is the corresponding interface of BACnetLightingCommand
type BACnetLightingCommand interface {
	utils.LengthAware
	utils.Serializable
	// GetLightningOperation returns LightningOperation (property field)
	GetLightningOperation() BACnetLightingOperationTagged
	// GetTargetLevel returns TargetLevel (property field)
	GetTargetLevel() BACnetContextTagReal
	// GetRampRate returns RampRate (property field)
	GetRampRate() BACnetContextTagReal
	// GetStepIncrement returns StepIncrement (property field)
	GetStepIncrement() BACnetContextTagReal
	// GetFadeTime returns FadeTime (property field)
	GetFadeTime() BACnetContextTagUnsignedInteger
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
}

// BACnetLightingCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetLightingCommand.
// This is useful for switch cases.
type BACnetLightingCommandExactly interface {
	BACnetLightingCommand
	isBACnetLightingCommand() bool
}

// _BACnetLightingCommand is the data-structure of this message
type _BACnetLightingCommand struct {
        LightningOperation BACnetLightingOperationTagged
        TargetLevel BACnetContextTagReal
        RampRate BACnetContextTagReal
        StepIncrement BACnetContextTagReal
        FadeTime BACnetContextTagUnsignedInteger
        Priority BACnetContextTagUnsignedInteger
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLightingCommand) GetLightningOperation() BACnetLightingOperationTagged {
	return m.LightningOperation
}

func (m *_BACnetLightingCommand) GetTargetLevel() BACnetContextTagReal {
	return m.TargetLevel
}

func (m *_BACnetLightingCommand) GetRampRate() BACnetContextTagReal {
	return m.RampRate
}

func (m *_BACnetLightingCommand) GetStepIncrement() BACnetContextTagReal {
	return m.StepIncrement
}

func (m *_BACnetLightingCommand) GetFadeTime() BACnetContextTagUnsignedInteger {
	return m.FadeTime
}

func (m *_BACnetLightingCommand) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetLightingCommand factory function for _BACnetLightingCommand
func NewBACnetLightingCommand( lightningOperation BACnetLightingOperationTagged , targetLevel BACnetContextTagReal , rampRate BACnetContextTagReal , stepIncrement BACnetContextTagReal , fadeTime BACnetContextTagUnsignedInteger , priority BACnetContextTagUnsignedInteger ) *_BACnetLightingCommand {
return &_BACnetLightingCommand{ LightningOperation: lightningOperation , TargetLevel: targetLevel , RampRate: rampRate , StepIncrement: stepIncrement , FadeTime: fadeTime , Priority: priority }
}

// Deprecated: use the interface for direct cast
func CastBACnetLightingCommand(structType interface{}) BACnetLightingCommand {
    if casted, ok := structType.(BACnetLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLightingCommand) GetTypeName() string {
	return "BACnetLightingCommand"
}

func (m *_BACnetLightingCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLightingCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (lightningOperation)
	lengthInBits += m.LightningOperation.GetLengthInBits()

	// Optional Field (targetLevel)
	if m.TargetLevel != nil {
		lengthInBits += m.TargetLevel.GetLengthInBits()
	}

	// Optional Field (rampRate)
	if m.RampRate != nil {
		lengthInBits += m.RampRate.GetLengthInBits()
	}

	// Optional Field (stepIncrement)
	if m.StepIncrement != nil {
		lengthInBits += m.StepIncrement.GetLengthInBits()
	}

	// Optional Field (fadeTime)
	if m.FadeTime != nil {
		lengthInBits += m.FadeTime.GetLengthInBits()
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits()
	}

	return lengthInBits
}


func (m *_BACnetLightingCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLightingCommandParse(readBuffer utils.ReadBuffer) (BACnetLightingCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lightningOperation)
	if pullErr := readBuffer.PullContext("lightningOperation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lightningOperation")
	}
_lightningOperation, _lightningOperationErr := BACnetLightingOperationTaggedParse(readBuffer , uint8( uint8(0) ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _lightningOperationErr != nil {
		return nil, errors.Wrap(_lightningOperationErr, "Error parsing 'lightningOperation' field of BACnetLightingCommand")
	}
	lightningOperation := _lightningOperation.(BACnetLightingOperationTagged)
	if closeErr := readBuffer.CloseContext("lightningOperation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lightningOperation")
	}

	// Optional Field (targetLevel) (Can be skipped, if a given expression evaluates to false)
	var targetLevel BACnetContextTagReal = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("targetLevel"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for targetLevel")
		}
_val, _err := BACnetContextTagParse(readBuffer , uint8(1) , BACnetDataType_REAL )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'targetLevel' field of BACnetLightingCommand")
		default:
			targetLevel = _val.(BACnetContextTagReal)
			if closeErr := readBuffer.CloseContext("targetLevel"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for targetLevel")
			}
		}
	}

	// Optional Field (rampRate) (Can be skipped, if a given expression evaluates to false)
	var rampRate BACnetContextTagReal = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("rampRate"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for rampRate")
		}
_val, _err := BACnetContextTagParse(readBuffer , uint8(2) , BACnetDataType_REAL )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'rampRate' field of BACnetLightingCommand")
		default:
			rampRate = _val.(BACnetContextTagReal)
			if closeErr := readBuffer.CloseContext("rampRate"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for rampRate")
			}
		}
	}

	// Optional Field (stepIncrement) (Can be skipped, if a given expression evaluates to false)
	var stepIncrement BACnetContextTagReal = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("stepIncrement"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for stepIncrement")
		}
_val, _err := BACnetContextTagParse(readBuffer , uint8(3) , BACnetDataType_REAL )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'stepIncrement' field of BACnetLightingCommand")
		default:
			stepIncrement = _val.(BACnetContextTagReal)
			if closeErr := readBuffer.CloseContext("stepIncrement"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for stepIncrement")
			}
		}
	}

	// Optional Field (fadeTime) (Can be skipped, if a given expression evaluates to false)
	var fadeTime BACnetContextTagUnsignedInteger = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("fadeTime"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for fadeTime")
		}
_val, _err := BACnetContextTagParse(readBuffer , uint8(4) , BACnetDataType_UNSIGNED_INTEGER )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'fadeTime' field of BACnetLightingCommand")
		default:
			fadeTime = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("fadeTime"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for fadeTime")
			}
		}
	}

	// Optional Field (priority) (Can be skipped, if a given expression evaluates to false)
	var priority BACnetContextTagUnsignedInteger = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for priority")
		}
_val, _err := BACnetContextTagParse(readBuffer , uint8(5) , BACnetDataType_UNSIGNED_INTEGER )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priority' field of BACnetLightingCommand")
		default:
			priority = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for priority")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLightingCommand")
	}

	// Create the instance
	return &_BACnetLightingCommand{
			LightningOperation: lightningOperation,
			TargetLevel: targetLevel,
			RampRate: rampRate,
			StepIncrement: stepIncrement,
			FadeTime: fadeTime,
			Priority: priority,
		}, nil
}

func (m *_BACnetLightingCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLightingCommand) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetLightingCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLightingCommand")
	}

	// Simple Field (lightningOperation)
	if pushErr := writeBuffer.PushContext("lightningOperation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for lightningOperation")
	}
	_lightningOperationErr := writeBuffer.WriteSerializable(m.GetLightningOperation())
	if popErr := writeBuffer.PopContext("lightningOperation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for lightningOperation")
	}
	if _lightningOperationErr != nil {
		return errors.Wrap(_lightningOperationErr, "Error serializing 'lightningOperation' field")
	}

	// Optional Field (targetLevel) (Can be skipped, if the value is null)
	var targetLevel BACnetContextTagReal = nil
	if m.GetTargetLevel() != nil {
		if pushErr := writeBuffer.PushContext("targetLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for targetLevel")
		}
		targetLevel = m.GetTargetLevel()
		_targetLevelErr := writeBuffer.WriteSerializable(targetLevel)
		if popErr := writeBuffer.PopContext("targetLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for targetLevel")
		}
		if _targetLevelErr != nil {
			return errors.Wrap(_targetLevelErr, "Error serializing 'targetLevel' field")
		}
	}

	// Optional Field (rampRate) (Can be skipped, if the value is null)
	var rampRate BACnetContextTagReal = nil
	if m.GetRampRate() != nil {
		if pushErr := writeBuffer.PushContext("rampRate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for rampRate")
		}
		rampRate = m.GetRampRate()
		_rampRateErr := writeBuffer.WriteSerializable(rampRate)
		if popErr := writeBuffer.PopContext("rampRate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for rampRate")
		}
		if _rampRateErr != nil {
			return errors.Wrap(_rampRateErr, "Error serializing 'rampRate' field")
		}
	}

	// Optional Field (stepIncrement) (Can be skipped, if the value is null)
	var stepIncrement BACnetContextTagReal = nil
	if m.GetStepIncrement() != nil {
		if pushErr := writeBuffer.PushContext("stepIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for stepIncrement")
		}
		stepIncrement = m.GetStepIncrement()
		_stepIncrementErr := writeBuffer.WriteSerializable(stepIncrement)
		if popErr := writeBuffer.PopContext("stepIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for stepIncrement")
		}
		if _stepIncrementErr != nil {
			return errors.Wrap(_stepIncrementErr, "Error serializing 'stepIncrement' field")
		}
	}

	// Optional Field (fadeTime) (Can be skipped, if the value is null)
	var fadeTime BACnetContextTagUnsignedInteger = nil
	if m.GetFadeTime() != nil {
		if pushErr := writeBuffer.PushContext("fadeTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fadeTime")
		}
		fadeTime = m.GetFadeTime()
		_fadeTimeErr := writeBuffer.WriteSerializable(fadeTime)
		if popErr := writeBuffer.PopContext("fadeTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fadeTime")
		}
		if _fadeTimeErr != nil {
			return errors.Wrap(_fadeTimeErr, "Error serializing 'fadeTime' field")
		}
	}

	// Optional Field (priority) (Can be skipped, if the value is null)
	var priority BACnetContextTagUnsignedInteger = nil
	if m.GetPriority() != nil {
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priority")
		}
		priority = m.GetPriority()
		_priorityErr := writeBuffer.WriteSerializable(priority)
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priority")
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetLightingCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLightingCommand")
	}
	return nil
}


func (m *_BACnetLightingCommand) isBACnetLightingCommand() bool {
	return true
}

func (m *_BACnetLightingCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



