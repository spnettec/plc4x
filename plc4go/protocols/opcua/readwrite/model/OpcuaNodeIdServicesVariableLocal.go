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

// OpcuaNodeIdServicesVariableLocal is an enum
type OpcuaNodeIdServicesVariableLocal int32

type IOpcuaNodeIdServicesVariableLocal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableLocal_LocalTime OpcuaNodeIdServicesVariableLocal = 3069
)

var OpcuaNodeIdServicesVariableLocalValues []OpcuaNodeIdServicesVariableLocal

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableLocalValues = []OpcuaNodeIdServicesVariableLocal {
		OpcuaNodeIdServicesVariableLocal_LocalTime,
	}
}

func OpcuaNodeIdServicesVariableLocalByValue(value int32) (enum OpcuaNodeIdServicesVariableLocal, ok bool) {
	switch value {
		case 3069:
			return OpcuaNodeIdServicesVariableLocal_LocalTime, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableLocalByName(value string) (enum OpcuaNodeIdServicesVariableLocal, ok bool) {
	switch value {
	case "LocalTime":
		return OpcuaNodeIdServicesVariableLocal_LocalTime, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableLocalKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableLocalValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableLocal(structType any) OpcuaNodeIdServicesVariableLocal {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableLocal {
		if sOpcuaNodeIdServicesVariableLocal, ok := typ.(OpcuaNodeIdServicesVariableLocal); ok {
			return sOpcuaNodeIdServicesVariableLocal
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableLocal) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableLocal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableLocalParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableLocal, error) {
	return OpcuaNodeIdServicesVariableLocalParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableLocalParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableLocal, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableLocal", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableLocal")
	}
	if enum, ok := OpcuaNodeIdServicesVariableLocalByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableLocal")
		return OpcuaNodeIdServicesVariableLocal(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableLocal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableLocal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableLocal", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableLocal) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableLocal_LocalTime:
		return "LocalTime"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableLocal) String() string {
	return e.PLC4XEnumName()
}

