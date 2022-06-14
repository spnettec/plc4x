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

// BACnetEventState is an enum
type BACnetEventState uint16

type IBACnetEventState interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetEventState_NORMAL                   BACnetEventState = 0
	BACnetEventState_FAULT                    BACnetEventState = 1
	BACnetEventState_OFFNORMAL                BACnetEventState = 2
	BACnetEventState_HIGH_LIMIT               BACnetEventState = 3
	BACnetEventState_LOW_LIMIT                BACnetEventState = 4
	BACnetEventState_LIFE_SAVETY_ALARM        BACnetEventState = 5
	BACnetEventState_VENDOR_PROPRIETARY_VALUE BACnetEventState = 0xFFFF
)

var BACnetEventStateValues []BACnetEventState

func init() {
	_ = errors.New
	BACnetEventStateValues = []BACnetEventState{
		BACnetEventState_NORMAL,
		BACnetEventState_FAULT,
		BACnetEventState_OFFNORMAL,
		BACnetEventState_HIGH_LIMIT,
		BACnetEventState_LOW_LIMIT,
		BACnetEventState_LIFE_SAVETY_ALARM,
		BACnetEventState_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetEventStateByValue(value uint16) BACnetEventState {
	switch value {
	case 0:
		return BACnetEventState_NORMAL
	case 0xFFFF:
		return BACnetEventState_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetEventState_FAULT
	case 2:
		return BACnetEventState_OFFNORMAL
	case 3:
		return BACnetEventState_HIGH_LIMIT
	case 4:
		return BACnetEventState_LOW_LIMIT
	case 5:
		return BACnetEventState_LIFE_SAVETY_ALARM
	}
	return 0
}

func BACnetEventStateByName(value string) BACnetEventState {
	switch value {
	case "NORMAL":
		return BACnetEventState_NORMAL
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetEventState_VENDOR_PROPRIETARY_VALUE
	case "FAULT":
		return BACnetEventState_FAULT
	case "OFFNORMAL":
		return BACnetEventState_OFFNORMAL
	case "HIGH_LIMIT":
		return BACnetEventState_HIGH_LIMIT
	case "LOW_LIMIT":
		return BACnetEventState_LOW_LIMIT
	case "LIFE_SAVETY_ALARM":
		return BACnetEventState_LIFE_SAVETY_ALARM
	}
	return 0
}

func BACnetEventStateKnows(value uint16) bool {
	for _, typeValue := range BACnetEventStateValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetEventState(structType interface{}) BACnetEventState {
	castFunc := func(typ interface{}) BACnetEventState {
		if sBACnetEventState, ok := typ.(BACnetEventState); ok {
			return sBACnetEventState
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetEventState) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetEventState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventStateParse(readBuffer utils.ReadBuffer) (BACnetEventState, error) {
	val, err := readBuffer.ReadUint16("BACnetEventState", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetEventStateByValue(val), nil
}

func (e BACnetEventState) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetEventState", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetEventState) name() string {
	switch e {
	case BACnetEventState_NORMAL:
		return "NORMAL"
	case BACnetEventState_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetEventState_FAULT:
		return "FAULT"
	case BACnetEventState_OFFNORMAL:
		return "OFFNORMAL"
	case BACnetEventState_HIGH_LIMIT:
		return "HIGH_LIMIT"
	case BACnetEventState_LOW_LIMIT:
		return "LOW_LIMIT"
	case BACnetEventState_LIFE_SAVETY_ALARM:
		return "LIFE_SAVETY_ALARM"
	}
	return ""
}

func (e BACnetEventState) String() string {
	return e.name()
}
