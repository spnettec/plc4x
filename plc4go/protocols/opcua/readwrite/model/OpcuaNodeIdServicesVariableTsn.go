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

// OpcuaNodeIdServicesVariableTsn is an enum
type OpcuaNodeIdServicesVariableTsn int32

type IOpcuaNodeIdServicesVariableTsn interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableTsn_TsnFailureCode_EnumValues OpcuaNodeIdServicesVariableTsn = 24239
	OpcuaNodeIdServicesVariableTsn_TsnStreamState_EnumValues OpcuaNodeIdServicesVariableTsn = 24240
	OpcuaNodeIdServicesVariableTsn_TsnTalkerStatus_EnumValues OpcuaNodeIdServicesVariableTsn = 24241
	OpcuaNodeIdServicesVariableTsn_TsnListenerStatus_EnumValues OpcuaNodeIdServicesVariableTsn = 24242
)

var OpcuaNodeIdServicesVariableTsnValues []OpcuaNodeIdServicesVariableTsn

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableTsnValues = []OpcuaNodeIdServicesVariableTsn {
		OpcuaNodeIdServicesVariableTsn_TsnFailureCode_EnumValues,
		OpcuaNodeIdServicesVariableTsn_TsnStreamState_EnumValues,
		OpcuaNodeIdServicesVariableTsn_TsnTalkerStatus_EnumValues,
		OpcuaNodeIdServicesVariableTsn_TsnListenerStatus_EnumValues,
	}
}

func OpcuaNodeIdServicesVariableTsnByValue(value int32) (enum OpcuaNodeIdServicesVariableTsn, ok bool) {
	switch value {
		case 24239:
			return OpcuaNodeIdServicesVariableTsn_TsnFailureCode_EnumValues, true
		case 24240:
			return OpcuaNodeIdServicesVariableTsn_TsnStreamState_EnumValues, true
		case 24241:
			return OpcuaNodeIdServicesVariableTsn_TsnTalkerStatus_EnumValues, true
		case 24242:
			return OpcuaNodeIdServicesVariableTsn_TsnListenerStatus_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTsnByName(value string) (enum OpcuaNodeIdServicesVariableTsn, ok bool) {
	switch value {
	case "TsnFailureCode_EnumValues":
		return OpcuaNodeIdServicesVariableTsn_TsnFailureCode_EnumValues, true
	case "TsnStreamState_EnumValues":
		return OpcuaNodeIdServicesVariableTsn_TsnStreamState_EnumValues, true
	case "TsnTalkerStatus_EnumValues":
		return OpcuaNodeIdServicesVariableTsn_TsnTalkerStatus_EnumValues, true
	case "TsnListenerStatus_EnumValues":
		return OpcuaNodeIdServicesVariableTsn_TsnListenerStatus_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTsnKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableTsnValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableTsn(structType any) OpcuaNodeIdServicesVariableTsn {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableTsn {
		if sOpcuaNodeIdServicesVariableTsn, ok := typ.(OpcuaNodeIdServicesVariableTsn); ok {
			return sOpcuaNodeIdServicesVariableTsn
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableTsn) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableTsn) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableTsnParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableTsn, error) {
	return OpcuaNodeIdServicesVariableTsnParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableTsnParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableTsn, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableTsn", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableTsn")
	}
	if enum, ok := OpcuaNodeIdServicesVariableTsnByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableTsn")
		return OpcuaNodeIdServicesVariableTsn(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableTsn) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableTsn) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableTsn", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableTsn) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableTsn_TsnFailureCode_EnumValues:
		return "TsnFailureCode_EnumValues"
	case OpcuaNodeIdServicesVariableTsn_TsnStreamState_EnumValues:
		return "TsnStreamState_EnumValues"
	case OpcuaNodeIdServicesVariableTsn_TsnTalkerStatus_EnumValues:
		return "TsnTalkerStatus_EnumValues"
	case OpcuaNodeIdServicesVariableTsn_TsnListenerStatus_EnumValues:
		return "TsnListenerStatus_EnumValues"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableTsn) String() string {
	return e.PLC4XEnumName()
}

