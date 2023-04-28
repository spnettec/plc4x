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
package org.apache.plc4x.java.opcua;

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.utils.cache.CachedPlcConnectionManager;

/**
 * This class serves only as a manual entry point for ad-hoc tests of the OPC UA PLC4J driver.
 * <p>
 * <p>
 * The current version is tested against a public server, which is to be replaced later by a separate instance of the Milo framework.
 * Afterwards the code represented here will be used as an example for the introduction page.
 * <p>
 *
 */
public class ManualPLC4XOpcua {

    public static void main(String args[]) {
        CachedPlcConnectionManager plcConnectionManager = CachedPlcConnectionManager.getBuilder().build();

        try (PlcConnection connection = plcConnectionManager.getConnection("opcua:tcp://127.0.0.1:12673?discovery=true")) {
            final PlcReadRequest.Builder builder = connection.readRequestBuilder();
            builder.addTagAddress("Simulated_BOOL", "ns=3;s=Simulated_BOOL");
            builder.addTagAddress("Simulated_BYTE", "ns=3;s=Simulated_BYTE");
            builder.addTagAddress("Simulated_WORD", "ns=3;s=Simulated_WORD");
            builder.addTagAddress("Simulated_DWORD", "ns=3;s=Simulated_DWORD");
            builder.addTagAddress("Simulated_SINT", "ns=3;s=Simulated_SINT");
            builder.addTagAddress("Simulated_USINT", "ns=3;s=Simulated_USINT");
            builder.addTagAddress("Simulated_INT", "ns=3;s=Simulated_INT");
            builder.addTagAddress("Simulated_UINT", "ns=3;s=Simulated_UINT");
            builder.addTagAddress("Simulated_DINT", "ns=3;s=Simulated_DINT");
            builder.addTagAddress("Simulated_UDINT", "ns=3;s=Simulated_UDINT");
            builder.addTagAddress("Simulated_LINT", "ns=3;s=Simulated_LINT");
            builder.addTagAddress("Simulated_ULINT", "ns=3;s=Simulated_ULINT");
            builder.addTagAddress("Simulated_REAL", "ns=3;s=Simulated_REAL");
            builder.addTagAddress("Simulated_LREAL", "ns=3;s=Simulated_LREAL");
            builder.addTagAddress("Simulated_WCHAR", "ns=3;s=Simulated_WCHAR");
            builder.addTagAddress("Simulated_STRING", "ns=3;s=Simulated_STRING");
            builder.addTagAddress("Simulated_WSTRING", "ns=3;s=Simulated_WSTRING");
            builder.addTagAddress("Simulated_BYTEARRAY", "ns=3;s=Simulated_BYTEARRAY");
            builder.addTagAddress("Simulated_INTARRAY", "ns=3;s=Simulated_INTARRAY");
            final PlcReadRequest readRequest = builder.build();

            final PlcReadResponse readResponse = readRequest.execute().get();
            System.out.println(readResponse.getAsPlcValue());

        } catch (Exception e) {
            System.out.println("error:" + e.getMessage());
        }
    }
}
