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
package org.apache.plc4x.java.s7.readwrite;

import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.java.utils.connectionpool.PooledPlcDriverManager;

public class DatatypesTest {

    private PlcReadResponse readResponse;

    public static void main(String[] args) throws Exception {
        PooledPlcDriverManager pooledPlcDriverManager = new PooledPlcDriverManager();

        try (PlcConnection connection = pooledPlcDriverManager.getConnection("s7://10.166.11.20?remote-rack=0&remote-slot=1")) {
            final PlcReadRequest.Builder builder = connection.readRequestBuilder();
            builder.addItem("CTray01_BarDiameter","%DB4:242:REAL");
            builder.addItem("CTray01_Length","%DB4:246:REAL");
            builder.addItem("CTray01_Status","%DB4:250:REAL");
            builder.addItem("CTray01_StickName","%DB4:200:STRING(40)");
            builder.addItem("CTray02_BarDiameter","%DB4:342:REAL");
            builder.addItem("CTray02_Length","%DB4:346:REAL");
            builder.addItem("CTray02_Status","%DB4:350:REAL");
            builder.addItem("CTray02_StickName","%DB4:300:STRING(40)");
            builder.addItem("MTray01_BarDiameter","%DB4:42:REAL");
            builder.addItem("MTray01_Length","%DB4:46:REAL");
            builder.addItem("MTray01_Status","%DB4:50:REAL");
            builder.addItem("MTray01_StickName","%DB4:0:STRING(40)");
            builder.addItem("MTray02_BarDiameter","%DB4:142:REAL");
            builder.addItem("MTray02_Length","%DB4:146:REAL");
            builder.addItem("MTray02_Status","%DB4:150:REAL");
            builder.addItem("MTray02_StickName","%DB4:100:STRING(40)");
            builder.addItem("PMS1301_(P|G)EndPoint","%DB131:54:INT");
            builder.addItem("PMS1301_BarDiameter","%DB131:56:REAL");
            builder.addItem("PMS1301_JobNum","%DB131:60:STRING(20)");
            builder.addItem("PMS1301_MeasurePoint","%DB131:166:STRING(200)");
            builder.addItem("PMS1301_StickName","%DB131:82:STRING(40)");
            builder.addItem("PMS1301_Template","%DB131:124:STRING(40)");
            builder.addItem("PMS1302_(P|G)EndPoint","%DB132:54:INT");
            builder.addItem("PMS1302_BarDiameter","%DB132:56:REAL");
            builder.addItem("PMS1302_JobNum","%DB132:60:STRING(20)");
            builder.addItem("PMS1302_StickName","%DB132:82:STRING(40)");
            builder.addItem("PMS1302_Template","%DB132:124:STRING(40)");
            builder.addItem("PMS1302_MeasurePoint","%DB132:166:STRING(200)");





            final PlcReadRequest readRequest = builder.build();
            final PlcReadResponse readResponse = readRequest.execute().get();
            // Object o = readResponse.getObject("bool-value-1");
            System.out.println(readResponse.getAsPlcValue());





            final PlcWriteRequest.Builder rbuilder = connection.writeRequestBuilder();
            rbuilder.addItem("PMS1302_Template", "%DB132:124:STRING(40)|GBK", "啊啊啊啊11"); // true

            final PlcWriteRequest writeRequest = rbuilder.build();

            final PlcWriteResponse writeResponse = writeRequest.execute().get();

            System.out.println(writeResponse);


        }
    }

}
