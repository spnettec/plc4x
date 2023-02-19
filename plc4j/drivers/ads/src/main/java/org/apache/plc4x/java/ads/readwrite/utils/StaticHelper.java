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

import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;

import java.io.BufferedInputStream;
import java.io.ByteArrayInputStream;
import java.nio.ByteBuffer;
import java.nio.charset.CharacterCodingException;
import java.nio.charset.Charset;
import java.nio.charset.CharsetDecoder;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.logging.Level;
import java.util.logging.Logger;

public class StaticHelper {
    private static final String[] DEFAULTCHARSETS = {"ASCII", "UTF-8", "GBK", "GB2312", "BIG5", "GB18030"};

    public static Charset detectCharset(String firstMaTch, byte[] bytes) {

        Charset charset = null;
        if (firstMaTch!=null && !"".equals(firstMaTch))
        {
            try {
                charset = Charset.forName(firstMaTch.replaceAll("[^a-zA-Z0-9]", ""));
            }catch (Exception ignored) {
            }
            return charset;
        }
        for (String charsetName : DEFAULTCHARSETS) {
            charset = detectCharset(bytes, Charset.forName(charsetName), bytes.length);
            if (charset != null) {
                break;
            }
        }

        return charset;
    }

    private static Charset detectCharset(byte[] bytes, Charset charset, int length) {
        try {
            BufferedInputStream input = new BufferedInputStream(new ByteArrayInputStream(bytes, 0, length));

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
    public static Charset getEncoding(String firstMaTch, String str) {
        if (str == null || str.trim().length() < 1) {
            return null;
        }
        Charset charset = null;
        if (firstMaTch!=null && !"".equals(firstMaTch))
        {
            try {
                charset = Charset.forName(firstMaTch.replaceAll("[^a-zA-Z0-9]", ""));
            }catch (Exception ignored) {
            }
            return charset;
        }
        for (String encode : DEFAULTCHARSETS) {
            try {
                Charset charset1 = Charset.forName(encode);
                if (str.equals(new String(str.getBytes(charset1), charset1))) {
                    return charset1;
                }
            } catch (Exception er) {
            }
        }
        return null;
    }

    public static String parseAmsString(ReadBuffer readBuffer, int stringLength, String encoding, String stringEncoding) {
        stringLength = Math.min(stringLength, 256);
        try {
            if ("UTF-8".equalsIgnoreCase(encoding)) {
                List<Byte> bytes = new ArrayList<>();
                for(int i = 0; (i < stringLength) && readBuffer.hasMore(8); i++) {
                    final byte curByte = readBuffer.readByte();
                    if (curByte != 0) {
                        bytes.add(curByte);
                    } else {
                        // Gobble up the remaining data, which is not added to the string.
                        i++;
                        for(; (i < stringLength) && readBuffer.hasMore(8); i++) {
                            readBuffer.readByte();
                        }
                        break;
                    }
                }
                // Read the terminating byte.
                readBuffer.readByte();
                final byte[] byteArray = new byte[bytes.size()];
                for (int i = 0; i < bytes.size(); i++) {
                    byteArray[i] = bytes.get(i);
                }

                Charset charset = detectCharset(null,byteArray);
                if (charset == null) {
                    try {
                        charset = Charset.forName(stringEncoding.replaceAll("[^a-zA-Z0-9]", ""));
                    }catch (Exception ignored) {
                    }
                    if (charset == null) {
                        charset = StandardCharsets.UTF_8;
                    }
                }
                String substr = new String(byteArray, charset);
                substr = substr.replaceAll("[^\u0020-\u9FA5]", "");
                return substr;
            } else if ("UTF-16".equalsIgnoreCase(encoding)) {
                List<Byte> bytes = new ArrayList<>();
                for(int i = 0; (i < stringLength) && readBuffer.hasMore(16); i++) {
                    final short curShort = readBuffer.readShort(16);
                    if (curShort != 0) {
                        bytes.add((byte) (curShort >>> 8));
                        bytes.add((byte) (curShort & 0xFF));
                    } else {
                        // Gobble up the remaining data, which is not added to the string.
                        i++;
                        for(; (i < stringLength) && readBuffer.hasMore(16); i++) {
                            readBuffer.readShort(16);
                        }
                        break;
                    }
                }
                // Read the terminating byte.
                readBuffer.readByte();
                final byte[] byteArray = new byte[bytes.size()];
                for (int i = 0; i < bytes.size(); i++) {
                    byteArray[i] = bytes.get(i);
                }
                Charset charset = detectCharset(stringEncoding,byteArray);
                if (charset == null) {
                    try {
                        charset = Charset.forName(stringEncoding.replaceAll("[^a-zA-Z0-9]", ""));
                    }catch (Exception ignored) {
                    }
                    if (charset == null) {
                        charset = StandardCharsets.UTF_16;
                    }
                }
                return new String(byteArray, charset);
            } else {
                throw new PlcRuntimeException("Unsupported string encoding " + encoding);
            }
        } catch (ParseException e) {
            throw new PlcRuntimeException("Error parsing string", e);
        }
    }

    public static void serializeAmsString(WriteBuffer io, PlcValue value, int stringLength, String encoding, String stringEncoding) {
        stringLength = Math.min(stringLength, 256);
        String valueString = (String) value.getObject();
        valueString = valueString == null ? "" : valueString;
        if ("AUTO".equalsIgnoreCase(stringEncoding))
        {
            stringEncoding = null;
        }
        Charset charsetTemp = getEncoding(stringEncoding,valueString);
        if ("UTF-8".equalsIgnoreCase(encoding)) {
            if (charsetTemp == null) {
                charsetTemp = StandardCharsets.UTF_8;
            }
            final byte[] raw = valueString.getBytes(charsetTemp);
            try {
                for (int i = 0; i < stringLength; i++) {
                    if (i >= raw.length) {
                        io.writeByte((byte) 0x00);
                    } else {
                        io.writeByte( raw[i]);
                    }
                }
            }
            catch (SerializationException ex) {
                Logger.getLogger(StaticHelper.class.getName()).log(Level.SEVERE, null, ex);
            }
        } else if ("UTF-16".equalsIgnoreCase(encoding)) {
            if (charsetTemp == null) {
                charsetTemp = StandardCharsets.UTF_16;
            }
            final byte[] raw = valueString.getBytes(charsetTemp);
            try {
                for (int i = 0; i < stringLength * 2; i++) {
                    if (i >= raw.length) {
                        io.writeByte((byte) 0x00);
                    } else {
                        io.writeByte( raw[i]);
                    }
                }
            }
            catch (SerializationException ex) {
                Logger.getLogger(StaticHelper.class.getName()).log(Level.SEVERE, null, ex);
            }
        } else {
            throw new PlcRuntimeException("Unsupported string encoding " + encoding);
        }
    }

}
