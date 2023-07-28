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

// TelephonyCommandTypeContainer is an enum
type TelephonyCommandTypeContainer uint8

type ITelephonyCommandTypeContainer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumBytes() uint8
	CommandType() TelephonyCommandType
}

const(
	TelephonyCommandTypeContainer_TelephonyCommandLineOnHook TelephonyCommandTypeContainer = 0x09
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_0Bytes TelephonyCommandTypeContainer = 0xA0
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_1Bytes TelephonyCommandTypeContainer = 0xA1
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_2Bytes TelephonyCommandTypeContainer = 0xA2
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_3Bytes TelephonyCommandTypeContainer = 0xA3
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_4Bytes TelephonyCommandTypeContainer = 0xA4
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_5Bytes TelephonyCommandTypeContainer = 0xA5
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_6Bytes TelephonyCommandTypeContainer = 0xA6
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_7Bytes TelephonyCommandTypeContainer = 0xA7
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_8Bytes TelephonyCommandTypeContainer = 0xA8
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_9Bytes TelephonyCommandTypeContainer = 0xA9
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_10Bytes TelephonyCommandTypeContainer = 0xAA
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_11Bytes TelephonyCommandTypeContainer = 0xAB
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_12Bytes TelephonyCommandTypeContainer = 0xAC
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_13Bytes TelephonyCommandTypeContainer = 0xAD
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_14Bytes TelephonyCommandTypeContainer = 0xAE
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_15Bytes TelephonyCommandTypeContainer = 0xAF
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_16Bytes TelephonyCommandTypeContainer = 0xB0
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_17Bytes TelephonyCommandTypeContainer = 0xB1
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_18Bytes TelephonyCommandTypeContainer = 0xB2
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_19Bytes TelephonyCommandTypeContainer = 0xB3
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_20Bytes TelephonyCommandTypeContainer = 0xB4
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_21Bytes TelephonyCommandTypeContainer = 0xB5
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_22Bytes TelephonyCommandTypeContainer = 0xB6
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_23Bytes TelephonyCommandTypeContainer = 0xB7
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_24Bytes TelephonyCommandTypeContainer = 0xB8
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_25Bytes TelephonyCommandTypeContainer = 0xB9
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_26Bytes TelephonyCommandTypeContainer = 0xBA
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_27Bytes TelephonyCommandTypeContainer = 0xBB
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_28Bytes TelephonyCommandTypeContainer = 0xBC
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_29Bytes TelephonyCommandTypeContainer = 0xBD
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_30Bytes TelephonyCommandTypeContainer = 0xBE
	TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_31Bytes TelephonyCommandTypeContainer = 0xBF
)

var TelephonyCommandTypeContainerValues []TelephonyCommandTypeContainer

func init() {
	_ = errors.New
	TelephonyCommandTypeContainerValues = []TelephonyCommandTypeContainer {
		TelephonyCommandTypeContainer_TelephonyCommandLineOnHook,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_0Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_1Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_2Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_3Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_4Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_5Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_6Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_7Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_8Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_9Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_10Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_11Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_12Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_13Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_14Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_15Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_16Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_17Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_18Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_19Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_20Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_21Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_22Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_23Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_24Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_25Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_26Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_27Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_28Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_29Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_30Bytes,
		TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_31Bytes,
	}
}


func (e TelephonyCommandTypeContainer) NumBytes() uint8 {
	switch e  {
		case 0x09: { /* '0x09' */
            return 1
		}
		case 0xA0: { /* '0xA0' */
            return 0
		}
		case 0xA1: { /* '0xA1' */
            return 1
		}
		case 0xA2: { /* '0xA2' */
            return 2
		}
		case 0xA3: { /* '0xA3' */
            return 3
		}
		case 0xA4: { /* '0xA4' */
            return 4
		}
		case 0xA5: { /* '0xA5' */
            return 5
		}
		case 0xA6: { /* '0xA6' */
            return 6
		}
		case 0xA7: { /* '0xA7' */
            return 7
		}
		case 0xA8: { /* '0xA8' */
            return 8
		}
		case 0xA9: { /* '0xA9' */
            return 9
		}
		case 0xAA: { /* '0xAA' */
            return 10
		}
		case 0xAB: { /* '0xAB' */
            return 11
		}
		case 0xAC: { /* '0xAC' */
            return 12
		}
		case 0xAD: { /* '0xAD' */
            return 13
		}
		case 0xAE: { /* '0xAE' */
            return 14
		}
		case 0xAF: { /* '0xAF' */
            return 15
		}
		case 0xB0: { /* '0xB0' */
            return 16
		}
		case 0xB1: { /* '0xB1' */
            return 17
		}
		case 0xB2: { /* '0xB2' */
            return 18
		}
		case 0xB3: { /* '0xB3' */
            return 19
		}
		case 0xB4: { /* '0xB4' */
            return 20
		}
		case 0xB5: { /* '0xB5' */
            return 21
		}
		case 0xB6: { /* '0xB6' */
            return 22
		}
		case 0xB7: { /* '0xB7' */
            return 23
		}
		case 0xB8: { /* '0xB8' */
            return 24
		}
		case 0xB9: { /* '0xB9' */
            return 25
		}
		case 0xBA: { /* '0xBA' */
            return 26
		}
		case 0xBB: { /* '0xBB' */
            return 27
		}
		case 0xBC: { /* '0xBC' */
            return 28
		}
		case 0xBD: { /* '0xBD' */
            return 29
		}
		case 0xBE: { /* '0xBE' */
            return 30
		}
		case 0xBF: { /* '0xBF' */
            return 31
		}
		default: {
			return 0
		}
	}
}

