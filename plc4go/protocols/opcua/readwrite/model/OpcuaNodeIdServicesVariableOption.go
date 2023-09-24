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

// OpcuaNodeIdServicesVariableOption is an enum
type OpcuaNodeIdServicesVariableOption int32

type IOpcuaNodeIdServicesVariableOption interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableOption_OptionSetType_OptionSetValues OpcuaNodeIdServicesVariableOption = 11488
	OpcuaNodeIdServicesVariableOption_OptionSetType_BitMask OpcuaNodeIdServicesVariableOption = 11701
	OpcuaNodeIdServicesVariableOption_OptionSetValues OpcuaNodeIdServicesVariableOption = 12745
	OpcuaNodeIdServicesVariableOption_OptionSetLength OpcuaNodeIdServicesVariableOption = 32750
)

var OpcuaNodeIdServicesVariableOptionValues []OpcuaNodeIdServicesVariableOption

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableOptionValues = []OpcuaNodeIdServicesVariableOption {
		OpcuaNodeIdServicesVariableOption_OptionSetType_OptionSetValues,
		OpcuaNodeIdServicesVariableOption_OptionSetType_BitMask,
		OpcuaNodeIdServicesVariableOption_OptionSetValues,
		OpcuaNodeIdServicesVariableOption_OptionSetLength,
	}
}

func OpcuaNodeIdServicesVariableOptionByValue(value int32) (enum OpcuaNodeIdServicesVariableOption, ok bool) {
	switch value {
		case 11488:
			return OpcuaNodeIdServicesVariableOption_OptionSetType_OptionSetValues, true
		case 11701:
			return OpcuaNodeIdServicesVariableOption_OptionSetType_BitMask, true
		case 12745:
			return OpcuaNodeIdServicesVariableOption_OptionSetValues, true
		case 32750:
			return OpcuaNodeIdServicesVariableOption_OptionSetLength, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOptionByName(value string) (enum OpcuaNodeIdServicesVariableOption, ok bool) {
	switch value {
	case "OptionSetType_OptionSetValues":
		return OpcuaNodeIdServicesVariableOption_OptionSetType_OptionSetValues, true
	case "OptionSetType_BitMask":
		return OpcuaNodeIdServicesVariableOption_OptionSetType_BitMask, true
	case "OptionSetValues":
		return OpcuaNodeIdServicesVariableOption_OptionSetValues, true
	case "OptionSetLength":
		return OpcuaNodeIdServicesVariableOption_OptionSetLength, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOptionKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableOptionValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableOption(structType any) OpcuaNodeIdServicesVariableOption {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableOption {
		if sOpcuaNodeIdServicesVariableOption, ok := typ.(OpcuaNodeIdServicesVariableOption); ok {
			return sOpcuaNodeIdServicesVariableOption
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableOption) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableOption) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableOptionParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableOption, error) {
	return OpcuaNodeIdServicesVariableOptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableOptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableOption, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableOption", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableOption")
	}
	if enum, ok := OpcuaNodeIdServicesVariableOptionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableOption")
		return OpcuaNodeIdServicesVariableOption(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableOption) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableOption) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableOption", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableOption) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableOption_OptionSetType_OptionSetValues:
		return "OptionSetType_OptionSetValues"
	case OpcuaNodeIdServicesVariableOption_OptionSetType_BitMask:
		return "OptionSetType_BitMask"
	case OpcuaNodeIdServicesVariableOption_OptionSetValues:
		return "OptionSetValues"
	case OpcuaNodeIdServicesVariableOption_OptionSetLength:
		return "OptionSetLength"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableOption) String() string {
	return e.PLC4XEnumName()
}

