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

// BACnetRestartReason is an enum
type BACnetRestartReason uint8

type IBACnetRestartReason interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetRestartReason_UNKNOWN                  BACnetRestartReason = 0
	BACnetRestartReason_COLDSTART                BACnetRestartReason = 1
	BACnetRestartReason_WARMSTART                BACnetRestartReason = 2
	BACnetRestartReason_DETECTED_POWER_LOST      BACnetRestartReason = 3
	BACnetRestartReason_DETECTED_POWERED_OFF     BACnetRestartReason = 4
	BACnetRestartReason_HARDWARE_WATCHDOG        BACnetRestartReason = 5
	BACnetRestartReason_SOFTWARE_WATCHDOG        BACnetRestartReason = 6
	BACnetRestartReason_SUSPENDED                BACnetRestartReason = 7
	BACnetRestartReason_ACTIVATE_CHANGES         BACnetRestartReason = 8
	BACnetRestartReason_VENDOR_PROPRIETARY_VALUE BACnetRestartReason = 0xFF
)

var BACnetRestartReasonValues []BACnetRestartReason

func init() {
	_ = errors.New
	BACnetRestartReasonValues = []BACnetRestartReason{
		BACnetRestartReason_UNKNOWN,
		BACnetRestartReason_COLDSTART,
		BACnetRestartReason_WARMSTART,
		BACnetRestartReason_DETECTED_POWER_LOST,
		BACnetRestartReason_DETECTED_POWERED_OFF,
		BACnetRestartReason_HARDWARE_WATCHDOG,
		BACnetRestartReason_SOFTWARE_WATCHDOG,
		BACnetRestartReason_SUSPENDED,
		BACnetRestartReason_ACTIVATE_CHANGES,
		BACnetRestartReason_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetRestartReasonByValue(value uint8) BACnetRestartReason {
	switch value {
	case 0:
		return BACnetRestartReason_UNKNOWN
	case 0xFF:
		return BACnetRestartReason_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetRestartReason_COLDSTART
	case 2:
		return BACnetRestartReason_WARMSTART
	case 3:
		return BACnetRestartReason_DETECTED_POWER_LOST
	case 4:
		return BACnetRestartReason_DETECTED_POWERED_OFF
	case 5:
		return BACnetRestartReason_HARDWARE_WATCHDOG
	case 6:
		return BACnetRestartReason_SOFTWARE_WATCHDOG
	case 7:
		return BACnetRestartReason_SUSPENDED
	case 8:
		return BACnetRestartReason_ACTIVATE_CHANGES
	}
	return 0
}

func BACnetRestartReasonByName(value string) BACnetRestartReason {
	switch value {
	case "UNKNOWN":
		return BACnetRestartReason_UNKNOWN
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetRestartReason_VENDOR_PROPRIETARY_VALUE
	case "COLDSTART":
		return BACnetRestartReason_COLDSTART
	case "WARMSTART":
		return BACnetRestartReason_WARMSTART
	case "DETECTED_POWER_LOST":
		return BACnetRestartReason_DETECTED_POWER_LOST
	case "DETECTED_POWERED_OFF":
		return BACnetRestartReason_DETECTED_POWERED_OFF
	case "HARDWARE_WATCHDOG":
		return BACnetRestartReason_HARDWARE_WATCHDOG
	case "SOFTWARE_WATCHDOG":
		return BACnetRestartReason_SOFTWARE_WATCHDOG
	case "SUSPENDED":
		return BACnetRestartReason_SUSPENDED
	case "ACTIVATE_CHANGES":
		return BACnetRestartReason_ACTIVATE_CHANGES
	}
	return 0
}

func BACnetRestartReasonKnows(value uint8) bool {
	for _, typeValue := range BACnetRestartReasonValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetRestartReason(structType interface{}) BACnetRestartReason {
	castFunc := func(typ interface{}) BACnetRestartReason {
		if sBACnetRestartReason, ok := typ.(BACnetRestartReason); ok {
			return sBACnetRestartReason
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetRestartReason) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetRestartReason) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRestartReasonParse(readBuffer utils.ReadBuffer) (BACnetRestartReason, error) {
	val, err := readBuffer.ReadUint8("BACnetRestartReason", 8)
	if err != nil {
		return 0, nil
	}
	return BACnetRestartReasonByValue(val), nil
}

func (e BACnetRestartReason) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetRestartReason", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetRestartReason) name() string {
	switch e {
	case BACnetRestartReason_UNKNOWN:
		return "UNKNOWN"
	case BACnetRestartReason_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetRestartReason_COLDSTART:
		return "COLDSTART"
	case BACnetRestartReason_WARMSTART:
		return "WARMSTART"
	case BACnetRestartReason_DETECTED_POWER_LOST:
		return "DETECTED_POWER_LOST"
	case BACnetRestartReason_DETECTED_POWERED_OFF:
		return "DETECTED_POWERED_OFF"
	case BACnetRestartReason_HARDWARE_WATCHDOG:
		return "HARDWARE_WATCHDOG"
	case BACnetRestartReason_SOFTWARE_WATCHDOG:
		return "SOFTWARE_WATCHDOG"
	case BACnetRestartReason_SUSPENDED:
		return "SUSPENDED"
	case BACnetRestartReason_ACTIVATE_CHANGES:
		return "ACTIVATE_CHANGES"
	}
	return ""
}

func (e BACnetRestartReason) String() string {
	return e.name()
}
