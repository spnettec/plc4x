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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetEscalatorMode is an enum
type BACnetEscalatorMode uint16

type IBACnetEscalatorMode interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetEscalatorMode_UNKNOWN                  BACnetEscalatorMode = 0
	BACnetEscalatorMode_STOP                     BACnetEscalatorMode = 1
	BACnetEscalatorMode_UP                       BACnetEscalatorMode = 2
	BACnetEscalatorMode_DOWN                     BACnetEscalatorMode = 3
	BACnetEscalatorMode_INSPECTION               BACnetEscalatorMode = 4
	BACnetEscalatorMode_OUT_OF_SERVICE           BACnetEscalatorMode = 5
	BACnetEscalatorMode_VENDOR_PROPRIETARY_VALUE BACnetEscalatorMode = 0xFFFF
)

var BACnetEscalatorModeValues []BACnetEscalatorMode

func init() {
	_ = errors.New
	BACnetEscalatorModeValues = []BACnetEscalatorMode{
		BACnetEscalatorMode_UNKNOWN,
		BACnetEscalatorMode_STOP,
		BACnetEscalatorMode_UP,
		BACnetEscalatorMode_DOWN,
		BACnetEscalatorMode_INSPECTION,
		BACnetEscalatorMode_OUT_OF_SERVICE,
		BACnetEscalatorMode_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetEscalatorModeByValue(value uint16) BACnetEscalatorMode {
	switch value {
	case 0:
		return BACnetEscalatorMode_UNKNOWN
	case 0xFFFF:
		return BACnetEscalatorMode_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetEscalatorMode_STOP
	case 2:
		return BACnetEscalatorMode_UP
	case 3:
		return BACnetEscalatorMode_DOWN
	case 4:
		return BACnetEscalatorMode_INSPECTION
	case 5:
		return BACnetEscalatorMode_OUT_OF_SERVICE
	}
	return 0
}

func BACnetEscalatorModeByName(value string) (enum BACnetEscalatorMode, ok bool) {
	ok = true
	switch value {
	case "UNKNOWN":
		enum = BACnetEscalatorMode_UNKNOWN
	case "VENDOR_PROPRIETARY_VALUE":
		enum = BACnetEscalatorMode_VENDOR_PROPRIETARY_VALUE
	case "STOP":
		enum = BACnetEscalatorMode_STOP
	case "UP":
		enum = BACnetEscalatorMode_UP
	case "DOWN":
		enum = BACnetEscalatorMode_DOWN
	case "INSPECTION":
		enum = BACnetEscalatorMode_INSPECTION
	case "OUT_OF_SERVICE":
		enum = BACnetEscalatorMode_OUT_OF_SERVICE
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetEscalatorModeKnows(value uint16) bool {
	for _, typeValue := range BACnetEscalatorModeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetEscalatorMode(structType interface{}) BACnetEscalatorMode {
	castFunc := func(typ interface{}) BACnetEscalatorMode {
		if sBACnetEscalatorMode, ok := typ.(BACnetEscalatorMode); ok {
			return sBACnetEscalatorMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetEscalatorMode) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetEscalatorMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEscalatorModeParse(readBuffer utils.ReadBuffer) (BACnetEscalatorMode, error) {
	val, err := readBuffer.ReadUint16("BACnetEscalatorMode", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetEscalatorModeByValue(val), nil
}

func (e BACnetEscalatorMode) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetEscalatorMode", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetEscalatorMode) PLC4XEnumName() string {
	switch e {
	case BACnetEscalatorMode_UNKNOWN:
		return "UNKNOWN"
	case BACnetEscalatorMode_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetEscalatorMode_STOP:
		return "STOP"
	case BACnetEscalatorMode_UP:
		return "UP"
	case BACnetEscalatorMode_DOWN:
		return "DOWN"
	case BACnetEscalatorMode_INSPECTION:
		return "INSPECTION"
	case BACnetEscalatorMode_OUT_OF_SERVICE:
		return "OUT_OF_SERVICE"
	}
	return ""
}

func (e BACnetEscalatorMode) String() string {
	return e.PLC4XEnumName()
}
