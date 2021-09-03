/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.s7.utils;

import org.apache.commons.lang3.NotImplementedException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.apache.plc4x.java.spi.values.PlcDATE_AND_TIME;
import org.apache.plc4x.java.spi.values.PlcList;
import org.apache.plc4x.java.spi.values.PlcWCHAR;

import java.io.BufferedInputStream;
import java.io.ByteArrayInputStream;
import java.io.UnsupportedEncodingException;
import java.nio.ByteBuffer;
import java.nio.charset.CharacterCodingException;
import java.nio.charset.Charset;
import java.nio.charset.CharsetDecoder;
import java.nio.charset.StandardCharsets;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.time.temporal.ChronoUnit;

public class StaticHelper {
    private static final String[] DEFAULTCHARSETS = {"ASCII", "UTF-8", "GBK", "GB2312", "BIG5", "GB18030"};

    public static Charset detectCharset(byte[] bytes) {

        Charset charset = null;

        for (String charsetName : DEFAULTCHARSETS) {
            charset = detectCharset(bytes, Charset.forName(charsetName), 0, bytes.length);
            if (charset != null) {
                break;
            }
        }

        return charset;
    }

    private static Charset detectCharset(byte[] bytes, Charset charset, int offset, int length) {
        try {
            BufferedInputStream input = new BufferedInputStream(new ByteArrayInputStream(bytes, offset, length));

            CharsetDecoder decoder = charset.newDecoder();
            decoder.reset();

            byte[] buffer = new byte[512];
            boolean identified = false;
            while (input.read(buffer) != -1 && !identified) {
                identified = identify(buffer, decoder);
            }

            input.close();

            if (identified) {
                return charset;
            } else {
                return null;
            }

        } catch (Exception e) {
            return null;
        }
    }

    private static boolean identify(byte[] bytes, CharsetDecoder decoder) {
        try {
            decoder.decode(ByteBuffer.wrap(bytes));
        } catch (CharacterCodingException e) {
            return false;
        }
        return true;
    }

    public static Charset getEncoding(String str) {
        if (str == null || str.trim().length() < 1) {
            return null;
        }
        for (String encode : DEFAULTCHARSETS) {
            try {
                Charset charset = Charset.forName(encode);
                if (str.equals(new String(str.getBytes(charset), charset))) {
                    return charset;
                }
            } catch (Exception er) {
            }
        }
        return null;
    }

    public static LocalTime parseTiaTime(ReadBuffer io) {
        try {
            int millisSinceMidnight = io.readInt(32);
            return LocalTime.now().withHour(0).withMinute(0).withSecond(0).withNano(0).plus(
                millisSinceMidnight, ChronoUnit.MILLIS);
        } catch (ParseException e) {
            return null;
        }
    }

    public static void serializeTiaTime(WriteBuffer io, PlcValue value) {
        throw new NotImplementedException("Serializing TIME not implemented");
    }

    public static LocalTime parseS5Time(ReadBuffer io) {
        try {
            int stuff = io.readInt(16);
            // TODO: Implement this correctly.
            throw new NotImplementedException("S5TIME not implemented");
        } catch (ParseException e) {
            return null;
        }
    }

    public static Short parseBcDYear(ReadBuffer io) {
        try {
            int num = Integer.parseInt(Integer.toHexString(io.readByte() & 0xFF));
            if (num < 90) {
                num += 2000;
            } else {
                num += 1900;
            }
            return (short) num;
        } catch (ParseException e) {
            return 0;
        }
    }

    public static Integer parseBcDNanos(ReadBuffer io) {
        try {
            return Integer.parseInt(Integer.toHexString(io.readInt(16)));
        } catch (ParseException e) {
            return 0;
        }
    }

    public static Short parseBcD(ReadBuffer io) {
        try {
            return Short.parseShort(Integer.toHexString(io.readByte() & 0xFF));
        } catch (ParseException e) {
            return 0;
        }
    }

    public static void serializeBcdNanos(WriteBuffer io, PlcValue value) {
        PlcDATE_AND_TIME dt = (PlcDATE_AND_TIME) value;
        LocalDateTime ldt = dt.getDateTime();
        try {
            int hexNano = Integer.parseInt("" + ldt.getNano(), 16);
            io.writeInt(16, hexNano);
        } catch (ParseException e) {
            throw new PlcRuntimeException("Parse DateTime error " + e.getMessage());
        }
    }

