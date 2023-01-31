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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetDataType is an enum
type BACnetDataType uint8

type IBACnetDataType interface {
	utils.Serializable
}

const(
	BACnetDataType_NULL BACnetDataType = 0
	BACnetDataType_BOOLEAN BACnetDataType = 1
	BACnetDataType_UNSIGNED_INTEGER BACnetDataType = 2
	BACnetDataType_SIGNED_INTEGER BACnetDataType = 3
	BACnetDataType_REAL BACnetDataType = 4
	BACnetDataType_DOUBLE BACnetDataType = 5
	BACnetDataType_OCTET_STRING BACnetDataType = 6
	BACnetDataType_CHARACTER_STRING BACnetDataType = 7
	BACnetDataType_BIT_STRING BACnetDataType = 8
	BACnetDataType_ENUMERATED BACnetDataType = 9
	BACnetDataType_DATE BACnetDataType = 10
	BACnetDataType_TIME BACnetDataType = 11
	BACnetDataType_BACNET_OBJECT_IDENTIFIER BACnetDataType = 12
	BACnetDataType_UNKNOWN BACnetDataType = 33
)

var BACnetDataTypeValues []BACnetDataType

func init() {
	_ = errors.New
	BACnetDataTypeValues = []BACnetDataType {
		BACnetDataType_NULL,
		BACnetDataType_BOOLEAN,
		BACnetDataType_UNSIGNED_INTEGER,
		BACnetDataType_SIGNED_INTEGER,
		BACnetDataType_REAL,
		BACnetDataType_DOUBLE,
		BACnetDataType_OCTET_STRING,
		BACnetDataType_CHARACTER_STRING,
		BACnetDataType_BIT_STRING,
		BACnetDataType_ENUMERATED,
		BACnetDataType_DATE,
		BACnetDataType_TIME,
		BACnetDataType_BACNET_OBJECT_IDENTIFIER,
		BACnetDataType_UNKNOWN,
	}
}

func BACnetDataTypeByValue(value uint8) (enum BACnetDataType, ok bool) {
	switch value {
		case 0:
			return BACnetDataType_NULL, true
		case 1:
			return BACnetDataType_BOOLEAN, true
		case 10:
			return BACnetDataType_DATE, true
		case 11:
			return BACnetDataType_TIME, true
		case 12:
			return BACnetDataType_BACNET_OBJECT_IDENTIFIER, true
		case 2:
			return BACnetDataType_UNSIGNED_INTEGER, true
		case 3:
			return BACnetDataType_SIGNED_INTEGER, true
		case 33:
			return BACnetDataType_UNKNOWN, true
		case 4:
			return BACnetDataType_REAL, true
		case 5:
			return BACnetDataType_DOUBLE, true
		case 6:
			return BACnetDataType_OCTET_STRING, true
		case 7:
			return BACnetDataType_CHARACTER_STRING, true
		case 8:
			return BACnetDataType_BIT_STRING, true
		case 9:
			return BACnetDataType_ENUMERATED, true
	}
	return 0, false
}

func BACnetDataTypeByName(value string) (enum BACnetDataType, ok bool) {
	switch value {
	case "NULL":
		return BACnetDataType_NULL, true
	case "BOOLEAN":
		return BACnetDataType_BOOLEAN, true
	case "DATE":
		return BACnetDataType_DATE, true
	case "TIME":
		return BACnetDataType_TIME, true
	case "BACNET_OBJECT_IDENTIFIER":
		return BACnetDataType_BACNET_OBJECT_IDENTIFIER, true
	case "UNSIGNED_INTEGER":
		return BACnetDataType_UNSIGNED_INTEGER, true
	case "SIGNED_INTEGER":
		return BACnetDataType_SIGNED_INTEGER, true
	case "UNKNOWN":
		return BACnetDataType_UNKNOWN, true
	case "REAL":
		return BACnetDataType_REAL, true
	case "DOUBLE":
		return BACnetDataType_DOUBLE, true
	case "OCTET_STRING":
		return BACnetDataType_OCTET_STRING, true
	case "CHARACTER_STRING":
		return BACnetDataType_CHARACTER_STRING, true
	case "BIT_STRING":
		return BACnetDataType_BIT_STRING, true
	case "ENUMERATED":
		return BACnetDataType_ENUMERATED, true
	}
	return 0, false
}

func BACnetDataTypeKnows(value uint8)  bool {
	for _, typeValue := range BACnetDataTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetDataType(structType interface{}) BACnetDataType {
	castFunc := func(typ interface{}) BACnetDataType {
		if sBACnetDataType, ok := typ.(BACnetDataType); ok {
			return sBACnetDataType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetDataType) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetDataType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDataTypeParse(theBytes []byte) (BACnetDataType, error) {
	return BACnetDataTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetDataTypeParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetDataType, error) {
	val, err := readBuffer.ReadUint8("BACnetDataType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetDataType")
	}
	if enum, ok := BACnetDataTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetDataType(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetDataType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetDataType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetDataType) PLC4XEnumName() string {
	switch e {
	case BACnetDataType_NULL:
		return "NULL"
	case BACnetDataType_BOOLEAN:
		return "BOOLEAN"
	case BACnetDataType_DATE:
		return "DATE"
	case BACnetDataType_TIME:
		return "TIME"
	case BACnetDataType_BACNET_OBJECT_IDENTIFIER:
		return "BACNET_OBJECT_IDENTIFIER"
	case BACnetDataType_UNSIGNED_INTEGER:
		return "UNSIGNED_INTEGER"
	case BACnetDataType_SIGNED_INTEGER:
		return "SIGNED_INTEGER"
	case BACnetDataType_UNKNOWN:
		return "UNKNOWN"
	case BACnetDataType_REAL:
		return "REAL"
	case BACnetDataType_DOUBLE:
		return "DOUBLE"
	case BACnetDataType_OCTET_STRING:
		return "OCTET_STRING"
	case BACnetDataType_CHARACTER_STRING:
		return "CHARACTER_STRING"
	case BACnetDataType_BIT_STRING:
		return "BIT_STRING"
	case BACnetDataType_ENUMERATED:
		return "ENUMERATED"
	}
	return ""
}

func (e BACnetDataType) String() string {
	return e.PLC4XEnumName()
}

