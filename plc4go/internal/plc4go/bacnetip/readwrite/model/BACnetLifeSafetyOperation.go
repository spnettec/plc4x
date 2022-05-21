/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLifeSafetyOperation is an enum
type BACnetLifeSafetyOperation uint16

type IBACnetLifeSafetyOperation interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetLifeSafetyOperation_NONE                     BACnetLifeSafetyOperation = 0
	BACnetLifeSafetyOperation_SILENCE                  BACnetLifeSafetyOperation = 1
	BACnetLifeSafetyOperation_SILENCE_AUDIBLE          BACnetLifeSafetyOperation = 2
	BACnetLifeSafetyOperation_SILENCE_VISUAL           BACnetLifeSafetyOperation = 3
	BACnetLifeSafetyOperation_RESET                    BACnetLifeSafetyOperation = 4
	BACnetLifeSafetyOperation_RESET_ALARM              BACnetLifeSafetyOperation = 5
	BACnetLifeSafetyOperation_RESET_FAULT              BACnetLifeSafetyOperation = 6
	BACnetLifeSafetyOperation_UNSILENCE                BACnetLifeSafetyOperation = 7
	BACnetLifeSafetyOperation_UNSILENCE_AUDIBLE        BACnetLifeSafetyOperation = 8
	BACnetLifeSafetyOperation_UNSILENCE_VISUAL         BACnetLifeSafetyOperation = 9
	BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE BACnetLifeSafetyOperation = 0xFFFF
)

var BACnetLifeSafetyOperationValues []BACnetLifeSafetyOperation

func init() {
	_ = errors.New
	BACnetLifeSafetyOperationValues = []BACnetLifeSafetyOperation{
		BACnetLifeSafetyOperation_NONE,
		BACnetLifeSafetyOperation_SILENCE,
		BACnetLifeSafetyOperation_SILENCE_AUDIBLE,
		BACnetLifeSafetyOperation_SILENCE_VISUAL,
		BACnetLifeSafetyOperation_RESET,
		BACnetLifeSafetyOperation_RESET_ALARM,
		BACnetLifeSafetyOperation_RESET_FAULT,
		BACnetLifeSafetyOperation_UNSILENCE,
		BACnetLifeSafetyOperation_UNSILENCE_AUDIBLE,
		BACnetLifeSafetyOperation_UNSILENCE_VISUAL,
		BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLifeSafetyOperationByValue(value uint16) BACnetLifeSafetyOperation {
	switch value {
	case 0:
		return BACnetLifeSafetyOperation_NONE
	case 0xFFFF:
		return BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetLifeSafetyOperation_SILENCE
	case 2:
		return BACnetLifeSafetyOperation_SILENCE_AUDIBLE
	case 3:
		return BACnetLifeSafetyOperation_SILENCE_VISUAL
	case 4:
		return BACnetLifeSafetyOperation_RESET
	case 5:
		return BACnetLifeSafetyOperation_RESET_ALARM
	case 6:
		return BACnetLifeSafetyOperation_RESET_FAULT
	case 7:
		return BACnetLifeSafetyOperation_UNSILENCE
	case 8:
		return BACnetLifeSafetyOperation_UNSILENCE_AUDIBLE
	case 9:
		return BACnetLifeSafetyOperation_UNSILENCE_VISUAL
	}
	return 0
}

func BACnetLifeSafetyOperationByName(value string) BACnetLifeSafetyOperation {
	switch value {
	case "NONE":
		return BACnetLifeSafetyOperation_NONE
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE
	case "SILENCE":
		return BACnetLifeSafetyOperation_SILENCE
	case "SILENCE_AUDIBLE":
		return BACnetLifeSafetyOperation_SILENCE_AUDIBLE
	case "SILENCE_VISUAL":
		return BACnetLifeSafetyOperation_SILENCE_VISUAL
	case "RESET":
		return BACnetLifeSafetyOperation_RESET
	case "RESET_ALARM":
		return BACnetLifeSafetyOperation_RESET_ALARM
	case "RESET_FAULT":
		return BACnetLifeSafetyOperation_RESET_FAULT
	case "UNSILENCE":
		return BACnetLifeSafetyOperation_UNSILENCE
	case "UNSILENCE_AUDIBLE":
		return BACnetLifeSafetyOperation_UNSILENCE_AUDIBLE
	case "UNSILENCE_VISUAL":
		return BACnetLifeSafetyOperation_UNSILENCE_VISUAL
	}
	return 0
}

func BACnetLifeSafetyOperationKnows(value uint16) bool {
	for _, typeValue := range BACnetLifeSafetyOperationValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLifeSafetyOperation(structType interface{}) BACnetLifeSafetyOperation {
	castFunc := func(typ interface{}) BACnetLifeSafetyOperation {
		if sBACnetLifeSafetyOperation, ok := typ.(BACnetLifeSafetyOperation); ok {
			return sBACnetLifeSafetyOperation
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLifeSafetyOperation) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetLifeSafetyOperation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLifeSafetyOperationParse(readBuffer utils.ReadBuffer) (BACnetLifeSafetyOperation, error) {
	val, err := readBuffer.ReadUint16("BACnetLifeSafetyOperation", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetLifeSafetyOperationByValue(val), nil
}

func (e BACnetLifeSafetyOperation) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetLifeSafetyOperation", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetLifeSafetyOperation) name() string {
	switch e {
	case BACnetLifeSafetyOperation_NONE:
		return "NONE"
	case BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLifeSafetyOperation_SILENCE:
		return "SILENCE"
	case BACnetLifeSafetyOperation_SILENCE_AUDIBLE:
		return "SILENCE_AUDIBLE"
	case BACnetLifeSafetyOperation_SILENCE_VISUAL:
		return "SILENCE_VISUAL"
	case BACnetLifeSafetyOperation_RESET:
		return "RESET"
	case BACnetLifeSafetyOperation_RESET_ALARM:
		return "RESET_ALARM"
	case BACnetLifeSafetyOperation_RESET_FAULT:
		return "RESET_FAULT"
	case BACnetLifeSafetyOperation_UNSILENCE:
		return "UNSILENCE"
	case BACnetLifeSafetyOperation_UNSILENCE_AUDIBLE:
		return "UNSILENCE_AUDIBLE"
	case BACnetLifeSafetyOperation_UNSILENCE_VISUAL:
		return "UNSILENCE_VISUAL"
	}
	return ""
}

func (e BACnetLifeSafetyOperation) String() string {
	return e.name()
}
