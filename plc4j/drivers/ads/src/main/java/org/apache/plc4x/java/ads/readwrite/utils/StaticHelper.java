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
package org.apache.plc4x.java.ads.readwrite.utils;

import org.apache.plc4x.java.spi.buffers.api.ReadBuffer;
import org.apache.plc4x.java.spi.buffers.api.WithOption;
import org.apache.plc4x.java.spi.buffers.api.WriteBuffer;
import org.apache.plc4x.java.spi.buffers.api.exceptions.BufferException;
import org.apache.plc4x.java.api.value.PlcValue;

import java.nio.ByteBuffer;
import java.nio.charset.CharacterCodingException;
import java.nio.charset.Charset;
import java.nio.charset.CodingErrorAction;
import java.nio.charset.StandardCharsets;
import java.util.Arrays;

public class StaticHelper {

    private static final String[] DEFAULT_CHARSETS = {"ASCII", "UTF-8", "GBK", "GB2312", "BIG5", "GB18030"};
    private static final Charset WINDOWS_1252 = Charset.forName("Windows-1252");

    public static String parseZeroTerminatedString(ReadBuffer io, int stringValueLength) throws BufferException {
        byte[] bytes = io.readBits(stringValueLength * 8, WithOption.WithName("stringValueLength"));
        String stringValue = new String(bytes, 0, bytes.length);
        // Consume the zero terminator
        byte terminatorByte = io.readSignedByte(8);
        if (terminatorByte != (byte) 0x00) {
            throw new BufferException("Expected 0x00, but found " + terminatorByte);
        }
        return stringValue;
    }

    public static void serializeZeroTerminatedString(WriteBuffer io, String data) throws BufferException {
        io.writeString(data.length() * 8, data, WithOption.WithName("stringValue"), WithOption.WithEncoding("ASCII"));
        io.writeSignedByte(8, (byte) 0x00, WithOption.WithName("terminator"));
    }

    public static int lengthZeroTerminatedString(String data) {
        return (data.length() + 1) * 8;
    }

    public static String parseAmsString(ReadBuffer io, int stringLength, String encoding, String stringEncoding) throws BufferException {
        boolean wide = isWideEncoding(encoding);
        byte[] raw = new byte[wide ? stringLength * 2 : stringLength];
        for (int i = 0; i < raw.length; i++) {
            raw[i] = io.readSignedByte(8);
        }
        if (wide) {
            int terminator = io.readUnsignedInt(16);
            if (terminator != 0x0000) {
                throw new BufferException("Expected 0x0000, but found " + terminator);
            }
        } else {
            short terminator = io.readUnsignedShort(8);
            if (terminator != 0x00) {
                throw new BufferException("Expected 0x00, but found " + terminator);
            }
        }
        byte[] content = wide ? trimWideTerminated(raw) : trimZeroTerminated(raw);
        return new String(content, resolveReadCharset(stringEncoding, wide, content));
    }

    public static void serializeAmsString(WriteBuffer io, PlcValue value, int stringLength, String encoding, String stringEncoding) throws BufferException {
        boolean wide = isWideEncoding(encoding);
        String stringValue = value.getString();
        stringValue = stringValue == null ? "" : stringValue;
        Charset charset = resolveWriteCharset(stringEncoding, wide, stringValue);
        if (wide && stringValue.length() > stringLength) {
            stringValue = stringValue.substring(0, stringLength);
        }
        byte[] encoded = stringValue.getBytes(charset);
        int dataLength = wide ? stringLength * 2 : stringLength;
        for (int i = 0; i < dataLength; i++) {
            io.writeSignedByte(8, i < encoded.length ? encoded[i] : 0);
        }
        if (wide) {
            io.writeUnsignedInt(16, 0x0000);
        } else {
            io.writeUnsignedShort(8, (short) 0x00);
        }
    }

    private static byte[] trimZeroTerminated(byte[] bytes) {
        int length = 0;
        while (length < bytes.length && bytes[length] != 0) {
            length++;
        }
        return Arrays.copyOf(bytes, length);
    }

    private static byte[] trimWideTerminated(byte[] bytes) {
        int length = bytes.length;
        for (int i = 0; i + 1 < bytes.length; i += 2) {
            if (bytes[i] == 0 && bytes[i + 1] == 0) {
                length = i;
                break;
            }
        }
        return Arrays.copyOf(bytes, length);
    }

    private static boolean isWideEncoding(String encoding) {
        return encoding != null && encoding.replace("-", "").equalsIgnoreCase("UTF16");
    }

    private static Charset resolveReadCharset(String stringEncoding, boolean wide, byte[] content) {
        if (stringEncoding == null || stringEncoding.isEmpty() || "AUTO".equalsIgnoreCase(stringEncoding)) {
            return wide ? StandardCharsets.UTF_16LE : detectCharset(content);
        }
        return resolveCharset(stringEncoding, wide);
    }

    private static Charset resolveWriteCharset(String stringEncoding, boolean wide, String value) {
        if (stringEncoding == null || stringEncoding.isEmpty() || "AUTO".equalsIgnoreCase(stringEncoding)) {
            return wide ? StandardCharsets.UTF_16LE : detectWriteCharset(value);
        }
        return resolveCharset(stringEncoding, wide);
    }

    private static Charset resolveCharset(String charsetName, boolean wide) {
        String normalized = charsetName.replace('_', '-');
        if ("UTF8".equalsIgnoreCase(normalized) || "UTF-8".equalsIgnoreCase(normalized)) {
            return StandardCharsets.UTF_8;
        }
        if ("UTF16".equalsIgnoreCase(normalized) || "UTF-16".equalsIgnoreCase(normalized)) {
            return wide ? StandardCharsets.UTF_16LE : StandardCharsets.UTF_16;
        }
        if ("UTF16LE".equalsIgnoreCase(normalized) || "UTF-16LE".equalsIgnoreCase(normalized)) {
            return StandardCharsets.UTF_16LE;
        }
        if ("UTF16BE".equalsIgnoreCase(normalized) || "UTF-16BE".equalsIgnoreCase(normalized)) {
            return StandardCharsets.UTF_16BE;
        }
        if ("WINDOWS1252".equalsIgnoreCase(normalized) || "WINDOWS-1252".equalsIgnoreCase(normalized)) {
            return WINDOWS_1252;
        }
        return Charset.forName(normalized);
    }

    private static Charset detectCharset(byte[] bytes) {
        for (String charsetName : DEFAULT_CHARSETS) {
            Charset charset = Charset.forName(charsetName);
            if (canDecode(bytes, charset)) {
                return charset;
            }
        }
        return WINDOWS_1252;
    }

    private static boolean canDecode(byte[] bytes, Charset charset) {
        try {
            charset.newDecoder()
                .onMalformedInput(CodingErrorAction.REPORT)
                .onUnmappableCharacter(CodingErrorAction.REPORT)
                .decode(ByteBuffer.wrap(bytes));
            return true;
        } catch (CharacterCodingException e) {
            return false;
        }
    }

    private static Charset detectWriteCharset(String value) {
        for (String charsetName : DEFAULT_CHARSETS) {
            Charset charset = Charset.forName(charsetName);
            if (value.equals(new String(value.getBytes(charset), charset))) {
                return charset;
            }
        }
        return WINDOWS_1252;
    }

}
