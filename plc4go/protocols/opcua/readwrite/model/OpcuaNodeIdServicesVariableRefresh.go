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

// OpcuaNodeIdServicesVariableRefresh is an enum
type OpcuaNodeIdServicesVariableRefresh int32

type IOpcuaNodeIdServicesVariableRefresh interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_EventId OpcuaNodeIdServicesVariableRefresh = 3969
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_EventType OpcuaNodeIdServicesVariableRefresh = 3970
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_SourceNode OpcuaNodeIdServicesVariableRefresh = 3971
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_SourceName OpcuaNodeIdServicesVariableRefresh = 3972
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Time OpcuaNodeIdServicesVariableRefresh = 3973
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ReceiveTime OpcuaNodeIdServicesVariableRefresh = 3974
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_LocalTime OpcuaNodeIdServicesVariableRefresh = 3975
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Message OpcuaNodeIdServicesVariableRefresh = 3976
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Severity OpcuaNodeIdServicesVariableRefresh = 3977
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_EventId OpcuaNodeIdServicesVariableRefresh = 3978
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_EventType OpcuaNodeIdServicesVariableRefresh = 3979
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_SourceNode OpcuaNodeIdServicesVariableRefresh = 3980
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_SourceName OpcuaNodeIdServicesVariableRefresh = 3981
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Time OpcuaNodeIdServicesVariableRefresh = 3982
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ReceiveTime OpcuaNodeIdServicesVariableRefresh = 3983
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_LocalTime OpcuaNodeIdServicesVariableRefresh = 3984
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Message OpcuaNodeIdServicesVariableRefresh = 3985
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Severity OpcuaNodeIdServicesVariableRefresh = 3986
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_EventId OpcuaNodeIdServicesVariableRefresh = 3987
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_EventType OpcuaNodeIdServicesVariableRefresh = 3988
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_SourceNode OpcuaNodeIdServicesVariableRefresh = 3989
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_SourceName OpcuaNodeIdServicesVariableRefresh = 3990
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Time OpcuaNodeIdServicesVariableRefresh = 3991
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ReceiveTime OpcuaNodeIdServicesVariableRefresh = 3992
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_LocalTime OpcuaNodeIdServicesVariableRefresh = 3993
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Message OpcuaNodeIdServicesVariableRefresh = 3994
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Severity OpcuaNodeIdServicesVariableRefresh = 3995
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionClassId OpcuaNodeIdServicesVariableRefresh = 31975
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionClassName OpcuaNodeIdServicesVariableRefresh = 31976
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionSubClassId OpcuaNodeIdServicesVariableRefresh = 31977
	OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionSubClassName OpcuaNodeIdServicesVariableRefresh = 31978
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionClassId OpcuaNodeIdServicesVariableRefresh = 31979
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionClassName OpcuaNodeIdServicesVariableRefresh = 31980
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionSubClassId OpcuaNodeIdServicesVariableRefresh = 31981
	OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionSubClassName OpcuaNodeIdServicesVariableRefresh = 31982
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionClassId OpcuaNodeIdServicesVariableRefresh = 31983
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionClassName OpcuaNodeIdServicesVariableRefresh = 31984
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionSubClassId OpcuaNodeIdServicesVariableRefresh = 31985
	OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionSubClassName OpcuaNodeIdServicesVariableRefresh = 31986
)

var OpcuaNodeIdServicesVariableRefreshValues []OpcuaNodeIdServicesVariableRefresh

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableRefreshValues = []OpcuaNodeIdServicesVariableRefresh {
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_EventId,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_EventType,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_SourceNode,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_SourceName,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Time,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ReceiveTime,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_LocalTime,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Message,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Severity,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_EventId,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_EventType,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_SourceNode,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_SourceName,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Time,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ReceiveTime,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_LocalTime,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Message,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Severity,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_EventId,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_EventType,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_SourceNode,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_SourceName,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Time,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ReceiveTime,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_LocalTime,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Message,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Severity,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionClassId,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionClassName,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionSubClassId,
		OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionSubClassName,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionClassId,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionClassName,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionSubClassId,
		OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionSubClassName,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionClassId,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionClassName,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionSubClassId,
		OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionSubClassName,
	}
}

