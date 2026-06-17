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
package org.apache.plc4x.java.s7.optimizer;

import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.s7.readwrite.MemoryArea;
import org.apache.plc4x.java.s7.readwrite.TransportSize;
import org.apache.plc4x.java.s7.context.S7DriverContext;
import org.apache.plc4x.java.s7.tag.S7StringTag;
import org.apache.plc4x.java.s7.tag.S7StringTag;
import org.apache.plc4x.java.s7.tag.S7Tag;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/**
 * Read-optimizer that merges adjacent same-area tags into single block-reads. If two
 * tags in the same memory area sit close enough that an extra S7 address-item
 * (12 bytes) costs more than the gap, they are fetched as one byte-block and the
 * decoder splits the response back into per-tag values.
 *
 * <p>The merge threshold is controlled by the {@code gap} configuration parameter:
 * <ul>
 *   <li>{@code gap < 0} (default -1) — auto cost-based: merge when the merged
 *       byte-range item is cheaper than reading the tags individually.</li>
 *   <li>{@code gap == 0} — disabled; each tag is its own item.</li>
 *   <li>{@code gap > 0} — fixed-gap: merge when the byte distance between
 *       adjacent tags is less than {@code gap}.</li>
 * </ul>
 *
 * <p>Tags that cannot benefit from block-merging (Bool/BIT, STRING/WSTRING,
 * var-length strings) fall through to the base {@link S7Optimizer}.
 */
public class S7BlockReadOptimizer extends S7Optimizer {

    @Override
    public List<S7ReadChunk> splitReadRequest(PlcReadRequest request, S7DriverContext context) {
        int gap = context.getGap();

        // gap == 0 disables block merging — delegate to base optimizer directly.
        if (gap == 0) {
            return super.splitReadRequest(request, context);
        }

        // 1. Group tags by memory area (and DB number where applicable).
        Map<String, List<TagEntry>> tagsPerArea = new LinkedHashMap<>();
        LinkedHashMap<String, PlcTag> passthrough = new LinkedHashMap<>();
        for (String tagName : request.getTagNames()) {
            PlcTag plcTag = request.getTag(tagName);
            if (!(plcTag instanceof S7Tag s7Tag) || plcTag instanceof S7StringTag) {
                // Block-merging is unsafe for var-length strings (response size is dynamic);
                // delegate them to the base optimizer.
                passthrough.put(tagName, plcTag);
                continue;
            }
            // Bool and STRING/WSTRING are excluded from block merging.
            if (s7Tag.getDataType() == TransportSize.BOOL
                    || s7Tag.getDataType() == TransportSize.STRING
                    || s7Tag.getDataType() == TransportSize.WSTRING) {
                passthrough.put(tagName, plcTag);
                continue;
            }
            String areaKey = areaKey(s7Tag);
            tagsPerArea.computeIfAbsent(areaKey, k -> new ArrayList<>()).add(new TagEntry(tagName, s7Tag));
        }

        // 2. Within each area, sort by byte offset and build merge groups.
        LinkedHashMap<String, PlcTag> merged = new LinkedHashMap<>();
        Map<String, List<S7ReadChunk.Binding>> blockBindings = new LinkedHashMap<>();
        int blockCounter = 0;

        for (Map.Entry<String, List<TagEntry>> e : tagsPerArea.entrySet()) {
            List<TagEntry> entries = e.getValue();
            entries.sort(Comparator.comparingInt(t -> t.tag.getByteOffset()));

            boolean autoMerge = gap < 0;
            int idx = 0;
            while (idx < entries.size()) {
                TagEntry first = entries.get(idx);
                int blockStart = first.tag.getByteOffset();
                int blockEnd = blockStart + tagSizeInBytes(first.tag);
                List<TagEntry> group = new ArrayList<>();
                group.add(first);
                // Accumulated individual cost for auto mode.
                int groupIndividualCost = computeReadRequestCost(first.tag);
                idx++;
                while (idx < entries.size()) {
                    TagEntry next = entries.get(idx);
                    int nextStart = next.tag.getByteOffset();
                    int nextEnd = nextStart + tagSizeInBytes(next.tag);
                    if (autoMerge) {
                        // Auto cost-based: close the group when merging is no longer cheaper.
                        int candidateStart = Math.min(blockStart, nextStart);
                        int candidateEnd = Math.max(blockEnd, nextEnd);
                        int candidateBlockSize = candidateEnd - candidateStart;
                        int candidateIndividualCost = groupIndividualCost + computeReadRequestCost(next.tag);
                        int candidateMergedCost = S7_ADDRESS_ANY_SIZE + candidateBlockSize;
                        if (candidateMergedCost < candidateIndividualCost) {
                            blockStart = candidateStart;
                            blockEnd = candidateEnd;
                            groupIndividualCost = candidateIndividualCost;
                        } else {
                            break;  // close current group
                        }
                    } else {
                        // Fixed-gap mode: merge when gap < configured threshold.
                        int gapBytes = nextStart - blockEnd;
                        if (gapBytes >= gap) {
                            break;
                        }
                        blockEnd = Math.max(blockEnd, nextEnd);
                    }
                    group.add(next);
                    idx++;
                }
                if (group.size() == 1) {
                    merged.put(first.tagName, first.tag);
                } else {
                    String blockName = "__block__" + (blockCounter++);
                    int blockBytes = blockEnd - blockStart;
                    S7Tag blockTag = new S7Tag(TransportSize.BYTE, first.tag.getMemoryArea(),
                        first.tag.getBlockNumber(), blockStart, (byte) 0, blockBytes);
                    merged.put(blockName, blockTag);
                    List<S7ReadChunk.Binding> bindings = new ArrayList<>(group.size());
                    for (TagEntry te : group) {
                        bindings.add(new S7ReadChunk.Binding(te.tagName, te.tag,
                            te.tag.getByteOffset() - blockStart, 0, false));
                    }
                    blockBindings.put(blockName, bindings);
                }
            }
        }
        // Append the passthrough (un-mergeable) tags last in original order.
        merged.putAll(passthrough);

        // 3. Hand off to the base optimizer with the rewritten map.
        List<S7ReadChunk> baseChunks = splitReadFromMap(merged, context);

        // 4. Replace synthetic-block bindings with their per-tag bindings.
        if (blockBindings.isEmpty()) {
            return baseChunks;
        }
        List<S7ReadChunk> out = new ArrayList<>(baseChunks.size());
        for (S7ReadChunk chunk : baseChunks) {
            List<S7ReadChunk.Slot> rewritten = new ArrayList<>(chunk.slots().size());
            for (S7ReadChunk.Slot slot : chunk.slots()) {
                S7ReadChunk.Binding b0 = slot.bindings().get(0);
                List<S7ReadChunk.Binding> blockBs = blockBindings.get(b0.tagName());
                if (blockBs != null) {
                    rewritten.add(new S7ReadChunk.Slot(slot.requestItem(), slot.fragmentTag(), blockBs));
                } else {
                    rewritten.add(slot);
                }
            }
            out.add(new S7ReadChunk(rewritten));
        }
        return out;
    }

