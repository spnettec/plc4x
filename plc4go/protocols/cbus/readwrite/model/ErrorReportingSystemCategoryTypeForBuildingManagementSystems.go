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

// ErrorReportingSystemCategoryTypeForBuildingManagementSystems is an enum
type ErrorReportingSystemCategoryTypeForBuildingManagementSystems uint8

type IErrorReportingSystemCategoryTypeForBuildingManagementSystems interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_BMS_DIAGNOSTIC_REPORTING ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0x0
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_1 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0x1
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_2 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0x2
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_3 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0x3
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_4 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0x4
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_5 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0x5
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_6 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0x6
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_7 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0x7
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_8 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0x8
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_9 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0x9
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_10 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0xA
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_11 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0xB
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_12 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0xC
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_13 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0xD
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_14 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0xE
	ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_15 ErrorReportingSystemCategoryTypeForBuildingManagementSystems = 0xF
)

var ErrorReportingSystemCategoryTypeForBuildingManagementSystemsValues []ErrorReportingSystemCategoryTypeForBuildingManagementSystems

func init() {
	_ = errors.New
	ErrorReportingSystemCategoryTypeForBuildingManagementSystemsValues = []ErrorReportingSystemCategoryTypeForBuildingManagementSystems {
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_BMS_DIAGNOSTIC_REPORTING,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_1,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_2,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_3,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_4,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_5,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_6,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_7,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_8,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_9,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_10,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_11,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_12,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_13,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_14,
		ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_15,
	}
}

func ErrorReportingSystemCategoryTypeForBuildingManagementSystemsByValue(value uint8) (enum ErrorReportingSystemCategoryTypeForBuildingManagementSystems, ok bool) {
	switch value {
		case 0x0:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_BMS_DIAGNOSTIC_REPORTING, true
		case 0x1:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_1, true
		case 0x2:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_2, true
		case 0x3:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_3, true
		case 0x4:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_4, true
		case 0x5:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_5, true
		case 0x6:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_6, true
		case 0x7:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_7, true
		case 0x8:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_8, true
		case 0x9:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_9, true
		case 0xA:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_10, true
		case 0xB:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_11, true
		case 0xC:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_12, true
		case 0xD:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_13, true
		case 0xE:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_14, true
		case 0xF:
			return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_15, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryTypeForBuildingManagementSystemsByName(value string) (enum ErrorReportingSystemCategoryTypeForBuildingManagementSystems, ok bool) {
	switch value {
	case "BMS_DIAGNOSTIC_REPORTING":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_BMS_DIAGNOSTIC_REPORTING, true
	case "RESERVED_1":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_1, true
	case "RESERVED_2":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_2, true
	case "RESERVED_3":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_3, true
	case "RESERVED_4":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_4, true
	case "RESERVED_5":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_5, true
	case "RESERVED_6":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_6, true
	case "RESERVED_7":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_7, true
	case "RESERVED_8":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_8, true
	case "RESERVED_9":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_9, true
	case "RESERVED_10":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_10, true
	case "RESERVED_11":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_11, true
	case "RESERVED_12":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_12, true
	case "RESERVED_13":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_13, true
	case "RESERVED_14":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_14, true
	case "RESERVED_15":
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_15, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryTypeForBuildingManagementSystemsKnows(value uint8)  bool {
	for _, typeValue := range ErrorReportingSystemCategoryTypeForBuildingManagementSystemsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastErrorReportingSystemCategoryTypeForBuildingManagementSystems(structType any) ErrorReportingSystemCategoryTypeForBuildingManagementSystems {
	castFunc := func(typ any) ErrorReportingSystemCategoryTypeForBuildingManagementSystems {
		if sErrorReportingSystemCategoryTypeForBuildingManagementSystems, ok := typ.(ErrorReportingSystemCategoryTypeForBuildingManagementSystems); ok {
			return sErrorReportingSystemCategoryTypeForBuildingManagementSystems
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorReportingSystemCategoryTypeForBuildingManagementSystems) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m ErrorReportingSystemCategoryTypeForBuildingManagementSystems) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryTypeForBuildingManagementSystemsParse(ctx context.Context, theBytes []byte) (ErrorReportingSystemCategoryTypeForBuildingManagementSystems, error) {
	return ErrorReportingSystemCategoryTypeForBuildingManagementSystemsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingSystemCategoryTypeForBuildingManagementSystemsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSystemCategoryTypeForBuildingManagementSystems, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("ErrorReportingSystemCategoryTypeForBuildingManagementSystems", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorReportingSystemCategoryTypeForBuildingManagementSystems")
	}
	if enum, ok := ErrorReportingSystemCategoryTypeForBuildingManagementSystemsByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return ErrorReportingSystemCategoryTypeForBuildingManagementSystems(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorReportingSystemCategoryTypeForBuildingManagementSystems) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ErrorReportingSystemCategoryTypeForBuildingManagementSystems) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("ErrorReportingSystemCategoryTypeForBuildingManagementSystems", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorReportingSystemCategoryTypeForBuildingManagementSystems) PLC4XEnumName() string {
	switch e {
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_BMS_DIAGNOSTIC_REPORTING:
		return "BMS_DIAGNOSTIC_REPORTING"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_1:
		return "RESERVED_1"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_2:
		return "RESERVED_2"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_3:
		return "RESERVED_3"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_4:
		return "RESERVED_4"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_5:
		return "RESERVED_5"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_6:
		return "RESERVED_6"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_7:
		return "RESERVED_7"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_8:
		return "RESERVED_8"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_9:
		return "RESERVED_9"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_10:
		return "RESERVED_10"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_11:
		return "RESERVED_11"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_12:
		return "RESERVED_12"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_13:
		return "RESERVED_13"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_14:
		return "RESERVED_14"
	case ErrorReportingSystemCategoryTypeForBuildingManagementSystems_RESERVED_15:
		return "RESERVED_15"
	}
	return ""
}

func (e ErrorReportingSystemCategoryTypeForBuildingManagementSystems) String() string {
	return e.PLC4XEnumName()
}

