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

// AirConditioningDataZoneHumidity is the corresponding interface of AirConditioningDataZoneHumidity
type AirConditioningDataZoneHumidity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHumidity returns Humidity (property field)
	GetHumidity() HVACHumidity
	// GetSensorStatus returns SensorStatus (property field)
	GetSensorStatus() HVACSensorStatus
}

// AirConditioningDataZoneHumidityExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataZoneHumidity.
// This is useful for switch cases.
type AirConditioningDataZoneHumidityExactly interface {
	AirConditioningDataZoneHumidity
	isAirConditioningDataZoneHumidity() bool
}

// _AirConditioningDataZoneHumidity is the data-structure of this message
type _AirConditioningDataZoneHumidity struct {
	*_AirConditioningData
	ZoneGroup    byte
	ZoneList     HVACZoneList
	Humidity     HVACHumidity
	SensorStatus HVACSensorStatus
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataZoneHumidity) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataZoneHumidity) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataZoneHumidity) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataZoneHumidity) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataZoneHumidity) GetHumidity() HVACHumidity {
	return m.Humidity
}

func (m *_AirConditioningDataZoneHumidity) GetSensorStatus() HVACSensorStatus {
	return m.SensorStatus
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataZoneHumidity factory function for _AirConditioningDataZoneHumidity
func NewAirConditioningDataZoneHumidity(zoneGroup byte, zoneList HVACZoneList, humidity HVACHumidity, sensorStatus HVACSensorStatus, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataZoneHumidity {
	_result := &_AirConditioningDataZoneHumidity{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Humidity:             humidity,
		SensorStatus:         sensorStatus,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataZoneHumidity(structType any) AirConditioningDataZoneHumidity {
	if casted, ok := structType.(AirConditioningDataZoneHumidity); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataZoneHumidity); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataZoneHumidity) GetTypeName() string {
	return "AirConditioningDataZoneHumidity"
}

func (m *_AirConditioningDataZoneHumidity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (humidity)
	lengthInBits += m.Humidity.GetLengthInBits(ctx)

	// Simple field (sensorStatus)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataZoneHumidity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningDataZoneHumidityParse(ctx context.Context, theBytes []byte) (AirConditioningDataZoneHumidity, error) {
	return AirConditioningDataZoneHumidityParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningDataZoneHumidityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataZoneHumidity, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AirConditioningDataZoneHumidity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataZoneHumidity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
	_zoneGroup, _zoneGroupErr := readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataZoneHumidity")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
	_zoneList, _zoneListErr := HVACZoneListParseWithBuffer(ctx, readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataZoneHumidity")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (humidity)
	if pullErr := readBuffer.PullContext("humidity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for humidity")
	}
	_humidity, _humidityErr := HVACHumidityParseWithBuffer(ctx, readBuffer)
	if _humidityErr != nil {
		return nil, errors.Wrap(_humidityErr, "Error parsing 'humidity' field of AirConditioningDataZoneHumidity")
	}
	humidity := _humidity.(HVACHumidity)
	if closeErr := readBuffer.CloseContext("humidity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for humidity")
	}

	// Simple Field (sensorStatus)
	if pullErr := readBuffer.PullContext("sensorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sensorStatus")
	}
	_sensorStatus, _sensorStatusErr := HVACSensorStatusParseWithBuffer(ctx, readBuffer)
	if _sensorStatusErr != nil {
		return nil, errors.Wrap(_sensorStatusErr, "Error parsing 'sensorStatus' field of AirConditioningDataZoneHumidity")
	}
	sensorStatus := _sensorStatus
	if closeErr := readBuffer.CloseContext("sensorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sensorStatus")
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataZoneHumidity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataZoneHumidity")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataZoneHumidity{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Humidity:             humidity,
		SensorStatus:         sensorStatus,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataZoneHumidity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataZoneHumidity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataZoneHumidity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataZoneHumidity")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		// Simple Field (zoneList)
		if pushErr := writeBuffer.PushContext("zoneList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneList")
		}
		_zoneListErr := writeBuffer.WriteSerializable(ctx, m.GetZoneList())
		if popErr := writeBuffer.PopContext("zoneList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneList")
		}
		if _zoneListErr != nil {
			return errors.Wrap(_zoneListErr, "Error serializing 'zoneList' field")
		}

		// Simple Field (humidity)
		if pushErr := writeBuffer.PushContext("humidity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for humidity")
		}
		_humidityErr := writeBuffer.WriteSerializable(ctx, m.GetHumidity())
		if popErr := writeBuffer.PopContext("humidity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for humidity")
		}
		if _humidityErr != nil {
			return errors.Wrap(_humidityErr, "Error serializing 'humidity' field")
		}

		// Simple Field (sensorStatus)
		if pushErr := writeBuffer.PushContext("sensorStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sensorStatus")
		}
		_sensorStatusErr := writeBuffer.WriteSerializable(ctx, m.GetSensorStatus())
		if popErr := writeBuffer.PopContext("sensorStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sensorStatus")
		}
		if _sensorStatusErr != nil {
			return errors.Wrap(_sensorStatusErr, "Error serializing 'sensorStatus' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataZoneHumidity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataZoneHumidity")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataZoneHumidity) isAirConditioningDataZoneHumidity() bool {
	return true
}

func (m *_AirConditioningDataZoneHumidity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
