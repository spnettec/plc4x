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

// TemperatureBroadcastData is the corresponding interface of TemperatureBroadcastData
type TemperatureBroadcastData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() TemperatureBroadcastCommandTypeContainer
	// GetTemperatureGroup returns TemperatureGroup (property field)
	GetTemperatureGroup() byte
	// GetTemperatureByte returns TemperatureByte (property field)
	GetTemperatureByte() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() TemperatureBroadcastCommandType
	// GetTemperatureInCelsius returns TemperatureInCelsius (virtual field)
	GetTemperatureInCelsius() float32
}

// TemperatureBroadcastDataExactly can be used when we want exactly this type and not a type which fulfills TemperatureBroadcastData.
// This is useful for switch cases.
type TemperatureBroadcastDataExactly interface {
	TemperatureBroadcastData
	isTemperatureBroadcastData() bool
}

// _TemperatureBroadcastData is the data-structure of this message
type _TemperatureBroadcastData struct {
	CommandTypeContainer TemperatureBroadcastCommandTypeContainer
	TemperatureGroup     byte
	TemperatureByte      byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TemperatureBroadcastData) GetCommandTypeContainer() TemperatureBroadcastCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_TemperatureBroadcastData) GetTemperatureGroup() byte {
	return m.TemperatureGroup
}

func (m *_TemperatureBroadcastData) GetTemperatureByte() byte {
	return m.TemperatureByte
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_TemperatureBroadcastData) GetCommandType() TemperatureBroadcastCommandType {
	ctx := context.Background()
	_ = ctx
	return CastTemperatureBroadcastCommandType(m.GetCommandTypeContainer().CommandType())
}

func (m *_TemperatureBroadcastData) GetTemperatureInCelsius() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(float32(m.GetTemperatureByte()) / float32(float32(4)))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTemperatureBroadcastData factory function for _TemperatureBroadcastData
func NewTemperatureBroadcastData(commandTypeContainer TemperatureBroadcastCommandTypeContainer, temperatureGroup byte, temperatureByte byte) *_TemperatureBroadcastData {
	return &_TemperatureBroadcastData{CommandTypeContainer: commandTypeContainer, TemperatureGroup: temperatureGroup, TemperatureByte: temperatureByte}
}

// Deprecated: use the interface for direct cast
func CastTemperatureBroadcastData(structType any) TemperatureBroadcastData {
	if casted, ok := structType.(TemperatureBroadcastData); ok {
		return casted
	}
	if casted, ok := structType.(*TemperatureBroadcastData); ok {
		return *casted
	}
	return nil
}

func (m *_TemperatureBroadcastData) GetTypeName() string {
	return "TemperatureBroadcastData"
}

func (m *_TemperatureBroadcastData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (temperatureGroup)
	lengthInBits += 8

	// Simple field (temperatureByte)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_TemperatureBroadcastData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TemperatureBroadcastDataParse(ctx context.Context, theBytes []byte) (TemperatureBroadcastData, error) {
	return TemperatureBroadcastDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TemperatureBroadcastDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TemperatureBroadcastData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TemperatureBroadcastData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TemperatureBroadcastData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsTemperatureBroadcastCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{"no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := TemperatureBroadcastCommandTypeContainerParseWithBuffer(ctx, readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of TemperatureBroadcastData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := TemperatureBroadcastCommandType(_commandType)
	_ = commandType

	// Simple Field (temperatureGroup)
	_temperatureGroup, _temperatureGroupErr := readBuffer.ReadByte("temperatureGroup")
	if _temperatureGroupErr != nil {
		return nil, errors.Wrap(_temperatureGroupErr, "Error parsing 'temperatureGroup' field of TemperatureBroadcastData")
	}
	temperatureGroup := _temperatureGroup

	// Simple Field (temperatureByte)
	_temperatureByte, _temperatureByteErr := readBuffer.ReadByte("temperatureByte")
	if _temperatureByteErr != nil {
		return nil, errors.Wrap(_temperatureByteErr, "Error parsing 'temperatureByte' field of TemperatureBroadcastData")
	}
	temperatureByte := _temperatureByte

	// Virtual field
	_temperatureInCelsius := float32(temperatureByte) / float32(float32(4))
	temperatureInCelsius := float32(_temperatureInCelsius)
	_ = temperatureInCelsius

	if closeErr := readBuffer.CloseContext("TemperatureBroadcastData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TemperatureBroadcastData")
	}

	// Create the instance
	return &_TemperatureBroadcastData{
		CommandTypeContainer: commandTypeContainer,
		TemperatureGroup:     temperatureGroup,
		TemperatureByte:      temperatureByte,
	}, nil
}

func (m *_TemperatureBroadcastData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TemperatureBroadcastData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TemperatureBroadcastData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TemperatureBroadcastData")
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandTypeContainer")
	}
	_commandTypeContainerErr := writeBuffer.WriteSerializable(ctx, m.GetCommandTypeContainer())
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandTypeContainer")
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Simple Field (temperatureGroup)
	temperatureGroup := byte(m.GetTemperatureGroup())
	_temperatureGroupErr := writeBuffer.WriteByte("temperatureGroup", (temperatureGroup))
	if _temperatureGroupErr != nil {
		return errors.Wrap(_temperatureGroupErr, "Error serializing 'temperatureGroup' field")
	}

	// Simple Field (temperatureByte)
	temperatureByte := byte(m.GetTemperatureByte())
	_temperatureByteErr := writeBuffer.WriteByte("temperatureByte", (temperatureByte))
	if _temperatureByteErr != nil {
		return errors.Wrap(_temperatureByteErr, "Error serializing 'temperatureByte' field")
	}
	// Virtual field
	if _temperatureInCelsiusErr := writeBuffer.WriteVirtual(ctx, "temperatureInCelsius", m.GetTemperatureInCelsius()); _temperatureInCelsiusErr != nil {
		return errors.Wrap(_temperatureInCelsiusErr, "Error serializing 'temperatureInCelsius' field")
	}

	if popErr := writeBuffer.PopContext("TemperatureBroadcastData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TemperatureBroadcastData")
	}
	return nil
}

func (m *_TemperatureBroadcastData) isTemperatureBroadcastData() bool {
	return true
}

func (m *_TemperatureBroadcastData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
