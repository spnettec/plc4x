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

import org.apache.plc4x.java.s7.readwrite.TransportSize;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

/**
 * S7 strings are variable-length on the wire: the actual length is read from the data, so no
 * declared length belongs in the address. The declared-length form {@code STRING(n)} was an
 * upstream workaround for not handling that and has been dropped; these tests pin the
 * variable-length form, YOFC's {@code |encoding} suffix, and the two array notations an
 * address may carry.
 */
class S7StringTagTest {

    @Test
    void parseVarLengthString() {
        S7StringVarLengthTag tag = S7StringVarLengthTag.of("%DB1.DB0:STRING");
        assertNotNull(tag);
        assertEquals(TransportSize.STRING, tag.getDataType());
    }

    @Test
    void varLengthMatches() {
        assertTrue(S7StringVarLengthTag.matches("%DB1.DB0:STRING"));
        assertTrue(S7StringVarLengthTag.matches("%DB1:0:STRING"));
        assertTrue(S7StringVarLengthTag.matches("%DB1:0:STRING[2]"));
        assertFalse(S7StringVarLengthTag.matches("%DB1.DBW0:INT"));
        assertFalse(S7StringVarLengthTag.matches("%DB1.DB0:STRING(80)"));
    }

    @Test
    void varLengthEqualityAndHashCode() {
        S7StringVarLengthTag a = S7StringVarLengthTag.of("%DB1.DB0:STRING");
        S7StringVarLengthTag b = S7StringVarLengthTag.of("%DB1.DB0:STRING");
        S7StringVarLengthTag c = S7StringVarLengthTag.of("%DB1.DB0:WSTRING");
        assertEquals(a, b);
        assertNotEquals(a, c);
        assertEquals(a.hashCode(), b.hashCode());
        assertNotNull(a.toString());
    }

    @Test
    void varLengthToStringContainsType() {
        S7StringVarLengthTag tag = S7StringVarLengthTag.of("%DB1.DB0:WSTRING");
        assertNotNull(tag);
        assertTrue(tag.toString().contains("WSTRING"));
    }

    @Test
    void varLengthShortForm() {
        S7StringVarLengthTag tag = S7StringVarLengthTag.of("%DB1:0:STRING");
        assertNotNull(tag);
        assertEquals(TransportSize.STRING, tag.getDataType());
    }

    /**
     * The legacy count suffix {@code TYPE[n]} must parse with count semantics: {@code n} is the
     * number of elements, and the address renders back in the same shape.
     */
    @Test
    void postfixCountSuffixOnStrings() {
        S7Tag tag = S7Tag.of("%DB1:312:STRING[2]");
        assertInstanceOf(S7StringVarLengthTag.class, tag);
        assertEquals(2, tag.getNumberOfElements());
        assertEquals("%DB1.DB312:STRING[2]", tag.getAddressString());
    }

    @Test
    void postfixCountSuffixOnPlainTypes() {
        S7Tag tag = S7Tag.of("%DB1:36:DINT[2]");
        assertEquals(2, tag.getNumberOfElements());
        assertEquals("%DB1.DB36:DINT[2]", tag.getAddressString());
    }

    @Test
    void postfixBoolArrayNeedsNoBitOffset() {
        S7Tag tag = S7Tag.of("%DB1:2:BOOL[10]");
        assertEquals(10, tag.getNumberOfElements());
    }

    @Test
    void prefixRangeStillParsesAndRendersPrefix() {
        S7Tag tag = S7Tag.of("%DB1.DB36[0..1]:DINT");
        assertEquals(2, tag.getNumberOfElements());
        assertEquals("%DB1.DB36[0..1]:DINT", tag.getAddressString());
    }

    @Test
    void scalarStringHasNoCountSuffix() {
        assertEquals("%DB1.DB56:STRING", S7Tag.of("%DB1:56:STRING").getAddressString());
    }

    @Test
    void s7TagMatchesAcceptsVarLengthStrings() {
        assertTrue(S7Tag.matches("%DB69:68:STRING"));
        assertTrue(S7Tag.matches("%DB1.DB0[0..2]:WSTRING"));
        assertTrue(S7Tag.matches("%DB1:312:STRING[2]"));
        assertFalse(S7Tag.matches("not-a-tag"));
    }

