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
package org.apache.plc4x.java.s7.readwrite;

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.utils.cache.CachedPlcConnectionManager;

import java.util.concurrent.TimeUnit;

/**
 * Verify BIT array (%DB1:2:BOOL[10]) write → read consistency,
 * and that read-modify-write preserves unused bits in the last byte.
 */
public class BitArrayTest {

    private static final String TAG = "%DB1:2:BOOL[10]";
    private static final String RAW_TAG = "%DB1:2:BYTE[2]";
    private static final int BIT_COUNT = 10;

    public static void main(String[] args) throws Exception {
        CachedPlcConnectionManager plcConnectionManager = CachedPlcConnectionManager.getBuilder().build();
        String url = args.length > 0 ? args[0] : "s7://10.80.41.57";

        int failures = 0;
        int total = 0;

        // ── Step 1: Save original state ─────────────────────────────
        System.out.println("=== Step 1: Save original ===");
        boolean[] originalBits = new boolean[BIT_COUNT];
        byte[] originalRaw;
        try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
            PlcReadResponse resp = conn.readRequestBuilder()
                    .addTagAddress("bits", TAG)
                    .addTagAddress("raw", RAW_TAG)
                    .build().execute().get(5, TimeUnit.SECONDS);
            PlcValue bitsValue = resp.getPlcValue("bits");
            for (int i = 0; i < BIT_COUNT; i++) originalBits[i] = bitsValue.getIndex(i).getBoolean();
            originalRaw = resp.getPlcValue("raw").getRaw();
            System.out.printf("  original bits: %s  raw: %s%n", bs(originalBits), bh(originalRaw));
        }

        // ── Step 2: RMW verification ───────────────────────────────
        System.out.println("\n=== Step 2: RMW (preserve unused bits) ===");

        try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
            // Write 0xFF to both bytes (sets unused bits 2-7 of byte1)
            conn.writeRequestBuilder()
                    .addTagAddress("raw", RAW_TAG, new byte[]{(byte) 0xFF, (byte) 0xFF})
                    .build().execute().get(5, TimeUnit.SECONDS);
            System.out.println("  wrote raw [0xFF, 0xFF]");

            // Write BIT pattern [1,0,0,0,0,0,0,0,0,0] (only bit 0 set)
            conn.writeRequestBuilder()
                    .addTagAddress("bits", TAG, (Object[]) new Boolean[]{true, false, false, false, false,
                            false, false, false, false, false})
                    .build().execute().get(5, TimeUnit.SECONDS);
            System.out.println("  wrote BIT[10] = [1,0,0,0,0,0,0,0,0,0]");

            // Read back
            PlcReadResponse resp = conn.readRequestBuilder()
                    .addTagAddress("bits", TAG)
                    .addTagAddress("raw", RAW_TAG)
                    .build().execute().get(5, TimeUnit.SECONDS);

            boolean[] readBits = new boolean[BIT_COUNT];
            PlcValue bitsValue = resp.getPlcValue("bits");
            for (int i = 0; i < BIT_COUNT; i++) readBits[i] = bitsValue.getIndex(i).getBoolean();
            byte[] readRaw = resp.getPlcValue("raw").getRaw();

            // Check BIT values
            total++;
            if (readBits[0] && !readBits[1] && !readBits[2]) {
                System.out.printf("  ✅ BIT values correct: %s%n", bs(readBits));
            } else {
                failures++;
                System.out.printf("  ❌ BIT values wrong: %s%n", bs(readBits));
            }

            // Check RMW: byte0 should be 0x01 (bit0=1 from our write, bits1-7=0xFF was overwritten... but actually byte0 has all 10 bits covered!)

