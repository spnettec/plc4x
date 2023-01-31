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

// AdsTransMode is an enum
type AdsTransMode uint32

type IAdsTransMode interface {
	utils.Serializable
}

const(
	AdsTransMode_NONE AdsTransMode = 0
	AdsTransMode_CLIENT_CYCLE AdsTransMode = 1
	AdsTransMode_CLIENT_ON_CHANGE AdsTransMode = 2
	AdsTransMode_CYCLIC AdsTransMode = 3
	AdsTransMode_ON_CHANGE AdsTransMode = 4
	AdsTransMode_CYCLIC_IN_CONTEXT AdsTransMode = 5
	AdsTransMode_ON_CHANGE_IN_CONTEXT AdsTransMode = 6
)

var AdsTransModeValues []AdsTransMode

func init() {
	_ = errors.New
	AdsTransModeValues = []AdsTransMode {
		AdsTransMode_NONE,
		AdsTransMode_CLIENT_CYCLE,
		AdsTransMode_CLIENT_ON_CHANGE,
		AdsTransMode_CYCLIC,
		AdsTransMode_ON_CHANGE,
		AdsTransMode_CYCLIC_IN_CONTEXT,
		AdsTransMode_ON_CHANGE_IN_CONTEXT,
	}
}

func AdsTransModeByValue(value uint32) (enum AdsTransMode, ok bool) {
	switch value {
		case 0:
			return AdsTransMode_NONE, true
		case 1:
			return AdsTransMode_CLIENT_CYCLE, true
		case 2:
			return AdsTransMode_CLIENT_ON_CHANGE, true
		case 3:
			return AdsTransMode_CYCLIC, true
		case 4:
			return AdsTransMode_ON_CHANGE, true
		case 5:
			return AdsTransMode_CYCLIC_IN_CONTEXT, true
		case 6:
			return AdsTransMode_ON_CHANGE_IN_CONTEXT, true
	}
	return 0, false
}

func AdsTransModeByName(value string) (enum AdsTransMode, ok bool) {
	switch value {
	case "NONE":
		return AdsTransMode_NONE, true
	case "CLIENT_CYCLE":
		return AdsTransMode_CLIENT_CYCLE, true
	case "CLIENT_ON_CHANGE":
		return AdsTransMode_CLIENT_ON_CHANGE, true
	case "CYCLIC":
		return AdsTransMode_CYCLIC, true
	case "ON_CHANGE":
		return AdsTransMode_ON_CHANGE, true
	case "CYCLIC_IN_CONTEXT":
		return AdsTransMode_CYCLIC_IN_CONTEXT, true
	case "ON_CHANGE_IN_CONTEXT":
		return AdsTransMode_ON_CHANGE_IN_CONTEXT, true
	}
	return 0, false
}

func AdsTransModeKnows(value uint32)  bool {
	for _, typeValue := range AdsTransModeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastAdsTransMode(structType interface{}) AdsTransMode {
	castFunc := func(typ interface{}) AdsTransMode {
		if sAdsTransMode, ok := typ.(AdsTransMode); ok {
			return sAdsTransMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m AdsTransMode) GetLengthInBits() uint16 {
	return 32
}

func (m AdsTransMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsTransModeParse(theBytes []byte) (AdsTransMode, error) {
	return AdsTransModeParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func AdsTransModeParseWithBuffer(readBuffer utils.ReadBuffer) (AdsTransMode, error) {
	val, err := readBuffer.ReadUint32("AdsTransMode", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AdsTransMode")
	}
	if enum, ok := AdsTransModeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return AdsTransMode(val), nil
	} else {
		return enum, nil
	}
}

func (e AdsTransMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AdsTransMode) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint32("AdsTransMode", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AdsTransMode) PLC4XEnumName() string {
	switch e {
	case AdsTransMode_NONE:
		return "NONE"
	case AdsTransMode_CLIENT_CYCLE:
		return "CLIENT_CYCLE"
	case AdsTransMode_CLIENT_ON_CHANGE:
		return "CLIENT_ON_CHANGE"
	case AdsTransMode_CYCLIC:
		return "CYCLIC"
	case AdsTransMode_ON_CHANGE:
		return "ON_CHANGE"
	case AdsTransMode_CYCLIC_IN_CONTEXT:
		return "CYCLIC_IN_CONTEXT"
	case AdsTransMode_ON_CHANGE_IN_CONTEXT:
		return "ON_CHANGE_IN_CONTEXT"
	}
	return ""
}

func (e AdsTransMode) String() string {
	return e.PLC4XEnumName()
}

