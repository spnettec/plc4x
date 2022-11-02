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

// DeviceDescriptorMediumType is an enum
type DeviceDescriptorMediumType uint8

type IDeviceDescriptorMediumType interface {
	utils.Serializable
}

const(
	DeviceDescriptorMediumType_TP1 DeviceDescriptorMediumType = 0x0
	DeviceDescriptorMediumType_PL110 DeviceDescriptorMediumType = 0x1
	DeviceDescriptorMediumType_RF DeviceDescriptorMediumType = 0x2
	DeviceDescriptorMediumType_TP0 DeviceDescriptorMediumType = 0x3
	DeviceDescriptorMediumType_PL132 DeviceDescriptorMediumType = 0x4
	DeviceDescriptorMediumType_KNX_IP DeviceDescriptorMediumType = 0x5
)

var DeviceDescriptorMediumTypeValues []DeviceDescriptorMediumType

func init() {
	_ = errors.New
	DeviceDescriptorMediumTypeValues = []DeviceDescriptorMediumType {
		DeviceDescriptorMediumType_TP1,
		DeviceDescriptorMediumType_PL110,
		DeviceDescriptorMediumType_RF,
		DeviceDescriptorMediumType_TP0,
		DeviceDescriptorMediumType_PL132,
		DeviceDescriptorMediumType_KNX_IP,
	}
}

func DeviceDescriptorMediumTypeByValue(value uint8) (enum DeviceDescriptorMediumType, ok bool) {
	switch value {
		case 0x0:
			return DeviceDescriptorMediumType_TP1, true
		case 0x1:
			return DeviceDescriptorMediumType_PL110, true
		case 0x2:
			return DeviceDescriptorMediumType_RF, true
		case 0x3:
			return DeviceDescriptorMediumType_TP0, true
		case 0x4:
			return DeviceDescriptorMediumType_PL132, true
		case 0x5:
			return DeviceDescriptorMediumType_KNX_IP, true
	}
	return 0, false
}

func DeviceDescriptorMediumTypeByName(value string) (enum DeviceDescriptorMediumType, ok bool) {
	switch value {
	case "TP1":
		return DeviceDescriptorMediumType_TP1, true
	case "PL110":
		return DeviceDescriptorMediumType_PL110, true
	case "RF":
		return DeviceDescriptorMediumType_RF, true
	case "TP0":
		return DeviceDescriptorMediumType_TP0, true
	case "PL132":
		return DeviceDescriptorMediumType_PL132, true
	case "KNX_IP":
		return DeviceDescriptorMediumType_KNX_IP, true
	}
	return 0, false
}

func DeviceDescriptorMediumTypeKnows(value uint8)  bool {
	for _, typeValue := range DeviceDescriptorMediumTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastDeviceDescriptorMediumType(structType interface{}) DeviceDescriptorMediumType {
	castFunc := func(typ interface{}) DeviceDescriptorMediumType {
		if sDeviceDescriptorMediumType, ok := typ.(DeviceDescriptorMediumType); ok {
			return sDeviceDescriptorMediumType
		}
		return 0
	}
	return castFunc(structType)
}

func (m DeviceDescriptorMediumType) GetLengthInBits() uint16 {
	return 4
}

func (m DeviceDescriptorMediumType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DeviceDescriptorMediumTypeParse(readBuffer utils.ReadBuffer) (DeviceDescriptorMediumType, error) {
	val, err := readBuffer.ReadUint8("DeviceDescriptorMediumType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DeviceDescriptorMediumType")
	}
	if enum, ok := DeviceDescriptorMediumTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return DeviceDescriptorMediumType(val), nil
	} else {
		return enum, nil
	}
}

func (e DeviceDescriptorMediumType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DeviceDescriptorMediumType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("DeviceDescriptorMediumType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DeviceDescriptorMediumType) PLC4XEnumName() string {
	switch e {
	case DeviceDescriptorMediumType_TP1:
		return "TP1"
	case DeviceDescriptorMediumType_PL110:
		return "PL110"
	case DeviceDescriptorMediumType_RF:
		return "RF"
	case DeviceDescriptorMediumType_TP0:
		return "TP0"
	case DeviceDescriptorMediumType_PL132:
		return "PL132"
	case DeviceDescriptorMediumType_KNX_IP:
		return "KNX_IP"
	}
	return ""
}

func (e DeviceDescriptorMediumType) String() string {
	return e.PLC4XEnumName()
}

