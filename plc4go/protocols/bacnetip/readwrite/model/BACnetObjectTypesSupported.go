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
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetObjectTypesSupported is an enum
type BACnetObjectTypesSupported uint8

type IBACnetObjectTypesSupported interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	BACnetObjectTypesSupported_ANALOG_INPUT BACnetObjectTypesSupported = 0
	BACnetObjectTypesSupported_ANALOG_OUTPUT BACnetObjectTypesSupported = 1
	BACnetObjectTypesSupported_ANALOG_VALUE BACnetObjectTypesSupported = 2
	BACnetObjectTypesSupported_BINARY_INPUT BACnetObjectTypesSupported = 3
	BACnetObjectTypesSupported_BINARY_OUTPUT BACnetObjectTypesSupported = 4
	BACnetObjectTypesSupported_BINARY_VALUE BACnetObjectTypesSupported = 5
	BACnetObjectTypesSupported_CALENDAR BACnetObjectTypesSupported = 6
	BACnetObjectTypesSupported_COMMAND BACnetObjectTypesSupported = 7
	BACnetObjectTypesSupported_DEVICE BACnetObjectTypesSupported = 8
	BACnetObjectTypesSupported_EVENT_ENROLLMENT BACnetObjectTypesSupported = 9
	BACnetObjectTypesSupported_FILE BACnetObjectTypesSupported = 10
	BACnetObjectTypesSupported_GROUP BACnetObjectTypesSupported = 11
	BACnetObjectTypesSupported_LOOP BACnetObjectTypesSupported = 12
	BACnetObjectTypesSupported_MULTI_STATE_INPUT BACnetObjectTypesSupported = 13
	BACnetObjectTypesSupported_MULTI_STATE_OUTPUT BACnetObjectTypesSupported = 14
	BACnetObjectTypesSupported_NOTIFICATION_CLASS BACnetObjectTypesSupported = 15
	BACnetObjectTypesSupported_PROGRAM BACnetObjectTypesSupported = 16
	BACnetObjectTypesSupported_SCHEDULE BACnetObjectTypesSupported = 17
	BACnetObjectTypesSupported_AVERAGING BACnetObjectTypesSupported = 18
	BACnetObjectTypesSupported_MULTI_STATE_VALUE BACnetObjectTypesSupported = 19
	BACnetObjectTypesSupported_TREND_LOG BACnetObjectTypesSupported = 20
	BACnetObjectTypesSupported_LIFE_SAFETY_POINT BACnetObjectTypesSupported = 21
	BACnetObjectTypesSupported_LIFE_SAFETY_ZONE BACnetObjectTypesSupported = 22
	BACnetObjectTypesSupported_ACCUMULATOR BACnetObjectTypesSupported = 23
	BACnetObjectTypesSupported_PULSE_CONVERTER BACnetObjectTypesSupported = 24
	BACnetObjectTypesSupported_EVENT_LOG BACnetObjectTypesSupported = 25
	BACnetObjectTypesSupported_GLOBAL_GROUP BACnetObjectTypesSupported = 26
	BACnetObjectTypesSupported_TREND_LOG_MULTIPLE BACnetObjectTypesSupported = 27
	BACnetObjectTypesSupported_LOAD_CONTROL BACnetObjectTypesSupported = 28
	BACnetObjectTypesSupported_STRUCTURED_VIEW BACnetObjectTypesSupported = 29
	BACnetObjectTypesSupported_ACCESS_DOOR BACnetObjectTypesSupported = 30
	BACnetObjectTypesSupported_TIMER BACnetObjectTypesSupported = 31
	BACnetObjectTypesSupported_ACCESS_CREDENTIAL BACnetObjectTypesSupported = 32
	BACnetObjectTypesSupported_ACCESS_POINT BACnetObjectTypesSupported = 33
	BACnetObjectTypesSupported_ACCESS_RIGHTS BACnetObjectTypesSupported = 34
	BACnetObjectTypesSupported_ACCESS_USER BACnetObjectTypesSupported = 35
	BACnetObjectTypesSupported_ACCESS_ZONE BACnetObjectTypesSupported = 36
	BACnetObjectTypesSupported_CREDENTIAL_DATA_INPUT BACnetObjectTypesSupported = 37
	BACnetObjectTypesSupported_NETWORK_SECURITY BACnetObjectTypesSupported = 38
	BACnetObjectTypesSupported_BITSTRING_VALUE BACnetObjectTypesSupported = 39
	BACnetObjectTypesSupported_CHARACTERSTRING_VALUE BACnetObjectTypesSupported = 40
	BACnetObjectTypesSupported_DATEPATTERN_VALUE BACnetObjectTypesSupported = 41
	BACnetObjectTypesSupported_DATE_VALUE BACnetObjectTypesSupported = 42
	BACnetObjectTypesSupported_DATETIMEPATTERN_VALUE BACnetObjectTypesSupported = 43
	BACnetObjectTypesSupported_DATETIME_VALUE BACnetObjectTypesSupported = 44
	BACnetObjectTypesSupported_INTEGER_VALUE BACnetObjectTypesSupported = 45
	BACnetObjectTypesSupported_LARGE_ANALOG_VALUE BACnetObjectTypesSupported = 46
	BACnetObjectTypesSupported_OCTETSTRING_VALUE BACnetObjectTypesSupported = 47
	BACnetObjectTypesSupported_POSITIVE_INTEGER_VALUE BACnetObjectTypesSupported = 48
	BACnetObjectTypesSupported_TIMEPATTERN_VALUE BACnetObjectTypesSupported = 49
	BACnetObjectTypesSupported_TIME_VALUE BACnetObjectTypesSupported = 50
	BACnetObjectTypesSupported_NOTIFICATION_FORWARDER BACnetObjectTypesSupported = 51
	BACnetObjectTypesSupported_ALERT_ENROLLMENT BACnetObjectTypesSupported = 52
	BACnetObjectTypesSupported_CHANNEL BACnetObjectTypesSupported = 53
	BACnetObjectTypesSupported_LIGHTING_OUTPUT BACnetObjectTypesSupported = 54
	BACnetObjectTypesSupported_BINARY_LIGHTING_OUTPUT BACnetObjectTypesSupported = 55
	BACnetObjectTypesSupported_NETWORK_PORT BACnetObjectTypesSupported = 56
	BACnetObjectTypesSupported_ELEVATOR_GROUP BACnetObjectTypesSupported = 57
	BACnetObjectTypesSupported_ESCALATOR BACnetObjectTypesSupported = 58
	BACnetObjectTypesSupported_LIFT BACnetObjectTypesSupported = 59
)