    public static void serializeBcd(WriteBuffer io, PlcValue value, String seq) {
        PlcDATE_AND_TIME dt = (PlcDATE_AND_TIME) value;
        LocalDateTime ldt = dt.getDateTime();
        try {
            switch (seq) {
                case "1":
                    int year = ldt.getYear();
                    if (year < 2000) {
                        year -= 1900;
                    } else {
                        year -= 2000;
                    }
                    int ret = Integer.parseInt("" + year, 16);
                    io.writeByte((byte) ret);
                    break;
                case "2":
                    ret = Integer.parseInt("" + ldt.getMonthValue(), 16);
                    io.writeByte((byte) ret);
                    break;
                case "3":
                    ret = Integer.parseInt("" + ldt.getDayOfMonth(), 16);
                    io.writeByte((byte) ret);
                    break;
                case "4":
                    ret = Integer.parseInt("" + ldt.getHour(), 16);
                    io.writeByte((byte) ret);
                    break;
                case "5":
                    ret = Integer.parseInt("" + ldt.getMinute(), 16);
                    io.writeByte((byte) ret);
                    break;
                case "6":
                    ret = Integer.parseInt("" + ldt.getSecond(), 16);
                    io.writeByte((byte) ret);
                    break;
                default:
                    throw new IllegalStateException("Unexpected value: " + seq);
            }
        } catch (ParseException e) {
            throw new PlcRuntimeException("Parse DateTime error " + e.getMessage());
        }

    }

    public static void serializeDateAndTime(WriteBuffer io, PlcValue value) {
        throw new NotImplementedException("Serializing DateAndTime not implemented");
    }

    public static void serializeS5Time(WriteBuffer io, PlcValue value) {
        throw new NotImplementedException("Serializing S5TIME not implemented");

    }

    public static LocalTime parseTiaLTime(ReadBuffer io) {
        throw new NotImplementedException("LTIME not implemented");
    }

    public static void serializeTiaLTime(WriteBuffer io, PlcValue value) {
        throw new NotImplementedException("Serializing LTIME not implemented");
    }

    public static LocalTime parseTiaTimeOfDay(ReadBuffer io) {
        try {
            long millisSinceMidnight = io.readUnsignedLong(32);
            return LocalTime.now().withHour(0).withMinute(0).withSecond(0).withNano(0).plus(
                millisSinceMidnight, ChronoUnit.MILLIS);
        } catch (ParseException e) {
            return null;
        }
    }

    public static void serializeTiaTimeOfDay(WriteBuffer io, PlcValue value) {
        throw new NotImplementedException("Serializing TIME_OF_DAY not implemented");
    }

    public static LocalDate parseTiaDate(ReadBuffer io) {
        try {
            int daysSince1990 = io.readUnsignedInt(16);
            return LocalDate.now().withYear(1990).withDayOfMonth(1).withMonth(1).plus(daysSince1990, ChronoUnit.DAYS);
        } catch (ParseException e) {
            return null;
        }
    }

    public static void serializeTiaDate(WriteBuffer io, PlcValue value) {
        throw new NotImplementedException("Serializing DATE not implemented");
    }

    //TODO: Call BCD converter
    public static LocalDateTime parseTiaDateTime(ReadBuffer io) {
        try {
            int year = io.readUnsignedInt(16);
            int month = io.readUnsignedInt(8);
            int day = io.readUnsignedInt(8);
            // Skip day-of-week
            io.readByte();
            int hour = io.readByte();
            int minute = io.readByte();
            int second = io.readByte();
            int nanosecond = io.readUnsignedInt(24);
            // Skip day-of-week
            io.readByte();

            return LocalDateTime.of(year, month, day, hour, minute, second, nanosecond);
        } catch (Exception e) {
            return null;
        }
    }

    public static void serializeTiaDateTime(WriteBuffer io, PlcValue value) {
        throw new NotImplementedException("Serializing DATE_AND_TIME not implemented");
    }

    public static String parseS7Char(ReadBuffer io, String encoding) {
        if ("UTF-8".equalsIgnoreCase(encoding)) {
            return io.readString(8, encoding);
        } else if ("UTF-16".equalsIgnoreCase(encoding)) {
            return io.readString(16, encoding);
        } else {
            throw new PlcRuntimeException("Unsupported encoding");
        }
    }

    public static Integer parseDate(ReadBuffer io) {
        try {
            Integer value = io.readUnsignedInt("", 16);
            return value + 7305;
        } catch (ParseException e) {
            throw new PlcRuntimeException("read Date parse error");
        }

    }

    public static void serializeDate(WriteBuffer io, PlcValue value) {
        int v = value.getInt();
        try {
            io.writeUnsignedInt("", 16, v - 7305);
        } catch (ParseException e) {
            throw new PlcRuntimeException("wriet Date parse error");
        }
    }

