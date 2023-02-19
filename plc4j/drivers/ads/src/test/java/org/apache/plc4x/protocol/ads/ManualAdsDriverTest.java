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
package org.apache.plc4x.protocol.ads;

import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;
import org.apache.plc4x.java.api.value.PlcValue;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.function.Consumer;

public class ManualAdsDriverTest {
    private static final Logger LOGGER = LoggerFactory.getLogger(ManualAdsDriverTest.class);
    public static void main(String[] args) throws Exception {
        String ip = "10.80.41.238";

        String sourceAmsNetId = "10.211.55.4.1.1";
        int sourceAmsPort = 65534;
        String targetAmsNetId = "5.81.202.72.1.1";
        int targetAmsPort = 851;
        String connectionString = String.format("ads:tcp://%s?sourceAmsNetId=%s&sourceAmsPort=%d&targetAmsNetId=%s&targetAmsPort=%d", ip, sourceAmsNetId, sourceAmsPort, targetAmsNetId, targetAmsPort);
        try (PlcConnection plcConnection = new DefaultPlcDriverManager().getConnection(connectionString)) {

            final PlcReadRequest.Builder builder = plcConnection.readRequestBuilder();
            builder.addTagAddress("errorMsg","GVLMES.sErrorMessage:STRING(20)|GBK");
            builder.addTagAddress("sDrumID","GVLMES.sDrumID:STRING(20)");
            builder.addTagAddress("status","GVLMES.iConnectionStatus:INT");
            final PlcReadRequest readRequest = builder.build();
            final PlcReadResponse readResponse = readRequest.execute().get();
            System.out.println(readResponse.getAsPlcValue());


            final PlcWriteRequest.Builder rbuilder = plcConnection.writeRequestBuilder();
            rbuilder.addTagAddress("errorMsg", "GVLMES.sErrorMessage:STRING(20)|GBK", "啊啊啊啊");
            rbuilder.addTagAddress("sDrumID", "GVLMES.sDrumID:STRING(20)", "bbbbb");
            rbuilder.addTagAddress("status","GVLMES.iConnectionStatus:INT",2);
            final PlcWriteRequest writeRequest = rbuilder.build();

            final PlcWriteResponse writeResponse = writeRequest.execute().get();

            System.out.println(writeResponse);


            final PlcSubscriptionRequest.Builder sbuilder = plcConnection.subscriptionRequestBuilder();
            sbuilder.addChangeOfStateTagAddress("errorMsg", "GVLMES.sErrorMessage:STRING(20)");
            sbuilder.addChangeOfStateTagAddress("sDrumID", "GVLMES.sDrumID:STRING(20)");
            sbuilder.addChangeOfStateTagAddress("status", "GVLMES.iConnectionStatus:INT");
            PlcSubscriptionRequest subscriptionRequest = sbuilder.build();
            final PlcSubscriptionResponse subscriptionResponse = subscriptionRequest.execute().get();
            for (String subscriptionName : subscriptionResponse.getTagNames()) {
                final PlcSubscriptionHandle subscriptionHandle = subscriptionResponse.getSubscriptionHandle(subscriptionName);
                subscriptionHandle.register(new ValueChangeHandler());
            }

        }

    }
    private static class ValueChangeHandler implements Consumer<PlcSubscriptionEvent> {

        @Override
        public void accept(PlcSubscriptionEvent plcSubscriptionEvent) {
            LOGGER.info("Incoming event:");
            for (String fieldName : plcSubscriptionEvent.getTagNames()) {
                final PlcValue plcValue = plcSubscriptionEvent.getPlcValue(fieldName);
                if(plcValue.isList()) {
                    StringBuilder sb = new StringBuilder(String.format("Field '%s' value:", fieldName));
                    for (PlcValue value : plcValue.getList()) {
                        sb.append(" ").append(value.getString());
                    }
                    LOGGER.info(sb.toString());
                } else if (plcValue.isStruct()) {
                    StringBuilder sb = new StringBuilder(String.format("Field '%s' value:", fieldName));
                    plcValue.getStruct().forEach((name, value) ->
                        sb.append(" ").append(name).append("=").append(value.getString())
                    );
                    LOGGER.info(sb.toString());
                } else {
                    LOGGER.info(String.format("Field '%s' value: %s", fieldName, plcValue.getString()));
                }
            }
        }
    }

}
