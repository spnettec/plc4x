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

// TriggerControlCommandTypeContainer is an enum
type TriggerControlCommandTypeContainer uint8

type ITriggerControlCommandTypeContainer interface {
	utils.Serializable
	NumBytes() uint8
	CommandType() TriggerControlCommandType
}

const(
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerMin_1Bytes TriggerControlCommandTypeContainer = 0x01
	TriggerControlCommandTypeContainer_TriggerControlCommandIndicatorKill_1Bytes TriggerControlCommandTypeContainer = 0x09
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerMax_1Bytes TriggerControlCommandTypeContainer = 0x79
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent0_2Bytes TriggerControlCommandTypeContainer = 0x02
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent1_2Bytes TriggerControlCommandTypeContainer = 0x0A
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent2_2Bytes TriggerControlCommandTypeContainer = 0x12
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent3_2Bytes TriggerControlCommandTypeContainer = 0x1A
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent4_2Bytes TriggerControlCommandTypeContainer = 0x22
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent5_2Bytes TriggerControlCommandTypeContainer = 0x2A
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent6_2Bytes TriggerControlCommandTypeContainer = 0x32
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent7_2Bytes TriggerControlCommandTypeContainer = 0x3A
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent8_2Bytes TriggerControlCommandTypeContainer = 0x42
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent9_2Bytes TriggerControlCommandTypeContainer = 0x4A
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent10_2Bytes TriggerControlCommandTypeContainer = 0x52
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent11_2Bytes TriggerControlCommandTypeContainer = 0x5A
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent12_2Bytes TriggerControlCommandTypeContainer = 0x62
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent13_2Bytes TriggerControlCommandTypeContainer = 0x6A
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent14_2Bytes TriggerControlCommandTypeContainer = 0x72
	TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent15_2Bytes TriggerControlCommandTypeContainer = 0x7A
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_0Bytes TriggerControlCommandTypeContainer = 0xA0
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_1Bytes TriggerControlCommandTypeContainer = 0xA1
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_2Bytes TriggerControlCommandTypeContainer = 0xA2
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_3Bytes TriggerControlCommandTypeContainer = 0xA3
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_4Bytes TriggerControlCommandTypeContainer = 0xA4
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_5Bytes TriggerControlCommandTypeContainer = 0xA5
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_6Bytes TriggerControlCommandTypeContainer = 0xA6
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_7Bytes TriggerControlCommandTypeContainer = 0xA7
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_8Bytes TriggerControlCommandTypeContainer = 0xA8
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_9Bytes TriggerControlCommandTypeContainer = 0xA9
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_10Bytes TriggerControlCommandTypeContainer = 0xAA
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_11Bytes TriggerControlCommandTypeContainer = 0xAB
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_12Bytes TriggerControlCommandTypeContainer = 0xAC
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_13Bytes TriggerControlCommandTypeContainer = 0xAD
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_14Bytes TriggerControlCommandTypeContainer = 0xAE
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_15Bytes TriggerControlCommandTypeContainer = 0xAF
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_16Bytes TriggerControlCommandTypeContainer = 0xB0
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_17Bytes TriggerControlCommandTypeContainer = 0xB1
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_18Bytes TriggerControlCommandTypeContainer = 0xB2
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_19Bytes TriggerControlCommandTypeContainer = 0xB3
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_20Bytes TriggerControlCommandTypeContainer = 0xB4
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_21Bytes TriggerControlCommandTypeContainer = 0xB5
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_22Bytes TriggerControlCommandTypeContainer = 0xB6
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_23Bytes TriggerControlCommandTypeContainer = 0xB7
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_24Bytes TriggerControlCommandTypeContainer = 0xB8
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_25Bytes TriggerControlCommandTypeContainer = 0xB9
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_26Bytes TriggerControlCommandTypeContainer = 0xBA
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_27Bytes TriggerControlCommandTypeContainer = 0xBB
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_28Bytes TriggerControlCommandTypeContainer = 0xBC
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_29Bytes TriggerControlCommandTypeContainer = 0xBD
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_30Bytes TriggerControlCommandTypeContainer = 0xBE
	TriggerControlCommandTypeContainer_TriggerControlCommandLabel_31Bytes TriggerControlCommandTypeContainer = 0xBF
)

