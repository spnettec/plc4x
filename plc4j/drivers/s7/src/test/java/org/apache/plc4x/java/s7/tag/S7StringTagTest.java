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

import org.apache.plc4x.java.s7.readwrite.MemoryArea;
import org.apache.plc4x.java.s7.readwrite.TransportSize;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class S7StringTagTest {

    @Test
    void parseStringWithExplicitLength() {
        S7StringTag tag = S7StringTag.of("%DB1.DB0:STRING(80)");
        assertNotNull(tag);
        assertEquals(80, tag.getStringLength());
        assertEquals(TransportSize.STRING, tag.getDataType());
        assertEquals(1, tag.getBlockNumber());
        assertEquals(0, tag.getByteOffset());
    }

    @Test
    void parseStringWithExplicitLengthAndCount() {
        S7StringTag tag = S7StringTag.of("%DB1.DB0:STRING(40)[3]");
        assertNotNull(tag);
        assertEquals(40, tag.getStringLength());
        assertEquals(3, tag.getNumberOfElements());
    }

    @Test
    void parseStringShortForm() {
        S7StringTag tag = S7StringTag.of("%DB1:0:STRING(40)");
        assertNotNull(tag);
        assertEquals(40, tag.getStringLength());
    }

    @Test
    void parseStringDefaultLength() {
        S7StringTag tag = S7StringTag.of("%DB1.DB0:STRING");
        assertNotNull(tag);
        assertEquals(254, tag.getStringLength());
        assertEquals(TransportSize.STRING, tag.getDataType());
    }

    @Test
    void parseStringDefaultLengthShortForm() {
        S7StringTag tag = S7StringTag.of("%DB1:0:STRING");
        assertNotNull(tag);
        assertEquals(254, tag.getStringLength());
        assertEquals(TransportSize.STRING, tag.getDataType());
    }

    @Test
    void parseWStringDefaultLength() {
        S7StringTag tag = S7StringTag.of("%DB1.DB0:WSTRING");
        assertNotNull(tag);
        assertEquals(254, tag.getStringLength());
        assertEquals(TransportSize.WSTRING, tag.getDataType());
        assertTrue(tag.toString().contains("WSTRING"));
    }

    @Test
    void stringMatches() {
        assertTrue(S7StringTag.matches("%DB1.DB0:STRING(80)"));
        assertTrue(S7StringTag.matches("%DB1:0:STRING(80)"));
        assertTrue(S7StringTag.matches("%DB1.DB0:STRING"));
        assertTrue(S7StringTag.matches("%DB1:0:WSTRING"));
        assertTrue(S7StringTag.matches("%DB1:0:STRING[3]"));
        assertTrue(S7StringTag.matches("%DB1:0:STRING(20)[2]"));
        assertFalse(S7StringTag.matches("%DB1.DBW0:INT"));
    }

    @Test
    void equalityAndHashCode() {
        S7StringTag a = new S7StringTag(TransportSize.STRING, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, 80);
        S7StringTag b = new S7StringTag(TransportSize.STRING, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, 80);
        S7StringTag c = new S7StringTag(TransportSize.STRING, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, 40);
        assertEquals(a, b);
        assertEquals(a.hashCode(), b.hashCode());
        assertNotEquals(a, c);
    }

    @Test
    void toStringMentionsLength() {
        S7StringTag tag = new S7StringTag(TransportSize.STRING, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, 80);
        assertTrue(tag.toString().contains("80"));
    }

    @Test
    void tagHandlerRoutesToSingleStringTag() {
        S7PlcTagHandler handler = new S7PlcTagHandler();
        assertInstanceOf(S7StringTag.class, handler.parseTag("%DB1.DB0:STRING(80)"));
        assertEquals(80, ((S7StringTag) handler.parseTag("%DB1.DB0:STRING(80)")).getStringLength());
        assertInstanceOf(S7StringTag.class, handler.parseTag("%DB1.DB0:STRING"));
        assertEquals(254, ((S7StringTag) handler.parseTag("%DB1.DB0:STRING")).getStringLength());
        assertInstanceOf(S7Tag.class, handler.parseTag("%MW0:INT"));
        assertThrows(org.apache.plc4x.java.api.exceptions.PlcInvalidTagException.class,
            () -> handler.parseTag("not-a-tag"));
        assertNull(handler.parseQuery("any"));
    }
}
