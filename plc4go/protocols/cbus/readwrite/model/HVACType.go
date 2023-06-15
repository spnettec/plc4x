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

// HVACType is an enum
type HVACType uint8

type IHVACType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	HVACType_NONE HVACType = 0x00
	HVACType_FURNACE_GAS_OIL_ELECTRIC HVACType = 0x01
	HVACType_EVAPORATIVE HVACType = 0x02
	HVACType_HEAT_PUMP_REVERSE_CYCLE HVACType = 0x03
	HVACType_HEAT_PUMP_HEATING_ONLY HVACType = 0x04
	HVACType_HEAT_PUMP_COOLING_ONLY HVACType = 0x05
	HVACType_FURNANCE_EVAP_COOLING HVACType = 0x06
	HVACType_FURNANCE_HEAT_PUMP_COOLING_ONLY HVACType = 0x07
	HVACType_HYDRONIC HVACType = 0x08
	HVACType_HYDRONIC_HEAT_PUMP_COOLING_ONLY HVACType = 0x09
	HVACType_HYDRONIC_EVAPORATIVE HVACType = 0x0A
	HVACType_ANY HVACType = 0xFF
)

var HVACTypeValues []HVACType

func init() {
	_ = errors.New
	HVACTypeValues = []HVACType {
		HVACType_NONE,
		HVACType_FURNACE_GAS_OIL_ELECTRIC,
		HVACType_EVAPORATIVE,
		HVACType_HEAT_PUMP_REVERSE_CYCLE,
		HVACType_HEAT_PUMP_HEATING_ONLY,
		HVACType_HEAT_PUMP_COOLING_ONLY,
		HVACType_FURNANCE_EVAP_COOLING,
		HVACType_FURNANCE_HEAT_PUMP_COOLING_ONLY,
		HVACType_HYDRONIC,
		HVACType_HYDRONIC_HEAT_PUMP_COOLING_ONLY,
		HVACType_HYDRONIC_EVAPORATIVE,
		HVACType_ANY,
	}
}

func HVACTypeByValue(value uint8) (enum HVACType, ok bool) {
	switch value {
		case 0x00:
			return HVACType_NONE, true
		case 0x01:
			return HVACType_FURNACE_GAS_OIL_ELECTRIC, true
		case 0x02:
			return HVACType_EVAPORATIVE, true
		case 0x03:
			return HVACType_HEAT_PUMP_REVERSE_CYCLE, true
		case 0x04:
			return HVACType_HEAT_PUMP_HEATING_ONLY, true
		case 0x05:
			return HVACType_HEAT_PUMP_COOLING_ONLY, true
		case 0x06:
			return HVACType_FURNANCE_EVAP_COOLING, true
		case 0x07:
			return HVACType_FURNANCE_HEAT_PUMP_COOLING_ONLY, true
		case 0x08:
			return HVACType_HYDRONIC, true
		case 0x09:
			return HVACType_HYDRONIC_HEAT_PUMP_COOLING_ONLY, true
		case 0x0A:
			return HVACType_HYDRONIC_EVAPORATIVE, true
		case 0xFF:
			return HVACType_ANY, true
	}
	return 0, false
}

func HVACTypeByName(value string) (enum HVACType, ok bool) {
	switch value {
	case "NONE":
		return HVACType_NONE, true
	case "FURNACE_GAS_OIL_ELECTRIC":
		return HVACType_FURNACE_GAS_OIL_ELECTRIC, true
	case "EVAPORATIVE":
		return HVACType_EVAPORATIVE, true
	case "HEAT_PUMP_REVERSE_CYCLE":
		return HVACType_HEAT_PUMP_REVERSE_CYCLE, true
	case "HEAT_PUMP_HEATING_ONLY":
		return HVACType_HEAT_PUMP_HEATING_ONLY, true
	case "HEAT_PUMP_COOLING_ONLY":
		return HVACType_HEAT_PUMP_COOLING_ONLY, true
	case "FURNANCE_EVAP_COOLING":
		return HVACType_FURNANCE_EVAP_COOLING, true
	case "FURNANCE_HEAT_PUMP_COOLING_ONLY":
		return HVACType_FURNANCE_HEAT_PUMP_COOLING_ONLY, true
	case "HYDRONIC":
		return HVACType_HYDRONIC, true
	case "HYDRONIC_HEAT_PUMP_COOLING_ONLY":
		return HVACType_HYDRONIC_HEAT_PUMP_COOLING_ONLY, true
	case "HYDRONIC_EVAPORATIVE":
		return HVACType_HYDRONIC_EVAPORATIVE, true
	case "ANY":
		return HVACType_ANY, true
	}
	return 0, false
}

func HVACTypeKnows(value uint8)  bool {
	for _, typeValue := range HVACTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastHVACType(structType any) HVACType {
	castFunc := func(typ any) HVACType {
		if sHVACType, ok := typ.(HVACType); ok {
			return sHVACType
		}
		return 0
	}
	return castFunc(structType)
}

func (m HVACType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m HVACType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACTypeParse(ctx context.Context, theBytes []byte) (HVACType, error) {
	return HVACTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("HVACType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading HVACType")
	}
	if enum, ok := HVACTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return HVACType(val), nil
	} else {
		return enum, nil
	}
}

func (e HVACType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e HVACType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("HVACType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e HVACType) PLC4XEnumName() string {
	switch e {
	case HVACType_NONE:
		return "NONE"
	case HVACType_FURNACE_GAS_OIL_ELECTRIC:
		return "FURNACE_GAS_OIL_ELECTRIC"
	case HVACType_EVAPORATIVE:
		return "EVAPORATIVE"
	case HVACType_HEAT_PUMP_REVERSE_CYCLE:
		return "HEAT_PUMP_REVERSE_CYCLE"
	case HVACType_HEAT_PUMP_HEATING_ONLY:
		return "HEAT_PUMP_HEATING_ONLY"
	case HVACType_HEAT_PUMP_COOLING_ONLY:
		return "HEAT_PUMP_COOLING_ONLY"
	case HVACType_FURNANCE_EVAP_COOLING:
		return "FURNANCE_EVAP_COOLING"
	case HVACType_FURNANCE_HEAT_PUMP_COOLING_ONLY:
		return "FURNANCE_HEAT_PUMP_COOLING_ONLY"
	case HVACType_HYDRONIC:
		return "HYDRONIC"
	case HVACType_HYDRONIC_HEAT_PUMP_COOLING_ONLY:
		return "HYDRONIC_HEAT_PUMP_COOLING_ONLY"
	case HVACType_HYDRONIC_EVAPORATIVE:
		return "HYDRONIC_EVAPORATIVE"
	case HVACType_ANY:
		return "ANY"
	}
	return ""
}

func (e HVACType) String() string {
	return e.PLC4XEnumName()
}

