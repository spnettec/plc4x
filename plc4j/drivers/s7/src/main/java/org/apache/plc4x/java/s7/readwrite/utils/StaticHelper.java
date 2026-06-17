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
package org.apache.plc4x.java.s7.readwrite.utils;

import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.buffers.api.ReadBuffer;
import org.apache.plc4x.java.spi.buffers.api.WithOption;
import org.apache.plc4x.java.spi.buffers.api.WriteBuffer;
import org.apache.plc4x.java.spi.buffers.api.exceptions.BufferException;

import org.apache.plc4x.java.s7.readwrite.DataTransportSize;

import java.nio.charset.Charset;
import java.nio.charset.StandardCharsets;
import java.time.LocalDate;
import java.time.temporal.ChronoUnit;

/**
 * Hand-written helper methods called from the mspec-generated DataItem via STATIC_CALL,
 * and from S7CotpConnection directly (bypassing DataItem for STRING encoding support).
 */
public class StaticHelper {

    private static final WithOption UINT_OPT = WithOption.WithUnsignedIntegerEncoding("unsigned-binary");
    private static final WithOption SINT_OPT = WithOption.WithSignedIntegerEncoding("twos-complement");

    // ════════════════════════════════════════════════════════════════════════
    //  S7 STRING / WSTRING
    // ════════════════════════════════════════════════════════════════════════

    /**
     * Parse an S7 STRING or WSTRING from the buffer.
     * <pre>
     *  STRING:  1-byte maxLen + 1-byte curLen + stringLength bytes
     *  WSTRING: 2-byte maxLen + 2-byte curLen + stringLength×2 bytes
     * </pre>
     *
     * @param encoding charset name ({@code "UTF8"}, {@code "GBK"}, {@code "UTF16BE"}, etc.)
     *                 or {@code null} for auto-detection
     */
    public static String parseS7String(ReadBuffer io, int stringLength, String encoding) {
        try {
            boolean isWide = isWideEncoding(encoding);

            if (!isWide) {
                int maxLen = io.readUnsignedShort(8, UINT_OPT) & 0xFF;
                int curLen = io.readUnsignedShort(8, UINT_OPT) & 0xFF;

                byte[] raw = new byte[stringLength];
                int bytesRead = 0;
                for (; bytesRead < stringLength && io.getRemainingBits() >= 8; bytesRead++) {
                    raw[bytesRead] = io.readSignedByte(8, SINT_OPT);
                }

                int actualLen = Math.min(curLen, bytesRead);
                byte[] content = java.util.Arrays.copyOf(raw, actualLen);

                Charset charset = (encoding == null || encoding.isEmpty())
                    ? detectCharset(content) : resolveCharset(encoding);
                return new String(content, charset);
            } else {
                int maxLen = io.readUnsignedInt(16, UINT_OPT);
                int curLen = io.readUnsignedInt(16, UINT_OPT);

                byte[] raw = new byte[stringLength * 2];
                int bytesRead = 0;
                for (; bytesRead < raw.length && io.getRemainingBits() >= 8; bytesRead++) {
                    raw[bytesRead] = io.readSignedByte(8, SINT_OPT);
                }

                int actualByteLen = Math.min(curLen * 2, bytesRead);
                byte[] content = java.util.Arrays.copyOf(raw, actualByteLen);

                Charset charset = (encoding == null || encoding.isEmpty())
                    ? StandardCharsets.UTF_16BE : resolveCharset(encoding);
                return new String(content, charset);
            }
        } catch (BufferException e) {
            throw new PlcRuntimeException("Error parsing S7 string", e);
        }
    }

    /**
     * Serialize a string value into S7 STRING/WSTRING format.
     */
    public static void serializeS7String(WriteBuffer io, PlcValue value, int stringLength, String encoding) {
        boolean isWide = isWideEncoding(encoding);
        String str = value.getString();
        str = (str == null) ? "" : str;

        Charset charset;
        if (encoding == null || encoding.isEmpty()) {
            charset = isWide ? StandardCharsets.UTF_16BE : detectWriteCharset(str);
        } else {
            charset = resolveCharset(encoding);
        }

        try {
            byte[] encoded = str.getBytes(charset);

            if (!isWide) {
                int actLen = Math.min(encoded.length, stringLength);
                io.writeSignedByte(8, (byte) stringLength, SINT_OPT);
                io.writeSignedByte(8, (byte) actLen, SINT_OPT);
                for (int i = 0; i < stringLength; i++) {
                    io.writeSignedByte(8, (i < actLen) ? encoded[i] : 0, SINT_OPT);
                }
            } else {
                int charCount = str.length();
                int actChars = Math.min(charCount, stringLength);
                io.writeUnsignedInt(16, stringLength, UINT_OPT);
                io.writeUnsignedInt(16, actChars, UINT_OPT);
                byte[] wideBytes = str.substring(0, actChars).getBytes(charset);
                int totalBytes = stringLength * 2;
                for (int i = 0; i < totalBytes; i++) {
                    io.writeSignedByte(8, (i < wideBytes.length) ? wideBytes[i] : 0, SINT_OPT);
                }
            }
        } catch (BufferException e) {
            throw new PlcRuntimeException("Error serializing S7 string", e);
        }
    }

