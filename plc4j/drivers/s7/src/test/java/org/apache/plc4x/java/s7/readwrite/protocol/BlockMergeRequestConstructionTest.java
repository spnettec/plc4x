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

import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.s7.readwrite.MemoryArea;
import org.apache.plc4x.java.s7.readwrite.S7AddressAny;
import org.apache.plc4x.java.s7.readwrite.S7VarRequestParameterItem;
import org.apache.plc4x.java.s7.readwrite.S7VarRequestParameterItemAddress;
import org.apache.plc4x.java.s7.readwrite.TransportSize;
import org.apache.plc4x.java.s7.readwrite.configuration.S7Configuration;
import org.apache.plc4x.java.s7.readwrite.context.S7DriverContext;
import org.apache.plc4x.java.s7.readwrite.tag.S7Tag;
import org.apache.plc4x.java.spi.messages.DefaultPlcReadRequest;
import org.apache.plc4x.java.spi.messages.PlcReader;
import org.apache.plc4x.java.spi.messages.utils.DefaultPlcTagItem;
import org.apache.plc4x.java.spi.messages.utils.PlcTagItem;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;

import java.lang.reflect.Field;
import java.util.*;
import java.util.concurrent.CompletableFuture;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.*;

/**
 * Tests that drive {@code performOrdinaryReadRequest} via Mockito spy to verify
 * request-item construction, block-merge mapping (returned via
 * {@code ReadRequestContext}), and PDU budget behaviour. No PLC connection required.
 */
class BlockMergeRequestConstructionTest {

    private S7NonHProtocolLogic logic;
    private S7Configuration config;
    private S7DriverContext driverCtx;

    private ArgumentCaptor<List<S7VarRequestParameterItem>> itemsCaptor;

    @BeforeEach
    @SuppressWarnings("unchecked")
    void setUp() throws Exception {
        driverCtx = mock(S7DriverContext.class);
        when(driverCtx.getPduSize()).thenReturn(240);

        config = new S7Configuration();
        config.gap = 16;

        // spy() calls the real constructor → tpduGenerator is initialized
        logic = spy(new S7NonHProtocolLogic());
        logic.setConfiguration(config);

        Field ctxField = S7NonHProtocolLogic.class.getDeclaredField("s7DriverContext");
        ctxField.setAccessible(true);
        ctxField.set(logic, driverCtx);

        itemsCaptor = ArgumentCaptor.forClass(List.class);
        doReturn(CompletableFuture.completedFuture(null))
                .when(logic)
                .sendReadMessage(any(DefaultPlcReadRequest.class), itemsCaptor.capture(), anyInt());
    }

    /** Build a DefaultPlcReadRequest with the given S7 tags. */
    private static DefaultPlcReadRequest readRequest(Map<String, S7Tag> tags) {
        PlcReader reader = mock(PlcReader.class);
        LinkedHashMap<String, PlcTagItem<PlcTag>> items = new LinkedHashMap<>();
        tags.forEach((name, tag) -> items.put(name, new DefaultPlcTagItem<>(tag)));
        return new DefaultPlcReadRequest(reader, items);
    }

    /** Cast S7Address to S7AddressAny for field access. */
    private static S7AddressAny asAny(S7VarRequestParameterItem item) {
        return (S7AddressAny) ((S7VarRequestParameterItemAddress) item).getAddress();
    }

    @Nested
    @DisplayName("Block merge: adjacent tags")
    class AdjacentMerge {

