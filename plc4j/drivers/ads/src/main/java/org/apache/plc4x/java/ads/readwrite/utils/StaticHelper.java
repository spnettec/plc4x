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
import java.util.List;
import java.util.logging.Level;
import java.util.logging.Logger;

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
    public static String parseAmsString(ReadBuffer readBuffer, int stringLength, String encoding) {
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
                final byte[] byteArray = new byte[bytes.size()];
                for (int i = 0; i < bytes.size(); i++) {
                    byteArray[i] = bytes.get(i);
                }
                return new String(byteArray, StandardCharsets.UTF_8);
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
                final byte[] byteArray = new byte[bytes.size()];
                for (int i = 0; i < bytes.size(); i++) {
                    byteArray[i] = bytes.get(i);
                }
                return new String(byteArray, StandardCharsets.UTF_16);
            } else {
                throw new PlcRuntimeException("Unsupported string encoding " + encoding);
            }
        } catch (ParseException e) {
            throw new PlcRuntimeException("Error parsing string", e);
        }
    }

    public static void serializeAmsString(WriteBuffer io, PlcValue value, int stringLength, String encoding) {
        String valueString = (String) value.getObject();
        valueString = valueString == null ? "" : valueString;
        Charset charsetTemp = getEncoding(valueString);
        if ("UTF-8".equalsIgnoreCase(encoding)) {
            if (charsetTemp == null) {
                charsetTemp = StandardCharsets.UTF_8;
            }
            final byte[] raw = valueString.getBytes(charsetTemp);
            valueString = new String(raw,StandardCharsets.UTF_8);
            try {
                io.writeString(stringLength * 8, encoding, valueString);
                //io.writeByte((byte) 0x00);
            }
            catch (SerializationException ex) {
                Logger.getLogger(StaticHelper.class.getName()).log(Level.SEVERE, null, ex);
            }
        } else if ("UTF-16".equalsIgnoreCase(encoding)) {
            if (charsetTemp == null) {
                charsetTemp = StandardCharsets.UTF_16;
            }
            final byte[] raw = valueString.getBytes(charsetTemp);
            valueString = new String(raw,StandardCharsets.UTF_16);
            try {
                io.writeString(stringLength * 16, encoding, valueString);
                //io.writeByte((byte) 0x00);
                //io.writeByte((byte) 0x00);
            }
            catch (SerializationException ex) {
                Logger.getLogger(StaticHelper.class.getName()).log(Level.SEVERE, null, ex);
            }
        } else {
            throw new PlcRuntimeException("Unsupported string encoding " + encoding);
        }
    }

}
