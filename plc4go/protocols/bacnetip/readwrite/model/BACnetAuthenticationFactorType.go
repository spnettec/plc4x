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

// BACnetAuthenticationFactorType is an enum
type BACnetAuthenticationFactorType uint8

type IBACnetAuthenticationFactorType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	BACnetAuthenticationFactorType_UNDEFINED BACnetAuthenticationFactorType = 0
	BACnetAuthenticationFactorType_ERROR BACnetAuthenticationFactorType = 1
	BACnetAuthenticationFactorType_CUSTOM BACnetAuthenticationFactorType = 2
	BACnetAuthenticationFactorType_SIMPLE_NUMBER16 BACnetAuthenticationFactorType = 3
	BACnetAuthenticationFactorType_SIMPLE_NUMBER32 BACnetAuthenticationFactorType = 4
	BACnetAuthenticationFactorType_SIMPLE_NUMBER56 BACnetAuthenticationFactorType = 5
	BACnetAuthenticationFactorType_SIMPLE_ALPHA_NUMERIC BACnetAuthenticationFactorType = 6
	BACnetAuthenticationFactorType_ABA_TRACK2 BACnetAuthenticationFactorType = 7
	BACnetAuthenticationFactorType_WIEGAND26 BACnetAuthenticationFactorType = 8
	BACnetAuthenticationFactorType_WIEGAND37 BACnetAuthenticationFactorType = 9
	BACnetAuthenticationFactorType_WIEGAND37_FACILITY BACnetAuthenticationFactorType = 10
	BACnetAuthenticationFactorType_FACILITY16_CARD32 BACnetAuthenticationFactorType = 11
	BACnetAuthenticationFactorType_FACILITY32_CARD32 BACnetAuthenticationFactorType = 12
	BACnetAuthenticationFactorType_FASC_N BACnetAuthenticationFactorType = 13
	BACnetAuthenticationFactorType_FASC_N_BCD BACnetAuthenticationFactorType = 14
	BACnetAuthenticationFactorType_FASC_N_LARGE BACnetAuthenticationFactorType = 15
	BACnetAuthenticationFactorType_FASC_N_LARGE_BCD BACnetAuthenticationFactorType = 16
	BACnetAuthenticationFactorType_GSA75 BACnetAuthenticationFactorType = 17
	BACnetAuthenticationFactorType_CHUID BACnetAuthenticationFactorType = 18
	BACnetAuthenticationFactorType_CHUID_FULL BACnetAuthenticationFactorType = 19
	BACnetAuthenticationFactorType_GUID BACnetAuthenticationFactorType = 20
	BACnetAuthenticationFactorType_CBEFF_A BACnetAuthenticationFactorType = 21
	BACnetAuthenticationFactorType_CBEFF_B BACnetAuthenticationFactorType = 22
	BACnetAuthenticationFactorType_CBEFF_C BACnetAuthenticationFactorType = 23
	BACnetAuthenticationFactorType_USER_PASSWORD BACnetAuthenticationFactorType = 24
)

var BACnetAuthenticationFactorTypeValues []BACnetAuthenticationFactorType

func init() {
	_ = errors.New
	BACnetAuthenticationFactorTypeValues = []BACnetAuthenticationFactorType {
		BACnetAuthenticationFactorType_UNDEFINED,
		BACnetAuthenticationFactorType_ERROR,
		BACnetAuthenticationFactorType_CUSTOM,
		BACnetAuthenticationFactorType_SIMPLE_NUMBER16,
		BACnetAuthenticationFactorType_SIMPLE_NUMBER32,
		BACnetAuthenticationFactorType_SIMPLE_NUMBER56,
		BACnetAuthenticationFactorType_SIMPLE_ALPHA_NUMERIC,
		BACnetAuthenticationFactorType_ABA_TRACK2,
		BACnetAuthenticationFactorType_WIEGAND26,
		BACnetAuthenticationFactorType_WIEGAND37,
		BACnetAuthenticationFactorType_WIEGAND37_FACILITY,
		BACnetAuthenticationFactorType_FACILITY16_CARD32,
		BACnetAuthenticationFactorType_FACILITY32_CARD32,
		BACnetAuthenticationFactorType_FASC_N,
		BACnetAuthenticationFactorType_FASC_N_BCD,
		BACnetAuthenticationFactorType_FASC_N_LARGE,
		BACnetAuthenticationFactorType_FASC_N_LARGE_BCD,
		BACnetAuthenticationFactorType_GSA75,
		BACnetAuthenticationFactorType_CHUID,
		BACnetAuthenticationFactorType_CHUID_FULL,
		BACnetAuthenticationFactorType_GUID,
		BACnetAuthenticationFactorType_CBEFF_A,
		BACnetAuthenticationFactorType_CBEFF_B,
		BACnetAuthenticationFactorType_CBEFF_C,
		BACnetAuthenticationFactorType_USER_PASSWORD,
	}
}

