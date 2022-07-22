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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// HVACModeAndFlagsMode is an enum
type HVACModeAndFlagsMode uint8

type IHVACModeAndFlagsMode interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	HVACModeAndFlagsMode_OFF           HVACModeAndFlagsMode = 0x0
	HVACModeAndFlagsMode_HEAT_ONLY     HVACModeAndFlagsMode = 0x1
	HVACModeAndFlagsMode_COOL_ONLY     HVACModeAndFlagsMode = 0x2
	HVACModeAndFlagsMode_HEAT_AND_COOL HVACModeAndFlagsMode = 0x3
	HVACModeAndFlagsMode_VENT_FAN_ONLY HVACModeAndFlagsMode = 0x4
)

var HVACModeAndFlagsModeValues []HVACModeAndFlagsMode

func init() {
	_ = errors.New
	HVACModeAndFlagsModeValues = []HVACModeAndFlagsMode{
		HVACModeAndFlagsMode_OFF,
		HVACModeAndFlagsMode_HEAT_ONLY,
		HVACModeAndFlagsMode_COOL_ONLY,
		HVACModeAndFlagsMode_HEAT_AND_COOL,
		HVACModeAndFlagsMode_VENT_FAN_ONLY,
	}
}

func HVACModeAndFlagsModeByValue(value uint8) (enum HVACModeAndFlagsMode, ok bool) {
	switch value {
	case 0x0:
		return HVACModeAndFlagsMode_OFF, true
	case 0x1:
		return HVACModeAndFlagsMode_HEAT_ONLY, true
	case 0x2:
		return HVACModeAndFlagsMode_COOL_ONLY, true
	case 0x3:
		return HVACModeAndFlagsMode_HEAT_AND_COOL, true
	case 0x4:
		return HVACModeAndFlagsMode_VENT_FAN_ONLY, true
	}
	return 0, false
}

func HVACModeAndFlagsModeByName(value string) (enum HVACModeAndFlagsMode, ok bool) {
	switch value {
	case "OFF":
		return HVACModeAndFlagsMode_OFF, true
	case "HEAT_ONLY":
		return HVACModeAndFlagsMode_HEAT_ONLY, true
	case "COOL_ONLY":
		return HVACModeAndFlagsMode_COOL_ONLY, true
	case "HEAT_AND_COOL":
		return HVACModeAndFlagsMode_HEAT_AND_COOL, true
	case "VENT_FAN_ONLY":
		return HVACModeAndFlagsMode_VENT_FAN_ONLY, true
	}
	return 0, false
}

func HVACModeAndFlagsModeKnows(value uint8) bool {
	for _, typeValue := range HVACModeAndFlagsModeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastHVACModeAndFlagsMode(structType interface{}) HVACModeAndFlagsMode {
	castFunc := func(typ interface{}) HVACModeAndFlagsMode {
		if sHVACModeAndFlagsMode, ok := typ.(HVACModeAndFlagsMode); ok {
			return sHVACModeAndFlagsMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m HVACModeAndFlagsMode) GetLengthInBits() uint16 {
	return 3
}

func (m HVACModeAndFlagsMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func HVACModeAndFlagsModeParse(readBuffer utils.ReadBuffer) (HVACModeAndFlagsMode, error) {
	val, err := readBuffer.ReadUint8("HVACModeAndFlagsMode", 3)
	if err != nil {
		return 0, errors.Wrap(err, "error reading HVACModeAndFlagsMode")
	}
	if enum, ok := HVACModeAndFlagsModeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return HVACModeAndFlagsMode(val), nil
	} else {
		return enum, nil
	}
}

func (e HVACModeAndFlagsMode) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("HVACModeAndFlagsMode", 3, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e HVACModeAndFlagsMode) PLC4XEnumName() string {
	switch e {
	case HVACModeAndFlagsMode_OFF:
		return "OFF"
	case HVACModeAndFlagsMode_HEAT_ONLY:
		return "HEAT_ONLY"
	case HVACModeAndFlagsMode_COOL_ONLY:
		return "COOL_ONLY"
	case HVACModeAndFlagsMode_HEAT_AND_COOL:
		return "HEAT_AND_COOL"
	case HVACModeAndFlagsMode_VENT_FAN_ONLY:
		return "VENT_FAN_ONLY"
	}
	return ""
}

func (e HVACModeAndFlagsMode) String() string {
	return e.PLC4XEnumName()
}
