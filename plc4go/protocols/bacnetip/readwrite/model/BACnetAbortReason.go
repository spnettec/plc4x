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

// BACnetAbortReason is an enum
type BACnetAbortReason uint8

type IBACnetAbortReason interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	BACnetAbortReason_OTHER BACnetAbortReason = 0
	BACnetAbortReason_BUFFER_OVERFLOW BACnetAbortReason = 1
	BACnetAbortReason_INVALID_APDU_IN_THIS_STATE BACnetAbortReason = 2
	BACnetAbortReason_PREEMPTED_BY_HIGHER_PRIORITY_TASK BACnetAbortReason = 3
	BACnetAbortReason_SEGMENTATION_NOT_SUPPORTED BACnetAbortReason = 4
	BACnetAbortReason_SECURITY_ERROR BACnetAbortReason = 5
	BACnetAbortReason_INSUFFICIENT_SECURITY BACnetAbortReason = 6
	BACnetAbortReason_WINDOW_SIZE_OUT_OF_RANGE BACnetAbortReason = 7
	BACnetAbortReason_APPLICATION_EXCEEDED_REPLY_TIME BACnetAbortReason = 8
	BACnetAbortReason_OUT_OF_RESOURCES BACnetAbortReason = 9
	BACnetAbortReason_TSM_TIMEOUT BACnetAbortReason = 10
	BACnetAbortReason_APDU_TOO_LONG BACnetAbortReason = 11
	BACnetAbortReason_VENDOR_PROPRIETARY_VALUE BACnetAbortReason = 0xFF
)

var BACnetAbortReasonValues []BACnetAbortReason