        @Test
        @DisplayName("3 adjacent DINTs merge into 1 block item → 1 request item")
        void threeAdjacentDintsMergeToOneItem() {
            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("dint0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            tags.put("dint1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 1, "UTF-8"));
            tags.put("dint2", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 40, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            S7NonHProtocolLogic.ReadRequestContext ctx = logic.performOrdinaryReadRequest(request);
            assertNotNull(ctx.blockMapping(), "blockMapping should be non-null for gap>0");

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(1, items.size(), "3 adjacent DINTs should merge into 1 block item");

            S7AddressAny addr = asAny(items.get(0));
            assertEquals(TransportSize.BYTE, addr.getTransportSize());
            assertEquals(12, addr.getNumberOfElements(), "Block should be 3×4 = 12 bytes");
            assertEquals(32, addr.getByteAddress());

            // Verify mapping via ReadRequestContext
            Map<Integer, S7NonHProtocolLogic.BlockMergeMapping> mapping = ctx.blockMapping();
            assertEquals(1, mapping.size());
            S7NonHProtocolLogic.BlockMergeMapping bm = mapping.values().iterator().next();
            assertTrue(bm.isMerged());
            assertEquals(3, bm.tagSlices.size());
            assertEquals("dint0", bm.tagSlices.get(0).tagName());
            assertEquals(0, bm.tagSlices.get(0).offsetInBlock());
            assertEquals(4, bm.tagSlices.get(0).byteSize());
            assertEquals("dint1", bm.tagSlices.get(1).tagName());
            assertEquals(4, bm.tagSlices.get(1).offsetInBlock());
            assertEquals("dint2", bm.tagSlices.get(2).tagName());
            assertEquals(8, bm.tagSlices.get(2).offsetInBlock());
        }

