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

// OpcuaNodeIdServicesVariableSubscribed is an enum
type OpcuaNodeIdServicesVariableSubscribed int32

type IOpcuaNodeIdServicesVariableSubscribed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_InputArguments    OpcuaNodeIdServicesVariableSubscribed = 23798
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_OutputArguments   OpcuaNodeIdServicesVariableSubscribed = 23799
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveSubscribedDataSet_InputArguments OpcuaNodeIdServicesVariableSubscribed = 23801
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_InputArguments        OpcuaNodeIdServicesVariableSubscribed = 23803
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_OutputArguments       OpcuaNodeIdServicesVariableSubscribed = 23804
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveDataSetFolder_InputArguments     OpcuaNodeIdServicesVariableSubscribed = 23806
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_DataSetMetaData                    OpcuaNodeIdServicesVariableSubscribed = 23809
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_IsConnected                        OpcuaNodeIdServicesVariableSubscribed = 23810
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddSubscribedDataSet_InputArguments                                            OpcuaNodeIdServicesVariableSubscribed = 23812
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddSubscribedDataSet_OutputArguments                                           OpcuaNodeIdServicesVariableSubscribed = 23813
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_RemoveSubscribedDataSet_InputArguments                                         OpcuaNodeIdServicesVariableSubscribed = 23815
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddDataSetFolder_InputArguments                                                OpcuaNodeIdServicesVariableSubscribed = 23817
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddDataSetFolder_OutputArguments                                               OpcuaNodeIdServicesVariableSubscribed = 23818
	OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_RemoveDataSetFolder_InputArguments                                             OpcuaNodeIdServicesVariableSubscribed = 23820
)

var OpcuaNodeIdServicesVariableSubscribedValues []OpcuaNodeIdServicesVariableSubscribed

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableSubscribedValues = []OpcuaNodeIdServicesVariableSubscribed{
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_InputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_OutputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveSubscribedDataSet_InputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_InputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_OutputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveDataSetFolder_InputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_DataSetMetaData,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_IsConnected,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddSubscribedDataSet_InputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddSubscribedDataSet_OutputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_RemoveSubscribedDataSet_InputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddDataSetFolder_InputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddDataSetFolder_OutputArguments,
		OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_RemoveDataSetFolder_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableSubscribedByValue(value int32) (enum OpcuaNodeIdServicesVariableSubscribed, ok bool) {
	switch value {
	case 23798:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_InputArguments, true
	case 23799:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_OutputArguments, true
	case 23801:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveSubscribedDataSet_InputArguments, true
	case 23803:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_InputArguments, true
	case 23804:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_OutputArguments, true
	case 23806:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveDataSetFolder_InputArguments, true
	case 23809:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_DataSetMetaData, true
	case 23810:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_IsConnected, true
	case 23812:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddSubscribedDataSet_InputArguments, true
	case 23813:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddSubscribedDataSet_OutputArguments, true
	case 23815:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_RemoveSubscribedDataSet_InputArguments, true
	case 23817:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddDataSetFolder_InputArguments, true
	case 23818:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddDataSetFolder_OutputArguments, true
	case 23820:
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_RemoveDataSetFolder_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableSubscribedByName(value string) (enum OpcuaNodeIdServicesVariableSubscribed, ok bool) {
	switch value {
	case "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_InputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_InputArguments, true
	case "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_OutputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_OutputArguments, true
	case "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveSubscribedDataSet_InputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveSubscribedDataSet_InputArguments, true
	case "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_InputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_InputArguments, true
	case "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_OutputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_OutputArguments, true
	case "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveDataSetFolder_InputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveDataSetFolder_InputArguments, true
	case "SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_DataSetMetaData":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_DataSetMetaData, true
	case "SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_IsConnected":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_IsConnected, true
	case "SubscribedDataSetFolderType_AddSubscribedDataSet_InputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddSubscribedDataSet_InputArguments, true
	case "SubscribedDataSetFolderType_AddSubscribedDataSet_OutputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddSubscribedDataSet_OutputArguments, true
	case "SubscribedDataSetFolderType_RemoveSubscribedDataSet_InputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_RemoveSubscribedDataSet_InputArguments, true
	case "SubscribedDataSetFolderType_AddDataSetFolder_InputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddDataSetFolder_InputArguments, true
	case "SubscribedDataSetFolderType_AddDataSetFolder_OutputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddDataSetFolder_OutputArguments, true
	case "SubscribedDataSetFolderType_RemoveDataSetFolder_InputArguments":
		return OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_RemoveDataSetFolder_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableSubscribedKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableSubscribedValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableSubscribed(structType any) OpcuaNodeIdServicesVariableSubscribed {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableSubscribed {
		if sOpcuaNodeIdServicesVariableSubscribed, ok := typ.(OpcuaNodeIdServicesVariableSubscribed); ok {
			return sOpcuaNodeIdServicesVariableSubscribed
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableSubscribed) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableSubscribed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableSubscribedParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableSubscribed, error) {
	return OpcuaNodeIdServicesVariableSubscribedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableSubscribedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableSubscribed, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableSubscribed", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableSubscribed")
	}
	if enum, ok := OpcuaNodeIdServicesVariableSubscribedByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableSubscribed")
		return OpcuaNodeIdServicesVariableSubscribed(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableSubscribed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableSubscribed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableSubscribed", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableSubscribed) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_InputArguments:
		return "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_InputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_OutputArguments:
		return "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_OutputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveSubscribedDataSet_InputArguments:
		return "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveSubscribedDataSet_InputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_InputArguments:
		return "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_InputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_OutputArguments:
		return "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_OutputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveDataSetFolder_InputArguments:
		return "SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveDataSetFolder_InputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_DataSetMetaData:
		return "SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_DataSetMetaData"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_IsConnected:
		return "SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_IsConnected"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddSubscribedDataSet_InputArguments:
		return "SubscribedDataSetFolderType_AddSubscribedDataSet_InputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddSubscribedDataSet_OutputArguments:
		return "SubscribedDataSetFolderType_AddSubscribedDataSet_OutputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_RemoveSubscribedDataSet_InputArguments:
		return "SubscribedDataSetFolderType_RemoveSubscribedDataSet_InputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddDataSetFolder_InputArguments:
		return "SubscribedDataSetFolderType_AddDataSetFolder_InputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_AddDataSetFolder_OutputArguments:
		return "SubscribedDataSetFolderType_AddDataSetFolder_OutputArguments"
	case OpcuaNodeIdServicesVariableSubscribed_SubscribedDataSetFolderType_RemoveDataSetFolder_InputArguments:
		return "SubscribedDataSetFolderType_RemoveDataSetFolder_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableSubscribed) String() string {
	return e.PLC4XEnumName()
}
