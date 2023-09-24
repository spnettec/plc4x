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

// OpcuaNodeIdServicesVariableGenerate is an enum
type OpcuaNodeIdServicesVariableGenerate int32

type IOpcuaNodeIdServicesVariableGenerate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableGenerate_GenerateFileForReadMethodType_InputArguments OpcuaNodeIdServicesVariableGenerate = 15796
	OpcuaNodeIdServicesVariableGenerate_GenerateFileForReadMethodType_OutputArguments OpcuaNodeIdServicesVariableGenerate = 15797
	OpcuaNodeIdServicesVariableGenerate_GenerateFileForWriteMethodType_OutputArguments OpcuaNodeIdServicesVariableGenerate = 15799
	OpcuaNodeIdServicesVariableGenerate_GenerateFileForWriteMethodType_InputArguments OpcuaNodeIdServicesVariableGenerate = 16360
)

var OpcuaNodeIdServicesVariableGenerateValues []OpcuaNodeIdServicesVariableGenerate

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableGenerateValues = []OpcuaNodeIdServicesVariableGenerate {
		OpcuaNodeIdServicesVariableGenerate_GenerateFileForReadMethodType_InputArguments,
		OpcuaNodeIdServicesVariableGenerate_GenerateFileForReadMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGenerate_GenerateFileForWriteMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableGenerate_GenerateFileForWriteMethodType_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableGenerateByValue(value int32) (enum OpcuaNodeIdServicesVariableGenerate, ok bool) {
	switch value {
		case 15796:
			return OpcuaNodeIdServicesVariableGenerate_GenerateFileForReadMethodType_InputArguments, true
		case 15797:
			return OpcuaNodeIdServicesVariableGenerate_GenerateFileForReadMethodType_OutputArguments, true
		case 15799:
			return OpcuaNodeIdServicesVariableGenerate_GenerateFileForWriteMethodType_OutputArguments, true
		case 16360:
			return OpcuaNodeIdServicesVariableGenerate_GenerateFileForWriteMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableGenerateByName(value string) (enum OpcuaNodeIdServicesVariableGenerate, ok bool) {
	switch value {
	case "GenerateFileForReadMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGenerate_GenerateFileForReadMethodType_InputArguments, true
	case "GenerateFileForReadMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGenerate_GenerateFileForReadMethodType_OutputArguments, true
	case "GenerateFileForWriteMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableGenerate_GenerateFileForWriteMethodType_OutputArguments, true
	case "GenerateFileForWriteMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableGenerate_GenerateFileForWriteMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableGenerateKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableGenerateValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableGenerate(structType any) OpcuaNodeIdServicesVariableGenerate {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableGenerate {
		if sOpcuaNodeIdServicesVariableGenerate, ok := typ.(OpcuaNodeIdServicesVariableGenerate); ok {
			return sOpcuaNodeIdServicesVariableGenerate
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableGenerate) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableGenerate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableGenerateParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableGenerate, error) {
	return OpcuaNodeIdServicesVariableGenerateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableGenerateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableGenerate, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableGenerate", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableGenerate")
	}
	if enum, ok := OpcuaNodeIdServicesVariableGenerateByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableGenerate")
		return OpcuaNodeIdServicesVariableGenerate(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableGenerate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableGenerate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableGenerate", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableGenerate) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableGenerate_GenerateFileForReadMethodType_InputArguments:
		return "GenerateFileForReadMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableGenerate_GenerateFileForReadMethodType_OutputArguments:
		return "GenerateFileForReadMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGenerate_GenerateFileForWriteMethodType_OutputArguments:
		return "GenerateFileForWriteMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableGenerate_GenerateFileForWriteMethodType_InputArguments:
		return "GenerateFileForWriteMethodType_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableGenerate) String() string {
	return e.PLC4XEnumName()
}

