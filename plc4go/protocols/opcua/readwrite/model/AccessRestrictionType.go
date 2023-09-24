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

// AccessRestrictionType is an enum
type AccessRestrictionType uint16

type IAccessRestrictionType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	AccessRestrictionType_accessRestrictionTypeNone AccessRestrictionType = 0
	AccessRestrictionType_accessRestrictionTypeSigningRequired AccessRestrictionType = 1
	AccessRestrictionType_accessRestrictionTypeEncryptionRequired AccessRestrictionType = 2
	AccessRestrictionType_accessRestrictionTypeSessionRequired AccessRestrictionType = 4
	AccessRestrictionType_accessRestrictionTypeApplyRestrictionsToBrowse AccessRestrictionType = 8
)

var AccessRestrictionTypeValues []AccessRestrictionType

func init() {
	_ = errors.New
	AccessRestrictionTypeValues = []AccessRestrictionType {
		AccessRestrictionType_accessRestrictionTypeNone,
		AccessRestrictionType_accessRestrictionTypeSigningRequired,
		AccessRestrictionType_accessRestrictionTypeEncryptionRequired,
		AccessRestrictionType_accessRestrictionTypeSessionRequired,
		AccessRestrictionType_accessRestrictionTypeApplyRestrictionsToBrowse,
	}
}

func AccessRestrictionTypeByValue(value uint16) (enum AccessRestrictionType, ok bool) {
	switch value {
		case 0:
			return AccessRestrictionType_accessRestrictionTypeNone, true
		case 1:
			return AccessRestrictionType_accessRestrictionTypeSigningRequired, true
		case 2:
			return AccessRestrictionType_accessRestrictionTypeEncryptionRequired, true
		case 4:
			return AccessRestrictionType_accessRestrictionTypeSessionRequired, true
		case 8:
			return AccessRestrictionType_accessRestrictionTypeApplyRestrictionsToBrowse, true
	}
	return 0, false
}

func AccessRestrictionTypeByName(value string) (enum AccessRestrictionType, ok bool) {
	switch value {
	case "accessRestrictionTypeNone":
		return AccessRestrictionType_accessRestrictionTypeNone, true
	case "accessRestrictionTypeSigningRequired":
		return AccessRestrictionType_accessRestrictionTypeSigningRequired, true
	case "accessRestrictionTypeEncryptionRequired":
		return AccessRestrictionType_accessRestrictionTypeEncryptionRequired, true
	case "accessRestrictionTypeSessionRequired":
		return AccessRestrictionType_accessRestrictionTypeSessionRequired, true
	case "accessRestrictionTypeApplyRestrictionsToBrowse":
		return AccessRestrictionType_accessRestrictionTypeApplyRestrictionsToBrowse, true
	}
	return 0, false
}

func AccessRestrictionTypeKnows(value uint16)  bool {
	for _, typeValue := range AccessRestrictionTypeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastAccessRestrictionType(structType any) AccessRestrictionType {
	castFunc := func(typ any) AccessRestrictionType {
		if sAccessRestrictionType, ok := typ.(AccessRestrictionType); ok {
			return sAccessRestrictionType
		}
		return 0
	}
	return castFunc(structType)
}

func (m AccessRestrictionType) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m AccessRestrictionType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AccessRestrictionTypeParse(ctx context.Context, theBytes []byte) (AccessRestrictionType, error) {
	return AccessRestrictionTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AccessRestrictionTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AccessRestrictionType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("AccessRestrictionType", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AccessRestrictionType")
	}
	if enum, ok := AccessRestrictionTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for AccessRestrictionType")
		return AccessRestrictionType(val), nil
	} else {
		return enum, nil
	}
}

func (e AccessRestrictionType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AccessRestrictionType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("AccessRestrictionType", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AccessRestrictionType) PLC4XEnumName() string {
	switch e {
	case AccessRestrictionType_accessRestrictionTypeNone:
		return "accessRestrictionTypeNone"
	case AccessRestrictionType_accessRestrictionTypeSigningRequired:
		return "accessRestrictionTypeSigningRequired"
	case AccessRestrictionType_accessRestrictionTypeEncryptionRequired:
		return "accessRestrictionTypeEncryptionRequired"
	case AccessRestrictionType_accessRestrictionTypeSessionRequired:
		return "accessRestrictionTypeSessionRequired"
	case AccessRestrictionType_accessRestrictionTypeApplyRestrictionsToBrowse:
		return "accessRestrictionTypeApplyRestrictionsToBrowse"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e AccessRestrictionType) String() string {
	return e.PLC4XEnumName()
}