    // ════════════════════════════════════════════════════════════════════════
    //  S5TIME
    // ════════════════════════════════════════════════════════════════════════

    /** Parse S5TIME (2 bytes: 2-bit time-base + 3×4-bit BCD = 0–999). Returns milliseconds. */
    public static long parseS5Time(ReadBuffer io) {
        try {
            int raw = io.readUnsignedInt(16, UINT_OPT);
            int timeBase = (raw >> 12) & 0x03;
            int hundreds = (raw >> 8) & 0x0F;
            int tens     = (raw >> 4) & 0x0F;
            int ones     = raw & 0x0F;
            int bcdValue = hundreds * 100 + tens * 10 + ones;
            long[] multipliers = {10, 100, 1000, 10000};
            return bcdValue * multipliers[timeBase];
        } catch (BufferException e) {
            throw new PlcRuntimeException("Error parsing S5TIME", e);
        }
    }

    /** Serialize S5TIME (milliseconds → 2-byte BCD + time-base). */
    public static void serializeS5Time(WriteBuffer io, PlcValue value) {
        try {
            long ms = value.getDuration().toMillis();
            int timeBase;
            int bcdValue;
            if (ms <= 9990) {
                timeBase = 0; bcdValue = (int) (ms / 10);
            } else if (ms <= 99900) {
                timeBase = 1; bcdValue = (int) (ms / 100);
            } else if (ms <= 999000) {
                timeBase = 2; bcdValue = (int) (ms / 1000);
            } else {
                timeBase = 3; bcdValue = Math.min((int) (ms / 10000), 999);
            }
            int h = (bcdValue / 100) % 10;
            int t = (bcdValue / 10) % 10;
            int o = bcdValue % 10;
            io.writeUnsignedInt(16, (timeBase << 12) | (h << 8) | (t << 4) | o, UINT_OPT);
        } catch (BufferException e) {
            throw new PlcRuntimeException("Error serializing S5TIME", e);
        }
    }

    // ════════════════════════════════════════════════════════════════════════
    //  Siemens Year (BCD) and TIA Date
    // ════════════════════════════════════════════════════════════════════════

    /** Parse 1-byte BCD Siemens year: 0–89 → 2000–2089, 90–99 → 1990–1999. */
    public static int parseSiemensYear(ReadBuffer io) {
        try {
            short raw = io.readUnsignedShort(8, UINT_OPT);
            int year = ((raw >> 4) & 0x0F) * 10 + (raw & 0x0F);
            return (year >= 90) ? 1900 + year : 2000 + year;
        } catch (BufferException e) {
            throw new PlcRuntimeException("Error parsing Siemens year", e);
        }
    }

    /** Serialize year to 1-byte BCD Siemens format. */
    public static void serializeSiemensYear(WriteBuffer io, PlcValue value) {
        try {
            int twoDigit = value.getDateTime().getYear() % 100;
            io.writeUnsignedShort(8, (short) (((twoDigit / 10) << 4) | (twoDigit % 10)), UINT_OPT);
        } catch (BufferException e) {
            throw new PlcRuntimeException("Error serializing Siemens year", e);
        }
    }

    /** Parse TIA DATE: 2 bytes = days since 1990-01-01. */
    public static long parseTiaDate(ReadBuffer io) {
        try {
            return io.readUnsignedInt(16, UINT_OPT);
        } catch (BufferException e) {
            throw new PlcRuntimeException("Error parsing TIA date", e);
        }
    }

    /** Serialize TIA DATE: days since 1990-01-01. */
    public static void serializeTiaDate(WriteBuffer io, PlcValue value) {
        try {
            long days = ChronoUnit.DAYS.between(LocalDate.of(1990, 1, 1), value.getDate());
            io.writeUnsignedInt(16, (int) days, UINT_OPT);
        } catch (BufferException e) {
            throw new PlcRuntimeException("Error serializing TIA date", e);
        }
    }

    // ════════════════════════════════════════════════════════════════════════
    //  Charset / Encoding Utilities
    // ════════════════════════════════════════════════════════════════════════

