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

// IdentityCriteriaType is an enum
type IdentityCriteriaType uint32

type IIdentityCriteriaType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	IdentityCriteriaType_identityCriteriaTypeUserName IdentityCriteriaType = 1
	IdentityCriteriaType_identityCriteriaTypeThumbprint IdentityCriteriaType = 2
	IdentityCriteriaType_identityCriteriaTypeRole IdentityCriteriaType = 3
	IdentityCriteriaType_identityCriteriaTypeGroupId IdentityCriteriaType = 4
	IdentityCriteriaType_identityCriteriaTypeAnonymous IdentityCriteriaType = 5
	IdentityCriteriaType_identityCriteriaTypeAuthenticatedUser IdentityCriteriaType = 6
	IdentityCriteriaType_identityCriteriaTypeApplication IdentityCriteriaType = 7
	IdentityCriteriaType_identityCriteriaTypeX509Subject IdentityCriteriaType = 8
)

var IdentityCriteriaTypeValues []IdentityCriteriaType

func init() {
	_ = errors.New
	IdentityCriteriaTypeValues = []IdentityCriteriaType {
		IdentityCriteriaType_identityCriteriaTypeUserName,
		IdentityCriteriaType_identityCriteriaTypeThumbprint,
		IdentityCriteriaType_identityCriteriaTypeRole,
		IdentityCriteriaType_identityCriteriaTypeGroupId,
		IdentityCriteriaType_identityCriteriaTypeAnonymous,
		IdentityCriteriaType_identityCriteriaTypeAuthenticatedUser,
		IdentityCriteriaType_identityCriteriaTypeApplication,
		IdentityCriteriaType_identityCriteriaTypeX509Subject,
	}
}

func IdentityCriteriaTypeByValue(value uint32) (enum IdentityCriteriaType, ok bool) {
	switch value {
		case 1:
			return IdentityCriteriaType_identityCriteriaTypeUserName, true
		case 2:
			return IdentityCriteriaType_identityCriteriaTypeThumbprint, true
		case 3:
			return IdentityCriteriaType_identityCriteriaTypeRole, true
		case 4:
			return IdentityCriteriaType_identityCriteriaTypeGroupId, true
		case 5:
			return IdentityCriteriaType_identityCriteriaTypeAnonymous, true
		case 6:
			return IdentityCriteriaType_identityCriteriaTypeAuthenticatedUser, true
		case 7:
			return IdentityCriteriaType_identityCriteriaTypeApplication, true
		case 8:
			return IdentityCriteriaType_identityCriteriaTypeX509Subject, true
	}
	return 0, false
}

func IdentityCriteriaTypeByName(value string) (enum IdentityCriteriaType, ok bool) {
	switch value {
	case "identityCriteriaTypeUserName":
		return IdentityCriteriaType_identityCriteriaTypeUserName, true
	case "identityCriteriaTypeThumbprint":
		return IdentityCriteriaType_identityCriteriaTypeThumbprint, true
	case "identityCriteriaTypeRole":
		return IdentityCriteriaType_identityCriteriaTypeRole, true
	case "identityCriteriaTypeGroupId":
		return IdentityCriteriaType_identityCriteriaTypeGroupId, true
	case "identityCriteriaTypeAnonymous":
		return IdentityCriteriaType_identityCriteriaTypeAnonymous, true
	case "identityCriteriaTypeAuthenticatedUser":
		return IdentityCriteriaType_identityCriteriaTypeAuthenticatedUser, true
	case "identityCriteriaTypeApplication":
		return IdentityCriteriaType_identityCriteriaTypeApplication, true
	case "identityCriteriaTypeX509Subject":
		return IdentityCriteriaType_identityCriteriaTypeX509Subject, true
	}
	return 0, false
}

func IdentityCriteriaTypeKnows(value uint32)  bool {
	for _, typeValue := range IdentityCriteriaTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastIdentityCriteriaType(structType any) IdentityCriteriaType {
	castFunc := func(typ any) IdentityCriteriaType {
		if sIdentityCriteriaType, ok := typ.(IdentityCriteriaType); ok {
			return sIdentityCriteriaType
		}
		return 0
	}
	return castFunc(structType)
}

func (m IdentityCriteriaType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m IdentityCriteriaType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentityCriteriaTypeParse(ctx context.Context, theBytes []byte) (IdentityCriteriaType, error) {
	return IdentityCriteriaTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func IdentityCriteriaTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (IdentityCriteriaType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("IdentityCriteriaType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading IdentityCriteriaType")
	}
	if enum, ok := IdentityCriteriaTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for IdentityCriteriaType")
		return IdentityCriteriaType(val), nil
	} else {
		return enum, nil
	}
}

func (e IdentityCriteriaType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e IdentityCriteriaType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("IdentityCriteriaType", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e IdentityCriteriaType) PLC4XEnumName() string {
	switch e {
	case IdentityCriteriaType_identityCriteriaTypeUserName:
		return "identityCriteriaTypeUserName"
	case IdentityCriteriaType_identityCriteriaTypeThumbprint:
		return "identityCriteriaTypeThumbprint"
	case IdentityCriteriaType_identityCriteriaTypeRole:
		return "identityCriteriaTypeRole"
	case IdentityCriteriaType_identityCriteriaTypeGroupId:
		return "identityCriteriaTypeGroupId"
	case IdentityCriteriaType_identityCriteriaTypeAnonymous:
		return "identityCriteriaTypeAnonymous"
	case IdentityCriteriaType_identityCriteriaTypeAuthenticatedUser:
		return "identityCriteriaTypeAuthenticatedUser"
	case IdentityCriteriaType_identityCriteriaTypeApplication:
		return "identityCriteriaTypeApplication"
	case IdentityCriteriaType_identityCriteriaTypeX509Subject:
		return "identityCriteriaTypeX509Subject"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e IdentityCriteriaType) String() string {
	return e.PLC4XEnumName()
}

