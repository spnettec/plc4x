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

// BACnetAccumulatorRecordAccumulatorStatus is an enum
type BACnetAccumulatorRecordAccumulatorStatus uint8

type IBACnetAccumulatorRecordAccumulatorStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	BACnetAccumulatorRecordAccumulatorStatus_NORMAL BACnetAccumulatorRecordAccumulatorStatus = 0
	BACnetAccumulatorRecordAccumulatorStatus_STARTING BACnetAccumulatorRecordAccumulatorStatus = 1
	BACnetAccumulatorRecordAccumulatorStatus_RECOVERED BACnetAccumulatorRecordAccumulatorStatus = 2
	BACnetAccumulatorRecordAccumulatorStatus_ABNORMAL BACnetAccumulatorRecordAccumulatorStatus = 3
	BACnetAccumulatorRecordAccumulatorStatus_FAILED BACnetAccumulatorRecordAccumulatorStatus = 4
)

var BACnetAccumulatorRecordAccumulatorStatusValues []BACnetAccumulatorRecordAccumulatorStatus

func init() {
	_ = errors.New
	BACnetAccumulatorRecordAccumulatorStatusValues = []BACnetAccumulatorRecordAccumulatorStatus {
		BACnetAccumulatorRecordAccumulatorStatus_NORMAL,
		BACnetAccumulatorRecordAccumulatorStatus_STARTING,
		BACnetAccumulatorRecordAccumulatorStatus_RECOVERED,
		BACnetAccumulatorRecordAccumulatorStatus_ABNORMAL,
		BACnetAccumulatorRecordAccumulatorStatus_FAILED,
	}
}

func BACnetAccumulatorRecordAccumulatorStatusByValue(value uint8) (enum BACnetAccumulatorRecordAccumulatorStatus, ok bool) {
	switch value {
		case 0:
			return BACnetAccumulatorRecordAccumulatorStatus_NORMAL, true
		case 1:
			return BACnetAccumulatorRecordAccumulatorStatus_STARTING, true
		case 2:
			return BACnetAccumulatorRecordAccumulatorStatus_RECOVERED, true
		case 3:
			return BACnetAccumulatorRecordAccumulatorStatus_ABNORMAL, true
		case 4:
			return BACnetAccumulatorRecordAccumulatorStatus_FAILED, true
	}
	return 0, false
}

func BACnetAccumulatorRecordAccumulatorStatusByName(value string) (enum BACnetAccumulatorRecordAccumulatorStatus, ok bool) {
	switch value {
	case "NORMAL":
		return BACnetAccumulatorRecordAccumulatorStatus_NORMAL, true
	case "STARTING":
		return BACnetAccumulatorRecordAccumulatorStatus_STARTING, true
	case "RECOVERED":
		return BACnetAccumulatorRecordAccumulatorStatus_RECOVERED, true
	case "ABNORMAL":
		return BACnetAccumulatorRecordAccumulatorStatus_ABNORMAL, true
	case "FAILED":
		return BACnetAccumulatorRecordAccumulatorStatus_FAILED, true
	}
	return 0, false
}

func BACnetAccumulatorRecordAccumulatorStatusKnows(value uint8)  bool {
	for _, typeValue := range BACnetAccumulatorRecordAccumulatorStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetAccumulatorRecordAccumulatorStatus(structType any) BACnetAccumulatorRecordAccumulatorStatus {
	castFunc := func(typ any) BACnetAccumulatorRecordAccumulatorStatus {
		if sBACnetAccumulatorRecordAccumulatorStatus, ok := typ.(BACnetAccumulatorRecordAccumulatorStatus); ok {
			return sBACnetAccumulatorRecordAccumulatorStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccumulatorRecordAccumulatorStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetAccumulatorRecordAccumulatorStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccumulatorRecordAccumulatorStatusParse(ctx context.Context, theBytes []byte) (BACnetAccumulatorRecordAccumulatorStatus, error) {
	return BACnetAccumulatorRecordAccumulatorStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccumulatorRecordAccumulatorStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccumulatorRecordAccumulatorStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetAccumulatorRecordAccumulatorStatus", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAccumulatorRecordAccumulatorStatus")
	}
	if enum, ok := BACnetAccumulatorRecordAccumulatorStatusByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetAccumulatorRecordAccumulatorStatus")
		return BACnetAccumulatorRecordAccumulatorStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAccumulatorRecordAccumulatorStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAccumulatorRecordAccumulatorStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetAccumulatorRecordAccumulatorStatus", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccumulatorRecordAccumulatorStatus) PLC4XEnumName() string {
	switch e {
	case BACnetAccumulatorRecordAccumulatorStatus_NORMAL:
		return "NORMAL"
	case BACnetAccumulatorRecordAccumulatorStatus_STARTING:
		return "STARTING"
	case BACnetAccumulatorRecordAccumulatorStatus_RECOVERED:
		return "RECOVERED"
	case BACnetAccumulatorRecordAccumulatorStatus_ABNORMAL:
		return "ABNORMAL"
	case BACnetAccumulatorRecordAccumulatorStatus_FAILED:
		return "FAILED"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetAccumulatorRecordAccumulatorStatus) String() string {
	return e.PLC4XEnumName()
}

