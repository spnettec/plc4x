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

// BACnetVTClass is an enum
type BACnetVTClass uint16

type IBACnetVTClass interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetVTClass_DEFAULT_TERMINAL         BACnetVTClass = 0
	BACnetVTClass_ANSI_X3_64               BACnetVTClass = 1
	BACnetVTClass_DEC_VT52                 BACnetVTClass = 2
	BACnetVTClass_DEC_VT100                BACnetVTClass = 3
	BACnetVTClass_DEC_VT220                BACnetVTClass = 4
	BACnetVTClass_HP_700_94                BACnetVTClass = 5
	BACnetVTClass_IBM_3130                 BACnetVTClass = 6
	BACnetVTClass_VENDOR_PROPRIETARY_VALUE BACnetVTClass = 0xFFFF
)

var BACnetVTClassValues []BACnetVTClass

func init() {
	_ = errors.New
	BACnetVTClassValues = []BACnetVTClass{
		BACnetVTClass_DEFAULT_TERMINAL,
		BACnetVTClass_ANSI_X3_64,
		BACnetVTClass_DEC_VT52,
		BACnetVTClass_DEC_VT100,
		BACnetVTClass_DEC_VT220,
		BACnetVTClass_HP_700_94,
		BACnetVTClass_IBM_3130,
		BACnetVTClass_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetVTClassByValue(value uint16) (enum BACnetVTClass, ok bool) {
	switch value {
	case 0:
		return BACnetVTClass_DEFAULT_TERMINAL, true
	case 0xFFFF:
		return BACnetVTClass_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetVTClass_ANSI_X3_64, true
	case 2:
		return BACnetVTClass_DEC_VT52, true
	case 3:
		return BACnetVTClass_DEC_VT100, true
	case 4:
		return BACnetVTClass_DEC_VT220, true
	case 5:
		return BACnetVTClass_HP_700_94, true
	case 6:
		return BACnetVTClass_IBM_3130, true
	}
	return 0, false
}

func BACnetVTClassByName(value string) (enum BACnetVTClass, ok bool) {
	switch value {
	case "DEFAULT_TERMINAL":
		return BACnetVTClass_DEFAULT_TERMINAL, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetVTClass_VENDOR_PROPRIETARY_VALUE, true
	case "ANSI_X3_64":
		return BACnetVTClass_ANSI_X3_64, true
	case "DEC_VT52":
		return BACnetVTClass_DEC_VT52, true
	case "DEC_VT100":
		return BACnetVTClass_DEC_VT100, true
	case "DEC_VT220":
		return BACnetVTClass_DEC_VT220, true
	case "HP_700_94":
		return BACnetVTClass_HP_700_94, true
	case "IBM_3130":
		return BACnetVTClass_IBM_3130, true
	}
	return 0, false
}

func BACnetVTClassKnows(value uint16) bool {
	for _, typeValue := range BACnetVTClassValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetVTClass(structType any) BACnetVTClass {
	castFunc := func(typ any) BACnetVTClass {
		if sBACnetVTClass, ok := typ.(BACnetVTClass); ok {
			return sBACnetVTClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetVTClass) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetVTClass) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetVTClassParse(ctx context.Context, theBytes []byte) (BACnetVTClass, error) {
	return BACnetVTClassParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetVTClassParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetVTClass, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("BACnetVTClass", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetVTClass")
	}
	if enum, ok := BACnetVTClassByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetVTClass(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetVTClass) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetVTClass) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("BACnetVTClass", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetVTClass) PLC4XEnumName() string {
	switch e {
	case BACnetVTClass_DEFAULT_TERMINAL:
		return "DEFAULT_TERMINAL"
	case BACnetVTClass_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetVTClass_ANSI_X3_64:
		return "ANSI_X3_64"
	case BACnetVTClass_DEC_VT52:
		return "DEC_VT52"
	case BACnetVTClass_DEC_VT100:
		return "DEC_VT100"
	case BACnetVTClass_DEC_VT220:
		return "DEC_VT220"
	case BACnetVTClass_HP_700_94:
		return "HP_700_94"
	case BACnetVTClass_IBM_3130:
		return "IBM_3130"
	}
	return ""
}

func (e BACnetVTClass) String() string {
	return e.PLC4XEnumName()
}
