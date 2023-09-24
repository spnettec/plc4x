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

// ModbusDataType is an enum
type ModbusDataType uint8

type IModbusDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	DataTypeSize() uint8
}

const(
	ModbusDataType_BOOL ModbusDataType = 1
	ModbusDataType_BYTE ModbusDataType = 2
	ModbusDataType_WORD ModbusDataType = 3
	ModbusDataType_DWORD ModbusDataType = 4
	ModbusDataType_LWORD ModbusDataType = 5
	ModbusDataType_SINT ModbusDataType = 6
	ModbusDataType_INT ModbusDataType = 7
	ModbusDataType_DINT ModbusDataType = 8
	ModbusDataType_LINT ModbusDataType = 9
	ModbusDataType_USINT ModbusDataType = 10
	ModbusDataType_UINT ModbusDataType = 11
	ModbusDataType_UDINT ModbusDataType = 12
	ModbusDataType_ULINT ModbusDataType = 13
	ModbusDataType_REAL ModbusDataType = 14
	ModbusDataType_LREAL ModbusDataType = 15
	ModbusDataType_TIME ModbusDataType = 16
	ModbusDataType_LTIME ModbusDataType = 17
	ModbusDataType_DATE ModbusDataType = 18
	ModbusDataType_LDATE ModbusDataType = 19
	ModbusDataType_TIME_OF_DAY ModbusDataType = 20
	ModbusDataType_LTIME_OF_DAY ModbusDataType = 21
	ModbusDataType_DATE_AND_TIME ModbusDataType = 22
	ModbusDataType_LDATE_AND_TIME ModbusDataType = 23
	ModbusDataType_CHAR ModbusDataType = 24
	ModbusDataType_WCHAR ModbusDataType = 25
	ModbusDataType_STRING ModbusDataType = 26
	ModbusDataType_WSTRING ModbusDataType = 27
)

var ModbusDataTypeValues []ModbusDataType

func init() {
	_ = errors.New
	ModbusDataTypeValues = []ModbusDataType {
		ModbusDataType_BOOL,
		ModbusDataType_BYTE,
		ModbusDataType_WORD,
		ModbusDataType_DWORD,
		ModbusDataType_LWORD,
		ModbusDataType_SINT,
		ModbusDataType_INT,
		ModbusDataType_DINT,
		ModbusDataType_LINT,
		ModbusDataType_USINT,
		ModbusDataType_UINT,
		ModbusDataType_UDINT,
		ModbusDataType_ULINT,
		ModbusDataType_REAL,
		ModbusDataType_LREAL,
		ModbusDataType_TIME,
		ModbusDataType_LTIME,
		ModbusDataType_DATE,
		ModbusDataType_LDATE,
		ModbusDataType_TIME_OF_DAY,
		ModbusDataType_LTIME_OF_DAY,
		ModbusDataType_DATE_AND_TIME,
		ModbusDataType_LDATE_AND_TIME,
		ModbusDataType_CHAR,
		ModbusDataType_WCHAR,
		ModbusDataType_STRING,
		ModbusDataType_WSTRING,
	}
}


func (e ModbusDataType) DataTypeSize() uint8 {
	switch e  {
		case 1: { /* '1' */
            return 2
		}
		case 10: { /* '10' */
            return 2
		}
		case 11: { /* '11' */
            return 2
		}
		case 12: { /* '12' */
            return 4
		}
		case 13: { /* '13' */
            return 8
		}
		case 14: { /* '14' */
            return 4
		}
		case 15: { /* '15' */
            return 8
		}
		case 16: { /* '16' */
            return 8
		}
		case 17: { /* '17' */
            return 8
		}
		case 18: { /* '18' */
            return 8
		}
		case 19: { /* '19' */
            return 8
		}
		case 2: { /* '2' */
            return 2
		}
		case 20: { /* '20' */
            return 8
		}
		case 21: { /* '21' */
            return 8
		}
		case 22: { /* '22' */
            return 8
		}
		case 23: { /* '23' */
            return 8
		}
		case 24: { /* '24' */
            return 1
		}
		case 25: { /* '25' */
            return 2
		}
		case 26: { /* '26' */
            return 1
		}
		case 27: { /* '27' */
            return 2
		}
		case 3: { /* '3' */
            return 2
		}
		case 4: { /* '4' */
            return 4
		}
		case 5: { /* '5' */
            return 8
		}
		case 6: { /* '6' */
            return 2
		}
		case 7: { /* '7' */
            return 2
		}
		case 8: { /* '8' */
            return 4
		}
		case 9: { /* '9' */
            return 8
		}
		default: {
			return 0
		}
	}
}

