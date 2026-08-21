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
import org.apache.plc4x.java.api.PlcDriverManager;
import org.apache.plc4x.java.utils.cache.PlcConnectionCache;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.TimeUnit;

/**
 * Performance comparison: no block merge vs fixed gap block merge vs automatic cost-based block merge.
 * <p>
 * Reads the same set of adjacent/mixed tags N times with each configuration
 * and reports the average wall-clock time.
 */
public class BlockPerfTest {

    // Same dataset as DatatypesTest.
    private static final String[][] TAGS = {
        {"bool-value-1", "%DB1:0.0:BOOL"},
        {"bool-value-2", "%DB1:0.1:BOOL"},
        {"bool-array", "%DB1:2:BOOL[10]"},
        {"byte-value", "%DB1:4:BYTE"},
        {"byte-array", "%DB1:6:BYTE[2]"},
        {"word-value", "%DB1:8:WORD"},
        {"word-array", "%DB1:10:WORD[2]"},
        {"dword-value", "%DB1:14:DWORD"},
        {"dword-array", "%DB1:18:DWORD[2]"},
        {"int-value", "%DB1:26:INT"},
        {"int-array", "%DB1:28:INT[2]"},
        {"dint-value", "%DB1:32:DINT"},
        {"dint-array", "%DB1:36:DINT[2]"},
        {"real-value", "%DB1:44:REAL"},
        {"real-array", "%DB1:48:REAL[2]"},
        {"string-value", "%DB1:56:STRING"},
        {"string-array", "%DB1:312:STRING[2]"},
        {"time-value", "%DB1:824:TIME"},
        {"time-array", "%DB1:828:TIME[2]"},
        {"date-value", "%DB1:836:DATE"},
        {"date-array", "%DB1:838:DATE[2]"},
        {"time-of-day-value", "%DB1:842:TIME_OF_DAY"},
        {"time-of-day-array", "%DB1:846:TIME_OF_DAY[2]"},
        {"date-and-time-value", "%DB1:854:DTL"},
        {"date-and-time-array", "%DB1:866:DTL[2]"},
        {"char-value", "%DB1:890:CHAR"},
        {"char-array", "%DB1:892:CHAR[2]"},
    };

    private static final int WARMUP = 20;
    private static final int ITERATIONS = 200;

    public static void main(String[] args) throws Exception {
        String host = args.length > 0 ? args[0] : "10.80.41.57";
        String noMergeUrl = "s7://" + host + "?gap=0";
        String fixedGapUrl = "s7://" + host + "?gap=16";
        String autoUrl = "s7://" + host + "?gap=-1";

        // ── No merge: baseline ───────────────────────────────────────
        PlcConnectionCache mgr0 = PlcConnectionCache.getBuilder().withConnectionFactory(PlcDriverManager.getDefault().getConnectionFactory()).build();
        System.out.println("=== No merge: gap=0 ===");
        double noMerge = benchmark(mgr0, noMergeUrl, "gap=0");
        mgr0.close();

        // ── Fixed gap: gap=16 ────────────────────────────────────────
        PlcConnectionCache mgr1 = PlcConnectionCache.getBuilder().withConnectionFactory(PlcDriverManager.getDefault().getConnectionFactory()).build();
        System.out.println("\n=== Fixed gap: gap=16 ===");
        double fixedGap = benchmark(mgr1, fixedGapUrl, "gap=16");
        mgr1.close();

        // ── Auto: cost-based merge ───────────────────────────────────
        PlcConnectionCache mgr2 = PlcConnectionCache.getBuilder().withConnectionFactory(PlcDriverManager.getDefault().getConnectionFactory()).build();
        System.out.println("\n=== Auto: cost-based block merge ===");
        double auto = benchmark(mgr2, autoUrl, "auto");
        mgr2.close();

        // ── Report ────────────────────────────────────────────────────
        System.out.println("\n==============================================");
        System.out.printf("Tags per request : %d%n", TAGS.length);
        System.out.printf("Iterations       : %d (after %d warmup)%n", ITERATIONS, WARMUP);
        System.out.printf("No merge (gap=0): %.2f ms avg%n", noMerge);
        System.out.printf("Fixed   (gap=16): %.2f ms avg%n", fixedGap);
        System.out.printf("Auto            : %.2f ms avg%n", auto);
        double fixedPct = (noMerge - fixedGap) / noMerge * 100;
        double autoPct = (noMerge - auto) / noMerge * 100;
        double autoVsFixedPct = (fixedGap - auto) / fixedGap * 100;
        System.out.printf("gap=16 vs gap=0 : %.1f%% %s%n", fixedPct, fixedPct > 0 ? "faster" : "slower");
        System.out.printf("Auto vs gap=0   : %.1f%% %s%n", autoPct, autoPct > 0 ? "faster" : "slower");
        System.out.printf("Auto vs gap=16  : %.1f%% %s%n", autoVsFixedPct, autoVsFixedPct > 0 ? "faster" : "slower");
        System.out.println("==============================================");
    }

    private static double benchmark(PlcConnectionCache mgr, String url, String label)
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