    public static String parseS7String(ReadBuffer io, int stringLength, String encoding) {
        try {
            if ("UTF-8".equalsIgnoreCase(encoding)) {
                // This is the maximum number of bytes a string can be long.
                short maxLength = io.readUnsignedShort(8);
                // This is the total length of the string on the PLC (Not necessarily the number of characters read)
                short totalStringLength = io.readUnsignedShort(8);
                totalStringLength = (short) Math.min(maxLength, totalStringLength);
                final byte[] byteArray = new byte[totalStringLength];
                for (int i = 0; (i < stringLength) && io.hasMore(8); i++) {
                    final byte curByte = io.readByte();
                    if (i < totalStringLength) {
                        byteArray[i] = curByte;
                    } else {
                        // Gobble up the remaining data, which is not added to the string.
                        i++;
                        for (; (i < stringLength) && io.hasMore(8); i++) {
                            io.readByte();
                        }
                        break;
                    }
                }
                Charset charset = detectCharset(byteArray);
                if (charset == null) {
                    charset = StandardCharsets.US_ASCII;
                }
                String substr = new String(byteArray, charset);
                substr = substr.replaceAll("[^\u0020-\u9FA5]", "");
                return substr;
            } else if ("UTF-16".equalsIgnoreCase(encoding)) {
                // This is the maximum number of bytes a string can be long.
                int maxLength = io.readUnsignedInt(16);
                // This is the total length of the string on the PLC (Not necessarily the number of characters read)
                int totalStringLength = io.readUnsignedInt(16);
                totalStringLength = Math.min(maxLength, totalStringLength);
                final byte[] byteArray = new byte[totalStringLength * 2];
                for (int i = 0; (i < stringLength) && io.hasMore(16); i++) {
                    final short curShort = io.readShort(16);
                    if (i < totalStringLength) {
                        byteArray[i * 2] = (byte) (curShort >>> 8);
                        byteArray[(i * 2) + 1] = (byte) (curShort & 0xFF);
                    } else {
                        // Gobble up the remaining data, which is not added to the string.
                        i++;
                        for (; (i < stringLength) && io.hasMore(16); i++) {
                            io.readShort(16);
                        }
                        break;
                    }
                }
                return new String(byteArray, StandardCharsets.UTF_16);
            } else {
                throw new PlcRuntimeException("Unsupported string encoding " + encoding);
            }
        } catch (ParseException e) {
            throw new PlcRuntimeException("Error parsing string", e);
        }
    }

    public static void serializeS7Char(WriteBuffer io, PlcValue value, String encoding) {
        if (value instanceof PlcList) {
            PlcList list = (PlcList) value;
            list.getList().forEach(v -> writeChar(io, v, encoding));
        } else {
            writeChar(io, value, encoding);
        }
    }

    private static void writeChar(WriteBuffer io, PlcValue value, String encoding) {
        if ("UTF-8".equalsIgnoreCase(encoding)) {
            try {
                byte valueByte = value.getByte();
                io.writeByte(valueByte);
            } catch (ParseException e) {
                throw new PlcRuntimeException("writeChar error");
            }
        } else if ("UTF-16".equalsIgnoreCase(encoding)) {
            try {
                byte[] bytes = ((PlcWCHAR) value).getBytes();
                io.writeByteArray(bytes);
            } catch (ParseException e) {
                throw new PlcRuntimeException("writeWChar error");
            }
        } else {
            throw new PlcRuntimeException("Unsupported encoding");
        }
    }

    public static void serializeS7String(WriteBuffer io, PlcValue value, int stringLength, String encoding) throws ParseException {
        String valueString = (String) value.getObject();
        if ("UTF-8".equalsIgnoreCase(encoding)) {
            valueString = valueString == null ? "" : valueString;
            Charset charsetTemp = getEncoding(valueString);
            if (charsetTemp == null) {
                charsetTemp = StandardCharsets.UTF_8;
            }
            final byte[] raw = valueString.getBytes(charsetTemp);
            io.writeByte((byte) stringLength);
            io.writeByte((byte) raw.length);
            io.writeByteArray(raw);
        } else if ("UTF-16".equalsIgnoreCase(encoding)) {
            io.writeUnsignedInt(16, stringLength);
            try {
                byte[] bytes = valueString.getBytes(encoding);
                io.writeUnsignedInt(16, bytes.length / 2);
            } catch (UnsupportedEncodingException e) {
                throw new ParseException("Parse wstring error");
            }

            io.writeString(stringLength * 2 * 8, encoding, valueString);
        } else {
            throw new PlcRuntimeException("Unsupported string encoding " + encoding);
        }
    }
}
