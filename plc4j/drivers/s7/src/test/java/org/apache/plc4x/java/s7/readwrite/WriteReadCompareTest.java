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
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.values.PlcCHAR;
import org.apache.plc4x.java.spi.values.PlcDATE;
import org.apache.plc4x.java.spi.values.PlcDATE_AND_TIME;
import org.apache.plc4x.java.spi.values.PlcTIME_OF_DAY;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

/**
 * Write/read round-trip comparison for every datatype of the factory test DB.
 *
 * <p>For each entry: write a value that differs from what the DB holds, read it back and
 * compare the rendered form. Afterwards the canonical values (the ones WriteDatatypesTest
 * seeds and the other harnesses expect) are restored and verified once.</p>
 */
public class WriteReadCompareTest {

    private static final String URL = "s7://10.80.41.57";

    private record Entry(String name, String address, Object[] write, String expected) {}

    private static final Entry[] ENTRIES = {
        new Entry("bool-value-1", "%DB1:0.0:BOOL", new Object[]{false}, "false"),
        new Entry("bool-value-2", "%DB1:0.1:BOOL", new Object[]{true}, "true"),
        new Entry("bool-array", "%DB1:2:BOOL[16]",
            new Object[]{true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false},
            "[true,false,true,false,true,false,true,false,true,false,true,false,true,false,true,false]"),
        new Entry("byte-value", "%DB1:4:BYTE", new Object[]{31}, "31"),
        // hex digits 1/2 only - the byte-array read renders as hex, whose letter case is not pinned
        new Entry("byte-array", "%DB1:6:BYTE[2]", new Object[]{0x11, 0x22}, "1122"),
        new Entry("word-value", "%DB1:8:WORD", new Object[]{4660}, "4660"),
        new Entry("word-array", "%DB1:10:WORD[2]", new Object[]{4661, 4662}, "[4661,4662]"),
        new Entry("dword-value", "%DB1:14:DWORD", new Object[]{3735928559L}, "3735928559"),
        new Entry("dword-array", "%DB1:18:DWORD[2]", new Object[]{3735928558L, 3735928557L}, "[3735928558,3735928557]"),
        new Entry("int-value", "%DB1:26:INT", new Object[]{777}, "777"),
        new Entry("int-array", "%DB1:28:INT[2]", new Object[]{778, -779}, "[778,-779]"),
        new Entry("dint-value", "%DB1:32:DINT", new Object[]{77777}, "77777"),
        new Entry("dint-array", "%DB1:36:DINT[2]", new Object[]{77778, -77779}, "[77778,-77779]"),
        // float-exact values, so the rendered form cannot drift
        new Entry("real-value", "%DB1:44:REAL", new Object[]{100.25}, "100.25"),
        new Entry("real-array", "%DB1:48:REAL[2]", new Object[]{100.25, -0.5}, "[100.25,-0.5]"),
        // exercises the CJK path of YOFC's string encoding handling
        new Entry("string-value", "%DB1:56:STRING", new Object[]{"回环测试"}, "回环测试"),
        new Entry("string-array", "%DB1:312:STRING[2]", new Object[]{"甲", "乙"}, "[甲,乙]"),
        new Entry("time-value", "%DB1:824:TIME", new Object[]{"PT0.25S"}, "PT0.25S"),
        new Entry("time-array", "%DB1:828:TIME[2]", new Object[]{"PT0.25S", "PT1.5S"}, "[PT0.25S,PT1.5S]"),
        new Entry("date-value", "%DB1:836:DATE",
            new Object[]{new PlcDATE(LocalDate.parse("2026-09-28"))}, "2026-09-28"),
        new Entry("date-array", "%DB1:838:DATE[2]",
            new Object[]{new PlcDATE(LocalDate.parse("2026-09-28")), new PlcDATE(LocalDate.parse("2026-09-29"))},
            "[2026-09-28,2026-09-29]"),
        new Entry("time-of-day-value", "%DB1:842:TIME_OF_DAY",
            new Object[]{new PlcTIME_OF_DAY(LocalTime.parse("23:59:59"))}, "23:59:59"),
        new Entry("time-of-day-array", "%DB1:846:TIME_OF_DAY[2]",
            new Object[]{new PlcTIME_OF_DAY(LocalTime.parse("23:59:59")), new PlcTIME_OF_DAY(LocalTime.parse("00:00:01"))},
            "[23:59:59,00:00:01]"),
        new Entry("date-and-time-value", "%DB1:854:DTL",
            new Object[]{new PlcDATE_AND_TIME(LocalDateTime.parse("2026-09-28T23:59:59"))}, "2026-09-28T23:59:59"),
        new Entry("date-and-time-array", "%DB1:866:DTL[2]",
            new Object[]{new PlcDATE_AND_TIME(LocalDateTime.parse("2026-09-28T23:59:59")),
                new PlcDATE_AND_TIME(LocalDateTime.parse("2026-09-29T00:00:01"))},
            "[2026-09-28T23:59:59,2026-09-29T00:00:01]"),
        new Entry("char-value", "%DB1:890:CHAR", new Object[]{new PlcCHAR("Z")}, "Z"),
        new Entry("char-array", "%DB1:892:CHAR[2]", new Object[]{new PlcCHAR("Y"), new PlcCHAR("x")}, "[Y,x]"),
    };