    /**
     * Compute the approximate read-request cost (request address + response payload) for one
     * tag, used by the auto-gap mode to decide whether a block-merge is cheaper than
     * individual reads. Does not need to match the wire exactly; a conservative estimate
     * is sufficient because the base optimizer does the final PDU layout.
     */
    private static int computeReadRequestCost(S7Tag tag) {
        int response = 4 + tagSizeInBytes(tag);
        if (response % 2 == 1) response++;
        return S7_ADDRESS_ANY_SIZE + response;
    }

    private static String areaKey(S7Tag s7Tag) {
        MemoryArea area = s7Tag.getMemoryArea();
        if (area == MemoryArea.DATA_BLOCKS || area == MemoryArea.INSTANCE_DATA_BLOCKS) {
            return area.getShortName() + "/" + s7Tag.getBlockNumber();
        }
        return area.getShortName();
    }

    private static int tagSizeInBytes(S7Tag tag) {
        if (tag.getDataType() == TransportSize.BOOL) {
            return Math.max(1, (tag.getNumberOfElements() + 7) / 8);
        }
        if (tag instanceof S7StringTag fixed) {
            int bytesPerChar = fixed.getDataType() == TransportSize.WSTRING ? 2 : 1;
            return tag.getNumberOfElements() * (fixed.getStringLength() + 2) * bytesPerChar;
        }
        return tag.getNumberOfElements() * tag.getDataType().getSizeInBytes();
    }

    private record TagEntry(String tagName, S7Tag tag) {}
}
