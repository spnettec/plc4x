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

// AccessLevelType is an enum
type AccessLevelType uint8

type IAccessLevelType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	AccessLevelType_accessLevelTypeNone AccessLevelType = 0
	AccessLevelType_accessLevelTypeCurrentRead AccessLevelType = 1
	AccessLevelType_accessLevelTypeCurrentWrite AccessLevelType = 2
	AccessLevelType_accessLevelTypeHistoryRead AccessLevelType = 4
	AccessLevelType_accessLevelTypeHistoryWrite AccessLevelType = 8
	AccessLevelType_accessLevelTypeSemanticChange AccessLevelType = 16
	AccessLevelType_accessLevelTypeStatusWrite AccessLevelType = 32
	AccessLevelType_accessLevelTypeTimestampWrite AccessLevelType = 64
)

var AccessLevelTypeValues []AccessLevelType

func init() {
	_ = errors.New
	AccessLevelTypeValues = []AccessLevelType {
		AccessLevelType_accessLevelTypeNone,
		AccessLevelType_accessLevelTypeCurrentRead,
		AccessLevelType_accessLevelTypeCurrentWrite,
		AccessLevelType_accessLevelTypeHistoryRead,
		AccessLevelType_accessLevelTypeHistoryWrite,
		AccessLevelType_accessLevelTypeSemanticChange,
		AccessLevelType_accessLevelTypeStatusWrite,
		AccessLevelType_accessLevelTypeTimestampWrite,
	}
}

func AccessLevelTypeByValue(value uint8) (enum AccessLevelType, ok bool) {
	switch value {
		case 0:
			return AccessLevelType_accessLevelTypeNone, true
		case 1:
			return AccessLevelType_accessLevelTypeCurrentRead, true
		case 16:
			return AccessLevelType_accessLevelTypeSemanticChange, true
		case 2:
			return AccessLevelType_accessLevelTypeCurrentWrite, true
		case 32:
			return AccessLevelType_accessLevelTypeStatusWrite, true
		case 4:
			return AccessLevelType_accessLevelTypeHistoryRead, true
		case 64:
			return AccessLevelType_accessLevelTypeTimestampWrite, true
		case 8:
			return AccessLevelType_accessLevelTypeHistoryWrite, true
	}
	return 0, false
}

func AccessLevelTypeByName(value string) (enum AccessLevelType, ok bool) {
	switch value {
	case "accessLevelTypeNone":
		return AccessLevelType_accessLevelTypeNone, true
	case "accessLevelTypeCurrentRead":
		return AccessLevelType_accessLevelTypeCurrentRead, true
	case "accessLevelTypeSemanticChange":
		return AccessLevelType_accessLevelTypeSemanticChange, true
	case "accessLevelTypeCurrentWrite":
		return AccessLevelType_accessLevelTypeCurrentWrite, true
	case "accessLevelTypeStatusWrite":
		return AccessLevelType_accessLevelTypeStatusWrite, true
	case "accessLevelTypeHistoryRead":
		return AccessLevelType_accessLevelTypeHistoryRead, true
	case "accessLevelTypeTimestampWrite":
		return AccessLevelType_accessLevelTypeTimestampWrite, true
	case "accessLevelTypeHistoryWrite":
		return AccessLevelType_accessLevelTypeHistoryWrite, true
	}
	return 0, false
}

func AccessLevelTypeKnows(value uint8)  bool {
	for _, typeValue := range AccessLevelTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastAccessLevelType(structType any) AccessLevelType {
	castFunc := func(typ any) AccessLevelType {
		if sAccessLevelType, ok := typ.(AccessLevelType); ok {
			return sAccessLevelType
		}
		return 0
	}
	return castFunc(structType)
}

func (m AccessLevelType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m AccessLevelType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AccessLevelTypeParse(ctx context.Context, theBytes []byte) (AccessLevelType, error) {
	return AccessLevelTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AccessLevelTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AccessLevelType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("AccessLevelType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AccessLevelType")
	}
	if enum, ok := AccessLevelTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return AccessLevelType(val), nil
	} else {
		return enum, nil
	}
}

func (e AccessLevelType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AccessLevelType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("AccessLevelType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AccessLevelType) PLC4XEnumName() string {
	switch e {
	case AccessLevelType_accessLevelTypeNone:
		return "accessLevelTypeNone"
	case AccessLevelType_accessLevelTypeCurrentRead:
		return "accessLevelTypeCurrentRead"
	case AccessLevelType_accessLevelTypeSemanticChange:
		return "accessLevelTypeSemanticChange"
	case AccessLevelType_accessLevelTypeCurrentWrite:
		return "accessLevelTypeCurrentWrite"
	case AccessLevelType_accessLevelTypeStatusWrite:
		return "accessLevelTypeStatusWrite"
	case AccessLevelType_accessLevelTypeHistoryRead:
		return "accessLevelTypeHistoryRead"
	case AccessLevelType_accessLevelTypeTimestampWrite:
		return "accessLevelTypeTimestampWrite"
	case AccessLevelType_accessLevelTypeHistoryWrite:
		return "accessLevelTypeHistoryWrite"
	}
	return ""
}

func (e AccessLevelType) String() string {
	return e.PLC4XEnumName()
}

