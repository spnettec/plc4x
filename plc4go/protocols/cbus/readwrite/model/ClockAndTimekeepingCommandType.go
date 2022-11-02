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
	"encoding/binary"

	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ClockAndTimekeepingCommandType is an enum
type ClockAndTimekeepingCommandType uint8

type IClockAndTimekeepingCommandType interface {
	utils.Serializable
	NumberOfArguments() uint8
}

const(
	ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE ClockAndTimekeepingCommandType = 0x00
	ClockAndTimekeepingCommandType_REQUEST_REFRESH ClockAndTimekeepingCommandType = 0x01
)

var ClockAndTimekeepingCommandTypeValues []ClockAndTimekeepingCommandType

func init() {
	_ = errors.New
	ClockAndTimekeepingCommandTypeValues = []ClockAndTimekeepingCommandType {
		ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE,
		ClockAndTimekeepingCommandType_REQUEST_REFRESH,
	}
}


func (e ClockAndTimekeepingCommandType) NumberOfArguments() uint8 {
	switch e  {
		case 0x00: { /* '0x00' */
            return 0xFF
		}
		case 0x01: { /* '0x01' */
            return 0
		}
		default: {
			return 0
		}
	}
}

func ClockAndTimekeepingCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (ClockAndTimekeepingCommandType, error) {
	for _, sizeValue := range ClockAndTimekeepingCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumberOfArguments not found", value)
}
func ClockAndTimekeepingCommandTypeByValue(value uint8) (enum ClockAndTimekeepingCommandType, ok bool) {
	switch value {
		case 0x00:
			return ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE, true
		case 0x01:
			return ClockAndTimekeepingCommandType_REQUEST_REFRESH, true
	}
	return 0, false
}

func ClockAndTimekeepingCommandTypeByName(value string) (enum ClockAndTimekeepingCommandType, ok bool) {
	switch value {
	case "UPDATE_NETWORK_VARIABLE":
		return ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE, true
	case "REQUEST_REFRESH":
		return ClockAndTimekeepingCommandType_REQUEST_REFRESH, true
	}
	return 0, false
}

func ClockAndTimekeepingCommandTypeKnows(value uint8)  bool {
	for _, typeValue := range ClockAndTimekeepingCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastClockAndTimekeepingCommandType(structType interface{}) ClockAndTimekeepingCommandType {
	castFunc := func(typ interface{}) ClockAndTimekeepingCommandType {
		if sClockAndTimekeepingCommandType, ok := typ.(ClockAndTimekeepingCommandType); ok {
			return sClockAndTimekeepingCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ClockAndTimekeepingCommandType) GetLengthInBits() uint16 {
	return 4
}

func (m ClockAndTimekeepingCommandType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ClockAndTimekeepingCommandTypeParse(readBuffer utils.ReadBuffer) (ClockAndTimekeepingCommandType, error) {
	val, err := readBuffer.ReadUint8("ClockAndTimekeepingCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ClockAndTimekeepingCommandType")
	}
	if enum, ok := ClockAndTimekeepingCommandTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ClockAndTimekeepingCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e ClockAndTimekeepingCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ClockAndTimekeepingCommandType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ClockAndTimekeepingCommandType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ClockAndTimekeepingCommandType) PLC4XEnumName() string {
	switch e {
	case ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE:
		return "UPDATE_NETWORK_VARIABLE"
	case ClockAndTimekeepingCommandType_REQUEST_REFRESH:
		return "REQUEST_REFRESH"
	}
	return ""
}

func (e ClockAndTimekeepingCommandType) String() string {
	return e.PLC4XEnumName()
}

