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

// OpcuaNodeIdServicesVariableAlias is an enum
type OpcuaNodeIdServicesVariableAlias int32

type IOpcuaNodeIdServicesVariableAlias interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_InputArguments OpcuaNodeIdServicesVariableAlias = 23460
	OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_OutputArguments OpcuaNodeIdServicesVariableAlias = 23461
	OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_FindAlias_InputArguments OpcuaNodeIdServicesVariableAlias = 23463
	OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_FindAlias_OutputArguments OpcuaNodeIdServicesVariableAlias = 23464
)

var OpcuaNodeIdServicesVariableAliasValues []OpcuaNodeIdServicesVariableAlias

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableAliasValues = []OpcuaNodeIdServicesVariableAlias {
		OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_InputArguments,
		OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_OutputArguments,
		OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_FindAlias_InputArguments,
		OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_FindAlias_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableAliasByValue(value int32) (enum OpcuaNodeIdServicesVariableAlias, ok bool) {
	switch value {
		case 23460:
			return OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_InputArguments, true
		case 23461:
			return OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_OutputArguments, true
		case 23463:
			return OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_FindAlias_InputArguments, true
		case 23464:
			return OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_FindAlias_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAliasByName(value string) (enum OpcuaNodeIdServicesVariableAlias, ok bool) {
	switch value {
	case "AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_InputArguments":
		return OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_InputArguments, true
	case "AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_OutputArguments":
		return OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_OutputArguments, true
	case "AliasNameCategoryType_FindAlias_InputArguments":
		return OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_FindAlias_InputArguments, true
	case "AliasNameCategoryType_FindAlias_OutputArguments":
		return OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_FindAlias_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAliasKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableAliasValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableAlias(structType any) OpcuaNodeIdServicesVariableAlias {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableAlias {
		if sOpcuaNodeIdServicesVariableAlias, ok := typ.(OpcuaNodeIdServicesVariableAlias); ok {
			return sOpcuaNodeIdServicesVariableAlias
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableAlias) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableAlias) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableAliasParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableAlias, error) {
	return OpcuaNodeIdServicesVariableAliasParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableAliasParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableAlias, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableAlias", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableAlias")
	}
	if enum, ok := OpcuaNodeIdServicesVariableAliasByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableAlias")
		return OpcuaNodeIdServicesVariableAlias(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableAlias) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableAlias) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableAlias", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableAlias) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_InputArguments:
		return "AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_InputArguments"
	case OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_OutputArguments:
		return "AliasNameCategoryType_SubAliasNameCategories_Placeholder_FindAlias_OutputArguments"
	case OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_FindAlias_InputArguments:
		return "AliasNameCategoryType_FindAlias_InputArguments"
	case OpcuaNodeIdServicesVariableAlias_AliasNameCategoryType_FindAlias_OutputArguments:
		return "AliasNameCategoryType_FindAlias_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableAlias) String() string {
	return e.PLC4XEnumName()
}