func TelephonyCommandTypeContainerFirstEnumForFieldNumBytes(value uint8) (TelephonyCommandTypeContainer, error) {
	for _, sizeValue := range TelephonyCommandTypeContainerValues {
		if sizeValue.NumBytes() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumBytes not found", value)
}

func (e TelephonyCommandTypeContainer) CommandType() TelephonyCommandType {
	switch e  {
		case 0x09: { /* '0x09' */
			return TelephonyCommandType_EVENT
		}
		case 0xA0: { /* '0xA0' */
			return TelephonyCommandType_EVENT
		}
		case 0xA1: { /* '0xA1' */
			return TelephonyCommandType_EVENT
		}
		case 0xA2: { /* '0xA2' */
			return TelephonyCommandType_EVENT
		}
		case 0xA3: { /* '0xA3' */
			return TelephonyCommandType_EVENT
		}
		case 0xA4: { /* '0xA4' */
			return TelephonyCommandType_EVENT
		}
		case 0xA5: { /* '0xA5' */
			return TelephonyCommandType_EVENT
		}
		case 0xA6: { /* '0xA6' */
			return TelephonyCommandType_EVENT
		}
		case 0xA7: { /* '0xA7' */
			return TelephonyCommandType_EVENT
		}
		case 0xA8: { /* '0xA8' */
			return TelephonyCommandType_EVENT
		}
		case 0xA9: { /* '0xA9' */
			return TelephonyCommandType_EVENT
		}
		case 0xAA: { /* '0xAA' */
			return TelephonyCommandType_EVENT
		}
		case 0xAB: { /* '0xAB' */
			return TelephonyCommandType_EVENT
		}
		case 0xAC: { /* '0xAC' */
			return TelephonyCommandType_EVENT
		}
		case 0xAD: { /* '0xAD' */
			return TelephonyCommandType_EVENT
		}
		case 0xAE: { /* '0xAE' */
			return TelephonyCommandType_EVENT
		}
		case 0xAF: { /* '0xAF' */
			return TelephonyCommandType_EVENT
		}
		case 0xB0: { /* '0xB0' */
			return TelephonyCommandType_EVENT
		}
		case 0xB1: { /* '0xB1' */
			return TelephonyCommandType_EVENT
		}
		case 0xB2: { /* '0xB2' */
			return TelephonyCommandType_EVENT
		}
		case 0xB3: { /* '0xB3' */
			return TelephonyCommandType_EVENT
		}
		case 0xB4: { /* '0xB4' */
			return TelephonyCommandType_EVENT
		}
		case 0xB5: { /* '0xB5' */
			return TelephonyCommandType_EVENT
		}
		case 0xB6: { /* '0xB6' */
			return TelephonyCommandType_EVENT
		}
		case 0xB7: { /* '0xB7' */
			return TelephonyCommandType_EVENT
		}
		case 0xB8: { /* '0xB8' */
			return TelephonyCommandType_EVENT
		}
		case 0xB9: { /* '0xB9' */
			return TelephonyCommandType_EVENT
		}
		case 0xBA: { /* '0xBA' */
			return TelephonyCommandType_EVENT
		}
		case 0xBB: { /* '0xBB' */
			return TelephonyCommandType_EVENT
		}
		case 0xBC: { /* '0xBC' */
			return TelephonyCommandType_EVENT
		}
		case 0xBD: { /* '0xBD' */
			return TelephonyCommandType_EVENT
		}
		case 0xBE: { /* '0xBE' */
			return TelephonyCommandType_EVENT
		}
		case 0xBF: { /* '0xBF' */
			return TelephonyCommandType_EVENT
		}
		default: {
			return 0
		}
	}
}

func TelephonyCommandTypeContainerFirstEnumForFieldCommandType(value TelephonyCommandType) (TelephonyCommandTypeContainer, error) {
	for _, sizeValue := range TelephonyCommandTypeContainerValues {
		if sizeValue.CommandType() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing CommandType not found", value)
}
func TelephonyCommandTypeContainerByValue(value uint8) (enum TelephonyCommandTypeContainer, ok bool) {
	switch value {
		case 0x09:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOnHook, true
		case 0xA0:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_0Bytes, true
		case 0xA1:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_1Bytes, true
		case 0xA2:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_2Bytes, true
		case 0xA3:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_3Bytes, true
		case 0xA4:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_4Bytes, true
		case 0xA5:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_5Bytes, true
		case 0xA6:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_6Bytes, true
		case 0xA7:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_7Bytes, true
		case 0xA8:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_8Bytes, true
		case 0xA9:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_9Bytes, true
		case 0xAA:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_10Bytes, true
		case 0xAB:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_11Bytes, true
		case 0xAC:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_12Bytes, true
		case 0xAD:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_13Bytes, true
		case 0xAE:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_14Bytes, true
		case 0xAF:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_15Bytes, true
		case 0xB0:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_16Bytes, true
		case 0xB1:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_17Bytes, true
		case 0xB2:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_18Bytes, true
		case 0xB3:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_19Bytes, true
		case 0xB4:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_20Bytes, true
		case 0xB5:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_21Bytes, true
		case 0xB6:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_22Bytes, true
		case 0xB7:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_23Bytes, true
		case 0xB8:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_24Bytes, true
		case 0xB9:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_25Bytes, true
		case 0xBA:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_26Bytes, true
		case 0xBB:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_27Bytes, true
		case 0xBC:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_28Bytes, true
		case 0xBD:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_29Bytes, true
		case 0xBE:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_30Bytes, true
		case 0xBF:
			return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_31Bytes, true
	}
	return 0, false
}

func TelephonyCommandTypeContainerByName(value string) (enum TelephonyCommandTypeContainer, ok bool) {
	switch value {
	case "TelephonyCommandLineOnHook":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOnHook, true
	case "TelephonyCommandLineOffHook_0Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_0Bytes, true
	case "TelephonyCommandLineOffHook_1Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_1Bytes, true
	case "TelephonyCommandLineOffHook_2Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_2Bytes, true
	case "TelephonyCommandLineOffHook_3Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_3Bytes, true
	case "TelephonyCommandLineOffHook_4Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_4Bytes, true
	case "TelephonyCommandLineOffHook_5Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_5Bytes, true
	case "TelephonyCommandLineOffHook_6Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_6Bytes, true
	case "TelephonyCommandLineOffHook_7Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_7Bytes, true
	case "TelephonyCommandLineOffHook_8Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_8Bytes, true
	case "TelephonyCommandLineOffHook_9Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_9Bytes, true
	case "TelephonyCommandLineOffHook_10Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_10Bytes, true
	case "TelephonyCommandLineOffHook_11Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_11Bytes, true
	case "TelephonyCommandLineOffHook_12Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_12Bytes, true
	case "TelephonyCommandLineOffHook_13Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_13Bytes, true
	case "TelephonyCommandLineOffHook_14Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_14Bytes, true
	case "TelephonyCommandLineOffHook_15Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_15Bytes, true
	case "TelephonyCommandLineOffHook_16Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_16Bytes, true
	case "TelephonyCommandLineOffHook_17Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_17Bytes, true
	case "TelephonyCommandLineOffHook_18Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_18Bytes, true
	case "TelephonyCommandLineOffHook_19Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_19Bytes, true
	case "TelephonyCommandLineOffHook_20Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_20Bytes, true
	case "TelephonyCommandLineOffHook_21Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_21Bytes, true
	case "TelephonyCommandLineOffHook_22Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_22Bytes, true
	case "TelephonyCommandLineOffHook_23Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_23Bytes, true
	case "TelephonyCommandLineOffHook_24Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_24Bytes, true
	case "TelephonyCommandLineOffHook_25Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_25Bytes, true
	case "TelephonyCommandLineOffHook_26Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_26Bytes, true
	case "TelephonyCommandLineOffHook_27Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_27Bytes, true
	case "TelephonyCommandLineOffHook_28Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_28Bytes, true
	case "TelephonyCommandLineOffHook_29Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_29Bytes, true
	case "TelephonyCommandLineOffHook_30Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_30Bytes, true
	case "TelephonyCommandLineOffHook_31Bytes":
		return TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_31Bytes, true
	}
	return 0, false
}

func TelephonyCommandTypeContainerKnows(value uint8)  bool {
	for _, typeValue := range TelephonyCommandTypeContainerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastTelephonyCommandTypeContainer(structType any) TelephonyCommandTypeContainer {
	castFunc := func(typ any) TelephonyCommandTypeContainer {
		if sTelephonyCommandTypeContainer, ok := typ.(TelephonyCommandTypeContainer); ok {
			return sTelephonyCommandTypeContainer
		}
		return 0
	}
	return castFunc(structType)
}

func (m TelephonyCommandTypeContainer) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m TelephonyCommandTypeContainer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TelephonyCommandTypeContainerParse(ctx context.Context, theBytes []byte) (TelephonyCommandTypeContainer, error) {
	return TelephonyCommandTypeContainerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TelephonyCommandTypeContainerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TelephonyCommandTypeContainer, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("TelephonyCommandTypeContainer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TelephonyCommandTypeContainer")
	}
	if enum, ok := TelephonyCommandTypeContainerByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return TelephonyCommandTypeContainer(val), nil
	} else {
		return enum, nil
	}
}

func (e TelephonyCommandTypeContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TelephonyCommandTypeContainer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("TelephonyCommandTypeContainer", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TelephonyCommandTypeContainer) PLC4XEnumName() string {
	switch e {
	case TelephonyCommandTypeContainer_TelephonyCommandLineOnHook:
		return "TelephonyCommandLineOnHook"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_0Bytes:
		return "TelephonyCommandLineOffHook_0Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_1Bytes:
		return "TelephonyCommandLineOffHook_1Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_2Bytes:
		return "TelephonyCommandLineOffHook_2Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_3Bytes:
		return "TelephonyCommandLineOffHook_3Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_4Bytes:
		return "TelephonyCommandLineOffHook_4Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_5Bytes:
		return "TelephonyCommandLineOffHook_5Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_6Bytes:
		return "TelephonyCommandLineOffHook_6Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_7Bytes:
		return "TelephonyCommandLineOffHook_7Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_8Bytes:
		return "TelephonyCommandLineOffHook_8Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_9Bytes:
		return "TelephonyCommandLineOffHook_9Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_10Bytes:
		return "TelephonyCommandLineOffHook_10Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_11Bytes:
		return "TelephonyCommandLineOffHook_11Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_12Bytes:
		return "TelephonyCommandLineOffHook_12Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_13Bytes:
		return "TelephonyCommandLineOffHook_13Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_14Bytes:
		return "TelephonyCommandLineOffHook_14Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_15Bytes:
		return "TelephonyCommandLineOffHook_15Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_16Bytes:
		return "TelephonyCommandLineOffHook_16Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_17Bytes:
		return "TelephonyCommandLineOffHook_17Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_18Bytes:
		return "TelephonyCommandLineOffHook_18Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_19Bytes:
		return "TelephonyCommandLineOffHook_19Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_20Bytes:
		return "TelephonyCommandLineOffHook_20Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_21Bytes:
		return "TelephonyCommandLineOffHook_21Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_22Bytes:
		return "TelephonyCommandLineOffHook_22Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_23Bytes:
		return "TelephonyCommandLineOffHook_23Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_24Bytes:
		return "TelephonyCommandLineOffHook_24Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_25Bytes:
		return "TelephonyCommandLineOffHook_25Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_26Bytes:
		return "TelephonyCommandLineOffHook_26Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_27Bytes:
		return "TelephonyCommandLineOffHook_27Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_28Bytes:
		return "TelephonyCommandLineOffHook_28Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_29Bytes:
		return "TelephonyCommandLineOffHook_29Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_30Bytes:
		return "TelephonyCommandLineOffHook_30Bytes"
	case TelephonyCommandTypeContainer_TelephonyCommandLineOffHook_31Bytes:
		return "TelephonyCommandLineOffHook_31Bytes"
	}
	return ""
}

func (e TelephonyCommandTypeContainer) String() string {
	return e.PLC4XEnumName()
}

