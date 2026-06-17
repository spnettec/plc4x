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
package org.apache.plc4x.java.plc4x.readwrite.utils;

import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.buffers.api.ReadBuffer;
import org.apache.plc4x.java.spi.buffers.api.WithOption;
import org.apache.plc4x.java.spi.buffers.api.WriteBuffer;
import org.apache.plc4x.java.spi.buffers.api.exceptions.BufferException;

import java.nio.charset.StandardCharsets;

public class StaticHelper {

    private static final WithOption UINT_OPT = WithOption.WithUnsignedIntegerEncoding("unsigned-binary");
    private static final WithOption SINT_OPT = WithOption.WithSignedIntegerEncoding("twos-complement");

    public static String parseString(ReadBuffer io, String encoding) {
        try {
            if ("UTF-8".equalsIgnoreCase(encoding)) {
                short stringLength = io.readUnsignedShort(8, UINT_OPT);
                byte[] byteArray = new byte[stringLength];
                for (int i = 0; i < stringLength && io.getRemainingBits() >= 8; i++) {
                    byteArray[i] = io.readSignedByte(8, SINT_OPT);
                }
                return new String(byteArray, StandardCharsets.UTF_8);
            } else if ("UTF-16".equalsIgnoreCase(encoding)) {
                int stringLength = io.readUnsignedInt(16, UINT_OPT);
                byte[] byteArray = new byte[stringLength * 2];
                for (int i = 0; i < stringLength && io.getRemainingBits() >= 16; i++) {
                    short curShort = io.readSignedShort(16, SINT_OPT);
                    byteArray[i * 2] = (byte) (curShort >>> 8);
                    byteArray[(i * 2) + 1] = (byte) (curShort & 0xFF);
                }
                return new String(byteArray, StandardCharsets.UTF_16);
            } else {
                throw new PlcRuntimeException("Unsupported string encoding " + encoding);
            }
        } catch (BufferException e) {
            throw new PlcRuntimeException("Error parsing string", e);
        }
    }

    public static void serializeString(WriteBuffer io, PlcValue value, String encoding) {
        String str = value.getString();
        str = (str == null) ? "" : str;
        try {
            if ("UTF-8".equalsIgnoreCase(encoding)) {
                byte[] raw = str.getBytes(StandardCharsets.UTF_8);
                io.writeSignedByte(8, (byte) raw.length, SINT_OPT);
                for (byte b : raw) {
                    io.writeSignedByte(8, b, SINT_OPT);
                }
            } else if ("UTF-16".equalsIgnoreCase(encoding)) {
                byte[] raw = str.getBytes(StandardCharsets.UTF_16);
                io.writeUnsignedInt(16, raw.length, UINT_OPT);
                for (byte b : raw) {
                    io.writeSignedByte(8, b, SINT_OPT);
                }
            } else {
                throw new PlcRuntimeException("Unsupported string encoding " + encoding);
            }
        } catch (BufferException e) {
            throw new PlcRuntimeException("Error serializing string", e);
        }
    }

    public static int stringLengthInBits(PlcValue value, String encoding) {
        String s = (value != null) ? value.getString() : "";
        if (s == null) s = "";
        if ("UTF-8".equalsIgnoreCase(encoding)) {
            return 8 + s.getBytes(StandardCharsets.UTF_8).length * 8;
        } else if ("UTF-16".equalsIgnoreCase(encoding)) {
            return 16 + s.getBytes(StandardCharsets.UTF_16).length * 8;
        }
        return 8;
    }
}
