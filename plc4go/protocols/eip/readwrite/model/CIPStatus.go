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

// CIPStatus is an enum
type CIPStatus uint32

type ICIPStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	CIPStatus_Success CIPStatus = 0x00000000
	CIPStatus_ConnectionFailure CIPStatus = 0x00000001
	CIPStatus_ResourceUnAvailable CIPStatus = 0x00000002
	CIPStatus_InvalidParameterValue CIPStatus = 0x00000003
	CIPStatus_PathSegmentError CIPStatus = 0x00000004
	CIPStatus_PathDestinationUnknown CIPStatus = 0x00000005
	CIPStatus_PartialTransfer CIPStatus = 0x00000006
	CIPStatus_ConnectionIDNotValid CIPStatus = 0x00000007
	CIPStatus_ServiceNotSupported CIPStatus = 0x00000008
	CIPStatus_InvalidAttributeValue CIPStatus = 0x00000009
	CIPStatus_AttributeListError CIPStatus = 0x0000000A
	CIPStatus_AlreadyInRequestedState CIPStatus = 0x0000000B
	CIPStatus_ObjectStateConflict CIPStatus = 0x0000000C
	CIPStatus_ObjectAlreadyExists CIPStatus = 0x0000000D
	CIPStatus_AttributeNotSettable CIPStatus = 0x0000000E
	CIPStatus_PrivilegeViolation CIPStatus = 0x0000000F
	CIPStatus_DeviceStateConflict CIPStatus = 0x00000010
	CIPStatus_ReplyDataTooLarge CIPStatus = 0x00000011
	CIPStatus_FragmentationOfPrimitiveValue CIPStatus = 0x00000012
	CIPStatus_NotEnoughData CIPStatus = 0x00000013
	CIPStatus_AttributeNotSupported CIPStatus = 0x00000014
	CIPStatus_TooMuchData CIPStatus = 0x00000015
	CIPStatus_ObjectDoesNotExist CIPStatus = 0x00000016
	CIPStatus_ServiceFragmentation CIPStatus = 0x00000017
	CIPStatus_NoStoredAttributeData CIPStatus = 0x00000018
	CIPStatus_StoreOperationFailure CIPStatus = 0x00000019
	CIPStatus_RequestPacketTooLarge CIPStatus = 0x0000001A
	CIPStatus_ResponsePacketTooLarge CIPStatus = 0x0000001B
	CIPStatus_MissingAttributeListEntryData CIPStatus = 0x0000001C
	CIPStatus_InvalidAttributeValueList CIPStatus = 0x0000001D
	CIPStatus_EmbeddedServiceError CIPStatus = 0x0000001E
	CIPStatus_VendorSpecificError CIPStatus = 0x0000001F
	CIPStatus_InvalidCommandWithWrongEndianess CIPStatus = 0x01000000
)

var CIPStatusValues []CIPStatus

func init() {
	_ = errors.New
	CIPStatusValues = []CIPStatus {
		CIPStatus_Success,
		CIPStatus_ConnectionFailure,
		CIPStatus_ResourceUnAvailable,
		CIPStatus_InvalidParameterValue,
		CIPStatus_PathSegmentError,
		CIPStatus_PathDestinationUnknown,
		CIPStatus_PartialTransfer,
		CIPStatus_ConnectionIDNotValid,
		CIPStatus_ServiceNotSupported,
		CIPStatus_InvalidAttributeValue,
		CIPStatus_AttributeListError,
		CIPStatus_AlreadyInRequestedState,
		CIPStatus_ObjectStateConflict,
		CIPStatus_ObjectAlreadyExists,
		CIPStatus_AttributeNotSettable,
		CIPStatus_PrivilegeViolation,
		CIPStatus_DeviceStateConflict,
		CIPStatus_ReplyDataTooLarge,
		CIPStatus_FragmentationOfPrimitiveValue,
		CIPStatus_NotEnoughData,
		CIPStatus_AttributeNotSupported,
		CIPStatus_TooMuchData,
		CIPStatus_ObjectDoesNotExist,
		CIPStatus_ServiceFragmentation,
		CIPStatus_NoStoredAttributeData,
		CIPStatus_StoreOperationFailure,
		CIPStatus_RequestPacketTooLarge,
		CIPStatus_ResponsePacketTooLarge,
		CIPStatus_MissingAttributeListEntryData,
		CIPStatus_InvalidAttributeValueList,
		CIPStatus_EmbeddedServiceError,
		CIPStatus_VendorSpecificError,
		CIPStatus_InvalidCommandWithWrongEndianess,
	}
}

