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
package org.apache.plc4x.java.plc4x;

import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;

public class Plc4xReadClient {

    public static void main(String[] args) throws Exception {
        try (final PlcConnection connection = new PlcDriverManager().getConnection("plc4x://localhost?remote-connection-string=s7://10.166.11.18?remote-rack=0&remote-slot=1")) {
            final PlcReadRequest.Builder builder = connection.readRequestBuilder();
            builder.addItem("TestBool1","%DB1:4.1:BOOL");
            builder.addItem("TestBool2","%DB1:4.0:BOOL");
            builder.addItem("TestInt","%DB1:18:INT");
            builder.addItem("TestLReal","%DB1:20:LREAL");
            builder.addItem("TestReal1","%DB1:0:REAL");
            builder.addItem("TestReal2","%DB1:28:REAL");
            builder.addItem("TestSInt","%DB1:32:SINT");
            builder.addItem("TestString","%DB1:34:STRING(40)|GBK");
            builder.addItem("TestTypeArray1","%DB1:818:BYTE[10]");
            builder.addItem("TestUDInt","%DB1:294:UDINT");
            builder.addItem("TestWString","%DB1:304:STRING(100)");

            final PlcReadRequest readRequest = builder.build();
            final PlcReadResponse readResponse = readRequest.execute().get();
            System.out.println(readResponse);
        }
    }

}
