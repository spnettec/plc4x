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

// OpcuaNodeIdServicesVariableNegotiation is an enum
type OpcuaNodeIdServicesVariableNegotiation int32

type IOpcuaNodeIdServicesVariableNegotiation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableNegotiation_NegotiationStatus_EnumValues OpcuaNodeIdServicesVariableNegotiation = 24238
)

var OpcuaNodeIdServicesVariableNegotiationValues []OpcuaNodeIdServicesVariableNegotiation

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableNegotiationValues = []OpcuaNodeIdServicesVariableNegotiation {
		OpcuaNodeIdServicesVariableNegotiation_NegotiationStatus_EnumValues,
	}
}

func OpcuaNodeIdServicesVariableNegotiationByValue(value int32) (enum OpcuaNodeIdServicesVariableNegotiation, ok bool) {
	switch value {
		case 24238:
			return OpcuaNodeIdServicesVariableNegotiation_NegotiationStatus_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNegotiationByName(value string) (enum OpcuaNodeIdServicesVariableNegotiation, ok bool) {
	switch value {
	case "NegotiationStatus_EnumValues":
		return OpcuaNodeIdServicesVariableNegotiation_NegotiationStatus_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNegotiationKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableNegotiationValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableNegotiation(structType any) OpcuaNodeIdServicesVariableNegotiation {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableNegotiation {
		if sOpcuaNodeIdServicesVariableNegotiation, ok := typ.(OpcuaNodeIdServicesVariableNegotiation); ok {
			return sOpcuaNodeIdServicesVariableNegotiation
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableNegotiation) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableNegotiation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableNegotiationParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableNegotiation, error) {
	return OpcuaNodeIdServicesVariableNegotiationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableNegotiationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableNegotiation, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableNegotiation", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableNegotiation")
	}
	if enum, ok := OpcuaNodeIdServicesVariableNegotiationByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableNegotiation")
		return OpcuaNodeIdServicesVariableNegotiation(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableNegotiation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableNegotiation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableNegotiation", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableNegotiation) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableNegotiation_NegotiationStatus_EnumValues:
		return "NegotiationStatus_EnumValues"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableNegotiation) String() string {
	return e.PLC4XEnumName()
}

