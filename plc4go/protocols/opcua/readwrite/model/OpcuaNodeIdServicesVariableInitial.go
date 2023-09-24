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

// OpcuaNodeIdServicesVariableInitial is an enum
type OpcuaNodeIdServicesVariableInitial int32

type IOpcuaNodeIdServicesVariableInitial interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableInitial_InitialStateType_StateNumber OpcuaNodeIdServicesVariableInitial = 3736
)

var OpcuaNodeIdServicesVariableInitialValues []OpcuaNodeIdServicesVariableInitial

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableInitialValues = []OpcuaNodeIdServicesVariableInitial {
		OpcuaNodeIdServicesVariableInitial_InitialStateType_StateNumber,
	}
}

func OpcuaNodeIdServicesVariableInitialByValue(value int32) (enum OpcuaNodeIdServicesVariableInitial, ok bool) {
	switch value {
		case 3736:
			return OpcuaNodeIdServicesVariableInitial_InitialStateType_StateNumber, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableInitialByName(value string) (enum OpcuaNodeIdServicesVariableInitial, ok bool) {
	switch value {
	case "InitialStateType_StateNumber":
		return OpcuaNodeIdServicesVariableInitial_InitialStateType_StateNumber, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableInitialKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableInitialValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableInitial(structType any) OpcuaNodeIdServicesVariableInitial {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableInitial {
		if sOpcuaNodeIdServicesVariableInitial, ok := typ.(OpcuaNodeIdServicesVariableInitial); ok {
			return sOpcuaNodeIdServicesVariableInitial
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableInitial) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableInitial) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableInitialParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableInitial, error) {
	return OpcuaNodeIdServicesVariableInitialParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableInitialParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableInitial, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableInitial", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableInitial")
	}
	if enum, ok := OpcuaNodeIdServicesVariableInitialByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableInitial")
		return OpcuaNodeIdServicesVariableInitial(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableInitial) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableInitial) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableInitial", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableInitial) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableInitial_InitialStateType_StateNumber:
		return "InitialStateType_StateNumber"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableInitial) String() string {
	return e.PLC4XEnumName()
}