var BACnetObjectTypesSupportedValues []BACnetObjectTypesSupported

func init() {
	_ = errors.New
	BACnetObjectTypesSupportedValues = []BACnetObjectTypesSupported {
		BACnetObjectTypesSupported_ANALOG_INPUT,
		BACnetObjectTypesSupported_ANALOG_OUTPUT,
		BACnetObjectTypesSupported_ANALOG_VALUE,
		BACnetObjectTypesSupported_BINARY_INPUT,
		BACnetObjectTypesSupported_BINARY_OUTPUT,
		BACnetObjectTypesSupported_BINARY_VALUE,
		BACnetObjectTypesSupported_CALENDAR,
		BACnetObjectTypesSupported_COMMAND,
		BACnetObjectTypesSupported_DEVICE,
		BACnetObjectTypesSupported_EVENT_ENROLLMENT,
		BACnetObjectTypesSupported_FILE,
		BACnetObjectTypesSupported_GROUP,
		BACnetObjectTypesSupported_LOOP,
		BACnetObjectTypesSupported_MULTI_STATE_INPUT,
		BACnetObjectTypesSupported_MULTI_STATE_OUTPUT,
		BACnetObjectTypesSupported_NOTIFICATION_CLASS,
		BACnetObjectTypesSupported_PROGRAM,
		BACnetObjectTypesSupported_SCHEDULE,
		BACnetObjectTypesSupported_AVERAGING,
		BACnetObjectTypesSupported_MULTI_STATE_VALUE,
		BACnetObjectTypesSupported_TREND_LOG,
		BACnetObjectTypesSupported_LIFE_SAFETY_POINT,
		BACnetObjectTypesSupported_LIFE_SAFETY_ZONE,
		BACnetObjectTypesSupported_ACCUMULATOR,
		BACnetObjectTypesSupported_PULSE_CONVERTER,
		BACnetObjectTypesSupported_EVENT_LOG,
		BACnetObjectTypesSupported_GLOBAL_GROUP,
		BACnetObjectTypesSupported_TREND_LOG_MULTIPLE,
		BACnetObjectTypesSupported_LOAD_CONTROL,
		BACnetObjectTypesSupported_STRUCTURED_VIEW,
		BACnetObjectTypesSupported_ACCESS_DOOR,
		BACnetObjectTypesSupported_TIMER,
		BACnetObjectTypesSupported_ACCESS_CREDENTIAL,
		BACnetObjectTypesSupported_ACCESS_POINT,
		BACnetObjectTypesSupported_ACCESS_RIGHTS,
		BACnetObjectTypesSupported_ACCESS_USER,
		BACnetObjectTypesSupported_ACCESS_ZONE,
		BACnetObjectTypesSupported_CREDENTIAL_DATA_INPUT,
		BACnetObjectTypesSupported_NETWORK_SECURITY,
		BACnetObjectTypesSupported_BITSTRING_VALUE,
		BACnetObjectTypesSupported_CHARACTERSTRING_VALUE,
		BACnetObjectTypesSupported_DATEPATTERN_VALUE,
		BACnetObjectTypesSupported_DATE_VALUE,
		BACnetObjectTypesSupported_DATETIMEPATTERN_VALUE,
		BACnetObjectTypesSupported_DATETIME_VALUE,
		BACnetObjectTypesSupported_INTEGER_VALUE,
		BACnetObjectTypesSupported_LARGE_ANALOG_VALUE,
		BACnetObjectTypesSupported_OCTETSTRING_VALUE,
		BACnetObjectTypesSupported_POSITIVE_INTEGER_VALUE,
		BACnetObjectTypesSupported_TIMEPATTERN_VALUE,
		BACnetObjectTypesSupported_TIME_VALUE,
		BACnetObjectTypesSupported_NOTIFICATION_FORWARDER,
		BACnetObjectTypesSupported_ALERT_ENROLLMENT,
		BACnetObjectTypesSupported_CHANNEL,
		BACnetObjectTypesSupported_LIGHTING_OUTPUT,
		BACnetObjectTypesSupported_BINARY_LIGHTING_OUTPUT,
		BACnetObjectTypesSupported_NETWORK_PORT,
		BACnetObjectTypesSupported_ELEVATOR_GROUP,
		BACnetObjectTypesSupported_ESCALATOR,
		BACnetObjectTypesSupported_LIFT,
	}
}

