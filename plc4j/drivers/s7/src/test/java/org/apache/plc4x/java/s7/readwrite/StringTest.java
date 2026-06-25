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
 * Verify STRING and STRING[2] read/write consistency in NonH driver.
 */
public class StringTest {

    public static void main(String[] args) throws Exception {
        CachedPlcConnectionManager plcConnectionManager = CachedPlcConnectionManager.getBuilder().build();
        String url = args.length > 0 ? args[0] : "s7://10.80.41.57";

        // ── Step 1: Read original values ──────────────────────────────
        System.out.println("=== Step 1: Read original strings ===");
        String origSingle, origArr0, origArr1;
        try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
            PlcReadResponse resp = conn.readRequestBuilder()
                    .addTagAddress("s1", "%DB1:56:STRING")
                    .addTagAddress("s2", "%DB1:312:STRING[2]")
                    .build()
                    .execute()
                    .get(5, TimeUnit.SECONDS);

            PlcValue v1 = resp.getPlcValue("s1");
            origSingle = v1 != null ? v1.getString() : null;
            System.out.printf("  STRING:       \"%s\" (type: %s)%n", origSingle,
                    v1 != null ? v1.getClass().getSimpleName() : "null");

            PlcValue v2 = resp.getPlcValue("s2");
            if (v2 != null) {
                int nv = resp.getNumberOfValues("s2");
                System.out.printf("  STRING[2]:    #values=%d (type: %s)%n", nv,
                        v2.getClass().getSimpleName());
                origArr0 = nv > 0 ? v2.getIndex(0).getString() : null;
                origArr1 = nv > 1 ? v2.getIndex(1).getString() : null;
                System.out.printf("    [0]: \"%s\"%n", origArr0);
                System.out.printf("    [1]: \"%s\"%n", origArr1);
            } else {
                origArr0 = origArr1 = null;
                System.out.println("  STRING[2]:    null");
            }
        }

        // ── Step 2: Write + Read back ─────────────────────────────────
        System.out.println("\n=== Step 2: Write → Read consistency ===");

        String[] testStrings = {"Hello世界", "Wolf", "Lamm", "Abc123"};
        int failures = 0;
        int total = 0;

        // ── Single STRING ──
        for (String test : testStrings) {
            total++;
            try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
                // Write
                PlcWriteResponse wresp = conn.writeRequestBuilder()
                        .addTagAddress("s", "%DB1:56:STRING", test)
                        .build()
                        .execute()
                        .get(5, TimeUnit.SECONDS);
                System.out.printf("[S%d] WRITE \"%s\" → %s%n", total, test,
                        wresp.getResponseCode("s"));

                // Read back
                PlcReadResponse rresp = conn.readRequestBuilder()
                        .addTagAddress("s", "%DB1:56:STRING")
                        .build()
                        .execute()
                        .get(5, TimeUnit.SECONDS);
                PlcValue v = rresp.getPlcValue("s");
                String readBack = v.getString();

                if (test.equals(readBack)) {
                    System.out.printf("[S%d] READ  \"%s\" ✅ MATCH%n", total, readBack);
                } else {
                    failures++;
                    System.out.printf("[S%d] READ  \"%s\" ❌ MISMATCH (expected \"%s\")%n",
                            total, readBack, test);
                }
            }
        }

        // ── STRING[2] ──
        String[][] testArrPairs = {
            {"你好世界", "HelloWorld"},
            {"Wolf", "Lamm"},
            {"abc", "xyz"},
        };
        for (String[] pair : testArrPairs) {
            total++;
            try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
                PlcWriteResponse wresp = conn.writeRequestBuilder()
                        .addTagAddress("s", "%DB1:312:STRING[2]", (Object[]) pair)
                        .build()
                        .execute()
                        .get(5, TimeUnit.SECONDS);
                System.out.printf("[A%d] WRITE [\"%s\", \"%s\"] → %s%n",
                        total, pair[0], pair[1], wresp.getResponseCode("s"));

                PlcReadResponse rresp = conn.readRequestBuilder()
                        .addTagAddress("s", "%DB1:312:STRING[2]")
                        .build()
                        .execute()
                        .get(5, TimeUnit.SECONDS);
                PlcValue v = rresp.getPlcValue("s");
                int nv = rresp.getNumberOfValues("s");
                String r0 = nv > 0 ? v.getIndex(0).getString() : null;
                String r1 = nv > 1 ? v.getIndex(1).getString() : null;

                boolean m0 = pair[0].equals(r0);
                boolean m1 = pair[1].equals(r1);
                if (m0 && m1) {
                    System.out.printf("[A%d] READ  [\"%s\", \"%s\"] ✅ MATCH%n",
                            total, r0, r1);
                } else {
                    failures++;
                    if (!m0) System.err.printf("  ❌ STRING[0]: wrote=\"%s\", read=\"%s\"%n",
                            pair[0], r0);
                    if (!m1) System.err.printf("  ❌ STRING[1]: wrote=\"%s\", read=\"%s\"%n",
                            pair[1], r1);
                }
            }
        }

        // ── Step 3: Restore ───────────────────────────────────────────
        System.out.println("\n=== Step 3: Restore original ===");
        try (PlcConnection conn = plcConnectionManager.getConnection(url)) {
            if (origSingle != null) {
                conn.writeRequestBuilder()
                        .addTagAddress("s", "%DB1:56:STRING", origSingle)
                        .build()
                        .execute()
                        .get(5, TimeUnit.SECONDS);
                System.out.printf("  Restored STRING: \"%s\"%n", origSingle);
            }
            if (origArr0 != null) {
                conn.writeRequestBuilder()
                        .addTagAddress("s", "%DB1:312:STRING[2]",
                                (Object[]) new String[]{origArr0, origArr1})
                        .build()
                        .execute()
                        .get(5, TimeUnit.SECONDS);
                System.out.printf("  Restored STRING[2]: [\"%s\", \"%s\"]%n", origArr0, origArr1);
            }
        }

        System.out.printf("%n=== String Test: %d/%d passed ===%n",
                total - failures, total);
        plcConnectionManager.close();
        if (failures > 0) System.exit(1);
    }
}
