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

// BACnetLifeSafetyState is an enum
type BACnetLifeSafetyState uint16

type IBACnetLifeSafetyState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetLifeSafetyState_QUIET                    BACnetLifeSafetyState = 0
	BACnetLifeSafetyState_PRE_ALARM                BACnetLifeSafetyState = 1
	BACnetLifeSafetyState_ALARM                    BACnetLifeSafetyState = 2
	BACnetLifeSafetyState_FAULT                    BACnetLifeSafetyState = 3
	BACnetLifeSafetyState_FAULT_PRE_ALARM          BACnetLifeSafetyState = 4
	BACnetLifeSafetyState_FAULT_ALARM              BACnetLifeSafetyState = 5
	BACnetLifeSafetyState_NOT_READY                BACnetLifeSafetyState = 6
	BACnetLifeSafetyState_ACTIVE                   BACnetLifeSafetyState = 7
	BACnetLifeSafetyState_TAMPER                   BACnetLifeSafetyState = 8
	BACnetLifeSafetyState_TEST_ALARM               BACnetLifeSafetyState = 9
	BACnetLifeSafetyState_TEST_ACTIVE              BACnetLifeSafetyState = 10
	BACnetLifeSafetyState_TEST_FAULT               BACnetLifeSafetyState = 11
	BACnetLifeSafetyState_TEST_FAULT_ALARM         BACnetLifeSafetyState = 12
	BACnetLifeSafetyState_HOLDUP                   BACnetLifeSafetyState = 13
	BACnetLifeSafetyState_DURESS                   BACnetLifeSafetyState = 14
	BACnetLifeSafetyState_TAMPER_ALARM             BACnetLifeSafetyState = 15
	BACnetLifeSafetyState_ABNORMAL                 BACnetLifeSafetyState = 16
	BACnetLifeSafetyState_EMERGENCY_POWER          BACnetLifeSafetyState = 17
	BACnetLifeSafetyState_DELAYED                  BACnetLifeSafetyState = 18
	BACnetLifeSafetyState_BLOCKED                  BACnetLifeSafetyState = 19
	BACnetLifeSafetyState_LOCAL_ALARM              BACnetLifeSafetyState = 20
	BACnetLifeSafetyState_GENERAL_ALARM            BACnetLifeSafetyState = 21
	BACnetLifeSafetyState_SUPERVISORY              BACnetLifeSafetyState = 22
	BACnetLifeSafetyState_TEST_SUPERVISORY         BACnetLifeSafetyState = 23
	BACnetLifeSafetyState_VENDOR_PROPRIETARY_VALUE BACnetLifeSafetyState = 0xFFFF
)

var BACnetLifeSafetyStateValues []BACnetLifeSafetyState

