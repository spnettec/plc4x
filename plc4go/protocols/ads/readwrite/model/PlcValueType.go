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

// PlcValueType is an enum
type PlcValueType uint8

type IPlcValueType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	PlcValueType_NULL PlcValueType = 0x00
	PlcValueType_BOOL PlcValueType = 0x01
	PlcValueType_BYTE PlcValueType = 0x02
	PlcValueType_WORD PlcValueType = 0x03
	PlcValueType_DWORD PlcValueType = 0x04
	PlcValueType_LWORD PlcValueType = 0x05
	PlcValueType_USINT PlcValueType = 0x11
	PlcValueType_UINT PlcValueType = 0x12
	PlcValueType_UDINT PlcValueType = 0x13
	PlcValueType_ULINT PlcValueType = 0x14
	PlcValueType_SINT PlcValueType = 0x21
	PlcValueType_INT PlcValueType = 0x22
	PlcValueType_DINT PlcValueType = 0x23
	PlcValueType_LINT PlcValueType = 0x24
	PlcValueType_REAL PlcValueType = 0x31
	PlcValueType_LREAL PlcValueType = 0x32
	PlcValueType_CHAR PlcValueType = 0x41
	PlcValueType_WCHAR PlcValueType = 0x42
	PlcValueType_STRING PlcValueType = 0x43
	PlcValueType_WSTRING PlcValueType = 0x44
	PlcValueType_TIME PlcValueType = 0x51
	PlcValueType_LTIME PlcValueType = 0x52
	PlcValueType_DATE PlcValueType = 0x53
	PlcValueType_LDATE PlcValueType = 0x54
	PlcValueType_TIME_OF_DAY PlcValueType = 0x55
	PlcValueType_LTIME_OF_DAY PlcValueType = 0x56
	PlcValueType_DATE_AND_TIME PlcValueType = 0x57
	PlcValueType_LDATE_AND_TIME PlcValueType = 0x58
	PlcValueType_Struct PlcValueType = 0x61
	PlcValueType_List PlcValueType = 0x62
	PlcValueType_RAW_BYTE_ARRAY PlcValueType = 0x71
)

var PlcValueTypeValues []PlcValueType

func init() {
	_ = errors.New
	PlcValueTypeValues = []PlcValueType {
		PlcValueType_NULL,
		PlcValueType_BOOL,
		PlcValueType_BYTE,
		PlcValueType_WORD,
		PlcValueType_DWORD,
		PlcValueType_LWORD,
		PlcValueType_USINT,
		PlcValueType_UINT,
		PlcValueType_UDINT,
		PlcValueType_ULINT,
		PlcValueType_SINT,
		PlcValueType_INT,
		PlcValueType_DINT,
		PlcValueType_LINT,
		PlcValueType_REAL,
		PlcValueType_LREAL,
		PlcValueType_CHAR,
		PlcValueType_WCHAR,
		PlcValueType_STRING,
		PlcValueType_WSTRING,
		PlcValueType_TIME,
		PlcValueType_LTIME,
		PlcValueType_DATE,
		PlcValueType_LDATE,
		PlcValueType_TIME_OF_DAY,
		PlcValueType_LTIME_OF_DAY,
		PlcValueType_DATE_AND_TIME,
		PlcValueType_LDATE_AND_TIME,
		PlcValueType_Struct,
		PlcValueType_List,
		PlcValueType_RAW_BYTE_ARRAY,
	}
}

func PlcValueTypeByValue(value uint8) (enum PlcValueType, ok bool) {
	switch value {
		case 0x00:
			return PlcValueType_NULL, true
		case 0x01:
			return PlcValueType_BOOL, true
		case 0x02:
			return PlcValueType_BYTE, true
		case 0x03:
			return PlcValueType_WORD, true
		case 0x04:
			return PlcValueType_DWORD, true
		case 0x05:
			return PlcValueType_LWORD, true
		case 0x11:
			return PlcValueType_USINT, true
		case 0x12:
			return PlcValueType_UINT, true
		case 0x13:
			return PlcValueType_UDINT, true
		case 0x14:
			return PlcValueType_ULINT, true
		case 0x21:
			return PlcValueType_SINT, true
		case 0x22:
			return PlcValueType_INT, true
		case 0x23:
			return PlcValueType_DINT, true
		case 0x24:
			return PlcValueType_LINT, true
		case 0x31:
			return PlcValueType_REAL, true
		case 0x32:
			return PlcValueType_LREAL, true
		case 0x41:
			return PlcValueType_CHAR, true
		case 0x42:
			return PlcValueType_WCHAR, true
		case 0x43:
			return PlcValueType_STRING, true
		case 0x44:
			return PlcValueType_WSTRING, true
		case 0x51:
			return PlcValueType_TIME, true
		case 0x52:
			return PlcValueType_LTIME, true
		case 0x53:
			return PlcValueType_DATE, true
		case 0x54:
			return PlcValueType_LDATE, true
		case 0x55:
			return PlcValueType_TIME_OF_DAY, true
		case 0x56:
			return PlcValueType_LTIME_OF_DAY, true
		case 0x57:
			return PlcValueType_DATE_AND_TIME, true
		case 0x58:
			return PlcValueType_LDATE_AND_TIME, true
		case 0x61:
			return PlcValueType_Struct, true
		case 0x62:
			return PlcValueType_List, true
		case 0x71:
			return PlcValueType_RAW_BYTE_ARRAY, true
	}
	return 0, false
}

