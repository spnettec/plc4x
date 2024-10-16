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
package org.apache.plc4x.java.plc4x.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.math.BigInteger;
import java.time.*;
import java.time.temporal.ChronoField;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;
import org.apache.plc4x.java.spi.values.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

// Code generated by code-generation. DO NOT EDIT.

public class Plc4xValue {

  private static final Logger LOGGER = LoggerFactory.getLogger(Plc4xValue.class);

  public static PlcValue staticParse(ReadBuffer readBuffer, Plc4xValueType valueType)
      throws ParseException {
    if (EvaluationHelper.equals(valueType, Plc4xValueType.BOOL)) { // BOOL
      Byte reservedField0 =
          readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

      boolean value = readSimpleField("value", readBoolean(readBuffer));
      return new PlcBOOL(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.BYTE)) { // BYTE
      short value = readSimpleField("value", readUnsignedShort(readBuffer, 8));
      return new PlcBYTE(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.WORD)) { // WORD
      int value = readSimpleField("value", readUnsignedInt(readBuffer, 16));
      return new PlcWORD(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DWORD)) { // DWORD
      long value = readSimpleField("value", readUnsignedLong(readBuffer, 32));
      return new PlcDWORD(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LWORD)) { // LWORD
      BigInteger value = readSimpleField("value", readUnsignedBigInteger(readBuffer, 64));
      return new PlcLWORD(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.USINT)) { // USINT
      short value = readSimpleField("value", readUnsignedShort(readBuffer, 8));
      return new PlcUSINT(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.UINT)) { // UINT
      int value = readSimpleField("value", readUnsignedInt(readBuffer, 16));
      return new PlcUINT(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.UDINT)) { // UDINT
      long value = readSimpleField("value", readUnsignedLong(readBuffer, 32));
      return new PlcUDINT(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.ULINT)) { // ULINT
      BigInteger value = readSimpleField("value", readUnsignedBigInteger(readBuffer, 64));
      return new PlcULINT(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.SINT)) { // SINT
      byte value = readSimpleField("value", readSignedByte(readBuffer, 8));
      return new PlcSINT(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.INT)) { // INT
      short value = readSimpleField("value", readSignedShort(readBuffer, 16));
      return new PlcINT(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DINT)) { // DINT
      int value = readSimpleField("value", readSignedInt(readBuffer, 32));
      return new PlcDINT(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LINT)) { // LINT
      long value = readSimpleField("value", readSignedLong(readBuffer, 64));
      return new PlcLINT(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.REAL)) { // REAL
      float value = readSimpleField("value", readFloat(readBuffer, 32));
      return new PlcREAL(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LREAL)) { // LREAL
      double value = readSimpleField("value", readDouble(readBuffer, 64));
      return new PlcLREAL(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.CHAR)) { // STRING
      String value = readSimpleField("value", readString(readBuffer, 8));
      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.WCHAR)) { // STRING
      String value =
          readSimpleField("value", readString(readBuffer, 16), WithOption.WithEncoding("UTF-16"));
      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.STRING)) { // STRING
      String value =
          readManualField(
              "value",
              readBuffer,
              () ->
                  (String)
                      (org.apache.plc4x.java.plc4x.readwrite.utils.StaticHelper.parseString(
                          readBuffer, "UTF-8")));
      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.WSTRING)) { // STRING
      String value =
          readManualField(
              "value",
              readBuffer,
              () ->
                  (String)
                      (org.apache.plc4x.java.plc4x.readwrite.utils.StaticHelper.parseString(
                          readBuffer, "UTF-16")),
              WithOption.WithEncoding("UTF-16"));
      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.TIME)) { // TIME
      long milliseconds = readSimpleField("milliseconds", readUnsignedLong(readBuffer, 32));
      return PlcTIME.ofMilliseconds(milliseconds);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LTIME)) { // LTIME
      BigInteger nanoseconds =
          readSimpleField("nanoseconds", readUnsignedBigInteger(readBuffer, 64));
      return PlcLTIME.ofNanoseconds(nanoseconds);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DATE)) { // DATE
      long secondsSinceEpoch =
          readSimpleField("secondsSinceEpoch", readUnsignedLong(readBuffer, 32));
      return PlcDATE.ofSecondsSinceEpoch(secondsSinceEpoch);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LDATE)) { // LDATE
      BigInteger nanosecondsSinceEpoch =
          readSimpleField("nanosecondsSinceEpoch", readUnsignedBigInteger(readBuffer, 64));
      return PlcLDATE.ofNanosecondsSinceEpoch(nanosecondsSinceEpoch);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.TIME_OF_DAY)) { // TIME_OF_DAY
      long millisecondsSinceMidnight =
          readSimpleField("millisecondsSinceMidnight", readUnsignedLong(readBuffer, 32));
      return PlcTIME_OF_DAY.ofMillisecondsSinceMidnight(millisecondsSinceMidnight);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LTIME_OF_DAY)) { // LTIME_OF_DAY
      BigInteger nanosecondsSinceMidnight =
          readSimpleField("nanosecondsSinceMidnight", readUnsignedBigInteger(readBuffer, 64));
      return PlcLTIME_OF_DAY.ofNanosecondsSinceMidnight(nanosecondsSinceMidnight);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DATE_AND_TIME)) { // DATE_AND_TIME
      long secondsSinceEpoch =
          readSimpleField("secondsSinceEpoch", readUnsignedLong(readBuffer, 32));
      return PlcDATE_AND_TIME.ofSecondsSinceEpoch(secondsSinceEpoch);
    } else if (EvaluationHelper.equals(
        valueType, Plc4xValueType.DATE_AND_LTIME)) { // DATE_AND_LTIME
      BigInteger nanosecondsSinceEpoch =
          readSimpleField("nanosecondsSinceEpoch", readUnsignedBigInteger(readBuffer, 64));
      return PlcDATE_AND_LTIME.ofNanosecondsSinceEpoch(nanosecondsSinceEpoch);
    } else if (EvaluationHelper.equals(
        valueType, Plc4xValueType.LDATE_AND_TIME)) { // LDATE_AND_TIME
      BigInteger nanosecondsSinceEpoch =
          readSimpleField("nanosecondsSinceEpoch", readUnsignedBigInteger(readBuffer, 64));
      return PlcLDATE_AND_TIME.ofNanosecondsSinceEpoch(nanosecondsSinceEpoch);
    }
    return null;
  }

  public static int getLengthInBytes(PlcValue _value, Plc4xValueType valueType) {
    return (int) Math.ceil((float) getLengthInBits(_value, valueType) / 8.0);
  }

  public static int getLengthInBits(PlcValue _value, Plc4xValueType valueType) {
    int lengthInBits = 0;
    if (EvaluationHelper.equals(valueType, Plc4xValueType.BOOL)) { // BOOL
      // Reserved Field (reserved)
      lengthInBits += 7;

      // Simple field (value)
      lengthInBits += 1;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.BYTE)) { // BYTE
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.WORD)) { // WORD
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DWORD)) { // DWORD
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LWORD)) { // LWORD
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.USINT)) { // USINT
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.UINT)) { // UINT
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.UDINT)) { // UDINT
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.ULINT)) { // ULINT
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.SINT)) { // SINT
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.INT)) { // INT
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DINT)) { // DINT
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LINT)) { // LINT
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.REAL)) { // REAL
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LREAL)) { // LREAL
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.CHAR)) { // STRING
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.WCHAR)) { // STRING
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.STRING)) { // STRING
      // Manual Field (value)
      lengthInBits += (((STR_LEN(_value)) + (1))) * (8);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.WSTRING)) { // STRING
      // Manual Field (value)
      lengthInBits += (((STR_LEN(_value)) + (1))) * (16);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.TIME)) { // TIME
      // Simple field (milliseconds)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LTIME)) { // LTIME
      // Simple field (nanoseconds)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DATE)) { // DATE
      // Simple field (secondsSinceEpoch)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LDATE)) { // LDATE
      // Simple field (nanosecondsSinceEpoch)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.TIME_OF_DAY)) { // TIME_OF_DAY
      // Simple field (millisecondsSinceMidnight)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LTIME_OF_DAY)) { // LTIME_OF_DAY
      // Simple field (nanosecondsSinceMidnight)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DATE_AND_TIME)) { // DATE_AND_TIME
      // Simple field (secondsSinceEpoch)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(
        valueType, Plc4xValueType.DATE_AND_LTIME)) { // DATE_AND_LTIME
      // Simple field (nanosecondsSinceEpoch)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(
        valueType, Plc4xValueType.LDATE_AND_TIME)) { // LDATE_AND_TIME
      // Simple field (nanosecondsSinceEpoch)
      lengthInBits += 64;
    }

    return lengthInBits;
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer, PlcValue _value, Plc4xValueType valueType)
      throws SerializationException {
    staticSerialize(writeBuffer, _value, valueType, ByteOrder.BIG_ENDIAN);
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer, PlcValue _value, Plc4xValueType valueType, ByteOrder byteOrder)
      throws SerializationException {
    if (EvaluationHelper.equals(valueType, Plc4xValueType.BOOL)) { // BOOL
      // Reserved Field (reserved)
      writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

      // Simple Field (value)
      writeSimpleField("value", (boolean) _value.getBoolean(), writeBoolean(writeBuffer));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.BYTE)) { // BYTE
      // Simple Field (value)
      writeSimpleField("value", (short) _value.getShort(), writeUnsignedShort(writeBuffer, 8));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.WORD)) { // WORD
      // Simple Field (value)
      writeSimpleField("value", (int) _value.getInteger(), writeUnsignedInt(writeBuffer, 16));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DWORD)) { // DWORD
      // Simple Field (value)
      writeSimpleField("value", (long) _value.getLong(), writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LWORD)) { // LWORD
      // Simple Field (value)
      writeSimpleField(
          "value", (BigInteger) _value.getBigInteger(), writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.USINT)) { // USINT
      // Simple Field (value)
      writeSimpleField("value", (short) _value.getShort(), writeUnsignedShort(writeBuffer, 8));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.UINT)) { // UINT
      // Simple Field (value)
      writeSimpleField("value", (int) _value.getInteger(), writeUnsignedInt(writeBuffer, 16));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.UDINT)) { // UDINT
      // Simple Field (value)
      writeSimpleField("value", (long) _value.getLong(), writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.ULINT)) { // ULINT
      // Simple Field (value)
      writeSimpleField(
          "value", (BigInteger) _value.getBigInteger(), writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.SINT)) { // SINT
      // Simple Field (value)
      writeSimpleField("value", (byte) _value.getByte(), writeSignedByte(writeBuffer, 8));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.INT)) { // INT
      // Simple Field (value)
      writeSimpleField("value", (short) _value.getShort(), writeSignedShort(writeBuffer, 16));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DINT)) { // DINT
      // Simple Field (value)
      writeSimpleField("value", (int) _value.getInteger(), writeSignedInt(writeBuffer, 32));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LINT)) { // LINT
      // Simple Field (value)
      writeSimpleField("value", (long) _value.getLong(), writeSignedLong(writeBuffer, 64));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.REAL)) { // REAL
      // Simple Field (value)
      writeSimpleField("value", (float) _value.getFloat(), writeFloat(writeBuffer, 32));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LREAL)) { // LREAL
      // Simple Field (value)
      writeSimpleField("value", (double) _value.getDouble(), writeDouble(writeBuffer, 64));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.CHAR)) { // STRING
      // Simple Field (value)
      writeSimpleField("value", (String) _value.getString(), writeString(writeBuffer, 8));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.WCHAR)) { // STRING
      // Simple Field (value)
      writeSimpleField(
          "value",
          (String) _value.getString(),
          writeString(writeBuffer, 16),
          WithOption.WithEncoding("UTF-16"));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.STRING)) { // STRING
      // Manual Field (value)
      writeManualField(
          "value",
          () ->
              org.apache.plc4x.java.plc4x.readwrite.utils.StaticHelper.serializeString(
                  writeBuffer, _value, "UTF-8"),
          writeBuffer);
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.WSTRING)) { // STRING
      // Manual Field (value)
      writeManualField(
          "value",
          () ->
              org.apache.plc4x.java.plc4x.readwrite.utils.StaticHelper.serializeString(
                  writeBuffer, _value, "UTF-16"),
          writeBuffer,
          WithOption.WithEncoding("UTF-16"));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.TIME)) { // TIME
      // Simple Field (milliseconds)
      writeSimpleField(
          "milliseconds",
          (long) _value.getDuration().toMillis(),
          writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LTIME)) { // LTIME
      // Simple Field (nanoseconds)
      writeSimpleField(
          "nanoseconds",
          (BigInteger) BigInteger.valueOf(_value.getDuration().toNanos()),
          writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DATE)) { // DATE
      // Simple Field (secondsSinceEpoch)
      writeSimpleField(
          "secondsSinceEpoch",
          (long) _value.getDateTime().toEpochSecond(ZoneOffset.UTC),
          writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LDATE)) { // LDATE
      // Simple Field (nanosecondsSinceEpoch)
      writeSimpleField(
          "nanosecondsSinceEpoch",
          (BigInteger)
              BigInteger.valueOf(_value.getDateTime().toEpochSecond(ZoneOffset.UTC))
                  .multiply(BigInteger.valueOf(1000000000))
                  .add(BigInteger.valueOf(_value.getDateTime().getNano())),
          writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.TIME_OF_DAY)) { // TIME_OF_DAY
      // Simple Field (millisecondsSinceMidnight)
      writeSimpleField(
          "millisecondsSinceMidnight",
          (long) _value.getTime().getLong(ChronoField.MILLI_OF_DAY),
          writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.LTIME_OF_DAY)) { // LTIME_OF_DAY
      // Simple Field (nanosecondsSinceMidnight)
      writeSimpleField(
          "nanosecondsSinceMidnight",
          (BigInteger) BigInteger.valueOf(_value.getTime().getLong(ChronoField.NANO_OF_DAY)),
          writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(valueType, Plc4xValueType.DATE_AND_TIME)) { // DATE_AND_TIME
      // Simple Field (secondsSinceEpoch)
      writeSimpleField(
          "secondsSinceEpoch",
          (long) _value.getDateTime().toEpochSecond(ZoneOffset.UTC),
          writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(
        valueType, Plc4xValueType.DATE_AND_LTIME)) { // DATE_AND_LTIME
      // Simple Field (nanosecondsSinceEpoch)
      writeSimpleField(
          "nanosecondsSinceEpoch",
          (BigInteger)
              BigInteger.valueOf(_value.getDateTime().toEpochSecond(ZoneOffset.UTC))
                  .multiply(BigInteger.valueOf(1000000000))
                  .add(BigInteger.valueOf(_value.getDateTime().getNano())),
          writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(
        valueType, Plc4xValueType.LDATE_AND_TIME)) { // LDATE_AND_TIME
      // Simple Field (nanosecondsSinceEpoch)
      writeSimpleField(
          "nanosecondsSinceEpoch",
          (BigInteger)
              BigInteger.valueOf(_value.getDateTime().toEpochSecond(ZoneOffset.UTC))
                  .multiply(BigInteger.valueOf(1000000000))
                  .add(BigInteger.valueOf(_value.getDateTime().getNano())),
          writeUnsignedBigInteger(writeBuffer, 64));
    }
  }
}