            // byte0 has bits 0-7 all used by BIT[10], so it gets fully overwritten:
            // BIT pattern: [1,0,0,0,0,0,0,0, 0,0] → byte0=0x01, byte1 bits0-1=00
            // Expected byte0 = 0x01 (fully overwritten by BIT pattern — OK, no unused bits in byte0)
            // Expected byte1: bit0=0 (from BIT[8]), bit1=0 (from BIT[9]), bits2-7 preserved from 0xFF
            // → byte1 = 0xFC (11111100)
            total++;
            int expectedByte1 = 0xFC;
            int actualByte1 = readRaw[1] & 0xFF;
            if (actualByte1 == expectedByte1) {
                System.out.printf("  ✅ RMW: byte1=0x%02X (bits 2-7 preserved = 0x%02X)%n",
                        actualByte1, actualByte1 & 0xFC);
            } else {
                failures++;
                System.out.printf("  ❌ RMW: byte1=0x%02X expected 0x%02X (unused bits NOT preserved!)%n",
                        actualByte1, expectedByte1);
            }
            System.out.printf("     raw bytes: %s%n", bh(readRaw));
        }

        // ── Step 3: Regular patterns test ──────────────────────────
        System.out.println("\n=== Step 3: Pattern consistency ===");
        boolean[][] patterns = {
            {true,  false, true,  false, true,  false, true,  false, true,  false},
            {false, false, false, false, false, false, false, false, false, false},
            {true,  true,  true,  true,  true,  true,  true,  true,  true,  true},
            {false, false, true,  true,  false, false, true,  true,  false, false},
            {true,  false, false, false, true,  false, false, false, false, true},
        };

        for (int p = 0; p < patterns.length; p++) {
            total++;
            boolean[] pattern = patterns[p];
            try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
                conn.writeRequestBuilder()
                        .addTagAddress("bits", TAG, (Object[]) box(pattern))
                        .build().execute().get(5, TimeUnit.SECONDS);

                PlcReadResponse resp = conn.readRequestBuilder()
                        .addTagAddress("bits", TAG)
                        .build().execute().get(5, TimeUnit.SECONDS);
                PlcValue v = resp.getPlcValue("bits");
                boolean[] rb = new boolean[BIT_COUNT];
                for (int i = 0; i < BIT_COUNT; i++) rb[i] = v.getIndex(i).getBoolean();
                boolean match = true;
                for (int i = 0; i < BIT_COUNT; i++) if (rb[i] != pattern[i]) { match = false; failures++; break; }
                System.out.printf("[P%d] %s → %s %s%n", p, bs(pattern), bs(rb), match ? "✅" : "❌");
            }
        }

        // ── Step 4: Restore ────────────────────────────────────────
        System.out.println("\n=== Step 4: Restore original ===");
        try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
            conn.writeRequestBuilder()
                    .addTagAddress("bits", TAG, (Object[]) box(originalBits))
                    .build().execute().get(5, TimeUnit.SECONDS);
            conn.writeRequestBuilder()
                    .addTagAddress("raw", RAW_TAG, originalRaw)
                    .build().execute().get(5, TimeUnit.SECONDS);
            System.out.println("  restored.");
        }

        System.out.printf("%n=== BIT Array Test: %d/%d passed ===%n", total - failures, total);
        plcConnectionManager.close();
        if (failures > 0) System.exit(1);
    }

    private static Boolean[] box(boolean[] bits) {
        Boolean[] r = new Boolean[bits.length];
        for (int i = 0; i < bits.length; i++) r[i] = bits[i];
        return r;
    }

    private static String bs(boolean[] bits) {
        StringBuilder sb = new StringBuilder("[");
        for (int i = 0; i < bits.length; i++) sb.append(bits[i] ? "1" : "0");
        return sb.append("]").toString();
    }

    private static String bh(byte[] b) {
        if (b == null) return "null";
        StringBuilder sb = new StringBuilder("[");
        for (int i = 0; i < b.length; i++)
            sb.append(String.format("0x%02X", b[i] & 0xFF)).append(i < b.length - 1 ? ", " : "");
        return sb.append("]").toString();
    }
}