func BACnetAuthenticationFactorTypeByValue(value uint8) (enum BACnetAuthenticationFactorType, ok bool) {
	switch value {
		case 0:
			return BACnetAuthenticationFactorType_UNDEFINED, true
		case 1:
			return BACnetAuthenticationFactorType_ERROR, true
		case 10:
			return BACnetAuthenticationFactorType_WIEGAND37_FACILITY, true
		case 11:
			return BACnetAuthenticationFactorType_FACILITY16_CARD32, true
		case 12:
			return BACnetAuthenticationFactorType_FACILITY32_CARD32, true
		case 13:
			return BACnetAuthenticationFactorType_FASC_N, true
		case 14:
			return BACnetAuthenticationFactorType_FASC_N_BCD, true
		case 15:
			return BACnetAuthenticationFactorType_FASC_N_LARGE, true
		case 16:
			return BACnetAuthenticationFactorType_FASC_N_LARGE_BCD, true
		case 17:
			return BACnetAuthenticationFactorType_GSA75, true
		case 18:
			return BACnetAuthenticationFactorType_CHUID, true
		case 19:
			return BACnetAuthenticationFactorType_CHUID_FULL, true
		case 2:
			return BACnetAuthenticationFactorType_CUSTOM, true
		case 20:
			return BACnetAuthenticationFactorType_GUID, true
		case 21:
			return BACnetAuthenticationFactorType_CBEFF_A, true
		case 22:
			return BACnetAuthenticationFactorType_CBEFF_B, true
		case 23:
			return BACnetAuthenticationFactorType_CBEFF_C, true
		case 24:
			return BACnetAuthenticationFactorType_USER_PASSWORD, true
		case 3:
			return BACnetAuthenticationFactorType_SIMPLE_NUMBER16, true
		case 4:
			return BACnetAuthenticationFactorType_SIMPLE_NUMBER32, true
		case 5:
			return BACnetAuthenticationFactorType_SIMPLE_NUMBER56, true
		case 6:
			return BACnetAuthenticationFactorType_SIMPLE_ALPHA_NUMERIC, true
		case 7:
			return BACnetAuthenticationFactorType_ABA_TRACK2, true
		case 8:
			return BACnetAuthenticationFactorType_WIEGAND26, true
		case 9:
			return BACnetAuthenticationFactorType_WIEGAND37, true
	}
	return 0, false
}

func BACnetAuthenticationFactorTypeByName(value string) (enum BACnetAuthenticationFactorType, ok bool) {
	switch value {
	case "UNDEFINED":
		return BACnetAuthenticationFactorType_UNDEFINED, true
	case "ERROR":
		return BACnetAuthenticationFactorType_ERROR, true
	case "WIEGAND37_FACILITY":
		return BACnetAuthenticationFactorType_WIEGAND37_FACILITY, true
	case "FACILITY16_CARD32":
		return BACnetAuthenticationFactorType_FACILITY16_CARD32, true
	case "FACILITY32_CARD32":
		return BACnetAuthenticationFactorType_FACILITY32_CARD32, true
	case "FASC_N":
		return BACnetAuthenticationFactorType_FASC_N, true
	case "FASC_N_BCD":
		return BACnetAuthenticationFactorType_FASC_N_BCD, true
	case "FASC_N_LARGE":
		return BACnetAuthenticationFactorType_FASC_N_LARGE, true
	case "FASC_N_LARGE_BCD":
		return BACnetAuthenticationFactorType_FASC_N_LARGE_BCD, true
	case "GSA75":
		return BACnetAuthenticationFactorType_GSA75, true
	case "CHUID":
		return BACnetAuthenticationFactorType_CHUID, true
	case "CHUID_FULL":
		return BACnetAuthenticationFactorType_CHUID_FULL, true
	case "CUSTOM":
		return BACnetAuthenticationFactorType_CUSTOM, true
	case "GUID":
		return BACnetAuthenticationFactorType_GUID, true
	case "CBEFF_A":
		return BACnetAuthenticationFactorType_CBEFF_A, true
	case "CBEFF_B":
		return BACnetAuthenticationFactorType_CBEFF_B, true
	case "CBEFF_C":
		return BACnetAuthenticationFactorType_CBEFF_C, true
	case "USER_PASSWORD":
		return BACnetAuthenticationFactorType_USER_PASSWORD, true
	case "SIMPLE_NUMBER16":
		return BACnetAuthenticationFactorType_SIMPLE_NUMBER16, true
	case "SIMPLE_NUMBER32":
		return BACnetAuthenticationFactorType_SIMPLE_NUMBER32, true
	case "SIMPLE_NUMBER56":
		return BACnetAuthenticationFactorType_SIMPLE_NUMBER56, true
	case "SIMPLE_ALPHA_NUMERIC":
		return BACnetAuthenticationFactorType_SIMPLE_ALPHA_NUMERIC, true
	case "ABA_TRACK2":
		return BACnetAuthenticationFactorType_ABA_TRACK2, true
	case "WIEGAND26":
		return BACnetAuthenticationFactorType_WIEGAND26, true
	case "WIEGAND37":
		return BACnetAuthenticationFactorType_WIEGAND37, true
	}
	return 0, false
}

