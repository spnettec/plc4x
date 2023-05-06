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
package org.apache.plc4x.java.profinet;

import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcBrowseRequest;
import org.apache.plc4x.java.api.messages.PlcBrowseResponse;
import org.apache.plc4x.java.api.messages.PlcSubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcSubscriptionResponse;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.profinet.device.ProfinetSubscriptionHandle;
import org.apache.plc4x.java.profinet.tag.ProfinetTag;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class ManualProfinetIoTest {

    private static final Logger LOGGER = LoggerFactory.getLogger(ManualProfinetIoTest.class);

    public static void main(String[] args) throws Exception {
        final PlcConnection connection = new DefaultPlcDriverManager().getConnection("profinet://192.168.54.2?gsddirectory=/home/missy/Documents/Profinet/gsd&devices=[[simocodexbpn156e,DAP%201,(1,),192.168.54.23]]&reductionratio=16&sendclockfactor=32&dataholdfactor=3&watchdogfactor=3");
        PlcBrowseRequest browseRequest = connection.browseRequestBuilder().addQuery("Browse", "").build();
        final PlcBrowseResponse browseResponse = browseRequest.execute().get();
        PlcSubscriptionRequest.Builder builder = connection.subscriptionRequestBuilder();
        builder.addChangeOfStateTag("Input 4", ProfinetTag.of("simocodexbpn156e.1.1.Inputs.2:BOOL"));
        PlcSubscriptionRequest request = builder.build();

        final PlcSubscriptionResponse response = request.execute().get();

        // Get result of creating subscription
        final ProfinetSubscriptionHandle subscriptionHandle = (ProfinetSubscriptionHandle) response.getSubscriptionHandle("Input 4");

        // Create handler for returned value
        subscriptionHandle.register(plcSubscriptionEvent -> {
            assert plcSubscriptionEvent.getResponseCode("Input 4").equals(PlcResponseCode.OK);
            LOGGER.info("Received a response from {} test {}", "Input 4", plcSubscriptionEvent.getPlcValue("Input 4").toString());
        });
    }

}
