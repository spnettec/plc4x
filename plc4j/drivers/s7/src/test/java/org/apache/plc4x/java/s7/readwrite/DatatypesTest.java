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
import org.apache.plc4x.java.utils.cache.CachedPlcConnectionManager;

public class DatatypesTest {

    public static void main(String[] args) throws Exception {
        CachedPlcConnectionManager plcConnectionManager = CachedPlcConnectionManager.getBuilder().build();
        while (true) {
            try (PlcConnection connection = plcConnectionManager.getConnection("s7://192.168.2.22")) {
                final PlcReadRequest.Builder builder = connection.readRequestBuilder();
                builder.addTagAddress("value-1", "%DB53:2:3:BOOL");
                builder.addTagAddress("value-2", "%DB15:18:3:BOOL");
                builder.addTagAddress("value-3", "%DB34:114:REAL");
                builder.addTagAddress("value-4", "%DB34:118:REAL");
                builder.addTagAddress("value-5", "%DB53:1:4:BOOL");
                builder.addTagAddress("value-6", "%DB34:102:REAL");
                builder.addTagAddress("value-7", "%DB34:98:REAL");
                builder.addTagAddress("value-8", "%DB25:0:REAL");
                builder.addTagAddress("value-9", "%DB25:4:REAL");
                builder.addTagAddress("value-10", "%DB14:4:0:BOOL");
                builder.addTagAddress("value-11", "%DB71:0:REAL");
                builder.addTagAddress("value-12", "%DB71:4:REAL");
                builder.addTagAddress("value-13", "%DB17:14:REAL");
                builder.addTagAddress("value-14", "%DB17:38:REAL");
                builder.addTagAddress("value-15", "%DB17:0:1:BOOL");
                builder.addTagAddress("value-16", "%DB17:24:0:BOOL");
                builder.addTagAddress("value-17", "%DB17:24:1:BOOL");
                builder.addTagAddress("value-18", "%DB17:0:0:BOOL");
                builder.addTagAddress("value-19", "%DB17:10:REAL");
                builder.addTagAddress("value-20", "%DB17:34:REAL");
                builder.addTagAddress("value-21", "%DB71:62:1:BOOL");
                builder.addTagAddress("value-22", "%DB71:10:REAL");
                builder.addTagAddress("value-23", "%DB1:8:0:BOOL");
                builder.addTagAddress("value-24", "%DB1:116:0:BOOL");
                builder.addTagAddress("value-25", "%DB34:42:REAL");
                builder.addTagAddress("value-26", "%DB34:46:REAL");
                builder.addTagAddress("value-27", "%DB34:50:REAL");
                builder.addTagAddress("value-28", "%DB34:54:REAL");
                builder.addTagAddress("value-29", "%DB1:0:REAL");
                builder.addTagAddress("value-30", "%DB1:4:REAL");
                builder.addTagAddress("value-31", "%DB1:92:REAL");
                builder.addTagAddress("value-32", "%DB1:96:REAL");
                builder.addTagAddress("value-33", "%DB34:66:REAL");
                builder.addTagAddress("value-34", "%DB34:70:REAL");
                builder.addTagAddress("value-35", "%DB34:74:REAL");
                builder.addTagAddress("value-36", "%DB34:78:REAL");
                builder.addTagAddress("value-37", "%DB42:30:5:BOOL");
                builder.addTagAddress("value-38", "%DB53:1:6:BOOL");
                builder.addTagAddress("value-39", "%DB34:82:REAL");
                builder.addTagAddress("value-40", "%DB34:86:REAL");
                builder.addTagAddress("value-41", "%DB18:0:1:BOOL");
                builder.addTagAddress("value-42", "%DB18:2:REAL");
                builder.addTagAddress("value-43", "%DB18:10:REAL");
                builder.addTagAddress("value-44", "%DB53:2:2:BOOL");
                builder.addTagAddress("value-45", "%DB34:106:REAL");
                builder.addTagAddress("value-46", "%DB34:110:REAL");
                builder.addTagAddress("value-47", "%DB15:14:REAL");
                builder.addTagAddress("value-48", "%DB15:18:2:BOOL");
                builder.addTagAddress("value-49", "%DB53:0:0:BOOL");
                builder.addTagAddress("value-50", "%DB34:58:REAL");
                builder.addTagAddress("value-51", "%DB34:62:REAL");
                builder.addTagAddress("value-52", "%DB34:108:INT");
                builder.addTagAddress("value-53", "%DB23:44:REAL");
                builder.addTagAddress("value-54", "%DB34:0:0:BOOL");
                builder.addTagAddress("value-55", "%DB34:0:1:BOOL");
                builder.addTagAddress("value-56", "%DB53:2:1:BOOL");
                builder.addTagAddress("value-57", "%DB23:8:REAL");
                builder.addTagAddress("value-58", "%DB32:0:1:BOOL");
                builder.addTagAddress("value-59", "%DB34:154:REAL");
                builder.addTagAddress("value-60", "%DB34:158:REAL");
                builder.addTagAddress("value-61", "%DB34:162:REAL");
                builder.addTagAddress("value-62", "%DB34:166:REAL");
                builder.addTagAddress("value-63", "%DB34:170:REAL");
                builder.addTagAddress("value-64", "%DB34:174:REAL");
                builder.addTagAddress("value-65", "%DB34:178:REAL");
                builder.addTagAddress("value-66", "%DB34:182:REAL");
                builder.addTagAddress("value-67", "%DB53:0:2:BOOL");
                builder.addTagAddress("value-68", "%DB17:62:REAL");
                builder.addTagAddress("value-69", "%DB17:48:1:BOOL");
                builder.addTagAddress("value-70", "%DB17:48:0:BOOL");
                builder.addTagAddress("value-71", "%DB17:58:REAL");
                builder.addTagAddress("value-72", "%DB32:14:REAL");
                builder.addTagAddress("value-73", "%DB32:6:REAL");
                builder.addTagAddress("value-74", "%DB53:0:3:BOOL");
                builder.addTagAddress("value-75", "%DB34:122:REAL");
                builder.addTagAddress("value-76", "%DB34:126:REAL");
                builder.addTagAddress("value-77", "%DB34:130:REAL");
                builder.addTagAddress("value-78", "%DB34:134:REAL");
                builder.addTagAddress("value-79", "%DB34:138:REAL");
                builder.addTagAddress("value-80", "%DB34:142:REAL");
                builder.addTagAddress("value-81", "%DB34:146:REAL");
                builder.addTagAddress("value-82", "%DB34:150:REAL");
                builder.addTagAddress("value-83", "%DB23:26:REAL");
                builder.addTagAddress("value-84", "%DB53:1:7:BOOL");
                builder.addTagAddress("value-85", "%DB34:90:REAL");
                builder.addTagAddress("value-86", "%DB34:94:REAL");
                builder.addTagAddress("value-87", "%DB53:0:1:BOOL");
                builder.addTagAddress("value-88", "%DB34:2:REAL");
                builder.addTagAddress("value-89", "%DB34:6:REAL");
                builder.addTagAddress("value-90", "%DB34:10:REAL");
                builder.addTagAddress("value-91", "%DB34:14:REAL");
                builder.addTagAddress("value-92", "%DB34:18:REAL");
                builder.addTagAddress("value-93", "%DB34:22:REAL");
                builder.addTagAddress("value-94", "%DB34:26:REAL");
                builder.addTagAddress("value-95", "%DB34:30:REAL");
                builder.addTagAddress("value-96", "%DB34:34:REAL");
                builder.addTagAddress("value-97", "%DB34:38:REAL");
                builder.addTagAddress("value-98", "%DB16:20:0:BOOL");

                //builder.addTagAddress("huge-array", "%DB2:0:UINT[4000]");
                final PlcReadRequest readRequest = builder.build();

                final PlcReadResponse readResponse = readRequest.execute().get();

                System.out.println(readResponse.getAsPlcValue());
            }
        }
    }

}