func BACnetObjectTypesSupportedByValue(value uint8) (enum BACnetObjectTypesSupported, ok bool) {
	switch value {
		case 0:
			return BACnetObjectTypesSupported_ANALOG_INPUT, true
		case 1:
			return BACnetObjectTypesSupported_ANALOG_OUTPUT, true
		case 10:
			return BACnetObjectTypesSupported_FILE, true
		case 11:
			return BACnetObjectTypesSupported_GROUP, true
		case 12:
			return BACnetObjectTypesSupported_LOOP, true
		case 13:
			return BACnetObjectTypesSupported_MULTI_STATE_INPUT, true
		case 14:
			return BACnetObjectTypesSupported_MULTI_STATE_OUTPUT, true
		case 15:
			return BACnetObjectTypesSupported_NOTIFICATION_CLASS, true
		case 16:
			return BACnetObjectTypesSupported_PROGRAM, true
		case 17:
			return BACnetObjectTypesSupported_SCHEDULE, true
		case 18:
			return BACnetObjectTypesSupported_AVERAGING, true
		case 19:
			return BACnetObjectTypesSupported_MULTI_STATE_VALUE, true
		case 2:
			return BACnetObjectTypesSupported_ANALOG_VALUE, true
		case 20:
			return BACnetObjectTypesSupported_TREND_LOG, true
		case 21:
			return BACnetObjectTypesSupported_LIFE_SAFETY_POINT, true
		case 22:
			return BACnetObjectTypesSupported_LIFE_SAFETY_ZONE, true
		case 23:
			return BACnetObjectTypesSupported_ACCUMULATOR, true
		case 24:
			return BACnetObjectTypesSupported_PULSE_CONVERTER, true
		case 25:
			return BACnetObjectTypesSupported_EVENT_LOG, true
		case 26:
			return BACnetObjectTypesSupported_GLOBAL_GROUP, true
		case 27:
			return BACnetObjectTypesSupported_TREND_LOG_MULTIPLE, true
		case 28:
			return BACnetObjectTypesSupported_LOAD_CONTROL, true
		case 29:
			return BACnetObjectTypesSupported_STRUCTURED_VIEW, true
		case 3:
			return BACnetObjectTypesSupported_BINARY_INPUT, true
		case 30:
			return BACnetObjectTypesSupported_ACCESS_DOOR, true
		case 31:
			return BACnetObjectTypesSupported_TIMER, true
		case 32:
			return BACnetObjectTypesSupported_ACCESS_CREDENTIAL, true
		case 33:
			return BACnetObjectTypesSupported_ACCESS_POINT, true
		case 34:
			return BACnetObjectTypesSupported_ACCESS_RIGHTS, true
		case 35:
			return BACnetObjectTypesSupported_ACCESS_USER, true
		case 36:
			return BACnetObjectTypesSupported_ACCESS_ZONE, true
		case 37:
			return BACnetObjectTypesSupported_CREDENTIAL_DATA_INPUT, true
		case 38:
			return BACnetObjectTypesSupported_NETWORK_SECURITY, true
		case 39:
			return BACnetObjectTypesSupported_BITSTRING_VALUE, true
		case 4:
			return BACnetObjectTypesSupported_BINARY_OUTPUT, true
		case 40:
			return BACnetObjectTypesSupported_CHARACTERSTRING_VALUE, true
		case 41:
			return BACnetObjectTypesSupported_DATEPATTERN_VALUE, true
		case 42:
			return BACnetObjectTypesSupported_DATE_VALUE, true
		case 43:
			return BACnetObjectTypesSupported_DATETIMEPATTERN_VALUE, true
		case 44:
			return BACnetObjectTypesSupported_DATETIME_VALUE, true
		case 45:
			return BACnetObjectTypesSupported_INTEGER_VALUE, true
		case 46:
			return BACnetObjectTypesSupported_LARGE_ANALOG_VALUE, true
		case 47:
			return BACnetObjectTypesSupported_OCTETSTRING_VALUE, true
		case 48:
			return BACnetObjectTypesSupported_POSITIVE_INTEGER_VALUE, true
		case 49:
			return BACnetObjectTypesSupported_TIMEPATTERN_VALUE, true
		case 5:
			return BACnetObjectTypesSupported_BINARY_VALUE, true
		case 50:
			return BACnetObjectTypesSupported_TIME_VALUE, true
		case 51:
			return BACnetObjectTypesSupported_NOTIFICATION_FORWARDER, true
		case 52:
			return BACnetObjectTypesSupported_ALERT_ENROLLMENT, true
		case 53:
			return BACnetObjectTypesSupported_CHANNEL, true
		case 54:
			return BACnetObjectTypesSupported_LIGHTING_OUTPUT, true
		case 55:
			return BACnetObjectTypesSupported_BINARY_LIGHTING_OUTPUT, true
		case 56:
			return BACnetObjectTypesSupported_NETWORK_PORT, true
		case 57:
			return BACnetObjectTypesSupported_ELEVATOR_GROUP, true
		case 58:
			return BACnetObjectTypesSupported_ESCALATOR, true
		case 59:
			return BACnetObjectTypesSupported_LIFT, true
		case 6:
			return BACnetObjectTypesSupported_CALENDAR, true
		case 7:
			return BACnetObjectTypesSupported_COMMAND, true
		case 8:
			return BACnetObjectTypesSupported_DEVICE, true
		case 9:
			return BACnetObjectTypesSupported_EVENT_ENROLLMENT, true
	}
	return 0, false
}

