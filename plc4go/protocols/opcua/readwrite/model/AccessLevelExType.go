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

// AccessLevelExType is an enum
type AccessLevelExType uint32

type IAccessLevelExType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	AccessLevelExType_accessLevelExTypeNone AccessLevelExType = 0
	AccessLevelExType_accessLevelExTypeCurrentRead AccessLevelExType = 1
	AccessLevelExType_accessLevelExTypeCurrentWrite AccessLevelExType = 2
	AccessLevelExType_accessLevelExTypeHistoryRead AccessLevelExType = 4
	AccessLevelExType_accessLevelExTypeHistoryWrite AccessLevelExType = 8
	AccessLevelExType_accessLevelExTypeSemanticChange AccessLevelExType = 16
	AccessLevelExType_accessLevelExTypeStatusWrite AccessLevelExType = 32
	AccessLevelExType_accessLevelExTypeTimestampWrite AccessLevelExType = 64
	AccessLevelExType_accessLevelExTypeNonatomicRead AccessLevelExType = 256
	AccessLevelExType_accessLevelExTypeNonatomicWrite AccessLevelExType = 512
	AccessLevelExType_accessLevelExTypeWriteFullArrayOnly AccessLevelExType = 1024
	AccessLevelExType_accessLevelExTypeNoSubDataTypes AccessLevelExType = 2048
)

var AccessLevelExTypeValues []AccessLevelExType

func init() {
	_ = errors.New
	AccessLevelExTypeValues = []AccessLevelExType {
		AccessLevelExType_accessLevelExTypeNone,
		AccessLevelExType_accessLevelExTypeCurrentRead,
		AccessLevelExType_accessLevelExTypeCurrentWrite,
		AccessLevelExType_accessLevelExTypeHistoryRead,
		AccessLevelExType_accessLevelExTypeHistoryWrite,
		AccessLevelExType_accessLevelExTypeSemanticChange,
		AccessLevelExType_accessLevelExTypeStatusWrite,
		AccessLevelExType_accessLevelExTypeTimestampWrite,
		AccessLevelExType_accessLevelExTypeNonatomicRead,
		AccessLevelExType_accessLevelExTypeNonatomicWrite,
		AccessLevelExType_accessLevelExTypeWriteFullArrayOnly,
		AccessLevelExType_accessLevelExTypeNoSubDataTypes,
	}
}

func AccessLevelExTypeByValue(value uint32) (enum AccessLevelExType, ok bool) {
	switch value {
		case 0:
			return AccessLevelExType_accessLevelExTypeNone, true
		case 1:
			return AccessLevelExType_accessLevelExTypeCurrentRead, true
		case 1024:
			return AccessLevelExType_accessLevelExTypeWriteFullArrayOnly, true
		case 16:
			return AccessLevelExType_accessLevelExTypeSemanticChange, true
		case 2:
			return AccessLevelExType_accessLevelExTypeCurrentWrite, true
		case 2048:
			return AccessLevelExType_accessLevelExTypeNoSubDataTypes, true
		case 256:
			return AccessLevelExType_accessLevelExTypeNonatomicRead, true
		case 32:
			return AccessLevelExType_accessLevelExTypeStatusWrite, true
		case 4:
			return AccessLevelExType_accessLevelExTypeHistoryRead, true
		case 512:
			return AccessLevelExType_accessLevelExTypeNonatomicWrite, true
		case 64:
			return AccessLevelExType_accessLevelExTypeTimestampWrite, true
		case 8:
			return AccessLevelExType_accessLevelExTypeHistoryWrite, true
	}
	return 0, false
}

