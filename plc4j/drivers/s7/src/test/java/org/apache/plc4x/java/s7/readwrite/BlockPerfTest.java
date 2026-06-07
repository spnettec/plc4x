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
import org.apache.plc4x.java.utils.cache.CachedPlcConnectionManager;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.TimeUnit;

/**
 * Performance comparison: gap=0 (original, no merge) vs gap=N (block merge enabled).
 * <p>
 * Reads the same set of adjacent/mixed tags N times with each configuration
 * and reports the average wall-clock time.
 */
public class BlockPerfTest {

    // Adjacent groups — with gap=16 these should merge from ~14 items → ~5 blocks
    private static final String[][] TAGS = {
        {"byte-val",   "%DB1:4:BYTE"},
        {"byte-arr",   "%DB1:6:BYTE[2]"},
        {"word-val",   "%DB1:8:WORD"},
        {"word-arr",   "%DB1:10:WORD[2]"},
        {"dword-val",  "%DB1:14:DWORD"},
        {"dword-arr",  "%DB1:18:DWORD[2]"},
        {"int-val",    "%DB1:26:INT"},
        {"int-arr",    "%DB1:28:INT[2]"},
        {"dint-val",   "%DB1:32:DINT"},
        {"dint-arr",   "%DB1:36:DINT[2]"},
        {"real-val",   "%DB1:44:REAL"},
        {"real-arr",   "%DB1:48:REAL[2]"},
        {"date-val",   "%DB1:836:DATE"},
        {"date-arr",   "%DB1:838:DATE[2]"},
    };

    private static final int WARMUP = 20;
    private static final int ITERATIONS = 200;

    public static void main(String[] args) throws Exception {
        String host = args.length > 0 ? args[0] : "10.80.41.57";
        String baselineUrl = "s7://" + host + "?block-merge-min-gap=0";
        String mergedUrl  = "s7://" + host + "?block-merge-min-gap=16";

        // ── Baseline: gap=0 (original behaviour) ─────────────────────
        CachedPlcConnectionManager mgr0 = CachedPlcConnectionManager.getBuilder().build();
        System.out.println("=== Baseline: gap=0 (no merge) ===");
        double baseline = benchmark(mgr0, baselineUrl, "gap=0");
        mgr0.destroy();

        // ── Optimized: gap=16 ────────────────────────────────────────
        CachedPlcConnectionManager mgr1 = CachedPlcConnectionManager.getBuilder().build();
        System.out.println("\n=== Optimized: gap=16 (block merge) ===");
        double merged = benchmark(mgr1, mergedUrl, "gap=16");
        mgr1.destroy();

        // ── Report ────────────────────────────────────────────────────
        System.out.println("\n==============================================");
        System.out.printf("Tags per request : %d%n", TAGS.length);
        System.out.printf("Iterations       : %d (after %d warmup)%n", ITERATIONS, WARMUP);
        System.out.printf("Baseline (gap=0) : %.2f ms avg%n", baseline);
        System.out.printf("Merged  (gap=16): %.2f ms avg%n", merged);
        double pct = (baseline - merged) / baseline * 100;
        System.out.printf("Improvement      : %.1f%% %s%n", pct, pct > 0 ? "faster" : "slower");
        System.out.println("==============================================");
    }

    private static double benchmark(CachedPlcConnectionManager mgr, String url, String label)
            throws Exception {

        List<Long> times = new ArrayList<>(ITERATIONS);

        for (int i = -WARMUP; i < ITERATIONS; i++) {
            long t0 = System.nanoTime();
            try (PlcConnection conn = mgr.getConnection(url)) {
                PlcReadRequest.Builder b = conn.readRequestBuilder();
                for (String[] tag : TAGS) {
                    b.addTagAddress(tag[0], tag[1]);
                }
                PlcReadResponse resp = b.build().execute().get(10, TimeUnit.SECONDS);
                for (String[] tag : TAGS) {
                    resp.getObject(tag[0]); // ensure all values parsed
                }
            }
            long t1 = System.nanoTime();
            if (i >= 0) {
                times.add(t1 - t0);
            }
            if (i < 0 && (i + WARMUP) % 5 == 0) {
                System.out.printf("  warmup %d/%d ...%n", -i, WARMUP);
            }
        }

        double avgNs = times.stream().mapToLong(Long::longValue).average().orElse(0);
        double avgMs = avgNs / 1_000_000.0;
        long minNs = times.stream().mapToLong(Long::longValue).min().orElse(0);
        long maxNs = times.stream().mapToLong(Long::longValue).max().orElse(0);
        System.out.printf("  avg: %.2f ms  min: %.2f ms  max: %.2f ms%n",
                avgMs, minNs / 1_000_000.0, maxNs / 1_000_000.0);
        return avgMs;
    }
}
