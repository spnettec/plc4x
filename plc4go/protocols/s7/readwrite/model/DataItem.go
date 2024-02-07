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
	api "github.com/apache/plc4x/plc4go/pkg/api/values"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/apache/plc4x/plc4go/spi/values"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

func DataItemParse(ctx context.Context, theBytes []byte, dataProtocolId string, controllerType ControllerType, stringLength int32, stringEncoding string) (api.PlcValue, error) {
	return DataItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), dataProtocolId, controllerType, stringLength, stringEncoding)
}

func DataItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataProtocolId string, controllerType ControllerType, stringLength int32, stringEncoding string) (api.PlcValue, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	readBuffer.PullContext("DataItem")
	switch {
	case dataProtocolId == "IEC61131_BOOL": // BOOL
		// Reserved Field (Just skip the bytes)
		if _, _err := readBuffer.ReadUint8("reserved", 7); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadBit("value")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBOOL(value), nil
	case dataProtocolId == "IEC61131_BYTE": // BYTE
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBYTE(value), nil
	case dataProtocolId == "IEC61131_WORD": // WORD
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcWORD(value), nil
	case dataProtocolId == "IEC61131_DWORD": // DWORD
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDWORD(value), nil
	case dataProtocolId == "IEC61131_LWORD": // LWORD
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLWORD(value), nil
	case dataProtocolId == "IEC61131_SINT": // SINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSINT(value), nil
	case dataProtocolId == "IEC61131_USINT": // USINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUSINT(value), nil
	case dataProtocolId == "IEC61131_INT": // INT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcINT(value), nil
	case dataProtocolId == "IEC61131_UINT": // UINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUINT(value), nil
	case dataProtocolId == "IEC61131_DINT": // DINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDINT(value), nil
	case dataProtocolId == "IEC61131_UDINT": // UDINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUDINT(value), nil
	case dataProtocolId == "IEC61131_LINT": // LINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLINT(value), nil
	case dataProtocolId == "IEC61131_ULINT": // ULINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcULINT(value), nil
	case dataProtocolId == "IEC61131_REAL": // REAL
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadFloat32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcREAL(value), nil
	case dataProtocolId == "IEC61131_LREAL": // LREAL
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadFloat64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLREAL(value), nil
	case dataProtocolId == "IEC61131_CHAR": // CHAR
		// Manual Field (value)
		value, _valueErr := ParseS7Char(ctx, readBuffer, "UTF-8", stringEncoding)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcCHAR(value), nil
	case dataProtocolId == "IEC61131_WCHAR": // CHAR
		// Manual Field (value)
		value, _valueErr := ParseS7Char(ctx, readBuffer, "UTF-16", stringEncoding)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcCHAR(value), nil
	case dataProtocolId == "IEC61131_STRING": // STRING
		// Manual Field (value)
		value, _valueErr := ParseS7String(ctx, readBuffer, stringLength, "UTF-8", stringEncoding)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSTRING(value), nil
	case dataProtocolId == "IEC61131_WSTRING": // STRING
		// Manual Field (value)
		value, _valueErr := ParseS7String(ctx, readBuffer, stringLength, "UTF-16", stringEncoding)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		_ = value // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSTRING(value), nil
	case dataProtocolId == "IEC61131_TIME": // TIME
		// Simple Field (milliseconds)
		milliseconds, _millisecondsErr := readBuffer.ReadUint32("milliseconds", 32)
		if _millisecondsErr != nil {
			return nil, errors.Wrap(_millisecondsErr, "Error parsing 'milliseconds' field")
		}
		_ = milliseconds // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcTIMEFromMilliseconds(int64(milliseconds)), nil
	case dataProtocolId == "S7_S5TIME": // TIME
		// Manual Field (milliseconds)
		milliseconds, _millisecondsErr := ParseS5Time(ctx, readBuffer)
		if _millisecondsErr != nil {
			return nil, errors.Wrap(_millisecondsErr, "Error parsing 'milliseconds' field")
		}
		_ = milliseconds // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcTIMEFromMilliseconds(int64(milliseconds)), nil
	case dataProtocolId == "IEC61131_LTIME": // LTIME
		// Simple Field (nanoseconds)
		nanoseconds, _nanosecondsErr := readBuffer.ReadUint64("nanoseconds", 64)
		if _nanosecondsErr != nil {
			return nil, errors.Wrap(_nanosecondsErr, "Error parsing 'nanoseconds' field")
		}
		_ = nanoseconds // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLTIMEFromNanoseconds(nanoseconds), nil
	case dataProtocolId == "IEC61131_DATE": // DATE
		// Manual Field (daysSinceEpoch)
		daysSinceEpoch, _daysSinceEpochErr := ParseTiaDate(ctx, readBuffer)
		if _daysSinceEpochErr != nil {
			return nil, errors.Wrap(_daysSinceEpochErr, "Error parsing 'daysSinceEpoch' field")
		}
		_ = daysSinceEpoch // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDATEFromDaysSinceEpoch(daysSinceEpoch), nil
	case dataProtocolId == "IEC61131_TIME_OF_DAY": // TIME_OF_DAY
		// Simple Field (millisecondsSinceMidnight)
		millisecondsSinceMidnight, _millisecondsSinceMidnightErr := readBuffer.ReadUint32("millisecondsSinceMidnight", 32)
		if _millisecondsSinceMidnightErr != nil {
			return nil, errors.Wrap(_millisecondsSinceMidnightErr, "Error parsing 'millisecondsSinceMidnight' field")
		}
		_ = millisecondsSinceMidnight // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcTIME_OF_DAYFromMillisecondsSinceMidnight(millisecondsSinceMidnight), nil
	case dataProtocolId == "IEC61131_LTIME_OF_DAY": // LTIME_OF_DAY
		// Simple Field (nanosecondsSinceMidnight)
		nanosecondsSinceMidnight, _nanosecondsSinceMidnightErr := readBuffer.ReadUint64("nanosecondsSinceMidnight", 64)
		if _nanosecondsSinceMidnightErr != nil {
			return nil, errors.Wrap(_nanosecondsSinceMidnightErr, "Error parsing 'nanosecondsSinceMidnight' field")
		}
		_ = nanosecondsSinceMidnight // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLTIME_OF_DAYFromNanosecondsSinceMidnight(nanosecondsSinceMidnight), nil
	case dataProtocolId == "IEC61131_DATE_AND_TIME": // DATE_AND_TIME
		// Manual Field (year)
		year, _yearErr := ParseSiemensYear(ctx, readBuffer)
		if _yearErr != nil {
			return nil, errors.Wrap(_yearErr, "Error parsing 'year' field")
		}
		_ = year // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (month)
		month, _monthErr := readBuffer.ReadUint8("month", 8)
		if _monthErr != nil {
			return nil, errors.Wrap(_monthErr, "Error parsing 'month' field")
		}
		_ = month // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (day)
		day, _dayErr := readBuffer.ReadUint8("day", 8)
		if _dayErr != nil {
			return nil, errors.Wrap(_dayErr, "Error parsing 'day' field")
		}
		_ = day // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (hour)
		hour, _hourErr := readBuffer.ReadUint8("hour", 8)
		if _hourErr != nil {
			return nil, errors.Wrap(_hourErr, "Error parsing 'hour' field")
		}
		_ = hour // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (minutes)
		minutes, _minutesErr := readBuffer.ReadUint8("minutes", 8)
		if _minutesErr != nil {
			return nil, errors.Wrap(_minutesErr, "Error parsing 'minutes' field")
		}
		_ = minutes // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (seconds)
		seconds, _secondsErr := readBuffer.ReadUint8("seconds", 8)
		if _secondsErr != nil {
			return nil, errors.Wrap(_secondsErr, "Error parsing 'seconds' field")
		}
		_ = seconds // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (millisecondsOfSecond)
		millisecondsOfSecond, _millisecondsOfSecondErr := readBuffer.ReadUint16("millisecondsOfSecond", 12)
		if _millisecondsOfSecondErr != nil {
			return nil, errors.Wrap(_millisecondsOfSecondErr, "Error parsing 'millisecondsOfSecond' field")
		}
		_ = millisecondsOfSecond // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (dayOfWeek)
		dayOfWeek, _dayOfWeekErr := readBuffer.ReadUint8("dayOfWeek", 4)
		if _dayOfWeekErr != nil {
			return nil, errors.Wrap(_dayOfWeekErr, "Error parsing 'dayOfWeek' field")
		}
		_ = dayOfWeek // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
	case dataProtocolId == "IEC61131_DATE_AND_LTIME": // DATE_AND_LTIME
		// Simple Field (nanosecondsSinceEpoch)
		nanosecondsSinceEpoch, _nanosecondsSinceEpochErr := readBuffer.ReadUint64("nanosecondsSinceEpoch", 64)
		if _nanosecondsSinceEpochErr != nil {
			return nil, errors.Wrap(_nanosecondsSinceEpochErr, "Error parsing 'nanosecondsSinceEpoch' field")
		}
		_ = nanosecondsSinceEpoch // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDATA_AND_LTIMEFromNanosecondsSinceEpoch(nanosecondsSinceEpoch), nil
	case dataProtocolId == "IEC61131_DTL": // DATE_AND_LTIME
		// Simple Field (year)
		year, _yearErr := readBuffer.ReadUint16("year", 16)
		if _yearErr != nil {
			return nil, errors.Wrap(_yearErr, "Error parsing 'year' field")
		}
		_ = year // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (month)
		month, _monthErr := readBuffer.ReadUint8("month", 8)
		if _monthErr != nil {
			return nil, errors.Wrap(_monthErr, "Error parsing 'month' field")
		}
		_ = month // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (day)
		day, _dayErr := readBuffer.ReadUint8("day", 8)
		if _dayErr != nil {
			return nil, errors.Wrap(_dayErr, "Error parsing 'day' field")
		}
		_ = day // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (dayOfWeek)
		dayOfWeek, _dayOfWeekErr := readBuffer.ReadUint8("dayOfWeek", 8)
		if _dayOfWeekErr != nil {
			return nil, errors.Wrap(_dayOfWeekErr, "Error parsing 'dayOfWeek' field")
		}
		_ = dayOfWeek // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (hour)
		hour, _hourErr := readBuffer.ReadUint8("hour", 8)
		if _hourErr != nil {
			return nil, errors.Wrap(_hourErr, "Error parsing 'hour' field")
		}
		_ = hour // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (minutes)
		minutes, _minutesErr := readBuffer.ReadUint8("minutes", 8)
		if _minutesErr != nil {
			return nil, errors.Wrap(_minutesErr, "Error parsing 'minutes' field")
		}
		_ = minutes // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (seconds)
		seconds, _secondsErr := readBuffer.ReadUint8("seconds", 8)
		if _secondsErr != nil {
			return nil, errors.Wrap(_secondsErr, "Error parsing 'seconds' field")
		}
		_ = seconds // TODO: temporary till we fix TIME stuff in golang (see above in the template)

		// Simple Field (nannosecondsOfSecond)
		nannosecondsOfSecond, _nannosecondsOfSecondErr := readBuffer.ReadUint32("nannosecondsOfSecond", 32)
		if _nannosecondsOfSecondErr != nil {
			return nil, errors.Wrap(_nannosecondsOfSecondErr, "Error parsing 'nannosecondsOfSecond' field")
		}
		_ = nannosecondsOfSecond // TODO: temporary till we fix TIME stuff in golang (see above in the template)
		readBuffer.CloseContext("DataItem")
	}
	// TODO: add more info which type it is actually
	return nil, errors.New("unsupported type")
}

