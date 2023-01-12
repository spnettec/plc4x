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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority is an enum
type BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority uint8

type IBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority interface {
	utils.Serializable
}

const(
	BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_NORMAL BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority = 0
	BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_URGENT BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority = 1
)

var BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityValues []BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority

func init() {
	_ = errors.New
	BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityValues = []BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority {
		BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_NORMAL,
		BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_URGENT,
	}
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityByValue(value uint8) (enum BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority, ok bool) {
	switch value {
		case 0:
			return BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_NORMAL, true
		case 1:
			return BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_URGENT, true
	}
	return 0, false
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityByName(value string) (enum BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority, ok bool) {
	switch value {
	case "NORMAL":
		return BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_NORMAL, true
	case "URGENT":
		return BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_URGENT, true
	}
	return 0, false
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityKnows(value uint8)  bool {
	for _, typeValue := range BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority(structType interface{}) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority {
	castFunc := func(typ interface{}) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority {
		if sBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority, ok := typ.(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority); ok {
			return sBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityParse(theBytes []byte) (BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority, error) {
	return BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority, error) {
	val, err := readBuffer.ReadUint8("BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority")
	}
	if enum, ok := BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority) PLC4XEnumName() string {
	switch e {
	case BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_NORMAL:
		return "NORMAL"
	case BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_URGENT:
		return "URGENT"
	}
	return ""
}

func (e BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority) String() string {
	return e.PLC4XEnumName()
}

