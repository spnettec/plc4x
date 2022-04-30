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

package readwrite

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/bacnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type BacnetipParserHelper struct {
}

func (m BacnetipParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (interface{}, error) {
	switch typeName {
	case "BACnetContextTag":
		tagNumberArgument, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		dataType := model.BACnetDataTypeByName(arguments[1])
		return model.BACnetContextTagParse(io, tagNumberArgument, dataType)
	case "BACnetStatusFlags":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetStatusFlagsParse(io, tagNumber)
	case "BACnetReadAccessPropertyError":
		return model.BACnetReadAccessPropertyErrorParse(io)
	case "BACnetTagPayloadReal":
		return model.BACnetTagPayloadRealParse(io)
	case "BVLCForeignDeviceTableEntry":
		return model.BVLCForeignDeviceTableEntryParse(io)
	case "NLM":
		apduLength, err := utils.StrToUint16(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.NLMParse(io, apduLength)
	case "BACnetActionCommand":
		return model.BACnetActionCommandParse(io)
	case "BACnetTagPayloadDate":
		return model.BACnetTagPayloadDateParse(io)
	case "BACnetNotificationParametersExtendedParameters":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetNotificationParametersExtendedParametersParse(io, tagNumber)
	case "BACnetReadAccessProperty":
		objectType := model.BACnetObjectTypeByName(arguments[0])
		return model.BACnetReadAccessPropertyParse(io, objectType)
	case "BACnetNotificationParametersChangeOfValueNewValue":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetNotificationParametersChangeOfValueNewValueParse(io, tagNumber)
	case "BACnetTagPayloadEnumerated":
		actualLength, err := utils.StrToUint32(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetTagPayloadEnumeratedParse(io, actualLength)
	case "BACnetTagPayloadOctetString":
		actualLength, err := utils.StrToUint32(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetTagPayloadOctetStringParse(io, actualLength)
	case "BACnetServiceAckAtomicReadFileStreamOrRecord":
		return model.BACnetServiceAckAtomicReadFileStreamOrRecordParse(io)
	case "NPDUControl":
		return model.NPDUControlParse(io)
	case "BACnetDeviceObjectPropertyReferenceEnclosed":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetDeviceObjectPropertyReferenceEnclosedParse(io, tagNumber)
	case "BACnetPropertyStates":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetPropertyStatesParse(io, tagNumber)
	case "BACnetReadAccessSpecification":
		return model.BACnetReadAccessSpecificationParse(io)
	case "BACnetReadAccessResult":
		return model.BACnetReadAccessResultParse(io)
	case "BACnetConstructedData":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		objectType := model.BACnetObjectTypeByName(arguments[1])
		var propertyIdentifierArgument model.BACnetContextTagPropertyIdentifier
		return model.BACnetConstructedDataParse(io, tagNumber, objectType, &propertyIdentifierArgument)
	case "BACnetSegmentation":
		return model.BACnetSegmentationParse(io)
	case "BACnetTagPayloadTime":
		return model.BACnetTagPayloadTimeParse(io)
	case "BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableParse(io, tagNumber)
	case "BACnetTagPayloadSignedInteger":
		actualLength, err := utils.StrToUint32(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetTagPayloadSignedIntegerParse(io, actualLength)
	case "BACnetUnconfirmedServiceRequest":
		serviceRequestLength, err := utils.StrToUint16(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetUnconfirmedServiceRequestParse(io, serviceRequestLength)
	case "BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord":
		return model.BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParse(io)
	case "BVLC":
		return model.BVLCParse(io)
	case "BACnetTagPayloadObjectIdentifier":
		return model.BACnetTagPayloadObjectIdentifierParse(io)
	case "BVLCBroadcastDistributionTableEntry":
		return model.BVLCBroadcastDistributionTableEntryParse(io)
	case "BACnetPropertyWriteDefinition":
		objectType := model.BACnetObjectTypeByName(arguments[0])
		return model.BACnetPropertyWriteDefinitionParse(io, objectType)
	case "BACnetDateTime":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetDateTimeParse(io, tagNumber)
	case "APDU":
		apduLength, err := utils.StrToUint16(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.APDUParse(io, apduLength)
	case "BACnetTagPayloadCharacterString":
		actualLength, err := utils.StrToUint32(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetTagPayloadCharacterStringParse(io, actualLength)
	case "BACnetError":
		return model.BACnetErrorParse(io)
	case "BACnetTimeStamp":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetTimeStampParse(io, tagNumber)
	case "BACnetNotificationParameters":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		objectType := model.BACnetObjectTypeByName(arguments[1])
		return model.BACnetNotificationParametersParse(io, tagNumber, objectType)
	case "BACnetConfirmedServiceRequest":
		serviceRequestLength, err := utils.StrToUint16(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetConfirmedServiceRequestParse(io, serviceRequestLength)
	case "BACnetAddress":
		return model.BACnetAddressParse(io)
	case "BACnetTagPayloadUnsignedInteger":
		actualLength, err := utils.StrToUint32(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetTagPayloadUnsignedIntegerParse(io, actualLength)
	case "BACnetApplicationTag":
		return model.BACnetApplicationTagParse(io)
	case "BACnetTagPayloadBitString":
		actualLength, err := utils.StrToUint32(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetTagPayloadBitStringParse(io, actualLength)
	case "BACnetDeviceObjectPropertyReference":
		return model.BACnetDeviceObjectPropertyReferenceParse(io)
	case "BACnetConstructedDataElement":
		objectType := model.BACnetObjectTypeByName(arguments[0])
		var propertyIdentifier model.BACnetContextTagPropertyIdentifier
		return model.BACnetConstructedDataElementParse(io, objectType, &propertyIdentifier)
	case "BACnetPropertyValues":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		objectType := model.BACnetObjectTypeByName(arguments[1])
		return model.BACnetPropertyValuesParse(io, tagNumber, objectType)
	case "BACnetTagHeader":
		return model.BACnetTagHeaderParse(io)
	case "BACnetTagPayloadBoolean":
		actualLength, err := utils.StrToUint32(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetTagPayloadBooleanParse(io, actualLength)
	case "BACnetTagPayloadDouble":
		return model.BACnetTagPayloadDoubleParse(io)
	case "BACnetPropertyValue":
		objectType := model.BACnetObjectTypeByName(arguments[0])
		return model.BACnetPropertyValueParse(io, objectType)
	case "NLMInitalizeRoutingTablePortMapping":
		return model.NLMInitalizeRoutingTablePortMappingParse(io)
	case "BACnetWriteAccessSpecification":
		return model.BACnetWriteAccessSpecificationParse(io)
	case "BACnetServiceAck":
		serviceRequestLength, err := utils.StrToUint16(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetServiceAckParse(io, serviceRequestLength)
	case "BACnetBinaryPV":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetBinaryPVParse(io, tagNumber)
	case "BACnetAction":
		tagNumber, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.BACnetActionParse(io, tagNumber)
	case "NPDU":
		npduLength, err := utils.StrToUint16(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.NPDUParse(io, npduLength)
	case "BACnetPropertyReference":
		return model.BACnetPropertyReferenceParse(io)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
