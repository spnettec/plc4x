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

// BACnetDoorValue is an enum
type BACnetDoorValue uint8

type IBACnetDoorValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	BACnetDoorValue_LOCK BACnetDoorValue = 0
	BACnetDoorValue_UNLOCK BACnetDoorValue = 1
	BACnetDoorValue_PULSE_UNLOCK BACnetDoorValue = 2
	BACnetDoorValue_EXTENDED_PULSE_UNLOCK BACnetDoorValue = 3
)

var BACnetDoorValueValues []BACnetDoorValue

func init() {
	_ = errors.New
	BACnetDoorValueValues = []BACnetDoorValue {
		BACnetDoorValue_LOCK,
		BACnetDoorValue_UNLOCK,
		BACnetDoorValue_PULSE_UNLOCK,
		BACnetDoorValue_EXTENDED_PULSE_UNLOCK,
	}
}

func BACnetDoorValueByValue(value uint8) (enum BACnetDoorValue, ok bool) {
	switch value {
		case 0:
			return BACnetDoorValue_LOCK, true
		case 1:
			return BACnetDoorValue_UNLOCK, true
		case 2:
			return BACnetDoorValue_PULSE_UNLOCK, true
		case 3:
			return BACnetDoorValue_EXTENDED_PULSE_UNLOCK, true
	}
	return 0, false
}

func BACnetDoorValueByName(value string) (enum BACnetDoorValue, ok bool) {
	switch value {
	case "LOCK":
		return BACnetDoorValue_LOCK, true
	case "UNLOCK":
		return BACnetDoorValue_UNLOCK, true
	case "PULSE_UNLOCK":
		return BACnetDoorValue_PULSE_UNLOCK, true
	case "EXTENDED_PULSE_UNLOCK":
		return BACnetDoorValue_EXTENDED_PULSE_UNLOCK, true
	}
	return 0, false
}

func BACnetDoorValueKnows(value uint8)  bool {
	for _, typeValue := range BACnetDoorValueValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetDoorValue(structType any) BACnetDoorValue {
	castFunc := func(typ any) BACnetDoorValue {
		if sBACnetDoorValue, ok := typ.(BACnetDoorValue); ok {
			return sBACnetDoorValue
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetDoorValue) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetDoorValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDoorValueParse(ctx context.Context, theBytes []byte) (BACnetDoorValue, error) {
	return BACnetDoorValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDoorValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDoorValue, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetDoorValue", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetDoorValue")
	}
	if enum, ok := BACnetDoorValueByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetDoorValue(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetDoorValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetDoorValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetDoorValue", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetDoorValue) PLC4XEnumName() string {
	switch e {
	case BACnetDoorValue_LOCK:
		return "LOCK"
	case BACnetDoorValue_UNLOCK:
		return "UNLOCK"
	case BACnetDoorValue_PULSE_UNLOCK:
		return "PULSE_UNLOCK"
	case BACnetDoorValue_EXTENDED_PULSE_UNLOCK:
		return "EXTENDED_PULSE_UNLOCK"
	}
	return ""
}

func (e BACnetDoorValue) String() string {
	return e.PLC4XEnumName()
}

