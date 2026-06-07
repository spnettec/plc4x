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
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.utils.cache.CachedPlcConnectionManager;

public class DatatypesTest {

    public static void main(String[] args) throws Exception {
        CachedPlcConnectionManager plcConnectionManager = CachedPlcConnectionManager.getBuilder().build();
        //PlcConnectionManager plcConnectionManager = new DefaultPlcDriverManager();

        Runtime runtime = Runtime.getRuntime();
        long loop = 0;
        long memLogInterval = 100;  // print memory every 100 loops

        boolean lastError = false;
        while (true) {
            loop++;
            try (PlcConnection connection = plcConnectionManager.getConnection("s7://10.80.41.57?block-merge-min-gap=16")) {
                final PlcReadRequest.Builder builder = connection.readRequestBuilder();
                builder.addTagAddress("bool-value-1", "%DB1:0.0:BOOL"); // true
                builder.addTagAddress("bool-value-2", "%DB1:0.1:BOOL"); // false
                builder.addTagAddress("bool-array", "%DB1:2:BIT[10]");
                builder.addTagAddress("byte-value", "%DB1:4:BYTE");
                builder.addTagAddress("byte-array", "%DB1:6:BYTE[2]");
                builder.addTagAddress("word-value", "%DB1:8:WORD");
                builder.addTagAddress("word-array", "%DB1:10:WORD[2]");
                builder.addTagAddress("dword-value", "%DB1:14:DWORD");
                builder.addTagAddress("dword-array", "%DB1:18:DWORD[2]");
                builder.addTagAddress("int-value", "%DB1:26:INT"); // 23
                builder.addTagAddress("int-array", "%DB1:28:INT[2]"); // 123, -142
                builder.addTagAddress("dint-value", "%DB1:32:DINT"); // 24
                builder.addTagAddress("dint-array", "%DB1:36:DINT[2]"); // 1234, -2345
                builder.addTagAddress("real-value", "%DB1:44:REAL"); // 3.14159
                builder.addTagAddress("real-array", "%DB1:48:REAL[2]"); // 12.345, 12.345
                builder.addTagAddress("string-value", "%DB1:56:STRING"); // "Hurz"
                builder.addTagAddress("string-array", "%DB1:312:STRING[2]"); // "Wolf", "Lamm"
                builder.addTagAddress("time-value", "%DB1:824:TIME"); // 1234ms
                builder.addTagAddress("time-array", "%DB1:828:TIME[2]"); // 123ms, 234ms
                builder.addTagAddress("date-value", "%DB1:836:DATE"); // D#2020-08-20
                builder.addTagAddress("date-array", "%DB1:838:DATE[2]"); // D#1990-03-28, D#2020-10-25
                builder.addTagAddress("time-of-day-value", "%DB1:842:TIME_OF_DAY"); // TOD#12:34:56
                builder.addTagAddress("time-of-day-array", "%DB1:846:TIME_OF_DAY[2]"); // TOD#16:34:56, TOD#08:15:00
                builder.addTagAddress("date-and-time-value", "%DB1:854:DTL"); // DTL#1978-03-28-12:34:56
                builder.addTagAddress("date-and-time-array", "%DB1:866:DTL[2]"); // DTL#1978-03-28-12:34:56, DTL#1978-03-28-12:34:56
                builder.addTagAddress("char-value", "%DB1:890:CHAR"); // "H"
                builder.addTagAddress("char-array", "%DB1:892:CHAR[2]"); // "H", "u", "r", "z"

                final PlcReadResponse readResponse = builder.build().execute().get(5, java.util.concurrent.TimeUnit.SECONDS);
                if (lastError) {
                    System.out.printf("[%d] RECOVERED%n", loop);
                    lastError = false;
                }
                System.out.printf("[%d] OK  %s%n", loop, readResponse.getAsPlcValue());

            } catch (Exception e) {
                if (!lastError) {
                    System.err.printf("[%d] ERR %s: %s%n", loop, e.getClass().getSimpleName(), e.getMessage());
                    lastError = true;
                }
            }

            // Memory monitoring every N loops
            if (loop % memLogInterval == 0) {
                runtime.gc();
                long used = runtime.totalMemory() - runtime.freeMemory();
                System.out.printf("[%d] MEM  heap-used=%.1fMB  total=%.1fMB%n",
                    loop, used / 1048576.0, runtime.totalMemory() / 1048576.0);
            }
        }
    }

}
