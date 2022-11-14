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
import org.apache.plc4x.java.utils.connectionpool.PooledPlcDriverManager;

public class DatatypesTest {

    public static void main(String[] args) throws Exception {
        PooledPlcDriverManager pooledPlcDriverManager = new PooledPlcDriverManager();

        try (PlcConnection connection = pooledPlcDriverManager.getConnection("s7://10.166.11.20?remote-rack=0&remote-slot=1")) {
            final PlcReadRequest.Builder builder = connection.readRequestBuilder();
            builder.addTagAddress("CTray01_BarDiameter","%DB3:192.0:BOOL");
            builder.addTagAddress("CTray01_Length","%DB3:0:STRING(22)");
            builder.addTagAddress("CTray01_Status","%DB3:48:STRING(22)");
            builder.addTagAddress("CTray01_StickName","%DB3:72:STRING(22)");
            builder.addTagAddress("CTray02_BarDiameter","%DB3:96:STRING(22)");
            builder.addTagAddress("CTray02_Length","%DB3:120:STRING(22)");
            builder.addTagAddress("CTray02_Status","%DB3:144:STRING(22)");
            builder.addTagAddress("CTray02_StickName","%DB3:168:STRING(22)");
            builder.addTagAddress("MTray01_BarDiameter","%DB3:24:STRING(22)");





            final PlcReadRequest readRequest = builder.build();
            final PlcReadResponse readResponse = readRequest.execute().get();
            // Object o = readResponse.getObject("bool-value-1");
            System.out.println(readResponse.getAsPlcValue());


        }
    }

}
