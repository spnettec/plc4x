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

// BACnetSilencedState is an enum
type BACnetSilencedState uint16

type IBACnetSilencedState interface {
	utils.Serializable
}

const(
	BACnetSilencedState_UNSILENCED BACnetSilencedState = 0
	BACnetSilencedState_AUDIBLE_SILENCED BACnetSilencedState = 1
	BACnetSilencedState_VISIBLE_SILENCED BACnetSilencedState = 2
	BACnetSilencedState_ALL_SILENCED BACnetSilencedState = 3
	BACnetSilencedState_VENDOR_PROPRIETARY_VALUE BACnetSilencedState = 0XFFFF
)

var BACnetSilencedStateValues []BACnetSilencedState

func init() {
	_ = errors.New
	BACnetSilencedStateValues = []BACnetSilencedState {
		BACnetSilencedState_UNSILENCED,
		BACnetSilencedState_AUDIBLE_SILENCED,
		BACnetSilencedState_VISIBLE_SILENCED,
		BACnetSilencedState_ALL_SILENCED,
		BACnetSilencedState_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetSilencedStateByValue(value uint16) (enum BACnetSilencedState, ok bool) {
	switch value {
		case 0:
			return BACnetSilencedState_UNSILENCED, true
		case 0XFFFF:
			return BACnetSilencedState_VENDOR_PROPRIETARY_VALUE, true
		case 1:
			return BACnetSilencedState_AUDIBLE_SILENCED, true
		case 2:
			return BACnetSilencedState_VISIBLE_SILENCED, true
		case 3:
			return BACnetSilencedState_ALL_SILENCED, true
	}
	return 0, false
}

func BACnetSilencedStateByName(value string) (enum BACnetSilencedState, ok bool) {
	switch value {
	case "UNSILENCED":
		return BACnetSilencedState_UNSILENCED, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetSilencedState_VENDOR_PROPRIETARY_VALUE, true
	case "AUDIBLE_SILENCED":
		return BACnetSilencedState_AUDIBLE_SILENCED, true
	case "VISIBLE_SILENCED":
		return BACnetSilencedState_VISIBLE_SILENCED, true
	case "ALL_SILENCED":
		return BACnetSilencedState_ALL_SILENCED, true
	}
	return 0, false
}

func BACnetSilencedStateKnows(value uint16)  bool {
	for _, typeValue := range BACnetSilencedStateValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetSilencedState(structType interface{}) BACnetSilencedState {
	castFunc := func(typ interface{}) BACnetSilencedState {
		if sBACnetSilencedState, ok := typ.(BACnetSilencedState); ok {
			return sBACnetSilencedState
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetSilencedState) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetSilencedState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetSilencedStateParse(theBytes []byte) (BACnetSilencedState, error) {
	return BACnetSilencedStateParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetSilencedStateParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetSilencedState, error) {
	val, err := readBuffer.ReadUint16("BACnetSilencedState", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetSilencedState")
	}
	if enum, ok := BACnetSilencedStateByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetSilencedState(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetSilencedState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetSilencedState) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetSilencedState", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetSilencedState) PLC4XEnumName() string {
	switch e {
	case BACnetSilencedState_UNSILENCED:
		return "UNSILENCED"
	case BACnetSilencedState_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetSilencedState_AUDIBLE_SILENCED:
		return "AUDIBLE_SILENCED"
	case BACnetSilencedState_VISIBLE_SILENCED:
		return "VISIBLE_SILENCED"
	case BACnetSilencedState_ALL_SILENCED:
		return "ALL_SILENCED"
	}
	return ""
}

func (e BACnetSilencedState) String() string {
	return e.PLC4XEnumName()
}

