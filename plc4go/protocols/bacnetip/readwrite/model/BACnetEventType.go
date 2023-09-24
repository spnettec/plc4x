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

// BACnetEventType is an enum
type BACnetEventType uint16

type IBACnetEventType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	BACnetEventType_CHANGE_OF_BITSTRING BACnetEventType = 0
	BACnetEventType_CHANGE_OF_STATE BACnetEventType = 1
	BACnetEventType_CHANGE_OF_VALUE BACnetEventType = 2
	BACnetEventType_COMMAND_FAILURE BACnetEventType = 3
	BACnetEventType_FLOATING_LIMIT BACnetEventType = 4
	BACnetEventType_OUT_OF_RANGE BACnetEventType = 5
	BACnetEventType_CHANGE_OF_LIFE_SAFETY BACnetEventType = 8
	BACnetEventType_EXTENDED BACnetEventType = 9
	BACnetEventType_BUFFER_READY BACnetEventType = 10
	BACnetEventType_UNSIGNED_RANGE BACnetEventType = 11
	BACnetEventType_ACCESS_EVENT BACnetEventType = 13
	BACnetEventType_DOUBLE_OUT_OF_RANGE BACnetEventType = 14
	BACnetEventType_SIGNED_OUT_OF_RANGE BACnetEventType = 15
	BACnetEventType_UNSIGNED_OUT_OF_RANGE BACnetEventType = 16
	BACnetEventType_CHANGE_OF_CHARACTERSTRING BACnetEventType = 17
	BACnetEventType_CHANGE_OF_STATUS_FLAGS BACnetEventType = 18
	BACnetEventType_CHANGE_OF_RELIABILITY BACnetEventType = 19
	BACnetEventType_NONE BACnetEventType = 20
	BACnetEventType_CHANGE_OF_DISCRETE_VALUE BACnetEventType = 21
	BACnetEventType_CHANGE_OF_TIMER BACnetEventType = 22
	BACnetEventType_VENDOR_PROPRIETARY_VALUE BACnetEventType = 0xFFFF
)

var BACnetEventTypeValues []BACnetEventType

