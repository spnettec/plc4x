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

// BACnetLightingTransition is an enum
type BACnetLightingTransition uint8

type IBACnetLightingTransition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	BACnetLightingTransition_NONE BACnetLightingTransition = 0
	BACnetLightingTransition_FADE BACnetLightingTransition = 1
	BACnetLightingTransition_RAMP BACnetLightingTransition = 2
	BACnetLightingTransition_VENDOR_PROPRIETARY_VALUE BACnetLightingTransition = 0XFF
)

var BACnetLightingTransitionValues []BACnetLightingTransition

func init() {
	_ = errors.New
	BACnetLightingTransitionValues = []BACnetLightingTransition {
		BACnetLightingTransition_NONE,
		BACnetLightingTransition_FADE,
		BACnetLightingTransition_RAMP,
		BACnetLightingTransition_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLightingTransitionByValue(value uint8) (enum BACnetLightingTransition, ok bool) {
	switch value {
		case 0:
			return BACnetLightingTransition_NONE, true
		case 0XFF:
			return BACnetLightingTransition_VENDOR_PROPRIETARY_VALUE, true
		case 1:
			return BACnetLightingTransition_FADE, true
		case 2:
			return BACnetLightingTransition_RAMP, true
	}
	return 0, false
}

func BACnetLightingTransitionByName(value string) (enum BACnetLightingTransition, ok bool) {
	switch value {
	case "NONE":
		return BACnetLightingTransition_NONE, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetLightingTransition_VENDOR_PROPRIETARY_VALUE, true
	case "FADE":
		return BACnetLightingTransition_FADE, true
	case "RAMP":
		return BACnetLightingTransition_RAMP, true
	}
	return 0, false
}

func BACnetLightingTransitionKnows(value uint8)  bool {
	for _, typeValue := range BACnetLightingTransitionValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetLightingTransition(structType any) BACnetLightingTransition {
	castFunc := func(typ any) BACnetLightingTransition {
		if sBACnetLightingTransition, ok := typ.(BACnetLightingTransition); ok {
			return sBACnetLightingTransition
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLightingTransition) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetLightingTransition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLightingTransitionParse(ctx context.Context, theBytes []byte) (BACnetLightingTransition, error) {
	return BACnetLightingTransitionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLightingTransitionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingTransition, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetLightingTransition", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLightingTransition")
	}
	if enum, ok := BACnetLightingTransitionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetLightingTransition")
		return BACnetLightingTransition(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLightingTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLightingTransition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetLightingTransition", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLightingTransition) PLC4XEnumName() string {
	switch e {
	case BACnetLightingTransition_NONE:
		return "NONE"
	case BACnetLightingTransition_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLightingTransition_FADE:
		return "FADE"
	case BACnetLightingTransition_RAMP:
		return "RAMP"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetLightingTransition) String() string {
	return e.PLC4XEnumName()
}

