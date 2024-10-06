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
        //try (PlcConnection connection = new DefaultPlcDriverManager().getConnection("s7://192.168.24.83")) {
        try (PlcConnection connection = new DefaultPlcDriverManager().getConnection("s7://192.168.2.22")) {
                final PlcReadRequest.Builder builder = connection.readRequestBuilder();
            builder.addTagAddress("bool-value-1", "%DB14:4:0:BOOL"); // true
            builder.addTagAddress("bool-value-2", "%DB34:0:0:1BOOL");

            //builder.addTagAddress("huge-array", "%DB2:0:UINT[4000]");
            final PlcReadRequest readRequest = builder.build();

            final PlcReadResponse readResponse = readRequest.execute().get();

            System.out.println(readResponse);

        }
    }

}
