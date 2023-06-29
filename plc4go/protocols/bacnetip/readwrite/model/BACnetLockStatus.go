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

// BACnetLockStatus is an enum
type BACnetLockStatus uint8

type IBACnetLockStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetLockStatus_LOCKED     BACnetLockStatus = 0
	BACnetLockStatus_UNLOCKED   BACnetLockStatus = 1
	BACnetLockStatus_LOCK_FAULT BACnetLockStatus = 2
	BACnetLockStatus_UNUSED     BACnetLockStatus = 3
	BACnetLockStatus_UNKNOWN    BACnetLockStatus = 4
)

var BACnetLockStatusValues []BACnetLockStatus

func init() {
	_ = errors.New
	BACnetLockStatusValues = []BACnetLockStatus{
		BACnetLockStatus_LOCKED,
		BACnetLockStatus_UNLOCKED,
		BACnetLockStatus_LOCK_FAULT,
		BACnetLockStatus_UNUSED,
		BACnetLockStatus_UNKNOWN,
	}
}

func BACnetLockStatusByValue(value uint8) (enum BACnetLockStatus, ok bool) {
	switch value {
	case 0:
		return BACnetLockStatus_LOCKED, true
	case 1:
		return BACnetLockStatus_UNLOCKED, true
	case 2:
		return BACnetLockStatus_LOCK_FAULT, true
	case 3:
		return BACnetLockStatus_UNUSED, true
	case 4:
		return BACnetLockStatus_UNKNOWN, true
	}
	return 0, false
}

func BACnetLockStatusByName(value string) (enum BACnetLockStatus, ok bool) {
	switch value {
	case "LOCKED":
		return BACnetLockStatus_LOCKED, true
	case "UNLOCKED":
		return BACnetLockStatus_UNLOCKED, true
	case "LOCK_FAULT":
		return BACnetLockStatus_LOCK_FAULT, true
	case "UNUSED":
		return BACnetLockStatus_UNUSED, true
	case "UNKNOWN":
		return BACnetLockStatus_UNKNOWN, true
	}
	return 0, false
}

func BACnetLockStatusKnows(value uint8) bool {
	for _, typeValue := range BACnetLockStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLockStatus(structType any) BACnetLockStatus {
	castFunc := func(typ any) BACnetLockStatus {
		if sBACnetLockStatus, ok := typ.(BACnetLockStatus); ok {
			return sBACnetLockStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLockStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetLockStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLockStatusParse(ctx context.Context, theBytes []byte) (BACnetLockStatus, error) {
	return BACnetLockStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLockStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLockStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetLockStatus", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLockStatus")
	}
	if enum, ok := BACnetLockStatusByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetLockStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLockStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLockStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetLockStatus", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLockStatus) PLC4XEnumName() string {
	switch e {
	case BACnetLockStatus_LOCKED:
		return "LOCKED"
	case BACnetLockStatus_UNLOCKED:
		return "UNLOCKED"
	case BACnetLockStatus_LOCK_FAULT:
		return "LOCK_FAULT"
	case BACnetLockStatus_UNUSED:
		return "UNUSED"
	case BACnetLockStatus_UNKNOWN:
		return "UNKNOWN"
	}
	return ""
}

func (e BACnetLockStatus) String() string {
	return e.PLC4XEnumName()
}
