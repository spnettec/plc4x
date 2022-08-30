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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// AirConditioningDataSetHvacUpperGuardLimit is the corresponding interface of AirConditioningDataSetHvacUpperGuardLimit
type AirConditioningDataSetHvacUpperGuardLimit interface {
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetLimit returns Limit (property field)
	GetLimit() HVACTemperature
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACModeAndFlags
}

// AirConditioningDataSetHvacUpperGuardLimitExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataSetHvacUpperGuardLimit.
// This is useful for switch cases.
type AirConditioningDataSetHvacUpperGuardLimitExactly interface {
	AirConditioningDataSetHvacUpperGuardLimit
	isAirConditioningDataSetHvacUpperGuardLimit() bool
}

// _AirConditioningDataSetHvacUpperGuardLimit is the data-structure of this message
type _AirConditioningDataSetHvacUpperGuardLimit struct {
	*_AirConditioningData
        ZoneGroup byte
        ZoneList HVACZoneList
        Limit HVACTemperature
        HvacModeAndFlags HVACModeAndFlags
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetHvacUpperGuardLimit) InitializeParent(parent AirConditioningData , commandTypeContainer AirConditioningCommandTypeContainer ) {	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit)  GetParent() AirConditioningData {
	return m._AirConditioningData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetLimit() HVACTemperature {
	return m.Limit
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetHvacModeAndFlags() HVACModeAndFlags {
	return m.HvacModeAndFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewAirConditioningDataSetHvacUpperGuardLimit factory function for _AirConditioningDataSetHvacUpperGuardLimit
func NewAirConditioningDataSetHvacUpperGuardLimit( zoneGroup byte , zoneList HVACZoneList , limit HVACTemperature , hvacModeAndFlags HVACModeAndFlags , commandTypeContainer AirConditioningCommandTypeContainer ) *_AirConditioningDataSetHvacUpperGuardLimit {
	_result := &_AirConditioningDataSetHvacUpperGuardLimit{
		ZoneGroup: zoneGroup,
		ZoneList: zoneList,
		Limit: limit,
		HvacModeAndFlags: hvacModeAndFlags,
    	_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetHvacUpperGuardLimit(structType interface{}) AirConditioningDataSetHvacUpperGuardLimit {
    if casted, ok := structType.(AirConditioningDataSetHvacUpperGuardLimit); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetHvacUpperGuardLimit); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetTypeName() string {
	return "AirConditioningDataSetHvacUpperGuardLimit"
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneGroup)
	lengthInBits += 8;

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits()

	// Simple field (limit)
	lengthInBits += m.Limit.GetLengthInBits()

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits()

	return lengthInBits
}


func (m *_AirConditioningDataSetHvacUpperGuardLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AirConditioningDataSetHvacUpperGuardLimitParse(readBuffer utils.ReadBuffer) (AirConditioningDataSetHvacUpperGuardLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetHvacUpperGuardLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetHvacUpperGuardLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
_zoneGroup, _zoneGroupErr := readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataSetHvacUpperGuardLimit")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
_zoneList, _zoneListErr := HVACZoneListParse(readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataSetHvacUpperGuardLimit")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (limit)
	if pullErr := readBuffer.PullContext("limit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for limit")
	}
_limit, _limitErr := HVACTemperatureParse(readBuffer)
	if _limitErr != nil {
		return nil, errors.Wrap(_limitErr, "Error parsing 'limit' field of AirConditioningDataSetHvacUpperGuardLimit")
	}
	limit := _limit.(HVACTemperature)
	if closeErr := readBuffer.CloseContext("limit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for limit")
	}

	// Simple Field (hvacModeAndFlags)
	if pullErr := readBuffer.PullContext("hvacModeAndFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hvacModeAndFlags")
	}
_hvacModeAndFlags, _hvacModeAndFlagsErr := HVACModeAndFlagsParse(readBuffer)
	if _hvacModeAndFlagsErr != nil {
		return nil, errors.Wrap(_hvacModeAndFlagsErr, "Error parsing 'hvacModeAndFlags' field of AirConditioningDataSetHvacUpperGuardLimit")
	}
	hvacModeAndFlags := _hvacModeAndFlags.(HVACModeAndFlags)
	if closeErr := readBuffer.CloseContext("hvacModeAndFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hvacModeAndFlags")
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetHvacUpperGuardLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetHvacUpperGuardLimit")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataSetHvacUpperGuardLimit{
		_AirConditioningData: &_AirConditioningData{
		},
		ZoneGroup: zoneGroup,
		ZoneList: zoneList,
		Limit: limit,
		HvacModeAndFlags: hvacModeAndFlags,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetHvacUpperGuardLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetHvacUpperGuardLimit")
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
	_zoneListErr := writeBuffer.WriteSerializable(m.GetZoneList())
	if popErr := writeBuffer.PopContext("zoneList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for zoneList")
	}
	if _zoneListErr != nil {
		return errors.Wrap(_zoneListErr, "Error serializing 'zoneList' field")
	}

	// Simple Field (limit)
	if pushErr := writeBuffer.PushContext("limit"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for limit")
	}
	_limitErr := writeBuffer.WriteSerializable(m.GetLimit())
	if popErr := writeBuffer.PopContext("limit"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for limit")
	}
	if _limitErr != nil {
		return errors.Wrap(_limitErr, "Error serializing 'limit' field")
	}

	// Simple Field (hvacModeAndFlags)
	if pushErr := writeBuffer.PushContext("hvacModeAndFlags"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for hvacModeAndFlags")
	}
	_hvacModeAndFlagsErr := writeBuffer.WriteSerializable(m.GetHvacModeAndFlags())
	if popErr := writeBuffer.PopContext("hvacModeAndFlags"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for hvacModeAndFlags")
	}
	if _hvacModeAndFlagsErr != nil {
		return errors.Wrap(_hvacModeAndFlagsErr, "Error serializing 'hvacModeAndFlags' field")
	}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetHvacUpperGuardLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetHvacUpperGuardLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_AirConditioningDataSetHvacUpperGuardLimit) isAirConditioningDataSetHvacUpperGuardLimit() bool {
	return true
}

func (m *_AirConditioningDataSetHvacUpperGuardLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