func init() {
	_ = errors.New
	BACnetAbortReasonValues = []BACnetAbortReason {
		BACnetAbortReason_OTHER,
		BACnetAbortReason_BUFFER_OVERFLOW,
		BACnetAbortReason_INVALID_APDU_IN_THIS_STATE,
		BACnetAbortReason_PREEMPTED_BY_HIGHER_PRIORITY_TASK,
		BACnetAbortReason_SEGMENTATION_NOT_SUPPORTED,
		BACnetAbortReason_SECURITY_ERROR,
		BACnetAbortReason_INSUFFICIENT_SECURITY,
		BACnetAbortReason_WINDOW_SIZE_OUT_OF_RANGE,
		BACnetAbortReason_APPLICATION_EXCEEDED_REPLY_TIME,
		BACnetAbortReason_OUT_OF_RESOURCES,
		BACnetAbortReason_TSM_TIMEOUT,
		BACnetAbortReason_APDU_TOO_LONG,
		BACnetAbortReason_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAbortReasonByValue(value uint8) (enum BACnetAbortReason, ok bool) {
	switch value {
		case 0:
			return BACnetAbortReason_OTHER, true
		case 0xFF:
			return BACnetAbortReason_VENDOR_PROPRIETARY_VALUE, true
		case 1:
			return BACnetAbortReason_BUFFER_OVERFLOW, true
		case 10:
			return BACnetAbortReason_TSM_TIMEOUT, true
		case 11:
			return BACnetAbortReason_APDU_TOO_LONG, true
		case 2:
			return BACnetAbortReason_INVALID_APDU_IN_THIS_STATE, true
		case 3:
			return BACnetAbortReason_PREEMPTED_BY_HIGHER_PRIORITY_TASK, true
		case 4:
			return BACnetAbortReason_SEGMENTATION_NOT_SUPPORTED, true
		case 5:
			return BACnetAbortReason_SECURITY_ERROR, true
		case 6:
			return BACnetAbortReason_INSUFFICIENT_SECURITY, true
		case 7:
			return BACnetAbortReason_WINDOW_SIZE_OUT_OF_RANGE, true
		case 8:
			return BACnetAbortReason_APPLICATION_EXCEEDED_REPLY_TIME, true
		case 9:
			return BACnetAbortReason_OUT_OF_RESOURCES, true
	}
	return 0, false
}

func BACnetAbortReasonByName(value string) (enum BACnetAbortReason, ok bool) {
	switch value {
	case "OTHER":
		return BACnetAbortReason_OTHER, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetAbortReason_VENDOR_PROPRIETARY_VALUE, true
	case "BUFFER_OVERFLOW":
		return BACnetAbortReason_BUFFER_OVERFLOW, true
	case "TSM_TIMEOUT":
		return BACnetAbortReason_TSM_TIMEOUT, true
	case "APDU_TOO_LONG":
		return BACnetAbortReason_APDU_TOO_LONG, true
	case "INVALID_APDU_IN_THIS_STATE":
		return BACnetAbortReason_INVALID_APDU_IN_THIS_STATE, true
	case "PREEMPTED_BY_HIGHER_PRIORITY_TASK":
		return BACnetAbortReason_PREEMPTED_BY_HIGHER_PRIORITY_TASK, true
	case "SEGMENTATION_NOT_SUPPORTED":
		return BACnetAbortReason_SEGMENTATION_NOT_SUPPORTED, true
	case "SECURITY_ERROR":
		return BACnetAbortReason_SECURITY_ERROR, true
	case "INSUFFICIENT_SECURITY":
		return BACnetAbortReason_INSUFFICIENT_SECURITY, true
	case "WINDOW_SIZE_OUT_OF_RANGE":
		return BACnetAbortReason_WINDOW_SIZE_OUT_OF_RANGE, true
	case "APPLICATION_EXCEEDED_REPLY_TIME":
		return BACnetAbortReason_APPLICATION_EXCEEDED_REPLY_TIME, true
	case "OUT_OF_RESOURCES":
		return BACnetAbortReason_OUT_OF_RESOURCES, true
	}
	return 0, false
}

func BACnetAbortReasonKnows(value uint8)  bool {
	for _, typeValue := range BACnetAbortReasonValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetAbortReason(structType any) BACnetAbortReason {
	castFunc := func(typ any) BACnetAbortReason {
		if sBACnetAbortReason, ok := typ.(BACnetAbortReason); ok {
			return sBACnetAbortReason
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAbortReason) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetAbortReason) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAbortReasonParse(ctx context.Context, theBytes []byte) (BACnetAbortReason, error) {
	return BACnetAbortReasonParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAbortReasonParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAbortReason, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetAbortReason", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAbortReason")
	}
	if enum, ok := BACnetAbortReasonByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetAbortReason")
		return BACnetAbortReason(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAbortReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAbortReason) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetAbortReason", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAbortReason) PLC4XEnumName() string {
	switch e {
	case BACnetAbortReason_OTHER:
		return "OTHER"
	case BACnetAbortReason_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAbortReason_BUFFER_OVERFLOW:
		return "BUFFER_OVERFLOW"
	case BACnetAbortReason_TSM_TIMEOUT:
		return "TSM_TIMEOUT"
	case BACnetAbortReason_APDU_TOO_LONG:
		return "APDU_TOO_LONG"
	case BACnetAbortReason_INVALID_APDU_IN_THIS_STATE:
		return "INVALID_APDU_IN_THIS_STATE"
	case BACnetAbortReason_PREEMPTED_BY_HIGHER_PRIORITY_TASK:
		return "PREEMPTED_BY_HIGHER_PRIORITY_TASK"
	case BACnetAbortReason_SEGMENTATION_NOT_SUPPORTED:
		return "SEGMENTATION_NOT_SUPPORTED"
	case BACnetAbortReason_SECURITY_ERROR:
		return "SECURITY_ERROR"
	case BACnetAbortReason_INSUFFICIENT_SECURITY:
		return "INSUFFICIENT_SECURITY"
	case BACnetAbortReason_WINDOW_SIZE_OUT_OF_RANGE:
		return "WINDOW_SIZE_OUT_OF_RANGE"
	case BACnetAbortReason_APPLICATION_EXCEEDED_REPLY_TIME:
		return "APPLICATION_EXCEEDED_REPLY_TIME"
	case BACnetAbortReason_OUT_OF_RESOURCES:
		return "OUT_OF_RESOURCES"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetAbortReason) String() string {
	return e.PLC4XEnumName()
}

