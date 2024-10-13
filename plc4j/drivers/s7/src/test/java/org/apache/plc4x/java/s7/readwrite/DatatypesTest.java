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
import org.apache.plc4x.java.api.PlcConnectionManager;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.utils.cache.CachedPlcConnectionManager;

public class DatatypesTest {

    public static void main(String[] args) throws Exception {
        CachedPlcConnectionManager plcConnectionManager = CachedPlcConnectionManager.getBuilder().build();
        //PlcConnectionManager plcConnectionManager = new DefaultPlcDriverManager();
        while(true) {
            try (PlcConnection connection = plcConnectionManager.getConnection("s7://100.64.0.7")) {
                final PlcReadRequest.Builder builder = connection.readRequestBuilder();
                builder.addTagAddress("bool-value-1", "%DB1:0.0:BOOL"); // true
                builder.addTagAddress("bool-value-2", "%DB1:0.1:BOOL"); // false
                // It seems S7 PLCs ignores the array notation for BOOL
                builder.addTagAddress("bool-array", "%DB1:2:BOOL[16]");
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
                final PlcReadRequest readRequest = builder.build();

                final PlcReadResponse readResponse = readRequest.execute().get();

                System.out.println(readResponse.getAsPlcValue());

            }


        }

        /*
        while(true) {
        try (PlcConnection connection = plcConnectionManager.getConnection("s7://10.110.20.104?controller-type=S7_200&remote-rack=0&remote-slot=2")) {
            final PlcReadRequest.Builder builder = connection.readRequestBuilder();
            builder.addTagAddress("string", "%DB1.DBD28:REAL"); // true
            builder.addTagAddress("Caterpillar-ControllerFailure",	"%M81.3:BOOL");
            builder.addTagAddress("Caterpillar-Current",	"%DB1.DBD28:REAL");
            builder.addTagAddress("Caterpillar-DownStatus",	"%DB12.DBX32.1:BOOL");
            builder.addTagAddress("Caterpillar-EmergencyStop",	"%M80.7:BOOL");
            builder.addTagAddress("Caterpillar-HighSpeed-Selection",	"%M35.4:BOOL");
            builder.addTagAddress("Caterpillar-LowSpeed-Selection",	"%M35.5:BOOL");
            builder.addTagAddress("Caterpillar-Speed",	"%DB1.DBD20:REAL");
            builder.addTagAddress("Caterpillar-Status",	"%M12.7:BOOL");
            builder.addTagAddress("Charge-Mode","%M13.1:BOOL");
            builder.addTagAddress("Clearance-Mode","%M13.0:BOOL");
            builder.addTagAddress("Clearance-Switch","%M19.3:BOOL");
            builder.addTagAddress("CoolingFan-Failure","%M81.7:BOOL");
            builder.addTagAddress("Diameter-Setting","%DB1.DBD174:REAL");
            builder.addTagAddress("DiameterAverage-Display","%DB66.DBD40:REAL");
            builder.addTagAddress("DiameterDeviation-AlarmCount","%DB1.DBW244:INT");
            builder.addTagAddress("DiameterDeviation-Value","%DB1.DBD306:REAL");
            builder.addTagAddress("Extruder-Current","%DB1.DBD44:REAL");
            builder.addTagAddress("Extruder-Failure","%M81.1:BOOL");
            builder.addTagAddress("Extruder-Speed","%DB1.DBD36:REAL");
            builder.addTagAddress("Extruder-Status","%M12.0:BOOL");
            builder.addTagAddress("FuelSupply-Switch","%M19.4:BOOL");
            builder.addTagAddress("Heat-Mode","%M100.0:BOOL");
            builder.addTagAddress("Length-Product","%DB1.DBD224:REAL");
            builder.addTagAddress("Lump-AlarmCount","%DB1.DBW246:INT");
            builder.addTagAddress("MainCabinet-EmergencyStop","%M81.4:BOOL");
            builder.addTagAddress("Neck-AlarmCount","%DB1.DBW248:INT");
            builder.addTagAddress("PO-Alarm","%M80.6:BOOL");
            builder.addTagAddress("PO-DanceAlarm","%M82.1:BOOL");
            builder.addTagAddress("Separate-or-Sychrom","%M35.6:BOOL");
            builder.addTagAddress("Spark-Alarm","%M80.2:BOOL");
            builder.addTagAddress("Spark-AlarmCount","%DB1.DBW250:INT");
            builder.addTagAddress("TU-Alarm","%M80.5:BOOL");
            builder.addTagAddress("Temperature-Alarm","%M65.0:BOOL");
            builder.addTagAddress("Vacuum-Failure","%M81.6:BOOL");
            builder.addTagAddress("Vacuum-Switch","%M30.4:BOOL");
            builder.addTagAddress("WaterPump-Failure","%M80.1:BOOL");
            builder.addTagAddress("WaterPump-Switch","%M30.1:BOOL");
            builder.addTagAddress("WaterTank-Temperature","%DB32.DBW74:INT");
            builder.addTagAddress("Xaxis-Display","%DB66.DBD44:REAL");
            builder.addTagAddress("Yaxis-Display","%DB66.DBD48:REAL");
            builder.addTagAddress("Zone1-Temperature","%DB32.DBW56:INT");
            builder.addTagAddress("Zone2-Temperature","%DB32.DBW58:INT");
            builder.addTagAddress("Zone3-Temperature","%DB32.DBW60:INT");
            builder.addTagAddress("Zone4-Temperature","%DB32.DBW62:INT");
            builder.addTagAddress("Zone5-Temperature","%DB32.DBW64:INT");
            builder.addTagAddress("Zone6-Temperature","%DB32.DBW66:INT");
            builder.addTagAddress("Zone7-Temperature","%DB32.DBW68:INT");
            builder.addTagAddress("Zone8-Temperature","%DB32.DBW70:INT");
            builder.addTagAddress("Zone9-Temperature","%DB32.DBW72:INT");
            builder.addTagAddress("sb_23_pv","%DB1.DBD20:REAL");
            builder.addTagAddress("sb_25_pv","%DB1.DBD224:REAL");
            builder.addTagAddress("sb_27_pv","%DB1.DBW244:INT");
            builder.addTagAddress("sb_28_pv","%DB1.DBW246:INT");
            builder.addTagAddress("sb_29_pv","%DB1.DBW248:INT");
            builder.addTagAddress("sb_30_pv","%DB1.DBW250:INT");
            final PlcReadRequest readRequest = builder.build();

            final PlcReadResponse readResponse = readRequest.execute().get();

            System.out.println(readResponse.getAsPlcValue());
        } catch (Exception e){
            System.out.println("error:"+e.getMessage());
        }

        }


        while(true) {
            try (PlcConnection connection = plcConnectionManager.getConnection("s7://10.80.41.47")) {
                final PlcReadRequest.Builder builder = connection.readRequestBuilder();
                builder.addTagAddress("string", "%DB4:340:STRING(256)"); // true

                final PlcReadRequest readRequest = builder.build();

                final PlcReadResponse readResponse = readRequest.execute().get();

                System.out.println(readResponse.getAsPlcValue());
            }
        }

         */
    }

}
