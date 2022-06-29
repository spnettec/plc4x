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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceChoice is an enum
type BACnetConfirmedServiceChoice uint8

type IBACnetConfirmedServiceChoice interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM                   BACnetConfirmedServiceChoice = 0x00
	BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION          BACnetConfirmedServiceChoice = 0x01
	BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE BACnetConfirmedServiceChoice = 0x1F
	BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION        BACnetConfirmedServiceChoice = 0x02
	BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY                   BACnetConfirmedServiceChoice = 0x03
	BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY              BACnetConfirmedServiceChoice = 0x04
	BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION               BACnetConfirmedServiceChoice = 0x1D
	BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION               BACnetConfirmedServiceChoice = 0x1B
	BACnetConfirmedServiceChoice_SUBSCRIBE_COV                       BACnetConfirmedServiceChoice = 0x05
	BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY              BACnetConfirmedServiceChoice = 0x1C
	BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE     BACnetConfirmedServiceChoice = 0x1E
	BACnetConfirmedServiceChoice_ATOMIC_READ_FILE                    BACnetConfirmedServiceChoice = 0x06
	BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE                   BACnetConfirmedServiceChoice = 0x07
	BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT                    BACnetConfirmedServiceChoice = 0x08
	BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT                 BACnetConfirmedServiceChoice = 0x09
	BACnetConfirmedServiceChoice_CREATE_OBJECT                       BACnetConfirmedServiceChoice = 0x0A
	BACnetConfirmedServiceChoice_DELETE_OBJECT                       BACnetConfirmedServiceChoice = 0x0B
	BACnetConfirmedServiceChoice_READ_PROPERTY                       BACnetConfirmedServiceChoice = 0x0C
	BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE              BACnetConfirmedServiceChoice = 0x0E
	BACnetConfirmedServiceChoice_READ_RANGE                          BACnetConfirmedServiceChoice = 0x1A
	BACnetConfirmedServiceChoice_WRITE_PROPERTY                      BACnetConfirmedServiceChoice = 0x0F
	BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE             BACnetConfirmedServiceChoice = 0x10
	BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL        BACnetConfirmedServiceChoice = 0x11
	BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER          BACnetConfirmedServiceChoice = 0x12
	BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE              BACnetConfirmedServiceChoice = 0x13
	BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE                 BACnetConfirmedServiceChoice = 0x14
	BACnetConfirmedServiceChoice_VT_OPEN                             BACnetConfirmedServiceChoice = 0x15
	BACnetConfirmedServiceChoice_VT_CLOSE                            BACnetConfirmedServiceChoice = 0x16
	BACnetConfirmedServiceChoice_VT_DATA                             BACnetConfirmedServiceChoice = 0x17
	BACnetConfirmedServiceChoice_AUTHENTICATE                        BACnetConfirmedServiceChoice = 0x18
	BACnetConfirmedServiceChoice_REQUEST_KEY                         BACnetConfirmedServiceChoice = 0x19
	BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL           BACnetConfirmedServiceChoice = 0x0D
)

var BACnetConfirmedServiceChoiceValues []BACnetConfirmedServiceChoice

func init() {
	_ = errors.New
	BACnetConfirmedServiceChoiceValues = []BACnetConfirmedServiceChoice{
		BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM,
		BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION,
		BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE,
		BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION,
		BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY,
		BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY,
		BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION,
		BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION,
		BACnetConfirmedServiceChoice_SUBSCRIBE_COV,
		BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY,
		BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE,
		BACnetConfirmedServiceChoice_ATOMIC_READ_FILE,
		BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE,
		BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT,
		BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT,
		BACnetConfirmedServiceChoice_CREATE_OBJECT,
		BACnetConfirmedServiceChoice_DELETE_OBJECT,
		BACnetConfirmedServiceChoice_READ_PROPERTY,
		BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE,
		BACnetConfirmedServiceChoice_READ_RANGE,
		BACnetConfirmedServiceChoice_WRITE_PROPERTY,
		BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE,
		BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL,
		BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER,
		BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE,
		BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE,
		BACnetConfirmedServiceChoice_VT_OPEN,
		BACnetConfirmedServiceChoice_VT_CLOSE,
		BACnetConfirmedServiceChoice_VT_DATA,
		BACnetConfirmedServiceChoice_AUTHENTICATE,
		BACnetConfirmedServiceChoice_REQUEST_KEY,
		BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL,
	}
}

