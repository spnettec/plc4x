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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCResultCode is an enum
type BVLCResultCode uint16

type IBVLCResultCode interface {
	utils.Serializable
}

const(
	BVLCResultCode_SUCCESSFUL_COMPLETION BVLCResultCode = 0x0000
	BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK BVLCResultCode = 0x0010
	BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK BVLCResultCode = 0x0020
	BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK BVLCResultCode = 0x0030
	BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK BVLCResultCode = 0x0040
	BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK BVLCResultCode = 0x0050
	BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK BVLCResultCode = 0x0060
)

var BVLCResultCodeValues []BVLCResultCode

func init() {
	_ = errors.New
	BVLCResultCodeValues = []BVLCResultCode {
		BVLCResultCode_SUCCESSFUL_COMPLETION,
		BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK,
		BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK,
		BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK,
		BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK,
		BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK,
		BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK,
	}
}

func BVLCResultCodeByValue(value uint16) (enum BVLCResultCode, ok bool) {
	switch value {
		case 0x0000:
			return BVLCResultCode_SUCCESSFUL_COMPLETION, true
		case 0x0010:
			return BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK, true
		case 0x0020:
			return BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK, true
		case 0x0030:
			return BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK, true
		case 0x0040:
			return BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK, true
		case 0x0050:
			return BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK, true
		case 0x0060:
			return BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK, true
	}
	return 0, false
}

func BVLCResultCodeByName(value string) (enum BVLCResultCode, ok bool) {
	switch value {
	case "SUCCESSFUL_COMPLETION":
		return BVLCResultCode_SUCCESSFUL_COMPLETION, true
	case "WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK":
		return BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK, true
	case "READ_BROADCAST_DISTRIBUTION_TABLE_NAK":
		return BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK, true
	case "REGISTER_FOREIGN_DEVICE_NAK":
		return BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK, true
	case "READ_FOREIGN_DEVICE_TABLE_NAK":
		return BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK, true
	case "DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK":
		return BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK, true
	case "DISTRIBUTE_BROADCAST_TO_NETWORK_NAK":
		return BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK, true
	}
	return 0, false
}

func BVLCResultCodeKnows(value uint16)  bool {
	for _, typeValue := range BVLCResultCodeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBVLCResultCode(structType interface{}) BVLCResultCode {
	castFunc := func(typ interface{}) BVLCResultCode {
		if sBVLCResultCode, ok := typ.(BVLCResultCode); ok {
			return sBVLCResultCode
		}
		return 0
	}
	return castFunc(structType)
}

func (m BVLCResultCode) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BVLCResultCode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BVLCResultCodeParse(ctx context.Context, theBytes []byte) (BVLCResultCode, error) {
	return BVLCResultCodeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BVLCResultCodeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCResultCode, error) {
	val, err := readBuffer.ReadUint16("BVLCResultCode", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BVLCResultCode")
	}
	if enum, ok := BVLCResultCodeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BVLCResultCode(val), nil
	} else {
		return enum, nil
	}
}

func (e BVLCResultCode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BVLCResultCode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BVLCResultCode", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BVLCResultCode) PLC4XEnumName() string {
	switch e {
	case BVLCResultCode_SUCCESSFUL_COMPLETION:
		return "SUCCESSFUL_COMPLETION"
	case BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK:
		return "WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK"
	case BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK:
		return "READ_BROADCAST_DISTRIBUTION_TABLE_NAK"
	case BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK:
		return "REGISTER_FOREIGN_DEVICE_NAK"
	case BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK:
		return "READ_FOREIGN_DEVICE_TABLE_NAK"
	case BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK:
		return "DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK"
	case BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK:
		return "DISTRIBUTE_BROADCAST_TO_NETWORK_NAK"
	}
	return ""
}

func (e BVLCResultCode) String() string {
	return e.PLC4XEnumName()
}

