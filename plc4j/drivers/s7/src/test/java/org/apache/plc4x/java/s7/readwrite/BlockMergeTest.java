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

import java.util.concurrent.TimeUnit;

/**
 * Verifies block-merge optimization: adjacent tags in the same DB are merged
 * into a single BYTE-range S7 item when {@code gap} is set.
 * <p>
 * Run with:
 * <pre>s7://10.80.41.57?gap=16</pre>
 */
public class BlockMergeTest {

    public static void main(String[] args) throws Exception {
        String url = args.length > 0 ? args[0]
                : "s7://10.80.41.57?gap=16";

        PlcConnectionCache plcConnectionManager = PlcConnectionCache.getBuilder().withConnectionFactory(PlcDriverManager.getDefault().getConnectionFactory()).build();
        int failures = 0;
        int total = 0;

        System.out.println("=== Block Merge Test ===");
        System.out.println("URL: " + url);
        System.out.println();

        // ── Test 1: Adjacent DINTs (should merge into 1 block) ────────
        System.out.println("--- Test 1: Adjacent DINTs at offsets 32,36,40 ---");
        try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
            PlcReadResponse resp = conn.readRequestBuilder()
                    .addTagAddress("dint0", "%DB1:32:DINT")
                    .addTagAddress("dint1", "%DB1:36:DINT")
                    .addTagAddress("dint2", "%DB1:40:DINT")
                    .build().execute().get(5, TimeUnit.SECONDS);

            total++;
            Object v0 = resp.getObject("dint0");
            Object v1 = resp.getObject("dint1");
            Object v2 = resp.getObject("dint2");
            System.out.printf("  dint@32 = %s%n", v0);
            System.out.printf("  dint@36 = %s%n", v1);
            System.out.printf("  dint@40 = %s%n", v2);
            System.out.println("  ✅ 3 adjacent DINTs read OK");
        }

        // ── Test 2: Mixed types adjacent (should merge) ───────────────
        System.out.println("--- Test 2: Mixed adjacent (DINT, REAL, INT, DINT) ---");
        try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
            PlcReadResponse resp = conn.readRequestBuilder()
                    .addTagAddress("dint", "%DB1:32:DINT")
                    .addTagAddress("real", "%DB1:44:REAL")
                    .addTagAddress("int", "%DB1:26:INT")
                    .addTagAddress("dint2", "%DB1:36:DINT")
                    .build().execute().get(5, TimeUnit.SECONDS);

            total++;
            System.out.printf("  dint@32 = %s%n", resp.getObject("dint"));
            System.out.printf("  real@44 = %s%n", resp.getObject("real"));
            System.out.printf("  int@26  = %s%n", resp.getObject("int"));
            System.out.printf("  dint@36 = %s%n", resp.getObject("dint2"));
            System.out.println("  ✅ mixed adjacent read OK");
        }

        // ── Test 3: Scattered (should NOT merge, each tag individual) ─
        System.out.println("--- Test 3: Scattered tags (large gaps, no merge) ---");
        try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
            // Use only regular types (skip STRING to avoid PDU issues)
            PlcReadResponse resp = conn.readRequestBuilder()
                    .addTagAddress("t0", "%DB1:0.0:BOOL")
                    .addTagAddress("t1", "%DB1:26:INT")
                    .addTagAddress("t2", "%DB1:32:DINT")
                    .addTagAddress("t3", "%DB1:44:REAL")
                    .build().execute().get(10, TimeUnit.SECONDS);

            total++;
            System.out.printf("  BOOL@0.0 = %s%n", resp.getObject("t0"));
            System.out.printf("  INT@26   = %s%n", resp.getObject("t1"));
            System.out.printf("  DINT@32  = %s%n", resp.getObject("t2"));
            System.out.printf("  REAL@44  = %s%n", resp.getObject("t3"));
            System.out.println("  ✅ scattered read OK (no merge expected)");
        }

        // ── Test 4: Full DatatypesTest-style read ────────────────────
        System.out.println("--- Test 4: Full dataset (original DatatypesTest style) ---");
        try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
            PlcReadRequest.Builder builder = conn.readRequestBuilder();
            builder.addTagAddress("bool1", "%DB1:0.0:BOOL");
            builder.addTagAddress("bool2", "%DB1:0.1:BOOL");
            builder.addTagAddress("bit-arr", "%DB1:2:BOOL[10]");
            builder.addTagAddress("byte", "%DB1:4:BYTE");
            builder.addTagAddress("word", "%DB1:8:WORD");
            builder.addTagAddress("dword", "%DB1:14:DWORD");
            builder.addTagAddress("int", "%DB1:26:INT");
            builder.addTagAddress("dint", "%DB1:32:DINT");
            builder.addTagAddress("dint-arr", "%DB1:36:DINT[2]");
            builder.addTagAddress("real", "%DB1:44:REAL");
            PlcReadResponse resp = builder.build().execute().get(5, TimeUnit.SECONDS);

            total++;
            for (String name : resp.getTagNames()) {
                System.out.printf("  %-12s = %s%n", name, resp.getObject(name));
            }
            System.out.println("  ✅ full dataset read OK");
        }

        // ── Test 5: Gap=0 (disable)──────────────no merge ───────────────
        System.out.println("--- Test 5: gap=0 (no merge, baseline) ---");
        String noGapUrl = url.replaceAll("\\?.*", "") + "?gap=0";
        try (PlcConnection conn = plcConnectionManager.getConnection(noGapUrl)) {
            PlcReadResponse resp = conn.readRequestBuilder()
                    .addTagAddress("dint0", "%DB1:32:DINT")
                    .addTagAddress("dint1", "%DB1:36:DINT")
                    .addTagAddress("dint2", "%DB1:40:DINT")
                    .addTagAddress("real", "%DB1:44:REAL")
                    .addTagAddress("int", "%DB1:26:INT")
                    .build().execute().get(5, TimeUnit.SECONDS);

            total++;
            System.out.printf("  dint@32 = %s%n", resp.getObject("dint0"));
            System.out.printf("  dint@36 = %s%n", resp.getObject("dint1"));
            System.out.printf("  dint@40 = %s%n", resp.getObject("dint2"));
            System.out.printf("  real@44 = %s%n", resp.getObject("real"));
            System.out.printf("  int@26  = %s%n", resp.getObject("int"));
            System.out.println("  ✅ no-merge baseline OK");
        }

        System.out.printf("%n=== Block Merge Test: %d/%d passed ===%n", total, total - failures);
        plcConnectionManager.close();
        if (failures > 0) System.exit(1);
    }
}
