/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedData is the data-structure of this message
type BACnetConstructedData struct {
	OpeningTag *BACnetOpeningTag
	ClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber uint8
	Child     IBACnetConstructedDataChild
}

// IBACnetConstructedData is the corresponding interface of BACnetConstructedData
type IBACnetConstructedData interface {
	// GetObjectType returns ObjectType (discriminator field)
	GetObjectType() BACnetObjectType
	// GetPropertyIdentifierArgument returns PropertyIdentifierArgument (discriminator field)
	GetPropertyIdentifierArgument() BACnetPropertyIdentifier
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetConstructedDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetConstructedData, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetConstructedDataChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag)
	GetParent() *BACnetConstructedData

	GetTypeName() string
	IBACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedData) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetConstructedData) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedData factory function for BACnetConstructedData
func NewBACnetConstructedData(openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedData {
	return &BACnetConstructedData{OpeningTag: openingTag, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetConstructedData(structType interface{}) *BACnetConstructedData {
	if casted, ok := structType.(BACnetConstructedData); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetConstructedDataChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetConstructedData) GetTypeName() string {
	return "BACnetConstructedData"
}

func (m *BACnetConstructedData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedData) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetConstructedData) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedData"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetConstructedDataChild interface {
		InitializeParent(*BACnetConstructedData, *BACnetOpeningTag, *BACnetClosingTag)
		GetParent() *BACnetConstructedData
	}
	var _child BACnetConstructedDataChild
	var typeSwitchError error
	switch {
	case true && propertyIdentifierArgument == BACnetPropertyIdentifier_ACCEPTED_MODES: // BACnetConstructedDataAcceptedModes
		_child, typeSwitchError = BACnetConstructedDataAcceptedModesParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LOOP && propertyIdentifierArgument == BACnetPropertyIdentifier_ACTION: // BACnetConstructedDataLoopAction
		_child, typeSwitchError = BACnetConstructedDataLoopActionParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true && propertyIdentifierArgument == BACnetPropertyIdentifier_ACTION: // BACnetConstructedDataAction
		_child, typeSwitchError = BACnetConstructedDataActionParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LIFE_SAFETY_POINT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALARM_VALUES: // BACnetConstructedDataLifeSafetyPointAlarmValues
		_child, typeSwitchError = BACnetConstructedDataLifeSafetyPointAlarmValuesParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ACCESS_CREDENTIAL && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAccessCredentialAl
		_child, typeSwitchError = BACnetConstructedDataAccessCredentialAlParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ACCESS_DOOR && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAccessDoorAll
		_child, typeSwitchError = BACnetConstructedDataAccessDoorAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ACCESS_POINT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAccessPointAll
		_child, typeSwitchError = BACnetConstructedDataAccessPointAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ACCESS_RIGHTS && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAccessRightsAll
		_child, typeSwitchError = BACnetConstructedDataAccessRightsAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ACCESS_USER && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAccessUserAll
		_child, typeSwitchError = BACnetConstructedDataAccessUserAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ACCESS_ZONE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAccessZoneAll
		_child, typeSwitchError = BACnetConstructedDataAccessZoneAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ACCUMULATOR && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAccumulatorAll
		_child, typeSwitchError = BACnetConstructedDataAccumulatorAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ALERT_ENROLLMENT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAlertEnrollmentAll
		_child, typeSwitchError = BACnetConstructedDataAlertEnrollmentAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ANALOG_INPUT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAnalogInputAll
		_child, typeSwitchError = BACnetConstructedDataAnalogInputAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ANALOG_OUTPUT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAnalogOutputAll
		_child, typeSwitchError = BACnetConstructedDataAnalogOutputAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ANALOG_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAnalogValueAll
		_child, typeSwitchError = BACnetConstructedDataAnalogValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_AVERAGING && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataAveragingAll
		_child, typeSwitchError = BACnetConstructedDataAveragingAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_BINARY_INPUT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataBinaryInputAll
		_child, typeSwitchError = BACnetConstructedDataBinaryInputAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_BINARY_LIGHTING_OUTPUT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataBinaryLightingOutputAll
		_child, typeSwitchError = BACnetConstructedDataBinaryLightingOutputAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_BINARY_OUTPUT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataBinaryOutputAll
		_child, typeSwitchError = BACnetConstructedDataBinaryOutputAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_BINARY_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataBinaryValueAll
		_child, typeSwitchError = BACnetConstructedDataBinaryValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_BITSTRING_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataBitstringValueAll
		_child, typeSwitchError = BACnetConstructedDataBitstringValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_CALENDAR && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataCalendarAll
		_child, typeSwitchError = BACnetConstructedDataCalendarAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_CHANNEL && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataChannelAll
		_child, typeSwitchError = BACnetConstructedDataChannelAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_CHARACTERSTRING_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataCharacterstringValueAll
		_child, typeSwitchError = BACnetConstructedDataCharacterstringValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_COMMAND && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataCommandAll
		_child, typeSwitchError = BACnetConstructedDataCommandAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_CREDENTIAL_DATA_INPUT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataCredentialDataInputAll
		_child, typeSwitchError = BACnetConstructedDataCredentialDataInputAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_DATEPATTERN_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataDatepatternValueAll
		_child, typeSwitchError = BACnetConstructedDataDatepatternValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_DATE_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataDateValueAll
		_child, typeSwitchError = BACnetConstructedDataDateValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_DATETIMEPATTERN_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataDatetimepatternValueAll
		_child, typeSwitchError = BACnetConstructedDataDatetimepatternValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_DATETIME_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataDatetimeValueAll
		_child, typeSwitchError = BACnetConstructedDataDatetimeValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_DEVICE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataDeviceAll
		_child, typeSwitchError = BACnetConstructedDataDeviceAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ELEVATOR_GROUP && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataElevatorGroupAll
		_child, typeSwitchError = BACnetConstructedDataElevatorGroupAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_ESCALATOR && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataEscalatorAll
		_child, typeSwitchError = BACnetConstructedDataEscalatorAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_EVENT_ENROLLMENT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataEventEnrollmentAll
		_child, typeSwitchError = BACnetConstructedDataEventEnrollmentAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_EVENT_LOG && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataEventLogAll
		_child, typeSwitchError = BACnetConstructedDataEventLogAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_FILE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataFileAll
		_child, typeSwitchError = BACnetConstructedDataFileAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_GLOBAL_GROUP && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataGlobalGroupAll
		_child, typeSwitchError = BACnetConstructedDataGlobalGroupAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_GROUP && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataGroupAll
		_child, typeSwitchError = BACnetConstructedDataGroupAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_INTEGER_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataIntegerValueAll
		_child, typeSwitchError = BACnetConstructedDataIntegerValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LARGE_ANALOG_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataLargeAnalogValueAll
		_child, typeSwitchError = BACnetConstructedDataLargeAnalogValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LIFE_SAFETY_POINT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataLifeSafetyPointAll
		_child, typeSwitchError = BACnetConstructedDataLifeSafetyPointAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LIFE_SAFETY_ZONE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataLifeSafetyZoneAll
		_child, typeSwitchError = BACnetConstructedDataLifeSafetyZoneAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LIFT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataLiftAll
		_child, typeSwitchError = BACnetConstructedDataLiftAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LIGHTING_OUTPUT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataLightingOutputAll
		_child, typeSwitchError = BACnetConstructedDataLightingOutputAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LOAD_CONTROL && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataLoadControlAll
		_child, typeSwitchError = BACnetConstructedDataLoadControlAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LOOP && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataLoopAll
		_child, typeSwitchError = BACnetConstructedDataLoopAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_MULTI_STATE_INPUT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataMultiStateInputAll
		_child, typeSwitchError = BACnetConstructedDataMultiStateInputAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_MULTI_STATE_OUTPUT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataMultiStateOutputAll
		_child, typeSwitchError = BACnetConstructedDataMultiStateOutputAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_MULTI_STATE_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataMultiStateValueAll
		_child, typeSwitchError = BACnetConstructedDataMultiStateValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_NETWORK_PORT && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataNetworkPortAll
		_child, typeSwitchError = BACnetConstructedDataNetworkPortAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_NETWORK_SECURITY && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataNetworkSecurityAll
		_child, typeSwitchError = BACnetConstructedDataNetworkSecurityAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_NOTIFICATION_CLASS && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataNotificationClassAll
		_child, typeSwitchError = BACnetConstructedDataNotificationClassAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_NOTIFICATION_FORWARDER && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataNotificationForwarderAll
		_child, typeSwitchError = BACnetConstructedDataNotificationForwarderAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_OCTETSTRING_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataOctetstringValueAll
		_child, typeSwitchError = BACnetConstructedDataOctetstringValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_POSITIVE_INTEGER_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataPositiveIntegerValueAll
		_child, typeSwitchError = BACnetConstructedDataPositiveIntegerValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_PROGRAM && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataProgramAll
		_child, typeSwitchError = BACnetConstructedDataProgramAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_PULSE_CONVERTER && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataPulseConverterAll
		_child, typeSwitchError = BACnetConstructedDataPulseConverterAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_SCHEDULE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataScheduleAll
		_child, typeSwitchError = BACnetConstructedDataScheduleAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_STRUCTURED_VIEW && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataStructuredViewAll
		_child, typeSwitchError = BACnetConstructedDataStructuredViewAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_TIMEPATTERN_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataTimepatternValueAll
		_child, typeSwitchError = BACnetConstructedDataTimepatternValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_TIME_VALUE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataTimeValueAll
		_child, typeSwitchError = BACnetConstructedDataTimeValueAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_TIMER && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataTimerAll
		_child, typeSwitchError = BACnetConstructedDataTimerAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_TREND_LOG && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataTrendLogAll
		_child, typeSwitchError = BACnetConstructedDataTrendLogAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_TREND_LOG_MULTIPLE && propertyIdentifierArgument == BACnetPropertyIdentifier_ALL: // BACnetConstructedDataTrendLogMultipleAll
		_child, typeSwitchError = BACnetConstructedDataTrendLogMultipleAllParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true && propertyIdentifierArgument == BACnetPropertyIdentifier_EVENT_TIME_STAMPS: // BACnetConstructedDataEventTimestamps
		_child, typeSwitchError = BACnetConstructedDataEventTimestampsParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case objectType == BACnetObjectType_LIFE_SAFETY_POINT && propertyIdentifierArgument == BACnetPropertyIdentifier_FAULT_VALUES: // BACnetConstructedDataLifeSafetyPointFaultValues
		_child, typeSwitchError = BACnetConstructedDataLifeSafetyPointFaultValuesParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true && propertyIdentifierArgument == BACnetPropertyIdentifier_LIFE_SAFETY_ALARM_VALUES: // BACnetConstructedDataLifeSafetyAlarmValues
		_child, typeSwitchError = BACnetConstructedDataLifeSafetyAlarmValuesParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true && propertyIdentifierArgument == BACnetPropertyIdentifier_LIST_OF_OBJECT_PROPERTY_REFERENCES: // BACnetConstructedDataListOfObjectPropertyReferences
		_child, typeSwitchError = BACnetConstructedDataListOfObjectPropertyReferencesParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true && propertyIdentifierArgument == BACnetPropertyIdentifier_MEMBER_OF: // BACnetConstructedDataMemberOf
		_child, typeSwitchError = BACnetConstructedDataMemberOfParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true && propertyIdentifierArgument == BACnetPropertyIdentifier_RELIABILITY: // BACnetConstructedDataReliability
		_child, typeSwitchError = BACnetConstructedDataReliabilityParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true && propertyIdentifierArgument == BACnetPropertyIdentifier_SUBORDINATE_LIST: // BACnetConstructedDataSubordinateList
		_child, typeSwitchError = BACnetConstructedDataSubordinateListParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true && propertyIdentifierArgument == BACnetPropertyIdentifier_ZONE_MEMBERS: // BACnetConstructedDataZoneMembers
		_child, typeSwitchError = BACnetConstructedDataZoneMembersParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true: // BACnetConstructedDataUnspecified
		_child, typeSwitchError = BACnetConstructedDataUnspecifiedParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedData"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), openingTag, closingTag)
	return _child.GetParent(), nil
}

func (m *BACnetConstructedData) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetConstructedData) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetConstructedData, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetConstructedData"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConstructedData"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConstructedData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
