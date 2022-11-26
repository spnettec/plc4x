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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetProtocolLevel is an enum
type BACnetProtocolLevel uint8

type IBACnetProtocolLevel interface {
	utils.Serializable
}

const (
	BACnetProtocolLevel_PHYSICAL               BACnetProtocolLevel = 0
	BACnetProtocolLevel_PROTOCOL               BACnetProtocolLevel = 1
	BACnetProtocolLevel_BACNET_APPLICATION     BACnetProtocolLevel = 2
	BACnetProtocolLevel_NON_BACNET_APPLICATION BACnetProtocolLevel = 3
)

var BACnetProtocolLevelValues []BACnetProtocolLevel

func init() {
	_ = errors.New
	BACnetProtocolLevelValues = []BACnetProtocolLevel{
		BACnetProtocolLevel_PHYSICAL,
		BACnetProtocolLevel_PROTOCOL,
		BACnetProtocolLevel_BACNET_APPLICATION,
		BACnetProtocolLevel_NON_BACNET_APPLICATION,
	}
}

func BACnetProtocolLevelByValue(value uint8) (enum BACnetProtocolLevel, ok bool) {
	switch value {
	case 0:
		return BACnetProtocolLevel_PHYSICAL, true
	case 1:
		return BACnetProtocolLevel_PROTOCOL, true
	case 2:
		return BACnetProtocolLevel_BACNET_APPLICATION, true
	case 3:
		return BACnetProtocolLevel_NON_BACNET_APPLICATION, true
	}
	return 0, false
}

func BACnetProtocolLevelByName(value string) (enum BACnetProtocolLevel, ok bool) {
	switch value {
	case "PHYSICAL":
		return BACnetProtocolLevel_PHYSICAL, true
	case "PROTOCOL":
		return BACnetProtocolLevel_PROTOCOL, true
	case "BACNET_APPLICATION":
		return BACnetProtocolLevel_BACNET_APPLICATION, true
	case "NON_BACNET_APPLICATION":
		return BACnetProtocolLevel_NON_BACNET_APPLICATION, true
	}
	return 0, false
}

func BACnetProtocolLevelKnows(value uint8) bool {
	for _, typeValue := range BACnetProtocolLevelValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetProtocolLevel(structType interface{}) BACnetProtocolLevel {
	castFunc := func(typ interface{}) BACnetProtocolLevel {
		if sBACnetProtocolLevel, ok := typ.(BACnetProtocolLevel); ok {
			return sBACnetProtocolLevel
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetProtocolLevel) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetProtocolLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetProtocolLevelParse(theBytes []byte) (BACnetProtocolLevel, error) {
	return BACnetProtocolLevelParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetProtocolLevelParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetProtocolLevel, error) {
	val, err := readBuffer.ReadUint8("BACnetProtocolLevel", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetProtocolLevel")
	}
	if enum, ok := BACnetProtocolLevelByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetProtocolLevel(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetProtocolLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetProtocolLevel) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetProtocolLevel", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetProtocolLevel) PLC4XEnumName() string {
	switch e {
	case BACnetProtocolLevel_PHYSICAL:
		return "PHYSICAL"
	case BACnetProtocolLevel_PROTOCOL:
		return "PROTOCOL"
	case BACnetProtocolLevel_BACNET_APPLICATION:
		return "BACNET_APPLICATION"
	case BACnetProtocolLevel_NON_BACNET_APPLICATION:
		return "NON_BACNET_APPLICATION"
	}
	return ""
}

func (e BACnetProtocolLevel) String() string {
	return e.PLC4XEnumName()
}
