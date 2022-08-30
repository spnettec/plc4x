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


// AirConditioningDataZoneHvacPlantStatus is the corresponding interface of AirConditioningDataZoneHvacPlantStatus
type AirConditioningDataZoneHvacPlantStatus interface {
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHvacType returns HvacType (property field)
	GetHvacType() HVACType
	// GetHvacStatus returns HvacStatus (property field)
	GetHvacStatus() HVACStatusFlags
	// GetHvacErrorCode returns HvacErrorCode (property field)
	GetHvacErrorCode() HVACError
}

// AirConditioningDataZoneHvacPlantStatusExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataZoneHvacPlantStatus.
// This is useful for switch cases.
type AirConditioningDataZoneHvacPlantStatusExactly interface {
	AirConditioningDataZoneHvacPlantStatus
	isAirConditioningDataZoneHvacPlantStatus() bool
}

// _AirConditioningDataZoneHvacPlantStatus is the data-structure of this message
type _AirConditioningDataZoneHvacPlantStatus struct {
	*_AirConditioningData
        ZoneGroup byte
        ZoneList HVACZoneList
        HvacType HVACType
        HvacStatus HVACStatusFlags
        HvacErrorCode HVACError
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataZoneHvacPlantStatus) InitializeParent(parent AirConditioningData , commandTypeContainer AirConditioningCommandTypeContainer ) {	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataZoneHvacPlantStatus)  GetParent() AirConditioningData {
	return m._AirConditioningData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataZoneHvacPlantStatus) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetHvacType() HVACType {
	return m.HvacType
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetHvacStatus() HVACStatusFlags {
	return m.HvacStatus
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetHvacErrorCode() HVACError {
	return m.HvacErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewAirConditioningDataZoneHvacPlantStatus factory function for _AirConditioningDataZoneHvacPlantStatus
func NewAirConditioningDataZoneHvacPlantStatus( zoneGroup byte , zoneList HVACZoneList , hvacType HVACType , hvacStatus HVACStatusFlags , hvacErrorCode HVACError , commandTypeContainer AirConditioningCommandTypeContainer ) *_AirConditioningDataZoneHvacPlantStatus {
	_result := &_AirConditioningDataZoneHvacPlantStatus{
		ZoneGroup: zoneGroup,
		ZoneList: zoneList,
		HvacType: hvacType,
		HvacStatus: hvacStatus,
		HvacErrorCode: hvacErrorCode,
    	_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataZoneHvacPlantStatus(structType interface{}) AirConditioningDataZoneHvacPlantStatus {
    if casted, ok := structType.(AirConditioningDataZoneHvacPlantStatus); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataZoneHvacPlantStatus); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetTypeName() string {
	return "AirConditioningDataZoneHvacPlantStatus"
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneGroup)
	lengthInBits += 8;

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits()

	// Simple field (hvacType)
	lengthInBits += 8

	// Simple field (hvacStatus)
	lengthInBits += m.HvacStatus.GetLengthInBits()

	// Simple field (hvacErrorCode)
	lengthInBits += 8

	return lengthInBits
}


func (m *_AirConditioningDataZoneHvacPlantStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AirConditioningDataZoneHvacPlantStatusParse(readBuffer utils.ReadBuffer) (AirConditioningDataZoneHvacPlantStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataZoneHvacPlantStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataZoneHvacPlantStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
_zoneGroup, _zoneGroupErr := readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataZoneHvacPlantStatus")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
_zoneList, _zoneListErr := HVACZoneListParse(readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataZoneHvacPlantStatus")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (hvacType)
	if pullErr := readBuffer.PullContext("hvacType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hvacType")
	}
_hvacType, _hvacTypeErr := HVACTypeParse(readBuffer)
	if _hvacTypeErr != nil {
		return nil, errors.Wrap(_hvacTypeErr, "Error parsing 'hvacType' field of AirConditioningDataZoneHvacPlantStatus")
	}
	hvacType := _hvacType
	if closeErr := readBuffer.CloseContext("hvacType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hvacType")
	}

	// Simple Field (hvacStatus)
	if pullErr := readBuffer.PullContext("hvacStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hvacStatus")
	}
_hvacStatus, _hvacStatusErr := HVACStatusFlagsParse(readBuffer)
	if _hvacStatusErr != nil {
		return nil, errors.Wrap(_hvacStatusErr, "Error parsing 'hvacStatus' field of AirConditioningDataZoneHvacPlantStatus")
	}
	hvacStatus := _hvacStatus.(HVACStatusFlags)
	if closeErr := readBuffer.CloseContext("hvacStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hvacStatus")
	}

	// Simple Field (hvacErrorCode)
	if pullErr := readBuffer.PullContext("hvacErrorCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hvacErrorCode")
	}
_hvacErrorCode, _hvacErrorCodeErr := HVACErrorParse(readBuffer)
	if _hvacErrorCodeErr != nil {
		return nil, errors.Wrap(_hvacErrorCodeErr, "Error parsing 'hvacErrorCode' field of AirConditioningDataZoneHvacPlantStatus")
	}
	hvacErrorCode := _hvacErrorCode
	if closeErr := readBuffer.CloseContext("hvacErrorCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hvacErrorCode")
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataZoneHvacPlantStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataZoneHvacPlantStatus")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataZoneHvacPlantStatus{
		_AirConditioningData: &_AirConditioningData{
		},
		ZoneGroup: zoneGroup,
		ZoneList: zoneList,
		HvacType: hvacType,
		HvacStatus: hvacStatus,
		HvacErrorCode: hvacErrorCode,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataZoneHvacPlantStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataZoneHvacPlantStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataZoneHvacPlantStatus")
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

	// Simple Field (hvacType)
	if pushErr := writeBuffer.PushContext("hvacType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for hvacType")
	}
	_hvacTypeErr := writeBuffer.WriteSerializable(m.GetHvacType())
	if popErr := writeBuffer.PopContext("hvacType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for hvacType")
	}
	if _hvacTypeErr != nil {
		return errors.Wrap(_hvacTypeErr, "Error serializing 'hvacType' field")
	}

	// Simple Field (hvacStatus)
	if pushErr := writeBuffer.PushContext("hvacStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for hvacStatus")
	}
	_hvacStatusErr := writeBuffer.WriteSerializable(m.GetHvacStatus())
	if popErr := writeBuffer.PopContext("hvacStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for hvacStatus")
	}
	if _hvacStatusErr != nil {
		return errors.Wrap(_hvacStatusErr, "Error serializing 'hvacStatus' field")
	}

	// Simple Field (hvacErrorCode)
	if pushErr := writeBuffer.PushContext("hvacErrorCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for hvacErrorCode")
	}
	_hvacErrorCodeErr := writeBuffer.WriteSerializable(m.GetHvacErrorCode())
	if popErr := writeBuffer.PopContext("hvacErrorCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for hvacErrorCode")
	}
	if _hvacErrorCodeErr != nil {
		return errors.Wrap(_hvacErrorCodeErr, "Error serializing 'hvacErrorCode' field")
	}

		if popErr := writeBuffer.PopContext("AirConditioningDataZoneHvacPlantStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataZoneHvacPlantStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_AirConditioningDataZoneHvacPlantStatus) isAirConditioningDataZoneHvacPlantStatus() bool {
	return true
}

func (m *_AirConditioningDataZoneHvacPlantStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



