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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNetworkNumberQuality is an enum
type BACnetNetworkNumberQuality uint8

type IBACnetNetworkNumberQuality interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const(
	BACnetNetworkNumberQuality_UNKNOWN BACnetNetworkNumberQuality = 0
	BACnetNetworkNumberQuality_LEARNED BACnetNetworkNumberQuality = 1
	BACnetNetworkNumberQuality_LEARNED_CONFIGURED BACnetNetworkNumberQuality = 2
	BACnetNetworkNumberQuality_CONFIGURED BACnetNetworkNumberQuality = 3
)

var BACnetNetworkNumberQualityValues []BACnetNetworkNumberQuality

func init() {
	_ = errors.New
	BACnetNetworkNumberQualityValues = []BACnetNetworkNumberQuality {
		BACnetNetworkNumberQuality_UNKNOWN,
		BACnetNetworkNumberQuality_LEARNED,
		BACnetNetworkNumberQuality_LEARNED_CONFIGURED,
		BACnetNetworkNumberQuality_CONFIGURED,
	}
}

func BACnetNetworkNumberQualityByValue(value uint8) (enum BACnetNetworkNumberQuality, ok bool) {
	switch value {
		case 0:
			return BACnetNetworkNumberQuality_UNKNOWN, true
		case 1:
			return BACnetNetworkNumberQuality_LEARNED, true
		case 2:
			return BACnetNetworkNumberQuality_LEARNED_CONFIGURED, true
		case 3:
			return BACnetNetworkNumberQuality_CONFIGURED, true
	}
	return 0, false
}

func BACnetNetworkNumberQualityByName(value string) (enum BACnetNetworkNumberQuality, ok bool) {
	switch value {
	case "UNKNOWN":
		return BACnetNetworkNumberQuality_UNKNOWN, true
	case "LEARNED":
		return BACnetNetworkNumberQuality_LEARNED, true
	case "LEARNED_CONFIGURED":
		return BACnetNetworkNumberQuality_LEARNED_CONFIGURED, true
	case "CONFIGURED":
		return BACnetNetworkNumberQuality_CONFIGURED, true
	}
	return 0, false
}

func BACnetNetworkNumberQualityKnows(value uint8)  bool {
	for _, typeValue := range BACnetNetworkNumberQualityValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetNetworkNumberQuality(structType interface{}) BACnetNetworkNumberQuality {
	castFunc := func(typ interface{}) BACnetNetworkNumberQuality {
		if sBACnetNetworkNumberQuality, ok := typ.(BACnetNetworkNumberQuality); ok {
			return sBACnetNetworkNumberQuality
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetNetworkNumberQuality) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetNetworkNumberQuality) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNetworkNumberQualityParse(readBuffer utils.ReadBuffer) (BACnetNetworkNumberQuality, error) {
	val, err := readBuffer.ReadUint8("BACnetNetworkNumberQuality", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetNetworkNumberQuality")
	}
	if enum, ok := BACnetNetworkNumberQualityByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetNetworkNumberQuality(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetNetworkNumberQuality) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetNetworkNumberQuality", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetNetworkNumberQuality) PLC4XEnumName() string {
	switch e {
	case BACnetNetworkNumberQuality_UNKNOWN:
		return "UNKNOWN"
	case BACnetNetworkNumberQuality_LEARNED:
		return "LEARNED"
	case BACnetNetworkNumberQuality_LEARNED_CONFIGURED:
		return "LEARNED_CONFIGURED"
	case BACnetNetworkNumberQuality_CONFIGURED:
		return "CONFIGURED"
	}
	return ""
}

func (e BACnetNetworkNumberQuality) String() string {
	return e.PLC4XEnumName()
}

