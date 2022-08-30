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

// BACnetFileAccessMethod is an enum
type BACnetFileAccessMethod uint8

type IBACnetFileAccessMethod interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const(
	BACnetFileAccessMethod_RECORD_ACCESS BACnetFileAccessMethod = 0
	BACnetFileAccessMethod_STREAM_ACCESS BACnetFileAccessMethod = 1
)

var BACnetFileAccessMethodValues []BACnetFileAccessMethod

func init() {
	_ = errors.New
	BACnetFileAccessMethodValues = []BACnetFileAccessMethod {
		BACnetFileAccessMethod_RECORD_ACCESS,
		BACnetFileAccessMethod_STREAM_ACCESS,
	}
}

func BACnetFileAccessMethodByValue(value uint8) (enum BACnetFileAccessMethod, ok bool) {
	switch value {
		case 0:
			return BACnetFileAccessMethod_RECORD_ACCESS, true
		case 1:
			return BACnetFileAccessMethod_STREAM_ACCESS, true
	}
	return 0, false
}

func BACnetFileAccessMethodByName(value string) (enum BACnetFileAccessMethod, ok bool) {
	switch value {
	case "RECORD_ACCESS":
		return BACnetFileAccessMethod_RECORD_ACCESS, true
	case "STREAM_ACCESS":
		return BACnetFileAccessMethod_STREAM_ACCESS, true
	}
	return 0, false
}

func BACnetFileAccessMethodKnows(value uint8)  bool {
	for _, typeValue := range BACnetFileAccessMethodValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetFileAccessMethod(structType interface{}) BACnetFileAccessMethod {
	castFunc := func(typ interface{}) BACnetFileAccessMethod {
		if sBACnetFileAccessMethod, ok := typ.(BACnetFileAccessMethod); ok {
			return sBACnetFileAccessMethod
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetFileAccessMethod) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetFileAccessMethod) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFileAccessMethodParse(readBuffer utils.ReadBuffer) (BACnetFileAccessMethod, error) {
	val, err := readBuffer.ReadUint8("BACnetFileAccessMethod", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetFileAccessMethod")
	}
	if enum, ok := BACnetFileAccessMethodByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetFileAccessMethod(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetFileAccessMethod) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetFileAccessMethod", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetFileAccessMethod) PLC4XEnumName() string {
	switch e {
	case BACnetFileAccessMethod_RECORD_ACCESS:
		return "RECORD_ACCESS"
	case BACnetFileAccessMethod_STREAM_ACCESS:
		return "STREAM_ACCESS"
	}
	return ""
}

func (e BACnetFileAccessMethod) String() string {
	return e.PLC4XEnumName()
}

