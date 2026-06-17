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
package org.apache.plc4x.java.s7.tag;

import org.apache.plc4x.java.api.exceptions.PlcInvalidTagException;
import org.apache.plc4x.java.s7.readwrite.MemoryArea;
import org.apache.plc4x.java.s7.readwrite.TransportSize;
import org.apache.plc4x.java.spi.buffers.api.WithOption;
import org.apache.plc4x.java.spi.buffers.api.WriteBuffer;
import org.apache.plc4x.java.spi.buffers.api.exceptions.BufferException;

import java.nio.charset.StandardCharsets;
import java.util.Objects;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Single tag class for S7 STRING and WSTRING, with or without an explicit
 * string-length suffix and optional character encoding.
 *
 * <p>Address formats:
 * <ul>
 *   <li>{@code %DB1.DB0:STRING} — default length 254, default encoding UTF-8</li>
 *   <li>{@code %DB1:56:STRING(20)} — explicit length 20</li>
 *   <li>{@code %DB1:56:STRING(20)|GBK} — explicit length 20, GBK encoding</li>
 *   <li>{@code %DB1:56:WSTRING|UTF-16} — WSTRING with explicit encoding</li>
 *   <li>{@code %DB1:56:STRING[3]} — array of 3 strings</li>
 *   <li>{@code %DB1:56:STRING(20)[3]|GBK} — array with length and encoding</li>
 * </ul>
 */
public class S7StringTag extends S7Tag {

    private static final String ENCODING_SUFFIX = "(\\|(?<stringEncoding>[a-zA-Z0-9_-]+))?";

    /** e.g. %DB1.DB0:STRING or %DB1.DB0:STRING(20)|GBK */
    public static final Pattern DATA_BLOCK_STRING_PATTERN =
        Pattern.compile("^%DB(?<blockNumber>\\d{1,5})\\.DB(?<transferSizeCode>[XBWD]?)"
            + "(?<byteOffset>\\d{1,7})(\\.(?<bitOffset>[0-7]))?:"
            + "(?<dataType>STRING|WSTRING)"
            + "(?:\\((?<stringLength>\\d{1,3})\\))?"
            + "(?:\\[(?<numElements>\\d+)])?"
            + ENCODING_SUFFIX);

    /** e.g. %DB1:56:STRING or %DB1:56:STRING(20)|GBK */
    public static final Pattern DATA_BLOCK_STRING_SHORT_PATTERN =
        Pattern.compile("^%DB(?<blockNumber>\\d{1,5}):"
            + "(?<byteOffset>\\d{1,7})(\\.(?<bitOffset>[0-7]))?:"
            + "(?<dataType>STRING|WSTRING)"
            + "(?:\\((?<stringLength>\\d{1,3})\\))?"
            + "(?:\\[(?<numElements>\\d+)])?"
            + ENCODING_SUFFIX);

    private static final int DEFAULT_STRING_LENGTH = 254;
    /** null means auto-detect on read, default to UTF-8 on write */
    private static final String DEFAULT_STRING_ENCODING = null;
    private static final String DEFAULT_WSTRING_ENCODING = "UTF-16";

    private final int stringLength;
    private final String stringEncoding;

    public S7StringTag(TransportSize dataType, MemoryArea memoryArea,
                       int blockNumber, int byteOffset,
                       byte bitOffset, int numElements,
                       int stringLength) {
        this(dataType, memoryArea, blockNumber, byteOffset, bitOffset, numElements,
            stringLength, dataType == TransportSize.WSTRING ? DEFAULT_WSTRING_ENCODING : DEFAULT_STRING_ENCODING);
    }

    public S7StringTag(TransportSize dataType, MemoryArea memoryArea,
                       int blockNumber, int byteOffset,
                       byte bitOffset, int numElements,
                       int stringLength, String stringEncoding) {
        super(dataType, memoryArea, blockNumber, byteOffset, bitOffset, numElements);
        this.stringLength = stringLength;
        this.stringEncoding = stringEncoding != null ? stringEncoding
            : (dataType == TransportSize.WSTRING ? DEFAULT_WSTRING_ENCODING : DEFAULT_STRING_ENCODING);
    }

    public static boolean matches(String address) {
        return DATA_BLOCK_STRING_PATTERN.matcher(address).matches()
            || DATA_BLOCK_STRING_SHORT_PATTERN.matcher(address).matches();
    }