func BACnetObjectTypesSupportedByName(value string) (enum BACnetObjectTypesSupported, ok bool) {
	switch value {
	case "ANALOG_INPUT":
		return BACnetObjectTypesSupported_ANALOG_INPUT, true
	case "ANALOG_OUTPUT":
		return BACnetObjectTypesSupported_ANALOG_OUTPUT, true
	case "FILE":
		return BACnetObjectTypesSupported_FILE, true
	case "GROUP":
		return BACnetObjectTypesSupported_GROUP, true
	case "LOOP":
		return BACnetObjectTypesSupported_LOOP, true
	case "MULTI_STATE_INPUT":
		return BACnetObjectTypesSupported_MULTI_STATE_INPUT, true
	case "MULTI_STATE_OUTPUT":
		return BACnetObjectTypesSupported_MULTI_STATE_OUTPUT, true
	case "NOTIFICATION_CLASS":
		return BACnetObjectTypesSupported_NOTIFICATION_CLASS, true
	case "PROGRAM":
		return BACnetObjectTypesSupported_PROGRAM, true
	case "SCHEDULE":
		return BACnetObjectTypesSupported_SCHEDULE, true
	case "AVERAGING":
		return BACnetObjectTypesSupported_AVERAGING, true
	case "MULTI_STATE_VALUE":
		return BACnetObjectTypesSupported_MULTI_STATE_VALUE, true
	case "ANALOG_VALUE":
		return BACnetObjectTypesSupported_ANALOG_VALUE, true
	case "TREND_LOG":
		return BACnetObjectTypesSupported_TREND_LOG, true
	case "LIFE_SAFETY_POINT":
		return BACnetObjectTypesSupported_LIFE_SAFETY_POINT, true
	case "LIFE_SAFETY_ZONE":
		return BACnetObjectTypesSupported_LIFE_SAFETY_ZONE, true
	case "ACCUMULATOR":
		return BACnetObjectTypesSupported_ACCUMULATOR, true
	case "PULSE_CONVERTER":
		return BACnetObjectTypesSupported_PULSE_CONVERTER, true
	case "EVENT_LOG":
		return BACnetObjectTypesSupported_EVENT_LOG, true
	case "GLOBAL_GROUP":
		return BACnetObjectTypesSupported_GLOBAL_GROUP, true
	case "TREND_LOG_MULTIPLE":
		return BACnetObjectTypesSupported_TREND_LOG_MULTIPLE, true
	case "LOAD_CONTROL":
		return BACnetObjectTypesSupported_LOAD_CONTROL, true
	case "STRUCTURED_VIEW":
		return BACnetObjectTypesSupported_STRUCTURED_VIEW, true
	case "BINARY_INPUT":
		return BACnetObjectTypesSupported_BINARY_INPUT, true
	case "ACCESS_DOOR":
		return BACnetObjectTypesSupported_ACCESS_DOOR, true
	case "TIMER":
		return BACnetObjectTypesSupported_TIMER, true
	case "ACCESS_CREDENTIAL":
		return BACnetObjectTypesSupported_ACCESS_CREDENTIAL, true
	case "ACCESS_POINT":
		return BACnetObjectTypesSupported_ACCESS_POINT, true
	case "ACCESS_RIGHTS":
		return BACnetObjectTypesSupported_ACCESS_RIGHTS, true
	case "ACCESS_USER":
		return BACnetObjectTypesSupported_ACCESS_USER, true
	case "ACCESS_ZONE":
		return BACnetObjectTypesSupported_ACCESS_ZONE, true
	case "CREDENTIAL_DATA_INPUT":
		return BACnetObjectTypesSupported_CREDENTIAL_DATA_INPUT, true
	case "NETWORK_SECURITY":
		return BACnetObjectTypesSupported_NETWORK_SECURITY, true
	case "BITSTRING_VALUE":
		return BACnetObjectTypesSupported_BITSTRING_VALUE, true
	case "BINARY_OUTPUT":
		return BACnetObjectTypesSupported_BINARY_OUTPUT, true
	case "CHARACTERSTRING_VALUE":
		return BACnetObjectTypesSupported_CHARACTERSTRING_VALUE, true
	case "DATEPATTERN_VALUE":
		return BACnetObjectTypesSupported_DATEPATTERN_VALUE, true
	case "DATE_VALUE":
		return BACnetObjectTypesSupported_DATE_VALUE, true
	case "DATETIMEPATTERN_VALUE":
		return BACnetObjectTypesSupported_DATETIMEPATTERN_VALUE, true
	case "DATETIME_VALUE":
		return BACnetObjectTypesSupported_DATETIME_VALUE, true
	case "INTEGER_VALUE":
		return BACnetObjectTypesSupported_INTEGER_VALUE, true
	case "LARGE_ANALOG_VALUE":
		return BACnetObjectTypesSupported_LARGE_ANALOG_VALUE, true
	case "OCTETSTRING_VALUE":
		return BACnetObjectTypesSupported_OCTETSTRING_VALUE, true
	case "POSITIVE_INTEGER_VALUE":
		return BACnetObjectTypesSupported_POSITIVE_INTEGER_VALUE, true
	case "TIMEPATTERN_VALUE":
		return BACnetObjectTypesSupported_TIMEPATTERN_VALUE, true
	case "BINARY_VALUE":
		return BACnetObjectTypesSupported_BINARY_VALUE, true
	case "TIME_VALUE":
		return BACnetObjectTypesSupported_TIME_VALUE, true
	case "NOTIFICATION_FORWARDER":
		return BACnetObjectTypesSupported_NOTIFICATION_FORWARDER, true
	case "ALERT_ENROLLMENT":
		return BACnetObjectTypesSupported_ALERT_ENROLLMENT, true
	case "CHANNEL":
		return BACnetObjectTypesSupported_CHANNEL, true
	case "LIGHTING_OUTPUT":
		return BACnetObjectTypesSupported_LIGHTING_OUTPUT, true
	case "BINARY_LIGHTING_OUTPUT":
		return BACnetObjectTypesSupported_BINARY_LIGHTING_OUTPUT, true
	case "NETWORK_PORT":
		return BACnetObjectTypesSupported_NETWORK_PORT, true
	case "ELEVATOR_GROUP":
		return BACnetObjectTypesSupported_ELEVATOR_GROUP, true
	case "ESCALATOR":
		return BACnetObjectTypesSupported_ESCALATOR, true
	case "LIFT":
		return BACnetObjectTypesSupported_LIFT, true
	case "CALENDAR":
		return BACnetObjectTypesSupported_CALENDAR, true
	case "COMMAND":
		return BACnetObjectTypesSupported_COMMAND, true
	case "DEVICE":
		return BACnetObjectTypesSupported_DEVICE, true
	case "EVENT_ENROLLMENT":
		return BACnetObjectTypesSupported_EVENT_ENROLLMENT, true
	}
	return 0, false
}