func ModbusDataTypeFirstEnumForFieldDataTypeSize(value uint8) (ModbusDataType, error) {
	for _, sizeValue := range ModbusDataTypeValues {
		if sizeValue.DataTypeSize() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing DataTypeSize not found", value)
}
func ModbusDataTypeByValue(value uint8) (enum ModbusDataType, ok bool) {
	switch value {
		case 1:
			return ModbusDataType_BOOL, true
		case 10:
			return ModbusDataType_USINT, true
		case 11:
			return ModbusDataType_UINT, true
		case 12:
			return ModbusDataType_UDINT, true
		case 13:
			return ModbusDataType_ULINT, true
		case 14:
			return ModbusDataType_REAL, true
		case 15:
			return ModbusDataType_LREAL, true
		case 16:
			return ModbusDataType_TIME, true
		case 17:
			return ModbusDataType_LTIME, true
		case 18:
			return ModbusDataType_DATE, true
		case 19:
			return ModbusDataType_LDATE, true
		case 2:
			return ModbusDataType_BYTE, true
		case 20:
			return ModbusDataType_TIME_OF_DAY, true
		case 21:
			return ModbusDataType_LTIME_OF_DAY, true
		case 22:
			return ModbusDataType_DATE_AND_TIME, true
		case 23:
			return ModbusDataType_LDATE_AND_TIME, true
		case 24:
			return ModbusDataType_CHAR, true
		case 25:
			return ModbusDataType_WCHAR, true
		case 26:
			return ModbusDataType_STRING, true
		case 27:
			return ModbusDataType_WSTRING, true
		case 3:
			return ModbusDataType_WORD, true
		case 4:
			return ModbusDataType_DWORD, true
		case 5:
			return ModbusDataType_LWORD, true
		case 6:
			return ModbusDataType_SINT, true
		case 7:
			return ModbusDataType_INT, true
		case 8:
			return ModbusDataType_DINT, true
		case 9:
			return ModbusDataType_LINT, true
	}
	return 0, false
}

func ModbusDataTypeByName(value string) (enum ModbusDataType, ok bool) {
	switch value {
	case "BOOL":
		return ModbusDataType_BOOL, true
	case "USINT":
		return ModbusDataType_USINT, true
	case "UINT":
		return ModbusDataType_UINT, true
	case "UDINT":
		return ModbusDataType_UDINT, true
	case "ULINT":
		return ModbusDataType_ULINT, true
	case "REAL":
		return ModbusDataType_REAL, true
	case "LREAL":
		return ModbusDataType_LREAL, true
	case "TIME":
		return ModbusDataType_TIME, true
	case "LTIME":
		return ModbusDataType_LTIME, true
	case "DATE":
		return ModbusDataType_DATE, true
	case "LDATE":
		return ModbusDataType_LDATE, true
	case "BYTE":
		return ModbusDataType_BYTE, true
	case "TIME_OF_DAY":
		return ModbusDataType_TIME_OF_DAY, true
	case "LTIME_OF_DAY":
		return ModbusDataType_LTIME_OF_DAY, true
	case "DATE_AND_TIME":
		return ModbusDataType_DATE_AND_TIME, true
	case "LDATE_AND_TIME":
		return ModbusDataType_LDATE_AND_TIME, true
	case "CHAR":
		return ModbusDataType_CHAR, true
	case "WCHAR":
		return ModbusDataType_WCHAR, true
	case "STRING":
		return ModbusDataType_STRING, true
	case "WSTRING":
		return ModbusDataType_WSTRING, true
	case "WORD":
		return ModbusDataType_WORD, true
	case "DWORD":
		return ModbusDataType_DWORD, true
	case "LWORD":
		return ModbusDataType_LWORD, true
	case "SINT":
		return ModbusDataType_SINT, true
	case "INT":
		return ModbusDataType_INT, true
	case "DINT":
		return ModbusDataType_DINT, true
	case "LINT":
		return ModbusDataType_LINT, true
	}
	return 0, false
}

func ModbusDataTypeKnows(value uint8)  bool {
	for _, typeValue := range ModbusDataTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastModbusDataType(structType any) ModbusDataType {
	castFunc := func(typ any) ModbusDataType {
		if sModbusDataType, ok := typ.(ModbusDataType); ok {
			return sModbusDataType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ModbusDataType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m ModbusDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusDataTypeParse(ctx context.Context, theBytes []byte) (ModbusDataType, error) {
	return ModbusDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModbusDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusDataType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("ModbusDataType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ModbusDataType")
	}
	if enum, ok := ModbusDataTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ModbusDataType")
		return ModbusDataType(val), nil
	} else {
		return enum, nil
	}
}

func (e ModbusDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ModbusDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("ModbusDataType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ModbusDataType) PLC4XEnumName() string {
	switch e {
	case ModbusDataType_BOOL:
		return "BOOL"
	case ModbusDataType_USINT:
		return "USINT"
	case ModbusDataType_UINT:
		return "UINT"
	case ModbusDataType_UDINT:
		return "UDINT"
	case ModbusDataType_ULINT:
		return "ULINT"
	case ModbusDataType_REAL:
		return "REAL"
	case ModbusDataType_LREAL:
		return "LREAL"
	case ModbusDataType_TIME:
		return "TIME"
	case ModbusDataType_LTIME:
		return "LTIME"
	case ModbusDataType_DATE:
		return "DATE"
	case ModbusDataType_LDATE:
		return "LDATE"
	case ModbusDataType_BYTE:
		return "BYTE"
	case ModbusDataType_TIME_OF_DAY:
		return "TIME_OF_DAY"
	case ModbusDataType_LTIME_OF_DAY:
		return "LTIME_OF_DAY"
	case ModbusDataType_DATE_AND_TIME:
		return "DATE_AND_TIME"
	case ModbusDataType_LDATE_AND_TIME:
		return "LDATE_AND_TIME"
	case ModbusDataType_CHAR:
		return "CHAR"
	case ModbusDataType_WCHAR:
		return "WCHAR"
	case ModbusDataType_STRING:
		return "STRING"
	case ModbusDataType_WSTRING:
		return "WSTRING"
	case ModbusDataType_WORD:
		return "WORD"
	case ModbusDataType_DWORD:
		return "DWORD"
	case ModbusDataType_LWORD:
		return "LWORD"
	case ModbusDataType_SINT:
		return "SINT"
	case ModbusDataType_INT:
		return "INT"
	case ModbusDataType_DINT:
		return "DINT"
	case ModbusDataType_LINT:
		return "LINT"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e ModbusDataType) String() string {
	return e.PLC4XEnumName()
}

