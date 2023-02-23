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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/apache/plc4x/plc4go/spi/values"
	"github.com/pkg/errors"
	api "github.com/apache/plc4x/plc4go/pkg/api/values"
)

	// Code generated by code-generation. DO NOT EDIT.
	
func DataItemParse(ctx context.Context, theBytes []byte, plcValueType PlcValueType, stringLength int32, stringEncoding string) (api.PlcValue, error) {
	return DataItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), plcValueType, stringLength, stringEncoding)
}

func DataItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, plcValueType PlcValueType, stringLength int32, stringEncoding string) (api.PlcValue, error) {
	readBuffer.PullContext("DataItem")
	switch {
case plcValueType == PlcValueType_BOOL : // BOOL
			// Reserved Field (Just skip the bytes)
			if _, _err := readBuffer.ReadUint8("reserved", 7); _err != nil {
				return nil, errors.Wrap(_err, "Error parsing reserved field")
			}

			// Simple Field (value)
			value, _valueErr := readBuffer.ReadBit("value")
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcBOOL(value), nil
case plcValueType == PlcValueType_BYTE : // BYTE
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadUint8("value", 8)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcBYTE(value), nil
case plcValueType == PlcValueType_WORD : // WORD
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadUint16("value", 16)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcWORD(value), nil
case plcValueType == PlcValueType_DWORD : // DWORD
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadUint32("value", 32)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcDWORD(value), nil
case plcValueType == PlcValueType_LWORD : // LWORD
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadUint64("value", 64)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcLWORD(value), nil
case plcValueType == PlcValueType_SINT : // SINT
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadInt8("value", 8)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcSINT(value), nil
case plcValueType == PlcValueType_USINT : // USINT
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadUint8("value", 8)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcUSINT(value), nil
case plcValueType == PlcValueType_INT : // INT
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadInt16("value", 16)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcINT(value), nil
case plcValueType == PlcValueType_UINT : // UINT
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadUint16("value", 16)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcUINT(value), nil
case plcValueType == PlcValueType_DINT : // DINT
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadInt32("value", 32)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcDINT(value), nil
case plcValueType == PlcValueType_UDINT : // UDINT
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadUint32("value", 32)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcUDINT(value), nil
case plcValueType == PlcValueType_LINT : // LINT
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadInt64("value", 64)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcLINT(value), nil
case plcValueType == PlcValueType_ULINT : // ULINT
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadUint64("value", 64)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcULINT(value), nil
case plcValueType == PlcValueType_REAL : // REAL
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadFloat32("value", 32)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcREAL(value), nil
case plcValueType == PlcValueType_LREAL : // LREAL
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadFloat64("value", 64)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcLREAL(value), nil
case plcValueType == PlcValueType_CHAR : // CHAR
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadString("value", uint32(8), "UTF-8")
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcCHAR(value), nil
case plcValueType == PlcValueType_WCHAR : // WCHAR
			// Simple Field (value)
			value, _valueErr := readBuffer.ReadString("value", uint32(16), "UTF-16LE")
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcWCHAR(value), nil
case plcValueType == PlcValueType_STRING : // STRING
			// Manual Field (value)
			value, _valueErr := ParseAmsString(readBuffer, stringLength, "UTF-8", stringEncoding)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcSTRING(value), nil
case plcValueType == PlcValueType_WSTRING : // STRING
			// Manual Field (value)
			value, _valueErr := ParseAmsString(readBuffer, stringLength, "UTF-16", stringEncoding)
			if _valueErr != nil {
				return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcSTRING(value), nil
case plcValueType == PlcValueType_TIME : // TIME
			// Simple Field (milliseconds)
			milliseconds, _millisecondsErr := readBuffer.ReadUint32("milliseconds", 32)
			if _millisecondsErr != nil {
				return nil, errors.Wrap(_millisecondsErr, "Error parsing 'milliseconds' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcTIMEFromMilliseconds(milliseconds), nil
case plcValueType == PlcValueType_LTIME : // LTIME
			// Simple Field (nanoseconds)
			nanoseconds, _nanosecondsErr := readBuffer.ReadUint64("nanoseconds", 64)
			if _nanosecondsErr != nil {
				return nil, errors.Wrap(_nanosecondsErr, "Error parsing 'nanoseconds' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcLTIMEFromNanoseconds(nanoseconds), nil
case plcValueType == PlcValueType_DATE : // DATE
			// Simple Field (secondsSinceEpoch)
			secondsSinceEpoch, _secondsSinceEpochErr := readBuffer.ReadUint32("secondsSinceEpoch", 32)
			if _secondsSinceEpochErr != nil {
				return nil, errors.Wrap(_secondsSinceEpochErr, "Error parsing 'secondsSinceEpoch' field")
			}
 			readBuffer.CloseContext("DataItem")
			return values.NewPlcDATEFromSecondsSinceEpoch(uint32(secondsSinceEpoch)), nil
case plcValueType == PlcValueType_LDATE : // LDATE
			// Simple Field (nanosecondsSinceEpoch)
			nanosecondsSinceEpoch, _nanosecondsSinceEpochErr := readBuffer.ReadUint64("nanosecondsSinceEpoch", 64)
			if _nanosecondsSinceEpochErr != nil {
				return nil, errors.Wrap(_nanosecondsSinceEpochErr, "Error parsing 'nanosecondsSinceEpoch' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcLDATEFromNanosecondsSinceEpoch(uint64(nanosecondsSinceEpoch)), nil
case plcValueType == PlcValueType_TIME_OF_DAY : // TIME_OF_DAY
			// Simple Field (millisecondsSinceMidnight)
			millisecondsSinceMidnight, _millisecondsSinceMidnightErr := readBuffer.ReadUint32("millisecondsSinceMidnight", 32)
			if _millisecondsSinceMidnightErr != nil {
				return nil, errors.Wrap(_millisecondsSinceMidnightErr, "Error parsing 'millisecondsSinceMidnight' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcTIME_OF_DAYFromMillisecondsSinceMidnight(millisecondsSinceMidnight), nil
case plcValueType == PlcValueType_LTIME_OF_DAY : // LTIME_OF_DAY
			// Simple Field (nanosecondsSinceMidnight)
			nanosecondsSinceMidnight, _nanosecondsSinceMidnightErr := readBuffer.ReadUint64("nanosecondsSinceMidnight", 64)
			if _nanosecondsSinceMidnightErr != nil {
				return nil, errors.Wrap(_nanosecondsSinceMidnightErr, "Error parsing 'nanosecondsSinceMidnight' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcLTIME_OF_DAYFromNanosecondsSinceMidnight(nanosecondsSinceMidnight), nil
case plcValueType == PlcValueType_DATE_AND_TIME : // DATE_AND_TIME
			// Simple Field (secondsSinceEpoch)
			secondsSinceEpoch, _secondsSinceEpochErr := readBuffer.ReadUint32("secondsSinceEpoch", 32)
			if _secondsSinceEpochErr != nil {
				return nil, errors.Wrap(_secondsSinceEpochErr, "Error parsing 'secondsSinceEpoch' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcDATA_AND_TIMEFromSecondsSinceEpoch(secondsSinceEpoch), nil
case plcValueType == PlcValueType_LDATE_AND_TIME : // LDATE_AND_TIME
			// Simple Field (nanosecondsSinceEpoch)
			nanosecondsSinceEpoch, _nanosecondsSinceEpochErr := readBuffer.ReadUint64("nanosecondsSinceEpoch", 64)
			if _nanosecondsSinceEpochErr != nil {
				return nil, errors.Wrap(_nanosecondsSinceEpochErr, "Error parsing 'nanosecondsSinceEpoch' field")
			}
			readBuffer.CloseContext("DataItem")
			return values.NewPlcLDATE_AND_TIMEFromNanosecondsSinceEpoch(uint64(nanosecondsSinceEpoch)), nil
	}
    // TODO: add more info which type it is actually
	return nil, errors.New("unsupported type")
}

func DataItemSerialize(value api.PlcValue, plcValueType PlcValueType, stringLength int32, stringEncoding string) ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := DataItemSerializeWithWriteBuffer(context.Background(), wb, value, plcValueType, stringLength, stringEncoding); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func DataItemSerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer, value api.PlcValue, plcValueType PlcValueType, stringLength int32, stringEncoding string) error {
	m := struct {
			PlcValueType PlcValueType
			StringLength int32
			StringEncoding string
	}{
			PlcValueType: plcValueType,
			StringLength: stringLength,
			StringEncoding: stringEncoding,
	}
	_ = m
	writeBuffer.PushContext("DataItem")
	switch {
case plcValueType == PlcValueType_BOOL : // BOOL
			// Reserved Field (Just skip the bytes)
			if _err := writeBuffer.WriteUint8("reserved", 7, uint8(0x00)); _err != nil {
				return errors.Wrap(_err, "Error serializing reserved field")
			}

			// Simple Field (value)
			if _err := writeBuffer.WriteBit("value", value.GetBool()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_BYTE : // BYTE
			// Simple Field (value)
			if _err := writeBuffer.WriteUint8("value", 8, value.GetUint8()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_WORD : // WORD
			// Simple Field (value)
			if _err := writeBuffer.WriteUint16("value", 16, value.GetUint16()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_DWORD : // DWORD
			// Simple Field (value)
			if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_LWORD : // LWORD
			// Simple Field (value)
			if _err := writeBuffer.WriteUint64("value", 64, value.GetUint64()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_SINT : // SINT
			// Simple Field (value)
			if _err := writeBuffer.WriteInt8("value", 8, value.GetInt8()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_USINT : // USINT
			// Simple Field (value)
			if _err := writeBuffer.WriteUint8("value", 8, value.GetUint8()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_INT : // INT
			// Simple Field (value)
			if _err := writeBuffer.WriteInt16("value", 16, value.GetInt16()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_UINT : // UINT
			// Simple Field (value)
			if _err := writeBuffer.WriteUint16("value", 16, value.GetUint16()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_DINT : // DINT
			// Simple Field (value)
			if _err := writeBuffer.WriteInt32("value", 32, value.GetInt32()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_UDINT : // UDINT
			// Simple Field (value)
			if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_LINT : // LINT
			// Simple Field (value)
			if _err := writeBuffer.WriteInt64("value", 64, value.GetInt64()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_ULINT : // ULINT
			// Simple Field (value)
			if _err := writeBuffer.WriteUint64("value", 64, value.GetUint64()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_REAL : // REAL
			// Simple Field (value)
			if _err := writeBuffer.WriteFloat32("value", 32, value.GetFloat32()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_LREAL : // LREAL
			// Simple Field (value)
			if _err := writeBuffer.WriteFloat64("value", 64, value.GetFloat64()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_CHAR : // CHAR
			// Simple Field (value)
			if _err := writeBuffer.WriteString("value", uint32(8), "UTF-8", value.GetString()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_WCHAR : // WCHAR
			// Simple Field (value)
			if _err := writeBuffer.WriteString("value", uint32(16), "UTF-16LE", value.GetString()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_STRING : // STRING
			// Manual Field (value)
			_valueErr := SerializeAmsString(writeBuffer, value, stringLength, "UTF-8", m.StringEncoding)
			if _valueErr != nil {
				return errors.Wrap(_valueErr, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_WSTRING : // STRING
			// Manual Field (value)
			_valueErr := SerializeAmsString(writeBuffer, value, stringLength, "UTF-16", m.StringEncoding)
			if _valueErr != nil {
				return errors.Wrap(_valueErr, "Error serializing 'value' field")
			}
case plcValueType == PlcValueType_TIME : // TIME
			// Simple Field (milliseconds)
			if _err := writeBuffer.WriteUint32("milliseconds", 32, value.(values.PlcTIME).GetMilliseconds()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'milliseconds' field")
			}
case plcValueType == PlcValueType_LTIME : // LTIME
			// Simple Field (nanoseconds)
			if _err := writeBuffer.WriteUint64("nanoseconds", 64, value.(values.PlcLTIME).GetNanoseconds()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'nanoseconds' field")
			}
case plcValueType == PlcValueType_DATE : // DATE
			// Simple Field (secondsSinceEpoch)
			if _err := writeBuffer.WriteUint32("secondsSinceEpoch", 32, value.(values.PlcDATE).GetSecondsSinceEpoch()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'secondsSinceEpoch' field")
			}
case plcValueType == PlcValueType_LDATE : // LDATE
			// Simple Field (nanosecondsSinceEpoch)
			if _err := writeBuffer.WriteUint64("nanosecondsSinceEpoch", 64, value.(values.PlcLDATE).GetNanosecondsSinceEpoch()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'nanosecondsSinceEpoch' field")
			}
case plcValueType == PlcValueType_TIME_OF_DAY : // TIME_OF_DAY
			// Simple Field (millisecondsSinceMidnight)
			if _err := writeBuffer.WriteUint32("millisecondsSinceMidnight", 32, value.(values.PlcTIME_OF_DAY).GetMillisecondsSinceMidnight()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'millisecondsSinceMidnight' field")
			}
case plcValueType == PlcValueType_LTIME_OF_DAY : // LTIME_OF_DAY
			// Simple Field (nanosecondsSinceMidnight)
			if _err := writeBuffer.WriteUint64("nanosecondsSinceMidnight", 64, value.(values.PlcLTIME_OF_DAY).GetNanosecondsSinceMidnight()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'nanosecondsSinceMidnight' field")
			}
case plcValueType == PlcValueType_DATE_AND_TIME : // DATE_AND_TIME
			// Simple Field (secondsSinceEpoch)
			if _err := writeBuffer.WriteUint32("secondsSinceEpoch", 32, value.(values.PlcDATE_AND_TIME).GetSecondsSinceEpoch()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'secondsSinceEpoch' field")
			}
case plcValueType == PlcValueType_LDATE_AND_TIME : // LDATE_AND_TIME
			// Simple Field (nanosecondsSinceEpoch)
			if _err := writeBuffer.WriteUint64("nanosecondsSinceEpoch", 64, value.(values.PlcLDATE_AND_TIME).GetNanosecondsSinceEpoch()); _err != nil {
				return errors.Wrap(_err, "Error serializing 'nanosecondsSinceEpoch' field")
			}
		default:
            // TODO: add more info which type it is actually
			return errors.New("unsupported type")
	}
	writeBuffer.PopContext("DataItem")
	return nil
}


