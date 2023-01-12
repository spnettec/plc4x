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

package readwrite

import (
	"github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
    "strings"
    "strconv"
)

// Code generated by code-generation. DO NOT EDIT.

type CbusXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m CbusXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
    switch typeName {
        case "HVACStatusFlags":
			return model.HVACStatusFlagsParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "ParameterValue":
            parameterType, _ := model.ParameterTypeByName(parserArguments[0])
			parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 8)
			if err!=nil {
				return nil, err
			}
			numBytes := uint8(parsedUint1)
            return model.ParameterValueParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), parameterType,  numBytes  )
        case "ReplyOrConfirmation":
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
			// TODO: find a way to parse the sub types
            var requestContext model.RequestContext
            return model.ReplyOrConfirmationParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions,  requestContext  )
        case "CBusOptions":
			return model.CBusOptionsParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "TemperatureBroadcastData":
			return model.TemperatureBroadcastDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "PanicStatus":
			return model.PanicStatusParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "IdentifyReplyCommandUnitSummary":
			return model.IdentifyReplyCommandUnitSummaryParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "InterfaceOptions1PowerUpSettings":
			return model.InterfaceOptions1PowerUpSettingsParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "MonitoredSAL":
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
            return model.MonitoredSALParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions  )
        case "ReplyNetwork":
			return model.ReplyNetworkParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "SerialNumber":
			return model.SerialNumberParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "CBusPointToMultiPointCommand":
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
            return model.CBusPointToMultiPointCommandParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions  )
        case "StatusRequest":
			return model.StatusRequestParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "InterfaceOptions3":
			return model.InterfaceOptions3ParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "InterfaceOptions1":
			return model.InterfaceOptions1ParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "InterfaceOptions2":
			return model.InterfaceOptions2ParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "HVACModeAndFlags":
			return model.HVACModeAndFlagsParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "LightingData":
			return model.LightingDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "SALData":
            applicationId, _ := model.ApplicationIdByName(parserArguments[0])
            return model.SALDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), applicationId  )
        case "CBusCommand":
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
            return model.CBusCommandParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions  )
        case "HVACHumidity":
			return model.HVACHumidityParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "HVACHumidityModeAndFlags":
			return model.HVACHumidityModeAndFlagsParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "CBusConstants":
			return model.CBusConstantsParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "SerialInterfaceAddress":
			return model.SerialInterfaceAddressParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "MeasurementData":
			return model.MeasurementDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "HVACZoneList":
			return model.HVACZoneListParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "MediaTransportControlData":
			return model.MediaTransportControlDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "StatusByte":
			return model.StatusByteParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "TriggerControlLabelOptions":
			return model.TriggerControlLabelOptionsParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "HVACAuxiliaryLevel":
			return model.HVACAuxiliaryLevelParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "ErrorReportingData":
			return model.ErrorReportingDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "UnitAddress":
			return model.UnitAddressParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "SecurityArmCode":
			return model.SecurityArmCodeParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "MeteringData":
			return model.MeteringDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "EnableControlData":
			return model.EnableControlDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "ApplicationAddress2":
			return model.ApplicationAddress2ParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "ApplicationAddress1":
			return model.ApplicationAddress1ParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "RequestContext":
			return model.RequestContextParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "TriggerControlData":
			return model.TriggerControlDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "HVACStartTime":
			return model.HVACStartTimeParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "HVACTemperature":
			return model.HVACTemperatureParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "RequestTermination":
			return model.RequestTerminationParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "CBusMessage":
            isResponse := parserArguments[0] == "true"
			// TODO: find a way to parse the sub types
            var requestContext model.RequestContext
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
            return model.CBusMessageParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), isResponse,  requestContext,  cBusOptions  )
        case "ErrorReportingSystemCategory":
			return model.ErrorReportingSystemCategoryParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "PowerUp":
			return model.PowerUpParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "Reply":
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
			// TODO: find a way to parse the sub types
            var requestContext model.RequestContext
            return model.ReplyParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions,  requestContext  )
        case "TelephonyData":
			return model.TelephonyDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "HVACHumidityStatusFlags":
			return model.HVACHumidityStatusFlagsParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "ParameterChange":
			return model.ParameterChangeParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "ErrorReportingSystemCategoryType":
            errorReportingSystemCategoryClass, _ := model.ErrorReportingSystemCategoryClassByName(parserArguments[0])
            return model.ErrorReportingSystemCategoryTypeParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), errorReportingSystemCategoryClass  )
        case "Confirmation":
			return model.ConfirmationParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "SecurityData":
			return model.SecurityDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "NetworkProtocolControlInformation":
			return model.NetworkProtocolControlInformationParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "CBusHeader":
			return model.CBusHeaderParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "Request":
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
            return model.RequestParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions  )
        case "Alpha":
			return model.AlphaParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "CALData":
			// TODO: find a way to parse the sub types
            var requestContext model.RequestContext
            return model.CALDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), requestContext  )
        case "Checksum":
			return model.ChecksumParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "CALReply":
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
			// TODO: find a way to parse the sub types
            var requestContext model.RequestContext
            return model.CALReplyParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions,  requestContext  )
        case "CustomManufacturer":
			parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
			if err!=nil {
				return nil, err
			}
			numBytes := uint8(parsedUint0)
            return model.CustomManufacturerParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), numBytes  )
        case "AccessControlData":
			return model.AccessControlDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "ClockAndTimekeepingData":
			return model.ClockAndTimekeepingDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "NetworkRoute":
			return model.NetworkRouteParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "ResponseTermination":
			return model.ResponseTerminationParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "LevelInformation":
			return model.LevelInformationParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "TamperStatus":
			return model.TamperStatusParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "IdentifyReplyCommand":
            attribute, _ := model.AttributeByName(parserArguments[0])
			parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 5)
			if err!=nil {
				return nil, err
			}
			numBytes := uint8(parsedUint1)
            return model.IdentifyReplyCommandParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), attribute,  numBytes  )
        case "HVACRawLevels":
			return model.HVACRawLevelsParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "ZoneStatus":
			return model.ZoneStatusParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "BridgeAddress":
			return model.BridgeAddressParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "LightingLabelOptions":
			return model.LightingLabelOptionsParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "CustomTypes":
			parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
			if err!=nil {
				return nil, err
			}
			numBytes := uint8(parsedUint0)
            return model.CustomTypesParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), numBytes  )
        case "EncodedReply":
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
			// TODO: find a way to parse the sub types
            var requestContext model.RequestContext
            return model.EncodedReplyParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions,  requestContext  )
        case "CBusPointToPointToMultiPointCommand":
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
            return model.CBusPointToPointToMultiPointCommandParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions  )
        case "CBusPointToPointCommand":
			// TODO: find a way to parse the sub types
            var cBusOptions model.CBusOptions
            return model.CBusPointToPointCommandParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions  )
        case "AirConditioningData":
			return model.AirConditioningDataParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "LogicAssignment":
			return model.LogicAssignmentParseWithBuffer(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
    }
    return nil, errors.Errorf("Unsupported type %s", typeName)
}
