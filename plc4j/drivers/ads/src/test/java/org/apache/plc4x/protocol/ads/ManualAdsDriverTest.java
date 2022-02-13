/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.protocol.ads;

import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.test.manual.ManualTest;

import java.util.Arrays;

public class ManualAdsDriverTest extends ManualTest {

    /*
     * Test program code on the PLC with the test-data.
     *
     * Located in "main"
     *

    hurz_BOOL  := TRUE;
	hurz_BYTE  := 42;
	hurz_WORD  := 42424;
	hurz_DWORD := 4242442424;
	hurz_LWORD := 4242442424242424242;
	hurz_SINT  := -42;
	hurz_USINT := 42;
	hurz_INT   := -2424;
	hurz_UINT  := 42424;
	hurz_DINT  := -242442424;
	hurz_UDINT := 4242442424;
	hurz_LINT  := -4242442424242424242;
	hurz_ULINT := 4242442424242424242;
	hurz_REAL  := 3.14159265359;
	hurz_LREAL := 2.71828182846;
	hurz_TIME  := T#1S234MS;
	hurz_LTIME := LTIME#1000D15H23M12S34MS2US44NS;
	hurz_DATE  := D#1978-03-28;
	//hurz_LDATE:LDATE;
	hurz_TIME_OF_DAY 	:= TIME_OF_DAY#15:36:30.123;
	hurz_TOD         	:= TOD#16:17:18.123;
	//hurz_LTIME_OF_DAY:LTIME_OF_DAY;
	//hurz_LTOD:LTOD;
	hurz_DATE_AND_TIME 	:= DATE_AND_TIME#1996-05-06-15:36:30;
	hurz_DT				:= DT#1972-03-29-00:00:00;
	//hurz_LDATE_AND_TIME:LDATE_AND_TIME;
	//hurz_LDT:LDT;
	hurz_STRING			:= 'hurz';
	hurz_WSTRING		:= "wolf";

     *
     */

    public ManualAdsDriverTest(String connectionString) {
        super(connectionString);
    }

    public static void main(String[] args) throws Exception {
        String ip = "127.0.0.1";

        String sourceAmsNetId = "10.80.41.10.1.1";
        int sourceAmsPort = 65534;
        String targetAmsNetId = "5.81.202.72.1.1";
        int targetAmsPort = 851;
        String connectionString = String.format("ads:tcp://%s?sourceAmsNetId=%s&sourceAmsPort=%d&targetAmsNetId=%s&targetAmsPort=%d", ip, sourceAmsNetId, sourceAmsPort, targetAmsNetId, targetAmsPort);
        try (PlcConnection plcConnection = new PlcDriverManager().getConnection(connectionString)) {

            final PlcReadRequest.Builder builder = plcConnection.readRequestBuilder();
            builder.addItem("errorMsg","GVLMES.sErrorMessage:STRING(20)");
           // builder.addItem("sDrumID","GVLMES.sDrumID:STRING(20)");
            final PlcReadRequest readRequest = builder.build();
            final PlcReadResponse readResponse = readRequest.execute().get();
            System.out.println(readResponse.getAsPlcValue());


            final PlcWriteRequest.Builder rbuilder = plcConnection.writeRequestBuilder();
            rbuilder.addItem("errorMsg", "GVLMES.sErrorMessage:STRING(20)", "啊啊啊啊");

            final PlcWriteRequest writeRequest = rbuilder.build();

            final PlcWriteResponse writeResponse = writeRequest.execute().get();

            System.out.println(writeResponse);
        }

    }

}
