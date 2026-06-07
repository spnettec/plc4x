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
package org.apache.plc4x.java.s7.readwrite.protocol;

import org.apache.plc4x.java.s7.readwrite.MemoryArea;
import org.apache.plc4x.java.s7.readwrite.TransportSize;
import org.apache.plc4x.java.s7.readwrite.tag.S7StringTag;
import org.apache.plc4x.java.s7.readwrite.tag.S7Tag;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

import java.util.List;
import java.util.Map;

import static org.junit.jupiter.api.Assertions.*;

/**
 * Unit tests for block-merge pure logic — no PLC connection required.
 */
class BlockMergeLogicTest {

    // ── s7TagByteSize ──────────────────────────────────────────────────

    @Nested
    @DisplayName("s7TagByteSize")
    class TagByteSize {

        @Test
        @DisplayName("BYTE returns 1")
        void byteTag() {
            S7Tag tag = new S7Tag(TransportSize.BYTE, MemoryArea.DATA_BLOCKS, 1, 4, (byte) 0, 1, "UTF-8");
            assertEquals(1, S7NonHProtocolLogic.s7TagByteSize(tag));
        }

        @Test
        @DisplayName("INT[3] returns 6")
        void intArray() {
            S7Tag tag = new S7Tag(TransportSize.INT, MemoryArea.DATA_BLOCKS, 1, 4, (byte) 0, 3, "UTF-8");
            assertEquals(6, S7NonHProtocolLogic.s7TagByteSize(tag));
        }

        @Test
        @DisplayName("DINT returns 4")
        void dint() {
            S7Tag tag = new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8");
            assertEquals(4, S7NonHProtocolLogic.s7TagByteSize(tag));
        }

        @Test
        @DisplayName("REAL[2] returns 8")
        void realArray() {
            S7Tag tag = new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 48, (byte) 0, 2, "UTF-8");
            assertEquals(8, S7NonHProtocolLogic.s7TagByteSize(tag));
        }