func OpcuaNodeIdServicesVariableRefreshByValue(value int32) (enum OpcuaNodeIdServicesVariableRefresh, ok bool) {
	switch value {
		case 31975:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionClassId, true
		case 31976:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionClassName, true
		case 31977:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionSubClassId, true
		case 31978:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionSubClassName, true
		case 31979:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionClassId, true
		case 31980:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionClassName, true
		case 31981:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionSubClassId, true
		case 31982:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionSubClassName, true
		case 31983:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionClassId, true
		case 31984:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionClassName, true
		case 31985:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionSubClassId, true
		case 31986:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionSubClassName, true
		case 3969:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_EventId, true
		case 3970:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_EventType, true
		case 3971:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_SourceNode, true
		case 3972:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_SourceName, true
		case 3973:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Time, true
		case 3974:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ReceiveTime, true
		case 3975:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_LocalTime, true
		case 3976:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Message, true
		case 3977:
			return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Severity, true
		case 3978:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_EventId, true
		case 3979:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_EventType, true
		case 3980:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_SourceNode, true
		case 3981:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_SourceName, true
		case 3982:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Time, true
		case 3983:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ReceiveTime, true
		case 3984:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_LocalTime, true
		case 3985:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Message, true
		case 3986:
			return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Severity, true
		case 3987:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_EventId, true
		case 3988:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_EventType, true
		case 3989:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_SourceNode, true
		case 3990:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_SourceName, true
		case 3991:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Time, true
		case 3992:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ReceiveTime, true
		case 3993:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_LocalTime, true
		case 3994:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Message, true
		case 3995:
			return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Severity, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableRefreshByName(value string) (enum OpcuaNodeIdServicesVariableRefresh, ok bool) {
	switch value {
	case "RefreshStartEventType_ConditionClassId":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionClassId, true
	case "RefreshStartEventType_ConditionClassName":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionClassName, true
	case "RefreshStartEventType_ConditionSubClassId":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionSubClassId, true
	case "RefreshStartEventType_ConditionSubClassName":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionSubClassName, true
	case "RefreshEndEventType_ConditionClassId":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionClassId, true
	case "RefreshEndEventType_ConditionClassName":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionClassName, true
	case "RefreshEndEventType_ConditionSubClassId":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionSubClassId, true
	case "RefreshEndEventType_ConditionSubClassName":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionSubClassName, true
	case "RefreshRequiredEventType_ConditionClassId":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionClassId, true
	case "RefreshRequiredEventType_ConditionClassName":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionClassName, true
	case "RefreshRequiredEventType_ConditionSubClassId":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionSubClassId, true
	case "RefreshRequiredEventType_ConditionSubClassName":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionSubClassName, true
	case "RefreshStartEventType_EventId":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_EventId, true
	case "RefreshStartEventType_EventType":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_EventType, true
	case "RefreshStartEventType_SourceNode":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_SourceNode, true
	case "RefreshStartEventType_SourceName":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_SourceName, true
	case "RefreshStartEventType_Time":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Time, true
	case "RefreshStartEventType_ReceiveTime":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ReceiveTime, true
	case "RefreshStartEventType_LocalTime":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_LocalTime, true
	case "RefreshStartEventType_Message":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Message, true
	case "RefreshStartEventType_Severity":
		return OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Severity, true
	case "RefreshEndEventType_EventId":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_EventId, true
	case "RefreshEndEventType_EventType":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_EventType, true
	case "RefreshEndEventType_SourceNode":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_SourceNode, true
	case "RefreshEndEventType_SourceName":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_SourceName, true
	case "RefreshEndEventType_Time":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Time, true
	case "RefreshEndEventType_ReceiveTime":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ReceiveTime, true
	case "RefreshEndEventType_LocalTime":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_LocalTime, true
	case "RefreshEndEventType_Message":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Message, true
	case "RefreshEndEventType_Severity":
		return OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Severity, true
	case "RefreshRequiredEventType_EventId":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_EventId, true
	case "RefreshRequiredEventType_EventType":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_EventType, true
	case "RefreshRequiredEventType_SourceNode":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_SourceNode, true
	case "RefreshRequiredEventType_SourceName":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_SourceName, true
	case "RefreshRequiredEventType_Time":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Time, true
	case "RefreshRequiredEventType_ReceiveTime":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ReceiveTime, true
	case "RefreshRequiredEventType_LocalTime":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_LocalTime, true
	case "RefreshRequiredEventType_Message":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Message, true
	case "RefreshRequiredEventType_Severity":
		return OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Severity, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableRefreshKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableRefreshValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableRefresh(structType any) OpcuaNodeIdServicesVariableRefresh {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableRefresh {
		if sOpcuaNodeIdServicesVariableRefresh, ok := typ.(OpcuaNodeIdServicesVariableRefresh); ok {
			return sOpcuaNodeIdServicesVariableRefresh
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableRefresh) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableRefresh) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableRefreshParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableRefresh, error) {
	return OpcuaNodeIdServicesVariableRefreshParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableRefreshParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableRefresh, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableRefresh", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableRefresh")
	}
	if enum, ok := OpcuaNodeIdServicesVariableRefreshByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableRefresh")
		return OpcuaNodeIdServicesVariableRefresh(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableRefresh) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableRefresh) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableRefresh", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableRefresh) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionClassId:
		return "RefreshStartEventType_ConditionClassId"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionClassName:
		return "RefreshStartEventType_ConditionClassName"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionSubClassId:
		return "RefreshStartEventType_ConditionSubClassId"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ConditionSubClassName:
		return "RefreshStartEventType_ConditionSubClassName"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionClassId:
		return "RefreshEndEventType_ConditionClassId"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionClassName:
		return "RefreshEndEventType_ConditionClassName"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionSubClassId:
		return "RefreshEndEventType_ConditionSubClassId"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ConditionSubClassName:
		return "RefreshEndEventType_ConditionSubClassName"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionClassId:
		return "RefreshRequiredEventType_ConditionClassId"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionClassName:
		return "RefreshRequiredEventType_ConditionClassName"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionSubClassId:
		return "RefreshRequiredEventType_ConditionSubClassId"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ConditionSubClassName:
		return "RefreshRequiredEventType_ConditionSubClassName"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_EventId:
		return "RefreshStartEventType_EventId"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_EventType:
		return "RefreshStartEventType_EventType"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_SourceNode:
		return "RefreshStartEventType_SourceNode"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_SourceName:
		return "RefreshStartEventType_SourceName"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Time:
		return "RefreshStartEventType_Time"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_ReceiveTime:
		return "RefreshStartEventType_ReceiveTime"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_LocalTime:
		return "RefreshStartEventType_LocalTime"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Message:
		return "RefreshStartEventType_Message"
	case OpcuaNodeIdServicesVariableRefresh_RefreshStartEventType_Severity:
		return "RefreshStartEventType_Severity"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_EventId:
		return "RefreshEndEventType_EventId"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_EventType:
		return "RefreshEndEventType_EventType"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_SourceNode:
		return "RefreshEndEventType_SourceNode"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_SourceName:
		return "RefreshEndEventType_SourceName"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Time:
		return "RefreshEndEventType_Time"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_ReceiveTime:
		return "RefreshEndEventType_ReceiveTime"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_LocalTime:
		return "RefreshEndEventType_LocalTime"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Message:
		return "RefreshEndEventType_Message"
	case OpcuaNodeIdServicesVariableRefresh_RefreshEndEventType_Severity:
		return "RefreshEndEventType_Severity"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_EventId:
		return "RefreshRequiredEventType_EventId"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_EventType:
		return "RefreshRequiredEventType_EventType"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_SourceNode:
		return "RefreshRequiredEventType_SourceNode"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_SourceName:
		return "RefreshRequiredEventType_SourceName"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Time:
		return "RefreshRequiredEventType_Time"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_ReceiveTime:
		return "RefreshRequiredEventType_ReceiveTime"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_LocalTime:
		return "RefreshRequiredEventType_LocalTime"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Message:
		return "RefreshRequiredEventType_Message"
	case OpcuaNodeIdServicesVariableRefresh_RefreshRequiredEventType_Severity:
		return "RefreshRequiredEventType_Severity"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableRefresh) String() string {
	return e.PLC4XEnumName()
}

