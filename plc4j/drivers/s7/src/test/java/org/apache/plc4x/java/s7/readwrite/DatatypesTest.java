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

import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;

public class DatatypesTest {

    public static void main(String[] args) throws Exception {
        try (PlcConnection connection = new DefaultPlcDriverManager().getConnection("s7://10.15.72.52")) {
            final PlcReadRequest.Builder builder = connection.readRequestBuilder();
            builder.addTagAddress("bool-value-1", "%DB100:0.0:BOOL"); // true
            builder.addTagAddress("bool-value-2", "%DB100:2.1:BOOL"); // false
            // It seems S7 PLCs ignores the array notation for BOOL
            builder.addTagAddress("bool-array", "%DB100:2:BOOL[16]");
            builder.addTagAddress("byte-value", "%DB100:4:BYTE");
            builder.addTagAddress("byte-array", "%DB100:6:BYTE[2]");
            builder.addTagAddress("word-value", "%DB100:8:WORD");
            builder.addTagAddress("word-array", "%DB100:10:WORD[2]");
            builder.addTagAddress("dword-value", "%DB100:14:DWORD");
            builder.addTagAddress("dword-array", "%DB100:18:DWORD[2]");
            //builder.addTagAddress("sint-value", "%DB100:12:SINT"); // 7
            //builder.addTagAddress("sint-array", "%DB100:14:SINT[2]"); // 1, -2
            builder.addTagAddress("int-value", "%DB100:26:INT"); // 23
            builder.addTagAddress("int-array", "%DB100:28:INT[2]"); // 123, -142
            builder.addTagAddress("dint-value", "%DB100:32:DINT"); // 24
            builder.addTagAddress("dint-array", "%DB100:36:DINT[2]"); // 1234, -2345
            //builder.addTagAddress("usint-value", "%DB100:36:USINT"); // 42
            //builder.addTagAddress("usint-array", "%DB100:38:USINT[2]"); // 3, 4
            //builder.addTagAddress("uint-value", "%DB100:40:UINT"); // 3
            //builder.addTagAddress("uint-array", "%DB100:42:UINT[2]"); // 242, 223
            //builder.addTagAddress("udint-value", "%DB100:46:UDINT"); // 815
            //builder.addTagAddress("udint-array", "%DB100:50:UDINT[2]"); // 12345, 23456
            builder.addTagAddress("real-value", "%DB100:44:REAL"); // 3.14159
            builder.addTagAddress("real-array", "%DB100:48:REAL[2]"); // 12.345, 12.345
            //builder.addTagAddress("lreal-value", "%DB100:70:LREAL"); // 3.14159265358979
            //builder.addTagAddress("lreal-array", "%DB100:78:LREAL[2]"); // 1.2345, -1.2345
            builder.addTagAddress("string-value", "%DB100:56:STRING"); // "Hurz"
            // When reading a sized STRING string array, this has to be translated into multiple items
            builder.addTagAddress("string-array", "%DB100:312:STRING[2]"); // "Wolf", "Lamm"
            builder.addTagAddress("time-value", "%DB100:824:TIME"); // 1234ms
            builder.addTagAddress("time-array", "%DB100:828:TIME[2]"); // 123ms, 234ms
            builder.addTagAddress("date-value", "%DB100:836:DATE"); // D#2020-08-20
            builder.addTagAddress("date-array", "%DB100:838:DATE[2]"); // D#1990-03-28, D#2020-10-25
            builder.addTagAddress("time-of-day-value", "%DB100:842:TIME_OF_DAY"); // TOD#12:34:56
            builder.addTagAddress("time-of-day-array", "%DB100:846:TIME_OF_DAY[2]"); // TOD#16:34:56, TOD#08:15:00
            builder.addTagAddress("date-and-time-value", "%DB100:854:DATE_AND_TIME"); // DTL#1978-03-28-12:34:56
            builder.addTagAddress("date-and-time-array", "%DB100:862:DATE_AND_TIME[2]"); // DTL#1978-03-28-12:34:56, DTL#1978-03-28-12:34:56
            builder.addTagAddress("char-value", "%DB100:870:CHAR"); // "H"
            builder.addTagAddress("char-array", "%DB100:872:CHAR[2]"); // "H", "u", "r", "z"
            final PlcReadRequest readRequest = builder.build();

            final PlcReadResponse readResponse = readRequest.execute().get();

            System.out.println(readResponse.getAsPlcValue());

        }
    }

}
