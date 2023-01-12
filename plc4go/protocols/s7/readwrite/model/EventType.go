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

// EventType is an enum
type EventType uint8

type IEventType interface {
	utils.Serializable
}

const(
	EventType_MODE EventType = 0x01
	EventType_SYS EventType = 0x02
	EventType_USR EventType = 0x04
	EventType_ALM EventType = 0x80
)

var EventTypeValues []EventType

func init() {
	_ = errors.New
	EventTypeValues = []EventType {
		EventType_MODE,
		EventType_SYS,
		EventType_USR,
		EventType_ALM,
	}
}

func EventTypeByValue(value uint8) (enum EventType, ok bool) {
	switch value {
		case 0x01:
			return EventType_MODE, true
		case 0x02:
			return EventType_SYS, true
		case 0x04:
			return EventType_USR, true
		case 0x80:
			return EventType_ALM, true
	}
	return 0, false
}

func EventTypeByName(value string) (enum EventType, ok bool) {
	switch value {
	case "MODE":
		return EventType_MODE, true
	case "SYS":
		return EventType_SYS, true
	case "USR":
		return EventType_USR, true
	case "ALM":
		return EventType_ALM, true
	}
	return 0, false
}

func EventTypeKnows(value uint8)  bool {
	for _, typeValue := range EventTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastEventType(structType interface{}) EventType {
	castFunc := func(typ interface{}) EventType {
		if sEventType, ok := typ.(EventType); ok {
			return sEventType
		}
		return 0
	}
	return castFunc(structType)
}

func (m EventType) GetLengthInBits() uint16 {
	return 8
}

func (m EventType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func EventTypeParse(theBytes []byte) (EventType, error) {
	return EventTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func EventTypeParseWithBuffer(readBuffer utils.ReadBuffer) (EventType, error) {
	val, err := readBuffer.ReadUint8("EventType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading EventType")
	}
	if enum, ok := EventTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return EventType(val), nil
	} else {
		return enum, nil
	}
}

func (e EventType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e EventType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("EventType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e EventType) PLC4XEnumName() string {
	switch e {
	case EventType_MODE:
		return "MODE"
	case EventType_SYS:
		return "SYS"
	case EventType_USR:
		return "USR"
	case EventType_ALM:
		return "ALM"
	}
	return ""
}

func (e EventType) String() string {
	return e.PLC4XEnumName()
}