func CIPStatusByValue(value uint32) (enum CIPStatus, ok bool) {
	switch value {
		case 0x00000000:
			return CIPStatus_Success, true
		case 0x00000001:
			return CIPStatus_ConnectionFailure, true
		case 0x00000002:
			return CIPStatus_ResourceUnAvailable, true
		case 0x00000003:
			return CIPStatus_InvalidParameterValue, true
		case 0x00000004:
			return CIPStatus_PathSegmentError, true
		case 0x00000005:
			return CIPStatus_PathDestinationUnknown, true
		case 0x00000006:
			return CIPStatus_PartialTransfer, true
		case 0x00000007:
			return CIPStatus_ConnectionIDNotValid, true
		case 0x00000008:
			return CIPStatus_ServiceNotSupported, true
		case 0x00000009:
			return CIPStatus_InvalidAttributeValue, true
		case 0x0000000A:
			return CIPStatus_AttributeListError, true
		case 0x0000000B:
			return CIPStatus_AlreadyInRequestedState, true
		case 0x0000000C:
			return CIPStatus_ObjectStateConflict, true
		case 0x0000000D:
			return CIPStatus_ObjectAlreadyExists, true
		case 0x0000000E:
			return CIPStatus_AttributeNotSettable, true
		case 0x0000000F:
			return CIPStatus_PrivilegeViolation, true
		case 0x00000010:
			return CIPStatus_DeviceStateConflict, true
		case 0x00000011:
			return CIPStatus_ReplyDataTooLarge, true
		case 0x00000012:
			return CIPStatus_FragmentationOfPrimitiveValue, true
		case 0x00000013:
			return CIPStatus_NotEnoughData, true
		case 0x00000014:
			return CIPStatus_AttributeNotSupported, true
		case 0x00000015:
			return CIPStatus_TooMuchData, true
		case 0x00000016:
			return CIPStatus_ObjectDoesNotExist, true
		case 0x00000017:
			return CIPStatus_ServiceFragmentation, true
		case 0x00000018:
			return CIPStatus_NoStoredAttributeData, true
		case 0x00000019:
			return CIPStatus_StoreOperationFailure, true
		case 0x0000001A:
			return CIPStatus_RequestPacketTooLarge, true
		case 0x0000001B:
			return CIPStatus_ResponsePacketTooLarge, true
		case 0x0000001C:
			return CIPStatus_MissingAttributeListEntryData, true
		case 0x0000001D:
			return CIPStatus_InvalidAttributeValueList, true
		case 0x0000001E:
			return CIPStatus_EmbeddedServiceError, true
		case 0x0000001F:
			return CIPStatus_VendorSpecificError, true
		case 0x01000000:
			return CIPStatus_InvalidCommandWithWrongEndianess, true
	}
	return 0, false
}

func CIPStatusByName(value string) (enum CIPStatus, ok bool) {
	switch value {
	case "Success":
		return CIPStatus_Success, true
	case "ConnectionFailure":
		return CIPStatus_ConnectionFailure, true
	case "ResourceUnAvailable":
		return CIPStatus_ResourceUnAvailable, true
	case "InvalidParameterValue":
		return CIPStatus_InvalidParameterValue, true
	case "PathSegmentError":
		return CIPStatus_PathSegmentError, true
	case "PathDestinationUnknown":
		return CIPStatus_PathDestinationUnknown, true
	case "PartialTransfer":
		return CIPStatus_PartialTransfer, true
	case "ConnectionIDNotValid":
		return CIPStatus_ConnectionIDNotValid, true
	case "ServiceNotSupported":
		return CIPStatus_ServiceNotSupported, true
	case "InvalidAttributeValue":
		return CIPStatus_InvalidAttributeValue, true
	case "AttributeListError":
		return CIPStatus_AttributeListError, true
	case "AlreadyInRequestedState":
		return CIPStatus_AlreadyInRequestedState, true
	case "ObjectStateConflict":
		return CIPStatus_ObjectStateConflict, true
	case "ObjectAlreadyExists":
		return CIPStatus_ObjectAlreadyExists, true
	case "AttributeNotSettable":
		return CIPStatus_AttributeNotSettable, true
	case "PrivilegeViolation":
		return CIPStatus_PrivilegeViolation, true
	case "DeviceStateConflict":
		return CIPStatus_DeviceStateConflict, true
	case "ReplyDataTooLarge":
		return CIPStatus_ReplyDataTooLarge, true
	case "FragmentationOfPrimitiveValue":
		return CIPStatus_FragmentationOfPrimitiveValue, true
	case "NotEnoughData":
		return CIPStatus_NotEnoughData, true
	case "AttributeNotSupported":
		return CIPStatus_AttributeNotSupported, true
	case "TooMuchData":
		return CIPStatus_TooMuchData, true
	case "ObjectDoesNotExist":
		return CIPStatus_ObjectDoesNotExist, true
	case "ServiceFragmentation":
		return CIPStatus_ServiceFragmentation, true
	case "NoStoredAttributeData":
		return CIPStatus_NoStoredAttributeData, true
	case "StoreOperationFailure":
		return CIPStatus_StoreOperationFailure, true
	case "RequestPacketTooLarge":
		return CIPStatus_RequestPacketTooLarge, true
	case "ResponsePacketTooLarge":
		return CIPStatus_ResponsePacketTooLarge, true
	case "MissingAttributeListEntryData":
		return CIPStatus_MissingAttributeListEntryData, true
	case "InvalidAttributeValueList":
		return CIPStatus_InvalidAttributeValueList, true
	case "EmbeddedServiceError":
		return CIPStatus_EmbeddedServiceError, true
	case "VendorSpecificError":
		return CIPStatus_VendorSpecificError, true
	case "InvalidCommandWithWrongEndianess":
		return CIPStatus_InvalidCommandWithWrongEndianess, true
	}
	return 0, false
}