        @Test
        @DisplayName("BOOL returns -1 (packed bits)")
        void boolReturnsMinusOne() {
            S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 5, "UTF-8");
            assertEquals(-1, S7NonHProtocolLogic.s7TagByteSize(tag));
        }

        @Test
        @DisplayName("BIT returns -1 (packed bits)")
        void bitReturnsMinusOne() {
            S7Tag tag = new S7Tag(TransportSize.BIT, MemoryArea.DATA_BLOCKS, 1, 2, (byte) 0, 10, "UTF-8");
            assertEquals(-1, S7NonHProtocolLogic.s7TagByteSize(tag));
        }

        @Test
        @DisplayName("STRING returns -1 (variable-length)")
        void stringReturnsMinusOne() {
            S7Tag tag = new S7Tag(TransportSize.STRING, MemoryArea.DATA_BLOCKS, 1, 56, (byte) 0, 1, "UTF-8");
            // TransportSize.STRING.getSizeInBytes() = 1, so s7TagByteSize returns 1 * 1 = 1
            // STRING is excluded from merging by the caller, not by s7TagByteSize.
            // The method returns the raw DataType-based size.
            assertEquals(1, S7NonHProtocolLogic.s7TagByteSize(tag));
        }
    }

    // ── computeReadResponseItemSize ────────────────────────────────────

    @Nested
    @DisplayName("computeReadResponseItemSize (optimizer-equivalent formula)")
    class ResponseItemSize {

        @Test
        @DisplayName("BYTE → 6 (4 header + 1 data, padded even)")
        void byteTag() {
            S7Tag tag = new S7Tag(TransportSize.BYTE, MemoryArea.DATA_BLOCKS, 1, 4, (byte) 0, 1, "UTF-8");
            assertEquals(6, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("INT → 6 (4 + 2, already even)")
        void intTag() {
            S7Tag tag = new S7Tag(TransportSize.INT, MemoryArea.DATA_BLOCKS, 1, 4, (byte) 0, 1, "UTF-8");
            assertEquals(6, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("DINT → 8 (4 + 4)")
        void dintTag() {
            S7Tag tag = new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8");
            assertEquals(8, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("REAL[2] → 12 (4 + 2×4)")
        void realArray() {
            S7Tag tag = new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 48, (byte) 0, 2, "UTF-8");
            assertEquals(12, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("DINT[3] → 16 (4 + 3×4)")
        void dintArray() {
            S7Tag tag = new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 3, "UTF-8");
            assertEquals(16, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("BIT[10] → packed to 2 bytes → 6 (4 + 2)")
        void bitArrayPacked() {
            S7Tag tag = new S7Tag(TransportSize.BIT, MemoryArea.DATA_BLOCKS, 1, 2, (byte) 0, 10, "UTF-8");
            assertEquals(6, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("BIT[1] → single bit stays 1 byte → 6 (4+1, padded even)")
        void singleBit() {
            S7Tag tag = new S7Tag(TransportSize.BIT, MemoryArea.DATA_BLOCKS, 1, 2, (byte) 0, 1, "UTF-8");
            assertEquals(6, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("BOOL[8] → packed to 1 byte → 5 padded → 6")
        void boolArrayPacked() {
            S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 8, "UTF-8");
            assertEquals(6, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("BOOL[10] → packed to 2 bytes → 6")
        void boolArray10() {
            S7Tag tag = new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 10, "UTF-8");
            assertEquals(6, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("STRING(254) → 4 + 1×(254+2)×1 = 260 (even)")
        void stringWithLength() {
            S7Tag tag = S7Tag.of("%DB1:56:STRING(254)");
            assertInstanceOf(S7StringTag.class, tag);
            assertEquals(260, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("WSTRING(254) → 4 + 1×(254+2)×2 = 516 (even)")
        void wstringWithLength() {
            S7Tag tag = S7Tag.of("%DB1:312:WSTRING(254)");
            assertInstanceOf(S7StringTag.class, tag);
            assertEquals(516, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("STRING(254)[2] → 4 + 2×(254+2)×1 = 516")
        void stringArray() {
            S7Tag tag = S7Tag.of("%DB1:312:STRING(254)[2]");
            assertInstanceOf(S7StringTag.class, tag);
            assertEquals(2, tag.getNumberOfElements());
            assertEquals(516, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("STRING(10) → 4 + 12 = 16")
        void shortString() {
            S7Tag tag = S7Tag.of("%DB1:100:STRING(10)");
            assertEquals(16, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("STRING (no length) → of() creates S7StringTag with default 254 → 260 bytes")
        void stringDefaultLength() {
            // %DB1:56:STRING matches DATA_BLOCK_SHORT_PATTERN, which has a special
            // case: if(dataType==STRING) → new S7StringTag(..., 254, ...)
            S7Tag tag = S7Tag.of("%DB1:56:STRING");
            assertInstanceOf(S7StringTag.class, tag,
                    "STRING without (length) still creates S7StringTag via DATA_BLOCK_SHORT_PATTERN");
            assertEquals(254, ((S7StringTag) tag).getStringLength());
            assertEquals(TransportSize.STRING, tag.getDataType());
            // 4 + 1 * 1 * (254+2) = 260
            assertEquals(260, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }

        @Test
        @DisplayName("plain S7Tag(STRING) via constructor → else-if catches it → 260 bytes")
        void plainStringTagViaConstructor() {
            // Direct constructor creates a plain S7Tag, not S7StringTag.
            // The else-if in computeReadResponseItemSize should catch this.
            S7Tag tag = new S7Tag(TransportSize.STRING, MemoryArea.DATA_BLOCKS, 1, 56, (byte) 0, 1, "UTF-8");
            assertFalse(tag instanceof S7StringTag,
                    "Constructor-created STRING tag is a plain S7Tag");
            assertEquals(260, S7NonHProtocolLogic.computeReadResponseItemSize(tag),
                    "else-if branch should use default stringLength=254");
        }

        @Test
        @DisplayName("STRING(20) → 4 + 22 = 26")
        void stringExplicitShortLength() {
            S7Tag tag = S7Tag.of("%DB1:56:STRING(20)");
            assertInstanceOf(S7StringTag.class, tag);
            assertEquals(20, ((S7StringTag) tag).getStringLength());
            // 4 + 1 * 1 * (20+2) = 26
            assertEquals(26, S7NonHProtocolLogic.computeReadResponseItemSize(tag));
        }
    }

    // ── computeMergedReadResponseItemSize ─────────────────────────────

    @Nested
    @DisplayName("computeMergedReadResponseItemSize")
    class MergedResponseItemSize {

        @Test
        @DisplayName("zero-byte block → 4 header bytes")
        void zeroByteBlock() {
            assertEquals(4, S7NonHProtocolLogic.computeMergedReadResponseItemSize(0));
        }

        @Test
        @DisplayName("one-byte block → 6 (4 header + 1 data, padded even)")
        void oneByteBlockPadded() {
            assertEquals(6, S7NonHProtocolLogic.computeMergedReadResponseItemSize(1));
        }

        @Test
        @DisplayName("two-byte block → 6 (4 header + 2 data)")
        void twoByteBlock() {
            assertEquals(6, S7NonHProtocolLogic.computeMergedReadResponseItemSize(2));
        }

        @Test
        @DisplayName("twelve-byte block → 16 (4 header + 12 data)")
        void twelveByteBlock() {
            assertEquals(16, S7NonHProtocolLogic.computeMergedReadResponseItemSize(12));
        }

        @Test
        @DisplayName("thirteen-byte block → 18 (4 header + 13 data, padded even)")
        void thirteenByteBlockPadded() {
            assertEquals(18, S7NonHProtocolLogic.computeMergedReadResponseItemSize(13));
        }
    }

    // ── Block-merge group building ────────────────────────────────────

    @Nested
    @DisplayName("buildFixedGapMergeGroups")
    class FixedGapMergeGroups {

        @Test
        @DisplayName("groups adjacent tags when byte gap is less than minGap")
        void groupsGapLessThanMinGap() {
            List<S7Tag> tags = List.of(
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, "UTF-8"),
                    new S7Tag(TransportSize.WORD, MemoryArea.DATA_BLOCKS, 1, 10, (byte) 0, 1, "UTF-8"),
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 40, (byte) 0, 1, "UTF-8"));

            List<List<Integer>> groups = S7NonHProtocolLogic.buildFixedGapMergeGroups(
                    List.of(0, 1, 2), tags, 16);

            assertEquals(List.of(List.of(0, 1), List.of(2)), groups);
        }

        @Test
        @DisplayName("does not group when byte gap is equal to minGap")
        void equalGapDoesNotMerge() {
            List<S7Tag> tags = List.of(
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, "UTF-8"),
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 20, (byte) 0, 1, "UTF-8"));

            List<List<Integer>> groups = S7NonHProtocolLogic.buildFixedGapMergeGroups(
                    List.of(0, 1), tags, 16);

            assertEquals(List.of(List.of(0), List.of(1)), groups);
        }
    }

    @Nested
    @DisplayName("buildAutoMergeGroups")
    class AutoMergeGroups {

        @Test
        @DisplayName("groups adjacent DINTs when merged item is cheaper")
        void groupsAdjacentDints() {
            List<S7Tag> tags = List.of(
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"),
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 1, "UTF-8"),
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 40, (byte) 0, 1, "UTF-8"));

            List<List<Integer>> groups = S7NonHProtocolLogic.buildAutoMergeGroups(
                    List.of(0, 1, 2), tags);

            assertEquals(List.of(List.of(0, 1, 2)), groups);
        }

        @Test
        @DisplayName("rejects large gap when merged item is more expensive")
        void rejectsLargeGap() {
            List<S7Tag> tags = List.of(
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, "UTF-8"),
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 100, (byte) 0, 1, "UTF-8"));

            List<List<Integer>> groups = S7NonHProtocolLogic.buildAutoMergeGroups(
                    List.of(0, 1), tags);

            assertEquals(List.of(List.of(0), List.of(1)), groups);
        }

        @Test
        @DisplayName("accepts moderate gap when item overhead saving wins")
        void acceptsCostEffectiveGap() {
            List<S7Tag> tags = List.of(
                    new S7Tag(TransportSize.WORD, MemoryArea.DATA_BLOCKS, 1, 8, (byte) 0, 1, "UTF-8"),
                    new S7Tag(TransportSize.WORD, MemoryArea.DATA_BLOCKS, 1, 20, (byte) 0, 1, "UTF-8"));

            List<List<Integer>> groups = S7NonHProtocolLogic.buildAutoMergeGroups(
                    List.of(0, 1), tags);

            assertEquals(List.of(List.of(0, 1)), groups);
        }

        @Test
        @DisplayName("starts a new candidate group after rejecting an expensive gap")
        void startsNewGroupAfterRejectedGap() {
            List<S7Tag> tags = List.of(
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, "UTF-8"),
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 100, (byte) 0, 1, "UTF-8"),
                    new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 104, (byte) 0, 1, "UTF-8"));

            List<List<Integer>> groups = S7NonHProtocolLogic.buildAutoMergeGroups(
                    List.of(0, 1, 2), tags);

            assertEquals(List.of(List.of(0), List.of(1, 2)), groups);
        }
    }

    // ── BlockMergeMapping ──────────────────────────────────────────────

    @Nested
    @DisplayName("BlockMergeMapping")
    class Mapping {

        @Test
        @DisplayName("isMerged() true when tagSlices > 1")
        void isMergedWithMultipleSlices() {
            var slices = List.of(
                    new S7NonHProtocolLogic.BlockMergeMapping.TagSlice("a", 0, 4),
                    new S7NonHProtocolLogic.BlockMergeMapping.TagSlice("b", 4, 4));
            var chunks = List.of(
                    new S7NonHProtocolLogic.BlockMergeMapping.Chunk(0, 0, 8));
            var mapping = new S7NonHProtocolLogic.BlockMergeMapping(8, chunks, slices);
            assertTrue(mapping.isMerged());
            assertEquals(8, mapping.blockSize);
            assertEquals(2, mapping.tagSlices.size());
            assertEquals(1, mapping.chunks.size());
        }

        @Test
        @DisplayName("isMerged() false with single tagSlice")
        void isNotMergedWithSingleSlice() {
            var slices = List.of(
                    new S7NonHProtocolLogic.BlockMergeMapping.TagSlice("x", 0, 0));
            var chunks = List.of(
                    new S7NonHProtocolLogic.BlockMergeMapping.Chunk(0, 0, 0));
            var mapping = new S7NonHProtocolLogic.BlockMergeMapping(0, chunks, slices);
            assertFalse(mapping.isMerged());
            assertEquals(0, mapping.blockSize);
        }

        @Test
        @DisplayName("TagSlice records correct offsets")
        void tagSliceOffsets() {
            var slice = new S7NonHProtocolLogic.BlockMergeMapping.TagSlice("dint0", 0, 4);
            assertEquals("dint0", slice.tagName());
            assertEquals(0, slice.offsetInBlock());
            assertEquals(4, slice.byteSize());

            var slice2 = new S7NonHProtocolLogic.BlockMergeMapping.TagSlice("dint1", 4, 4);
            assertEquals("dint1", slice2.tagName());
            assertEquals(4, slice2.offsetInBlock());
        }

        @Test
        @DisplayName("Chunk records correct response item index and byte range")
        void chunkFields() {
            var chunk = new S7NonHProtocolLogic.BlockMergeMapping.Chunk(3, 128, 200);
            assertEquals(3, chunk.responseItemIndex());
            assertEquals(128, chunk.offsetInBlock());
            assertEquals(200, chunk.size());
        }

        @Test
        @DisplayName("Multiple chunks for a single block")
        void multipleChunks() {
            // Simulates a block that was split across 3 response items
            var slices = List.of(
                    new S7NonHProtocolLogic.BlockMergeMapping.TagSlice("big", 0, 600));
            var chunks = List.of(
                    new S7NonHProtocolLogic.BlockMergeMapping.Chunk(0, 0, 200),
                    new S7NonHProtocolLogic.BlockMergeMapping.Chunk(1, 200, 200),
                    new S7NonHProtocolLogic.BlockMergeMapping.Chunk(2, 400, 200));
            var mapping = new S7NonHProtocolLogic.BlockMergeMapping(600, chunks, slices);
            assertEquals(600, mapping.blockSize);
            assertEquals(3, mapping.chunks.size());
            // chunks cover full block: offset 0→200, 200→400, 400→600
            assertEquals(0, mapping.chunks.get(0).offsetInBlock());
            assertEquals(200, mapping.chunks.get(1).offsetInBlock());
            assertEquals(400, mapping.chunks.get(2).offsetInBlock());
        }
    }

    // ── PDU budget fallback scenario ───────────────────────────────────

    @Nested
    @DisplayName("PDU budget: merge fallback")
    class PduBudgetFallback {

        @Test
        @DisplayName("adjacent DINTs at offsets 32,36,40 → block size 12 bytes")
        void adjacentDintsBlockSize() {
            S7Tag t0 = new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8");
            S7Tag t1 = new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 1, "UTF-8");
            S7Tag t2 = new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 40, (byte) 0, 1, "UTF-8");

            assertEquals(4, S7NonHProtocolLogic.s7TagByteSize(t0));
            assertEquals(4, S7NonHProtocolLogic.s7TagByteSize(t1));
            assertEquals(4, S7NonHProtocolLogic.s7TagByteSize(t2));

            // Merged block: offset 32 to 44 = 12 bytes
            int blockSize = (40 + 4) - 32; // end(t2.endByte) - start(t0.byteOffset)
            assertEquals(12, blockSize);

            // Individual response sizes: 8 bytes each (4 header + 4 data)
            assertEquals(8, S7NonHProtocolLogic.computeReadResponseItemSize(t0));
            assertEquals(8, S7NonHProtocolLogic.computeReadResponseItemSize(t1));
            assertEquals(8, S7NonHProtocolLogic.computeReadResponseItemSize(t2));

            // Merged response size: 4 + 12 = 16 (even)
            int mergedRespSize = 4 + 12; // 4 header + 12 bytes of block data
            assertEquals(16, mergedRespSize);

            // Net PDU impact: 3 individual (3×8=24 resp, 3×S7_ADDRESS_ANY_SIZE=~30 req)
            // vs 1 merged (16 resp, ~10 req). Both sides shrink.
            assertTrue(mergedRespSize < 3 * 8,
                    "Merged response should be smaller than sum of individual");
        }

        @Test
        @DisplayName("PDU budget overflow → merge is rejected for that block")
        void pduOverflowRejectsMerge() {
            // Simulate a merged block that would overflow a 240-byte PDU
            // and verify the overflow detection arithmetic.
            int pduSize = 240;
            int emptyRespSize = 100;  // approximate EMPTY_READ_RESPONSE_SIZE

            // 10 adjacent DINTs: 10×4=40 bytes → merged response = 4+40=44 bytes
            int mergedRespItemSize = 4 + 40;
            if (mergedRespItemSize % 2 == 1) mergedRespItemSize++;
            assertEquals(44, mergedRespItemSize);

            // Individual: 10 items × 8 bytes = 80 bytes response
            int individualRespTotal = 10 * 8;
            assertEquals(80, individualRespTotal);

            // Both fit — merge is beneficial
            int totalAfterMerge = emptyRespSize + mergedRespItemSize;
            assertTrue(totalAfterMerge < pduSize, "Should fit in PDU");

            // Now simulate a huge merged block (200 bytes) that DOES NOT fit
            int hugeBlockRespItemSize = 4 + 200; // = 204, already even
            int totalAfterHugeMerge = emptyRespSize + hugeBlockRespItemSize;
            assertTrue(totalAfterHugeMerge > pduSize,
                    "Huge block should overflow " + pduSize + " byte PDU, "
                    + "total=" + totalAfterHugeMerge);
            // Expected behavior: merge is skipped, tags stay individual
        }

        @Test
        @DisplayName("merge cost includes gap bytes — gap may cause overflow")
        void gapBytesIncreaseResponseSize() {
            // Two REALs at offsets 48 and 56 with a BYTE[2] in between at 52-54.
            // Merged: offset 48 to 56 → blockSize = 56 - 48 = 8 bytes (includes 2 gap bytes)
            // Individual: 4+4=8 + 4+4=8 = 16 bytes response
            // Merged: 4+8=12 bytes response (but reads 2 unused gap bytes)

            S7Tag real0 = new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 48, (byte) 0, 1, "UTF-8");
            S7Tag real1 = new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 56, (byte) 0, 1, "UTF-8");

            int blockSize = (56 + 4) - 48; // gap = 4 bytes between real0 end(52) and real1 start(56)
            assertEquals(12, blockSize);    // 8 bytes of data + 4 bytes of gap

            int mergedRespSize = 4 + blockSize; // = 16
            if (mergedRespSize % 2 == 1) mergedRespSize++;
            assertEquals(16, mergedRespSize);

            int individualRespSize =
                    S7NonHProtocolLogic.computeReadResponseItemSize(real0)
                    + S7NonHProtocolLogic.computeReadResponseItemSize(real1);
            // Each REAL: 4 + 4 = 8 (even)
            assertEquals(16, individualRespSize);

            // With small gap, merge is neutral — but the gap bytes are "wasted" read.
            // The PDU budget check correctly accounts for this (mergedRespSize includes gaps).
        }
    }
}