func BACnetConfirmedServiceChoiceByValue(value uint8) BACnetConfirmedServiceChoice {
	switch value {
	case 0x00:
		return BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM
	case 0x01:
		return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION
	case 0x02:
		return BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION
	case 0x03:
		return BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY
	case 0x04:
		return BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY
	case 0x05:
		return BACnetConfirmedServiceChoice_SUBSCRIBE_COV
	case 0x06:
		return BACnetConfirmedServiceChoice_ATOMIC_READ_FILE
	case 0x07:
		return BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE
	case 0x08:
		return BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT
	case 0x09:
		return BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT
	case 0x0A:
		return BACnetConfirmedServiceChoice_CREATE_OBJECT
	case 0x0B:
		return BACnetConfirmedServiceChoice_DELETE_OBJECT
	case 0x0C:
		return BACnetConfirmedServiceChoice_READ_PROPERTY
	case 0x0D:
		return BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL
	case 0x0E:
		return BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE
	case 0x0F:
		return BACnetConfirmedServiceChoice_WRITE_PROPERTY
	case 0x10:
		return BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE
	case 0x11:
		return BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL
	case 0x12:
		return BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER
	case 0x13:
		return BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE
	case 0x14:
		return BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE
	case 0x15:
		return BACnetConfirmedServiceChoice_VT_OPEN
	case 0x16:
		return BACnetConfirmedServiceChoice_VT_CLOSE
	case 0x17:
		return BACnetConfirmedServiceChoice_VT_DATA
	case 0x18:
		return BACnetConfirmedServiceChoice_AUTHENTICATE
	case 0x19:
		return BACnetConfirmedServiceChoice_REQUEST_KEY
	case 0x1A:
		return BACnetConfirmedServiceChoice_READ_RANGE
	case 0x1B:
		return BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION
	case 0x1C:
		return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY
	case 0x1D:
		return BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION
	case 0x1E:
		return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE
	case 0x1F:
		return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE
	}
	return 0
}

