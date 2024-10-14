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
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.java.spi.values.PlcCHAR;
import org.apache.plc4x.java.spi.values.PlcDATE;
import org.apache.plc4x.java.spi.values.PlcDATE_AND_TIME;
import org.apache.plc4x.java.spi.values.PlcTIME_OF_DAY;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

public class WriteDatatypesTest {

    public static void main(String[] args) throws Exception {
        try (PlcConnection connection = new DefaultPlcDriverManager().getConnection("s7://100.64.0.7")) {
            final PlcWriteRequest.Builder builder = connection.writeRequestBuilder();
            builder.addTagAddress("bool-value-1", "%DB1:0.0:BOOL",true); // true
            builder.addTagAddress("bool-value-2", "%DB1:0.1:BOOL",true); // false
            //builder.addTagAddress("bool-array", "%DB1:2:BOOL[16]",true,false,true,true,false,true,true,false,true,true,false,true,true,false,true,true);
            builder.addTagAddress("byte-value", "%DB1:4:BYTE",'a');
            builder.addTagAddress("byte-array", "%DB1:6:BYTE[2]",'a','b');
            builder.addTagAddress("word-value", "%DB1:8:WORD",42424);
            builder.addTagAddress("word-array", "%DB1:10:WORD[2]",42423,42425);
            builder.addTagAddress("dword-value", "%DB1:14:DWORD",4242442424L);
            builder.addTagAddress("dword-array", "%DB1:18:DWORD[2]",4242442423L,4242442425L);
            builder.addTagAddress("int-value", "%DB1:26:INT",12); // 23
            builder.addTagAddress("int-array", "%DB1:28:INT[2]",123,-142); // 123, -142
            builder.addTagAddress("dint-value", "%DB1:32:DINT",24); // 24
            builder.addTagAddress("dint-array", "%DB1:36:DINT[2]",1234,-2345); // 1234, -2345
            builder.addTagAddress("real-value", "%DB1:44:REAL",3.14159); // 3.14159
            builder.addTagAddress("real-array", "%DB1:48:REAL[2]",12.345,12.345); // 12.345, 12.345
            builder.addTagAddress("string-value", "%DB1:56:STRING","Hurz"); // "Hurz"
            builder.addTagAddress("string-array", "%DB1:312:STRING[2]","Wolf", "Lamm"); // "Wolf", "Lamm"
            builder.addTagAddress("time-value", "%DB1:824:TIME","PT1.234S"); // 1234ms
            builder.addTagAddress("time-array", "%DB1:828:TIME[2]","PT1.234S","PT1.235S"); // 123ms, 234ms
            builder.addTagAddress("date-value", "%DB1:836:DATE",new PlcDATE(LocalDate.parse("1998-03-28"))); // D#2020-08-20
            builder.addTagAddress("date-array", "%DB1:838:DATE[2]",new PlcDATE(LocalDate.parse("1998-03-28")),new PlcDATE(LocalDate.parse("1998-04-28"))); // D#1990-03-28, D#2020-10-25
            builder.addTagAddress("time-of-day-value", "%DB1:842:TIME_OF_DAY",new PlcTIME_OF_DAY(LocalTime.parse("15:36:30.123"))); // TOD#12:34:56
            builder.addTagAddress("time-of-day-array", "%DB1:846:TIME_OF_DAY[2]",new PlcTIME_OF_DAY(LocalTime.parse("15:36:30.123")),new PlcTIME_OF_DAY(LocalTime.parse("15:36:30.124"))); // TOD#16:34:56, TOD#08:15:00
            builder.addTagAddress("date-and-time-value", "%DB1:854:DTL",new PlcDATE_AND_TIME(LocalDateTime.parse("1996-05-06T15:36:30"))); // DTL#1978-03-28-12:34:56
            builder.addTagAddress("date-and-time-array", "%DB1:866:DTL[2]",new PlcDATE_AND_TIME(LocalDateTime.parse("1996-05-06T15:36:30")),new PlcDATE_AND_TIME(LocalDateTime.parse("1996-06-06T15:36:30"))); // DTL#1978-03-28-12:34:56, DTL#1978-03-28-12:34:56
            builder.addTagAddress("char-value", "%DB1:890:CHAR",new PlcCHAR("H")); // "H"
            builder.addTagAddress("char-array", "%DB1:892:CHAR[2]",new PlcCHAR("K"),new PlcCHAR("v")); // "H", "u", "r", "z"
            final PlcWriteRequest writeRequest = builder.build();

            final PlcWriteResponse writeResponse = writeRequest.execute().get();

            writeResponse.getTagNames().forEach(name->{
                System.out.println(name+":"+writeResponse.getResponseCode(name));
            });

        }catch (Exception e) {
            e.printStackTrace();
        }
    }

}