func BACnetObjectTypesSupportedKnows(value uint8)  bool {
	for _, typeValue := range BACnetObjectTypesSupportedValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetObjectTypesSupported(structType any) BACnetObjectTypesSupported {
	castFunc := func(typ any) BACnetObjectTypesSupported {
		if sBACnetObjectTypesSupported, ok := typ.(BACnetObjectTypesSupported); ok {
			return sBACnetObjectTypesSupported
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetObjectTypesSupported) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetObjectTypesSupported) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetObjectTypesSupportedParse(ctx context.Context, theBytes []byte) (BACnetObjectTypesSupported, error) {
	return BACnetObjectTypesSupportedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetObjectTypesSupportedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetObjectTypesSupported, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetObjectTypesSupported", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetObjectTypesSupported")
	}
	if enum, ok := BACnetObjectTypesSupportedByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetObjectTypesSupported")
		return BACnetObjectTypesSupported(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetObjectTypesSupported) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetObjectTypesSupported) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetObjectTypesSupported", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetObjectTypesSupported) PLC4XEnumName() string {
	switch e {
	case BACnetObjectTypesSupported_ANALOG_INPUT:
		return "ANALOG_INPUT"
	case BACnetObjectTypesSupported_ANALOG_OUTPUT:
		return "ANALOG_OUTPUT"
	case BACnetObjectTypesSupported_FILE:
		return "FILE"
	case BACnetObjectTypesSupported_GROUP:
		return "GROUP"
	case BACnetObjectTypesSupported_LOOP:
		return "LOOP"
	case BACnetObjectTypesSupported_MULTI_STATE_INPUT:
		return "MULTI_STATE_INPUT"
	case BACnetObjectTypesSupported_MULTI_STATE_OUTPUT:
		return "MULTI_STATE_OUTPUT"
	case BACnetObjectTypesSupported_NOTIFICATION_CLASS:
		return "NOTIFICATION_CLASS"
	case BACnetObjectTypesSupported_PROGRAM:
		return "PROGRAM"
	case BACnetObjectTypesSupported_SCHEDULE:
		return "SCHEDULE"
	case BACnetObjectTypesSupported_AVERAGING:
		return "AVERAGING"
	case BACnetObjectTypesSupported_MULTI_STATE_VALUE:
		return "MULTI_STATE_VALUE"
	case BACnetObjectTypesSupported_ANALOG_VALUE:
		return "ANALOG_VALUE"
	case BACnetObjectTypesSupported_TREND_LOG:
		return "TREND_LOG"
	case BACnetObjectTypesSupported_LIFE_SAFETY_POINT:
		return "LIFE_SAFETY_POINT"
	case BACnetObjectTypesSupported_LIFE_SAFETY_ZONE:
		return "LIFE_SAFETY_ZONE"
	case BACnetObjectTypesSupported_ACCUMULATOR:
		return "ACCUMULATOR"
	case BACnetObjectTypesSupported_PULSE_CONVERTER:
		return "PULSE_CONVERTER"
	case BACnetObjectTypesSupported_EVENT_LOG:
		return "EVENT_LOG"
	case BACnetObjectTypesSupported_GLOBAL_GROUP:
		return "GLOBAL_GROUP"
	case BACnetObjectTypesSupported_TREND_LOG_MULTIPLE:
		return "TREND_LOG_MULTIPLE"
	case BACnetObjectTypesSupported_LOAD_CONTROL:
		return "LOAD_CONTROL"
	case BACnetObjectTypesSupported_STRUCTURED_VIEW:
		return "STRUCTURED_VIEW"
	case BACnetObjectTypesSupported_BINARY_INPUT:
		return "BINARY_INPUT"
	case BACnetObjectTypesSupported_ACCESS_DOOR:
		return "ACCESS_DOOR"
	case BACnetObjectTypesSupported_TIMER:
		return "TIMER"
	case BACnetObjectTypesSupported_ACCESS_CREDENTIAL:
		return "ACCESS_CREDENTIAL"
	case BACnetObjectTypesSupported_ACCESS_POINT:
		return "ACCESS_POINT"
	case BACnetObjectTypesSupported_ACCESS_RIGHTS:
		return "ACCESS_RIGHTS"
	case BACnetObjectTypesSupported_ACCESS_USER:
		return "ACCESS_USER"
	case BACnetObjectTypesSupported_ACCESS_ZONE:
		return "ACCESS_ZONE"
	case BACnetObjectTypesSupported_CREDENTIAL_DATA_INPUT:
		return "CREDENTIAL_DATA_INPUT"
	case BACnetObjectTypesSupported_NETWORK_SECURITY:
		return "NETWORK_SECURITY"
	case BACnetObjectTypesSupported_BITSTRING_VALUE:
		return "BITSTRING_VALUE"
	case BACnetObjectTypesSupported_BINARY_OUTPUT:
		return "BINARY_OUTPUT"
	case BACnetObjectTypesSupported_CHARACTERSTRING_VALUE:
		return "CHARACTERSTRING_VALUE"
	case BACnetObjectTypesSupported_DATEPATTERN_VALUE:
		return "DATEPATTERN_VALUE"
	case BACnetObjectTypesSupported_DATE_VALUE:
		return "DATE_VALUE"
	case BACnetObjectTypesSupported_DATETIMEPATTERN_VALUE:
		return "DATETIMEPATTERN_VALUE"
	case BACnetObjectTypesSupported_DATETIME_VALUE:
		return "DATETIME_VALUE"
	case BACnetObjectTypesSupported_INTEGER_VALUE:
		return "INTEGER_VALUE"
	case BACnetObjectTypesSupported_LARGE_ANALOG_VALUE:
		return "LARGE_ANALOG_VALUE"
	case BACnetObjectTypesSupported_OCTETSTRING_VALUE:
		return "OCTETSTRING_VALUE"
	case BACnetObjectTypesSupported_POSITIVE_INTEGER_VALUE:
		return "POSITIVE_INTEGER_VALUE"
	case BACnetObjectTypesSupported_TIMEPATTERN_VALUE:
		return "TIMEPATTERN_VALUE"
	case BACnetObjectTypesSupported_BINARY_VALUE:
		return "BINARY_VALUE"
	case BACnetObjectTypesSupported_TIME_VALUE:
		return "TIME_VALUE"
	case BACnetObjectTypesSupported_NOTIFICATION_FORWARDER:
		return "NOTIFICATION_FORWARDER"
	case BACnetObjectTypesSupported_ALERT_ENROLLMENT:
		return "ALERT_ENROLLMENT"
	case BACnetObjectTypesSupported_CHANNEL:
		return "CHANNEL"
	case BACnetObjectTypesSupported_LIGHTING_OUTPUT:
		return "LIGHTING_OUTPUT"
	case BACnetObjectTypesSupported_BINARY_LIGHTING_OUTPUT:
		return "BINARY_LIGHTING_OUTPUT"
	case BACnetObjectTypesSupported_NETWORK_PORT:
		return "NETWORK_PORT"
	case BACnetObjectTypesSupported_ELEVATOR_GROUP:
		return "ELEVATOR_GROUP"
	case BACnetObjectTypesSupported_ESCALATOR:
		return "ESCALATOR"
	case BACnetObjectTypesSupported_LIFT:
		return "LIFT"
	case BACnetObjectTypesSupported_CALENDAR:
		return "CALENDAR"
	case BACnetObjectTypesSupported_COMMAND:
		return "COMMAND"
	case BACnetObjectTypesSupported_DEVICE:
		return "DEVICE"
	case BACnetObjectTypesSupported_EVENT_ENROLLMENT:
		return "EVENT_ENROLLMENT"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetObjectTypesSupported) String() string {
	return e.PLC4XEnumName()
}