func BACnetConfirmedServiceChoiceByName(value string) (enum BACnetConfirmedServiceChoice, ok bool) {
	ok = true
	switch value {
	case "ACKNOWLEDGE_ALARM":
		enum = BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM
	case "CONFIRMED_COV_NOTIFICATION":
		enum = BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION
	case "CONFIRMED_EVENT_NOTIFICATION":
		enum = BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION
	case "GET_ALARM_SUMMARY":
		enum = BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY
	case "GET_ENROLLMENT_SUMMARY":
		enum = BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY
	case "SUBSCRIBE_COV":
		enum = BACnetConfirmedServiceChoice_SUBSCRIBE_COV
	case "ATOMIC_READ_FILE":
		enum = BACnetConfirmedServiceChoice_ATOMIC_READ_FILE
	case "ATOMIC_WRITE_FILE":
		enum = BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE
	case "ADD_LIST_ELEMENT":
		enum = BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT
	case "REMOVE_LIST_ELEMENT":
		enum = BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT
	case "CREATE_OBJECT":
		enum = BACnetConfirmedServiceChoice_CREATE_OBJECT
	case "DELETE_OBJECT":
		enum = BACnetConfirmedServiceChoice_DELETE_OBJECT
	case "READ_PROPERTY":
		enum = BACnetConfirmedServiceChoice_READ_PROPERTY
	case "READ_PROPERTY_CONDITIONAL":
		enum = BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL
	case "READ_PROPERTY_MULTIPLE":
		enum = BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE
	case "WRITE_PROPERTY":
		enum = BACnetConfirmedServiceChoice_WRITE_PROPERTY
	case "WRITE_PROPERTY_MULTIPLE":
		enum = BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE
	case "DEVICE_COMMUNICATION_CONTROL":
		enum = BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL
	case "CONFIRMED_PRIVATE_TRANSFER":
		enum = BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER
	case "CONFIRMED_TEXT_MESSAGE":
		enum = BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE
	case "REINITIALIZE_DEVICE":
		enum = BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE
	case "VT_OPEN":
		enum = BACnetConfirmedServiceChoice_VT_OPEN
	case "VT_CLOSE":
		enum = BACnetConfirmedServiceChoice_VT_CLOSE
	case "VT_DATA":
		enum = BACnetConfirmedServiceChoice_VT_DATA
	case "AUTHENTICATE":
		enum = BACnetConfirmedServiceChoice_AUTHENTICATE
	case "REQUEST_KEY":
		enum = BACnetConfirmedServiceChoice_REQUEST_KEY
	case "READ_RANGE":
		enum = BACnetConfirmedServiceChoice_READ_RANGE
	case "LIFE_SAFETY_OPERATION":
		enum = BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION
	case "SUBSCRIBE_COV_PROPERTY":
		enum = BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY
	case "GET_EVENT_INFORMATION":
		enum = BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION
	case "SUBSCRIBE_COV_PROPERTY_MULTIPLE":
		enum = BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE
	case "CONFIRMED_COV_NOTIFICATION_MULTIPLE":
		enum = BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetConfirmedServiceChoiceKnows(value uint8) bool {
	for _, typeValue := range BACnetConfirmedServiceChoiceValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetConfirmedServiceChoice(structType interface{}) BACnetConfirmedServiceChoice {
	castFunc := func(typ interface{}) BACnetConfirmedServiceChoice {
		if sBACnetConfirmedServiceChoice, ok := typ.(BACnetConfirmedServiceChoice); ok {
			return sBACnetConfirmedServiceChoice
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceChoice) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetConfirmedServiceChoice) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceChoiceParse(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceChoice, error) {
	val, err := readBuffer.ReadUint8("BACnetConfirmedServiceChoice", 8)
	if err != nil {
		return 0, nil
	}
	return BACnetConfirmedServiceChoiceByValue(val), nil
}

func (e BACnetConfirmedServiceChoice) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetConfirmedServiceChoice", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetConfirmedServiceChoice) PLC4XEnumName() string {
	switch e {
	case BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM:
		return "ACKNOWLEDGE_ALARM"
	case BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION:
		return "CONFIRMED_COV_NOTIFICATION"
	case BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION:
		return "CONFIRMED_EVENT_NOTIFICATION"
	case BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY:
		return "GET_ALARM_SUMMARY"
	case BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY:
		return "GET_ENROLLMENT_SUMMARY"
	case BACnetConfirmedServiceChoice_SUBSCRIBE_COV:
		return "SUBSCRIBE_COV"
	case BACnetConfirmedServiceChoice_ATOMIC_READ_FILE:
		return "ATOMIC_READ_FILE"
	case BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE:
		return "ATOMIC_WRITE_FILE"
	case BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT:
		return "ADD_LIST_ELEMENT"
	case BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT:
		return "REMOVE_LIST_ELEMENT"
	case BACnetConfirmedServiceChoice_CREATE_OBJECT:
		return "CREATE_OBJECT"
	case BACnetConfirmedServiceChoice_DELETE_OBJECT:
		return "DELETE_OBJECT"
	case BACnetConfirmedServiceChoice_READ_PROPERTY:
		return "READ_PROPERTY"
	case BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL:
		return "READ_PROPERTY_CONDITIONAL"
	case BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE:
		return "READ_PROPERTY_MULTIPLE"
	case BACnetConfirmedServiceChoice_WRITE_PROPERTY:
		return "WRITE_PROPERTY"
	case BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE:
		return "WRITE_PROPERTY_MULTIPLE"
	case BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL:
		return "DEVICE_COMMUNICATION_CONTROL"
	case BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER:
		return "CONFIRMED_PRIVATE_TRANSFER"
	case BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE:
		return "CONFIRMED_TEXT_MESSAGE"
	case BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE:
		return "REINITIALIZE_DEVICE"
	case BACnetConfirmedServiceChoice_VT_OPEN:
		return "VT_OPEN"
	case BACnetConfirmedServiceChoice_VT_CLOSE:
		return "VT_CLOSE"
	case BACnetConfirmedServiceChoice_VT_DATA:
		return "VT_DATA"
	case BACnetConfirmedServiceChoice_AUTHENTICATE:
		return "AUTHENTICATE"
	case BACnetConfirmedServiceChoice_REQUEST_KEY:
		return "REQUEST_KEY"
	case BACnetConfirmedServiceChoice_READ_RANGE:
		return "READ_RANGE"
	case BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION:
		return "LIFE_SAFETY_OPERATION"
	case BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY:
		return "SUBSCRIBE_COV_PROPERTY"
	case BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION:
		return "GET_EVENT_INFORMATION"
	case BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE:
		return "SUBSCRIBE_COV_PROPERTY_MULTIPLE"
	case BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE:
		return "CONFIRMED_COV_NOTIFICATION_MULTIPLE"
	}
	return ""
}

func (e BACnetConfirmedServiceChoice) String() string {
	return e.PLC4XEnumName()
}
