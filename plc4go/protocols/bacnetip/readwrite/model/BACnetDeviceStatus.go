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
	"encoding/binary"

	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetDeviceStatus is an enum
type BACnetDeviceStatus uint16

type IBACnetDeviceStatus interface {
	utils.Serializable
}

const(
	BACnetDeviceStatus_OPERATIONAL BACnetDeviceStatus = 0
	BACnetDeviceStatus_OPERATIONAL_READ_ONLY BACnetDeviceStatus = 1
	BACnetDeviceStatus_DOWNLOAD_REQUIRED BACnetDeviceStatus = 2
	BACnetDeviceStatus_DOWNLOAD_IN_PROGRESS BACnetDeviceStatus = 3
	BACnetDeviceStatus_NON_OPERATIONAL BACnetDeviceStatus = 4
	BACnetDeviceStatus_BACKUP_IN_PROGRESS BACnetDeviceStatus = 5
	BACnetDeviceStatus_VENDOR_PROPRIETARY_VALUE BACnetDeviceStatus = 0XFFFF
)

var BACnetDeviceStatusValues []BACnetDeviceStatus

func init() {
	_ = errors.New
	BACnetDeviceStatusValues = []BACnetDeviceStatus {
		BACnetDeviceStatus_OPERATIONAL,
		BACnetDeviceStatus_OPERATIONAL_READ_ONLY,
		BACnetDeviceStatus_DOWNLOAD_REQUIRED,
		BACnetDeviceStatus_DOWNLOAD_IN_PROGRESS,
		BACnetDeviceStatus_NON_OPERATIONAL,
		BACnetDeviceStatus_BACKUP_IN_PROGRESS,
		BACnetDeviceStatus_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetDeviceStatusByValue(value uint16) (enum BACnetDeviceStatus, ok bool) {
	switch value {
		case 0:
			return BACnetDeviceStatus_OPERATIONAL, true
		case 0XFFFF:
			return BACnetDeviceStatus_VENDOR_PROPRIETARY_VALUE, true
		case 1:
			return BACnetDeviceStatus_OPERATIONAL_READ_ONLY, true
		case 2:
			return BACnetDeviceStatus_DOWNLOAD_REQUIRED, true
		case 3:
			return BACnetDeviceStatus_DOWNLOAD_IN_PROGRESS, true
		case 4:
			return BACnetDeviceStatus_NON_OPERATIONAL, true
		case 5:
			return BACnetDeviceStatus_BACKUP_IN_PROGRESS, true
	}
	return 0, false
}

func BACnetDeviceStatusByName(value string) (enum BACnetDeviceStatus, ok bool) {
	switch value {
	case "OPERATIONAL":
		return BACnetDeviceStatus_OPERATIONAL, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetDeviceStatus_VENDOR_PROPRIETARY_VALUE, true
	case "OPERATIONAL_READ_ONLY":
		return BACnetDeviceStatus_OPERATIONAL_READ_ONLY, true
	case "DOWNLOAD_REQUIRED":
		return BACnetDeviceStatus_DOWNLOAD_REQUIRED, true
	case "DOWNLOAD_IN_PROGRESS":
		return BACnetDeviceStatus_DOWNLOAD_IN_PROGRESS, true
	case "NON_OPERATIONAL":
		return BACnetDeviceStatus_NON_OPERATIONAL, true
	case "BACKUP_IN_PROGRESS":
		return BACnetDeviceStatus_BACKUP_IN_PROGRESS, true
	}
	return 0, false
}

func BACnetDeviceStatusKnows(value uint16)  bool {
	for _, typeValue := range BACnetDeviceStatusValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetDeviceStatus(structType interface{}) BACnetDeviceStatus {
	castFunc := func(typ interface{}) BACnetDeviceStatus {
		if sBACnetDeviceStatus, ok := typ.(BACnetDeviceStatus); ok {
			return sBACnetDeviceStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetDeviceStatus) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetDeviceStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDeviceStatusParse(readBuffer utils.ReadBuffer) (BACnetDeviceStatus, error) {
	val, err := readBuffer.ReadUint16("BACnetDeviceStatus", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetDeviceStatus")
	}
	if enum, ok := BACnetDeviceStatusByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetDeviceStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetDeviceStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetDeviceStatus) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetDeviceStatus", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetDeviceStatus) PLC4XEnumName() string {
	switch e {
	case BACnetDeviceStatus_OPERATIONAL:
		return "OPERATIONAL"
	case BACnetDeviceStatus_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetDeviceStatus_OPERATIONAL_READ_ONLY:
		return "OPERATIONAL_READ_ONLY"
	case BACnetDeviceStatus_DOWNLOAD_REQUIRED:
		return "DOWNLOAD_REQUIRED"
	case BACnetDeviceStatus_DOWNLOAD_IN_PROGRESS:
		return "DOWNLOAD_IN_PROGRESS"
	case BACnetDeviceStatus_NON_OPERATIONAL:
		return "NON_OPERATIONAL"
	case BACnetDeviceStatus_BACKUP_IN_PROGRESS:
		return "BACKUP_IN_PROGRESS"
	}
	return ""
}

func (e BACnetDeviceStatus) String() string {
	return e.PLC4XEnumName()
}