    /** Resolve encoding name to Java Charset. Handles PLC4X names and standard names. */
    public static Charset resolveCharset(String encoding) {
        if (encoding == null || encoding.isEmpty()) {
            return StandardCharsets.UTF_8;
        }
        return switch (encoding.toUpperCase()) {
            case "UTF8", "UTF-8" -> StandardCharsets.UTF_8;
            case "UTF16", "UTF-16", "UTF16BE", "UTF-16BE" -> StandardCharsets.UTF_16BE;
            case "UTF16LE", "UTF-16LE" -> StandardCharsets.UTF_16LE;
            case "ASCII", "US-ASCII" -> StandardCharsets.US_ASCII;
            case "ISO-8859-1", "ISO8859-1", "LATIN1" -> StandardCharsets.ISO_8859_1;
            default -> {
                try {
                    yield Charset.forName(encoding);
                } catch (Exception e) {
                    yield StandardCharsets.UTF_8;
                }
            }
        };
    }

    /** Auto-detect charset from raw bytes. Try UTF-8, then GBK, fallback ISO-8859-1. */
    public static Charset detectCharset(byte[] data) {
        if (data == null || data.length == 0) return StandardCharsets.UTF_8;
        if (isValidUtf8(data)) return StandardCharsets.UTF_8;
        if (isLikelyGbk(data)) {
            try { return Charset.forName("GBK"); } catch (Exception ignored) {}
        }
        return StandardCharsets.ISO_8859_1;
    }

    /** Auto-detect charset for writing. Non-ASCII + GBK shorter → GBK, else UTF-8. */
    public static Charset detectWriteCharset(String str) {
        if (str == null || str.isEmpty()) return StandardCharsets.UTF_8;
        for (int i = 0; i < str.length(); i++) {
            if (str.charAt(i) > 0x7F) {
                try {
                    Charset gbk = Charset.forName("GBK");
                    if (str.getBytes(gbk).length <= str.getBytes(StandardCharsets.UTF_8).length) return gbk;
                } catch (Exception ignored) {}
                return StandardCharsets.UTF_8;
            }
        }
        return StandardCharsets.UTF_8;
    }

    /** Check if byte array is valid UTF-8. */
    public static boolean isValidUtf8(byte[] data) {
        int i = 0;
        while (i < data.length) {
            int b = data[i] & 0xFF;
            int len;
            if      (b <= 0x7F)                  len = 1;
            else if (b >= 0xC2 && b <= 0xDF)     len = 2;
            else if (b >= 0xE0 && b <= 0xEF)     len = 3;
            else if (b >= 0xF0 && b <= 0xF4)     len = 4;
            else return false;
            if (i + len > data.length) return false;
            for (int j = 1; j < len; j++) {
                if ((data[i + j] & 0xC0) != 0x80) return false;
            }
            i += len;
        }
        return true;
    }

    /** Check if bytes look like GBK (double-byte Chinese encoding). */
    public static boolean isLikelyGbk(byte[] data) {
        if (data.length < 2) return false;
        int i = 0, pairs = 0;
        while (i < data.length) {
            int b = data[i] & 0xFF;
            if (b <= 0x7F) { i++; continue; }
            if (b >= 0x81 && b <= 0xFE && i + 1 < data.length) {
                int b2 = data[i + 1] & 0xFF;
                if ((b2 >= 0x40 && b2 <= 0x7E) || (b2 >= 0x80 && b2 <= 0xFE)) {
                    pairs++; i += 2; continue;
                }
            }
            return false;
        }
        return pairs > 0;
    }

    /** Check if encoding name is a wide (2-byte) encoding. */
    public static boolean isWideEncoding(String encoding) {
        if (encoding == null || encoding.isEmpty()) return false;
        String upper = encoding.toUpperCase();
        return upper.contains("UTF16") || upper.contains("UTF-16") || upper.contains("UNICODE");
    }

    // ════════════════════════════════════════════════════════════════════════
    //  Alarm / Associated-value helpers (used by generated S7 protocol code)
    // ════════════════════════════════════════════════════════════════════════

    public static void leftShift3(WriteBuffer buffer, int value) throws BufferException {
        buffer.writeUnsignedInt(16, value << 3, UINT_OPT);
    }

    public static int rightShift3(ReadBuffer buffer) throws BufferException {
        return buffer.readUnsignedInt(16, UINT_OPT) >> 3;
    }

    public static int rightShift3(ReadBuffer buffer, DataTransportSize tsize) throws BufferException {
        int value = buffer.readUnsignedInt(16, UINT_OPT);
        if (tsize == DataTransportSize.OCTET_STRING
                || tsize == DataTransportSize.REAL
                || tsize == DataTransportSize.BIT) {
            return value;
        }
        return value >> 3;
    }

    public static int eventItemLength(ReadBuffer buffer, int valueLength) {
        return ((valueLength % 2 == 0) || (buffer.getRemainingBits() < (valueLength + 1) * 8))
            ? valueLength : valueLength + 1;
    }
}
