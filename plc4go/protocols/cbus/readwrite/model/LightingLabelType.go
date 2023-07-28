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

// LightingLabelType is an enum
type LightingLabelType uint8

type ILightingLabelType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	LightingLabelType_TEXT_LABEL LightingLabelType = 0
	LightingLabelType_PREDEFINED_ICON LightingLabelType = 1
	LightingLabelType_LOAD_DYNAMIC_ICON LightingLabelType = 2
	LightingLabelType_SET_PREFERRED_LANGUAGE LightingLabelType = 3
)

var LightingLabelTypeValues []LightingLabelType

func init() {
	_ = errors.New
	LightingLabelTypeValues = []LightingLabelType {
		LightingLabelType_TEXT_LABEL,
		LightingLabelType_PREDEFINED_ICON,
		LightingLabelType_LOAD_DYNAMIC_ICON,
		LightingLabelType_SET_PREFERRED_LANGUAGE,
	}
}

func LightingLabelTypeByValue(value uint8) (enum LightingLabelType, ok bool) {
	switch value {
		case 0:
			return LightingLabelType_TEXT_LABEL, true
		case 1:
			return LightingLabelType_PREDEFINED_ICON, true
		case 2:
			return LightingLabelType_LOAD_DYNAMIC_ICON, true
		case 3:
			return LightingLabelType_SET_PREFERRED_LANGUAGE, true
	}
	return 0, false
}

func LightingLabelTypeByName(value string) (enum LightingLabelType, ok bool) {
	switch value {
	case "TEXT_LABEL":
		return LightingLabelType_TEXT_LABEL, true
	case "PREDEFINED_ICON":
		return LightingLabelType_PREDEFINED_ICON, true
	case "LOAD_DYNAMIC_ICON":
		return LightingLabelType_LOAD_DYNAMIC_ICON, true
	case "SET_PREFERRED_LANGUAGE":
		return LightingLabelType_SET_PREFERRED_LANGUAGE, true
	}
	return 0, false
}

func LightingLabelTypeKnows(value uint8)  bool {
	for _, typeValue := range LightingLabelTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastLightingLabelType(structType any) LightingLabelType {
	castFunc := func(typ any) LightingLabelType {
		if sLightingLabelType, ok := typ.(LightingLabelType); ok {
			return sLightingLabelType
		}
		return 0
	}
	return castFunc(structType)
}

func (m LightingLabelType) GetLengthInBits(ctx context.Context) uint16 {
	return 2
}

func (m LightingLabelType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LightingLabelTypeParse(ctx context.Context, theBytes []byte) (LightingLabelType, error) {
	return LightingLabelTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LightingLabelTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LightingLabelType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("LightingLabelType", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading LightingLabelType")
	}
	if enum, ok := LightingLabelTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return LightingLabelType(val), nil
	} else {
		return enum, nil
	}
}

func (e LightingLabelType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e LightingLabelType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("LightingLabelType", 2, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e LightingLabelType) PLC4XEnumName() string {
	switch e {
	case LightingLabelType_TEXT_LABEL:
		return "TEXT_LABEL"
	case LightingLabelType_PREDEFINED_ICON:
		return "PREDEFINED_ICON"
	case LightingLabelType_LOAD_DYNAMIC_ICON:
		return "LOAD_DYNAMIC_ICON"
	case LightingLabelType_SET_PREFERRED_LANGUAGE:
		return "SET_PREFERRED_LANGUAGE"
	}
	return ""
}

func (e LightingLabelType) String() string {
	return e.PLC4XEnumName()
}