func init() {
	_ = errors.New
	BACnetLifeSafetyStateValues = []BACnetLifeSafetyState{
		BACnetLifeSafetyState_QUIET,
		BACnetLifeSafetyState_PRE_ALARM,
		BACnetLifeSafetyState_ALARM,
		BACnetLifeSafetyState_FAULT,
		BACnetLifeSafetyState_FAULT_PRE_ALARM,
		BACnetLifeSafetyState_FAULT_ALARM,
		BACnetLifeSafetyState_NOT_READY,
		BACnetLifeSafetyState_ACTIVE,
		BACnetLifeSafetyState_TAMPER,
		BACnetLifeSafetyState_TEST_ALARM,
		BACnetLifeSafetyState_TEST_ACTIVE,
		BACnetLifeSafetyState_TEST_FAULT,
		BACnetLifeSafetyState_TEST_FAULT_ALARM,
		BACnetLifeSafetyState_HOLDUP,
		BACnetLifeSafetyState_DURESS,
		BACnetLifeSafetyState_TAMPER_ALARM,
		BACnetLifeSafetyState_ABNORMAL,
		BACnetLifeSafetyState_EMERGENCY_POWER,
		BACnetLifeSafetyState_DELAYED,
		BACnetLifeSafetyState_BLOCKED,
		BACnetLifeSafetyState_LOCAL_ALARM,
		BACnetLifeSafetyState_GENERAL_ALARM,
		BACnetLifeSafetyState_SUPERVISORY,
		BACnetLifeSafetyState_TEST_SUPERVISORY,
		BACnetLifeSafetyState_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLifeSafetyStateByValue(value uint16) (enum BACnetLifeSafetyState, ok bool) {
	switch value {
	case 0:
		return BACnetLifeSafetyState_QUIET, true
	case 0xFFFF:
		return BACnetLifeSafetyState_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetLifeSafetyState_PRE_ALARM, true
	case 10:
		return BACnetLifeSafetyState_TEST_ACTIVE, true
	case 11:
		return BACnetLifeSafetyState_TEST_FAULT, true
	case 12:
		return BACnetLifeSafetyState_TEST_FAULT_ALARM, true
	case 13:
		return BACnetLifeSafetyState_HOLDUP, true
	case 14:
		return BACnetLifeSafetyState_DURESS, true
	case 15:
		return BACnetLifeSafetyState_TAMPER_ALARM, true
	case 16:
		return BACnetLifeSafetyState_ABNORMAL, true
	case 17:
		return BACnetLifeSafetyState_EMERGENCY_POWER, true
	case 18:
		return BACnetLifeSafetyState_DELAYED, true
	case 19:
		return BACnetLifeSafetyState_BLOCKED, true
	case 2:
		return BACnetLifeSafetyState_ALARM, true
	case 20:
		return BACnetLifeSafetyState_LOCAL_ALARM, true
	case 21:
		return BACnetLifeSafetyState_GENERAL_ALARM, true
	case 22:
		return BACnetLifeSafetyState_SUPERVISORY, true
	case 23:
		return BACnetLifeSafetyState_TEST_SUPERVISORY, true
	case 3:
		return BACnetLifeSafetyState_FAULT, true
	case 4:
		return BACnetLifeSafetyState_FAULT_PRE_ALARM, true
	case 5:
		return BACnetLifeSafetyState_FAULT_ALARM, true
	case 6:
		return BACnetLifeSafetyState_NOT_READY, true
	case 7:
		return BACnetLifeSafetyState_ACTIVE, true
	case 8:
		return BACnetLifeSafetyState_TAMPER, true
	case 9:
		return BACnetLifeSafetyState_TEST_ALARM, true
	}
	return 0, false
}

func BACnetLifeSafetyStateByName(value string) (enum BACnetLifeSafetyState, ok bool) {
	switch value {
	case "QUIET":
		return BACnetLifeSafetyState_QUIET, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetLifeSafetyState_VENDOR_PROPRIETARY_VALUE, true
	case "PRE_ALARM":
		return BACnetLifeSafetyState_PRE_ALARM, true
	case "TEST_ACTIVE":
		return BACnetLifeSafetyState_TEST_ACTIVE, true
	case "TEST_FAULT":
		return BACnetLifeSafetyState_TEST_FAULT, true
	case "TEST_FAULT_ALARM":
		return BACnetLifeSafetyState_TEST_FAULT_ALARM, true
	case "HOLDUP":
		return BACnetLifeSafetyState_HOLDUP, true
	case "DURESS":
		return BACnetLifeSafetyState_DURESS, true
	case "TAMPER_ALARM":
		return BACnetLifeSafetyState_TAMPER_ALARM, true
	case "ABNORMAL":
		return BACnetLifeSafetyState_ABNORMAL, true
	case "EMERGENCY_POWER":
		return BACnetLifeSafetyState_EMERGENCY_POWER, true
	case "DELAYED":
		return BACnetLifeSafetyState_DELAYED, true
	case "BLOCKED":
		return BACnetLifeSafetyState_BLOCKED, true
	case "ALARM":
		return BACnetLifeSafetyState_ALARM, true
	case "LOCAL_ALARM":
		return BACnetLifeSafetyState_LOCAL_ALARM, true
	case "GENERAL_ALARM":
		return BACnetLifeSafetyState_GENERAL_ALARM, true
	case "SUPERVISORY":
		return BACnetLifeSafetyState_SUPERVISORY, true
	case "TEST_SUPERVISORY":
		return BACnetLifeSafetyState_TEST_SUPERVISORY, true
	case "FAULT":
		return BACnetLifeSafetyState_FAULT, true
	case "FAULT_PRE_ALARM":
		return BACnetLifeSafetyState_FAULT_PRE_ALARM, true
	case "FAULT_ALARM":
		return BACnetLifeSafetyState_FAULT_ALARM, true
	case "NOT_READY":
		return BACnetLifeSafetyState_NOT_READY, true
	case "ACTIVE":
		return BACnetLifeSafetyState_ACTIVE, true
	case "TAMPER":
		return BACnetLifeSafetyState_TAMPER, true
	case "TEST_ALARM":
		return BACnetLifeSafetyState_TEST_ALARM, true
	}
	return 0, false
}

func BACnetLifeSafetyStateKnows(value uint16) bool {
	for _, typeValue := range BACnetLifeSafetyStateValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLifeSafetyState(structType any) BACnetLifeSafetyState {
	castFunc := func(typ any) BACnetLifeSafetyState {
		if sBACnetLifeSafetyState, ok := typ.(BACnetLifeSafetyState); ok {
			return sBACnetLifeSafetyState
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLifeSafetyState) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetLifeSafetyState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLifeSafetyStateParse(ctx context.Context, theBytes []byte) (BACnetLifeSafetyState, error) {
	return BACnetLifeSafetyStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLifeSafetyStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLifeSafetyState, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("BACnetLifeSafetyState", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLifeSafetyState")
	}
	if enum, ok := BACnetLifeSafetyStateByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetLifeSafetyState")
		return BACnetLifeSafetyState(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLifeSafetyState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLifeSafetyState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("BACnetLifeSafetyState", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLifeSafetyState) PLC4XEnumName() string {
	switch e {
	case BACnetLifeSafetyState_QUIET:
		return "QUIET"
	case BACnetLifeSafetyState_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLifeSafetyState_PRE_ALARM:
		return "PRE_ALARM"
	case BACnetLifeSafetyState_TEST_ACTIVE:
		return "TEST_ACTIVE"
	case BACnetLifeSafetyState_TEST_FAULT:
		return "TEST_FAULT"
	case BACnetLifeSafetyState_TEST_FAULT_ALARM:
		return "TEST_FAULT_ALARM"
	case BACnetLifeSafetyState_HOLDUP:
		return "HOLDUP"
	case BACnetLifeSafetyState_DURESS:
		return "DURESS"
	case BACnetLifeSafetyState_TAMPER_ALARM:
		return "TAMPER_ALARM"
	case BACnetLifeSafetyState_ABNORMAL:
		return "ABNORMAL"
	case BACnetLifeSafetyState_EMERGENCY_POWER:
		return "EMERGENCY_POWER"
	case BACnetLifeSafetyState_DELAYED:
		return "DELAYED"
	case BACnetLifeSafetyState_BLOCKED:
		return "BLOCKED"
	case BACnetLifeSafetyState_ALARM:
		return "ALARM"
	case BACnetLifeSafetyState_LOCAL_ALARM:
		return "LOCAL_ALARM"
	case BACnetLifeSafetyState_GENERAL_ALARM:
		return "GENERAL_ALARM"
	case BACnetLifeSafetyState_SUPERVISORY:
		return "SUPERVISORY"
	case BACnetLifeSafetyState_TEST_SUPERVISORY:
		return "TEST_SUPERVISORY"
	case BACnetLifeSafetyState_FAULT:
		return "FAULT"
	case BACnetLifeSafetyState_FAULT_PRE_ALARM:
		return "FAULT_PRE_ALARM"
	case BACnetLifeSafetyState_FAULT_ALARM:
		return "FAULT_ALARM"
	case BACnetLifeSafetyState_NOT_READY:
		return "NOT_READY"
	case BACnetLifeSafetyState_ACTIVE:
		return "ACTIVE"
	case BACnetLifeSafetyState_TAMPER:
		return "TAMPER"
	case BACnetLifeSafetyState_TEST_ALARM:
		return "TEST_ALARM"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetLifeSafetyState) String() string {
	return e.PLC4XEnumName()
}
