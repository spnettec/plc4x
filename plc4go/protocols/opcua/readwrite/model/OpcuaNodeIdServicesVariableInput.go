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

// OpcuaNodeIdServicesVariableInput is an enum
type OpcuaNodeIdServicesVariableInput int32

type IOpcuaNodeIdServicesVariableInput interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableInput_InputArguments OpcuaNodeIdServicesVariableInput = 3072
)

var OpcuaNodeIdServicesVariableInputValues []OpcuaNodeIdServicesVariableInput

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableInputValues = []OpcuaNodeIdServicesVariableInput{
		OpcuaNodeIdServicesVariableInput_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableInputByValue(value int32) (enum OpcuaNodeIdServicesVariableInput, ok bool) {
	switch value {
	case 3072:
		return OpcuaNodeIdServicesVariableInput_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableInputByName(value string) (enum OpcuaNodeIdServicesVariableInput, ok bool) {
	switch value {
	case "InputArguments":
		return OpcuaNodeIdServicesVariableInput_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableInputKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableInputValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableInput(structType any) OpcuaNodeIdServicesVariableInput {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableInput {
		if sOpcuaNodeIdServicesVariableInput, ok := typ.(OpcuaNodeIdServicesVariableInput); ok {
			return sOpcuaNodeIdServicesVariableInput
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableInput) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableInput) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableInputParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableInput, error) {
	return OpcuaNodeIdServicesVariableInputParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableInputParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableInput, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableInput", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableInput")
	}
	if enum, ok := OpcuaNodeIdServicesVariableInputByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableInput")
		return OpcuaNodeIdServicesVariableInput(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableInput) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableInput) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableInput", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableInput) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableInput_InputArguments:
		return "InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableInput) String() string {
	return e.PLC4XEnumName()
}