func init() {
	_ = errors.New
	BACnetEventTypeValues = []BACnetEventType {
		BACnetEventType_CHANGE_OF_BITSTRING,
		BACnetEventType_CHANGE_OF_STATE,
		BACnetEventType_CHANGE_OF_VALUE,
		BACnetEventType_COMMAND_FAILURE,
		BACnetEventType_FLOATING_LIMIT,
		BACnetEventType_OUT_OF_RANGE,
		BACnetEventType_CHANGE_OF_LIFE_SAFETY,
		BACnetEventType_EXTENDED,
		BACnetEventType_BUFFER_READY,
		BACnetEventType_UNSIGNED_RANGE,
		BACnetEventType_ACCESS_EVENT,
		BACnetEventType_DOUBLE_OUT_OF_RANGE,
		BACnetEventType_SIGNED_OUT_OF_RANGE,
		BACnetEventType_UNSIGNED_OUT_OF_RANGE,
		BACnetEventType_CHANGE_OF_CHARACTERSTRING,
		BACnetEventType_CHANGE_OF_STATUS_FLAGS,
		BACnetEventType_CHANGE_OF_RELIABILITY,
		BACnetEventType_NONE,
		BACnetEventType_CHANGE_OF_DISCRETE_VALUE,
		BACnetEventType_CHANGE_OF_TIMER,
		BACnetEventType_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetEventTypeByValue(value uint16) (enum BACnetEventType, ok bool) {
	switch value {
		case 0:
			return BACnetEventType_CHANGE_OF_BITSTRING, true
		case 0xFFFF:
			return BACnetEventType_VENDOR_PROPRIETARY_VALUE, true
		case 1:
			return BACnetEventType_CHANGE_OF_STATE, true
		case 10:
			return BACnetEventType_BUFFER_READY, true
		case 11:
			return BACnetEventType_UNSIGNED_RANGE, true
		case 13:
			return BACnetEventType_ACCESS_EVENT, true
		case 14:
			return BACnetEventType_DOUBLE_OUT_OF_RANGE, true
		case 15:
			return BACnetEventType_SIGNED_OUT_OF_RANGE, true
		case 16:
			return BACnetEventType_UNSIGNED_OUT_OF_RANGE, true
		case 17:
			return BACnetEventType_CHANGE_OF_CHARACTERSTRING, true
		case 18:
			return BACnetEventType_CHANGE_OF_STATUS_FLAGS, true
		case 19:
			return BACnetEventType_CHANGE_OF_RELIABILITY, true
		case 2:
			return BACnetEventType_CHANGE_OF_VALUE, true
		case 20:
			return BACnetEventType_NONE, true
		case 21:
			return BACnetEventType_CHANGE_OF_DISCRETE_VALUE, true
		case 22:
			return BACnetEventType_CHANGE_OF_TIMER, true
		case 3:
			return BACnetEventType_COMMAND_FAILURE, true
		case 4:
			return BACnetEventType_FLOATING_LIMIT, true
		case 5:
			return BACnetEventType_OUT_OF_RANGE, true
		case 8:
			return BACnetEventType_CHANGE_OF_LIFE_SAFETY, true
		case 9:
			return BACnetEventType_EXTENDED, true
	}
	return 0, false
}

func BACnetEventTypeByName(value string) (enum BACnetEventType, ok bool) {
	switch value {
	case "CHANGE_OF_BITSTRING":
		return BACnetEventType_CHANGE_OF_BITSTRING, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetEventType_VENDOR_PROPRIETARY_VALUE, true
	case "CHANGE_OF_STATE":
		return BACnetEventType_CHANGE_OF_STATE, true
	case "BUFFER_READY":
		return BACnetEventType_BUFFER_READY, true
	case "UNSIGNED_RANGE":
		return BACnetEventType_UNSIGNED_RANGE, true
	case "ACCESS_EVENT":
		return BACnetEventType_ACCESS_EVENT, true
	case "DOUBLE_OUT_OF_RANGE":
		return BACnetEventType_DOUBLE_OUT_OF_RANGE, true
	case "SIGNED_OUT_OF_RANGE":
		return BACnetEventType_SIGNED_OUT_OF_RANGE, true
	case "UNSIGNED_OUT_OF_RANGE":
		return BACnetEventType_UNSIGNED_OUT_OF_RANGE, true
	case "CHANGE_OF_CHARACTERSTRING":
		return BACnetEventType_CHANGE_OF_CHARACTERSTRING, true
	case "CHANGE_OF_STATUS_FLAGS":
		return BACnetEventType_CHANGE_OF_STATUS_FLAGS, true
	case "CHANGE_OF_RELIABILITY":
		return BACnetEventType_CHANGE_OF_RELIABILITY, true
	case "CHANGE_OF_VALUE":
		return BACnetEventType_CHANGE_OF_VALUE, true
	case "NONE":
		return BACnetEventType_NONE, true
	case "CHANGE_OF_DISCRETE_VALUE":
		return BACnetEventType_CHANGE_OF_DISCRETE_VALUE, true
	case "CHANGE_OF_TIMER":
		return BACnetEventType_CHANGE_OF_TIMER, true
	case "COMMAND_FAILURE":
		return BACnetEventType_COMMAND_FAILURE, true
	case "FLOATING_LIMIT":
		return BACnetEventType_FLOATING_LIMIT, true
	case "OUT_OF_RANGE":
		return BACnetEventType_OUT_OF_RANGE, true
	case "CHANGE_OF_LIFE_SAFETY":
		return BACnetEventType_CHANGE_OF_LIFE_SAFETY, true
	case "EXTENDED":
		return BACnetEventType_EXTENDED, true
	}
	return 0, false
}

func BACnetEventTypeKnows(value uint16)  bool {
	for _, typeValue := range BACnetEventTypeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetEventType(structType any) BACnetEventType {
	castFunc := func(typ any) BACnetEventType {
		if sBACnetEventType, ok := typ.(BACnetEventType); ok {
			return sBACnetEventType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetEventType) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetEventType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventTypeParse(ctx context.Context, theBytes []byte) (BACnetEventType, error) {
	return BACnetEventTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("BACnetEventType", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetEventType")
	}
	if enum, ok := BACnetEventTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetEventType")
		return BACnetEventType(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetEventType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetEventType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("BACnetEventType", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetEventType) PLC4XEnumName() string {
	switch e {
	case BACnetEventType_CHANGE_OF_BITSTRING:
		return "CHANGE_OF_BITSTRING"
	case BACnetEventType_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetEventType_CHANGE_OF_STATE:
		return "CHANGE_OF_STATE"
	case BACnetEventType_BUFFER_READY:
		return "BUFFER_READY"
	case BACnetEventType_UNSIGNED_RANGE:
		return "UNSIGNED_RANGE"
	case BACnetEventType_ACCESS_EVENT:
		return "ACCESS_EVENT"
	case BACnetEventType_DOUBLE_OUT_OF_RANGE:
		return "DOUBLE_OUT_OF_RANGE"
	case BACnetEventType_SIGNED_OUT_OF_RANGE:
		return "SIGNED_OUT_OF_RANGE"
	case BACnetEventType_UNSIGNED_OUT_OF_RANGE:
		return "UNSIGNED_OUT_OF_RANGE"
	case BACnetEventType_CHANGE_OF_CHARACTERSTRING:
		return "CHANGE_OF_CHARACTERSTRING"
	case BACnetEventType_CHANGE_OF_STATUS_FLAGS:
		return "CHANGE_OF_STATUS_FLAGS"
	case BACnetEventType_CHANGE_OF_RELIABILITY:
		return "CHANGE_OF_RELIABILITY"
	case BACnetEventType_CHANGE_OF_VALUE:
		return "CHANGE_OF_VALUE"
	case BACnetEventType_NONE:
		return "NONE"
	case BACnetEventType_CHANGE_OF_DISCRETE_VALUE:
		return "CHANGE_OF_DISCRETE_VALUE"
	case BACnetEventType_CHANGE_OF_TIMER:
		return "CHANGE_OF_TIMER"
	case BACnetEventType_COMMAND_FAILURE:
		return "COMMAND_FAILURE"
	case BACnetEventType_FLOATING_LIMIT:
		return "FLOATING_LIMIT"
	case BACnetEventType_OUT_OF_RANGE:
		return "OUT_OF_RANGE"
	case BACnetEventType_CHANGE_OF_LIFE_SAFETY:
		return "CHANGE_OF_LIFE_SAFETY"
	case BACnetEventType_EXTENDED:
		return "EXTENDED"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetEventType) String() string {
	return e.PLC4XEnumName()
}