        @Test
        @DisplayName("Mixed adjacent types (DINT,REAL,INT,DINT) merge into 1 block")
        void mixedAdjacentTypesMerge() {
            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("d0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            tags.put("r0", new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 1, "UTF-8"));
            tags.put("i0", new S7Tag(TransportSize.INT,  MemoryArea.DATA_BLOCKS, 1, 40, (byte) 0, 1, "UTF-8"));
            tags.put("d1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 42, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            S7NonHProtocolLogic.ReadRequestContext ctx = logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(1, items.size(), "4 adjacent tags should merge into 1 block");

            S7AddressAny addr = asAny(items.get(0));
            assertEquals(14, addr.getNumberOfElements(), "DINT(4)+REAL(4)+INT(2)+DINT(4)=14 bytes");

            S7NonHProtocolLogic.BlockMergeMapping bm = ctx.blockMapping().values().iterator().next();
            assertEquals(4, bm.tagSlices.size());
        }
    }

    @Nested
    @DisplayName("Block merge: non-mergeable tags")
    class NonMergeableTags {

        @Test
        @DisplayName("Scattered tags (gap > minGap) stay as individual items")
        void scatteredTagsNotMerged() {
            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("d0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            tags.put("d1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 100, (byte) 0, 1, "UTF-8"));
            tags.put("d2", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 200, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            S7NonHProtocolLogic.ReadRequestContext ctx = logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(3, items.size(), "Scattered tags should stay as 3 individual items");

            for (S7VarRequestParameterItem item : items) {
                assertEquals(TransportSize.DINT, asAny(item).getTransportSize());
            }

            Map<Integer, S7NonHProtocolLogic.BlockMergeMapping> mapping = ctx.blockMapping();
            assertEquals(3, mapping.size());
            for (S7NonHProtocolLogic.BlockMergeMapping bm : mapping.values()) {
                assertFalse(bm.isMerged());
                assertEquals(0, bm.blockSize);
                assertEquals(1, bm.tagSlices.size());
            }
        }

        @Test
        @DisplayName("BOOL tags are excluded from merge, stay individual")
        void boolTagsNotMerged() {
            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("b0", new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, "UTF-8"));
            tags.put("b1", new S7Tag(TransportSize.BOOL, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 1, 1, "UTF-8"));
            tags.put("d0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(3, items.size(), "2 BOOL + 1 DINT → 3 individual items");

            long boolItems = 0;
            for (S7VarRequestParameterItem item : items) {
                if (asAny(item).getTransportSize() == TransportSize.BOOL) boolItems++;
            }
            assertEquals(2, boolItems);
        }

        @Test
        @DisplayName("BIT array tags are excluded from merge")
        void bitTagsNotMerged() {
            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("bits", new S7Tag(TransportSize.BIT, MemoryArea.DATA_BLOCKS, 1, 2, (byte) 0, 10, "UTF-8"));
            tags.put("d0",   new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(2, items.size(), "BIT[10] + DINT → 2 individual items");
        }

        @Test
        @DisplayName("STRING tags stay individual, don't distort PDU baseline")
        void stringTagsNotMerged() {
            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("d0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            tags.put("str", S7Tag.of("%DB1:56:STRING(254)"));
            tags.put("d1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 312, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(3, items.size(), "STRING excluded → all 3 individual");
        }

        @Test
        @DisplayName("Multiple independent merge groups in same request")
        void independentMergeGroups() {
            Map<String, S7Tag> tags = new LinkedHashMap<>();
            // Group A: 3 DINTs at 32-44
            tags.put("a0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            tags.put("a1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 1, "UTF-8"));
            tags.put("a2", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 40, (byte) 0, 1, "UTF-8"));
            // Group B: 2 REALs at 48-56 (gap 4 bytes from group A, minGap=16 covers it)
            tags.put("b0", new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 48, (byte) 0, 1, "UTF-8"));
            tags.put("b1", new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 52, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            S7NonHProtocolLogic.ReadRequestContext ctx = logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(1, items.size(), "All 5 tags within gap<16 → 1 merged block");

            assertEquals(1, ctx.blockMapping().size());
            S7NonHProtocolLogic.BlockMergeMapping bm = ctx.blockMapping().values().iterator().next();
            assertTrue(bm.isMerged());
            assertEquals(5, bm.tagSlices.size());
        }
    }

    @Nested
    @DisplayName("Block merge disabled")
    class GapDisabled {

        @Test
        @DisplayName("gap=0 returns null blockMapping")
        void gapZeroNoMerge() {
            config.gap = 0;
            logic.setConfiguration(config);

            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("d0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            tags.put("d1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 1, "UTF-8"));
            tags.put("d2", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 40, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            S7NonHProtocolLogic.ReadRequestContext ctx = logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(3, items.size(), "gap=0 → 3 individual items, no merge");

            assertNull(ctx.blockMapping(), "blockMapping should be null when block merging is disabled");
        }
    }

    @Nested
    @DisplayName("Block merge: automatic gap detection")
    class AutoGapDetection {

        @Test
        @DisplayName("gap=-1 auto mode merges adjacent tags")
        void autoMergeWorksWithGapMinusOne() {
            config.gap = -1;
            logic.setConfiguration(config);

            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("dint0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            tags.put("dint1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 1, "UTF-8"));
            tags.put("dint2", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 40, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            S7NonHProtocolLogic.ReadRequestContext ctx = logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(1, items.size(), "auto mode should merge adjacent DINTs with gap=-1");

            S7AddressAny addr = asAny(items.get(0));
            assertEquals(TransportSize.BYTE, addr.getTransportSize());
            assertEquals(12, addr.getNumberOfElements());
            assertNotNull(ctx.blockMapping());
            assertTrue(ctx.blockMapping().get(0).isMerged());
        }

        @Test
        @DisplayName("auto mode rejects a gap when reading gap bytes is not cheaper")
        void autoMergeRejectsTooLargeGap() {
            config.gap = -1;
            logic.setConfiguration(config);

            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("d0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, "UTF-8"));
            tags.put("d1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 100, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            S7NonHProtocolLogic.ReadRequestContext ctx = logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(2, items.size(), "auto mode should not merge when the BYTE range is more expensive");
            assertEquals(2, ctx.blockMapping().size());
            assertFalse(ctx.blockMapping().get(0).isMerged());
            assertFalse(ctx.blockMapping().get(1).isMerged());
        }

        @Test
        @DisplayName("auto mode accepts a moderate gap when protocol overhead saving wins")
        void autoMergeAcceptsCostEffectiveGap() {
            config.gap = -1;
            logic.setConfiguration(config);

            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("w0", new S7Tag(TransportSize.WORD, MemoryArea.DATA_BLOCKS, 1, 8, (byte) 0, 1, "UTF-8"));
            tags.put("w1", new S7Tag(TransportSize.WORD, MemoryArea.DATA_BLOCKS, 1, 20, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            S7NonHProtocolLogic.ReadRequestContext ctx = logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(1, items.size(), "auto mode should merge when gap bytes cost less than item overhead");

            S7AddressAny addr = asAny(items.get(0));
            assertEquals(TransportSize.BYTE, addr.getTransportSize());
            assertEquals(14, addr.getNumberOfElements());
            assertTrue(ctx.blockMapping().get(0).isMerged());
        }

        @Test
        @DisplayName("auto mode still rejects a cost-effective group when it exceeds PDU")
        void autoMergeStillRespectsPduBudget() {
            config.gap = -1;
            logic.setConfiguration(config);
            when(driverCtx.getPduSize()).thenReturn(28);

            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("w0", new S7Tag(TransportSize.WORD, MemoryArea.DATA_BLOCKS, 1, 8, (byte) 0, 1, "UTF-8"));
            tags.put("w1", new S7Tag(TransportSize.WORD, MemoryArea.DATA_BLOCKS, 1, 20, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(2, items.size(), "auto-detected merge should be skipped when total PDU would overflow");
            assertEquals(TransportSize.WORD, asAny(items.get(0)).getTransportSize());
            assertEquals(TransportSize.WORD, asAny(items.get(1)).getTransportSize());
        }
    }

    @Nested
    @DisplayName("PDU budget: merge vs individual")
    class PduBudget {

        @Test
        @DisplayName("merged block response size ≤ individual total")
        void mergedResponseSmallerThanIndividual() {
            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("t0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            tags.put("t1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 1, "UTF-8"));
            tags.put("t2", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 40, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(1, items.size());

            int mergedResp = 4 + asAny(items.get(0)).getNumberOfElements();
            if (mergedResp % 2 == 1) mergedResp++;

            S7Tag sampleTag = new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 0, (byte) 0, 1, "UTF-8");
            int individualResp = 3 * S7NonHProtocolLogic.computeReadResponseItemSize(sampleTag);

            assertTrue(mergedResp < individualResp,
                    "Merged response " + mergedResp + " should be < individual " + individualResp);
        }

        @Test
        @DisplayName("merge rejected when gap bytes cause PDU overflow")
        void mergeRejectedBeyondPdu() {
            when(driverCtx.getPduSize()).thenReturn(28);

            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("w0", new S7Tag(TransportSize.WORD, MemoryArea.DATA_BLOCKS, 1, 8, (byte) 0, 1, "UTF-8"));
            tags.put("w1", new S7Tag(TransportSize.WORD, MemoryArea.DATA_BLOCKS, 1, 20, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            logic.performOrdinaryReadRequest(request);

            List<S7VarRequestParameterItem> items = itemsCaptor.getValue();
            assertEquals(2, items.size(),
                    "Gapped WORDs should stay individual when merge would overflow PDU");

            for (S7VarRequestParameterItem item : items) {
                assertEquals(TransportSize.WORD, asAny(item).getTransportSize());
            }
        }
    }

    @Nested
    @DisplayName("Mapping travels with ReadRequestContext closure")
    class ClosureSafety {

        @Test
        @DisplayName("mapping survives transport error — no shared-state leak")
        void mappingOnError() {
            RuntimeException error = new RuntimeException("transport down");
            doReturn(CompletableFuture.failedFuture(error))
                    .when(logic)
                    .sendReadMessage(any(DefaultPlcReadRequest.class), anyList(), anyInt());

            Map<String, S7Tag> tags = new LinkedHashMap<>();
            tags.put("d0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            DefaultPlcReadRequest request = readRequest(tags);

            config.gap = 16;
            logic.setConfiguration(config);

            S7NonHProtocolLogic.ReadRequestContext ctx = logic.performOrdinaryReadRequest(request);

            assertTrue(ctx.future().isCompletedExceptionally());
            // The mapping is still held by the closure — no shared state to leak from.
            // If the caller wanted to decode an error response, the mapping is available.
            assertNotNull(ctx.blockMapping(), "blockMapping survives even on transport error");
        }

        @Test
        @DisplayName("two concurrent requests have independent mappings")
        void concurrentRequestsIndependentMappings() {
            Map<String, S7Tag> tagsA = new LinkedHashMap<>();
            tagsA.put("a0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            tagsA.put("a1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 1, "UTF-8"));

            Map<String, S7Tag> tagsB = new LinkedHashMap<>();
            tagsB.put("b0", new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 100, (byte) 0, 1, "UTF-8"));
            tagsB.put("b1", new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 104, (byte) 0, 1, "UTF-8"));

            // Request A: 2 adjacent DINTs merge
            S7NonHProtocolLogic.ReadRequestContext ctxA =
                    logic.performOrdinaryReadRequest(readRequest(tagsA));
            // Request B: 2 adjacent REALs merge
            S7NonHProtocolLogic.ReadRequestContext ctxB =
                    logic.performOrdinaryReadRequest(readRequest(tagsB));

            // Each context holds its own mapping — no interference
            assertNotNull(ctxA.blockMapping());
            assertNotNull(ctxB.blockMapping());
            assertNotSame(ctxA.blockMapping(), ctxB.blockMapping());

            // ctxA has DINT merged block (8 bytes)
            S7NonHProtocolLogic.BlockMergeMapping bmA = ctxA.blockMapping().values().iterator().next();
            assertEquals(8, bmA.blockSize);
            assertEquals("a0", bmA.tagSlices.get(0).tagName());

            // ctxB has REAL merged block (8 bytes)
            S7NonHProtocolLogic.BlockMergeMapping bmB = ctxB.blockMapping().values().iterator().next();
            assertEquals(8, bmB.blockSize);
            assertEquals("b0", bmB.tagSlices.get(0).tagName());
        }

        @Test
        @DisplayName("S7-200 scenario: same tpduId won't collide (closure-based)")
        void sameTpduIdNoCollision() throws Exception {
            // Simulate S7-200: force getTpduId() to always return 0 by
            // setting the controller type in the driver context
            when(driverCtx.getControllerType())
                    .thenReturn(org.apache.plc4x.java.s7.readwrite.ControllerType.S7_200);

            Map<String, S7Tag> tagsA = new LinkedHashMap<>();
            tagsA.put("a0", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 32, (byte) 0, 1, "UTF-8"));
            tagsA.put("a1", new S7Tag(TransportSize.DINT, MemoryArea.DATA_BLOCKS, 1, 36, (byte) 0, 1, "UTF-8"));

            Map<String, S7Tag> tagsB = new LinkedHashMap<>();
            tagsB.put("b0", new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 48, (byte) 0, 1, "UTF-8"));
            tagsB.put("b1", new S7Tag(TransportSize.REAL, MemoryArea.DATA_BLOCKS, 1, 52, (byte) 0, 1, "UTF-8"));

            S7NonHProtocolLogic.ReadRequestContext ctxA =
                    logic.performOrdinaryReadRequest(readRequest(tagsA));
            S7NonHProtocolLogic.ReadRequestContext ctxB =
                    logic.performOrdinaryReadRequest(readRequest(tagsB));

            // Both requests used tpduId=0, but mappings are independent via closure
            assertNotNull(ctxA.blockMapping());
            assertNotNull(ctxB.blockMapping());
            assertNotSame(ctxA.blockMapping(), ctxB.blockMapping(),
                    "S7-200 with tpduId=0: mappings must be independent");
        }
    }
}