func PlcValueTypeByName(value string) (enum PlcValueType, ok bool) {
	switch value {
	case "NULL":
		return PlcValueType_NULL, true
	case "BOOL":
		return PlcValueType_BOOL, true
	case "BYTE":
		return PlcValueType_BYTE, true
	case "WORD":
		return PlcValueType_WORD, true
	case "DWORD":
		return PlcValueType_DWORD, true
	case "LWORD":
		return PlcValueType_LWORD, true
	case "USINT":
		return PlcValueType_USINT, true
	case "UINT":
		return PlcValueType_UINT, true
	case "UDINT":
		return PlcValueType_UDINT, true
	case "ULINT":
		return PlcValueType_ULINT, true
	case "SINT":
		return PlcValueType_SINT, true
	case "INT":
		return PlcValueType_INT, true
	case "DINT":
		return PlcValueType_DINT, true
	case "LINT":
		return PlcValueType_LINT, true
	case "REAL":
		return PlcValueType_REAL, true
	case "LREAL":
		return PlcValueType_LREAL, true
	case "CHAR":
		return PlcValueType_CHAR, true
	case "WCHAR":
		return PlcValueType_WCHAR, true
	case "STRING":
		return PlcValueType_STRING, true
	case "WSTRING":
		return PlcValueType_WSTRING, true
	case "TIME":
		return PlcValueType_TIME, true
	case "LTIME":
		return PlcValueType_LTIME, true
	case "DATE":
		return PlcValueType_DATE, true
	case "LDATE":
		return PlcValueType_LDATE, true
	case "TIME_OF_DAY":
		return PlcValueType_TIME_OF_DAY, true
	case "LTIME_OF_DAY":
		return PlcValueType_LTIME_OF_DAY, true
	case "DATE_AND_TIME":
		return PlcValueType_DATE_AND_TIME, true
	case "LDATE_AND_TIME":
		return PlcValueType_LDATE_AND_TIME, true
	case "Struct":
		return PlcValueType_Struct, true
	case "List":
		return PlcValueType_List, true
	case "RAW_BYTE_ARRAY":
		return PlcValueType_RAW_BYTE_ARRAY, true
	}
	return 0, false
}

func PlcValueTypeKnows(value uint8)  bool {
	for _, typeValue := range PlcValueTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastPlcValueType(structType any) PlcValueType {
	castFunc := func(typ any) PlcValueType {
		if sPlcValueType, ok := typ.(PlcValueType); ok {
			return sPlcValueType
		}
		return 0
	}
	return castFunc(structType)
}

func (m PlcValueType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m PlcValueType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PlcValueTypeParse(ctx context.Context, theBytes []byte) (PlcValueType, error) {
	return PlcValueTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PlcValueTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PlcValueType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("PlcValueType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading PlcValueType")
	}
	if enum, ok := PlcValueTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for PlcValueType")
		return PlcValueType(val), nil
	} else {
		return enum, nil
	}
}

func (e PlcValueType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e PlcValueType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("PlcValueType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e PlcValueType) PLC4XEnumName() string {
	switch e {
	case PlcValueType_NULL:
		return "NULL"
	case PlcValueType_BOOL:
		return "BOOL"
	case PlcValueType_BYTE:
		return "BYTE"
	case PlcValueType_WORD:
		return "WORD"
	case PlcValueType_DWORD:
		return "DWORD"
	case PlcValueType_LWORD:
		return "LWORD"
	case PlcValueType_USINT:
		return "USINT"
	case PlcValueType_UINT:
		return "UINT"
	case PlcValueType_UDINT:
		return "UDINT"
	case PlcValueType_ULINT:
		return "ULINT"
	case PlcValueType_SINT:
		return "SINT"
	case PlcValueType_INT:
		return "INT"
	case PlcValueType_DINT:
		return "DINT"
	case PlcValueType_LINT:
		return "LINT"
	case PlcValueType_REAL:
		return "REAL"
	case PlcValueType_LREAL:
		return "LREAL"
	case PlcValueType_CHAR:
		return "CHAR"
	case PlcValueType_WCHAR:
		return "WCHAR"
	case PlcValueType_STRING:
		return "STRING"
	case PlcValueType_WSTRING:
		return "WSTRING"
	case PlcValueType_TIME:
		return "TIME"
	case PlcValueType_LTIME:
		return "LTIME"
	case PlcValueType_DATE:
		return "DATE"
	case PlcValueType_LDATE:
		return "LDATE"
	case PlcValueType_TIME_OF_DAY:
		return "TIME_OF_DAY"
	case PlcValueType_LTIME_OF_DAY:
		return "LTIME_OF_DAY"
	case PlcValueType_DATE_AND_TIME:
		return "DATE_AND_TIME"
	case PlcValueType_LDATE_AND_TIME:
		return "LDATE_AND_TIME"
	case PlcValueType_Struct:
		return "Struct"
	case PlcValueType_List:
		return "List"
	case PlcValueType_RAW_BYTE_ARRAY:
		return "RAW_BYTE_ARRAY"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e PlcValueType) String() string {
	return e.PLC4XEnumName()
}

