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

// OpcuaNodeIdServicesVariableDuplex is an enum
type OpcuaNodeIdServicesVariableDuplex int32

type IOpcuaNodeIdServicesVariableDuplex interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableDuplex_Duplex_EnumValues OpcuaNodeIdServicesVariableDuplex = 24235
)

var OpcuaNodeIdServicesVariableDuplexValues []OpcuaNodeIdServicesVariableDuplex

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableDuplexValues = []OpcuaNodeIdServicesVariableDuplex{
		OpcuaNodeIdServicesVariableDuplex_Duplex_EnumValues,
	}
}

func OpcuaNodeIdServicesVariableDuplexByValue(value int32) (enum OpcuaNodeIdServicesVariableDuplex, ok bool) {
	switch value {
	case 24235:
		return OpcuaNodeIdServicesVariableDuplex_Duplex_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDuplexByName(value string) (enum OpcuaNodeIdServicesVariableDuplex, ok bool) {
	switch value {
	case "Duplex_EnumValues":
		return OpcuaNodeIdServicesVariableDuplex_Duplex_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDuplexKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableDuplexValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableDuplex(structType any) OpcuaNodeIdServicesVariableDuplex {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableDuplex {
		if sOpcuaNodeIdServicesVariableDuplex, ok := typ.(OpcuaNodeIdServicesVariableDuplex); ok {
			return sOpcuaNodeIdServicesVariableDuplex
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableDuplex) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableDuplex) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableDuplexParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableDuplex, error) {
	return OpcuaNodeIdServicesVariableDuplexParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableDuplexParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableDuplex, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableDuplex", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableDuplex")
	}
	if enum, ok := OpcuaNodeIdServicesVariableDuplexByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableDuplex")
		return OpcuaNodeIdServicesVariableDuplex(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableDuplex) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableDuplex) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableDuplex", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableDuplex) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableDuplex_Duplex_EnumValues:
		return "Duplex_EnumValues"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableDuplex) String() string {
	return e.PLC4XEnumName()
}