var TriggerControlCommandTypeContainerValues []TriggerControlCommandTypeContainer

func init() {
	_ = errors.New
	TriggerControlCommandTypeContainerValues = []TriggerControlCommandTypeContainer {
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerMin_1Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandIndicatorKill_1Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerMax_1Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent0_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent1_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent2_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent3_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent4_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent5_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent6_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent7_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent8_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent9_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent10_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent11_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent12_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent13_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent14_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent15_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_0Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_1Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_2Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_3Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_4Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_5Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_6Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_7Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_8Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_9Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_10Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_11Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_12Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_13Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_14Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_15Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_16Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_17Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_18Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_19Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_20Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_21Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_22Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_23Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_24Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_25Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_26Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_27Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_28Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_29Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_30Bytes,
		TriggerControlCommandTypeContainer_TriggerControlCommandLabel_31Bytes,
	}
}


func (e TriggerControlCommandTypeContainer) NumBytes() uint8 {
	switch e  {
		case 0x01: { /* '0x01' */
            return 1
		}
		case 0x02: { /* '0x02' */
            return 2
		}
		case 0x09: { /* '0x09' */
            return 1
		}
		case 0x0A: { /* '0x0A' */
            return 2
		}
		case 0x12: { /* '0x12' */
            return 2
		}
		case 0x1A: { /* '0x1A' */
            return 2
		}
		case 0x22: { /* '0x22' */
            return 2
		}
		case 0x2A: { /* '0x2A' */
            return 2
		}
		case 0x32: { /* '0x32' */
            return 2
		}
		case 0x3A: { /* '0x3A' */
            return 2
		}
		case 0x42: { /* '0x42' */
            return 2
		}
		case 0x4A: { /* '0x4A' */
            return 2
		}
		case 0x52: { /* '0x52' */
            return 2
		}
		case 0x5A: { /* '0x5A' */
            return 2
		}
		case 0x62: { /* '0x62' */
            return 2
		}
		case 0x6A: { /* '0x6A' */
            return 2
		}
		case 0x72: { /* '0x72' */
            return 2
		}
		case 0x79: { /* '0x79' */
            return 1
		}
		case 0x7A: { /* '0x7A' */
            return 2
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

func TriggerControlCommandTypeContainerFirstEnumForFieldNumBytes(value uint8) (TriggerControlCommandTypeContainer, error) {
	for _, sizeValue := range TriggerControlCommandTypeContainerValues {
		if sizeValue.NumBytes() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumBytes not found", value)
}

func (e TriggerControlCommandTypeContainer) CommandType() TriggerControlCommandType {
	switch e  {
		case 0x01: { /* '0x01' */
			return TriggerControlCommandType_TRIGGER_MIN
		}
		case 0x02: { /* '0x02' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x09: { /* '0x09' */
			return TriggerControlCommandType_INDICATOR_KILL
		}
		case 0x0A: { /* '0x0A' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x12: { /* '0x12' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x1A: { /* '0x1A' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x22: { /* '0x22' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x2A: { /* '0x2A' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x32: { /* '0x32' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x3A: { /* '0x3A' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x42: { /* '0x42' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x4A: { /* '0x4A' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x52: { /* '0x52' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x5A: { /* '0x5A' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x62: { /* '0x62' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x6A: { /* '0x6A' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x72: { /* '0x72' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0x79: { /* '0x79' */
			return TriggerControlCommandType_TRIGGER_MAX
		}
		case 0x7A: { /* '0x7A' */
			return TriggerControlCommandType_TRIGGER_EVENT
		}
		case 0xA0: { /* '0xA0' */
			return TriggerControlCommandType_LABEL
		}
		case 0xA1: { /* '0xA1' */
			return TriggerControlCommandType_LABEL
		}
		case 0xA2: { /* '0xA2' */
			return TriggerControlCommandType_LABEL
		}
		case 0xA3: { /* '0xA3' */
			return TriggerControlCommandType_LABEL
		}
		case 0xA4: { /* '0xA4' */
			return TriggerControlCommandType_LABEL
		}
		case 0xA5: { /* '0xA5' */
			return TriggerControlCommandType_LABEL
		}
		case 0xA6: { /* '0xA6' */
			return TriggerControlCommandType_LABEL
		}
		case 0xA7: { /* '0xA7' */
			return TriggerControlCommandType_LABEL
		}
		case 0xA8: { /* '0xA8' */
			return TriggerControlCommandType_LABEL
		}
		case 0xA9: { /* '0xA9' */
			return TriggerControlCommandType_LABEL
		}
		case 0xAA: { /* '0xAA' */
			return TriggerControlCommandType_LABEL
		}
		case 0xAB: { /* '0xAB' */
			return TriggerControlCommandType_LABEL
		}
		case 0xAC: { /* '0xAC' */
			return TriggerControlCommandType_LABEL
		}
		case 0xAD: { /* '0xAD' */
			return TriggerControlCommandType_LABEL
		}
		case 0xAE: { /* '0xAE' */
			return TriggerControlCommandType_LABEL
		}
		case 0xAF: { /* '0xAF' */
			return TriggerControlCommandType_LABEL
		}
		case 0xB0: { /* '0xB0' */
			return TriggerControlCommandType_LABEL
		}
		case 0xB1: { /* '0xB1' */
			return TriggerControlCommandType_LABEL
		}
		case 0xB2: { /* '0xB2' */
			return TriggerControlCommandType_LABEL
		}
		case 0xB3: { /* '0xB3' */
			return TriggerControlCommandType_LABEL
		}
		case 0xB4: { /* '0xB4' */
			return TriggerControlCommandType_LABEL
		}
		case 0xB5: { /* '0xB5' */
			return TriggerControlCommandType_LABEL
		}
		case 0xB6: { /* '0xB6' */
			return TriggerControlCommandType_LABEL
		}
		case 0xB7: { /* '0xB7' */
			return TriggerControlCommandType_LABEL
		}
		case 0xB8: { /* '0xB8' */
			return TriggerControlCommandType_LABEL
		}
		case 0xB9: { /* '0xB9' */
			return TriggerControlCommandType_LABEL
		}
		case 0xBA: { /* '0xBA' */
			return TriggerControlCommandType_LABEL
		}
		case 0xBB: { /* '0xBB' */
			return TriggerControlCommandType_LABEL
		}
		case 0xBC: { /* '0xBC' */
			return TriggerControlCommandType_LABEL
		}
		case 0xBD: { /* '0xBD' */
			return TriggerControlCommandType_LABEL
		}
		case 0xBE: { /* '0xBE' */
			return TriggerControlCommandType_LABEL
		}
		case 0xBF: { /* '0xBF' */
			return TriggerControlCommandType_LABEL
		}
		default: {
			return 0
		}
	}
}

func TriggerControlCommandTypeContainerFirstEnumForFieldCommandType(value TriggerControlCommandType) (TriggerControlCommandTypeContainer, error) {
	for _, sizeValue := range TriggerControlCommandTypeContainerValues {
		if sizeValue.CommandType() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing CommandType not found", value)
}
func TriggerControlCommandTypeContainerByValue(value uint8) (enum TriggerControlCommandTypeContainer, ok bool) {
	switch value {
		case 0x01:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerMin_1Bytes, true
		case 0x02:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent0_2Bytes, true
		case 0x09:
			return TriggerControlCommandTypeContainer_TriggerControlCommandIndicatorKill_1Bytes, true
		case 0x0A:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent1_2Bytes, true
		case 0x12:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent2_2Bytes, true
		case 0x1A:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent3_2Bytes, true
		case 0x22:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent4_2Bytes, true
		case 0x2A:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent5_2Bytes, true
		case 0x32:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent6_2Bytes, true
		case 0x3A:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent7_2Bytes, true
		case 0x42:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent8_2Bytes, true
		case 0x4A:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent9_2Bytes, true
		case 0x52:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent10_2Bytes, true
		case 0x5A:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent11_2Bytes, true
		case 0x62:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent12_2Bytes, true
		case 0x6A:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent13_2Bytes, true
		case 0x72:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent14_2Bytes, true
		case 0x79:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerMax_1Bytes, true
		case 0x7A:
			return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent15_2Bytes, true
		case 0xA0:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_0Bytes, true
		case 0xA1:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_1Bytes, true
		case 0xA2:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_2Bytes, true
		case 0xA3:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_3Bytes, true
		case 0xA4:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_4Bytes, true
		case 0xA5:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_5Bytes, true
		case 0xA6:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_6Bytes, true
		case 0xA7:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_7Bytes, true
		case 0xA8:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_8Bytes, true
		case 0xA9:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_9Bytes, true
		case 0xAA:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_10Bytes, true
		case 0xAB:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_11Bytes, true
		case 0xAC:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_12Bytes, true
		case 0xAD:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_13Bytes, true
		case 0xAE:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_14Bytes, true
		case 0xAF:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_15Bytes, true
		case 0xB0:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_16Bytes, true
		case 0xB1:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_17Bytes, true
		case 0xB2:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_18Bytes, true
		case 0xB3:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_19Bytes, true
		case 0xB4:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_20Bytes, true
		case 0xB5:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_21Bytes, true
		case 0xB6:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_22Bytes, true
		case 0xB7:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_23Bytes, true
		case 0xB8:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_24Bytes, true
		case 0xB9:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_25Bytes, true
		case 0xBA:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_26Bytes, true
		case 0xBB:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_27Bytes, true
		case 0xBC:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_28Bytes, true
		case 0xBD:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_29Bytes, true
		case 0xBE:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_30Bytes, true
		case 0xBF:
			return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_31Bytes, true
	}
	return 0, false
}

func TriggerControlCommandTypeContainerByName(value string) (enum TriggerControlCommandTypeContainer, ok bool) {
	switch value {
	case "TriggerControlCommandTriggerMin_1Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerMin_1Bytes, true
	case "TriggerControlCommandTriggerEvent0_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent0_2Bytes, true
	case "TriggerControlCommandIndicatorKill_1Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandIndicatorKill_1Bytes, true
	case "TriggerControlCommandTriggerEvent1_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent1_2Bytes, true
	case "TriggerControlCommandTriggerEvent2_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent2_2Bytes, true
	case "TriggerControlCommandTriggerEvent3_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent3_2Bytes, true
	case "TriggerControlCommandTriggerEvent4_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent4_2Bytes, true
	case "TriggerControlCommandTriggerEvent5_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent5_2Bytes, true
	case "TriggerControlCommandTriggerEvent6_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent6_2Bytes, true
	case "TriggerControlCommandTriggerEvent7_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent7_2Bytes, true
	case "TriggerControlCommandTriggerEvent8_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent8_2Bytes, true
	case "TriggerControlCommandTriggerEvent9_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent9_2Bytes, true
	case "TriggerControlCommandTriggerEvent10_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent10_2Bytes, true
	case "TriggerControlCommandTriggerEvent11_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent11_2Bytes, true
	case "TriggerControlCommandTriggerEvent12_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent12_2Bytes, true
	case "TriggerControlCommandTriggerEvent13_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent13_2Bytes, true
	case "TriggerControlCommandTriggerEvent14_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent14_2Bytes, true
	case "TriggerControlCommandTriggerMax_1Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerMax_1Bytes, true
	case "TriggerControlCommandTriggerEvent15_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent15_2Bytes, true
	case "TriggerControlCommandLabel_0Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_0Bytes, true
	case "TriggerControlCommandLabel_1Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_1Bytes, true
	case "TriggerControlCommandLabel_2Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_2Bytes, true
	case "TriggerControlCommandLabel_3Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_3Bytes, true
	case "TriggerControlCommandLabel_4Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_4Bytes, true
	case "TriggerControlCommandLabel_5Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_5Bytes, true
	case "TriggerControlCommandLabel_6Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_6Bytes, true
	case "TriggerControlCommandLabel_7Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_7Bytes, true
	case "TriggerControlCommandLabel_8Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_8Bytes, true
	case "TriggerControlCommandLabel_9Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_9Bytes, true
	case "TriggerControlCommandLabel_10Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_10Bytes, true
	case "TriggerControlCommandLabel_11Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_11Bytes, true
	case "TriggerControlCommandLabel_12Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_12Bytes, true
	case "TriggerControlCommandLabel_13Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_13Bytes, true
	case "TriggerControlCommandLabel_14Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_14Bytes, true
	case "TriggerControlCommandLabel_15Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_15Bytes, true
	case "TriggerControlCommandLabel_16Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_16Bytes, true
	case "TriggerControlCommandLabel_17Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_17Bytes, true
	case "TriggerControlCommandLabel_18Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_18Bytes, true
	case "TriggerControlCommandLabel_19Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_19Bytes, true
	case "TriggerControlCommandLabel_20Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_20Bytes, true
	case "TriggerControlCommandLabel_21Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_21Bytes, true
	case "TriggerControlCommandLabel_22Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_22Bytes, true
	case "TriggerControlCommandLabel_23Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_23Bytes, true
	case "TriggerControlCommandLabel_24Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_24Bytes, true
	case "TriggerControlCommandLabel_25Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_25Bytes, true
	case "TriggerControlCommandLabel_26Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_26Bytes, true
	case "TriggerControlCommandLabel_27Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_27Bytes, true
	case "TriggerControlCommandLabel_28Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_28Bytes, true
	case "TriggerControlCommandLabel_29Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_29Bytes, true
	case "TriggerControlCommandLabel_30Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_30Bytes, true
	case "TriggerControlCommandLabel_31Bytes":
		return TriggerControlCommandTypeContainer_TriggerControlCommandLabel_31Bytes, true
	}
	return 0, false
}

func TriggerControlCommandTypeContainerKnows(value uint8)  bool {
	for _, typeValue := range TriggerControlCommandTypeContainerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastTriggerControlCommandTypeContainer(structType interface{}) TriggerControlCommandTypeContainer {
	castFunc := func(typ interface{}) TriggerControlCommandTypeContainer {
		if sTriggerControlCommandTypeContainer, ok := typ.(TriggerControlCommandTypeContainer); ok {
			return sTriggerControlCommandTypeContainer
		}
		return 0
	}
	return castFunc(structType)
}

func (m TriggerControlCommandTypeContainer) GetLengthInBits() uint16 {
	return 8
}

func (m TriggerControlCommandTypeContainer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TriggerControlCommandTypeContainerParse(readBuffer utils.ReadBuffer) (TriggerControlCommandTypeContainer, error) {
	val, err := readBuffer.ReadUint8("TriggerControlCommandTypeContainer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TriggerControlCommandTypeContainer")
	}
	if enum, ok := TriggerControlCommandTypeContainerByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return TriggerControlCommandTypeContainer(val), nil
	} else {
		return enum, nil
	}
}

func (e TriggerControlCommandTypeContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TriggerControlCommandTypeContainer) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("TriggerControlCommandTypeContainer", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TriggerControlCommandTypeContainer) PLC4XEnumName() string {
	switch e {
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerMin_1Bytes:
		return "TriggerControlCommandTriggerMin_1Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent0_2Bytes:
		return "TriggerControlCommandTriggerEvent0_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandIndicatorKill_1Bytes:
		return "TriggerControlCommandIndicatorKill_1Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent1_2Bytes:
		return "TriggerControlCommandTriggerEvent1_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent2_2Bytes:
		return "TriggerControlCommandTriggerEvent2_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent3_2Bytes:
		return "TriggerControlCommandTriggerEvent3_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent4_2Bytes:
		return "TriggerControlCommandTriggerEvent4_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent5_2Bytes:
		return "TriggerControlCommandTriggerEvent5_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent6_2Bytes:
		return "TriggerControlCommandTriggerEvent6_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent7_2Bytes:
		return "TriggerControlCommandTriggerEvent7_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent8_2Bytes:
		return "TriggerControlCommandTriggerEvent8_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent9_2Bytes:
		return "TriggerControlCommandTriggerEvent9_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent10_2Bytes:
		return "TriggerControlCommandTriggerEvent10_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent11_2Bytes:
		return "TriggerControlCommandTriggerEvent11_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent12_2Bytes:
		return "TriggerControlCommandTriggerEvent12_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent13_2Bytes:
		return "TriggerControlCommandTriggerEvent13_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent14_2Bytes:
		return "TriggerControlCommandTriggerEvent14_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerMax_1Bytes:
		return "TriggerControlCommandTriggerMax_1Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandTriggerEvent15_2Bytes:
		return "TriggerControlCommandTriggerEvent15_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_0Bytes:
		return "TriggerControlCommandLabel_0Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_1Bytes:
		return "TriggerControlCommandLabel_1Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_2Bytes:
		return "TriggerControlCommandLabel_2Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_3Bytes:
		return "TriggerControlCommandLabel_3Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_4Bytes:
		return "TriggerControlCommandLabel_4Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_5Bytes:
		return "TriggerControlCommandLabel_5Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_6Bytes:
		return "TriggerControlCommandLabel_6Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_7Bytes:
		return "TriggerControlCommandLabel_7Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_8Bytes:
		return "TriggerControlCommandLabel_8Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_9Bytes:
		return "TriggerControlCommandLabel_9Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_10Bytes:
		return "TriggerControlCommandLabel_10Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_11Bytes:
		return "TriggerControlCommandLabel_11Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_12Bytes:
		return "TriggerControlCommandLabel_12Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_13Bytes:
		return "TriggerControlCommandLabel_13Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_14Bytes:
		return "TriggerControlCommandLabel_14Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_15Bytes:
		return "TriggerControlCommandLabel_15Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_16Bytes:
		return "TriggerControlCommandLabel_16Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_17Bytes:
		return "TriggerControlCommandLabel_17Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_18Bytes:
		return "TriggerControlCommandLabel_18Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_19Bytes:
		return "TriggerControlCommandLabel_19Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_20Bytes:
		return "TriggerControlCommandLabel_20Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_21Bytes:
		return "TriggerControlCommandLabel_21Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_22Bytes:
		return "TriggerControlCommandLabel_22Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_23Bytes:
		return "TriggerControlCommandLabel_23Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_24Bytes:
		return "TriggerControlCommandLabel_24Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_25Bytes:
		return "TriggerControlCommandLabel_25Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_26Bytes:
		return "TriggerControlCommandLabel_26Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_27Bytes:
		return "TriggerControlCommandLabel_27Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_28Bytes:
		return "TriggerControlCommandLabel_28Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_29Bytes:
		return "TriggerControlCommandLabel_29Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_30Bytes:
		return "TriggerControlCommandLabel_30Bytes"
	case TriggerControlCommandTypeContainer_TriggerControlCommandLabel_31Bytes:
		return "TriggerControlCommandLabel_31Bytes"
	}
	return ""
}

func (e TriggerControlCommandTypeContainer) String() string {
	return e.PLC4XEnumName()
}

