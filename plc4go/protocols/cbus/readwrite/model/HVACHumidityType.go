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

// HVACHumidityType is an enum
type HVACHumidityType uint8

type IHVACHumidityType interface {
	utils.Serializable
}

const(
	HVACHumidityType_NONE HVACHumidityType = 0x00
	HVACHumidityType_EVAPORATOR HVACHumidityType = 0x01
	HVACHumidityType_REFRIGERATIVE HVACHumidityType = 0x02
	HVACHumidityType_EVAPORATOR_REFRIGERATIVE HVACHumidityType = 0x03
)

var HVACHumidityTypeValues []HVACHumidityType

func init() {
	_ = errors.New
	HVACHumidityTypeValues = []HVACHumidityType {
		HVACHumidityType_NONE,
		HVACHumidityType_EVAPORATOR,
		HVACHumidityType_REFRIGERATIVE,
		HVACHumidityType_EVAPORATOR_REFRIGERATIVE,
	}
}

func HVACHumidityTypeByValue(value uint8) (enum HVACHumidityType, ok bool) {
	switch value {
		case 0x00:
			return HVACHumidityType_NONE, true
		case 0x01:
			return HVACHumidityType_EVAPORATOR, true
		case 0x02:
			return HVACHumidityType_REFRIGERATIVE, true
		case 0x03:
			return HVACHumidityType_EVAPORATOR_REFRIGERATIVE, true
	}
	return 0, false
}

func HVACHumidityTypeByName(value string) (enum HVACHumidityType, ok bool) {
	switch value {
	case "NONE":
		return HVACHumidityType_NONE, true
	case "EVAPORATOR":
		return HVACHumidityType_EVAPORATOR, true
	case "REFRIGERATIVE":
		return HVACHumidityType_REFRIGERATIVE, true
	case "EVAPORATOR_REFRIGERATIVE":
		return HVACHumidityType_EVAPORATOR_REFRIGERATIVE, true
	}
	return 0, false
}

func HVACHumidityTypeKnows(value uint8)  bool {
	for _, typeValue := range HVACHumidityTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastHVACHumidityType(structType interface{}) HVACHumidityType {
	castFunc := func(typ interface{}) HVACHumidityType {
		if sHVACHumidityType, ok := typ.(HVACHumidityType); ok {
			return sHVACHumidityType
		}
		return 0
	}
	return castFunc(structType)
}

func (m HVACHumidityType) GetLengthInBits() uint16 {
	return 8
}

func (m HVACHumidityType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func HVACHumidityTypeParse(readBuffer utils.ReadBuffer) (HVACHumidityType, error) {
	val, err := readBuffer.ReadUint8("HVACHumidityType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading HVACHumidityType")
	}
	if enum, ok := HVACHumidityTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return HVACHumidityType(val), nil
	} else {
		return enum, nil
	}
}

func (e HVACHumidityType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e HVACHumidityType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("HVACHumidityType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e HVACHumidityType) PLC4XEnumName() string {
	switch e {
	case HVACHumidityType_NONE:
		return "NONE"
	case HVACHumidityType_EVAPORATOR:
		return "EVAPORATOR"
	case HVACHumidityType_REFRIGERATIVE:
		return "REFRIGERATIVE"
	case HVACHumidityType_EVAPORATOR_REFRIGERATIVE:
		return "EVAPORATOR_REFRIGERATIVE"
	}
	return ""
}

func (e HVACHumidityType) String() string {
	return e.PLC4XEnumName()
}