func CIPStatusKnows(value uint32)  bool {
	for _, typeValue := range CIPStatusValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastCIPStatus(structType any) CIPStatus {
	castFunc := func(typ any) CIPStatus {
		if sCIPStatus, ok := typ.(CIPStatus); ok {
			return sCIPStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m CIPStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m CIPStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPStatusParse(ctx context.Context, theBytes []byte) (CIPStatus, error) {
	return CIPStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CIPStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CIPStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("CIPStatus", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading CIPStatus")
	}
	if enum, ok := CIPStatusByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for CIPStatus")
		return CIPStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e CIPStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e CIPStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("CIPStatus", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e CIPStatus) PLC4XEnumName() string {
	switch e {
	case CIPStatus_Success:
		return "Success"
	case CIPStatus_ConnectionFailure:
		return "ConnectionFailure"
	case CIPStatus_ResourceUnAvailable:
		return "ResourceUnAvailable"
	case CIPStatus_InvalidParameterValue:
		return "InvalidParameterValue"
	case CIPStatus_PathSegmentError:
		return "PathSegmentError"
	case CIPStatus_PathDestinationUnknown:
		return "PathDestinationUnknown"
	case CIPStatus_PartialTransfer:
		return "PartialTransfer"
	case CIPStatus_ConnectionIDNotValid:
		return "ConnectionIDNotValid"
	case CIPStatus_ServiceNotSupported:
		return "ServiceNotSupported"
	case CIPStatus_InvalidAttributeValue:
		return "InvalidAttributeValue"
	case CIPStatus_AttributeListError:
		return "AttributeListError"
	case CIPStatus_AlreadyInRequestedState:
		return "AlreadyInRequestedState"
	case CIPStatus_ObjectStateConflict:
		return "ObjectStateConflict"
	case CIPStatus_ObjectAlreadyExists:
		return "ObjectAlreadyExists"
	case CIPStatus_AttributeNotSettable:
		return "AttributeNotSettable"
	case CIPStatus_PrivilegeViolation:
		return "PrivilegeViolation"
	case CIPStatus_DeviceStateConflict:
		return "DeviceStateConflict"
	case CIPStatus_ReplyDataTooLarge:
		return "ReplyDataTooLarge"
	case CIPStatus_FragmentationOfPrimitiveValue:
		return "FragmentationOfPrimitiveValue"
	case CIPStatus_NotEnoughData:
		return "NotEnoughData"
	case CIPStatus_AttributeNotSupported:
		return "AttributeNotSupported"
	case CIPStatus_TooMuchData:
		return "TooMuchData"
	case CIPStatus_ObjectDoesNotExist:
		return "ObjectDoesNotExist"
	case CIPStatus_ServiceFragmentation:
		return "ServiceFragmentation"
	case CIPStatus_NoStoredAttributeData:
		return "NoStoredAttributeData"
	case CIPStatus_StoreOperationFailure:
		return "StoreOperationFailure"
	case CIPStatus_RequestPacketTooLarge:
		return "RequestPacketTooLarge"
	case CIPStatus_ResponsePacketTooLarge:
		return "ResponsePacketTooLarge"
	case CIPStatus_MissingAttributeListEntryData:
		return "MissingAttributeListEntryData"
	case CIPStatus_InvalidAttributeValueList:
		return "InvalidAttributeValueList"
	case CIPStatus_EmbeddedServiceError:
		return "EmbeddedServiceError"
	case CIPStatus_VendorSpecificError:
		return "VendorSpecificError"
	case CIPStatus_InvalidCommandWithWrongEndianess:
		return "InvalidCommandWithWrongEndianess"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e CIPStatus) String() string {
	return e.PLC4XEnumName()
}

