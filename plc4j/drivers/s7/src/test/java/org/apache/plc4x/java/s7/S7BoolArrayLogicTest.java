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
package org.apache.plc4x.java.s7;

import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.s7.readwrite.MemoryArea;
import org.apache.plc4x.java.s7.readwrite.TransportSize;
import org.apache.plc4x.java.s7.tag.S7Tag;
import org.apache.plc4x.java.spi.values.PlcBOOL;
import org.apache.plc4x.java.spi.values.PlcList;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

/**
 * Unit tests for BOOL array bit-packing/unpacking and Read-Modify-Write logic
 * in {@link S7CotpConnection}.
 */
class S7BoolArrayLogicTest {

    // ═══════════════════════════════════════════════════════════════════════════
    // needsReadModifyWrite
    // ═══════════════════════════════════════════════════════════════════════════

    @Test
    void needsRmw_boolArrayNotMultipleOf8_returnsTrue() {
        // BOOL[10] → 10 bits, last byte has 6 unused bits → needs RMW
        S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 10);
        assertTrue(S7CotpConnection.needsReadModifyWrite(tag));
    }

    @Test
    void needsRmw_boolArrayMultipleOf8_returnsFalse() {
        // BOOL[16] → 16 bits fills 2 bytes exactly → no RMW needed
        S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 16);
        assertFalse(S7CotpConnection.needsReadModifyWrite(tag));
    }

    @Test
    void needsRmw_singleBool_returnsFalse() {
        // Single BOOL (numElements=1) → native bit transport, no RMW
        S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1);
        assertFalse(S7CotpConnection.needsReadModifyWrite(tag));
    }

    @Test
    void needsRmw_nonBoolType_returnsFalse() {
        // INT[3] → not BOOL, no RMW regardless of count
        S7Tag tag = new S7Tag(TransportSize.INT, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 3);
        assertFalse(S7CotpConnection.needsReadModifyWrite(tag));
    }

    @Test
    void needsRmw_boolArray1Element_returnsFalse() {
        // Edge case: only 1 element, n%8 != 0 but n <= 1
        S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1);
        assertFalse(S7CotpConnection.needsReadModifyWrite(tag));
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // unpackBits
    // ═══════════════════════════════════════════════════════════════════════════

    @Test
    void unpackBits_singleByte_correctOrder() {
        // 0b10110001 = 0xB1 → bits[0]=1, bits[1]=0, bits[2]=0, bits[3]=0,
        //                      bits[4]=1, bits[5]=1, bits[6]=0, bits[7]=1
        byte[] data = {(byte) 0xB1};
        boolean[] bits = S7CotpConnection.unpackBits(data, 8);
        assertTrue(bits[0]);   // bit 0
        assertFalse(bits[1]);  // bit 1
        assertFalse(bits[2]);  // bit 2
        assertFalse(bits[3]);  // bit 3
        assertTrue(bits[4]);   // bit 4
        assertTrue(bits[5]);   // bit 5
        assertFalse(bits[6]);  // bit 6
        assertTrue(bits[7]);   // bit 7
    }

    @Test
    void unpackBits_multipleBytes_spansCorrectly() {
        // 2 bytes, 10 bits: 0xFF 0x03 → all first 10 bits should be true
        byte[] data = {(byte) 0xFF, (byte) 0x03};
        boolean[] bits = S7CotpConnection.unpackBits(data, 10);
        assertEquals(10, bits.length);
        for (int i = 0; i < 10; i++) {
            assertTrue(bits[i], "bit " + i + " should be true");
        }
    }

    @Test
    void unpackBits_partialByte_onlyRequestedBits() {
        // 1 byte, but only 3 bits requested: 0b00000101 → [true, false, true]
        byte[] data = {0x05};
        boolean[] bits = S7CotpConnection.unpackBits(data, 3);
        assertEquals(3, bits.length);
        assertTrue(bits[0]);
        assertFalse(bits[1]);
        assertTrue(bits[2]);
    }

    @Test
    void unpackBits_allZeros() {
        byte[] data = {0x00, 0x00};
        boolean[] bits = S7CotpConnection.unpackBits(data, 12);
        for (int i = 0; i < 12; i++) {
            assertFalse(bits[i]);
        }
    }

    @Test
    void unpackBits_allOnes() {
        byte[] data = {(byte) 0xFF, (byte) 0xFF};
        boolean[] bits = S7CotpConnection.unpackBits(data, 16);
        for (int i = 0; i < 16; i++) {
            assertTrue(bits[i]);
        }
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // mergeBits
    // ═══════════════════════════════════════════════════════════════════════════

    @Test
    void mergeBits_preservesUnusedBitsInLastByte() {
        // Scenario: BOOL[10] at some offset. Current PLC state: both bytes = 0xFF.
        // Write: all 10 bits to false. After merge, bits 10-15 in byte[1] must stay 1.
        S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 10);
        byte[] currentBytes = {(byte) 0xFF, (byte) 0xFF};

        // Write all false
        List<PlcValue> values = new ArrayList<>();
        for (int i = 0; i < 10; i++) {
            values.add(PlcBOOL.of(false));
        }
        PlcList plcList = new PlcList(values);

        byte[] merged = S7CotpConnection.mergeBits(currentBytes, tag, plcList);

        // Byte 0: all 8 bits cleared → 0x00
        assertEquals((byte) 0x00, merged[0]);
        // Byte 1: bits 0-1 cleared (elements 8-9), bits 2-7 preserved (0xFC)
        assertEquals((byte) 0xFC, merged[1]);
    }

    @Test
    void mergeBits_setsSpecificBits() {
        // Current: all zeros. Write: bits 0, 4, 9 to true, rest false.
        S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 10);
        byte[] currentBytes = {0x00, 0x00};

        List<PlcValue> values = new ArrayList<>();
        for (int i = 0; i < 10; i++) {
            values.add(PlcBOOL.of(i == 0 || i == 4 || i == 9));
        }
        PlcList plcList = new PlcList(values);

        byte[] merged = S7CotpConnection.mergeBits(currentBytes, tag, plcList);

        // Byte 0: bit0=1, bit4=1 → 0b00010001 = 0x11
        assertEquals((byte) 0x11, merged[0]);
        // Byte 1: bit1(element 9)=1 → 0b00000010 = 0x02
        assertEquals((byte) 0x02, merged[1]);
    }

    @Test
    void mergeBits_clearsBitsWhilePreservingOthers() {
        // Current: 0xFF 0xFF. Write: only bit 0 set, bits 1-9 cleared.
        // Unused bits 10-15 should stay as 1.
        S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 10);
        byte[] currentBytes = {(byte) 0xFF, (byte) 0xFF};

        List<PlcValue> values = new ArrayList<>();
        for (int i = 0; i < 10; i++) {
            values.add(PlcBOOL.of(i == 0));
        }
        PlcList plcList = new PlcList(values);

        byte[] merged = S7CotpConnection.mergeBits(currentBytes, tag, plcList);

        // Byte 0: only bit 0 set → 0x01
        assertEquals((byte) 0x01, merged[0]);
        // Byte 1: bits 0-1 cleared (elements 8-9), bits 2-7 preserved → 0xFC
        assertEquals((byte) 0xFC, merged[1]);
    }

    @Test
    void mergeBits_fullByte_noPreservation() {
        // BOOL[8] → exactly fills 1 byte, no unused bits.
        // Note: needsReadModifyWrite returns false for n%8==0, but mergeBits should still
        // work correctly if called — it simply overwrites all bits.
        S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 8);
        byte[] currentBytes = {(byte) 0xAA};  // 10101010

        List<PlcValue> values = new ArrayList<>();
        for (int i = 0; i < 8; i++) {
            values.add(PlcBOOL.of(i % 2 == 0));  // 01010101 = 0x55
        }
        PlcList plcList = new PlcList(values);

        byte[] merged = S7CotpConnection.mergeBits(currentBytes, tag, plcList);
        assertEquals((byte) 0x55, merged[0]);
    }

    @Test
    void mergeBits_3bits_preserves5() {
        // BOOL[3]: only 3 bits used in the byte. Bits 3-7 must be preserved.
        S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 3);
        byte[] currentBytes = {(byte) 0xF8};  // 11111000 — bits 3-7 set

        List<PlcValue> values = new ArrayList<>();
        values.add(PlcBOOL.of(true));   // bit 0
        values.add(PlcBOOL.of(true));   // bit 1
        values.add(PlcBOOL.of(false));  // bit 2
        PlcList plcList = new PlcList(values);

        byte[] merged = S7CotpConnection.mergeBits(currentBytes, tag, plcList);
        // bits 0-1 set, bit 2 clear, bits 3-7 preserved → 11111011 = 0xFB
        assertEquals((byte) 0xFB, merged[0]);
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Round-trip consistency: pack → unpack
    // ═══════════════════════════════════════════════════════════════════════════

    @Test
    void packUnpack_roundTrip_consistent() {
        // Write bits [T, F, T, T, F, F, T, F, T, T] into zeroed bytes using mergeBits,
        // then unpack and verify we get the same pattern back.
        boolean[] pattern = {true, false, true, true, false, false, true, false, true, true};
        S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 10);
        byte[] zeroBytes = {0x00, 0x00};

        List<PlcValue> values = new ArrayList<>();
        for (boolean b : pattern) {
            values.add(PlcBOOL.of(b));
        }
        PlcList plcList = new PlcList(values);

        byte[] packed = S7CotpConnection.mergeBits(zeroBytes, tag, plcList);
        boolean[] unpacked = S7CotpConnection.unpackBits(packed, 10);

        assertArrayEquals(pattern, unpacked);
    }
}