    /** The canonical values the other harnesses expect - same set WriteDatatypesTest seeds. */
    private static final Entry[] RESTORE = {
        new Entry("bool-value-1", "%DB1:0.0:BOOL", new Object[]{true}, "true"),
        new Entry("bool-value-2", "%DB1:0.1:BOOL", new Object[]{false}, "false"),
        new Entry("bool-array", "%DB1:2:BOOL[16]",
            new Object[]{true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true},
            "[true,false,true,true,false,true,true,false,true,true,false,true,true,false,true,true]"),
        new Entry("byte-value", "%DB1:4:BYTE", new Object[]{'a'}, "97"),
        new Entry("byte-array", "%DB1:6:BYTE[2]", new Object[]{'a', 'b'}, "6162"),
        new Entry("int-value", "%DB1:26:INT", new Object[]{23}, "23"),
        new Entry("int-array", "%DB1:28:INT[2]", new Object[]{123, -142}, "[123,-142]"),
        new Entry("dint-value", "%DB1:32:DINT", new Object[]{24}, "24"),
        new Entry("dint-array", "%DB1:36:DINT[2]", new Object[]{1234, -2345}, "[1234,-2345]"),
        new Entry("real-value", "%DB1:44:REAL", new Object[]{3.14159}, "3.14159"),
        new Entry("real-array", "%DB1:48:REAL[2]", new Object[]{12.345, 12.345}, "[12.345,12.345]"),
        new Entry("string-value", "%DB1:56:STRING", new Object[]{"Hurz"}, "Hurz"),
        new Entry("string-array", "%DB1:312:STRING[2]", new Object[]{"Wolf", "Lamm"}, "[Wolf,Lamm]"),
        new Entry("time-value", "%DB1:824:TIME", new Object[]{"PT1.234S"}, "PT1.234S"),
        new Entry("time-array", "%DB1:828:TIME[2]", new Object[]{"PT0.123S", "PT0.234S"}, "[PT0.123S,PT0.234S]"),
        new Entry("date-value", "%DB1:836:DATE",
            new Object[]{new PlcDATE(LocalDate.parse("2020-08-20"))}, "2020-08-20"),
        new Entry("date-array", "%DB1:838:DATE[2]",
            new Object[]{new PlcDATE(LocalDate.parse("1990-03-28")), new PlcDATE(LocalDate.parse("2020-10-25"))},
            "[1990-03-28,2020-10-25]"),
        new Entry("time-of-day-value", "%DB1:842:TIME_OF_DAY",
            new Object[]{new PlcTIME_OF_DAY(LocalTime.parse("12:34:56"))}, "12:34:56"),
        new Entry("time-of-day-array", "%DB1:846:TIME_OF_DAY[2]",
            new Object[]{new PlcTIME_OF_DAY(LocalTime.parse("16:34:56")), new PlcTIME_OF_DAY(LocalTime.parse("08:15"))},
            "[16:34:56,08:15]"),
        new Entry("date-and-time-value", "%DB1:854:DTL",
            new Object[]{new PlcDATE_AND_TIME(LocalDateTime.parse("1978-03-28T12:34:56"))}, "1978-03-28T12:34:56"),
        new Entry("date-and-time-array", "%DB1:866:DTL[2]",
            new Object[]{new PlcDATE_AND_TIME(LocalDateTime.parse("1978-03-28T12:34:56")),
                new PlcDATE_AND_TIME(LocalDateTime.parse("1978-03-28T12:34:56"))},
            "[1978-03-28T12:34:56,1978-03-28T12:34:56]"),
        new Entry("char-value", "%DB1:890:CHAR", new Object[]{new PlcCHAR("H")}, "H"),
        new Entry("char-array", "%DB1:892:CHAR[2]", new Object[]{new PlcCHAR("H"), new PlcCHAR("u")}, "[H,u]"),
    };

    public static void main(String[] args) throws Exception {
        try (PlcConnection connection = new org.apache.plc4x.java.DefaultPlcDriverManager().getConnection(URL)) {
            int pass = 0, fail = 0;
            System.out.println("================ 写入 -> 读出 对比 ================");
            for (Entry e : ENTRIES) {
                PlcWriteRequest.Builder wb = connection.writeRequestBuilder();
                wb.addTagAddress(e.name, e.address, e.write);
                String code = wb.build().execute().get().getResponseCode(e.name).toString();
                PlcReadResponse rd = connection.readRequestBuilder().addTagAddress(e.name, e.address).build().execute().get();
                String actual = render(rd, e.name);
                boolean ok = code.equals("OK") && expected(e.expected, actual);
                if (ok) pass++; else fail++;
                System.out.printf("%-22s %-8s write=%-6s read=%s%s%n", e.name, ok ? "✅" : "❌", code, actual,
                    ok ? "" : "  (expect " + e.expected + ")");
            }
            System.out.println("---------------- 恢复标准值并验证 ----------------");
            for (Entry e : RESTORE) {
                PlcWriteRequest.Builder wb = connection.writeRequestBuilder();
                wb.addTagAddress(e.name, e.address, e.write);
                String code = wb.build().execute().get().getResponseCode(e.name).toString();
                PlcReadResponse rd = connection.readRequestBuilder().addTagAddress(e.name, e.address).build().execute().get();
                String actual = render(rd, e.name);
                boolean ok = code.equals("OK") && expected(e.expected, actual);
                if (!ok) fail++;
                System.out.printf("%-22s %-8s write=%-6s read=%s%s%n", e.name, ok ? "✅" : "❌", code, actual,
                    ok ? "" : "  (expect " + e.expected + ")");
            }
            int total = ENTRIES.length + RESTORE.length;
            System.out.printf("%n============ 回环 %d 项，失败 %d 项 ============%n", total, fail);
        }
    }

    private static String render(PlcReadResponse rd, String name) {
        PlcValue v = rd.getPlcValue(name);
        return v == null ? "<null>" : v.toString().replace(" ", "");
    }

    private static boolean expected(String expected, String actual) {
        return expected.replace(" ", "").equals(actual);
    }
}
