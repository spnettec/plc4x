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

// OpcuaNodeIdServicesVariableArray is an enum
type OpcuaNodeIdServicesVariableArray int32

type IOpcuaNodeIdServicesVariableArray interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableArray_ArrayItemType_Definition OpcuaNodeIdServicesVariableArray = 12022
	OpcuaNodeIdServicesVariableArray_ArrayItemType_ValuePrecision OpcuaNodeIdServicesVariableArray = 12023
	OpcuaNodeIdServicesVariableArray_ArrayItemType_InstrumentRange OpcuaNodeIdServicesVariableArray = 12024
	OpcuaNodeIdServicesVariableArray_ArrayItemType_EURange OpcuaNodeIdServicesVariableArray = 12025
	OpcuaNodeIdServicesVariableArray_ArrayItemType_EngineeringUnits OpcuaNodeIdServicesVariableArray = 12026
	OpcuaNodeIdServicesVariableArray_ArrayItemType_Title OpcuaNodeIdServicesVariableArray = 12027
	OpcuaNodeIdServicesVariableArray_ArrayItemType_AxisScaleType OpcuaNodeIdServicesVariableArray = 12028
)

var OpcuaNodeIdServicesVariableArrayValues []OpcuaNodeIdServicesVariableArray

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableArrayValues = []OpcuaNodeIdServicesVariableArray {
		OpcuaNodeIdServicesVariableArray_ArrayItemType_Definition,
		OpcuaNodeIdServicesVariableArray_ArrayItemType_ValuePrecision,
		OpcuaNodeIdServicesVariableArray_ArrayItemType_InstrumentRange,
		OpcuaNodeIdServicesVariableArray_ArrayItemType_EURange,
		OpcuaNodeIdServicesVariableArray_ArrayItemType_EngineeringUnits,
		OpcuaNodeIdServicesVariableArray_ArrayItemType_Title,
		OpcuaNodeIdServicesVariableArray_ArrayItemType_AxisScaleType,
	}
}

func OpcuaNodeIdServicesVariableArrayByValue(value int32) (enum OpcuaNodeIdServicesVariableArray, ok bool) {
	switch value {
		case 12022:
			return OpcuaNodeIdServicesVariableArray_ArrayItemType_Definition, true
		case 12023:
			return OpcuaNodeIdServicesVariableArray_ArrayItemType_ValuePrecision, true
		case 12024:
			return OpcuaNodeIdServicesVariableArray_ArrayItemType_InstrumentRange, true
		case 12025:
			return OpcuaNodeIdServicesVariableArray_ArrayItemType_EURange, true
		case 12026:
			return OpcuaNodeIdServicesVariableArray_ArrayItemType_EngineeringUnits, true
		case 12027:
			return OpcuaNodeIdServicesVariableArray_ArrayItemType_Title, true
		case 12028:
			return OpcuaNodeIdServicesVariableArray_ArrayItemType_AxisScaleType, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableArrayByName(value string) (enum OpcuaNodeIdServicesVariableArray, ok bool) {
	switch value {
	case "ArrayItemType_Definition":
		return OpcuaNodeIdServicesVariableArray_ArrayItemType_Definition, true
	case "ArrayItemType_ValuePrecision":
		return OpcuaNodeIdServicesVariableArray_ArrayItemType_ValuePrecision, true
	case "ArrayItemType_InstrumentRange":
		return OpcuaNodeIdServicesVariableArray_ArrayItemType_InstrumentRange, true
	case "ArrayItemType_EURange":
		return OpcuaNodeIdServicesVariableArray_ArrayItemType_EURange, true
	case "ArrayItemType_EngineeringUnits":
		return OpcuaNodeIdServicesVariableArray_ArrayItemType_EngineeringUnits, true
	case "ArrayItemType_Title":
		return OpcuaNodeIdServicesVariableArray_ArrayItemType_Title, true
	case "ArrayItemType_AxisScaleType":
		return OpcuaNodeIdServicesVariableArray_ArrayItemType_AxisScaleType, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableArrayKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableArrayValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableArray(structType any) OpcuaNodeIdServicesVariableArray {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableArray {
		if sOpcuaNodeIdServicesVariableArray, ok := typ.(OpcuaNodeIdServicesVariableArray); ok {
			return sOpcuaNodeIdServicesVariableArray
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableArray) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableArray) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableArrayParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableArray, error) {
	return OpcuaNodeIdServicesVariableArrayParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableArrayParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableArray, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableArray", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableArray")
	}
	if enum, ok := OpcuaNodeIdServicesVariableArrayByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableArray")
		return OpcuaNodeIdServicesVariableArray(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableArray) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableArray) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableArray", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableArray) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableArray_ArrayItemType_Definition:
		return "ArrayItemType_Definition"
	case OpcuaNodeIdServicesVariableArray_ArrayItemType_ValuePrecision:
		return "ArrayItemType_ValuePrecision"
	case OpcuaNodeIdServicesVariableArray_ArrayItemType_InstrumentRange:
		return "ArrayItemType_InstrumentRange"
	case OpcuaNodeIdServicesVariableArray_ArrayItemType_EURange:
		return "ArrayItemType_EURange"
	case OpcuaNodeIdServicesVariableArray_ArrayItemType_EngineeringUnits:
		return "ArrayItemType_EngineeringUnits"
	case OpcuaNodeIdServicesVariableArray_ArrayItemType_Title:
		return "ArrayItemType_Title"
	case OpcuaNodeIdServicesVariableArray_ArrayItemType_AxisScaleType:
		return "ArrayItemType_AxisScaleType"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableArray) String() string {
	return e.PLC4XEnumName()
}