func DataItemSerialize(value api.PlcValue, dataProtocolId string, controllerType ControllerType, stringLength int32, stringEncoding string) ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := DataItemSerializeWithWriteBuffer(context.Background(), wb, value, dataProtocolId, controllerType, stringLength, stringEncoding); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func DataItemSerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer, value api.PlcValue, dataProtocolId string, controllerType ControllerType, stringLength int32, stringEncoding string) error {
	log := zerolog.Ctx(ctx)
	_ = log
	m := struct {
		DataProtocolId string
		ControllerType ControllerType
		StringLength   int32
		StringEncoding string
	}{
		DataProtocolId: dataProtocolId,
		ControllerType: controllerType,
		StringLength:   stringLength,
		StringEncoding: stringEncoding,
	}
	_ = m
	writeBuffer.PushContext("DataItem")
	switch {
	case dataProtocolId == "IEC61131_BOOL": // BOOL
		// Reserved Field (Just skip the bytes)
		if _err := writeBuffer.WriteUint8("reserved", 7, uint8(uint8(0x00))); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}

		// Simple Field (value)
		if _err := writeBuffer.WriteBit("value", value.GetBool()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_BYTE": // BYTE
		// Simple Field (value)
		if _err := writeBuffer.WriteUint8("value", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_WORD": // WORD
		// Simple Field (value)
		if _err := writeBuffer.WriteUint16("value", 16, uint16(value.GetUint16())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_DWORD": // DWORD
		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, uint32(value.GetUint32())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_LWORD": // LWORD
		// Simple Field (value)
		if _err := writeBuffer.WriteUint64("value", 64, uint64(value.GetUint64())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_SINT": // SINT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt8("value", 8, int8(value.GetInt8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_USINT": // USINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint8("value", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_INT": // INT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt16("value", 16, int16(value.GetInt16())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_UINT": // UINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint16("value", 16, uint16(value.GetUint16())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_DINT": // DINT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt32("value", 32, int32(value.GetInt32())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_UDINT": // UDINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, uint32(value.GetUint32())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_LINT": // LINT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt64("value", 64, int64(value.GetInt64())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_ULINT": // ULINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint64("value", 64, uint64(value.GetUint64())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_REAL": // REAL
		// Simple Field (value)
		if _err := writeBuffer.WriteFloat32("value", 32, value.GetFloat32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_LREAL": // LREAL
		// Simple Field (value)
		if _err := writeBuffer.WriteFloat64("value", 64, value.GetFloat64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_CHAR": // CHAR
		// Manual Field (value)
		_valueErr := SerializeS7Char(ctx, writeBuffer, value, "UTF-8", m.StringEncoding)
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_WCHAR": // CHAR
		// Manual Field (value)
		_valueErr := SerializeS7Char(ctx, writeBuffer, value, "UTF-16", m.StringEncoding)
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_STRING": // STRING
		// Manual Field (value)
		_valueErr := SerializeS7String(ctx, writeBuffer, value, stringLength, "UTF-8", m.StringEncoding)
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_WSTRING": // STRING
		// Manual Field (value)
		_valueErr := SerializeS7String(ctx, writeBuffer, value, stringLength, "UTF-16", m.StringEncoding)
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_TIME": // TIME
		// Simple Field (milliseconds)
		if _err := writeBuffer.WriteUint32("milliseconds", 32, uint32(value.(values.PlcTIME).GetMilliseconds())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'milliseconds' field")
		}
	case dataProtocolId == "S7_S5TIME": // TIME
		// Manual Field (milliseconds)
		_millisecondsErr := SerializeS5Time(ctx, writeBuffer, value)
		if _millisecondsErr != nil {
			return errors.Wrap(_millisecondsErr, "Error serializing 'milliseconds' field")
		}
	case dataProtocolId == "IEC61131_LTIME": // LTIME
		// Simple Field (nanoseconds)
		if _err := writeBuffer.WriteUint64("nanoseconds", 64, uint64(value.(values.PlcLTIME).GetNanoseconds())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'nanoseconds' field")
		}
	case dataProtocolId == "IEC61131_DATE": // DATE
		// Manual Field (daysSinceEpoch)
		_daysSinceEpochErr := SerializeTiaDate(ctx, writeBuffer, value)
		if _daysSinceEpochErr != nil {
			return errors.Wrap(_daysSinceEpochErr, "Error serializing 'daysSinceEpoch' field")
		}
	case dataProtocolId == "IEC61131_TIME_OF_DAY": // TIME_OF_DAY
		// Simple Field (millisecondsSinceMidnight)
		if _err := writeBuffer.WriteUint32("millisecondsSinceMidnight", 32, uint32(value.(values.PlcTIME_OF_DAY).GetMillisecondsSinceMidnight())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'millisecondsSinceMidnight' field")
		}
	case dataProtocolId == "IEC61131_LTIME_OF_DAY": // LTIME_OF_DAY
		// Simple Field (nanosecondsSinceMidnight)
		if _err := writeBuffer.WriteUint64("nanosecondsSinceMidnight", 64, uint64(value.(values.PlcLTIME_OF_DAY).GetNanosecondsSinceMidnight())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'nanosecondsSinceMidnight' field")
		}
	case dataProtocolId == "IEC61131_DATE_AND_TIME": // DATE_AND_TIME
		// Manual Field (year)
		_yearErr := SerializeSiemensYear(ctx, writeBuffer, value)
		if _yearErr != nil {
			return errors.Wrap(_yearErr, "Error serializing 'year' field")
		}

		// Simple Field (month)
		if _err := writeBuffer.WriteUint8("month", 8, uint8(value.(values.PlcDATE_AND_TIME).GetMonth())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'month' field")
		}

		// Simple Field (day)
		if _err := writeBuffer.WriteUint8("day", 8, uint8(value.(values.PlcDATE_AND_TIME).GetDay())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'day' field")
		}

		// Simple Field (hour)
		if _err := writeBuffer.WriteUint8("hour", 8, uint8(value.(values.PlcDATE_AND_TIME).GetHour())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'hour' field")
		}

		// Simple Field (minutes)
		if _err := writeBuffer.WriteUint8("minutes", 8, uint8(value.(values.PlcDATE_AND_TIME).GetMinutes())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'minutes' field")
		}

		// Simple Field (seconds)
		if _err := writeBuffer.WriteUint8("seconds", 8, uint8(value.(values.PlcDATE_AND_TIME).GetSeconds())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'seconds' field")
		}

		// Simple Field (millisecondsOfSecond)
		if _err := writeBuffer.WriteUint16("millisecondsOfSecond", 12, uint16(value.(values.PlcDATE_AND_TIME).GetMillisecondsOfSecond())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'millisecondsOfSecond' field")
		}

		// Simple Field (dayOfWeek)
		if _err := writeBuffer.WriteUint8("dayOfWeek", 4, uint8(value.(values.PlcDATE_AND_TIME).GetDayOfWeek())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'dayOfWeek' field")
		}
	case dataProtocolId == "IEC61131_DATE_AND_LTIME": // DATE_AND_LTIME
		// Simple Field (nanosecondsSinceEpoch)
		if _err := writeBuffer.WriteUint64("nanosecondsSinceEpoch", 64, uint64(value.GetUint64())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'nanosecondsSinceEpoch' field")
		}
	case dataProtocolId == "IEC61131_DTL": // DATE_AND_LTIME
		// Simple Field (year)
		if _err := writeBuffer.WriteUint16("year", 16, uint16(value.GetUint16())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'year' field")
		}

		// Simple Field (month)
		if _err := writeBuffer.WriteUint8("month", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'month' field")
		}

		// Simple Field (day)
		if _err := writeBuffer.WriteUint8("day", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'day' field")
		}

		// Simple Field (dayOfWeek)
		if _err := writeBuffer.WriteUint8("dayOfWeek", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'dayOfWeek' field")
		}

		// Simple Field (hour)
		if _err := writeBuffer.WriteUint8("hour", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'hour' field")
		}

		// Simple Field (minutes)
		if _err := writeBuffer.WriteUint8("minutes", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'minutes' field")
		}

		// Simple Field (seconds)
		if _err := writeBuffer.WriteUint8("seconds", 8, uint8(value.GetUint8())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'seconds' field")
		}

		// Simple Field (nannosecondsOfSecond)
		if _err := writeBuffer.WriteUint32("nannosecondsOfSecond", 32, uint32(value.GetUint32())); _err != nil {
			return errors.Wrap(_err, "Error serializing 'nannosecondsOfSecond' field")
		}
	default:
		// TODO: add more info which type it is actually
		return errors.New("unsupported type")
	}
	writeBuffer.PopContext("DataItem")
	return nil
}