    /**
     * Both entry points must return equal tags for the same address.
     */
    @Test
    void s7TagOfAgreesWithTagHandler() {
        S7PlcTagHandler handler = new S7PlcTagHandler();
        for (String address : new String[]{
            "%DB69:68:STRING",
            "%DB1.DB0:STRING",
            "%DB1.DB0[0..2]:STRING",
            "%DB1:312:STRING[2]",
            "%DB1:36:DINT[2]",
            "%MW0:INT",
            "%DB1.DBX0.0:BOOL"}) {
            assertEquals(handler.parseTag(address), S7Tag.of(address), address);
        }
    }

    @Test
    void s7TagOfStillRejectsGarbage() {
        assertThrows(org.apache.plc4x.java.api.exceptions.PlcInvalidTagException.class,
            () -> S7Tag.of("not-a-tag"));
    }

    @Test
    void tagHandlerRoutesToCorrectTagClass() {
        S7PlcTagHandler handler = new S7PlcTagHandler();
        assertInstanceOf(S7StringVarLengthTag.class, handler.parseTag("%DB1.DB0:STRING"));
        assertInstanceOf(S7StringVarLengthTag.class, handler.parseTag("%DB1:312:STRING[2]"));
        assertInstanceOf(S7Tag.class, handler.parseTag("%MW0:INT"));
        assertThrows(org.apache.plc4x.java.api.exceptions.PlcInvalidTagException.class,
            () -> handler.parseTag("not-a-tag"));
        // parseQuery returns null — there's no S7 query language; the connection's
        // onBrowse emits everything it knows regardless of the supplied query string.
        assertNull(handler.parseQuery("any"));
    }

    // ------------------------------------------------------------------------
    // YOFC stringEncoding extension: '|encoding' address suffix
    // ------------------------------------------------------------------------

    @Test
    void varLengthStringWithEncodingSuffix() {
        S7StringVarLengthTag tag = S7StringVarLengthTag.of("%DB1.DB0:STRING|GBK");
        assertNotNull(tag);
        assertEquals("GBK", tag.getStringEncoding());
    }

    @Test
    void wstringDefaultsToUtf16Encoding() {
        assertEquals("UTF-16", S7StringVarLengthTag.of("%DB1.DB0:WSTRING").getStringEncoding());
        assertNull(S7StringVarLengthTag.of("%DB1.DB0:STRING").getStringEncoding());
    }

    @Test
    void encodingParticipatesInEquality() {
        assertNotEquals(S7StringVarLengthTag.of("%DB1.DB0:STRING"),
            S7StringVarLengthTag.of("%DB1.DB0:STRING|GBK"));
        assertEquals(S7StringVarLengthTag.of("%DB1.DB0:STRING|GBK"),
            S7StringVarLengthTag.of("%DB1.DB0:STRING|GBK"));
    }

    @Test
    void s7TagOfParsesEncodingSuffix() {
        S7Tag tag = S7Tag.of("%DB1.DB0:STRING|GBK");
        assertInstanceOf(S7StringVarLengthTag.class, tag);
        assertEquals("GBK", tag.getStringEncoding());
    }

    @Test
    void encodingMatches() {
        assertTrue(S7StringVarLengthTag.matches("%DB1.DB0:STRING|GBK"));
        assertTrue(S7StringVarLengthTag.matches("%DB1:0:WSTRING|UTF-16"));
        assertFalse(S7StringVarLengthTag.matches("%DB1.DB0:STRING|GB K"));
    }

    @Test
    void encodingShowsUpInAddressString() {
        assertEquals("%DB1.DB0:STRING|GBK",
            S7StringVarLengthTag.of("%DB1.DB0:STRING|GBK").getAddressString());
        assertEquals("%DB1.DB312:STRING[2]|GBK",
            S7StringVarLengthTag.of("%DB1:312:STRING[2]|GBK").getAddressString());
    }
}
