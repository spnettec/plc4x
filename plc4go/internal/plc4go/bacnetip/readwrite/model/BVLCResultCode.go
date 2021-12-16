/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type BVLCResultCode uint16

type IBVLCResultCode interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BVLCResultCode_SUCCESSFUL_COMPLETION                  BVLCResultCode = 0x0000
	BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK BVLCResultCode = 0x0010
	BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK  BVLCResultCode = 0x0020
	BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK            BVLCResultCode = 0x0030
	BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK          BVLCResultCode = 0x0040
	BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK  BVLCResultCode = 0x0050
	BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK    BVLCResultCode = 0x0060
)

var BVLCResultCodeValues []BVLCResultCode

func init() {
	_ = errors.New
	BVLCResultCodeValues = []BVLCResultCode{
		BVLCResultCode_SUCCESSFUL_COMPLETION,
		BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK,
		BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK,
		BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK,
		BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK,
		BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK,
		BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK,
	}
}

func BVLCResultCodeByValue(value uint16) BVLCResultCode {
	switch value {
	case 0x0000:
		return BVLCResultCode_SUCCESSFUL_COMPLETION
	case 0x0010:
		return BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK
	case 0x0020:
		return BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK
	case 0x0030:
		return BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK
	case 0x0040:
		return BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK
	case 0x0050:
		return BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK
	case 0x0060:
		return BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK
	}
	return 0
}

func BVLCResultCodeByName(value string) BVLCResultCode {
	switch value {
	case "SUCCESSFUL_COMPLETION":
		return BVLCResultCode_SUCCESSFUL_COMPLETION
	case "WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK":
		return BVLCResultCode_WRITE_BROADCAST_DISTRIBUTION_TABLE_NAK
	case "READ_BROADCAST_DISTRIBUTION_TABLE_NAK":
		return BVLCResultCode_READ_BROADCAST_DISTRIBUTION_TABLE_NAK
	case "REGISTER_FOREIGN_DEVICE_NAK":
		return BVLCResultCode_REGISTER_FOREIGN_DEVICE_NAK
	case "READ_FOREIGN_DEVICE_TABLE_NAK":
		return BVLCResultCode_READ_FOREIGN_DEVICE_TABLE_NAK
	case "DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK":
		return BVLCResultCode_DELETE_FOREIGN_DEVICE_TABLE_ENTRY_NAK
	case "DISTRIBUTE_BROADCAST_TO_NETWORK_NAK":
		return BVLCResultCode_DISTRIBUTE_BROADCAST_TO_NETWORK_NAK
	}
	return 0
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

func (m BVLCResultCode) LengthInBits() uint16 {
	return 16
}

func (m BVLCResultCode) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCResultCodeParse(readBuffer utils.ReadBuffer) (BVLCResultCode, error) {
	val, err := readBuffer.ReadUint16("BVLCResultCode", 16)
	if err != nil {
		return 0, nil
	}
	return BVLCResultCodeByValue(val), nil
}

func (e BVLCResultCode) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BVLCResultCode", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BVLCResultCode) name() string {
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
	return e.name()
}
