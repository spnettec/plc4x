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
package org.apache.plc4x.java.modbus;

import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.PlcConnectionManager;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.utils.cache.CachedPlcConnectionManager;
import org.junit.jupiter.api.Disabled;

@Disabled("Manual Test")
public class ManualModbusTCPDriverTest  {


    public static void main(String[] args) throws Exception {
        CachedPlcConnectionManager plcConnectionManager = CachedPlcConnectionManager.getBuilder(new DefaultPlcDriverManager()).build();
        //PlcConnectionManager plcConnectionManager = new DefaultPlcDriverManager();
        while(true) {
            try (PlcConnection plcConnection = plcConnectionManager.getConnection(
                    "modbus-rtu:tcp://10.110.20.56:6000?request-timeout=1000")) {
                final PlcReadRequest readRequest = plcConnection.readRequestBuilder()
                        .addTagAddress("aa", "holding-register:1:UINT[11]").build();

                // Execute the read request
                try {
                    final PlcReadResponse readResponse = readRequest.execute().get();
                    System.out.println(readResponse.getAsPlcValue());
                }catch (Exception e) {
                    e.printStackTrace();
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }

}
