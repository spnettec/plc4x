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

// BACnetStatusFlags is an enum
type BACnetStatusFlags uint8

type IBACnetStatusFlags interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetStatusFlags_IN_ALARM       BACnetStatusFlags = 0
	BACnetStatusFlags_FAULT          BACnetStatusFlags = 1
	BACnetStatusFlags_OVERRIDDEN     BACnetStatusFlags = 2
	BACnetStatusFlags_OUT_OF_SERVICE BACnetStatusFlags = 3
)

var BACnetStatusFlagsValues []BACnetStatusFlags

func init() {
	_ = errors.New
	BACnetStatusFlagsValues = []BACnetStatusFlags{
		BACnetStatusFlags_IN_ALARM,
		BACnetStatusFlags_FAULT,
		BACnetStatusFlags_OVERRIDDEN,
		BACnetStatusFlags_OUT_OF_SERVICE,
	}
}

func BACnetStatusFlagsByValue(value uint8) BACnetStatusFlags {
	switch value {
	case 0:
		return BACnetStatusFlags_IN_ALARM
	case 1:
		return BACnetStatusFlags_FAULT
	case 2:
		return BACnetStatusFlags_OVERRIDDEN
	case 3:
		return BACnetStatusFlags_OUT_OF_SERVICE
	}
	return 0
}

func BACnetStatusFlagsByName(value string) (enum BACnetStatusFlags, ok bool) {
	ok = true
	switch value {
	case "IN_ALARM":
		enum = BACnetStatusFlags_IN_ALARM
	case "FAULT":
		enum = BACnetStatusFlags_FAULT
	case "OVERRIDDEN":
		enum = BACnetStatusFlags_OVERRIDDEN
	case "OUT_OF_SERVICE":
		enum = BACnetStatusFlags_OUT_OF_SERVICE
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetStatusFlagsKnows(value uint8) bool {
	for _, typeValue := range BACnetStatusFlagsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetStatusFlags(structType interface{}) BACnetStatusFlags {
	castFunc := func(typ interface{}) BACnetStatusFlags {
		if sBACnetStatusFlags, ok := typ.(BACnetStatusFlags); ok {
			return sBACnetStatusFlags
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetStatusFlags) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetStatusFlags) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetStatusFlagsParse(readBuffer utils.ReadBuffer) (BACnetStatusFlags, error) {
	val, err := readBuffer.ReadUint8("BACnetStatusFlags", 8)
	if err != nil {
		return 0, nil
	}
	return BACnetStatusFlagsByValue(val), nil
}

func (e BACnetStatusFlags) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetStatusFlags", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetStatusFlags) PLC4XEnumName() string {
	switch e {
	case BACnetStatusFlags_IN_ALARM:
		return "IN_ALARM"
	case BACnetStatusFlags_FAULT:
		return "FAULT"
	case BACnetStatusFlags_OVERRIDDEN:
		return "OVERRIDDEN"
	case BACnetStatusFlags_OUT_OF_SERVICE:
		return "OUT_OF_SERVICE"
	}
	return ""
}

func (e BACnetStatusFlags) String() string {
	return e.PLC4XEnumName()
}