    public static S7StringTag of(String address) {
        Matcher matcher = DATA_BLOCK_STRING_SHORT_PATTERN.matcher(address);
        if (!matcher.matches()) {
            matcher = DATA_BLOCK_STRING_PATTERN.matcher(address);
        }
        if (!matcher.matches()) {
            throw new PlcInvalidTagException(
                "Unable to parse address: " + address + ". Doesn't match S7StringTag patterns");
        }

        TransportSize dataType = TransportSize.valueOf(matcher.group("dataType"));
        MemoryArea memoryArea = MemoryArea.DATA_BLOCKS;

        int blockNumber = Integer.parseInt(matcher.group("blockNumber"));
        int byteOffset = Integer.parseInt(matcher.group("byteOffset"));
        byte bitOffset = 0;
        if (matcher.group("bitOffset") != null) {
            bitOffset = Byte.parseByte(matcher.group("bitOffset"));
        }
        int numElements = 1;
        if (matcher.group("numElements") != null) {
            numElements = Integer.parseInt(matcher.group("numElements"));
        }
        int stringLength = DEFAULT_STRING_LENGTH;
        if (matcher.group("stringLength") != null) {
            stringLength = Integer.parseInt(matcher.group("stringLength"));
        }
        String stringEncoding = matcher.group("stringEncoding");
        if (stringEncoding == null || stringEncoding.isEmpty()) {
            stringEncoding = dataType == TransportSize.WSTRING
                ? DEFAULT_WSTRING_ENCODING : DEFAULT_STRING_ENCODING;
        }

        return new S7StringTag(dataType, memoryArea, blockNumber, byteOffset,
            bitOffset, numElements, stringLength, stringEncoding);
    }

    public int getStringLength() {
        return stringLength;
    }

    public String getStringEncoding() {
        return stringEncoding;
    }

    @Override
    public String getAddressString() {
        String address = String.format("%%DB%d.DB%d:%s(%d)",
            getBlockNumber(), getByteOffset(), getDataType().name(), getStringLength());
        if (getNumberOfElements() != 1) {
            address += "[" + getNumberOfElements() + "]";
        }
        if (stringEncoding != null && !stringEncoding.isEmpty()) {
            address += "|" + stringEncoding;
        }
        return address;
    }

    @Override
    public String toString() {
        return "S7StringTag{" +
            "dataType=" + getDataType() +
            ", memoryArea=" + getMemoryArea() +
            ", blockNumber=" + getBlockNumber() +
            ", byteOffset=" + getByteOffset() +
            ", bitOffset=" + getBitOffset() +
            ", numElements=" + getNumberOfElements() +
            ", stringLength=" + stringLength +
            ", encoding=" + stringEncoding +
            '}';
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        if (!super.equals(o)) return false;
        S7StringTag that = (S7StringTag) o;
        return stringLength == that.stringLength
            && Objects.equals(stringEncoding, that.stringEncoding);
    }

    @Override
    public int hashCode() {
        return Objects.hash(super.hashCode(), stringLength, stringEncoding);
    }

    @Override
    public void serialize(WriteBuffer writeBuffer) throws BufferException {
        writeBuffer.pushContext();

        String memoryArea = getMemoryArea().name();
        writeBuffer.writeString(
            memoryArea.getBytes(StandardCharsets.UTF_8).length * 8,
            memoryArea, WithOption.WithEncoding("UTF8"));

        writeBuffer.writeUnsignedInt(16, getBlockNumber());
        writeBuffer.writeUnsignedInt(16, getByteOffset());
        writeBuffer.writeUnsignedInt(8, getBitOffset());
        writeBuffer.writeUnsignedInt(16, getNumberOfElements());
        writeBuffer.writeUnsignedInt(16, getStringLength());

        String dataType = getDataType().name();
        writeBuffer.writeString(
            dataType.getBytes(StandardCharsets.UTF_8).length * 8,
            dataType, WithOption.WithEncoding("UTF8"));
        if (getStringEncoding() != null && !getStringEncoding().isEmpty()) {
            writeBuffer.writeString(
                getStringEncoding().getBytes(StandardCharsets.UTF_8).length * 8,
                getStringEncoding(), WithOption.WithEncoding("UTF8"));
        }

        writeBuffer.popContext();
    }

}