func AccessLevelExTypeByName(value string) (enum AccessLevelExType, ok bool) {
	switch value {
	case "accessLevelExTypeNone":
		return AccessLevelExType_accessLevelExTypeNone, true
	case "accessLevelExTypeCurrentRead":
		return AccessLevelExType_accessLevelExTypeCurrentRead, true
	case "accessLevelExTypeWriteFullArrayOnly":
		return AccessLevelExType_accessLevelExTypeWriteFullArrayOnly, true
	case "accessLevelExTypeSemanticChange":
		return AccessLevelExType_accessLevelExTypeSemanticChange, true
	case "accessLevelExTypeCurrentWrite":
		return AccessLevelExType_accessLevelExTypeCurrentWrite, true
	case "accessLevelExTypeNoSubDataTypes":
		return AccessLevelExType_accessLevelExTypeNoSubDataTypes, true
	case "accessLevelExTypeNonatomicRead":
		return AccessLevelExType_accessLevelExTypeNonatomicRead, true
	case "accessLevelExTypeStatusWrite":
		return AccessLevelExType_accessLevelExTypeStatusWrite, true
	case "accessLevelExTypeHistoryRead":
		return AccessLevelExType_accessLevelExTypeHistoryRead, true
	case "accessLevelExTypeNonatomicWrite":
		return AccessLevelExType_accessLevelExTypeNonatomicWrite, true
	case "accessLevelExTypeTimestampWrite":
		return AccessLevelExType_accessLevelExTypeTimestampWrite, true
	case "accessLevelExTypeHistoryWrite":
		return AccessLevelExType_accessLevelExTypeHistoryWrite, true
	}
	return 0, false
}

func AccessLevelExTypeKnows(value uint32)  bool {
	for _, typeValue := range AccessLevelExTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastAccessLevelExType(structType any) AccessLevelExType {
	castFunc := func(typ any) AccessLevelExType {
		if sAccessLevelExType, ok := typ.(AccessLevelExType); ok {
			return sAccessLevelExType
		}
		return 0
	}
	return castFunc(structType)
}

func (m AccessLevelExType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m AccessLevelExType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AccessLevelExTypeParse(ctx context.Context, theBytes []byte) (AccessLevelExType, error) {
	return AccessLevelExTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AccessLevelExTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AccessLevelExType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("AccessLevelExType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AccessLevelExType")
	}
	if enum, ok := AccessLevelExTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return AccessLevelExType(val), nil
	} else {
		return enum, nil
	}
}

func (e AccessLevelExType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AccessLevelExType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("AccessLevelExType", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AccessLevelExType) PLC4XEnumName() string {
	switch e {
	case AccessLevelExType_accessLevelExTypeNone:
		return "accessLevelExTypeNone"
	case AccessLevelExType_accessLevelExTypeCurrentRead:
		return "accessLevelExTypeCurrentRead"
	case AccessLevelExType_accessLevelExTypeWriteFullArrayOnly:
		return "accessLevelExTypeWriteFullArrayOnly"
	case AccessLevelExType_accessLevelExTypeSemanticChange:
		return "accessLevelExTypeSemanticChange"
	case AccessLevelExType_accessLevelExTypeCurrentWrite:
		return "accessLevelExTypeCurrentWrite"
	case AccessLevelExType_accessLevelExTypeNoSubDataTypes:
		return "accessLevelExTypeNoSubDataTypes"
	case AccessLevelExType_accessLevelExTypeNonatomicRead:
		return "accessLevelExTypeNonatomicRead"
	case AccessLevelExType_accessLevelExTypeStatusWrite:
		return "accessLevelExTypeStatusWrite"
	case AccessLevelExType_accessLevelExTypeHistoryRead:
		return "accessLevelExTypeHistoryRead"
	case AccessLevelExType_accessLevelExTypeNonatomicWrite:
		return "accessLevelExTypeNonatomicWrite"
	case AccessLevelExType_accessLevelExTypeTimestampWrite:
		return "accessLevelExTypeTimestampWrite"
	case AccessLevelExType_accessLevelExTypeHistoryWrite:
		return "accessLevelExTypeHistoryWrite"
	}
	return ""
}

func (e AccessLevelExType) String() string {
	return e.PLC4XEnumName()
}

