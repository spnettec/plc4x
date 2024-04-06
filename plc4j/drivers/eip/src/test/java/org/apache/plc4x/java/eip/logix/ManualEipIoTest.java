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
package org.apache.plc4x.java.eip.logix;

import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class ManualEipIoTest {

    public static void main(String[] args) throws Exception {
        /*
        final PlcConnection connection = new DefaultPlcDriverManager().getConnection("logix://192.168.50.36");
        final PlcReadRequest readRequest = connection.readRequestBuilder()
            .addTagAddress("B1", "B1")
            .addTagAddress("B2", "B2")
            .addTagAddress("DI1", "DI1")
            .addTagAddress("DI2", "DI2")
            .addTagAddress("I1", "I1")
            .addTagAddress("I2", "I2")
            .addTagAddress("R1", "R1")
            .addTagAddress("R2", "R2")
            .build();
        final PlcReadResponse plcReadResponse = readRequest.execute().get();
        connection.close();
        System.out.println(plcReadResponse.getAsPlcValue());
        */

        final PlcConnection connection = new DefaultPlcDriverManager().getConnection("logix://192.168.50.36");
        final PlcWriteRequest readRequest = connection.writeRequestBuilder()
            .addTagAddress("B2", "B2:BOOL",true)
            .addTagAddress("I2", "I2:INT",5)
            .build();
        final PlcWriteResponse plcReadResponse = readRequest.execute().get();
        connection.close();
        System.out.println(plcReadResponse);

    }
}
