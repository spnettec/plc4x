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

// BACnetDaysOfWeek is an enum
type BACnetDaysOfWeek uint8

type IBACnetDaysOfWeek interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetDaysOfWeek_MONDAY    BACnetDaysOfWeek = 0
	BACnetDaysOfWeek_TUESDAY   BACnetDaysOfWeek = 1
	BACnetDaysOfWeek_WEDNESDAY BACnetDaysOfWeek = 2
	BACnetDaysOfWeek_THURSDAY  BACnetDaysOfWeek = 3
	BACnetDaysOfWeek_FRIDAY    BACnetDaysOfWeek = 4
	BACnetDaysOfWeek_SATURDAY  BACnetDaysOfWeek = 5
	BACnetDaysOfWeek_SUNDAY    BACnetDaysOfWeek = 6
)

var BACnetDaysOfWeekValues []BACnetDaysOfWeek

func init() {
	_ = errors.New
	BACnetDaysOfWeekValues = []BACnetDaysOfWeek{
		BACnetDaysOfWeek_MONDAY,
		BACnetDaysOfWeek_TUESDAY,
		BACnetDaysOfWeek_WEDNESDAY,
		BACnetDaysOfWeek_THURSDAY,
		BACnetDaysOfWeek_FRIDAY,
		BACnetDaysOfWeek_SATURDAY,
		BACnetDaysOfWeek_SUNDAY,
	}
}

func BACnetDaysOfWeekByValue(value uint8) BACnetDaysOfWeek {
	switch value {
	case 0:
		return BACnetDaysOfWeek_MONDAY
	case 1:
		return BACnetDaysOfWeek_TUESDAY
	case 2:
		return BACnetDaysOfWeek_WEDNESDAY
	case 3:
		return BACnetDaysOfWeek_THURSDAY
	case 4:
		return BACnetDaysOfWeek_FRIDAY
	case 5:
		return BACnetDaysOfWeek_SATURDAY
	case 6:
		return BACnetDaysOfWeek_SUNDAY
	}
	return 0
}

func BACnetDaysOfWeekByName(value string) (enum BACnetDaysOfWeek, ok bool) {
	ok = true
	switch value {
	case "MONDAY":
		enum = BACnetDaysOfWeek_MONDAY
	case "TUESDAY":
		enum = BACnetDaysOfWeek_TUESDAY
	case "WEDNESDAY":
		enum = BACnetDaysOfWeek_WEDNESDAY
	case "THURSDAY":
		enum = BACnetDaysOfWeek_THURSDAY
	case "FRIDAY":
		enum = BACnetDaysOfWeek_FRIDAY
	case "SATURDAY":
		enum = BACnetDaysOfWeek_SATURDAY
	case "SUNDAY":
		enum = BACnetDaysOfWeek_SUNDAY
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetDaysOfWeekKnows(value uint8) bool {
	for _, typeValue := range BACnetDaysOfWeekValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetDaysOfWeek(structType interface{}) BACnetDaysOfWeek {
	castFunc := func(typ interface{}) BACnetDaysOfWeek {
		if sBACnetDaysOfWeek, ok := typ.(BACnetDaysOfWeek); ok {
			return sBACnetDaysOfWeek
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetDaysOfWeek) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetDaysOfWeek) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDaysOfWeekParse(readBuffer utils.ReadBuffer) (BACnetDaysOfWeek, error) {
	val, err := readBuffer.ReadUint8("BACnetDaysOfWeek", 8)
	if err != nil {
		return 0, nil
	}
	return BACnetDaysOfWeekByValue(val), nil
}

func (e BACnetDaysOfWeek) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetDaysOfWeek", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetDaysOfWeek) PLC4XEnumName() string {
	switch e {
	case BACnetDaysOfWeek_MONDAY:
		return "MONDAY"
	case BACnetDaysOfWeek_TUESDAY:
		return "TUESDAY"
	case BACnetDaysOfWeek_WEDNESDAY:
		return "WEDNESDAY"
	case BACnetDaysOfWeek_THURSDAY:
		return "THURSDAY"
	case BACnetDaysOfWeek_FRIDAY:
		return "FRIDAY"
	case BACnetDaysOfWeek_SATURDAY:
		return "SATURDAY"
	case BACnetDaysOfWeek_SUNDAY:
		return "SUNDAY"
	}
	return ""
}

func (e BACnetDaysOfWeek) String() string {
	return e.PLC4XEnumName()
}
