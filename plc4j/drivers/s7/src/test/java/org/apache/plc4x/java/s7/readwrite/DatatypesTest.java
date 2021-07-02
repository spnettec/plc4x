/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.s7.readwrite;

import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;

public class DatatypesTest {

    private PlcReadResponse readResponse;

    public static void main(String[] args) throws Exception {
        try (PlcConnection connection = new PlcDriverManager().getConnection("s7://10.166.11.20?local-rack=0")) {
            final PlcReadRequest.Builder builder = connection.readRequestBuilder();
            builder.addItem("bool-value-1", "%DB1:16:DATE"); // true

            final PlcReadRequest readRequest = builder.build();

            final PlcReadResponse readResponse = readRequest.execute().get();
            Object o = readResponse.getObject("bool-value-1");
            System.out.println(o);


            final PlcWriteRequest.Builder rbuilder = connection.writeRequestBuilder();
            rbuilder.addItem("bool-value-1", "%DB1:16:DATE", "2083-03-23"); // true

            final PlcWriteRequest writeRequest = rbuilder.build();

            final PlcWriteResponse writeResponse = writeRequest.execute().get();

            System.out.println(writeResponse);


        }
    }

}