func BACnetAuthenticationFactorTypeKnows(value uint8)  bool {
	for _, typeValue := range BACnetAuthenticationFactorTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetAuthenticationFactorType(structType any) BACnetAuthenticationFactorType {
	castFunc := func(typ any) BACnetAuthenticationFactorType {
		if sBACnetAuthenticationFactorType, ok := typ.(BACnetAuthenticationFactorType); ok {
			return sBACnetAuthenticationFactorType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAuthenticationFactorType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetAuthenticationFactorType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationFactorTypeParse(ctx context.Context, theBytes []byte) (BACnetAuthenticationFactorType, error) {
	return BACnetAuthenticationFactorTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAuthenticationFactorTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactorType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetAuthenticationFactorType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAuthenticationFactorType")
	}
	if enum, ok := BACnetAuthenticationFactorTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetAuthenticationFactorType(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAuthenticationFactorType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAuthenticationFactorType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetAuthenticationFactorType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAuthenticationFactorType) PLC4XEnumName() string {
	switch e {
	case BACnetAuthenticationFactorType_UNDEFINED:
		return "UNDEFINED"
	case BACnetAuthenticationFactorType_ERROR:
		return "ERROR"
	case BACnetAuthenticationFactorType_WIEGAND37_FACILITY:
		return "WIEGAND37_FACILITY"
	case BACnetAuthenticationFactorType_FACILITY16_CARD32:
		return "FACILITY16_CARD32"
	case BACnetAuthenticationFactorType_FACILITY32_CARD32:
		return "FACILITY32_CARD32"
	case BACnetAuthenticationFactorType_FASC_N:
		return "FASC_N"
	case BACnetAuthenticationFactorType_FASC_N_BCD:
		return "FASC_N_BCD"
	case BACnetAuthenticationFactorType_FASC_N_LARGE:
		return "FASC_N_LARGE"
	case BACnetAuthenticationFactorType_FASC_N_LARGE_BCD:
		return "FASC_N_LARGE_BCD"
	case BACnetAuthenticationFactorType_GSA75:
		return "GSA75"
	case BACnetAuthenticationFactorType_CHUID:
		return "CHUID"
	case BACnetAuthenticationFactorType_CHUID_FULL:
		return "CHUID_FULL"
	case BACnetAuthenticationFactorType_CUSTOM:
		return "CUSTOM"
	case BACnetAuthenticationFactorType_GUID:
		return "GUID"
	case BACnetAuthenticationFactorType_CBEFF_A:
		return "CBEFF_A"
	case BACnetAuthenticationFactorType_CBEFF_B:
		return "CBEFF_B"
	case BACnetAuthenticationFactorType_CBEFF_C:
		return "CBEFF_C"
	case BACnetAuthenticationFactorType_USER_PASSWORD:
		return "USER_PASSWORD"
	case BACnetAuthenticationFactorType_SIMPLE_NUMBER16:
		return "SIMPLE_NUMBER16"
	case BACnetAuthenticationFactorType_SIMPLE_NUMBER32:
		return "SIMPLE_NUMBER32"
	case BACnetAuthenticationFactorType_SIMPLE_NUMBER56:
		return "SIMPLE_NUMBER56"
	case BACnetAuthenticationFactorType_SIMPLE_ALPHA_NUMERIC:
		return "SIMPLE_ALPHA_NUMERIC"
	case BACnetAuthenticationFactorType_ABA_TRACK2:
		return "ABA_TRACK2"
	case BACnetAuthenticationFactorType_WIEGAND26:
		return "WIEGAND26"
	case BACnetAuthenticationFactorType_WIEGAND37:
		return "WIEGAND37"
	}
	return ""
}

func (e BACnetAuthenticationFactorType) String() string {
	return e.PLC4XEnumName()
}

