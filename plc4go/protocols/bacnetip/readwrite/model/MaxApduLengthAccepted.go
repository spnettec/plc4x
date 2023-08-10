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

// MaxApduLengthAccepted is an enum
type MaxApduLengthAccepted uint8

type IMaxApduLengthAccepted interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumberOfOctets() uint16
}

const (
	MaxApduLengthAccepted_MINIMUM_MESSAGE_SIZE  MaxApduLengthAccepted = 0x0
	MaxApduLengthAccepted_NUM_OCTETS_128        MaxApduLengthAccepted = 0x1
	MaxApduLengthAccepted_NUM_OCTETS_206        MaxApduLengthAccepted = 0x2
	MaxApduLengthAccepted_NUM_OCTETS_480        MaxApduLengthAccepted = 0x3
	MaxApduLengthAccepted_NUM_OCTETS_1024       MaxApduLengthAccepted = 0x4
	MaxApduLengthAccepted_NUM_OCTETS_1476       MaxApduLengthAccepted = 0x5
	MaxApduLengthAccepted_RESERVED_BY_ASHRAE_01 MaxApduLengthAccepted = 0x6
	MaxApduLengthAccepted_RESERVED_BY_ASHRAE_02 MaxApduLengthAccepted = 0x7
	MaxApduLengthAccepted_RESERVED_BY_ASHRAE_03 MaxApduLengthAccepted = 0x8
	MaxApduLengthAccepted_RESERVED_BY_ASHRAE_04 MaxApduLengthAccepted = 0x9
	MaxApduLengthAccepted_RESERVED_BY_ASHRAE_05 MaxApduLengthAccepted = 0xA
	MaxApduLengthAccepted_RESERVED_BY_ASHRAE_06 MaxApduLengthAccepted = 0xB
	MaxApduLengthAccepted_RESERVED_BY_ASHRAE_07 MaxApduLengthAccepted = 0xC
	MaxApduLengthAccepted_RESERVED_BY_ASHRAE_08 MaxApduLengthAccepted = 0xD
	MaxApduLengthAccepted_RESERVED_BY_ASHRAE_09 MaxApduLengthAccepted = 0xE
	MaxApduLengthAccepted_RESERVED_BY_ASHRAE_10 MaxApduLengthAccepted = 0xF
)

var MaxApduLengthAcceptedValues []MaxApduLengthAccepted

func init() {
	_ = errors.New
	MaxApduLengthAcceptedValues = []MaxApduLengthAccepted{
		MaxApduLengthAccepted_MINIMUM_MESSAGE_SIZE,
		MaxApduLengthAccepted_NUM_OCTETS_128,
		MaxApduLengthAccepted_NUM_OCTETS_206,
		MaxApduLengthAccepted_NUM_OCTETS_480,
		MaxApduLengthAccepted_NUM_OCTETS_1024,
		MaxApduLengthAccepted_NUM_OCTETS_1476,
		MaxApduLengthAccepted_RESERVED_BY_ASHRAE_01,
		MaxApduLengthAccepted_RESERVED_BY_ASHRAE_02,
		MaxApduLengthAccepted_RESERVED_BY_ASHRAE_03,
		MaxApduLengthAccepted_RESERVED_BY_ASHRAE_04,
		MaxApduLengthAccepted_RESERVED_BY_ASHRAE_05,
		MaxApduLengthAccepted_RESERVED_BY_ASHRAE_06,
		MaxApduLengthAccepted_RESERVED_BY_ASHRAE_07,
		MaxApduLengthAccepted_RESERVED_BY_ASHRAE_08,
		MaxApduLengthAccepted_RESERVED_BY_ASHRAE_09,
		MaxApduLengthAccepted_RESERVED_BY_ASHRAE_10,
	}
}

func (e MaxApduLengthAccepted) NumberOfOctets() uint16 {
	switch e {
	case 0x0:
		{ /* '0x0' */
			return 50
		}
	case 0x1:
		{ /* '0x1' */
			return 128
		}
	case 0x2:
		{ /* '0x2' */
			return 206
		}
	case 0x3:
		{ /* '0x3' */
			return 480
		}
	case 0x4:
		{ /* '0x4' */
			return 1024
		}
	case 0x5:
		{ /* '0x5' */
			return 1476
		}
	case 0x6:
		{ /* '0x6' */
			return 0
		}
	case 0x7:
		{ /* '0x7' */
			return 0
		}
	case 0x8:
		{ /* '0x8' */
			return 0
		}
	case 0x9:
		{ /* '0x9' */
			return 0
		}
	case 0xA:
		{ /* '0xA' */
			return 0
		}
	case 0xB:
		{ /* '0xB' */
			return 0
		}
	case 0xC:
		{ /* '0xC' */
			return 0
		}
	case 0xD:
		{ /* '0xD' */
			return 0
		}
	case 0xE:
		{ /* '0xE' */
			return 0
		}
	case 0xF:
		{ /* '0xF' */
			return 0
		}
	default:
		{
			return 0
		}
	}
}

