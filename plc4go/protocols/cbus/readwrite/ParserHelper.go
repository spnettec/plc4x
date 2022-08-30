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
)

// Code generated by code-generation. DO NOT EDIT.

type CbusParserHelper struct {
}

func (m CbusParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (interface{}, error) {
	switch typeName {
	case "HVACStatusFlags":
		return model.HVACStatusFlagsParse(io)
	case "ParameterValue":
        parameterType, _ := model.ParameterTypeByName(arguments[0])
		numBytes, err := utils.StrToUint8(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.ParameterValueParse(io, parameterType, numBytes)
	case "ReplyOrConfirmation":
		var cBusOptions model.CBusOptions
		var requestContext model.RequestContext
		return model.ReplyOrConfirmationParse(io, cBusOptions, requestContext)
	case "CBusOptions":
		return model.CBusOptionsParse(io)
	case "TemperatureBroadcastData":
		return model.TemperatureBroadcastDataParse(io)
	case "PanicStatus":
		return model.PanicStatusParse(io)
	case "IdentifyReplyCommandUnitSummary":
		return model.IdentifyReplyCommandUnitSummaryParse(io)
	case "InterfaceOptions1PowerUpSettings":
		return model.InterfaceOptions1PowerUpSettingsParse(io)
	case "MonitoredSAL":
		var cBusOptions model.CBusOptions
		return model.MonitoredSALParse(io, cBusOptions)
	case "ReplyNetwork":
		return model.ReplyNetworkParse(io)
	case "SerialNumber":
		return model.SerialNumberParse(io)
	case "CBusPointToMultiPointCommand":
		var cBusOptions model.CBusOptions
		return model.CBusPointToMultiPointCommandParse(io, cBusOptions)
	case "StatusRequest":
		return model.StatusRequestParse(io)
	case "InterfaceOptions3":
		return model.InterfaceOptions3Parse(io)
	case "InterfaceOptions1":
		return model.InterfaceOptions1Parse(io)
	case "InterfaceOptions2":
		return model.InterfaceOptions2Parse(io)
	case "HVACModeAndFlags":
		return model.HVACModeAndFlagsParse(io)
	case "LightingData":
		return model.LightingDataParse(io)
	case "SALData":
        applicationId, _ := model.ApplicationIdByName(arguments[0])
		return model.SALDataParse(io, applicationId)
	case "CBusCommand":
		var cBusOptions model.CBusOptions
		return model.CBusCommandParse(io, cBusOptions)
	case "HVACHumidity":
		return model.HVACHumidityParse(io)
	case "HVACHumidityModeAndFlags":
		return model.HVACHumidityModeAndFlagsParse(io)
	case "CBusConstants":
		return model.CBusConstantsParse(io)
	case "SerialInterfaceAddress":
		return model.SerialInterfaceAddressParse(io)
	case "MeasurementData":
		return model.MeasurementDataParse(io)
	case "HVACZoneList":
		return model.HVACZoneListParse(io)
	case "MediaTransportControlData":
		return model.MediaTransportControlDataParse(io)
	case "StatusByte":
		return model.StatusByteParse(io)
	case "TriggerControlLabelOptions":
		return model.TriggerControlLabelOptionsParse(io)
	case "HVACAuxiliaryLevel":
		return model.HVACAuxiliaryLevelParse(io)
	case "ErrorReportingData":
		return model.ErrorReportingDataParse(io)
	case "UnitAddress":
		return model.UnitAddressParse(io)
	case "SecurityArmCode":
		return model.SecurityArmCodeParse(io)
	case "MeteringData":
		return model.MeteringDataParse(io)
	case "EnableControlData":
		return model.EnableControlDataParse(io)
	case "ApplicationAddress2":
		return model.ApplicationAddress2Parse(io)
	case "ApplicationAddress1":
		return model.ApplicationAddress1Parse(io)
	case "RequestContext":
		return model.RequestContextParse(io)
	case "TriggerControlData":
		return model.TriggerControlDataParse(io)
	case "HVACStartTime":
		return model.HVACStartTimeParse(io)
	case "HVACTemperature":
		return model.HVACTemperatureParse(io)
	case "RequestTermination":
		return model.RequestTerminationParse(io)
	case "CBusMessage":
		isResponse, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		var requestContext model.RequestContext
		var cBusOptions model.CBusOptions
		return model.CBusMessageParse(io, isResponse, requestContext, cBusOptions)
	case "ErrorReportingSystemCategory":
		return model.ErrorReportingSystemCategoryParse(io)
	case "PowerUp":
		return model.PowerUpParse(io)
	case "Reply":
		var cBusOptions model.CBusOptions
		var requestContext model.RequestContext
		return model.ReplyParse(io, cBusOptions, requestContext)
	case "TelephonyData":
		return model.TelephonyDataParse(io)
	case "HVACHumidityStatusFlags":
		return model.HVACHumidityStatusFlagsParse(io)
	case "ParameterChange":
		return model.ParameterChangeParse(io)
	case "ErrorReportingSystemCategoryType":
        errorReportingSystemCategoryClass, _ := model.ErrorReportingSystemCategoryClassByName(arguments[0])
		return model.ErrorReportingSystemCategoryTypeParse(io, errorReportingSystemCategoryClass)
	case "Confirmation":
		return model.ConfirmationParse(io)
	case "SecurityData":
		return model.SecurityDataParse(io)
	case "NetworkProtocolControlInformation":
		return model.NetworkProtocolControlInformationParse(io)
	case "CBusHeader":
		return model.CBusHeaderParse(io)
	case "Request":
		var cBusOptions model.CBusOptions
		return model.RequestParse(io, cBusOptions)
	case "Alpha":
		return model.AlphaParse(io)
	case "CALData":
		var requestContext model.RequestContext
		return model.CALDataParse(io, requestContext)
	case "Checksum":
		return model.ChecksumParse(io)
	case "CALReply":
		var cBusOptions model.CBusOptions
		var requestContext model.RequestContext
		return model.CALReplyParse(io, cBusOptions, requestContext)
	case "CustomManufacturer":
		numBytes, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.CustomManufacturerParse(io, numBytes)
	case "AccessControlData":
		return model.AccessControlDataParse(io)
	case "ClockAndTimekeepingData":
		return model.ClockAndTimekeepingDataParse(io)
	case "NetworkRoute":
		return model.NetworkRouteParse(io)
	case "ResponseTermination":
		return model.ResponseTerminationParse(io)
	case "LevelInformation":
		return model.LevelInformationParse(io)
	case "TamperStatus":
		return model.TamperStatusParse(io)
	case "IdentifyReplyCommand":
        attribute, _ := model.AttributeByName(arguments[0])
		numBytes, err := utils.StrToUint8(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.IdentifyReplyCommandParse(io, attribute, numBytes)
	case "HVACRawLevels":
		return model.HVACRawLevelsParse(io)
	case "ZoneStatus":
		return model.ZoneStatusParse(io)
	case "BridgeAddress":
		return model.BridgeAddressParse(io)
	case "LightingLabelOptions":
		return model.LightingLabelOptionsParse(io)
	case "CustomTypes":
		numBytes, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.CustomTypesParse(io, numBytes)
	case "EncodedReply":
		var cBusOptions model.CBusOptions
		var requestContext model.RequestContext
		return model.EncodedReplyParse(io, cBusOptions, requestContext)
	case "CBusPointToPointToMultiPointCommand":
		var cBusOptions model.CBusOptions
		return model.CBusPointToPointToMultiPointCommandParse(io, cBusOptions)
	case "CBusPointToPointCommand":
		var cBusOptions model.CBusOptions
		return model.CBusPointToPointCommandParse(io, cBusOptions)
	case "AirConditioningData":
		return model.AirConditioningDataParse(io)
	case "LogicAssignment":
		return model.LogicAssignmentParse(io)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
