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

// UserTokenType is an enum
type UserTokenType uint32

type IUserTokenType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	UserTokenType_userTokenTypeAnonymous UserTokenType = 0
	UserTokenType_userTokenTypeUserName UserTokenType = 1
	UserTokenType_userTokenTypeCertificate UserTokenType = 2
	UserTokenType_userTokenTypeIssuedToken UserTokenType = 3
)

var UserTokenTypeValues []UserTokenType

func init() {
	_ = errors.New
	UserTokenTypeValues = []UserTokenType {
		UserTokenType_userTokenTypeAnonymous,
		UserTokenType_userTokenTypeUserName,
		UserTokenType_userTokenTypeCertificate,
		UserTokenType_userTokenTypeIssuedToken,
	}
}

func UserTokenTypeByValue(value uint32) (enum UserTokenType, ok bool) {
	switch value {
		case 0:
			return UserTokenType_userTokenTypeAnonymous, true
		case 1:
			return UserTokenType_userTokenTypeUserName, true
		case 2:
			return UserTokenType_userTokenTypeCertificate, true
		case 3:
			return UserTokenType_userTokenTypeIssuedToken, true
	}
	return 0, false
}

func UserTokenTypeByName(value string) (enum UserTokenType, ok bool) {
	switch value {
	case "userTokenTypeAnonymous":
		return UserTokenType_userTokenTypeAnonymous, true
	case "userTokenTypeUserName":
		return UserTokenType_userTokenTypeUserName, true
	case "userTokenTypeCertificate":
		return UserTokenType_userTokenTypeCertificate, true
	case "userTokenTypeIssuedToken":
		return UserTokenType_userTokenTypeIssuedToken, true
	}
	return 0, false
}

func UserTokenTypeKnows(value uint32)  bool {
	for _, typeValue := range UserTokenTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastUserTokenType(structType any) UserTokenType {
	castFunc := func(typ any) UserTokenType {
		if sUserTokenType, ok := typ.(UserTokenType); ok {
			return sUserTokenType
		}
		return 0
	}
	return castFunc(structType)
}

func (m UserTokenType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m UserTokenType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UserTokenTypeParse(ctx context.Context, theBytes []byte) (UserTokenType, error) {
	return UserTokenTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func UserTokenTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (UserTokenType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("UserTokenType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading UserTokenType")
	}
	if enum, ok := UserTokenTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return UserTokenType(val), nil
	} else {
		return enum, nil
	}
}

func (e UserTokenType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e UserTokenType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("UserTokenType", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e UserTokenType) PLC4XEnumName() string {
	switch e {
	case UserTokenType_userTokenTypeAnonymous:
		return "userTokenTypeAnonymous"
	case UserTokenType_userTokenTypeUserName:
		return "userTokenTypeUserName"
	case UserTokenType_userTokenTypeCertificate:
		return "userTokenTypeCertificate"
	case UserTokenType_userTokenTypeIssuedToken:
		return "userTokenTypeIssuedToken"
	}
	return ""
}

func (e UserTokenType) String() string {
	return e.PLC4XEnumName()
}