func MaxApduLengthAcceptedFirstEnumForFieldNumberOfOctets(value uint16) (MaxApduLengthAccepted, error) {
	for _, sizeValue := range MaxApduLengthAcceptedValues {
		if sizeValue.NumberOfOctets() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumberOfOctets not found", value)
}
func MaxApduLengthAcceptedByValue(value uint8) (enum MaxApduLengthAccepted, ok bool) {
	switch value {
	case 0x0:
		return MaxApduLengthAccepted_MINIMUM_MESSAGE_SIZE, true
	case 0x1:
		return MaxApduLengthAccepted_NUM_OCTETS_128, true
	case 0x2:
		return MaxApduLengthAccepted_NUM_OCTETS_206, true
	case 0x3:
		return MaxApduLengthAccepted_NUM_OCTETS_480, true
	case 0x4:
		return MaxApduLengthAccepted_NUM_OCTETS_1024, true
	case 0x5:
		return MaxApduLengthAccepted_NUM_OCTETS_1476, true
	case 0x6:
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_01, true
	case 0x7:
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_02, true
	case 0x8:
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_03, true
	case 0x9:
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_04, true
	case 0xA:
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_05, true
	case 0xB:
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_06, true
	case 0xC:
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_07, true
	case 0xD:
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_08, true
	case 0xE:
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_09, true
	case 0xF:
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_10, true
	}
	return 0, false
}

func MaxApduLengthAcceptedByName(value string) (enum MaxApduLengthAccepted, ok bool) {
	switch value {
	case "MINIMUM_MESSAGE_SIZE":
		return MaxApduLengthAccepted_MINIMUM_MESSAGE_SIZE, true
	case "NUM_OCTETS_128":
		return MaxApduLengthAccepted_NUM_OCTETS_128, true
	case "NUM_OCTETS_206":
		return MaxApduLengthAccepted_NUM_OCTETS_206, true
	case "NUM_OCTETS_480":
		return MaxApduLengthAccepted_NUM_OCTETS_480, true
	case "NUM_OCTETS_1024":
		return MaxApduLengthAccepted_NUM_OCTETS_1024, true
	case "NUM_OCTETS_1476":
		return MaxApduLengthAccepted_NUM_OCTETS_1476, true
	case "RESERVED_BY_ASHRAE_01":
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_01, true
	case "RESERVED_BY_ASHRAE_02":
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_02, true
	case "RESERVED_BY_ASHRAE_03":
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_03, true
	case "RESERVED_BY_ASHRAE_04":
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_04, true
	case "RESERVED_BY_ASHRAE_05":
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_05, true
	case "RESERVED_BY_ASHRAE_06":
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_06, true
	case "RESERVED_BY_ASHRAE_07":
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_07, true
	case "RESERVED_BY_ASHRAE_08":
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_08, true
	case "RESERVED_BY_ASHRAE_09":
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_09, true
	case "RESERVED_BY_ASHRAE_10":
		return MaxApduLengthAccepted_RESERVED_BY_ASHRAE_10, true
	}
	return 0, false
}

func MaxApduLengthAcceptedKnows(value uint8) bool {
	for _, typeValue := range MaxApduLengthAcceptedValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMaxApduLengthAccepted(structType any) MaxApduLengthAccepted {
	castFunc := func(typ any) MaxApduLengthAccepted {
		if sMaxApduLengthAccepted, ok := typ.(MaxApduLengthAccepted); ok {
			return sMaxApduLengthAccepted
		}
		return 0
	}
	return castFunc(structType)
}

func (m MaxApduLengthAccepted) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m MaxApduLengthAccepted) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MaxApduLengthAcceptedParse(ctx context.Context, theBytes []byte) (MaxApduLengthAccepted, error) {
	return MaxApduLengthAcceptedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MaxApduLengthAcceptedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MaxApduLengthAccepted, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("MaxApduLengthAccepted", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MaxApduLengthAccepted")
	}
	if enum, ok := MaxApduLengthAcceptedByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for MaxApduLengthAccepted")
		return MaxApduLengthAccepted(val), nil
	} else {
		return enum, nil
	}
}

func (e MaxApduLengthAccepted) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e MaxApduLengthAccepted) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("MaxApduLengthAccepted", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MaxApduLengthAccepted) PLC4XEnumName() string {
	switch e {
	case MaxApduLengthAccepted_MINIMUM_MESSAGE_SIZE:
		return "MINIMUM_MESSAGE_SIZE"
	case MaxApduLengthAccepted_NUM_OCTETS_128:
		return "NUM_OCTETS_128"
	case MaxApduLengthAccepted_NUM_OCTETS_206:
		return "NUM_OCTETS_206"
	case MaxApduLengthAccepted_NUM_OCTETS_480:
		return "NUM_OCTETS_480"
	case MaxApduLengthAccepted_NUM_OCTETS_1024:
		return "NUM_OCTETS_1024"
	case MaxApduLengthAccepted_NUM_OCTETS_1476:
		return "NUM_OCTETS_1476"
	case MaxApduLengthAccepted_RESERVED_BY_ASHRAE_01:
		return "RESERVED_BY_ASHRAE_01"
	case MaxApduLengthAccepted_RESERVED_BY_ASHRAE_02:
		return "RESERVED_BY_ASHRAE_02"
	case MaxApduLengthAccepted_RESERVED_BY_ASHRAE_03:
		return "RESERVED_BY_ASHRAE_03"
	case MaxApduLengthAccepted_RESERVED_BY_ASHRAE_04:
		return "RESERVED_BY_ASHRAE_04"
	case MaxApduLengthAccepted_RESERVED_BY_ASHRAE_05:
		return "RESERVED_BY_ASHRAE_05"
	case MaxApduLengthAccepted_RESERVED_BY_ASHRAE_06:
		return "RESERVED_BY_ASHRAE_06"
	case MaxApduLengthAccepted_RESERVED_BY_ASHRAE_07:
		return "RESERVED_BY_ASHRAE_07"
	case MaxApduLengthAccepted_RESERVED_BY_ASHRAE_08:
		return "RESERVED_BY_ASHRAE_08"
	case MaxApduLengthAccepted_RESERVED_BY_ASHRAE_09:
		return "RESERVED_BY_ASHRAE_09"
	case MaxApduLengthAccepted_RESERVED_BY_ASHRAE_10:
		return "RESERVED_BY_ASHRAE_10"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e MaxApduLengthAccepted) String() string {
	return e.PLC4XEnumName()
}
